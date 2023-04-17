// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var AppearanceClass = _AppearanceClass{objc.GetClass("NSAppearance")}

type _AppearanceClass struct {
	objc.Class
}

type IAppearance interface {
	objc.IObject
	BestMatchFromAppearancesWithNames(appearances []AppearanceName) AppearanceName
	PerformAsCurrentDrawingAppearance(block func())
	Name() AppearanceName
	AllowsVibrancy() bool
}

type Appearance struct {
	objc.Object
}

func MakeAppearance(ptr unsafe.Pointer) Appearance {
	return Appearance{
		Object: objc.MakeObject(ptr),
	}
}

func (a_ Appearance) InitWithAppearanceNamed_Bundle(name AppearanceName, bundle foundation.IBundle) Appearance {
	rv := ffi.CallMethod[Appearance](a_, "initWithAppearanceNamed:bundle:", name, bundle)
	return rv
}

func (ac _AppearanceClass) Alloc() Appearance {
	rv := ffi.CallMethod[Appearance](ac, "alloc")
	return rv
}

func (ac _AppearanceClass) New() Appearance {
	rv := ffi.CallMethod[Appearance](ac, "new")
	rv.Autorelease()
	return rv
}

func NewAppearance() Appearance {
	return AppearanceClass.New()
}

func (a_ Appearance) Init() Appearance {
	rv := ffi.CallMethod[Appearance](a_, "init")
	return rv
}

func (ac _AppearanceClass) AppearanceNamed(name AppearanceName) Appearance {
	rv := ffi.CallMethod[Appearance](ac, "appearanceNamed:", name)
	return rv
}

func (a_ Appearance) BestMatchFromAppearancesWithNames(appearances []AppearanceName) AppearanceName {
	rv := ffi.CallMethod[AppearanceName](a_, "bestMatchFromAppearancesWithNames:", appearances)
	return rv
}

func (a_ Appearance) PerformAsCurrentDrawingAppearance(block func()) {
	ffi.CallMethod[ffi.Void](a_, "performAsCurrentDrawingAppearance:", block)
}

func (a_ Appearance) Name() AppearanceName {
	rv := ffi.CallMethod[AppearanceName](a_, "name")
	return rv
}

func (ac _AppearanceClass) CurrentDrawingAppearance() Appearance {
	rv := ffi.CallMethod[Appearance](ac, "currentDrawingAppearance")
	return rv
}

// deprecated
func (ac _AppearanceClass) CurrentAppearance() Appearance {
	rv := ffi.CallMethod[Appearance](ac, "currentAppearance")
	return rv
}

// deprecated
func (ac _AppearanceClass) SetCurrentAppearance(value IAppearance) {
	ffi.CallMethod[ffi.Void](ac, "setCurrentAppearance:", value)
}

func (a_ Appearance) AllowsVibrancy() bool {
	rv := ffi.CallMethod[bool](a_, "allowsVibrancy")
	return rv
}
