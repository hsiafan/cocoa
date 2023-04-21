// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ScrollViewClass = _ScrollViewClass{objc.GetClass("NSScrollView")}

type _ScrollViewClass struct {
	objc.Class
}

type IScrollView interface {
	IView
	AddFloatingSubview_ForAxis(view IView, axis EventGestureAxis)
	Tile()
	FlashScrollers()
	MagnifyToFitRect(rect foundation.Rect)
	SetMagnification_CenteredAtPoint(magnification float64, point foundation.Point)
	ContentSize() foundation.Size
	DocumentVisibleRect() foundation.Rect
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	BorderType() BorderType
	SetBorderType(value BorderType)
	DocumentCursor() Cursor
	SetDocumentCursor(value ICursor)
	ContentView() ClipView
	SetContentView(value IClipView)
	DocumentView() View
	SetDocumentView(value IView)
	HorizontalScroller() Scroller
	SetHorizontalScroller(value IScroller)
	HasHorizontalScroller() bool
	SetHasHorizontalScroller(value bool)
	VerticalScroller() Scroller
	SetVerticalScroller(value IScroller)
	HasVerticalScroller() bool
	SetHasVerticalScroller(value bool)
	AutohidesScrollers() bool
	SetAutohidesScrollers(value bool)
	HasHorizontalRuler() bool
	SetHasHorizontalRuler(value bool)
	HorizontalRulerView() RulerView
	SetHorizontalRulerView(value IRulerView)
	HasVerticalRuler() bool
	SetHasVerticalRuler(value bool)
	VerticalRulerView() RulerView
	SetVerticalRulerView(value IRulerView)
	RulersVisible() bool
	SetRulersVisible(value bool)
	AutomaticallyAdjustsContentInsets() bool
	SetAutomaticallyAdjustsContentInsets(value bool)
	ContentInsets() foundation.EdgeInsets
	SetContentInsets(value foundation.EdgeInsets)
	ScrollerInsets() foundation.EdgeInsets
	SetScrollerInsets(value foundation.EdgeInsets)
	ScrollerKnobStyle() ScrollerKnobStyle
	SetScrollerKnobStyle(value ScrollerKnobStyle)
	ScrollerStyle() ScrollerStyle
	SetScrollerStyle(value ScrollerStyle)
	LineScroll() float64
	SetLineScroll(value float64)
	HorizontalLineScroll() float64
	SetHorizontalLineScroll(value float64)
	VerticalLineScroll() float64
	SetVerticalLineScroll(value float64)
	PageScroll() float64
	SetPageScroll(value float64)
	HorizontalPageScroll() float64
	SetHorizontalPageScroll(value float64)
	VerticalPageScroll() float64
	SetVerticalPageScroll(value float64)
	ScrollsDynamically() bool
	SetScrollsDynamically(value bool)
	FindBarPosition() ScrollViewFindBarPosition
	SetFindBarPosition(value ScrollViewFindBarPosition)
	UsesPredominantAxisScrolling() bool
	SetUsesPredominantAxisScrolling(value bool)
	HorizontalScrollElasticity() ScrollElasticity
	SetHorizontalScrollElasticity(value ScrollElasticity)
	VerticalScrollElasticity() ScrollElasticity
	SetVerticalScrollElasticity(value ScrollElasticity)
	AllowsMagnification() bool
	SetAllowsMagnification(value bool)
	Magnification() float64
	SetMagnification(value float64)
	MaxMagnification() float64
	SetMaxMagnification(value float64)
	MinMagnification() float64
	SetMinMagnification(value float64)
}

type ScrollView struct {
	View
}

func MakeScrollView(ptr unsafe.Pointer) ScrollView {
	return ScrollView{
		View: MakeView(ptr),
	}
}

func (s_ ScrollView) InitWithFrame(frameRect foundation.Rect) ScrollView {
	rv := objc.CallMethod[ScrollView](s_, "initWithFrame:", frameRect)
	return rv
}

