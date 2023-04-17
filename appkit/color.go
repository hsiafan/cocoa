// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[Color](c_, "init")
	return rv
}

func (cc _ColorClass) Alloc() Color {
	rv := ffi.CallMethod[Color](cc, "alloc")
	return rv
}

func (cc _ColorClass) New() Color {
	rv := ffi.CallMethod[Color](cc, "new")
	rv.Autorelease()
	return rv
}

func NewColor() Color {
	return ColorClass.New()
}

func (c_ Color) ColorWithSystemEffect(systemEffect ColorSystemEffect) Color {
	rv := ffi.CallMethod[Color](c_, "colorWithSystemEffect:", systemEffect)
	return rv
}

func (c_ Color) ColorUsingColorSpace(space IColorSpace) Color {
	rv := ffi.CallMethod[Color](c_, "colorUsingColorSpace:", space)
	return rv
}

func (c_ Color) BlendedColorWithFraction_OfColor(fraction float64, color IColor) Color {
	rv := ffi.CallMethod[Color](c_, "blendedColorWithFraction:ofColor:", fraction, color)
	return rv
}

func (c_ Color) ColorWithAlphaComponent(alpha float64) Color {
	rv := ffi.CallMethod[Color](c_, "colorWithAlphaComponent:", alpha)
	return rv
}

func (c_ Color) HighlightWithLevel(val float64) Color {
	rv := ffi.CallMethod[Color](c_, "highlightWithLevel:", val)
	return rv
}

func (c_ Color) ShadowWithLevel(val float64) Color {
	rv := ffi.CallMethod[Color](c_, "shadowWithLevel:", val)
	return rv
}

// deprecated
func (c_ Color) ColorUsingColorSpaceName(name ColorSpaceName) Color {
	rv := ffi.CallMethod[Color](c_, "colorUsingColorSpaceName:", name)
	return rv
}

// deprecated
func (c_ Color) ColorUsingColorSpaceName_Device(name ColorSpaceName, deviceDescription map[DeviceDescriptionKey]objc.IObject) Color {
	rv := ffi.CallMethod[Color](c_, "colorUsingColorSpaceName:device:", name, deviceDescription)
	return rv
}

func (cc _ColorClass) ColorFromPasteboard(pasteBoard IPasteboard) Color {
	rv := ffi.CallMethod[Color](cc, "colorFromPasteboard:", pasteBoard)
	return rv
}

func (c_ Color) WriteToPasteboard(pasteBoard IPasteboard) {
	ffi.CallMethod[ffi.Void](c_, "writeToPasteboard:", pasteBoard)
}

func (c_ Color) GetCyan_Magenta_Yellow_Black_Alpha(cyan *float64, magenta *float64, yellow *float64, black *float64, alpha *float64) {
	ffi.CallMethod[ffi.Void](c_, "getCyan:magenta:yellow:black:alpha:", cyan, magenta, yellow, black, alpha)
}

func (c_ Color) GetHue_Saturation_Brightness_Alpha(hue *float64, saturation *float64, brightness *float64, alpha *float64) {
	ffi.CallMethod[ffi.Void](c_, "getHue:saturation:brightness:alpha:", hue, saturation, brightness, alpha)
}

func (c_ Color) GetRed_Green_Blue_Alpha(red *float64, green *float64, blue *float64, alpha *float64) {
	ffi.CallMethod[ffi.Void](c_, "getRed:green:blue:alpha:", red, green, blue, alpha)
}

func (c_ Color) GetWhite_Alpha(white *float64, alpha *float64) {
	ffi.CallMethod[ffi.Void](c_, "getWhite:alpha:", white, alpha)
}

func (c_ Color) GetComponents(components *float64) {
	ffi.CallMethod[ffi.Void](c_, "getComponents:", components)
}

func (c_ Color) ColorUsingType(type_ ColorType) Color {
	rv := ffi.CallMethod[Color](c_, "colorUsingType:", type_)
	return rv
}

func (c_ Color) DrawSwatchInRect(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](c_, "drawSwatchInRect:", rect)
}

func (c_ Color) Set() {
	ffi.CallMethod[ffi.Void](c_, "set")
}

func (c_ Color) SetFill() {
	ffi.CallMethod[ffi.Void](c_, "setFill")
}

func (c_ Color) SetStroke() {
	ffi.CallMethod[ffi.Void](c_, "setStroke")
}

