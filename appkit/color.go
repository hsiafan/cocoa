// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ColorClass = _ColorClass{objc.GetClass("NSColor")}

type _ColorClass struct {
	objc.Class
}

type IColor interface {
	objc.IObject
	ColorWithSystemEffect(systemEffect ColorSystemEffect) Color
	ColorUsingColorSpace(space IColorSpace) Color
	BlendedColorWithFraction_OfColor(fraction float64, color IColor) Color
	ColorWithAlphaComponent(alpha float64) Color
	HighlightWithLevel(val float64) Color
	ShadowWithLevel(val float64) Color
	// deprecated
	ColorUsingColorSpaceName(name ColorSpaceName) Color
	// deprecated
	ColorUsingColorSpaceName_Device(name ColorSpaceName, deviceDescription map[DeviceDescriptionKey]objc.IObject) Color
	WriteToPasteboard(pasteBoard IPasteboard)
	GetCyan_Magenta_Yellow_Black_Alpha(cyan *float64, magenta *float64, yellow *float64, black *float64, alpha *float64)
	GetHue_Saturation_Brightness_Alpha(hue *float64, saturation *float64, brightness *float64, alpha *float64)
	GetRed_Green_Blue_Alpha(red *float64, green *float64, blue *float64, alpha *float64)
	GetWhite_Alpha(white *float64, alpha *float64)
	GetComponents(components *float64)
	ColorUsingType(type_ ColorType) Color
	DrawSwatchInRect(rect foundation.Rect)
	Set()
	SetFill()
	SetStroke()
	NumberOfComponents() int
	AlphaComponent() float64
	WhiteComponent() float64
	RedComponent() float64
	GreenComponent() float64
	BlueComponent() float64
	CyanComponent() float64
	MagentaComponent() float64
	YellowComponent() float64
	BlackComponent() float64
	HueComponent() float64
	SaturationComponent() float64
	BrightnessComponent() float64
	CatalogNameComponent() ColorListName
	LocalizedCatalogNameComponent() string
	ColorNameComponent() ColorName
	LocalizedColorNameComponent() string
	Type() ColorType
	ColorSpace() ColorSpace
	// deprecated
	ColorSpaceName() ColorSpaceName
	CGColor() coregraphics.ColorRef
	PatternImage() Image
}

type Color struct {
	objc.Object
}

func MakeColor(ptr unsafe.Pointer) Color {
	return Color{
		Object: objc.MakeObject(ptr),
	}
}

func (c_ Color) Init() Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("init"))
	return rv
}

func (cc _ColorClass) Alloc() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("alloc"))
	return rv
}

func (cc _ColorClass) New() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewColor() Color {
	return ColorClass.New()
}

func (c_ Color) ColorWithSystemEffect(systemEffect ColorSystemEffect) Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("colorWithSystemEffect:"), systemEffect)
	return rv
}

func (c_ Color) ColorUsingColorSpace(space IColorSpace) Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("colorUsingColorSpace:"), space)
	return rv
}

func (c_ Color) BlendedColorWithFraction_OfColor(fraction float64, color IColor) Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("blendedColorWithFraction:ofColor:"), fraction, color)
	return rv
}

func (c_ Color) ColorWithAlphaComponent(alpha float64) Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("colorWithAlphaComponent:"), alpha)
	return rv
}

func (c_ Color) HighlightWithLevel(val float64) Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("highlightWithLevel:"), val)
	return rv
}

func (c_ Color) ShadowWithLevel(val float64) Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("shadowWithLevel:"), val)
	return rv
}

// deprecated
func (c_ Color) ColorUsingColorSpaceName(name ColorSpaceName) Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("colorUsingColorSpaceName:"), name)
	return rv
}

// deprecated
func (c_ Color) ColorUsingColorSpaceName_Device(name ColorSpaceName, deviceDescription map[DeviceDescriptionKey]objc.IObject) Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("colorUsingColorSpaceName:device:"), name, deviceDescription)
	return rv
}

func (cc _ColorClass) ColorFromPasteboard(pasteBoard IPasteboard) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorFromPasteboard:"), pasteBoard)
	return rv
}

