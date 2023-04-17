// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[PDFPanel](pc, "alloc")
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

func (p_ PDFPanel) Init() PDFPanel {
	rv := ffi.CallMethod[PDFPanel](p_, "init")
	return rv
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
