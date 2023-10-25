// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[Stepper](s_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (s_ Stepper) Init() Stepper {
	rv := objc.CallMethod[Stepper](s_, objc.SEL("init"))
	return rv
}

func (sc _StepperClass) Alloc() Stepper {
	rv := objc.CallMethod[Stepper](sc, objc.SEL("alloc"))
	return rv
}

func (sc _StepperClass) New() Stepper {
	rv := objc.CallMethod[Stepper](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewStepper() Stepper {
	return StepperClass.New()
}

func (s_ Stepper) MaxValue() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("maxValue"))
	return rv
}

func (s_ Stepper) SetMaxValue(value float64) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setMaxValue:"), value)
}

func (s_ Stepper) MinValue() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("minValue"))
	return rv
}

func (s_ Stepper) SetMinValue(value float64) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setMinValue:"), value)
}

func (s_ Stepper) Increment() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("increment"))
	return rv
}

func (s_ Stepper) SetIncrement(value float64) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setIncrement:"), value)
}

func (s_ Stepper) Autorepeat() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("autorepeat"))
	return rv
}

func (s_ Stepper) SetAutorepeat(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setAutorepeat:"), value)
}

func (s_ Stepper) ValueWraps() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("valueWraps"))
	return rv
}

func (s_ Stepper) SetValueWraps(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setValueWraps:"), value)
}
