// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var ContentWorldClass = _ContentWorldClass{objc.GetClass("WKContentWorld")}

type _ContentWorldClass struct {
	objc.Class
}

type IContentWorld interface {
	objc.IObject
	Name() string
}

type ContentWorld struct {
	objc.Object
}

func MakeContentWorld(ptr unsafe.Pointer) ContentWorld {
	return ContentWorld{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _ContentWorldClass) Alloc() ContentWorld {
	rv := objc.CallMethod[ContentWorld](cc, objc.SEL("alloc"))
	return rv
}

func (cc _ContentWorldClass) New() ContentWorld {
	rv := objc.CallMethod[ContentWorld](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewContentWorld() ContentWorld {
	return ContentWorldClass.New()
}

func (c_ ContentWorld) Init() ContentWorld {
	rv := objc.CallMethod[ContentWorld](c_, objc.SEL("init"))
	return rv
}

func (cc _ContentWorldClass) WorldWithName(name string) ContentWorld {
	rv := objc.CallMethod[ContentWorld](cc, objc.SEL("worldWithName:"), name)
	return rv
}

func (cc _ContentWorldClass) DefaultClientWorld() ContentWorld {
	rv := objc.CallMethod[ContentWorld](cc, objc.SEL("defaultClientWorld"))
	return rv
}

func (cc _ContentWorldClass) PageWorld() ContentWorld {
	rv := objc.CallMethod[ContentWorld](cc, objc.SEL("pageWorld"))
	return rv
}

func (c_ ContentWorld) Name() string {
	rv := objc.CallMethod[string](c_, objc.SEL("name"))
	return rv
}
