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

type GestureRecognizerDelegateImpl struct {
	_GestureRecognizer_ShouldAttemptToRecognizeWithEvent                  func(gestureRecognizer GestureRecognizer, event Event) bool
	_GestureRecognizerShouldBegin                                         func(gestureRecognizer GestureRecognizer) bool
	_GestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	_GestureRecognizer_ShouldRequireFailureOfGestureRecognizer            func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	_GestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer          func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	_GestureRecognizer_ShouldReceiveTouch                                 func(gestureRecognizer GestureRecognizer, touch Touch) bool
}

func (di *GestureRecognizerDelegateImpl) ImplementsGestureRecognizer_ShouldAttemptToRecognizeWithEvent() bool {
	return di._GestureRecognizer_ShouldAttemptToRecognizeWithEvent != nil
}

func (di *GestureRecognizerDelegateImpl) SetGestureRecognizer_ShouldAttemptToRecognizeWithEvent(f func(gestureRecognizer GestureRecognizer, event Event) bool) {
	di._GestureRecognizer_ShouldAttemptToRecognizeWithEvent = f
}

func (di *GestureRecognizerDelegateImpl) GestureRecognizer_ShouldAttemptToRecognizeWithEvent(gestureRecognizer GestureRecognizer, event Event) bool {
	return di._GestureRecognizer_ShouldAttemptToRecognizeWithEvent(gestureRecognizer, event)
}
func (di *GestureRecognizerDelegateImpl) ImplementsGestureRecognizerShouldBegin() bool {
	return di._GestureRecognizerShouldBegin != nil
}

func (di *GestureRecognizerDelegateImpl) SetGestureRecognizerShouldBegin(f func(gestureRecognizer GestureRecognizer) bool) {
	di._GestureRecognizerShouldBegin = f
}

func (di *GestureRecognizerDelegateImpl) GestureRecognizerShouldBegin(gestureRecognizer GestureRecognizer) bool {
	return di._GestureRecognizerShouldBegin(gestureRecognizer)
}
func (di *GestureRecognizerDelegateImpl) ImplementsGestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer() bool {
	return di._GestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer != nil
}

func (di *GestureRecognizerDelegateImpl) SetGestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer(f func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool) {
	di._GestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer = f
}

func (di *GestureRecognizerDelegateImpl) GestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool {
	return di._GestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer(gestureRecognizer, otherGestureRecognizer)
}
func (di *GestureRecognizerDelegateImpl) ImplementsGestureRecognizer_ShouldRequireFailureOfGestureRecognizer() bool {
	return di._GestureRecognizer_ShouldRequireFailureOfGestureRecognizer != nil
}

func (di *GestureRecognizerDelegateImpl) SetGestureRecognizer_ShouldRequireFailureOfGestureRecognizer(f func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool) {
	di._GestureRecognizer_ShouldRequireFailureOfGestureRecognizer = f
}

func (di *GestureRecognizerDelegateImpl) GestureRecognizer_ShouldRequireFailureOfGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool {
	return di._GestureRecognizer_ShouldRequireFailureOfGestureRecognizer(gestureRecognizer, otherGestureRecognizer)
}
func (di *GestureRecognizerDelegateImpl) ImplementsGestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer() bool {
	return di._GestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer != nil
}

func (di *GestureRecognizerDelegateImpl) SetGestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer(f func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool) {
	di._GestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer = f
}

func (di *GestureRecognizerDelegateImpl) GestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool {
	return di._GestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer(gestureRecognizer, otherGestureRecognizer)
}
func (di *GestureRecognizerDelegateImpl) ImplementsGestureRecognizer_ShouldReceiveTouch() bool {
	return di._GestureRecognizer_ShouldReceiveTouch != nil
}

func (di *GestureRecognizerDelegateImpl) SetGestureRecognizer_ShouldReceiveTouch(f func(gestureRecognizer GestureRecognizer, touch Touch) bool) {
	di._GestureRecognizer_ShouldReceiveTouch = f
}

func (di *GestureRecognizerDelegateImpl) GestureRecognizer_ShouldReceiveTouch(gestureRecognizer GestureRecognizer, touch Touch) bool {
	return di._GestureRecognizer_ShouldReceiveTouch(gestureRecognizer, touch)
}

type GestureRecognizerDelegateWrapper struct {
	objc.Object
}

func (g_ *GestureRecognizerDelegateWrapper) ImplementsGestureRecognizer_ShouldAttemptToRecognizeWithEvent() bool {
	return g_.RespondsToSelector(objc.GetSelector("gestureRecognizer:shouldAttemptToRecognizeWithEvent:"))
}

func (g_ GestureRecognizerDelegateWrapper) GestureRecognizer_ShouldAttemptToRecognizeWithEvent(gestureRecognizer IGestureRecognizer, event IEvent) bool {
	rv := objc.CallMethod[bool](g_, "gestureRecognizer:shouldAttemptToRecognizeWithEvent:", gestureRecognizer, event)
	return rv
}

func (g_ *GestureRecognizerDelegateWrapper) ImplementsGestureRecognizerShouldBegin() bool {
	return g_.RespondsToSelector(objc.GetSelector("gestureRecognizerShouldBegin:"))
}

func (g_ GestureRecognizerDelegateWrapper) GestureRecognizerShouldBegin(gestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, "gestureRecognizerShouldBegin:", gestureRecognizer)
	return rv
}

func (g_ *GestureRecognizerDelegateWrapper) ImplementsGestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer() bool {
	return g_.RespondsToSelector(objc.GetSelector("gestureRecognizer:shouldRecognizeSimultaneouslyWithGestureRecognizer:"))
}

func (g_ GestureRecognizerDelegateWrapper) GestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer(gestureRecognizer IGestureRecognizer, otherGestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, "gestureRecognizer:shouldRecognizeSimultaneouslyWithGestureRecognizer:", gestureRecognizer, otherGestureRecognizer)
	return rv
}

func (g_ *GestureRecognizerDelegateWrapper) ImplementsGestureRecognizer_ShouldRequireFailureOfGestureRecognizer() bool {
	return g_.RespondsToSelector(objc.GetSelector("gestureRecognizer:shouldRequireFailureOfGestureRecognizer:"))
}

func (g_ GestureRecognizerDelegateWrapper) GestureRecognizer_ShouldRequireFailureOfGestureRecognizer(gestureRecognizer IGestureRecognizer, otherGestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, "gestureRecognizer:shouldRequireFailureOfGestureRecognizer:", gestureRecognizer, otherGestureRecognizer)
	return rv
}

func (g_ *GestureRecognizerDelegateWrapper) ImplementsGestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer() bool {
	return g_.RespondsToSelector(objc.GetSelector("gestureRecognizer:shouldBeRequiredToFailByGestureRecognizer:"))
}

func (g_ GestureRecognizerDelegateWrapper) GestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer(gestureRecognizer IGestureRecognizer, otherGestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, "gestureRecognizer:shouldBeRequiredToFailByGestureRecognizer:", gestureRecognizer, otherGestureRecognizer)
	return rv
}

func (g_ *GestureRecognizerDelegateWrapper) ImplementsGestureRecognizer_ShouldReceiveTouch() bool {
	return g_.RespondsToSelector(objc.GetSelector("gestureRecognizer:shouldReceiveTouch:"))
}

func (g_ GestureRecognizerDelegateWrapper) GestureRecognizer_ShouldReceiveTouch(gestureRecognizer IGestureRecognizer, touch ITouch) bool {
	rv := objc.CallMethod[bool](g_, "gestureRecognizer:shouldReceiveTouch:", gestureRecognizer, touch)
	return rv
}
