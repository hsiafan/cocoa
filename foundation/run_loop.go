// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var RunLoopClass = _RunLoopClass{objc.GetClass("NSRunLoop")}

type _RunLoopClass struct {
	objc.Class
}

type IRunLoop interface {
	objc.IObject
	LimitDateForMode(mode RunLoopMode) Date
	AddTimer_ForMode(timer ITimer, mode RunLoopMode)
	AddPort_ForMode(aPort IPort, mode RunLoopMode)
	RemovePort_ForMode(aPort IPort, mode RunLoopMode)
	// deprecated
	ConfigureAsServer()
	Run()
	RunMode_BeforeDate(mode RunLoopMode, limitDate IDate) bool
	RunUntilDate(limitDate IDate)
	AcceptInputForMode_BeforeDate(mode RunLoopMode, limitDate IDate)
	PerformBlock(block func())
	PerformInModes_Block(modes []RunLoopMode, block func())
	PerformSelector_Target_Argument_Order_Modes(aSelector objc.Selector, target objc.IObject, arg objc.IObject, order uint, modes []RunLoopMode)
	CancelPerformSelector_Target_Argument(aSelector objc.Selector, target objc.IObject, arg objc.IObject)
	CancelPerformSelectorsWithTarget(target objc.IObject)
	CurrentMode() RunLoopMode
}

type RunLoop struct {
	objc.Object
}

func MakeRunLoop(ptr unsafe.Pointer) RunLoop {
	return RunLoop{
		Object: objc.MakeObject(ptr),
	}
}

func (rc _RunLoopClass) Alloc() RunLoop {
	rv := objc.CallMethod[RunLoop](rc, "alloc")
	return rv
}

func (rc _RunLoopClass) New() RunLoop {
	rv := objc.CallMethod[RunLoop](rc, "new")
	rv.Autorelease()
	return rv
}

func NewRunLoop() RunLoop {
	return RunLoopClass.New()
}

func (r_ RunLoop) Init() RunLoop {
	rv := objc.CallMethod[RunLoop](r_, "init")
	return rv
}

func (r_ RunLoop) LimitDateForMode(mode RunLoopMode) Date {
	rv := objc.CallMethod[Date](r_, "limitDateForMode:", mode)
	return rv
}

func (r_ RunLoop) AddTimer_ForMode(timer ITimer, mode RunLoopMode) {
	objc.CallMethod[objc.Void](r_, "addTimer:forMode:", timer, mode)
}

func (r_ RunLoop) AddPort_ForMode(aPort IPort, mode RunLoopMode) {
	objc.CallMethod[objc.Void](r_, "addPort:forMode:", aPort, mode)
}

func (r_ RunLoop) RemovePort_ForMode(aPort IPort, mode RunLoopMode) {
	objc.CallMethod[objc.Void](r_, "removePort:forMode:", aPort, mode)
}

// deprecated
func (r_ RunLoop) ConfigureAsServer() {
	objc.CallMethod[objc.Void](r_, "configureAsServer")
}

func (r_ RunLoop) Run() {
	objc.CallMethod[objc.Void](r_, "run")
}

func (r_ RunLoop) RunMode_BeforeDate(mode RunLoopMode, limitDate IDate) bool {
	rv := objc.CallMethod[bool](r_, "runMode:beforeDate:", mode, limitDate)
	return rv
}

func (r_ RunLoop) RunUntilDate(limitDate IDate) {
	objc.CallMethod[objc.Void](r_, "runUntilDate:", limitDate)
}

func (r_ RunLoop) AcceptInputForMode_BeforeDate(mode RunLoopMode, limitDate IDate) {
	objc.CallMethod[objc.Void](r_, "acceptInputForMode:beforeDate:", mode, limitDate)
}

func (r_ RunLoop) PerformBlock(block func()) {
	objc.CallMethod[objc.Void](r_, "performBlock:", block)
}

func (r_ RunLoop) PerformInModes_Block(modes []RunLoopMode, block func()) {
	objc.CallMethod[objc.Void](r_, "performInModes:block:", modes, block)
}

func (r_ RunLoop) PerformSelector_Target_Argument_Order_Modes(aSelector objc.Selector, target objc.IObject, arg objc.IObject, order uint, modes []RunLoopMode) {
	objc.CallMethod[objc.Void](r_, "performSelector:target:argument:order:modes:", aSelector, target, arg, order, modes)
}

func (r_ RunLoop) CancelPerformSelector_Target_Argument(aSelector objc.Selector, target objc.IObject, arg objc.IObject) {
	objc.CallMethod[objc.Void](r_, "cancelPerformSelector:target:argument:", aSelector, target, arg)
}

func (r_ RunLoop) CancelPerformSelectorsWithTarget(target objc.IObject) {
	objc.CallMethod[objc.Void](r_, "cancelPerformSelectorsWithTarget:", target)
}

func (rc _RunLoopClass) CurrentRunLoop() RunLoop {
	rv := objc.CallMethod[RunLoop](rc, "currentRunLoop")
	return rv
}

func (r_ RunLoop) CurrentMode() RunLoopMode {
	rv := objc.CallMethod[RunLoopMode](r_, "currentMode")
	return rv
}

func (rc _RunLoopClass) MainRunLoop() RunLoop {
	rv := objc.CallMethod[RunLoop](rc, "mainRunLoop")
	return rv
}
