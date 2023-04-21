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
	rv := objc.CallMethod[Slider](sc, "sliderWithTarget:action:", target, action)
	return rv
}

func (sc _SliderClass) SliderWithValue_MinValue_MaxValue_Target_Action(value float64, minValue float64, maxValue float64, target objc.IObject, action objc.Selector) Slider {
	rv := objc.CallMethod[Slider](sc, "sliderWithValue:minValue:maxValue:target:action:", value, minValue, maxValue, target, action)
	return rv
}

func (s_ Slider) InitWithFrame(frameRect foundation.Rect) Slider {
	rv := objc.CallMethod[Slider](s_, "initWithFrame:", frameRect)
	return rv
}

func (s_ Slider) Init() Slider {
	rv := objc.CallMethod[Slider](s_, "init")
	return rv
}

func (sc _SliderClass) Alloc() Slider {
	rv := objc.CallMethod[Slider](sc, "alloc")
	return rv
}

func (sc _SliderClass) New() Slider {
	rv := objc.CallMethod[Slider](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSlider() Slider {
	return SliderClass.New()
}

// deprecated
func (s_ Slider) SetKnobThickness(thickness float64) {
	objc.CallMethod[objc.Void](s_, "setKnobThickness:", thickness)
}

func (s_ Slider) ClosestTickMarkValueToValue(value float64) float64 {
	rv := objc.CallMethod[float64](s_, "closestTickMarkValueToValue:", value)
	return rv
}

func (s_ Slider) IndexOfTickMarkAtPoint(point foundation.Point) int {
	rv := objc.CallMethod[int](s_, "indexOfTickMarkAtPoint:", point)
	return rv
}

func (s_ Slider) RectOfTickMarkAtIndex(index int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, "rectOfTickMarkAtIndex:", index)
	return rv
}

func (s_ Slider) TickMarkValueAtIndex(index int) float64 {
	rv := objc.CallMethod[float64](s_, "tickMarkValueAtIndex:", index)
	return rv
}

// deprecated
func (s_ Slider) SetImage(backgroundImage IImage) {
	objc.CallMethod[objc.Void](s_, "setImage:", backgroundImage)
}

// deprecated
func (s_ Slider) Image() Image {
	rv := objc.CallMethod[Image](s_, "image")
	return rv
}

// deprecated
func (s_ Slider) Title() string {
	rv := objc.CallMethod[string](s_, "title")
	return rv
}

// deprecated
func (s_ Slider) TitleCell() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, "titleCell")
	return rv
}

// deprecated
func (s_ Slider) TitleColor() Color {
	rv := objc.CallMethod[Color](s_, "titleColor")
	return rv
}

// deprecated
func (s_ Slider) TitleFont() Font {
	rv := objc.CallMethod[Font](s_, "titleFont")
	return rv
}

// deprecated
func (s_ Slider) SetTitle(string_ string) {
	objc.CallMethod[objc.Void](s_, "setTitle:", string_)
}

// deprecated
func (s_ Slider) SetTitleCell(cell ICell) {
	objc.CallMethod[objc.Void](s_, "setTitleCell:", cell)
}

// deprecated
func (s_ Slider) SetTitleColor(newColor IColor) {
	objc.CallMethod[objc.Void](s_, "setTitleColor:", newColor)
}

// deprecated
func (s_ Slider) SetTitleFont(fontObj IFont) {
	objc.CallMethod[objc.Void](s_, "setTitleFont:", fontObj)
}

func (s_ Slider) SliderType() SliderType {
	rv := objc.CallMethod[SliderType](s_, "sliderType")
	return rv
}

func (s_ Slider) SetSliderType(value SliderType) {
	objc.CallMethod[objc.Void](s_, "setSliderType:", value)
}

func (s_ Slider) AltIncrementValue() float64 {
	rv := objc.CallMethod[float64](s_, "altIncrementValue")
	return rv
}

func (s_ Slider) SetAltIncrementValue(value float64) {
	objc.CallMethod[objc.Void](s_, "setAltIncrementValue:", value)
}

func (s_ Slider) KnobThickness() float64 {
	rv := objc.CallMethod[float64](s_, "knobThickness")
	return rv
}

func (s_ Slider) IsVertical() bool {
	rv := objc.CallMethod[bool](s_, "isVertical")
	return rv
}

func (s_ Slider) SetVertical(value bool) {
	objc.CallMethod[objc.Void](s_, "setVertical:", value)
}

func (s_ Slider) TrackFillColor() Color {
	rv := objc.CallMethod[Color](s_, "trackFillColor")
	return rv
}

func (s_ Slider) SetTrackFillColor(value IColor) {
	objc.CallMethod[objc.Void](s_, "setTrackFillColor:", value)
}

func (s_ Slider) MaxValue() float64 {
	rv := objc.CallMethod[float64](s_, "maxValue")
	return rv
}

func (s_ Slider) SetMaxValue(value float64) {
	objc.CallMethod[objc.Void](s_, "setMaxValue:", value)
}

func (s_ Slider) MinValue() float64 {
	rv := objc.CallMethod[float64](s_, "minValue")
	return rv
}

func (s_ Slider) SetMinValue(value float64) {
	objc.CallMethod[objc.Void](s_, "setMinValue:", value)
}

func (s_ Slider) AllowsTickMarkValuesOnly() bool {
	rv := objc.CallMethod[bool](s_, "allowsTickMarkValuesOnly")
	return rv
}

func (s_ Slider) SetAllowsTickMarkValuesOnly(value bool) {
	objc.CallMethod[objc.Void](s_, "setAllowsTickMarkValuesOnly:", value)
}

func (s_ Slider) NumberOfTickMarks() int {
	rv := objc.CallMethod[int](s_, "numberOfTickMarks")
	return rv
}

func (s_ Slider) SetNumberOfTickMarks(value int) {
	objc.CallMethod[objc.Void](s_, "setNumberOfTickMarks:", value)
}

func (s_ Slider) TickMarkPosition() TickMarkPosition {
	rv := objc.CallMethod[TickMarkPosition](s_, "tickMarkPosition")
	return rv
}

func (s_ Slider) SetTickMarkPosition(value TickMarkPosition) {
	objc.CallMethod[objc.Void](s_, "setTickMarkPosition:", value)
}
