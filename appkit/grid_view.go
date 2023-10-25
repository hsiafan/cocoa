// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var GridViewClass = _GridViewClass{objc.GetClass("NSGridView")}

type _GridViewClass struct {
	objc.Class
}

type IGridView interface {
	IView
	IndexOfColumn(column IGridColumn) int
	RowAtIndex(index int) GridRow
	ColumnAtIndex(index int) GridColumn
	IndexOfRow(row IGridRow) int
	AddRowWithViews(views []IView) GridRow
	InsertRowAtIndex_WithViews(index int, views []IView) GridRow
	RemoveRowAtIndex(index int)
	MoveRowAtIndex_ToIndex(fromIndex int, toIndex int)
	AddColumnWithViews(views []IView) GridColumn
	InsertColumnAtIndex_WithViews(index int, views []IView) GridColumn
	RemoveColumnAtIndex(index int)
	MoveColumnAtIndex_ToIndex(fromIndex int, toIndex int)
	CellAtColumnIndex_RowIndex(columnIndex int, rowIndex int) GridCell
	CellForView(view IView) GridCell
	MergeCellsInHorizontalRange_VerticalRange(hRange foundation.Range, vRange foundation.Range)
	NumberOfRows() int
	NumberOfColumns() int
	ColumnSpacing() float64
	SetColumnSpacing(value float64)
	RowSpacing() float64
	SetRowSpacing(value float64)
	RowAlignment() GridRowAlignment
	SetRowAlignment(value GridRowAlignment)
	XPlacement() GridCellPlacement
	SetXPlacement(value GridCellPlacement)
	YPlacement() GridCellPlacement
	SetYPlacement(value GridCellPlacement)
}

type GridView struct {
	View
}

func MakeGridView(ptr unsafe.Pointer) GridView {
	return GridView{
		View: MakeView(ptr),
	}
}

func (gc _GridViewClass) GridViewWithNumberOfColumns_Rows(columnCount int, rowCount int) GridView {
	rv := objc.CallMethod[GridView](gc, objc.SEL("gridViewWithNumberOfColumns:rows:"), columnCount, rowCount)
	return rv
}

func (gc _GridViewClass) GridViewWithViews(rows [][]IView) GridView {
	rv := objc.CallMethod[GridView](gc, objc.SEL("gridViewWithViews:"), rows)
	return rv
}