func (cc _ColorClass) ColorNamed(name ColorName) Color {
	rv := ffi.CallMethod[Color](cc, "colorNamed:", name)
	return rv
}

func (cc _ColorClass) ColorNamed_Bundle(name ColorName, bundle foundation.IBundle) Color {
	rv := ffi.CallMethod[Color](cc, "colorNamed:bundle:", name, bundle)
	return rv
}

func (cc _ColorClass) ColorWithCatalogName_ColorName(listName ColorListName, colorName ColorName) Color {
	rv := ffi.CallMethod[Color](cc, "colorWithCatalogName:colorName:", listName, colorName)
	return rv
}

func (cc _ColorClass) ColorWithSRGBRed_Green_Blue_Alpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := ffi.CallMethod[Color](cc, "colorWithSRGBRed:green:blue:alpha:", red, green, blue, alpha)
	return rv
}

func (cc _ColorClass) ColorWithDisplayP3Red_Green_Blue_Alpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := ffi.CallMethod[Color](cc, "colorWithDisplayP3Red:green:blue:alpha:", red, green, blue, alpha)
	return rv
}

func (cc _ColorClass) ColorWithRed_Green_Blue_Alpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := ffi.CallMethod[Color](cc, "colorWithRed:green:blue:alpha:", red, green, blue, alpha)
	return rv
}

func (cc _ColorClass) ColorWithCalibratedRed_Green_Blue_Alpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := ffi.CallMethod[Color](cc, "colorWithCalibratedRed:green:blue:alpha:", red, green, blue, alpha)
	return rv
}

func (cc _ColorClass) ColorWithDeviceRed_Green_Blue_Alpha(red float64, green float64, blue float64, alpha float64) Color {
	rv := ffi.CallMethod[Color](cc, "colorWithDeviceRed:green:blue:alpha:", red, green, blue, alpha)
	return rv
}

func (cc _ColorClass) ColorWithCalibratedHue_Saturation_Brightness_Alpha(hue float64, saturation float64, brightness float64, alpha float64) Color {
	rv := ffi.CallMethod[Color](cc, "colorWithCalibratedHue:saturation:brightness:alpha:", hue, saturation, brightness, alpha)
	return rv
}

func (cc _ColorClass) ColorWithDeviceHue_Saturation_Brightness_Alpha(hue float64, saturation float64, brightness float64, alpha float64) Color {
	rv := ffi.CallMethod[Color](cc, "colorWithDeviceHue:saturation:brightness:alpha:", hue, saturation, brightness, alpha)
	return rv
}

func (cc _ColorClass) ColorWithHue_Saturation_Brightness_Alpha(hue float64, saturation float64, brightness float64, alpha float64) Color {
	rv := ffi.CallMethod[Color](cc, "colorWithHue:saturation:brightness:alpha:", hue, saturation, brightness, alpha)
	return rv
}

func (cc _ColorClass) ColorWithColorSpace_Hue_Saturation_Brightness_Alpha(space IColorSpace, hue float64, saturation float64, brightness float64, alpha float64) Color {
	rv := ffi.CallMethod[Color](cc, "colorWithColorSpace:hue:saturation:brightness:alpha:", space, hue, saturation, brightness, alpha)
	return rv
}

func (cc _ColorClass) ColorWithDeviceCyan_Magenta_Yellow_Black_Alpha(cyan float64, magenta float64, yellow float64, black float64, alpha float64) Color {
	rv := ffi.CallMethod[Color](cc, "colorWithDeviceCyan:magenta:yellow:black:alpha:", cyan, magenta, yellow, black, alpha)
	return rv
}

func (cc _ColorClass) ColorWithWhite_Alpha(white float64, alpha float64) Color {
	rv := ffi.CallMethod[Color](cc, "colorWithWhite:alpha:", white, alpha)
	return rv
}

func (cc _ColorClass) ColorWithCalibratedWhite_Alpha(white float64, alpha float64) Color {
	rv := ffi.CallMethod[Color](cc, "colorWithCalibratedWhite:alpha:", white, alpha)
	return rv
}

func (cc _ColorClass) ColorWithDeviceWhite_Alpha(white float64, alpha float64) Color {
	rv := ffi.CallMethod[Color](cc, "colorWithDeviceWhite:alpha:", white, alpha)
	return rv
}

