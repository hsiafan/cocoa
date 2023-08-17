// auto-generated code, do not modify
package appkit

import (
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

func WrapFontChanging(v FontChanging) objc.Object {
	return objc.WrapAsProtocol("NSFontChanging", v)
}

type FontChangingBase struct {
}

func (p *FontChangingBase) ImplementsChangeFont() bool {
	return false
}

func (p *FontChangingBase) ChangeFont(sender FontManager) {
	panic("unimpemented protocol method")
}

func (p *FontChangingBase) ImplementsValidModesForFontPanel() bool {
	return false
}

func (p *FontChangingBase) ValidModesForFontPanel(fontPanel FontPanel) FontPanelModeMask {
	panic("unimpemented protocol method")
}

type FontChangingWrapper struct {
	objc.Object
}

func (f_ FontChangingWrapper) ChangeFont(sender IFontManager) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("changeFont:"), objc.ExtractPtr(sender))
}

func (f_ FontChangingWrapper) ValidModesForFontPanel(fontPanel IFontPanel) FontPanelModeMask {
	rv := objc.CallMethod[FontPanelModeMask](f_, objc.GetSelector("validModesForFontPanel:"), objc.ExtractPtr(fontPanel))
	return rv
}
