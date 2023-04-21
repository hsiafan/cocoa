// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var NullClass = _NullClass{objc.GetClass("NSNull")}

type _NullClass struct {
	objc.Class
}

type INull interface {
	objc.IObject
}

type Null struct {
	objc.Object
}

func MakeNull(ptr unsafe.Pointer) Null {
	return Null{
		Object: objc.MakeObject(ptr),
	}
}

func (nc _NullClass) Alloc() Null {
	rv := objc.CallMethod[Null](nc, "alloc")
	return rv
}

func (nc _NullClass) New() Null {
	rv := objc.CallMethod[Null](nc, "new")
	rv.Autorelease()
	return rv
}

func NewNull() Null {
	return NullClass.New()
}

func (n_ Null) Init() Null {
	rv := objc.CallMethod[Null](n_, "init")
	return rv
}

func (nc _NullClass) Null() Null {
	rv := objc.CallMethod[Null](nc, "null")
	return rv
}
