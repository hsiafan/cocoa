// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionViewLayoutClass = _CollectionViewLayoutClass{objc.GetClass("NSCollectionViewLayout")}

type _CollectionViewLayoutClass struct {
	objc.Class
}

type ICollectionViewLayout interface {
	objc.IObject
	PrepareLayout()
	LayoutAttributesForElementsInRect(rect foundation.Rect) []CollectionViewLayoutAttributes
	LayoutAttributesForItemAtIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	LayoutAttributesForSupplementaryViewOfKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	LayoutAttributesForDecorationViewOfKind_AtIndexPath(elementKind CollectionViewDecorationElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	LayoutAttributesForDropTargetAtPoint(pointInCollectionView foundation.Point) CollectionViewLayoutAttributes
	LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	TargetContentOffsetForProposedContentOffset(proposedContentOffset foundation.Point) foundation.Point
	TargetContentOffsetForProposedContentOffset_WithScrollingVelocity(proposedContentOffset foundation.Point, velocity foundation.Point) foundation.Point
	PrepareForCollectionViewUpdates(updateItems []ICollectionViewUpdateItem)
	FinalizeCollectionViewUpdates()
	IndexPathsToInsertForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set
	IndexPathsToInsertForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set
	InitialLayoutAttributesForAppearingItemAtIndexPath(itemIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	InitialLayoutAttributesForAppearingSupplementaryElementOfKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, elementIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	InitialLayoutAttributesForAppearingDecorationElementOfKind_AtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	IndexPathsToDeleteForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set
	IndexPathsToDeleteForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set
	FinalLayoutAttributesForDisappearingItemAtIndexPath(itemIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	FinalLayoutAttributesForDisappearingSupplementaryElementOfKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, elementIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	FinalLayoutAttributesForDisappearingDecorationElementOfKind_AtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	InvalidateLayout()
	InvalidateLayoutWithContext(context ICollectionViewLayoutInvalidationContext)
	ShouldInvalidateLayoutForBoundsChange(newBounds foundation.Rect) bool
	ShouldInvalidateLayoutForPreferredLayoutAttributes_WithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) bool
	InvalidationContextForBoundsChange(newBounds foundation.Rect) CollectionViewLayoutInvalidationContext
	InvalidationContextForPreferredLayoutAttributes_WithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) CollectionViewLayoutInvalidationContext
	PrepareForAnimatedBoundsChange(oldBounds foundation.Rect)
	FinalizeAnimatedBoundsChange()
	RegisterClass_ForDecorationViewOfKind(viewClass objc.IClass, elementKind CollectionViewDecorationElementKind)
	RegisterNib_ForDecorationViewOfKind(nib INib, elementKind CollectionViewDecorationElementKind)
	PrepareForTransitionFromLayout(oldLayout ICollectionViewLayout)
	PrepareForTransitionToLayout(newLayout ICollectionViewLayout)
	FinalizeLayoutTransition()
	CollectionView() CollectionView
	CollectionViewContentSize() foundation.Size
}

type CollectionViewLayout struct {
	objc.Object
}

func MakeCollectionViewLayout(ptr unsafe.Pointer) CollectionViewLayout {
	return CollectionViewLayout{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionViewLayoutClass) Alloc() CollectionViewLayout {
	rv := objc.CallMethod[CollectionViewLayout](cc, "alloc")
	return rv
}

func (cc _CollectionViewLayoutClass) New() CollectionViewLayout {
	rv := objc.CallMethod[CollectionViewLayout](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewLayout() CollectionViewLayout {
	return CollectionViewLayoutClass.New()
}

func (c_ CollectionViewLayout) Init() CollectionViewLayout {
	rv := objc.CallMethod[CollectionViewLayout](c_, "init")
	return rv
}

func (c_ CollectionViewLayout) PrepareLayout() {
	objc.CallMethod[objc.Void](c_, "prepareLayout")
}

func (c_ CollectionViewLayout) LayoutAttributesForElementsInRect(rect foundation.Rect) []CollectionViewLayoutAttributes {
	rv := objc.CallMethod[[]CollectionViewLayoutAttributes](c_, "layoutAttributesForElementsInRect:", rect)
	return rv
}

func (c_ CollectionViewLayout) LayoutAttributesForItemAtIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, "layoutAttributesForItemAtIndexPath:", indexPath)
	return rv
}

func (c_ CollectionViewLayout) LayoutAttributesForSupplementaryViewOfKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, "layoutAttributesForSupplementaryViewOfKind:atIndexPath:", elementKind, indexPath)
	return rv
}

