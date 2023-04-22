// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[Animation](a_, objc.GetSelector("initWithDuration:animationCurve:"), duration, animationCurve)
	return rv
}

func (ac _AnimationClass) Alloc() Animation {
	rv := objc.CallMethod[Animation](ac, objc.GetSelector("alloc"))
	return rv
}

func (ac _AnimationClass) New() Animation {
	rv := objc.CallMethod[Animation](ac, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewAnimation() Animation {
	return AnimationClass.New()
}

func (a_ Animation) Init() Animation {
	rv := objc.CallMethod[Animation](a_, objc.GetSelector("init"))
	return rv
}

func (a_ Animation) StartAnimation() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("startAnimation"))
}

func (a_ Animation) StopAnimation() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("stopAnimation"))
}

func (a_ Animation) AddProgressMark(progressMark AnimationProgress) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("addProgressMark:"), progressMark)
}

func (a_ Animation) RemoveProgressMark(progressMark AnimationProgress) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("removeProgressMark:"), progressMark)
}

func (a_ Animation) StartWhenAnimation_ReachesProgress(animation IAnimation, startProgress AnimationProgress) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("startWhenAnimation:reachesProgress:"), animation, startProgress)
}

func (a_ Animation) StopWhenAnimation_ReachesProgress(animation IAnimation, stopProgress AnimationProgress) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("stopWhenAnimation:reachesProgress:"), animation, stopProgress)
}

func (a_ Animation) ClearStartAnimation() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("clearStartAnimation"))
}

func (a_ Animation) ClearStopAnimation() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("clearStopAnimation"))
}

func (a_ Animation) AnimationBlockingMode() AnimationBlockingMode {
	rv := objc.CallMethod[AnimationBlockingMode](a_, objc.GetSelector("animationBlockingMode"))
	return rv
}

func (a_ Animation) SetAnimationBlockingMode(value AnimationBlockingMode) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setAnimationBlockingMode:"), value)
}

func (a_ Animation) RunLoopModesForAnimating() []foundation.RunLoopMode {
	rv := objc.CallMethod[[]foundation.RunLoopMode](a_, objc.GetSelector("runLoopModesForAnimating"))
	return rv
}

func (a_ Animation) AnimationCurve() AnimationCurve {
	rv := objc.CallMethod[AnimationCurve](a_, objc.GetSelector("animationCurve"))
	return rv
}

func (a_ Animation) SetAnimationCurve(value AnimationCurve) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setAnimationCurve:"), value)
}

func (a_ Animation) Duration() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](a_, objc.GetSelector("duration"))
	return rv
}

func (a_ Animation) SetDuration(value foundation.TimeInterval) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setDuration:"), value)
}

func (a_ Animation) FrameRate() float32 {
	rv := objc.CallMethod[float32](a_, objc.GetSelector("frameRate"))
	return rv
}

func (a_ Animation) SetFrameRate(value float32) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setFrameRate:"), value)
}

func (a_ Animation) Delegate() AnimationDelegateWrapper {
	rv := objc.CallMethod[AnimationDelegateWrapper](a_, objc.GetSelector("delegate"))
	return rv
}

func (a_ Animation) SetDelegate(value AnimationDelegate) {
	po := objc.CreateProtocol("NSAnimationDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(a_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setDelegate:"), po)
}

func (a_ Animation) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setDelegate:"), value)
}

func (a_ Animation) IsAnimating() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("isAnimating"))
	return rv
}

func (a_ Animation) CurrentProgress() AnimationProgress {
	rv := objc.CallMethod[AnimationProgress](a_, objc.GetSelector("currentProgress"))
	return rv
}

func (a_ Animation) SetCurrentProgress(value AnimationProgress) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setCurrentProgress:"), value)
}

func (a_ Animation) CurrentValue() float32 {
	rv := objc.CallMethod[float32](a_, objc.GetSelector("currentValue"))
	return rv
}

func (a_ Animation) ProgressMarks() []foundation.Number {
	rv := objc.CallMethod[[]foundation.Number](a_, objc.GetSelector("progressMarks"))
	return rv
}

func (a_ Animation) SetProgressMarks(value []foundation.INumber) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setProgressMarks:"), value)
}
