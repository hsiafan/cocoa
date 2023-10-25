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
	objc.IObject
}

type Coder struct {
	objc.Object
}

func MakeCoder(ptr unsafe.Pointer) Coder {
	return Coder{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CoderClass) Alloc() Coder {
	rv := objc.CallMethod[Coder](cc, objc.SEL("alloc"))
	return rv
}

func (cc _CoderClass) New() Coder {
	rv := objc.CallMethod[Coder](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewCoder() Coder {
	return CoderClass.New()
}

func (c_ Coder) Init() Coder {
	rv := objc.CallMethod[Coder](c_, objc.SEL("init"))
	return rv
}
