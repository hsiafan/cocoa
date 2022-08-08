// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var RulerViewClass = _RulerViewClass{objc.GetClass("NSRulerView")}

type _RulerViewClass struct {
	objc.Class
}

type IRulerView interface {
	IView
	AddMarker(marker IRulerMarker)
	RemoveMarker(marker IRulerMarker)
	TrackMarker_WithMouseEvent(marker IRulerMarker, event IEvent) bool
	MoveRulerlineFromLocation_ToLocation(oldLocation float64, newLocation float64)
	DrawHashMarksAndLabelsInRect(rect foundation.Rect)
	DrawMarkersInRect(rect foundation.Rect)
	InvalidateHashMarks()
	MeasurementUnits() RulerViewUnitName
	SetMeasurementUnits(value RulerViewUnitName)
	ClientView() View
	SetClientView(value IView)
	AccessoryView() View
	SetAccessoryView(value IView)
	OriginOffset() float64
	SetOriginOffset(value float64)
	Markers() []RulerMarker
	SetMarkers(value []IRulerMarker)
	ScrollView() ScrollView
	SetScrollView(value IScrollView)
	Orientation() RulerOrientation
	SetOrientation(value RulerOrientation)
	ReservedThicknessForAccessoryView() float64
	SetReservedThicknessForAccessoryView(value float64)
	ReservedThicknessForMarkers() float64
	SetReservedThicknessForMarkers(value float64)
	RuleThickness() float64
	SetRuleThickness(value float64)
	RequiredThickness() float64
	BaselineLocation() float64
}

type RulerView struct {
	View
}

func MakeRulerView(ptr unsafe.Pointer) RulerView {
	return RulerView{
		View: MakeView(ptr),
	}
}

func (r_ RulerView) InitWithScrollView_Orientation(scrollView IScrollView, orientation RulerOrientation) RulerView {
	rv := ffi.CallMethod[RulerView](r_, "initWithScrollView:orientation:", scrollView, orientation)
	return rv
}

func (r_ RulerView) InitWithFrame(frameRect foundation.Rect) RulerView {
	rv := ffi.CallMethod[RulerView](r_, "initWithFrame:", frameRect)
	return rv
}

func (r_ RulerView) Init() RulerView {
	rv := ffi.CallMethod[RulerView](r_, "init")
	return rv
}

func (rc _RulerViewClass) Alloc() RulerView {
	rv := ffi.CallMethod[RulerView](rc, "alloc")
	return rv
}

func (rc _RulerViewClass) New() RulerView {
	rv := ffi.CallMethod[RulerView](rc, "new")
	rv.Autorelease()
	return rv
}

func NewRulerView() RulerView {
	return RulerViewClass.New()
}

func (rc _RulerViewClass) RegisterUnitWithName_Abbreviation_UnitToPointsConversionFactor_StepUpCycle_StepDownCycle(unitName RulerViewUnitName, abbreviation string, conversionFactor float64, stepUpCycle []foundation.INumber, stepDownCycle []foundation.INumber) {
	ffi.CallMethod[ffi.Void](rc, "registerUnitWithName:abbreviation:unitToPointsConversionFactor:stepUpCycle:stepDownCycle:", unitName, abbreviation, conversionFactor, stepUpCycle, stepDownCycle)
}

func (r_ RulerView) AddMarker(marker IRulerMarker) {
	ffi.CallMethod[ffi.Void](r_, "addMarker:", marker)
}

func (r_ RulerView) RemoveMarker(marker IRulerMarker) {
	ffi.CallMethod[ffi.Void](r_, "removeMarker:", marker)
}

func (r_ RulerView) TrackMarker_WithMouseEvent(marker IRulerMarker, event IEvent) bool {
	rv := ffi.CallMethod[bool](r_, "trackMarker:withMouseEvent:", marker, event)
	return rv
}

func (r_ RulerView) MoveRulerlineFromLocation_ToLocation(oldLocation float64, newLocation float64) {
	ffi.CallMethod[ffi.Void](r_, "moveRulerlineFromLocation:toLocation:", oldLocation, newLocation)
}

