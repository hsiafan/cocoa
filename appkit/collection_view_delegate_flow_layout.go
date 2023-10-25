// auto-generated code, do not modify
package appkit

import (
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
