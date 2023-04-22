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
	rv := objc.CallMethod[Cursor](c_, objc.GetSelector("initWithImage:hotSpot:"), newImage, point)
	return rv
}

func (c_ Cursor) InitWithImage_ForegroundColorHint_BackgroundColorHint_HotSpot(newImage IImage, fg IColor, bg IColor, hotSpot foundation.Point) Cursor {
	rv := objc.CallMethod[Cursor](c_, objc.GetSelector("initWithImage:foregroundColorHint:backgroundColorHint:hotSpot:"), newImage, fg, bg, hotSpot)
	return rv
}

func (cc _CursorClass) Alloc() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("alloc"))
	return rv
}

func (cc _CursorClass) New() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCursor() Cursor {
	return CursorClass.New()
}

func (c_ Cursor) Init() Cursor {
	rv := objc.CallMethod[Cursor](c_, objc.GetSelector("init"))
	return rv
}

func (cc _CursorClass) Hide() {
	objc.CallMethod[objc.Void](cc, objc.GetSelector("hide"))
}

func (cc _CursorClass) Unhide() {
	objc.CallMethod[objc.Void](cc, objc.GetSelector("unhide"))
}

func (cc _CursorClass) SetHiddenUntilMouseMoves(flag bool) {
	objc.CallMethod[objc.Void](cc, objc.GetSelector("setHiddenUntilMouseMoves:"), flag)
}

func (cc _CursorClass) Pop() {
	objc.CallMethod[objc.Void](cc, objc.GetSelector("pop"))
}

func (c_ Cursor) Push() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("push"))
}

func (c_ Cursor) Set() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("set"))
}

// deprecated
func (c_ Cursor) MouseEntered(event IEvent) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("mouseEntered:"), event)
}

// deprecated
func (c_ Cursor) SetOnMouseEntered(flag bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setOnMouseEntered:"), flag)
}

// deprecated
func (c_ Cursor) MouseExited(event IEvent) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("mouseExited:"), event)
}

// deprecated
func (c_ Cursor) SetOnMouseExited(flag bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setOnMouseExited:"), flag)
}

func (c_ Cursor) Image() Image {
	rv := objc.CallMethod[Image](c_, objc.GetSelector("image"))
	return rv
}

func (c_ Cursor) HotSpot() foundation.Point {
	rv := objc.CallMethod[foundation.Point](c_, objc.GetSelector("hotSpot"))
	return rv
}

// deprecated
func (c_ Cursor) IsSetOnMouseEntered() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isSetOnMouseEntered"))
	return rv
}

// deprecated
func (c_ Cursor) IsSetOnMouseExited() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isSetOnMouseExited"))
	return rv
}

func (cc _CursorClass) CurrentCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("currentCursor"))
	return rv
}

func (cc _CursorClass) CurrentSystemCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("currentSystemCursor"))
	return rv
}

func (cc _CursorClass) ArrowCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("arrowCursor"))
	return rv
}

func (cc _CursorClass) ContextualMenuCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("contextualMenuCursor"))
	return rv
}

func (cc _CursorClass) ClosedHandCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("closedHandCursor"))
	return rv
}

func (cc _CursorClass) CrosshairCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("crosshairCursor"))
	return rv
}

func (cc _CursorClass) DisappearingItemCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("disappearingItemCursor"))
	return rv
}

func (cc _CursorClass) DragCopyCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("dragCopyCursor"))
	return rv
}

func (cc _CursorClass) DragLinkCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("dragLinkCursor"))
	return rv
}

func (cc _CursorClass) IBeamCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("IBeamCursor"))
	return rv
}

func (cc _CursorClass) OpenHandCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("openHandCursor"))
	return rv
}

func (cc _CursorClass) OperationNotAllowedCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("operationNotAllowedCursor"))
	return rv
}

func (cc _CursorClass) PointingHandCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("pointingHandCursor"))
	return rv
}

func (cc _CursorClass) ResizeDownCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("resizeDownCursor"))
	return rv
}

func (cc _CursorClass) ResizeLeftCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("resizeLeftCursor"))
	return rv
}

func (cc _CursorClass) ResizeLeftRightCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("resizeLeftRightCursor"))
	return rv
}

func (cc _CursorClass) ResizeRightCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("resizeRightCursor"))
	return rv
}

func (cc _CursorClass) ResizeUpCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("resizeUpCursor"))
	return rv
}

func (cc _CursorClass) ResizeUpDownCursor() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("resizeUpDownCursor"))
	return rv
}

func (cc _CursorClass) IBeamCursorForVerticalLayout() Cursor {
	rv := objc.CallMethod[Cursor](cc, objc.GetSelector("IBeamCursorForVerticalLayout"))
	return rv
}