func (g_ GridView) InitWithFrame(frameRect foundation.Rect) GridView {
	rv := objc.CallMethod[GridView](g_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (g_ GridView) Init() GridView {
	rv := objc.CallMethod[GridView](g_, objc.SEL("init"))
	return rv
}

func (gc _GridViewClass) Alloc() GridView {
	rv := objc.CallMethod[GridView](gc, objc.SEL("alloc"))
	return rv
}

func (gc _GridViewClass) New() GridView {
	rv := objc.CallMethod[GridView](gc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewGridView() GridView {
	return GridViewClass.New()
}

func (g_ GridView) IndexOfColumn(column IGridColumn) int {
	rv := objc.CallMethod[int](g_, objc.SEL("indexOfColumn:"), objc.ExtractPtr(column))
	return rv
}

func (g_ GridView) RowAtIndex(index int) GridRow {
	rv := objc.CallMethod[GridRow](g_, objc.SEL("rowAtIndex:"), index)
	return rv
}

func (g_ GridView) ColumnAtIndex(index int) GridColumn {
	rv := objc.CallMethod[GridColumn](g_, objc.SEL("columnAtIndex:"), index)
	return rv
}

func (g_ GridView) IndexOfRow(row IGridRow) int {
	rv := objc.CallMethod[int](g_, objc.SEL("indexOfRow:"), objc.ExtractPtr(row))
	return rv
}

func (g_ GridView) AddRowWithViews(views []IView) GridRow {
	rv := objc.CallMethod[GridRow](g_, objc.SEL("addRowWithViews:"), views)
	return rv
}

func (g_ GridView) InsertRowAtIndex_WithViews(index int, views []IView) GridRow {
	rv := objc.CallMethod[GridRow](g_, objc.SEL("insertRowAtIndex:withViews:"), index, views)
	return rv
}

func (g_ GridView) RemoveRowAtIndex(index int) {
	objc.CallMethod[objc.Void](g_, objc.SEL("removeRowAtIndex:"), index)
}

func (g_ GridView) MoveRowAtIndex_ToIndex(fromIndex int, toIndex int) {
	objc.CallMethod[objc.Void](g_, objc.SEL("moveRowAtIndex:toIndex:"), fromIndex, toIndex)
}

func (g_ GridView) AddColumnWithViews(views []IView) GridColumn {
	rv := objc.CallMethod[GridColumn](g_, objc.SEL("addColumnWithViews:"), views)
	return rv
}

func (g_ GridView) InsertColumnAtIndex_WithViews(index int, views []IView) GridColumn {
	rv := objc.CallMethod[GridColumn](g_, objc.SEL("insertColumnAtIndex:withViews:"), index, views)
	return rv
}

func (g_ GridView) RemoveColumnAtIndex(index int) {
	objc.CallMethod[objc.Void](g_, objc.SEL("removeColumnAtIndex:"), index)
}

func (g_ GridView) MoveColumnAtIndex_ToIndex(fromIndex int, toIndex int) {
	objc.CallMethod[objc.Void](g_, objc.SEL("moveColumnAtIndex:toIndex:"), fromIndex, toIndex)
}

func (g_ GridView) CellAtColumnIndex_RowIndex(columnIndex int, rowIndex int) GridCell {
	rv := objc.CallMethod[GridCell](g_, objc.SEL("cellAtColumnIndex:rowIndex:"), columnIndex, rowIndex)
	return rv
}

func (g_ GridView) CellForView(view IView) GridCell {
	rv := objc.CallMethod[GridCell](g_, objc.SEL("cellForView:"), objc.ExtractPtr(view))
	return rv
}

func (g_ GridView) MergeCellsInHorizontalRange_VerticalRange(hRange foundation.Range, vRange foundation.Range) {
	objc.CallMethod[objc.Void](g_, objc.SEL("mergeCellsInHorizontalRange:verticalRange:"), hRange, vRange)
}

func (g_ GridView) NumberOfRows() int {
	rv := objc.CallMethod[int](g_, objc.SEL("numberOfRows"))
	return rv
}

func (g_ GridView) NumberOfColumns() int {
	rv := objc.CallMethod[int](g_, objc.SEL("numberOfColumns"))
	return rv
}

func (g_ GridView) ColumnSpacing() float64 {
	rv := objc.CallMethod[float64](g_, objc.SEL("columnSpacing"))
	return rv
}

func (g_ GridView) SetColumnSpacing(value float64) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setColumnSpacing:"), value)
}

func (g_ GridView) RowSpacing() float64 {
	rv := objc.CallMethod[float64](g_, objc.SEL("rowSpacing"))
	return rv
}

func (g_ GridView) SetRowSpacing(value float64) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setRowSpacing:"), value)
}

func (g_ GridView) RowAlignment() GridRowAlignment {
	rv := objc.CallMethod[GridRowAlignment](g_, objc.SEL("rowAlignment"))
	return rv
}

func (g_ GridView) SetRowAlignment(value GridRowAlignment) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setRowAlignment:"), value)
}

func (g_ GridView) XPlacement() GridCellPlacement {
	rv := objc.CallMethod[GridCellPlacement](g_, objc.SEL("xPlacement"))
	return rv
}

func (g_ GridView) SetXPlacement(value GridCellPlacement) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setXPlacement:"), value)
}

func (g_ GridView) YPlacement() GridCellPlacement {
	rv := objc.CallMethod[GridCellPlacement](g_, objc.SEL("yPlacement"))
	return rv
}

func (g_ GridView) SetYPlacement(value GridCellPlacement) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setYPlacement:"), value)
}
