// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[Switch](s_, "initWithFrame:", frameRect)
	return rv
}

func (s_ Switch) Init() Switch {
	rv := ffi.CallMethod[Switch](s_, "init")
	return rv
}

func (sc _SwitchClass) Alloc() Switch {
	rv := ffi.CallMethod[Switch](sc, "alloc")
	return rv
}

func (sc _SwitchClass) New() Switch {
	rv := ffi.CallMethod[Switch](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSwitch() Switch {
	return SwitchClass.New()
}

func (s_ Switch) State() ControlStateValue {
	rv := ffi.CallMethod[ControlStateValue](s_, "state")
	return rv
}

func (s_ Switch) SetState(value ControlStateValue) {
	ffi.CallMethod[ffi.Void](s_, "setState:", value)
}
