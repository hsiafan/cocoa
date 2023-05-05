// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[ColorPanel](cc, objc.GetSelector("windowWithContentViewController:"), objc.ExtractPtr(contentViewController))
	return rv
}

func (c_ ColorPanel) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) ColorPanel {
	rv := objc.CallMethod[ColorPanel](c_, objc.GetSelector("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return rv
}

func (c_ ColorPanel) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) ColorPanel {
	rv := objc.CallMethod[ColorPanel](c_, objc.GetSelector("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.ExtractPtr(screen))
	return rv
}

func (c_ ColorPanel) Init() ColorPanel {
	rv := objc.CallMethod[ColorPanel](c_, objc.GetSelector("init"))
	return rv
}

func (cc _ColorPanelClass) Alloc() ColorPanel {
	rv := objc.CallMethod[ColorPanel](cc, objc.GetSelector("alloc"))
	return rv
}

func (cc _ColorPanelClass) New() ColorPanel {
	rv := objc.CallMethod[ColorPanel](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewColorPanel() ColorPanel {
	return ColorPanelClass.New()
}

func (cc _ColorPanelClass) SetPickerMode(mode ColorPanelMode) {
	objc.CallMethod[objc.Void](cc, objc.GetSelector("setPickerMode:"), mode)
}

func (cc _ColorPanelClass) SetPickerMask(mask ColorPanelOptions) {
	objc.CallMethod[objc.Void](cc, objc.GetSelector("setPickerMask:"), mask)
}

func (c_ ColorPanel) SetAction(selector objc.Selector) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setAction:"), selector)
}

func (c_ ColorPanel) SetTarget(target objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setTarget:"), objc.ExtractPtr(target))
}

func (c_ ColorPanel) AttachColorList(colorList IColorList) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("attachColorList:"), objc.ExtractPtr(colorList))
}

func (c_ ColorPanel) DetachColorList(colorList IColorList) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("detachColorList:"), objc.ExtractPtr(colorList))
}

func (cc _ColorPanelClass) DragColor_WithEvent_FromView(color IColor, event IEvent, sourceView IView) bool {
	rv := objc.CallMethod[bool](cc, objc.GetSelector("dragColor:withEvent:fromView:"), objc.ExtractPtr(color), objc.ExtractPtr(event), objc.ExtractPtr(sourceView))
	return rv
}

// deprecated
func (c_ ColorPanel) ChangeColor(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("changeColor:"), objc.ExtractPtr(sender))
}

func (cc _ColorPanelClass) SharedColorPanel() ColorPanel {
	rv := objc.CallMethod[ColorPanel](cc, objc.GetSelector("sharedColorPanel"))
	return rv
}

func (cc _ColorPanelClass) SharedColorPanelExists() bool {
	rv := objc.CallMethod[bool](cc, objc.GetSelector("sharedColorPanelExists"))
	return rv
}

func (c_ ColorPanel) Mode() ColorPanelMode {
	rv := objc.CallMethod[ColorPanelMode](c_, objc.GetSelector("mode"))
	return rv
}

func (c_ ColorPanel) SetMode(value ColorPanelMode) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setMode:"), value)
}

func (c_ ColorPanel) AccessoryView() View {
	rv := objc.CallMethod[View](c_, objc.GetSelector("accessoryView"))
	return rv
}

func (c_ ColorPanel) SetAccessoryView(value IView) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setAccessoryView:"), objc.ExtractPtr(value))
}

func (c_ ColorPanel) IsContinuous() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isContinuous"))
	return rv
}

func (c_ ColorPanel) SetContinuous(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setContinuous:"), value)
}

func (c_ ColorPanel) ShowsAlpha() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("showsAlpha"))
	return rv
}

func (c_ ColorPanel) SetShowsAlpha(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setShowsAlpha:"), value)
}

func (c_ ColorPanel) Color() Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("color"))
	return rv
}

func (c_ ColorPanel) SetColor(value IColor) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setColor:"), objc.ExtractPtr(value))
}

func (c_ ColorPanel) Alpha() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("alpha"))
	return rv
}
