// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[TextAttachmentCell](t_, objc.SEL("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func (t_ TextAttachmentCell) InitTextCell(string_ string) TextAttachmentCell {
	rv := objc.CallMethod[TextAttachmentCell](t_, objc.SEL("initTextCell:"), string_)
	return rv
}

func (t_ TextAttachmentCell) Init() TextAttachmentCell {
	rv := objc.CallMethod[TextAttachmentCell](t_, objc.SEL("init"))
	return rv
}

func (tc _TextAttachmentCellClass) Alloc() TextAttachmentCell {
	rv := objc.CallMethod[TextAttachmentCell](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TextAttachmentCellClass) New() TextAttachmentCell {
	rv := objc.CallMethod[TextAttachmentCell](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTextAttachmentCell() TextAttachmentCell {
	return TextAttachmentCellClass.New()
}
