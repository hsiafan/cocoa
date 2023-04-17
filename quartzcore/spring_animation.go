// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var SpringAnimationClass = _SpringAnimationClass{objc.GetClass("CASpringAnimation")}

type _SpringAnimationClass struct {
	objc.Class
}

type ISpringAnimation interface {
	IBasicAnimation
	Damping() float64
	SetDamping(value float64)
	InitialVelocity() float64
	SetInitialVelocity(value float64)
	Mass() float64
	SetMass(value float64)
	Stiffness() float64
	SetStiffness(value float64)
}

type SpringAnimation struct {
	BasicAnimation
}

func MakeSpringAnimation(ptr unsafe.Pointer) SpringAnimation {
	return SpringAnimation{
		BasicAnimation: MakeBasicAnimation(ptr),
	}
}

func (sc _SpringAnimationClass) AnimationWithKeyPath(path string) SpringAnimation {
	rv := ffi.CallMethod[SpringAnimation](sc, "animationWithKeyPath:", path)
	return rv
}

func (sc _SpringAnimationClass) Animation() SpringAnimation {
	rv := ffi.CallMethod[SpringAnimation](sc, "animation")
	return rv
}

func (sc _SpringAnimationClass) Alloc() SpringAnimation {
	rv := ffi.CallMethod[SpringAnimation](sc, "alloc")
	return rv
}

func (sc _SpringAnimationClass) New() SpringAnimation {
	rv := ffi.CallMethod[SpringAnimation](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSpringAnimation() SpringAnimation {
	return SpringAnimationClass.New()
}

func (s_ SpringAnimation) Init() SpringAnimation {
	rv := ffi.CallMethod[SpringAnimation](s_, "init")
	return rv
}

func (s_ SpringAnimation) Damping() float64 {
	rv := ffi.CallMethod[float64](s_, "damping")
	return rv
}

func (s_ SpringAnimation) SetDamping(value float64) {
	ffi.CallMethod[ffi.Void](s_, "setDamping:", value)
}

func (s_ SpringAnimation) InitialVelocity() float64 {
	rv := ffi.CallMethod[float64](s_, "initialVelocity")
	return rv
}

func (s_ SpringAnimation) SetInitialVelocity(value float64) {
	ffi.CallMethod[ffi.Void](s_, "setInitialVelocity:", value)
}

func (s_ SpringAnimation) Mass() float64 {
	rv := ffi.CallMethod[float64](s_, "mass")
	return rv
}

func (s_ SpringAnimation) SetMass(value float64) {
	ffi.CallMethod[ffi.Void](s_, "setMass:", value)
}

func (s_ SpringAnimation) Stiffness() float64 {
	rv := ffi.CallMethod[float64](s_, "stiffness")
	return rv
}

func (s_ SpringAnimation) SetStiffness(value float64) {
	ffi.CallMethod[ffi.Void](s_, "setStiffness:", value)
}
