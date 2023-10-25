// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type OutlineViewDataSource interface {
	ImplementsOutlineView_AcceptDrop_Item_ChildIndex() bool
	// optional
	OutlineView_AcceptDrop_Item_ChildIndex(outlineView OutlineView, info objc.Object, item objc.Object, index int) bool
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
	OutlineView_UpdateDraggingItemsForDrag(outlineView OutlineView, draggingInfo objc.Object)
	ImplementsOutlineView_ValidateDrop_ProposedItem_ProposedChildIndex() bool
	// optional
	OutlineView_ValidateDrop_ProposedItem_ProposedChildIndex(outlineView OutlineView, info objc.Object, item objc.Object, index int) DragOperation
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

func (p *OutlineViewDataSourceBase) OutlineView_AcceptDrop_Item_ChildIndex(outlineView OutlineView, info objc.Object, item objc.Object, index int) bool {
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

func (p *OutlineViewDataSourceBase) OutlineView_UpdateDraggingItemsForDrag(outlineView OutlineView, draggingInfo objc.Object) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDataSourceBase) ImplementsOutlineView_ValidateDrop_ProposedItem_ProposedChildIndex() bool {
	return false
}

func (p *OutlineViewDataSourceBase) OutlineView_ValidateDrop_ProposedItem_ProposedChildIndex(outlineView OutlineView, info objc.Object, item objc.Object, index int) DragOperation {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDataSourceBase) ImplementsOutlineView_WriteItems_ToPasteboard() bool {
	return false
}

// deprecated
func (p *OutlineViewDataSourceBase) OutlineView_WriteItems_ToPasteboard(outlineView OutlineView, items []objc.Object, pasteboard Pasteboard) bool {
	panic("unimpemented protocol method")
}
