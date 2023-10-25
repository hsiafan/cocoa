// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ClipViewClass = _ClipViewClass{objc.GetClass("NSClipView")}

type _ClipViewClass struct {
	objc.Class
}

type IClipView interface {
	IView
	ScrollToPoint(newOrigin foundation.Point)
	// deprecated
	ConstrainScrollPoint(newOrigin foundation.Point) foundation.Point
	ConstrainBoundsRect(proposedBounds foundation.Rect) foundation.Rect
	ViewBoundsChanged(notification foundation.INotification)
	ViewFrameChanged(notification foundation.INotification)
	DocumentView() View
	SetDocumentView(value IView)
	// deprecated
	CopiesOnScroll() bool
	// deprecated
	SetCopiesOnScroll(value bool)
	ContentInsets() foundation.EdgeInsets
	SetContentInsets(value foundation.EdgeInsets)
	AutomaticallyAdjustsContentInsets() bool
	SetAutomaticallyAdjustsContentInsets(value bool)
	DocumentRect() foundation.Rect
	DocumentVisibleRect() foundation.Rect
	DocumentCursor() Cursor
	SetDocumentCursor(value ICursor)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
}

type ClipView struct {
	View
}

func MakeClipView(ptr unsafe.Pointer) ClipView {
	return ClipView{
		View: MakeView(ptr),
	}
}

func (c_ ClipView) InitWithFrame(frameRect foundation.Rect) ClipView {
	rv := objc.CallMethod[ClipView](c_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (c_ ClipView) Init() ClipView {
	rv := objc.CallMethod[ClipView](c_, objc.SEL("init"))
	return rv
}

func (cc _ClipViewClass) Alloc() ClipView {
	rv := objc.CallMethod[ClipView](cc, objc.SEL("alloc"))
	return rv
}

func (cc _ClipViewClass) New() ClipView {
	rv := objc.CallMethod[ClipView](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewClipView() ClipView {
	return ClipViewClass.New()
}

func (c_ ClipView) ScrollToPoint(newOrigin foundation.Point) {
	objc.CallMethod[objc.Void](c_, objc.SEL("scrollToPoint:"), newOrigin)
}

// deprecated
func (c_ ClipView) ConstrainScrollPoint(newOrigin foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](c_, objc.SEL("constrainScrollPoint:"), newOrigin)
	return rv
}

func (c_ ClipView) ConstrainBoundsRect(proposedBounds foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.SEL("constrainBoundsRect:"), proposedBounds)
	return rv
}

func (c_ ClipView) ViewBoundsChanged(notification foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.SEL("viewBoundsChanged:"), objc.ExtractPtr(notification))
}

func (c_ ClipView) ViewFrameChanged(notification foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.SEL("viewFrameChanged:"), objc.ExtractPtr(notification))
}

func (c_ ClipView) DocumentView() View {
	rv := objc.CallMethod[View](c_, objc.SEL("documentView"))
	return rv
}

func (c_ ClipView) SetDocumentView(value IView) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setDocumentView:"), objc.ExtractPtr(value))
}

// deprecated
func (c_ ClipView) CopiesOnScroll() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("copiesOnScroll"))
	return rv
}

// deprecated
func (c_ ClipView) SetCopiesOnScroll(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setCopiesOnScroll:"), value)
}

func (c_ ClipView) ContentInsets() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](c_, objc.SEL("contentInsets"))
	return rv
}

func (c_ ClipView) SetContentInsets(value foundation.EdgeInsets) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setContentInsets:"), value)
}

func (c_ ClipView) AutomaticallyAdjustsContentInsets() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("automaticallyAdjustsContentInsets"))
	return rv
}

func (c_ ClipView) SetAutomaticallyAdjustsContentInsets(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setAutomaticallyAdjustsContentInsets:"), value)
}

func (c_ ClipView) DocumentRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.SEL("documentRect"))
	return rv
}

func (c_ ClipView) DocumentVisibleRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.SEL("documentVisibleRect"))
	return rv
}

func (c_ ClipView) DocumentCursor() Cursor {
	rv := objc.CallMethod[Cursor](c_, objc.SEL("documentCursor"))
	return rv
}

func (c_ ClipView) SetDocumentCursor(value ICursor) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setDocumentCursor:"), objc.ExtractPtr(value))
}

func (c_ ClipView) DrawsBackground() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("drawsBackground"))
	return rv
}

func (c_ ClipView) SetDrawsBackground(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setDrawsBackground:"), value)
}

func (c_ ClipView) BackgroundColor() Color {
	rv := objc.CallMethod[Color](c_, objc.SEL("backgroundColor"))
	return rv
}

func (c_ ClipView) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setBackgroundColor:"), objc.ExtractPtr(value))
}