func (c_ CollectionViewLayout) LayoutAttributesForDecorationViewOfKind_AtIndexPath(elementKind CollectionViewDecorationElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, "layoutAttributesForDecorationViewOfKind:atIndexPath:", elementKind, indexPath)
	return rv
}

func (c_ CollectionViewLayout) LayoutAttributesForDropTargetAtPoint(pointInCollectionView foundation.Point) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, "layoutAttributesForDropTargetAtPoint:", pointInCollectionView)
	return rv
}

func (c_ CollectionViewLayout) LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, "layoutAttributesForInterItemGapBeforeIndexPath:", indexPath)
	return rv
}

func (c_ CollectionViewLayout) TargetContentOffsetForProposedContentOffset(proposedContentOffset foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](c_, "targetContentOffsetForProposedContentOffset:", proposedContentOffset)
	return rv
}

func (c_ CollectionViewLayout) TargetContentOffsetForProposedContentOffset_WithScrollingVelocity(proposedContentOffset foundation.Point, velocity foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](c_, "targetContentOffsetForProposedContentOffset:withScrollingVelocity:", proposedContentOffset, velocity)
	return rv
}

func (c_ CollectionViewLayout) PrepareForCollectionViewUpdates(updateItems []ICollectionViewUpdateItem) {
	objc.CallMethod[objc.Void](c_, "prepareForCollectionViewUpdates:", updateItems)
}

func (c_ CollectionViewLayout) FinalizeCollectionViewUpdates() {
	objc.CallMethod[objc.Void](c_, "finalizeCollectionViewUpdates")
}

func (c_ CollectionViewLayout) IndexPathsToInsertForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, "indexPathsToInsertForSupplementaryViewOfKind:", elementKind)
	return rv
}

func (c_ CollectionViewLayout) IndexPathsToInsertForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, "indexPathsToInsertForDecorationViewOfKind:", elementKind)
	return rv
}

func (c_ CollectionViewLayout) InitialLayoutAttributesForAppearingItemAtIndexPath(itemIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, "initialLayoutAttributesForAppearingItemAtIndexPath:", itemIndexPath)
	return rv
}

func (c_ CollectionViewLayout) InitialLayoutAttributesForAppearingSupplementaryElementOfKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, elementIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, "initialLayoutAttributesForAppearingSupplementaryElementOfKind:atIndexPath:", elementKind, elementIndexPath)
	return rv
}

func (c_ CollectionViewLayout) InitialLayoutAttributesForAppearingDecorationElementOfKind_AtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, "initialLayoutAttributesForAppearingDecorationElementOfKind:atIndexPath:", elementKind, decorationIndexPath)
	return rv
}

func (c_ CollectionViewLayout) IndexPathsToDeleteForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, "indexPathsToDeleteForSupplementaryViewOfKind:", elementKind)
	return rv
}

func (c_ CollectionViewLayout) IndexPathsToDeleteForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, "indexPathsToDeleteForDecorationViewOfKind:", elementKind)
	return rv
}

func (c_ CollectionViewLayout) FinalLayoutAttributesForDisappearingItemAtIndexPath(itemIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, "finalLayoutAttributesForDisappearingItemAtIndexPath:", itemIndexPath)
	return rv
}

func (c_ CollectionViewLayout) FinalLayoutAttributesForDisappearingSupplementaryElementOfKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, elementIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, "finalLayoutAttributesForDisappearingSupplementaryElementOfKind:atIndexPath:", elementKind, elementIndexPath)
	return rv
}

