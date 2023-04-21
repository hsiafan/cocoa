// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[TextTable](t_, "init")
	return rv
}

func (tc _TextTableClass) Alloc() TextTable {
	rv := objc.CallMethod[TextTable](tc, "alloc")
	return rv
}

func (tc _TextTableClass) New() TextTable {
	rv := objc.CallMethod[TextTable](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextTable() TextTable {
	return TextTableClass.New()
}

func (t_ TextTable) NumberOfColumns() uint {
	rv := objc.CallMethod[uint](t_, "numberOfColumns")
	return rv
}

func (t_ TextTable) SetNumberOfColumns(value uint) {
	objc.CallMethod[objc.Void](t_, "setNumberOfColumns:", value)
}

func (t_ TextTable) LayoutAlgorithm() TextTableLayoutAlgorithm {
	rv := objc.CallMethod[TextTableLayoutAlgorithm](t_, "layoutAlgorithm")
	return rv
}

func (t_ TextTable) SetLayoutAlgorithm(value TextTableLayoutAlgorithm) {
	objc.CallMethod[objc.Void](t_, "setLayoutAlgorithm:", value)
}

func (t_ TextTable) CollapsesBorders() bool {
	rv := objc.CallMethod[bool](t_, "collapsesBorders")
	return rv
}

func (t_ TextTable) SetCollapsesBorders(value bool) {
	objc.CallMethod[objc.Void](t_, "setCollapsesBorders:", value)
}

func (t_ TextTable) HidesEmptyCells() bool {
	rv := objc.CallMethod[bool](t_, "hidesEmptyCells")
	return rv
}

func (t_ TextTable) SetHidesEmptyCells(value bool) {
	objc.CallMethod[objc.Void](t_, "setHidesEmptyCells:", value)
}
