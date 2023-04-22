// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type OutlineViewDataSource interface {
	ImplementsOutlineView_AcceptDrop_Item_ChildIndex() bool
	// optional
	OutlineView_AcceptDrop_Item_ChildIndex(outlineView OutlineView, info DraggingInfoWrapper, item objc.Object, index int) bool
	ImplementsOutlineView_Child_OfItem() bool
	// optional
	OutlineView_Child_OfItem(outlineView OutlineView, index int, item objc.Object) objc.IObject
	ImplementsOutlineView_DraggingSession_EndedAtPoint_Operation() bool
	// optional
	OutlineView_DraggingSession_EndedAtPoint_Operation(outlineView OutlineView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	ImplementsOutlineView_DraggingSession_WillBeginAtPoint_ForItems() bool
	// optional
	OutlineView_DraggingSession_WillBeginAtPoint_ForItems(outlineView OutlineView, session DraggingSession, screenPoint foundation.Point, draggedItems []objc.Object)
	ImplementsOutlineView_IsItemExpandable() bool
	// optional
	OutlineView_IsItemExpandable(outlineView OutlineView, item objc.Object) bool
	ImplementsOutlineView_ItemForPersistentObject() bool
	// optional
	OutlineView_ItemForPersistentObject(outlineView OutlineView, object objc.Object) objc.IObject
	ImplementsOutlineView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItems() bool
	// optional
	// deprecated
	OutlineView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItems(outlineView OutlineView, dropDestination foundation.URL, items []objc.Object) []string
	ImplementsOutlineView_NumberOfChildrenOfItem() bool
	// optional
	OutlineView_NumberOfChildrenOfItem(outlineView OutlineView, item objc.Object) int
	ImplementsOutlineView_ObjectValueForTableColumn_ByItem() bool
	// optional
	OutlineView_ObjectValueForTableColumn_ByItem(outlineView OutlineView, tableColumn TableColumn, item objc.Object) objc.IObject
	ImplementsOutlineView_PasteboardWriterForItem() bool
	// optional
	OutlineView_PasteboardWriterForItem(outlineView OutlineView, item objc.Object) PasteboardWriting
	ImplementsOutlineView_PersistentObjectForItem() bool
	// optional
	OutlineView_PersistentObjectForItem(outlineView OutlineView, item objc.Object) objc.IObject
	ImplementsOutlineView_SetObjectValue_ForTableColumn_ByItem() bool
	// optional
	OutlineView_SetObjectValue_ForTableColumn_ByItem(outlineView OutlineView, object objc.Object, tableColumn TableColumn, item objc.Object)
	ImplementsOutlineView_SortDescriptorsDidChange() bool
	// optional
	OutlineView_SortDescriptorsDidChange(outlineView OutlineView, oldDescriptors []foundation.SortDescriptor)
	ImplementsOutlineView_UpdateDraggingItemsForDrag() bool
	// optional
	OutlineView_UpdateDraggingItemsForDrag(outlineView OutlineView, draggingInfo DraggingInfoWrapper)
	ImplementsOutlineView_ValidateDrop_ProposedItem_ProposedChildIndex() bool
	// optional
	OutlineView_ValidateDrop_ProposedItem_ProposedChildIndex(outlineView OutlineView, info DraggingInfoWrapper, item objc.Object, index int) DragOperation
	ImplementsOutlineView_WriteItems_ToPasteboard() bool
	// optional
	// deprecated
	OutlineView_WriteItems_ToPasteboard(outlineView OutlineView, items []objc.Object, pasteboard Pasteboard) bool
}

