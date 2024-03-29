// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[CollectionLayoutItem](cc, objc.SEL("itemWithLayoutSize:"), objc.ExtractPtr(layoutSize))
	return rv
}

func (cc _CollectionLayoutItemClass) ItemWithLayoutSize_SupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutItem {
	rv := objc.CallMethod[CollectionLayoutItem](cc, objc.SEL("itemWithLayoutSize:supplementaryItems:"), objc.ExtractPtr(layoutSize), supplementaryItems)
	return rv
}

func (cc _CollectionLayoutItemClass) Alloc() CollectionLayoutItem {
	rv := objc.CallMethod[CollectionLayoutItem](cc, objc.SEL("alloc"))
	return rv
}

func (cc _CollectionLayoutItemClass) New() CollectionLayoutItem {
	rv := objc.CallMethod[CollectionLayoutItem](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutItem() CollectionLayoutItem {
	return CollectionLayoutItemClass.New()
}

func (c_ CollectionLayoutItem) Init() CollectionLayoutItem {
	rv := objc.CallMethod[CollectionLayoutItem](c_, objc.SEL("init"))
	return rv
}

func (c_ CollectionLayoutItem) LayoutSize() CollectionLayoutSize {
	rv := objc.CallMethod[CollectionLayoutSize](c_, objc.SEL("layoutSize"))
	return rv
}

func (c_ CollectionLayoutItem) SupplementaryItems() []CollectionLayoutSupplementaryItem {
	rv := objc.CallMethod[[]CollectionLayoutSupplementaryItem](c_, objc.SEL("supplementaryItems"))
	return rv
}

func (c_ CollectionLayoutItem) EdgeSpacing() CollectionLayoutEdgeSpacing {
	rv := objc.CallMethod[CollectionLayoutEdgeSpacing](c_, objc.SEL("edgeSpacing"))
	return rv
}

func (c_ CollectionLayoutItem) SetEdgeSpacing(value ICollectionLayoutEdgeSpacing) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setEdgeSpacing:"), objc.ExtractPtr(value))
}

func (c_ CollectionLayoutItem) ContentInsets() DirectionalEdgeInsets {
	rv := objc.CallMethod[DirectionalEdgeInsets](c_, objc.SEL("contentInsets"))
	return rv
}

func (c_ CollectionLayoutItem) SetContentInsets(value DirectionalEdgeInsets) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setContentInsets:"), value)
}
