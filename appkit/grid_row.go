// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var GridRowClass = _GridRowClass{objc.GetClass("NSGridRow")}

type _GridRowClass struct {
	objc.Class
}

type IGridRow interface {
	objc.IObject
	CellAtIndex(index int) GridCell
	MergeCellsInRange(range_ foundation.Range)
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
	rv := objc.CallMethod[GridRow](gc, "alloc")
	return rv
}

func (gc _GridRowClass) New() GridRow {
	rv := objc.CallMethod[GridRow](gc, "new")
	rv.Autorelease()
	return rv
}

func NewGridRow() GridRow {
	return GridRowClass.New()
}

func (g_ GridRow) Init() GridRow {
	rv := objc.CallMethod[GridRow](g_, "init")
	return rv
}

func (g_ GridRow) CellAtIndex(index int) GridCell {
	rv := objc.CallMethod[GridCell](g_, "cellAtIndex:", index)
	return rv
}

func (g_ GridRow) MergeCellsInRange(range_ foundation.Range) {
	objc.CallMethod[objc.Void](g_, "mergeCellsInRange:", range_)
}

func (g_ GridRow) NumberOfCells() int {
	rv := objc.CallMethod[int](g_, "numberOfCells")
	return rv
}

func (g_ GridRow) IsHidden() bool {
	rv := objc.CallMethod[bool](g_, "isHidden")
	return rv
}

func (g_ GridRow) SetHidden(value bool) {
	objc.CallMethod[objc.Void](g_, "setHidden:", value)
}

func (g_ GridRow) TopPadding() float64 {
	rv := objc.CallMethod[float64](g_, "topPadding")
	return rv
}

func (g_ GridRow) SetTopPadding(value float64) {
	objc.CallMethod[objc.Void](g_, "setTopPadding:", value)
}

func (g_ GridRow) BottomPadding() float64 {
	rv := objc.CallMethod[float64](g_, "bottomPadding")
	return rv
}

func (g_ GridRow) SetBottomPadding(value float64) {
	objc.CallMethod[objc.Void](g_, "setBottomPadding:", value)
}

func (g_ GridRow) Height() float64 {
	rv := objc.CallMethod[float64](g_, "height")
	return rv
}

func (g_ GridRow) SetHeight(value float64) {
	objc.CallMethod[objc.Void](g_, "setHeight:", value)
}

func (g_ GridRow) RowAlignment() GridRowAlignment {
	rv := objc.CallMethod[GridRowAlignment](g_, "rowAlignment")
	return rv
}

func (g_ GridRow) SetRowAlignment(value GridRowAlignment) {
	objc.CallMethod[objc.Void](g_, "setRowAlignment:", value)
}

func (g_ GridRow) YPlacement() GridCellPlacement {
	rv := objc.CallMethod[GridCellPlacement](g_, "yPlacement")
	return rv
}

func (g_ GridRow) SetYPlacement(value GridCellPlacement) {
	objc.CallMethod[objc.Void](g_, "setYPlacement:", value)
}

func (g_ GridRow) GridView() GridView {
	rv := objc.CallMethod[GridView](g_, "gridView")
	return rv
}
