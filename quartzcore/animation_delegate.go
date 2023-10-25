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

type AnimationDelegateCreator struct {
	className string
	class     objc.Class
}

func NewAnimationDelegateCreator(name string) *AnimationDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("ProtocolBase"), name, 0)
	objc.RegisterClassPair(class)
	return &AnimationDelegateCreator{className: name, class: class}
}

func (c *AnimationDelegateCreator) SetAnimationDidStart(handle func(o objc.ProtocolBase, anim Animation)) {
	objc.AddMethod(c.class, objc.SEL("animationDidStart:"), handle)
}

func (c *AnimationDelegateCreator) SetAnimationDidStop_Finished(handle func(o objc.ProtocolBase, anim Animation, flag bool)) {
	objc.AddMethod(c.class, objc.SEL("animationDidStop:finished:"), handle)
}

func (c *AnimationDelegateCreator) Create() objc.ProtocolBase {
	return objc.ProtocolBase{Object: c.class.CreateInstance(0)}
}
