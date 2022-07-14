// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[Shadow](s_, "init")
	rv.Autorelease()
	return rv
}

func (sc _ShadowClass) Alloc() Shadow {
	rv := ffi.CallMethod[Shadow](sc, "alloc")
	return rv
}

func (sc _ShadowClass) New() Shadow {
	rv := ffi.CallMethod[Shadow](sc, "new")
	rv.Autorelease()
	return rv
}

func NewShadow() Shadow {
	return ShadowClass.New()
}

func (s_ Shadow) Set() {
	ffi.CallMethod[ffi.Void](s_, "set")
}

func (s_ Shadow) ShadowOffset() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](s_, "shadowOffset")
	return rv
}

func (s_ Shadow) SetShadowOffset(value foundation.Size) {
	ffi.CallMethod[ffi.Void](s_, "setShadowOffset:", value)
}

func (s_ Shadow) ShadowBlurRadius() float64 {
	rv := ffi.CallMethod[float64](s_, "shadowBlurRadius")
	return rv
}

func (s_ Shadow) SetShadowBlurRadius(value float64) {
	ffi.CallMethod[ffi.Void](s_, "setShadowBlurRadius:", value)
}

func (s_ Shadow) ShadowColor() Color {
	rv := ffi.CallMethod[Color](s_, "shadowColor")
	return rv
}

func (s_ Shadow) SetShadowColor(value IColor) {
	ffi.CallMethod[ffi.Void](s_, "setShadowColor:", value)
}
