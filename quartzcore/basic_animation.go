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
	rv := objc.CallMethod[BasicAnimation](bc, objc.SEL("animationWithKeyPath:"), path)
	return rv
}

func (bc _BasicAnimationClass) Animation() BasicAnimation {
	rv := objc.CallMethod[BasicAnimation](bc, objc.SEL("animation"))
	return rv
}

func (bc _BasicAnimationClass) Alloc() BasicAnimation {
	rv := objc.CallMethod[BasicAnimation](bc, objc.SEL("alloc"))
	return rv
}

func (bc _BasicAnimationClass) New() BasicAnimation {
	rv := objc.CallMethod[BasicAnimation](bc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewBasicAnimation() BasicAnimation {
	return BasicAnimationClass.New()
}

func (b_ BasicAnimation) Init() BasicAnimation {
	rv := objc.CallMethod[BasicAnimation](b_, objc.SEL("init"))
	return rv
}

func (b_ BasicAnimation) FromValue() objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.SEL("fromValue"))
	return rv
}

func (b_ BasicAnimation) SetFromValue(value objc.IObject) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setFromValue:"), objc.ExtractPtr(value))
}

func (b_ BasicAnimation) ToValue() objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.SEL("toValue"))
	return rv
}

func (b_ BasicAnimation) SetToValue(value objc.IObject) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setToValue:"), objc.ExtractPtr(value))
}

func (b_ BasicAnimation) ByValue() objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.SEL("byValue"))
	return rv
}

func (b_ BasicAnimation) SetByValue(value objc.IObject) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setByValue:"), objc.ExtractPtr(value))
}
