// auto-generated code, do not modify
package quartzcore

import (
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

func WrapAnimationDelegate(v AnimationDelegate) objc.Object {
	return objc.WrapAsProtocol("CAAnimationDelegate", v)
}

type AnimationDelegateBase struct {
}

func (p *AnimationDelegateBase) ImplementsAnimationDidStart() bool {
	return false
}

func (p *AnimationDelegateBase) AnimationDidStart(anim Animation) {
	panic("unimpemented protocol method")
}

func (p *AnimationDelegateBase) ImplementsAnimationDidStop_Finished() bool {
	return false
}

func (p *AnimationDelegateBase) AnimationDidStop_Finished(anim Animation, flag bool) {
	panic("unimpemented protocol method")
}

type AnimationDelegateWrapper struct {
	objc.Object
}

func (a_ AnimationDelegateWrapper) AnimationDidStart(anim IAnimation) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("animationDidStart:"), objc.ExtractPtr(anim))
}

func (a_ AnimationDelegateWrapper) AnimationDidStop_Finished(anim IAnimation, flag bool) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("animationDidStop:finished:"), objc.ExtractPtr(anim), flag)
}
