// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
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

type AppearanceCustomizationWrapper struct {
	objc.Object
}

func (a_ *AppearanceCustomizationWrapper) ImplementsSetAppearance() bool {
	return a_.RespondsToSelector(objc.GetSelector("setAppearance:"))
}

func (a_ AppearanceCustomizationWrapper) SetAppearance(value IAppearance) {
	ffi.CallMethod[ffi.Void](a_, "setAppearance:", value)
}

func (a_ *AppearanceCustomizationWrapper) ImplementsAppearance() bool {
	return a_.RespondsToSelector(objc.GetSelector("appearance"))
}

func (a_ AppearanceCustomizationWrapper) Appearance() Appearance {
	rv := ffi.CallMethod[Appearance](a_, "appearance")
	return rv
}

func (a_ *AppearanceCustomizationWrapper) ImplementsEffectiveAppearance() bool {
	return a_.RespondsToSelector(objc.GetSelector("effectiveAppearance"))
}

func (a_ AppearanceCustomizationWrapper) EffectiveAppearance() Appearance {
	rv := ffi.CallMethod[Appearance](a_, "effectiveAppearance")
	return rv
}
