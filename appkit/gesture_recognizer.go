// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var GestureRecognizerClass = _GestureRecognizerClass{objc.GetClass("NSGestureRecognizer")}

type _GestureRecognizerClass struct {
	objc.Class
}

type IGestureRecognizer interface {
	objc.IObject
	LocationInView(view IView) foundation.Point
	Reset()
	MouseDown(event IEvent)
	MouseDragged(event IEvent)
	MouseUp(event IEvent)
	OtherMouseDown(event IEvent)
	OtherMouseDragged(event IEvent)
	OtherMouseUp(event IEvent)
	RightMouseDown(event IEvent)
	RightMouseDragged(event IEvent)
	RightMouseUp(event IEvent)
	MagnifyWithEvent(event IEvent)
	RotateWithEvent(event IEvent)
	CanBePreventedByGestureRecognizer(preventingGestureRecognizer IGestureRecognizer) bool
	CanPreventGestureRecognizer(preventedGestureRecognizer IGestureRecognizer) bool
	ShouldBeRequiredToFailByGestureRecognizer(otherGestureRecognizer IGestureRecognizer) bool
	ShouldRequireFailureOfGestureRecognizer(otherGestureRecognizer IGestureRecognizer) bool
	KeyDown(event IEvent)
	KeyUp(event IEvent)
	TabletPoint(event IEvent)
	FlagsChanged(event IEvent)
	PressureChangeWithEvent(event IEvent)
	TouchesBeganWithEvent(event IEvent)
	TouchesCancelledWithEvent(event IEvent)
	TouchesEndedWithEvent(event IEvent)
	TouchesMovedWithEvent(event IEvent)
	Action() objc.Selector
	SetAction(value objc.Selector)
	Target() objc.Object
	SetTarget(value objc.IObject)
	State() GestureRecognizerState
	View() View
	IsEnabled() bool
	SetEnabled(value bool)
	DelaysPrimaryMouseButtonEvents() bool
	SetDelaysPrimaryMouseButtonEvents(value bool)
	DelaysSecondaryMouseButtonEvents() bool
	SetDelaysSecondaryMouseButtonEvents(value bool)
	DelaysOtherMouseButtonEvents() bool
	SetDelaysOtherMouseButtonEvents(value bool)
	DelaysKeyEvents() bool
	SetDelaysKeyEvents(value bool)
	DelaysMagnificationEvents() bool
	SetDelaysMagnificationEvents(value bool)
	DelaysRotationEvents() bool
	SetDelaysRotationEvents(value bool)
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
	PressureConfiguration() PressureConfiguration
	SetPressureConfiguration(value IPressureConfiguration)
	AllowedTouchTypes() TouchTypeMask
	SetAllowedTouchTypes(value TouchTypeMask)
}

type GestureRecognizer struct {
	objc.Object
}

func MakeGestureRecognizer(ptr unsafe.Pointer) GestureRecognizer {
	return GestureRecognizer{
		Object: objc.MakeObject(ptr),
	}
}

func (g_ GestureRecognizer) InitWithTarget_Action(target objc.IObject, action objc.Selector) GestureRecognizer {
	rv := objc.CallMethod[GestureRecognizer](g_, objc.SEL("initWithTarget:action:"), objc.ExtractPtr(target), action)
	return rv
}

func (gc _GestureRecognizerClass) Alloc() GestureRecognizer {
	rv := objc.CallMethod[GestureRecognizer](gc, objc.SEL("alloc"))
	return rv
}

