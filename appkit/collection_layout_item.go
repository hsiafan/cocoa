// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionLayoutItemClass = _CollectionLayoutItemClass{objc.GetClass("NSCollectionLayoutItem")}

type _CollectionLayoutItemClass struct {
	objc.Class
}

type ICollectionLayoutItem interface {
	objc.IObject
	LayoutSize() CollectionLayoutSize
	SupplementaryItems() []CollectionLayoutSupplementaryItem
	EdgeSpacing() CollectionLayoutEdgeSpacing
	SetEdgeSpacing(value ICollectionLayoutEdgeSpacing)
	ContentInsets() DirectionalEdgeInsets
	SetContentInsets(value DirectionalEdgeInsets)
}

type CollectionLayoutItem struct {
	objc.Object
}

func MakeCollectionLayoutItem(ptr unsafe.Pointer) CollectionLayoutItem {
	return CollectionLayoutItem{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutItem {
	rv := ffi.CallMethod[CollectionLayoutItem](cc, "itemWithLayoutSize:", layoutSize)
	return rv
}

func (cc _CollectionLayoutItemClass) ItemWithLayoutSize_SupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutItem {
	rv := ffi.CallMethod[CollectionLayoutItem](cc, "itemWithLayoutSize:supplementaryItems:", layoutSize, supplementaryItems)
	return rv
}

func (cc _CollectionLayoutItemClass) Alloc() CollectionLayoutItem {
	rv := ffi.CallMethod[CollectionLayoutItem](cc, "alloc")
	return rv
}

func (cc _CollectionLayoutItemClass) New() CollectionLayoutItem {
	rv := ffi.CallMethod[CollectionLayoutItem](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutItem() CollectionLayoutItem {
	return CollectionLayoutItemClass.New()
}

func (c_ CollectionLayoutItem) Init() CollectionLayoutItem {
	rv := ffi.CallMethod[CollectionLayoutItem](c_, "init")
	return rv
}

func (c_ CollectionLayoutItem) LayoutSize() CollectionLayoutSize {
	rv := ffi.CallMethod[CollectionLayoutSize](c_, "layoutSize")
	return rv
}

func (c_ CollectionLayoutItem) SupplementaryItems() []CollectionLayoutSupplementaryItem {
	rv := ffi.CallMethod[[]CollectionLayoutSupplementaryItem](c_, "supplementaryItems")
	return rv
}

func (c_ CollectionLayoutItem) EdgeSpacing() CollectionLayoutEdgeSpacing {
	rv := ffi.CallMethod[CollectionLayoutEdgeSpacing](c_, "edgeSpacing")
	return rv
}

func (c_ CollectionLayoutItem) SetEdgeSpacing(value ICollectionLayoutEdgeSpacing) {
	ffi.CallMethod[ffi.Void](c_, "setEdgeSpacing:", value)
}

func (c_ CollectionLayoutItem) ContentInsets() DirectionalEdgeInsets {
	rv := ffi.CallMethod[DirectionalEdgeInsets](c_, "contentInsets")
	return rv
}

func (c_ CollectionLayoutItem) SetContentInsets(value DirectionalEdgeInsets) {
	ffi.CallMethod[ffi.Void](c_, "setContentInsets:", value)
}