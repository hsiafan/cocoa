// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ProgressIndicatorClass = _ProgressIndicatorClass{objc.GetClass("NSProgressIndicator")}

type _ProgressIndicatorClass struct {
	objc.Class
}

type IProgressIndicator interface {
	IView
	StartAnimation(sender objc.IObject)
	StopAnimation(sender objc.IObject)
	// deprecated
	Animate(sender objc.IObject)
	// deprecated
	AnimationDelay() foundation.TimeInterval
	// deprecated
	SetAnimationDelay(delay foundation.TimeInterval)
	IncrementBy(delta float64)
	SizeToFit()
	UsesThreadedAnimation() bool
	SetUsesThreadedAnimation(value bool)
	DoubleValue() float64
	SetDoubleValue(value float64)
	MinValue() float64
	SetMinValue(value float64)
	MaxValue() float64
	SetMaxValue(value float64)
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	ControlTint() ControlTint
	SetControlTint(value ControlTint)
	IsBezeled() bool
	SetBezeled(value bool)
	IsIndeterminate() bool
	SetIndeterminate(value bool)
	Style() ProgressIndicatorStyle
	SetStyle(value ProgressIndicatorStyle)
	IsDisplayedWhenStopped() bool
	SetDisplayedWhenStopped(value bool)
}

type ProgressIndicator struct {
	View
}

func MakeProgressIndicator(ptr unsafe.Pointer) ProgressIndicator {
	return ProgressIndicator{
		View: MakeView(ptr),
	}
}

func (p_ ProgressIndicator) InitWithFrame(frameRect foundation.Rect) ProgressIndicator {
	rv := objc.CallMethod[ProgressIndicator](p_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (p_ ProgressIndicator) Init() ProgressIndicator {
	rv := objc.CallMethod[ProgressIndicator](p_, objc.SEL("init"))
	return rv
}

func (pc _ProgressIndicatorClass) Alloc() ProgressIndicator {
	rv := objc.CallMethod[ProgressIndicator](pc, objc.SEL("alloc"))
	return rv
}

func (pc _ProgressIndicatorClass) New() ProgressIndicator {
	rv := objc.CallMethod[ProgressIndicator](pc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewProgressIndicator() ProgressIndicator {
	return ProgressIndicatorClass.New()
}

func (p_ ProgressIndicator) StartAnimation(sender objc.IObject) {
	objc.CallMethod[objc.Void](p_, objc.SEL("startAnimation:"), objc.ExtractPtr(sender))
}

func (p_ ProgressIndicator) StopAnimation(sender objc.IObject) {
	objc.CallMethod[objc.Void](p_, objc.SEL("stopAnimation:"), objc.ExtractPtr(sender))
}

// deprecated
func (p_ ProgressIndicator) Animate(sender objc.IObject) {
	objc.CallMethod[objc.Void](p_, objc.SEL("animate:"), objc.ExtractPtr(sender))
}

// deprecated
func (p_ ProgressIndicator) AnimationDelay() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](p_, objc.SEL("animationDelay"))
	return rv
}

// deprecated
func (p_ ProgressIndicator) SetAnimationDelay(delay foundation.TimeInterval) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setAnimationDelay:"), delay)
}

func (p_ ProgressIndicator) IncrementBy(delta float64) {
	objc.CallMethod[objc.Void](p_, objc.SEL("incrementBy:"), delta)
}

func (p_ ProgressIndicator) SizeToFit() {
	objc.CallMethod[objc.Void](p_, objc.SEL("sizeToFit"))
}

func (p_ ProgressIndicator) UsesThreadedAnimation() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("usesThreadedAnimation"))
	return rv
}

func (p_ ProgressIndicator) SetUsesThreadedAnimation(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setUsesThreadedAnimation:"), value)
}

func (p_ ProgressIndicator) DoubleValue() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("doubleValue"))
	return rv
}

func (p_ ProgressIndicator) SetDoubleValue(value float64) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setDoubleValue:"), value)
}

func (p_ ProgressIndicator) MinValue() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("minValue"))
	return rv
}

func (p_ ProgressIndicator) SetMinValue(value float64) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setMinValue:"), value)
}

func (p_ ProgressIndicator) MaxValue() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("maxValue"))
	return rv
}

func (p_ ProgressIndicator) SetMaxValue(value float64) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setMaxValue:"), value)
}

func (p_ ProgressIndicator) ControlSize() ControlSize {
	rv := objc.CallMethod[ControlSize](p_, objc.SEL("controlSize"))
	return rv
}

func (p_ ProgressIndicator) SetControlSize(value ControlSize) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setControlSize:"), value)
}

func (p_ ProgressIndicator) ControlTint() ControlTint {
	rv := objc.CallMethod[ControlTint](p_, objc.SEL("controlTint"))
	return rv
}

func (p_ ProgressIndicator) SetControlTint(value ControlTint) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setControlTint:"), value)
}

func (p_ ProgressIndicator) IsBezeled() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isBezeled"))
	return rv
}

func (p_ ProgressIndicator) SetBezeled(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setBezeled:"), value)
}

func (p_ ProgressIndicator) IsIndeterminate() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isIndeterminate"))
	return rv
}

func (p_ ProgressIndicator) SetIndeterminate(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setIndeterminate:"), value)
}

func (p_ ProgressIndicator) Style() ProgressIndicatorStyle {
	rv := objc.CallMethod[ProgressIndicatorStyle](p_, objc.SEL("style"))
	return rv
}

func (p_ ProgressIndicator) SetStyle(value ProgressIndicatorStyle) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setStyle:"), value)
}

func (p_ ProgressIndicator) IsDisplayedWhenStopped() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isDisplayedWhenStopped"))
	return rv
}

func (p_ ProgressIndicator) SetDisplayedWhenStopped(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setDisplayedWhenStopped:"), value)
}
