// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[ClipView](c_, "initWithFrame:", frameRect)
	rv.Autorelease()
	return rv
}

func (c_ ClipView) Init() ClipView {
	rv := ffi.CallMethod[ClipView](c_, "init")
	rv.Autorelease()
	return rv
}

func (cc _ClipViewClass) Alloc() ClipView {
	rv := ffi.CallMethod[ClipView](cc, "alloc")
	return rv
}

func (cc _ClipViewClass) New() ClipView {
	rv := ffi.CallMethod[ClipView](cc, "new")
	rv.Autorelease()
	return rv
}

func NewClipView() ClipView {
	return ClipViewClass.New()
}

func (c_ ClipView) ScrollToPoint(newOrigin foundation.Point) {
	ffi.CallMethod[ffi.Void](c_, "scrollToPoint:", newOrigin)
}

// deprecated
func (c_ ClipView) ConstrainScrollPoint(newOrigin foundation.Point) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](c_, "constrainScrollPoint:", newOrigin)
	return rv
}

func (c_ ClipView) ConstrainBoundsRect(proposedBounds foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](c_, "constrainBoundsRect:", proposedBounds)
	return rv
}

func (c_ ClipView) ViewBoundsChanged(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](c_, "viewBoundsChanged:", notification)
}

func (c_ ClipView) ViewFrameChanged(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](c_, "viewFrameChanged:", notification)
}

func (c_ ClipView) DocumentView() View {
	rv := ffi.CallMethod[View](c_, "documentView")
	return rv
}

func (c_ ClipView) SetDocumentView(value IView) {
	ffi.CallMethod[ffi.Void](c_, "setDocumentView:", value)
}

// deprecated
func (c_ ClipView) CopiesOnScroll() bool {
	rv := ffi.CallMethod[bool](c_, "copiesOnScroll")
	return rv
}

// deprecated
func (c_ ClipView) SetCopiesOnScroll(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setCopiesOnScroll:", value)
}

func (c_ ClipView) ContentInsets() foundation.EdgeInsets {
	rv := ffi.CallMethod[foundation.EdgeInsets](c_, "contentInsets")
	return rv
}

func (c_ ClipView) SetContentInsets(value foundation.EdgeInsets) {
	ffi.CallMethod[ffi.Void](c_, "setContentInsets:", value)
}

func (c_ ClipView) AutomaticallyAdjustsContentInsets() bool {
	rv := ffi.CallMethod[bool](c_, "automaticallyAdjustsContentInsets")
	return rv
}

func (c_ ClipView) SetAutomaticallyAdjustsContentInsets(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setAutomaticallyAdjustsContentInsets:", value)
}

func (c_ ClipView) DocumentRect() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](c_, "documentRect")
	return rv
}

func (c_ ClipView) DocumentVisibleRect() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](c_, "documentVisibleRect")
	return rv
}

func (c_ ClipView) DocumentCursor() Cursor {
	rv := ffi.CallMethod[Cursor](c_, "documentCursor")
	return rv
}

func (c_ ClipView) SetDocumentCursor(value ICursor) {
	ffi.CallMethod[ffi.Void](c_, "setDocumentCursor:", value)
}

func (c_ ClipView) DrawsBackground() bool {
	rv := ffi.CallMethod[bool](c_, "drawsBackground")
	return rv
}

func (c_ ClipView) SetDrawsBackground(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setDrawsBackground:", value)
}

func (c_ ClipView) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](c_, "backgroundColor")
	return rv
}

func (c_ ClipView) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](c_, "setBackgroundColor:", value)
}