func (cc _ColorClass) ColorWithGenericGamma22White_Alpha(white float64, alpha float64) Color {
	rv := ffi.CallMethod[Color](cc, "colorWithGenericGamma22White:alpha:", white, alpha)
	return rv
}

func (cc _ColorClass) ColorWithPatternImage(image IImage) Color {
	rv := ffi.CallMethod[Color](cc, "colorWithPatternImage:", image)
	return rv
}

func (cc _ColorClass) ColorWithName_DynamicProvider(colorName ColorName, dynamicProvider func(param1 Appearance) IColor) Color {
	rv := ffi.CallMethod[Color](cc, "colorWithName:dynamicProvider:", colorName, dynamicProvider)
	return rv
}

func (cc _ColorClass) ColorWithColorSpace_Components_Count(space IColorSpace, components *float64, numberOfComponents int) Color {
	rv := ffi.CallMethod[Color](cc, "colorWithColorSpace:components:count:", space, components, numberOfComponents)
	return rv
}

// deprecated
func (cc _ColorClass) ColorForControlTint(controlTint ControlTint) Color {
	rv := ffi.CallMethod[Color](cc, "colorForControlTint:", controlTint)
	return rv
}

func (cc _ColorClass) ColorWithCGColor(cgColor coregraphics.ColorRef) Color {
	rv := ffi.CallMethod[Color](cc, "colorWithCGColor:", cgColor)
	return rv
}

func (cc _ColorClass) IgnoresAlpha() bool {
	rv := ffi.CallMethod[bool](cc, "ignoresAlpha")
	return rv
}

func (cc _ColorClass) SetIgnoresAlpha(value bool) {
	ffi.CallMethod[ffi.Void](cc, "setIgnoresAlpha:", value)
}

func (c_ Color) NumberOfComponents() int {
	rv := ffi.CallMethod[int](c_, "numberOfComponents")
	return rv
}

func (c_ Color) AlphaComponent() float64 {
	rv := ffi.CallMethod[float64](c_, "alphaComponent")
	return rv
}

func (c_ Color) WhiteComponent() float64 {
	rv := ffi.CallMethod[float64](c_, "whiteComponent")
	return rv
}

func (c_ Color) RedComponent() float64 {
	rv := ffi.CallMethod[float64](c_, "redComponent")
	return rv
}

func (c_ Color) GreenComponent() float64 {
	rv := ffi.CallMethod[float64](c_, "greenComponent")
	return rv
}

func (c_ Color) BlueComponent() float64 {
	rv := ffi.CallMethod[float64](c_, "blueComponent")
	return rv
}

func (c_ Color) CyanComponent() float64 {
	rv := ffi.CallMethod[float64](c_, "cyanComponent")
	return rv
}

func (c_ Color) MagentaComponent() float64 {
	rv := ffi.CallMethod[float64](c_, "magentaComponent")
	return rv
}

func (c_ Color) YellowComponent() float64 {
	rv := ffi.CallMethod[float64](c_, "yellowComponent")
	return rv
}

func (c_ Color) BlackComponent() float64 {
	rv := ffi.CallMethod[float64](c_, "blackComponent")
	return rv
}

func (c_ Color) HueComponent() float64 {
	rv := ffi.CallMethod[float64](c_, "hueComponent")
	return rv
}

func (c_ Color) SaturationComponent() float64 {
	rv := ffi.CallMethod[float64](c_, "saturationComponent")
	return rv
}

func (c_ Color) BrightnessComponent() float64 {
	rv := ffi.CallMethod[float64](c_, "brightnessComponent")
	return rv
}

func (c_ Color) CatalogNameComponent() ColorListName {
	rv := ffi.CallMethod[ColorListName](c_, "catalogNameComponent")
	return rv
}

func (c_ Color) LocalizedCatalogNameComponent() string {
	rv := ffi.CallMethod[string](c_, "localizedCatalogNameComponent")
	return rv
}

func (c_ Color) ColorNameComponent() ColorName {
	rv := ffi.CallMethod[ColorName](c_, "colorNameComponent")
	return rv
}

func (c_ Color) LocalizedColorNameComponent() string {
	rv := ffi.CallMethod[string](c_, "localizedColorNameComponent")
	return rv
}

func (c_ Color) Type() ColorType {
	rv := ffi.CallMethod[ColorType](c_, "type")
	return rv
}

