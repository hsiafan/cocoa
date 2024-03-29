// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[Scroller](s_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (s_ Scroller) Init() Scroller {
	rv := objc.CallMethod[Scroller](s_, objc.SEL("init"))
	return rv
}

func (sc _ScrollerClass) Alloc() Scroller {
	rv := objc.CallMethod[Scroller](sc, objc.SEL("alloc"))
	return rv
}

func (sc _ScrollerClass) New() Scroller {
	rv := objc.CallMethod[Scroller](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewScroller() Scroller {
	return ScrollerClass.New()
}

func (sc _ScrollerClass) ScrollerWidthForControlSize_ScrollerStyle(controlSize ControlSize, scrollerStyle ScrollerStyle) float64 {
	rv := objc.CallMethod[float64](sc, objc.SEL("scrollerWidthForControlSize:scrollerStyle:"), controlSize, scrollerStyle)
	return rv
}

// deprecated
func (sc _ScrollerClass) ScrollerWidth() float64 {
	rv := objc.CallMethod[float64](sc, objc.SEL("scrollerWidth"))
	return rv
}

// deprecated
func (sc _ScrollerClass) ScrollerWidthForControlSize(controlSize ControlSize) float64 {
	rv := objc.CallMethod[float64](sc, objc.SEL("scrollerWidthForControlSize:"), controlSize)
	return rv
}

// deprecated
func (s_ Scroller) SetFloatValue_KnobProportion(value float32, proportion float64) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setFloatValue:knobProportion:"), value, proportion)
}

func (s_ Scroller) SetKnobProportion(proportion float64) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setKnobProportion:"), proportion)
}

func (s_ Scroller) RectForPart(partCode ScrollerPart) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.SEL("rectForPart:"), partCode)
	return rv
}

func (s_ Scroller) TestPart(point foundation.Point) ScrollerPart {
	rv := objc.CallMethod[ScrollerPart](s_, objc.SEL("testPart:"), point)
	return rv
}

func (s_ Scroller) CheckSpaceForParts() {
	objc.CallMethod[objc.Void](s_, objc.SEL("checkSpaceForParts"))
}

// deprecated
func (s_ Scroller) DrawArrow_Highlight(whichArrow ScrollerArrow, flag bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("drawArrow:highlight:"), whichArrow, flag)
}

func (s_ Scroller) DrawKnobSlotInRect_Highlight(slotRect foundation.Rect, flag bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("drawKnobSlotInRect:highlight:"), slotRect, flag)
}

func (s_ Scroller) DrawKnob() {
	objc.CallMethod[objc.Void](s_, objc.SEL("drawKnob"))
}

// deprecated
func (s_ Scroller) DrawParts() {
	objc.CallMethod[objc.Void](s_, objc.SEL("drawParts"))
}

// deprecated
func (s_ Scroller) Highlight(flag bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("highlight:"), flag)
}

func (s_ Scroller) TrackKnob(event IEvent) {
	objc.CallMethod[objc.Void](s_, objc.SEL("trackKnob:"), objc.ExtractPtr(event))
}

// deprecated
func (s_ Scroller) TrackScrollButtons(event IEvent) {
	objc.CallMethod[objc.Void](s_, objc.SEL("trackScrollButtons:"), objc.ExtractPtr(event))
}

// deprecated
func (s_ Scroller) ArrowsPosition() ScrollArrowPosition {
	rv := objc.CallMethod[ScrollArrowPosition](s_, objc.SEL("arrowsPosition"))
	return rv
}

// deprecated
func (s_ Scroller) SetArrowsPosition(value ScrollArrowPosition) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setArrowsPosition:"), value)
}

func (s_ Scroller) UsableParts() UsableScrollerParts {
	rv := objc.CallMethod[UsableScrollerParts](s_, objc.SEL("usableParts"))
	return rv
}

func (s_ Scroller) HitPart() ScrollerPart {
	rv := objc.CallMethod[ScrollerPart](s_, objc.SEL("hitPart"))
	return rv
}

// deprecated
func (s_ Scroller) ControlTint() ControlTint {
	rv := objc.CallMethod[ControlTint](s_, objc.SEL("controlTint"))
	return rv
}

// deprecated
func (s_ Scroller) SetControlTint(value ControlTint) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setControlTint:"), value)
}

func (sc _ScrollerClass) PreferredScrollerStyle() ScrollerStyle {
	rv := objc.CallMethod[ScrollerStyle](sc, objc.SEL("preferredScrollerStyle"))
	return rv
}

func (s_ Scroller) ScrollerStyle() ScrollerStyle {
	rv := objc.CallMethod[ScrollerStyle](s_, objc.SEL("scrollerStyle"))
	return rv
}

func (s_ Scroller) SetScrollerStyle(value ScrollerStyle) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setScrollerStyle:"), value)
}

func (s_ Scroller) KnobStyle() ScrollerKnobStyle {
	rv := objc.CallMethod[ScrollerKnobStyle](s_, objc.SEL("knobStyle"))
	return rv
}

func (s_ Scroller) SetKnobStyle(value ScrollerKnobStyle) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setKnobStyle:"), value)
}

func (s_ Scroller) KnobProportion() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("knobProportion"))
	return rv
}

func (sc _ScrollerClass) IsCompatibleWithOverlayScrollers() bool {
	rv := objc.CallMethod[bool](sc, objc.SEL("isCompatibleWithOverlayScrollers"))
	return rv
}
