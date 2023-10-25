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

type CollectionViewDelegateFlowLayoutCreator struct {
	className string
	class     objc.Class
}

func NewCollectionViewDelegateFlowLayoutCreator(name string) *CollectionViewDelegateFlowLayoutCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &CollectionViewDelegateFlowLayoutCreator{className: name, class: class}
}

func (c *CollectionViewDelegateFlowLayoutCreator) SetCollectionView_Layout_SizeForItemAtIndexPath(handle func(o objc.Object, collectionView CollectionView, collectionViewLayout CollectionViewLayout, indexPath foundation.IndexPath) foundation.Size) {
	objc.AddMethod(c.class, objc.SEL("collectionView:layout:sizeForItemAtIndexPath:"), handle)
}

func (c *CollectionViewDelegateFlowLayoutCreator) SetCollectionView_Layout_InsetForSectionAtIndex(handle func(o objc.Object, collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.EdgeInsets) {
	objc.AddMethod(c.class, objc.SEL("collectionView:layout:insetForSectionAtIndex:"), handle)
}

func (c *CollectionViewDelegateFlowLayoutCreator) SetCollectionView_Layout_MinimumLineSpacingForSectionAtIndex(handle func(o objc.Object, collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64) {
	objc.AddMethod(c.class, objc.SEL("collectionView:layout:minimumLineSpacingForSectionAtIndex:"), handle)
}

func (c *CollectionViewDelegateFlowLayoutCreator) SetCollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex(handle func(o objc.Object, collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64) {
	objc.AddMethod(c.class, objc.SEL("collectionView:layout:minimumInteritemSpacingForSectionAtIndex:"), handle)
}

func (c *CollectionViewDelegateFlowLayoutCreator) SetCollectionView_Layout_ReferenceSizeForHeaderInSection(handle func(o objc.Object, collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size) {
	objc.AddMethod(c.class, objc.SEL("collectionView:layout:referenceSizeForHeaderInSection:"), handle)
}

func (c *CollectionViewDelegateFlowLayoutCreator) SetCollectionView_Layout_ReferenceSizeForFooterInSection(handle func(o objc.Object, collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size) {
	objc.AddMethod(c.class, objc.SEL("collectionView:layout:referenceSizeForFooterInSection:"), handle)
}

func (c *CollectionViewDelegateFlowLayoutCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
