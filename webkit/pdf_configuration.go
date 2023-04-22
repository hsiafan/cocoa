// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
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
	rv := objc.CallMethod[PDFConfiguration](pc, objc.GetSelector("alloc"))
	return rv
}

func (pc _PDFConfigurationClass) New() PDFConfiguration {
	rv := objc.CallMethod[PDFConfiguration](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPDFConfiguration() PDFConfiguration {
	return PDFConfigurationClass.New()
}

func (p_ PDFConfiguration) Init() PDFConfiguration {
	rv := objc.CallMethod[PDFConfiguration](p_, objc.GetSelector("init"))
	return rv
}

func (p_ PDFConfiguration) Rect() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](p_, objc.GetSelector("rect"))
	return rv
}

func (p_ PDFConfiguration) SetRect(value coregraphics.Rect) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setRect:"), value)
}