func (s_ ScrollView) Init() ScrollView {
	rv := objc.CallMethod[ScrollView](s_, "init")
	return rv
}

func (sc _ScrollViewClass) Alloc() ScrollView {
	rv := objc.CallMethod[ScrollView](sc, "alloc")
	return rv
}

func (sc _ScrollViewClass) New() ScrollView {
	rv := objc.CallMethod[ScrollView](sc, "new")
	rv.Autorelease()
	return rv
}

func NewScrollView() ScrollView {
	return ScrollViewClass.New()
}

func (sc _ScrollViewClass) FrameSizeForContentSize_HorizontalScrollerClass_VerticalScrollerClass_BorderType_ControlSize_ScrollerStyle(cSize foundation.Size, horizontalScrollerClass objc.IClass, verticalScrollerClass objc.IClass, type_ BorderType, controlSize ControlSize, scrollerStyle ScrollerStyle) foundation.Size {
	rv := objc.CallMethod[foundation.Size](sc, "frameSizeForContentSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:", cSize, horizontalScrollerClass, verticalScrollerClass, type_, controlSize, scrollerStyle)
	return rv
}

func (sc _ScrollViewClass) ContentSizeForFrameSize_HorizontalScrollerClass_VerticalScrollerClass_BorderType_ControlSize_ScrollerStyle(fSize foundation.Size, horizontalScrollerClass objc.IClass, verticalScrollerClass objc.IClass, type_ BorderType, controlSize ControlSize, scrollerStyle ScrollerStyle) foundation.Size {
	rv := objc.CallMethod[foundation.Size](sc, "contentSizeForFrameSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:", fSize, horizontalScrollerClass, verticalScrollerClass, type_, controlSize, scrollerStyle)
	return rv
}

func (s_ ScrollView) AddFloatingSubview_ForAxis(view IView, axis EventGestureAxis) {
	objc.CallMethod[objc.Void](s_, "addFloatingSubview:forAxis:", view, axis)
}

func (s_ ScrollView) Tile() {
	objc.CallMethod[objc.Void](s_, "tile")
}

func (s_ ScrollView) FlashScrollers() {
	objc.CallMethod[objc.Void](s_, "flashScrollers")
}

func (s_ ScrollView) MagnifyToFitRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](s_, "magnifyToFitRect:", rect)
}

func (s_ ScrollView) SetMagnification_CenteredAtPoint(magnification float64, point foundation.Point) {
	objc.CallMethod[objc.Void](s_, "setMagnification:centeredAtPoint:", magnification, point)
}

// deprecated
func (sc _ScrollViewClass) FrameSizeForContentSize_HasHorizontalScroller_HasVerticalScroller_BorderType(cSize foundation.Size, hFlag bool, vFlag bool, type_ BorderType) foundation.Size {
	rv := objc.CallMethod[foundation.Size](sc, "frameSizeForContentSize:hasHorizontalScroller:hasVerticalScroller:borderType:", cSize, hFlag, vFlag, type_)
	return rv
}

// deprecated
func (sc _ScrollViewClass) ContentSizeForFrameSize_HasHorizontalScroller_HasVerticalScroller_BorderType(fSize foundation.Size, hFlag bool, vFlag bool, type_ BorderType) foundation.Size {
	rv := objc.CallMethod[foundation.Size](sc, "contentSizeForFrameSize:hasHorizontalScroller:hasVerticalScroller:borderType:", fSize, hFlag, vFlag, type_)
	return rv
}

func (s_ ScrollView) ContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](s_, "contentSize")
	return rv
}

func (s_ ScrollView) DocumentVisibleRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, "documentVisibleRect")
	return rv
}

func (s_ ScrollView) BackgroundColor() Color {
	rv := objc.CallMethod[Color](s_, "backgroundColor")
	return rv
}

func (s_ ScrollView) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](s_, "setBackgroundColor:", value)
}

