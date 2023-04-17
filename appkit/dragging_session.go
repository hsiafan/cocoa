// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var DraggingSessionClass = _DraggingSessionClass{objc.GetClass("NSDraggingSession")}

type _DraggingSessionClass struct {
	objc.Class
}

type IDraggingSession interface {
	objc.IObject
	EnumerateDraggingItemsWithOptions_ForView_Classes_SearchOptions_UsingBlock(enumOpts DraggingItemEnumerationOptions, view IView, classArray []objc.IClass, searchOptions map[PasteboardReadingOptionKey]objc.IObject, block func(draggingItem DraggingItem, idx int, stop *bool))
	DraggingPasteboard() Pasteboard
	AnimatesToStartingPositionsOnCancelOrFail() bool
	SetAnimatesToStartingPositionsOnCancelOrFail(value bool)
	DraggingFormation() DraggingFormation
	SetDraggingFormation(value DraggingFormation)
	DraggingSequenceNumber() int
	DraggingLocation() foundation.Point
	DraggingLeaderIndex() int
	SetDraggingLeaderIndex(value int)
}

type DraggingSession struct {
	objc.Object
}

func MakeDraggingSession(ptr unsafe.Pointer) DraggingSession {
	return DraggingSession{
		Object: objc.MakeObject(ptr),
	}
}

func (dc _DraggingSessionClass) Alloc() DraggingSession {
	rv := ffi.CallMethod[DraggingSession](dc, "alloc")
	return rv
}

func (dc _DraggingSessionClass) New() DraggingSession {
	rv := ffi.CallMethod[DraggingSession](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDraggingSession() DraggingSession {
	return DraggingSessionClass.New()
}

func (d_ DraggingSession) Init() DraggingSession {
	rv := ffi.CallMethod[DraggingSession](d_, "init")
	return rv
}

func (d_ DraggingSession) EnumerateDraggingItemsWithOptions_ForView_Classes_SearchOptions_UsingBlock(enumOpts DraggingItemEnumerationOptions, view IView, classArray []objc.IClass, searchOptions map[PasteboardReadingOptionKey]objc.IObject, block func(draggingItem DraggingItem, idx int, stop *bool)) {
	ffi.CallMethod[ffi.Void](d_, "enumerateDraggingItemsWithOptions:forView:classes:searchOptions:usingBlock:", enumOpts, view, classArray, searchOptions, block)
}

func (d_ DraggingSession) DraggingPasteboard() Pasteboard {
	rv := ffi.CallMethod[Pasteboard](d_, "draggingPasteboard")
	return rv
}

func (d_ DraggingSession) AnimatesToStartingPositionsOnCancelOrFail() bool {
	rv := ffi.CallMethod[bool](d_, "animatesToStartingPositionsOnCancelOrFail")
	return rv
}

func (d_ DraggingSession) SetAnimatesToStartingPositionsOnCancelOrFail(value bool) {
	ffi.CallMethod[ffi.Void](d_, "setAnimatesToStartingPositionsOnCancelOrFail:", value)
}

func (d_ DraggingSession) DraggingFormation() DraggingFormation {
	rv := ffi.CallMethod[DraggingFormation](d_, "draggingFormation")
	return rv
}

func (d_ DraggingSession) SetDraggingFormation(value DraggingFormation) {
	ffi.CallMethod[ffi.Void](d_, "setDraggingFormation:", value)
}

func (d_ DraggingSession) DraggingSequenceNumber() int {
	rv := ffi.CallMethod[int](d_, "draggingSequenceNumber")
	return rv
}

func (d_ DraggingSession) DraggingLocation() foundation.Point {
	rv := ffi.CallMethod[foundation.Point](d_, "draggingLocation")
	return rv
}

func (d_ DraggingSession) DraggingLeaderIndex() int {
	rv := ffi.CallMethod[int](d_, "draggingLeaderIndex")
	return rv
}

func (d_ DraggingSession) SetDraggingLeaderIndex(value int) {
	ffi.CallMethod[ffi.Void](d_, "setDraggingLeaderIndex:", value)
}
