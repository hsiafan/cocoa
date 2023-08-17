// auto-generated code, do not modify
package quartzcore

import (
	"github.com/hsiafan/cocoa/coregraphics"
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

func WrapLayoutManager(v LayoutManager) objc.Object {
	return objc.WrapAsProtocol("CALayoutManager", v)
}

type LayoutManagerBase struct {
}

func (p *LayoutManagerBase) ImplementsInvalidateLayoutOfLayer() bool {
	return false
}

func (p *LayoutManagerBase) InvalidateLayoutOfLayer(layer Layer) {
	panic("unimpemented protocol method")
}

func (p *LayoutManagerBase) ImplementsLayoutSublayersOfLayer() bool {
	return false
}

func (p *LayoutManagerBase) LayoutSublayersOfLayer(layer Layer) {
	panic("unimpemented protocol method")
}

func (p *LayoutManagerBase) ImplementsPreferredSizeOfLayer() bool {
	return false
}

func (p *LayoutManagerBase) PreferredSizeOfLayer(layer Layer) coregraphics.Size {
	panic("unimpemented protocol method")
}

type LayoutManagerWrapper struct {
	objc.Object
}

func (l_ LayoutManagerWrapper) InvalidateLayoutOfLayer(layer ILayer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("invalidateLayoutOfLayer:"), objc.ExtractPtr(layer))
}

func (l_ LayoutManagerWrapper) LayoutSublayersOfLayer(layer ILayer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("layoutSublayersOfLayer:"), objc.ExtractPtr(layer))
}

func (l_ LayoutManagerWrapper) PreferredSizeOfLayer(layer ILayer) coregraphics.Size {
	rv := objc.CallMethod[coregraphics.Size](l_, objc.GetSelector("preferredSizeOfLayer:"), objc.ExtractPtr(layer))
	return rv
}
