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
	RepresentedObject() objc.Object
	SetRepresentedObject(value objc.IObject)
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
	rv := objc.CallMethod[RulerMarker](r_, objc.SEL("initWithRulerView:markerLocation:image:imageOrigin:"), objc.ExtractPtr(ruler), location, objc.ExtractPtr(image), imageOrigin)
	return rv
}

func (rc _RulerMarkerClass) Alloc() RulerMarker {
	rv := objc.CallMethod[RulerMarker](rc, objc.SEL("alloc"))
	return rv
}

func (rc _RulerMarkerClass) New() RulerMarker {
	rv := objc.CallMethod[RulerMarker](rc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewRulerMarker() RulerMarker {
	return RulerMarkerClass.New()
}

func (r_ RulerMarker) Init() RulerMarker {
	rv := objc.CallMethod[RulerMarker](r_, objc.SEL("init"))
	return rv
}

func (r_ RulerMarker) DrawRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](r_, objc.SEL("drawRect:"), rect)
}

func (r_ RulerMarker) TrackMouse_Adding(mouseDownEvent IEvent, isAdding bool) bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("trackMouse:adding:"), objc.ExtractPtr(mouseDownEvent), isAdding)
	return rv
}

// weak property
func (r_ RulerMarker) Ruler() RulerView {
	rv := objc.CallMethod[RulerView](r_, objc.SEL("ruler"))
	return rv
}

func (r_ RulerMarker) Image() Image {
	rv := objc.CallMethod[Image](r_, objc.SEL("image"))
	return rv
}

func (r_ RulerMarker) SetImage(value IImage) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setImage:"), objc.ExtractPtr(value))
}

func (r_ RulerMarker) ImageOrigin() foundation.Point {
	rv := objc.CallMethod[foundation.Point](r_, objc.SEL("imageOrigin"))
	return rv
}

func (r_ RulerMarker) SetImageOrigin(value foundation.Point) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setImageOrigin:"), value)
}

func (r_ RulerMarker) ImageRectInRuler() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](r_, objc.SEL("imageRectInRuler"))
	return rv
}

func (r_ RulerMarker) ThicknessRequiredInRuler() float64 {
	rv := objc.CallMethod[float64](r_, objc.SEL("thicknessRequiredInRuler"))
	return rv
}

func (r_ RulerMarker) IsMovable() bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("isMovable"))
	return rv
}

func (r_ RulerMarker) SetMovable(value bool) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setMovable:"), value)
}

func (r_ RulerMarker) IsRemovable() bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("isRemovable"))
	return rv
}

func (r_ RulerMarker) SetRemovable(value bool) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setRemovable:"), value)
}

func (r_ RulerMarker) MarkerLocation() float64 {
	rv := objc.CallMethod[float64](r_, objc.SEL("markerLocation"))
	return rv
}

func (r_ RulerMarker) SetMarkerLocation(value float64) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setMarkerLocation:"), value)
}

func (r_ RulerMarker) RepresentedObject() objc.Object {
	rv := objc.CallMethod[objc.Object](r_, objc.SEL("representedObject"))
	return rv
}

func (r_ RulerMarker) SetRepresentedObject(value objc.IObject) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setRepresentedObject:"), objc.ExtractPtr(value))
}

func (r_ RulerMarker) IsDragging() bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("isDragging"))
	return rv
}