func (c_ Color) WriteToPasteboard(pasteBoard IPasteboard) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("writeToPasteboard:"), pasteBoard)
}

func (c_ Color) GetCyan_Magenta_Yellow_Black_Alpha(cyan *float64, magenta *float64, yellow *float64, black *float64, alpha *float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("getCyan:magenta:yellow:black:alpha:"), cyan, magenta, yellow, black, alpha)
}

func (c_ Color) GetHue_Saturation_Brightness_Alpha(hue *float64, saturation *float64, brightness *float64, alpha *float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("getHue:saturation:brightness:alpha:"), hue, saturation, brightness, alpha)
}

func (c_ Color) GetRed_Green_Blue_Alpha(red *float64, green *float64, blue *float64, alpha *float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("getRed:green:blue:alpha:"), red, green, blue, alpha)
}

func (c_ Color) GetWhite_Alpha(white *float64, alpha *float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("getWhite:alpha:"), white, alpha)
}

func (c_ Color) GetComponents(components *float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("getComponents:"), components)
}

func (c_ Color) ColorUsingType(type_ ColorType) Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("colorUsingType:"), type_)
	return rv
}

func (c_ Color) DrawSwatchInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("drawSwatchInRect:"), rect)
}

func (c_ Color) Set() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("set"))
}

func (c_ Color) SetFill() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setFill"))
}

func (c_ Color) SetStroke() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setStroke"))
}

func (cc _ColorClass) ColorNamed(name ColorName) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorNamed:"), name)
	return rv
}

func (cc _ColorClass) ColorNamed_Bundle(name ColorName, bundle foundation.IBundle) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorNamed:bundle:"), name, bundle)
	return rv
}

func (cc _ColorClass) ColorWithCatalogName_ColorName(listName ColorListName, colorName ColorName) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithCatalogName:colorName:"), listName, colorName)
	return rv
}

func (cc _ColorClass) ColorWithSRGBRed_Green_Blue_Alpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithSRGBRed:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}

func (cc _ColorClass) ColorWithDisplayP3Red_Green_Blue_Alpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithDisplayP3Red:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}

func (cc _ColorClass) ColorWithRed_Green_Blue_Alpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithRed:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}

func (cc _ColorClass) ColorWithCalibratedRed_Green_Blue_Alpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithCalibratedRed:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}

func (cc _ColorClass) ColorWithDeviceRed_Green_Blue_Alpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithDeviceRed:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}

func (cc _ColorClass) ColorWithCalibratedHue_Saturation_Brightness_Alpha(hue float64, saturation float64, brightness float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithCalibratedHue:saturation:brightness:alpha:"), hue, saturation, brightness, alpha)
	return rv
}

func (cc _ColorClass) ColorWithDeviceHue_Saturation_Brightness_Alpha(hue float64, saturation float64, brightness float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithDeviceHue:saturation:brightness:alpha:"), hue, saturation, brightness, alpha)
	return rv
}

func (cc _ColorClass) ColorWithHue_Saturation_Brightness_Alpha(hue float64, saturation float64, brightness float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithHue:saturation:brightness:alpha:"), hue, saturation, brightness, alpha)
	return rv
}

func (cc _ColorClass) ColorWithColorSpace_Hue_Saturation_Brightness_Alpha(space IColorSpace, hue float64, saturation float64, brightness float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithColorSpace:hue:saturation:brightness:alpha:"), space, hue, saturation, brightness, alpha)
	return rv
}

func (cc _ColorClass) ColorWithDeviceCyan_Magenta_Yellow_Black_Alpha(cyan float64, magenta float64, yellow float64, black float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithDeviceCyan:magenta:yellow:black:alpha:"), cyan, magenta, yellow, black, alpha)
	return rv
}

func (cc _ColorClass) ColorWithWhite_Alpha(white float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithWhite:alpha:"), white, alpha)
	return rv
}

func (cc _ColorClass) ColorWithCalibratedWhite_Alpha(white float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithCalibratedWhite:alpha:"), white, alpha)
	return rv
}