func (c_ Color) ColorSpace() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](c_, "colorSpace")
	return rv
}

// deprecated
func (c_ Color) ColorSpaceName() ColorSpaceName {
	rv := ffi.CallMethod[ColorSpaceName](c_, "colorSpaceName")
	return rv
}

func (c_ Color) CGColor() coregraphics.ColorRef {
	rv := ffi.CallMethod[coregraphics.ColorRef](c_, "CGColor")
	return rv
}

func (cc _ColorClass) SystemCyanColor() Color {
	rv := ffi.CallMethod[Color](cc, "systemCyanColor")
	return rv
}

func (cc _ColorClass) SystemMintColor() Color {
	rv := ffi.CallMethod[Color](cc, "systemMintColor")
	return rv
}

func (cc _ColorClass) LabelColor() Color {
	rv := ffi.CallMethod[Color](cc, "labelColor")
	return rv
}

func (cc _ColorClass) SecondaryLabelColor() Color {
	rv := ffi.CallMethod[Color](cc, "secondaryLabelColor")
	return rv
}

func (cc _ColorClass) TertiaryLabelColor() Color {
	rv := ffi.CallMethod[Color](cc, "tertiaryLabelColor")
	return rv
}

func (cc _ColorClass) QuaternaryLabelColor() Color {
	rv := ffi.CallMethod[Color](cc, "quaternaryLabelColor")
	return rv
}

func (cc _ColorClass) TextColor() Color {
	rv := ffi.CallMethod[Color](cc, "textColor")
	return rv
}

func (cc _ColorClass) PlaceholderTextColor() Color {
	rv := ffi.CallMethod[Color](cc, "placeholderTextColor")
	return rv
}

func (cc _ColorClass) SelectedTextColor() Color {
	rv := ffi.CallMethod[Color](cc, "selectedTextColor")
	return rv
}

func (cc _ColorClass) TextBackgroundColor() Color {
	rv := ffi.CallMethod[Color](cc, "textBackgroundColor")
	return rv
}

func (cc _ColorClass) SelectedTextBackgroundColor() Color {
	rv := ffi.CallMethod[Color](cc, "selectedTextBackgroundColor")
	return rv
}

func (cc _ColorClass) KeyboardFocusIndicatorColor() Color {
	rv := ffi.CallMethod[Color](cc, "keyboardFocusIndicatorColor")
	return rv
}

func (cc _ColorClass) UnemphasizedSelectedTextColor() Color {
	rv := ffi.CallMethod[Color](cc, "unemphasizedSelectedTextColor")
	return rv
}

func (cc _ColorClass) UnemphasizedSelectedTextBackgroundColor() Color {
	rv := ffi.CallMethod[Color](cc, "unemphasizedSelectedTextBackgroundColor")
	return rv
}

func (cc _ColorClass) LinkColor() Color {
	rv := ffi.CallMethod[Color](cc, "linkColor")
	return rv
}

func (cc _ColorClass) SeparatorColor() Color {
	rv := ffi.CallMethod[Color](cc, "separatorColor")
	return rv
}

func (cc _ColorClass) SelectedContentBackgroundColor() Color {
	rv := ffi.CallMethod[Color](cc, "selectedContentBackgroundColor")
	return rv
}

func (cc _ColorClass) UnemphasizedSelectedContentBackgroundColor() Color {
	rv := ffi.CallMethod[Color](cc, "unemphasizedSelectedContentBackgroundColor")
	return rv
}

func (cc _ColorClass) SelectedMenuItemTextColor() Color {
	rv := ffi.CallMethod[Color](cc, "selectedMenuItemTextColor")
	return rv
}

func (cc _ColorClass) GridColor() Color {
	rv := ffi.CallMethod[Color](cc, "gridColor")
	return rv
}

func (cc _ColorClass) HeaderTextColor() Color {
	rv := ffi.CallMethod[Color](cc, "headerTextColor")
	return rv
}

func (cc _ColorClass) AlternatingContentBackgroundColors() []Color {
	rv := ffi.CallMethod[[]Color](cc, "alternatingContentBackgroundColors")
	return rv
}

func (cc _ColorClass) ControlAccentColor() Color {
	rv := ffi.CallMethod[Color](cc, "controlAccentColor")
	return rv
}

