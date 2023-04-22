// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var SliderClass = _SliderClass{objc.GetClass("NSSlider")}

type _SliderClass struct {
	objc.Class
}

type ISlider interface {
	IControl
	// deprecated
	SetKnobThickness(thickness float64)
	ClosestTickMarkValueToValue(value float64) float64
	IndexOfTickMarkAtPoint(point foundation.Point) int
	RectOfTickMarkAtIndex(index int) foundation.Rect
	TickMarkValueAtIndex(index int) float64
	// deprecated
	SetImage(backgroundImage IImage)
	// deprecated
	Image() Image
	// deprecated
	Title() string
	// deprecated
	TitleCell() objc.Object
	// deprecated
	TitleColor() Color
	// deprecated
	TitleFont() Font
	// deprecated
	SetTitle(string_ string)
	// deprecated
	SetTitleCell(cell ICell)
	// deprecated
	SetTitleColor(newColor IColor)
	// deprecated
	SetTitleFont(fontObj IFont)
	SliderType() SliderType
	SetSliderType(value SliderType)
	AltIncrementValue() float64
	SetAltIncrementValue(value float64)
	KnobThickness() float64
	IsVertical() bool
	SetVertical(value bool)
	TrackFillColor() Color
	SetTrackFillColor(value IColor)
	MaxValue() float64
	SetMaxValue(value float64)
	MinValue() float64
	SetMinValue(value float64)
	AllowsTickMarkValuesOnly() bool
	SetAllowsTickMarkValuesOnly(value bool)
	NumberOfTickMarks() int
	SetNumberOfTickMarks(value int)
	TickMarkPosition() TickMarkPosition
	SetTickMarkPosition(value TickMarkPosition)
}

type Slider struct {
	Control
}

func MakeSlider(ptr unsafe.Pointer) Slider {
	return Slider{
		Control: MakeControl(ptr),
	}
}

func (sc _SliderClass) SliderWithTarget_Action(target objc.IObject, action objc.Selector) Slider {
	rv := objc.CallMethod[Slider](sc, objc.GetSelector("sliderWithTarget:action:"), target, action)
	return rv
}

func (sc _SliderClass) SliderWithValue_MinValue_MaxValue_Target_Action(value float64, minValue float64, maxValue float64, target objc.IObject, action objc.Selector) Slider {
	rv := objc.CallMethod[Slider](sc, objc.GetSelector("sliderWithValue:minValue:maxValue:target:action:"), value, minValue, maxValue, target, action)
	return rv
}

func (s_ Slider) InitWithFrame(frameRect foundation.Rect) Slider {
	rv := objc.CallMethod[Slider](s_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func (s_ Slider) Init() Slider {
	rv := objc.CallMethod[Slider](s_, objc.GetSelector("init"))
	return rv
}

func (sc _SliderClass) Alloc() Slider {
	rv := objc.CallMethod[Slider](sc, objc.GetSelector("alloc"))
	return rv
}

func (sc _SliderClass) New() Slider {
	rv := objc.CallMethod[Slider](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewSlider() Slider {
	return SliderClass.New()
}

// deprecated
func (s_ Slider) SetKnobThickness(thickness float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setKnobThickness:"), thickness)
}

func (s_ Slider) ClosestTickMarkValueToValue(value float64) float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("closestTickMarkValueToValue:"), value)
	return rv
}

func (s_ Slider) IndexOfTickMarkAtPoint(point foundation.Point) int {
	rv := objc.CallMethod[int](s_, objc.GetSelector("indexOfTickMarkAtPoint:"), point)
	return rv
}

func (s_ Slider) RectOfTickMarkAtIndex(index int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("rectOfTickMarkAtIndex:"), index)
	return rv
}

func (s_ Slider) TickMarkValueAtIndex(index int) float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("tickMarkValueAtIndex:"), index)
	return rv
}

// deprecated
func (s_ Slider) SetImage(backgroundImage IImage) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setImage:"), backgroundImage)
}

// deprecated
func (s_ Slider) Image() Image {
	rv := objc.CallMethod[Image](s_, objc.GetSelector("image"))
	return rv
}

// deprecated
func (s_ Slider) Title() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("title"))
	return rv
}

// deprecated
func (s_ Slider) TitleCell() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.GetSelector("titleCell"))
	return rv
}

// deprecated
func (s_ Slider) TitleColor() Color {
	rv := objc.CallMethod[Color](s_, objc.GetSelector("titleColor"))
	return rv
}

// deprecated
func (s_ Slider) TitleFont() Font {
	rv := objc.CallMethod[Font](s_, objc.GetSelector("titleFont"))
	return rv
}

// deprecated
func (s_ Slider) SetTitle(string_ string) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setTitle:"), string_)
}

// deprecated
func (s_ Slider) SetTitleCell(cell ICell) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setTitleCell:"), cell)
}

// deprecated
func (s_ Slider) SetTitleColor(newColor IColor) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setTitleColor:"), newColor)
}

// deprecated
func (s_ Slider) SetTitleFont(fontObj IFont) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setTitleFont:"), fontObj)
}

func (s_ Slider) SliderType() SliderType {
	rv := objc.CallMethod[SliderType](s_, objc.GetSelector("sliderType"))
	return rv
}

func (s_ Slider) SetSliderType(value SliderType) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSliderType:"), value)
}

func (s_ Slider) AltIncrementValue() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("altIncrementValue"))
	return rv
}

func (s_ Slider) SetAltIncrementValue(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAltIncrementValue:"), value)
}

func (s_ Slider) KnobThickness() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("knobThickness"))
	return rv
}

func (s_ Slider) IsVertical() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("isVertical"))
	return rv
}

func (s_ Slider) SetVertical(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setVertical:"), value)
}

func (s_ Slider) TrackFillColor() Color {
	rv := objc.CallMethod[Color](s_, objc.GetSelector("trackFillColor"))
	return rv
}

func (s_ Slider) SetTrackFillColor(value IColor) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setTrackFillColor:"), value)
}

func (s_ Slider) MaxValue() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("maxValue"))
	return rv
}

func (s_ Slider) SetMaxValue(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setMaxValue:"), value)
}

func (s_ Slider) MinValue() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("minValue"))
	return rv
}

func (s_ Slider) SetMinValue(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setMinValue:"), value)
}

func (s_ Slider) AllowsTickMarkValuesOnly() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("allowsTickMarkValuesOnly"))
	return rv
}

func (s_ Slider) SetAllowsTickMarkValuesOnly(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAllowsTickMarkValuesOnly:"), value)
}

func (s_ Slider) NumberOfTickMarks() int {
	rv := objc.CallMethod[int](s_, objc.GetSelector("numberOfTickMarks"))
	return rv
}

func (s_ Slider) SetNumberOfTickMarks(value int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setNumberOfTickMarks:"), value)
}

func (s_ Slider) TickMarkPosition() TickMarkPosition {
	rv := objc.CallMethod[TickMarkPosition](s_, objc.GetSelector("tickMarkPosition"))
	return rv
}

func (s_ Slider) SetTickMarkPosition(value TickMarkPosition) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setTickMarkPosition:"), value)
}
