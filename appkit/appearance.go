// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[Appearance](a_, objc.SEL("initWithAppearanceNamed:bundle:"), name, objc.ExtractPtr(bundle))
	return rv
}

func (ac _AppearanceClass) Alloc() Appearance {
	rv := objc.CallMethod[Appearance](ac, objc.SEL("alloc"))
	return rv
}

func (ac _AppearanceClass) New() Appearance {
	rv := objc.CallMethod[Appearance](ac, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewAppearance() Appearance {
	return AppearanceClass.New()
}

func (a_ Appearance) Init() Appearance {
	rv := objc.CallMethod[Appearance](a_, objc.SEL("init"))
	return rv
}

func (ac _AppearanceClass) AppearanceNamed(name AppearanceName) Appearance {
	rv := objc.CallMethod[Appearance](ac, objc.SEL("appearanceNamed:"), name)
	return rv
}

func (a_ Appearance) BestMatchFromAppearancesWithNames(appearances []AppearanceName) AppearanceName {
	rv := objc.CallMethod[AppearanceName](a_, objc.SEL("bestMatchFromAppearancesWithNames:"), appearances)
	return rv
}

func (a_ Appearance) PerformAsCurrentDrawingAppearance(block func()) {
	objc.CallMethod[objc.Void](a_, objc.SEL("performAsCurrentDrawingAppearance:"), block)
}

func (a_ Appearance) Name() AppearanceName {
	rv := objc.CallMethod[AppearanceName](a_, objc.SEL("name"))
	return rv
}

func (ac _AppearanceClass) CurrentDrawingAppearance() Appearance {
	rv := objc.CallMethod[Appearance](ac, objc.SEL("currentDrawingAppearance"))
	return rv
}

// deprecated
func (ac _AppearanceClass) CurrentAppearance() Appearance {
	rv := objc.CallMethod[Appearance](ac, objc.SEL("currentAppearance"))
	return rv
}

// deprecated
func (ac _AppearanceClass) SetCurrentAppearance(value IAppearance) {
	objc.CallMethod[objc.Void](ac, objc.SEL("setCurrentAppearance:"), objc.ExtractPtr(value))
}

func (a_ Appearance) AllowsVibrancy() bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("allowsVibrancy"))
	return rv
}
