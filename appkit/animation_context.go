// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[AnimationContext](ac, "alloc")
	return rv
}

func (ac _AnimationContextClass) New() AnimationContext {
	rv := objc.CallMethod[AnimationContext](ac, "new")
	rv.Autorelease()
	return rv
}

func NewAnimationContext() AnimationContext {
	return AnimationContextClass.New()
}

func (a_ AnimationContext) Init() AnimationContext {
	rv := objc.CallMethod[AnimationContext](a_, "init")
	return rv
}

func (ac _AnimationContextClass) BeginGrouping() {
	objc.CallMethod[objc.Void](ac, "beginGrouping")
}

func (ac _AnimationContextClass) EndGrouping() {
	objc.CallMethod[objc.Void](ac, "endGrouping")
}

func (ac _AnimationContextClass) RunAnimationGroup_CompletionHandler(changes func(context AnimationContext), completionHandler func()) {
	objc.CallMethod[objc.Void](ac, "runAnimationGroup:completionHandler:", changes, completionHandler)
}

func (ac _AnimationContextClass) RunAnimationGroup(changes func(context AnimationContext)) {
	objc.CallMethod[objc.Void](ac, "runAnimationGroup:", changes)
}

func (ac _AnimationContextClass) CurrentContext() AnimationContext {
	rv := objc.CallMethod[AnimationContext](ac, "currentContext")
	return rv
}

func (a_ AnimationContext) CompletionHandler() func() {
	rv := objc.CallMethod[func()](a_, "completionHandler")
	return rv
}

func (a_ AnimationContext) SetCompletionHandler(value func()) {
	objc.CallMethod[objc.Void](a_, "setCompletionHandler:", value)
}

func (a_ AnimationContext) Duration() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](a_, "duration")
	return rv
}

func (a_ AnimationContext) SetDuration(value foundation.TimeInterval) {
	objc.CallMethod[objc.Void](a_, "setDuration:", value)
}

func (a_ AnimationContext) TimingFunction() quartzcore.MediaTimingFunction {
	rv := objc.CallMethod[quartzcore.MediaTimingFunction](a_, "timingFunction")
	return rv
}

func (a_ AnimationContext) SetTimingFunction(value quartzcore.IMediaTimingFunction) {
	objc.CallMethod[objc.Void](a_, "setTimingFunction:", value)
}

func (a_ AnimationContext) AllowsImplicitAnimation() bool {
	rv := objc.CallMethod[bool](a_, "allowsImplicitAnimation")
	return rv
}

func (a_ AnimationContext) SetAllowsImplicitAnimation(value bool) {
	objc.CallMethod[objc.Void](a_, "setAllowsImplicitAnimation:", value)
}