func (cc _ColorClass) ColorWithDeviceWhite_Alpha(white float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithDeviceWhite:alpha:"), white, alpha)
	return rv
}

func (cc _ColorClass) ColorWithGenericGamma22White_Alpha(white float64, alpha float64) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithGenericGamma22White:alpha:"), white, alpha)
	return rv
}

func (cc _ColorClass) ColorWithPatternImage(image IImage) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithPatternImage:"), image)
	return rv
}

func (cc _ColorClass) ColorWithName_DynamicProvider(colorName ColorName, dynamicProvider func(param1 Appearance) IColor) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithName:dynamicProvider:"), colorName, dynamicProvider)
	return rv
}

func (cc _ColorClass) ColorWithColorSpace_Components_Count(space IColorSpace, components *float64, numberOfComponents int) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithColorSpace:components:count:"), space, components, numberOfComponents)
	return rv
}

// deprecated
func (cc _ColorClass) ColorForControlTint(controlTint ControlTint) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorForControlTint:"), controlTint)
	return rv
}

func (cc _ColorClass) ColorWithCGColor(cgColor coregraphics.ColorRef) Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("colorWithCGColor:"), cgColor)
	return rv
}

func (cc _ColorClass) IgnoresAlpha() bool {
	rv := objc.CallMethod[bool](cc, objc.GetSelector("ignoresAlpha"))
	return rv
}

func (cc _ColorClass) SetIgnoresAlpha(value bool) {
	objc.CallMethod[objc.Void](cc, objc.GetSelector("setIgnoresAlpha:"), value)
}

func (c_ Color) NumberOfComponents() int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("numberOfComponents"))
	return rv
}

func (c_ Color) AlphaComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("alphaComponent"))
	return rv
}

func (c_ Color) WhiteComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("whiteComponent"))
	return rv
}

func (c_ Color) RedComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("redComponent"))
	return rv
}

func (c_ Color) GreenComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("greenComponent"))
	return rv
}

func (c_ Color) BlueComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("blueComponent"))
	return rv
}

func (c_ Color) CyanComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("cyanComponent"))
	return rv
}

func (c_ Color) MagentaComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("magentaComponent"))
	return rv
}

func (c_ Color) YellowComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("yellowComponent"))
	return rv
}

func (c_ Color) BlackComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("blackComponent"))
	return rv
}

func (c_ Color) HueComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("hueComponent"))
	return rv
}

func (c_ Color) SaturationComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("saturationComponent"))
	return rv
}

func (c_ Color) BrightnessComponent() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("brightnessComponent"))
	return rv
}

func (c_ Color) CatalogNameComponent() ColorListName {
	rv := objc.CallMethod[ColorListName](c_, objc.GetSelector("catalogNameComponent"))
	return rv
}

func (c_ Color) LocalizedCatalogNameComponent() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("localizedCatalogNameComponent"))
	return rv
}

func (c_ Color) ColorNameComponent() ColorName {
	rv := objc.CallMethod[ColorName](c_, objc.GetSelector("colorNameComponent"))
	return rv
}

func (c_ Color) LocalizedColorNameComponent() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("localizedColorNameComponent"))
	return rv
}

func (c_ Color) Type() ColorType {
	rv := objc.CallMethod[ColorType](c_, objc.GetSelector("type"))
	return rv
}

func (c_ Color) ColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](c_, objc.GetSelector("colorSpace"))
	return rv
}

// deprecated
func (c_ Color) ColorSpaceName() ColorSpaceName {
	rv := objc.CallMethod[ColorSpaceName](c_, objc.GetSelector("colorSpaceName"))
	return rv
}

func (c_ Color) CGColor() coregraphics.ColorRef {
	rv := objc.CallMethod[coregraphics.ColorRef](c_, objc.GetSelector("CGColor"))
	return rv
}

func (cc _ColorClass) SystemCyanColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemCyanColor"))
	return rv
}

