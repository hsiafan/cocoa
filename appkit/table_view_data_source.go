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
	TableView_AcceptDrop_Row_DropOperation(tableView TableView, info DraggingInfoWrapper, row int, dropOperation TableViewDropOperation) bool
	ImplementsTableView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes() bool
	// optional
	// deprecated
	TableView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes(tableView TableView, dropDestination foundation.URL, indexSet foundation.IndexSet) []string
	ImplementsTableView_ValidateDrop_ProposedRow_ProposedDropOperation() bool
	// optional
	TableView_ValidateDrop_ProposedRow_ProposedDropOperation(tableView TableView, info DraggingInfoWrapper, row int, dropOperation TableViewDropOperation) DragOperation
	ImplementsTableView_WriteRowsWithIndexes_ToPasteboard() bool
	// optional
	// deprecated
	TableView_WriteRowsWithIndexes_ToPasteboard(tableView TableView, rowIndexes foundation.IndexSet, pboard Pasteboard) bool
	ImplementsTableView_DraggingSession_WillBeginAtPoint_ForRowIndexes() bool
	// optional
	TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes(tableView TableView, session DraggingSession, screenPoint foundation.Point, rowIndexes foundation.IndexSet)
	ImplementsTableView_UpdateDraggingItemsForDrag() bool
	// optional
	TableView_UpdateDraggingItemsForDrag(tableView TableView, draggingInfo DraggingInfoWrapper)
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

func (p *TableViewDataSourceBase) TableView_AcceptDrop_Row_DropOperation(tableView TableView, info DraggingInfoWrapper, row int, dropOperation TableViewDropOperation) bool {
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

func (p *TableViewDataSourceBase) TableView_ValidateDrop_ProposedRow_ProposedDropOperation(tableView TableView, info DraggingInfoWrapper, row int, dropOperation TableViewDropOperation) DragOperation {
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

func (p *TableViewDataSourceBase) TableView_UpdateDraggingItemsForDrag(tableView TableView, draggingInfo DraggingInfoWrapper) {
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

type TableViewDataSourceWrapper struct {
	objc.Object
}

func (t_ TableViewDataSourceWrapper) NumberOfRowsInTableView(tableView ITableView) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("numberOfRowsInTableView:"), objc.ExtractPtr(tableView))
	return rv
}

func (t_ TableViewDataSourceWrapper) TableView_ObjectValueForTableColumn_Row(tableView ITableView, tableColumn ITableColumn, row int) objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.GetSelector("tableView:objectValueForTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn), row)
	return rv
}

func (t_ TableViewDataSourceWrapper) TableView_SetObjectValue_ForTableColumn_Row(tableView ITableView, object objc.IObject, tableColumn ITableColumn, row int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:setObjectValue:forTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(object), objc.ExtractPtr(tableColumn), row)
}

func (t_ TableViewDataSourceWrapper) TableView_PasteboardWriterForRow(tableView ITableView, row int) PasteboardWritingWrapper {
	rv := objc.CallMethod[PasteboardWritingWrapper](t_, objc.GetSelector("tableView:pasteboardWriterForRow:"), objc.ExtractPtr(tableView), row)
	return rv
}

func (t_ TableViewDataSourceWrapper) TableView_AcceptDrop_Row_DropOperation(tableView ITableView, info objc.IObject, row int, dropOperation TableViewDropOperation) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:acceptDrop:row:dropOperation:"), objc.ExtractPtr(tableView), objc.ExtractPtr(info), row, dropOperation)
	return rv
}

// deprecated
func (t_ TableViewDataSourceWrapper) TableView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes(tableView ITableView, dropDestination foundation.IURL, indexSet foundation.IIndexSet) []string {
	rv := objc.CallMethod[[]string](t_, objc.GetSelector("tableView:namesOfPromisedFilesDroppedAtDestination:forDraggedRowsWithIndexes:"), objc.ExtractPtr(tableView), objc.ExtractPtr(dropDestination), objc.ExtractPtr(indexSet))
	return rv
}

func (t_ TableViewDataSourceWrapper) TableView_ValidateDrop_ProposedRow_ProposedDropOperation(tableView ITableView, info objc.IObject, row int, dropOperation TableViewDropOperation) DragOperation {
	rv := objc.CallMethod[DragOperation](t_, objc.GetSelector("tableView:validateDrop:proposedRow:proposedDropOperation:"), objc.ExtractPtr(tableView), objc.ExtractPtr(info), row, dropOperation)
	return rv
}

// deprecated
func (t_ TableViewDataSourceWrapper) TableView_WriteRowsWithIndexes_ToPasteboard(tableView ITableView, rowIndexes foundation.IIndexSet, pboard IPasteboard) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:writeRowsWithIndexes:toPasteboard:"), objc.ExtractPtr(tableView), objc.ExtractPtr(rowIndexes), objc.ExtractPtr(pboard))
	return rv
}

func (t_ TableViewDataSourceWrapper) TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes(tableView ITableView, session IDraggingSession, screenPoint foundation.Point, rowIndexes foundation.IIndexSet) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:draggingSession:willBeginAtPoint:forRowIndexes:"), objc.ExtractPtr(tableView), objc.ExtractPtr(session), screenPoint, objc.ExtractPtr(rowIndexes))
}

func (t_ TableViewDataSourceWrapper) TableView_UpdateDraggingItemsForDrag(tableView ITableView, draggingInfo objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:updateDraggingItemsForDrag:"), objc.ExtractPtr(tableView), objc.ExtractPtr(draggingInfo))
}

func (t_ TableViewDataSourceWrapper) TableView_DraggingSession_EndedAtPoint_Operation(tableView ITableView, session IDraggingSession, screenPoint foundation.Point, operation DragOperation) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:draggingSession:endedAtPoint:operation:"), objc.ExtractPtr(tableView), objc.ExtractPtr(session), screenPoint, operation)
}

func (t_ TableViewDataSourceWrapper) TableView_SortDescriptorsDidChange(tableView ITableView, oldDescriptors []foundation.ISortDescriptor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:sortDescriptorsDidChange:"), objc.ExtractPtr(tableView), oldDescriptors)
}
