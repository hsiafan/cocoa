// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionLayoutGroupClass = _CollectionLayoutGroupClass{objc.GetClass("NSCollectionLayoutGroup")}

type _CollectionLayoutGroupClass struct {
	objc.Class
}

type ICollectionLayoutGroup interface {
	ICollectionLayoutItem
	VisualDescription() string
	Subitems() []CollectionLayoutItem
	SetSupplementaryItems(value []ICollectionLayoutSupplementaryItem)
	InterItemSpacing() CollectionLayoutSpacing
	SetInterItemSpacing(value ICollectionLayoutSpacing)
}

type CollectionLayoutGroup struct {
	CollectionLayoutItem
}

func MakeCollectionLayoutGroup(ptr unsafe.Pointer) CollectionLayoutGroup {
	return CollectionLayoutGroup{
		CollectionLayoutItem: MakeCollectionLayoutItem(ptr),
	}
}

func (cc _CollectionLayoutGroupClass) HorizontalGroupWithLayoutSize_Subitems(layoutSize ICollectionLayoutSize, subitems []ICollectionLayoutItem) CollectionLayoutGroup {
	rv := ffi.CallMethod[CollectionLayoutGroup](cc, "horizontalGroupWithLayoutSize:subitems:", layoutSize, subitems)
	return rv
}

func (cc _CollectionLayoutGroupClass) VerticalGroupWithLayoutSize_Subitems(layoutSize ICollectionLayoutSize, subitems []ICollectionLayoutItem) CollectionLayoutGroup {
	rv := ffi.CallMethod[CollectionLayoutGroup](cc, "verticalGroupWithLayoutSize:subitems:", layoutSize, subitems)
	return rv
}

func (cc _CollectionLayoutGroupClass) CustomGroupWithLayoutSize_ItemProvider(layoutSize ICollectionLayoutSize, itemProvider func(layoutEnvironment CollectionLayoutEnvironmentWrapper) []ICollectionLayoutGroupCustomItem) CollectionLayoutGroup {
	rv := ffi.CallMethod[CollectionLayoutGroup](cc, "customGroupWithLayoutSize:itemProvider:", layoutSize, itemProvider)
	return rv
}

func (cc _CollectionLayoutGroupClass) HorizontalGroupWithLayoutSize_Subitem_Count(layoutSize ICollectionLayoutSize, subitem ICollectionLayoutItem, count int) CollectionLayoutGroup {
	rv := ffi.CallMethod[CollectionLayoutGroup](cc, "horizontalGroupWithLayoutSize:subitem:count:", layoutSize, subitem, count)
	return rv
}

func (cc _CollectionLayoutGroupClass) VerticalGroupWithLayoutSize_Subitem_Count(layoutSize ICollectionLayoutSize, subitem ICollectionLayoutItem, count int) CollectionLayoutGroup {
	rv := ffi.CallMethod[CollectionLayoutGroup](cc, "verticalGroupWithLayoutSize:subitem:count:", layoutSize, subitem, count)
	return rv
}

func (cc _CollectionLayoutGroupClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutGroup {
	rv := ffi.CallMethod[CollectionLayoutGroup](cc, "itemWithLayoutSize:", layoutSize)
	return rv
}

func (cc _CollectionLayoutGroupClass) ItemWithLayoutSize_SupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutGroup {
	rv := ffi.CallMethod[CollectionLayoutGroup](cc, "itemWithLayoutSize:supplementaryItems:", layoutSize, supplementaryItems)
	return rv
}

func (cc _CollectionLayoutGroupClass) Alloc() CollectionLayoutGroup {
	rv := ffi.CallMethod[CollectionLayoutGroup](cc, "alloc")
	return rv
}

func (cc _CollectionLayoutGroupClass) New() CollectionLayoutGroup {
	rv := ffi.CallMethod[CollectionLayoutGroup](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutGroup() CollectionLayoutGroup {
	return CollectionLayoutGroupClass.New()
}

func (c_ CollectionLayoutGroup) Init() CollectionLayoutGroup {
	rv := ffi.CallMethod[CollectionLayoutGroup](c_, "init")
	return rv
}

func (c_ CollectionLayoutGroup) VisualDescription() string {
	rv := ffi.CallMethod[string](c_, "visualDescription")
	return rv
}

func (c_ CollectionLayoutGroup) Subitems() []CollectionLayoutItem {
	rv := ffi.CallMethod[[]CollectionLayoutItem](c_, "subitems")
	return rv
}

func (c_ CollectionLayoutGroup) SetSupplementaryItems(value []ICollectionLayoutSupplementaryItem) {
	ffi.CallMethod[ffi.Void](c_, "setSupplementaryItems:", value)
}

func (c_ CollectionLayoutGroup) InterItemSpacing() CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](c_, "interItemSpacing")
	return rv
}

func (c_ CollectionLayoutGroup) SetInterItemSpacing(value ICollectionLayoutSpacing) {
	ffi.CallMethod[ffi.Void](c_, "setInterItemSpacing:", value)
}
