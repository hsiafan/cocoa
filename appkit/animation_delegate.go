// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

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
