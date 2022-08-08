// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[FontPanel](fc, "windowWithContentViewController:", contentViewController)
	return rv
}

func (f_ FontPanel) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) FontPanel {
	rv := ffi.CallMethod[FontPanel](f_, "initWithContentRect:styleMask:backing:defer:", contentRect, style, backingStoreType, flag)
	return rv
}

func (f_ FontPanel) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) FontPanel {
	rv := ffi.CallMethod[FontPanel](f_, "initWithContentRect:styleMask:backing:defer:screen:", contentRect, style, backingStoreType, flag, screen)
	return rv
}

func (f_ FontPanel) Init() FontPanel {
	rv := ffi.CallMethod[FontPanel](f_, "init")
	return rv
}

func (fc _FontPanelClass) Alloc() FontPanel {
	rv := ffi.CallMethod[FontPanel](fc, "alloc")
	return rv
}

func (fc _FontPanelClass) New() FontPanel {
	rv := ffi.CallMethod[FontPanel](fc, "new")
	rv.Autorelease()
	return rv
}

func NewFontPanel() FontPanel {
	return FontPanelClass.New()
}

func (f_ FontPanel) ReloadDefaultFontFamilies() {
	ffi.CallMethod[ffi.Void](f_, "reloadDefaultFontFamilies")
}

func (f_ FontPanel) SetPanelFont_IsMultiple(fontObj IFont, flag bool) {
	ffi.CallMethod[ffi.Void](f_, "setPanelFont:isMultiple:", fontObj, flag)
}

func (f_ FontPanel) PanelConvertFont(fontObj IFont) Font {
	rv := ffi.CallMethod[Font](f_, "panelConvertFont:", fontObj)
	return rv
}

func (fc _FontPanelClass) SharedFontPanel() FontPanel {
	rv := ffi.CallMethod[FontPanel](fc, "sharedFontPanel")
	return rv
}

func (fc _FontPanelClass) SharedFontPanelExists() bool {
	rv := ffi.CallMethod[bool](fc, "sharedFontPanelExists")
	return rv
}

func (f_ FontPanel) IsEnabled() bool {
	rv := ffi.CallMethod[bool](f_, "isEnabled")
	return rv
}

func (f_ FontPanel) SetEnabled(value bool) {
	ffi.CallMethod[ffi.Void](f_, "setEnabled:", value)
}

func (f_ FontPanel) AccessoryView() View {
	rv := ffi.CallMethod[View](f_, "accessoryView")
	return rv
}

func (f_ FontPanel) SetAccessoryView(value IView) {
	ffi.CallMethod[ffi.Void](f_, "setAccessoryView:", value)
}

type FontChanging interface {
	ImplementsChangeFont() bool
	// optional
	ChangeFont(sender FontManager)
	ImplementsValidModesForFontPanel() bool
	// optional
	ValidModesForFontPanel(fontPanel FontPanel) FontPanelModeMask
}

type FontChangingWrapper struct {
	objc.Object
}

func (f_ *FontChangingWrapper) ImplementsChangeFont() bool {
	return f_.RespondsToSelector(objc.GetSelector("changeFont:"))
}

func (f_ FontChangingWrapper) ChangeFont(sender IFontManager) {
	ffi.CallMethod[ffi.Void](f_, "changeFont:", sender)
}

func (f_ *FontChangingWrapper) ImplementsValidModesForFontPanel() bool {
	return f_.RespondsToSelector(objc.GetSelector("validModesForFontPanel:"))
}

func (f_ FontChangingWrapper) ValidModesForFontPanel(fontPanel IFontPanel) FontPanelModeMask {
	rv := ffi.CallMethod[FontPanelModeMask](f_, "validModesForFontPanel:", fontPanel)
	return rv
}
