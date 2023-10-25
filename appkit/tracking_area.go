// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TrackingAreaClass = _TrackingAreaClass{objc.GetClass("NSTrackingArea")}

type _TrackingAreaClass struct {
	objc.Class
}

type ITrackingArea interface {
	objc.IObject
	Options() TrackingAreaOptions
	Owner() objc.Object
	Rect() foundation.Rect
	UserInfo() foundation.Dictionary
}

type TrackingArea struct {
	objc.Object
}

func MakeTrackingArea(ptr unsafe.Pointer) TrackingArea {
	return TrackingArea{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TrackingArea) InitWithRect_Options_Owner_UserInfo(rect foundation.Rect, options TrackingAreaOptions, owner objc.IObject, userInfo foundation.IDictionary) TrackingArea {
	rv := objc.CallMethod[TrackingArea](t_, objc.SEL("initWithRect:options:owner:userInfo:"), rect, options, objc.ExtractPtr(owner), userInfo)
	return rv
}

func (tc _TrackingAreaClass) Alloc() TrackingArea {
	rv := objc.CallMethod[TrackingArea](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TrackingAreaClass) New() TrackingArea {
	rv := objc.CallMethod[TrackingArea](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTrackingArea() TrackingArea {
	return TrackingAreaClass.New()
}

func (t_ TrackingArea) Init() TrackingArea {
	rv := objc.CallMethod[TrackingArea](t_, objc.SEL("init"))
	return rv
}

func (t_ TrackingArea) Options() TrackingAreaOptions {
	rv := objc.CallMethod[TrackingAreaOptions](t_, objc.SEL("options"))
	return rv
}

// weak property
func (t_ TrackingArea) Owner() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.SEL("owner"))
	return rv
}

func (t_ TrackingArea) Rect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.SEL("rect"))
	return rv
}

func (t_ TrackingArea) UserInfo() foundation.Dictionary {
	rv := objc.CallMethod[foundation.Dictionary](t_, objc.SEL("userInfo"))
	return rv
}
