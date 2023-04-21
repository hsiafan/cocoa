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
	CollectionView_PasteboardWriterForItemAtIndexPath(collectionView CollectionView, indexPath foundation.IndexPath) PasteboardWriting
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
	CollectionView_PasteboardWriterForItemAtIndex(collectionView CollectionView, index uint) PasteboardWriting
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

type CollectionViewDelegateImpl struct {
	_CollectionView_ShouldSelectItemsAtIndexPaths                                        func(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet
	_CollectionView_DidSelectItemsAtIndexPaths                                           func(collectionView CollectionView, indexPaths foundation.Set)
	_CollectionView_ShouldDeselectItemsAtIndexPaths                                      func(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet
	_CollectionView_DidDeselectItemsAtIndexPaths                                         func(collectionView CollectionView, indexPaths foundation.Set)
	_CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState                       func(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) foundation.ISet
	_CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState                          func(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState)
	_CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath                      func(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)
	_CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath                 func(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)
	_CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath              func(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)
	_CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath       func(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)
	_CollectionView_TransitionLayoutForOldLayout_NewLayout                               func(collectionView CollectionView, fromLayout CollectionViewLayout, toLayout CollectionViewLayout) ICollectionViewTransitionLayout
	_CollectionView_CanDragItemsAtIndexPaths_WithEvent                                   func(collectionView CollectionView, indexPaths foundation.Set, event Event) bool
	_CollectionView_PasteboardWriterForItemAtIndexPath                                   func(collectionView CollectionView, indexPath foundation.IndexPath) PasteboardWriting
	_CollectionView_WriteItemsAtIndexPaths_ToPasteboard                                  func(collectionView CollectionView, indexPaths foundation.Set, pasteboard Pasteboard) bool
	_CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths func(collectionView CollectionView, dropURL foundation.URL, indexPaths foundation.Set) []string
	_CollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset                   func(collectionView CollectionView, indexPaths foundation.Set, event Event, dragImageOffset *foundation.Point) IImage
	_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths                func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexPaths foundation.Set)
	_CollectionView_DraggingSession_EndedAtPoint_DragOperation                           func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	_CollectionView_UpdateDraggingItemsForDrag                                           func(collectionView CollectionView, draggingInfo DraggingInfoWrapper)
	_CollectionView_ValidateDrop_ProposedIndexPath_DropOperation                         func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation
	_CollectionView_AcceptDrop_IndexPath_DropOperation                                   func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool
	_CollectionView_CanDragItemsAtIndexes_WithEvent                                      func(collectionView CollectionView, indexes foundation.IndexSet, event Event) bool
	_CollectionView_PasteboardWriterForItemAtIndex                                       func(collectionView CollectionView, index uint) PasteboardWriting
	_CollectionView_WriteItemsAtIndexes_ToPasteboard                                     func(collectionView CollectionView, indexes foundation.IndexSet, pasteboard Pasteboard) bool
	_CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes    func(collectionView CollectionView, dropURL foundation.URL, indexes foundation.IndexSet) []string
	_CollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset                      func(collectionView CollectionView, indexes foundation.IndexSet, event Event, dragImageOffset *foundation.Point) IImage
	_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes                   func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexes foundation.IndexSet)
	_CollectionView_ValidateDrop_ProposedIndex_DropOperation                             func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation
	_CollectionView_AcceptDrop_Index_DropOperation                                       func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, index int, dropOperation CollectionViewDropOperation) bool
}