func (gc _GestureRecognizerClass) New() GestureRecognizer {
	rv := objc.CallMethod[GestureRecognizer](gc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewGestureRecognizer() GestureRecognizer {
	return GestureRecognizerClass.New()
}

func (g_ GestureRecognizer) Init() GestureRecognizer {
	rv := objc.CallMethod[GestureRecognizer](g_, objc.SEL("init"))
	return rv
}

func (g_ GestureRecognizer) LocationInView(view IView) foundation.Point {
	rv := objc.CallMethod[foundation.Point](g_, objc.SEL("locationInView:"), objc.ExtractPtr(view))
	return rv
}

func (g_ GestureRecognizer) Reset() {
	objc.CallMethod[objc.Void](g_, objc.SEL("reset"))
}

func (g_ GestureRecognizer) MouseDown(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("mouseDown:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) MouseDragged(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("mouseDragged:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) MouseUp(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("mouseUp:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) OtherMouseDown(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("otherMouseDown:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) OtherMouseDragged(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("otherMouseDragged:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) OtherMouseUp(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("otherMouseUp:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) RightMouseDown(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("rightMouseDown:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) RightMouseDragged(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("rightMouseDragged:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) RightMouseUp(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("rightMouseUp:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) MagnifyWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("magnifyWithEvent:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) RotateWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("rotateWithEvent:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) CanBePreventedByGestureRecognizer(preventingGestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, objc.SEL("canBePreventedByGestureRecognizer:"), objc.ExtractPtr(preventingGestureRecognizer))
	return rv
}

func (g_ GestureRecognizer) CanPreventGestureRecognizer(preventedGestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, objc.SEL("canPreventGestureRecognizer:"), objc.ExtractPtr(preventedGestureRecognizer))
	return rv
}

func (g_ GestureRecognizer) ShouldBeRequiredToFailByGestureRecognizer(otherGestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, objc.SEL("shouldBeRequiredToFailByGestureRecognizer:"), objc.ExtractPtr(otherGestureRecognizer))
	return rv
}

func (g_ GestureRecognizer) ShouldRequireFailureOfGestureRecognizer(otherGestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, objc.SEL("shouldRequireFailureOfGestureRecognizer:"), objc.ExtractPtr(otherGestureRecognizer))
	return rv
}

func (g_ GestureRecognizer) KeyDown(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("keyDown:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) KeyUp(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("keyUp:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) TabletPoint(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("tabletPoint:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) FlagsChanged(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("flagsChanged:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) PressureChangeWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("pressureChangeWithEvent:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) TouchesBeganWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("touchesBeganWithEvent:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) TouchesCancelledWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("touchesCancelledWithEvent:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) TouchesEndedWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("touchesEndedWithEvent:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) TouchesMovedWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("touchesMovedWithEvent:"), objc.ExtractPtr(event))
}

func (g_ GestureRecognizer) Action() objc.Selector {
	rv := objc.CallMethod[objc.Selector](g_, objc.SEL("action"))
	return rv
}

func (g_ GestureRecognizer) SetAction(value objc.Selector) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setAction:"), value)
}

// weak property
func (g_ GestureRecognizer) Target() objc.Object {
	rv := objc.CallMethod[objc.Object](g_, objc.SEL("target"))
	return rv
}

// weak property
func (g_ GestureRecognizer) SetTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setTarget:"), objc.ExtractPtr(value))
}

func (g_ GestureRecognizer) State() GestureRecognizerState {
	rv := objc.CallMethod[GestureRecognizerState](g_, objc.SEL("state"))
	return rv
}

func (g_ GestureRecognizer) View() View {
	rv := objc.CallMethod[View](g_, objc.SEL("view"))
	return rv
}

func (g_ GestureRecognizer) IsEnabled() bool {
	rv := objc.CallMethod[bool](g_, objc.SEL("isEnabled"))
	return rv
}

func (g_ GestureRecognizer) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setEnabled:"), value)
}

func (g_ GestureRecognizer) DelaysPrimaryMouseButtonEvents() bool {
	rv := objc.CallMethod[bool](g_, objc.SEL("delaysPrimaryMouseButtonEvents"))
	return rv
}

func (g_ GestureRecognizer) SetDelaysPrimaryMouseButtonEvents(value bool) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setDelaysPrimaryMouseButtonEvents:"), value)
}

func (g_ GestureRecognizer) DelaysSecondaryMouseButtonEvents() bool {
	rv := objc.CallMethod[bool](g_, objc.SEL("delaysSecondaryMouseButtonEvents"))
	return rv
}

func (g_ GestureRecognizer) SetDelaysSecondaryMouseButtonEvents(value bool) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setDelaysSecondaryMouseButtonEvents:"), value)
}

func (g_ GestureRecognizer) DelaysOtherMouseButtonEvents() bool {
	rv := objc.CallMethod[bool](g_, objc.SEL("delaysOtherMouseButtonEvents"))
	return rv
}

func (g_ GestureRecognizer) SetDelaysOtherMouseButtonEvents(value bool) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setDelaysOtherMouseButtonEvents:"), value)
}

func (g_ GestureRecognizer) DelaysKeyEvents() bool {
	rv := objc.CallMethod[bool](g_, objc.SEL("delaysKeyEvents"))
	return rv
}

func (g_ GestureRecognizer) SetDelaysKeyEvents(value bool) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setDelaysKeyEvents:"), value)
}

func (g_ GestureRecognizer) DelaysMagnificationEvents() bool {
	rv := objc.CallMethod[bool](g_, objc.SEL("delaysMagnificationEvents"))
	return rv
}

func (g_ GestureRecognizer) SetDelaysMagnificationEvents(value bool) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setDelaysMagnificationEvents:"), value)
}

func (g_ GestureRecognizer) DelaysRotationEvents() bool {
	rv := objc.CallMethod[bool](g_, objc.SEL("delaysRotationEvents"))
	return rv
}

func (g_ GestureRecognizer) SetDelaysRotationEvents(value bool) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setDelaysRotationEvents:"), value)
}

// weak property
func (g_ GestureRecognizer) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](g_, objc.SEL("delegate"))
	return rv
}

// weak property
func (g_ GestureRecognizer) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}

func (g_ GestureRecognizer) PressureConfiguration() PressureConfiguration {
	rv := objc.CallMethod[PressureConfiguration](g_, objc.SEL("pressureConfiguration"))
	return rv
}

func (g_ GestureRecognizer) SetPressureConfiguration(value IPressureConfiguration) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setPressureConfiguration:"), objc.ExtractPtr(value))
}

func (g_ GestureRecognizer) AllowedTouchTypes() TouchTypeMask {
	rv := objc.CallMethod[TouchTypeMask](g_, objc.SEL("allowedTouchTypes"))
	return rv
}

func (g_ GestureRecognizer) SetAllowedTouchTypes(value TouchTypeMask) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setAllowedTouchTypes:"), value)
}
