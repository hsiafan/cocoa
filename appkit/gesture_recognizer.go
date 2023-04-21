// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
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
	Delegate() GestureRecognizerDelegateWrapper
	SetDelegate(value GestureRecognizerDelegate)
	SetDelegate0(value objc.IObject)
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
	rv := objc.CallMethod[GestureRecognizer](g_, "initWithTarget:action:", target, action)
	return rv
}

func (gc _GestureRecognizerClass) Alloc() GestureRecognizer {
	rv := objc.CallMethod[GestureRecognizer](gc, "alloc")
	return rv
}

func (gc _GestureRecognizerClass) New() GestureRecognizer {
	rv := objc.CallMethod[GestureRecognizer](gc, "new")
	rv.Autorelease()
	return rv
}

func NewGestureRecognizer() GestureRecognizer {
	return GestureRecognizerClass.New()
}

func (g_ GestureRecognizer) Init() GestureRecognizer {
	rv := objc.CallMethod[GestureRecognizer](g_, "init")
	return rv
}

func (g_ GestureRecognizer) LocationInView(view IView) foundation.Point {
	rv := objc.CallMethod[foundation.Point](g_, "locationInView:", view)
	return rv
}

func (g_ GestureRecognizer) Reset() {
	objc.CallMethod[objc.Void](g_, "reset")
}

func (g_ GestureRecognizer) MouseDown(event IEvent) {
	objc.CallMethod[objc.Void](g_, "mouseDown:", event)
}

func (g_ GestureRecognizer) MouseDragged(event IEvent) {
	objc.CallMethod[objc.Void](g_, "mouseDragged:", event)
}

func (g_ GestureRecognizer) MouseUp(event IEvent) {
	objc.CallMethod[objc.Void](g_, "mouseUp:", event)
}

func (g_ GestureRecognizer) OtherMouseDown(event IEvent) {
	objc.CallMethod[objc.Void](g_, "otherMouseDown:", event)
}

func (g_ GestureRecognizer) OtherMouseDragged(event IEvent) {
	objc.CallMethod[objc.Void](g_, "otherMouseDragged:", event)
}

func (g_ GestureRecognizer) OtherMouseUp(event IEvent) {
	objc.CallMethod[objc.Void](g_, "otherMouseUp:", event)
}

func (g_ GestureRecognizer) RightMouseDown(event IEvent) {
	objc.CallMethod[objc.Void](g_, "rightMouseDown:", event)
}

func (g_ GestureRecognizer) RightMouseDragged(event IEvent) {
	objc.CallMethod[objc.Void](g_, "rightMouseDragged:", event)
}

func (g_ GestureRecognizer) RightMouseUp(event IEvent) {
	objc.CallMethod[objc.Void](g_, "rightMouseUp:", event)
}

func (g_ GestureRecognizer) MagnifyWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, "magnifyWithEvent:", event)
}

func (g_ GestureRecognizer) RotateWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, "rotateWithEvent:", event)
}

func (g_ GestureRecognizer) CanBePreventedByGestureRecognizer(preventingGestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, "canBePreventedByGestureRecognizer:", preventingGestureRecognizer)
	return rv
}

func (g_ GestureRecognizer) CanPreventGestureRecognizer(preventedGestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, "canPreventGestureRecognizer:", preventedGestureRecognizer)
	return rv
}

func (g_ GestureRecognizer) ShouldBeRequiredToFailByGestureRecognizer(otherGestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, "shouldBeRequiredToFailByGestureRecognizer:", otherGestureRecognizer)
	return rv
}

func (g_ GestureRecognizer) ShouldRequireFailureOfGestureRecognizer(otherGestureRecognizer IGestureRecognizer) bool {
	rv := objc.CallMethod[bool](g_, "shouldRequireFailureOfGestureRecognizer:", otherGestureRecognizer)
	return rv
}

func (g_ GestureRecognizer) KeyDown(event IEvent) {
	objc.CallMethod[objc.Void](g_, "keyDown:", event)
}

func (g_ GestureRecognizer) KeyUp(event IEvent) {
	objc.CallMethod[objc.Void](g_, "keyUp:", event)
}

func (g_ GestureRecognizer) TabletPoint(event IEvent) {
	objc.CallMethod[objc.Void](g_, "tabletPoint:", event)
}

func (g_ GestureRecognizer) FlagsChanged(event IEvent) {
	objc.CallMethod[objc.Void](g_, "flagsChanged:", event)
}

func (g_ GestureRecognizer) PressureChangeWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, "pressureChangeWithEvent:", event)
}

func (g_ GestureRecognizer) TouchesBeganWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, "touchesBeganWithEvent:", event)
}

func (g_ GestureRecognizer) TouchesCancelledWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, "touchesCancelledWithEvent:", event)
}