func (di *CollectionViewDelegateImpl) ImplementsCollectionView_ShouldSelectItemsAtIndexPaths() bool {
	return di._CollectionView_ShouldSelectItemsAtIndexPaths != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_ShouldSelectItemsAtIndexPaths(f func(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet) {
	di._CollectionView_ShouldSelectItemsAtIndexPaths = f
}

func (di *CollectionViewDelegateImpl) CollectionView_ShouldSelectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet {
	return di._CollectionView_ShouldSelectItemsAtIndexPaths(collectionView, indexPaths)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DidSelectItemsAtIndexPaths() bool {
	return di._CollectionView_DidSelectItemsAtIndexPaths != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DidSelectItemsAtIndexPaths(f func(collectionView CollectionView, indexPaths foundation.Set)) {
	di._CollectionView_DidSelectItemsAtIndexPaths = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DidSelectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) {
	di._CollectionView_DidSelectItemsAtIndexPaths(collectionView, indexPaths)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_ShouldDeselectItemsAtIndexPaths() bool {
	return di._CollectionView_ShouldDeselectItemsAtIndexPaths != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_ShouldDeselectItemsAtIndexPaths(f func(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet) {
	di._CollectionView_ShouldDeselectItemsAtIndexPaths = f
}

func (di *CollectionViewDelegateImpl) CollectionView_ShouldDeselectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet {
	return di._CollectionView_ShouldDeselectItemsAtIndexPaths(collectionView, indexPaths)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DidDeselectItemsAtIndexPaths() bool {
	return di._CollectionView_DidDeselectItemsAtIndexPaths != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DidDeselectItemsAtIndexPaths(f func(collectionView CollectionView, indexPaths foundation.Set)) {
	di._CollectionView_DidDeselectItemsAtIndexPaths = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DidDeselectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) {
	di._CollectionView_DidDeselectItemsAtIndexPaths(collectionView, indexPaths)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState() bool {
	return di._CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(f func(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) foundation.ISet) {
	di._CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState = f
}

func (di *CollectionViewDelegateImpl) CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) foundation.ISet {
	return di._CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(collectionView, indexPaths, highlightState)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DidChangeItemsAtIndexPaths_ToHighlightState() bool {
	return di._CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(f func(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState)) {
	di._CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) {
	di._CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(collectionView, indexPaths, highlightState)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath() bool {
	return di._CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(f func(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)) {
	di._CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath = f
}

func (di *CollectionViewDelegateImpl) CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath) {
	di._CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(collectionView, item, indexPath)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath() bool {
	return di._CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(f func(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)) {
	di._CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath) {
	di._CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(collectionView, item, indexPath)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath() bool {
	return di._CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(f func(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)) {
	di._CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath = f
}

func (di *CollectionViewDelegateImpl) CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) {
	di._CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(collectionView, view, elementKind, indexPath)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath() bool {
	return di._CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(f func(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)) {
	di._CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) {
	di._CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(collectionView, view, elementKind, indexPath)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_TransitionLayoutForOldLayout_NewLayout() bool {
	return di._CollectionView_TransitionLayoutForOldLayout_NewLayout != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_TransitionLayoutForOldLayout_NewLayout(f func(collectionView CollectionView, fromLayout CollectionViewLayout, toLayout CollectionViewLayout) ICollectionViewTransitionLayout) {
	di._CollectionView_TransitionLayoutForOldLayout_NewLayout = f
}

func (di *CollectionViewDelegateImpl) CollectionView_TransitionLayoutForOldLayout_NewLayout(collectionView CollectionView, fromLayout CollectionViewLayout, toLayout CollectionViewLayout) ICollectionViewTransitionLayout {
	return di._CollectionView_TransitionLayoutForOldLayout_NewLayout(collectionView, fromLayout, toLayout)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_CanDragItemsAtIndexPaths_WithEvent() bool {
	return di._CollectionView_CanDragItemsAtIndexPaths_WithEvent != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_CanDragItemsAtIndexPaths_WithEvent(f func(collectionView CollectionView, indexPaths foundation.Set, event Event) bool) {
	di._CollectionView_CanDragItemsAtIndexPaths_WithEvent = f
}

func (di *CollectionViewDelegateImpl) CollectionView_CanDragItemsAtIndexPaths_WithEvent(collectionView CollectionView, indexPaths foundation.Set, event Event) bool {
	return di._CollectionView_CanDragItemsAtIndexPaths_WithEvent(collectionView, indexPaths, event)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_PasteboardWriterForItemAtIndexPath() bool {
	return di._CollectionView_PasteboardWriterForItemAtIndexPath != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_PasteboardWriterForItemAtIndexPath(f func(collectionView CollectionView, indexPath foundation.IndexPath) PasteboardWriting) {
	di._CollectionView_PasteboardWriterForItemAtIndexPath = f
}

func (di *CollectionViewDelegateImpl) CollectionView_PasteboardWriterForItemAtIndexPath(collectionView CollectionView, indexPath foundation.IndexPath) PasteboardWriting {
	return di._CollectionView_PasteboardWriterForItemAtIndexPath(collectionView, indexPath)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_WriteItemsAtIndexPaths_ToPasteboard() bool {
	return di._CollectionView_WriteItemsAtIndexPaths_ToPasteboard != nil
}

// deprecated
func (di *CollectionViewDelegateImpl) SetCollectionView_WriteItemsAtIndexPaths_ToPasteboard(f func(collectionView CollectionView, indexPaths foundation.Set, pasteboard Pasteboard) bool) {
	di._CollectionView_WriteItemsAtIndexPaths_ToPasteboard = f
}

// deprecated
func (di *CollectionViewDelegateImpl) CollectionView_WriteItemsAtIndexPaths_ToPasteboard(collectionView CollectionView, indexPaths foundation.Set, pasteboard Pasteboard) bool {
	return di._CollectionView_WriteItemsAtIndexPaths_ToPasteboard(collectionView, indexPaths, pasteboard)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths() bool {
	return di._CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths != nil
}

// deprecated
func (di *CollectionViewDelegateImpl) SetCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths(f func(collectionView CollectionView, dropURL foundation.URL, indexPaths foundation.Set) []string) {
	di._CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths = f
}

// deprecated
func (di *CollectionViewDelegateImpl) CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths(collectionView CollectionView, dropURL foundation.URL, indexPaths foundation.Set) []string {
	return di._CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths(collectionView, dropURL, indexPaths)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset() bool {
	return di._CollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset(f func(collectionView CollectionView, indexPaths foundation.Set, event Event, dragImageOffset *foundation.Point) IImage) {
	di._CollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset(collectionView CollectionView, indexPaths foundation.Set, event Event, dragImageOffset *foundation.Point) IImage {
	return di._CollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset(collectionView, indexPaths, event, dragImageOffset)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths() bool {
	return di._CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(f func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexPaths foundation.Set)) {
	di._CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexPaths foundation.Set) {
	di._CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(collectionView, session, screenPoint, indexPaths)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DraggingSession_EndedAtPoint_DragOperation() bool {
	return di._CollectionView_DraggingSession_EndedAtPoint_DragOperation != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DraggingSession_EndedAtPoint_DragOperation(f func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)) {
	di._CollectionView_DraggingSession_EndedAtPoint_DragOperation = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DraggingSession_EndedAtPoint_DragOperation(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, operation DragOperation) {
	di._CollectionView_DraggingSession_EndedAtPoint_DragOperation(collectionView, session, screenPoint, operation)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_UpdateDraggingItemsForDrag() bool {
	return di._CollectionView_UpdateDraggingItemsForDrag != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_UpdateDraggingItemsForDrag(f func(collectionView CollectionView, draggingInfo DraggingInfoWrapper)) {
	di._CollectionView_UpdateDraggingItemsForDrag = f
}

func (di *CollectionViewDelegateImpl) CollectionView_UpdateDraggingItemsForDrag(collectionView CollectionView, draggingInfo DraggingInfoWrapper) {
	di._CollectionView_UpdateDraggingItemsForDrag(collectionView, draggingInfo)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_ValidateDrop_ProposedIndexPath_DropOperation() bool {
	return di._CollectionView_ValidateDrop_ProposedIndexPath_DropOperation != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_ValidateDrop_ProposedIndexPath_DropOperation(f func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation) {
	di._CollectionView_ValidateDrop_ProposedIndexPath_DropOperation = f
}

func (di *CollectionViewDelegateImpl) CollectionView_ValidateDrop_ProposedIndexPath_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	return di._CollectionView_ValidateDrop_ProposedIndexPath_DropOperation(collectionView, draggingInfo, proposedDropIndexPath, proposedDropOperation)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_AcceptDrop_IndexPath_DropOperation() bool {
	return di._CollectionView_AcceptDrop_IndexPath_DropOperation != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_AcceptDrop_IndexPath_DropOperation(f func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool) {
	di._CollectionView_AcceptDrop_IndexPath_DropOperation = f
}

func (di *CollectionViewDelegateImpl) CollectionView_AcceptDrop_IndexPath_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool {
	return di._CollectionView_AcceptDrop_IndexPath_DropOperation(collectionView, draggingInfo, indexPath, dropOperation)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_CanDragItemsAtIndexes_WithEvent() bool {
	return di._CollectionView_CanDragItemsAtIndexes_WithEvent != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_CanDragItemsAtIndexes_WithEvent(f func(collectionView CollectionView, indexes foundation.IndexSet, event Event) bool) {
	di._CollectionView_CanDragItemsAtIndexes_WithEvent = f
}

func (di *CollectionViewDelegateImpl) CollectionView_CanDragItemsAtIndexes_WithEvent(collectionView CollectionView, indexes foundation.IndexSet, event Event) bool {
	return di._CollectionView_CanDragItemsAtIndexes_WithEvent(collectionView, indexes, event)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_PasteboardWriterForItemAtIndex() bool {
	return di._CollectionView_PasteboardWriterForItemAtIndex != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_PasteboardWriterForItemAtIndex(f func(collectionView CollectionView, index uint) PasteboardWriting) {
	di._CollectionView_PasteboardWriterForItemAtIndex = f
}

func (di *CollectionViewDelegateImpl) CollectionView_PasteboardWriterForItemAtIndex(collectionView CollectionView, index uint) PasteboardWriting {
	return di._CollectionView_PasteboardWriterForItemAtIndex(collectionView, index)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_WriteItemsAtIndexes_ToPasteboard() bool {
	return di._CollectionView_WriteItemsAtIndexes_ToPasteboard != nil
}

// deprecated
func (di *CollectionViewDelegateImpl) SetCollectionView_WriteItemsAtIndexes_ToPasteboard(f func(collectionView CollectionView, indexes foundation.IndexSet, pasteboard Pasteboard) bool) {
	di._CollectionView_WriteItemsAtIndexes_ToPasteboard = f
}

// deprecated
func (di *CollectionViewDelegateImpl) CollectionView_WriteItemsAtIndexes_ToPasteboard(collectionView CollectionView, indexes foundation.IndexSet, pasteboard Pasteboard) bool {
	return di._CollectionView_WriteItemsAtIndexes_ToPasteboard(collectionView, indexes, pasteboard)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes() bool {
	return di._CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes != nil
}

// deprecated
func (di *CollectionViewDelegateImpl) SetCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes(f func(collectionView CollectionView, dropURL foundation.URL, indexes foundation.IndexSet) []string) {
	di._CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes = f
}

// deprecated
func (di *CollectionViewDelegateImpl) CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes(collectionView CollectionView, dropURL foundation.URL, indexes foundation.IndexSet) []string {
	return di._CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes(collectionView, dropURL, indexes)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset() bool {
	return di._CollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset(f func(collectionView CollectionView, indexes foundation.IndexSet, event Event, dragImageOffset *foundation.Point) IImage) {
	di._CollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset(collectionView CollectionView, indexes foundation.IndexSet, event Event, dragImageOffset *foundation.Point) IImage {
	return di._CollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset(collectionView, indexes, event, dragImageOffset)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes() bool {
	return di._CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(f func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexes foundation.IndexSet)) {
	di._CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexes foundation.IndexSet) {
	di._CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(collectionView, session, screenPoint, indexes)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_ValidateDrop_ProposedIndex_DropOperation() bool {
	return di._CollectionView_ValidateDrop_ProposedIndex_DropOperation != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_ValidateDrop_ProposedIndex_DropOperation(f func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation) {
	di._CollectionView_ValidateDrop_ProposedIndex_DropOperation = f
}

func (di *CollectionViewDelegateImpl) CollectionView_ValidateDrop_ProposedIndex_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	return di._CollectionView_ValidateDrop_ProposedIndex_DropOperation(collectionView, draggingInfo, proposedDropIndex, proposedDropOperation)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_AcceptDrop_Index_DropOperation() bool {
	return di._CollectionView_AcceptDrop_Index_DropOperation != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_AcceptDrop_Index_DropOperation(f func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, index int, dropOperation CollectionViewDropOperation) bool) {
	di._CollectionView_AcceptDrop_Index_DropOperation = f
}

func (di *CollectionViewDelegateImpl) CollectionView_AcceptDrop_Index_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, index int, dropOperation CollectionViewDropOperation) bool {
	return di._CollectionView_AcceptDrop_Index_DropOperation(collectionView, draggingInfo, index, dropOperation)
}

type CollectionViewDelegateWrapper struct {
	objc.Object
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_ShouldSelectItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:shouldSelectItemsAtIndexPaths:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_ShouldSelectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, "collectionView:shouldSelectItemsAtIndexPaths:", collectionView, indexPaths)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DidSelectItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:didSelectItemsAtIndexPaths:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DidSelectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, "collectionView:didSelectItemsAtIndexPaths:", collectionView, indexPaths)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_ShouldDeselectItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:shouldDeselectItemsAtIndexPaths:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_ShouldDeselectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, "collectionView:shouldDeselectItemsAtIndexPaths:", collectionView, indexPaths)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DidDeselectItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:didDeselectItemsAtIndexPaths:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DidDeselectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, "collectionView:didDeselectItemsAtIndexPaths:", collectionView, indexPaths)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:shouldChangeItemsAtIndexPaths:toHighlightState:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(collectionView ICollectionView, indexPaths foundation.ISet, highlightState CollectionViewItemHighlightState) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, "collectionView:shouldChangeItemsAtIndexPaths:toHighlightState:", collectionView, indexPaths, highlightState)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DidChangeItemsAtIndexPaths_ToHighlightState() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:didChangeItemsAtIndexPaths:toHighlightState:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(collectionView ICollectionView, indexPaths foundation.ISet, highlightState CollectionViewItemHighlightState) {
	objc.CallMethod[objc.Void](c_, "collectionView:didChangeItemsAtIndexPaths:toHighlightState:", collectionView, indexPaths, highlightState)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:willDisplayItem:forRepresentedObjectAtIndexPath:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(collectionView ICollectionView, item ICollectionViewItem, indexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, "collectionView:willDisplayItem:forRepresentedObjectAtIndexPath:", collectionView, item, indexPath)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:didEndDisplayingItem:forRepresentedObjectAtIndexPath:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(collectionView ICollectionView, item ICollectionViewItem, indexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, "collectionView:didEndDisplayingItem:forRepresentedObjectAtIndexPath:", collectionView, item, indexPath)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:willDisplaySupplementaryView:forElementKind:atIndexPath:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(collectionView ICollectionView, view IView, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, "collectionView:willDisplaySupplementaryView:forElementKind:atIndexPath:", collectionView, view, elementKind, indexPath)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:didEndDisplayingSupplementaryView:forElementOfKind:atIndexPath:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(collectionView ICollectionView, view IView, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, "collectionView:didEndDisplayingSupplementaryView:forElementOfKind:atIndexPath:", collectionView, view, elementKind, indexPath)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_TransitionLayoutForOldLayout_NewLayout() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:transitionLayoutForOldLayout:newLayout:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_TransitionLayoutForOldLayout_NewLayout(collectionView ICollectionView, fromLayout ICollectionViewLayout, toLayout ICollectionViewLayout) CollectionViewTransitionLayout {
	rv := objc.CallMethod[CollectionViewTransitionLayout](c_, "collectionView:transitionLayoutForOldLayout:newLayout:", collectionView, fromLayout, toLayout)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_CanDragItemsAtIndexPaths_WithEvent() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:canDragItemsAtIndexPaths:withEvent:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_CanDragItemsAtIndexPaths_WithEvent(collectionView ICollectionView, indexPaths foundation.ISet, event IEvent) bool {
	rv := objc.CallMethod[bool](c_, "collectionView:canDragItemsAtIndexPaths:withEvent:", collectionView, indexPaths, event)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_PasteboardWriterForItemAtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:pasteboardWriterForItemAtIndexPath:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_PasteboardWriterForItemAtIndexPath(collectionView ICollectionView, indexPath foundation.IIndexPath) PasteboardWritingWrapper {
	rv := objc.CallMethod[PasteboardWritingWrapper](c_, "collectionView:pasteboardWriterForItemAtIndexPath:", collectionView, indexPath)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_WriteItemsAtIndexPaths_ToPasteboard() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:writeItemsAtIndexPaths:toPasteboard:"))
}

// deprecated
func (c_ CollectionViewDelegateWrapper) CollectionView_WriteItemsAtIndexPaths_ToPasteboard(collectionView ICollectionView, indexPaths foundation.ISet, pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](c_, "collectionView:writeItemsAtIndexPaths:toPasteboard:", collectionView, indexPaths, pasteboard)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:namesOfPromisedFilesDroppedAtDestination:forDraggedItemsAtIndexPaths:"))
}

// deprecated
func (c_ CollectionViewDelegateWrapper) CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths(collectionView ICollectionView, dropURL foundation.IURL, indexPaths foundation.ISet) []string {
	rv := objc.CallMethod[[]string](c_, "collectionView:namesOfPromisedFilesDroppedAtDestination:forDraggedItemsAtIndexPaths:", collectionView, dropURL, indexPaths)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:draggingImageForItemsAtIndexPaths:withEvent:offset:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset(collectionView ICollectionView, indexPaths foundation.ISet, event IEvent, dragImageOffset *foundation.Point) Image {
	rv := objc.CallMethod[Image](c_, "collectionView:draggingImageForItemsAtIndexPaths:withEvent:offset:", collectionView, indexPaths, event, dragImageOffset)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexPaths:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(collectionView ICollectionView, session IDraggingSession, screenPoint foundation.Point, indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, "collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexPaths:", collectionView, session, screenPoint, indexPaths)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DraggingSession_EndedAtPoint_DragOperation() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:draggingSession:endedAtPoint:dragOperation:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DraggingSession_EndedAtPoint_DragOperation(collectionView ICollectionView, session IDraggingSession, screenPoint foundation.Point, operation DragOperation) {
	objc.CallMethod[objc.Void](c_, "collectionView:draggingSession:endedAtPoint:dragOperation:", collectionView, session, screenPoint, operation)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_UpdateDraggingItemsForDrag() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:updateDraggingItemsForDrag:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_UpdateDraggingItemsForDrag(collectionView ICollectionView, draggingInfo DraggingInfo) {
	po := objc.CreateProtocol("NSDraggingInfo", draggingInfo)
	defer po.Release()
	objc.CallMethod[objc.Void](c_, "collectionView:updateDraggingItemsForDrag:", collectionView, po)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_ValidateDrop_ProposedIndexPath_DropOperation() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:validateDrop:proposedIndexPath:dropOperation:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_ValidateDrop_ProposedIndexPath_DropOperation(collectionView ICollectionView, draggingInfo DraggingInfo, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	po := objc.CreateProtocol("NSDraggingInfo", draggingInfo)
	defer po.Release()
	rv := objc.CallMethod[DragOperation](c_, "collectionView:validateDrop:proposedIndexPath:dropOperation:", collectionView, po, unsafe.Pointer(proposedDropIndexPath), proposedDropOperation)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_AcceptDrop_IndexPath_DropOperation() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:acceptDrop:indexPath:dropOperation:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_AcceptDrop_IndexPath_DropOperation(collectionView ICollectionView, draggingInfo DraggingInfo, indexPath foundation.IIndexPath, dropOperation CollectionViewDropOperation) bool {
	po := objc.CreateProtocol("NSDraggingInfo", draggingInfo)
	defer po.Release()
	rv := objc.CallMethod[bool](c_, "collectionView:acceptDrop:indexPath:dropOperation:", collectionView, po, indexPath, dropOperation)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_CanDragItemsAtIndexes_WithEvent() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:canDragItemsAtIndexes:withEvent:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_CanDragItemsAtIndexes_WithEvent(collectionView ICollectionView, indexes foundation.IIndexSet, event IEvent) bool {
	rv := objc.CallMethod[bool](c_, "collectionView:canDragItemsAtIndexes:withEvent:", collectionView, indexes, event)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_PasteboardWriterForItemAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:pasteboardWriterForItemAtIndex:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_PasteboardWriterForItemAtIndex(collectionView ICollectionView, index uint) PasteboardWritingWrapper {
	rv := objc.CallMethod[PasteboardWritingWrapper](c_, "collectionView:pasteboardWriterForItemAtIndex:", collectionView, index)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_WriteItemsAtIndexes_ToPasteboard() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:writeItemsAtIndexes:toPasteboard:"))
}

// deprecated
func (c_ CollectionViewDelegateWrapper) CollectionView_WriteItemsAtIndexes_ToPasteboard(collectionView ICollectionView, indexes foundation.IIndexSet, pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](c_, "collectionView:writeItemsAtIndexes:toPasteboard:", collectionView, indexes, pasteboard)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:namesOfPromisedFilesDroppedAtDestination:forDraggedItemsAtIndexes:"))
}

// deprecated
func (c_ CollectionViewDelegateWrapper) CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes(collectionView ICollectionView, dropURL foundation.IURL, indexes foundation.IIndexSet) []string {
	rv := objc.CallMethod[[]string](c_, "collectionView:namesOfPromisedFilesDroppedAtDestination:forDraggedItemsAtIndexes:", collectionView, dropURL, indexes)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:draggingImageForItemsAtIndexes:withEvent:offset:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset(collectionView ICollectionView, indexes foundation.IIndexSet, event IEvent, dragImageOffset *foundation.Point) Image {
	rv := objc.CallMethod[Image](c_, "collectionView:draggingImageForItemsAtIndexes:withEvent:offset:", collectionView, indexes, event, dragImageOffset)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexes:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(collectionView ICollectionView, session IDraggingSession, screenPoint foundation.Point, indexes foundation.IIndexSet) {
	objc.CallMethod[objc.Void](c_, "collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexes:", collectionView, session, screenPoint, indexes)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_ValidateDrop_ProposedIndex_DropOperation() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:validateDrop:proposedIndex:dropOperation:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_ValidateDrop_ProposedIndex_DropOperation(collectionView ICollectionView, draggingInfo DraggingInfo, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	po := objc.CreateProtocol("NSDraggingInfo", draggingInfo)
	defer po.Release()
	rv := objc.CallMethod[DragOperation](c_, "collectionView:validateDrop:proposedIndex:dropOperation:", collectionView, po, proposedDropIndex, proposedDropOperation)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_AcceptDrop_Index_DropOperation() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:acceptDrop:index:dropOperation:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_AcceptDrop_Index_DropOperation(collectionView ICollectionView, draggingInfo DraggingInfo, index int, dropOperation CollectionViewDropOperation) bool {
	po := objc.CreateProtocol("NSDraggingInfo", draggingInfo)
	defer po.Release()
	rv := objc.CallMethod[bool](c_, "collectionView:acceptDrop:index:dropOperation:", collectionView, po, index, dropOperation)
	return rv
}
