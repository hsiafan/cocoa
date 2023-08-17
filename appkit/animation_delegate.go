// auto-generated code, do not modify
package appkit

import (
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

func WrapAnimationDelegate(v AnimationDelegate) objc.Object {
	return objc.WrapAsProtocol("NSAnimationDelegate", v)
}

type AnimationDelegateBase struct {
}

func (p *AnimationDelegateBase) ImplementsAnimationDidEnd() bool {
	return false
}

func (p *AnimationDelegateBase) AnimationDidEnd(animation Animation) {
	panic("unimpemented protocol method")
}

func (p *AnimationDelegateBase) ImplementsAnimationDidStop() bool {
	return false
}

func (p *AnimationDelegateBase) AnimationDidStop(animation Animation) {
	panic("unimpemented protocol method")
}

func (p *AnimationDelegateBase) ImplementsAnimationShouldStart() bool {
	return false
}

func (p *AnimationDelegateBase) AnimationShouldStart(animation Animation) bool {
	panic("unimpemented protocol method")
}

func (p *AnimationDelegateBase) ImplementsAnimation_ValueForProgress() bool {
	return false
}

func (p *AnimationDelegateBase) Animation_ValueForProgress(animation Animation, progress AnimationProgress) float32 {
	panic("unimpemented protocol method")
}

func (p *AnimationDelegateBase) ImplementsAnimation_DidReachProgressMark() bool {
	return false
}

func (p *AnimationDelegateBase) Animation_DidReachProgressMark(animation Animation, progress AnimationProgress) {
	panic("unimpemented protocol method")
}

type AnimationDelegateWrapper struct {
	objc.Object
}

func (a_ AnimationDelegateWrapper) AnimationDidEnd(animation IAnimation) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("animationDidEnd:"), objc.ExtractPtr(animation))
}

func (a_ AnimationDelegateWrapper) AnimationDidStop(animation IAnimation) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("animationDidStop:"), objc.ExtractPtr(animation))
}

func (a_ AnimationDelegateWrapper) AnimationShouldStart(animation IAnimation) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("animationShouldStart:"), objc.ExtractPtr(animation))
	return rv
}

func (a_ AnimationDelegateWrapper) Animation_ValueForProgress(animation IAnimation, progress AnimationProgress) float32 {
	rv := objc.CallMethod[float32](a_, objc.GetSelector("animation:valueForProgress:"), objc.ExtractPtr(animation), progress)
	return rv
}

func (a_ AnimationDelegateWrapper) Animation_DidReachProgressMark(animation IAnimation, progress AnimationProgress) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("animation:didReachProgressMark:"), objc.ExtractPtr(animation), progress)
}
