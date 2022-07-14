// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var PrinterClass = _PrinterClass{objc.GetClass("NSPrinter")}

type _PrinterClass struct {
	objc.Class
}

type IPrinter interface {
	objc.IObject
	PageSizeForPaper(paperName PrinterPaperName) foundation.Size
	// deprecated
	IsKey_InTable(key string, table string) bool
	// deprecated
	StringForKey_InTable(key string, table string) string
	// deprecated
	StringListForKey_InTable(key string, table string) []objc.Object
	// deprecated
	BooleanForKey_InTable(key string, table string) bool
	// deprecated
	FloatForKey_InTable(key string, table string) float32
	// deprecated
	IntForKey_InTable(key string, table string) int32
	// deprecated
	RectForKey_InTable(key string, table string) foundation.Rect
	// deprecated
	SizeForKey_InTable(key string, table string) foundation.Size
	// deprecated
	StatusForTable(tableName string) PrinterTableStatus
	// deprecated
	AcceptsBinary() bool
	// deprecated
	Domain() string
	// deprecated
	Host() string
	// deprecated
	ImageRectForPaper(paperName string) foundation.Rect
	// deprecated
	IsColor() bool
	// deprecated
	IsFontAvailable(faceName string) bool
	// deprecated
	IsOutputStackInReverseOrder() bool
	// deprecated
	Note() string
	Name() string
	Type() PrinterTypeName
	LanguageLevel() int
	DeviceDescription() map[DeviceDescriptionKey]objc.Object
}

type Printer struct {
	objc.Object
}

func MakePrinter(ptr unsafe.Pointer) Printer {
	return Printer{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PrinterClass) Alloc() Printer {
	rv := ffi.CallMethod[Printer](pc, "alloc")
	return rv
}

func (p_ Printer) Init() Printer {
	rv := ffi.CallMethod[Printer](p_, "init")
	rv.Autorelease()
	return rv
}

func (pc _PrinterClass) New() Printer {
	rv := ffi.CallMethod[Printer](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPrinter() Printer {
	return PrinterClass.New()
}

func (pc _PrinterClass) PrinterWithName(name string) Printer {
	rv := ffi.CallMethod[Printer](pc, "printerWithName:", name)
	return rv
}

func (pc _PrinterClass) PrinterWithType(_type PrinterTypeName) Printer {
	rv := ffi.CallMethod[Printer](pc, "printerWithType:", _type)
	return rv
}

func (p_ Printer) PageSizeForPaper(paperName PrinterPaperName) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](p_, "pageSizeForPaper:", paperName)
	return rv
}

// deprecated
func (p_ Printer) IsKey_InTable(key string, table string) bool {
	rv := ffi.CallMethod[bool](p_, "isKey:inTable:", key, table)
	return rv
}

// deprecated
func (p_ Printer) StringForKey_InTable(key string, table string) string {
	rv := ffi.CallMethod[string](p_, "stringForKey:inTable:", key, table)
	return rv
}

// deprecated
func (p_ Printer) StringListForKey_InTable(key string, table string) []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](p_, "stringListForKey:inTable:", key, table)
	return rv
}

// deprecated
func (p_ Printer) BooleanForKey_InTable(key string, table string) bool {
	rv := ffi.CallMethod[bool](p_, "booleanForKey:inTable:", key, table)
	return rv
}

// deprecated
func (p_ Printer) FloatForKey_InTable(key string, table string) float32 {
	rv := ffi.CallMethod[float32](p_, "floatForKey:inTable:", key, table)
	return rv
}

// deprecated
func (p_ Printer) IntForKey_InTable(key string, table string) int32 {
	rv := ffi.CallMethod[int32](p_, "intForKey:inTable:", key, table)
	return rv
}

// deprecated
func (p_ Printer) RectForKey_InTable(key string, table string) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](p_, "rectForKey:inTable:", key, table)
	return rv
}

