// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[CollectionLayoutDimension](cc, "absoluteDimension:", absoluteDimension)
	return rv
}

func (cc _CollectionLayoutDimensionClass) EstimatedDimension(estimatedDimension float64) CollectionLayoutDimension {
	rv := ffi.CallMethod[CollectionLayoutDimension](cc, "estimatedDimension:", estimatedDimension)
	return rv
}

func (cc _CollectionLayoutDimensionClass) FractionalHeightDimension(fractionalHeight float64) CollectionLayoutDimension {
	rv := ffi.CallMethod[CollectionLayoutDimension](cc, "fractionalHeightDimension:", fractionalHeight)
	return rv
}

func (cc _CollectionLayoutDimensionClass) FractionalWidthDimension(fractionalWidth float64) CollectionLayoutDimension {
	rv := ffi.CallMethod[CollectionLayoutDimension](cc, "fractionalWidthDimension:", fractionalWidth)
	return rv
}

func (cc _CollectionLayoutDimensionClass) Alloc() CollectionLayoutDimension {
	rv := ffi.CallMethod[CollectionLayoutDimension](cc, "alloc")
	return rv
}

func (cc _CollectionLayoutDimensionClass) New() CollectionLayoutDimension {
	rv := ffi.CallMethod[CollectionLayoutDimension](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutDimension() CollectionLayoutDimension {
	return CollectionLayoutDimensionClass.New()
}

func (c_ CollectionLayoutDimension) Init() CollectionLayoutDimension {
	rv := ffi.CallMethod[CollectionLayoutDimension](c_, "init")
	return rv
}

func (c_ CollectionLayoutDimension) Dimension() float64 {
	rv := ffi.CallMethod[float64](c_, "dimension")
	return rv
}

func (c_ CollectionLayoutDimension) IsAbsolute() bool {
	rv := ffi.CallMethod[bool](c_, "isAbsolute")
	return rv
}

func (c_ CollectionLayoutDimension) IsEstimated() bool {
	rv := ffi.CallMethod[bool](c_, "isEstimated")
	return rv
}

func (c_ CollectionLayoutDimension) IsFractionalHeight() bool {
	rv := ffi.CallMethod[bool](c_, "isFractionalHeight")
	return rv
}

func (c_ CollectionLayoutDimension) IsFractionalWidth() bool {
	rv := ffi.CallMethod[bool](c_, "isFractionalWidth")
	return rv
}
