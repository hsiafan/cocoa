// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var SecurityOriginClass = _SecurityOriginClass{objc.GetClass("WKSecurityOrigin")}

type _SecurityOriginClass struct {
	objc.Class
}

type ISecurityOrigin interface {
	objc.IObject
	Host() string
	Port() int
	Protocol() string
}

type SecurityOrigin struct {
	objc.Object
}

func MakeSecurityOrigin(ptr unsafe.Pointer) SecurityOrigin {
	return SecurityOrigin{
		Object: objc.MakeObject(ptr),
	}
}

func (sc _SecurityOriginClass) Alloc() SecurityOrigin {
	rv := objc.CallMethod[SecurityOrigin](sc, "alloc")
	return rv
}

func (sc _SecurityOriginClass) New() SecurityOrigin {
	rv := objc.CallMethod[SecurityOrigin](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSecurityOrigin() SecurityOrigin {
	return SecurityOriginClass.New()
}

func (s_ SecurityOrigin) Init() SecurityOrigin {
	rv := objc.CallMethod[SecurityOrigin](s_, "init")
	return rv
}

func (s_ SecurityOrigin) Host() string {
	rv := objc.CallMethod[string](s_, "host")
	return rv
}

func (s_ SecurityOrigin) Port() int {
	rv := objc.CallMethod[int](s_, "port")
	return rv
}

func (s_ SecurityOrigin) Protocol() string {
	rv := objc.CallMethod[string](s_, "protocol")
	return rv
}
