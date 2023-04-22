// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[ActionCell](a_, objc.GetSelector("initImageCell:"), image)
	return rv
}

func (a_ ActionCell) InitTextCell(string_ string) ActionCell {
	rv := objc.CallMethod[ActionCell](a_, objc.GetSelector("initTextCell:"), string_)
	return rv
}

func (a_ ActionCell) Init() ActionCell {
	rv := objc.CallMethod[ActionCell](a_, objc.GetSelector("init"))
	return rv
}

func (ac _ActionCellClass) Alloc() ActionCell {
	rv := objc.CallMethod[ActionCell](ac, objc.GetSelector("alloc"))
	return rv
}

func (ac _ActionCellClass) New() ActionCell {
	rv := objc.CallMethod[ActionCell](ac, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewActionCell() ActionCell {
	return ActionCellClass.New()
}
