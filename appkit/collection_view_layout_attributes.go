// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionViewLayoutAttributesClass = _CollectionViewLayoutAttributesClass{objc.GetClass("NSCollectionViewLayoutAttributes")}

type _CollectionViewLayoutAttributesClass struct {
	objc.Class
}

type ICollectionViewLayoutAttributes interface {
	objc.IObject
	RepresentedElementCategory() CollectionElementCategory
	IndexPath() foundation.IndexPath
	SetIndexPath(value foundation.IIndexPath)
	RepresentedElementKind() string
	Frame() foundation.Rect
	SetFrame(value foundation.Rect)
	Size() foundation.Size
	SetSize(value foundation.Size)
	Alpha() float64
	SetAlpha(value float64)
	IsHidden() bool
	SetHidden(value bool)
	ZIndex() int
	SetZIndex(value int)
}

type CollectionViewLayoutAttributes struct {
	objc.Object
}

func MakeCollectionViewLayoutAttributes(ptr unsafe.Pointer) CollectionViewLayoutAttributes {
	return CollectionViewLayoutAttributes{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForItemWithIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](cc, "layoutAttributesForItemWithIndexPath:", indexPath)
	return rv
}

func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForSupplementaryViewOfKind_WithIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](cc, "layoutAttributesForSupplementaryViewOfKind:withIndexPath:", elementKind, indexPath)
	return rv
}

func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForDecorationViewOfKind_WithIndexPath(decorationViewKind CollectionViewDecorationElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](cc, "layoutAttributesForDecorationViewOfKind:withIndexPath:", decorationViewKind, indexPath)
	return rv
}

func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](cc, "layoutAttributesForInterItemGapBeforeIndexPath:", indexPath)
	return rv
}

func (cc _CollectionViewLayoutAttributesClass) Alloc() CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](cc, "alloc")
	return rv
}

func (cc _CollectionViewLayoutAttributesClass) New() CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewLayoutAttributes() CollectionViewLayoutAttributes {
	return CollectionViewLayoutAttributesClass.New()
}

func (c_ CollectionViewLayoutAttributes) Init() CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](c_, "init")
	return rv
}

func (c_ CollectionViewLayoutAttributes) RepresentedElementCategory() CollectionElementCategory {
	rv := ffi.CallMethod[CollectionElementCategory](c_, "representedElementCategory")
	return rv
}

func (c_ CollectionViewLayoutAttributes) IndexPath() foundation.IndexPath {
	rv := ffi.CallMethod[foundation.IndexPath](c_, "indexPath")
	return rv
}

func (c_ CollectionViewLayoutAttributes) SetIndexPath(value foundation.IIndexPath) {
	ffi.CallMethod[ffi.Void](c_, "setIndexPath:", value)
}

func (c_ CollectionViewLayoutAttributes) RepresentedElementKind() string {
	rv := ffi.CallMethod[string](c_, "representedElementKind")
	return rv
}

func (c_ CollectionViewLayoutAttributes) Frame() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](c_, "frame")
	return rv
}

func (c_ CollectionViewLayoutAttributes) SetFrame(value foundation.Rect) {
	ffi.CallMethod[ffi.Void](c_, "setFrame:", value)
}

func (c_ CollectionViewLayoutAttributes) Size() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "size")
	return rv
}

func (c_ CollectionViewLayoutAttributes) SetSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](c_, "setSize:", value)
}

func (c_ CollectionViewLayoutAttributes) Alpha() float64 {
	rv := ffi.CallMethod[float64](c_, "alpha")
	return rv
}

func (c_ CollectionViewLayoutAttributes) SetAlpha(value float64) {
	ffi.CallMethod[ffi.Void](c_, "setAlpha:", value)
}

func (c_ CollectionViewLayoutAttributes) IsHidden() bool {
	rv := ffi.CallMethod[bool](c_, "isHidden")
	return rv
}

func (c_ CollectionViewLayoutAttributes) SetHidden(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setHidden:", value)
}

func (c_ CollectionViewLayoutAttributes) ZIndex() int {
	rv := ffi.CallMethod[int](c_, "zIndex")
	return rv
}

func (c_ CollectionViewLayoutAttributes) SetZIndex(value int) {
	ffi.CallMethod[ffi.Void](c_, "setZIndex:", value)
}
