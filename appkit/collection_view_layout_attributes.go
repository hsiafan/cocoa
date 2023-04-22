// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[CollectionViewLayoutAttributes](cc, objc.GetSelector("layoutAttributesForItemWithIndexPath:"), indexPath)
	return rv
}

func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForSupplementaryViewOfKind_WithIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](cc, objc.GetSelector("layoutAttributesForSupplementaryViewOfKind:withIndexPath:"), elementKind, indexPath)
	return rv
}

func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForDecorationViewOfKind_WithIndexPath(decorationViewKind CollectionViewDecorationElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](cc, objc.GetSelector("layoutAttributesForDecorationViewOfKind:withIndexPath:"), decorationViewKind, indexPath)
	return rv
}

func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](cc, objc.GetSelector("layoutAttributesForInterItemGapBeforeIndexPath:"), indexPath)
	return rv
}

func (cc _CollectionViewLayoutAttributesClass) Alloc() CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](cc, objc.GetSelector("alloc"))
	return rv
}

func (cc _CollectionViewLayoutAttributesClass) New() CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewLayoutAttributes() CollectionViewLayoutAttributes {
	return CollectionViewLayoutAttributesClass.New()
}

func (c_ CollectionViewLayoutAttributes) Init() CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, objc.GetSelector("init"))
	return rv
}

func (c_ CollectionViewLayoutAttributes) RepresentedElementCategory() CollectionElementCategory {
	rv := objc.CallMethod[CollectionElementCategory](c_, objc.GetSelector("representedElementCategory"))
	return rv
}

func (c_ CollectionViewLayoutAttributes) IndexPath() foundation.IndexPath {
	rv := objc.CallMethod[foundation.IndexPath](c_, objc.GetSelector("indexPath"))
	return rv
}

func (c_ CollectionViewLayoutAttributes) SetIndexPath(value foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setIndexPath:"), value)
}

func (c_ CollectionViewLayoutAttributes) RepresentedElementKind() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("representedElementKind"))
	return rv
}

func (c_ CollectionViewLayoutAttributes) Frame() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.GetSelector("frame"))
	return rv
}

func (c_ CollectionViewLayoutAttributes) SetFrame(value foundation.Rect) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setFrame:"), value)
}

func (c_ CollectionViewLayoutAttributes) Size() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("size"))
	return rv
}

func (c_ CollectionViewLayoutAttributes) SetSize(value foundation.Size) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setSize:"), value)
}

func (c_ CollectionViewLayoutAttributes) Alpha() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("alpha"))
	return rv
}

func (c_ CollectionViewLayoutAttributes) SetAlpha(value float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setAlpha:"), value)
}

func (c_ CollectionViewLayoutAttributes) IsHidden() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isHidden"))
	return rv
}

func (c_ CollectionViewLayoutAttributes) SetHidden(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setHidden:"), value)
}

func (c_ CollectionViewLayoutAttributes) ZIndex() int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("zIndex"))
	return rv
}

func (c_ CollectionViewLayoutAttributes) SetZIndex(value int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setZIndex:"), value)
}
