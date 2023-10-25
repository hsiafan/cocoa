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
	rv := objc.CallMethod[RulerView](r_, objc.SEL("initWithScrollView:orientation:"), objc.ExtractPtr(scrollView), orientation)
	return rv
}

func (r_ RulerView) InitWithFrame(frameRect foundation.Rect) RulerView {
	rv := objc.CallMethod[RulerView](r_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (r_ RulerView) Init() RulerView {
	rv := objc.CallMethod[RulerView](r_, objc.SEL("init"))
	return rv
}

func (rc _RulerViewClass) Alloc() RulerView {
	rv := objc.CallMethod[RulerView](rc, objc.SEL("alloc"))
	return rv
}

func (rc _RulerViewClass) New() RulerView {
	rv := objc.CallMethod[RulerView](rc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewRulerView() RulerView {
	return RulerViewClass.New()
}

func (rc _RulerViewClass) RegisterUnitWithName_Abbreviation_UnitToPointsConversionFactor_StepUpCycle_StepDownCycle(unitName RulerViewUnitName, abbreviation string, conversionFactor float64, stepUpCycle []foundation.INumber, stepDownCycle []foundation.INumber) {
	objc.CallMethod[objc.Void](rc, objc.SEL("registerUnitWithName:abbreviation:unitToPointsConversionFactor:stepUpCycle:stepDownCycle:"), unitName, abbreviation, conversionFactor, stepUpCycle, stepDownCycle)
}

func (r_ RulerView) AddMarker(marker IRulerMarker) {
	objc.CallMethod[objc.Void](r_, objc.SEL("addMarker:"), objc.ExtractPtr(marker))
}

func (r_ RulerView) RemoveMarker(marker IRulerMarker) {
	objc.CallMethod[objc.Void](r_, objc.SEL("removeMarker:"), objc.ExtractPtr(marker))
}

func (r_ RulerView) TrackMarker_WithMouseEvent(marker IRulerMarker, event IEvent) bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("trackMarker:withMouseEvent:"), objc.ExtractPtr(marker), objc.ExtractPtr(event))
	return rv
}

func (r_ RulerView) MoveRulerlineFromLocation_ToLocation(oldLocation float64, newLocation float64) {
	objc.CallMethod[objc.Void](r_, objc.SEL("moveRulerlineFromLocation:toLocation:"), oldLocation, newLocation)
}

func (r_ RulerView) DrawHashMarksAndLabelsInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](r_, objc.SEL("drawHashMarksAndLabelsInRect:"), rect)
}

func (r_ RulerView) DrawMarkersInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](r_, objc.SEL("drawMarkersInRect:"), rect)
}

func (r_ RulerView) InvalidateHashMarks() {
	objc.CallMethod[objc.Void](r_, objc.SEL("invalidateHashMarks"))
}

func (r_ RulerView) MeasurementUnits() RulerViewUnitName {
	rv := objc.CallMethod[RulerViewUnitName](r_, objc.SEL("measurementUnits"))
	return rv
}

func (r_ RulerView) SetMeasurementUnits(value RulerViewUnitName) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setMeasurementUnits:"), value)
}

// weak property
func (r_ RulerView) ClientView() View {
	rv := objc.CallMethod[View](r_, objc.SEL("clientView"))
	return rv
}

// weak property
func (r_ RulerView) SetClientView(value IView) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setClientView:"), objc.ExtractPtr(value))
}

func (r_ RulerView) AccessoryView() View {
	rv := objc.CallMethod[View](r_, objc.SEL("accessoryView"))
	return rv
}

func (r_ RulerView) SetAccessoryView(value IView) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setAccessoryView:"), objc.ExtractPtr(value))
}

func (r_ RulerView) OriginOffset() float64 {
	rv := objc.CallMethod[float64](r_, objc.SEL("originOffset"))
	return rv
}

func (r_ RulerView) SetOriginOffset(value float64) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setOriginOffset:"), value)
}

func (r_ RulerView) Markers() []RulerMarker {
	rv := objc.CallMethod[[]RulerMarker](r_, objc.SEL("markers"))
	return rv
}

func (r_ RulerView) SetMarkers(value []IRulerMarker) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setMarkers:"), value)
}

// weak property
func (r_ RulerView) ScrollView() ScrollView {
	rv := objc.CallMethod[ScrollView](r_, objc.SEL("scrollView"))
	return rv
}

// weak property
func (r_ RulerView) SetScrollView(value IScrollView) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setScrollView:"), objc.ExtractPtr(value))
}

func (r_ RulerView) Orientation() RulerOrientation {
	rv := objc.CallMethod[RulerOrientation](r_, objc.SEL("orientation"))
	return rv
}

func (r_ RulerView) SetOrientation(value RulerOrientation) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setOrientation:"), value)
}

func (r_ RulerView) ReservedThicknessForAccessoryView() float64 {
	rv := objc.CallMethod[float64](r_, objc.SEL("reservedThicknessForAccessoryView"))
	return rv
}

func (r_ RulerView) SetReservedThicknessForAccessoryView(value float64) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setReservedThicknessForAccessoryView:"), value)
}

func (r_ RulerView) ReservedThicknessForMarkers() float64 {
	rv := objc.CallMethod[float64](r_, objc.SEL("reservedThicknessForMarkers"))
	return rv
}

func (r_ RulerView) SetReservedThicknessForMarkers(value float64) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setReservedThicknessForMarkers:"), value)
}

func (r_ RulerView) RuleThickness() float64 {
	rv := objc.CallMethod[float64](r_, objc.SEL("ruleThickness"))
	return rv
}

func (r_ RulerView) SetRuleThickness(value float64) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setRuleThickness:"), value)
}

func (r_ RulerView) RequiredThickness() float64 {
	rv := objc.CallMethod[float64](r_, objc.SEL("requiredThickness"))
	return rv
}

func (r_ RulerView) BaselineLocation() float64 {
	rv := objc.CallMethod[float64](r_, objc.SEL("baselineLocation"))
	return rv
}
