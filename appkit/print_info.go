// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var PrintInfoClass = _PrintInfoClass{objc.GetClass("NSPrintInfo")}

type _PrintInfoClass struct {
	objc.Class
}

type IPrintInfo interface {
	objc.IObject
	SetUpPrintOperationDefaultValues()
	Dictionary() foundation.MutableDictionary
	PMPrintSession() unsafe.Pointer
	PMPageFormat() unsafe.Pointer
	PMPrintSettings() unsafe.Pointer
	UpdateFromPMPageFormat()
	UpdateFromPMPrintSettings()
	TakeSettingsFromPDFInfo(inPDFInfo IPDFInfo)
	PaperSize() foundation.Size
	SetPaperSize(value foundation.Size)
	TopMargin() float64
	SetTopMargin(value float64)
	BottomMargin() float64
	SetBottomMargin(value float64)
	LeftMargin() float64
	SetLeftMargin(value float64)
	RightMargin() float64
	SetRightMargin(value float64)
	ImageablePageBounds() foundation.Rect
	Orientation() PaperOrientation
	SetOrientation(value PaperOrientation)
	PaperName() PrinterPaperName
	SetPaperName(value PrinterPaperName)
	LocalizedPaperName() string
	HorizontalPagination() PrintingPaginationMode
	SetHorizontalPagination(value PrintingPaginationMode)
	VerticalPagination() PrintingPaginationMode
	SetVerticalPagination(value PrintingPaginationMode)
	IsHorizontallyCentered() bool
	SetHorizontallyCentered(value bool)
	IsVerticallyCentered() bool
	SetVerticallyCentered(value bool)
	Printer() Printer
	SetPrinter(value IPrinter)
	JobDisposition() PrintJobDispositionValue
	SetJobDisposition(value PrintJobDispositionValue)
	IsSelectionOnly() bool
	SetSelectionOnly(value bool)
	ScalingFactor() float64
	SetScalingFactor(value float64)
	PrintSettings() foundation.MutableDictionary
}

type PrintInfo struct {
	objc.Object
}

func MakePrintInfo(ptr unsafe.Pointer) PrintInfo {
	return PrintInfo{
		Object: objc.MakeObject(ptr),
	}
}

func (p_ PrintInfo) InitWithDictionary(attributes map[PrintInfoAttributeKey]objc.IObject) PrintInfo {
	rv := objc.CallMethod[PrintInfo](p_, objc.SEL("initWithDictionary:"), attributes)
	return rv
}

func (p_ PrintInfo) Init() PrintInfo {
	rv := objc.CallMethod[PrintInfo](p_, objc.SEL("init"))
	return rv
}

func (pc _PrintInfoClass) Alloc() PrintInfo {
	rv := objc.CallMethod[PrintInfo](pc, objc.SEL("alloc"))
	return rv
}

func (pc _PrintInfoClass) New() PrintInfo {
	rv := objc.CallMethod[PrintInfo](pc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewPrintInfo() PrintInfo {
	return PrintInfoClass.New()
}

func (p_ PrintInfo) SetUpPrintOperationDefaultValues() {
	objc.CallMethod[objc.Void](p_, objc.SEL("setUpPrintOperationDefaultValues"))
}

func (p_ PrintInfo) Dictionary() foundation.MutableDictionary {
	rv := objc.CallMethod[foundation.MutableDictionary](p_, objc.SEL("dictionary"))
	return rv
}

func (p_ PrintInfo) PMPrintSession() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](p_, objc.SEL("PMPrintSession"))
	return rv
}

func (p_ PrintInfo) PMPageFormat() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](p_, objc.SEL("PMPageFormat"))
	return rv
}

func (p_ PrintInfo) PMPrintSettings() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](p_, objc.SEL("PMPrintSettings"))
	return rv
}

func (p_ PrintInfo) UpdateFromPMPageFormat() {
	objc.CallMethod[objc.Void](p_, objc.SEL("updateFromPMPageFormat"))
}

func (p_ PrintInfo) UpdateFromPMPrintSettings() {
	objc.CallMethod[objc.Void](p_, objc.SEL("updateFromPMPrintSettings"))
}

func (p_ PrintInfo) TakeSettingsFromPDFInfo(inPDFInfo IPDFInfo) {
	objc.CallMethod[objc.Void](p_, objc.SEL("takeSettingsFromPDFInfo:"), objc.ExtractPtr(inPDFInfo))
}

// deprecated
func (pc _PrintInfoClass) SetDefaultPrinter(printer IPrinter) {
	objc.CallMethod[objc.Void](pc, objc.SEL("setDefaultPrinter:"), objc.ExtractPtr(printer))
}

// deprecated
func (pc _PrintInfoClass) SizeForPaperName(name PrinterPaperName) foundation.Size {
	rv := objc.CallMethod[foundation.Size](pc, objc.SEL("sizeForPaperName:"), name)
	return rv
}

