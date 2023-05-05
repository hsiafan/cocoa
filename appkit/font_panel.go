// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var FontPanelClass = _FontPanelClass{objc.GetClass("NSFontPanel")}

type _FontPanelClass struct {
	objc.Class
}

type IFontPanel interface {
	IPanel
	ReloadDefaultFontFamilies()
	SetPanelFont_IsMultiple(fontObj IFont, flag bool)
	PanelConvertFont(fontObj IFont) Font
	IsEnabled() bool
	SetEnabled(value bool)
	AccessoryView() View
	SetAccessoryView(value IView)
}

type FontPanel struct {
	Panel
}

func MakeFontPanel(ptr unsafe.Pointer) FontPanel {
	return FontPanel{
		Panel: MakePanel(ptr),
	}
}

func (fc _FontPanelClass) WindowWithContentViewController(contentViewController IViewController) FontPanel {
	rv := objc.CallMethod[FontPanel](fc, objc.GetSelector("windowWithContentViewController:"), objc.ExtractPtr(contentViewController))
	return rv
}

func (f_ FontPanel) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) FontPanel {
	rv := objc.CallMethod[FontPanel](f_, objc.GetSelector("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return rv
}

func (f_ FontPanel) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) FontPanel {
	rv := objc.CallMethod[FontPanel](f_, objc.GetSelector("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.ExtractPtr(screen))
	return rv
}

func (f_ FontPanel) Init() FontPanel {
	rv := objc.CallMethod[FontPanel](f_, objc.GetSelector("init"))
	return rv
}

func (fc _FontPanelClass) Alloc() FontPanel {
	rv := objc.CallMethod[FontPanel](fc, objc.GetSelector("alloc"))
	return rv
}

func (fc _FontPanelClass) New() FontPanel {
	rv := objc.CallMethod[FontPanel](fc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewFontPanel() FontPanel {
	return FontPanelClass.New()
}

func (f_ FontPanel) ReloadDefaultFontFamilies() {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("reloadDefaultFontFamilies"))
}

func (f_ FontPanel) SetPanelFont_IsMultiple(fontObj IFont, flag bool) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setPanelFont:isMultiple:"), objc.ExtractPtr(fontObj), flag)
}

func (f_ FontPanel) PanelConvertFont(fontObj IFont) Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("panelConvertFont:"), objc.ExtractPtr(fontObj))
	return rv
}

func (fc _FontPanelClass) SharedFontPanel() FontPanel {
	rv := objc.CallMethod[FontPanel](fc, objc.GetSelector("sharedFontPanel"))
	return rv
}

func (fc _FontPanelClass) SharedFontPanelExists() bool {
	rv := objc.CallMethod[bool](fc, objc.GetSelector("sharedFontPanelExists"))
	return rv
}

func (f_ FontPanel) IsEnabled() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("isEnabled"))
	return rv
}

func (f_ FontPanel) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setEnabled:"), value)
}

func (f_ FontPanel) AccessoryView() View {
	rv := objc.CallMethod[View](f_, objc.GetSelector("accessoryView"))
	return rv
}

func (f_ FontPanel) SetAccessoryView(value IView) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setAccessoryView:"), objc.ExtractPtr(value))
}
