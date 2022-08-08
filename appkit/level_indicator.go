// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var LevelIndicatorClass = _LevelIndicatorClass{objc.GetClass("NSLevelIndicator")}

type _LevelIndicatorClass struct {
	objc.Class
}

type ILevelIndicator interface {
	IControl
	TickMarkValueAtIndex(index int) float64
	RectOfTickMarkAtIndex(index int) foundation.Rect
	MinValue() float64
	SetMinValue(value float64)
	MaxValue() float64
	SetMaxValue(value float64)
	WarningValue() float64
	SetWarningValue(value float64)
	CriticalValue() float64
	SetCriticalValue(value float64)
	TickMarkPosition() TickMarkPosition
	SetTickMarkPosition(value TickMarkPosition)
	NumberOfTickMarks() int
	SetNumberOfTickMarks(value int)
	NumberOfMajorTickMarks() int
	SetNumberOfMajorTickMarks(value int)
	LevelIndicatorStyle() LevelIndicatorStyle
	SetLevelIndicatorStyle(value LevelIndicatorStyle)
	RatingImage() Image
	SetRatingImage(value IImage)
	DrawsTieredCapacityLevels() bool
	SetDrawsTieredCapacityLevels(value bool)
	FillColor() Color
	SetFillColor(value IColor)
	WarningFillColor() Color
	SetWarningFillColor(value IColor)
	CriticalFillColor() Color
	SetCriticalFillColor(value IColor)
	RatingPlaceholderImage() Image
	SetRatingPlaceholderImage(value IImage)
	PlaceholderVisibility() LevelIndicatorPlaceholderVisibility
	SetPlaceholderVisibility(value LevelIndicatorPlaceholderVisibility)
	IsEditable() bool
	SetEditable(value bool)
}

type LevelIndicator struct {
	Control
}

func MakeLevelIndicator(ptr unsafe.Pointer) LevelIndicator {
	return LevelIndicator{
		Control: MakeControl(ptr),
	}
}

func (l_ LevelIndicator) InitWithFrame(frameRect foundation.Rect) LevelIndicator {
	rv := ffi.CallMethod[LevelIndicator](l_, "initWithFrame:", frameRect)
	return rv
}

func (l_ LevelIndicator) Init() LevelIndicator {
	rv := ffi.CallMethod[LevelIndicator](l_, "init")
	return rv
}

func (lc _LevelIndicatorClass) Alloc() LevelIndicator {
	rv := ffi.CallMethod[LevelIndicator](lc, "alloc")
	return rv
}

func (lc _LevelIndicatorClass) New() LevelIndicator {
	rv := ffi.CallMethod[LevelIndicator](lc, "new")
	rv.Autorelease()
	return rv
}

func NewLevelIndicator() LevelIndicator {
	return LevelIndicatorClass.New()
}

func (l_ LevelIndicator) TickMarkValueAtIndex(index int) float64 {
	rv := ffi.CallMethod[float64](l_, "tickMarkValueAtIndex:", index)
	return rv
}

func (l_ LevelIndicator) RectOfTickMarkAtIndex(index int) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](l_, "rectOfTickMarkAtIndex:", index)
	return rv
}

func (l_ LevelIndicator) MinValue() float64 {
	rv := ffi.CallMethod[float64](l_, "minValue")
	return rv
}

func (l_ LevelIndicator) SetMinValue(value float64) {
	ffi.CallMethod[ffi.Void](l_, "setMinValue:", value)
}

func (l_ LevelIndicator) MaxValue() float64 {
	rv := ffi.CallMethod[float64](l_, "maxValue")
	return rv
}

func (l_ LevelIndicator) SetMaxValue(value float64) {
	ffi.CallMethod[ffi.Void](l_, "setMaxValue:", value)
}

func (l_ LevelIndicator) WarningValue() float64 {
	rv := ffi.CallMethod[float64](l_, "warningValue")
	return rv
}

func (l_ LevelIndicator) SetWarningValue(value float64) {
	ffi.CallMethod[ffi.Void](l_, "setWarningValue:", value)
}

func (l_ LevelIndicator) CriticalValue() float64 {
	rv := ffi.CallMethod[float64](l_, "criticalValue")
	return rv
}

func (l_ LevelIndicator) SetCriticalValue(value float64) {
	ffi.CallMethod[ffi.Void](l_, "setCriticalValue:", value)
}

