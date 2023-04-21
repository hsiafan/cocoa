// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ResponderClass = _ResponderClass{objc.GetClass("NSResponder")}

type _ResponderClass struct {
	objc.Class
}

type IResponder interface {
	objc.IObject
	BecomeFirstResponder() bool
	ResignFirstResponder() bool
	ValidateProposedFirstResponder_ForEvent(responder IResponder, event IEvent) bool
	MouseDown(event IEvent)
	MouseDragged(event IEvent)
	MouseUp(event IEvent)
	MouseMoved(event IEvent)
	MouseEntered(event IEvent)
	MouseExited(event IEvent)
	RightMouseDown(event IEvent)
	RightMouseDragged(event IEvent)
	RightMouseUp(event IEvent)
	OtherMouseDown(event IEvent)
	OtherMouseDragged(event IEvent)
	OtherMouseUp(event IEvent)
	KeyDown(event IEvent)
	KeyUp(event IEvent)
	InterpretKeyEvents(eventArray []IEvent)
	PerformKeyEquivalent(event IEvent) bool
	// deprecated
	PerformMnemonic(string_ string) bool
	FlushBufferedKeyEvents()
	PressureChangeWithEvent(event IEvent)
	CursorUpdate(event IEvent)
	FlagsChanged(event IEvent)
	TabletPoint(event IEvent)
	TabletProximity(event IEvent)
	HelpRequested(eventPtr IEvent)
	ScrollWheel(event IEvent)
	QuickLookWithEvent(event IEvent)
	ChangeModeWithEvent(event IEvent)
	SupplementalTargetForAction_Sender(action objc.Selector, sender objc.IObject) objc.Object
	EncodeRestorableStateWithCoder(coder foundation.ICoder)
	EncodeRestorableStateWithCoder_BackgroundQueue(coder foundation.ICoder, queue foundation.IOperationQueue)
	RestoreStateWithCoder(coder foundation.ICoder)
	InvalidateRestorableState()
	UpdateUserActivityState(userActivity foundation.IUserActivity)
	PresentError(error foundation.IError) bool
	PresentError_ModalForWindow_Delegate_DidPresentSelector_ContextInfo(error foundation.IError, window IWindow, delegate objc.IObject, didPresentSelector objc.Selector, contextInfo unsafe.Pointer)
	WillPresentError(error foundation.IError) foundation.Error
	TryToPerform_With(action objc.Selector, object objc.IObject) bool
	ValidRequestorForSendType_ReturnType(sendType PasteboardType, returnType PasteboardType) objc.Object
	ShouldBeTreatedAsInkEvent(event IEvent) bool
	NoResponderFor(eventSelector objc.Selector)
	// deprecated
	SetInterfaceStyle(interfaceStyle InterfaceStyle)
	// deprecated
	InterfaceStyle() InterfaceStyle
	BeginGestureWithEvent(event IEvent)
	EndGestureWithEvent(event IEvent)
	MagnifyWithEvent(event IEvent)
	RotateWithEvent(event IEvent)
	SwipeWithEvent(event IEvent)
	TouchesBeganWithEvent(event IEvent)
	TouchesMovedWithEvent(event IEvent)
	TouchesCancelledWithEvent(event IEvent)
	TouchesEndedWithEvent(event IEvent)
	WantsForwardedScrollEventsForAxis(axis EventGestureAxis) bool
	SmartMagnifyWithEvent(event IEvent)
	WantsScrollEventsForSwipeTrackingOnAxis(axis EventGestureAxis) bool
	MakeTouchBar() TouchBar
	PerformTextFinderAction(sender objc.IObject)
	NewWindowForTab(sender objc.IObject)
	ShowContextHelp(sender objc.IObject)
	AcceptsFirstResponder() bool
	NextResponder() Responder
	SetNextResponder(value IResponder)
	UserActivity() foundation.UserActivity
	SetUserActivity(value foundation.IUserActivity)
	Menu() Menu
	SetMenu(value IMenu)
	UndoManager() foundation.UndoManager
	TouchBar() TouchBar
	SetTouchBar(value ITouchBar)
}

