// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var CollectionLayoutDimensionClass = _CollectionLayoutDimensionClass{objc.GetClass("NSCollectionLayoutDimension")}

type _CollectionLayoutDimensionClass struct {
	objc.Class
}

type ICollectionLayoutDimension interface {
	objc.IObject
	Dimension() float64
	IsAbsolute() bool
	IsEstimated() bool
	IsFractionalHeight() bool
	IsFractionalWidth() bool
}

type CollectionLayoutDimension struct {
	objc.Object
}

func MakeCollectionLayoutDimension(ptr unsafe.Pointer) CollectionLayoutDimension {
	return CollectionLayoutDimension{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutDimensionClass) AbsoluteDimension(absoluteDimension float64) CollectionLayoutDimension {
	rv := objc.CallMethod[CollectionLayoutDimension](cc, objc.SEL("absoluteDimension:"), absoluteDimension)
	return rv
}

func (cc _CollectionLayoutDimensionClass) EstimatedDimension(estimatedDimension float64) CollectionLayoutDimension {
	rv := objc.CallMethod[CollectionLayoutDimension](cc, objc.SEL("estimatedDimension:"), estimatedDimension)
	return rv
}

func (cc _CollectionLayoutDimensionClass) FractionalHeightDimension(fractionalHeight float64) CollectionLayoutDimension {
	rv := objc.CallMethod[CollectionLayoutDimension](cc, objc.SEL("fractionalHeightDimension:"), fractionalHeight)
	return rv
}

func (cc _CollectionLayoutDimensionClass) FractionalWidthDimension(fractionalWidth float64) CollectionLayoutDimension {
	rv := objc.CallMethod[CollectionLayoutDimension](cc, objc.SEL("fractionalWidthDimension:"), fractionalWidth)
	return rv
}

func (cc _CollectionLayoutDimensionClass) Alloc() CollectionLayoutDimension {
	rv := objc.CallMethod[CollectionLayoutDimension](cc, objc.SEL("alloc"))
	return rv
}

func (cc _CollectionLayoutDimensionClass) New() CollectionLayoutDimension {
	rv := objc.CallMethod[CollectionLayoutDimension](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutDimension() CollectionLayoutDimension {
	return CollectionLayoutDimensionClass.New()
}

func (c_ CollectionLayoutDimension) Init() CollectionLayoutDimension {
	rv := objc.CallMethod[CollectionLayoutDimension](c_, objc.SEL("init"))
	return rv
}

func (c_ CollectionLayoutDimension) Dimension() float64 {
	rv := objc.CallMethod[float64](c_, objc.SEL("dimension"))
	return rv
}

func (c_ CollectionLayoutDimension) IsAbsolute() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isAbsolute"))
	return rv
}

func (c_ CollectionLayoutDimension) IsEstimated() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isEstimated"))
	return rv
}

func (c_ CollectionLayoutDimension) IsFractionalHeight() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isFractionalHeight"))
	return rv
}

func (c_ CollectionLayoutDimension) IsFractionalWidth() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isFractionalWidth"))
	return rv
}
