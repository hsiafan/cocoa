// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[CollectionViewLayoutInvalidationContext](cc, "alloc")
	return rv
}

func (cc _CollectionViewLayoutInvalidationContextClass) New() CollectionViewLayoutInvalidationContext {
	rv := ffi.CallMethod[CollectionViewLayoutInvalidationContext](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewLayoutInvalidationContext() CollectionViewLayoutInvalidationContext {
	return CollectionViewLayoutInvalidationContextClass.New()
}

func (c_ CollectionViewLayoutInvalidationContext) Init() CollectionViewLayoutInvalidationContext {
	rv := ffi.CallMethod[CollectionViewLayoutInvalidationContext](c_, "init")
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateItemsAtIndexPaths(indexPaths foundation.ISet) {
	ffi.CallMethod[ffi.Void](c_, "invalidateItemsAtIndexPaths:", indexPaths)
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateSupplementaryElementsOfKind_AtIndexPaths(elementKind CollectionViewSupplementaryElementKind, indexPaths foundation.ISet) {
	ffi.CallMethod[ffi.Void](c_, "invalidateSupplementaryElementsOfKind:atIndexPaths:", elementKind, indexPaths)
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateDecorationElementsOfKind_AtIndexPaths(elementKind CollectionViewDecorationElementKind, indexPaths foundation.ISet) {
	ffi.CallMethod[ffi.Void](c_, "invalidateDecorationElementsOfKind:atIndexPaths:", elementKind, indexPaths)
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateEverything() bool {
	rv := ffi.CallMethod[bool](c_, "invalidateEverything")
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateDataSourceCounts() bool {
	rv := ffi.CallMethod[bool](c_, "invalidateDataSourceCounts")
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) ContentOffsetAdjustment() foundation.Point {
	rv := ffi.CallMethod[foundation.Point](c_, "contentOffsetAdjustment")
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) SetContentOffsetAdjustment(value foundation.Point) {
	ffi.CallMethod[ffi.Void](c_, "setContentOffsetAdjustment:", value)
}

func (c_ CollectionViewLayoutInvalidationContext) ContentSizeAdjustment() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "contentSizeAdjustment")
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) SetContentSizeAdjustment(value foundation.Size) {
	ffi.CallMethod[ffi.Void](c_, "setContentSizeAdjustment:", value)
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidatedItemIndexPaths() foundation.Set {
	rv := ffi.CallMethod[foundation.Set](c_, "invalidatedItemIndexPaths")
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidatedSupplementaryIndexPaths() map[CollectionViewSupplementaryElementKind]foundation.Set {
	rv := ffi.CallMethod[map[CollectionViewSupplementaryElementKind]foundation.Set](c_, "invalidatedSupplementaryIndexPaths")
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidatedDecorationIndexPaths() map[CollectionViewDecorationElementKind]foundation.Set {
	rv := ffi.CallMethod[map[CollectionViewDecorationElementKind]foundation.Set](c_, "invalidatedDecorationIndexPaths")
	return rv
}
