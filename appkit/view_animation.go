// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ViewAnimationClass = _ViewAnimationClass{objc.GetClass("NSViewAnimation")}

type _ViewAnimationClass struct {
	objc.Class
}

type IViewAnimation interface {
	IAnimation
	ViewAnimations() []map[ViewAnimationKey]objc.Object
	SetViewAnimations(value []map[ViewAnimationKey]objc.IObject)
}

type ViewAnimation struct {
	Animation
}

func MakeViewAnimation(ptr unsafe.Pointer) ViewAnimation {
	return ViewAnimation{
		Animation: MakeAnimation(ptr),
	}
}

func (v_ ViewAnimation) InitWithViewAnimations(viewAnimations []map[ViewAnimationKey]objc.IObject) ViewAnimation {
	rv := objc.CallMethod[ViewAnimation](v_, objc.SEL("initWithViewAnimations:"), viewAnimations)
	return rv
}

func (v_ ViewAnimation) InitWithDuration_AnimationCurve(duration foundation.TimeInterval, animationCurve AnimationCurve) ViewAnimation {
	rv := objc.CallMethod[ViewAnimation](v_, objc.SEL("initWithDuration:animationCurve:"), duration, animationCurve)
	return rv
}

func (vc _ViewAnimationClass) Alloc() ViewAnimation {
	rv := objc.CallMethod[ViewAnimation](vc, objc.SEL("alloc"))
	return rv
}

func (vc _ViewAnimationClass) New() ViewAnimation {
	rv := objc.CallMethod[ViewAnimation](vc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewViewAnimation() ViewAnimation {
	return ViewAnimationClass.New()
}

func (v_ ViewAnimation) Init() ViewAnimation {
	rv := objc.CallMethod[ViewAnimation](v_, objc.SEL("init"))
	return rv
}

func (v_ ViewAnimation) ViewAnimations() []map[ViewAnimationKey]objc.Object {
	rv := objc.CallMethod[[]map[ViewAnimationKey]objc.Object](v_, objc.SEL("viewAnimations"))
	return rv
}

func (v_ ViewAnimation) SetViewAnimations(value []map[ViewAnimationKey]objc.IObject) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setViewAnimations:"), value)
}
