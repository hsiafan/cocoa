// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var BasicAnimationClass = _BasicAnimationClass{objc.GetClass("CABasicAnimation")}

type _BasicAnimationClass struct {
	objc.Class
}

type IBasicAnimation interface {
	IPropertyAnimation
	FromValue() objc.Object
	SetFromValue(value objc.IObject)
	ToValue() objc.Object
	SetToValue(value objc.IObject)
	ByValue() objc.Object
	SetByValue(value objc.IObject)
}

type BasicAnimation struct {
	PropertyAnimation
}

func MakeBasicAnimation(ptr unsafe.Pointer) BasicAnimation {
	return BasicAnimation{
		PropertyAnimation: MakePropertyAnimation(ptr),
	}
}

func (bc _BasicAnimationClass) AnimationWithKeyPath(path string) BasicAnimation {
	rv := ffi.CallMethod[BasicAnimation](bc, "animationWithKeyPath:", path)
	return rv
}

func (bc _BasicAnimationClass) Animation() BasicAnimation {
	rv := ffi.CallMethod[BasicAnimation](bc, "animation")
	return rv
}

func (bc _BasicAnimationClass) Alloc() BasicAnimation {
	rv := ffi.CallMethod[BasicAnimation](bc, "alloc")
	return rv
}

func (bc _BasicAnimationClass) New() BasicAnimation {
	rv := ffi.CallMethod[BasicAnimation](bc, "new")
	rv.Autorelease()
	return rv
}

func NewBasicAnimation() BasicAnimation {
	return BasicAnimationClass.New()
}

func (b_ BasicAnimation) Init() BasicAnimation {
	rv := ffi.CallMethod[BasicAnimation](b_, "init")
	return rv
}

func (b_ BasicAnimation) FromValue() objc.Object {
	rv := ffi.CallMethod[objc.Object](b_, "fromValue")
	return rv
}

func (b_ BasicAnimation) SetFromValue(value objc.IObject) {
	ffi.CallMethod[ffi.Void](b_, "setFromValue:", value)
}

func (b_ BasicAnimation) ToValue() objc.Object {
	rv := ffi.CallMethod[objc.Object](b_, "toValue")
	return rv
}

func (b_ BasicAnimation) SetToValue(value objc.IObject) {
	ffi.CallMethod[ffi.Void](b_, "setToValue:", value)
}

func (b_ BasicAnimation) ByValue() objc.Object {
	rv := ffi.CallMethod[objc.Object](b_, "byValue")
	return rv
}

func (b_ BasicAnimation) SetByValue(value objc.IObject) {
	ffi.CallMethod[ffi.Void](b_, "setByValue:", value)
}
