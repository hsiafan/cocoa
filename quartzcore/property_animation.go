// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var PropertyAnimationClass = _PropertyAnimationClass{objc.GetClass("CAPropertyAnimation")}

type _PropertyAnimationClass struct {
	objc.Class
}

type IPropertyAnimation interface {
	IAnimation
	KeyPath() string
	SetKeyPath(value string)
	IsCumulative() bool
	SetCumulative(value bool)
	IsAdditive() bool
	SetAdditive(value bool)
	ValueFunction() ValueFunction
	SetValueFunction(value IValueFunction)
}

type PropertyAnimation struct {
	Animation
}

func MakePropertyAnimation(ptr unsafe.Pointer) PropertyAnimation {
	return PropertyAnimation{
		Animation: MakeAnimation(ptr),
	}
}

func (pc _PropertyAnimationClass) AnimationWithKeyPath(path string) PropertyAnimation {
	rv := ffi.CallMethod[PropertyAnimation](pc, "animationWithKeyPath:", path)
	return rv
}

func (pc _PropertyAnimationClass) Animation() PropertyAnimation {
	rv := ffi.CallMethod[PropertyAnimation](pc, "animation")
	return rv
}

func (pc _PropertyAnimationClass) Alloc() PropertyAnimation {
	rv := ffi.CallMethod[PropertyAnimation](pc, "alloc")
	return rv
}

func (pc _PropertyAnimationClass) New() PropertyAnimation {
	rv := ffi.CallMethod[PropertyAnimation](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPropertyAnimation() PropertyAnimation {
	return PropertyAnimationClass.New()
}

func (p_ PropertyAnimation) Init() PropertyAnimation {
	rv := ffi.CallMethod[PropertyAnimation](p_, "init")
	return rv
}

func (p_ PropertyAnimation) KeyPath() string {
	rv := ffi.CallMethod[string](p_, "keyPath")
	return rv
}

func (p_ PropertyAnimation) SetKeyPath(value string) {
	ffi.CallMethod[ffi.Void](p_, "setKeyPath:", value)
}

func (p_ PropertyAnimation) IsCumulative() bool {
	rv := ffi.CallMethod[bool](p_, "isCumulative")
	return rv
}

func (p_ PropertyAnimation) SetCumulative(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setCumulative:", value)
}

func (p_ PropertyAnimation) IsAdditive() bool {
	rv := ffi.CallMethod[bool](p_, "isAdditive")
	return rv
}

func (p_ PropertyAnimation) SetAdditive(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setAdditive:", value)
}

func (p_ PropertyAnimation) ValueFunction() ValueFunction {
	rv := ffi.CallMethod[ValueFunction](p_, "valueFunction")
	return rv
}

func (p_ PropertyAnimation) SetValueFunction(value IValueFunction) {
	ffi.CallMethod[ffi.Void](p_, "setValueFunction:", value)
}
