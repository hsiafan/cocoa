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
	TableView_PasteboardWriterForRow(tableView TableView, row int) PasteboardWriting
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

type TableViewDataSourceWrapper struct {
	objc.Object
}

func (t_ *TableViewDataSourceWrapper) ImplementsNumberOfRowsInTableView() bool {
	return t_.RespondsToSelector(objc.GetSelector("numberOfRowsInTableView:"))
}

func (t_ TableViewDataSourceWrapper) NumberOfRowsInTableView(tableView ITableView) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("numberOfRowsInTableView:"), objc.ExtractPtr(tableView))
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_ObjectValueForTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:objectValueForTableColumn:row:"))
}

func (t_ TableViewDataSourceWrapper) TableView_ObjectValueForTableColumn_Row(tableView ITableView, tableColumn ITableColumn, row int) objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.GetSelector("tableView:objectValueForTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(tableColumn), row)
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_SetObjectValue_ForTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:setObjectValue:forTableColumn:row:"))
}

func (t_ TableViewDataSourceWrapper) TableView_SetObjectValue_ForTableColumn_Row(tableView ITableView, object objc.IObject, tableColumn ITableColumn, row int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:setObjectValue:forTableColumn:row:"), objc.ExtractPtr(tableView), objc.ExtractPtr(object), objc.ExtractPtr(tableColumn), row)
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_PasteboardWriterForRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:pasteboardWriterForRow:"))
}

func (t_ TableViewDataSourceWrapper) TableView_PasteboardWriterForRow(tableView ITableView, row int) PasteboardWritingWrapper {
	rv := objc.CallMethod[PasteboardWritingWrapper](t_, objc.GetSelector("tableView:pasteboardWriterForRow:"), objc.ExtractPtr(tableView), row)
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_AcceptDrop_Row_DropOperation() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:acceptDrop:row:dropOperation:"))
}

func (t_ TableViewDataSourceWrapper) TableView_AcceptDrop_Row_DropOperation(tableView ITableView, info DraggingInfo, row int, dropOperation TableViewDropOperation) bool {
	po := objc.WrapAsProtocol("NSDraggingInfo", info)
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:acceptDrop:row:dropOperation:"), objc.ExtractPtr(tableView), po, row, dropOperation)
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:namesOfPromisedFilesDroppedAtDestination:forDraggedRowsWithIndexes:"))
}

// deprecated
func (t_ TableViewDataSourceWrapper) TableView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes(tableView ITableView, dropDestination foundation.IURL, indexSet foundation.IIndexSet) []string {
	rv := objc.CallMethod[[]string](t_, objc.GetSelector("tableView:namesOfPromisedFilesDroppedAtDestination:forDraggedRowsWithIndexes:"), objc.ExtractPtr(tableView), objc.ExtractPtr(dropDestination), objc.ExtractPtr(indexSet))
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_ValidateDrop_ProposedRow_ProposedDropOperation() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:validateDrop:proposedRow:proposedDropOperation:"))
}

func (t_ TableViewDataSourceWrapper) TableView_ValidateDrop_ProposedRow_ProposedDropOperation(tableView ITableView, info DraggingInfo, row int, dropOperation TableViewDropOperation) DragOperation {
	po := objc.WrapAsProtocol("NSDraggingInfo", info)
	rv := objc.CallMethod[DragOperation](t_, objc.GetSelector("tableView:validateDrop:proposedRow:proposedDropOperation:"), objc.ExtractPtr(tableView), po, row, dropOperation)
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_WriteRowsWithIndexes_ToPasteboard() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:writeRowsWithIndexes:toPasteboard:"))
}

// deprecated
func (t_ TableViewDataSourceWrapper) TableView_WriteRowsWithIndexes_ToPasteboard(tableView ITableView, rowIndexes foundation.IIndexSet, pboard IPasteboard) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:writeRowsWithIndexes:toPasteboard:"), objc.ExtractPtr(tableView), objc.ExtractPtr(rowIndexes), objc.ExtractPtr(pboard))
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_DraggingSession_WillBeginAtPoint_ForRowIndexes() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:draggingSession:willBeginAtPoint:forRowIndexes:"))
}

func (t_ TableViewDataSourceWrapper) TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes(tableView ITableView, session IDraggingSession, screenPoint foundation.Point, rowIndexes foundation.IIndexSet) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:draggingSession:willBeginAtPoint:forRowIndexes:"), objc.ExtractPtr(tableView), objc.ExtractPtr(session), screenPoint, objc.ExtractPtr(rowIndexes))
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_UpdateDraggingItemsForDrag() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:updateDraggingItemsForDrag:"))
}

func (t_ TableViewDataSourceWrapper) TableView_UpdateDraggingItemsForDrag(tableView ITableView, draggingInfo DraggingInfo) {
	po := objc.WrapAsProtocol("NSDraggingInfo", draggingInfo)
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:updateDraggingItemsForDrag:"), objc.ExtractPtr(tableView), po)
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_DraggingSession_EndedAtPoint_Operation() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:draggingSession:endedAtPoint:operation:"))
}

func (t_ TableViewDataSourceWrapper) TableView_DraggingSession_EndedAtPoint_Operation(tableView ITableView, session IDraggingSession, screenPoint foundation.Point, operation DragOperation) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:draggingSession:endedAtPoint:operation:"), objc.ExtractPtr(tableView), objc.ExtractPtr(session), screenPoint, operation)
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_SortDescriptorsDidChange() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:sortDescriptorsDidChange:"))
}

func (t_ TableViewDataSourceWrapper) TableView_SortDescriptorsDidChange(tableView ITableView, oldDescriptors []foundation.ISortDescriptor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tableView:sortDescriptorsDidChange:"), objc.ExtractPtr(tableView), oldDescriptors)
}
