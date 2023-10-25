// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type CollectionViewDelegate interface {
	ImplementsCollectionView_ShouldSelectItemsAtIndexPaths() bool
	// optional
	CollectionView_ShouldSelectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet
	ImplementsCollectionView_DidSelectItemsAtIndexPaths() bool
	// optional
	CollectionView_DidSelectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set)
	ImplementsCollectionView_ShouldDeselectItemsAtIndexPaths() bool
	// optional
	CollectionView_ShouldDeselectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet
	ImplementsCollectionView_DidDeselectItemsAtIndexPaths() bool
	// optional
	CollectionView_DidDeselectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set)
	ImplementsCollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState() bool
	// optional
	CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) foundation.ISet
	ImplementsCollectionView_DidChangeItemsAtIndexPaths_ToHighlightState() bool
	// optional
	CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState)
	ImplementsCollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath() bool
	// optional
	CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)
	ImplementsCollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath() bool
	// optional
	CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)
	ImplementsCollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath() bool
	// optional
	CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)
	ImplementsCollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath() bool
	// optional
	CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)
	ImplementsCollectionView_TransitionLayoutForOldLayout_NewLayout() bool
	// optional
	CollectionView_TransitionLayoutForOldLayout_NewLayout(collectionView CollectionView, fromLayout CollectionViewLayout, toLayout CollectionViewLayout) ICollectionViewTransitionLayout
	ImplementsCollectionView_CanDragItemsAtIndexPaths_WithEvent() bool
	// optional
	CollectionView_CanDragItemsAtIndexPaths_WithEvent(collectionView CollectionView, indexPaths foundation.Set, event Event) bool
	ImplementsCollectionView_PasteboardWriterForItemAtIndexPath() bool
	// optional
	CollectionView_PasteboardWriterForItemAtIndexPath(collectionView CollectionView, indexPath foundation.IndexPath) objc.IObject
	ImplementsCollectionView_WriteItemsAtIndexPaths_ToPasteboard() bool
	// optional
	// deprecated
	CollectionView_WriteItemsAtIndexPaths_ToPasteboard(collectionView CollectionView, indexPaths foundation.Set, pasteboard Pasteboard) bool
	ImplementsCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths() bool
	// optional
	// deprecated
	CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths(collectionView CollectionView, dropURL foundation.URL, indexPaths foundation.Set) []string
	ImplementsCollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset() bool
	// optional
	CollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset(collectionView CollectionView, indexPaths foundation.Set, event Event, dragImageOffset *foundation.Point) IImage
	ImplementsCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths() bool
	// optional
	CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexPaths foundation.Set)
	ImplementsCollectionView_DraggingSession_EndedAtPoint_DragOperation() bool
	// optional
	CollectionView_DraggingSession_EndedAtPoint_DragOperation(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	ImplementsCollectionView_UpdateDraggingItemsForDrag() bool
	// optional
	CollectionView_UpdateDraggingItemsForDrag(collectionView CollectionView, draggingInfo objc.Object)
	ImplementsCollectionView_ValidateDrop_ProposedIndexPath_DropOperation() bool
	// optional
	CollectionView_ValidateDrop_ProposedIndexPath_DropOperation(collectionView CollectionView, draggingInfo objc.Object, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation
	ImplementsCollectionView_AcceptDrop_IndexPath_DropOperation() bool
	// optional
	CollectionView_AcceptDrop_IndexPath_DropOperation(collectionView CollectionView, draggingInfo objc.Object, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool
	ImplementsCollectionView_CanDragItemsAtIndexes_WithEvent() bool
	// optional
	CollectionView_CanDragItemsAtIndexes_WithEvent(collectionView CollectionView, indexes foundation.IndexSet, event Event) bool
	ImplementsCollectionView_PasteboardWriterForItemAtIndex() bool
	// optional
	CollectionView_PasteboardWriterForItemAtIndex(collectionView CollectionView, index uint) objc.IObject
	ImplementsCollectionView_WriteItemsAtIndexes_ToPasteboard() bool
	// optional
	// deprecated
	CollectionView_WriteItemsAtIndexes_ToPasteboard(collectionView CollectionView, indexes foundation.IndexSet, pasteboard Pasteboard) bool
	ImplementsCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes() bool
	// optional
	// deprecated
	CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes(collectionView CollectionView, dropURL foundation.URL, indexes foundation.IndexSet) []string
	ImplementsCollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset() bool
	// optional
	CollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset(collectionView CollectionView, indexes foundation.IndexSet, event Event, dragImageOffset *foundation.Point) IImage
	ImplementsCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes() bool
	// optional
	CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexes foundation.IndexSet)
	ImplementsCollectionView_ValidateDrop_ProposedIndex_DropOperation() bool
	// optional
	CollectionView_ValidateDrop_ProposedIndex_DropOperation(collectionView CollectionView, draggingInfo objc.Object, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation
	ImplementsCollectionView_AcceptDrop_Index_DropOperation() bool
	// optional
	CollectionView_AcceptDrop_Index_DropOperation(collectionView CollectionView, draggingInfo objc.Object, index int, dropOperation CollectionViewDropOperation) bool
}

func WrapCollectionViewDelegate(v CollectionViewDelegate) objc.Object {
	return objc.WrapAsProtocol("NSCollectionViewDelegate", v)
}

type CollectionViewDelegateBase struct {
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_ShouldSelectItemsAtIndexPaths() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_ShouldSelectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_DidSelectItemsAtIndexPaths() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_DidSelectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_ShouldDeselectItemsAtIndexPaths() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_ShouldDeselectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_DidDeselectItemsAtIndexPaths() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_DidDeselectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) foundation.ISet {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_DidChangeItemsAtIndexPaths_ToHighlightState() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath) {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath) {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_TransitionLayoutForOldLayout_NewLayout() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_TransitionLayoutForOldLayout_NewLayout(collectionView CollectionView, fromLayout CollectionViewLayout, toLayout CollectionViewLayout) ICollectionViewTransitionLayout {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_CanDragItemsAtIndexPaths_WithEvent() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_CanDragItemsAtIndexPaths_WithEvent(collectionView CollectionView, indexPaths foundation.Set, event Event) bool {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_PasteboardWriterForItemAtIndexPath() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_PasteboardWriterForItemAtIndexPath(collectionView CollectionView, indexPath foundation.IndexPath) objc.IObject {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_WriteItemsAtIndexPaths_ToPasteboard() bool {
	return false
}

// deprecated
func (p *CollectionViewDelegateBase) CollectionView_WriteItemsAtIndexPaths_ToPasteboard(collectionView CollectionView, indexPaths foundation.Set, pasteboard Pasteboard) bool {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths() bool {
	return false
}

// deprecated
func (p *CollectionViewDelegateBase) CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths(collectionView CollectionView, dropURL foundation.URL, indexPaths foundation.Set) []string {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset(collectionView CollectionView, indexPaths foundation.Set, event Event, dragImageOffset *foundation.Point) IImage {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexPaths foundation.Set) {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_DraggingSession_EndedAtPoint_DragOperation() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_DraggingSession_EndedAtPoint_DragOperation(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, operation DragOperation) {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_UpdateDraggingItemsForDrag() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_UpdateDraggingItemsForDrag(collectionView CollectionView, draggingInfo objc.Object) {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_ValidateDrop_ProposedIndexPath_DropOperation() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_ValidateDrop_ProposedIndexPath_DropOperation(collectionView CollectionView, draggingInfo objc.Object, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_AcceptDrop_IndexPath_DropOperation() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_AcceptDrop_IndexPath_DropOperation(collectionView CollectionView, draggingInfo objc.Object, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_CanDragItemsAtIndexes_WithEvent() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_CanDragItemsAtIndexes_WithEvent(collectionView CollectionView, indexes foundation.IndexSet, event Event) bool {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_PasteboardWriterForItemAtIndex() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_PasteboardWriterForItemAtIndex(collectionView CollectionView, index uint) objc.IObject {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_WriteItemsAtIndexes_ToPasteboard() bool {
	return false
}

// deprecated
func (p *CollectionViewDelegateBase) CollectionView_WriteItemsAtIndexes_ToPasteboard(collectionView CollectionView, indexes foundation.IndexSet, pasteboard Pasteboard) bool {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes() bool {
	return false
}

// deprecated
func (p *CollectionViewDelegateBase) CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes(collectionView CollectionView, dropURL foundation.URL, indexes foundation.IndexSet) []string {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset(collectionView CollectionView, indexes foundation.IndexSet, event Event, dragImageOffset *foundation.Point) IImage {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexes foundation.IndexSet) {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_ValidateDrop_ProposedIndex_DropOperation() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_ValidateDrop_ProposedIndex_DropOperation(collectionView CollectionView, draggingInfo objc.Object, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_AcceptDrop_Index_DropOperation() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_AcceptDrop_Index_DropOperation(collectionView CollectionView, draggingInfo objc.Object, index int, dropOperation CollectionViewDropOperation) bool {
	panic("unimpemented protocol method")
}

type CollectionViewDelegateCreator struct {
	className string
	class     objc.Class
}

func NewCollectionViewDelegateCreator(name string) *CollectionViewDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("ProtocolBase"), name, 0)
	objc.RegisterClassPair(class)
	return &CollectionViewDelegateCreator{className: name, class: class}
}

func (c *CollectionViewDelegateCreator) SetCollectionView_ShouldSelectItemsAtIndexPaths(handle func(o objc.ProtocolBase, collectionView CollectionView, indexPaths foundation.Set) foundation.ISet) {
	objc.AddMethod(c.class, objc.SEL("collectionView:shouldSelectItemsAtIndexPaths:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_DidSelectItemsAtIndexPaths(handle func(o objc.ProtocolBase, collectionView CollectionView, indexPaths foundation.Set)) {
	objc.AddMethod(c.class, objc.SEL("collectionView:didSelectItemsAtIndexPaths:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_ShouldDeselectItemsAtIndexPaths(handle func(o objc.ProtocolBase, collectionView CollectionView, indexPaths foundation.Set) foundation.ISet) {
	objc.AddMethod(c.class, objc.SEL("collectionView:shouldDeselectItemsAtIndexPaths:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_DidDeselectItemsAtIndexPaths(handle func(o objc.ProtocolBase, collectionView CollectionView, indexPaths foundation.Set)) {
	objc.AddMethod(c.class, objc.SEL("collectionView:didDeselectItemsAtIndexPaths:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(handle func(o objc.ProtocolBase, collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) foundation.ISet) {
	objc.AddMethod(c.class, objc.SEL("collectionView:shouldChangeItemsAtIndexPaths:toHighlightState:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(handle func(o objc.ProtocolBase, collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState)) {
	objc.AddMethod(c.class, objc.SEL("collectionView:didChangeItemsAtIndexPaths:toHighlightState:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(handle func(o objc.ProtocolBase, collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)) {
	objc.AddMethod(c.class, objc.SEL("collectionView:willDisplayItem:forRepresentedObjectAtIndexPath:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(handle func(o objc.ProtocolBase, collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)) {
	objc.AddMethod(c.class, objc.SEL("collectionView:didEndDisplayingItem:forRepresentedObjectAtIndexPath:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(handle func(o objc.ProtocolBase, collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)) {
	objc.AddMethod(c.class, objc.SEL("collectionView:willDisplaySupplementaryView:forElementKind:atIndexPath:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(handle func(o objc.ProtocolBase, collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)) {
	objc.AddMethod(c.class, objc.SEL("collectionView:didEndDisplayingSupplementaryView:forElementOfKind:atIndexPath:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_TransitionLayoutForOldLayout_NewLayout(handle func(o objc.ProtocolBase, collectionView CollectionView, fromLayout CollectionViewLayout, toLayout CollectionViewLayout) ICollectionViewTransitionLayout) {
	objc.AddMethod(c.class, objc.SEL("collectionView:transitionLayoutForOldLayout:newLayout:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_CanDragItemsAtIndexPaths_WithEvent(handle func(o objc.ProtocolBase, collectionView CollectionView, indexPaths foundation.Set, event Event) bool) {
	objc.AddMethod(c.class, objc.SEL("collectionView:canDragItemsAtIndexPaths:withEvent:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_PasteboardWriterForItemAtIndexPath(handle func(o objc.ProtocolBase, collectionView CollectionView, indexPath foundation.IndexPath) objc.IObject) {
	objc.AddMethod(c.class, objc.SEL("collectionView:pasteboardWriterForItemAtIndexPath:"), handle)
}

// deprecated
func (c *CollectionViewDelegateCreator) SetCollectionView_WriteItemsAtIndexPaths_ToPasteboard(handle func(o objc.ProtocolBase, collectionView CollectionView, indexPaths foundation.Set, pasteboard Pasteboard) bool) {
	objc.AddMethod(c.class, objc.SEL("collectionView:writeItemsAtIndexPaths:toPasteboard:"), handle)
}

// deprecated
func (c *CollectionViewDelegateCreator) SetCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths(handle func(o objc.ProtocolBase, collectionView CollectionView, dropURL foundation.URL, indexPaths foundation.Set) []string) {
	objc.AddMethod(c.class, objc.SEL("collectionView:namesOfPromisedFilesDroppedAtDestination:forDraggedItemsAtIndexPaths:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset(handle func(o objc.ProtocolBase, collectionView CollectionView, indexPaths foundation.Set, event Event, dragImageOffset *foundation.Point) IImage) {
	objc.AddMethod(c.class, objc.SEL("collectionView:draggingImageForItemsAtIndexPaths:withEvent:offset:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(handle func(o objc.ProtocolBase, collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexPaths foundation.Set)) {
	objc.AddMethod(c.class, objc.SEL("collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexPaths:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_DraggingSession_EndedAtPoint_DragOperation(handle func(o objc.ProtocolBase, collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)) {
	objc.AddMethod(c.class, objc.SEL("collectionView:draggingSession:endedAtPoint:dragOperation:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_UpdateDraggingItemsForDrag(handle func(o objc.ProtocolBase, collectionView CollectionView, draggingInfo objc.Object)) {
	objc.AddMethod(c.class, objc.SEL("collectionView:updateDraggingItemsForDrag:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_ValidateDrop_ProposedIndexPath_DropOperation(handle func(o objc.ProtocolBase, collectionView CollectionView, draggingInfo objc.Object, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation) {
	objc.AddMethod(c.class, objc.SEL("collectionView:validateDrop:proposedIndexPath:dropOperation:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_AcceptDrop_IndexPath_DropOperation(handle func(o objc.ProtocolBase, collectionView CollectionView, draggingInfo objc.Object, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool) {
	objc.AddMethod(c.class, objc.SEL("collectionView:acceptDrop:indexPath:dropOperation:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_CanDragItemsAtIndexes_WithEvent(handle func(o objc.ProtocolBase, collectionView CollectionView, indexes foundation.IndexSet, event Event) bool) {
	objc.AddMethod(c.class, objc.SEL("collectionView:canDragItemsAtIndexes:withEvent:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_PasteboardWriterForItemAtIndex(handle func(o objc.ProtocolBase, collectionView CollectionView, index uint) objc.IObject) {
	objc.AddMethod(c.class, objc.SEL("collectionView:pasteboardWriterForItemAtIndex:"), handle)
}

// deprecated
func (c *CollectionViewDelegateCreator) SetCollectionView_WriteItemsAtIndexes_ToPasteboard(handle func(o objc.ProtocolBase, collectionView CollectionView, indexes foundation.IndexSet, pasteboard Pasteboard) bool) {
	objc.AddMethod(c.class, objc.SEL("collectionView:writeItemsAtIndexes:toPasteboard:"), handle)
}

// deprecated
func (c *CollectionViewDelegateCreator) SetCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes(handle func(o objc.ProtocolBase, collectionView CollectionView, dropURL foundation.URL, indexes foundation.IndexSet) []string) {
	objc.AddMethod(c.class, objc.SEL("collectionView:namesOfPromisedFilesDroppedAtDestination:forDraggedItemsAtIndexes:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset(handle func(o objc.ProtocolBase, collectionView CollectionView, indexes foundation.IndexSet, event Event, dragImageOffset *foundation.Point) IImage) {
	objc.AddMethod(c.class, objc.SEL("collectionView:draggingImageForItemsAtIndexes:withEvent:offset:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(handle func(o objc.ProtocolBase, collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexes foundation.IndexSet)) {
	objc.AddMethod(c.class, objc.SEL("collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexes:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_ValidateDrop_ProposedIndex_DropOperation(handle func(o objc.ProtocolBase, collectionView CollectionView, draggingInfo objc.Object, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation) {
	objc.AddMethod(c.class, objc.SEL("collectionView:validateDrop:proposedIndex:dropOperation:"), handle)
}

func (c *CollectionViewDelegateCreator) SetCollectionView_AcceptDrop_Index_DropOperation(handle func(o objc.ProtocolBase, collectionView CollectionView, draggingInfo objc.Object, index int, dropOperation CollectionViewDropOperation) bool) {
	objc.AddMethod(c.class, objc.SEL("collectionView:acceptDrop:index:dropOperation:"), handle)
}

func (c *CollectionViewDelegateCreator) Create() objc.ProtocolBase {
	return objc.ProtocolBase{Object: c.class.CreateInstance(0)}
}
