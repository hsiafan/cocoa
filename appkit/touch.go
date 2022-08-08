// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	Identity() objc.Object
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
	rv := ffi.CallMethod[Touch](tc, "alloc")
	return rv
}

func (t_ Touch) Init() Touch {
	rv := ffi.CallMethod[Touch](t_, "init")
	return rv
}

func (tc _TouchClass) New() Touch {
	rv := ffi.CallMethod[Touch](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTouch() Touch {
	return TouchClass.New()
}

func (t_ Touch) LocationInView(view IView) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](t_, "locationInView:", view)
	return rv
}

func (t_ Touch) PreviousLocationInView(view IView) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](t_, "previousLocationInView:", view)
	return rv
}

func (t_ Touch) Type() TouchType {
	rv := ffi.CallMethod[TouchType](t_, "type")
	return rv
}

func (t_ Touch) Identity() objc.Object {
	rv := ffi.CallMethod[objc.Object](t_, "identity")
	return rv
}

func (t_ Touch) Phase() TouchPhase {
	rv := ffi.CallMethod[TouchPhase](t_, "phase")
	return rv
}

func (t_ Touch) NormalizedPosition() foundation.Point {
	rv := ffi.CallMethod[foundation.Point](t_, "normalizedPosition")
	return rv
}

func (t_ Touch) IsResting() bool {
	rv := ffi.CallMethod[bool](t_, "isResting")
	return rv
}

func (t_ Touch) Device() objc.Object {
	rv := ffi.CallMethod[objc.Object](t_, "device")
	return rv
}

func (t_ Touch) DeviceSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](t_, "deviceSize")
	return rv
}