func (cc _ColorClass) ControlColor() Color {
	rv := ffi.CallMethod[Color](cc, "controlColor")
	return rv
}

func (cc _ColorClass) ControlBackgroundColor() Color {
	rv := ffi.CallMethod[Color](cc, "controlBackgroundColor")
	return rv
}

func (cc _ColorClass) ControlTextColor() Color {
	rv := ffi.CallMethod[Color](cc, "controlTextColor")
	return rv
}

func (cc _ColorClass) DisabledControlTextColor() Color {
	rv := ffi.CallMethod[Color](cc, "disabledControlTextColor")
	return rv
}

func (cc _ColorClass) CurrentControlTint() ControlTint {
	rv := ffi.CallMethod[ControlTint](cc, "currentControlTint")
	return rv
}

func (cc _ColorClass) SelectedControlColor() Color {
	rv := ffi.CallMethod[Color](cc, "selectedControlColor")
	return rv
}

func (cc _ColorClass) SelectedControlTextColor() Color {
	rv := ffi.CallMethod[Color](cc, "selectedControlTextColor")
	return rv
}

func (cc _ColorClass) AlternateSelectedControlTextColor() Color {
	rv := ffi.CallMethod[Color](cc, "alternateSelectedControlTextColor")
	return rv
}

func (cc _ColorClass) ScrubberTexturedBackgroundColor() Color {
	rv := ffi.CallMethod[Color](cc, "scrubberTexturedBackgroundColor")
	return rv
}

func (cc _ColorClass) WindowBackgroundColor() Color {
	rv := ffi.CallMethod[Color](cc, "windowBackgroundColor")
	return rv
}

func (cc _ColorClass) WindowFrameTextColor() Color {
	rv := ffi.CallMethod[Color](cc, "windowFrameTextColor")
	return rv
}

func (cc _ColorClass) UnderPageBackgroundColor() Color {
	rv := ffi.CallMethod[Color](cc, "underPageBackgroundColor")
	return rv
}

func (cc _ColorClass) FindHighlightColor() Color {
	rv := ffi.CallMethod[Color](cc, "findHighlightColor")
	return rv
}

func (cc _ColorClass) HighlightColor() Color {
	rv := ffi.CallMethod[Color](cc, "highlightColor")
	return rv
}

func (cc _ColorClass) ShadowColor() Color {
	rv := ffi.CallMethod[Color](cc, "shadowColor")
	return rv
}

// deprecated
func (cc _ColorClass) AlternateSelectedControlColor() Color {
	rv := ffi.CallMethod[Color](cc, "alternateSelectedControlColor")
	return rv
}

// deprecated
func (cc _ColorClass) ControlAlternatingRowBackgroundColors() []Color {
	rv := ffi.CallMethod[[]Color](cc, "controlAlternatingRowBackgroundColors")
	return rv
}

// deprecated
func (cc _ColorClass) ControlHighlightColor() Color {
	rv := ffi.CallMethod[Color](cc, "controlHighlightColor")
	return rv
}

// deprecated
func (cc _ColorClass) ControlLightHighlightColor() Color {
	rv := ffi.CallMethod[Color](cc, "controlLightHighlightColor")
	return rv
}

// deprecated
func (cc _ColorClass) ControlShadowColor() Color {
	rv := ffi.CallMethod[Color](cc, "controlShadowColor")
	return rv
}

// deprecated
func (cc _ColorClass) ControlDarkShadowColor() Color {
	rv := ffi.CallMethod[Color](cc, "controlDarkShadowColor")
	return rv
}

// deprecated
func (cc _ColorClass) HeaderColor() Color {
	rv := ffi.CallMethod[Color](cc, "headerColor")
	return rv
}

// deprecated
func (cc _ColorClass) KnobColor() Color {
	rv := ffi.CallMethod[Color](cc, "knobColor")
	return rv
}

// deprecated
func (cc _ColorClass) SelectedKnobColor() Color {
	rv := ffi.CallMethod[Color](cc, "selectedKnobColor")
	return rv
}

// deprecated
func (cc _ColorClass) ScrollBarColor() Color {
	rv := ffi.CallMethod[Color](cc, "scrollBarColor")
	return rv
}

// deprecated
func (cc _ColorClass) SecondarySelectedControlColor() Color {
	rv := ffi.CallMethod[Color](cc, "secondarySelectedControlColor")
	return rv
}

