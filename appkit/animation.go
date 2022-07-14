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
	rv.Autorelease()
	return rv
}

func (ac _AnimationClass) Alloc() Animation {
	rv := ffi.CallMethod[Animation](ac, "alloc")
	return rv
}

func (a_ Animation) Init() Animation {
	rv := ffi.CallMethod[Animation](a_, "init")
	rv.Autorelease()
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
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(a_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](a_, "setDelegate:", po)
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

var ViewAnimationClass = _ViewAnimationClass{objc.GetClass("NSViewAnimation")}

type _ViewAnimationClass struct {
	objc.Class
}

type IViewAnimation interface {
	IAnimation
	ViewAnimations() []map[ViewAnimationKey]objc.Object
	SetViewAnimations(value []map[ViewAnimationKey]objc.IObject)
}

type ViewAnimation struct {
	Animation
}

func MakeViewAnimation(ptr unsafe.Pointer) ViewAnimation {
	return ViewAnimation{
		Animation: MakeAnimation(ptr),
	}
}

func (v_ ViewAnimation) InitWithViewAnimations(viewAnimations []map[ViewAnimationKey]objc.IObject) ViewAnimation {
	rv := ffi.CallMethod[ViewAnimation](v_, "initWithViewAnimations:", viewAnimations)
	rv.Autorelease()
	return rv
}

func (v_ ViewAnimation) InitWithDuration_AnimationCurve(duration foundation.TimeInterval, animationCurve AnimationCurve) ViewAnimation {
	rv := ffi.CallMethod[ViewAnimation](v_, "initWithDuration:animationCurve:", duration, animationCurve)
	rv.Autorelease()
	return rv
}

func (vc _ViewAnimationClass) Alloc() ViewAnimation {
	rv := ffi.CallMethod[ViewAnimation](vc, "alloc")
	return rv
}

func (v_ ViewAnimation) Init() ViewAnimation {
	rv := ffi.CallMethod[ViewAnimation](v_, "init")
	rv.Autorelease()
	return rv
}

func (vc _ViewAnimationClass) New() ViewAnimation {
	rv := ffi.CallMethod[ViewAnimation](vc, "new")
	rv.Autorelease()
	return rv
}

func NewViewAnimation() ViewAnimation {
	return ViewAnimationClass.New()
}

func (v_ ViewAnimation) ViewAnimations() []map[ViewAnimationKey]objc.Object {
	rv := ffi.CallMethod[[]map[ViewAnimationKey]objc.Object](v_, "viewAnimations")
	return rv
}

func (v_ ViewAnimation) SetViewAnimations(value []map[ViewAnimationKey]objc.IObject) {
	ffi.CallMethod[ffi.Void](v_, "setViewAnimations:", value)
}

type AnimationDelegate interface {
	ImplementsAnimationDidEnd() bool
	// optional
	AnimationDidEnd(animation Animation)
	ImplementsAnimationDidStop() bool
	// optional
	AnimationDidStop(animation Animation)
	ImplementsAnimationShouldStart() bool
	// optional
	AnimationShouldStart(animation Animation) bool
	ImplementsAnimation_ValueForProgress() bool
	// optional
	Animation_ValueForProgress(animation Animation, progress AnimationProgress) float32
	ImplementsAnimation_DidReachProgressMark() bool
	// optional
	Animation_DidReachProgressMark(animation Animation, progress AnimationProgress)
}

type AnimationDelegateImpl struct {
	_AnimationDidEnd                func(animation Animation)
	_AnimationDidStop               func(animation Animation)
	_AnimationShouldStart           func(animation Animation) bool
	_Animation_ValueForProgress     func(animation Animation, progress AnimationProgress) float32
	_Animation_DidReachProgressMark func(animation Animation, progress AnimationProgress)
}

func (di *AnimationDelegateImpl) ImplementsAnimationDidEnd() bool {
	return di._AnimationDidEnd != nil
}

func (di *AnimationDelegateImpl) SetAnimationDidEnd(f func(animation Animation)) {
	di._AnimationDidEnd = f
}

func (di *AnimationDelegateImpl) AnimationDidEnd(animation Animation) {
	di._AnimationDidEnd(animation)
}
func (di *AnimationDelegateImpl) ImplementsAnimationDidStop() bool {
	return di._AnimationDidStop != nil
}

func (di *AnimationDelegateImpl) SetAnimationDidStop(f func(animation Animation)) {
	di._AnimationDidStop = f
}

func (di *AnimationDelegateImpl) AnimationDidStop(animation Animation) {
	di._AnimationDidStop(animation)
}
func (di *AnimationDelegateImpl) ImplementsAnimationShouldStart() bool {
	return di._AnimationShouldStart != nil
}

func (di *AnimationDelegateImpl) SetAnimationShouldStart(f func(animation Animation) bool) {
	di._AnimationShouldStart = f
}

func (di *AnimationDelegateImpl) AnimationShouldStart(animation Animation) bool {
	return di._AnimationShouldStart(animation)
}
func (di *AnimationDelegateImpl) ImplementsAnimation_ValueForProgress() bool {
	return di._Animation_ValueForProgress != nil
}

func (di *AnimationDelegateImpl) SetAnimation_ValueForProgress(f func(animation Animation, progress AnimationProgress) float32) {
	di._Animation_ValueForProgress = f
}

func (di *AnimationDelegateImpl) Animation_ValueForProgress(animation Animation, progress AnimationProgress) float32 {
	return di._Animation_ValueForProgress(animation, progress)
}
func (di *AnimationDelegateImpl) ImplementsAnimation_DidReachProgressMark() bool {
	return di._Animation_DidReachProgressMark != nil
}

func (di *AnimationDelegateImpl) SetAnimation_DidReachProgressMark(f func(animation Animation, progress AnimationProgress)) {
	di._Animation_DidReachProgressMark = f
}

func (di *AnimationDelegateImpl) Animation_DidReachProgressMark(animation Animation, progress AnimationProgress) {
	di._Animation_DidReachProgressMark(animation, progress)
}

type AnimationDelegateWrapper struct {
	objc.Object
}

func (a_ *AnimationDelegateWrapper) ImplementsAnimationDidEnd() bool {
	return a_.RespondsToSelector(objc.GetSelector("animationDidEnd:"))
}

func (a_ AnimationDelegateWrapper) AnimationDidEnd(animation IAnimation) {
	ffi.CallMethod[ffi.Void](a_, "animationDidEnd:", animation)
}

func (a_ *AnimationDelegateWrapper) ImplementsAnimationDidStop() bool {
	return a_.RespondsToSelector(objc.GetSelector("animationDidStop:"))
}

func (a_ AnimationDelegateWrapper) AnimationDidStop(animation IAnimation) {
	ffi.CallMethod[ffi.Void](a_, "animationDidStop:", animation)
}

func (a_ *AnimationDelegateWrapper) ImplementsAnimationShouldStart() bool {
	return a_.RespondsToSelector(objc.GetSelector("animationShouldStart:"))
}

func (a_ AnimationDelegateWrapper) AnimationShouldStart(animation IAnimation) bool {
	rv := ffi.CallMethod[bool](a_, "animationShouldStart:", animation)
	return rv
}

func (a_ *AnimationDelegateWrapper) ImplementsAnimation_ValueForProgress() bool {
	return a_.RespondsToSelector(objc.GetSelector("animation:valueForProgress:"))
}

func (a_ AnimationDelegateWrapper) Animation_ValueForProgress(animation IAnimation, progress AnimationProgress) float32 {
	rv := ffi.CallMethod[float32](a_, "animation:valueForProgress:", animation, progress)
	return rv
}

func (a_ *AnimationDelegateWrapper) ImplementsAnimation_DidReachProgressMark() bool {
	return a_.RespondsToSelector(objc.GetSelector("animation:didReachProgressMark:"))
}

func (a_ AnimationDelegateWrapper) Animation_DidReachProgressMark(animation IAnimation, progress AnimationProgress) {
	ffi.CallMethod[ffi.Void](a_, "animation:didReachProgressMark:", animation, progress)
}

var AnimationContextClass = _AnimationContextClass{objc.GetClass("NSAnimationContext")}

type _AnimationContextClass struct {
	objc.Class
}

type IAnimationContext interface {
	objc.IObject
	CompletionHandler() func()
	SetCompletionHandler(value func())
	Duration() foundation.TimeInterval
	SetDuration(value foundation.TimeInterval)
	AllowsImplicitAnimation() bool
	SetAllowsImplicitAnimation(value bool)
}

type AnimationContext struct {
	objc.Object
}

func MakeAnimationContext(ptr unsafe.Pointer) AnimationContext {
	return AnimationContext{
		Object: objc.MakeObject(ptr),
	}
}

func (ac _AnimationContextClass) Alloc() AnimationContext {
	rv := ffi.CallMethod[AnimationContext](ac, "alloc")
	return rv
}

func (a_ AnimationContext) Init() AnimationContext {
	rv := ffi.CallMethod[AnimationContext](a_, "init")
	rv.Autorelease()
	return rv
}

func (ac _AnimationContextClass) New() AnimationContext {
	rv := ffi.CallMethod[AnimationContext](ac, "new")
	rv.Autorelease()
	return rv
}

func NewAnimationContext() AnimationContext {
	return AnimationContextClass.New()
}

func (ac _AnimationContextClass) BeginGrouping() {
	ffi.CallMethod[ffi.Void](ac, "beginGrouping")
}

func (ac _AnimationContextClass) EndGrouping() {
	ffi.CallMethod[ffi.Void](ac, "endGrouping")
}

func (ac _AnimationContextClass) RunAnimationGroup_CompletionHandler(changes func(context AnimationContext), completionHandler func()) {
	ffi.CallMethod[ffi.Void](ac, "runAnimationGroup:completionHandler:", changes, completionHandler)
}

func (ac _AnimationContextClass) RunAnimationGroup(changes func(context AnimationContext)) {
	ffi.CallMethod[ffi.Void](ac, "runAnimationGroup:", changes)
}

func (ac _AnimationContextClass) CurrentContext() AnimationContext {
	rv := ffi.CallMethod[AnimationContext](ac, "currentContext")
	return rv
}

func (a_ AnimationContext) CompletionHandler() func() {
	rv := ffi.CallMethod[func()](a_, "completionHandler")
	return rv
}

func (a_ AnimationContext) SetCompletionHandler(value func()) {
	ffi.CallMethod[ffi.Void](a_, "setCompletionHandler:", value)
}

func (a_ AnimationContext) Duration() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](a_, "duration")
	return rv
}

func (a_ AnimationContext) SetDuration(value foundation.TimeInterval) {
	ffi.CallMethod[ffi.Void](a_, "setDuration:", value)
}

func (a_ AnimationContext) AllowsImplicitAnimation() bool {
	rv := ffi.CallMethod[bool](a_, "allowsImplicitAnimation")
	return rv
}

func (a_ AnimationContext) SetAllowsImplicitAnimation(value bool) {
	ffi.CallMethod[ffi.Void](a_, "setAllowsImplicitAnimation:", value)
}
