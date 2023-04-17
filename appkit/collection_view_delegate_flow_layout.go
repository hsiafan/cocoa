// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
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

type CollectionViewDelegateFlowLayoutImpl struct {
	CollectionViewDelegateImpl
	_CollectionView_Layout_SizeForItemAtIndexPath                   func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, indexPath foundation.IndexPath) foundation.Size
	_CollectionView_Layout_InsetForSectionAtIndex                   func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.EdgeInsets
	_CollectionView_Layout_MinimumLineSpacingForSectionAtIndex      func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64
	_CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64
	_CollectionView_Layout_ReferenceSizeForHeaderInSection          func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size
	_CollectionView_Layout_ReferenceSizeForFooterInSection          func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size
}

func (di *CollectionViewDelegateFlowLayoutImpl) ImplementsCollectionView_Layout_SizeForItemAtIndexPath() bool {
	return di._CollectionView_Layout_SizeForItemAtIndexPath != nil
}

func (di *CollectionViewDelegateFlowLayoutImpl) SetCollectionView_Layout_SizeForItemAtIndexPath(f func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, indexPath foundation.IndexPath) foundation.Size) {
	di._CollectionView_Layout_SizeForItemAtIndexPath = f
}

func (di *CollectionViewDelegateFlowLayoutImpl) CollectionView_Layout_SizeForItemAtIndexPath(collectionView CollectionView, collectionViewLayout CollectionViewLayout, indexPath foundation.IndexPath) foundation.Size {
	return di._CollectionView_Layout_SizeForItemAtIndexPath(collectionView, collectionViewLayout, indexPath)
}
func (di *CollectionViewDelegateFlowLayoutImpl) ImplementsCollectionView_Layout_InsetForSectionAtIndex() bool {
	return di._CollectionView_Layout_InsetForSectionAtIndex != nil
}

func (di *CollectionViewDelegateFlowLayoutImpl) SetCollectionView_Layout_InsetForSectionAtIndex(f func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.EdgeInsets) {
	di._CollectionView_Layout_InsetForSectionAtIndex = f
}

func (di *CollectionViewDelegateFlowLayoutImpl) CollectionView_Layout_InsetForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.EdgeInsets {
	return di._CollectionView_Layout_InsetForSectionAtIndex(collectionView, collectionViewLayout, section)
}
func (di *CollectionViewDelegateFlowLayoutImpl) ImplementsCollectionView_Layout_MinimumLineSpacingForSectionAtIndex() bool {
	return di._CollectionView_Layout_MinimumLineSpacingForSectionAtIndex != nil
}

func (di *CollectionViewDelegateFlowLayoutImpl) SetCollectionView_Layout_MinimumLineSpacingForSectionAtIndex(f func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64) {
	di._CollectionView_Layout_MinimumLineSpacingForSectionAtIndex = f
}

func (di *CollectionViewDelegateFlowLayoutImpl) CollectionView_Layout_MinimumLineSpacingForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64 {
	return di._CollectionView_Layout_MinimumLineSpacingForSectionAtIndex(collectionView, collectionViewLayout, section)
}
func (di *CollectionViewDelegateFlowLayoutImpl) ImplementsCollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex() bool {
	return di._CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex != nil
}

func (di *CollectionViewDelegateFlowLayoutImpl) SetCollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex(f func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64) {
	di._CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex = f
}

func (di *CollectionViewDelegateFlowLayoutImpl) CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64 {
	return di._CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex(collectionView, collectionViewLayout, section)
}
func (di *CollectionViewDelegateFlowLayoutImpl) ImplementsCollectionView_Layout_ReferenceSizeForHeaderInSection() bool {
	return di._CollectionView_Layout_ReferenceSizeForHeaderInSection != nil
}

func (di *CollectionViewDelegateFlowLayoutImpl) SetCollectionView_Layout_ReferenceSizeForHeaderInSection(f func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size) {
	di._CollectionView_Layout_ReferenceSizeForHeaderInSection = f
}

func (di *CollectionViewDelegateFlowLayoutImpl) CollectionView_Layout_ReferenceSizeForHeaderInSection(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size {
	return di._CollectionView_Layout_ReferenceSizeForHeaderInSection(collectionView, collectionViewLayout, section)
}
func (di *CollectionViewDelegateFlowLayoutImpl) ImplementsCollectionView_Layout_ReferenceSizeForFooterInSection() bool {
	return di._CollectionView_Layout_ReferenceSizeForFooterInSection != nil
}

func (di *CollectionViewDelegateFlowLayoutImpl) SetCollectionView_Layout_ReferenceSizeForFooterInSection(f func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size) {
	di._CollectionView_Layout_ReferenceSizeForFooterInSection = f
}

func (di *CollectionViewDelegateFlowLayoutImpl) CollectionView_Layout_ReferenceSizeForFooterInSection(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size {
	return di._CollectionView_Layout_ReferenceSizeForFooterInSection(collectionView, collectionViewLayout, section)
}

type CollectionViewDelegateFlowLayoutWrapper struct {
	CollectionViewDelegateWrapper
}

func (c_ *CollectionViewDelegateFlowLayoutWrapper) ImplementsCollectionView_Layout_SizeForItemAtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:layout:sizeForItemAtIndexPath:"))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_Layout_SizeForItemAtIndexPath(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, indexPath foundation.IIndexPath) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "collectionView:layout:sizeForItemAtIndexPath:", collectionView, collectionViewLayout, indexPath)
	return rv
}

func (c_ *CollectionViewDelegateFlowLayoutWrapper) ImplementsCollectionView_Layout_InsetForSectionAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:layout:insetForSectionAtIndex:"))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_Layout_InsetForSectionAtIndex(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) foundation.EdgeInsets {
	rv := ffi.CallMethod[foundation.EdgeInsets](c_, "collectionView:layout:insetForSectionAtIndex:", collectionView, collectionViewLayout, section)
	return rv
}

func (c_ *CollectionViewDelegateFlowLayoutWrapper) ImplementsCollectionView_Layout_MinimumLineSpacingForSectionAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:layout:minimumLineSpacingForSectionAtIndex:"))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_Layout_MinimumLineSpacingForSectionAtIndex(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) float64 {
	rv := ffi.CallMethod[float64](c_, "collectionView:layout:minimumLineSpacingForSectionAtIndex:", collectionView, collectionViewLayout, section)
	return rv
}

func (c_ *CollectionViewDelegateFlowLayoutWrapper) ImplementsCollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:layout:minimumInteritemSpacingForSectionAtIndex:"))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) float64 {
	rv := ffi.CallMethod[float64](c_, "collectionView:layout:minimumInteritemSpacingForSectionAtIndex:", collectionView, collectionViewLayout, section)
	return rv
}

func (c_ *CollectionViewDelegateFlowLayoutWrapper) ImplementsCollectionView_Layout_ReferenceSizeForHeaderInSection() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:layout:referenceSizeForHeaderInSection:"))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_Layout_ReferenceSizeForHeaderInSection(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "collectionView:layout:referenceSizeForHeaderInSection:", collectionView, collectionViewLayout, section)
	return rv
}

func (c_ *CollectionViewDelegateFlowLayoutWrapper) ImplementsCollectionView_Layout_ReferenceSizeForFooterInSection() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:layout:referenceSizeForFooterInSection:"))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_Layout_ReferenceSizeForFooterInSection(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "collectionView:layout:referenceSizeForFooterInSection:", collectionView, collectionViewLayout, section)
	return rv
}
