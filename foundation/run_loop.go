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
	rv := objc.CallMethod[RunLoop](rc, objc.SEL("alloc"))
	return rv
}

func (rc _RunLoopClass) New() RunLoop {
	rv := objc.CallMethod[RunLoop](rc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewRunLoop() RunLoop {
	return RunLoopClass.New()
}

func (r_ RunLoop) Init() RunLoop {
	rv := objc.CallMethod[RunLoop](r_, objc.SEL("init"))
	return rv
}

func (r_ RunLoop) LimitDateForMode(mode RunLoopMode) Date {
	rv := objc.CallMethod[Date](r_, objc.SEL("limitDateForMode:"), mode)
	return rv
}

func (r_ RunLoop) AddTimer_ForMode(timer ITimer, mode RunLoopMode) {
	objc.CallMethod[objc.Void](r_, objc.SEL("addTimer:forMode:"), objc.ExtractPtr(timer), mode)
}

func (r_ RunLoop) AddPort_ForMode(aPort IPort, mode RunLoopMode) {
	objc.CallMethod[objc.Void](r_, objc.SEL("addPort:forMode:"), objc.ExtractPtr(aPort), mode)
}

func (r_ RunLoop) RemovePort_ForMode(aPort IPort, mode RunLoopMode) {
	objc.CallMethod[objc.Void](r_, objc.SEL("removePort:forMode:"), objc.ExtractPtr(aPort), mode)
}

// deprecated
func (r_ RunLoop) ConfigureAsServer() {
	objc.CallMethod[objc.Void](r_, objc.SEL("configureAsServer"))
}

func (r_ RunLoop) Run() {
	objc.CallMethod[objc.Void](r_, objc.SEL("run"))
}

func (r_ RunLoop) RunMode_BeforeDate(mode RunLoopMode, limitDate IDate) bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("runMode:beforeDate:"), mode, objc.ExtractPtr(limitDate))
	return rv
}

func (r_ RunLoop) RunUntilDate(limitDate IDate) {
	objc.CallMethod[objc.Void](r_, objc.SEL("runUntilDate:"), objc.ExtractPtr(limitDate))
}

func (r_ RunLoop) AcceptInputForMode_BeforeDate(mode RunLoopMode, limitDate IDate) {
	objc.CallMethod[objc.Void](r_, objc.SEL("acceptInputForMode:beforeDate:"), mode, objc.ExtractPtr(limitDate))
}

func (r_ RunLoop) PerformBlock(block func()) {
	objc.CallMethod[objc.Void](r_, objc.SEL("performBlock:"), block)
}

func (r_ RunLoop) PerformInModes_Block(modes []RunLoopMode, block func()) {
	objc.CallMethod[objc.Void](r_, objc.SEL("performInModes:block:"), modes, block)
}

func (r_ RunLoop) PerformSelector_Target_Argument_Order_Modes(aSelector objc.Selector, target objc.IObject, arg objc.IObject, order uint, modes []RunLoopMode) {
	objc.CallMethod[objc.Void](r_, objc.SEL("performSelector:target:argument:order:modes:"), aSelector, objc.ExtractPtr(target), objc.ExtractPtr(arg), order, modes)
}

func (r_ RunLoop) CancelPerformSelector_Target_Argument(aSelector objc.Selector, target objc.IObject, arg objc.IObject) {
	objc.CallMethod[objc.Void](r_, objc.SEL("cancelPerformSelector:target:argument:"), aSelector, objc.ExtractPtr(target), objc.ExtractPtr(arg))
}

func (r_ RunLoop) CancelPerformSelectorsWithTarget(target objc.IObject) {
	objc.CallMethod[objc.Void](r_, objc.SEL("cancelPerformSelectorsWithTarget:"), objc.ExtractPtr(target))
}

func (rc _RunLoopClass) CurrentRunLoop() RunLoop {
	rv := objc.CallMethod[RunLoop](rc, objc.SEL("currentRunLoop"))
	return rv
}

func (r_ RunLoop) CurrentMode() RunLoopMode {
	rv := objc.CallMethod[RunLoopMode](r_, objc.SEL("currentMode"))
	return rv
}

func (rc _RunLoopClass) MainRunLoop() RunLoop {
	rv := objc.CallMethod[RunLoop](rc, objc.SEL("mainRunLoop"))
	return rv
}
