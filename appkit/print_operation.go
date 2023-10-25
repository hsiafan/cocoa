// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var PrintOperationClass = _PrintOperationClass{objc.GetClass("NSPrintOperation")}

type _PrintOperationClass struct {
	objc.Class
}

type IPrintOperation interface {
	objc.IObject
	RunOperation() bool
	RunOperationModalForWindow_Delegate_DidRunSelector_ContextInfo(docWindow IWindow, delegate objc.IObject, didRunSelector objc.Selector, contextInfo unsafe.Pointer)
	CleanUpOperation()
	DeliverResult() bool
	CreateContext() GraphicsContext
	DestroyContext()
	// deprecated
	JobStyleHint() string
	// deprecated
	SetJobStyleHint(hint string)
	// deprecated
	AccessoryView() View
	// deprecated
	SetAccessoryView(view IView)
	// deprecated
	ShowPanels() bool
	// deprecated
	SetShowPanels(flag bool)
	IsCopyingOperation() bool
	PrintInfo() PrintInfo
	SetPrintInfo(value IPrintInfo)
	View() View
	PreferredRenderingQuality() PrintRenderingQuality
	ShowsPrintPanel() bool
	SetShowsPrintPanel(value bool)
	ShowsProgressPanel() bool
	SetShowsProgressPanel(value bool)
	JobTitle() string
	SetJobTitle(value string)
	PrintPanel() PrintPanel
	SetPrintPanel(value IPrintPanel)
	PDFPanel() PDFPanel
	SetPDFPanel(value IPDFPanel)
	Context() GraphicsContext
	CurrentPage() int
	PageRange() foundation.Range
	PageOrder() PrintingPageOrder
	SetPageOrder(value PrintingPageOrder)
	CanSpawnSeparateThread() bool
	SetCanSpawnSeparateThread(value bool)
}

type PrintOperation struct {
	objc.Object
}

func MakePrintOperation(ptr unsafe.Pointer) PrintOperation {
	return PrintOperation{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PrintOperationClass) Alloc() PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.SEL("alloc"))
	return rv
}

func (pc _PrintOperationClass) New() PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewPrintOperation() PrintOperation {
	return PrintOperationClass.New()
}

func (p_ PrintOperation) Init() PrintOperation {
	rv := objc.CallMethod[PrintOperation](p_, objc.SEL("init"))
	return rv
}

func (pc _PrintOperationClass) EPSOperationWithView_InsideRect_ToData(view IView, rect foundation.Rect, data foundation.IMutableData) PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.SEL("EPSOperationWithView:insideRect:toData:"), objc.ExtractPtr(view), rect, objc.ExtractPtr(data))
	return rv
}

func (pc _PrintOperationClass) EPSOperationWithView_InsideRect_ToData_PrintInfo(view IView, rect foundation.Rect, data foundation.IMutableData, printInfo IPrintInfo) PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.SEL("EPSOperationWithView:insideRect:toData:printInfo:"), objc.ExtractPtr(view), rect, objc.ExtractPtr(data), objc.ExtractPtr(printInfo))
	return rv
}

func (pc _PrintOperationClass) EPSOperationWithView_InsideRect_ToPath_PrintInfo(view IView, rect foundation.Rect, path string, printInfo IPrintInfo) PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.SEL("EPSOperationWithView:insideRect:toPath:printInfo:"), objc.ExtractPtr(view), rect, path, objc.ExtractPtr(printInfo))
	return rv
}

func (pc _PrintOperationClass) PDFOperationWithView_InsideRect_ToData(view IView, rect foundation.Rect, data foundation.IMutableData) PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.SEL("PDFOperationWithView:insideRect:toData:"), objc.ExtractPtr(view), rect, objc.ExtractPtr(data))
	return rv
}

func (pc _PrintOperationClass) PDFOperationWithView_InsideRect_ToData_PrintInfo(view IView, rect foundation.Rect, data foundation.IMutableData, printInfo IPrintInfo) PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.SEL("PDFOperationWithView:insideRect:toData:printInfo:"), objc.ExtractPtr(view), rect, objc.ExtractPtr(data), objc.ExtractPtr(printInfo))
	return rv
}

func (pc _PrintOperationClass) PDFOperationWithView_InsideRect_ToPath_PrintInfo(view IView, rect foundation.Rect, path string, printInfo IPrintInfo) PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.SEL("PDFOperationWithView:insideRect:toPath:printInfo:"), objc.ExtractPtr(view), rect, path, objc.ExtractPtr(printInfo))
	return rv
}

func (pc _PrintOperationClass) PrintOperationWithView(view IView) PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.SEL("printOperationWithView:"), objc.ExtractPtr(view))
	return rv
}

func (pc _PrintOperationClass) PrintOperationWithView_PrintInfo(view IView, printInfo IPrintInfo) PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.SEL("printOperationWithView:printInfo:"), objc.ExtractPtr(view), objc.ExtractPtr(printInfo))
	return rv
}

func (p_ PrintOperation) RunOperation() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("runOperation"))
	return rv
}

