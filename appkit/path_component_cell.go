// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[PathComponentCell](p_, "initTextCell:", string_)
	return rv
}

func (p_ PathComponentCell) InitImageCell(image IImage) PathComponentCell {
	rv := ffi.CallMethod[PathComponentCell](p_, "initImageCell:", image)
	return rv
}

func (p_ PathComponentCell) Init() PathComponentCell {
	rv := ffi.CallMethod[PathComponentCell](p_, "init")
	return rv
}

func (pc _PathComponentCellClass) Alloc() PathComponentCell {
	rv := ffi.CallMethod[PathComponentCell](pc, "alloc")
	return rv
}

func (pc _PathComponentCellClass) New() PathComponentCell {
	rv := ffi.CallMethod[PathComponentCell](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPathComponentCell() PathComponentCell {
	return PathComponentCellClass.New()
}

func (p_ PathComponentCell) URL() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](p_, "URL")
	return rv
}

func (p_ PathComponentCell) SetURL(value foundation.IURL) {
	ffi.CallMethod[ffi.Void](p_, "setURL:", value)
}
