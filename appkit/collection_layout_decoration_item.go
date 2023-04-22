// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var CollectionLayoutDecorationItemClass = _CollectionLayoutDecorationItemClass{objc.GetClass("NSCollectionLayoutDecorationItem")}

type _CollectionLayoutDecorationItemClass struct {
	objc.Class
}

type ICollectionLayoutDecorationItem interface {
	ICollectionLayoutItem
	ElementKind() string
	ZIndex() int
	SetZIndex(value int)
}

type CollectionLayoutDecorationItem struct {
	CollectionLayoutItem
}

func MakeCollectionLayoutDecorationItem(ptr unsafe.Pointer) CollectionLayoutDecorationItem {
	return CollectionLayoutDecorationItem{
		CollectionLayoutItem: MakeCollectionLayoutItem(ptr),
	}
}

func (cc _CollectionLayoutDecorationItemClass) BackgroundDecorationItemWithElementKind(elementKind string) CollectionLayoutDecorationItem {
	rv := objc.CallMethod[CollectionLayoutDecorationItem](cc, objc.GetSelector("backgroundDecorationItemWithElementKind:"), elementKind)
	return rv
}

func (cc _CollectionLayoutDecorationItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutDecorationItem {
	rv := objc.CallMethod[CollectionLayoutDecorationItem](cc, objc.GetSelector("itemWithLayoutSize:"), layoutSize)
	return rv
}

func (cc _CollectionLayoutDecorationItemClass) ItemWithLayoutSize_SupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutDecorationItem {
	rv := objc.CallMethod[CollectionLayoutDecorationItem](cc, objc.GetSelector("itemWithLayoutSize:supplementaryItems:"), layoutSize, supplementaryItems)
	return rv
}

func (cc _CollectionLayoutDecorationItemClass) Alloc() CollectionLayoutDecorationItem {
	rv := objc.CallMethod[CollectionLayoutDecorationItem](cc, objc.GetSelector("alloc"))
	return rv
}

func (cc _CollectionLayoutDecorationItemClass) New() CollectionLayoutDecorationItem {
	rv := objc.CallMethod[CollectionLayoutDecorationItem](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutDecorationItem() CollectionLayoutDecorationItem {
	return CollectionLayoutDecorationItemClass.New()
}

func (c_ CollectionLayoutDecorationItem) Init() CollectionLayoutDecorationItem {
	rv := objc.CallMethod[CollectionLayoutDecorationItem](c_, objc.GetSelector("init"))
	return rv
}

func (c_ CollectionLayoutDecorationItem) ElementKind() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("elementKind"))
	return rv
}

func (c_ CollectionLayoutDecorationItem) ZIndex() int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("zIndex"))
	return rv
}

func (c_ CollectionLayoutDecorationItem) SetZIndex(value int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setZIndex:"), value)
}