func (cc _ColorClass) SystemMintColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemMintColor"))
	return rv
}

func (cc _ColorClass) LabelColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("labelColor"))
	return rv
}

func (cc _ColorClass) SecondaryLabelColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("secondaryLabelColor"))
	return rv
}

func (cc _ColorClass) TertiaryLabelColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("tertiaryLabelColor"))
	return rv
}

func (cc _ColorClass) QuaternaryLabelColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("quaternaryLabelColor"))
	return rv
}

func (cc _ColorClass) TextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("textColor"))
	return rv
}

func (cc _ColorClass) PlaceholderTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("placeholderTextColor"))
	return rv
}

func (cc _ColorClass) SelectedTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("selectedTextColor"))
	return rv
}

func (cc _ColorClass) TextBackgroundColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("textBackgroundColor"))
	return rv
}

func (cc _ColorClass) SelectedTextBackgroundColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("selectedTextBackgroundColor"))
	return rv
}

func (cc _ColorClass) KeyboardFocusIndicatorColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("keyboardFocusIndicatorColor"))
	return rv
}

func (cc _ColorClass) UnemphasizedSelectedTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("unemphasizedSelectedTextColor"))
	return rv
}

func (cc _ColorClass) UnemphasizedSelectedTextBackgroundColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("unemphasizedSelectedTextBackgroundColor"))
	return rv
}

func (cc _ColorClass) LinkColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("linkColor"))
	return rv
}

func (cc _ColorClass) SeparatorColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("separatorColor"))
	return rv
}

func (cc _ColorClass) SelectedContentBackgroundColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("selectedContentBackgroundColor"))
	return rv
}

func (cc _ColorClass) UnemphasizedSelectedContentBackgroundColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("unemphasizedSelectedContentBackgroundColor"))
	return rv
}

func (cc _ColorClass) SelectedMenuItemTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("selectedMenuItemTextColor"))
	return rv
}

func (cc _ColorClass) GridColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("gridColor"))
	return rv
}

func (cc _ColorClass) HeaderTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("headerTextColor"))
	return rv
}

func (cc _ColorClass) AlternatingContentBackgroundColors() []Color {
	rv := objc.CallMethod[[]Color](cc, objc.GetSelector("alternatingContentBackgroundColors"))
	return rv
}

func (cc _ColorClass) ControlAccentColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("controlAccentColor"))
	return rv
}

func (cc _ColorClass) ControlColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("controlColor"))
	return rv
}

func (cc _ColorClass) ControlBackgroundColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("controlBackgroundColor"))
	return rv
}

func (cc _ColorClass) ControlTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("controlTextColor"))
	return rv
}

func (cc _ColorClass) DisabledControlTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("disabledControlTextColor"))
	return rv
}

func (cc _ColorClass) CurrentControlTint() ControlTint {
	rv := objc.CallMethod[ControlTint](cc, objc.GetSelector("currentControlTint"))
	return rv
}

func (cc _ColorClass) SelectedControlColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("selectedControlColor"))
	return rv
}

func (cc _ColorClass) SelectedControlTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("selectedControlTextColor"))
	return rv
}

func (cc _ColorClass) AlternateSelectedControlTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("alternateSelectedControlTextColor"))
	return rv
}

func (cc _ColorClass) ScrubberTexturedBackgroundColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("scrubberTexturedBackgroundColor"))
	return rv
}

func (cc _ColorClass) WindowBackgroundColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("windowBackgroundColor"))
	return rv
}

func (cc _ColorClass) WindowFrameTextColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("windowFrameTextColor"))
	return rv
}

func (cc _ColorClass) UnderPageBackgroundColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("underPageBackgroundColor"))
	return rv
}

func (cc _ColorClass) FindHighlightColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("findHighlightColor"))
	return rv
}

func (cc _ColorClass) HighlightColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("highlightColor"))
	return rv
}

func (cc _ColorClass) ShadowColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("shadowColor"))
	return rv
}

// deprecated
func (cc _ColorClass) AlternateSelectedControlColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("alternateSelectedControlColor"))
	return rv
}