type OutlineViewDataSourceWrapper struct {
	objc.Object
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_AcceptDrop_Item_ChildIndex() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:acceptDrop:item:childIndex:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_AcceptDrop_Item_ChildIndex(outlineView IOutlineView, info DraggingInfo, item objc.IObject, index int) bool {
	po := objc.CreateProtocol("NSDraggingInfo", info)
	defer po.Release()
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:acceptDrop:item:childIndex:"), outlineView, po, item, index)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_Child_OfItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:child:ofItem:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_Child_OfItem(outlineView IOutlineView, index int, item objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("outlineView:child:ofItem:"), outlineView, index, item)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_DraggingSession_EndedAtPoint_Operation() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:draggingSession:endedAtPoint:operation:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_DraggingSession_EndedAtPoint_Operation(outlineView IOutlineView, session IDraggingSession, screenPoint foundation.Point, operation DragOperation) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:draggingSession:endedAtPoint:operation:"), outlineView, session, screenPoint, operation)
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_DraggingSession_WillBeginAtPoint_ForItems() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:draggingSession:willBeginAtPoint:forItems:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_DraggingSession_WillBeginAtPoint_ForItems(outlineView IOutlineView, session IDraggingSession, screenPoint foundation.Point, draggedItems []objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:draggingSession:willBeginAtPoint:forItems:"), outlineView, session, screenPoint, draggedItems)
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_IsItemExpandable() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:isItemExpandable:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_IsItemExpandable(outlineView IOutlineView, item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:isItemExpandable:"), outlineView, item)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_ItemForPersistentObject() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:itemForPersistentObject:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_ItemForPersistentObject(outlineView IOutlineView, object objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("outlineView:itemForPersistentObject:"), outlineView, object)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItems() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:namesOfPromisedFilesDroppedAtDestination:forDraggedItems:"))
}

// deprecated
func (o_ OutlineViewDataSourceWrapper) OutlineView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItems(outlineView IOutlineView, dropDestination foundation.IURL, items []objc.IObject) []string {
	rv := objc.CallMethod[[]string](o_, objc.GetSelector("outlineView:namesOfPromisedFilesDroppedAtDestination:forDraggedItems:"), outlineView, dropDestination, items)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_NumberOfChildrenOfItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:numberOfChildrenOfItem:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_NumberOfChildrenOfItem(outlineView IOutlineView, item objc.IObject) int {
	rv := objc.CallMethod[int](o_, objc.GetSelector("outlineView:numberOfChildrenOfItem:"), outlineView, item)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_ObjectValueForTableColumn_ByItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:objectValueForTableColumn:byItem:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_ObjectValueForTableColumn_ByItem(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("outlineView:objectValueForTableColumn:byItem:"), outlineView, tableColumn, item)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_PasteboardWriterForItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:pasteboardWriterForItem:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_PasteboardWriterForItem(outlineView IOutlineView, item objc.IObject) PasteboardWritingWrapper {
	rv := objc.CallMethod[PasteboardWritingWrapper](o_, objc.GetSelector("outlineView:pasteboardWriterForItem:"), outlineView, item)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_PersistentObjectForItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:persistentObjectForItem:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_PersistentObjectForItem(outlineView IOutlineView, item objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("outlineView:persistentObjectForItem:"), outlineView, item)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_SetObjectValue_ForTableColumn_ByItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:setObjectValue:forTableColumn:byItem:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_SetObjectValue_ForTableColumn_ByItem(outlineView IOutlineView, object objc.IObject, tableColumn ITableColumn, item objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:setObjectValue:forTableColumn:byItem:"), outlineView, object, tableColumn, item)
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_SortDescriptorsDidChange() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:sortDescriptorsDidChange:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_SortDescriptorsDidChange(outlineView IOutlineView, oldDescriptors []foundation.ISortDescriptor) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:sortDescriptorsDidChange:"), outlineView, oldDescriptors)
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_UpdateDraggingItemsForDrag() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:updateDraggingItemsForDrag:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_UpdateDraggingItemsForDrag(outlineView IOutlineView, draggingInfo DraggingInfo) {
	po := objc.CreateProtocol("NSDraggingInfo", draggingInfo)
	defer po.Release()
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:updateDraggingItemsForDrag:"), outlineView, po)
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_ValidateDrop_ProposedItem_ProposedChildIndex() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:validateDrop:proposedItem:proposedChildIndex:"))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_ValidateDrop_ProposedItem_ProposedChildIndex(outlineView IOutlineView, info DraggingInfo, item objc.IObject, index int) DragOperation {
	po := objc.CreateProtocol("NSDraggingInfo", info)
	defer po.Release()
	rv := objc.CallMethod[DragOperation](o_, objc.GetSelector("outlineView:validateDrop:proposedItem:proposedChildIndex:"), outlineView, po, item, index)
	return rv
}

func (o_ *OutlineViewDataSourceWrapper) ImplementsOutlineView_WriteItems_ToPasteboard() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:writeItems:toPasteboard:"))
}

// deprecated
func (o_ OutlineViewDataSourceWrapper) OutlineView_WriteItems_ToPasteboard(outlineView IOutlineView, items []objc.IObject, pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:writeItems:toPasteboard:"), outlineView, items, pasteboard)
	return rv
}
