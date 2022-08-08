// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[ProgressIndicator](p_, "initWithFrame:", frameRect)
	return rv
}

func (p_ ProgressIndicator) Init() ProgressIndicator {
	rv := ffi.CallMethod[ProgressIndicator](p_, "init")
	return rv
}

func (pc _ProgressIndicatorClass) Alloc() ProgressIndicator {
	rv := ffi.CallMethod[ProgressIndicator](pc, "alloc")
	return rv
}

func (pc _ProgressIndicatorClass) New() ProgressIndicator {
	rv := ffi.CallMethod[ProgressIndicator](pc, "new")
	rv.Autorelease()
	return rv
}

func NewProgressIndicator() ProgressIndicator {
	return ProgressIndicatorClass.New()
}

func (p_ ProgressIndicator) StartAnimation(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](p_, "startAnimation:", sender)
}

func (p_ ProgressIndicator) StopAnimation(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](p_, "stopAnimation:", sender)
}

// deprecated
func (p_ ProgressIndicator) Animate(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](p_, "animate:", sender)
}

// deprecated
func (p_ ProgressIndicator) AnimationDelay() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](p_, "animationDelay")
	return rv
}

// deprecated
func (p_ ProgressIndicator) SetAnimationDelay(delay foundation.TimeInterval) {
	ffi.CallMethod[ffi.Void](p_, "setAnimationDelay:", delay)
}

func (p_ ProgressIndicator) IncrementBy(delta float64) {
	ffi.CallMethod[ffi.Void](p_, "incrementBy:", delta)
}

func (p_ ProgressIndicator) SizeToFit() {
	ffi.CallMethod[ffi.Void](p_, "sizeToFit")
}

func (p_ ProgressIndicator) UsesThreadedAnimation() bool {
	rv := ffi.CallMethod[bool](p_, "usesThreadedAnimation")
	return rv
}

func (p_ ProgressIndicator) SetUsesThreadedAnimation(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setUsesThreadedAnimation:", value)
}

func (p_ ProgressIndicator) DoubleValue() float64 {
	rv := ffi.CallMethod[float64](p_, "doubleValue")
	return rv
}

func (p_ ProgressIndicator) SetDoubleValue(value float64) {
	ffi.CallMethod[ffi.Void](p_, "setDoubleValue:", value)
}

func (p_ ProgressIndicator) MinValue() float64 {
	rv := ffi.CallMethod[float64](p_, "minValue")
	return rv
}

func (p_ ProgressIndicator) SetMinValue(value float64) {
	ffi.CallMethod[ffi.Void](p_, "setMinValue:", value)
}

func (p_ ProgressIndicator) MaxValue() float64 {
	rv := ffi.CallMethod[float64](p_, "maxValue")
	return rv
}

func (p_ ProgressIndicator) SetMaxValue(value float64) {
	ffi.CallMethod[ffi.Void](p_, "setMaxValue:", value)
}

func (p_ ProgressIndicator) ControlSize() ControlSize {
	rv := ffi.CallMethod[ControlSize](p_, "controlSize")
	return rv
}

func (p_ ProgressIndicator) SetControlSize(value ControlSize) {
	ffi.CallMethod[ffi.Void](p_, "setControlSize:", value)
}

func (p_ ProgressIndicator) ControlTint() ControlTint {
	rv := ffi.CallMethod[ControlTint](p_, "controlTint")
	return rv
}

func (p_ ProgressIndicator) SetControlTint(value ControlTint) {
	ffi.CallMethod[ffi.Void](p_, "setControlTint:", value)
}

func (p_ ProgressIndicator) IsBezeled() bool {
	rv := ffi.CallMethod[bool](p_, "isBezeled")
	return rv
}

func (p_ ProgressIndicator) SetBezeled(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setBezeled:", value)
}

func (p_ ProgressIndicator) IsIndeterminate() bool {
	rv := ffi.CallMethod[bool](p_, "isIndeterminate")
	return rv
}

func (p_ ProgressIndicator) SetIndeterminate(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setIndeterminate:", value)
}

func (p_ ProgressIndicator) Style() ProgressIndicatorStyle {
	rv := ffi.CallMethod[ProgressIndicatorStyle](p_, "style")
	return rv
}

func (p_ ProgressIndicator) SetStyle(value ProgressIndicatorStyle) {
	ffi.CallMethod[ffi.Void](p_, "setStyle:", value)
}

func (p_ ProgressIndicator) IsDisplayedWhenStopped() bool {
	rv := ffi.CallMethod[bool](p_, "isDisplayedWhenStopped")
	return rv
}

func (p_ ProgressIndicator) SetDisplayedWhenStopped(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setDisplayedWhenStopped:", value)
}
