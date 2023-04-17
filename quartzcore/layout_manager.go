// auto-generated code, do not modify
package quartzcore

import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

type LayoutManager interface {
	ImplementsInvalidateLayoutOfLayer() bool
	// optional
	InvalidateLayoutOfLayer(layer Layer)
	ImplementsLayoutSublayersOfLayer() bool
	// optional
	LayoutSublayersOfLayer(layer Layer)
	ImplementsPreferredSizeOfLayer() bool
	// optional
	PreferredSizeOfLayer(layer Layer) coregraphics.Size
}

type LayoutManagerWrapper struct {
	objc.Object
}

func (l_ *LayoutManagerWrapper) ImplementsInvalidateLayoutOfLayer() bool {
	return l_.RespondsToSelector(objc.GetSelector("invalidateLayoutOfLayer:"))
}

func (l_ LayoutManagerWrapper) InvalidateLayoutOfLayer(layer ILayer) {
	ffi.CallMethod[ffi.Void](l_, "invalidateLayoutOfLayer:", layer)
}

func (l_ *LayoutManagerWrapper) ImplementsLayoutSublayersOfLayer() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutSublayersOfLayer:"))
}

func (l_ LayoutManagerWrapper) LayoutSublayersOfLayer(layer ILayer) {
	ffi.CallMethod[ffi.Void](l_, "layoutSublayersOfLayer:", layer)
}

func (l_ *LayoutManagerWrapper) ImplementsPreferredSizeOfLayer() bool {
	return l_.RespondsToSelector(objc.GetSelector("preferredSizeOfLayer:"))
}

func (l_ LayoutManagerWrapper) PreferredSizeOfLayer(layer ILayer) coregraphics.Size {
	rv := ffi.CallMethod[coregraphics.Size](l_, "preferredSizeOfLayer:", layer)
	return rv
}