type Responder struct {
	objc.Object
}

func MakeResponder(ptr unsafe.Pointer) Responder {
	return Responder{
		Object: objc.MakeObject(ptr),
	}
}

func (r_ Responder) Init() Responder {
	rv := objc.CallMethod[Responder](r_, "init")
	return rv
}

func (rc _ResponderClass) Alloc() Responder {
	rv := objc.CallMethod[Responder](rc, "alloc")
	return rv
}

func (rc _ResponderClass) New() Responder {
	rv := objc.CallMethod[Responder](rc, "new")
	rv.Autorelease()
	return rv
}

func NewResponder() Responder {
	return ResponderClass.New()
}

func (r_ Responder) BecomeFirstResponder() bool {
	rv := objc.CallMethod[bool](r_, "becomeFirstResponder")
	return rv
}

func (r_ Responder) ResignFirstResponder() bool {
	rv := objc.CallMethod[bool](r_, "resignFirstResponder")
	return rv
}

func (r_ Responder) ValidateProposedFirstResponder_ForEvent(responder IResponder, event IEvent) bool {
	rv := objc.CallMethod[bool](r_, "validateProposedFirstResponder:forEvent:", responder, event)
	return rv
}

func (r_ Responder) MouseDown(event IEvent) {
	objc.CallMethod[objc.Void](r_, "mouseDown:", event)
}

func (r_ Responder) MouseDragged(event IEvent) {
	objc.CallMethod[objc.Void](r_, "mouseDragged:", event)
}

func (r_ Responder) MouseUp(event IEvent) {
	objc.CallMethod[objc.Void](r_, "mouseUp:", event)
}

func (r_ Responder) MouseMoved(event IEvent) {
	objc.CallMethod[objc.Void](r_, "mouseMoved:", event)
}

func (r_ Responder) MouseEntered(event IEvent) {
	objc.CallMethod[objc.Void](r_, "mouseEntered:", event)
}

func (r_ Responder) MouseExited(event IEvent) {
	objc.CallMethod[objc.Void](r_, "mouseExited:", event)
}

func (r_ Responder) RightMouseDown(event IEvent) {
	objc.CallMethod[objc.Void](r_, "rightMouseDown:", event)
}

func (r_ Responder) RightMouseDragged(event IEvent) {
	objc.CallMethod[objc.Void](r_, "rightMouseDragged:", event)
}

func (r_ Responder) RightMouseUp(event IEvent) {
	objc.CallMethod[objc.Void](r_, "rightMouseUp:", event)
}

func (r_ Responder) OtherMouseDown(event IEvent) {
	objc.CallMethod[objc.Void](r_, "otherMouseDown:", event)
}

func (r_ Responder) OtherMouseDragged(event IEvent) {
	objc.CallMethod[objc.Void](r_, "otherMouseDragged:", event)
}

func (r_ Responder) OtherMouseUp(event IEvent) {
	objc.CallMethod[objc.Void](r_, "otherMouseUp:", event)
}

func (r_ Responder) KeyDown(event IEvent) {
	objc.CallMethod[objc.Void](r_, "keyDown:", event)
}

func (r_ Responder) KeyUp(event IEvent) {
	objc.CallMethod[objc.Void](r_, "keyUp:", event)
}

func (r_ Responder) InterpretKeyEvents(eventArray []IEvent) {
	objc.CallMethod[objc.Void](r_, "interpretKeyEvents:", eventArray)
}

func (r_ Responder) PerformKeyEquivalent(event IEvent) bool {
	rv := objc.CallMethod[bool](r_, "performKeyEquivalent:", event)
	return rv
}

// deprecated
func (r_ Responder) PerformMnemonic(string_ string) bool {
	rv := objc.CallMethod[bool](r_, "performMnemonic:", string_)
	return rv
}

func (r_ Responder) FlushBufferedKeyEvents() {
	objc.CallMethod[objc.Void](r_, "flushBufferedKeyEvents")
}

func (r_ Responder) PressureChangeWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, "pressureChangeWithEvent:", event)
}

