// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var CoderClass = _CoderClass{objc.GetClass("NSCoder")}

type _CoderClass struct {
	objc.Class
}

type ICoder interface {
	IObject
}

type Coder struct {
	Object
}

func MakeCoder(ptr unsafe.Pointer) Coder {
	return Coder{
		Object: MakeObject(ptr),
	}
}
