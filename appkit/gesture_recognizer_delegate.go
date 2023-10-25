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
