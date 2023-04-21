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

type TableViewDelegateImpl struct {
	ControlTextEditingDelegateImpl
	_TableView_ViewForTableColumn_Row                            func(tableView TableView, tableColumn TableColumn, row int) IView
	_TableView_RowViewForRow                                     func(tableView TableView, row int) ITableRowView
	_TableView_DidAddRowView_ForRow                              func(tableView TableView, rowView TableRowView, row int)
	_TableView_DidRemoveRowView_ForRow                           func(tableView TableView, rowView TableRowView, row int)
	_TableView_IsGroupRow                                        func(tableView TableView, row int) bool
	_TableView_WillDisplayCell_ForTableColumn_Row                func(tableView TableView, cell objc.Object, tableColumn TableColumn, row int)
	_TableView_DataCellForTableColumn_Row                        func(tableView TableView, tableColumn TableColumn, row int) ICell
	_TableView_ShouldShowCellExpansionForTableColumn_Row         func(tableView TableView, tableColumn TableColumn, row int) bool
	_TableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation func(tableView TableView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, row int, mouseLocation foundation.Point) string
	_TableView_ShouldEditTableColumn_Row                         func(tableView TableView, tableColumn TableColumn, row int) bool
	_TableView_HeightOfRow                                       func(tableView TableView, row int) float64
	_TableView_SizeToFitWidthOfColumn                            func(tableView TableView, column int) float64
	_SelectionShouldChangeInTableView                            func(tableView TableView) bool
	_TableView_ShouldSelectRow                                   func(tableView TableView, row int) bool
	_TableView_SelectionIndexesForProposedSelection              func(tableView TableView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet
	_TableView_ShouldSelectTableColumn                           func(tableView TableView, tableColumn TableColumn) bool
	_TableViewSelectionIsChanging                                func(notification foundation.Notification)
	_TableViewSelectionDidChange                                 func(notification foundation.Notification)
	_TableView_ShouldTypeSelectForEvent_WithCurrentSearchString  func(tableView TableView, event Event, searchString string) bool
	_TableView_TypeSelectStringForTableColumn_Row                func(tableView TableView, tableColumn TableColumn, row int) string
	_TableView_NextTypeSelectMatchFromRow_ToRow_ForString        func(tableView TableView, startRow int, endRow int, searchString string) int
	_TableView_ShouldReorderColumn_ToColumn                      func(tableView TableView, columnIndex int, newColumnIndex int) bool
	_TableView_DidDragTableColumn                                func(tableView TableView, tableColumn TableColumn)
	_TableViewColumnDidMove                                      func(notification foundation.Notification)
	_TableViewColumnDidResize                                    func(notification foundation.Notification)
	_TableView_DidClickTableColumn                               func(tableView TableView, tableColumn TableColumn)
	_TableView_MouseDownInHeaderOfTableColumn                    func(tableView TableView, tableColumn TableColumn)
	_TableView_ShouldTrackCell_ForTableColumn_Row                func(tableView TableView, cell Cell, tableColumn TableColumn, row int) bool
	_TableView_RowActionsForRow_Edge                             func(tableView TableView, row int, edge TableRowActionEdge) []ITableViewRowAction
}

func (di *TableViewDelegateImpl) ImplementsTableView_ViewForTableColumn_Row() bool {
	return di._TableView_ViewForTableColumn_Row != nil
}

func (di *TableViewDelegateImpl) SetTableView_ViewForTableColumn_Row(f func(tableView TableView, tableColumn TableColumn, row int) IView) {
	di._TableView_ViewForTableColumn_Row = f
}

func (di *TableViewDelegateImpl) TableView_ViewForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) IView {
	return di._TableView_ViewForTableColumn_Row(tableView, tableColumn, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_RowViewForRow() bool {
	return di._TableView_RowViewForRow != nil
}

func (di *TableViewDelegateImpl) SetTableView_RowViewForRow(f func(tableView TableView, row int) ITableRowView) {
	di._TableView_RowViewForRow = f
}

func (di *TableViewDelegateImpl) TableView_RowViewForRow(tableView TableView, row int) ITableRowView {
	return di._TableView_RowViewForRow(tableView, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_DidAddRowView_ForRow() bool {
	return di._TableView_DidAddRowView_ForRow != nil
}

func (di *TableViewDelegateImpl) SetTableView_DidAddRowView_ForRow(f func(tableView TableView, rowView TableRowView, row int)) {
	di._TableView_DidAddRowView_ForRow = f
}

func (di *TableViewDelegateImpl) TableView_DidAddRowView_ForRow(tableView TableView, rowView TableRowView, row int) {
	di._TableView_DidAddRowView_ForRow(tableView, rowView, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_DidRemoveRowView_ForRow() bool {
	return di._TableView_DidRemoveRowView_ForRow != nil
}

func (di *TableViewDelegateImpl) SetTableView_DidRemoveRowView_ForRow(f func(tableView TableView, rowView TableRowView, row int)) {
	di._TableView_DidRemoveRowView_ForRow = f
}

func (di *TableViewDelegateImpl) TableView_DidRemoveRowView_ForRow(tableView TableView, rowView TableRowView, row int) {
	di._TableView_DidRemoveRowView_ForRow(tableView, rowView, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_IsGroupRow() bool {
	return di._TableView_IsGroupRow != nil
}

func (di *TableViewDelegateImpl) SetTableView_IsGroupRow(f func(tableView TableView, row int) bool) {
	di._TableView_IsGroupRow = f
}

func (di *TableViewDelegateImpl) TableView_IsGroupRow(tableView TableView, row int) bool {
	return di._TableView_IsGroupRow(tableView, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_WillDisplayCell_ForTableColumn_Row() bool {
	return di._TableView_WillDisplayCell_ForTableColumn_Row != nil
}

func (di *TableViewDelegateImpl) SetTableView_WillDisplayCell_ForTableColumn_Row(f func(tableView TableView, cell objc.Object, tableColumn TableColumn, row int)) {
	di._TableView_WillDisplayCell_ForTableColumn_Row = f
}

func (di *TableViewDelegateImpl) TableView_WillDisplayCell_ForTableColumn_Row(tableView TableView, cell objc.Object, tableColumn TableColumn, row int) {
	di._TableView_WillDisplayCell_ForTableColumn_Row(tableView, cell, tableColumn, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_DataCellForTableColumn_Row() bool {
	return di._TableView_DataCellForTableColumn_Row != nil
}

func (di *TableViewDelegateImpl) SetTableView_DataCellForTableColumn_Row(f func(tableView TableView, tableColumn TableColumn, row int) ICell) {
	di._TableView_DataCellForTableColumn_Row = f
}

func (di *TableViewDelegateImpl) TableView_DataCellForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) ICell {
	return di._TableView_DataCellForTableColumn_Row(tableView, tableColumn, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_ShouldShowCellExpansionForTableColumn_Row() bool {
	return di._TableView_ShouldShowCellExpansionForTableColumn_Row != nil
}

func (di *TableViewDelegateImpl) SetTableView_ShouldShowCellExpansionForTableColumn_Row(f func(tableView TableView, tableColumn TableColumn, row int) bool) {
	di._TableView_ShouldShowCellExpansionForTableColumn_Row = f
}

func (di *TableViewDelegateImpl) TableView_ShouldShowCellExpansionForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) bool {
	return di._TableView_ShouldShowCellExpansionForTableColumn_Row(tableView, tableColumn, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation() bool {
	return di._TableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation != nil
}

func (di *TableViewDelegateImpl) SetTableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation(f func(tableView TableView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, row int, mouseLocation foundation.Point) string) {
	di._TableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation = f
}

func (di *TableViewDelegateImpl) TableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation(tableView TableView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, row int, mouseLocation foundation.Point) string {
	return di._TableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation(tableView, cell, rect, tableColumn, row, mouseLocation)
}
func (di *TableViewDelegateImpl) ImplementsTableView_ShouldEditTableColumn_Row() bool {
	return di._TableView_ShouldEditTableColumn_Row != nil
}

func (di *TableViewDelegateImpl) SetTableView_ShouldEditTableColumn_Row(f func(tableView TableView, tableColumn TableColumn, row int) bool) {
	di._TableView_ShouldEditTableColumn_Row = f
}

func (di *TableViewDelegateImpl) TableView_ShouldEditTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) bool {
	return di._TableView_ShouldEditTableColumn_Row(tableView, tableColumn, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_HeightOfRow() bool {
	return di._TableView_HeightOfRow != nil
}

func (di *TableViewDelegateImpl) SetTableView_HeightOfRow(f func(tableView TableView, row int) float64) {
	di._TableView_HeightOfRow = f
}

func (di *TableViewDelegateImpl) TableView_HeightOfRow(tableView TableView, row int) float64 {
	return di._TableView_HeightOfRow(tableView, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_SizeToFitWidthOfColumn() bool {
	return di._TableView_SizeToFitWidthOfColumn != nil
}

func (di *TableViewDelegateImpl) SetTableView_SizeToFitWidthOfColumn(f func(tableView TableView, column int) float64) {
	di._TableView_SizeToFitWidthOfColumn = f
}

func (di *TableViewDelegateImpl) TableView_SizeToFitWidthOfColumn(tableView TableView, column int) float64 {
	return di._TableView_SizeToFitWidthOfColumn(tableView, column)
}
func (di *TableViewDelegateImpl) ImplementsSelectionShouldChangeInTableView() bool {
	return di._SelectionShouldChangeInTableView != nil
}

func (di *TableViewDelegateImpl) SetSelectionShouldChangeInTableView(f func(tableView TableView) bool) {
	di._SelectionShouldChangeInTableView = f
}

func (di *TableViewDelegateImpl) SelectionShouldChangeInTableView(tableView TableView) bool {
	return di._SelectionShouldChangeInTableView(tableView)
}
func (di *TableViewDelegateImpl) ImplementsTableView_ShouldSelectRow() bool {
	return di._TableView_ShouldSelectRow != nil
}

func (di *TableViewDelegateImpl) SetTableView_ShouldSelectRow(f func(tableView TableView, row int) bool) {
	di._TableView_ShouldSelectRow = f
}

func (di *TableViewDelegateImpl) TableView_ShouldSelectRow(tableView TableView, row int) bool {
	return di._TableView_ShouldSelectRow(tableView, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_SelectionIndexesForProposedSelection() bool {
	return di._TableView_SelectionIndexesForProposedSelection != nil
}

func (di *TableViewDelegateImpl) SetTableView_SelectionIndexesForProposedSelection(f func(tableView TableView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet) {
	di._TableView_SelectionIndexesForProposedSelection = f
}

func (di *TableViewDelegateImpl) TableView_SelectionIndexesForProposedSelection(tableView TableView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet {
	return di._TableView_SelectionIndexesForProposedSelection(tableView, proposedSelectionIndexes)
}
func (di *TableViewDelegateImpl) ImplementsTableView_ShouldSelectTableColumn() bool {
	return di._TableView_ShouldSelectTableColumn != nil
}

func (di *TableViewDelegateImpl) SetTableView_ShouldSelectTableColumn(f func(tableView TableView, tableColumn TableColumn) bool) {
	di._TableView_ShouldSelectTableColumn = f
}

func (di *TableViewDelegateImpl) TableView_ShouldSelectTableColumn(tableView TableView, tableColumn TableColumn) bool {
	return di._TableView_ShouldSelectTableColumn(tableView, tableColumn)
}
func (di *TableViewDelegateImpl) ImplementsTableViewSelectionIsChanging() bool {
	return di._TableViewSelectionIsChanging != nil
}

func (di *TableViewDelegateImpl) SetTableViewSelectionIsChanging(f func(notification foundation.Notification)) {
	di._TableViewSelectionIsChanging = f
}

func (di *TableViewDelegateImpl) TableViewSelectionIsChanging(notification foundation.Notification) {
	di._TableViewSelectionIsChanging(notification)
}
func (di *TableViewDelegateImpl) ImplementsTableViewSelectionDidChange() bool {
	return di._TableViewSelectionDidChange != nil
}

func (di *TableViewDelegateImpl) SetTableViewSelectionDidChange(f func(notification foundation.Notification)) {
	di._TableViewSelectionDidChange = f
}

func (di *TableViewDelegateImpl) TableViewSelectionDidChange(notification foundation.Notification) {
	di._TableViewSelectionDidChange(notification)
}
func (di *TableViewDelegateImpl) ImplementsTableView_ShouldTypeSelectForEvent_WithCurrentSearchString() bool {
	return di._TableView_ShouldTypeSelectForEvent_WithCurrentSearchString != nil
}

func (di *TableViewDelegateImpl) SetTableView_ShouldTypeSelectForEvent_WithCurrentSearchString(f func(tableView TableView, event Event, searchString string) bool) {
	di._TableView_ShouldTypeSelectForEvent_WithCurrentSearchString = f
}

func (di *TableViewDelegateImpl) TableView_ShouldTypeSelectForEvent_WithCurrentSearchString(tableView TableView, event Event, searchString string) bool {
	return di._TableView_ShouldTypeSelectForEvent_WithCurrentSearchString(tableView, event, searchString)
}
func (di *TableViewDelegateImpl) ImplementsTableView_TypeSelectStringForTableColumn_Row() bool {
	return di._TableView_TypeSelectStringForTableColumn_Row != nil
}

func (di *TableViewDelegateImpl) SetTableView_TypeSelectStringForTableColumn_Row(f func(tableView TableView, tableColumn TableColumn, row int) string) {
	di._TableView_TypeSelectStringForTableColumn_Row = f
}

func (di *TableViewDelegateImpl) TableView_TypeSelectStringForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) string {
	return di._TableView_TypeSelectStringForTableColumn_Row(tableView, tableColumn, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_NextTypeSelectMatchFromRow_ToRow_ForString() bool {
	return di._TableView_NextTypeSelectMatchFromRow_ToRow_ForString != nil
}

func (di *TableViewDelegateImpl) SetTableView_NextTypeSelectMatchFromRow_ToRow_ForString(f func(tableView TableView, startRow int, endRow int, searchString string) int) {
	di._TableView_NextTypeSelectMatchFromRow_ToRow_ForString = f
}

func (di *TableViewDelegateImpl) TableView_NextTypeSelectMatchFromRow_ToRow_ForString(tableView TableView, startRow int, endRow int, searchString string) int {
	return di._TableView_NextTypeSelectMatchFromRow_ToRow_ForString(tableView, startRow, endRow, searchString)
}
func (di *TableViewDelegateImpl) ImplementsTableView_ShouldReorderColumn_ToColumn() bool {
	return di._TableView_ShouldReorderColumn_ToColumn != nil
}

func (di *TableViewDelegateImpl) SetTableView_ShouldReorderColumn_ToColumn(f func(tableView TableView, columnIndex int, newColumnIndex int) bool) {
	di._TableView_ShouldReorderColumn_ToColumn = f
}

func (di *TableViewDelegateImpl) TableView_ShouldReorderColumn_ToColumn(tableView TableView, columnIndex int, newColumnIndex int) bool {
	return di._TableView_ShouldReorderColumn_ToColumn(tableView, columnIndex, newColumnIndex)
}
func (di *TableViewDelegateImpl) ImplementsTableView_DidDragTableColumn() bool {
	return di._TableView_DidDragTableColumn != nil
}

func (di *TableViewDelegateImpl) SetTableView_DidDragTableColumn(f func(tableView TableView, tableColumn TableColumn)) {
	di._TableView_DidDragTableColumn = f
}

func (di *TableViewDelegateImpl) TableView_DidDragTableColumn(tableView TableView, tableColumn TableColumn) {
	di._TableView_DidDragTableColumn(tableView, tableColumn)
}
func (di *TableViewDelegateImpl) ImplementsTableViewColumnDidMove() bool {
	return di._TableViewColumnDidMove != nil
}

func (di *TableViewDelegateImpl) SetTableViewColumnDidMove(f func(notification foundation.Notification)) {
	di._TableViewColumnDidMove = f
}

func (di *TableViewDelegateImpl) TableViewColumnDidMove(notification foundation.Notification) {
	di._TableViewColumnDidMove(notification)
}
func (di *TableViewDelegateImpl) ImplementsTableViewColumnDidResize() bool {
	return di._TableViewColumnDidResize != nil
}

func (di *TableViewDelegateImpl) SetTableViewColumnDidResize(f func(notification foundation.Notification)) {
	di._TableViewColumnDidResize = f
}

func (di *TableViewDelegateImpl) TableViewColumnDidResize(notification foundation.Notification) {
	di._TableViewColumnDidResize(notification)
}
func (di *TableViewDelegateImpl) ImplementsTableView_DidClickTableColumn() bool {
	return di._TableView_DidClickTableColumn != nil
}

func (di *TableViewDelegateImpl) SetTableView_DidClickTableColumn(f func(tableView TableView, tableColumn TableColumn)) {
	di._TableView_DidClickTableColumn = f
}

func (di *TableViewDelegateImpl) TableView_DidClickTableColumn(tableView TableView, tableColumn TableColumn) {
	di._TableView_DidClickTableColumn(tableView, tableColumn)
}
func (di *TableViewDelegateImpl) ImplementsTableView_MouseDownInHeaderOfTableColumn() bool {
	return di._TableView_MouseDownInHeaderOfTableColumn != nil
}

func (di *TableViewDelegateImpl) SetTableView_MouseDownInHeaderOfTableColumn(f func(tableView TableView, tableColumn TableColumn)) {
	di._TableView_MouseDownInHeaderOfTableColumn = f
}

func (di *TableViewDelegateImpl) TableView_MouseDownInHeaderOfTableColumn(tableView TableView, tableColumn TableColumn) {
	di._TableView_MouseDownInHeaderOfTableColumn(tableView, tableColumn)
}
func (di *TableViewDelegateImpl) ImplementsTableView_ShouldTrackCell_ForTableColumn_Row() bool {
	return di._TableView_ShouldTrackCell_ForTableColumn_Row != nil
}

func (di *TableViewDelegateImpl) SetTableView_ShouldTrackCell_ForTableColumn_Row(f func(tableView TableView, cell Cell, tableColumn TableColumn, row int) bool) {
	di._TableView_ShouldTrackCell_ForTableColumn_Row = f
}

func (di *TableViewDelegateImpl) TableView_ShouldTrackCell_ForTableColumn_Row(tableView TableView, cell Cell, tableColumn TableColumn, row int) bool {
	return di._TableView_ShouldTrackCell_ForTableColumn_Row(tableView, cell, tableColumn, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_RowActionsForRow_Edge() bool {
	return di._TableView_RowActionsForRow_Edge != nil
}

func (di *TableViewDelegateImpl) SetTableView_RowActionsForRow_Edge(f func(tableView TableView, row int, edge TableRowActionEdge) []ITableViewRowAction) {
	di._TableView_RowActionsForRow_Edge = f
}

func (di *TableViewDelegateImpl) TableView_RowActionsForRow_Edge(tableView TableView, row int, edge TableRowActionEdge) []ITableViewRowAction {
	return di._TableView_RowActionsForRow_Edge(tableView, row, edge)
}

type TableViewDelegateWrapper struct {
	ControlTextEditingDelegateWrapper
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_ViewForTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:viewForTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableView_ViewForTableColumn_Row(tableView ITableView, tableColumn ITableColumn, row int) View {
	rv := objc.CallMethod[View](t_, "tableView:viewForTableColumn:row:", tableView, tableColumn, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_RowViewForRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:rowViewForRow:"))
}

func (t_ TableViewDelegateWrapper) TableView_RowViewForRow(tableView ITableView, row int) TableRowView {
	rv := objc.CallMethod[TableRowView](t_, "tableView:rowViewForRow:", tableView, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_DidAddRowView_ForRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:didAddRowView:forRow:"))
}

func (t_ TableViewDelegateWrapper) TableView_DidAddRowView_ForRow(tableView ITableView, rowView ITableRowView, row int) {
	objc.CallMethod[objc.Void](t_, "tableView:didAddRowView:forRow:", tableView, rowView, row)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_DidRemoveRowView_ForRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:didRemoveRowView:forRow:"))
}

func (t_ TableViewDelegateWrapper) TableView_DidRemoveRowView_ForRow(tableView ITableView, rowView ITableRowView, row int) {
	objc.CallMethod[objc.Void](t_, "tableView:didRemoveRowView:forRow:", tableView, rowView, row)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_IsGroupRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:isGroupRow:"))
}

func (t_ TableViewDelegateWrapper) TableView_IsGroupRow(tableView ITableView, row int) bool {
	rv := objc.CallMethod[bool](t_, "tableView:isGroupRow:", tableView, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_WillDisplayCell_ForTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:willDisplayCell:forTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableView_WillDisplayCell_ForTableColumn_Row(tableView ITableView, cell objc.IObject, tableColumn ITableColumn, row int) {
	objc.CallMethod[objc.Void](t_, "tableView:willDisplayCell:forTableColumn:row:", tableView, cell, tableColumn, row)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_DataCellForTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:dataCellForTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableView_DataCellForTableColumn_Row(tableView ITableView, tableColumn ITableColumn, row int) Cell {
	rv := objc.CallMethod[Cell](t_, "tableView:dataCellForTableColumn:row:", tableView, tableColumn, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_ShouldShowCellExpansionForTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldShowCellExpansionForTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableView_ShouldShowCellExpansionForTableColumn_Row(tableView ITableView, tableColumn ITableColumn, row int) bool {
	rv := objc.CallMethod[bool](t_, "tableView:shouldShowCellExpansionForTableColumn:row:", tableView, tableColumn, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:toolTipForCell:rect:tableColumn:row:mouseLocation:"))
}

func (t_ TableViewDelegateWrapper) TableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation(tableView ITableView, cell ICell, rect *foundation.Rect, tableColumn ITableColumn, row int, mouseLocation foundation.Point) string {
	rv := objc.CallMethod[string](t_, "tableView:toolTipForCell:rect:tableColumn:row:mouseLocation:", tableView, cell, rect, tableColumn, row, mouseLocation)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_ShouldEditTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldEditTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableView_ShouldEditTableColumn_Row(tableView ITableView, tableColumn ITableColumn, row int) bool {
	rv := objc.CallMethod[bool](t_, "tableView:shouldEditTableColumn:row:", tableView, tableColumn, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_HeightOfRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:heightOfRow:"))
}

func (t_ TableViewDelegateWrapper) TableView_HeightOfRow(tableView ITableView, row int) float64 {
	rv := objc.CallMethod[float64](t_, "tableView:heightOfRow:", tableView, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_SizeToFitWidthOfColumn() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:sizeToFitWidthOfColumn:"))
}

func (t_ TableViewDelegateWrapper) TableView_SizeToFitWidthOfColumn(tableView ITableView, column int) float64 {
	rv := objc.CallMethod[float64](t_, "tableView:sizeToFitWidthOfColumn:", tableView, column)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsSelectionShouldChangeInTableView() bool {
	return t_.RespondsToSelector(objc.GetSelector("selectionShouldChangeInTableView:"))
}

func (t_ TableViewDelegateWrapper) SelectionShouldChangeInTableView(tableView ITableView) bool {
	rv := objc.CallMethod[bool](t_, "selectionShouldChangeInTableView:", tableView)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_ShouldSelectRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldSelectRow:"))
}

func (t_ TableViewDelegateWrapper) TableView_ShouldSelectRow(tableView ITableView, row int) bool {
	rv := objc.CallMethod[bool](t_, "tableView:shouldSelectRow:", tableView, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_SelectionIndexesForProposedSelection() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:selectionIndexesForProposedSelection:"))
}

func (t_ TableViewDelegateWrapper) TableView_SelectionIndexesForProposedSelection(tableView ITableView, proposedSelectionIndexes foundation.IIndexSet) foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](t_, "tableView:selectionIndexesForProposedSelection:", tableView, proposedSelectionIndexes)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_ShouldSelectTableColumn() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldSelectTableColumn:"))
}

func (t_ TableViewDelegateWrapper) TableView_ShouldSelectTableColumn(tableView ITableView, tableColumn ITableColumn) bool {
	rv := objc.CallMethod[bool](t_, "tableView:shouldSelectTableColumn:", tableView, tableColumn)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableViewSelectionIsChanging() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableViewSelectionIsChanging:"))
}

func (t_ TableViewDelegateWrapper) TableViewSelectionIsChanging(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, "tableViewSelectionIsChanging:", notification)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableViewSelectionDidChange() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableViewSelectionDidChange:"))
}

func (t_ TableViewDelegateWrapper) TableViewSelectionDidChange(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, "tableViewSelectionDidChange:", notification)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_ShouldTypeSelectForEvent_WithCurrentSearchString() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldTypeSelectForEvent:withCurrentSearchString:"))
}

func (t_ TableViewDelegateWrapper) TableView_ShouldTypeSelectForEvent_WithCurrentSearchString(tableView ITableView, event IEvent, searchString string) bool {
	rv := objc.CallMethod[bool](t_, "tableView:shouldTypeSelectForEvent:withCurrentSearchString:", tableView, event, searchString)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_TypeSelectStringForTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:typeSelectStringForTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableView_TypeSelectStringForTableColumn_Row(tableView ITableView, tableColumn ITableColumn, row int) string {
	rv := objc.CallMethod[string](t_, "tableView:typeSelectStringForTableColumn:row:", tableView, tableColumn, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_NextTypeSelectMatchFromRow_ToRow_ForString() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:nextTypeSelectMatchFromRow:toRow:forString:"))
}

func (t_ TableViewDelegateWrapper) TableView_NextTypeSelectMatchFromRow_ToRow_ForString(tableView ITableView, startRow int, endRow int, searchString string) int {
	rv := objc.CallMethod[int](t_, "tableView:nextTypeSelectMatchFromRow:toRow:forString:", tableView, startRow, endRow, searchString)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_ShouldReorderColumn_ToColumn() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldReorderColumn:toColumn:"))
}

func (t_ TableViewDelegateWrapper) TableView_ShouldReorderColumn_ToColumn(tableView ITableView, columnIndex int, newColumnIndex int) bool {
	rv := objc.CallMethod[bool](t_, "tableView:shouldReorderColumn:toColumn:", tableView, columnIndex, newColumnIndex)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_DidDragTableColumn() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:didDragTableColumn:"))
}

func (t_ TableViewDelegateWrapper) TableView_DidDragTableColumn(tableView ITableView, tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](t_, "tableView:didDragTableColumn:", tableView, tableColumn)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableViewColumnDidMove() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableViewColumnDidMove:"))
}

func (t_ TableViewDelegateWrapper) TableViewColumnDidMove(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, "tableViewColumnDidMove:", notification)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableViewColumnDidResize() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableViewColumnDidResize:"))
}

func (t_ TableViewDelegateWrapper) TableViewColumnDidResize(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, "tableViewColumnDidResize:", notification)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_DidClickTableColumn() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:didClickTableColumn:"))
}

func (t_ TableViewDelegateWrapper) TableView_DidClickTableColumn(tableView ITableView, tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](t_, "tableView:didClickTableColumn:", tableView, tableColumn)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_MouseDownInHeaderOfTableColumn() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:mouseDownInHeaderOfTableColumn:"))
}

func (t_ TableViewDelegateWrapper) TableView_MouseDownInHeaderOfTableColumn(tableView ITableView, tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](t_, "tableView:mouseDownInHeaderOfTableColumn:", tableView, tableColumn)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_ShouldTrackCell_ForTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldTrackCell:forTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableView_ShouldTrackCell_ForTableColumn_Row(tableView ITableView, cell ICell, tableColumn ITableColumn, row int) bool {
	rv := objc.CallMethod[bool](t_, "tableView:shouldTrackCell:forTableColumn:row:", tableView, cell, tableColumn, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_RowActionsForRow_Edge() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:rowActionsForRow:edge:"))
}

func (t_ TableViewDelegateWrapper) TableView_RowActionsForRow_Edge(tableView ITableView, row int, edge TableRowActionEdge) []TableViewRowAction {
	rv := objc.CallMethod[[]TableViewRowAction](t_, "tableView:rowActionsForRow:edge:", tableView, row, edge)
	return rv
}
