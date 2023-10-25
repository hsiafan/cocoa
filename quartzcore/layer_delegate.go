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

type LayerDelegateCreator struct {
	className string
	class     objc.Class
}

func NewLayerDelegateCreator(name string) *LayerDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("ProtocolBase"), name, 0)
	objc.RegisterClassPair(class)
	return &LayerDelegateCreator{className: name, class: class}
}

func (c *LayerDelegateCreator) SetDisplayLayer(handle func(o objc.ProtocolBase, layer Layer)) {
	objc.AddMethod(c.class, objc.SEL("displayLayer:"), handle)
}

func (c *LayerDelegateCreator) SetDrawLayer_InContext(handle func(o objc.ProtocolBase, layer Layer, ctx coregraphics.ContextRef)) {
	objc.AddMethod(c.class, objc.SEL("drawLayer:inContext:"), handle)
}

func (c *LayerDelegateCreator) SetLayerWillDraw(handle func(o objc.ProtocolBase, layer Layer)) {
	objc.AddMethod(c.class, objc.SEL("layerWillDraw:"), handle)
}

func (c *LayerDelegateCreator) SetLayoutSublayersOfLayer(handle func(o objc.ProtocolBase, layer Layer)) {
	objc.AddMethod(c.class, objc.SEL("layoutSublayersOfLayer:"), handle)
}

func (c *LayerDelegateCreator) SetActionForLayer_ForKey(handle func(o objc.ProtocolBase, layer Layer, event string) objc.IObject) {
	objc.AddMethod(c.class, objc.SEL("actionForLayer:forKey:"), handle)
}

func (c *LayerDelegateCreator) Create() objc.ProtocolBase {
	return objc.ProtocolBase{Object: c.class.CreateInstance(0)}
}
