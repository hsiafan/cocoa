// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[GridColumn](gc, objc.SEL("alloc"))
	return rv
}

func (gc _GridColumnClass) New() GridColumn {
	rv := objc.CallMethod[GridColumn](gc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewGridColumn() GridColumn {
	return GridColumnClass.New()
}

func (g_ GridColumn) Init() GridColumn {
	rv := objc.CallMethod[GridColumn](g_, objc.SEL("init"))
	return rv
}

func (g_ GridColumn) CellAtIndex(index int) GridCell {
	rv := objc.CallMethod[GridCell](g_, objc.SEL("cellAtIndex:"), index)
	return rv
}

func (g_ GridColumn) MergeCellsInRange(range_ foundation.Range) {
	objc.CallMethod[objc.Void](g_, objc.SEL("mergeCellsInRange:"), range_)
}

// weak property
func (g_ GridColumn) GridView() GridView {
	rv := objc.CallMethod[GridView](g_, objc.SEL("gridView"))
	return rv
}

func (g_ GridColumn) IsHidden() bool {
	rv := objc.CallMethod[bool](g_, objc.SEL("isHidden"))
	return rv
}

func (g_ GridColumn) SetHidden(value bool) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setHidden:"), value)
}

func (g_ GridColumn) LeadingPadding() float64 {
	rv := objc.CallMethod[float64](g_, objc.SEL("leadingPadding"))
	return rv
}

func (g_ GridColumn) SetLeadingPadding(value float64) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setLeadingPadding:"), value)
}

func (g_ GridColumn) NumberOfCells() int {
	rv := objc.CallMethod[int](g_, objc.SEL("numberOfCells"))
	return rv
}

func (g_ GridColumn) TrailingPadding() float64 {
	rv := objc.CallMethod[float64](g_, objc.SEL("trailingPadding"))
	return rv
}

func (g_ GridColumn) SetTrailingPadding(value float64) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setTrailingPadding:"), value)
}

func (g_ GridColumn) Width() float64 {
	rv := objc.CallMethod[float64](g_, objc.SEL("width"))
	return rv
}

func (g_ GridColumn) SetWidth(value float64) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setWidth:"), value)
}

func (g_ GridColumn) XPlacement() GridCellPlacement {
	rv := objc.CallMethod[GridCellPlacement](g_, objc.SEL("xPlacement"))
	return rv
}

func (g_ GridColumn) SetXPlacement(value GridCellPlacement) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setXPlacement:"), value)
}