func (g_ GestureRecognizer) TouchesEndedWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, "touchesEndedWithEvent:", event)
}

func (g_ GestureRecognizer) TouchesMovedWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](g_, "touchesMovedWithEvent:", event)
}

func (g_ GestureRecognizer) Action() objc.Selector {
	rv := objc.CallMethod[objc.Selector](g_, "action")
	return rv
}

func (g_ GestureRecognizer) SetAction(value objc.Selector) {
	objc.CallMethod[objc.Void](g_, "setAction:", value)
}

func (g_ GestureRecognizer) Target() objc.Object {
	rv := objc.CallMethod[objc.Object](g_, "target")
	return rv
}

func (g_ GestureRecognizer) SetTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](g_, "setTarget:", value)
}

func (g_ GestureRecognizer) State() GestureRecognizerState {
	rv := objc.CallMethod[GestureRecognizerState](g_, "state")
	return rv
}

func (g_ GestureRecognizer) View() View {
	rv := objc.CallMethod[View](g_, "view")
	return rv
}

func (g_ GestureRecognizer) IsEnabled() bool {
	rv := objc.CallMethod[bool](g_, "isEnabled")
	return rv
}

func (g_ GestureRecognizer) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](g_, "setEnabled:", value)
}

func (g_ GestureRecognizer) DelaysPrimaryMouseButtonEvents() bool {
	rv := objc.CallMethod[bool](g_, "delaysPrimaryMouseButtonEvents")
	return rv
}

func (g_ GestureRecognizer) SetDelaysPrimaryMouseButtonEvents(value bool) {
	objc.CallMethod[objc.Void](g_, "setDelaysPrimaryMouseButtonEvents:", value)
}

func (g_ GestureRecognizer) DelaysSecondaryMouseButtonEvents() bool {
	rv := objc.CallMethod[bool](g_, "delaysSecondaryMouseButtonEvents")
	return rv
}

func (g_ GestureRecognizer) SetDelaysSecondaryMouseButtonEvents(value bool) {
	objc.CallMethod[objc.Void](g_, "setDelaysSecondaryMouseButtonEvents:", value)
}

func (g_ GestureRecognizer) DelaysOtherMouseButtonEvents() bool {
	rv := objc.CallMethod[bool](g_, "delaysOtherMouseButtonEvents")
	return rv
}

func (g_ GestureRecognizer) SetDelaysOtherMouseButtonEvents(value bool) {
	objc.CallMethod[objc.Void](g_, "setDelaysOtherMouseButtonEvents:", value)
}

func (g_ GestureRecognizer) DelaysKeyEvents() bool {
	rv := objc.CallMethod[bool](g_, "delaysKeyEvents")
	return rv
}

func (g_ GestureRecognizer) SetDelaysKeyEvents(value bool) {
	objc.CallMethod[objc.Void](g_, "setDelaysKeyEvents:", value)
}

func (g_ GestureRecognizer) DelaysMagnificationEvents() bool {
	rv := objc.CallMethod[bool](g_, "delaysMagnificationEvents")
	return rv
}

func (g_ GestureRecognizer) SetDelaysMagnificationEvents(value bool) {
	objc.CallMethod[objc.Void](g_, "setDelaysMagnificationEvents:", value)
}

func (g_ GestureRecognizer) DelaysRotationEvents() bool {
	rv := objc.CallMethod[bool](g_, "delaysRotationEvents")
	return rv
}

func (g_ GestureRecognizer) SetDelaysRotationEvents(value bool) {
	objc.CallMethod[objc.Void](g_, "setDelaysRotationEvents:", value)
}

func (g_ GestureRecognizer) Delegate() GestureRecognizerDelegateWrapper {
	rv := objc.CallMethod[GestureRecognizerDelegateWrapper](g_, "delegate")
	return rv
}

func (g_ GestureRecognizer) SetDelegate(value GestureRecognizerDelegate) {
	po := objc.CreateProtocol("NSGestureRecognizerDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(g_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](g_, "setDelegate:", po)
}

func (g_ GestureRecognizer) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](g_, "setDelegate:", value)
}

func (g_ GestureRecognizer) PressureConfiguration() PressureConfiguration {
	rv := objc.CallMethod[PressureConfiguration](g_, "pressureConfiguration")
	return rv
}

func (g_ GestureRecognizer) SetPressureConfiguration(value IPressureConfiguration) {
	objc.CallMethod[objc.Void](g_, "setPressureConfiguration:", value)
}

func (g_ GestureRecognizer) AllowedTouchTypes() TouchTypeMask {
	rv := objc.CallMethod[TouchTypeMask](g_, "allowedTouchTypes")
	return rv
}

func (g_ GestureRecognizer) SetAllowedTouchTypes(value TouchTypeMask) {
	objc.CallMethod[objc.Void](g_, "setAllowedTouchTypes:", value)
}
