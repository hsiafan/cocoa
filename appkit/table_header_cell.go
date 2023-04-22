// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TableHeaderCellClass = _TableHeaderCellClass{objc.GetClass("NSTableHeaderCell")}

type _TableHeaderCellClass struct {
	objc.Class
}

type ITableHeaderCell interface {
	ITextFieldCell
	DrawSortIndicatorWithFrame_InView_Ascending_Priority(cellFrame foundation.Rect, controlView IView, ascending bool, priority int)
	SortIndicatorRectForBounds(rect foundation.Rect) foundation.Rect
}

type TableHeaderCell struct {
	TextFieldCell
}

func MakeTableHeaderCell(ptr unsafe.Pointer) TableHeaderCell {
	return TableHeaderCell{
		TextFieldCell: MakeTextFieldCell(ptr),
	}
}

func (t_ TableHeaderCell) InitTextCell(string_ string) TableHeaderCell {
	rv := objc.CallMethod[TableHeaderCell](t_, objc.GetSelector("initTextCell:"), string_)
	return rv
}

func (t_ TableHeaderCell) InitImageCell(image IImage) TableHeaderCell {
	rv := objc.CallMethod[TableHeaderCell](t_, objc.GetSelector("initImageCell:"), image)
	return rv
}

func (t_ TableHeaderCell) Init() TableHeaderCell {
	rv := objc.CallMethod[TableHeaderCell](t_, objc.GetSelector("init"))
	return rv
}

func (tc _TableHeaderCellClass) Alloc() TableHeaderCell {
	rv := objc.CallMethod[TableHeaderCell](tc, objc.GetSelector("alloc"))
	return rv
}

func (tc _TableHeaderCellClass) New() TableHeaderCell {
	rv := objc.CallMethod[TableHeaderCell](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTableHeaderCell() TableHeaderCell {
	return TableHeaderCellClass.New()
}

func (t_ TableHeaderCell) DrawSortIndicatorWithFrame_InView_Ascending_Priority(cellFrame foundation.Rect, controlView IView, ascending bool, priority int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("drawSortIndicatorWithFrame:inView:ascending:priority:"), cellFrame, controlView, ascending, priority)
}

func (t_ TableHeaderCell) SortIndicatorRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.GetSelector("sortIndicatorRectForBounds:"), rect)
	return rv
}
