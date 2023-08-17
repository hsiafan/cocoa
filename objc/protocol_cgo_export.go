package objc

import "C"
import (
	"reflect"
	"runtime/cgo"
	"unsafe"
)

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
