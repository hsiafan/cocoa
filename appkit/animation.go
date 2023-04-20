// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var AnimationClass = _AnimationClass{objc.GetClass("NSAnimation")}

type _AnimationClass struct {
	objc.Class
}

type IAnimation interface {
	objc.IObject
	StartAnimation()
	StopAnimation()
	AddProgressMark(progressMark AnimationProgress)
	RemoveProgressMark(progressMark AnimationProgress)
	StartWhenAnimation_ReachesProgress(animation IAnimation, startProgress AnimationProgress)
	StopWhenAnimation_ReachesProgress(animation IAnimation, stopProgress AnimationProgress)
	ClearStartAnimation()
	ClearStopAnimation()
	AnimationBlockingMode() AnimationBlockingMode
	SetAnimationBlockingMode(value AnimationBlockingMode)
	RunLoopModesForAnimating() []foundation.RunLoopMode
	AnimationCurve() AnimationCurve
	SetAnimationCurve(value AnimationCurve)
	Duration() foundation.TimeInterval
	SetDuration(value foundation.TimeInterval)
	FrameRate() float32
	SetFrameRate(value float32)
	Delegate() AnimationDelegateWrapper
	SetDelegate(value AnimationDelegate)
	SetDelegate0(value objc.IObject)
	IsAnimating() bool
	CurrentProgress() AnimationProgress
	SetCurrentProgress(value AnimationProgress)
	CurrentValue() float32
	ProgressMarks() []foundation.Number
	SetProgressMarks(value []foundation.INumber)
}

type Animation struct {
	objc.Object
}

func MakeAnimation(ptr unsafe.Pointer) Animation {
	return Animation{
		Object: objc.MakeObject(ptr),
	}
}

func (a_ Animation) InitWithDuration_AnimationCurve(duration foundation.TimeInterval, animationCurve AnimationCurve) Animation {
	rv := ffi.CallMethod[Animation](a_, "initWithDuration:animationCurve:", duration, animationCurve)
	return rv
}

func (ac _AnimationClass) Alloc() Animation {
	rv := ffi.CallMethod[Animation](ac, "alloc")
	return rv
}

func (ac _AnimationClass) New() Animation {
	rv := ffi.CallMethod[Animation](ac, "new")
	rv.Autorelease()
	return rv
}

func NewAnimation() Animation {
	return AnimationClass.New()
}

func (a_ Animation) Init() Animation {
	rv := ffi.CallMethod[Animation](a_, "init")
	return rv
}

func (a_ Animation) StartAnimation() {
	ffi.CallMethod[ffi.Void](a_, "startAnimation")
}

func (a_ Animation) StopAnimation() {
	ffi.CallMethod[ffi.Void](a_, "stopAnimation")
}

func (a_ Animation) AddProgressMark(progressMark AnimationProgress) {
	ffi.CallMethod[ffi.Void](a_, "addProgressMark:", progressMark)
}

func (a_ Animation) RemoveProgressMark(progressMark AnimationProgress) {
	ffi.CallMethod[ffi.Void](a_, "removeProgressMark:", progressMark)
}

func (a_ Animation) StartWhenAnimation_ReachesProgress(animation IAnimation, startProgress AnimationProgress) {
	ffi.CallMethod[ffi.Void](a_, "startWhenAnimation:reachesProgress:", animation, startProgress)
}

func (a_ Animation) StopWhenAnimation_ReachesProgress(animation IAnimation, stopProgress AnimationProgress) {
	ffi.CallMethod[ffi.Void](a_, "stopWhenAnimation:reachesProgress:", animation, stopProgress)
}

func (a_ Animation) ClearStartAnimation() {
	ffi.CallMethod[ffi.Void](a_, "clearStartAnimation")
}

func (a_ Animation) ClearStopAnimation() {
	ffi.CallMethod[ffi.Void](a_, "clearStopAnimation")
}

func (a_ Animation) AnimationBlockingMode() AnimationBlockingMode {
	rv := ffi.CallMethod[AnimationBlockingMode](a_, "animationBlockingMode")
	return rv
}

func (a_ Animation) SetAnimationBlockingMode(value AnimationBlockingMode) {
	ffi.CallMethod[ffi.Void](a_, "setAnimationBlockingMode:", value)
}

func (a_ Animation) RunLoopModesForAnimating() []foundation.RunLoopMode {
	rv := ffi.CallMethod[[]foundation.RunLoopMode](a_, "runLoopModesForAnimating")
	return rv
}

func (a_ Animation) AnimationCurve() AnimationCurve {
	rv := ffi.CallMethod[AnimationCurve](a_, "animationCurve")
	return rv
}

func (a_ Animation) SetAnimationCurve(value AnimationCurve) {
	ffi.CallMethod[ffi.Void](a_, "setAnimationCurve:", value)
}

func (a_ Animation) Duration() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](a_, "duration")
	return rv
}

func (a_ Animation) SetDuration(value foundation.TimeInterval) {
	ffi.CallMethod[ffi.Void](a_, "setDuration:", value)
}

func (a_ Animation) FrameRate() float32 {
	rv := ffi.CallMethod[float32](a_, "frameRate")
	return rv
}

func (a_ Animation) SetFrameRate(value float32) {
	ffi.CallMethod[ffi.Void](a_, "setFrameRate:", value)
}

func (a_ Animation) Delegate() AnimationDelegateWrapper {
	rv := ffi.CallMethod[AnimationDelegateWrapper](a_, "delegate")
	return rv
}

func (a_ Animation) SetDelegate(value AnimationDelegate) {
	po := ffi.CreateProtocol("NSAnimationDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(a_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](a_, "setDelegate:", po)
}

func (a_ Animation) SetDelegate0(value objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "setDelegate:", value)
}

func (a_ Animation) IsAnimating() bool {
	rv := ffi.CallMethod[bool](a_, "isAnimating")
	return rv
}

func (a_ Animation) CurrentProgress() AnimationProgress {
	rv := ffi.CallMethod[AnimationProgress](a_, "currentProgress")
	return rv
}

func (a_ Animation) SetCurrentProgress(value AnimationProgress) {
	ffi.CallMethod[ffi.Void](a_, "setCurrentProgress:", value)
}

func (a_ Animation) CurrentValue() float32 {
	rv := ffi.CallMethod[float32](a_, "currentValue")
	return rv
}

func (a_ Animation) ProgressMarks() []foundation.Number {
	rv := ffi.CallMethod[[]foundation.Number](a_, "progressMarks")
	return rv
}

func (a_ Animation) SetProgressMarks(value []foundation.INumber) {
	ffi.CallMethod[ffi.Void](a_, "setProgressMarks:", value)
}
