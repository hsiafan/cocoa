// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var GridColumnClass = _GridColumnClass{objc.GetClass("NSGridColumn")}

type _GridColumnClass struct {
	objc.Class
}

type IGridColumn interface {
	objc.IObject
	CellAtIndex(index int) GridCell
	MergeCellsInRange(range_ foundation.Range)
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

func (gc _GridColumnClass) New() GridColumn {
	rv := ffi.CallMethod[GridColumn](gc, "new")
	rv.Autorelease()
	return rv
}

func NewGridColumn() GridColumn {
	return GridColumnClass.New()
}

func (g_ GridColumn) Init() GridColumn {
	rv := ffi.CallMethod[GridColumn](g_, "init")
	return rv
}

func (g_ GridColumn) CellAtIndex(index int) GridCell {
	rv := ffi.CallMethod[GridCell](g_, "cellAtIndex:", index)
	return rv
}

func (g_ GridColumn) MergeCellsInRange(range_ foundation.Range) {
	ffi.CallMethod[ffi.Void](g_, "mergeCellsInRange:", range_)
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
