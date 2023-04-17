// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, "boundarySupplementaryItemWithLayoutSize:elementKind:alignment:", layoutSize, elementKind, alignment)
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment_AbsoluteOffset(layoutSize ICollectionLayoutSize, elementKind string, alignment RectAlignment, absoluteOffset foundation.Point) CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, "boundarySupplementaryItemWithLayoutSize:elementKind:alignment:absoluteOffset:", layoutSize, elementKind, alignment, absoluteOffset)
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor) CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, "supplementaryItemWithLayoutSize:elementKind:containerAnchor:", layoutSize, elementKind, containerAnchor)
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor, itemAnchor ICollectionLayoutAnchor) CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, "supplementaryItemWithLayoutSize:elementKind:containerAnchor:itemAnchor:", layoutSize, elementKind, containerAnchor, itemAnchor)
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, "itemWithLayoutSize:", layoutSize)
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) ItemWithLayoutSize_SupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, "itemWithLayoutSize:supplementaryItems:", layoutSize, supplementaryItems)
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) Alloc() CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, "alloc")
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) New() CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutBoundarySupplementaryItem() CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItemClass.New()
}

func (c_ CollectionLayoutBoundarySupplementaryItem) Init() CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutBoundarySupplementaryItem](c_, "init")
	return rv
}

func (c_ CollectionLayoutBoundarySupplementaryItem) PinToVisibleBounds() bool {
	rv := ffi.CallMethod[bool](c_, "pinToVisibleBounds")
	return rv
}

func (c_ CollectionLayoutBoundarySupplementaryItem) SetPinToVisibleBounds(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setPinToVisibleBounds:", value)
}

func (c_ CollectionLayoutBoundarySupplementaryItem) Offset() foundation.Point {
	rv := ffi.CallMethod[foundation.Point](c_, "offset")
	return rv
}

func (c_ CollectionLayoutBoundarySupplementaryItem) Alignment() RectAlignment {
	rv := ffi.CallMethod[RectAlignment](c_, "alignment")
	return rv
}

func (c_ CollectionLayoutBoundarySupplementaryItem) ExtendsBoundary() bool {
	rv := ffi.CallMethod[bool](c_, "extendsBoundary")
	return rv
}

func (c_ CollectionLayoutBoundarySupplementaryItem) SetExtendsBoundary(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setExtendsBoundary:", value)
}
