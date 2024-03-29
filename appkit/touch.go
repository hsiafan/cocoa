// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TouchClass = _TouchClass{objc.GetClass("NSTouch")}

type _TouchClass struct {
	objc.Class
}

type ITouch interface {
	objc.IObject
	LocationInView(view IView) foundation.Point
	PreviousLocationInView(view IView) foundation.Point
	Type() TouchType
	Phase() TouchPhase
	NormalizedPosition() foundation.Point
	IsResting() bool
	Device() objc.Object
	DeviceSize() foundation.Size
}

type Touch struct {
	objc.Object
}

func MakeTouch(ptr unsafe.Pointer) Touch {
	return Touch{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TouchClass) Alloc() Touch {
	rv := objc.CallMethod[Touch](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TouchClass) New() Touch {
	rv := objc.CallMethod[Touch](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTouch() Touch {
	return TouchClass.New()
}

func (t_ Touch) Init() Touch {
	rv := objc.CallMethod[Touch](t_, objc.SEL("init"))
	return rv
}

func (t_ Touch) LocationInView(view IView) foundation.Point {
	rv := objc.CallMethod[foundation.Point](t_, objc.SEL("locationInView:"), objc.ExtractPtr(view))
	return rv
}

func (t_ Touch) PreviousLocationInView(view IView) foundation.Point {
	rv := objc.CallMethod[foundation.Point](t_, objc.SEL("previousLocationInView:"), objc.ExtractPtr(view))
	return rv
}

func (t_ Touch) Type() TouchType {
	rv := objc.CallMethod[TouchType](t_, objc.SEL("type"))
	return rv
}

func (t_ Touch) Phase() TouchPhase {
	rv := objc.CallMethod[TouchPhase](t_, objc.SEL("phase"))
	return rv
}

func (t_ Touch) NormalizedPosition() foundation.Point {
	rv := objc.CallMethod[foundation.Point](t_, objc.SEL("normalizedPosition"))
	return rv
}

func (t_ Touch) IsResting() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isResting"))
	return rv
}

func (t_ Touch) Device() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.SEL("device"))
	return rv
}

func (t_ Touch) DeviceSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.SEL("deviceSize"))
	return rv
}
