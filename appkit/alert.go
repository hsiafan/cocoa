// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	SetDelegate0(value objc.IObject)
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
	rv := objc.CallMethod[Alert](ac, objc.GetSelector("alloc"))
	return rv
}

func (ac _AlertClass) New() Alert {
	rv := objc.CallMethod[Alert](ac, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewAlert() Alert {
	return AlertClass.New()
}

func (a_ Alert) Init() Alert {
	rv := objc.CallMethod[Alert](a_, objc.GetSelector("init"))
	return rv
}

func (ac _AlertClass) AlertWithError(error foundation.IError) Alert {
	rv := objc.CallMethod[Alert](ac, objc.GetSelector("alertWithError:"), objc.ExtractPtr(error))
	return rv
}

func (a_ Alert) Layout() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("layout"))
}

func (a_ Alert) RunModal() ModalResponse {
	rv := objc.CallMethod[ModalResponse](a_, objc.GetSelector("runModal"))
	return rv
}

func (a_ Alert) BeginSheetModalForWindow_CompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse)) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("beginSheetModalForWindow:completionHandler:"), objc.ExtractPtr(sheetWindow), handler)
}

// deprecated
func (a_ Alert) BeginSheetModalForWindow_ModalDelegate_DidEndSelector_ContextInfo(window IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("beginSheetModalForWindow:modalDelegate:didEndSelector:contextInfo:"), objc.ExtractPtr(window), objc.ExtractPtr(delegate), didEndSelector, contextInfo)
}

func (a_ Alert) AddButtonWithTitle(title string) Button {
	rv := objc.CallMethod[Button](a_, objc.GetSelector("addButtonWithTitle:"), title)
	return rv
}

func (a_ Alert) AlertStyle() AlertStyle {
	rv := objc.CallMethod[AlertStyle](a_, objc.GetSelector("alertStyle"))
	return rv
}

func (a_ Alert) SetAlertStyle(value AlertStyle) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setAlertStyle:"), value)
}

func (a_ Alert) AccessoryView() View {
	rv := objc.CallMethod[View](a_, objc.GetSelector("accessoryView"))
	return rv
}

func (a_ Alert) SetAccessoryView(value IView) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setAccessoryView:"), objc.ExtractPtr(value))
}

func (a_ Alert) ShowsHelp() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("showsHelp"))
	return rv
}

func (a_ Alert) SetShowsHelp(value bool) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setShowsHelp:"), value)
}

func (a_ Alert) HelpAnchor() HelpAnchorName {
	rv := objc.CallMethod[HelpAnchorName](a_, objc.GetSelector("helpAnchor"))
	return rv
}

func (a_ Alert) SetHelpAnchor(value HelpAnchorName) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setHelpAnchor:"), value)
}

func (a_ Alert) Delegate() AlertDelegateWrapper {
	rv := objc.CallMethod[AlertDelegateWrapper](a_, objc.GetSelector("delegate"))
	return rv
}

func (a_ Alert) SetDelegate(value AlertDelegate) {
	po := objc.CreateProtocol("NSAlertDelegate", value)
	objc.SetAssociatedObject(a_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setDelegate:"), po)
}

func (a_ Alert) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (a_ Alert) SuppressionButton() Button {
	rv := objc.CallMethod[Button](a_, objc.GetSelector("suppressionButton"))
	return rv
}

func (a_ Alert) ShowsSuppressionButton() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("showsSuppressionButton"))
	return rv
}

func (a_ Alert) SetShowsSuppressionButton(value bool) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setShowsSuppressionButton:"), value)
}

func (a_ Alert) InformativeText() string {
	rv := objc.CallMethod[string](a_, objc.GetSelector("informativeText"))
	return rv
}

func (a_ Alert) SetInformativeText(value string) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setInformativeText:"), value)
}

func (a_ Alert) MessageText() string {
	rv := objc.CallMethod[string](a_, objc.GetSelector("messageText"))
	return rv
}

func (a_ Alert) SetMessageText(value string) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setMessageText:"), value)
}

func (a_ Alert) Icon() Image {
	rv := objc.CallMethod[Image](a_, objc.GetSelector("icon"))
	return rv
}

func (a_ Alert) SetIcon(value IImage) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setIcon:"), objc.ExtractPtr(value))
}

func (a_ Alert) Buttons() []Button {
	rv := objc.CallMethod[[]Button](a_, objc.GetSelector("buttons"))
	return rv
}

func (a_ Alert) Window() Window {
	rv := objc.CallMethod[Window](a_, objc.GetSelector("window"))
	return rv
}
