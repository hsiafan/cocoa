// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type CollectionViewDelegateFlowLayout interface {
	CollectionViewDelegate
	ImplementsCollectionView_Layout_SizeForItemAtIndexPath() bool
	// optional
	CollectionView_Layout_SizeForItemAtIndexPath(collectionView CollectionView, collectionViewLayout CollectionViewLayout, indexPath foundation.IndexPath) foundation.Size
	ImplementsCollectionView_Layout_InsetForSectionAtIndex() bool
	// optional
	CollectionView_Layout_InsetForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.EdgeInsets
	ImplementsCollectionView_Layout_MinimumLineSpacingForSectionAtIndex() bool
	// optional
	CollectionView_Layout_MinimumLineSpacingForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64
	ImplementsCollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex() bool
	// optional
	CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64
	ImplementsCollectionView_Layout_ReferenceSizeForHeaderInSection() bool
	// optional
	CollectionView_Layout_ReferenceSizeForHeaderInSection(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size
	ImplementsCollectionView_Layout_ReferenceSizeForFooterInSection() bool
	// optional
	CollectionView_Layout_ReferenceSizeForFooterInSection(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size
}

func WrapCollectionViewDelegateFlowLayout(v CollectionViewDelegateFlowLayout) objc.Object {
	return objc.WrapAsProtocol("NSCollectionViewDelegateFlowLayout", v)
}

type CollectionViewDelegateFlowLayoutBase struct {
	CollectionViewDelegateBase
}

func (p *CollectionViewDelegateFlowLayoutBase) ImplementsCollectionView_Layout_SizeForItemAtIndexPath() bool {
	return false
}

func (p *CollectionViewDelegateFlowLayoutBase) CollectionView_Layout_SizeForItemAtIndexPath(collectionView CollectionView, collectionViewLayout CollectionViewLayout, indexPath foundation.IndexPath) foundation.Size {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateFlowLayoutBase) ImplementsCollectionView_Layout_InsetForSectionAtIndex() bool {
	return false
}

func (p *CollectionViewDelegateFlowLayoutBase) CollectionView_Layout_InsetForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.EdgeInsets {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateFlowLayoutBase) ImplementsCollectionView_Layout_MinimumLineSpacingForSectionAtIndex() bool {
	return false
}

func (p *CollectionViewDelegateFlowLayoutBase) CollectionView_Layout_MinimumLineSpacingForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64 {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateFlowLayoutBase) ImplementsCollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex() bool {
	return false
}

func (p *CollectionViewDelegateFlowLayoutBase) CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64 {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateFlowLayoutBase) ImplementsCollectionView_Layout_ReferenceSizeForHeaderInSection() bool {
	return false
}

func (p *CollectionViewDelegateFlowLayoutBase) CollectionView_Layout_ReferenceSizeForHeaderInSection(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size {
	panic("unimpemented protocol method")
}

func (p *CollectionViewDelegateFlowLayoutBase) ImplementsCollectionView_Layout_ReferenceSizeForFooterInSection() bool {
	return false
}

func (p *CollectionViewDelegateFlowLayoutBase) CollectionView_Layout_ReferenceSizeForFooterInSection(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size {
	panic("unimpemented protocol method")
}

type CollectionViewDelegateFlowLayoutWrapper struct {
	objc.Object
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_Layout_SizeForItemAtIndexPath(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, indexPath foundation.IIndexPath) foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("collectionView:layout:sizeForItemAtIndexPath:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(collectionViewLayout), objc.ExtractPtr(indexPath))
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_Layout_InsetForSectionAtIndex(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](c_, objc.GetSelector("collectionView:layout:insetForSectionAtIndex:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(collectionViewLayout), section)
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_Layout_MinimumLineSpacingForSectionAtIndex(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("collectionView:layout:minimumLineSpacingForSectionAtIndex:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(collectionViewLayout), section)
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("collectionView:layout:minimumInteritemSpacingForSectionAtIndex:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(collectionViewLayout), section)
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_Layout_ReferenceSizeForHeaderInSection(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("collectionView:layout:referenceSizeForHeaderInSection:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(collectionViewLayout), section)
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_Layout_ReferenceSizeForFooterInSection(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("collectionView:layout:referenceSizeForFooterInSection:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(collectionViewLayout), section)
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_ShouldSelectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, objc.GetSelector("collectionView:shouldSelectItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths))
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_DidSelectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:didSelectItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_ShouldDeselectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, objc.GetSelector("collectionView:shouldDeselectItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths))
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_DidDeselectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:didDeselectItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(collectionView ICollectionView, indexPaths foundation.ISet, highlightState CollectionViewItemHighlightState) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, objc.GetSelector("collectionView:shouldChangeItemsAtIndexPaths:toHighlightState:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths), highlightState)
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(collectionView ICollectionView, indexPaths foundation.ISet, highlightState CollectionViewItemHighlightState) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:didChangeItemsAtIndexPaths:toHighlightState:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths), highlightState)
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(collectionView ICollectionView, item ICollectionViewItem, indexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:willDisplayItem:forRepresentedObjectAtIndexPath:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(item), objc.ExtractPtr(indexPath))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(collectionView ICollectionView, item ICollectionViewItem, indexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:didEndDisplayingItem:forRepresentedObjectAtIndexPath:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(item), objc.ExtractPtr(indexPath))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(collectionView ICollectionView, view IView, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:willDisplaySupplementaryView:forElementKind:atIndexPath:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(view), elementKind, objc.ExtractPtr(indexPath))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(collectionView ICollectionView, view IView, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:didEndDisplayingSupplementaryView:forElementOfKind:atIndexPath:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(view), elementKind, objc.ExtractPtr(indexPath))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_TransitionLayoutForOldLayout_NewLayout(collectionView ICollectionView, fromLayout ICollectionViewLayout, toLayout ICollectionViewLayout) CollectionViewTransitionLayout {
	rv := objc.CallMethod[CollectionViewTransitionLayout](c_, objc.GetSelector("collectionView:transitionLayoutForOldLayout:newLayout:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(fromLayout), objc.ExtractPtr(toLayout))
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_CanDragItemsAtIndexPaths_WithEvent(collectionView ICollectionView, indexPaths foundation.ISet, event IEvent) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("collectionView:canDragItemsAtIndexPaths:withEvent:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths), objc.ExtractPtr(event))
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_PasteboardWriterForItemAtIndexPath(collectionView ICollectionView, indexPath foundation.IIndexPath) PasteboardWritingWrapper {
	rv := objc.CallMethod[PasteboardWritingWrapper](c_, objc.GetSelector("collectionView:pasteboardWriterForItemAtIndexPath:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPath))
	return rv
}

// deprecated
func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_WriteItemsAtIndexPaths_ToPasteboard(collectionView ICollectionView, indexPaths foundation.ISet, pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("collectionView:writeItemsAtIndexPaths:toPasteboard:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths), objc.ExtractPtr(pasteboard))
	return rv
}

// deprecated
func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths(collectionView ICollectionView, dropURL foundation.IURL, indexPaths foundation.ISet) []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("collectionView:namesOfPromisedFilesDroppedAtDestination:forDraggedItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(dropURL), objc.ExtractPtr(indexPaths))
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset(collectionView ICollectionView, indexPaths foundation.ISet, event IEvent, dragImageOffset *foundation.Point) Image {
	rv := objc.CallMethod[Image](c_, objc.GetSelector("collectionView:draggingImageForItemsAtIndexPaths:withEvent:offset:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexPaths), objc.ExtractPtr(event), dragImageOffset)
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(collectionView ICollectionView, session IDraggingSession, screenPoint foundation.Point, indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexPaths:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(session), screenPoint, objc.ExtractPtr(indexPaths))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_DraggingSession_EndedAtPoint_DragOperation(collectionView ICollectionView, session IDraggingSession, screenPoint foundation.Point, operation DragOperation) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:draggingSession:endedAtPoint:dragOperation:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(session), screenPoint, operation)
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_UpdateDraggingItemsForDrag(collectionView ICollectionView, draggingInfo objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:updateDraggingItemsForDrag:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(draggingInfo))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_ValidateDrop_ProposedIndexPath_DropOperation(collectionView ICollectionView, draggingInfo objc.IObject, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	rv := objc.CallMethod[DragOperation](c_, objc.GetSelector("collectionView:validateDrop:proposedIndexPath:dropOperation:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(draggingInfo), unsafe.Pointer(proposedDropIndexPath), proposedDropOperation)
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_AcceptDrop_IndexPath_DropOperation(collectionView ICollectionView, draggingInfo objc.IObject, indexPath foundation.IIndexPath, dropOperation CollectionViewDropOperation) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("collectionView:acceptDrop:indexPath:dropOperation:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(draggingInfo), objc.ExtractPtr(indexPath), dropOperation)
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_CanDragItemsAtIndexes_WithEvent(collectionView ICollectionView, indexes foundation.IIndexSet, event IEvent) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("collectionView:canDragItemsAtIndexes:withEvent:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexes), objc.ExtractPtr(event))
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_PasteboardWriterForItemAtIndex(collectionView ICollectionView, index uint) PasteboardWritingWrapper {
	rv := objc.CallMethod[PasteboardWritingWrapper](c_, objc.GetSelector("collectionView:pasteboardWriterForItemAtIndex:"), objc.ExtractPtr(collectionView), index)
	return rv
}

// deprecated
func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_WriteItemsAtIndexes_ToPasteboard(collectionView ICollectionView, indexes foundation.IIndexSet, pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("collectionView:writeItemsAtIndexes:toPasteboard:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexes), objc.ExtractPtr(pasteboard))
	return rv
}

// deprecated
func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes(collectionView ICollectionView, dropURL foundation.IURL, indexes foundation.IIndexSet) []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("collectionView:namesOfPromisedFilesDroppedAtDestination:forDraggedItemsAtIndexes:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(dropURL), objc.ExtractPtr(indexes))
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset(collectionView ICollectionView, indexes foundation.IIndexSet, event IEvent, dragImageOffset *foundation.Point) Image {
	rv := objc.CallMethod[Image](c_, objc.GetSelector("collectionView:draggingImageForItemsAtIndexes:withEvent:offset:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(indexes), objc.ExtractPtr(event), dragImageOffset)
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(collectionView ICollectionView, session IDraggingSession, screenPoint foundation.Point, indexes foundation.IIndexSet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexes:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(session), screenPoint, objc.ExtractPtr(indexes))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_ValidateDrop_ProposedIndex_DropOperation(collectionView ICollectionView, draggingInfo objc.IObject, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	rv := objc.CallMethod[DragOperation](c_, objc.GetSelector("collectionView:validateDrop:proposedIndex:dropOperation:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(draggingInfo), proposedDropIndex, proposedDropOperation)
	return rv
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_AcceptDrop_Index_DropOperation(collectionView ICollectionView, draggingInfo objc.IObject, index int, dropOperation CollectionViewDropOperation) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("collectionView:acceptDrop:index:dropOperation:"), objc.ExtractPtr(collectionView), objc.ExtractPtr(draggingInfo), index, dropOperation)
	return rv
}
