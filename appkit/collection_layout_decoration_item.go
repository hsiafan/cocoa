// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[CollectionLayoutDecorationItem](cc, "backgroundDecorationItemWithElementKind:", elementKind)
	return rv
}

func (cc _CollectionLayoutDecorationItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutDecorationItem {
	rv := ffi.CallMethod[CollectionLayoutDecorationItem](cc, "itemWithLayoutSize:", layoutSize)
	return rv
}

func (cc _CollectionLayoutDecorationItemClass) ItemWithLayoutSize_SupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutDecorationItem {
	rv := ffi.CallMethod[CollectionLayoutDecorationItem](cc, "itemWithLayoutSize:supplementaryItems:", layoutSize, supplementaryItems)
	return rv
}

func (cc _CollectionLayoutDecorationItemClass) Alloc() CollectionLayoutDecorationItem {
	rv := ffi.CallMethod[CollectionLayoutDecorationItem](cc, "alloc")
	return rv
}

func (cc _CollectionLayoutDecorationItemClass) New() CollectionLayoutDecorationItem {
	rv := ffi.CallMethod[CollectionLayoutDecorationItem](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutDecorationItem() CollectionLayoutDecorationItem {
	return CollectionLayoutDecorationItemClass.New()
}

func (c_ CollectionLayoutDecorationItem) Init() CollectionLayoutDecorationItem {
	rv := ffi.CallMethod[CollectionLayoutDecorationItem](c_, "init")
	return rv
}

func (c_ CollectionLayoutDecorationItem) ElementKind() string {
	rv := ffi.CallMethod[string](c_, "elementKind")
	return rv
}

func (c_ CollectionLayoutDecorationItem) ZIndex() int {
	rv := ffi.CallMethod[int](c_, "zIndex")
	return rv
}

func (c_ CollectionLayoutDecorationItem) SetZIndex(value int) {
	ffi.CallMethod[ffi.Void](c_, "setZIndex:", value)
}
