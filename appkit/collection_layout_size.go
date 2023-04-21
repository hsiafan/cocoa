// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var CollectionLayoutSizeClass = _CollectionLayoutSizeClass{objc.GetClass("NSCollectionLayoutSize")}

type _CollectionLayoutSizeClass struct {
	objc.Class
}

type ICollectionLayoutSize interface {
	objc.IObject
	WidthDimension() CollectionLayoutDimension
	HeightDimension() CollectionLayoutDimension
}

type CollectionLayoutSize struct {
	objc.Object
}

func MakeCollectionLayoutSize(ptr unsafe.Pointer) CollectionLayoutSize {
	return CollectionLayoutSize{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutSizeClass) SizeWithWidthDimension_HeightDimension(width ICollectionLayoutDimension, height ICollectionLayoutDimension) CollectionLayoutSize {
	rv := objc.CallMethod[CollectionLayoutSize](cc, "sizeWithWidthDimension:heightDimension:", width, height)
	return rv
}

func (cc _CollectionLayoutSizeClass) Alloc() CollectionLayoutSize {
	rv := objc.CallMethod[CollectionLayoutSize](cc, "alloc")
	return rv
}

func (cc _CollectionLayoutSizeClass) New() CollectionLayoutSize {
	rv := objc.CallMethod[CollectionLayoutSize](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutSize() CollectionLayoutSize {
	return CollectionLayoutSizeClass.New()
}

func (c_ CollectionLayoutSize) Init() CollectionLayoutSize {
	rv := objc.CallMethod[CollectionLayoutSize](c_, "init")
	return rv
}

func (c_ CollectionLayoutSize) WidthDimension() CollectionLayoutDimension {
	rv := objc.CallMethod[CollectionLayoutDimension](c_, "widthDimension")
	return rv
}

func (c_ CollectionLayoutSize) HeightDimension() CollectionLayoutDimension {
	rv := objc.CallMethod[CollectionLayoutDimension](c_, "heightDimension")
	return rv
}
