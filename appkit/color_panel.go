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
	rv := objc.CallMethod[ColorPanel](cc, objc.SEL("windowWithContentViewController:"), objc.ExtractPtr(contentViewController))
	return rv
}

func (c_ ColorPanel) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) ColorPanel {
	rv := objc.CallMethod[ColorPanel](c_, objc.SEL("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return rv
}

func (c_ ColorPanel) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) ColorPanel {
	rv := objc.CallMethod[ColorPanel](c_, objc.SEL("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.ExtractPtr(screen))
	return rv
}

func (c_ ColorPanel) Init() ColorPanel {
	rv := objc.CallMethod[ColorPanel](c_, objc.SEL("init"))
	return rv
}

func (cc _ColorPanelClass) Alloc() ColorPanel {
	rv := objc.CallMethod[ColorPanel](cc, objc.SEL("alloc"))
	return rv
}

func (cc _ColorPanelClass) New() ColorPanel {
	rv := objc.CallMethod[ColorPanel](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewColorPanel() ColorPanel {
	return ColorPanelClass.New()
}

func (cc _ColorPanelClass) SetPickerMode(mode ColorPanelMode) {
	objc.CallMethod[objc.Void](cc, objc.SEL("setPickerMode:"), mode)
}

func (cc _ColorPanelClass) SetPickerMask(mask ColorPanelOptions) {
	objc.CallMethod[objc.Void](cc, objc.SEL("setPickerMask:"), mask)
}

func (c_ ColorPanel) SetAction(selector objc.Selector) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setAction:"), selector)
}

func (c_ ColorPanel) SetTarget(target objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setTarget:"), objc.ExtractPtr(target))
}

func (c_ ColorPanel) AttachColorList(colorList IColorList) {
	objc.CallMethod[objc.Void](c_, objc.SEL("attachColorList:"), objc.ExtractPtr(colorList))
}

func (c_ ColorPanel) DetachColorList(colorList IColorList) {
	objc.CallMethod[objc.Void](c_, objc.SEL("detachColorList:"), objc.ExtractPtr(colorList))
}

func (cc _ColorPanelClass) DragColor_WithEvent_FromView(color IColor, event IEvent, sourceView IView) bool {
	rv := objc.CallMethod[bool](cc, objc.SEL("dragColor:withEvent:fromView:"), objc.ExtractPtr(color), objc.ExtractPtr(event), objc.ExtractPtr(sourceView))
	return rv
}

// deprecated
func (c_ ColorPanel) ChangeColor(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("changeColor:"), objc.ExtractPtr(sender))
}

func (cc _ColorPanelClass) SharedColorPanel() ColorPanel {
	rv := objc.CallMethod[ColorPanel](cc, objc.SEL("sharedColorPanel"))
	return rv
}

func (cc _ColorPanelClass) SharedColorPanelExists() bool {
	rv := objc.CallMethod[bool](cc, objc.SEL("sharedColorPanelExists"))
	return rv
}

func (c_ ColorPanel) Mode() ColorPanelMode {
	rv := objc.CallMethod[ColorPanelMode](c_, objc.SEL("mode"))
	return rv
}

func (c_ ColorPanel) SetMode(value ColorPanelMode) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setMode:"), value)
}

func (c_ ColorPanel) AccessoryView() View {
	rv := objc.CallMethod[View](c_, objc.SEL("accessoryView"))
	return rv
}

func (c_ ColorPanel) SetAccessoryView(value IView) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setAccessoryView:"), objc.ExtractPtr(value))
}

func (c_ ColorPanel) IsContinuous() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isContinuous"))
	return rv
}

func (c_ ColorPanel) SetContinuous(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setContinuous:"), value)
}

func (c_ ColorPanel) ShowsAlpha() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("showsAlpha"))
	return rv
}

func (c_ ColorPanel) SetShowsAlpha(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setShowsAlpha:"), value)
}

func (c_ ColorPanel) Color() Color {
	rv := objc.CallMethod[Color](c_, objc.SEL("color"))
	return rv
}

func (c_ ColorPanel) SetColor(value IColor) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setColor:"), objc.ExtractPtr(value))
}

func (c_ ColorPanel) Alpha() float64 {
	rv := objc.CallMethod[float64](c_, objc.SEL("alpha"))
	return rv
}
