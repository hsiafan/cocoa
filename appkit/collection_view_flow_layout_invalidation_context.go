// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionViewFlowLayoutInvalidationContextClass = _CollectionViewFlowLayoutInvalidationContextClass{objc.GetClass("NSCollectionViewFlowLayoutInvalidationContext")}

type _CollectionViewFlowLayoutInvalidationContextClass struct {
	objc.Class
}

type ICollectionViewFlowLayoutInvalidationContext interface {
	ICollectionViewLayoutInvalidationContext
	InvalidateFlowLayoutAttributes() bool
	SetInvalidateFlowLayoutAttributes(value bool)
	InvalidateFlowLayoutDelegateMetrics() bool
	SetInvalidateFlowLayoutDelegateMetrics(value bool)
}

type CollectionViewFlowLayoutInvalidationContext struct {
	CollectionViewLayoutInvalidationContext
}

func MakeCollectionViewFlowLayoutInvalidationContext(ptr unsafe.Pointer) CollectionViewFlowLayoutInvalidationContext {
	return CollectionViewFlowLayoutInvalidationContext{
		CollectionViewLayoutInvalidationContext: MakeCollectionViewLayoutInvalidationContext(ptr),
	}
}

func (cc _CollectionViewFlowLayoutInvalidationContextClass) Alloc() CollectionViewFlowLayoutInvalidationContext {
	rv := ffi.CallMethod[CollectionViewFlowLayoutInvalidationContext](cc, "alloc")
	return rv
}

func (cc _CollectionViewFlowLayoutInvalidationContextClass) New() CollectionViewFlowLayoutInvalidationContext {
	rv := ffi.CallMethod[CollectionViewFlowLayoutInvalidationContext](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewFlowLayoutInvalidationContext() CollectionViewFlowLayoutInvalidationContext {
	return CollectionViewFlowLayoutInvalidationContextClass.New()
}

func (c_ CollectionViewFlowLayoutInvalidationContext) Init() CollectionViewFlowLayoutInvalidationContext {
	rv := ffi.CallMethod[CollectionViewFlowLayoutInvalidationContext](c_, "init")
	return rv
}

func (c_ CollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutAttributes() bool {
	rv := ffi.CallMethod[bool](c_, "invalidateFlowLayoutAttributes")
	return rv
}

func (c_ CollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutAttributes(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setInvalidateFlowLayoutAttributes:", value)
}

func (c_ CollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutDelegateMetrics() bool {
	rv := ffi.CallMethod[bool](c_, "invalidateFlowLayoutDelegateMetrics")
	return rv
}

func (c_ CollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutDelegateMetrics(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setInvalidateFlowLayoutDelegateMetrics:", value)
}