func (s_ ScrollView) DrawsBackground() bool {
	rv := objc.CallMethod[bool](s_, "drawsBackground")
	return rv
}

func (s_ ScrollView) SetDrawsBackground(value bool) {
	objc.CallMethod[objc.Void](s_, "setDrawsBackground:", value)
}

func (s_ ScrollView) BorderType() BorderType {
	rv := objc.CallMethod[BorderType](s_, "borderType")
	return rv
}

func (s_ ScrollView) SetBorderType(value BorderType) {
	objc.CallMethod[objc.Void](s_, "setBorderType:", value)
}

func (s_ ScrollView) DocumentCursor() Cursor {
	rv := objc.CallMethod[Cursor](s_, "documentCursor")
	return rv
}

func (s_ ScrollView) SetDocumentCursor(value ICursor) {
	objc.CallMethod[objc.Void](s_, "setDocumentCursor:", value)
}

func (s_ ScrollView) ContentView() ClipView {
	rv := objc.CallMethod[ClipView](s_, "contentView")
	return rv
}

func (s_ ScrollView) SetContentView(value IClipView) {
	objc.CallMethod[objc.Void](s_, "setContentView:", value)
}

func (s_ ScrollView) DocumentView() View {
	rv := objc.CallMethod[View](s_, "documentView")
	return rv
}

func (s_ ScrollView) SetDocumentView(value IView) {
	objc.CallMethod[objc.Void](s_, "setDocumentView:", value)
}

func (s_ ScrollView) HorizontalScroller() Scroller {
	rv := objc.CallMethod[Scroller](s_, "horizontalScroller")
	return rv
}

func (s_ ScrollView) SetHorizontalScroller(value IScroller) {
	objc.CallMethod[objc.Void](s_, "setHorizontalScroller:", value)
}

func (s_ ScrollView) HasHorizontalScroller() bool {
	rv := objc.CallMethod[bool](s_, "hasHorizontalScroller")
	return rv
}

func (s_ ScrollView) SetHasHorizontalScroller(value bool) {
	objc.CallMethod[objc.Void](s_, "setHasHorizontalScroller:", value)
}

func (s_ ScrollView) VerticalScroller() Scroller {
	rv := objc.CallMethod[Scroller](s_, "verticalScroller")
	return rv
}

func (s_ ScrollView) SetVerticalScroller(value IScroller) {
	objc.CallMethod[objc.Void](s_, "setVerticalScroller:", value)
}

func (s_ ScrollView) HasVerticalScroller() bool {
	rv := objc.CallMethod[bool](s_, "hasVerticalScroller")
	return rv
}

func (s_ ScrollView) SetHasVerticalScroller(value bool) {
	objc.CallMethod[objc.Void](s_, "setHasVerticalScroller:", value)
}

func (s_ ScrollView) AutohidesScrollers() bool {
	rv := objc.CallMethod[bool](s_, "autohidesScrollers")
	return rv
}

func (s_ ScrollView) SetAutohidesScrollers(value bool) {
	objc.CallMethod[objc.Void](s_, "setAutohidesScrollers:", value)
}

func (sc _ScrollViewClass) RulerViewClass() objc.Class {
	rv := objc.CallMethod[objc.Class](sc, "rulerViewClass")
	return rv
}

func (sc _ScrollViewClass) SetRulerViewClass(value objc.IClass) {
	objc.CallMethod[objc.Void](sc, "setRulerViewClass:", value)
}

func (s_ ScrollView) HasHorizontalRuler() bool {
	rv := objc.CallMethod[bool](s_, "hasHorizontalRuler")
	return rv
}

func (s_ ScrollView) SetHasHorizontalRuler(value bool) {
	objc.CallMethod[objc.Void](s_, "setHasHorizontalRuler:", value)
}

func (s_ ScrollView) HorizontalRulerView() RulerView {
	rv := objc.CallMethod[RulerView](s_, "horizontalRulerView")
	return rv
}

