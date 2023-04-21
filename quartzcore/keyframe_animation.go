// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var KeyframeAnimationClass = _KeyframeAnimationClass{objc.GetClass("CAKeyframeAnimation")}

type _KeyframeAnimationClass struct {
	objc.Class
}

type IKeyframeAnimation interface {
	IPropertyAnimation
	Values() []objc.Object
	SetValues(value []objc.IObject)
	Path() coregraphics.PathRef
	SetPath(value coregraphics.PathRef)
	KeyTimes() []foundation.Number
	SetKeyTimes(value []foundation.INumber)
	TimingFunctions() []MediaTimingFunction
	SetTimingFunctions(value []IMediaTimingFunction)
	CalculationMode() AnimationCalculationMode
	SetCalculationMode(value AnimationCalculationMode)
	RotationMode() AnimationRotationMode
	SetRotationMode(value AnimationRotationMode)
	TensionValues() []foundation.Number
	SetTensionValues(value []foundation.INumber)
	ContinuityValues() []foundation.Number
	SetContinuityValues(value []foundation.INumber)
	BiasValues() []foundation.Number
	SetBiasValues(value []foundation.INumber)
}

type KeyframeAnimation struct {
	PropertyAnimation
}

func MakeKeyframeAnimation(ptr unsafe.Pointer) KeyframeAnimation {
	return KeyframeAnimation{
		PropertyAnimation: MakePropertyAnimation(ptr),
	}
}

func (kc _KeyframeAnimationClass) AnimationWithKeyPath(path string) KeyframeAnimation {
	rv := objc.CallMethod[KeyframeAnimation](kc, "animationWithKeyPath:", path)
	return rv
}

func (kc _KeyframeAnimationClass) Animation() KeyframeAnimation {
	rv := objc.CallMethod[KeyframeAnimation](kc, "animation")
	return rv
}

func (kc _KeyframeAnimationClass) Alloc() KeyframeAnimation {
	rv := objc.CallMethod[KeyframeAnimation](kc, "alloc")
	return rv
}

func (kc _KeyframeAnimationClass) New() KeyframeAnimation {
	rv := objc.CallMethod[KeyframeAnimation](kc, "new")
	rv.Autorelease()
	return rv
}

func NewKeyframeAnimation() KeyframeAnimation {
	return KeyframeAnimationClass.New()
}

func (k_ KeyframeAnimation) Init() KeyframeAnimation {
	rv := objc.CallMethod[KeyframeAnimation](k_, "init")
	return rv
}

func (k_ KeyframeAnimation) Values() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](k_, "values")
	return rv
}

func (k_ KeyframeAnimation) SetValues(value []objc.IObject) {
	objc.CallMethod[objc.Void](k_, "setValues:", value)
}

func (k_ KeyframeAnimation) Path() coregraphics.PathRef {
	rv := objc.CallMethod[coregraphics.PathRef](k_, "path")
	return rv
}

func (k_ KeyframeAnimation) SetPath(value coregraphics.PathRef) {
	objc.CallMethod[objc.Void](k_, "setPath:", value)
}

func (k_ KeyframeAnimation) KeyTimes() []foundation.Number {
	rv := objc.CallMethod[[]foundation.Number](k_, "keyTimes")
	return rv
}

func (k_ KeyframeAnimation) SetKeyTimes(value []foundation.INumber) {
	objc.CallMethod[objc.Void](k_, "setKeyTimes:", value)
}

func (k_ KeyframeAnimation) TimingFunctions() []MediaTimingFunction {
	rv := objc.CallMethod[[]MediaTimingFunction](k_, "timingFunctions")
	return rv
}

func (k_ KeyframeAnimation) SetTimingFunctions(value []IMediaTimingFunction) {
	objc.CallMethod[objc.Void](k_, "setTimingFunctions:", value)
}

func (k_ KeyframeAnimation) CalculationMode() AnimationCalculationMode {
	rv := objc.CallMethod[AnimationCalculationMode](k_, "calculationMode")
	return rv
}

func (k_ KeyframeAnimation) SetCalculationMode(value AnimationCalculationMode) {
	objc.CallMethod[objc.Void](k_, "setCalculationMode:", value)
}

func (k_ KeyframeAnimation) RotationMode() AnimationRotationMode {
	rv := objc.CallMethod[AnimationRotationMode](k_, "rotationMode")
	return rv
}

func (k_ KeyframeAnimation) SetRotationMode(value AnimationRotationMode) {
	objc.CallMethod[objc.Void](k_, "setRotationMode:", value)
}

func (k_ KeyframeAnimation) TensionValues() []foundation.Number {
	rv := objc.CallMethod[[]foundation.Number](k_, "tensionValues")
	return rv
}

func (k_ KeyframeAnimation) SetTensionValues(value []foundation.INumber) {
	objc.CallMethod[objc.Void](k_, "setTensionValues:", value)
}

func (k_ KeyframeAnimation) ContinuityValues() []foundation.Number {
	rv := objc.CallMethod[[]foundation.Number](k_, "continuityValues")
	return rv
}

func (k_ KeyframeAnimation) SetContinuityValues(value []foundation.INumber) {
	objc.CallMethod[objc.Void](k_, "setContinuityValues:", value)
}

func (k_ KeyframeAnimation) BiasValues() []foundation.Number {
	rv := objc.CallMethod[[]foundation.Number](k_, "biasValues")
	return rv
}

func (k_ KeyframeAnimation) SetBiasValues(value []foundation.INumber) {
	objc.CallMethod[objc.Void](k_, "setBiasValues:", value)
}
