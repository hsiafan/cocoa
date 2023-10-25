// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var AnimationClass = _AnimationClass{objc.GetClass("CAAnimation")}

type _AnimationClass struct {
	objc.Class
}

type IAnimation interface {
	objc.IObject
	ShouldArchiveValueForKey(key string) bool
	IsRemovedOnCompletion() bool
	SetRemovedOnCompletion(value bool)
	TimingFunction() MediaTimingFunction
	SetTimingFunction(value IMediaTimingFunction)
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
	PreferredFrameRateRange() FrameRateRange
	SetPreferredFrameRateRange(value FrameRateRange)
}

type Animation struct {
	objc.Object
}

func MakeAnimation(ptr unsafe.Pointer) Animation {
	return Animation{
		Object: objc.MakeObject(ptr),
	}
}

func (ac _AnimationClass) Animation() Animation {
	rv := objc.CallMethod[Animation](ac, objc.SEL("animation"))
	return rv
}

func (ac _AnimationClass) Alloc() Animation {
	rv := objc.CallMethod[Animation](ac, objc.SEL("alloc"))
	return rv
}

func (ac _AnimationClass) New() Animation {
	rv := objc.CallMethod[Animation](ac, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewAnimation() Animation {
	return AnimationClass.New()
}

func (a_ Animation) Init() Animation {
	rv := objc.CallMethod[Animation](a_, objc.SEL("init"))
	return rv
}

func (ac _AnimationClass) DefaultValueForKey(key string) objc.Object {
	rv := objc.CallMethod[objc.Object](ac, objc.SEL("defaultValueForKey:"), key)
	return rv
}

func (a_ Animation) ShouldArchiveValueForKey(key string) bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("shouldArchiveValueForKey:"), key)
	return rv
}

func (a_ Animation) IsRemovedOnCompletion() bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("isRemovedOnCompletion"))
	return rv
}

func (a_ Animation) SetRemovedOnCompletion(value bool) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setRemovedOnCompletion:"), value)
}

func (a_ Animation) TimingFunction() MediaTimingFunction {
	rv := objc.CallMethod[MediaTimingFunction](a_, objc.SEL("timingFunction"))
	return rv
}

func (a_ Animation) SetTimingFunction(value IMediaTimingFunction) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setTimingFunction:"), objc.ExtractPtr(value))
}

func (a_ Animation) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](a_, objc.SEL("delegate"))
	return rv
}

func (a_ Animation) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}

func (a_ Animation) PreferredFrameRateRange() FrameRateRange {
	rv := objc.CallMethod[FrameRateRange](a_, objc.SEL("preferredFrameRateRange"))
	return rv
}

func (a_ Animation) SetPreferredFrameRateRange(value FrameRateRange) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setPreferredFrameRateRange:"), value)
}
