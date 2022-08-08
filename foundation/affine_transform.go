// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var AffineTransformClass = _AffineTransformClass{objc.GetClass("NSAffineTransform")}

type _AffineTransformClass struct {
	objc.Class
}

type IAffineTransform interface {
	objc.IObject
	RotateByDegrees(angle float64)
	RotateByRadians(angle float64)
	ScaleBy(scale float64)
	ScaleXBy_YBy(scaleX float64, scaleY float64)
	TranslateXBy_YBy(deltaX float64, deltaY float64)
	AppendTransform(transform IAffineTransform)
	PrependTransform(transform IAffineTransform)
	Invert()
	TransformPoint(aPoint Point) Point
	TransformSize(aSize Size) Size
	TransformStruct() AffineTransformStruct
	SetTransformStruct(value AffineTransformStruct)
}

type AffineTransform struct {
	objc.Object
}

func MakeAffineTransform(ptr unsafe.Pointer) AffineTransform {
	return AffineTransform{
		Object: objc.MakeObject(ptr),
	}
}

func (a_ AffineTransform) Init() AffineTransform {
	rv := ffi.CallMethod[AffineTransform](a_, "init")
	return rv
}

func (a_ AffineTransform) InitWithTransform(transform IAffineTransform) AffineTransform {
	rv := ffi.CallMethod[AffineTransform](a_, "initWithTransform:", transform)
	return rv
}

func (ac _AffineTransformClass) Alloc() AffineTransform {
	rv := ffi.CallMethod[AffineTransform](ac, "alloc")
	return rv
}

func (ac _AffineTransformClass) New() AffineTransform {
	rv := ffi.CallMethod[AffineTransform](ac, "new")
	rv.Autorelease()
	return rv
}

func NewAffineTransform() AffineTransform {
	return AffineTransformClass.New()
}

func (ac _AffineTransformClass) Transform() AffineTransform {
	rv := ffi.CallMethod[AffineTransform](ac, "transform")
	return rv
}

func (a_ AffineTransform) RotateByDegrees(angle float64) {
	ffi.CallMethod[ffi.Void](a_, "rotateByDegrees:", angle)
}

func (a_ AffineTransform) RotateByRadians(angle float64) {
	ffi.CallMethod[ffi.Void](a_, "rotateByRadians:", angle)
}

func (a_ AffineTransform) ScaleBy(scale float64) {
	ffi.CallMethod[ffi.Void](a_, "scaleBy:", scale)
}

func (a_ AffineTransform) ScaleXBy_YBy(scaleX float64, scaleY float64) {
	ffi.CallMethod[ffi.Void](a_, "scaleXBy:yBy:", scaleX, scaleY)
}

func (a_ AffineTransform) TranslateXBy_YBy(deltaX float64, deltaY float64) {
	ffi.CallMethod[ffi.Void](a_, "translateXBy:yBy:", deltaX, deltaY)
}

func (a_ AffineTransform) AppendTransform(transform IAffineTransform) {
	ffi.CallMethod[ffi.Void](a_, "appendTransform:", transform)
}

func (a_ AffineTransform) PrependTransform(transform IAffineTransform) {
	ffi.CallMethod[ffi.Void](a_, "prependTransform:", transform)
}

func (a_ AffineTransform) Invert() {
	ffi.CallMethod[ffi.Void](a_, "invert")
}

func (a_ AffineTransform) TransformPoint(aPoint Point) Point {
	rv := ffi.CallMethod[Point](a_, "transformPoint:", aPoint)
	return rv
}

func (a_ AffineTransform) TransformSize(aSize Size) Size {
	rv := ffi.CallMethod[Size](a_, "transformSize:", aSize)
	return rv
}

func (a_ AffineTransform) TransformStruct() AffineTransformStruct {
	rv := ffi.CallMethod[AffineTransformStruct](a_, "transformStruct")
	return rv
}

func (a_ AffineTransform) SetTransformStruct(value AffineTransformStruct) {
	ffi.CallMethod[ffi.Void](a_, "setTransformStruct:", value)
}
