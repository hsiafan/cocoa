// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	CollectionView_UpdateDraggingItemsForDrag(collectionView CollectionView, draggingInfo DraggingInfoWrapper)
	ImplementsCollectionView_ValidateDrop_ProposedIndexPath_DropOperation() bool
	// optional
	CollectionView_ValidateDrop_ProposedIndexPath_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation
	ImplementsCollectionView_AcceptDrop_IndexPath_DropOperation() bool
	// optional
	CollectionView_AcceptDrop_IndexPath_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool
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
	CollectionView_ValidateDrop_ProposedIndex_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation
	ImplementsCollectionView_AcceptDrop_Index_DropOperation() bool
	// optional
	CollectionView_AcceptDrop_Index_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, index int, dropOperation CollectionViewDropOperation) bool
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

func (p *CollectionViewDelegateBase) CollectionView_UpdateDraggingItemsForDrag(collectionView CollectionView, draggingInfo DraggingInfoWrapper) {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_ValidateDrop_ProposedIndexPath_DropOperation() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_ValidateDrop_ProposedIndexPath_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_AcceptDrop_IndexPath_DropOperation() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_AcceptDrop_IndexPath_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool {
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

func (p *CollectionViewDelegateBase) CollectionView_ValidateDrop_ProposedIndex_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateBase) ImplementsCollectionView_AcceptDrop_Index_DropOperation() bool {
	return false
}

func (p *CollectionViewDelegateBase) CollectionView_AcceptDrop_Index_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, index int, dropOperation CollectionViewDropOperation) bool {
	panic("unimpemented protocol method")
}

type CollectionViewDelegateWrapper struct {
	objc.Object
}

func (c_ CollectionViewDelegateWrapper) CollectionView_ShouldSelectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, objc.GetSelector("collectionView:shouldSelectItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths))
	return rv
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DidSelectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:didSelectItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_ShouldDeselectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, objc.GetSelector("collectionView:shouldDeselectItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths))
	return rv
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DidDeselectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:didDeselectItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(collectionView ICollectionView, indexPaths foundation.ISet, highlightState CollectionViewItemHighlightState) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, objc.GetSelector("collectionView:shouldChangeItemsAtIndexPaths:toHighlightState:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths), highlightState)
	return rv
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(collectionView ICollectionView, indexPaths foundation.ISet, highlightState CollectionViewItemHighlightState) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:didChangeItemsAtIndexPaths:toHighlightState:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths), highlightState)
}

