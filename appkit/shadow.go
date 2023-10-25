// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ShadowClass = _ShadowClass{objc.GetClass("NSShadow")}

type _ShadowClass struct {
	objc.Class
}

type IShadow interface {
	objc.IObject
	Set()
	ShadowOffset() foundation.Size
	SetShadowOffset(value foundation.Size)
	ShadowBlurRadius() float64
	SetShadowBlurRadius(value float64)
	ShadowColor() Color
	SetShadowColor(value IColor)
}

type Shadow struct {
	objc.Object
}

func MakeShadow(ptr unsafe.Pointer) Shadow {
	return Shadow{
		Object: objc.MakeObject(ptr),
	}
}

func (s_ Shadow) Init() Shadow {
	rv := objc.CallMethod[Shadow](s_, objc.SEL("init"))
	return rv
}

func (sc _ShadowClass) Alloc() Shadow {
	rv := objc.CallMethod[Shadow](sc, objc.SEL("alloc"))
	return rv
}

func (sc _ShadowClass) New() Shadow {
	rv := objc.CallMethod[Shadow](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewShadow() Shadow {
	return ShadowClass.New()
}

func (s_ Shadow) Set() {
	objc.CallMethod[objc.Void](s_, objc.SEL("set"))
}

func (s_ Shadow) ShadowOffset() foundation.Size {
	rv := objc.CallMethod[foundation.Size](s_, objc.SEL("shadowOffset"))
	return rv
}

func (s_ Shadow) SetShadowOffset(value foundation.Size) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setShadowOffset:"), value)
}

func (s_ Shadow) ShadowBlurRadius() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("shadowBlurRadius"))
	return rv
}

func (s_ Shadow) SetShadowBlurRadius(value float64) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setShadowBlurRadius:"), value)
}

func (s_ Shadow) ShadowColor() Color {
	rv := objc.CallMethod[Color](s_, objc.SEL("shadowColor"))
	return rv
}

func (s_ Shadow) SetShadowColor(value IColor) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setShadowColor:"), objc.ExtractPtr(value))
}