func (r_ Responder) CursorUpdate(event IEvent) {
	objc.CallMethod[objc.Void](r_, "cursorUpdate:", event)
}

func (r_ Responder) FlagsChanged(event IEvent) {
	objc.CallMethod[objc.Void](r_, "flagsChanged:", event)
}

func (r_ Responder) TabletPoint(event IEvent) {
	objc.CallMethod[objc.Void](r_, "tabletPoint:", event)
}

func (r_ Responder) TabletProximity(event IEvent) {
	objc.CallMethod[objc.Void](r_, "tabletProximity:", event)
}

func (r_ Responder) HelpRequested(eventPtr IEvent) {
	objc.CallMethod[objc.Void](r_, "helpRequested:", eventPtr)
}

func (r_ Responder) ScrollWheel(event IEvent) {
	objc.CallMethod[objc.Void](r_, "scrollWheel:", event)
}

func (r_ Responder) QuickLookWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, "quickLookWithEvent:", event)
}

func (r_ Responder) ChangeModeWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, "changeModeWithEvent:", event)
}

func (r_ Responder) SupplementalTargetForAction_Sender(action objc.Selector, sender objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](r_, "supplementalTargetForAction:sender:", action, sender)
	return rv
}

func (rc _ResponderClass) AllowedClassesForRestorableStateKeyPath(keyPath string) []objc.Class {
	rv := objc.CallMethod[[]objc.Class](rc, "allowedClassesForRestorableStateKeyPath:", keyPath)
	return rv
}

func (r_ Responder) EncodeRestorableStateWithCoder(coder foundation.ICoder) {
	objc.CallMethod[objc.Void](r_, "encodeRestorableStateWithCoder:", coder)
}

func (r_ Responder) EncodeRestorableStateWithCoder_BackgroundQueue(coder foundation.ICoder, queue foundation.IOperationQueue) {
	objc.CallMethod[objc.Void](r_, "encodeRestorableStateWithCoder:backgroundQueue:", coder, queue)
}

func (r_ Responder) RestoreStateWithCoder(coder foundation.ICoder) {
	objc.CallMethod[objc.Void](r_, "restoreStateWithCoder:", coder)
}

func (r_ Responder) InvalidateRestorableState() {
	objc.CallMethod[objc.Void](r_, "invalidateRestorableState")
}

func (r_ Responder) UpdateUserActivityState(userActivity foundation.IUserActivity) {
	objc.CallMethod[objc.Void](r_, "updateUserActivityState:", userActivity)
}

func (r_ Responder) PresentError(error foundation.IError) bool {
	rv := objc.CallMethod[bool](r_, "presentError:", error)
	return rv
}

func (r_ Responder) PresentError_ModalForWindow_Delegate_DidPresentSelector_ContextInfo(error foundation.IError, window IWindow, delegate objc.IObject, didPresentSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](r_, "presentError:modalForWindow:delegate:didPresentSelector:contextInfo:", error, window, delegate, didPresentSelector, contextInfo)
}

func (r_ Responder) WillPresentError(error foundation.IError) foundation.Error {
	rv := objc.CallMethod[foundation.Error](r_, "willPresentError:", error)
	return rv
}

func (r_ Responder) TryToPerform_With(action objc.Selector, object objc.IObject) bool {
	rv := objc.CallMethod[bool](r_, "tryToPerform:with:", action, object)
	return rv
}

func (r_ Responder) ValidRequestorForSendType_ReturnType(sendType PasteboardType, returnType PasteboardType) objc.Object {
	rv := objc.CallMethod[objc.Object](r_, "validRequestorForSendType:returnType:", sendType, returnType)
	return rv
}

func (r_ Responder) ShouldBeTreatedAsInkEvent(event IEvent) bool {
	rv := objc.CallMethod[bool](r_, "shouldBeTreatedAsInkEvent:", event)
	return rv
}

func (r_ Responder) NoResponderFor(eventSelector objc.Selector) {
	objc.CallMethod[objc.Void](r_, "noResponderFor:", eventSelector)
}

