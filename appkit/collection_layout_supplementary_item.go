// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[CollectionLayoutSupplementaryItem](cc, objc.SEL("supplementaryItemWithLayoutSize:elementKind:containerAnchor:"), objc.ExtractPtr(layoutSize), elementKind, objc.ExtractPtr(containerAnchor))
	return rv
}

func (cc _CollectionLayoutSupplementaryItemClass) SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor, itemAnchor ICollectionLayoutAnchor) CollectionLayoutSupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutSupplementaryItem](cc, objc.SEL("supplementaryItemWithLayoutSize:elementKind:containerAnchor:itemAnchor:"), objc.ExtractPtr(layoutSize), elementKind, objc.ExtractPtr(containerAnchor), objc.ExtractPtr(itemAnchor))
	return rv
}

func (cc _CollectionLayoutSupplementaryItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutSupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutSupplementaryItem](cc, objc.SEL("itemWithLayoutSize:"), objc.ExtractPtr(layoutSize))
	return rv
}

func (cc _CollectionLayoutSupplementaryItemClass) ItemWithLayoutSize_SupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutSupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutSupplementaryItem](cc, objc.SEL("itemWithLayoutSize:supplementaryItems:"), objc.ExtractPtr(layoutSize), supplementaryItems)
	return rv
}

func (cc _CollectionLayoutSupplementaryItemClass) Alloc() CollectionLayoutSupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutSupplementaryItem](cc, objc.SEL("alloc"))
	return rv
}

func (cc _CollectionLayoutSupplementaryItemClass) New() CollectionLayoutSupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutSupplementaryItem](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutSupplementaryItem() CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItemClass.New()
}

func (c_ CollectionLayoutSupplementaryItem) Init() CollectionLayoutSupplementaryItem {
	rv := objc.CallMethod[CollectionLayoutSupplementaryItem](c_, objc.SEL("init"))
	return rv
}

func (c_ CollectionLayoutSupplementaryItem) ItemAnchor() CollectionLayoutAnchor {
	rv := objc.CallMethod[CollectionLayoutAnchor](c_, objc.SEL("itemAnchor"))
	return rv
}

func (c_ CollectionLayoutSupplementaryItem) ContainerAnchor() CollectionLayoutAnchor {
	rv := objc.CallMethod[CollectionLayoutAnchor](c_, objc.SEL("containerAnchor"))
	return rv
}

func (c_ CollectionLayoutSupplementaryItem) ElementKind() string {
	rv := objc.CallMethod[string](c_, objc.SEL("elementKind"))
	return rv
}

func (c_ CollectionLayoutSupplementaryItem) ZIndex() int {
	rv := objc.CallMethod[int](c_, objc.SEL("zIndex"))
	return rv
}

func (c_ CollectionLayoutSupplementaryItem) SetZIndex(value int) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setZIndex:"), value)
}
