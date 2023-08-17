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
	OutlineView_PasteboardWriterForItem(outlineView OutlineView, item objc.Object) objc.IObject
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

func WrapOutlineViewDataSource(v OutlineViewDataSource) objc.Object {
	return objc.WrapAsProtocol("NSOutlineViewDataSource", v)
}

type OutlineViewDataSourceBase struct {
}

func (p *OutlineViewDataSourceBase) ImplementsOutlineView_AcceptDrop_Item_ChildIndex() bool {
	return false
}

func (p *OutlineViewDataSourceBase) OutlineView_AcceptDrop_Item_ChildIndex(outlineView OutlineView, info DraggingInfoWrapper, item objc.Object, index int) bool {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDataSourceBase) ImplementsOutlineView_Child_OfItem() bool {
	return false
}

func (p *OutlineViewDataSourceBase) OutlineView_Child_OfItem(outlineView OutlineView, index int, item objc.Object) objc.IObject {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDataSourceBase) ImplementsOutlineView_DraggingSession_EndedAtPoint_Operation() bool {
	return false
}

func (p *OutlineViewDataSourceBase) OutlineView_DraggingSession_EndedAtPoint_Operation(outlineView OutlineView, session DraggingSession, screenPoint foundation.Point, operation DragOperation) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDataSourceBase) ImplementsOutlineView_DraggingSession_WillBeginAtPoint_ForItems() bool {
	return false
}

func (p *OutlineViewDataSourceBase) OutlineView_DraggingSession_WillBeginAtPoint_ForItems(outlineView OutlineView, session DraggingSession, screenPoint foundation.Point, draggedItems []objc.Object) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDataSourceBase) ImplementsOutlineView_IsItemExpandable() bool {
	return false
}

func (p *OutlineViewDataSourceBase) OutlineView_IsItemExpandable(outlineView OutlineView, item objc.Object) bool {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDataSourceBase) ImplementsOutlineView_ItemForPersistentObject() bool {
	return false
}

func (p *OutlineViewDataSourceBase) OutlineView_ItemForPersistentObject(outlineView OutlineView, object objc.Object) objc.IObject {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDataSourceBase) ImplementsOutlineView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItems() bool {
	return false
}

// deprecated
func (p *OutlineViewDataSourceBase) OutlineView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItems(outlineView OutlineView, dropDestination foundation.URL, items []objc.Object) []string {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDataSourceBase) ImplementsOutlineView_NumberOfChildrenOfItem() bool {
	return false
}

func (p *OutlineViewDataSourceBase) OutlineView_NumberOfChildrenOfItem(outlineView OutlineView, item objc.Object) int {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDataSourceBase) ImplementsOutlineView_ObjectValueForTableColumn_ByItem() bool {
	return false
}

func (p *OutlineViewDataSourceBase) OutlineView_ObjectValueForTableColumn_ByItem(outlineView OutlineView, tableColumn TableColumn, item objc.Object) objc.IObject {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDataSourceBase) ImplementsOutlineView_PasteboardWriterForItem() bool {
	return false
}

func (p *OutlineViewDataSourceBase) OutlineView_PasteboardWriterForItem(outlineView OutlineView, item objc.Object) objc.IObject {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDataSourceBase) ImplementsOutlineView_PersistentObjectForItem() bool {
	return false
}

func (p *OutlineViewDataSourceBase) OutlineView_PersistentObjectForItem(outlineView OutlineView, item objc.Object) objc.IObject {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDataSourceBase) ImplementsOutlineView_SetObjectValue_ForTableColumn_ByItem() bool {
	return false
}

func (p *OutlineViewDataSourceBase) OutlineView_SetObjectValue_ForTableColumn_ByItem(outlineView OutlineView, object objc.Object, tableColumn TableColumn, item objc.Object) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDataSourceBase) ImplementsOutlineView_SortDescriptorsDidChange() bool {
	return false
}

func (p *OutlineViewDataSourceBase) OutlineView_SortDescriptorsDidChange(outlineView OutlineView, oldDescriptors []foundation.SortDescriptor) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDataSourceBase) ImplementsOutlineView_UpdateDraggingItemsForDrag() bool {
	return false
}

