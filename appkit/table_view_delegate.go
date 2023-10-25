// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type TableViewDelegate interface {
	ControlTextEditingDelegate
	ImplementsTableView_ViewForTableColumn_Row() bool
	// optional
	TableView_ViewForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) IView
	ImplementsTableView_RowViewForRow() bool
	// optional
	TableView_RowViewForRow(tableView TableView, row int) ITableRowView
	ImplementsTableView_DidAddRowView_ForRow() bool
	// optional
	TableView_DidAddRowView_ForRow(tableView TableView, rowView TableRowView, row int)
	ImplementsTableView_DidRemoveRowView_ForRow() bool
	// optional
	TableView_DidRemoveRowView_ForRow(tableView TableView, rowView TableRowView, row int)
	ImplementsTableView_IsGroupRow() bool
	// optional
	TableView_IsGroupRow(tableView TableView, row int) bool
	ImplementsTableView_WillDisplayCell_ForTableColumn_Row() bool
	// optional
	TableView_WillDisplayCell_ForTableColumn_Row(tableView TableView, cell objc.Object, tableColumn TableColumn, row int)
	ImplementsTableView_DataCellForTableColumn_Row() bool
	// optional
	TableView_DataCellForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) ICell
	ImplementsTableView_ShouldShowCellExpansionForTableColumn_Row() bool
	// optional
	TableView_ShouldShowCellExpansionForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) bool
	ImplementsTableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation() bool
	// optional
	TableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation(tableView TableView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, row int, mouseLocation foundation.Point) string
	ImplementsTableView_ShouldEditTableColumn_Row() bool
	// optional
	TableView_ShouldEditTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) bool
	ImplementsTableView_HeightOfRow() bool
	// optional
	TableView_HeightOfRow(tableView TableView, row int) float64
	ImplementsTableView_SizeToFitWidthOfColumn() bool
	// optional
	TableView_SizeToFitWidthOfColumn(tableView TableView, column int) float64
	ImplementsSelectionShouldChangeInTableView() bool
	// optional
	SelectionShouldChangeInTableView(tableView TableView) bool
	ImplementsTableView_ShouldSelectRow() bool
	// optional
	TableView_ShouldSelectRow(tableView TableView, row int) bool
	ImplementsTableView_SelectionIndexesForProposedSelection() bool
	// optional
	TableView_SelectionIndexesForProposedSelection(tableView TableView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet
	ImplementsTableView_ShouldSelectTableColumn() bool
	// optional
	TableView_ShouldSelectTableColumn(tableView TableView, tableColumn TableColumn) bool
	ImplementsTableViewSelectionIsChanging() bool
	// optional
	TableViewSelectionIsChanging(notification foundation.Notification)
	ImplementsTableViewSelectionDidChange() bool
	// optional
	TableViewSelectionDidChange(notification foundation.Notification)
	ImplementsTableView_ShouldTypeSelectForEvent_WithCurrentSearchString() bool
	// optional
	TableView_ShouldTypeSelectForEvent_WithCurrentSearchString(tableView TableView, event Event, searchString string) bool
	ImplementsTableView_TypeSelectStringForTableColumn_Row() bool
	// optional
	TableView_TypeSelectStringForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) string
	ImplementsTableView_NextTypeSelectMatchFromRow_ToRow_ForString() bool
	// optional
	TableView_NextTypeSelectMatchFromRow_ToRow_ForString(tableView TableView, startRow int, endRow int, searchString string) int
	ImplementsTableView_ShouldReorderColumn_ToColumn() bool
	// optional
	TableView_ShouldReorderColumn_ToColumn(tableView TableView, columnIndex int, newColumnIndex int) bool
	ImplementsTableView_DidDragTableColumn() bool
	// optional
	TableView_DidDragTableColumn(tableView TableView, tableColumn TableColumn)
	ImplementsTableViewColumnDidMove() bool
	// optional
	TableViewColumnDidMove(notification foundation.Notification)
	ImplementsTableViewColumnDidResize() bool
	// optional
	TableViewColumnDidResize(notification foundation.Notification)
	ImplementsTableView_DidClickTableColumn() bool
	// optional
	TableView_DidClickTableColumn(tableView TableView, tableColumn TableColumn)
	ImplementsTableView_MouseDownInHeaderOfTableColumn() bool
	// optional
	TableView_MouseDownInHeaderOfTableColumn(tableView TableView, tableColumn TableColumn)
	ImplementsTableView_ShouldTrackCell_ForTableColumn_Row() bool
	// optional
	TableView_ShouldTrackCell_ForTableColumn_Row(tableView TableView, cell Cell, tableColumn TableColumn, row int) bool
	ImplementsTableView_RowActionsForRow_Edge() bool
	// optional
	TableView_RowActionsForRow_Edge(tableView TableView, row int, edge TableRowActionEdge) []ITableViewRowAction
}

