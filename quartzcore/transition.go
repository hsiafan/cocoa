// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

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
	rv := objc.CallMethod[Transition](tc, objc.SEL("animation"))
	return rv
}

func (tc _TransitionClass) Alloc() Transition {
	rv := objc.CallMethod[Transition](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TransitionClass) New() Transition {
	rv := objc.CallMethod[Transition](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTransition() Transition {
	return TransitionClass.New()
}

func (t_ Transition) Init() Transition {
	rv := objc.CallMethod[Transition](t_, objc.SEL("init"))
	return rv
}

func (t_ Transition) StartProgress() float32 {
	rv := objc.CallMethod[float32](t_, objc.SEL("startProgress"))
	return rv
}

func (t_ Transition) SetStartProgress(value float32) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setStartProgress:"), value)
}

func (t_ Transition) EndProgress() float32 {
	rv := objc.CallMethod[float32](t_, objc.SEL("endProgress"))
	return rv
}

func (t_ Transition) SetEndProgress(value float32) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setEndProgress:"), value)
}

func (t_ Transition) Type() TransitionType {
	rv := objc.CallMethod[TransitionType](t_, objc.SEL("type"))
	return rv
}

func (t_ Transition) SetType(value TransitionType) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setType:"), value)
}

func (t_ Transition) Subtype() TransitionSubtype {
	rv := objc.CallMethod[TransitionSubtype](t_, objc.SEL("subtype"))
	return rv
}

func (t_ Transition) SetSubtype(value TransitionSubtype) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSubtype:"), value)
}

func (t_ Transition) Filter() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.SEL("filter"))
	return rv
}

func (t_ Transition) SetFilter(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setFilter:"), objc.ExtractPtr(value))
}
