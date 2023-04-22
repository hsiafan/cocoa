package objc

// #import <stdint.h>
//
// void* New_ProtocolImpl(void* class, uintptr_t goID);
import "C"
import (
	"reflect"
	"runtime/cgo"
	"sync"
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/internal"
)

var classCache = map[string]*classInfo{} //go protocol interface name to ClassInfo
var classLock sync.Mutex
var baseClass = GetClass("ProtocolImplBase")

type instanceInfo struct {
	instance     any        // the go instance
	protocolName string     // go protocol interface name
	classInfo    *classInfo // the class info for this instance
}

type classInfo struct {
	class       Class
	methodInfos map[string]*methodInfo // selector name to MethodInfo
}

type methodInfo struct {
	required bool          // if is required protocol method
	hasFunc  reflect.Value // the hasXXX func for this method
}

// param protocolName: the delegate go interface name
// param d: the delegate go implementaion
func CreateProtocol[T any](protocolName string, d T) Object {
	dv := reflect.ValueOf(d)
	ci := createClass(dv.Type(), protocolName)
	ii := &instanceInfo{
		classInfo:    ci,
		instance:     d,
		protocolName: protocolName,
	}
	h := cgo.NewHandle(ii)
	return MakeObject(C.New_ProtocolImpl(ci.class.Ptr(), C.uintptr_t(h)))
}

func createClass(t reflect.Type, protocolName string) *classInfo {
	if t.Kind() == reflect.Interface {
		panic("should not be interface type")
	}

	classLock.Lock()
	defer classLock.Unlock()
	if ci, ok := classCache[protocolName]; ok {
		return ci
	}
	class := AllocateClassPair(baseClass, protocolName+"Adaptor", 0)
	protocol := GetProtocol(protocolName)
	class.AddProtocol(protocol)

	var methodInfos = map[string]*methodInfo{} // selector name to method signature
	for _, md := range getProtocolMethods(protocol) {
		selector := md.Name
		selName := selector.GetName()
		goFuncName := internal.SelectorToGoName(selName)
		goMethod, ok := t.MethodByName(goFuncName)
		if !ok {
			if md.required {
				panic("required method not implemented:" + selName)
			} else {
				continue
			}
		}
		addProtocolMethod(class, md, goMethod)
		hasFunc, _ := t.MethodByName("Implements" + goFuncName)

		methodInfos[selName] = &methodInfo{
			required: md.required,
			hasFunc:  hasFunc.Func,
		}
	}

	RegisterClassPair(class)
	ci := &classInfo{
		class:       class,
		methodInfos: methodInfos,
	}

	classCache[protocolName] = ci
	return ci
}

type methodDescription struct {
	MethodDescription
	required       bool
	instanceMethod bool
}

func getProtocolMethods(protocol Protocol) []methodDescription {
	if protocol.GetName() == "NSObject" {
		return nil
	}
	var mds []methodDescription
	for _, md := range protocol.CopyMethodDescriptionList(true, true) {
		mds = append(mds, methodDescription{
			MethodDescription: md,
			required:          true,
			instanceMethod:    true,
		})
	}

	for _, md := range protocol.CopyMethodDescriptionList(false, true) {
		mds = append(mds, methodDescription{
			MethodDescription: md,
			required:          false,
			instanceMethod:    true,
		})
	}

	for _, parent := range protocol.CopyProtocolList() {
		mds = append(mds, getProtocolMethods(parent)...)
	}
	return mds
}

func addProtocolMethod(class Class, md methodDescription, method reflect.Method) {
	//TODO: class method
	if !md.instanceMethod {
		return
	}

	rt := method.Type
	if rt.NumOut() > 1 {
		panic("too many return value")
	}
	goArgc := rt.NumIn()
	var objcArgTypes = make([]*ffi.Type, goArgc+1)
	objcArgTypes[0] = ffi.TypePointer // objc instance or class pointer
	objcArgTypes[1] = ffi.TypePointer // selector pointer
	// first go fun param is receiver
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
		o := MakeObject(*(*unsafe.Pointer)(objcArgs[0]))
		handle := CallMethod[uintptr](o, GetSelector("goID"))
		instance := cgo.Handle(handle).Value().(*instanceInfo)

		var goArgs = make([]reflect.Value, len(objcArgs)-1)
		goArgs[0] = reflect.ValueOf(instance.instance)
		// skip selector param(the second param of objc method)
		for i := 1; i < len(goArgs); i++ {
			goArgs[i] = convertToGoValue(objcArgs[i+1], rt.In(i))
		}
		results := method.Func.Call(goArgs)
		if len(results) == 1 {
			setGoValueToObjcPointer(results[0], ret)
		}
	})
	_ = handle // never free
	if status != ffi.OK {
		panic("ffi prep closure status not ok")
	}
	flag := class.AddMethod(md.Name, MakeIMP(fn), md.Types)
	if !flag {
		panic("add method to protocol failed")
	}
}

//export respondsTo
func respondsTo(goID uintptr, sel unsafe.Pointer) bool {
	selName := MakeSelector(sel).GetName()
	ii := cgo.Handle(goID).Value().(*instanceInfo)
	mi := ii.classInfo.methodInfos[selName]
	if mi == nil {
		return false
	}

	v := reflect.ValueOf(ii.instance)
	if mi.required {
		return true
	}

	return mi.hasFunc.Call([]reflect.Value{v})[0].Bool()
}
