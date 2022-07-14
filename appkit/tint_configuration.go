// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var TintConfigurationClass = _TintConfigurationClass{objc.GetClass("NSTintConfiguration")}

type _TintConfigurationClass struct {
	objc.Class
}

type ITintConfiguration interface {
	objc.IObject
	AdaptsToUserAccentColor() bool
	BaseTintColor() Color
	EquivalentContentTintColor() Color
}

type TintConfiguration struct {
	objc.Object
}

func MakeTintConfiguration(ptr unsafe.Pointer) TintConfiguration {
	return TintConfiguration{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TintConfigurationClass) TintConfigurationWithFixedColor(color IColor) TintConfiguration {
	rv := ffi.CallMethod[TintConfiguration](tc, "tintConfigurationWithFixedColor:", color)
	return rv
}

func (tc _TintConfigurationClass) TintConfigurationWithPreferredColor(color IColor) TintConfiguration {
	rv := ffi.CallMethod[TintConfiguration](tc, "tintConfigurationWithPreferredColor:", color)
	return rv
}

func (tc _TintConfigurationClass) Alloc() TintConfiguration {
	rv := ffi.CallMethod[TintConfiguration](tc, "alloc")
	return rv
}

func (t_ TintConfiguration) Init() TintConfiguration {
	rv := ffi.CallMethod[TintConfiguration](t_, "init")
	rv.Autorelease()
	return rv
}

func (tc _TintConfigurationClass) New() TintConfiguration {
	rv := ffi.CallMethod[TintConfiguration](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTintConfiguration() TintConfiguration {
	return TintConfigurationClass.New()
}

func (t_ TintConfiguration) AdaptsToUserAccentColor() bool {
	rv := ffi.CallMethod[bool](t_, "adaptsToUserAccentColor")
	return rv
}

func (tc _TintConfigurationClass) DefaultTintConfiguration() TintConfiguration {
	rv := ffi.CallMethod[TintConfiguration](tc, "defaultTintConfiguration")
	return rv
}

func (tc _TintConfigurationClass) MonochromeTintConfiguration() TintConfiguration {
	rv := ffi.CallMethod[TintConfiguration](tc, "monochromeTintConfiguration")
	return rv
}

func (t_ TintConfiguration) BaseTintColor() Color {
	rv := ffi.CallMethod[Color](t_, "baseTintColor")
	return rv
}

func (t_ TintConfiguration) EquivalentContentTintColor() Color {
	rv := ffi.CallMethod[Color](t_, "equivalentContentTintColor")
	return rv
}
