// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type AppearanceCustomization interface {
	ImplementsSetAppearance() bool
	// optional
	SetAppearance(value Appearance)
	ImplementsAppearance() bool
	// optional
	Appearance() IAppearance
	ImplementsEffectiveAppearance() bool
	// optional
	EffectiveAppearance() IAppearance
}

func WrapAppearanceCustomization(v AppearanceCustomization) objc.Object {
	return objc.WrapAsProtocol("NSAppearanceCustomization", v)
}

type AppearanceCustomizationBase struct {
}

func (p *AppearanceCustomizationBase) ImplementsSetAppearance() bool {
	return false
}

func (p *AppearanceCustomizationBase) SetAppearance(value Appearance) {
	panic("unimpemented protocol method")
}

func (p *AppearanceCustomizationBase) ImplementsAppearance() bool {
	return false
}

func (p *AppearanceCustomizationBase) Appearance() IAppearance {
	panic("unimpemented protocol method")
}

func (p *AppearanceCustomizationBase) ImplementsEffectiveAppearance() bool {
	return false
}

func (p *AppearanceCustomizationBase) EffectiveAppearance() IAppearance {
	panic("unimpemented protocol method")
}

type AppearanceCustomizationWrapper struct {
	objc.Object
}

func (a_ AppearanceCustomizationWrapper) SetAppearance(value IAppearance) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setAppearance:"), objc.ExtractPtr(value))
}

func (a_ AppearanceCustomizationWrapper) Appearance() Appearance {
	rv := objc.CallMethod[Appearance](a_, objc.GetSelector("appearance"))
	return rv
}

func (a_ AppearanceCustomizationWrapper) EffectiveAppearance() Appearance {
	rv := objc.CallMethod[Appearance](a_, objc.GetSelector("effectiveAppearance"))
	return rv
}
