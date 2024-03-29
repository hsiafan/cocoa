// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

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
	rv := objc.CallMethod[PDFPanel](pc, objc.SEL("alloc"))
	return rv
}

func (pc _PDFPanelClass) New() PDFPanel {
	rv := objc.CallMethod[PDFPanel](pc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewPDFPanel() PDFPanel {
	return PDFPanelClass.New()
}

func (p_ PDFPanel) Init() PDFPanel {
	rv := objc.CallMethod[PDFPanel](p_, objc.SEL("init"))
	return rv
}

func (pc _PDFPanelClass) Panel() PDFPanel {
	rv := objc.CallMethod[PDFPanel](pc, objc.SEL("panel"))
	return rv
}

func (p_ PDFPanel) BeginSheetWithPDFInfo_ModalForWindow_CompletionHandler(pdfInfo IPDFInfo, docWindow IWindow, completionHandler func(param1 int)) {
	objc.CallMethod[objc.Void](p_, objc.SEL("beginSheetWithPDFInfo:modalForWindow:completionHandler:"), objc.ExtractPtr(pdfInfo), objc.ExtractPtr(docWindow), completionHandler)
}

func (p_ PDFPanel) AccessoryController() ViewController {
	rv := objc.CallMethod[ViewController](p_, objc.SEL("accessoryController"))
	return rv
}

func (p_ PDFPanel) SetAccessoryController(value IViewController) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setAccessoryController:"), objc.ExtractPtr(value))
}

func (p_ PDFPanel) Options() PDFPanelOptions {
	rv := objc.CallMethod[PDFPanelOptions](p_, objc.SEL("options"))
	return rv
}

func (p_ PDFPanel) SetOptions(value PDFPanelOptions) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setOptions:"), value)
}

func (p_ PDFPanel) DefaultFileName() string {
	rv := objc.CallMethod[string](p_, objc.SEL("defaultFileName"))
	return rv
}

func (p_ PDFPanel) SetDefaultFileName(value string) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setDefaultFileName:"), value)
}
