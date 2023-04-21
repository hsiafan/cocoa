// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var PrintPanelClass = _PrintPanelClass{objc.GetClass("NSPrintPanel")}

type _PrintPanelClass struct {
	objc.Class
}

type IPrintPanel interface {
	objc.IObject
	DefaultButtonTitle() string
	SetDefaultButtonTitle(defaultButtonTitle string)
	AddAccessoryController(accessoryController IViewController)
	RemoveAccessoryController(accessoryController IViewController)
	BeginSheetWithPrintInfo_ModalForWindow_Delegate_DidEndSelector_ContextInfo(printInfo IPrintInfo, docWindow IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer)
	RunModal() int
	RunModalWithPrintInfo(printInfo IPrintInfo) int
	// deprecated
	AccessoryView() View
	// deprecated
	SetAccessoryView(accessoryView IView)
	// deprecated
	UpdateFromPrintInfo()
	// deprecated
	FinalWritePrintInfo()
	JobStyleHint() PrintPanelJobStyleHint
	SetJobStyleHint(value PrintPanelJobStyleHint)
	Options() PrintPanelOptions
	SetOptions(value PrintPanelOptions)
	HelpAnchor() HelpAnchorName
	SetHelpAnchor(value HelpAnchorName)
	AccessoryControllers() []ViewController
	PrintInfo() PrintInfo
}

type PrintPanel struct {
	objc.Object
}

func MakePrintPanel(ptr unsafe.Pointer) PrintPanel {
	return PrintPanel{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PrintPanelClass) Alloc() PrintPanel {
	rv := objc.CallMethod[PrintPanel](pc, "alloc")
	return rv
}

func (pc _PrintPanelClass) New() PrintPanel {
	rv := objc.CallMethod[PrintPanel](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPrintPanel() PrintPanel {
	return PrintPanelClass.New()
}

func (p_ PrintPanel) Init() PrintPanel {
	rv := objc.CallMethod[PrintPanel](p_, "init")
	return rv
}

func (pc _PrintPanelClass) PrintPanel() PrintPanel {
	rv := objc.CallMethod[PrintPanel](pc, "printPanel")
	return rv
}

func (p_ PrintPanel) DefaultButtonTitle() string {
	rv := objc.CallMethod[string](p_, "defaultButtonTitle")
	return rv
}

func (p_ PrintPanel) SetDefaultButtonTitle(defaultButtonTitle string) {
	objc.CallMethod[objc.Void](p_, "setDefaultButtonTitle:", defaultButtonTitle)
}

func (p_ PrintPanel) AddAccessoryController(accessoryController IViewController) {
	objc.CallMethod[objc.Void](p_, "addAccessoryController:", accessoryController)
}

func (p_ PrintPanel) RemoveAccessoryController(accessoryController IViewController) {
	objc.CallMethod[objc.Void](p_, "removeAccessoryController:", accessoryController)
}

func (p_ PrintPanel) BeginSheetWithPrintInfo_ModalForWindow_Delegate_DidEndSelector_ContextInfo(printInfo IPrintInfo, docWindow IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](p_, "beginSheetWithPrintInfo:modalForWindow:delegate:didEndSelector:contextInfo:", printInfo, docWindow, delegate, didEndSelector, contextInfo)
}

func (p_ PrintPanel) RunModal() int {
	rv := objc.CallMethod[int](p_, "runModal")
	return rv
}

func (p_ PrintPanel) RunModalWithPrintInfo(printInfo IPrintInfo) int {
	rv := objc.CallMethod[int](p_, "runModalWithPrintInfo:", printInfo)
	return rv
}

// deprecated
func (p_ PrintPanel) AccessoryView() View {
	rv := objc.CallMethod[View](p_, "accessoryView")
	return rv
}

// deprecated
func (p_ PrintPanel) SetAccessoryView(accessoryView IView) {
	objc.CallMethod[objc.Void](p_, "setAccessoryView:", accessoryView)
}

// deprecated
func (p_ PrintPanel) UpdateFromPrintInfo() {
	objc.CallMethod[objc.Void](p_, "updateFromPrintInfo")
}

// deprecated
func (p_ PrintPanel) FinalWritePrintInfo() {
	objc.CallMethod[objc.Void](p_, "finalWritePrintInfo")
}

func (p_ PrintPanel) JobStyleHint() PrintPanelJobStyleHint {
	rv := objc.CallMethod[PrintPanelJobStyleHint](p_, "jobStyleHint")
	return rv
}

func (p_ PrintPanel) SetJobStyleHint(value PrintPanelJobStyleHint) {
	objc.CallMethod[objc.Void](p_, "setJobStyleHint:", value)
}

func (p_ PrintPanel) Options() PrintPanelOptions {
	rv := objc.CallMethod[PrintPanelOptions](p_, "options")
	return rv
}

func (p_ PrintPanel) SetOptions(value PrintPanelOptions) {
	objc.CallMethod[objc.Void](p_, "setOptions:", value)
}

func (p_ PrintPanel) HelpAnchor() HelpAnchorName {
	rv := objc.CallMethod[HelpAnchorName](p_, "helpAnchor")
	return rv
}

func (p_ PrintPanel) SetHelpAnchor(value HelpAnchorName) {
	objc.CallMethod[objc.Void](p_, "setHelpAnchor:", value)
}

func (p_ PrintPanel) AccessoryControllers() []ViewController {
	rv := objc.CallMethod[[]ViewController](p_, "accessoryControllers")
	return rv
}

func (p_ PrintPanel) PrintInfo() PrintInfo {
	rv := objc.CallMethod[PrintInfo](p_, "printInfo")
	return rv
}