func (s_ ScrollView) SetHorizontalRulerView(value IRulerView) {
	objc.CallMethod[objc.Void](s_, "setHorizontalRulerView:", value)
}

func (s_ ScrollView) HasVerticalRuler() bool {
	rv := objc.CallMethod[bool](s_, "hasVerticalRuler")
	return rv
}

func (s_ ScrollView) SetHasVerticalRuler(value bool) {
	objc.CallMethod[objc.Void](s_, "setHasVerticalRuler:", value)
}

func (s_ ScrollView) VerticalRulerView() RulerView {
	rv := objc.CallMethod[RulerView](s_, "verticalRulerView")
	return rv
}

func (s_ ScrollView) SetVerticalRulerView(value IRulerView) {
	objc.CallMethod[objc.Void](s_, "setVerticalRulerView:", value)
}

func (s_ ScrollView) RulersVisible() bool {
	rv := objc.CallMethod[bool](s_, "rulersVisible")
	return rv
}

func (s_ ScrollView) SetRulersVisible(value bool) {
	objc.CallMethod[objc.Void](s_, "setRulersVisible:", value)
}

func (s_ ScrollView) AutomaticallyAdjustsContentInsets() bool {
	rv := objc.CallMethod[bool](s_, "automaticallyAdjustsContentInsets")
	return rv
}

func (s_ ScrollView) SetAutomaticallyAdjustsContentInsets(value bool) {
	objc.CallMethod[objc.Void](s_, "setAutomaticallyAdjustsContentInsets:", value)
}

func (s_ ScrollView) ContentInsets() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](s_, "contentInsets")
	return rv
}

func (s_ ScrollView) SetContentInsets(value foundation.EdgeInsets) {
	objc.CallMethod[objc.Void](s_, "setContentInsets:", value)
}

func (s_ ScrollView) ScrollerInsets() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](s_, "scrollerInsets")
	return rv
}

func (s_ ScrollView) SetScrollerInsets(value foundation.EdgeInsets) {
	objc.CallMethod[objc.Void](s_, "setScrollerInsets:", value)
}

func (s_ ScrollView) ScrollerKnobStyle() ScrollerKnobStyle {
	rv := objc.CallMethod[ScrollerKnobStyle](s_, "scrollerKnobStyle")
	return rv
}

func (s_ ScrollView) SetScrollerKnobStyle(value ScrollerKnobStyle) {
	objc.CallMethod[objc.Void](s_, "setScrollerKnobStyle:", value)
}

func (s_ ScrollView) ScrollerStyle() ScrollerStyle {
	rv := objc.CallMethod[ScrollerStyle](s_, "scrollerStyle")
	return rv
}

func (s_ ScrollView) SetScrollerStyle(value ScrollerStyle) {
	objc.CallMethod[objc.Void](s_, "setScrollerStyle:", value)
}

func (s_ ScrollView) LineScroll() float64 {
	rv := objc.CallMethod[float64](s_, "lineScroll")
	return rv
}

func (s_ ScrollView) SetLineScroll(value float64) {
	objc.CallMethod[objc.Void](s_, "setLineScroll:", value)
}

func (s_ ScrollView) HorizontalLineScroll() float64 {
	rv := objc.CallMethod[float64](s_, "horizontalLineScroll")
	return rv
}

func (s_ ScrollView) SetHorizontalLineScroll(value float64) {
	objc.CallMethod[objc.Void](s_, "setHorizontalLineScroll:", value)
}

func (s_ ScrollView) VerticalLineScroll() float64 {
	rv := objc.CallMethod[float64](s_, "verticalLineScroll")
	return rv
}

func (s_ ScrollView) SetVerticalLineScroll(value float64) {
	objc.CallMethod[objc.Void](s_, "setVerticalLineScroll:", value)
}

func (s_ ScrollView) PageScroll() float64 {
	rv := objc.CallMethod[float64](s_, "pageScroll")
	return rv
}

