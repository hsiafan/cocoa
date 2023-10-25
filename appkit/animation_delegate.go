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

type AnimationDelegateCreator struct {
	className string
	class     objc.Class
}

func NewAnimationDelegateCreator(name string) *AnimationDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &AnimationDelegateCreator{className: name, class: class}
}

func (c *AnimationDelegateCreator) SetAnimationDidEnd(handle func(o objc.Object, animation Animation)) {
	objc.AddMethod(c.class, objc.GetSelector("animationDidEnd:"), handle)
}

func (c *AnimationDelegateCreator) SetAnimationDidStop(handle func(o objc.Object, animation Animation)) {
	objc.AddMethod(c.class, objc.GetSelector("animationDidStop:"), handle)
}

func (c *AnimationDelegateCreator) SetAnimationShouldStart(handle func(o objc.Object, animation Animation) bool) {
	objc.AddMethod(c.class, objc.GetSelector("animationShouldStart:"), handle)
}

func (c *AnimationDelegateCreator) SetAnimation_ValueForProgress(handle func(o objc.Object, animation Animation, progress AnimationProgress) float32) {
	objc.AddMethod(c.class, objc.GetSelector("animation:valueForProgress:"), handle)
}

func (c *AnimationDelegateCreator) SetAnimation_DidReachProgressMark(handle func(o objc.Object, animation Animation, progress AnimationProgress)) {
	objc.AddMethod(c.class, objc.GetSelector("animation:didReachProgressMark:"), handle)
}

func (c *AnimationDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
