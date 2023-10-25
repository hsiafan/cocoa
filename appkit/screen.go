// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ScreenClass = _ScreenClass{objc.GetClass("NSScreen")}

type _ScreenClass struct {
	objc.Class
}

type IScreen interface {
	objc.IObject
	// deprecated
	UserSpaceScaleFactor() float64
	CanRepresentDisplayGamut(displayGamut DisplayGamut) bool
	BackingAlignedRect_Options(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect
	ConvertRectFromBacking(rect foundation.Rect) foundation.Rect
	ConvertRectToBacking(rect foundation.Rect) foundation.Rect
	Depth() WindowDepth
	Frame() foundation.Rect
	SupportedWindowDepths() *WindowDepth
	DeviceDescription() map[DeviceDescriptionKey]objc.Object
	ColorSpace() ColorSpace
	LocalizedName() string
	BackingScaleFactor() float64
	VisibleFrame() foundation.Rect
	SafeAreaInsets() foundation.EdgeInsets
	MaximumPotentialExtendedDynamicRangeColorComponentValue() float64
	MaximumExtendedDynamicRangeColorComponentValue() float64
	MaximumReferenceExtendedDynamicRangeColorComponentValue() float64
	MaximumFramesPerSecond() int
	MinimumRefreshInterval() foundation.TimeInterval
	MaximumRefreshInterval() foundation.TimeInterval
	DisplayUpdateGranularity() foundation.TimeInterval
	LastDisplayUpdateTimestamp() foundation.TimeInterval
	AuxiliaryTopLeftArea() foundation.Rect
	AuxiliaryTopRightArea() foundation.Rect
}

type Screen struct {
	objc.Object
}

func MakeScreen(ptr unsafe.Pointer) Screen {
	return Screen{
		Object: objc.MakeObject(ptr),
	}
}

func (sc _ScreenClass) Alloc() Screen {
	rv := objc.CallMethod[Screen](sc, objc.SEL("alloc"))
	return rv
}

func (sc _ScreenClass) New() Screen {
	rv := objc.CallMethod[Screen](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewScreen() Screen {
	return ScreenClass.New()
}

func (s_ Screen) Init() Screen {
	rv := objc.CallMethod[Screen](s_, objc.SEL("init"))
	return rv
}

// deprecated
func (s_ Screen) UserSpaceScaleFactor() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("userSpaceScaleFactor"))
	return rv
}

func (s_ Screen) CanRepresentDisplayGamut(displayGamut DisplayGamut) bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("canRepresentDisplayGamut:"), displayGamut)
	return rv
}

func (s_ Screen) BackingAlignedRect_Options(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.SEL("backingAlignedRect:options:"), rect, options)
	return rv
}

func (s_ Screen) ConvertRectFromBacking(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.SEL("convertRectFromBacking:"), rect)
	return rv
}

func (s_ Screen) ConvertRectToBacking(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.SEL("convertRectToBacking:"), rect)
	return rv
}

func (sc _ScreenClass) MainScreen() Screen {
	rv := objc.CallMethod[Screen](sc, objc.SEL("mainScreen"))
	return rv
}

func (sc _ScreenClass) DeepestScreen() Screen {
	rv := objc.CallMethod[Screen](sc, objc.SEL("deepestScreen"))
	return rv
}

func (sc _ScreenClass) Screens() []Screen {
	rv := objc.CallMethod[[]Screen](sc, objc.SEL("screens"))
	return rv
}

func (s_ Screen) Depth() WindowDepth {
	rv := objc.CallMethod[WindowDepth](s_, objc.SEL("depth"))
	return rv
}

func (s_ Screen) Frame() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.SEL("frame"))
	return rv
}

func (s_ Screen) SupportedWindowDepths() *WindowDepth {
	rv := objc.CallMethod[*WindowDepth](s_, objc.SEL("supportedWindowDepths"))
	return rv
}

func (s_ Screen) DeviceDescription() map[DeviceDescriptionKey]objc.Object {
	rv := objc.CallMethod[map[DeviceDescriptionKey]objc.Object](s_, objc.SEL("deviceDescription"))
	return rv
}

func (s_ Screen) ColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](s_, objc.SEL("colorSpace"))
	return rv
}

func (s_ Screen) LocalizedName() string {
	rv := objc.CallMethod[string](s_, objc.SEL("localizedName"))
	return rv
}

func (sc _ScreenClass) ScreensHaveSeparateSpaces() bool {
	rv := objc.CallMethod[bool](sc, objc.SEL("screensHaveSeparateSpaces"))
	return rv
}

func (s_ Screen) BackingScaleFactor() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("backingScaleFactor"))
	return rv
}

func (s_ Screen) VisibleFrame() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.SEL("visibleFrame"))
	return rv
}

func (s_ Screen) SafeAreaInsets() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](s_, objc.SEL("safeAreaInsets"))
	return rv
}

func (s_ Screen) MaximumPotentialExtendedDynamicRangeColorComponentValue() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("maximumPotentialExtendedDynamicRangeColorComponentValue"))
	return rv
}

func (s_ Screen) MaximumExtendedDynamicRangeColorComponentValue() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("maximumExtendedDynamicRangeColorComponentValue"))
	return rv
}

func (s_ Screen) MaximumReferenceExtendedDynamicRangeColorComponentValue() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("maximumReferenceExtendedDynamicRangeColorComponentValue"))
	return rv
}

func (s_ Screen) MaximumFramesPerSecond() int {
	rv := objc.CallMethod[int](s_, objc.SEL("maximumFramesPerSecond"))
	return rv
}

func (s_ Screen) MinimumRefreshInterval() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](s_, objc.SEL("minimumRefreshInterval"))
	return rv
}

func (s_ Screen) MaximumRefreshInterval() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](s_, objc.SEL("maximumRefreshInterval"))
	return rv
}

func (s_ Screen) DisplayUpdateGranularity() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](s_, objc.SEL("displayUpdateGranularity"))
	return rv
}

func (s_ Screen) LastDisplayUpdateTimestamp() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](s_, objc.SEL("lastDisplayUpdateTimestamp"))
	return rv
}

func (s_ Screen) AuxiliaryTopLeftArea() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.SEL("auxiliaryTopLeftArea"))
	return rv
}

func (s_ Screen) AuxiliaryTopRightArea() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.SEL("auxiliaryTopRightArea"))
	return rv
}
