// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[CollectionLayoutSpacing](cc, objc.GetSelector("fixedSpacing:"), fixedSpacing)
	return rv
}

func (cc _CollectionLayoutSpacingClass) FlexibleSpacing(flexibleSpacing float64) CollectionLayoutSpacing {
	rv := objc.CallMethod[CollectionLayoutSpacing](cc, objc.GetSelector("flexibleSpacing:"), flexibleSpacing)
	return rv
}

func (cc _CollectionLayoutSpacingClass) Alloc() CollectionLayoutSpacing {
	rv := objc.CallMethod[CollectionLayoutSpacing](cc, objc.GetSelector("alloc"))
	return rv
}

func (cc _CollectionLayoutSpacingClass) New() CollectionLayoutSpacing {
	rv := objc.CallMethod[CollectionLayoutSpacing](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutSpacing() CollectionLayoutSpacing {
	return CollectionLayoutSpacingClass.New()
}

func (c_ CollectionLayoutSpacing) Init() CollectionLayoutSpacing {
	rv := objc.CallMethod[CollectionLayoutSpacing](c_, objc.GetSelector("init"))
	return rv
}

func (c_ CollectionLayoutSpacing) Spacing() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("spacing"))
	return rv
}

func (c_ CollectionLayoutSpacing) IsFixedSpacing() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isFixedSpacing"))
	return rv
}

func (c_ CollectionLayoutSpacing) IsFlexibleSpacing() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isFlexibleSpacing"))
	return rv
}