// deprecated
func (p_ Printer) SizeForKey_InTable(key string, table string) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](p_, "sizeForKey:inTable:", key, table)
	return rv
}

// deprecated
func (p_ Printer) StatusForTable(tableName string) PrinterTableStatus {
	rv := ffi.CallMethod[PrinterTableStatus](p_, "statusForTable:", tableName)
	return rv
}

// deprecated
func (pc _PrinterClass) PrinterWithName_Domain_IncludeUnavailable(name string, domain string, flag bool) Printer {
	rv := ffi.CallMethod[Printer](pc, "printerWithName:domain:includeUnavailable:", name, domain, flag)
	return rv
}

// deprecated
func (p_ Printer) AcceptsBinary() bool {
	rv := ffi.CallMethod[bool](p_, "acceptsBinary")
	return rv
}

// deprecated
func (p_ Printer) Domain() string {
	rv := ffi.CallMethod[string](p_, "domain")
	return rv
}

// deprecated
func (p_ Printer) Host() string {
	rv := ffi.CallMethod[string](p_, "host")
	return rv
}

// deprecated
func (p_ Printer) ImageRectForPaper(paperName string) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](p_, "imageRectForPaper:", paperName)
	return rv
}

// deprecated
func (p_ Printer) IsColor() bool {
	rv := ffi.CallMethod[bool](p_, "isColor")
	return rv
}

// deprecated
func (p_ Printer) IsFontAvailable(faceName string) bool {
	rv := ffi.CallMethod[bool](p_, "isFontAvailable:", faceName)
	return rv
}

// deprecated
func (p_ Printer) IsOutputStackInReverseOrder() bool {
	rv := ffi.CallMethod[bool](p_, "isOutputStackInReverseOrder")
	return rv
}

// deprecated
func (p_ Printer) Note() string {
	rv := ffi.CallMethod[string](p_, "note")
	return rv
}

func (pc _PrinterClass) PrinterNames() []string {
	rv := ffi.CallMethod[[]string](pc, "printerNames")
	return rv
}

func (pc _PrinterClass) PrinterTypes() []PrinterTypeName {
	rv := ffi.CallMethod[[]PrinterTypeName](pc, "printerTypes")
	return rv
}

func (p_ Printer) Name() string {
	rv := ffi.CallMethod[string](p_, "name")
	return rv
}

func (p_ Printer) Type() PrinterTypeName {
	rv := ffi.CallMethod[PrinterTypeName](p_, "type")
	return rv
}

func (p_ Printer) LanguageLevel() int {
	rv := ffi.CallMethod[int](p_, "languageLevel")
	return rv
}

func (p_ Printer) DeviceDescription() map[DeviceDescriptionKey]objc.Object {
	rv := ffi.CallMethod[map[DeviceDescriptionKey]objc.Object](p_, "deviceDescription")
	return rv
}

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

