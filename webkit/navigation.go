// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var NavigationClass = _NavigationClass{objc.GetClass("WKNavigation")}

type _NavigationClass struct {
	objc.Class
}

type INavigation interface {
	objc.IObject
	EffectiveContentMode() ContentMode
}

type Navigation struct {
	objc.Object
}

func MakeNavigation(ptr unsafe.Pointer) Navigation {
	return Navigation{
		Object: objc.MakeObject(ptr),
	}
}

func (nc _NavigationClass) Alloc() Navigation {
	rv := objc.CallMethod[Navigation](nc, "alloc")
	return rv
}

func (nc _NavigationClass) New() Navigation {
	rv := objc.CallMethod[Navigation](nc, "new")
	rv.Autorelease()
	return rv
}

func NewNavigation() Navigation {
	return NavigationClass.New()
}

func (n_ Navigation) Init() Navigation {
	rv := objc.CallMethod[Navigation](n_, "init")
	return rv
}

func (n_ Navigation) EffectiveContentMode() ContentMode {
	rv := objc.CallMethod[ContentMode](n_, "effectiveContentMode")
	return rv
}