// deprecated
func (r_ Responder) SetInterfaceStyle(interfaceStyle InterfaceStyle) {
	objc.CallMethod[objc.Void](r_, "setInterfaceStyle:", interfaceStyle)
}

// deprecated
func (r_ Responder) InterfaceStyle() InterfaceStyle {
	rv := objc.CallMethod[InterfaceStyle](r_, "interfaceStyle")
	return rv
}

func (r_ Responder) BeginGestureWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, "beginGestureWithEvent:", event)
}

func (r_ Responder) EndGestureWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, "endGestureWithEvent:", event)
}

func (r_ Responder) MagnifyWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, "magnifyWithEvent:", event)
}

func (r_ Responder) RotateWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, "rotateWithEvent:", event)
}

func (r_ Responder) SwipeWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, "swipeWithEvent:", event)
}

func (r_ Responder) TouchesBeganWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, "touchesBeganWithEvent:", event)
}

func (r_ Responder) TouchesMovedWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, "touchesMovedWithEvent:", event)
}

func (r_ Responder) TouchesCancelledWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, "touchesCancelledWithEvent:", event)
}

func (r_ Responder) TouchesEndedWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, "touchesEndedWithEvent:", event)
}

func (r_ Responder) WantsForwardedScrollEventsForAxis(axis EventGestureAxis) bool {
	rv := objc.CallMethod[bool](r_, "wantsForwardedScrollEventsForAxis:", axis)
	return rv
}

func (r_ Responder) SmartMagnifyWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, "smartMagnifyWithEvent:", event)
}

func (r_ Responder) WantsScrollEventsForSwipeTrackingOnAxis(axis EventGestureAxis) bool {
	rv := objc.CallMethod[bool](r_, "wantsScrollEventsForSwipeTrackingOnAxis:", axis)
	return rv
}

func (r_ Responder) MakeTouchBar() TouchBar {
	rv := objc.CallMethod[TouchBar](r_, "makeTouchBar")
	return rv
}

func (r_ Responder) PerformTextFinderAction(sender objc.IObject) {
	objc.CallMethod[objc.Void](r_, "performTextFinderAction:", sender)
}

func (r_ Responder) NewWindowForTab(sender objc.IObject) {
	objc.CallMethod[objc.Void](r_, "newWindowForTab:", sender)
}

func (r_ Responder) ShowContextHelp(sender objc.IObject) {
	objc.CallMethod[objc.Void](r_, "showContextHelp:", sender)
}

func (r_ Responder) AcceptsFirstResponder() bool {
	rv := objc.CallMethod[bool](r_, "acceptsFirstResponder")
	return rv
}

func (r_ Responder) NextResponder() Responder {
	rv := objc.CallMethod[Responder](r_, "nextResponder")
	return rv
}

func (r_ Responder) SetNextResponder(value IResponder) {
	objc.CallMethod[objc.Void](r_, "setNextResponder:", value)
}

func (rc _ResponderClass) RestorableStateKeyPaths() []string {
	rv := objc.CallMethod[[]string](rc, "restorableStateKeyPaths")
	return rv
}

func (r_ Responder) UserActivity() foundation.UserActivity {
	rv := objc.CallMethod[foundation.UserActivity](r_, "userActivity")
	return rv
}

func (r_ Responder) SetUserActivity(value foundation.IUserActivity) {
	objc.CallMethod[objc.Void](r_, "setUserActivity:", value)
}

func (r_ Responder) Menu() Menu {
	rv := objc.CallMethod[Menu](r_, "menu")
	return rv
}

func (r_ Responder) SetMenu(value IMenu) {
	objc.CallMethod[objc.Void](r_, "setMenu:", value)
}

func (r_ Responder) UndoManager() foundation.UndoManager {
	rv := objc.CallMethod[foundation.UndoManager](r_, "undoManager")
	return rv
}

func (r_ Responder) TouchBar() TouchBar {
	rv := objc.CallMethod[TouchBar](r_, "touchBar")
	return rv
}

func (r_ Responder) SetTouchBar(value ITouchBar) {
	objc.CallMethod[objc.Void](r_, "setTouchBar:", value)
}
