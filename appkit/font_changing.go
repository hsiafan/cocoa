// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

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
