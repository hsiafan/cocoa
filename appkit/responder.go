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
	rv := objc.CallMethod[Responder](r_, objc.SEL("init"))
	return rv
}

func (rc _ResponderClass) Alloc() Responder {
	rv := objc.CallMethod[Responder](rc, objc.SEL("alloc"))
	return rv
}

func (rc _ResponderClass) New() Responder {
	rv := objc.CallMethod[Responder](rc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewResponder() Responder {
	return ResponderClass.New()
}

func (r_ Responder) BecomeFirstResponder() bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("becomeFirstResponder"))
	return rv
}

func (r_ Responder) ResignFirstResponder() bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("resignFirstResponder"))
	return rv
}

func (r_ Responder) ValidateProposedFirstResponder_ForEvent(responder IResponder, event IEvent) bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("validateProposedFirstResponder:forEvent:"), objc.ExtractPtr(responder), objc.ExtractPtr(event))
	return rv
}

func (r_ Responder) MouseDown(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("mouseDown:"), objc.ExtractPtr(event))
}

func (r_ Responder) MouseDragged(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("mouseDragged:"), objc.ExtractPtr(event))
}

func (r_ Responder) MouseUp(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("mouseUp:"), objc.ExtractPtr(event))
}

func (r_ Responder) MouseMoved(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("mouseMoved:"), objc.ExtractPtr(event))
}

func (r_ Responder) MouseEntered(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("mouseEntered:"), objc.ExtractPtr(event))
}

func (r_ Responder) MouseExited(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("mouseExited:"), objc.ExtractPtr(event))
}

func (r_ Responder) RightMouseDown(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("rightMouseDown:"), objc.ExtractPtr(event))
}

func (r_ Responder) RightMouseDragged(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("rightMouseDragged:"), objc.ExtractPtr(event))
}

func (r_ Responder) RightMouseUp(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("rightMouseUp:"), objc.ExtractPtr(event))
}

func (r_ Responder) OtherMouseDown(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("otherMouseDown:"), objc.ExtractPtr(event))
}

func (r_ Responder) OtherMouseDragged(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("otherMouseDragged:"), objc.ExtractPtr(event))
}

func (r_ Responder) OtherMouseUp(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("otherMouseUp:"), objc.ExtractPtr(event))
}

func (r_ Responder) KeyDown(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("keyDown:"), objc.ExtractPtr(event))
}

func (r_ Responder) KeyUp(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("keyUp:"), objc.ExtractPtr(event))
}

func (r_ Responder) InterpretKeyEvents(eventArray []IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("interpretKeyEvents:"), eventArray)
}

func (r_ Responder) PerformKeyEquivalent(event IEvent) bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("performKeyEquivalent:"), objc.ExtractPtr(event))
	return rv
}

// deprecated
func (r_ Responder) PerformMnemonic(string_ string) bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("performMnemonic:"), string_)
	return rv
}

func (r_ Responder) FlushBufferedKeyEvents() {
	objc.CallMethod[objc.Void](r_, objc.SEL("flushBufferedKeyEvents"))
}

func (r_ Responder) PressureChangeWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("pressureChangeWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) CursorUpdate(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("cursorUpdate:"), objc.ExtractPtr(event))
}

func (r_ Responder) FlagsChanged(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("flagsChanged:"), objc.ExtractPtr(event))
}

func (r_ Responder) TabletPoint(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("tabletPoint:"), objc.ExtractPtr(event))
}

func (r_ Responder) TabletProximity(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("tabletProximity:"), objc.ExtractPtr(event))
}

func (r_ Responder) HelpRequested(eventPtr IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("helpRequested:"), objc.ExtractPtr(eventPtr))
}

func (r_ Responder) ScrollWheel(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("scrollWheel:"), objc.ExtractPtr(event))
}

func (r_ Responder) QuickLookWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("quickLookWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) ChangeModeWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("changeModeWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) SupplementalTargetForAction_Sender(action objc.Selector, sender objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](r_, objc.SEL("supplementalTargetForAction:sender:"), action, objc.ExtractPtr(sender))
	return rv
}

func (rc _ResponderClass) AllowedClassesForRestorableStateKeyPath(keyPath string) []objc.Class {
	rv := objc.CallMethod[[]objc.Class](rc, objc.SEL("allowedClassesForRestorableStateKeyPath:"), keyPath)
	return rv
}

func (r_ Responder) EncodeRestorableStateWithCoder(coder foundation.ICoder) {
	objc.CallMethod[objc.Void](r_, objc.SEL("encodeRestorableStateWithCoder:"), objc.ExtractPtr(coder))
}

func (r_ Responder) EncodeRestorableStateWithCoder_BackgroundQueue(coder foundation.ICoder, queue foundation.IOperationQueue) {
	objc.CallMethod[objc.Void](r_, objc.SEL("encodeRestorableStateWithCoder:backgroundQueue:"), objc.ExtractPtr(coder), objc.ExtractPtr(queue))
}

func (r_ Responder) RestoreStateWithCoder(coder foundation.ICoder) {
	objc.CallMethod[objc.Void](r_, objc.SEL("restoreStateWithCoder:"), objc.ExtractPtr(coder))
}

func (r_ Responder) InvalidateRestorableState() {
	objc.CallMethod[objc.Void](r_, objc.SEL("invalidateRestorableState"))
}

func (r_ Responder) UpdateUserActivityState(userActivity foundation.IUserActivity) {
	objc.CallMethod[objc.Void](r_, objc.SEL("updateUserActivityState:"), objc.ExtractPtr(userActivity))
}