func (p_ PrintPanel) Init() PrintPanel {
	rv := ffi.CallMethod[PrintPanel](p_, "init")
	rv.Autorelease()
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

var PDFInfoClass = _PDFInfoClass{objc.GetClass("NSPDFInfo")}

type _PDFInfoClass struct {
	objc.Class
}

type IPDFInfo interface {
	objc.IObject
	URL() foundation.URL
	SetURL(value foundation.IURL)
	IsFileExtensionHidden() bool
	SetFileExtensionHidden(value bool)
	TagNames() []string
	SetTagNames(value []string)
	Orientation() PaperOrientation
	SetOrientation(value PaperOrientation)
	PaperSize() foundation.Size
	SetPaperSize(value foundation.Size)
	Attributes() foundation.MutableDictionary
}

type PDFInfo struct {
	objc.Object
}

func MakePDFInfo(ptr unsafe.Pointer) PDFInfo {
	return PDFInfo{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PDFInfoClass) Alloc() PDFInfo {
	rv := ffi.CallMethod[PDFInfo](pc, "alloc")
	return rv
}

func (p_ PDFInfo) Init() PDFInfo {
	rv := ffi.CallMethod[PDFInfo](p_, "init")
	rv.Autorelease()
	return rv
}

func (pc _PDFInfoClass) New() PDFInfo {
	rv := ffi.CallMethod[PDFInfo](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPDFInfo() PDFInfo {
	return PDFInfoClass.New()
}

func (p_ PDFInfo) URL() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](p_, "URL")
	return rv
}

func (p_ PDFInfo) SetURL(value foundation.IURL) {
	ffi.CallMethod[ffi.Void](p_, "setURL:", value)
}

func (p_ PDFInfo) IsFileExtensionHidden() bool {
	rv := ffi.CallMethod[bool](p_, "isFileExtensionHidden")
	return rv
}

func (p_ PDFInfo) SetFileExtensionHidden(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setFileExtensionHidden:", value)
}

func (p_ PDFInfo) TagNames() []string {
	rv := ffi.CallMethod[[]string](p_, "tagNames")
	return rv
}

func (p_ PDFInfo) SetTagNames(value []string) {
	ffi.CallMethod[ffi.Void](p_, "setTagNames:", value)
}

func (p_ PDFInfo) Orientation() PaperOrientation {
	rv := ffi.CallMethod[PaperOrientation](p_, "orientation")
	return rv
}

func (p_ PDFInfo) SetOrientation(value PaperOrientation) {
	ffi.CallMethod[ffi.Void](p_, "setOrientation:", value)
}

func (p_ PDFInfo) PaperSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](p_, "paperSize")
	return rv
}

func (p_ PDFInfo) SetPaperSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](p_, "setPaperSize:", value)
}

func (p_ PDFInfo) Attributes() foundation.MutableDictionary {
	rv := ffi.CallMethod[foundation.MutableDictionary](p_, "attributes")
	return rv
}

var PDFPanelClass = _PDFPanelClass{objc.GetClass("NSPDFPanel")}

type _PDFPanelClass struct {
	objc.Class
}

type IPDFPanel interface {
	objc.IObject
	BeginSheetWithPDFInfo_ModalForWindow_CompletionHandler(pdfInfo IPDFInfo, docWindow IWindow, completionHandler func(param1 int))
	AccessoryController() ViewController
	SetAccessoryController(value IViewController)
	Options() PDFPanelOptions
	SetOptions(value PDFPanelOptions)
	DefaultFileName() string
	SetDefaultFileName(value string)
}

type PDFPanel struct {
	objc.Object
}

