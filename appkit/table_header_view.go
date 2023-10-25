// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TableHeaderViewClass = _TableHeaderViewClass{objc.GetClass("NSTableHeaderView")}

type _TableHeaderViewClass struct {
	objc.Class
}

type ITableHeaderView interface {
	IView
	ColumnAtPoint(point foundation.Point) int
	HeaderRectOfColumn(column int) foundation.Rect
	TableView() TableView
	SetTableView(value ITableView)
	DraggedColumn() int
	DraggedDistance() float64
	ResizedColumn() int
}

type TableHeaderView struct {
	View
}

func MakeTableHeaderView(ptr unsafe.Pointer) TableHeaderView {
	return TableHeaderView{
		View: MakeView(ptr),
	}
}

func (t_ TableHeaderView) InitWithFrame(frameRect foundation.Rect) TableHeaderView {
	rv := objc.CallMethod[TableHeaderView](t_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (t_ TableHeaderView) Init() TableHeaderView {
	rv := objc.CallMethod[TableHeaderView](t_, objc.SEL("init"))
	return rv
}

func (tc _TableHeaderViewClass) Alloc() TableHeaderView {
	rv := objc.CallMethod[TableHeaderView](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TableHeaderViewClass) New() TableHeaderView {
	rv := objc.CallMethod[TableHeaderView](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTableHeaderView() TableHeaderView {
	return TableHeaderViewClass.New()
}

func (t_ TableHeaderView) ColumnAtPoint(point foundation.Point) int {
	rv := objc.CallMethod[int](t_, objc.SEL("columnAtPoint:"), point)
	return rv
}

func (t_ TableHeaderView) HeaderRectOfColumn(column int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.SEL("headerRectOfColumn:"), column)
	return rv
}

// weak property
func (t_ TableHeaderView) TableView() TableView {
	rv := objc.CallMethod[TableView](t_, objc.SEL("tableView"))
	return rv
}

// weak property
func (t_ TableHeaderView) SetTableView(value ITableView) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTableView:"), objc.ExtractPtr(value))
}

func (t_ TableHeaderView) DraggedColumn() int {
	rv := objc.CallMethod[int](t_, objc.SEL("draggedColumn"))
	return rv
}

func (t_ TableHeaderView) DraggedDistance() float64 {
	rv := objc.CallMethod[float64](t_, objc.SEL("draggedDistance"))
	return rv
}

func (t_ TableHeaderView) ResizedColumn() int {
	rv := objc.CallMethod[int](t_, objc.SEL("resizedColumn"))
	return rv
}
