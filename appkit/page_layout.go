// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var PageLayoutClass = _PageLayoutClass{objc.GetClass("NSPageLayout")}

type _PageLayoutClass struct {
	objc.Class
}

type IPageLayout interface {
	objc.IObject
	BeginSheetWithPrintInfo_ModalForWindow_Delegate_DidEndSelector_ContextInfo(printInfo IPrintInfo, docWindow IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer)
	RunModal() int
	RunModalWithPrintInfo(printInfo IPrintInfo) int
	AddAccessoryController(accessoryController IViewController)
	RemoveAccessoryController(accessoryController IViewController)
	// deprecated
	AccessoryView() View
	// deprecated
	SetAccessoryView(accessoryView IView)
	// deprecated
	ReadPrintInfo()
	// deprecated
	WritePrintInfo()
	AccessoryControllers() []ViewController
	PrintInfo() PrintInfo
}

type PageLayout struct {
	objc.Object
}

func MakePageLayout(ptr unsafe.Pointer) PageLayout {
	return PageLayout{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PageLayoutClass) Alloc() PageLayout {
	rv := objc.CallMethod[PageLayout](pc, objc.GetSelector("alloc"))
	return rv
}

func (pc _PageLayoutClass) New() PageLayout {
	rv := objc.CallMethod[PageLayout](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPageLayout() PageLayout {
	return PageLayoutClass.New()
}

func (p_ PageLayout) Init() PageLayout {
	rv := objc.CallMethod[PageLayout](p_, objc.GetSelector("init"))
	return rv
}

func (pc _PageLayoutClass) PageLayout() PageLayout {
	rv := objc.CallMethod[PageLayout](pc, objc.GetSelector("pageLayout"))
	return rv
}

func (p_ PageLayout) BeginSheetWithPrintInfo_ModalForWindow_Delegate_DidEndSelector_ContextInfo(printInfo IPrintInfo, docWindow IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("beginSheetWithPrintInfo:modalForWindow:delegate:didEndSelector:contextInfo:"), printInfo, docWindow, delegate, didEndSelector, contextInfo)
}

func (p_ PageLayout) RunModal() int {
	rv := objc.CallMethod[int](p_, objc.GetSelector("runModal"))
	return rv
}

func (p_ PageLayout) RunModalWithPrintInfo(printInfo IPrintInfo) int {
	rv := objc.CallMethod[int](p_, objc.GetSelector("runModalWithPrintInfo:"), printInfo)
	return rv
}

func (p_ PageLayout) AddAccessoryController(accessoryController IViewController) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("addAccessoryController:"), accessoryController)
}

func (p_ PageLayout) RemoveAccessoryController(accessoryController IViewController) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("removeAccessoryController:"), accessoryController)
}

// deprecated
func (p_ PageLayout) AccessoryView() View {
	rv := objc.CallMethod[View](p_, objc.GetSelector("accessoryView"))
	return rv
}

// deprecated
func (p_ PageLayout) SetAccessoryView(accessoryView IView) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setAccessoryView:"), accessoryView)
}

// deprecated
func (p_ PageLayout) ReadPrintInfo() {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("readPrintInfo"))
}

// deprecated
func (p_ PageLayout) WritePrintInfo() {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("writePrintInfo"))
}

func (p_ PageLayout) AccessoryControllers() []ViewController {
	rv := objc.CallMethod[[]ViewController](p_, objc.GetSelector("accessoryControllers"))
	return rv
}

func (p_ PageLayout) PrintInfo() PrintInfo {
	rv := objc.CallMethod[PrintInfo](p_, objc.GetSelector("printInfo"))
	return rv
}
