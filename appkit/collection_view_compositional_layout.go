// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var CollectionViewCompositionalLayoutClass = _CollectionViewCompositionalLayoutClass{objc.GetClass("NSCollectionViewCompositionalLayout")}

type _CollectionViewCompositionalLayoutClass struct {
	objc.Class
}

type ICollectionViewCompositionalLayout interface {
	ICollectionViewLayout
	Configuration() CollectionViewCompositionalLayoutConfiguration
	SetConfiguration(value ICollectionViewCompositionalLayoutConfiguration)
}

type CollectionViewCompositionalLayout struct {
	CollectionViewLayout
}

func MakeCollectionViewCompositionalLayout(ptr unsafe.Pointer) CollectionViewCompositionalLayout {
	return CollectionViewCompositionalLayout{
		CollectionViewLayout: MakeCollectionViewLayout(ptr),
	}
}

func (c_ CollectionViewCompositionalLayout) InitWithSection(section ICollectionLayoutSection) CollectionViewCompositionalLayout {
	rv := objc.CallMethod[CollectionViewCompositionalLayout](c_, "initWithSection:", section)
	return rv
}

func (c_ CollectionViewCompositionalLayout) InitWithSection_Configuration(section ICollectionLayoutSection, configuration ICollectionViewCompositionalLayoutConfiguration) CollectionViewCompositionalLayout {
	rv := objc.CallMethod[CollectionViewCompositionalLayout](c_, "initWithSection:configuration:", section, configuration)
	return rv
}

func (c_ CollectionViewCompositionalLayout) InitWithSectionProvider(sectionProvider func(section int, param2 CollectionLayoutEnvironmentWrapper) ICollectionLayoutSection) CollectionViewCompositionalLayout {
	rv := objc.CallMethod[CollectionViewCompositionalLayout](c_, "initWithSectionProvider:", sectionProvider)
	return rv
}

func (c_ CollectionViewCompositionalLayout) InitWithSectionProvider_Configuration(sectionProvider func(section int, param2 CollectionLayoutEnvironmentWrapper) ICollectionLayoutSection, configuration ICollectionViewCompositionalLayoutConfiguration) CollectionViewCompositionalLayout {
	rv := objc.CallMethod[CollectionViewCompositionalLayout](c_, "initWithSectionProvider:configuration:", sectionProvider, configuration)
	return rv
}

func (cc _CollectionViewCompositionalLayoutClass) Alloc() CollectionViewCompositionalLayout {
	rv := objc.CallMethod[CollectionViewCompositionalLayout](cc, "alloc")
	return rv
}

func (cc _CollectionViewCompositionalLayoutClass) New() CollectionViewCompositionalLayout {
	rv := objc.CallMethod[CollectionViewCompositionalLayout](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewCompositionalLayout() CollectionViewCompositionalLayout {
	return CollectionViewCompositionalLayoutClass.New()
}

func (c_ CollectionViewCompositionalLayout) Init() CollectionViewCompositionalLayout {
	rv := objc.CallMethod[CollectionViewCompositionalLayout](c_, "init")
	return rv
}

func (c_ CollectionViewCompositionalLayout) Configuration() CollectionViewCompositionalLayoutConfiguration {
	rv := objc.CallMethod[CollectionViewCompositionalLayoutConfiguration](c_, "configuration")
	return rv
}

func (c_ CollectionViewCompositionalLayout) SetConfiguration(value ICollectionViewCompositionalLayoutConfiguration) {
	objc.CallMethod[objc.Void](c_, "setConfiguration:", value)
}
