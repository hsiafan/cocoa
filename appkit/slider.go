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
	rv := objc.CallMethod[Slider](sc, objc.SEL("sliderWithTarget:action:"), objc.ExtractPtr(target), action)
	return rv
}

func (sc _SliderClass) SliderWithValue_MinValue_MaxValue_Target_Action(value float64, minValue float64, maxValue float64, target objc.IObject, action objc.Selector) Slider {
	rv := objc.CallMethod[Slider](sc, objc.SEL("sliderWithValue:minValue:maxValue:target:action:"), value, minValue, maxValue, objc.ExtractPtr(target), action)
	return rv
}

func (s_ Slider) InitWithFrame(frameRect foundation.Rect) Slider {
	rv := objc.CallMethod[Slider](s_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (s_ Slider) Init() Slider {
	rv := objc.CallMethod[Slider](s_, objc.SEL("init"))
	return rv
}

func (sc _SliderClass) Alloc() Slider {
	rv := objc.CallMethod[Slider](sc, objc.SEL("alloc"))
	return rv
}

func (sc _SliderClass) New() Slider {
	rv := objc.CallMethod[Slider](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewSlider() Slider {
	return SliderClass.New()
}

// deprecated
func (s_ Slider) SetKnobThickness(thickness float64) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setKnobThickness:"), thickness)
}

func (s_ Slider) ClosestTickMarkValueToValue(value float64) float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("closestTickMarkValueToValue:"), value)
	return rv
}

func (s_ Slider) IndexOfTickMarkAtPoint(point foundation.Point) int {
	rv := objc.CallMethod[int](s_, objc.SEL("indexOfTickMarkAtPoint:"), point)
	return rv
}

func (s_ Slider) RectOfTickMarkAtIndex(index int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.SEL("rectOfTickMarkAtIndex:"), index)
	return rv
}

func (s_ Slider) TickMarkValueAtIndex(index int) float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("tickMarkValueAtIndex:"), index)
	return rv
}

// deprecated
func (s_ Slider) SetImage(backgroundImage IImage) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setImage:"), objc.ExtractPtr(backgroundImage))
}

// deprecated
func (s_ Slider) Image() Image {
	rv := objc.CallMethod[Image](s_, objc.SEL("image"))
	return rv
}

// deprecated
func (s_ Slider) Title() string {
	rv := objc.CallMethod[string](s_, objc.SEL("title"))
	return rv
}

// deprecated
func (s_ Slider) TitleCell() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.SEL("titleCell"))
	return rv
}

// deprecated
func (s_ Slider) TitleColor() Color {
	rv := objc.CallMethod[Color](s_, objc.SEL("titleColor"))
	return rv
}

// deprecated
func (s_ Slider) TitleFont() Font {
	rv := objc.CallMethod[Font](s_, objc.SEL("titleFont"))
	return rv
}

// deprecated
func (s_ Slider) SetTitle(string_ string) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setTitle:"), string_)
}

// deprecated
func (s_ Slider) SetTitleCell(cell ICell) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setTitleCell:"), objc.ExtractPtr(cell))
}

// deprecated
func (s_ Slider) SetTitleColor(newColor IColor) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setTitleColor:"), objc.ExtractPtr(newColor))
}

// deprecated
func (s_ Slider) SetTitleFont(fontObj IFont) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setTitleFont:"), objc.ExtractPtr(fontObj))
}

func (s_ Slider) SliderType() SliderType {
	rv := objc.CallMethod[SliderType](s_, objc.SEL("sliderType"))
	return rv
}

func (s_ Slider) SetSliderType(value SliderType) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setSliderType:"), value)
}

func (s_ Slider) AltIncrementValue() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("altIncrementValue"))
	return rv
}

func (s_ Slider) SetAltIncrementValue(value float64) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setAltIncrementValue:"), value)
}

func (s_ Slider) KnobThickness() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("knobThickness"))
	return rv
}

func (s_ Slider) IsVertical() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("isVertical"))
	return rv
}

func (s_ Slider) SetVertical(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setVertical:"), value)
}

func (s_ Slider) TrackFillColor() Color {
	rv := objc.CallMethod[Color](s_, objc.SEL("trackFillColor"))
	return rv
}

func (s_ Slider) SetTrackFillColor(value IColor) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setTrackFillColor:"), objc.ExtractPtr(value))
}

func (s_ Slider) MaxValue() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("maxValue"))
	return rv
}

func (s_ Slider) SetMaxValue(value float64) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setMaxValue:"), value)
}

func (s_ Slider) MinValue() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("minValue"))
	return rv
}

func (s_ Slider) SetMinValue(value float64) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setMinValue:"), value)
}

func (s_ Slider) AllowsTickMarkValuesOnly() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("allowsTickMarkValuesOnly"))
	return rv
}

func (s_ Slider) SetAllowsTickMarkValuesOnly(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setAllowsTickMarkValuesOnly:"), value)
}

func (s_ Slider) NumberOfTickMarks() int {
	rv := objc.CallMethod[int](s_, objc.SEL("numberOfTickMarks"))
	return rv
}

func (s_ Slider) SetNumberOfTickMarks(value int) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setNumberOfTickMarks:"), value)
}

func (s_ Slider) TickMarkPosition() TickMarkPosition {
	rv := objc.CallMethod[TickMarkPosition](s_, objc.SEL("tickMarkPosition"))
	return rv
}

func (s_ Slider) SetTickMarkPosition(value TickMarkPosition) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setTickMarkPosition:"), value)
}
