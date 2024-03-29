// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var GradientLayerClass = _GradientLayerClass{objc.GetClass("CAGradientLayer")}

type _GradientLayerClass struct {
	objc.Class
}

type IGradientLayer interface {
	ILayer
	Colors() []objc.Object
	SetColors(value []objc.IObject)
	Locations() []foundation.Number
	SetLocations(value []foundation.INumber)
	EndPoint() coregraphics.Point
	SetEndPoint(value coregraphics.Point)
	StartPoint() coregraphics.Point
	SetStartPoint(value coregraphics.Point)
	Type() GradientLayerType
	SetType(value GradientLayerType)
}

type GradientLayer struct {
	Layer
}

func MakeGradientLayer(ptr unsafe.Pointer) GradientLayer {
	return GradientLayer{
		Layer: MakeLayer(ptr),
	}
}

func (gc _GradientLayerClass) Layer() GradientLayer {
	rv := objc.CallMethod[GradientLayer](gc, objc.SEL("layer"))
	return rv
}

func (g_ GradientLayer) Init() GradientLayer {
	rv := objc.CallMethod[GradientLayer](g_, objc.SEL("init"))
	return rv
}

func (g_ GradientLayer) InitWithLayer(layer objc.IObject) GradientLayer {
	rv := objc.CallMethod[GradientLayer](g_, objc.SEL("initWithLayer:"), objc.ExtractPtr(layer))
	return rv
}

func (g_ GradientLayer) PresentationLayer() GradientLayer {
	rv := objc.CallMethod[GradientLayer](g_, objc.SEL("presentationLayer"))
	return rv
}

func (g_ GradientLayer) ModelLayer() GradientLayer {
	rv := objc.CallMethod[GradientLayer](g_, objc.SEL("modelLayer"))
	return rv
}

func (gc _GradientLayerClass) Alloc() GradientLayer {
	rv := objc.CallMethod[GradientLayer](gc, objc.SEL("alloc"))
	return rv
}

func (gc _GradientLayerClass) New() GradientLayer {
	rv := objc.CallMethod[GradientLayer](gc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewGradientLayer() GradientLayer {
	return GradientLayerClass.New()
}

func (g_ GradientLayer) Colors() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](g_, objc.SEL("colors"))
	return rv
}

func (g_ GradientLayer) SetColors(value []objc.IObject) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setColors:"), value)
}

func (g_ GradientLayer) Locations() []foundation.Number {
	rv := objc.CallMethod[[]foundation.Number](g_, objc.SEL("locations"))
	return rv
}

func (g_ GradientLayer) SetLocations(value []foundation.INumber) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setLocations:"), value)
}

func (g_ GradientLayer) EndPoint() coregraphics.Point {
	rv := objc.CallMethod[coregraphics.Point](g_, objc.SEL("endPoint"))
	return rv
}

func (g_ GradientLayer) SetEndPoint(value coregraphics.Point) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setEndPoint:"), value)
}

func (g_ GradientLayer) StartPoint() coregraphics.Point {
	rv := objc.CallMethod[coregraphics.Point](g_, objc.SEL("startPoint"))
	return rv
}

func (g_ GradientLayer) SetStartPoint(value coregraphics.Point) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setStartPoint:"), value)
}

func (g_ GradientLayer) Type() GradientLayerType {
	rv := objc.CallMethod[GradientLayerType](g_, objc.SEL("type"))
	return rv
}

func (g_ GradientLayer) SetType(value GradientLayerType) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setType:"), value)
}