func (pc _PrintInfoClass) SharedPrintInfo() PrintInfo {
	rv := objc.CallMethod[PrintInfo](pc, objc.SEL("sharedPrintInfo"))
	return rv
}

func (pc _PrintInfoClass) SetSharedPrintInfo(value IPrintInfo) {
	objc.CallMethod[objc.Void](pc, objc.SEL("setSharedPrintInfo:"), objc.ExtractPtr(value))
}

func (p_ PrintInfo) PaperSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](p_, objc.SEL("paperSize"))
	return rv
}

func (p_ PrintInfo) SetPaperSize(value foundation.Size) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setPaperSize:"), value)
}

func (p_ PrintInfo) TopMargin() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("topMargin"))
	return rv
}

func (p_ PrintInfo) SetTopMargin(value float64) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setTopMargin:"), value)
}

func (p_ PrintInfo) BottomMargin() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("bottomMargin"))
	return rv
}

func (p_ PrintInfo) SetBottomMargin(value float64) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setBottomMargin:"), value)
}

func (p_ PrintInfo) LeftMargin() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("leftMargin"))
	return rv
}

func (p_ PrintInfo) SetLeftMargin(value float64) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setLeftMargin:"), value)
}

func (p_ PrintInfo) RightMargin() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("rightMargin"))
	return rv
}

func (p_ PrintInfo) SetRightMargin(value float64) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setRightMargin:"), value)
}

func (p_ PrintInfo) ImageablePageBounds() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](p_, objc.SEL("imageablePageBounds"))
	return rv
}

func (p_ PrintInfo) Orientation() PaperOrientation {
	rv := objc.CallMethod[PaperOrientation](p_, objc.SEL("orientation"))
	return rv
}

func (p_ PrintInfo) SetOrientation(value PaperOrientation) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setOrientation:"), value)
}

func (p_ PrintInfo) PaperName() PrinterPaperName {
	rv := objc.CallMethod[PrinterPaperName](p_, objc.SEL("paperName"))
	return rv
}

func (p_ PrintInfo) SetPaperName(value PrinterPaperName) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setPaperName:"), value)
}

func (p_ PrintInfo) LocalizedPaperName() string {
	rv := objc.CallMethod[string](p_, objc.SEL("localizedPaperName"))
	return rv
}

func (p_ PrintInfo) HorizontalPagination() PrintingPaginationMode {
	rv := objc.CallMethod[PrintingPaginationMode](p_, objc.SEL("horizontalPagination"))
	return rv
}

func (p_ PrintInfo) SetHorizontalPagination(value PrintingPaginationMode) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setHorizontalPagination:"), value)
}

func (p_ PrintInfo) VerticalPagination() PrintingPaginationMode {
	rv := objc.CallMethod[PrintingPaginationMode](p_, objc.SEL("verticalPagination"))
	return rv
}

func (p_ PrintInfo) SetVerticalPagination(value PrintingPaginationMode) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setVerticalPagination:"), value)
}

func (p_ PrintInfo) IsHorizontallyCentered() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isHorizontallyCentered"))
	return rv
}

func (p_ PrintInfo) SetHorizontallyCentered(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setHorizontallyCentered:"), value)
}

func (p_ PrintInfo) IsVerticallyCentered() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isVerticallyCentered"))
	return rv
}

func (p_ PrintInfo) SetVerticallyCentered(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setVerticallyCentered:"), value)
}

func (p_ PrintInfo) Printer() Printer {
	rv := objc.CallMethod[Printer](p_, objc.SEL("printer"))
	return rv
}

func (p_ PrintInfo) SetPrinter(value IPrinter) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setPrinter:"), objc.ExtractPtr(value))
}

func (p_ PrintInfo) JobDisposition() PrintJobDispositionValue {
	rv := objc.CallMethod[PrintJobDispositionValue](p_, objc.SEL("jobDisposition"))
	return rv
}

func (p_ PrintInfo) SetJobDisposition(value PrintJobDispositionValue) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setJobDisposition:"), value)
}

func (p_ PrintInfo) IsSelectionOnly() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isSelectionOnly"))
	return rv
}

func (p_ PrintInfo) SetSelectionOnly(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setSelectionOnly:"), value)
}

func (p_ PrintInfo) ScalingFactor() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("scalingFactor"))
	return rv
}

func (p_ PrintInfo) SetScalingFactor(value float64) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setScalingFactor:"), value)
}

func (p_ PrintInfo) PrintSettings() foundation.MutableDictionary {
	rv := objc.CallMethod[foundation.MutableDictionary](p_, objc.SEL("printSettings"))
	return rv
}

func (pc _PrintInfoClass) DefaultPrinter() Printer {
	rv := objc.CallMethod[Printer](pc, objc.SEL("defaultPrinter"))
	return rv
}
