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

var GridCellClass = _GridCellClass{objc.GetClass("NSGridCell")}

type _GridCellClass struct {
	objc.Class
}

type IGridCell interface {
	objc.IObject
	Column() GridColumn
	Row() GridRow
	ContentView() View
	SetContentView(value IView)
	CustomPlacementConstraints() []LayoutConstraint
	SetCustomPlacementConstraints(value []ILayoutConstraint)
	RowAlignment() GridRowAlignment
	SetRowAlignment(value GridRowAlignment)
	XPlacement() GridCellPlacement
	SetXPlacement(value GridCellPlacement)
	YPlacement() GridCellPlacement
	SetYPlacement(value GridCellPlacement)
}

type GridCell struct {
	objc.Object
}

func MakeGridCell(ptr unsafe.Pointer) GridCell {
	return GridCell{
		Object: objc.MakeObject(ptr),
	}
}

func (gc _GridCellClass) Alloc() GridCell {
	rv := ffi.CallMethod[GridCell](gc, "alloc")
	return rv
}

func (g_ GridCell) Init() GridCell {
	rv := ffi.CallMethod[GridCell](g_, "init")
	return rv
}

func (gc _GridCellClass) New() GridCell {
	rv := ffi.CallMethod[GridCell](gc, "new")
	rv.Autorelease()
	return rv
}

func NewGridCell() GridCell {
	return GridCellClass.New()
}

func (g_ GridCell) Column() GridColumn {
	rv := ffi.CallMethod[GridColumn](g_, "column")
	return rv
}

func (g_ GridCell) Row() GridRow {
	rv := ffi.CallMethod[GridRow](g_, "row")
	return rv
}

func (g_ GridCell) ContentView() View {
	rv := ffi.CallMethod[View](g_, "contentView")
	return rv
}

func (g_ GridCell) SetContentView(value IView) {
	ffi.CallMethod[ffi.Void](g_, "setContentView:", value)
}

func (gc _GridCellClass) EmptyContentView() View {
	rv := ffi.CallMethod[View](gc, "emptyContentView")
	return rv
}

func (g_ GridCell) CustomPlacementConstraints() []LayoutConstraint {
	rv := ffi.CallMethod[[]LayoutConstraint](g_, "customPlacementConstraints")
	return rv
}

func (g_ GridCell) SetCustomPlacementConstraints(value []ILayoutConstraint) {
	ffi.CallMethod[ffi.Void](g_, "setCustomPlacementConstraints:", value)
}

func (g_ GridCell) RowAlignment() GridRowAlignment {
	rv := ffi.CallMethod[GridRowAlignment](g_, "rowAlignment")
	return rv
}

func (g_ GridCell) SetRowAlignment(value GridRowAlignment) {
	ffi.CallMethod[ffi.Void](g_, "setRowAlignment:", value)
}

func (g_ GridCell) XPlacement() GridCellPlacement {
	rv := ffi.CallMethod[GridCellPlacement](g_, "xPlacement")
	return rv
}

func (g_ GridCell) SetXPlacement(value GridCellPlacement) {
	ffi.CallMethod[ffi.Void](g_, "setXPlacement:", value)
}

func (g_ GridCell) YPlacement() GridCellPlacement {
	rv := ffi.CallMethod[GridCellPlacement](g_, "yPlacement")
	return rv
}

func (g_ GridCell) SetYPlacement(value GridCellPlacement) {
	ffi.CallMethod[ffi.Void](g_, "setYPlacement:", value)
}

var GridColumnClass = _GridColumnClass{objc.GetClass("NSGridColumn")}

type _GridColumnClass struct {
	objc.Class
}

type IGridColumn interface {
	objc.IObject
	CellAtIndex(index int) GridCell
	MergeCellsInRange(_range foundation.Range)
	GridView() GridView
	IsHidden() bool
	SetHidden(value bool)
	LeadingPadding() float64
	SetLeadingPadding(value float64)
	NumberOfCells() int
	TrailingPadding() float64
	SetTrailingPadding(value float64)
	Width() float64
	SetWidth(value float64)
	XPlacement() GridCellPlacement
	SetXPlacement(value GridCellPlacement)
}

type GridColumn struct {
	objc.Object
}

func MakeGridColumn(ptr unsafe.Pointer) GridColumn {
	return GridColumn{
		Object: objc.MakeObject(ptr),
	}
}

func (gc _GridColumnClass) Alloc() GridColumn {
	rv := ffi.CallMethod[GridColumn](gc, "alloc")
	return rv
}

func (g_ GridColumn) Init() GridColumn {
	rv := ffi.CallMethod[GridColumn](g_, "init")
	return rv
}

func (gc _GridColumnClass) New() GridColumn {
	rv := ffi.CallMethod[GridColumn](gc, "new")
	rv.Autorelease()
	return rv
}

func NewGridColumn() GridColumn {
	return GridColumnClass.New()
}

func (g_ GridColumn) CellAtIndex(index int) GridCell {
	rv := ffi.CallMethod[GridCell](g_, "cellAtIndex:", index)
	return rv
}

func (g_ GridColumn) MergeCellsInRange(_range foundation.Range) {
	ffi.CallMethod[ffi.Void](g_, "mergeCellsInRange:", _range)
}

func (g_ GridColumn) GridView() GridView {
	rv := ffi.CallMethod[GridView](g_, "gridView")
	return rv
}

