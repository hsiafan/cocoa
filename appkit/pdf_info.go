// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[PDFInfo](pc, "alloc")
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

func (p_ PDFInfo) Init() PDFInfo {
	rv := ffi.CallMethod[PDFInfo](p_, "init")
	return rv
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
