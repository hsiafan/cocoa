// auto-generated code, do not modify
package quartzcore

import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/objc"
)

type LayerDelegate interface {
	ImplementsDisplayLayer() bool
	// optional
	DisplayLayer(layer Layer)
	ImplementsDrawLayer_InContext() bool
	// optional
	DrawLayer_InContext(layer Layer, ctx coregraphics.ContextRef)
	ImplementsLayerWillDraw() bool
	// optional
	LayerWillDraw(layer Layer)
	ImplementsLayoutSublayersOfLayer() bool
	// optional
	LayoutSublayersOfLayer(layer Layer)
	ImplementsActionForLayer_ForKey() bool
	// optional
	ActionForLayer_ForKey(layer Layer, event string) objc.IObject
}

func WrapLayerDelegate(v LayerDelegate) objc.Object {
	return objc.WrapAsProtocol("CALayerDelegate", v)
}

type LayerDelegateBase struct {
}

func (p *LayerDelegateBase) ImplementsDisplayLayer() bool {
	return false
}

func (p *LayerDelegateBase) DisplayLayer(layer Layer) {
	panic("unimpemented protocol method")
}

func (p *LayerDelegateBase) ImplementsDrawLayer_InContext() bool {
	return false
}

func (p *LayerDelegateBase) DrawLayer_InContext(layer Layer, ctx coregraphics.ContextRef) {
	panic("unimpemented protocol method")
}

func (p *LayerDelegateBase) ImplementsLayerWillDraw() bool {
	return false
}

func (p *LayerDelegateBase) LayerWillDraw(layer Layer) {
	panic("unimpemented protocol method")
}

func (p *LayerDelegateBase) ImplementsLayoutSublayersOfLayer() bool {
	return false
}

func (p *LayerDelegateBase) LayoutSublayersOfLayer(layer Layer) {
	panic("unimpemented protocol method")
}

func (p *LayerDelegateBase) ImplementsActionForLayer_ForKey() bool {
	return false
}

func (p *LayerDelegateBase) ActionForLayer_ForKey(layer Layer, event string) objc.IObject {
	panic("unimpemented protocol method")
}

type LayerDelegateWrapper struct {
	objc.Object
}

func (l_ LayerDelegateWrapper) DisplayLayer(layer ILayer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("displayLayer:"), objc.ExtractPtr(layer))
}

func (l_ LayerDelegateWrapper) DrawLayer_InContext(layer ILayer, ctx coregraphics.ContextRef) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("drawLayer:inContext:"), objc.ExtractPtr(layer), ctx)
}

func (l_ LayerDelegateWrapper) LayerWillDraw(layer ILayer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("layerWillDraw:"), objc.ExtractPtr(layer))
}

func (l_ LayerDelegateWrapper) LayoutSublayersOfLayer(layer ILayer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("layoutSublayersOfLayer:"), objc.ExtractPtr(layer))
}

func (l_ LayerDelegateWrapper) ActionForLayer_ForKey(layer ILayer, event string) ActionWrapper {
	rv := objc.CallMethod[ActionWrapper](l_, objc.GetSelector("actionForLayer:forKey:"), objc.ExtractPtr(layer), event)
	return rv
}
