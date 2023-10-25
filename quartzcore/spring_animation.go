// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

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
	rv := objc.CallMethod[SpringAnimation](sc, objc.SEL("animationWithKeyPath:"), path)
	return rv
}

func (sc _SpringAnimationClass) Animation() SpringAnimation {
	rv := objc.CallMethod[SpringAnimation](sc, objc.SEL("animation"))
	return rv
}

func (sc _SpringAnimationClass) Alloc() SpringAnimation {
	rv := objc.CallMethod[SpringAnimation](sc, objc.SEL("alloc"))
	return rv
}

func (sc _SpringAnimationClass) New() SpringAnimation {
	rv := objc.CallMethod[SpringAnimation](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewSpringAnimation() SpringAnimation {
	return SpringAnimationClass.New()
}

func (s_ SpringAnimation) Init() SpringAnimation {
	rv := objc.CallMethod[SpringAnimation](s_, objc.SEL("init"))
	return rv
}

func (s_ SpringAnimation) Damping() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("damping"))
	return rv
}

func (s_ SpringAnimation) SetDamping(value float64) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setDamping:"), value)
}

func (s_ SpringAnimation) InitialVelocity() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("initialVelocity"))
	return rv
}

func (s_ SpringAnimation) SetInitialVelocity(value float64) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setInitialVelocity:"), value)
}

func (s_ SpringAnimation) Mass() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("mass"))
	return rv
}

func (s_ SpringAnimation) SetMass(value float64) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setMass:"), value)
}

func (s_ SpringAnimation) Stiffness() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("stiffness"))
	return rv
}

func (s_ SpringAnimation) SetStiffness(value float64) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setStiffness:"), value)
}
