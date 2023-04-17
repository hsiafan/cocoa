// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var TextTableClass = _TextTableClass{objc.GetClass("NSTextTable")}

type _TextTableClass struct {
	objc.Class
}

type ITextTable interface {
	ITextBlock
	NumberOfColumns() uint
	SetNumberOfColumns(value uint)
	LayoutAlgorithm() TextTableLayoutAlgorithm
	SetLayoutAlgorithm(value TextTableLayoutAlgorithm)
	CollapsesBorders() bool
	SetCollapsesBorders(value bool)
	HidesEmptyCells() bool
	SetHidesEmptyCells(value bool)
}

type TextTable struct {
	TextBlock
}

func MakeTextTable(ptr unsafe.Pointer) TextTable {
	return TextTable{
		TextBlock: MakeTextBlock(ptr),
	}
}

func (t_ TextTable) Init() TextTable {
	rv := ffi.CallMethod[TextTable](t_, "init")
	return rv
}

func (tc _TextTableClass) Alloc() TextTable {
	rv := ffi.CallMethod[TextTable](tc, "alloc")
	return rv
}

func (tc _TextTableClass) New() TextTable {
	rv := ffi.CallMethod[TextTable](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextTable() TextTable {
	return TextTableClass.New()
}

func (t_ TextTable) NumberOfColumns() uint {
	rv := ffi.CallMethod[uint](t_, "numberOfColumns")
	return rv
}

func (t_ TextTable) SetNumberOfColumns(value uint) {
	ffi.CallMethod[ffi.Void](t_, "setNumberOfColumns:", value)
}

func (t_ TextTable) LayoutAlgorithm() TextTableLayoutAlgorithm {
	rv := ffi.CallMethod[TextTableLayoutAlgorithm](t_, "layoutAlgorithm")
	return rv
}

func (t_ TextTable) SetLayoutAlgorithm(value TextTableLayoutAlgorithm) {
	ffi.CallMethod[ffi.Void](t_, "setLayoutAlgorithm:", value)
}

func (t_ TextTable) CollapsesBorders() bool {
	rv := ffi.CallMethod[bool](t_, "collapsesBorders")
	return rv
}

func (t_ TextTable) SetCollapsesBorders(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setCollapsesBorders:", value)
}

func (t_ TextTable) HidesEmptyCells() bool {
	rv := ffi.CallMethod[bool](t_, "hidesEmptyCells")
	return rv
}

func (t_ TextTable) SetHidesEmptyCells(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setHidesEmptyCells:", value)
}