func (l_ LevelIndicator) TickMarkPosition() TickMarkPosition {
	rv := ffi.CallMethod[TickMarkPosition](l_, "tickMarkPosition")
	return rv
}

func (l_ LevelIndicator) SetTickMarkPosition(value TickMarkPosition) {
	ffi.CallMethod[ffi.Void](l_, "setTickMarkPosition:", value)
}

func (l_ LevelIndicator) NumberOfTickMarks() int {
	rv := ffi.CallMethod[int](l_, "numberOfTickMarks")
	return rv
}

func (l_ LevelIndicator) SetNumberOfTickMarks(value int) {
	ffi.CallMethod[ffi.Void](l_, "setNumberOfTickMarks:", value)
}

func (l_ LevelIndicator) NumberOfMajorTickMarks() int {
	rv := ffi.CallMethod[int](l_, "numberOfMajorTickMarks")
	return rv
}

func (l_ LevelIndicator) SetNumberOfMajorTickMarks(value int) {
	ffi.CallMethod[ffi.Void](l_, "setNumberOfMajorTickMarks:", value)
}

func (l_ LevelIndicator) LevelIndicatorStyle() LevelIndicatorStyle {
	rv := ffi.CallMethod[LevelIndicatorStyle](l_, "levelIndicatorStyle")
	return rv
}

func (l_ LevelIndicator) SetLevelIndicatorStyle(value LevelIndicatorStyle) {
	ffi.CallMethod[ffi.Void](l_, "setLevelIndicatorStyle:", value)
}

func (l_ LevelIndicator) RatingImage() Image {
	rv := ffi.CallMethod[Image](l_, "ratingImage")
	return rv
}

func (l_ LevelIndicator) SetRatingImage(value IImage) {
	ffi.CallMethod[ffi.Void](l_, "setRatingImage:", value)
}

func (l_ LevelIndicator) DrawsTieredCapacityLevels() bool {
	rv := ffi.CallMethod[bool](l_, "drawsTieredCapacityLevels")
	return rv
}

func (l_ LevelIndicator) SetDrawsTieredCapacityLevels(value bool) {
	ffi.CallMethod[ffi.Void](l_, "setDrawsTieredCapacityLevels:", value)
}

func (l_ LevelIndicator) FillColor() Color {
	rv := ffi.CallMethod[Color](l_, "fillColor")
	return rv
}

func (l_ LevelIndicator) SetFillColor(value IColor) {
	ffi.CallMethod[ffi.Void](l_, "setFillColor:", value)
}

func (l_ LevelIndicator) WarningFillColor() Color {
	rv := ffi.CallMethod[Color](l_, "warningFillColor")
	return rv
}

func (l_ LevelIndicator) SetWarningFillColor(value IColor) {
	ffi.CallMethod[ffi.Void](l_, "setWarningFillColor:", value)
}

func (l_ LevelIndicator) CriticalFillColor() Color {
	rv := ffi.CallMethod[Color](l_, "criticalFillColor")
	return rv
}

func (l_ LevelIndicator) SetCriticalFillColor(value IColor) {
	ffi.CallMethod[ffi.Void](l_, "setCriticalFillColor:", value)
}

func (l_ LevelIndicator) RatingPlaceholderImage() Image {
	rv := ffi.CallMethod[Image](l_, "ratingPlaceholderImage")
	return rv
}

func (l_ LevelIndicator) SetRatingPlaceholderImage(value IImage) {
	ffi.CallMethod[ffi.Void](l_, "setRatingPlaceholderImage:", value)
}

func (l_ LevelIndicator) PlaceholderVisibility() LevelIndicatorPlaceholderVisibility {
	rv := ffi.CallMethod[LevelIndicatorPlaceholderVisibility](l_, "placeholderVisibility")
	return rv
}

func (l_ LevelIndicator) SetPlaceholderVisibility(value LevelIndicatorPlaceholderVisibility) {
	ffi.CallMethod[ffi.Void](l_, "setPlaceholderVisibility:", value)
}

func (l_ LevelIndicator) IsEditable() bool {
	rv := ffi.CallMethod[bool](l_, "isEditable")
	return rv
}

func (l_ LevelIndicator) SetEditable(value bool) {
	ffi.CallMethod[ffi.Void](l_, "setEditable:", value)
}
