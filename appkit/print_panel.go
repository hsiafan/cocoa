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
	rv := objc.CallMethod[PrintPanel](pc, objc.SEL("alloc"))
	return rv
}

func (pc _PrintPanelClass) New() PrintPanel {
	rv := objc.CallMethod[PrintPanel](pc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewPrintPanel() PrintPanel {
	return PrintPanelClass.New()
}

func (p_ PrintPanel) Init() PrintPanel {
	rv := objc.CallMethod[PrintPanel](p_, objc.SEL("init"))
	return rv
}

func (pc _PrintPanelClass) PrintPanel() PrintPanel {
	rv := objc.CallMethod[PrintPanel](pc, objc.SEL("printPanel"))
	return rv
}

func (p_ PrintPanel) DefaultButtonTitle() string {
	rv := objc.CallMethod[string](p_, objc.SEL("defaultButtonTitle"))
	return rv
}

func (p_ PrintPanel) SetDefaultButtonTitle(defaultButtonTitle string) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setDefaultButtonTitle:"), defaultButtonTitle)
}

func (p_ PrintPanel) AddAccessoryController(accessoryController IViewController) {
	objc.CallMethod[objc.Void](p_, objc.SEL("addAccessoryController:"), objc.ExtractPtr(accessoryController))
}

func (p_ PrintPanel) RemoveAccessoryController(accessoryController IViewController) {
	objc.CallMethod[objc.Void](p_, objc.SEL("removeAccessoryController:"), objc.ExtractPtr(accessoryController))
}

func (p_ PrintPanel) BeginSheetWithPrintInfo_ModalForWindow_Delegate_DidEndSelector_ContextInfo(printInfo IPrintInfo, docWindow IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](p_, objc.SEL("beginSheetWithPrintInfo:modalForWindow:delegate:didEndSelector:contextInfo:"), objc.ExtractPtr(printInfo), objc.ExtractPtr(docWindow), objc.ExtractPtr(delegate), didEndSelector, contextInfo)
}

func (p_ PrintPanel) RunModal() int {
	rv := objc.CallMethod[int](p_, objc.SEL("runModal"))
	return rv
}

func (p_ PrintPanel) RunModalWithPrintInfo(printInfo IPrintInfo) int {
	rv := objc.CallMethod[int](p_, objc.SEL("runModalWithPrintInfo:"), objc.ExtractPtr(printInfo))
	return rv
}

// deprecated
func (p_ PrintPanel) AccessoryView() View {
	rv := objc.CallMethod[View](p_, objc.SEL("accessoryView"))
	return rv
}

// deprecated
func (p_ PrintPanel) SetAccessoryView(accessoryView IView) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setAccessoryView:"), objc.ExtractPtr(accessoryView))
}

// deprecated
func (p_ PrintPanel) UpdateFromPrintInfo() {
	objc.CallMethod[objc.Void](p_, objc.SEL("updateFromPrintInfo"))
}

// deprecated
func (p_ PrintPanel) FinalWritePrintInfo() {
	objc.CallMethod[objc.Void](p_, objc.SEL("finalWritePrintInfo"))
}

func (p_ PrintPanel) JobStyleHint() PrintPanelJobStyleHint {
	rv := objc.CallMethod[PrintPanelJobStyleHint](p_, objc.SEL("jobStyleHint"))
	return rv
}

func (p_ PrintPanel) SetJobStyleHint(value PrintPanelJobStyleHint) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setJobStyleHint:"), value)
}

func (p_ PrintPanel) Options() PrintPanelOptions {
	rv := objc.CallMethod[PrintPanelOptions](p_, objc.SEL("options"))
	return rv
}

func (p_ PrintPanel) SetOptions(value PrintPanelOptions) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setOptions:"), value)
}

func (p_ PrintPanel) HelpAnchor() HelpAnchorName {
	rv := objc.CallMethod[HelpAnchorName](p_, objc.SEL("helpAnchor"))
	return rv
}

func (p_ PrintPanel) SetHelpAnchor(value HelpAnchorName) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setHelpAnchor:"), value)
}

func (p_ PrintPanel) AccessoryControllers() []ViewController {
	rv := objc.CallMethod[[]ViewController](p_, objc.SEL("accessoryControllers"))
	return rv
}

func (p_ PrintPanel) PrintInfo() PrintInfo {
	rv := objc.CallMethod[PrintInfo](p_, objc.SEL("printInfo"))
	return rv
}
