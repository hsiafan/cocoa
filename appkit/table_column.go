// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TableColumnClass = _TableColumnClass{objc.GetClass("NSTableColumn")}

type _TableColumnClass struct {
	objc.Class
}

type ITableColumn interface {
	objc.IObject
	SizeToFit()
	// deprecated
	IsResizable() bool
	// deprecated
	SetResizable(flag bool)
	// deprecated
	DataCellForRow(row int) objc.Object
	TableView() TableView
	SetTableView(value ITableView)
	Width() float64
	SetWidth(value float64)
	MinWidth() float64
	SetMinWidth(value float64)
	MaxWidth() float64
	SetMaxWidth(value float64)
	ResizingMask() TableColumnResizingOptions
	SetResizingMask(value TableColumnResizingOptions)
	Title() string
	SetTitle(value string)
	HeaderCell() TableHeaderCell
	SetHeaderCell(value ITableHeaderCell)
	Identifier() UserInterfaceItemIdentifier
	SetIdentifier(value UserInterfaceItemIdentifier)
	IsEditable() bool
	SetEditable(value bool)
	SortDescriptorPrototype() foundation.SortDescriptor
	SetSortDescriptorPrototype(value foundation.ISortDescriptor)
	IsHidden() bool
	SetHidden(value bool)
	HeaderToolTip() string
	SetHeaderToolTip(value string)
	// deprecated
	DataCell() objc.Object
	// deprecated
	SetDataCell(value objc.IObject)
}

type TableColumn struct {
	objc.Object
}

func MakeTableColumn(ptr unsafe.Pointer) TableColumn {
	return TableColumn{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TableColumn) InitWithIdentifier(identifier UserInterfaceItemIdentifier) TableColumn {
	rv := objc.CallMethod[TableColumn](t_, "initWithIdentifier:", identifier)
	return rv
}

func (tc _TableColumnClass) Alloc() TableColumn {
	rv := objc.CallMethod[TableColumn](tc, "alloc")
	return rv
}

func (tc _TableColumnClass) New() TableColumn {
	rv := objc.CallMethod[TableColumn](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTableColumn() TableColumn {
	return TableColumnClass.New()
}

func (t_ TableColumn) Init() TableColumn {
	rv := objc.CallMethod[TableColumn](t_, "init")
	return rv
}

func (t_ TableColumn) SizeToFit() {
	objc.CallMethod[objc.Void](t_, "sizeToFit")
}

// deprecated
func (t_ TableColumn) IsResizable() bool {
	rv := objc.CallMethod[bool](t_, "isResizable")
	return rv
}

// deprecated
func (t_ TableColumn) SetResizable(flag bool) {
	objc.CallMethod[objc.Void](t_, "setResizable:", flag)
}

// deprecated
func (t_ TableColumn) DataCellForRow(row int) objc.Object {
	rv := objc.CallMethod[objc.Object](t_, "dataCellForRow:", row)
	return rv
}

func (t_ TableColumn) TableView() TableView {
	rv := objc.CallMethod[TableView](t_, "tableView")
	return rv
}

func (t_ TableColumn) SetTableView(value ITableView) {
	objc.CallMethod[objc.Void](t_, "setTableView:", value)
}

func (t_ TableColumn) Width() float64 {
	rv := objc.CallMethod[float64](t_, "width")
	return rv
}

func (t_ TableColumn) SetWidth(value float64) {
	objc.CallMethod[objc.Void](t_, "setWidth:", value)
}

func (t_ TableColumn) MinWidth() float64 {
	rv := objc.CallMethod[float64](t_, "minWidth")
	return rv
}

func (t_ TableColumn) SetMinWidth(value float64) {
	objc.CallMethod[objc.Void](t_, "setMinWidth:", value)
}

func (t_ TableColumn) MaxWidth() float64 {
	rv := objc.CallMethod[float64](t_, "maxWidth")
	return rv
}

func (t_ TableColumn) SetMaxWidth(value float64) {
	objc.CallMethod[objc.Void](t_, "setMaxWidth:", value)
}

func (t_ TableColumn) ResizingMask() TableColumnResizingOptions {
	rv := objc.CallMethod[TableColumnResizingOptions](t_, "resizingMask")
	return rv
}

func (t_ TableColumn) SetResizingMask(value TableColumnResizingOptions) {
	objc.CallMethod[objc.Void](t_, "setResizingMask:", value)
}

func (t_ TableColumn) Title() string {
	rv := objc.CallMethod[string](t_, "title")
	return rv
}

func (t_ TableColumn) SetTitle(value string) {
	objc.CallMethod[objc.Void](t_, "setTitle:", value)
}

func (t_ TableColumn) HeaderCell() TableHeaderCell {
	rv := objc.CallMethod[TableHeaderCell](t_, "headerCell")
	return rv
}

func (t_ TableColumn) SetHeaderCell(value ITableHeaderCell) {
	objc.CallMethod[objc.Void](t_, "setHeaderCell:", value)
}

func (t_ TableColumn) Identifier() UserInterfaceItemIdentifier {
	rv := objc.CallMethod[UserInterfaceItemIdentifier](t_, "identifier")
	return rv
}

func (t_ TableColumn) SetIdentifier(value UserInterfaceItemIdentifier) {
	objc.CallMethod[objc.Void](t_, "setIdentifier:", value)
}

func (t_ TableColumn) IsEditable() bool {
	rv := objc.CallMethod[bool](t_, "isEditable")
	return rv
}

func (t_ TableColumn) SetEditable(value bool) {
	objc.CallMethod[objc.Void](t_, "setEditable:", value)
}

func (t_ TableColumn) SortDescriptorPrototype() foundation.SortDescriptor {
	rv := objc.CallMethod[foundation.SortDescriptor](t_, "sortDescriptorPrototype")
	return rv
}

func (t_ TableColumn) SetSortDescriptorPrototype(value foundation.ISortDescriptor) {
	objc.CallMethod[objc.Void](t_, "setSortDescriptorPrototype:", value)
}

func (t_ TableColumn) IsHidden() bool {
	rv := objc.CallMethod[bool](t_, "isHidden")
	return rv
}

func (t_ TableColumn) SetHidden(value bool) {
	objc.CallMethod[objc.Void](t_, "setHidden:", value)
}

func (t_ TableColumn) HeaderToolTip() string {
	rv := objc.CallMethod[string](t_, "headerToolTip")
	return rv
}

func (t_ TableColumn) SetHeaderToolTip(value string) {
	objc.CallMethod[objc.Void](t_, "setHeaderToolTip:", value)
}

// deprecated
func (t_ TableColumn) DataCell() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, "dataCell")
	return rv
}

// deprecated
func (t_ TableColumn) SetDataCell(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, "setDataCell:", value)
}
