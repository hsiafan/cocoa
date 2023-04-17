// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[PrintOperation](pc, "alloc")
	return rv
}

func (pc _PrintOperationClass) New() PrintOperation {
	rv := ffi.CallMethod[PrintOperation](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPrintOperation() PrintOperation {
	return PrintOperationClass.New()
}

func (p_ PrintOperation) Init() PrintOperation {
	rv := ffi.CallMethod[PrintOperation](p_, "init")
	return rv
}

func (pc _PrintOperationClass) EPSOperationWithView_InsideRect_ToData(view IView, rect foundation.Rect, data foundation.IMutableData) PrintOperation {
	rv := ffi.CallMethod[PrintOperation](pc, "EPSOperationWithView:insideRect:toData:", view, rect, data)
	return rv
}

func (pc _PrintOperationClass) EPSOperationWithView_InsideRect_ToData_PrintInfo(view IView, rect foundation.Rect, data foundation.IMutableData, printInfo IPrintInfo) PrintOperation {
	rv := ffi.CallMethod[PrintOperation](pc, "EPSOperationWithView:insideRect:toData:printInfo:", view, rect, data, printInfo)
	return rv
}

func (pc _PrintOperationClass) EPSOperationWithView_InsideRect_ToPath_PrintInfo(view IView, rect foundation.Rect, path string, printInfo IPrintInfo) PrintOperation {
	rv := ffi.CallMethod[PrintOperation](pc, "EPSOperationWithView:insideRect:toPath:printInfo:", view, rect, path, printInfo)
	return rv
}

func (pc _PrintOperationClass) PDFOperationWithView_InsideRect_ToData(view IView, rect foundation.Rect, data foundation.IMutableData) PrintOperation {
	rv := ffi.CallMethod[PrintOperation](pc, "PDFOperationWithView:insideRect:toData:", view, rect, data)
	return rv
}

func (pc _PrintOperationClass) PDFOperationWithView_InsideRect_ToData_PrintInfo(view IView, rect foundation.Rect, data foundation.IMutableData, printInfo IPrintInfo) PrintOperation {
	rv := ffi.CallMethod[PrintOperation](pc, "PDFOperationWithView:insideRect:toData:printInfo:", view, rect, data, printInfo)
	return rv
}

func (pc _PrintOperationClass) PDFOperationWithView_InsideRect_ToPath_PrintInfo(view IView, rect foundation.Rect, path string, printInfo IPrintInfo) PrintOperation {
	rv := ffi.CallMethod[PrintOperation](pc, "PDFOperationWithView:insideRect:toPath:printInfo:", view, rect, path, printInfo)
	return rv
}

func (pc _PrintOperationClass) PrintOperationWithView(view IView) PrintOperation {
	rv := ffi.CallMethod[PrintOperation](pc, "printOperationWithView:", view)
	return rv
}

func (pc _PrintOperationClass) PrintOperationWithView_PrintInfo(view IView, printInfo IPrintInfo) PrintOperation {
	rv := ffi.CallMethod[PrintOperation](pc, "printOperationWithView:printInfo:", view, printInfo)
	return rv
}

func (p_ PrintOperation) RunOperation() bool {
	rv := ffi.CallMethod[bool](p_, "runOperation")
	return rv
}

func (p_ PrintOperation) RunOperationModalForWindow_Delegate_DidRunSelector_ContextInfo(docWindow IWindow, delegate objc.IObject, didRunSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](p_, "runOperationModalForWindow:delegate:didRunSelector:contextInfo:", docWindow, delegate, didRunSelector, contextInfo)
}

func (p_ PrintOperation) CleanUpOperation() {
	ffi.CallMethod[ffi.Void](p_, "cleanUpOperation")
}

func (p_ PrintOperation) DeliverResult() bool {
	rv := ffi.CallMethod[bool](p_, "deliverResult")
	return rv
}

func (p_ PrintOperation) CreateContext() GraphicsContext {
	rv := ffi.CallMethod[GraphicsContext](p_, "createContext")
	return rv
}

func (p_ PrintOperation) DestroyContext() {
	ffi.CallMethod[ffi.Void](p_, "destroyContext")
}

// deprecated
func (p_ PrintOperation) JobStyleHint() string {
	rv := ffi.CallMethod[string](p_, "jobStyleHint")
	return rv
}

// deprecated
func (p_ PrintOperation) SetJobStyleHint(hint string) {
	ffi.CallMethod[ffi.Void](p_, "setJobStyleHint:", hint)
}

// deprecated
func (p_ PrintOperation) AccessoryView() View {
	rv := ffi.CallMethod[View](p_, "accessoryView")
	return rv
}

// deprecated
func (p_ PrintOperation) SetAccessoryView(view IView) {
	ffi.CallMethod[ffi.Void](p_, "setAccessoryView:", view)
}

// deprecated
func (p_ PrintOperation) ShowPanels() bool {
	rv := ffi.CallMethod[bool](p_, "showPanels")
	return rv
}

// deprecated
func (p_ PrintOperation) SetShowPanels(flag bool) {
	ffi.CallMethod[ffi.Void](p_, "setShowPanels:", flag)
}

func (pc _PrintOperationClass) CurrentOperation() PrintOperation {
	rv := ffi.CallMethod[PrintOperation](pc, "currentOperation")
	return rv
}

func (pc _PrintOperationClass) SetCurrentOperation(value IPrintOperation) {
	ffi.CallMethod[ffi.Void](pc, "setCurrentOperation:", value)
}

func (p_ PrintOperation) IsCopyingOperation() bool {
	rv := ffi.CallMethod[bool](p_, "isCopyingOperation")
	return rv
}

func (p_ PrintOperation) PrintInfo() PrintInfo {
	rv := ffi.CallMethod[PrintInfo](p_, "printInfo")
	return rv
}

func (p_ PrintOperation) SetPrintInfo(value IPrintInfo) {
	ffi.CallMethod[ffi.Void](p_, "setPrintInfo:", value)
}

func (p_ PrintOperation) View() View {
	rv := ffi.CallMethod[View](p_, "view")
	return rv
}

func (p_ PrintOperation) PreferredRenderingQuality() PrintRenderingQuality {
	rv := ffi.CallMethod[PrintRenderingQuality](p_, "preferredRenderingQuality")
	return rv
}

func (p_ PrintOperation) ShowsPrintPanel() bool {
	rv := ffi.CallMethod[bool](p_, "showsPrintPanel")
	return rv
}

func (p_ PrintOperation) SetShowsPrintPanel(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setShowsPrintPanel:", value)
}

func (p_ PrintOperation) ShowsProgressPanel() bool {
	rv := ffi.CallMethod[bool](p_, "showsProgressPanel")
	return rv
}

func (p_ PrintOperation) SetShowsProgressPanel(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setShowsProgressPanel:", value)
}

func (p_ PrintOperation) JobTitle() string {
	rv := ffi.CallMethod[string](p_, "jobTitle")
	return rv
}

func (p_ PrintOperation) SetJobTitle(value string) {
	ffi.CallMethod[ffi.Void](p_, "setJobTitle:", value)
}

func (p_ PrintOperation) PrintPanel() PrintPanel {
	rv := ffi.CallMethod[PrintPanel](p_, "printPanel")
	return rv
}

func (p_ PrintOperation) SetPrintPanel(value IPrintPanel) {
	ffi.CallMethod[ffi.Void](p_, "setPrintPanel:", value)
}

func (p_ PrintOperation) PDFPanel() PDFPanel {
	rv := ffi.CallMethod[PDFPanel](p_, "PDFPanel")
	return rv
}

func (p_ PrintOperation) SetPDFPanel(value IPDFPanel) {
	ffi.CallMethod[ffi.Void](p_, "setPDFPanel:", value)
}

func (p_ PrintOperation) Context() GraphicsContext {
	rv := ffi.CallMethod[GraphicsContext](p_, "context")
	return rv
}

func (p_ PrintOperation) CurrentPage() int {
	rv := ffi.CallMethod[int](p_, "currentPage")
	return rv
}

func (p_ PrintOperation) PageRange() foundation.Range {
	rv := ffi.CallMethod[foundation.Range](p_, "pageRange")
	return rv
}

func (p_ PrintOperation) PageOrder() PrintingPageOrder {
	rv := ffi.CallMethod[PrintingPageOrder](p_, "pageOrder")
	return rv
}

func (p_ PrintOperation) SetPageOrder(value PrintingPageOrder) {
	ffi.CallMethod[ffi.Void](p_, "setPageOrder:", value)
}

func (p_ PrintOperation) CanSpawnSeparateThread() bool {
	rv := ffi.CallMethod[bool](p_, "canSpawnSeparateThread")
	return rv
}

func (p_ PrintOperation) SetCanSpawnSeparateThread(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setCanSpawnSeparateThread:", value)
}