func (r_ Responder) PresentError(error foundation.IError) bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("presentError:"), objc.ExtractPtr(error))
	return rv
}

func (r_ Responder) PresentError_ModalForWindow_Delegate_DidPresentSelector_ContextInfo(error foundation.IError, window IWindow, delegate objc.IObject, didPresentSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](r_, objc.SEL("presentError:modalForWindow:delegate:didPresentSelector:contextInfo:"), objc.ExtractPtr(error), objc.ExtractPtr(window), objc.ExtractPtr(delegate), didPresentSelector, contextInfo)
}

func (r_ Responder) WillPresentError(error foundation.IError) foundation.Error {
	rv := objc.CallMethod[foundation.Error](r_, objc.SEL("willPresentError:"), objc.ExtractPtr(error))
	return rv
}

func (r_ Responder) TryToPerform_With(action objc.Selector, object objc.IObject) bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("tryToPerform:with:"), action, objc.ExtractPtr(object))
	return rv
}

func (r_ Responder) ValidRequestorForSendType_ReturnType(sendType PasteboardType, returnType PasteboardType) objc.Object {
	rv := objc.CallMethod[objc.Object](r_, objc.SEL("validRequestorForSendType:returnType:"), sendType, returnType)
	return rv
}

func (r_ Responder) ShouldBeTreatedAsInkEvent(event IEvent) bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("shouldBeTreatedAsInkEvent:"), objc.ExtractPtr(event))
	return rv
}

func (r_ Responder) NoResponderFor(eventSelector objc.Selector) {
	objc.CallMethod[objc.Void](r_, objc.SEL("noResponderFor:"), eventSelector)
}

// deprecated
func (r_ Responder) SetInterfaceStyle(interfaceStyle InterfaceStyle) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setInterfaceStyle:"), interfaceStyle)
}

// deprecated
func (r_ Responder) InterfaceStyle() InterfaceStyle {
	rv := objc.CallMethod[InterfaceStyle](r_, objc.SEL("interfaceStyle"))
	return rv
}

func (r_ Responder) BeginGestureWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("beginGestureWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) EndGestureWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("endGestureWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) MagnifyWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("magnifyWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) RotateWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("rotateWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) SwipeWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("swipeWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) TouchesBeganWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("touchesBeganWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) TouchesMovedWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("touchesMovedWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) TouchesCancelledWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("touchesCancelledWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) TouchesEndedWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("touchesEndedWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) WantsForwardedScrollEventsForAxis(axis EventGestureAxis) bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("wantsForwardedScrollEventsForAxis:"), axis)
	return rv
}

func (r_ Responder) SmartMagnifyWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](r_, objc.SEL("smartMagnifyWithEvent:"), objc.ExtractPtr(event))
}

func (r_ Responder) WantsScrollEventsForSwipeTrackingOnAxis(axis EventGestureAxis) bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("wantsScrollEventsForSwipeTrackingOnAxis:"), axis)
	return rv
}

func (r_ Responder) MakeTouchBar() TouchBar {
	rv := objc.CallMethod[TouchBar](r_, objc.SEL("makeTouchBar"))
	return rv
}

func (r_ Responder) PerformTextFinderAction(sender objc.IObject) {
	objc.CallMethod[objc.Void](r_, objc.SEL("performTextFinderAction:"), objc.ExtractPtr(sender))
}

func (r_ Responder) NewWindowForTab(sender objc.IObject) {
	objc.CallMethod[objc.Void](r_, objc.SEL("newWindowForTab:"), objc.ExtractPtr(sender))
}

func (r_ Responder) ShowContextHelp(sender objc.IObject) {
	objc.CallMethod[objc.Void](r_, objc.SEL("showContextHelp:"), objc.ExtractPtr(sender))
}

func (r_ Responder) AcceptsFirstResponder() bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("acceptsFirstResponder"))
	return rv
}

// weak property
func (r_ Responder) NextResponder() Responder {
	rv := objc.CallMethod[Responder](r_, objc.SEL("nextResponder"))
	return rv
}

// weak property
func (r_ Responder) SetNextResponder(value IResponder) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setNextResponder:"), objc.ExtractPtr(value))
}

func (rc _ResponderClass) RestorableStateKeyPaths() []string {
	rv := objc.CallMethod[[]string](rc, objc.SEL("restorableStateKeyPaths"))
	return rv
}

func (r_ Responder) UserActivity() foundation.UserActivity {
	rv := objc.CallMethod[foundation.UserActivity](r_, objc.SEL("userActivity"))
	return rv
}

func (r_ Responder) SetUserActivity(value foundation.IUserActivity) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setUserActivity:"), objc.ExtractPtr(value))
}

func (r_ Responder) Menu() Menu {
	rv := objc.CallMethod[Menu](r_, objc.SEL("menu"))
	return rv
}

func (r_ Responder) SetMenu(value IMenu) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setMenu:"), objc.ExtractPtr(value))
}

func (r_ Responder) UndoManager() foundation.UndoManager {
	rv := objc.CallMethod[foundation.UndoManager](r_, objc.SEL("undoManager"))
	return rv
}

func (r_ Responder) TouchBar() TouchBar {
	rv := objc.CallMethod[TouchBar](r_, objc.SEL("touchBar"))
	return rv
}

func (r_ Responder) SetTouchBar(value ITouchBar) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setTouchBar:"), objc.ExtractPtr(value))
}
