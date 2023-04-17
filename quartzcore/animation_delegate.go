// auto-generated code, do not modify
package quartzcore

import (
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

type AnimationDelegate interface {
	ImplementsAnimationDidStart() bool
	// optional
	AnimationDidStart(anim Animation)
	ImplementsAnimationDidStop_Finished() bool
	// optional
	AnimationDidStop_Finished(anim Animation, flag bool)
}

type AnimationDelegateImpl struct {
	_AnimationDidStart         func(anim Animation)
	_AnimationDidStop_Finished func(anim Animation, flag bool)
}

func (di *AnimationDelegateImpl) ImplementsAnimationDidStart() bool {
	return di._AnimationDidStart != nil
}

func (di *AnimationDelegateImpl) SetAnimationDidStart(f func(anim Animation)) {
	di._AnimationDidStart = f
}

func (di *AnimationDelegateImpl) AnimationDidStart(anim Animation) {
	di._AnimationDidStart(anim)
}
func (di *AnimationDelegateImpl) ImplementsAnimationDidStop_Finished() bool {
	return di._AnimationDidStop_Finished != nil
}

func (di *AnimationDelegateImpl) SetAnimationDidStop_Finished(f func(anim Animation, flag bool)) {
	di._AnimationDidStop_Finished = f
}

func (di *AnimationDelegateImpl) AnimationDidStop_Finished(anim Animation, flag bool) {
	di._AnimationDidStop_Finished(anim, flag)
}

type AnimationDelegateWrapper struct {
	objc.Object
}

func (a_ *AnimationDelegateWrapper) ImplementsAnimationDidStart() bool {
	return a_.RespondsToSelector(objc.GetSelector("animationDidStart:"))
}

func (a_ AnimationDelegateWrapper) AnimationDidStart(anim IAnimation) {
	ffi.CallMethod[ffi.Void](a_, "animationDidStart:", anim)
}

func (a_ *AnimationDelegateWrapper) ImplementsAnimationDidStop_Finished() bool {
	return a_.RespondsToSelector(objc.GetSelector("animationDidStop:finished:"))
}

func (a_ AnimationDelegateWrapper) AnimationDidStop_Finished(anim IAnimation, flag bool) {
	ffi.CallMethod[ffi.Void](a_, "animationDidStop:finished:", anim, flag)
}