func (c_ CollectionViewDelegateWrapper) CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(collectionView ICollectionView, item ICollectionViewItem, indexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:willDisplayItem:forRepresentedObjectAtIndexPath:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(item), objc.ExtractPtr(indexPath))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(collectionView ICollectionView, item ICollectionViewItem, indexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:didEndDisplayingItem:forRepresentedObjectAtIndexPath:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(item), objc.ExtractPtr(indexPath))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(collectionView ICollectionView, view IView, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:willDisplaySupplementaryView:forElementKind:atIndexPath:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(view), elementKind, objc.ExtractPtr(indexPath))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(collectionView ICollectionView, view IView, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:didEndDisplayingSupplementaryView:forElementOfKind:atIndexPath:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(view), elementKind, objc.ExtractPtr(indexPath))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_TransitionLayoutForOldLayout_NewLayout(collectionView ICollectionView, fromLayout ICollectionViewLayout, toLayout ICollectionViewLayout) CollectionViewTransitionLayout {
	rv := objc.CallMethod[CollectionViewTransitionLayout](c_, objc.GetSelector("collectionView:transitionLayoutForOldLayout:newLayout:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(fromLayout), objc.ExtractPtr(toLayout))
	return rv
}

func (c_ CollectionViewDelegateWrapper) CollectionView_CanDragItemsAtIndexPaths_WithEvent(collectionView ICollectionView, indexPaths foundation.ISet, event IEvent) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("collectionView:canDragItemsAtIndexPaths:withEvent:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths), objc.ExtractPtr(event))
	return rv
}

func (c_ CollectionViewDelegateWrapper) CollectionView_PasteboardWriterForItemAtIndexPath(collectionView ICollectionView, indexPath foundation.IIndexPath) PasteboardWritingWrapper {
	rv := objc.CallMethod[PasteboardWritingWrapper](c_, objc.GetSelector("collectionView:pasteboardWriterForItemAtIndexPath:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPath))
	return rv
}

// deprecated
func (c_ CollectionViewDelegateWrapper) CollectionView_WriteItemsAtIndexPaths_ToPasteboard(collectionView ICollectionView, indexPaths foundation.ISet, pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("collectionView:writeItemsAtIndexPaths:toPasteboard:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths), objc.ExtractPtr(pasteboard))
	return rv
}

// deprecated
func (c_ CollectionViewDelegateWrapper) CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths(collectionView ICollectionView, dropURL foundation.IURL, indexPaths foundation.ISet) []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("collectionView:namesOfPromisedFilesDroppedAtDestination:forDraggedItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(dropURL), objc.ExtractPtr(indexPaths))
	return rv
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset(collectionView ICollectionView, indexPaths foundation.ISet, event IEvent, dragImageOffset *foundation.Point) Image {
	rv := objc.CallMethod[Image](c_, objc.GetSelector("collectionView:draggingImageForItemsAtIndexPaths:withEvent:offset:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths), objc.ExtractPtr(event), dragImageOffset)
	return rv
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(collectionView ICollectionView, session IDraggingSession, screenPoint foundation.Point, indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(session), screenPoint, objc.ExtractPtr(indexPaths))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DraggingSession_EndedAtPoint_DragOperation(collectionView ICollectionView, session IDraggingSession, screenPoint foundation.Point, operation DragOperation) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:draggingSession:endedAtPoint:dragOperation:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(session), screenPoint, operation)
}

func (c_ CollectionViewDelegateWrapper) CollectionView_UpdateDraggingItemsForDrag(collectionView ICollectionView, draggingInfo objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:updateDraggingItemsForDrag:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(draggingInfo))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_ValidateDrop_ProposedIndexPath_DropOperation(collectionView ICollectionView, draggingInfo objc.IObject, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	rv := objc.CallMethod[DragOperation](c_, objc.GetSelector("collectionView:validateDrop:proposedIndexPath:dropOperation:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(draggingInfo), unsafe.Pointer(proposedDropIndexPath), proposedDropOperation)
	return rv
}

func (c_ CollectionViewDelegateWrapper) CollectionView_AcceptDrop_IndexPath_DropOperation(collectionView ICollectionView, draggingInfo objc.IObject, indexPath foundation.IIndexPath, dropOperation CollectionViewDropOperation) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("collectionView:acceptDrop:indexPath:dropOperation:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(draggingInfo), objc.ExtractPtr(indexPath), dropOperation)
	return rv
}

func (c_ CollectionViewDelegateWrapper) CollectionView_CanDragItemsAtIndexes_WithEvent(collectionView ICollectionView, indexes foundation.IIndexSet, event IEvent) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("collectionView:canDragItemsAtIndexes:withEvent:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexes), objc.ExtractPtr(event))
	return rv
}

func (c_ CollectionViewDelegateWrapper) CollectionView_PasteboardWriterForItemAtIndex(collectionView ICollectionView, index uint) PasteboardWritingWrapper {
	rv := objc.CallMethod[PasteboardWritingWrapper](c_, objc.GetSelector("collectionView:pasteboardWriterForItemAtIndex:"), objc.ExtractPtr(collectionView), index)
	return rv
}

// deprecated
func (c_ CollectionViewDelegateWrapper) CollectionView_WriteItemsAtIndexes_ToPasteboard(collectionView ICollectionView, indexes foundation.IIndexSet, pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("collectionView:writeItemsAtIndexes:toPasteboard:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexes), objc.ExtractPtr(pasteboard))
	return rv
}

// deprecated
func (c_ CollectionViewDelegateWrapper) CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes(collectionView ICollectionView, dropURL foundation.IURL, indexes foundation.IIndexSet) []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("collectionView:namesOfPromisedFilesDroppedAtDestination:forDraggedItemsAtIndexes:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(dropURL), objc.ExtractPtr(indexes))
	return rv
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset(collectionView ICollectionView, indexes foundation.IIndexSet, event IEvent, dragImageOffset *foundation.Point) Image {
	rv := objc.CallMethod[Image](c_, objc.GetSelector("collectionView:draggingImageForItemsAtIndexes:withEvent:offset:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexes), objc.ExtractPtr(event), dragImageOffset)
	return rv
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(collectionView ICollectionView, session IDraggingSession, screenPoint foundation.Point, indexes foundation.IIndexSet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexes:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(session), screenPoint, objc.ExtractPtr(indexes))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_ValidateDrop_ProposedIndex_DropOperation(collectionView ICollectionView, draggingInfo objc.IObject, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	rv := objc.CallMethod[DragOperation](c_, objc.GetSelector("collectionView:validateDrop:proposedIndex:dropOperation:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(draggingInfo), proposedDropIndex, proposedDropOperation)
	return rv
}

func (c_ CollectionViewDelegateWrapper) CollectionView_AcceptDrop_Index_DropOperation(collectionView ICollectionView, draggingInfo objc.IObject, index int, dropOperation CollectionViewDropOperation) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("collectionView:acceptDrop:index:dropOperation:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(draggingInfo), index, dropOperation)
	return rv
}
