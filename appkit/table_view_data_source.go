// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type TableViewDataSource interface {
	ImplementsNumberOfRowsInTableView() bool
	// optional
	NumberOfRowsInTableView(tableView TableView) int
	ImplementsTableView_ObjectValueForTableColumn_Row() bool
	// optional
	TableView_ObjectValueForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) objc.IObject
	ImplementsTableView_SetObjectValue_ForTableColumn_Row() bool
	// optional
	TableView_SetObjectValue_ForTableColumn_Row(tableView TableView, object objc.Object, tableColumn TableColumn, row int)
	ImplementsTableView_PasteboardWriterForRow() bool
	// optional
	TableView_PasteboardWriterForRow(tableView TableView, row int) objc.IObject
	ImplementsTableView_AcceptDrop_Row_DropOperation() bool
	// optional
	TableView_AcceptDrop_Row_DropOperation(tableView TableView, info objc.Object, row int, dropOperation TableViewDropOperation) bool
	ImplementsTableView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes() bool
	// optional
	// deprecated
	TableView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes(tableView TableView, dropDestination foundation.URL, indexSet foundation.IndexSet) []string
	ImplementsTableView_ValidateDrop_ProposedRow_ProposedDropOperation() bool
	// optional
	TableView_ValidateDrop_ProposedRow_ProposedDropOperation(tableView TableView, info objc.Object, row int, dropOperation TableViewDropOperation) DragOperation
	ImplementsTableView_WriteRowsWithIndexes_ToPasteboard() bool
	// optional
	// deprecated
	TableView_WriteRowsWithIndexes_ToPasteboard(tableView TableView, rowIndexes foundation.IndexSet, pboard Pasteboard) bool
	ImplementsTableView_DraggingSession_WillBeginAtPoint_ForRowIndexes() bool
	// optional
	TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes(tableView TableView, session DraggingSession, screenPoint foundation.Point, rowIndexes foundation.IndexSet)
	ImplementsTableView_UpdateDraggingItemsForDrag() bool
	// optional
	TableView_UpdateDraggingItemsForDrag(tableView TableView, draggingInfo objc.Object)
	ImplementsTableView_DraggingSession_EndedAtPoint_Operation() bool
	// optional
	TableView_DraggingSession_EndedAtPoint_Operation(tableView TableView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	ImplementsTableView_SortDescriptorsDidChange() bool
	// optional
	TableView_SortDescriptorsDidChange(tableView TableView, oldDescriptors []foundation.SortDescriptor)
}

func WrapTableViewDataSource(v TableViewDataSource) objc.Object {
	return objc.WrapAsProtocol("NSTableViewDataSource", v)
}

type TableViewDataSourceBase struct {
}

func (p *TableViewDataSourceBase) ImplementsNumberOfRowsInTableView() bool {
	return false
}

func (p *TableViewDataSourceBase) NumberOfRowsInTableView(tableView TableView) int {
	panic("unimpemented protocol method")
}

func (p *TableViewDataSourceBase) ImplementsTableView_ObjectValueForTableColumn_Row() bool {
	return false
}

func (p *TableViewDataSourceBase) TableView_ObjectValueForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) objc.IObject {
	panic("unimpemented protocol method")
}

func (p *TableViewDataSourceBase) ImplementsTableView_SetObjectValue_ForTableColumn_Row() bool {
	return false
}

func (p *TableViewDataSourceBase) TableView_SetObjectValue_ForTableColumn_Row(tableView TableView, object objc.Object, tableColumn TableColumn, row int) {
	panic("unimpemented protocol method")
}

func (p *TableViewDataSourceBase) ImplementsTableView_PasteboardWriterForRow() bool {
	return false
}

func (p *TableViewDataSourceBase) TableView_PasteboardWriterForRow(tableView TableView, row int) objc.IObject {
	panic("unimpemented protocol method")
}

func (p *TableViewDataSourceBase) ImplementsTableView_AcceptDrop_Row_DropOperation() bool {
	return false
}

func (p *TableViewDataSourceBase) TableView_AcceptDrop_Row_DropOperation(tableView TableView, info objc.Object, row int, dropOperation TableViewDropOperation) bool {
	panic("unimpemented protocol method")
}

func (p *TableViewDataSourceBase) ImplementsTableView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes() bool {
	return false
}

// deprecated
func (p *TableViewDataSourceBase) TableView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes(tableView TableView, dropDestination foundation.URL, indexSet foundation.IndexSet) []string {
	panic("unimpemented protocol method")
}

func (p *TableViewDataSourceBase) ImplementsTableView_ValidateDrop_ProposedRow_ProposedDropOperation() bool {
	return false
}

func (p *TableViewDataSourceBase) TableView_ValidateDrop_ProposedRow_ProposedDropOperation(tableView TableView, info objc.Object, row int, dropOperation TableViewDropOperation) DragOperation {
	panic("unimpemented protocol method")
}

func (p *TableViewDataSourceBase) ImplementsTableView_WriteRowsWithIndexes_ToPasteboard() bool {
	return false
}

// deprecated
func (p *TableViewDataSourceBase) TableView_WriteRowsWithIndexes_ToPasteboard(tableView TableView, rowIndexes foundation.IndexSet, pboard Pasteboard) bool {
	panic("unimpemented protocol method")
}

func (p *TableViewDataSourceBase) ImplementsTableView_DraggingSession_WillBeginAtPoint_ForRowIndexes() bool {
	return false
}

func (p *TableViewDataSourceBase) TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes(tableView TableView, session DraggingSession, screenPoint foundation.Point, rowIndexes foundation.IndexSet) {
	panic("unimpemented protocol method")
}

func (p *TableViewDataSourceBase) ImplementsTableView_UpdateDraggingItemsForDrag() bool {
	return false
}

func (p *TableViewDataSourceBase) TableView_UpdateDraggingItemsForDrag(tableView TableView, draggingInfo objc.Object) {
	panic("unimpemented protocol method")
}

func (p *TableViewDataSourceBase) ImplementsTableView_DraggingSession_EndedAtPoint_Operation() bool {
	return false
}

func (p *TableViewDataSourceBase) TableView_DraggingSession_EndedAtPoint_Operation(tableView TableView, session DraggingSession, screenPoint foundation.Point, operation DragOperation) {
	panic("unimpemented protocol method")
}

func (p *TableViewDataSourceBase) ImplementsTableView_SortDescriptorsDidChange() bool {
	return false
}

func (p *TableViewDataSourceBase) TableView_SortDescriptorsDidChange(tableView TableView, oldDescriptors []foundation.SortDescriptor) {
	panic("unimpemented protocol method")
}
