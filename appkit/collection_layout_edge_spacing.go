// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionLayoutEdgeSpacingClass = _CollectionLayoutEdgeSpacingClass{objc.GetClass("NSCollectionLayoutEdgeSpacing")}

type _CollectionLayoutEdgeSpacingClass struct {
	objc.Class
}

type ICollectionLayoutEdgeSpacing interface {
	objc.IObject
	Leading() CollectionLayoutSpacing
	Top() CollectionLayoutSpacing
	Trailing() CollectionLayoutSpacing
	Bottom() CollectionLayoutSpacing
}

type CollectionLayoutEdgeSpacing struct {
	objc.Object
}

func MakeCollectionLayoutEdgeSpacing(ptr unsafe.Pointer) CollectionLayoutEdgeSpacing {
	return CollectionLayoutEdgeSpacing{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutEdgeSpacingClass) SpacingForLeading_Top_Trailing_Bottom(leading ICollectionLayoutSpacing, top ICollectionLayoutSpacing, trailing ICollectionLayoutSpacing, bottom ICollectionLayoutSpacing) CollectionLayoutEdgeSpacing {
	rv := ffi.CallMethod[CollectionLayoutEdgeSpacing](cc, "spacingForLeading:top:trailing:bottom:", leading, top, trailing, bottom)
	return rv
}

func (cc _CollectionLayoutEdgeSpacingClass) Alloc() CollectionLayoutEdgeSpacing {
	rv := ffi.CallMethod[CollectionLayoutEdgeSpacing](cc, "alloc")
	return rv
}

func (cc _CollectionLayoutEdgeSpacingClass) New() CollectionLayoutEdgeSpacing {
	rv := ffi.CallMethod[CollectionLayoutEdgeSpacing](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutEdgeSpacing() CollectionLayoutEdgeSpacing {
	return CollectionLayoutEdgeSpacingClass.New()
}

func (c_ CollectionLayoutEdgeSpacing) Init() CollectionLayoutEdgeSpacing {
	rv := ffi.CallMethod[CollectionLayoutEdgeSpacing](c_, "init")
	return rv
}

func (c_ CollectionLayoutEdgeSpacing) Leading() CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](c_, "leading")
	return rv
}

func (c_ CollectionLayoutEdgeSpacing) Top() CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](c_, "top")
	return rv
}

func (c_ CollectionLayoutEdgeSpacing) Trailing() CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](c_, "trailing")
	return rv
}

func (c_ CollectionLayoutEdgeSpacing) Bottom() CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](c_, "bottom")
	return rv
}