func (c_ CollectionViewLayout) FinalLayoutAttributesForDisappearingDecorationElementOfKind_AtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, "finalLayoutAttributesForDisappearingDecorationElementOfKind:atIndexPath:", elementKind, decorationIndexPath)
	return rv
}

func (c_ CollectionViewLayout) InvalidateLayout() {
	objc.CallMethod[objc.Void](c_, "invalidateLayout")
}

func (c_ CollectionViewLayout) InvalidateLayoutWithContext(context ICollectionViewLayoutInvalidationContext) {
	objc.CallMethod[objc.Void](c_, "invalidateLayoutWithContext:", context)
}

func (c_ CollectionViewLayout) ShouldInvalidateLayoutForBoundsChange(newBounds foundation.Rect) bool {
	rv := objc.CallMethod[bool](c_, "shouldInvalidateLayoutForBoundsChange:", newBounds)
	return rv
}

func (c_ CollectionViewLayout) ShouldInvalidateLayoutForPreferredLayoutAttributes_WithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) bool {
	rv := objc.CallMethod[bool](c_, "shouldInvalidateLayoutForPreferredLayoutAttributes:withOriginalAttributes:", preferredAttributes, originalAttributes)
	return rv
}

func (c_ CollectionViewLayout) InvalidationContextForBoundsChange(newBounds foundation.Rect) CollectionViewLayoutInvalidationContext {
	rv := objc.CallMethod[CollectionViewLayoutInvalidationContext](c_, "invalidationContextForBoundsChange:", newBounds)
	return rv
}

func (c_ CollectionViewLayout) InvalidationContextForPreferredLayoutAttributes_WithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) CollectionViewLayoutInvalidationContext {
	rv := objc.CallMethod[CollectionViewLayoutInvalidationContext](c_, "invalidationContextForPreferredLayoutAttributes:withOriginalAttributes:", preferredAttributes, originalAttributes)
	return rv
}

func (c_ CollectionViewLayout) PrepareForAnimatedBoundsChange(oldBounds foundation.Rect) {
	objc.CallMethod[objc.Void](c_, "prepareForAnimatedBoundsChange:", oldBounds)
}

func (c_ CollectionViewLayout) FinalizeAnimatedBoundsChange() {
	objc.CallMethod[objc.Void](c_, "finalizeAnimatedBoundsChange")
}

func (c_ CollectionViewLayout) RegisterClass_ForDecorationViewOfKind(viewClass objc.IClass, elementKind CollectionViewDecorationElementKind) {
	objc.CallMethod[objc.Void](c_, "registerClass:forDecorationViewOfKind:", viewClass, elementKind)
}

func (c_ CollectionViewLayout) RegisterNib_ForDecorationViewOfKind(nib INib, elementKind CollectionViewDecorationElementKind) {
	objc.CallMethod[objc.Void](c_, "registerNib:forDecorationViewOfKind:", nib, elementKind)
}

func (c_ CollectionViewLayout) PrepareForTransitionFromLayout(oldLayout ICollectionViewLayout) {
	objc.CallMethod[objc.Void](c_, "prepareForTransitionFromLayout:", oldLayout)
}

func (c_ CollectionViewLayout) PrepareForTransitionToLayout(newLayout ICollectionViewLayout) {
	objc.CallMethod[objc.Void](c_, "prepareForTransitionToLayout:", newLayout)
}

func (c_ CollectionViewLayout) FinalizeLayoutTransition() {
	objc.CallMethod[objc.Void](c_, "finalizeLayoutTransition")
}

func (c_ CollectionViewLayout) CollectionView() CollectionView {
	rv := objc.CallMethod[CollectionView](c_, "collectionView")
	return rv
}

func (cc _CollectionViewLayoutClass) LayoutAttributesClass() objc.Class {
	rv := objc.CallMethod[objc.Class](cc, "layoutAttributesClass")
	return rv
}

func (c_ CollectionViewLayout) CollectionViewContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, "collectionViewContentSize")
	return rv
}

func (cc _CollectionViewLayoutClass) InvalidationContextClass() objc.Class {
	rv := objc.CallMethod[objc.Class](cc, "invalidationContextClass")
	return rv
}