func MakePDFPanel(ptr unsafe.Pointer) PDFPanel {
	return PDFPanel{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PDFPanelClass) Alloc() PDFPanel {
	rv := ffi.CallMethod[PDFPanel](pc, "alloc")
	return rv
}

func (p_ PDFPanel) Init() PDFPanel {
	rv := ffi.CallMethod[PDFPanel](p_, "init")
	rv.Autorelease()
	return rv
}

func (pc _PDFPanelClass) New() PDFPanel {
	rv := ffi.CallMethod[PDFPanel](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPDFPanel() PDFPanel {
	return PDFPanelClass.New()
}

func (pc _PDFPanelClass) Panel() PDFPanel {
	rv := ffi.CallMethod[PDFPanel](pc, "panel")
	return rv
}

func (p_ PDFPanel) BeginSheetWithPDFInfo_ModalForWindow_CompletionHandler(pdfInfo IPDFInfo, docWindow IWindow, completionHandler func(param1 int)) {
	ffi.CallMethod[ffi.Void](p_, "beginSheetWithPDFInfo:modalForWindow:completionHandler:", pdfInfo, docWindow, completionHandler)
}

func (p_ PDFPanel) AccessoryController() ViewController {
	rv := ffi.CallMethod[ViewController](p_, "accessoryController")
	return rv
}

func (p_ PDFPanel) SetAccessoryController(value IViewController) {
	ffi.CallMethod[ffi.Void](p_, "setAccessoryController:", value)
}

func (p_ PDFPanel) Options() PDFPanelOptions {
	rv := ffi.CallMethod[PDFPanelOptions](p_, "options")
	return rv
}

func (p_ PDFPanel) SetOptions(value PDFPanelOptions) {
	ffi.CallMethod[ffi.Void](p_, "setOptions:", value)
}

func (p_ PDFPanel) DefaultFileName() string {
	rv := ffi.CallMethod[string](p_, "defaultFileName")
	return rv
}

func (p_ PDFPanel) SetDefaultFileName(value string) {
	ffi.CallMethod[ffi.Void](p_, "setDefaultFileName:", value)
}

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
	rv := ffi.CallMethod[PrintInfo](p_, "initWithDictionary:", attributes)
	rv.Autorelease()
	return rv
}

func (p_ PrintInfo) Init() PrintInfo {
	rv := ffi.CallMethod[PrintInfo](p_, "init")
	rv.Autorelease()
	return rv
}

func (pc _PrintInfoClass) Alloc() PrintInfo {
	rv := ffi.CallMethod[PrintInfo](pc, "alloc")
	return rv
}

func (pc _PrintInfoClass) New() PrintInfo {
	rv := ffi.CallMethod[PrintInfo](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPrintInfo() PrintInfo {
	return PrintInfoClass.New()
}

func (p_ PrintInfo) SetUpPrintOperationDefaultValues() {
	ffi.CallMethod[ffi.Void](p_, "setUpPrintOperationDefaultValues")
}

func (p_ PrintInfo) Dictionary() foundation.MutableDictionary {
	rv := ffi.CallMethod[foundation.MutableDictionary](p_, "dictionary")
	return rv
}

func (p_ PrintInfo) PMPrintSession() unsafe.Pointer {
	rv := ffi.CallMethod[unsafe.Pointer](p_, "PMPrintSession")
	return rv
}

func (p_ PrintInfo) PMPageFormat() unsafe.Pointer {
	rv := ffi.CallMethod[unsafe.Pointer](p_, "PMPageFormat")
	return rv
}

func (p_ PrintInfo) PMPrintSettings() unsafe.Pointer {
	rv := ffi.CallMethod[unsafe.Pointer](p_, "PMPrintSettings")
	return rv
}

func (p_ PrintInfo) UpdateFromPMPageFormat() {
	ffi.CallMethod[ffi.Void](p_, "updateFromPMPageFormat")
}

func (p_ PrintInfo) UpdateFromPMPrintSettings() {
	ffi.CallMethod[ffi.Void](p_, "updateFromPMPrintSettings")
}

func (p_ PrintInfo) TakeSettingsFromPDFInfo(inPDFInfo IPDFInfo) {
	ffi.CallMethod[ffi.Void](p_, "takeSettingsFromPDFInfo:", inPDFInfo)
}

// deprecated
func (pc _PrintInfoClass) SetDefaultPrinter(printer IPrinter) {
	ffi.CallMethod[ffi.Void](pc, "setDefaultPrinter:", printer)
}

// deprecated
func (pc _PrintInfoClass) SizeForPaperName(name PrinterPaperName) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](pc, "sizeForPaperName:", name)
	return rv
}

func (pc _PrintInfoClass) SharedPrintInfo() PrintInfo {
	rv := ffi.CallMethod[PrintInfo](pc, "sharedPrintInfo")
	return rv
}

func (pc _PrintInfoClass) SetSharedPrintInfo(value IPrintInfo) {
	ffi.CallMethod[ffi.Void](pc, "setSharedPrintInfo:", value)
}

func (p_ PrintInfo) PaperSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](p_, "paperSize")
	return rv
}

func (p_ PrintInfo) SetPaperSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](p_, "setPaperSize:", value)
}

