// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[GridCell](gc, objc.SEL("alloc"))
	return rv
}

func (gc _GridCellClass) New() GridCell {
	rv := objc.CallMethod[GridCell](gc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewGridCell() GridCell {
	return GridCellClass.New()
}

func (g_ GridCell) Init() GridCell {
	rv := objc.CallMethod[GridCell](g_, objc.SEL("init"))
	return rv
}

// weak property
func (g_ GridCell) Column() GridColumn {
	rv := objc.CallMethod[GridColumn](g_, objc.SEL("column"))
	return rv
}

// weak property
func (g_ GridCell) Row() GridRow {
	rv := objc.CallMethod[GridRow](g_, objc.SEL("row"))
	return rv
}

func (g_ GridCell) ContentView() View {
	rv := objc.CallMethod[View](g_, objc.SEL("contentView"))
	return rv
}

func (g_ GridCell) SetContentView(value IView) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setContentView:"), objc.ExtractPtr(value))
}

func (gc _GridCellClass) EmptyContentView() View {
	rv := objc.CallMethod[View](gc, objc.SEL("emptyContentView"))
	return rv
}

func (g_ GridCell) CustomPlacementConstraints() []LayoutConstraint {
	rv := objc.CallMethod[[]LayoutConstraint](g_, objc.SEL("customPlacementConstraints"))
	return rv
}

func (g_ GridCell) SetCustomPlacementConstraints(value []ILayoutConstraint) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setCustomPlacementConstraints:"), value)
}

func (g_ GridCell) RowAlignment() GridRowAlignment {
	rv := objc.CallMethod[GridRowAlignment](g_, objc.SEL("rowAlignment"))
	return rv
}

func (g_ GridCell) SetRowAlignment(value GridRowAlignment) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setRowAlignment:"), value)
}

func (g_ GridCell) XPlacement() GridCellPlacement {
	rv := objc.CallMethod[GridCellPlacement](g_, objc.SEL("xPlacement"))
	return rv
}

func (g_ GridCell) SetXPlacement(value GridCellPlacement) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setXPlacement:"), value)
}

func (g_ GridCell) YPlacement() GridCellPlacement {
	rv := objc.CallMethod[GridCellPlacement](g_, objc.SEL("yPlacement"))
	return rv
}

func (g_ GridCell) SetYPlacement(value GridCellPlacement) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setYPlacement:"), value)
}
