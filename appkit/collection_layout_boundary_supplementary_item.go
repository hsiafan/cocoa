// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionLayoutBoundarySupplementaryItemClass = _CollectionLayoutBoundarySupplementaryItemClass{objc.GetClass("NSCollectionLayoutBoundarySupplementaryItem")}

type _CollectionLayoutBoundarySupplementaryItemClass struct {
	objc.Class
}

type ICollectionLayoutBoundarySupplementaryItem interface {
	ICollectionLayoutSupplementaryItem
	PinToVisibleBounds() bool
	SetPinToVisibleBounds(value bool)
	Offset() foundation.Point
	Alignment() RectAlignment
	ExtendsBoundary() bool
	SetExtendsBoundary(value bool)
}

type CollectionLayoutBoundarySupplementaryItem struct {
	CollectionLayoutSupplementaryItem
}

func MakeCollectionLayoutBoundarySupplementaryItem(ptr unsafe.Pointer) CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItem{
		CollectionLayoutSupplementaryItem: MakeCollectionLayoutSupplementaryItem(ptr),
	}
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment(layoutSize ICollectionLayoutSize, elementKind string, alignment RectAlignment) CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, objc.SEL("boundarySupplementaryItemWithLayoutSize:elementKind:alignment:"), objc.ExtractPtr(layoutSize), elementKind, alignment)
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment_AbsoluteOffset(layoutSize ICollectionLayoutSize, elementKind string, alignment RectAlignment, absoluteOffset foundation.Point) CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, objc.SEL("boundarySupplementaryItemWithLayoutSize:elementKind:alignment:absoluteOffset:"), objc.ExtractPtr(layoutSize), elementKind, alignment, absoluteOffset)
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor) CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, objc.SEL("supplementaryItemWithLayoutSize:elementKind:containerAnchor:"), objc.ExtractPtr(layoutSize), elementKind, objc.ExtractPtr(containerAnchor))
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor, itemAnchor ICollectionLayoutAnchor) CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, objc.SEL("supplementaryItemWithLayoutSize:elementKind:containerAnchor:itemAnchor:"), objc.ExtractPtr(layoutSize), elementKind, objc.ExtractPtr(containerAnchor), objc.ExtractPtr(itemAnchor))
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, objc.SEL("itemWithLayoutSize:"), objc.ExtractPtr(layoutSize))
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) ItemWithLayoutSize_SupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, objc.SEL("itemWithLayoutSize:supplementaryItems:"), objc.ExtractPtr(layoutSize), supplementaryItems)
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) Alloc() CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, objc.SEL("alloc"))
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) New() CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutBoundarySupplementaryItem() CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItemClass.New()
}

func (c_ CollectionLayoutBoundarySupplementaryItem) Init() CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutBoundarySupplementaryItem](c_, objc.SEL("init"))
	return rv
}

func (c_ CollectionLayoutBoundarySupplementaryItem) PinToVisibleBounds() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("pinToVisibleBounds"))
	return rv
}

func (c_ CollectionLayoutBoundarySupplementaryItem) SetPinToVisibleBounds(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setPinToVisibleBounds:"), value)
}

func (c_ CollectionLayoutBoundarySupplementaryItem) Offset() foundation.Point {
	rv := objc.CallMethod[foundation.Point](c_, objc.SEL("offset"))
	return rv
}

func (c_ CollectionLayoutBoundarySupplementaryItem) Alignment() RectAlignment {
	rv := objc.CallMethod[RectAlignment](c_, objc.SEL("alignment"))
	return rv
}

func (c_ CollectionLayoutBoundarySupplementaryItem) ExtendsBoundary() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("extendsBoundary"))
	return rv
}

func (c_ CollectionLayoutBoundarySupplementaryItem) SetExtendsBoundary(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setExtendsBoundary:"), value)
}
