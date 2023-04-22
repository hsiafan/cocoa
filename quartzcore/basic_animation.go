// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

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
	rv := objc.CallMethod[BasicAnimation](bc, objc.GetSelector("animationWithKeyPath:"), path)
	return rv
}

func (bc _BasicAnimationClass) Animation() BasicAnimation {
	rv := objc.CallMethod[BasicAnimation](bc, objc.GetSelector("animation"))
	return rv
}

func (bc _BasicAnimationClass) Alloc() BasicAnimation {
	rv := objc.CallMethod[BasicAnimation](bc, objc.GetSelector("alloc"))
	return rv
}

func (bc _BasicAnimationClass) New() BasicAnimation {
	rv := objc.CallMethod[BasicAnimation](bc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewBasicAnimation() BasicAnimation {
	return BasicAnimationClass.New()
}

func (b_ BasicAnimation) Init() BasicAnimation {
	rv := objc.CallMethod[BasicAnimation](b_, objc.GetSelector("init"))
	return rv
}

func (b_ BasicAnimation) FromValue() objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("fromValue"))
	return rv
}

func (b_ BasicAnimation) SetFromValue(value objc.IObject) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setFromValue:"), value)
}

func (b_ BasicAnimation) ToValue() objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("toValue"))
	return rv
}

func (b_ BasicAnimation) SetToValue(value objc.IObject) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setToValue:"), value)
}

func (b_ BasicAnimation) ByValue() objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("byValue"))
	return rv
}

func (b_ BasicAnimation) SetByValue(value objc.IObject) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setByValue:"), value)
}