func (p_ PrintInfo) TopMargin() float64 {
	rv := ffi.CallMethod[float64](p_, "topMargin")
	return rv
}

func (p_ PrintInfo) SetTopMargin(value float64) {
	ffi.CallMethod[ffi.Void](p_, "setTopMargin:", value)
}

func (p_ PrintInfo) BottomMargin() float64 {
	rv := ffi.CallMethod[float64](p_, "bottomMargin")
	return rv
}

func (p_ PrintInfo) SetBottomMargin(value float64) {
	ffi.CallMethod[ffi.Void](p_, "setBottomMargin:", value)
}

func (p_ PrintInfo) LeftMargin() float64 {
	rv := ffi.CallMethod[float64](p_, "leftMargin")
	return rv
}

func (p_ PrintInfo) SetLeftMargin(value float64) {
	ffi.CallMethod[ffi.Void](p_, "setLeftMargin:", value)
}

func (p_ PrintInfo) RightMargin() float64 {
	rv := ffi.CallMethod[float64](p_, "rightMargin")
	return rv
}

func (p_ PrintInfo) SetRightMargin(value float64) {
	ffi.CallMethod[ffi.Void](p_, "setRightMargin:", value)
}

func (p_ PrintInfo) ImageablePageBounds() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](p_, "imageablePageBounds")
	return rv
}

func (p_ PrintInfo) Orientation() PaperOrientation {
	rv := ffi.CallMethod[PaperOrientation](p_, "orientation")
	return rv
}

func (p_ PrintInfo) SetOrientation(value PaperOrientation) {
	ffi.CallMethod[ffi.Void](p_, "setOrientation:", value)
}

func (p_ PrintInfo) PaperName() PrinterPaperName {
	rv := ffi.CallMethod[PrinterPaperName](p_, "paperName")
	return rv
}

func (p_ PrintInfo) SetPaperName(value PrinterPaperName) {
	ffi.CallMethod[ffi.Void](p_, "setPaperName:", value)
}

func (p_ PrintInfo) LocalizedPaperName() string {
	rv := ffi.CallMethod[string](p_, "localizedPaperName")
	return rv
}

func (p_ PrintInfo) HorizontalPagination() PrintingPaginationMode {
	rv := ffi.CallMethod[PrintingPaginationMode](p_, "horizontalPagination")
	return rv
}

func (p_ PrintInfo) SetHorizontalPagination(value PrintingPaginationMode) {
	ffi.CallMethod[ffi.Void](p_, "setHorizontalPagination:", value)
}

func (p_ PrintInfo) VerticalPagination() PrintingPaginationMode {
	rv := ffi.CallMethod[PrintingPaginationMode](p_, "verticalPagination")
	return rv
}

func (p_ PrintInfo) SetVerticalPagination(value PrintingPaginationMode) {
	ffi.CallMethod[ffi.Void](p_, "setVerticalPagination:", value)
}

func (p_ PrintInfo) IsHorizontallyCentered() bool {
	rv := ffi.CallMethod[bool](p_, "isHorizontallyCentered")
	return rv
}

func (p_ PrintInfo) SetHorizontallyCentered(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setHorizontallyCentered:", value)
}

func (p_ PrintInfo) IsVerticallyCentered() bool {
	rv := ffi.CallMethod[bool](p_, "isVerticallyCentered")
	return rv
}

func (p_ PrintInfo) SetVerticallyCentered(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setVerticallyCentered:", value)
}

func (p_ PrintInfo) Printer() Printer {
	rv := ffi.CallMethod[Printer](p_, "printer")
	return rv
}

func (p_ PrintInfo) SetPrinter(value IPrinter) {
	ffi.CallMethod[ffi.Void](p_, "setPrinter:", value)
}

func (p_ PrintInfo) JobDisposition() PrintJobDispositionValue {
	rv := ffi.CallMethod[PrintJobDispositionValue](p_, "jobDisposition")
	return rv
}

