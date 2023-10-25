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
