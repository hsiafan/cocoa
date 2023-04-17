// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var TextAttachmentCellClass = _TextAttachmentCellClass{objc.GetClass("NSTextAttachmentCell")}

type _TextAttachmentCellClass struct {
	objc.Class
}

type ITextAttachmentCell interface {
	ICell
}

type TextAttachmentCell struct {
	Cell
}

func MakeTextAttachmentCell(ptr unsafe.Pointer) TextAttachmentCell {
	return TextAttachmentCell{
		Cell: MakeCell(ptr),
	}
}

func (t_ TextAttachmentCell) InitImageCell(image IImage) TextAttachmentCell {
	rv := ffi.CallMethod[TextAttachmentCell](t_, "initImageCell:", image)
	return rv
}

func (t_ TextAttachmentCell) InitTextCell(string_ string) TextAttachmentCell {
	rv := ffi.CallMethod[TextAttachmentCell](t_, "initTextCell:", string_)
	return rv
}

func (t_ TextAttachmentCell) Init() TextAttachmentCell {
	rv := ffi.CallMethod[TextAttachmentCell](t_, "init")
	return rv
}

func (tc _TextAttachmentCellClass) Alloc() TextAttachmentCell {
	rv := ffi.CallMethod[TextAttachmentCell](tc, "alloc")
	return rv
}

func (tc _TextAttachmentCellClass) New() TextAttachmentCell {
	rv := ffi.CallMethod[TextAttachmentCell](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextAttachmentCell() TextAttachmentCell {
	return TextAttachmentCellClass.New()
}