func (p_ PrintInfo) SetJobDisposition(value PrintJobDispositionValue) {
	ffi.CallMethod[ffi.Void](p_, "setJobDisposition:", value)
}

func (p_ PrintInfo) IsSelectionOnly() bool {
	rv := ffi.CallMethod[bool](p_, "isSelectionOnly")
	return rv
}

func (p_ PrintInfo) SetSelectionOnly(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setSelectionOnly:", value)
}

func (p_ PrintInfo) ScalingFactor() float64 {
	rv := ffi.CallMethod[float64](p_, "scalingFactor")
	return rv
}

func (p_ PrintInfo) SetScalingFactor(value float64) {
	ffi.CallMethod[ffi.Void](p_, "setScalingFactor:", value)
}

func (p_ PrintInfo) PrintSettings() foundation.MutableDictionary {
	rv := ffi.CallMethod[foundation.MutableDictionary](p_, "printSettings")
	return rv
}

func (pc _PrintInfoClass) DefaultPrinter() Printer {
	rv := ffi.CallMethod[Printer](pc, "defaultPrinter")
	return rv
}

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

func (p_ PrintOperation) Init() PrintOperation {
	rv := ffi.CallMethod[PrintOperation](p_, "init")
	rv.Autorelease()
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
	rv := ffi.CallMethod[PageLayout](pc, "alloc")
	return rv
}

func (p_ PageLayout) Init() PageLayout {
	rv := ffi.CallMethod[PageLayout](p_, "init")
	rv.Autorelease()
	return rv
}

func (pc _PageLayoutClass) New() PageLayout {
	rv := ffi.CallMethod[PageLayout](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPageLayout() PageLayout {
	return PageLayoutClass.New()
}

func (pc _PageLayoutClass) PageLayout() PageLayout {
	rv := ffi.CallMethod[PageLayout](pc, "pageLayout")
	return rv
}

func (p_ PageLayout) BeginSheetWithPrintInfo_ModalForWindow_Delegate_DidEndSelector_ContextInfo(printInfo IPrintInfo, docWindow IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](p_, "beginSheetWithPrintInfo:modalForWindow:delegate:didEndSelector:contextInfo:", printInfo, docWindow, delegate, didEndSelector, contextInfo)
}

func (p_ PageLayout) RunModal() int {
	rv := ffi.CallMethod[int](p_, "runModal")
	return rv
}

func (p_ PageLayout) RunModalWithPrintInfo(printInfo IPrintInfo) int {
	rv := ffi.CallMethod[int](p_, "runModalWithPrintInfo:", printInfo)
	return rv
}

func (p_ PageLayout) AddAccessoryController(accessoryController IViewController) {
	ffi.CallMethod[ffi.Void](p_, "addAccessoryController:", accessoryController)
}

func (p_ PageLayout) RemoveAccessoryController(accessoryController IViewController) {
	ffi.CallMethod[ffi.Void](p_, "removeAccessoryController:", accessoryController)
}

// deprecated
func (p_ PageLayout) AccessoryView() View {
	rv := ffi.CallMethod[View](p_, "accessoryView")
	return rv
}

// deprecated
func (p_ PageLayout) SetAccessoryView(accessoryView IView) {
	ffi.CallMethod[ffi.Void](p_, "setAccessoryView:", accessoryView)
}

// deprecated
func (p_ PageLayout) ReadPrintInfo() {
	ffi.CallMethod[ffi.Void](p_, "readPrintInfo")
}

// deprecated
func (p_ PageLayout) WritePrintInfo() {
	ffi.CallMethod[ffi.Void](p_, "writePrintInfo")
}

func (p_ PageLayout) AccessoryControllers() []ViewController {
	rv := ffi.CallMethod[[]ViewController](p_, "accessoryControllers")
	return rv
}

func (p_ PageLayout) PrintInfo() PrintInfo {
	rv := ffi.CallMethod[PrintInfo](p_, "printInfo")
	return rv
}
