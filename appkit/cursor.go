// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var CursorClass = _CursorClass{objc.GetClass("NSCursor")}

type _CursorClass struct {
	objc.Class
}

type ICursor interface {
	objc.IObject
	Push()
	Set()
	// deprecated
	MouseEntered(event IEvent)
	// deprecated
	SetOnMouseEntered(flag bool)
	// deprecated
	MouseExited(event IEvent)
	// deprecated
	SetOnMouseExited(flag bool)
	Image() Image
	HotSpot() foundation.Point
	// deprecated
	IsSetOnMouseEntered() bool
	// deprecated
	IsSetOnMouseExited() bool
}

type Cursor struct {
	objc.Object
}

func MakeCursor(ptr unsafe.Pointer) Cursor {
	return Cursor{
		Object: objc.MakeObject(ptr),
	}
}

func (c_ Cursor) InitWithImage_HotSpot(newImage IImage, point foundation.Point) Cursor {
	rv := objc.CallMethod[Cursor](c_, "initWithImage:hotSpot:", newImage, point)
	return rv
}

func (c_ Cursor) InitWithImage_ForegroundColorHint_BackgroundColorHint_HotSpot(newImage IImage, fg IColor, bg IColor, hotSpot foundation.Point) Cursor {
	rv := objc.CallMethod[Cursor](c_, "initWithImage:foregroundColorHint:backgroundColorHint:hotSpot:", newImage, fg, bg, hotSpot)
	return rv
}

func (cc _CursorClass) Alloc() Cursor {
	rv := objc.CallMethod[Cursor](cc, "alloc")
	return rv
}

func (cc _CursorClass) New() Cursor {
	rv := objc.CallMethod[Cursor](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCursor() Cursor {
	return CursorClass.New()
}

func (c_ Cursor) Init() Cursor {
	rv := objc.CallMethod[Cursor](c_, "init")
	return rv
}

func (cc _CursorClass) Hide() {
	objc.CallMethod[objc.Void](cc, "hide")
}

func (cc _CursorClass) Unhide() {
	objc.CallMethod[objc.Void](cc, "unhide")
}

func (cc _CursorClass) SetHiddenUntilMouseMoves(flag bool) {
	objc.CallMethod[objc.Void](cc, "setHiddenUntilMouseMoves:", flag)
}

func (cc _CursorClass) Pop() {
	objc.CallMethod[objc.Void](cc, "pop")
}

func (c_ Cursor) Push() {
	objc.CallMethod[objc.Void](c_, "push")
}

func (c_ Cursor) Set() {
	objc.CallMethod[objc.Void](c_, "set")
}

// deprecated
func (c_ Cursor) MouseEntered(event IEvent) {
	objc.CallMethod[objc.Void](c_, "mouseEntered:", event)
}

// deprecated
func (c_ Cursor) SetOnMouseEntered(flag bool) {
	objc.CallMethod[objc.Void](c_, "setOnMouseEntered:", flag)
}

// deprecated
func (c_ Cursor) MouseExited(event IEvent) {
	objc.CallMethod[objc.Void](c_, "mouseExited:", event)
}

// deprecated
func (c_ Cursor) SetOnMouseExited(flag bool) {
	objc.CallMethod[objc.Void](c_, "setOnMouseExited:", flag)
}

func (c_ Cursor) Image() Image {
	rv := objc.CallMethod[Image](c_, "image")
	return rv
}

func (c_ Cursor) HotSpot() foundation.Point {
	rv := objc.CallMethod[foundation.Point](c_, "hotSpot")
	return rv
}

// deprecated
func (c_ Cursor) IsSetOnMouseEntered() bool {
	rv := objc.CallMethod[bool](c_, "isSetOnMouseEntered")
	return rv
}

// deprecated
func (c_ Cursor) IsSetOnMouseExited() bool {
	rv := objc.CallMethod[bool](c_, "isSetOnMouseExited")
	return rv
}

func (cc _CursorClass) CurrentCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, "currentCursor")
	return rv
}

func (cc _CursorClass) CurrentSystemCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, "currentSystemCursor")
	return rv
}

func (cc _CursorClass) ArrowCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, "arrowCursor")
	return rv
}

func (cc _CursorClass) ContextualMenuCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, "contextualMenuCursor")
	return rv
}

func (cc _CursorClass) ClosedHandCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, "closedHandCursor")
	return rv
}

func (cc _CursorClass) CrosshairCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, "crosshairCursor")
	return rv
}

func (cc _CursorClass) DisappearingItemCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, "disappearingItemCursor")
	return rv
}

func (cc _CursorClass) DragCopyCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, "dragCopyCursor")
	return rv
}

func (cc _CursorClass) DragLinkCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, "dragLinkCursor")
	return rv
}

func (cc _CursorClass) IBeamCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, "IBeamCursor")
	return rv
}

func (cc _CursorClass) OpenHandCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, "openHandCursor")
	return rv
}

func (cc _CursorClass) OperationNotAllowedCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, "operationNotAllowedCursor")
	return rv
}

func (cc _CursorClass) PointingHandCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, "pointingHandCursor")
	return rv
}

func (cc _CursorClass) ResizeDownCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, "resizeDownCursor")
	return rv
}

func (cc _CursorClass) ResizeLeftCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, "resizeLeftCursor")
	return rv
}

func (cc _CursorClass) ResizeLeftRightCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, "resizeLeftRightCursor")
	return rv
}

func (cc _CursorClass) ResizeRightCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, "resizeRightCursor")
	return rv
}

func (cc _CursorClass) ResizeUpCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, "resizeUpCursor")
	return rv
}

func (cc _CursorClass) ResizeUpDownCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, "resizeUpDownCursor")
	return rv
}

func (cc _CursorClass) IBeamCursorForVerticalLayout() Cursor {
	rv := objc.CallMethod[Cursor](cc, "IBeamCursorForVerticalLayout")
	return rv
}