func (g_ GridColumn) IsHidden() bool {
	rv := ffi.CallMethod[bool](g_, "isHidden")
	return rv
}

func (g_ GridColumn) SetHidden(value bool) {
	ffi.CallMethod[ffi.Void](g_, "setHidden:", value)
}

func (g_ GridColumn) LeadingPadding() float64 {
	rv := ffi.CallMethod[float64](g_, "leadingPadding")
	return rv
}

func (g_ GridColumn) SetLeadingPadding(value float64) {
	ffi.CallMethod[ffi.Void](g_, "setLeadingPadding:", value)
}

func (g_ GridColumn) NumberOfCells() int {
	rv := ffi.CallMethod[int](g_, "numberOfCells")
	return rv
}

func (g_ GridColumn) TrailingPadding() float64 {
	rv := ffi.CallMethod[float64](g_, "trailingPadding")
	return rv
}

func (g_ GridColumn) SetTrailingPadding(value float64) {
	ffi.CallMethod[ffi.Void](g_, "setTrailingPadding:", value)
}

func (g_ GridColumn) Width() float64 {
	rv := ffi.CallMethod[float64](g_, "width")
	return rv
}

func (g_ GridColumn) SetWidth(value float64) {
	ffi.CallMethod[ffi.Void](g_, "setWidth:", value)
}

func (g_ GridColumn) XPlacement() GridCellPlacement {
	rv := ffi.CallMethod[GridCellPlacement](g_, "xPlacement")
	return rv
}

func (g_ GridColumn) SetXPlacement(value GridCellPlacement) {
	ffi.CallMethod[ffi.Void](g_, "setXPlacement:", value)
}

var GridRowClass = _GridRowClass{objc.GetClass("NSGridRow")}

type _GridRowClass struct {
	objc.Class
}

type IGridRow interface {
	objc.IObject
	CellAtIndex(index int) GridCell
	MergeCellsInRange(_range foundation.Range)
	NumberOfCells() int
	IsHidden() bool
	SetHidden(value bool)
	TopPadding() float64
	SetTopPadding(value float64)
	BottomPadding() float64
	SetBottomPadding(value float64)
	Height() float64
	SetHeight(value float64)
	RowAlignment() GridRowAlignment
	SetRowAlignment(value GridRowAlignment)
	YPlacement() GridCellPlacement
	SetYPlacement(value GridCellPlacement)
	GridView() GridView
}

type GridRow struct {
	objc.Object
}

func MakeGridRow(ptr unsafe.Pointer) GridRow {
	return GridRow{
		Object: objc.MakeObject(ptr),
	}
}

func (gc _GridRowClass) Alloc() GridRow {
	rv := ffi.CallMethod[GridRow](gc, "alloc")
	return rv
}

func (g_ GridRow) Init() GridRow {
	rv := ffi.CallMethod[GridRow](g_, "init")
	return rv
}

func (gc _GridRowClass) New() GridRow {
	rv := ffi.CallMethod[GridRow](gc, "new")
	rv.Autorelease()
	return rv
}

func NewGridRow() GridRow {
	return GridRowClass.New()
}

func (g_ GridRow) CellAtIndex(index int) GridCell {
	rv := ffi.CallMethod[GridCell](g_, "cellAtIndex:", index)
	return rv
}

func (g_ GridRow) MergeCellsInRange(_range foundation.Range) {
	ffi.CallMethod[ffi.Void](g_, "mergeCellsInRange:", _range)
}

func (g_ GridRow) NumberOfCells() int {
	rv := ffi.CallMethod[int](g_, "numberOfCells")
	return rv
}

func (g_ GridRow) IsHidden() bool {
	rv := ffi.CallMethod[bool](g_, "isHidden")
	return rv
}

func (g_ GridRow) SetHidden(value bool) {
	ffi.CallMethod[ffi.Void](g_, "setHidden:", value)
}

func (g_ GridRow) TopPadding() float64 {
	rv := ffi.CallMethod[float64](g_, "topPadding")
	return rv
}

func (g_ GridRow) SetTopPadding(value float64) {
	ffi.CallMethod[ffi.Void](g_, "setTopPadding:", value)
}

func (g_ GridRow) BottomPadding() float64 {
	rv := ffi.CallMethod[float64](g_, "bottomPadding")
	return rv
}

func (g_ GridRow) SetBottomPadding(value float64) {
	ffi.CallMethod[ffi.Void](g_, "setBottomPadding:", value)
}

func (g_ GridRow) Height() float64 {
	rv := ffi.CallMethod[float64](g_, "height")
	return rv
}

func (g_ GridRow) SetHeight(value float64) {
	ffi.CallMethod[ffi.Void](g_, "setHeight:", value)
}

func (g_ GridRow) RowAlignment() GridRowAlignment {
	rv := ffi.CallMethod[GridRowAlignment](g_, "rowAlignment")
	return rv
}

func (g_ GridRow) SetRowAlignment(value GridRowAlignment) {
	ffi.CallMethod[ffi.Void](g_, "setRowAlignment:", value)
}

func (g_ GridRow) YPlacement() GridCellPlacement {
	rv := ffi.CallMethod[GridCellPlacement](g_, "yPlacement")
	return rv
}

func (g_ GridRow) SetYPlacement(value GridCellPlacement) {
	ffi.CallMethod[ffi.Void](g_, "setYPlacement:", value)
}

func (g_ GridRow) GridView() GridView {
	rv := ffi.CallMethod[GridView](g_, "gridView")
	return rv
}
