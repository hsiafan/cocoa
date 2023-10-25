// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var SwitchClass = _SwitchClass{objc.GetClass("NSSwitch")}

type _SwitchClass struct {
	objc.Class
}

type ISwitch interface {
	IControl
	State() ControlStateValue
	SetState(value ControlStateValue)
}

type Switch struct {
	Control
}

func MakeSwitch(ptr unsafe.Pointer) Switch {
	return Switch{
		Control: MakeControl(ptr),
	}
}

func (s_ Switch) InitWithFrame(frameRect foundation.Rect) Switch {
	rv := objc.CallMethod[Switch](s_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (s_ Switch) Init() Switch {
	rv := objc.CallMethod[Switch](s_, objc.SEL("init"))
	return rv
}

func (sc _SwitchClass) Alloc() Switch {
	rv := objc.CallMethod[Switch](sc, objc.SEL("alloc"))
	return rv
}

func (sc _SwitchClass) New() Switch {
	rv := objc.CallMethod[Switch](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewSwitch() Switch {
	return SwitchClass.New()
}

func (s_ Switch) State() ControlStateValue {
	rv := objc.CallMethod[ControlStateValue](s_, objc.SEL("state"))
	return rv
}

func (s_ Switch) SetState(value ControlStateValue) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setState:"), value)
}
