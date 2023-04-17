package ffi

//#import <stdlib.h>
//#import <stdint.h>
//void* wrap_block(uintptr_t goID, const char *typeEncoding);
// void invocationGetArgs(void* ptr, int begin, int argc, uintptr_t argsPtr);
// void invocationSetRet(void* ptr, void* loc);
// void callBlock(void *bp, int argc, uintptr_t argsPtr, void* ret);
import "C"
import (
	"reflect"
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

func wrapBlockInGoFunc(bp unsafe.Pointer, funcType reflect.Type) reflect.Value {
	if bp == nil {
		panic("block is nil")
	}
	if funcType.NumOut() > 1 {
		panic("too many return values")
	}

	b := objc.MakeBlock(bp)
	b = b.Copy()
	var sentinel = new(int)
	fv := reflect.MakeFunc(funcType, func(args []reflect.Value) (results []reflect.Value) {
		*sentinel = 1
		if funcType.NumOut() == 0 {
			callBlock(b, args, nil)
			return nil
		} else {
			rvs := callBlock(b, args, funcType.Out(0))
			return rvs
		}
	})
	runtime.SetFinalizer(sentinel, func(v *int) {
		b.Release()
	})
	return fv
}

func callBlock(b objc.Block, params []reflect.Value, retType reflect.Type) []reflect.Value {
	argc := len(params)

	var argsPtr C.uintptr_t
	var args []unsafe.Pointer
	if argc > 0 {
		args = make([]unsafe.Pointer, argc)
		for i := 0; i < argc; i++ {
			args[i] = convertToObjcValue(params[i].Interface())
		}
		argsPtr = toUIntptr(&args[0])
	}

	var retPtr unsafe.Pointer
	if retType != nil {
		rv := reflect.New(retType).Elem()
		retPtr = convertToObjcValue(rv.Interface())
	}

	C.callBlock(b.Ptr(), C.int(argc), argsPtr, retPtr)

	runtime.KeepAlive(args)
	if retType != nil {
		return []reflect.Value{convertToGoValue(retPtr, retType)}
	} else {
		return nil
	}
}

// f is the go function to wrap as a objc block
func WrapBlock(f any) objc.Block {
	ft := reflect.TypeOf(f)
	if ft.Kind() != reflect.Func {
		panic("not func type")
	}

	typeEncoding := getBlockTypeEncoding(ft)
	cte := C.CString(typeEncoding)
	defer C.free(unsafe.Pointer(cte))
	goId := cgo.NewHandle(f)
	r := C.wrap_block(C.uintptr_t(goId), cte)
	return objc.MakeBlock(r)
}

//export handleBlockInvocation
func handleBlockInvocation(goID uintptr, invocationPtr unsafe.Pointer) {
	f := reflect.ValueOf(cgo.Handle(goID).Value())
	ft := f.Type()
	args := make([]reflect.Value, ft.NumIn())

	argsPtr := make([]unsafe.Pointer, len(args))
	for i := 0; i < ft.NumIn(); i++ {
		argsPtr[i] = parseGoType(ft.In(i))
	}

	invocation := makeInvocation(invocationPtr)
	if len(argsPtr) > 0 {
		invocation.getArguments(1, argsPtr)
	}
	for i := 0; i < ft.NumIn(); i++ {
		args[i] = convertToGoValue(argsPtr[i], ft.In(i))
	}

	rvs := f.Call(args)
	if len(rvs) > 1 {
		panic("too many return values")
	}
	if len(rvs) == 1 {
		rp := convertToObjcValue(rvs[0].Interface())
		invocation.setReturnValue(rp)
	}
}
