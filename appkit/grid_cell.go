// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

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

func (gc _GridCellClass) New() GridCell {
	rv := ffi.CallMethod[GridCell](gc, "new")
	rv.Autorelease()
	return rv
}

func NewGridCell() GridCell {
	return GridCellClass.New()
}

func (g_ GridCell) Init() GridCell {
	rv := ffi.CallMethod[GridCell](g_, "init")
	return rv
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
