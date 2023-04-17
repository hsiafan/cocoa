// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[PrintPanel](pc, "alloc")
	return rv
}

func (pc _PrintPanelClass) New() PrintPanel {
	rv := ffi.CallMethod[PrintPanel](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPrintPanel() PrintPanel {
	return PrintPanelClass.New()
}

func (p_ PrintPanel) Init() PrintPanel {
	rv := ffi.CallMethod[PrintPanel](p_, "init")
	return rv
}

func (pc _PrintPanelClass) PrintPanel() PrintPanel {
	rv := ffi.CallMethod[PrintPanel](pc, "printPanel")
	return rv
}

func (p_ PrintPanel) DefaultButtonTitle() string {
	rv := ffi.CallMethod[string](p_, "defaultButtonTitle")
	return rv
}

func (p_ PrintPanel) SetDefaultButtonTitle(defaultButtonTitle string) {
	ffi.CallMethod[ffi.Void](p_, "setDefaultButtonTitle:", defaultButtonTitle)
}

func (p_ PrintPanel) AddAccessoryController(accessoryController IViewController) {
	ffi.CallMethod[ffi.Void](p_, "addAccessoryController:", accessoryController)
}

func (p_ PrintPanel) RemoveAccessoryController(accessoryController IViewController) {
	ffi.CallMethod[ffi.Void](p_, "removeAccessoryController:", accessoryController)
}

func (p_ PrintPanel) BeginSheetWithPrintInfo_ModalForWindow_Delegate_DidEndSelector_ContextInfo(printInfo IPrintInfo, docWindow IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](p_, "beginSheetWithPrintInfo:modalForWindow:delegate:didEndSelector:contextInfo:", printInfo, docWindow, delegate, didEndSelector, contextInfo)
}

func (p_ PrintPanel) RunModal() int {
	rv := ffi.CallMethod[int](p_, "runModal")
	return rv
}

func (p_ PrintPanel) RunModalWithPrintInfo(printInfo IPrintInfo) int {
	rv := ffi.CallMethod[int](p_, "runModalWithPrintInfo:", printInfo)
	return rv
}

// deprecated
func (p_ PrintPanel) AccessoryView() View {
	rv := ffi.CallMethod[View](p_, "accessoryView")
	return rv
}

// deprecated
func (p_ PrintPanel) SetAccessoryView(accessoryView IView) {
	ffi.CallMethod[ffi.Void](p_, "setAccessoryView:", accessoryView)
}

// deprecated
func (p_ PrintPanel) UpdateFromPrintInfo() {
	ffi.CallMethod[ffi.Void](p_, "updateFromPrintInfo")
}

// deprecated
func (p_ PrintPanel) FinalWritePrintInfo() {
	ffi.CallMethod[ffi.Void](p_, "finalWritePrintInfo")
}

func (p_ PrintPanel) JobStyleHint() PrintPanelJobStyleHint {
	rv := ffi.CallMethod[PrintPanelJobStyleHint](p_, "jobStyleHint")
	return rv
}

func (p_ PrintPanel) SetJobStyleHint(value PrintPanelJobStyleHint) {
	ffi.CallMethod[ffi.Void](p_, "setJobStyleHint:", value)
}

func (p_ PrintPanel) Options() PrintPanelOptions {
	rv := ffi.CallMethod[PrintPanelOptions](p_, "options")
	return rv
}

func (p_ PrintPanel) SetOptions(value PrintPanelOptions) {
	ffi.CallMethod[ffi.Void](p_, "setOptions:", value)
}

func (p_ PrintPanel) HelpAnchor() HelpAnchorName {
	rv := ffi.CallMethod[HelpAnchorName](p_, "helpAnchor")
	return rv
}

func (p_ PrintPanel) SetHelpAnchor(value HelpAnchorName) {
	ffi.CallMethod[ffi.Void](p_, "setHelpAnchor:", value)
}

func (p_ PrintPanel) AccessoryControllers() []ViewController {
	rv := ffi.CallMethod[[]ViewController](p_, "accessoryControllers")
	return rv
}

func (p_ PrintPanel) PrintInfo() PrintInfo {
	rv := ffi.CallMethod[PrintInfo](p_, "printInfo")
	return rv
}
