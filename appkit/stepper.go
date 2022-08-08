// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var StepperClass = _StepperClass{objc.GetClass("NSStepper")}

type _StepperClass struct {
	objc.Class
}

type IStepper interface {
	IControl
	MaxValue() float64
	SetMaxValue(value float64)
	MinValue() float64
	SetMinValue(value float64)
	Increment() float64
	SetIncrement(value float64)
	Autorepeat() bool
	SetAutorepeat(value bool)
	ValueWraps() bool
	SetValueWraps(value bool)
}

type Stepper struct {
	Control
}

func MakeStepper(ptr unsafe.Pointer) Stepper {
	return Stepper{
		Control: MakeControl(ptr),
	}
}

func (s_ Stepper) InitWithFrame(frameRect foundation.Rect) Stepper {
	rv := ffi.CallMethod[Stepper](s_, "initWithFrame:", frameRect)
	return rv
}

func (s_ Stepper) Init() Stepper {
	rv := ffi.CallMethod[Stepper](s_, "init")
	return rv
}

func (sc _StepperClass) Alloc() Stepper {
	rv := ffi.CallMethod[Stepper](sc, "alloc")
	return rv
}

func (sc _StepperClass) New() Stepper {
	rv := ffi.CallMethod[Stepper](sc, "new")
	rv.Autorelease()
	return rv
}

func NewStepper() Stepper {
	return StepperClass.New()
}

func (s_ Stepper) MaxValue() float64 {
	rv := ffi.CallMethod[float64](s_, "maxValue")
	return rv
}

func (s_ Stepper) SetMaxValue(value float64) {
	ffi.CallMethod[ffi.Void](s_, "setMaxValue:", value)
}

func (s_ Stepper) MinValue() float64 {
	rv := ffi.CallMethod[float64](s_, "minValue")
	return rv
}

func (s_ Stepper) SetMinValue(value float64) {
	ffi.CallMethod[ffi.Void](s_, "setMinValue:", value)
}

func (s_ Stepper) Increment() float64 {
	rv := ffi.CallMethod[float64](s_, "increment")
	return rv
}

func (s_ Stepper) SetIncrement(value float64) {
	ffi.CallMethod[ffi.Void](s_, "setIncrement:", value)
}

func (s_ Stepper) Autorepeat() bool {
	rv := ffi.CallMethod[bool](s_, "autorepeat")
	return rv
}

func (s_ Stepper) SetAutorepeat(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setAutorepeat:", value)
}

func (s_ Stepper) ValueWraps() bool {
	rv := ffi.CallMethod[bool](s_, "valueWraps")
	return rv
}

func (s_ Stepper) SetValueWraps(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setValueWraps:", value)
}
