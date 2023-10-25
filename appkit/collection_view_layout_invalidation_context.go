// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionViewLayoutInvalidationContextClass = _CollectionViewLayoutInvalidationContextClass{objc.GetClass("NSCollectionViewLayoutInvalidationContext")}

type _CollectionViewLayoutInvalidationContextClass struct {
	objc.Class
}

type ICollectionViewLayoutInvalidationContext interface {
	objc.IObject
	InvalidateItemsAtIndexPaths(indexPaths foundation.ISet)
	InvalidateSupplementaryElementsOfKind_AtIndexPaths(elementKind CollectionViewSupplementaryElementKind, indexPaths foundation.ISet)
	InvalidateDecorationElementsOfKind_AtIndexPaths(elementKind CollectionViewDecorationElementKind, indexPaths foundation.ISet)
	InvalidateEverything() bool
	InvalidateDataSourceCounts() bool
	ContentOffsetAdjustment() foundation.Point
	SetContentOffsetAdjustment(value foundation.Point)
	ContentSizeAdjustment() foundation.Size
	SetContentSizeAdjustment(value foundation.Size)
	InvalidatedItemIndexPaths() foundation.Set
	InvalidatedSupplementaryIndexPaths() map[CollectionViewSupplementaryElementKind]foundation.Set
	InvalidatedDecorationIndexPaths() map[CollectionViewDecorationElementKind]foundation.Set
}

type CollectionViewLayoutInvalidationContext struct {
	objc.Object
}

func MakeCollectionViewLayoutInvalidationContext(ptr unsafe.Pointer) CollectionViewLayoutInvalidationContext {
	return CollectionViewLayoutInvalidationContext{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionViewLayoutInvalidationContextClass) Alloc() CollectionViewLayoutInvalidationContext {
	rv := objc.CallMethod[CollectionViewLayoutInvalidationContext](cc, objc.SEL("alloc"))
	return rv
}

func (cc _CollectionViewLayoutInvalidationContextClass) New() CollectionViewLayoutInvalidationContext {
	rv := objc.CallMethod[CollectionViewLayoutInvalidationContext](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewLayoutInvalidationContext() CollectionViewLayoutInvalidationContext {
	return CollectionViewLayoutInvalidationContextClass.New()
}

func (c_ CollectionViewLayoutInvalidationContext) Init() CollectionViewLayoutInvalidationContext {
	rv := objc.CallMethod[CollectionViewLayoutInvalidationContext](c_, objc.SEL("init"))
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateItemsAtIndexPaths(indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.SEL("invalidateItemsAtIndexPaths:"), objc.ExtractPtr(indexPaths))
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateSupplementaryElementsOfKind_AtIndexPaths(elementKind CollectionViewSupplementaryElementKind, indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.SEL("invalidateSupplementaryElementsOfKind:atIndexPaths:"), elementKind, objc.ExtractPtr(indexPaths))
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateDecorationElementsOfKind_AtIndexPaths(elementKind CollectionViewDecorationElementKind, indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.SEL("invalidateDecorationElementsOfKind:atIndexPaths:"), elementKind, objc.ExtractPtr(indexPaths))
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateEverything() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("invalidateEverything"))
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateDataSourceCounts() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("invalidateDataSourceCounts"))
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) ContentOffsetAdjustment() foundation.Point {
	rv := objc.CallMethod[foundation.Point](c_, objc.SEL("contentOffsetAdjustment"))
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) SetContentOffsetAdjustment(value foundation.Point) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setContentOffsetAdjustment:"), value)
}

func (c_ CollectionViewLayoutInvalidationContext) ContentSizeAdjustment() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.SEL("contentSizeAdjustment"))
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) SetContentSizeAdjustment(value foundation.Size) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setContentSizeAdjustment:"), value)
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidatedItemIndexPaths() foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, objc.SEL("invalidatedItemIndexPaths"))
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidatedSupplementaryIndexPaths() map[CollectionViewSupplementaryElementKind]foundation.Set {
	rv := objc.CallMethod[map[CollectionViewSupplementaryElementKind]foundation.Set](c_, objc.SEL("invalidatedSupplementaryIndexPaths"))
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidatedDecorationIndexPaths() map[CollectionViewDecorationElementKind]foundation.Set {
	rv := objc.CallMethod[map[CollectionViewDecorationElementKind]foundation.Set](c_, objc.SEL("invalidatedDecorationIndexPaths"))
	return rv
}