func (r_ RulerView) DrawHashMarksAndLabelsInRect(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](r_, "drawHashMarksAndLabelsInRect:", rect)
}

func (r_ RulerView) DrawMarkersInRect(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](r_, "drawMarkersInRect:", rect)
}

func (r_ RulerView) InvalidateHashMarks() {
	ffi.CallMethod[ffi.Void](r_, "invalidateHashMarks")
}

func (r_ RulerView) MeasurementUnits() RulerViewUnitName {
	rv := ffi.CallMethod[RulerViewUnitName](r_, "measurementUnits")
	return rv
}

func (r_ RulerView) SetMeasurementUnits(value RulerViewUnitName) {
	ffi.CallMethod[ffi.Void](r_, "setMeasurementUnits:", value)
}

func (r_ RulerView) ClientView() View {
	rv := ffi.CallMethod[View](r_, "clientView")
	return rv
}

func (r_ RulerView) SetClientView(value IView) {
	ffi.CallMethod[ffi.Void](r_, "setClientView:", value)
}

func (r_ RulerView) AccessoryView() View {
	rv := ffi.CallMethod[View](r_, "accessoryView")
	return rv
}

func (r_ RulerView) SetAccessoryView(value IView) {
	ffi.CallMethod[ffi.Void](r_, "setAccessoryView:", value)
}

func (r_ RulerView) OriginOffset() float64 {
	rv := ffi.CallMethod[float64](r_, "originOffset")
	return rv
}

func (r_ RulerView) SetOriginOffset(value float64) {
	ffi.CallMethod[ffi.Void](r_, "setOriginOffset:", value)
}

func (r_ RulerView) Markers() []RulerMarker {
	rv := ffi.CallMethod[[]RulerMarker](r_, "markers")
	return rv
}

func (r_ RulerView) SetMarkers(value []IRulerMarker) {
	ffi.CallMethod[ffi.Void](r_, "setMarkers:", value)
}

func (r_ RulerView) ScrollView() ScrollView {
	rv := ffi.CallMethod[ScrollView](r_, "scrollView")
	return rv
}

func (r_ RulerView) SetScrollView(value IScrollView) {
	ffi.CallMethod[ffi.Void](r_, "setScrollView:", value)
}

func (r_ RulerView) Orientation() RulerOrientation {
	rv := ffi.CallMethod[RulerOrientation](r_, "orientation")
	return rv
}

func (r_ RulerView) SetOrientation(value RulerOrientation) {
	ffi.CallMethod[ffi.Void](r_, "setOrientation:", value)
}

func (r_ RulerView) ReservedThicknessForAccessoryView() float64 {
	rv := ffi.CallMethod[float64](r_, "reservedThicknessForAccessoryView")
	return rv
}

func (r_ RulerView) SetReservedThicknessForAccessoryView(value float64) {
	ffi.CallMethod[ffi.Void](r_, "setReservedThicknessForAccessoryView:", value)
}

func (r_ RulerView) ReservedThicknessForMarkers() float64 {
	rv := ffi.CallMethod[float64](r_, "reservedThicknessForMarkers")
	return rv
}

func (r_ RulerView) SetReservedThicknessForMarkers(value float64) {
	ffi.CallMethod[ffi.Void](r_, "setReservedThicknessForMarkers:", value)
}

func (r_ RulerView) RuleThickness() float64 {
	rv := ffi.CallMethod[float64](r_, "ruleThickness")
	return rv
}

func (r_ RulerView) SetRuleThickness(value float64) {
	ffi.CallMethod[ffi.Void](r_, "setRuleThickness:", value)
}

func (r_ RulerView) RequiredThickness() float64 {
	rv := ffi.CallMethod[float64](r_, "requiredThickness")
	return rv
}

func (r_ RulerView) BaselineLocation() float64 {
	rv := ffi.CallMethod[float64](r_, "baselineLocation")
	return rv
}

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

func (r_ RulerMarker) Init() RulerMarker {
	rv := ffi.CallMethod[RulerMarker](r_, "init")
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

func (r_ RulerMarker) IsDragging() bool {
	rv := ffi.CallMethod[bool](r_, "isDragging")
	return rv
}
