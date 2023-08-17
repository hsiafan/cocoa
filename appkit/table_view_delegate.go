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

type TableViewDelegateWrapper struct {
	objc.Object
}

func (t_ TableViewDelegateWrapper) TableView_ViewForTableColumn_Row(tableView ITableView, tableColumn ITableColumn, row int) View {
	rv := objc.CallMethod[View](t_, objc.GetSelector("tableView:viewForTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn), row)
	return rv
}

func (t_ TableViewDelegateWrapper) TableView_RowViewForRow(tableView ITableView, row int) TableRowView {
	rv := objc.CallMethod[TableRowView](t_, objc.GetSelector("tableView:rowViewForRow:"), objc.ExtractPtr(tableView), row)
	return rv
}

func (t_ TableViewDelegateWrapper) TableView_DidAddRowView_ForRow(tableView ITableView, rowView ITableRowView, row int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:didAddRowView:forRow:"), objc.ExtractPtr(tableView), objc.ExtractPtr(rowView), row)
}

func (t_ TableViewDelegateWrapper) TableView_DidRemoveRowView_ForRow(tableView ITableView, rowView ITableRowView, row int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:didRemoveRowView:forRow:"), objc.ExtractPtr(tableView), objc.ExtractPtr(rowView), row)
}

func (t_ TableViewDelegateWrapper) TableView_IsGroupRow(tableView ITableView, row int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:isGroupRow:"), objc.ExtractPtr(tableView), row)
	return rv
}

func (t_ TableViewDelegateWrapper) TableView_WillDisplayCell_ForTableColumn_Row(tableView ITableView, cell objc.IObject, tableColumn ITableColumn, row int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:willDisplayCell:forTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(cell), objc.ExtractPtr(tableColumn), row)
}

func (t_ TableViewDelegateWrapper) TableView_DataCellForTableColumn_Row(tableView ITableView, tableColumn ITableColumn, row int) Cell {
	rv := objc.CallMethod[Cell](t_, objc.GetSelector("tableView:dataCellForTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn), row)
	return rv
}

func (t_ TableViewDelegateWrapper) TableView_ShouldShowCellExpansionForTableColumn_Row(tableView ITableView, tableColumn ITableColumn, row int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:shouldShowCellExpansionForTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn), row)
	return rv
}

func (t_ TableViewDelegateWrapper) TableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation(tableView ITableView, cell ICell, rect *foundation.Rect, tableColumn ITableColumn, row int, mouseLocation foundation.Point) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("tableView:toolTipForCell:rect:tableColumn:row:mouseLocation:"), objc.ExtractPtr(tableView), objc.ExtractPtr(cell), rect, objc.ExtractPtr(tableColumn), row, mouseLocation)
	return rv
}

func (t_ TableViewDelegateWrapper) TableView_ShouldEditTableColumn_Row(tableView ITableView, tableColumn ITableColumn, row int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:shouldEditTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn), row)
	return rv
}

func (t_ TableViewDelegateWrapper) TableView_HeightOfRow(tableView ITableView, row int) float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("tableView:heightOfRow:"), objc.ExtractPtr(tableView), row)
	return rv
}

func (t_ TableViewDelegateWrapper) TableView_SizeToFitWidthOfColumn(tableView ITableView, column int) float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("tableView:sizeToFitWidthOfColumn:"), objc.ExtractPtr(tableView), column)
	return rv
}

func (t_ TableViewDelegateWrapper) SelectionShouldChangeInTableView(tableView ITableView) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("selectionShouldChangeInTableView:"), objc.ExtractPtr(tableView))
	return rv
}

func (t_ TableViewDelegateWrapper) TableView_ShouldSelectRow(tableView ITableView, row int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:shouldSelectRow:"), objc.ExtractPtr(tableView), row)
	return rv
}

func (t_ TableViewDelegateWrapper) TableView_SelectionIndexesForProposedSelection(tableView ITableView, proposedSelectionIndexes foundation.IIndexSet) foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](t_, objc.GetSelector("tableView:selectionIndexesForProposedSelection:"), objc.ExtractPtr(tableView), objc.ExtractPtr(proposedSelectionIndexes))
	return rv
}

func (t_ TableViewDelegateWrapper) TableView_ShouldSelectTableColumn(tableView ITableView, tableColumn ITableColumn) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:shouldSelectTableColumn:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn))
	return rv
}

func (t_ TableViewDelegateWrapper) TableViewSelectionIsChanging(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableViewSelectionIsChanging:"), objc.ExtractPtr(notification))
}

func (t_ TableViewDelegateWrapper) TableViewSelectionDidChange(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableViewSelectionDidChange:"), objc.ExtractPtr(notification))
}

func (t_ TableViewDelegateWrapper) TableView_ShouldTypeSelectForEvent_WithCurrentSearchString(tableView ITableView, event IEvent, searchString string) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:shouldTypeSelectForEvent:withCurrentSearchString:"), objc.ExtractPtr(tableView), objc.ExtractPtr(event), searchString)
	return rv
}

