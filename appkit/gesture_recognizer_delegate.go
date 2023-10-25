// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type GestureRecognizerDelegate interface {
	ImplementsGestureRecognizer_ShouldAttemptToRecognizeWithEvent() bool
	// optional
	GestureRecognizer_ShouldAttemptToRecognizeWithEvent(gestureRecognizer GestureRecognizer, event Event) bool
	ImplementsGestureRecognizerShouldBegin() bool
	// optional
	GestureRecognizerShouldBegin(gestureRecognizer GestureRecognizer) bool
	ImplementsGestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer() bool
	// optional
	GestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	ImplementsGestureRecognizer_ShouldRequireFailureOfGestureRecognizer() bool
	// optional
	GestureRecognizer_ShouldRequireFailureOfGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	ImplementsGestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer() bool
	// optional
	GestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	ImplementsGestureRecognizer_ShouldReceiveTouch() bool
	// optional
	GestureRecognizer_ShouldReceiveTouch(gestureRecognizer GestureRecognizer, touch Touch) bool
}

func WrapGestureRecognizerDelegate(v GestureRecognizerDelegate) objc.Object {
	return objc.WrapAsProtocol("NSGestureRecognizerDelegate", v)
}

type GestureRecognizerDelegateBase struct {
}

func (p *GestureRecognizerDelegateBase) ImplementsGestureRecognizer_ShouldAttemptToRecognizeWithEvent() bool {
	return false
}

func (p *GestureRecognizerDelegateBase) GestureRecognizer_ShouldAttemptToRecognizeWithEvent(gestureRecognizer GestureRecognizer, event Event) bool {
	panic("unimpemented protocol method")
}

func (p *GestureRecognizerDelegateBase) ImplementsGestureRecognizerShouldBegin() bool {
	return false
}

func (p *GestureRecognizerDelegateBase) GestureRecognizerShouldBegin(gestureRecognizer GestureRecognizer) bool {
	panic("unimpemented protocol method")
}

func (p *GestureRecognizerDelegateBase) ImplementsGestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer() bool {
	return false
}

func (p *GestureRecognizerDelegateBase) GestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool {
	panic("unimpemented protocol method")
}

func (p *GestureRecognizerDelegateBase) ImplementsGestureRecognizer_ShouldRequireFailureOfGestureRecognizer() bool {
	return false
}

func (p *GestureRecognizerDelegateBase) GestureRecognizer_ShouldRequireFailureOfGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool {
	panic("unimpemented protocol method")
}

func (p *GestureRecognizerDelegateBase) ImplementsGestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer() bool {
	return false
}

func (p *GestureRecognizerDelegateBase) GestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool {
	panic("unimpemented protocol method")
}

func (p *GestureRecognizerDelegateBase) ImplementsGestureRecognizer_ShouldReceiveTouch() bool {
	return false
}

func (p *GestureRecognizerDelegateBase) GestureRecognizer_ShouldReceiveTouch(gestureRecognizer GestureRecognizer, touch Touch) bool {
	panic("unimpemented protocol method")
}

type GestureRecognizerDelegateCreator struct {
	className string
	class     objc.Class
}

func NewGestureRecognizerDelegateCreator(name string) *GestureRecognizerDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("ProtocolBase"), name, 0)
	objc.RegisterClassPair(class)
	return &GestureRecognizerDelegateCreator{className: name, class: class}
}

func (c *GestureRecognizerDelegateCreator) SetGestureRecognizer_ShouldAttemptToRecognizeWithEvent(handle func(o objc.ProtocolBase, gestureRecognizer GestureRecognizer, event Event) bool) {
	objc.AddMethod(c.class, objc.SEL("gestureRecognizer:shouldAttemptToRecognizeWithEvent:"), handle)
}

func (c *GestureRecognizerDelegateCreator) SetGestureRecognizerShouldBegin(handle func(o objc.ProtocolBase, gestureRecognizer GestureRecognizer) bool) {
	objc.AddMethod(c.class, objc.SEL("gestureRecognizerShouldBegin:"), handle)
}

func (c *GestureRecognizerDelegateCreator) SetGestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer(handle func(o objc.ProtocolBase, gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool) {
	objc.AddMethod(c.class, objc.SEL("gestureRecognizer:shouldRecognizeSimultaneouslyWithGestureRecognizer:"), handle)
}

func (c *GestureRecognizerDelegateCreator) SetGestureRecognizer_ShouldRequireFailureOfGestureRecognizer(handle func(o objc.ProtocolBase, gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool) {
	objc.AddMethod(c.class, objc.SEL("gestureRecognizer:shouldRequireFailureOfGestureRecognizer:"), handle)
}

func (c *GestureRecognizerDelegateCreator) SetGestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer(handle func(o objc.ProtocolBase, gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool) {
	objc.AddMethod(c.class, objc.SEL("gestureRecognizer:shouldBeRequiredToFailByGestureRecognizer:"), handle)
}

func (c *GestureRecognizerDelegateCreator) SetGestureRecognizer_ShouldReceiveTouch(handle func(o objc.ProtocolBase, gestureRecognizer GestureRecognizer, touch Touch) bool) {
	objc.AddMethod(c.class, objc.SEL("gestureRecognizer:shouldReceiveTouch:"), handle)
}

func (c *GestureRecognizerDelegateCreator) Create() objc.ProtocolBase {
	return objc.ProtocolBase{Object: c.class.CreateInstance(0)}
}
