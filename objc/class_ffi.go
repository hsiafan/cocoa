package objc

import (
	"reflect"
	"runtime/cgo"
	"strings"
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
)

// for void type
type Void struct{}

var voidType = reflect.TypeOf((*Void)(nil)).Elem()

// type T: the ret value type
func CallMethod[T any](o Holder, selector Selector, params ...any) T {
	ptr := o.Ptr()
	if ptr == nil {
		panic("object is nil")
	}
	argc := len(params)

	var args = make([]unsafe.Pointer, argc+2)
	var argTypes = make([]*ffi.Type, argc+2)
	args[0] = unsafe.Pointer(&ptr)
	argTypes[0] = ffi.TypePointer
	args[1] = unsafe.Pointer(&selector.ptr)
	argTypes[1] = ffi.TypePointer
	for i := 0; i < argc; i++ {
		args[i+2] = convertToObjcValue(reflect.ValueOf(params[i]))
		argTypes[i+2] = getFFIType(params[i])
	}

	var ret T
	var retPtr unsafe.Pointer
	rt := reflect.TypeOf(ret)
	if directPointer(rt) {
		retPtr = unsafe.Pointer(&ret)
	} else {
		var i uintptr
		retPtr = unsafe.Pointer(&i)
	}
	retType := toFFIType(rt)

	class := MakeObject(ptr).GetClass()
	var imp IMP
	if ffi.IsStruct(retType) {
		imp = class.GetMethodImplementationStret(selector)
	} else {
		imp = class.GetMethodImplementation(selector)
	}
	if imp.ptr == nil {
		return ret
	}
	cif, status := ffi.PrepCIF(retType, argTypes)
	if status != ffi.OK {
		panic("ffi prep cif status not ok")
	}
	ffi.Call(cif, imp.ptr, retPtr, args)
	if directPointer(rt) {
		return *(*T)(retPtr)
	}

	return convertToGoValue(retPtr, reflect.TypeOf(ret)).Interface().(T)
}

func directPointer(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Bool:
		return true
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Int, reflect.Uint, reflect.Int64, reflect.Uint64:
		return true
	case reflect.Float32, reflect.Float64:
		return true
	case reflect.UnsafePointer, reflect.Pointer:
		return true
	case reflect.Struct:
		return true
	case reflect.String:
		return false
	case reflect.Slice, reflect.Map:
		return false
	default:
		return false
	}
}

// AddMethod add a new objc instance method with a go function.
// The first param of go function should be the object instance, the second param should be the method selector.
func AddMethod(class Class, sel Selector, f any) bool {
	rf := reflect.ValueOf(f)

	typeEncoding := _getMethodTypeEncoding(rf.Type(), false)

	imp, handle := wrapGoFuncAsMethodIMP(rf)
	flag := class.AddMethod(sel, imp, typeEncoding)
	if !flag {
		handle.Delete()
	}
	return flag
}

// ReplaceMethod replace objc instance method with a go function.
// The first param of go function should be the object instance, the second param should be the method selector.
func ReplaceMethod(class Class, sel Selector, f any) {
	rf := reflect.ValueOf(f)
	typeEncoding := _getMethodTypeEncoding(rf.Type(), false)

	imp, handle := wrapGoFuncAsMethodIMP(rf)
	_ = handle
	oldIMP := class.ReplaceMethod(sel, imp, typeEncoding)
	if oldIMP.ptr != nil {
		//
	}
}

// AddClassMethod add a new objc class method with a go function.
// The first param of go function should be the class, the second param should be the method selector.
func AddClassMethod(class Class, sel Selector, f any) bool {
	rf := reflect.ValueOf(f)
	typeEncoding := _getMethodTypeEncoding(rf.Type(), true)

	imp, handle := wrapGoFuncAsMethodIMP(rf)
	metaClass := class.GetClass()
	if metaClass.ptr == nil {
		panic("no meta class")
	}
	flag := metaClass.AddMethod(sel, imp, typeEncoding)
	if !flag {
		handle.Delete()
	}
	return flag
}

// ReplaceClassMethod replace objc class method with a go function.
func ReplaceClassMethod(class Class, sel Selector, f any) {
	rf := reflect.ValueOf(f)
	typeEncoding := _getMethodTypeEncoding(rf.Type(), true)

	imp, handle := wrapGoFuncAsMethodIMP(rf)
	_ = handle

	oldIMP := MakeObject(class.ptr).GetClass().ReplaceMethod(sel, imp, typeEncoding)
	if oldIMP.ptr != nil {

	}
}

func _getMethodTypeEncoding(ft reflect.Type, classMethod bool) string {
	if ft.Kind() != reflect.Func {
		panic("not func type")
	}
	if ft.NumOut() > 1 {
		panic("to many return values")
	}
	if ft.NumIn() < 1 || !ft.In(0).AssignableTo(pointerHolderType) {
		panic("first param must be the objc object or class")
	}

	var sb strings.Builder
	if ft.NumOut() == 0 {
		sb.WriteByte('v')
	} else {
		sb.WriteString(getTypeEncoding(ft.Out(0)))
	}
	if classMethod {
		sb.WriteString("#") // class self as first parameter
	} else {
		sb.WriteString("@") // instance self as first parameter
	}
	sb.WriteString(":") // selector
	// skip first go param
	for i := 1; i < ft.NumIn(); i++ {
		sb.WriteString(getTypeEncoding(ft.In(i)))
	}
	return sb.String()
}

// the first param of go func must be o object instance or a class instance;
// other params are mapped from objc method params.
func wrapGoFuncAsMethodIMP(rf reflect.Value) (imp IMP, handle cgo.Handle) {
	if rf.Kind() != reflect.Func {
		panic("f should be a func")
	}
	rt := rf.Type()
	if rt.NumOut() > 1 {
		panic("too many return value")
	}

	goArgc := rt.NumIn()
	var objcArgTypes = make([]*ffi.Type, goArgc+1)
	objcArgTypes[0] = ffi.TypePointer // objc instance or class pointer
	objcArgTypes[1] = ffi.TypePointer // selector pointer
	// skip selector param(the second param of objc method)
	for i := 1; i < goArgc; i++ {
		objcArgTypes[i+1] = toFFIType(rt.In(i))
	}

	var retType *ffi.Type
	if rt.NumOut() == 0 {
		retType = ffi.TypeVoid
	} else {
		retType = getFFIType(rt.Out(0))
	}

	cif, status := ffi.PrepCIF(retType, objcArgTypes)
	if status != ffi.OK {
		panic("ffi prep cif status not ok")
	}

	fn, handle, status := ffi.CreateClosure(cif, func(cif *ffi.CIF, ret unsafe.Pointer, objcArgs []unsafe.Pointer) {
		var goArgs = make([]reflect.Value, len(objcArgs)-1)
		goArgs[0] = convertToGoValue(objcArgs[0], rt.In(0))
		// skip selector param(the second param of objc method)
		for i := 1; i < len(goArgs); i++ {
			goArgs[i] = convertToGoValue(objcArgs[i+1], rt.In(i))
		}
		results := rf.Call(goArgs)
		if len(results) == 1 {
			setGoValueToObjcPointer(results[0], ret)
		}
	})
	if status != ffi.OK {
		panic("ffi prep closure status not ok")
	}

	return MakeIMP(fn), handle
}