func (t_ TableViewDelegateWrapper) TableView_TypeSelectStringForTableColumn_Row(tableView ITableView, tableColumn ITableColumn, row int) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("tableView:typeSelectStringForTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn), row)
	return rv
}

func (t_ TableViewDelegateWrapper) TableView_NextTypeSelectMatchFromRow_ToRow_ForString(tableView ITableView, startRow int, endRow int, searchString string) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("tableView:nextTypeSelectMatchFromRow:toRow:forString:"), objc.ExtractPtr(tableView), startRow, endRow, searchString)
	return rv
}

func (t_ TableViewDelegateWrapper) TableView_ShouldReorderColumn_ToColumn(tableView ITableView, columnIndex int, newColumnIndex int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:shouldReorderColumn:toColumn:"), objc.ExtractPtr(tableView), columnIndex, newColumnIndex)
	return rv
}

func (t_ TableViewDelegateWrapper) TableView_DidDragTableColumn(tableView ITableView, tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:didDragTableColumn:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn))
}

func (t_ TableViewDelegateWrapper) TableViewColumnDidMove(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableViewColumnDidMove:"), objc.ExtractPtr(notification))
}

func (t_ TableViewDelegateWrapper) TableViewColumnDidResize(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableViewColumnDidResize:"), objc.ExtractPtr(notification))
}

func (t_ TableViewDelegateWrapper) TableView_DidClickTableColumn(tableView ITableView, tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:didClickTableColumn:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn))
}

func (t_ TableViewDelegateWrapper) TableView_MouseDownInHeaderOfTableColumn(tableView ITableView, tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:mouseDownInHeaderOfTableColumn:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn))
}

func (t_ TableViewDelegateWrapper) TableView_ShouldTrackCell_ForTableColumn_Row(tableView ITableView, cell ICell, tableColumn ITableColumn, row int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:shouldTrackCell:forTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(cell), objc.ExtractPtr(tableColumn), row)
	return rv
}

func (t_ TableViewDelegateWrapper) TableView_RowActionsForRow_Edge(tableView ITableView, row int, edge TableRowActionEdge) []TableViewRowAction {
	rv := objc.CallMethod[[]TableViewRowAction](t_, objc.GetSelector("tableView:rowActionsForRow:edge:"), objc.ExtractPtr(tableView), row, edge)
	return rv
}

func (t_ TableViewDelegateWrapper) Control_IsValidObject(control IControl, obj objc.IObject) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("control:isValidObject:"), objc.ExtractPtr(control), objc.ExtractPtr(obj))
	return rv
}

func (t_ TableViewDelegateWrapper) Control_DidFailToValidatePartialString_ErrorDescription(control IControl, string_ string, error string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("control:didFailToValidatePartialString:errorDescription:"), objc.ExtractPtr(control), string_, error)
}

func (t_ TableViewDelegateWrapper) Control_DidFailToFormatString_ErrorDescription(control IControl, string_ string, error string) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("control:didFailToFormatString:errorDescription:"), objc.ExtractPtr(control), string_, error)
	return rv
}

func (t_ TableViewDelegateWrapper) Control_TextShouldBeginEditing(control IControl, fieldEditor IText) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("control:textShouldBeginEditing:"), objc.ExtractPtr(control), objc.ExtractPtr(fieldEditor))
	return rv
}

func (t_ TableViewDelegateWrapper) Control_TextShouldEndEditing(control IControl, fieldEditor IText) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("control:textShouldEndEditing:"), objc.ExtractPtr(control), objc.ExtractPtr(fieldEditor))
	return rv
}

func (t_ TableViewDelegateWrapper) Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(control IControl, textView ITextView, words []string, charRange foundation.Range, index *int) []string {
	rv := objc.CallMethod[[]string](t_, objc.GetSelector("control:textView:completions:forPartialWordRange:indexOfSelectedItem:"), objc.ExtractPtr(control), objc.ExtractPtr(textView), words, charRange, index)
	return rv
}

func (t_ TableViewDelegateWrapper) Control_TextView_DoCommandBySelector(control IControl, textView ITextView, commandSelector objc.Selector) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("control:textView:doCommandBySelector:"), objc.ExtractPtr(control), objc.ExtractPtr(textView), commandSelector)
	return rv
}

func (t_ TableViewDelegateWrapper) ControlTextDidBeginEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("controlTextDidBeginEditing:"), objc.ExtractPtr(obj))
}

func (t_ TableViewDelegateWrapper) ControlTextDidChange(obj foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("controlTextDidChange:"), objc.ExtractPtr(obj))
}

func (t_ TableViewDelegateWrapper) ControlTextDidEndEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("controlTextDidEndEditing:"), objc.ExtractPtr(obj))
}
