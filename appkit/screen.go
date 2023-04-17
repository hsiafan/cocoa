// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[Screen](sc, "alloc")
	return rv
}

func (sc _ScreenClass) New() Screen {
	rv := ffi.CallMethod[Screen](sc, "new")
	rv.Autorelease()
	return rv
}

func NewScreen() Screen {
	return ScreenClass.New()
}

func (s_ Screen) Init() Screen {
	rv := ffi.CallMethod[Screen](s_, "init")
	return rv
}

// deprecated
func (s_ Screen) UserSpaceScaleFactor() float64 {
	rv := ffi.CallMethod[float64](s_, "userSpaceScaleFactor")
	return rv
}

func (s_ Screen) CanRepresentDisplayGamut(displayGamut DisplayGamut) bool {
	rv := ffi.CallMethod[bool](s_, "canRepresentDisplayGamut:", displayGamut)
	return rv
}

func (s_ Screen) BackingAlignedRect_Options(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "backingAlignedRect:options:", rect, options)
	return rv
}

func (s_ Screen) ConvertRectFromBacking(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "convertRectFromBacking:", rect)
	return rv
}

func (s_ Screen) ConvertRectToBacking(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "convertRectToBacking:", rect)
	return rv
}

func (sc _ScreenClass) MainScreen() Screen {
	rv := ffi.CallMethod[Screen](sc, "mainScreen")
	return rv
}

func (sc _ScreenClass) DeepestScreen() Screen {
	rv := ffi.CallMethod[Screen](sc, "deepestScreen")
	return rv
}

func (sc _ScreenClass) Screens() []Screen {
	rv := ffi.CallMethod[[]Screen](sc, "screens")
	return rv
}

func (s_ Screen) Depth() WindowDepth {
	rv := ffi.CallMethod[WindowDepth](s_, "depth")
	return rv
}

func (s_ Screen) Frame() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "frame")
	return rv
}

func (s_ Screen) SupportedWindowDepths() *WindowDepth {
	rv := ffi.CallMethod[*WindowDepth](s_, "supportedWindowDepths")
	return rv
}

func (s_ Screen) DeviceDescription() map[DeviceDescriptionKey]objc.Object {
	rv := ffi.CallMethod[map[DeviceDescriptionKey]objc.Object](s_, "deviceDescription")
	return rv
}

func (s_ Screen) ColorSpace() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](s_, "colorSpace")
	return rv
}

func (s_ Screen) LocalizedName() string {
	rv := ffi.CallMethod[string](s_, "localizedName")
	return rv
}

func (sc _ScreenClass) ScreensHaveSeparateSpaces() bool {
	rv := ffi.CallMethod[bool](sc, "screensHaveSeparateSpaces")
	return rv
}

func (s_ Screen) BackingScaleFactor() float64 {
	rv := ffi.CallMethod[float64](s_, "backingScaleFactor")
	return rv
}

func (s_ Screen) VisibleFrame() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "visibleFrame")
	return rv
}

func (s_ Screen) SafeAreaInsets() foundation.EdgeInsets {
	rv := ffi.CallMethod[foundation.EdgeInsets](s_, "safeAreaInsets")
	return rv
}

func (s_ Screen) MaximumPotentialExtendedDynamicRangeColorComponentValue() float64 {
	rv := ffi.CallMethod[float64](s_, "maximumPotentialExtendedDynamicRangeColorComponentValue")
	return rv
}

func (s_ Screen) MaximumExtendedDynamicRangeColorComponentValue() float64 {
	rv := ffi.CallMethod[float64](s_, "maximumExtendedDynamicRangeColorComponentValue")
	return rv
}

func (s_ Screen) MaximumReferenceExtendedDynamicRangeColorComponentValue() float64 {
	rv := ffi.CallMethod[float64](s_, "maximumReferenceExtendedDynamicRangeColorComponentValue")
	return rv
}

func (s_ Screen) MaximumFramesPerSecond() int {
	rv := ffi.CallMethod[int](s_, "maximumFramesPerSecond")
	return rv
}

func (s_ Screen) MinimumRefreshInterval() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](s_, "minimumRefreshInterval")
	return rv
}

func (s_ Screen) MaximumRefreshInterval() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](s_, "maximumRefreshInterval")
	return rv
}

func (s_ Screen) DisplayUpdateGranularity() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](s_, "displayUpdateGranularity")
	return rv
}

func (s_ Screen) LastDisplayUpdateTimestamp() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](s_, "lastDisplayUpdateTimestamp")
	return rv
}

func (s_ Screen) AuxiliaryTopLeftArea() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "auxiliaryTopLeftArea")
	return rv
}

func (s_ Screen) AuxiliaryTopRightArea() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "auxiliaryTopRightArea")
	return rv
}
