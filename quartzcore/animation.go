// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	Delegate() AnimationDelegateWrapper
	SetDelegate(value AnimationDelegate)
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
	rv := ffi.CallMethod[Animation](ac, "animation")
	return rv
}

func (ac _AnimationClass) Alloc() Animation {
	rv := ffi.CallMethod[Animation](ac, "alloc")
	return rv
}

func (ac _AnimationClass) New() Animation {
	rv := ffi.CallMethod[Animation](ac, "new")
	rv.Autorelease()
	return rv
}

func NewAnimation() Animation {
	return AnimationClass.New()
}

func (a_ Animation) Init() Animation {
	rv := ffi.CallMethod[Animation](a_, "init")
	return rv
}

func (ac _AnimationClass) DefaultValueForKey(key string) objc.Object {
	rv := ffi.CallMethod[objc.Object](ac, "defaultValueForKey:", key)
	return rv
}

func (a_ Animation) ShouldArchiveValueForKey(key string) bool {
	rv := ffi.CallMethod[bool](a_, "shouldArchiveValueForKey:", key)
	return rv
}

func (a_ Animation) IsRemovedOnCompletion() bool {
	rv := ffi.CallMethod[bool](a_, "isRemovedOnCompletion")
	return rv
}

func (a_ Animation) SetRemovedOnCompletion(value bool) {
	ffi.CallMethod[ffi.Void](a_, "setRemovedOnCompletion:", value)
}

func (a_ Animation) TimingFunction() MediaTimingFunction {
	rv := ffi.CallMethod[MediaTimingFunction](a_, "timingFunction")
	return rv
}

func (a_ Animation) SetTimingFunction(value IMediaTimingFunction) {
	ffi.CallMethod[ffi.Void](a_, "setTimingFunction:", value)
}

func (a_ Animation) Delegate() AnimationDelegateWrapper {
	rv := ffi.CallMethod[AnimationDelegateWrapper](a_, "delegate")
	return rv
}

func (a_ Animation) SetDelegate(value AnimationDelegate) {
	po := ffi.CreateProtocol("CAAnimationDelegate", value)
	defer po.Release()
	ffi.CallMethod[ffi.Void](a_, "setDelegate:", po)
}

func (a_ Animation) PreferredFrameRateRange() FrameRateRange {
	rv := ffi.CallMethod[FrameRateRange](a_, "preferredFrameRateRange")
	return rv
}

func (a_ Animation) SetPreferredFrameRateRange(value FrameRateRange) {
	ffi.CallMethod[ffi.Void](a_, "setPreferredFrameRateRange:", value)
}