func (p *OutlineViewDataSourceBase) OutlineView_UpdateDraggingItemsForDrag(outlineView OutlineView, draggingInfo DraggingInfoWrapper) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDataSourceBase) ImplementsOutlineView_ValidateDrop_ProposedItem_ProposedChildIndex() bool {
	return false
}

func (p *OutlineViewDataSourceBase) OutlineView_ValidateDrop_ProposedItem_ProposedChildIndex(outlineView OutlineView, info DraggingInfoWrapper, item objc.Object, index int) DragOperation {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDataSourceBase) ImplementsOutlineView_WriteItems_ToPasteboard() bool {
	return false
}

// deprecated
func (p *OutlineViewDataSourceBase) OutlineView_WriteItems_ToPasteboard(outlineView OutlineView, items []objc.Object, pasteboard Pasteboard) bool {
	panic("unimpemented protocol method")
}

type OutlineViewDataSourceWrapper struct {
	objc.Object
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_AcceptDrop_Item_ChildIndex(outlineView IOutlineView, info objc.IObject, item objc.IObject, index int) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:acceptDrop:item:childIndex:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(info), objc.ExtractPtr(item), index)
	return rv
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_Child_OfItem(outlineView IOutlineView, index int, item objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("outlineView:child:ofItem:"), objc.ExtractPtr(outlineView), index, objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_DraggingSession_EndedAtPoint_Operation(outlineView IOutlineView, session IDraggingSession, screenPoint foundation.Point, operation DragOperation) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:draggingSession:endedAtPoint:operation:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(session), screenPoint, operation)
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_DraggingSession_WillBeginAtPoint_ForItems(outlineView IOutlineView, session IDraggingSession, screenPoint foundation.Point, draggedItems []objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:draggingSession:willBeginAtPoint:forItems:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(session), screenPoint, draggedItems)
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_IsItemExpandable(outlineView IOutlineView, item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:isItemExpandable:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_ItemForPersistentObject(outlineView IOutlineView, object objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("outlineView:itemForPersistentObject:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(object))
	return rv
}

// deprecated
func (o_ OutlineViewDataSourceWrapper) OutlineView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItems(outlineView IOutlineView, dropDestination foundation.IURL, items []objc.IObject) []string {
	rv := objc.CallMethod[[]string](o_, objc.GetSelector("outlineView:namesOfPromisedFilesDroppedAtDestination:forDraggedItems:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(dropDestination), items)
	return rv
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_NumberOfChildrenOfItem(outlineView IOutlineView, item objc.IObject) int {
	rv := objc.CallMethod[int](o_, objc.GetSelector("outlineView:numberOfChildrenOfItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_ObjectValueForTableColumn_ByItem(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("outlineView:objectValueForTableColumn:byItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_PasteboardWriterForItem(outlineView IOutlineView, item objc.IObject) PasteboardWritingWrapper {
	rv := objc.CallMethod[PasteboardWritingWrapper](o_, objc.GetSelector("outlineView:pasteboardWriterForItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_PersistentObjectForItem(outlineView IOutlineView, item objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("outlineView:persistentObjectForItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_SetObjectValue_ForTableColumn_ByItem(outlineView IOutlineView, object objc.IObject, tableColumn ITableColumn, item objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:setObjectValue:forTableColumn:byItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(object), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_SortDescriptorsDidChange(outlineView IOutlineView, oldDescriptors []foundation.ISortDescriptor) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:sortDescriptorsDidChange:"), objc.ExtractPtr(outlineView), oldDescriptors)
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_UpdateDraggingItemsForDrag(outlineView IOutlineView, draggingInfo objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:updateDraggingItemsForDrag:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(draggingInfo))
}

func (o_ OutlineViewDataSourceWrapper) OutlineView_ValidateDrop_ProposedItem_ProposedChildIndex(outlineView IOutlineView, info objc.IObject, item objc.IObject, index int) DragOperation {
	rv := objc.CallMethod[DragOperation](o_, objc.GetSelector("outlineView:validateDrop:proposedItem:proposedChildIndex:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(info), objc.ExtractPtr(item), index)
	return rv
}

// deprecated
func (o_ OutlineViewDataSourceWrapper) OutlineView_WriteItems_ToPasteboard(outlineView IOutlineView, items []objc.IObject, pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:writeItems:toPasteboard:"), objc.ExtractPtr(outlineView), items, objc.ExtractPtr(pasteboard))
	return rv
}
