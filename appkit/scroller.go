// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ScrollerClass = _ScrollerClass{objc.GetClass("NSScroller")}

type _ScrollerClass struct {
	objc.Class
}

type IScroller interface {
	IControl
	// deprecated
	SetFloatValue_KnobProportion(value float32, proportion float64)
	SetKnobProportion(proportion float64)
	RectForPart(partCode ScrollerPart) foundation.Rect
	TestPart(point foundation.Point) ScrollerPart
	CheckSpaceForParts()
	// deprecated
	DrawArrow_Highlight(whichArrow ScrollerArrow, flag bool)
	DrawKnobSlotInRect_Highlight(slotRect foundation.Rect, flag bool)
	DrawKnob()
	// deprecated
	DrawParts()
	// deprecated
	Highlight(flag bool)
	TrackKnob(event IEvent)
	// deprecated
	TrackScrollButtons(event IEvent)
	// deprecated
	ArrowsPosition() ScrollArrowPosition
	// deprecated
	SetArrowsPosition(value ScrollArrowPosition)
	UsableParts() UsableScrollerParts
	HitPart() ScrollerPart
	// deprecated
	ControlTint() ControlTint
	// deprecated
	SetControlTint(value ControlTint)
	ScrollerStyle() ScrollerStyle
	SetScrollerStyle(value ScrollerStyle)
	KnobStyle() ScrollerKnobStyle
	SetKnobStyle(value ScrollerKnobStyle)
	KnobProportion() float64
}

type Scroller struct {
	Control
}

func MakeScroller(ptr unsafe.Pointer) Scroller {
	return Scroller{
		Control: MakeControl(ptr),
	}
}

func (s_ Scroller) InitWithFrame(frameRect foundation.Rect) Scroller {
	rv := ffi.CallMethod[Scroller](s_, "initWithFrame:", frameRect)
	return rv
}

func (s_ Scroller) Init() Scroller {
	rv := ffi.CallMethod[Scroller](s_, "init")
	return rv
}

func (sc _ScrollerClass) Alloc() Scroller {
	rv := ffi.CallMethod[Scroller](sc, "alloc")
	return rv
}

func (sc _ScrollerClass) New() Scroller {
	rv := ffi.CallMethod[Scroller](sc, "new")
	rv.Autorelease()
	return rv
}

func NewScroller() Scroller {
	return ScrollerClass.New()
}

func (sc _ScrollerClass) ScrollerWidthForControlSize_ScrollerStyle(controlSize ControlSize, scrollerStyle ScrollerStyle) float64 {
	rv := ffi.CallMethod[float64](sc, "scrollerWidthForControlSize:scrollerStyle:", controlSize, scrollerStyle)
	return rv
}

// deprecated
func (sc _ScrollerClass) ScrollerWidth() float64 {
	rv := ffi.CallMethod[float64](sc, "scrollerWidth")
	return rv
}

// deprecated
func (sc _ScrollerClass) ScrollerWidthForControlSize(controlSize ControlSize) float64 {
	rv := ffi.CallMethod[float64](sc, "scrollerWidthForControlSize:", controlSize)
	return rv
}

// deprecated
func (s_ Scroller) SetFloatValue_KnobProportion(value float32, proportion float64) {
	ffi.CallMethod[ffi.Void](s_, "setFloatValue:knobProportion:", value, proportion)
}

func (s_ Scroller) SetKnobProportion(proportion float64) {
	ffi.CallMethod[ffi.Void](s_, "setKnobProportion:", proportion)
}

func (s_ Scroller) RectForPart(partCode ScrollerPart) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "rectForPart:", partCode)
	return rv
}

func (s_ Scroller) TestPart(point foundation.Point) ScrollerPart {
	rv := ffi.CallMethod[ScrollerPart](s_, "testPart:", point)
	return rv
}

func (s_ Scroller) CheckSpaceForParts() {
	ffi.CallMethod[ffi.Void](s_, "checkSpaceForParts")
}

// deprecated
func (s_ Scroller) DrawArrow_Highlight(whichArrow ScrollerArrow, flag bool) {
	ffi.CallMethod[ffi.Void](s_, "drawArrow:highlight:", whichArrow, flag)
}

func (s_ Scroller) DrawKnobSlotInRect_Highlight(slotRect foundation.Rect, flag bool) {
	ffi.CallMethod[ffi.Void](s_, "drawKnobSlotInRect:highlight:", slotRect, flag)
}

func (s_ Scroller) DrawKnob() {
	ffi.CallMethod[ffi.Void](s_, "drawKnob")
}

// deprecated
func (s_ Scroller) DrawParts() {
	ffi.CallMethod[ffi.Void](s_, "drawParts")
}

// deprecated
func (s_ Scroller) Highlight(flag bool) {
	ffi.CallMethod[ffi.Void](s_, "highlight:", flag)
}

func (s_ Scroller) TrackKnob(event IEvent) {
	ffi.CallMethod[ffi.Void](s_, "trackKnob:", event)
}

// deprecated
func (s_ Scroller) TrackScrollButtons(event IEvent) {
	ffi.CallMethod[ffi.Void](s_, "trackScrollButtons:", event)
}

// deprecated
func (s_ Scroller) ArrowsPosition() ScrollArrowPosition {
	rv := ffi.CallMethod[ScrollArrowPosition](s_, "arrowsPosition")
	return rv
}

// deprecated
func (s_ Scroller) SetArrowsPosition(value ScrollArrowPosition) {
	ffi.CallMethod[ffi.Void](s_, "setArrowsPosition:", value)
}

func (s_ Scroller) UsableParts() UsableScrollerParts {
	rv := ffi.CallMethod[UsableScrollerParts](s_, "usableParts")
	return rv
}

func (s_ Scroller) HitPart() ScrollerPart {
	rv := ffi.CallMethod[ScrollerPart](s_, "hitPart")
	return rv
}

// deprecated
func (s_ Scroller) ControlTint() ControlTint {
	rv := ffi.CallMethod[ControlTint](s_, "controlTint")
	return rv
}

// deprecated
func (s_ Scroller) SetControlTint(value ControlTint) {
	ffi.CallMethod[ffi.Void](s_, "setControlTint:", value)
}

func (sc _ScrollerClass) PreferredScrollerStyle() ScrollerStyle {
	rv := ffi.CallMethod[ScrollerStyle](sc, "preferredScrollerStyle")
	return rv
}

func (s_ Scroller) ScrollerStyle() ScrollerStyle {
	rv := ffi.CallMethod[ScrollerStyle](s_, "scrollerStyle")
	return rv
}

func (s_ Scroller) SetScrollerStyle(value ScrollerStyle) {
	ffi.CallMethod[ffi.Void](s_, "setScrollerStyle:", value)
}

func (s_ Scroller) KnobStyle() ScrollerKnobStyle {
	rv := ffi.CallMethod[ScrollerKnobStyle](s_, "knobStyle")
	return rv
}

func (s_ Scroller) SetKnobStyle(value ScrollerKnobStyle) {
	ffi.CallMethod[ffi.Void](s_, "setKnobStyle:", value)
}

func (s_ Scroller) KnobProportion() float64 {
	rv := ffi.CallMethod[float64](s_, "knobProportion")
	return rv
}

func (sc _ScrollerClass) IsCompatibleWithOverlayScrollers() bool {
	rv := ffi.CallMethod[bool](sc, "isCompatibleWithOverlayScrollers")
	return rv
}
