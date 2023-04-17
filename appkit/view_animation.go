// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[ViewAnimation](v_, "initWithViewAnimations:", viewAnimations)
	return rv
}

func (v_ ViewAnimation) InitWithDuration_AnimationCurve(duration foundation.TimeInterval, animationCurve AnimationCurve) ViewAnimation {
	rv := ffi.CallMethod[ViewAnimation](v_, "initWithDuration:animationCurve:", duration, animationCurve)
	return rv
}

func (vc _ViewAnimationClass) Alloc() ViewAnimation {
	rv := ffi.CallMethod[ViewAnimation](vc, "alloc")
	return rv
}

func (vc _ViewAnimationClass) New() ViewAnimation {
	rv := ffi.CallMethod[ViewAnimation](vc, "new")
	rv.Autorelease()
	return rv
}

func NewViewAnimation() ViewAnimation {
	return ViewAnimationClass.New()
}

func (v_ ViewAnimation) Init() ViewAnimation {
	rv := ffi.CallMethod[ViewAnimation](v_, "init")
	return rv
}

func (v_ ViewAnimation) ViewAnimations() []map[ViewAnimationKey]objc.Object {
	rv := ffi.CallMethod[[]map[ViewAnimationKey]objc.Object](v_, "viewAnimations")
	return rv
}

func (v_ ViewAnimation) SetViewAnimations(value []map[ViewAnimationKey]objc.IObject) {
	ffi.CallMethod[ffi.Void](v_, "setViewAnimations:", value)
}
