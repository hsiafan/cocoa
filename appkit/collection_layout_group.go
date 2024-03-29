// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[CollectionLayoutGroup](cc, objc.SEL("horizontalGroupWithLayoutSize:subitems:"), objc.ExtractPtr(layoutSize), subitems)
	return rv
}

func (cc _CollectionLayoutGroupClass) VerticalGroupWithLayoutSize_Subitems(layoutSize ICollectionLayoutSize, subitems []ICollectionLayoutItem) CollectionLayoutGroup {
	rv := objc.CallMethod[CollectionLayoutGroup](cc, objc.SEL("verticalGroupWithLayoutSize:subitems:"), objc.ExtractPtr(layoutSize), subitems)
	return rv
}

func (cc _CollectionLayoutGroupClass) CustomGroupWithLayoutSize_ItemProvider(layoutSize ICollectionLayoutSize, itemProvider func(layoutEnvironment objc.Object) []ICollectionLayoutGroupCustomItem) CollectionLayoutGroup {
	rv := objc.CallMethod[CollectionLayoutGroup](cc, objc.SEL("customGroupWithLayoutSize:itemProvider:"), objc.ExtractPtr(layoutSize), itemProvider)
	return rv
}

func (cc _CollectionLayoutGroupClass) HorizontalGroupWithLayoutSize_Subitem_Count(layoutSize ICollectionLayoutSize, subitem ICollectionLayoutItem, count int) CollectionLayoutGroup {
	rv := objc.CallMethod[CollectionLayoutGroup](cc, objc.SEL("horizontalGroupWithLayoutSize:subitem:count:"), objc.ExtractPtr(layoutSize), objc.ExtractPtr(subitem), count)
	return rv
}

func (cc _CollectionLayoutGroupClass) VerticalGroupWithLayoutSize_Subitem_Count(layoutSize ICollectionLayoutSize, subitem ICollectionLayoutItem, count int) CollectionLayoutGroup {
	rv := objc.CallMethod[CollectionLayoutGroup](cc, objc.SEL("verticalGroupWithLayoutSize:subitem:count:"), objc.ExtractPtr(layoutSize), objc.ExtractPtr(subitem), count)
	return rv
}

func (cc _CollectionLayoutGroupClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutGroup {
	rv := objc.CallMethod[CollectionLayoutGroup](cc, objc.SEL("itemWithLayoutSize:"), objc.ExtractPtr(layoutSize))
	return rv
}

func (cc _CollectionLayoutGroupClass) ItemWithLayoutSize_SupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutGroup {
	rv := objc.CallMethod[CollectionLayoutGroup](cc, objc.SEL("itemWithLayoutSize:supplementaryItems:"), objc.ExtractPtr(layoutSize), supplementaryItems)
	return rv
}

func (cc _CollectionLayoutGroupClass) Alloc() CollectionLayoutGroup {
	rv := objc.CallMethod[CollectionLayoutGroup](cc, objc.SEL("alloc"))
	return rv
}

func (cc _CollectionLayoutGroupClass) New() CollectionLayoutGroup {
	rv := objc.CallMethod[CollectionLayoutGroup](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutGroup() CollectionLayoutGroup {
	return CollectionLayoutGroupClass.New()
}

func (c_ CollectionLayoutGroup) Init() CollectionLayoutGroup {
	rv := objc.CallMethod[CollectionLayoutGroup](c_, objc.SEL("init"))
	return rv
}

func (c_ CollectionLayoutGroup) VisualDescription() string {
	rv := objc.CallMethod[string](c_, objc.SEL("visualDescription"))
	return rv
}

func (c_ CollectionLayoutGroup) Subitems() []CollectionLayoutItem {
	rv := objc.CallMethod[[]CollectionLayoutItem](c_, objc.SEL("subitems"))
	return rv
}

func (c_ CollectionLayoutGroup) SetSupplementaryItems(value []ICollectionLayoutSupplementaryItem) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setSupplementaryItems:"), value)
}

func (c_ CollectionLayoutGroup) InterItemSpacing() CollectionLayoutSpacing {
	rv := objc.CallMethod[CollectionLayoutSpacing](c_, objc.SEL("interItemSpacing"))
	return rv
}

func (c_ CollectionLayoutGroup) SetInterItemSpacing(value ICollectionLayoutSpacing) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setInterItemSpacing:"), objc.ExtractPtr(value))
}
