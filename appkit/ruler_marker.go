// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	SetRepresentedObject0(value objc.IObject)
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
	rv := objc.CallMethod[RulerMarker](r_, "initWithRulerView:markerLocation:image:imageOrigin:", ruler, location, image, imageOrigin)
	return rv
}

func (rc _RulerMarkerClass) Alloc() RulerMarker {
	rv := objc.CallMethod[RulerMarker](rc, "alloc")
	return rv
}

func (rc _RulerMarkerClass) New() RulerMarker {
	rv := objc.CallMethod[RulerMarker](rc, "new")
	rv.Autorelease()
	return rv
}

func NewRulerMarker() RulerMarker {
	return RulerMarkerClass.New()
}

func (r_ RulerMarker) Init() RulerMarker {
	rv := objc.CallMethod[RulerMarker](r_, "init")
	return rv
}

func (r_ RulerMarker) DrawRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](r_, "drawRect:", rect)
}

func (r_ RulerMarker) TrackMouse_Adding(mouseDownEvent IEvent, isAdding bool) bool {
	rv := objc.CallMethod[bool](r_, "trackMouse:adding:", mouseDownEvent, isAdding)
	return rv
}

func (r_ RulerMarker) Ruler() RulerView {
	rv := objc.CallMethod[RulerView](r_, "ruler")
	return rv
}

func (r_ RulerMarker) Image() Image {
	rv := objc.CallMethod[Image](r_, "image")
	return rv
}

func (r_ RulerMarker) SetImage(value IImage) {
	objc.CallMethod[objc.Void](r_, "setImage:", value)
}

func (r_ RulerMarker) ImageOrigin() foundation.Point {
	rv := objc.CallMethod[foundation.Point](r_, "imageOrigin")
	return rv
}

func (r_ RulerMarker) SetImageOrigin(value foundation.Point) {
	objc.CallMethod[objc.Void](r_, "setImageOrigin:", value)
}

func (r_ RulerMarker) ImageRectInRuler() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](r_, "imageRectInRuler")
	return rv
}

func (r_ RulerMarker) ThicknessRequiredInRuler() float64 {
	rv := objc.CallMethod[float64](r_, "thicknessRequiredInRuler")
	return rv
}

func (r_ RulerMarker) IsMovable() bool {
	rv := objc.CallMethod[bool](r_, "isMovable")
	return rv
}

func (r_ RulerMarker) SetMovable(value bool) {
	objc.CallMethod[objc.Void](r_, "setMovable:", value)
}

func (r_ RulerMarker) IsRemovable() bool {
	rv := objc.CallMethod[bool](r_, "isRemovable")
	return rv
}

func (r_ RulerMarker) SetRemovable(value bool) {
	objc.CallMethod[objc.Void](r_, "setRemovable:", value)
}

func (r_ RulerMarker) MarkerLocation() float64 {
	rv := objc.CallMethod[float64](r_, "markerLocation")
	return rv
}

func (r_ RulerMarker) SetMarkerLocation(value float64) {
	objc.CallMethod[objc.Void](r_, "setMarkerLocation:", value)
}

func (r_ RulerMarker) RepresentedObject() CopyingWrapper {
	rv := objc.CallMethod[CopyingWrapper](r_, "representedObject")
	return rv
}

func (r_ RulerMarker) SetRepresentedObject(value Copying) {
	po := objc.CreateProtocol("NSCopying", value)
	defer po.Release()
	objc.CallMethod[objc.Void](r_, "setRepresentedObject:", po)
}

func (r_ RulerMarker) SetRepresentedObject0(value objc.IObject) {
	objc.CallMethod[objc.Void](r_, "setRepresentedObject:", value)
}

func (r_ RulerMarker) IsDragging() bool {
	rv := objc.CallMethod[bool](r_, "isDragging")
	return rv
}
