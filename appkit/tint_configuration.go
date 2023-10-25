// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[TintConfiguration](tc, objc.SEL("tintConfigurationWithFixedColor:"), objc.ExtractPtr(color))
	return rv
}

func (tc _TintConfigurationClass) TintConfigurationWithPreferredColor(color IColor) TintConfiguration {
	rv := objc.CallMethod[TintConfiguration](tc, objc.SEL("tintConfigurationWithPreferredColor:"), objc.ExtractPtr(color))
	return rv
}

func (tc _TintConfigurationClass) Alloc() TintConfiguration {
	rv := objc.CallMethod[TintConfiguration](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TintConfigurationClass) New() TintConfiguration {
	rv := objc.CallMethod[TintConfiguration](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTintConfiguration() TintConfiguration {
	return TintConfigurationClass.New()
}

func (t_ TintConfiguration) Init() TintConfiguration {
	rv := objc.CallMethod[TintConfiguration](t_, objc.SEL("init"))
	return rv
}

func (t_ TintConfiguration) AdaptsToUserAccentColor() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("adaptsToUserAccentColor"))
	return rv
}

func (tc _TintConfigurationClass) DefaultTintConfiguration() TintConfiguration {
	rv := objc.CallMethod[TintConfiguration](tc, objc.SEL("defaultTintConfiguration"))
	return rv
}

func (tc _TintConfigurationClass) MonochromeTintConfiguration() TintConfiguration {
	rv := objc.CallMethod[TintConfiguration](tc, objc.SEL("monochromeTintConfiguration"))
	return rv
}

func (t_ TintConfiguration) BaseTintColor() Color {
	rv := objc.CallMethod[Color](t_, objc.SEL("baseTintColor"))
	return rv
}

func (t_ TintConfiguration) EquivalentContentTintColor() Color {
	rv := objc.CallMethod[Color](t_, objc.SEL("equivalentContentTintColor"))
	return rv
}
