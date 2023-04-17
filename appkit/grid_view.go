// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[GridView](gc, "gridViewWithNumberOfColumns:rows:", columnCount, rowCount)
	return rv
}

func (gc _GridViewClass) GridViewWithViews(rows [][]IView) GridView {
	rv := ffi.CallMethod[GridView](gc, "gridViewWithViews:", rows)
	return rv
}

func (g_ GridView) InitWithFrame(frameRect foundation.Rect) GridView {
	rv := ffi.CallMethod[GridView](g_, "initWithFrame:", frameRect)
	return rv
}

func (g_ GridView) Init() GridView {
	rv := ffi.CallMethod[GridView](g_, "init")
	return rv
}

func (gc _GridViewClass) Alloc() GridView {
	rv := ffi.CallMethod[GridView](gc, "alloc")
	return rv
}

func (gc _GridViewClass) New() GridView {
	rv := ffi.CallMethod[GridView](gc, "new")
	rv.Autorelease()
	return rv
}

func NewGridView() GridView {
	return GridViewClass.New()
}

func (g_ GridView) IndexOfColumn(column IGridColumn) int {
	rv := ffi.CallMethod[int](g_, "indexOfColumn:", column)
	return rv
}

func (g_ GridView) RowAtIndex(index int) GridRow {
	rv := ffi.CallMethod[GridRow](g_, "rowAtIndex:", index)
	return rv
}

func (g_ GridView) ColumnAtIndex(index int) GridColumn {
	rv := ffi.CallMethod[GridColumn](g_, "columnAtIndex:", index)
	return rv
}

func (g_ GridView) IndexOfRow(row IGridRow) int {
	rv := ffi.CallMethod[int](g_, "indexOfRow:", row)
	return rv
}

func (g_ GridView) AddRowWithViews(views []IView) GridRow {
	rv := ffi.CallMethod[GridRow](g_, "addRowWithViews:", views)
	return rv
}

func (g_ GridView) InsertRowAtIndex_WithViews(index int, views []IView) GridRow {
	rv := ffi.CallMethod[GridRow](g_, "insertRowAtIndex:withViews:", index, views)
	return rv
}

func (g_ GridView) RemoveRowAtIndex(index int) {
	ffi.CallMethod[ffi.Void](g_, "removeRowAtIndex:", index)
}

func (g_ GridView) MoveRowAtIndex_ToIndex(fromIndex int, toIndex int) {
	ffi.CallMethod[ffi.Void](g_, "moveRowAtIndex:toIndex:", fromIndex, toIndex)
}

func (g_ GridView) AddColumnWithViews(views []IView) GridColumn {
	rv := ffi.CallMethod[GridColumn](g_, "addColumnWithViews:", views)
	return rv
}

func (g_ GridView) InsertColumnAtIndex_WithViews(index int, views []IView) GridColumn {
	rv := ffi.CallMethod[GridColumn](g_, "insertColumnAtIndex:withViews:", index, views)
	return rv
}

func (g_ GridView) RemoveColumnAtIndex(index int) {
	ffi.CallMethod[ffi.Void](g_, "removeColumnAtIndex:", index)
}

func (g_ GridView) MoveColumnAtIndex_ToIndex(fromIndex int, toIndex int) {
	ffi.CallMethod[ffi.Void](g_, "moveColumnAtIndex:toIndex:", fromIndex, toIndex)
}

func (g_ GridView) CellAtColumnIndex_RowIndex(columnIndex int, rowIndex int) GridCell {
	rv := ffi.CallMethod[GridCell](g_, "cellAtColumnIndex:rowIndex:", columnIndex, rowIndex)
	return rv
}

func (g_ GridView) CellForView(view IView) GridCell {
	rv := ffi.CallMethod[GridCell](g_, "cellForView:", view)
	return rv
}

func (g_ GridView) MergeCellsInHorizontalRange_VerticalRange(hRange foundation.Range, vRange foundation.Range) {
	ffi.CallMethod[ffi.Void](g_, "mergeCellsInHorizontalRange:verticalRange:", hRange, vRange)
}

func (g_ GridView) NumberOfRows() int {
	rv := ffi.CallMethod[int](g_, "numberOfRows")
	return rv
}

func (g_ GridView) NumberOfColumns() int {
	rv := ffi.CallMethod[int](g_, "numberOfColumns")
	return rv
}

func (g_ GridView) ColumnSpacing() float64 {
	rv := ffi.CallMethod[float64](g_, "columnSpacing")
	return rv
}

func (g_ GridView) SetColumnSpacing(value float64) {
	ffi.CallMethod[ffi.Void](g_, "setColumnSpacing:", value)
}

func (g_ GridView) RowSpacing() float64 {
	rv := ffi.CallMethod[float64](g_, "rowSpacing")
	return rv
}

func (g_ GridView) SetRowSpacing(value float64) {
	ffi.CallMethod[ffi.Void](g_, "setRowSpacing:", value)
}

func (g_ GridView) RowAlignment() GridRowAlignment {
	rv := ffi.CallMethod[GridRowAlignment](g_, "rowAlignment")
	return rv
}

func (g_ GridView) SetRowAlignment(value GridRowAlignment) {
	ffi.CallMethod[ffi.Void](g_, "setRowAlignment:", value)
}

func (g_ GridView) XPlacement() GridCellPlacement {
	rv := ffi.CallMethod[GridCellPlacement](g_, "xPlacement")
	return rv
}

func (g_ GridView) SetXPlacement(value GridCellPlacement) {
	ffi.CallMethod[ffi.Void](g_, "setXPlacement:", value)
}

func (g_ GridView) YPlacement() GridCellPlacement {
	rv := ffi.CallMethod[GridCellPlacement](g_, "yPlacement")
	return rv
}

func (g_ GridView) SetYPlacement(value GridCellPlacement) {
	ffi.CallMethod[ffi.Void](g_, "setYPlacement:", value)
}
