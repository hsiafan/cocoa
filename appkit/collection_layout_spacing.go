// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionLayoutSpacingClass = _CollectionLayoutSpacingClass{objc.GetClass("NSCollectionLayoutSpacing")}

type _CollectionLayoutSpacingClass struct {
	objc.Class
}

type ICollectionLayoutSpacing interface {
	objc.IObject
	Spacing() float64
	IsFixedSpacing() bool
	IsFlexibleSpacing() bool
}

type CollectionLayoutSpacing struct {
	objc.Object
}

func MakeCollectionLayoutSpacing(ptr unsafe.Pointer) CollectionLayoutSpacing {
	return CollectionLayoutSpacing{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutSpacingClass) FixedSpacing(fixedSpacing float64) CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](cc, "fixedSpacing:", fixedSpacing)
	return rv
}

func (cc _CollectionLayoutSpacingClass) FlexibleSpacing(flexibleSpacing float64) CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](cc, "flexibleSpacing:", flexibleSpacing)
	return rv
}

func (cc _CollectionLayoutSpacingClass) Alloc() CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](cc, "alloc")
	return rv
}

func (cc _CollectionLayoutSpacingClass) New() CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutSpacing() CollectionLayoutSpacing {
	return CollectionLayoutSpacingClass.New()
}

func (c_ CollectionLayoutSpacing) Init() CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](c_, "init")
	return rv
}

func (c_ CollectionLayoutSpacing) Spacing() float64 {
	rv := ffi.CallMethod[float64](c_, "spacing")
	return rv
}

func (c_ CollectionLayoutSpacing) IsFixedSpacing() bool {
	rv := ffi.CallMethod[bool](c_, "isFixedSpacing")
	return rv
}

func (c_ CollectionLayoutSpacing) IsFlexibleSpacing() bool {
	rv := ffi.CallMethod[bool](c_, "isFlexibleSpacing")
	return rv
}
