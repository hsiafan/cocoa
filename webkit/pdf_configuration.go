// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var PDFConfigurationClass = _PDFConfigurationClass{objc.GetClass("WKPDFConfiguration")}

type _PDFConfigurationClass struct {
	objc.Class
}

type IPDFConfiguration interface {
	objc.IObject
	Rect() coregraphics.Rect
	SetRect(value coregraphics.Rect)
}

type PDFConfiguration struct {
	objc.Object
}

func MakePDFConfiguration(ptr unsafe.Pointer) PDFConfiguration {
	return PDFConfiguration{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PDFConfigurationClass) Alloc() PDFConfiguration {
	rv := ffi.CallMethod[PDFConfiguration](pc, "alloc")
	return rv
}

func (p_ PDFConfiguration) Init() PDFConfiguration {
	rv := ffi.CallMethod[PDFConfiguration](p_, "init")
	return rv
}

func (pc _PDFConfigurationClass) New() PDFConfiguration {
	rv := ffi.CallMethod[PDFConfiguration](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPDFConfiguration() PDFConfiguration {
	return PDFConfigurationClass.New()
}

func (p_ PDFConfiguration) Rect() coregraphics.Rect {
	rv := ffi.CallMethod[coregraphics.Rect](p_, "rect")
	return rv
}

func (p_ PDFConfiguration) SetRect(value coregraphics.Rect) {
	ffi.CallMethod[ffi.Void](p_, "setRect:", value)
}
