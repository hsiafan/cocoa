// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var ActionCellClass = _ActionCellClass{objc.GetClass("NSActionCell")}

type _ActionCellClass struct {
	objc.Class
}

type IActionCell interface {
	ICell
}

type ActionCell struct {
	Cell
}

func MakeActionCell(ptr unsafe.Pointer) ActionCell {
	return ActionCell{
		Cell: MakeCell(ptr),
	}
}

func (a_ ActionCell) InitImageCell(image IImage) ActionCell {
	rv := ffi.CallMethod[ActionCell](a_, "initImageCell:", image)
	return rv
}

func (a_ ActionCell) InitTextCell(string_ string) ActionCell {
	rv := ffi.CallMethod[ActionCell](a_, "initTextCell:", string_)
	return rv
}

func (a_ ActionCell) Init() ActionCell {
	rv := ffi.CallMethod[ActionCell](a_, "init")
	return rv
}

func (ac _ActionCellClass) Alloc() ActionCell {
	rv := ffi.CallMethod[ActionCell](ac, "alloc")
	return rv
}

func (ac _ActionCellClass) New() ActionCell {
	rv := ffi.CallMethod[ActionCell](ac, "new")
	rv.Autorelease()
	return rv
}

func NewActionCell() ActionCell {
	return ActionCellClass.New()
}