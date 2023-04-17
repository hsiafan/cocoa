// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var RulerMarkerClass = _RulerMarkerClass{objc.GetClass("NSRulerMarker")}

type _RulerMarkerClass struct {
	objc.Class
}

type IRulerMarker interface {
	objc.IObject
	DrawRect(rect foundation.Rect)
	TrackMouse_Adding(mouseDownEvent IEvent, isAdding bool) bool
	Ruler() RulerView
	Image() Image
	SetImage(value IImage)
	ImageOrigin() foundation.Point
	SetImageOrigin(value foundation.Point)
	ImageRectInRuler() foundation.Rect
	ThicknessRequiredInRuler() float64
	IsMovable() bool
	SetMovable(value bool)
	IsRemovable() bool
	SetRemovable(value bool)
	MarkerLocation() float64
	SetMarkerLocation(value float64)
	RepresentedObject() CopyingWrapper
	SetRepresentedObject(value Copying)
	IsDragging() bool
}

type RulerMarker struct {
	objc.Object
}

func MakeRulerMarker(ptr unsafe.Pointer) RulerMarker {
	return RulerMarker{
		Object: objc.MakeObject(ptr),
	}
}

func (r_ RulerMarker) InitWithRulerView_MarkerLocation_Image_ImageOrigin(ruler IRulerView, location float64, image IImage, imageOrigin foundation.Point) RulerMarker {
	rv := ffi.CallMethod[RulerMarker](r_, "initWithRulerView:markerLocation:image:imageOrigin:", ruler, location, image, imageOrigin)
	return rv
}

func (rc _RulerMarkerClass) Alloc() RulerMarker {
	rv := ffi.CallMethod[RulerMarker](rc, "alloc")
	return rv
}

func (rc _RulerMarkerClass) New() RulerMarker {
	rv := ffi.CallMethod[RulerMarker](rc, "new")
	rv.Autorelease()
	return rv
}

func NewRulerMarker() RulerMarker {
	return RulerMarkerClass.New()
}

func (r_ RulerMarker) Init() RulerMarker {
	rv := ffi.CallMethod[RulerMarker](r_, "init")
	return rv
}

func (r_ RulerMarker) DrawRect(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](r_, "drawRect:", rect)
}

func (r_ RulerMarker) TrackMouse_Adding(mouseDownEvent IEvent, isAdding bool) bool {
	rv := ffi.CallMethod[bool](r_, "trackMouse:adding:", mouseDownEvent, isAdding)
	return rv
}

func (r_ RulerMarker) Ruler() RulerView {
	rv := ffi.CallMethod[RulerView](r_, "ruler")
	return rv
}

func (r_ RulerMarker) Image() Image {
	rv := ffi.CallMethod[Image](r_, "image")
	return rv
}

func (r_ RulerMarker) SetImage(value IImage) {
	ffi.CallMethod[ffi.Void](r_, "setImage:", value)
}

func (r_ RulerMarker) ImageOrigin() foundation.Point {
	rv := ffi.CallMethod[foundation.Point](r_, "imageOrigin")
	return rv
}

func (r_ RulerMarker) SetImageOrigin(value foundation.Point) {
	ffi.CallMethod[ffi.Void](r_, "setImageOrigin:", value)
}

func (r_ RulerMarker) ImageRectInRuler() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](r_, "imageRectInRuler")
	return rv
}

func (r_ RulerMarker) ThicknessRequiredInRuler() float64 {
	rv := ffi.CallMethod[float64](r_, "thicknessRequiredInRuler")
	return rv
}

func (r_ RulerMarker) IsMovable() bool {
	rv := ffi.CallMethod[bool](r_, "isMovable")
	return rv
}

func (r_ RulerMarker) SetMovable(value bool) {
	ffi.CallMethod[ffi.Void](r_, "setMovable:", value)
}

func (r_ RulerMarker) IsRemovable() bool {
	rv := ffi.CallMethod[bool](r_, "isRemovable")
	return rv
}

func (r_ RulerMarker) SetRemovable(value bool) {
	ffi.CallMethod[ffi.Void](r_, "setRemovable:", value)
}

func (r_ RulerMarker) MarkerLocation() float64 {
	rv := ffi.CallMethod[float64](r_, "markerLocation")
	return rv
}

func (r_ RulerMarker) SetMarkerLocation(value float64) {
	ffi.CallMethod[ffi.Void](r_, "setMarkerLocation:", value)
}

func (r_ RulerMarker) RepresentedObject() CopyingWrapper {
	rv := ffi.CallMethod[CopyingWrapper](r_, "representedObject")
	return rv
}

func (r_ RulerMarker) SetRepresentedObject(value Copying) {
	po := ffi.CreateProtocol("NSCopying", value)
	defer po.Release()
	ffi.CallMethod[ffi.Void](r_, "setRepresentedObject:", po)
}

func (r_ RulerMarker) IsDragging() bool {
	rv := ffi.CallMethod[bool](r_, "isDragging")
	return rv
}
