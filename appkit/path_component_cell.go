// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var PathComponentCellClass = _PathComponentCellClass{objc.GetClass("NSPathComponentCell")}

type _PathComponentCellClass struct {
	objc.Class
}

type IPathComponentCell interface {
	ITextFieldCell
	URL() foundation.URL
	SetURL(value foundation.IURL)
}

type PathComponentCell struct {
	TextFieldCell
}

func MakePathComponentCell(ptr unsafe.Pointer) PathComponentCell {
	return PathComponentCell{
		TextFieldCell: MakeTextFieldCell(ptr),
	}
}

func (p_ PathComponentCell) InitTextCell(string_ string) PathComponentCell {
	rv := objc.CallMethod[PathComponentCell](p_, objc.SEL("initTextCell:"), string_)
	return rv
}

func (p_ PathComponentCell) InitImageCell(image IImage) PathComponentCell {
	rv := objc.CallMethod[PathComponentCell](p_, objc.SEL("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func (p_ PathComponentCell) Init() PathComponentCell {
	rv := objc.CallMethod[PathComponentCell](p_, objc.SEL("init"))
	return rv
}

func (pc _PathComponentCellClass) Alloc() PathComponentCell {
	rv := objc.CallMethod[PathComponentCell](pc, objc.SEL("alloc"))
	return rv
}

func (pc _PathComponentCellClass) New() PathComponentCell {
	rv := objc.CallMethod[PathComponentCell](pc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewPathComponentCell() PathComponentCell {
	return PathComponentCellClass.New()
}

func (p_ PathComponentCell) URL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](p_, objc.SEL("URL"))
	return rv
}

func (p_ PathComponentCell) SetURL(value foundation.IURL) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setURL:"), objc.ExtractPtr(value))
}
