package objc

// #import <stdint.h>
// void callMethod(void *op, void* sp, int argc, uintptr_t argsPtr, void* ret);
import "C"
import (
	"reflect"
	"runtime"
	"unsafe"
)

// for void type
type Void struct{}

var voidType = reflect.TypeOf((*Void)(nil)).Elem()

// type T: the ret value type
func CallMethod[T any](o Holder, selName string, params ...any) T {
	if o.Ptr() == nil {
		panic("object is nil")
	}
	argc := len(params)
	selector := GetSelector(selName)

	var argsPtr C.uintptr_t
	var args []unsafe.Pointer
	if argc > 0 {
		args = make([]unsafe.Pointer, argc)
		for i := 0; i < argc; i++ {
			args[i] = convertToObjcValue(params[i])
		}
		argsPtr = toUIntptr(&args[0])
	}

	var retPtr unsafe.Pointer
	var ret T
	retType := reflect.TypeOf(ret)
	isDirect := directPointer(retType)
	if unsafe.Sizeof(ret) != 0 {
		if isDirect {
			retPtr = unsafe.Pointer(&ret)
		} else {
			retPtr = convertToObjcValue(ret)
		}
	}

	C.callMethod(o.Ptr(), selector.Ptr(), C.int(argc), argsPtr, retPtr)

	runtime.KeepAlive(args)
	if isDirect {
		return ret
	} else {
		return convertToGoValue(retPtr, retType).Interface().(T)
	}
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

func toUIntptr[T any](ptr *T) C.uintptr_t {
	return C.uintptr_t(uintptr(unsafe.Pointer(ptr)))
}
