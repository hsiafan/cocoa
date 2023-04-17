// auto-generated code, do not modify
package quartzcore

import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
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
	ActionForLayer_ForKey(layer Layer, event string) Action
}

type LayerDelegateImpl struct {
	_DisplayLayer           func(layer Layer)
	_DrawLayer_InContext    func(layer Layer, ctx coregraphics.ContextRef)
	_LayerWillDraw          func(layer Layer)
	_LayoutSublayersOfLayer func(layer Layer)
	_ActionForLayer_ForKey  func(layer Layer, event string) Action
}

func (di *LayerDelegateImpl) ImplementsDisplayLayer() bool {
	return di._DisplayLayer != nil
}

func (di *LayerDelegateImpl) SetDisplayLayer(f func(layer Layer)) {
	di._DisplayLayer = f
}

func (di *LayerDelegateImpl) DisplayLayer(layer Layer) {
	di._DisplayLayer(layer)
}
func (di *LayerDelegateImpl) ImplementsDrawLayer_InContext() bool {
	return di._DrawLayer_InContext != nil
}

func (di *LayerDelegateImpl) SetDrawLayer_InContext(f func(layer Layer, ctx coregraphics.ContextRef)) {
	di._DrawLayer_InContext = f
}

func (di *LayerDelegateImpl) DrawLayer_InContext(layer Layer, ctx coregraphics.ContextRef) {
	di._DrawLayer_InContext(layer, ctx)
}
func (di *LayerDelegateImpl) ImplementsLayerWillDraw() bool {
	return di._LayerWillDraw != nil
}

func (di *LayerDelegateImpl) SetLayerWillDraw(f func(layer Layer)) {
	di._LayerWillDraw = f
}

func (di *LayerDelegateImpl) LayerWillDraw(layer Layer) {
	di._LayerWillDraw(layer)
}
func (di *LayerDelegateImpl) ImplementsLayoutSublayersOfLayer() bool {
	return di._LayoutSublayersOfLayer != nil
}

func (di *LayerDelegateImpl) SetLayoutSublayersOfLayer(f func(layer Layer)) {
	di._LayoutSublayersOfLayer = f
}

func (di *LayerDelegateImpl) LayoutSublayersOfLayer(layer Layer) {
	di._LayoutSublayersOfLayer(layer)
}
func (di *LayerDelegateImpl) ImplementsActionForLayer_ForKey() bool {
	return di._ActionForLayer_ForKey != nil
}

func (di *LayerDelegateImpl) SetActionForLayer_ForKey(f func(layer Layer, event string) Action) {
	di._ActionForLayer_ForKey = f
}

func (di *LayerDelegateImpl) ActionForLayer_ForKey(layer Layer, event string) Action {
	return di._ActionForLayer_ForKey(layer, event)
}

type LayerDelegateWrapper struct {
	objc.Object
}

func (l_ *LayerDelegateWrapper) ImplementsDisplayLayer() bool {
	return l_.RespondsToSelector(objc.GetSelector("displayLayer:"))
}

func (l_ LayerDelegateWrapper) DisplayLayer(layer ILayer) {
	ffi.CallMethod[ffi.Void](l_, "displayLayer:", layer)
}

func (l_ *LayerDelegateWrapper) ImplementsDrawLayer_InContext() bool {
	return l_.RespondsToSelector(objc.GetSelector("drawLayer:inContext:"))
}

func (l_ LayerDelegateWrapper) DrawLayer_InContext(layer ILayer, ctx coregraphics.ContextRef) {
	ffi.CallMethod[ffi.Void](l_, "drawLayer:inContext:", layer, ctx)
}

func (l_ *LayerDelegateWrapper) ImplementsLayerWillDraw() bool {
	return l_.RespondsToSelector(objc.GetSelector("layerWillDraw:"))
}

func (l_ LayerDelegateWrapper) LayerWillDraw(layer ILayer) {
	ffi.CallMethod[ffi.Void](l_, "layerWillDraw:", layer)
}

func (l_ *LayerDelegateWrapper) ImplementsLayoutSublayersOfLayer() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutSublayersOfLayer:"))
}

func (l_ LayerDelegateWrapper) LayoutSublayersOfLayer(layer ILayer) {
	ffi.CallMethod[ffi.Void](l_, "layoutSublayersOfLayer:", layer)
}

func (l_ *LayerDelegateWrapper) ImplementsActionForLayer_ForKey() bool {
	return l_.RespondsToSelector(objc.GetSelector("actionForLayer:forKey:"))
}

func (l_ LayerDelegateWrapper) ActionForLayer_ForKey(layer ILayer, event string) ActionWrapper {
	rv := ffi.CallMethod[ActionWrapper](l_, "actionForLayer:forKey:", layer, event)
	return rv
}
