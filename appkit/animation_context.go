// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"github.com/hsiafan/cocoa/quartzcore"
)

var AnimationContextClass = _AnimationContextClass{objc.GetClass("NSAnimationContext")}

type _AnimationContextClass struct {
	objc.Class
}

type IAnimationContext interface {
	objc.IObject
	CompletionHandler() func()
	SetCompletionHandler(value func())
	Duration() foundation.TimeInterval
	SetDuration(value foundation.TimeInterval)
	TimingFunction() quartzcore.MediaTimingFunction
	SetTimingFunction(value quartzcore.IMediaTimingFunction)
	AllowsImplicitAnimation() bool
	SetAllowsImplicitAnimation(value bool)
}

type AnimationContext struct {
	objc.Object
}

func MakeAnimationContext(ptr unsafe.Pointer) AnimationContext {
	return AnimationContext{
		Object: objc.MakeObject(ptr),
	}
}

func (ac _AnimationContextClass) Alloc() AnimationContext {
	rv := ffi.CallMethod[AnimationContext](ac, "alloc")
	return rv
}

func (ac _AnimationContextClass) New() AnimationContext {
	rv := ffi.CallMethod[AnimationContext](ac, "new")
	rv.Autorelease()
	return rv
}

func NewAnimationContext() AnimationContext {
	return AnimationContextClass.New()
}

func (a_ AnimationContext) Init() AnimationContext {
	rv := ffi.CallMethod[AnimationContext](a_, "init")
	return rv
}

func (ac _AnimationContextClass) BeginGrouping() {
	ffi.CallMethod[ffi.Void](ac, "beginGrouping")
}

func (ac _AnimationContextClass) EndGrouping() {
	ffi.CallMethod[ffi.Void](ac, "endGrouping")
}

func (ac _AnimationContextClass) RunAnimationGroup_CompletionHandler(changes func(context AnimationContext), completionHandler func()) {
	ffi.CallMethod[ffi.Void](ac, "runAnimationGroup:completionHandler:", changes, completionHandler)
}

func (ac _AnimationContextClass) RunAnimationGroup(changes func(context AnimationContext)) {
	ffi.CallMethod[ffi.Void](ac, "runAnimationGroup:", changes)
}

func (ac _AnimationContextClass) CurrentContext() AnimationContext {
	rv := ffi.CallMethod[AnimationContext](ac, "currentContext")
	return rv
}

func (a_ AnimationContext) CompletionHandler() func() {
	rv := ffi.CallMethod[func()](a_, "completionHandler")
	return rv
}

func (a_ AnimationContext) SetCompletionHandler(value func()) {
	ffi.CallMethod[ffi.Void](a_, "setCompletionHandler:", value)
}

func (a_ AnimationContext) Duration() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](a_, "duration")
	return rv
}

func (a_ AnimationContext) SetDuration(value foundation.TimeInterval) {
	ffi.CallMethod[ffi.Void](a_, "setDuration:", value)
}

func (a_ AnimationContext) TimingFunction() quartzcore.MediaTimingFunction {
	rv := ffi.CallMethod[quartzcore.MediaTimingFunction](a_, "timingFunction")
	return rv
}

func (a_ AnimationContext) SetTimingFunction(value quartzcore.IMediaTimingFunction) {
	ffi.CallMethod[ffi.Void](a_, "setTimingFunction:", value)
}

func (a_ AnimationContext) AllowsImplicitAnimation() bool {
	rv := ffi.CallMethod[bool](a_, "allowsImplicitAnimation")
	return rv
}

func (a_ AnimationContext) SetAllowsImplicitAnimation(value bool) {
	ffi.CallMethod[ffi.Void](a_, "setAllowsImplicitAnimation:", value)
}
