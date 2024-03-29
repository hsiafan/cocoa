// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
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
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
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
	rv := objc.CallMethod[Alert](ac, objc.SEL("alloc"))
	return rv
}

func (ac _AlertClass) New() Alert {
	rv := objc.CallMethod[Alert](ac, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewAlert() Alert {
	return AlertClass.New()
}

func (a_ Alert) Init() Alert {
	rv := objc.CallMethod[Alert](a_, objc.SEL("init"))
	return rv
}

func (ac _AlertClass) AlertWithError(error foundation.IError) Alert {
	rv := objc.CallMethod[Alert](ac, objc.SEL("alertWithError:"), objc.ExtractPtr(error))
	return rv
}

func (a_ Alert) Layout() {
	objc.CallMethod[objc.Void](a_, objc.SEL("layout"))
}

func (a_ Alert) RunModal() ModalResponse {
	rv := objc.CallMethod[ModalResponse](a_, objc.SEL("runModal"))
	return rv
}

func (a_ Alert) BeginSheetModalForWindow_CompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse)) {
	objc.CallMethod[objc.Void](a_, objc.SEL("beginSheetModalForWindow:completionHandler:"), objc.ExtractPtr(sheetWindow), handler)
}

// deprecated
func (a_ Alert) BeginSheetModalForWindow_ModalDelegate_DidEndSelector_ContextInfo(window IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](a_, objc.SEL("beginSheetModalForWindow:modalDelegate:didEndSelector:contextInfo:"), objc.ExtractPtr(window), objc.ExtractPtr(delegate), didEndSelector, contextInfo)
}

func (a_ Alert) AddButtonWithTitle(title string) Button {
	rv := objc.CallMethod[Button](a_, objc.SEL("addButtonWithTitle:"), title)
	return rv
}

func (a_ Alert) AlertStyle() AlertStyle {
	rv := objc.CallMethod[AlertStyle](a_, objc.SEL("alertStyle"))
	return rv
}

func (a_ Alert) SetAlertStyle(value AlertStyle) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setAlertStyle:"), value)
}

func (a_ Alert) AccessoryView() View {
	rv := objc.CallMethod[View](a_, objc.SEL("accessoryView"))
	return rv
}

func (a_ Alert) SetAccessoryView(value IView) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setAccessoryView:"), objc.ExtractPtr(value))
}

func (a_ Alert) ShowsHelp() bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("showsHelp"))
	return rv
}

func (a_ Alert) SetShowsHelp(value bool) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setShowsHelp:"), value)
}

func (a_ Alert) HelpAnchor() HelpAnchorName {
	rv := objc.CallMethod[HelpAnchorName](a_, objc.SEL("helpAnchor"))
	return rv
}

func (a_ Alert) SetHelpAnchor(value HelpAnchorName) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setHelpAnchor:"), value)
}

// weak property
func (a_ Alert) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](a_, objc.SEL("delegate"))
	return rv
}

// weak property
func (a_ Alert) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}

func (a_ Alert) SuppressionButton() Button {
	rv := objc.CallMethod[Button](a_, objc.SEL("suppressionButton"))
	return rv
}

func (a_ Alert) ShowsSuppressionButton() bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("showsSuppressionButton"))
	return rv
}

func (a_ Alert) SetShowsSuppressionButton(value bool) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setShowsSuppressionButton:"), value)
}

func (a_ Alert) InformativeText() string {
	rv := objc.CallMethod[string](a_, objc.SEL("informativeText"))
	return rv
}

func (a_ Alert) SetInformativeText(value string) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setInformativeText:"), value)
}

func (a_ Alert) MessageText() string {
	rv := objc.CallMethod[string](a_, objc.SEL("messageText"))
	return rv
}

func (a_ Alert) SetMessageText(value string) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setMessageText:"), value)
}

func (a_ Alert) Icon() Image {
	rv := objc.CallMethod[Image](a_, objc.SEL("icon"))
	return rv
}

func (a_ Alert) SetIcon(value IImage) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setIcon:"), objc.ExtractPtr(value))
}

func (a_ Alert) Buttons() []Button {
	rv := objc.CallMethod[[]Button](a_, objc.SEL("buttons"))
	return rv
}

func (a_ Alert) Window() Window {
	rv := objc.CallMethod[Window](a_, objc.SEL("window"))
	return rv
}