func (s_ ScrollView) SetPageScroll(value float64) {
	objc.CallMethod[objc.Void](s_, "setPageScroll:", value)
}

func (s_ ScrollView) HorizontalPageScroll() float64 {
	rv := objc.CallMethod[float64](s_, "horizontalPageScroll")
	return rv
}

func (s_ ScrollView) SetHorizontalPageScroll(value float64) {
	objc.CallMethod[objc.Void](s_, "setHorizontalPageScroll:", value)
}

func (s_ ScrollView) VerticalPageScroll() float64 {
	rv := objc.CallMethod[float64](s_, "verticalPageScroll")
	return rv
}

func (s_ ScrollView) SetVerticalPageScroll(value float64) {
	objc.CallMethod[objc.Void](s_, "setVerticalPageScroll:", value)
}

func (s_ ScrollView) ScrollsDynamically() bool {
	rv := objc.CallMethod[bool](s_, "scrollsDynamically")
	return rv
}

func (s_ ScrollView) SetScrollsDynamically(value bool) {
	objc.CallMethod[objc.Void](s_, "setScrollsDynamically:", value)
}

func (s_ ScrollView) FindBarPosition() ScrollViewFindBarPosition {
	rv := objc.CallMethod[ScrollViewFindBarPosition](s_, "findBarPosition")
	return rv
}

func (s_ ScrollView) SetFindBarPosition(value ScrollViewFindBarPosition) {
	objc.CallMethod[objc.Void](s_, "setFindBarPosition:", value)
}

func (s_ ScrollView) UsesPredominantAxisScrolling() bool {
	rv := objc.CallMethod[bool](s_, "usesPredominantAxisScrolling")
	return rv
}

func (s_ ScrollView) SetUsesPredominantAxisScrolling(value bool) {
	objc.CallMethod[objc.Void](s_, "setUsesPredominantAxisScrolling:", value)
}

func (s_ ScrollView) HorizontalScrollElasticity() ScrollElasticity {
	rv := objc.CallMethod[ScrollElasticity](s_, "horizontalScrollElasticity")
	return rv
}

func (s_ ScrollView) SetHorizontalScrollElasticity(value ScrollElasticity) {
	objc.CallMethod[objc.Void](s_, "setHorizontalScrollElasticity:", value)
}

func (s_ ScrollView) VerticalScrollElasticity() ScrollElasticity {
	rv := objc.CallMethod[ScrollElasticity](s_, "verticalScrollElasticity")
	return rv
}

func (s_ ScrollView) SetVerticalScrollElasticity(value ScrollElasticity) {
	objc.CallMethod[objc.Void](s_, "setVerticalScrollElasticity:", value)
}

func (s_ ScrollView) AllowsMagnification() bool {
	rv := objc.CallMethod[bool](s_, "allowsMagnification")
	return rv
}

func (s_ ScrollView) SetAllowsMagnification(value bool) {
	objc.CallMethod[objc.Void](s_, "setAllowsMagnification:", value)
}

func (s_ ScrollView) Magnification() float64 {
	rv := objc.CallMethod[float64](s_, "magnification")
	return rv
}

func (s_ ScrollView) SetMagnification(value float64) {
	objc.CallMethod[objc.Void](s_, "setMagnification:", value)
}

func (s_ ScrollView) MaxMagnification() float64 {
	rv := objc.CallMethod[float64](s_, "maxMagnification")
	return rv
}

func (s_ ScrollView) SetMaxMagnification(value float64) {
	objc.CallMethod[objc.Void](s_, "setMaxMagnification:", value)
}

func (s_ ScrollView) MinMagnification() float64 {
	rv := objc.CallMethod[float64](s_, "minMagnification")
	return rv
}

func (s_ ScrollView) SetMinMagnification(value float64) {
	objc.CallMethod[objc.Void](s_, "setMinMagnification:", value)
}
