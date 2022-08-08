// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
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

func (a_ Animation) Init() Animation {
	rv := ffi.CallMethod[Animation](a_, "init")
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

func (a_ Animation) Delegate() AnimationDelegateWrapper {
	rv := ffi.CallMethod[AnimationDelegateWrapper](a_, "delegate")
	return rv
}

func (a_ Animation) SetDelegate(value AnimationDelegate) {
	po := ffi.CreateProtocol(value)
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

type AnimationDelegate interface {
	ImplementsAnimationDidStart() bool
	// optional
	AnimationDidStart(anim Animation)
	ImplementsAnimationDidStop_Finished() bool
	// optional
	AnimationDidStop_Finished(anim Animation, flag bool)
}

type AnimationDelegateImpl struct {
	_AnimationDidStart         func(anim Animation)
	_AnimationDidStop_Finished func(anim Animation, flag bool)
}

func (di *AnimationDelegateImpl) ImplementsAnimationDidStart() bool {
	return di._AnimationDidStart != nil
}

func (di *AnimationDelegateImpl) SetAnimationDidStart(f func(anim Animation)) {
	di._AnimationDidStart = f
}

func (di *AnimationDelegateImpl) AnimationDidStart(anim Animation) {
	di._AnimationDidStart(anim)
}
func (di *AnimationDelegateImpl) ImplementsAnimationDidStop_Finished() bool {
	return di._AnimationDidStop_Finished != nil
}

func (di *AnimationDelegateImpl) SetAnimationDidStop_Finished(f func(anim Animation, flag bool)) {
	di._AnimationDidStop_Finished = f
}

func (di *AnimationDelegateImpl) AnimationDidStop_Finished(anim Animation, flag bool) {
	di._AnimationDidStop_Finished(anim, flag)
}

type AnimationDelegateWrapper struct {
	objc.Object
}

func (a_ *AnimationDelegateWrapper) ImplementsAnimationDidStart() bool {
	return a_.RespondsToSelector(objc.GetSelector("animationDidStart:"))
}

func (a_ AnimationDelegateWrapper) AnimationDidStart(anim IAnimation) {
	ffi.CallMethod[ffi.Void](a_, "animationDidStart:", anim)
}

func (a_ *AnimationDelegateWrapper) ImplementsAnimationDidStop_Finished() bool {
	return a_.RespondsToSelector(objc.GetSelector("animationDidStop:finished:"))
}

func (a_ AnimationDelegateWrapper) AnimationDidStop_Finished(anim IAnimation, flag bool) {
	ffi.CallMethod[ffi.Void](a_, "animationDidStop:finished:", anim, flag)
}

var PropertyAnimationClass = _PropertyAnimationClass{objc.GetClass("CAPropertyAnimation")}

type _PropertyAnimationClass struct {
	objc.Class
}

type IPropertyAnimation interface {
	IAnimation
	KeyPath() string
	SetKeyPath(value string)
	IsCumulative() bool
	SetCumulative(value bool)
	IsAdditive() bool
	SetAdditive(value bool)
	ValueFunction() ValueFunction
	SetValueFunction(value IValueFunction)
}

type PropertyAnimation struct {
	Animation
}

func MakePropertyAnimation(ptr unsafe.Pointer) PropertyAnimation {
	return PropertyAnimation{
		Animation: MakeAnimation(ptr),
	}
}

func (pc _PropertyAnimationClass) AnimationWithKeyPath(path string) PropertyAnimation {
	rv := ffi.CallMethod[PropertyAnimation](pc, "animationWithKeyPath:", path)
	return rv
}

func (pc _PropertyAnimationClass) Animation() PropertyAnimation {
	rv := ffi.CallMethod[PropertyAnimation](pc, "animation")
	return rv
}

func (pc _PropertyAnimationClass) Alloc() PropertyAnimation {
	rv := ffi.CallMethod[PropertyAnimation](pc, "alloc")
	return rv
}

func (p_ PropertyAnimation) Init() PropertyAnimation {
	rv := ffi.CallMethod[PropertyAnimation](p_, "init")
	return rv
}

func (pc _PropertyAnimationClass) New() PropertyAnimation {
	rv := ffi.CallMethod[PropertyAnimation](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPropertyAnimation() PropertyAnimation {
	return PropertyAnimationClass.New()
}

func (p_ PropertyAnimation) KeyPath() string {
	rv := ffi.CallMethod[string](p_, "keyPath")
	return rv
}

func (p_ PropertyAnimation) SetKeyPath(value string) {
	ffi.CallMethod[ffi.Void](p_, "setKeyPath:", value)
}

func (p_ PropertyAnimation) IsCumulative() bool {
	rv := ffi.CallMethod[bool](p_, "isCumulative")
	return rv
}

func (p_ PropertyAnimation) SetCumulative(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setCumulative:", value)
}

func (p_ PropertyAnimation) IsAdditive() bool {
	rv := ffi.CallMethod[bool](p_, "isAdditive")
	return rv
}

func (p_ PropertyAnimation) SetAdditive(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setAdditive:", value)
}

func (p_ PropertyAnimation) ValueFunction() ValueFunction {
	rv := ffi.CallMethod[ValueFunction](p_, "valueFunction")
	return rv
}

func (p_ PropertyAnimation) SetValueFunction(value IValueFunction) {
	ffi.CallMethod[ffi.Void](p_, "setValueFunction:", value)
}

var BasicAnimationClass = _BasicAnimationClass{objc.GetClass("CABasicAnimation")}

type _BasicAnimationClass struct {
	objc.Class
}

type IBasicAnimation interface {
	IPropertyAnimation
	FromValue() objc.Object
	SetFromValue(value objc.IObject)
	ToValue() objc.Object
	SetToValue(value objc.IObject)
	ByValue() objc.Object
	SetByValue(value objc.IObject)
}

type BasicAnimation struct {
	PropertyAnimation
}

func MakeBasicAnimation(ptr unsafe.Pointer) BasicAnimation {
	return BasicAnimation{
		PropertyAnimation: MakePropertyAnimation(ptr),
	}
}

func (bc _BasicAnimationClass) AnimationWithKeyPath(path string) BasicAnimation {
	rv := ffi.CallMethod[BasicAnimation](bc, "animationWithKeyPath:", path)
	return rv
}

func (bc _BasicAnimationClass) Animation() BasicAnimation {
	rv := ffi.CallMethod[BasicAnimation](bc, "animation")
	return rv
}

func (bc _BasicAnimationClass) Alloc() BasicAnimation {
	rv := ffi.CallMethod[BasicAnimation](bc, "alloc")
	return rv
}

func (b_ BasicAnimation) Init() BasicAnimation {
	rv := ffi.CallMethod[BasicAnimation](b_, "init")
	return rv
}

func (bc _BasicAnimationClass) New() BasicAnimation {
	rv := ffi.CallMethod[BasicAnimation](bc, "new")
	rv.Autorelease()
	return rv
}

func NewBasicAnimation() BasicAnimation {
	return BasicAnimationClass.New()
}

func (b_ BasicAnimation) FromValue() objc.Object {
	rv := ffi.CallMethod[objc.Object](b_, "fromValue")
	return rv
}

func (b_ BasicAnimation) SetFromValue(value objc.IObject) {
	ffi.CallMethod[ffi.Void](b_, "setFromValue:", value)
}

func (b_ BasicAnimation) ToValue() objc.Object {
	rv := ffi.CallMethod[objc.Object](b_, "toValue")
	return rv
}

func (b_ BasicAnimation) SetToValue(value objc.IObject) {
	ffi.CallMethod[ffi.Void](b_, "setToValue:", value)
}

func (b_ BasicAnimation) ByValue() objc.Object {
	rv := ffi.CallMethod[objc.Object](b_, "byValue")
	return rv
}

func (b_ BasicAnimation) SetByValue(value objc.IObject) {
	ffi.CallMethod[ffi.Void](b_, "setByValue:", value)
}

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
	rv := ffi.CallMethod[KeyframeAnimation](kc, "animationWithKeyPath:", path)
	return rv
}

func (kc _KeyframeAnimationClass) Animation() KeyframeAnimation {
	rv := ffi.CallMethod[KeyframeAnimation](kc, "animation")
	return rv
}

func (kc _KeyframeAnimationClass) Alloc() KeyframeAnimation {
	rv := ffi.CallMethod[KeyframeAnimation](kc, "alloc")
	return rv
}

func (k_ KeyframeAnimation) Init() KeyframeAnimation {
	rv := ffi.CallMethod[KeyframeAnimation](k_, "init")
	return rv
}

func (kc _KeyframeAnimationClass) New() KeyframeAnimation {
	rv := ffi.CallMethod[KeyframeAnimation](kc, "new")
	rv.Autorelease()
	return rv
}

func NewKeyframeAnimation() KeyframeAnimation {
	return KeyframeAnimationClass.New()
}

func (k_ KeyframeAnimation) Values() []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](k_, "values")
	return rv
}

func (k_ KeyframeAnimation) SetValues(value []objc.IObject) {
	ffi.CallMethod[ffi.Void](k_, "setValues:", value)
}

func (k_ KeyframeAnimation) Path() coregraphics.PathRef {
	rv := ffi.CallMethod[coregraphics.PathRef](k_, "path")
	return rv
}

func (k_ KeyframeAnimation) SetPath(value coregraphics.PathRef) {
	ffi.CallMethod[ffi.Void](k_, "setPath:", value)
}

func (k_ KeyframeAnimation) KeyTimes() []foundation.Number {
	rv := ffi.CallMethod[[]foundation.Number](k_, "keyTimes")
	return rv
}

func (k_ KeyframeAnimation) SetKeyTimes(value []foundation.INumber) {
	ffi.CallMethod[ffi.Void](k_, "setKeyTimes:", value)
}

func (k_ KeyframeAnimation) CalculationMode() AnimationCalculationMode {
	rv := ffi.CallMethod[AnimationCalculationMode](k_, "calculationMode")
	return rv
}

func (k_ KeyframeAnimation) SetCalculationMode(value AnimationCalculationMode) {
	ffi.CallMethod[ffi.Void](k_, "setCalculationMode:", value)
}

func (k_ KeyframeAnimation) RotationMode() AnimationRotationMode {
	rv := ffi.CallMethod[AnimationRotationMode](k_, "rotationMode")
	return rv
}

func (k_ KeyframeAnimation) SetRotationMode(value AnimationRotationMode) {
	ffi.CallMethod[ffi.Void](k_, "setRotationMode:", value)
}

func (k_ KeyframeAnimation) TensionValues() []foundation.Number {
	rv := ffi.CallMethod[[]foundation.Number](k_, "tensionValues")
	return rv
}

func (k_ KeyframeAnimation) SetTensionValues(value []foundation.INumber) {
	ffi.CallMethod[ffi.Void](k_, "setTensionValues:", value)
}

func (k_ KeyframeAnimation) ContinuityValues() []foundation.Number {
	rv := ffi.CallMethod[[]foundation.Number](k_, "continuityValues")
	return rv
}

func (k_ KeyframeAnimation) SetContinuityValues(value []foundation.INumber) {
	ffi.CallMethod[ffi.Void](k_, "setContinuityValues:", value)
}

func (k_ KeyframeAnimation) BiasValues() []foundation.Number {
	rv := ffi.CallMethod[[]foundation.Number](k_, "biasValues")
	return rv
}

func (k_ KeyframeAnimation) SetBiasValues(value []foundation.INumber) {
	ffi.CallMethod[ffi.Void](k_, "setBiasValues:", value)
}

var SpringAnimationClass = _SpringAnimationClass{objc.GetClass("CASpringAnimation")}

type _SpringAnimationClass struct {
	objc.Class
}

type ISpringAnimation interface {
	IBasicAnimation
	Damping() float64
	SetDamping(value float64)
	InitialVelocity() float64
	SetInitialVelocity(value float64)
	Mass() float64
	SetMass(value float64)
	Stiffness() float64
	SetStiffness(value float64)
}

type SpringAnimation struct {
	BasicAnimation
}

func MakeSpringAnimation(ptr unsafe.Pointer) SpringAnimation {
	return SpringAnimation{
		BasicAnimation: MakeBasicAnimation(ptr),
	}
}

func (sc _SpringAnimationClass) AnimationWithKeyPath(path string) SpringAnimation {
	rv := ffi.CallMethod[SpringAnimation](sc, "animationWithKeyPath:", path)
	return rv
}

func (sc _SpringAnimationClass) Animation() SpringAnimation {
	rv := ffi.CallMethod[SpringAnimation](sc, "animation")
	return rv
}

func (sc _SpringAnimationClass) Alloc() SpringAnimation {
	rv := ffi.CallMethod[SpringAnimation](sc, "alloc")
	return rv
}

func (s_ SpringAnimation) Init() SpringAnimation {
	rv := ffi.CallMethod[SpringAnimation](s_, "init")
	return rv
}

func (sc _SpringAnimationClass) New() SpringAnimation {
	rv := ffi.CallMethod[SpringAnimation](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSpringAnimation() SpringAnimation {
	return SpringAnimationClass.New()
}

func (s_ SpringAnimation) Damping() float64 {
	rv := ffi.CallMethod[float64](s_, "damping")
	return rv
}

func (s_ SpringAnimation) SetDamping(value float64) {
	ffi.CallMethod[ffi.Void](s_, "setDamping:", value)
}

func (s_ SpringAnimation) InitialVelocity() float64 {
	rv := ffi.CallMethod[float64](s_, "initialVelocity")
	return rv
}

func (s_ SpringAnimation) SetInitialVelocity(value float64) {
	ffi.CallMethod[ffi.Void](s_, "setInitialVelocity:", value)
}

func (s_ SpringAnimation) Mass() float64 {
	rv := ffi.CallMethod[float64](s_, "mass")
	return rv
}

func (s_ SpringAnimation) SetMass(value float64) {
	ffi.CallMethod[ffi.Void](s_, "setMass:", value)
}

func (s_ SpringAnimation) Stiffness() float64 {
	rv := ffi.CallMethod[float64](s_, "stiffness")
	return rv
}

func (s_ SpringAnimation) SetStiffness(value float64) {
	ffi.CallMethod[ffi.Void](s_, "setStiffness:", value)
}
