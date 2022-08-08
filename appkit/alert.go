// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var AlertClass = _AlertClass{objc.GetClass("NSAlert")}

type _AlertClass struct {
	objc.Class
}

type IAlert interface {
	objc.IObject
	Layout()
	RunModal() ModalResponse
	BeginSheetModalForWindow_CompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse))
	// deprecated
	BeginSheetModalForWindow_ModalDelegate_DidEndSelector_ContextInfo(window IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer)
	AddButtonWithTitle(title string) Button
	AlertStyle() AlertStyle
	SetAlertStyle(value AlertStyle)
	AccessoryView() View
	SetAccessoryView(value IView)
	ShowsHelp() bool
	SetShowsHelp(value bool)
	HelpAnchor() HelpAnchorName
	SetHelpAnchor(value HelpAnchorName)
	Delegate() AlertDelegateWrapper
	SetDelegate(value AlertDelegate)
	SuppressionButton() Button
	ShowsSuppressionButton() bool
	SetShowsSuppressionButton(value bool)
	InformativeText() string
	SetInformativeText(value string)
	MessageText() string
	SetMessageText(value string)
	Icon() Image
	SetIcon(value IImage)
	Buttons() []Button
	Window() Window
}

type Alert struct {
	objc.Object
}

func MakeAlert(ptr unsafe.Pointer) Alert {
	return Alert{
		Object: objc.MakeObject(ptr),
	}
}

func (ac _AlertClass) Alloc() Alert {
	rv := ffi.CallMethod[Alert](ac, "alloc")
	return rv
}

func (a_ Alert) Init() Alert {
	rv := ffi.CallMethod[Alert](a_, "init")
	return rv
}

func (ac _AlertClass) New() Alert {
	rv := ffi.CallMethod[Alert](ac, "new")
	rv.Autorelease()
	return rv
}

func NewAlert() Alert {
	return AlertClass.New()
}

func (ac _AlertClass) AlertWithError(error foundation.IError) Alert {
	rv := ffi.CallMethod[Alert](ac, "alertWithError:", error)
	return rv
}

func (a_ Alert) Layout() {
	ffi.CallMethod[ffi.Void](a_, "layout")
}

func (a_ Alert) RunModal() ModalResponse {
	rv := ffi.CallMethod[ModalResponse](a_, "runModal")
	return rv
}

func (a_ Alert) BeginSheetModalForWindow_CompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse)) {
	ffi.CallMethod[ffi.Void](a_, "beginSheetModalForWindow:completionHandler:", sheetWindow, handler)
}

// deprecated
func (a_ Alert) BeginSheetModalForWindow_ModalDelegate_DidEndSelector_ContextInfo(window IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](a_, "beginSheetModalForWindow:modalDelegate:didEndSelector:contextInfo:", window, delegate, didEndSelector, contextInfo)
}

func (a_ Alert) AddButtonWithTitle(title string) Button {
	rv := ffi.CallMethod[Button](a_, "addButtonWithTitle:", title)
	return rv
}

func (a_ Alert) AlertStyle() AlertStyle {
	rv := ffi.CallMethod[AlertStyle](a_, "alertStyle")
	return rv
}

func (a_ Alert) SetAlertStyle(value AlertStyle) {
	ffi.CallMethod[ffi.Void](a_, "setAlertStyle:", value)
}

func (a_ Alert) AccessoryView() View {
	rv := ffi.CallMethod[View](a_, "accessoryView")
	return rv
}

func (a_ Alert) SetAccessoryView(value IView) {
	ffi.CallMethod[ffi.Void](a_, "setAccessoryView:", value)
}

func (a_ Alert) ShowsHelp() bool {
	rv := ffi.CallMethod[bool](a_, "showsHelp")
	return rv
}

func (a_ Alert) SetShowsHelp(value bool) {
	ffi.CallMethod[ffi.Void](a_, "setShowsHelp:", value)
}

func (a_ Alert) HelpAnchor() HelpAnchorName {
	rv := ffi.CallMethod[HelpAnchorName](a_, "helpAnchor")
	return rv
}

func (a_ Alert) SetHelpAnchor(value HelpAnchorName) {
	ffi.CallMethod[ffi.Void](a_, "setHelpAnchor:", value)
}

func (a_ Alert) Delegate() AlertDelegateWrapper {
	rv := ffi.CallMethod[AlertDelegateWrapper](a_, "delegate")
	return rv
}

func (a_ Alert) SetDelegate(value AlertDelegate) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(a_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](a_, "setDelegate:", po)
}

func (a_ Alert) SuppressionButton() Button {
	rv := ffi.CallMethod[Button](a_, "suppressionButton")
	return rv
}

func (a_ Alert) ShowsSuppressionButton() bool {
	rv := ffi.CallMethod[bool](a_, "showsSuppressionButton")
	return rv
}

func (a_ Alert) SetShowsSuppressionButton(value bool) {
	ffi.CallMethod[ffi.Void](a_, "setShowsSuppressionButton:", value)
}

func (a_ Alert) InformativeText() string {
	rv := ffi.CallMethod[string](a_, "informativeText")
	return rv
}

func (a_ Alert) SetInformativeText(value string) {
	ffi.CallMethod[ffi.Void](a_, "setInformativeText:", value)
}

func (a_ Alert) MessageText() string {
	rv := ffi.CallMethod[string](a_, "messageText")
	return rv
}

func (a_ Alert) SetMessageText(value string) {
	ffi.CallMethod[ffi.Void](a_, "setMessageText:", value)
}

func (a_ Alert) Icon() Image {
	rv := ffi.CallMethod[Image](a_, "icon")
	return rv
}

func (a_ Alert) SetIcon(value IImage) {
	ffi.CallMethod[ffi.Void](a_, "setIcon:", value)
}

func (a_ Alert) Buttons() []Button {
	rv := ffi.CallMethod[[]Button](a_, "buttons")
	return rv
}

func (a_ Alert) Window() Window {
	rv := ffi.CallMethod[Window](a_, "window")
	return rv
}

type AlertDelegate interface {
	ImplementsAlertShowHelp() bool
	// optional
	AlertShowHelp(alert Alert) bool
}

type AlertDelegateImpl struct {
	_AlertShowHelp func(alert Alert) bool
}

func (di *AlertDelegateImpl) ImplementsAlertShowHelp() bool {
	return di._AlertShowHelp != nil
}

func (di *AlertDelegateImpl) SetAlertShowHelp(f func(alert Alert) bool) {
	di._AlertShowHelp = f
}

func (di *AlertDelegateImpl) AlertShowHelp(alert Alert) bool {
	return di._AlertShowHelp(alert)
}

type AlertDelegateWrapper struct {
	objc.Object
}

func (a_ *AlertDelegateWrapper) ImplementsAlertShowHelp() bool {
	return a_.RespondsToSelector(objc.GetSelector("alertShowHelp:"))
}

func (a_ AlertDelegateWrapper) AlertShowHelp(alert IAlert) bool {
	rv := ffi.CallMethod[bool](a_, "alertShowHelp:", alert)
	return rv
}