func (p_ PrintOperation) RunOperationModalForWindow_Delegate_DidRunSelector_ContextInfo(docWindow IWindow, delegate objc.IObject, didRunSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](p_, objc.SEL("runOperationModalForWindow:delegate:didRunSelector:contextInfo:"), objc.ExtractPtr(docWindow), objc.ExtractPtr(delegate), didRunSelector, contextInfo)
}

func (p_ PrintOperation) CleanUpOperation() {
	objc.CallMethod[objc.Void](p_, objc.SEL("cleanUpOperation"))
}

func (p_ PrintOperation) DeliverResult() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("deliverResult"))
	return rv
}

func (p_ PrintOperation) CreateContext() GraphicsContext {
	rv := objc.CallMethod[GraphicsContext](p_, objc.SEL("createContext"))
	return rv
}

func (p_ PrintOperation) DestroyContext() {
	objc.CallMethod[objc.Void](p_, objc.SEL("destroyContext"))
}

// deprecated
func (p_ PrintOperation) JobStyleHint() string {
	rv := objc.CallMethod[string](p_, objc.SEL("jobStyleHint"))
	return rv
}

// deprecated
func (p_ PrintOperation) SetJobStyleHint(hint string) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setJobStyleHint:"), hint)
}

// deprecated
func (p_ PrintOperation) AccessoryView() View {
	rv := objc.CallMethod[View](p_, objc.SEL("accessoryView"))
	return rv
}

// deprecated
func (p_ PrintOperation) SetAccessoryView(view IView) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setAccessoryView:"), objc.ExtractPtr(view))
}

// deprecated
func (p_ PrintOperation) ShowPanels() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("showPanels"))
	return rv
}

// deprecated
func (p_ PrintOperation) SetShowPanels(flag bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setShowPanels:"), flag)
}

func (pc _PrintOperationClass) CurrentOperation() PrintOperation {
	rv := objc.CallMethod[PrintOperation](pc, objc.SEL("currentOperation"))
	return rv
}

func (pc _PrintOperationClass) SetCurrentOperation(value IPrintOperation) {
	objc.CallMethod[objc.Void](pc, objc.SEL("setCurrentOperation:"), objc.ExtractPtr(value))
}

func (p_ PrintOperation) IsCopyingOperation() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isCopyingOperation"))
	return rv
}

func (p_ PrintOperation) PrintInfo() PrintInfo {
	rv := objc.CallMethod[PrintInfo](p_, objc.SEL("printInfo"))
	return rv
}

func (p_ PrintOperation) SetPrintInfo(value IPrintInfo) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setPrintInfo:"), objc.ExtractPtr(value))
}

func (p_ PrintOperation) View() View {
	rv := objc.CallMethod[View](p_, objc.SEL("view"))
	return rv
}

func (p_ PrintOperation) PreferredRenderingQuality() PrintRenderingQuality {
	rv := objc.CallMethod[PrintRenderingQuality](p_, objc.SEL("preferredRenderingQuality"))
	return rv
}

func (p_ PrintOperation) ShowsPrintPanel() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("showsPrintPanel"))
	return rv
}

func (p_ PrintOperation) SetShowsPrintPanel(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setShowsPrintPanel:"), value)
}

func (p_ PrintOperation) ShowsProgressPanel() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("showsProgressPanel"))
	return rv
}

func (p_ PrintOperation) SetShowsProgressPanel(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setShowsProgressPanel:"), value)
}

func (p_ PrintOperation) JobTitle() string {
	rv := objc.CallMethod[string](p_, objc.SEL("jobTitle"))
	return rv
}

func (p_ PrintOperation) SetJobTitle(value string) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setJobTitle:"), value)
}

func (p_ PrintOperation) PrintPanel() PrintPanel {
	rv := objc.CallMethod[PrintPanel](p_, objc.SEL("printPanel"))
	return rv
}

func (p_ PrintOperation) SetPrintPanel(value IPrintPanel) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setPrintPanel:"), objc.ExtractPtr(value))
}

func (p_ PrintOperation) PDFPanel() PDFPanel {
	rv := objc.CallMethod[PDFPanel](p_, objc.SEL("PDFPanel"))
	return rv
}

func (p_ PrintOperation) SetPDFPanel(value IPDFPanel) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setPDFPanel:"), objc.ExtractPtr(value))
}

func (p_ PrintOperation) Context() GraphicsContext {
	rv := objc.CallMethod[GraphicsContext](p_, objc.SEL("context"))
	return rv
}

func (p_ PrintOperation) CurrentPage() int {
	rv := objc.CallMethod[int](p_, objc.SEL("currentPage"))
	return rv
}

func (p_ PrintOperation) PageRange() foundation.Range {
	rv := objc.CallMethod[foundation.Range](p_, objc.SEL("pageRange"))
	return rv
}

func (p_ PrintOperation) PageOrder() PrintingPageOrder {
	rv := objc.CallMethod[PrintingPageOrder](p_, objc.SEL("pageOrder"))
	return rv
}

func (p_ PrintOperation) SetPageOrder(value PrintingPageOrder) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setPageOrder:"), value)
}

func (p_ PrintOperation) CanSpawnSeparateThread() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("canSpawnSeparateThread"))
	return rv
}

func (p_ PrintOperation) SetCanSpawnSeparateThread(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setCanSpawnSeparateThread:"), value)
}