// deprecated
func (cc _ColorClass) SelectedMenuItemColor() Color {
	rv := ffi.CallMethod[Color](cc, "selectedMenuItemColor")
	return rv
}

// deprecated
func (cc _ColorClass) WindowFrameColor() Color {
	rv := ffi.CallMethod[Color](cc, "windowFrameColor")
	return rv
}

func (cc _ColorClass) SystemBlueColor() Color {
	rv := ffi.CallMethod[Color](cc, "systemBlueColor")
	return rv
}

func (cc _ColorClass) SystemBrownColor() Color {
	rv := ffi.CallMethod[Color](cc, "systemBrownColor")
	return rv
}

func (cc _ColorClass) SystemGrayColor() Color {
	rv := ffi.CallMethod[Color](cc, "systemGrayColor")
	return rv
}

func (cc _ColorClass) SystemGreenColor() Color {
	rv := ffi.CallMethod[Color](cc, "systemGreenColor")
	return rv
}

func (cc _ColorClass) SystemIndigoColor() Color {
	rv := ffi.CallMethod[Color](cc, "systemIndigoColor")
	return rv
}

func (cc _ColorClass) SystemOrangeColor() Color {
	rv := ffi.CallMethod[Color](cc, "systemOrangeColor")
	return rv
}

func (cc _ColorClass) SystemPinkColor() Color {
	rv := ffi.CallMethod[Color](cc, "systemPinkColor")
	return rv
}

func (cc _ColorClass) SystemPurpleColor() Color {
	rv := ffi.CallMethod[Color](cc, "systemPurpleColor")
	return rv
}

func (cc _ColorClass) SystemRedColor() Color {
	rv := ffi.CallMethod[Color](cc, "systemRedColor")
	return rv
}

func (cc _ColorClass) SystemTealColor() Color {
	rv := ffi.CallMethod[Color](cc, "systemTealColor")
	return rv
}

func (cc _ColorClass) SystemYellowColor() Color {
	rv := ffi.CallMethod[Color](cc, "systemYellowColor")
	return rv
}

func (cc _ColorClass) ClearColor() Color {
	rv := ffi.CallMethod[Color](cc, "clearColor")
	return rv
}

func (cc _ColorClass) BlackColor() Color {
	rv := ffi.CallMethod[Color](cc, "blackColor")
	return rv
}

func (cc _ColorClass) BlueColor() Color {
	rv := ffi.CallMethod[Color](cc, "blueColor")
	return rv
}

func (cc _ColorClass) BrownColor() Color {
	rv := ffi.CallMethod[Color](cc, "brownColor")
	return rv
}

func (cc _ColorClass) CyanColor() Color {
	rv := ffi.CallMethod[Color](cc, "cyanColor")
	return rv
}

func (cc _ColorClass) DarkGrayColor() Color {
	rv := ffi.CallMethod[Color](cc, "darkGrayColor")
	return rv
}

func (cc _ColorClass) GrayColor() Color {
	rv := ffi.CallMethod[Color](cc, "grayColor")
	return rv
}

func (cc _ColorClass) GreenColor() Color {
	rv := ffi.CallMethod[Color](cc, "greenColor")
	return rv
}

func (cc _ColorClass) LightGrayColor() Color {
	rv := ffi.CallMethod[Color](cc, "lightGrayColor")
	return rv
}

func (cc _ColorClass) MagentaColor() Color {
	rv := ffi.CallMethod[Color](cc, "magentaColor")
	return rv
}

func (cc _ColorClass) OrangeColor() Color {
	rv := ffi.CallMethod[Color](cc, "orangeColor")
	return rv
}

func (cc _ColorClass) PurpleColor() Color {
	rv := ffi.CallMethod[Color](cc, "purpleColor")
	return rv
}

func (cc _ColorClass) RedColor() Color {
	rv := ffi.CallMethod[Color](cc, "redColor")
	return rv
}

func (cc _ColorClass) WhiteColor() Color {
	rv := ffi.CallMethod[Color](cc, "whiteColor")
	return rv
}

func (cc _ColorClass) YellowColor() Color {
	rv := ffi.CallMethod[Color](cc, "yellowColor")
	return rv
}

func (c_ Color) PatternImage() Image {
	rv := ffi.CallMethod[Image](c_, "patternImage")
	return rv
}