// deprecated
func (cc _ColorClass) ControlAlternatingRowBackgroundColors() []Color {
	rv := objc.CallMethod[[]Color](cc, objc.GetSelector("controlAlternatingRowBackgroundColors"))
	return rv
}

// deprecated
func (cc _ColorClass) ControlHighlightColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("controlHighlightColor"))
	return rv
}

// deprecated
func (cc _ColorClass) ControlLightHighlightColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("controlLightHighlightColor"))
	return rv
}

// deprecated
func (cc _ColorClass) ControlShadowColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("controlShadowColor"))
	return rv
}

// deprecated
func (cc _ColorClass) ControlDarkShadowColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("controlDarkShadowColor"))
	return rv
}

// deprecated
func (cc _ColorClass) HeaderColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("headerColor"))
	return rv
}

// deprecated
func (cc _ColorClass) KnobColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("knobColor"))
	return rv
}

// deprecated
func (cc _ColorClass) SelectedKnobColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("selectedKnobColor"))
	return rv
}

// deprecated
func (cc _ColorClass) ScrollBarColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("scrollBarColor"))
	return rv
}

// deprecated
func (cc _ColorClass) SecondarySelectedControlColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("secondarySelectedControlColor"))
	return rv
}

// deprecated
func (cc _ColorClass) SelectedMenuItemColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("selectedMenuItemColor"))
	return rv
}

// deprecated
func (cc _ColorClass) WindowFrameColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("windowFrameColor"))
	return rv
}

func (cc _ColorClass) SystemBlueColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemBlueColor"))
	return rv
}

func (cc _ColorClass) SystemBrownColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemBrownColor"))
	return rv
}

func (cc _ColorClass) SystemGrayColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemGrayColor"))
	return rv
}

func (cc _ColorClass) SystemGreenColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemGreenColor"))
	return rv
}

func (cc _ColorClass) SystemIndigoColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemIndigoColor"))
	return rv
}

func (cc _ColorClass) SystemOrangeColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemOrangeColor"))
	return rv
}

func (cc _ColorClass) SystemPinkColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemPinkColor"))
	return rv
}

func (cc _ColorClass) SystemPurpleColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemPurpleColor"))
	return rv
}

func (cc _ColorClass) SystemRedColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemRedColor"))
	return rv
}

func (cc _ColorClass) SystemTealColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemTealColor"))
	return rv
}

func (cc _ColorClass) SystemYellowColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("systemYellowColor"))
	return rv
}

func (cc _ColorClass) ClearColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("clearColor"))
	return rv
}

func (cc _ColorClass) BlackColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("blackColor"))
	return rv
}

func (cc _ColorClass) BlueColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("blueColor"))
	return rv
}

func (cc _ColorClass) BrownColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("brownColor"))
	return rv
}

func (cc _ColorClass) CyanColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("cyanColor"))
	return rv
}

func (cc _ColorClass) DarkGrayColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("darkGrayColor"))
	return rv
}

func (cc _ColorClass) GrayColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("grayColor"))
	return rv
}

func (cc _ColorClass) GreenColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("greenColor"))
	return rv
}

func (cc _ColorClass) LightGrayColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("lightGrayColor"))
	return rv
}

func (cc _ColorClass) MagentaColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("magentaColor"))
	return rv
}

func (cc _ColorClass) OrangeColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("orangeColor"))
	return rv
}

func (cc _ColorClass) PurpleColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("purpleColor"))
	return rv
}

func (cc _ColorClass) RedColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("redColor"))
	return rv
}

func (cc _ColorClass) WhiteColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("whiteColor"))
	return rv
}

func (cc _ColorClass) YellowColor() Color {
	rv := objc.CallMethod[Color](cc, objc.GetSelector("yellowColor"))
	return rv
}

func (c_ Color) PatternImage() Image {
	rv := objc.CallMethod[Image](c_, objc.GetSelector("patternImage"))
	return rv
}
