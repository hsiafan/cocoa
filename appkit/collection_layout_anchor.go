// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionLayoutAnchorClass = _CollectionLayoutAnchorClass{objc.GetClass("NSCollectionLayoutAnchor")}

type _CollectionLayoutAnchorClass struct {
	objc.Class
}

type ICollectionLayoutAnchor interface {
	objc.IObject
	Edges() DirectionalRectEdge
	Offset() foundation.Point
	IsAbsoluteOffset() bool
	IsFractionalOffset() bool
}

type CollectionLayoutAnchor struct {
	objc.Object
}

func MakeCollectionLayoutAnchor(ptr unsafe.Pointer) CollectionLayoutAnchor {
	return CollectionLayoutAnchor{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutAnchorClass) LayoutAnchorWithEdges(edges DirectionalRectEdge) CollectionLayoutAnchor {
	rv := ffi.CallMethod[CollectionLayoutAnchor](cc, "layoutAnchorWithEdges:", edges)
	return rv
}

func (cc _CollectionLayoutAnchorClass) LayoutAnchorWithEdges_AbsoluteOffset(edges DirectionalRectEdge, absoluteOffset foundation.Point) CollectionLayoutAnchor {
	rv := ffi.CallMethod[CollectionLayoutAnchor](cc, "layoutAnchorWithEdges:absoluteOffset:", edges, absoluteOffset)
	return rv
}

func (cc _CollectionLayoutAnchorClass) LayoutAnchorWithEdges_FractionalOffset(edges DirectionalRectEdge, fractionalOffset foundation.Point) CollectionLayoutAnchor {
	rv := ffi.CallMethod[CollectionLayoutAnchor](cc, "layoutAnchorWithEdges:fractionalOffset:", edges, fractionalOffset)
	return rv
}

func (cc _CollectionLayoutAnchorClass) Alloc() CollectionLayoutAnchor {
	rv := ffi.CallMethod[CollectionLayoutAnchor](cc, "alloc")
	return rv
}

func (cc _CollectionLayoutAnchorClass) New() CollectionLayoutAnchor {
	rv := ffi.CallMethod[CollectionLayoutAnchor](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutAnchor() CollectionLayoutAnchor {
	return CollectionLayoutAnchorClass.New()
}

func (c_ CollectionLayoutAnchor) Init() CollectionLayoutAnchor {
	rv := ffi.CallMethod[CollectionLayoutAnchor](c_, "init")
	return rv
}

func (c_ CollectionLayoutAnchor) Edges() DirectionalRectEdge {
	rv := ffi.CallMethod[DirectionalRectEdge](c_, "edges")
	return rv
}

func (c_ CollectionLayoutAnchor) Offset() foundation.Point {
	rv := ffi.CallMethod[foundation.Point](c_, "offset")
	return rv
}

func (c_ CollectionLayoutAnchor) IsAbsoluteOffset() bool {
	rv := ffi.CallMethod[bool](c_, "isAbsoluteOffset")
	return rv
}

func (c_ CollectionLayoutAnchor) IsFractionalOffset() bool {
	rv := ffi.CallMethod[bool](c_, "isFractionalOffset")
	return rv
}
