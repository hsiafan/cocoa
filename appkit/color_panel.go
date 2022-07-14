// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ColorPanelClass = _ColorPanelClass{objc.GetClass("NSColorPanel")}

type _ColorPanelClass struct {
	objc.Class
}

type IColorPanel interface {
	IPanel
	SetAction(selector objc.Selector)
	SetTarget(target objc.IObject)
	AttachColorList(colorList IColorList)
	DetachColorList(colorList IColorList)
	// deprecated
	ChangeColor(sender objc.IObject)
	Mode() ColorPanelMode
	SetMode(value ColorPanelMode)
	AccessoryView() View
	SetAccessoryView(value IView)
	IsContinuous() bool
	SetContinuous(value bool)
	ShowsAlpha() bool
	SetShowsAlpha(value bool)
	Color() Color
	SetColor(value IColor)
	Alpha() float64
}

type ColorPanel struct {
	Panel
}

func MakeColorPanel(ptr unsafe.Pointer) ColorPanel {
	return ColorPanel{
		Panel: MakePanel(ptr),
	}
}

func (cc _ColorPanelClass) WindowWithContentViewController(contentViewController IViewController) ColorPanel {
	rv := ffi.CallMethod[ColorPanel](cc, "windowWithContentViewController:", contentViewController)
	return rv
}

func (c_ ColorPanel) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) ColorPanel {
	rv := ffi.CallMethod[ColorPanel](c_, "initWithContentRect:styleMask:backing:defer:", contentRect, style, backingStoreType, flag)
	rv.Autorelease()
	return rv
}

func (c_ ColorPanel) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) ColorPanel {
	rv := ffi.CallMethod[ColorPanel](c_, "initWithContentRect:styleMask:backing:defer:screen:", contentRect, style, backingStoreType, flag, screen)
	rv.Autorelease()
	return rv
}

func (c_ ColorPanel) Init() ColorPanel {
	rv := ffi.CallMethod[ColorPanel](c_, "init")
	rv.Autorelease()
	return rv
}

func (cc _ColorPanelClass) Alloc() ColorPanel {
	rv := ffi.CallMethod[ColorPanel](cc, "alloc")
	return rv
}

func (cc _ColorPanelClass) New() ColorPanel {
	rv := ffi.CallMethod[ColorPanel](cc, "new")
	rv.Autorelease()
	return rv
}

func NewColorPanel() ColorPanel {
	return ColorPanelClass.New()
}

func (cc _ColorPanelClass) SetPickerMode(mode ColorPanelMode) {
	ffi.CallMethod[ffi.Void](cc, "setPickerMode:", mode)
}

func (cc _ColorPanelClass) SetPickerMask(mask ColorPanelOptions) {
	ffi.CallMethod[ffi.Void](cc, "setPickerMask:", mask)
}

func (c_ ColorPanel) SetAction(selector objc.Selector) {
	ffi.CallMethod[ffi.Void](c_, "setAction:", selector)
}

func (c_ ColorPanel) SetTarget(target objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "setTarget:", target)
}

func (c_ ColorPanel) AttachColorList(colorList IColorList) {
	ffi.CallMethod[ffi.Void](c_, "attachColorList:", colorList)
}

func (c_ ColorPanel) DetachColorList(colorList IColorList) {
	ffi.CallMethod[ffi.Void](c_, "detachColorList:", colorList)
}

func (cc _ColorPanelClass) DragColor_WithEvent_FromView(color IColor, event IEvent, sourceView IView) bool {
	rv := ffi.CallMethod[bool](cc, "dragColor:withEvent:fromView:", color, event, sourceView)
	return rv
}

// deprecated
func (c_ ColorPanel) ChangeColor(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "changeColor:", sender)
}

func (cc _ColorPanelClass) SharedColorPanel() ColorPanel {
	rv := ffi.CallMethod[ColorPanel](cc, "sharedColorPanel")
	return rv
}

func (cc _ColorPanelClass) SharedColorPanelExists() bool {
	rv := ffi.CallMethod[bool](cc, "sharedColorPanelExists")
	return rv
}

func (c_ ColorPanel) Mode() ColorPanelMode {
	rv := ffi.CallMethod[ColorPanelMode](c_, "mode")
	return rv
}

func (c_ ColorPanel) SetMode(value ColorPanelMode) {
	ffi.CallMethod[ffi.Void](c_, "setMode:", value)
}

func (c_ ColorPanel) AccessoryView() View {
	rv := ffi.CallMethod[View](c_, "accessoryView")
	return rv
}

func (c_ ColorPanel) SetAccessoryView(value IView) {
	ffi.CallMethod[ffi.Void](c_, "setAccessoryView:", value)
}

func (c_ ColorPanel) IsContinuous() bool {
	rv := ffi.CallMethod[bool](c_, "isContinuous")
	return rv
}

func (c_ ColorPanel) SetContinuous(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setContinuous:", value)
}

func (c_ ColorPanel) ShowsAlpha() bool {
	rv := ffi.CallMethod[bool](c_, "showsAlpha")
	return rv
}

func (c_ ColorPanel) SetShowsAlpha(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setShowsAlpha:", value)
}

func (c_ ColorPanel) Color() Color {
	rv := ffi.CallMethod[Color](c_, "color")
	return rv
}

func (c_ ColorPanel) SetColor(value IColor) {
	ffi.CallMethod[ffi.Void](c_, "setColor:", value)
}

func (c_ ColorPanel) Alpha() float64 {
	rv := ffi.CallMethod[float64](c_, "alpha")
	return rv
}
