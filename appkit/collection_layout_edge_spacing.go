// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[CollectionLayoutEdgeSpacing](cc, objc.SEL("spacingForLeading:top:trailing:bottom:"), objc.ExtractPtr(leading), objc.ExtractPtr(top), objc.ExtractPtr(trailing), objc.ExtractPtr(bottom))
	return rv
}

func (cc _CollectionLayoutEdgeSpacingClass) Alloc() CollectionLayoutEdgeSpacing {
	rv := objc.CallMethod[CollectionLayoutEdgeSpacing](cc, objc.SEL("alloc"))
	return rv
}

func (cc _CollectionLayoutEdgeSpacingClass) New() CollectionLayoutEdgeSpacing {
	rv := objc.CallMethod[CollectionLayoutEdgeSpacing](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutEdgeSpacing() CollectionLayoutEdgeSpacing {
	return CollectionLayoutEdgeSpacingClass.New()
}

func (c_ CollectionLayoutEdgeSpacing) Init() CollectionLayoutEdgeSpacing {
	rv := objc.CallMethod[CollectionLayoutEdgeSpacing](c_, objc.SEL("init"))
	return rv
}

func (c_ CollectionLayoutEdgeSpacing) Leading() CollectionLayoutSpacing {
	rv := objc.CallMethod[CollectionLayoutSpacing](c_, objc.SEL("leading"))
	return rv
}

func (c_ CollectionLayoutEdgeSpacing) Top() CollectionLayoutSpacing {
	rv := objc.CallMethod[CollectionLayoutSpacing](c_, objc.SEL("top"))
	return rv
}

func (c_ CollectionLayoutEdgeSpacing) Trailing() CollectionLayoutSpacing {
	rv := objc.CallMethod[CollectionLayoutSpacing](c_, objc.SEL("trailing"))
	return rv
}

func (c_ CollectionLayoutEdgeSpacing) Bottom() CollectionLayoutSpacing {
	rv := objc.CallMethod[CollectionLayoutSpacing](c_, objc.SEL("bottom"))
	return rv
}
