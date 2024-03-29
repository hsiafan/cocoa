// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[DraggingSession](dc, objc.SEL("alloc"))
	return rv
}

func (dc _DraggingSessionClass) New() DraggingSession {
	rv := objc.CallMethod[DraggingSession](dc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewDraggingSession() DraggingSession {
	return DraggingSessionClass.New()
}

func (d_ DraggingSession) Init() DraggingSession {
	rv := objc.CallMethod[DraggingSession](d_, objc.SEL("init"))
	return rv
}

func (d_ DraggingSession) EnumerateDraggingItemsWithOptions_ForView_Classes_SearchOptions_UsingBlock(enumOpts DraggingItemEnumerationOptions, view IView, classArray []objc.IClass, searchOptions map[PasteboardReadingOptionKey]objc.IObject, block func(draggingItem DraggingItem, idx int, stop *bool)) {
	objc.CallMethod[objc.Void](d_, objc.SEL("enumerateDraggingItemsWithOptions:forView:classes:searchOptions:usingBlock:"), enumOpts, objc.ExtractPtr(view), classArray, searchOptions, block)
}

func (d_ DraggingSession) DraggingPasteboard() Pasteboard {
	rv := objc.CallMethod[Pasteboard](d_, objc.SEL("draggingPasteboard"))
	return rv
}

func (d_ DraggingSession) AnimatesToStartingPositionsOnCancelOrFail() bool {
	rv := objc.CallMethod[bool](d_, objc.SEL("animatesToStartingPositionsOnCancelOrFail"))
	return rv
}

func (d_ DraggingSession) SetAnimatesToStartingPositionsOnCancelOrFail(value bool) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setAnimatesToStartingPositionsOnCancelOrFail:"), value)
}

func (d_ DraggingSession) DraggingFormation() DraggingFormation {
	rv := objc.CallMethod[DraggingFormation](d_, objc.SEL("draggingFormation"))
	return rv
}

func (d_ DraggingSession) SetDraggingFormation(value DraggingFormation) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setDraggingFormation:"), value)
}

func (d_ DraggingSession) DraggingSequenceNumber() int {
	rv := objc.CallMethod[int](d_, objc.SEL("draggingSequenceNumber"))
	return rv
}

func (d_ DraggingSession) DraggingLocation() foundation.Point {
	rv := objc.CallMethod[foundation.Point](d_, objc.SEL("draggingLocation"))
	return rv
}

func (d_ DraggingSession) DraggingLeaderIndex() int {
	rv := objc.CallMethod[int](d_, objc.SEL("draggingLeaderIndex"))
	return rv
}

func (d_ DraggingSession) SetDraggingLeaderIndex(value int) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setDraggingLeaderIndex:"), value)
}
