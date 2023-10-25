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
