// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var CollectionViewCompositionalLayoutConfigurationClass = _CollectionViewCompositionalLayoutConfigurationClass{objc.GetClass("NSCollectionViewCompositionalLayoutConfiguration")}

type _CollectionViewCompositionalLayoutConfigurationClass struct {
	objc.Class
}

type ICollectionViewCompositionalLayoutConfiguration interface {
	objc.IObject
	ScrollDirection() CollectionViewScrollDirection
	SetScrollDirection(value CollectionViewScrollDirection)
	InterSectionSpacing() float64
	SetInterSectionSpacing(value float64)
	BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem
	SetBoundarySupplementaryItems(value []ICollectionLayoutBoundarySupplementaryItem)
}

type CollectionViewCompositionalLayoutConfiguration struct {
	objc.Object
}

func MakeCollectionViewCompositionalLayoutConfiguration(ptr unsafe.Pointer) CollectionViewCompositionalLayoutConfiguration {
	return CollectionViewCompositionalLayoutConfiguration{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionViewCompositionalLayoutConfigurationClass) Alloc() CollectionViewCompositionalLayoutConfiguration {
	rv := objc.CallMethod[CollectionViewCompositionalLayoutConfiguration](cc, objc.SEL("alloc"))
	return rv
}

func (cc _CollectionViewCompositionalLayoutConfigurationClass) New() CollectionViewCompositionalLayoutConfiguration {
	rv := objc.CallMethod[CollectionViewCompositionalLayoutConfiguration](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewCompositionalLayoutConfiguration() CollectionViewCompositionalLayoutConfiguration {
	return CollectionViewCompositionalLayoutConfigurationClass.New()
}

func (c_ CollectionViewCompositionalLayoutConfiguration) Init() CollectionViewCompositionalLayoutConfiguration {
	rv := objc.CallMethod[CollectionViewCompositionalLayoutConfiguration](c_, objc.SEL("init"))
	return rv
}

func (c_ CollectionViewCompositionalLayoutConfiguration) ScrollDirection() CollectionViewScrollDirection {
	rv := objc.CallMethod[CollectionViewScrollDirection](c_, objc.SEL("scrollDirection"))
	return rv
}

func (c_ CollectionViewCompositionalLayoutConfiguration) SetScrollDirection(value CollectionViewScrollDirection) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setScrollDirection:"), value)
}

func (c_ CollectionViewCompositionalLayoutConfiguration) InterSectionSpacing() float64 {
	rv := objc.CallMethod[float64](c_, objc.SEL("interSectionSpacing"))
	return rv
}

func (c_ CollectionViewCompositionalLayoutConfiguration) SetInterSectionSpacing(value float64) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setInterSectionSpacing:"), value)
}

func (c_ CollectionViewCompositionalLayoutConfiguration) BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[[]CollectionLayoutBoundarySupplementaryItem](c_, objc.SEL("boundarySupplementaryItems"))
	return rv
}

func (c_ CollectionViewCompositionalLayoutConfiguration) SetBoundarySupplementaryItems(value []ICollectionLayoutBoundarySupplementaryItem) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setBoundarySupplementaryItems:"), value)
}
