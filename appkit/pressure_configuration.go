// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var PressureConfigurationClass = _PressureConfigurationClass{objc.GetClass("NSPressureConfiguration")}

type _PressureConfigurationClass struct {
	objc.Class
}

type IPressureConfiguration interface {
	objc.IObject
	Set()
	PressureBehavior() PressureBehavior
}

type PressureConfiguration struct {
	objc.Object
}

func MakePressureConfiguration(ptr unsafe.Pointer) PressureConfiguration {
	return PressureConfiguration{
		Object: objc.MakeObject(ptr),
	}
}

func (p_ PressureConfiguration) InitWithPressureBehavior(pressureBehavior PressureBehavior) PressureConfiguration {
	rv := objc.CallMethod[PressureConfiguration](p_, objc.GetSelector("initWithPressureBehavior:"), pressureBehavior)
	return rv
}

func (pc _PressureConfigurationClass) Alloc() PressureConfiguration {
	rv := objc.CallMethod[PressureConfiguration](pc, objc.GetSelector("alloc"))
	return rv
}

func (pc _PressureConfigurationClass) New() PressureConfiguration {
	rv := objc.CallMethod[PressureConfiguration](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPressureConfiguration() PressureConfiguration {
	return PressureConfigurationClass.New()
}

func (p_ PressureConfiguration) Init() PressureConfiguration {
	rv := objc.CallMethod[PressureConfiguration](p_, objc.GetSelector("init"))
	return rv
}

func (p_ PressureConfiguration) Set() {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("set"))
}

func (p_ PressureConfiguration) PressureBehavior() PressureBehavior {
	rv := objc.CallMethod[PressureBehavior](p_, objc.GetSelector("pressureBehavior"))
	return rv
}
