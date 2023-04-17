// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionLayoutSupplementaryItemClass = _CollectionLayoutSupplementaryItemClass{objc.GetClass("NSCollectionLayoutSupplementaryItem")}

type _CollectionLayoutSupplementaryItemClass struct {
	objc.Class
}

type ICollectionLayoutSupplementaryItem interface {
	ICollectionLayoutItem
	ItemAnchor() CollectionLayoutAnchor
	ContainerAnchor() CollectionLayoutAnchor
	ElementKind() string
	ZIndex() int
	SetZIndex(value int)
}

type CollectionLayoutSupplementaryItem struct {
	CollectionLayoutItem
}

func MakeCollectionLayoutSupplementaryItem(ptr unsafe.Pointer) CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItem{
		CollectionLayoutItem: MakeCollectionLayoutItem(ptr),
	}
}

func (cc _CollectionLayoutSupplementaryItemClass) SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor) CollectionLayoutSupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutSupplementaryItem](cc, "supplementaryItemWithLayoutSize:elementKind:containerAnchor:", layoutSize, elementKind, containerAnchor)
	return rv
}

func (cc _CollectionLayoutSupplementaryItemClass) SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor, itemAnchor ICollectionLayoutAnchor) CollectionLayoutSupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutSupplementaryItem](cc, "supplementaryItemWithLayoutSize:elementKind:containerAnchor:itemAnchor:", layoutSize, elementKind, containerAnchor, itemAnchor)
	return rv
}

func (cc _CollectionLayoutSupplementaryItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutSupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutSupplementaryItem](cc, "itemWithLayoutSize:", layoutSize)
	return rv
}

func (cc _CollectionLayoutSupplementaryItemClass) ItemWithLayoutSize_SupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutSupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutSupplementaryItem](cc, "itemWithLayoutSize:supplementaryItems:", layoutSize, supplementaryItems)
	return rv
}

func (cc _CollectionLayoutSupplementaryItemClass) Alloc() CollectionLayoutSupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutSupplementaryItem](cc, "alloc")
	return rv
}

func (cc _CollectionLayoutSupplementaryItemClass) New() CollectionLayoutSupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutSupplementaryItem](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutSupplementaryItem() CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItemClass.New()
}

func (c_ CollectionLayoutSupplementaryItem) Init() CollectionLayoutSupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutSupplementaryItem](c_, "init")
	return rv
}

func (c_ CollectionLayoutSupplementaryItem) ItemAnchor() CollectionLayoutAnchor {
	rv := ffi.CallMethod[CollectionLayoutAnchor](c_, "itemAnchor")
	return rv
}

func (c_ CollectionLayoutSupplementaryItem) ContainerAnchor() CollectionLayoutAnchor {
	rv := ffi.CallMethod[CollectionLayoutAnchor](c_, "containerAnchor")
	return rv
}

func (c_ CollectionLayoutSupplementaryItem) ElementKind() string {
	rv := ffi.CallMethod[string](c_, "elementKind")
	return rv
}

func (c_ CollectionLayoutSupplementaryItem) ZIndex() int {
	rv := ffi.CallMethod[int](c_, "zIndex")
	return rv
}

func (c_ CollectionLayoutSupplementaryItem) SetZIndex(value int) {
	ffi.CallMethod[ffi.Void](c_, "setZIndex:", value)
}
