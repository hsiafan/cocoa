// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ShapeLayerClass = _ShapeLayerClass{objc.GetClass("CAShapeLayer")}

type _ShapeLayerClass struct {
	objc.Class
}

type IShapeLayer interface {
	ILayer
	Path() coregraphics.PathRef
	SetPath(value coregraphics.PathRef)
	FillColor() coregraphics.ColorRef
	SetFillColor(value coregraphics.ColorRef)
	FillRule() ShapeLayerFillRule
	SetFillRule(value ShapeLayerFillRule)
	LineCap() ShapeLayerLineCap
	SetLineCap(value ShapeLayerLineCap)
	LineDashPattern() []foundation.Number
	SetLineDashPattern(value []foundation.INumber)
	LineDashPhase() float64
	SetLineDashPhase(value float64)
	LineJoin() ShapeLayerLineJoin
	SetLineJoin(value ShapeLayerLineJoin)
	LineWidth() float64
	SetLineWidth(value float64)
	MiterLimit() float64
	SetMiterLimit(value float64)
	StrokeColor() coregraphics.ColorRef
	SetStrokeColor(value coregraphics.ColorRef)
	StrokeStart() float64
	SetStrokeStart(value float64)
	StrokeEnd() float64
	SetStrokeEnd(value float64)
}

type ShapeLayer struct {
	Layer
}

func MakeShapeLayer(ptr unsafe.Pointer) ShapeLayer {
	return ShapeLayer{
		Layer: MakeLayer(ptr),
	}
}

func (sc _ShapeLayerClass) Layer() ShapeLayer {
	rv := objc.CallMethod[ShapeLayer](sc, "layer")
	return rv
}

func (s_ ShapeLayer) Init() ShapeLayer {
	rv := objc.CallMethod[ShapeLayer](s_, "init")
	return rv
}

func (s_ ShapeLayer) InitWithLayer(layer objc.IObject) ShapeLayer {
	rv := objc.CallMethod[ShapeLayer](s_, "initWithLayer:", layer)
	return rv
}

func (s_ ShapeLayer) PresentationLayer() ShapeLayer {
	rv := objc.CallMethod[ShapeLayer](s_, "presentationLayer")
	return rv
}

func (s_ ShapeLayer) ModelLayer() ShapeLayer {
	rv := objc.CallMethod[ShapeLayer](s_, "modelLayer")
	return rv
}

func (sc _ShapeLayerClass) Alloc() ShapeLayer {
	rv := objc.CallMethod[ShapeLayer](sc, "alloc")
	return rv
}

func (sc _ShapeLayerClass) New() ShapeLayer {
	rv := objc.CallMethod[ShapeLayer](sc, "new")
	rv.Autorelease()
	return rv
}

func NewShapeLayer() ShapeLayer {
	return ShapeLayerClass.New()
}

func (s_ ShapeLayer) Path() coregraphics.PathRef {
	rv := objc.CallMethod[coregraphics.PathRef](s_, "path")
	return rv
}

func (s_ ShapeLayer) SetPath(value coregraphics.PathRef) {
	objc.CallMethod[objc.Void](s_, "setPath:", value)
}

func (s_ ShapeLayer) FillColor() coregraphics.ColorRef {
	rv := objc.CallMethod[coregraphics.ColorRef](s_, "fillColor")
	return rv
}

func (s_ ShapeLayer) SetFillColor(value coregraphics.ColorRef) {
	objc.CallMethod[objc.Void](s_, "setFillColor:", value)
}

func (s_ ShapeLayer) FillRule() ShapeLayerFillRule {
	rv := objc.CallMethod[ShapeLayerFillRule](s_, "fillRule")
	return rv
}

func (s_ ShapeLayer) SetFillRule(value ShapeLayerFillRule) {
	objc.CallMethod[objc.Void](s_, "setFillRule:", value)
}

func (s_ ShapeLayer) LineCap() ShapeLayerLineCap {
	rv := objc.CallMethod[ShapeLayerLineCap](s_, "lineCap")
	return rv
}

func (s_ ShapeLayer) SetLineCap(value ShapeLayerLineCap) {
	objc.CallMethod[objc.Void](s_, "setLineCap:", value)
}

func (s_ ShapeLayer) LineDashPattern() []foundation.Number {
	rv := objc.CallMethod[[]foundation.Number](s_, "lineDashPattern")
	return rv
}

func (s_ ShapeLayer) SetLineDashPattern(value []foundation.INumber) {
	objc.CallMethod[objc.Void](s_, "setLineDashPattern:", value)
}

func (s_ ShapeLayer) LineDashPhase() float64 {
	rv := objc.CallMethod[float64](s_, "lineDashPhase")
	return rv
}

func (s_ ShapeLayer) SetLineDashPhase(value float64) {
	objc.CallMethod[objc.Void](s_, "setLineDashPhase:", value)
}

func (s_ ShapeLayer) LineJoin() ShapeLayerLineJoin {
	rv := objc.CallMethod[ShapeLayerLineJoin](s_, "lineJoin")
	return rv
}

func (s_ ShapeLayer) SetLineJoin(value ShapeLayerLineJoin) {
	objc.CallMethod[objc.Void](s_, "setLineJoin:", value)
}

func (s_ ShapeLayer) LineWidth() float64 {
	rv := objc.CallMethod[float64](s_, "lineWidth")
	return rv
}

func (s_ ShapeLayer) SetLineWidth(value float64) {
	objc.CallMethod[objc.Void](s_, "setLineWidth:", value)
}

func (s_ ShapeLayer) MiterLimit() float64 {
	rv := objc.CallMethod[float64](s_, "miterLimit")
	return rv
}

func (s_ ShapeLayer) SetMiterLimit(value float64) {
	objc.CallMethod[objc.Void](s_, "setMiterLimit:", value)
}

func (s_ ShapeLayer) StrokeColor() coregraphics.ColorRef {
	rv := objc.CallMethod[coregraphics.ColorRef](s_, "strokeColor")
	return rv
}

func (s_ ShapeLayer) SetStrokeColor(value coregraphics.ColorRef) {
	objc.CallMethod[objc.Void](s_, "setStrokeColor:", value)
}

func (s_ ShapeLayer) StrokeStart() float64 {
	rv := objc.CallMethod[float64](s_, "strokeStart")
	return rv
}

func (s_ ShapeLayer) SetStrokeStart(value float64) {
	objc.CallMethod[objc.Void](s_, "setStrokeStart:", value)
}

func (s_ ShapeLayer) StrokeEnd() float64 {
	rv := objc.CallMethod[float64](s_, "strokeEnd")
	return rv
}

func (s_ ShapeLayer) SetStrokeEnd(value float64) {
	objc.CallMethod[objc.Void](s_, "setStrokeEnd:", value)
}
