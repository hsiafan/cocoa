// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[TrackingArea](t_, "initWithRect:options:owner:userInfo:", rect, options, owner, userInfo)
	rv.Autorelease()
	return rv
}

func (tc _TrackingAreaClass) Alloc() TrackingArea {
	rv := ffi.CallMethod[TrackingArea](tc, "alloc")
	return rv
}

func (t_ TrackingArea) Init() TrackingArea {
	rv := ffi.CallMethod[TrackingArea](t_, "init")
	rv.Autorelease()
	return rv
}

func (tc _TrackingAreaClass) New() TrackingArea {
	rv := ffi.CallMethod[TrackingArea](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTrackingArea() TrackingArea {
	return TrackingAreaClass.New()
}

func (t_ TrackingArea) Options() TrackingAreaOptions {
	rv := ffi.CallMethod[TrackingAreaOptions](t_, "options")
	return rv
}

func (t_ TrackingArea) Owner() objc.Object {
	rv := ffi.CallMethod[objc.Object](t_, "owner")
	return rv
}

func (t_ TrackingArea) Rect() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](t_, "rect")
	return rv
}

func (t_ TrackingArea) UserInfo() foundation.Dictionary {
	rv := ffi.CallMethod[foundation.Dictionary](t_, "userInfo")
	return rv
}