func WrapTableViewDelegate(v TableViewDelegate) objc.Object {
	return objc.WrapAsProtocol("NSTableViewDelegate", v)
}

type TableViewDelegateBase struct {
	ControlTextEditingDelegateBase
}

func (p *TableViewDelegateBase) ImplementsTableView_ViewForTableColumn_Row() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_ViewForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) IView {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_RowViewForRow() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_RowViewForRow(tableView TableView, row int) ITableRowView {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_DidAddRowView_ForRow() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_DidAddRowView_ForRow(tableView TableView, rowView TableRowView, row int) {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_DidRemoveRowView_ForRow() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_DidRemoveRowView_ForRow(tableView TableView, rowView TableRowView, row int) {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_IsGroupRow() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_IsGroupRow(tableView TableView, row int) bool {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_WillDisplayCell_ForTableColumn_Row() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_WillDisplayCell_ForTableColumn_Row(tableView TableView, cell objc.Object, tableColumn TableColumn, row int) {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_DataCellForTableColumn_Row() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_DataCellForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) ICell {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_ShouldShowCellExpansionForTableColumn_Row() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_ShouldShowCellExpansionForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) bool {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation(tableView TableView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, row int, mouseLocation foundation.Point) string {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_ShouldEditTableColumn_Row() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_ShouldEditTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) bool {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_HeightOfRow() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_HeightOfRow(tableView TableView, row int) float64 {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_SizeToFitWidthOfColumn() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_SizeToFitWidthOfColumn(tableView TableView, column int) float64 {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsSelectionShouldChangeInTableView() bool {
	return false
}

func (p *TableViewDelegateBase) SelectionShouldChangeInTableView(tableView TableView) bool {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_ShouldSelectRow() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_ShouldSelectRow(tableView TableView, row int) bool {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_SelectionIndexesForProposedSelection() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_SelectionIndexesForProposedSelection(tableView TableView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_ShouldSelectTableColumn() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_ShouldSelectTableColumn(tableView TableView, tableColumn TableColumn) bool {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableViewSelectionIsChanging() bool {
	return false
}

func (p *TableViewDelegateBase) TableViewSelectionIsChanging(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableViewSelectionDidChange() bool {
	return false
}

func (p *TableViewDelegateBase) TableViewSelectionDidChange(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_ShouldTypeSelectForEvent_WithCurrentSearchString() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_ShouldTypeSelectForEvent_WithCurrentSearchString(tableView TableView, event Event, searchString string) bool {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_TypeSelectStringForTableColumn_Row() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_TypeSelectStringForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) string {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_NextTypeSelectMatchFromRow_ToRow_ForString() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_NextTypeSelectMatchFromRow_ToRow_ForString(tableView TableView, startRow int, endRow int, searchString string) int {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_ShouldReorderColumn_ToColumn() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_ShouldReorderColumn_ToColumn(tableView TableView, columnIndex int, newColumnIndex int) bool {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_DidDragTableColumn() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_DidDragTableColumn(tableView TableView, tableColumn TableColumn) {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableViewColumnDidMove() bool {
	return false
}

func (p *TableViewDelegateBase) TableViewColumnDidMove(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableViewColumnDidResize() bool {
	return false
}

func (p *TableViewDelegateBase) TableViewColumnDidResize(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_DidClickTableColumn() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_DidClickTableColumn(tableView TableView, tableColumn TableColumn) {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_MouseDownInHeaderOfTableColumn() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_MouseDownInHeaderOfTableColumn(tableView TableView, tableColumn TableColumn) {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_ShouldTrackCell_ForTableColumn_Row() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_ShouldTrackCell_ForTableColumn_Row(tableView TableView, cell Cell, tableColumn TableColumn, row int) bool {
	panic("unimpemented protocol method")
}

func (p *TableViewDelegateBase) ImplementsTableView_RowActionsForRow_Edge() bool {
	return false
}

func (p *TableViewDelegateBase) TableView_RowActionsForRow_Edge(tableView TableView, row int, edge TableRowActionEdge) []ITableViewRowAction {
	panic("unimpemented protocol method")
}

type TableViewDelegateCreator struct {
	className string
	class     objc.Class
}

func NewTableViewDelegateCreator(name string) *TableViewDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("ProtocolBase"), name, 0)
	objc.RegisterClassPair(class)
	return &TableViewDelegateCreator{className: name, class: class}
}

func (c *TableViewDelegateCreator) SetTableView_ViewForTableColumn_Row(handle func(o objc.ProtocolBase, tableView TableView, tableColumn TableColumn, row int) IView) {
	objc.AddMethod(c.class, objc.SEL("tableView:viewForTableColumn:row:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_RowViewForRow(handle func(o objc.ProtocolBase, tableView TableView, row int) ITableRowView) {
	objc.AddMethod(c.class, objc.SEL("tableView:rowViewForRow:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_DidAddRowView_ForRow(handle func(o objc.ProtocolBase, tableView TableView, rowView TableRowView, row int)) {
	objc.AddMethod(c.class, objc.SEL("tableView:didAddRowView:forRow:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_DidRemoveRowView_ForRow(handle func(o objc.ProtocolBase, tableView TableView, rowView TableRowView, row int)) {
	objc.AddMethod(c.class, objc.SEL("tableView:didRemoveRowView:forRow:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_IsGroupRow(handle func(o objc.ProtocolBase, tableView TableView, row int) bool) {
	objc.AddMethod(c.class, objc.SEL("tableView:isGroupRow:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_WillDisplayCell_ForTableColumn_Row(handle func(o objc.ProtocolBase, tableView TableView, cell objc.Object, tableColumn TableColumn, row int)) {
	objc.AddMethod(c.class, objc.SEL("tableView:willDisplayCell:forTableColumn:row:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_DataCellForTableColumn_Row(handle func(o objc.ProtocolBase, tableView TableView, tableColumn TableColumn, row int) ICell) {
	objc.AddMethod(c.class, objc.SEL("tableView:dataCellForTableColumn:row:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_ShouldShowCellExpansionForTableColumn_Row(handle func(o objc.ProtocolBase, tableView TableView, tableColumn TableColumn, row int) bool) {
	objc.AddMethod(c.class, objc.SEL("tableView:shouldShowCellExpansionForTableColumn:row:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation(handle func(o objc.ProtocolBase, tableView TableView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, row int, mouseLocation foundation.Point) string) {
	objc.AddMethod(c.class, objc.SEL("tableView:toolTipForCell:rect:tableColumn:row:mouseLocation:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_ShouldEditTableColumn_Row(handle func(o objc.ProtocolBase, tableView TableView, tableColumn TableColumn, row int) bool) {
	objc.AddMethod(c.class, objc.SEL("tableView:shouldEditTableColumn:row:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_HeightOfRow(handle func(o objc.ProtocolBase, tableView TableView, row int) float64) {
	objc.AddMethod(c.class, objc.SEL("tableView:heightOfRow:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_SizeToFitWidthOfColumn(handle func(o objc.ProtocolBase, tableView TableView, column int) float64) {
	objc.AddMethod(c.class, objc.SEL("tableView:sizeToFitWidthOfColumn:"), handle)
}

func (c *TableViewDelegateCreator) SetSelectionShouldChangeInTableView(handle func(o objc.ProtocolBase, tableView TableView) bool) {
	objc.AddMethod(c.class, objc.SEL("selectionShouldChangeInTableView:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_ShouldSelectRow(handle func(o objc.ProtocolBase, tableView TableView, row int) bool) {
	objc.AddMethod(c.class, objc.SEL("tableView:shouldSelectRow:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_SelectionIndexesForProposedSelection(handle func(o objc.ProtocolBase, tableView TableView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet) {
	objc.AddMethod(c.class, objc.SEL("tableView:selectionIndexesForProposedSelection:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_ShouldSelectTableColumn(handle func(o objc.ProtocolBase, tableView TableView, tableColumn TableColumn) bool) {
	objc.AddMethod(c.class, objc.SEL("tableView:shouldSelectTableColumn:"), handle)
}

func (c *TableViewDelegateCreator) SetTableViewSelectionIsChanging(handle func(o objc.ProtocolBase, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("tableViewSelectionIsChanging:"), handle)
}

func (c *TableViewDelegateCreator) SetTableViewSelectionDidChange(handle func(o objc.ProtocolBase, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("tableViewSelectionDidChange:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_ShouldTypeSelectForEvent_WithCurrentSearchString(handle func(o objc.ProtocolBase, tableView TableView, event Event, searchString string) bool) {
	objc.AddMethod(c.class, objc.SEL("tableView:shouldTypeSelectForEvent:withCurrentSearchString:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_TypeSelectStringForTableColumn_Row(handle func(o objc.ProtocolBase, tableView TableView, tableColumn TableColumn, row int) string) {
	objc.AddMethod(c.class, objc.SEL("tableView:typeSelectStringForTableColumn:row:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_NextTypeSelectMatchFromRow_ToRow_ForString(handle func(o objc.ProtocolBase, tableView TableView, startRow int, endRow int, searchString string) int) {
	objc.AddMethod(c.class, objc.SEL("tableView:nextTypeSelectMatchFromRow:toRow:forString:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_ShouldReorderColumn_ToColumn(handle func(o objc.ProtocolBase, tableView TableView, columnIndex int, newColumnIndex int) bool) {
	objc.AddMethod(c.class, objc.SEL("tableView:shouldReorderColumn:toColumn:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_DidDragTableColumn(handle func(o objc.ProtocolBase, tableView TableView, tableColumn TableColumn)) {
	objc.AddMethod(c.class, objc.SEL("tableView:didDragTableColumn:"), handle)
}

func (c *TableViewDelegateCreator) SetTableViewColumnDidMove(handle func(o objc.ProtocolBase, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("tableViewColumnDidMove:"), handle)
}

func (c *TableViewDelegateCreator) SetTableViewColumnDidResize(handle func(o objc.ProtocolBase, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("tableViewColumnDidResize:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_DidClickTableColumn(handle func(o objc.ProtocolBase, tableView TableView, tableColumn TableColumn)) {
	objc.AddMethod(c.class, objc.SEL("tableView:didClickTableColumn:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_MouseDownInHeaderOfTableColumn(handle func(o objc.ProtocolBase, tableView TableView, tableColumn TableColumn)) {
	objc.AddMethod(c.class, objc.SEL("tableView:mouseDownInHeaderOfTableColumn:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_ShouldTrackCell_ForTableColumn_Row(handle func(o objc.ProtocolBase, tableView TableView, cell Cell, tableColumn TableColumn, row int) bool) {
	objc.AddMethod(c.class, objc.SEL("tableView:shouldTrackCell:forTableColumn:row:"), handle)
}

func (c *TableViewDelegateCreator) SetTableView_RowActionsForRow_Edge(handle func(o objc.ProtocolBase, tableView TableView, row int, edge TableRowActionEdge) []ITableViewRowAction) {
	objc.AddMethod(c.class, objc.SEL("tableView:rowActionsForRow:edge:"), handle)
}

func (c *TableViewDelegateCreator) Create() objc.ProtocolBase {
	return objc.ProtocolBase{Object: c.class.CreateInstance(0)}
}
