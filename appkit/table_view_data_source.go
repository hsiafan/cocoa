// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[int](t_, "numberOfRowsInTableView:", tableView)
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_ObjectValueForTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:objectValueForTableColumn:row:"))
}

func (t_ TableViewDataSourceWrapper) TableView_ObjectValueForTableColumn_Row(tableView ITableView, tableColumn ITableColumn, row int) objc.Object {
	rv := ffi.CallMethod[objc.Object](t_, "tableView:objectValueForTableColumn:row:", tableView, tableColumn, row)
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_SetObjectValue_ForTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:setObjectValue:forTableColumn:row:"))
}

func (t_ TableViewDataSourceWrapper) TableView_SetObjectValue_ForTableColumn_Row(tableView ITableView, object objc.IObject, tableColumn ITableColumn, row int) {
	ffi.CallMethod[ffi.Void](t_, "tableView:setObjectValue:forTableColumn:row:", tableView, object, tableColumn, row)
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_PasteboardWriterForRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:pasteboardWriterForRow:"))
}

func (t_ TableViewDataSourceWrapper) TableView_PasteboardWriterForRow(tableView ITableView, row int) PasteboardWritingWrapper {
	rv := ffi.CallMethod[PasteboardWritingWrapper](t_, "tableView:pasteboardWriterForRow:", tableView, row)
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_AcceptDrop_Row_DropOperation() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:acceptDrop:row:dropOperation:"))
}

func (t_ TableViewDataSourceWrapper) TableView_AcceptDrop_Row_DropOperation(tableView ITableView, info DraggingInfo, row int, dropOperation TableViewDropOperation) bool {
	po := ffi.CreateProtocol("NSDraggingInfo", info)
	defer po.Release()
	rv := ffi.CallMethod[bool](t_, "tableView:acceptDrop:row:dropOperation:", tableView, po, row, dropOperation)
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:namesOfPromisedFilesDroppedAtDestination:forDraggedRowsWithIndexes:"))
}

// deprecated
func (t_ TableViewDataSourceWrapper) TableView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes(tableView ITableView, dropDestination foundation.IURL, indexSet foundation.IIndexSet) []string {
	rv := ffi.CallMethod[[]string](t_, "tableView:namesOfPromisedFilesDroppedAtDestination:forDraggedRowsWithIndexes:", tableView, dropDestination, indexSet)
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_ValidateDrop_ProposedRow_ProposedDropOperation() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:validateDrop:proposedRow:proposedDropOperation:"))
}

func (t_ TableViewDataSourceWrapper) TableView_ValidateDrop_ProposedRow_ProposedDropOperation(tableView ITableView, info DraggingInfo, row int, dropOperation TableViewDropOperation) DragOperation {
	po := ffi.CreateProtocol("NSDraggingInfo", info)
	defer po.Release()
	rv := ffi.CallMethod[DragOperation](t_, "tableView:validateDrop:proposedRow:proposedDropOperation:", tableView, po, row, dropOperation)
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_WriteRowsWithIndexes_ToPasteboard() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:writeRowsWithIndexes:toPasteboard:"))
}

// deprecated
func (t_ TableViewDataSourceWrapper) TableView_WriteRowsWithIndexes_ToPasteboard(tableView ITableView, rowIndexes foundation.IIndexSet, pboard IPasteboard) bool {
	rv := ffi.CallMethod[bool](t_, "tableView:writeRowsWithIndexes:toPasteboard:", tableView, rowIndexes, pboard)
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_DraggingSession_WillBeginAtPoint_ForRowIndexes() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:draggingSession:willBeginAtPoint:forRowIndexes:"))
}

func (t_ TableViewDataSourceWrapper) TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes(tableView ITableView, session IDraggingSession, screenPoint foundation.Point, rowIndexes foundation.IIndexSet) {
	ffi.CallMethod[ffi.Void](t_, "tableView:draggingSession:willBeginAtPoint:forRowIndexes:", tableView, session, screenPoint, rowIndexes)
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_UpdateDraggingItemsForDrag() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:updateDraggingItemsForDrag:"))
}

func (t_ TableViewDataSourceWrapper) TableView_UpdateDraggingItemsForDrag(tableView ITableView, draggingInfo DraggingInfo) {
	po := ffi.CreateProtocol("NSDraggingInfo", draggingInfo)
	defer po.Release()
	ffi.CallMethod[ffi.Void](t_, "tableView:updateDraggingItemsForDrag:", tableView, po)
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_DraggingSession_EndedAtPoint_Operation() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:draggingSession:endedAtPoint:operation:"))
}

func (t_ TableViewDataSourceWrapper) TableView_DraggingSession_EndedAtPoint_Operation(tableView ITableView, session IDraggingSession, screenPoint foundation.Point, operation DragOperation) {
	ffi.CallMethod[ffi.Void](t_, "tableView:draggingSession:endedAtPoint:operation:", tableView, session, screenPoint, operation)
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_SortDescriptorsDidChange() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:sortDescriptorsDidChange:"))
}

func (t_ TableViewDataSourceWrapper) TableView_SortDescriptorsDidChange(tableView ITableView, oldDescriptors []foundation.ISortDescriptor) {
	ffi.CallMethod[ffi.Void](t_, "tableView:sortDescriptorsDidChange:", tableView, oldDescriptors)
}
