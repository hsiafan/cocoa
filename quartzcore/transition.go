// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var TransitionClass = _TransitionClass{objc.GetClass("CATransition")}

type _TransitionClass struct {
	objc.Class
}

type ITransition interface {
	IAnimation
	StartProgress() float32
	SetStartProgress(value float32)
	EndProgress() float32
	SetEndProgress(value float32)
	Type() TransitionType
	SetType(value TransitionType)
	Subtype() TransitionSubtype
	SetSubtype(value TransitionSubtype)
	Filter() objc.Object
	SetFilter(value objc.IObject)
}

type Transition struct {
	Animation
}

func MakeTransition(ptr unsafe.Pointer) Transition {
	return Transition{
		Animation: MakeAnimation(ptr),
	}
}

func (tc _TransitionClass) Animation() Transition {
	rv := ffi.CallMethod[Transition](tc, "animation")
	return rv
}

func (tc _TransitionClass) Alloc() Transition {
	rv := ffi.CallMethod[Transition](tc, "alloc")
	return rv
}

func (t_ Transition) Init() Transition {
	rv := ffi.CallMethod[Transition](t_, "init")
	return rv
}

func (tc _TransitionClass) New() Transition {
	rv := ffi.CallMethod[Transition](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTransition() Transition {
	return TransitionClass.New()
}

func (t_ Transition) StartProgress() float32 {
	rv := ffi.CallMethod[float32](t_, "startProgress")
	return rv
}

func (t_ Transition) SetStartProgress(value float32) {
	ffi.CallMethod[ffi.Void](t_, "setStartProgress:", value)
}

func (t_ Transition) EndProgress() float32 {
	rv := ffi.CallMethod[float32](t_, "endProgress")
	return rv
}

func (t_ Transition) SetEndProgress(value float32) {
	ffi.CallMethod[ffi.Void](t_, "setEndProgress:", value)
}

func (t_ Transition) Type() TransitionType {
	rv := ffi.CallMethod[TransitionType](t_, "type")
	return rv
}

func (t_ Transition) SetType(value TransitionType) {
	ffi.CallMethod[ffi.Void](t_, "setType:", value)
}

func (t_ Transition) Subtype() TransitionSubtype {
	rv := ffi.CallMethod[TransitionSubtype](t_, "subtype")
	return rv
}

func (t_ Transition) SetSubtype(value TransitionSubtype) {
	ffi.CallMethod[ffi.Void](t_, "setSubtype:", value)
}

func (t_ Transition) Filter() objc.Object {
	rv := ffi.CallMethod[objc.Object](t_, "filter")
	return rv
}

func (t_ Transition) SetFilter(value objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "setFilter:", value)
}
