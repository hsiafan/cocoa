// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

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
	rv := objc.CallMethod[PDFInfo](pc, "alloc")
	return rv
}

func (pc _PDFInfoClass) New() PDFInfo {
	rv := objc.CallMethod[PDFInfo](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPDFInfo() PDFInfo {
	return PDFInfoClass.New()
}

func (p_ PDFInfo) Init() PDFInfo {
	rv := objc.CallMethod[PDFInfo](p_, "init")
	return rv
}

func (p_ PDFInfo) URL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](p_, "URL")
	return rv
}

func (p_ PDFInfo) SetURL(value foundation.IURL) {
	objc.CallMethod[objc.Void](p_, "setURL:", value)
}

func (p_ PDFInfo) IsFileExtensionHidden() bool {
	rv := objc.CallMethod[bool](p_, "isFileExtensionHidden")
	return rv
}

func (p_ PDFInfo) SetFileExtensionHidden(value bool) {
	objc.CallMethod[objc.Void](p_, "setFileExtensionHidden:", value)
}

func (p_ PDFInfo) TagNames() []string {
	rv := objc.CallMethod[[]string](p_, "tagNames")
	return rv
}

func (p_ PDFInfo) SetTagNames(value []string) {
	objc.CallMethod[objc.Void](p_, "setTagNames:", value)
}

func (p_ PDFInfo) Orientation() PaperOrientation {
	rv := objc.CallMethod[PaperOrientation](p_, "orientation")
	return rv
}

func (p_ PDFInfo) SetOrientation(value PaperOrientation) {
	objc.CallMethod[objc.Void](p_, "setOrientation:", value)
}

func (p_ PDFInfo) PaperSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](p_, "paperSize")
	return rv
}

func (p_ PDFInfo) SetPaperSize(value foundation.Size) {
	objc.CallMethod[objc.Void](p_, "setPaperSize:", value)
}

func (p_ PDFInfo) Attributes() foundation.MutableDictionary {
	rv := objc.CallMethod[foundation.MutableDictionary](p_, "attributes")
	return rv
}
