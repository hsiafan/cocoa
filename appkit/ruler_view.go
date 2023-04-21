// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[RulerView](r_, "initWithScrollView:orientation:", scrollView, orientation)
	return rv
}

func (r_ RulerView) InitWithFrame(frameRect foundation.Rect) RulerView {
	rv := objc.CallMethod[RulerView](r_, "initWithFrame:", frameRect)
	return rv
}

func (r_ RulerView) Init() RulerView {
	rv := objc.CallMethod[RulerView](r_, "init")
	return rv
}

func (rc _RulerViewClass) Alloc() RulerView {
	rv := objc.CallMethod[RulerView](rc, "alloc")
	return rv
}

func (rc _RulerViewClass) New() RulerView {
	rv := objc.CallMethod[RulerView](rc, "new")
	rv.Autorelease()
	return rv
}

func NewRulerView() RulerView {
	return RulerViewClass.New()
}

func (rc _RulerViewClass) RegisterUnitWithName_Abbreviation_UnitToPointsConversionFactor_StepUpCycle_StepDownCycle(unitName RulerViewUnitName, abbreviation string, conversionFactor float64, stepUpCycle []foundation.INumber, stepDownCycle []foundation.INumber) {
	objc.CallMethod[objc.Void](rc, "registerUnitWithName:abbreviation:unitToPointsConversionFactor:stepUpCycle:stepDownCycle:", unitName, abbreviation, conversionFactor, stepUpCycle, stepDownCycle)
}

func (r_ RulerView) AddMarker(marker IRulerMarker) {
	objc.CallMethod[objc.Void](r_, "addMarker:", marker)
}

func (r_ RulerView) RemoveMarker(marker IRulerMarker) {
	objc.CallMethod[objc.Void](r_, "removeMarker:", marker)
}

func (r_ RulerView) TrackMarker_WithMouseEvent(marker IRulerMarker, event IEvent) bool {
	rv := objc.CallMethod[bool](r_, "trackMarker:withMouseEvent:", marker, event)
	return rv
}

func (r_ RulerView) MoveRulerlineFromLocation_ToLocation(oldLocation float64, newLocation float64) {
	objc.CallMethod[objc.Void](r_, "moveRulerlineFromLocation:toLocation:", oldLocation, newLocation)
}

func (r_ RulerView) DrawHashMarksAndLabelsInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](r_, "drawHashMarksAndLabelsInRect:", rect)
}

func (r_ RulerView) DrawMarkersInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](r_, "drawMarkersInRect:", rect)
}

func (r_ RulerView) InvalidateHashMarks() {
	objc.CallMethod[objc.Void](r_, "invalidateHashMarks")
}

func (r_ RulerView) MeasurementUnits() RulerViewUnitName {
	rv := objc.CallMethod[RulerViewUnitName](r_, "measurementUnits")
	return rv
}

func (r_ RulerView) SetMeasurementUnits(value RulerViewUnitName) {
	objc.CallMethod[objc.Void](r_, "setMeasurementUnits:", value)
}

func (r_ RulerView) ClientView() View {
	rv := objc.CallMethod[View](r_, "clientView")
	return rv
}

func (r_ RulerView) SetClientView(value IView) {
	objc.CallMethod[objc.Void](r_, "setClientView:", value)
}

func (r_ RulerView) AccessoryView() View {
	rv := objc.CallMethod[View](r_, "accessoryView")
	return rv
}

func (r_ RulerView) SetAccessoryView(value IView) {
	objc.CallMethod[objc.Void](r_, "setAccessoryView:", value)
}

func (r_ RulerView) OriginOffset() float64 {
	rv := objc.CallMethod[float64](r_, "originOffset")
	return rv
}

func (r_ RulerView) SetOriginOffset(value float64) {
	objc.CallMethod[objc.Void](r_, "setOriginOffset:", value)
}

func (r_ RulerView) Markers() []RulerMarker {
	rv := objc.CallMethod[[]RulerMarker](r_, "markers")
	return rv
}

func (r_ RulerView) SetMarkers(value []IRulerMarker) {
	objc.CallMethod[objc.Void](r_, "setMarkers:", value)
}

func (r_ RulerView) ScrollView() ScrollView {
	rv := objc.CallMethod[ScrollView](r_, "scrollView")
	return rv
}

func (r_ RulerView) SetScrollView(value IScrollView) {
	objc.CallMethod[objc.Void](r_, "setScrollView:", value)
}

func (r_ RulerView) Orientation() RulerOrientation {
	rv := objc.CallMethod[RulerOrientation](r_, "orientation")
	return rv
}

func (r_ RulerView) SetOrientation(value RulerOrientation) {
	objc.CallMethod[objc.Void](r_, "setOrientation:", value)
}

func (r_ RulerView) ReservedThicknessForAccessoryView() float64 {
	rv := objc.CallMethod[float64](r_, "reservedThicknessForAccessoryView")
	return rv
}

func (r_ RulerView) SetReservedThicknessForAccessoryView(value float64) {
	objc.CallMethod[objc.Void](r_, "setReservedThicknessForAccessoryView:", value)
}

func (r_ RulerView) ReservedThicknessForMarkers() float64 {
	rv := objc.CallMethod[float64](r_, "reservedThicknessForMarkers")
	return rv
}

func (r_ RulerView) SetReservedThicknessForMarkers(value float64) {
	objc.CallMethod[objc.Void](r_, "setReservedThicknessForMarkers:", value)
}

func (r_ RulerView) RuleThickness() float64 {
	rv := objc.CallMethod[float64](r_, "ruleThickness")
	return rv
}

func (r_ RulerView) SetRuleThickness(value float64) {
	objc.CallMethod[objc.Void](r_, "setRuleThickness:", value)
}

func (r_ RulerView) RequiredThickness() float64 {
	rv := objc.CallMethod[float64](r_, "requiredThickness")
	return rv
}

func (r_ RulerView) BaselineLocation() float64 {
	rv := objc.CallMethod[float64](r_, "baselineLocation")
	return rv
}
