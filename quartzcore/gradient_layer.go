// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[GradientLayer](gc, "layer")
	return rv
}

func (g_ GradientLayer) Init() GradientLayer {
	rv := ffi.CallMethod[GradientLayer](g_, "init")
	return rv
}

func (g_ GradientLayer) InitWithLayer(layer objc.IObject) GradientLayer {
	rv := ffi.CallMethod[GradientLayer](g_, "initWithLayer:", layer)
	return rv
}

func (g_ GradientLayer) PresentationLayer() GradientLayer {
	rv := ffi.CallMethod[GradientLayer](g_, "presentationLayer")
	return rv
}

func (g_ GradientLayer) ModelLayer() GradientLayer {
	rv := ffi.CallMethod[GradientLayer](g_, "modelLayer")
	return rv
}

func (gc _GradientLayerClass) Alloc() GradientLayer {
	rv := ffi.CallMethod[GradientLayer](gc, "alloc")
	return rv
}

func (gc _GradientLayerClass) New() GradientLayer {
	rv := ffi.CallMethod[GradientLayer](gc, "new")
	rv.Autorelease()
	return rv
}

func NewGradientLayer() GradientLayer {
	return GradientLayerClass.New()
}

func (g_ GradientLayer) Colors() []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](g_, "colors")
	return rv
}

func (g_ GradientLayer) SetColors(value []objc.IObject) {
	ffi.CallMethod[ffi.Void](g_, "setColors:", value)
}

func (g_ GradientLayer) Locations() []foundation.Number {
	rv := ffi.CallMethod[[]foundation.Number](g_, "locations")
	return rv
}

func (g_ GradientLayer) SetLocations(value []foundation.INumber) {
	ffi.CallMethod[ffi.Void](g_, "setLocations:", value)
}

func (g_ GradientLayer) EndPoint() coregraphics.Point {
	rv := ffi.CallMethod[coregraphics.Point](g_, "endPoint")
	return rv
}

func (g_ GradientLayer) SetEndPoint(value coregraphics.Point) {
	ffi.CallMethod[ffi.Void](g_, "setEndPoint:", value)
}

func (g_ GradientLayer) StartPoint() coregraphics.Point {
	rv := ffi.CallMethod[coregraphics.Point](g_, "startPoint")
	return rv
}

func (g_ GradientLayer) SetStartPoint(value coregraphics.Point) {
	ffi.CallMethod[ffi.Void](g_, "setStartPoint:", value)
}

func (g_ GradientLayer) Type() GradientLayerType {
	rv := ffi.CallMethod[GradientLayerType](g_, "type")
	return rv
}

func (g_ GradientLayer) SetType(value GradientLayerType) {
	ffi.CallMethod[ffi.Void](g_, "setType:", value)
}
