// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var PortClass = _PortClass{objc.GetClass("NSPort")}

type _PortClass struct {
	objc.Class
}

type IPort interface {
	objc.IObject
	Invalidate()
	SetDelegate(anObject PortDelegate)
	SetDelegate0(anObject objc.IObject)
	Delegate() PortDelegateWrapper
	RemoveFromRunLoop_ForMode(runLoop IRunLoop, mode RunLoopMode)
	ScheduleInRunLoop_ForMode(runLoop IRunLoop, mode RunLoopMode)
	IsValid() bool
	ReservedSpaceLength() uint
}

type Port struct {
	objc.Object
}

func MakePort(ptr unsafe.Pointer) Port {
	return Port{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PortClass) Alloc() Port {
	rv := objc.CallMethod[Port](pc, "alloc")
	return rv
}

func (pc _PortClass) New() Port {
	rv := objc.CallMethod[Port](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPort() Port {
	return PortClass.New()
}

func (p_ Port) Init() Port {
	rv := objc.CallMethod[Port](p_, "init")
	return rv
}

func (pc _PortClass) Port() Port {
	rv := objc.CallMethod[Port](pc, "port")
	return rv
}

func (p_ Port) Invalidate() {
	objc.CallMethod[objc.Void](p_, "invalidate")
}

func (p_ Port) SetDelegate(anObject PortDelegate) {
	po := objc.CreateProtocol("NSPortDelegate", anObject)
	defer po.Release()
	objc.CallMethod[objc.Void](p_, "setDelegate:", po)
}

func (p_ Port) SetDelegate0(anObject objc.IObject) {
	objc.CallMethod[objc.Void](p_, "setDelegate:", anObject)
}

func (p_ Port) Delegate() PortDelegateWrapper {
	rv := objc.CallMethod[PortDelegateWrapper](p_, "delegate")
	return rv
}

func (p_ Port) RemoveFromRunLoop_ForMode(runLoop IRunLoop, mode RunLoopMode) {
	objc.CallMethod[objc.Void](p_, "removeFromRunLoop:forMode:", runLoop, mode)
}

func (p_ Port) ScheduleInRunLoop_ForMode(runLoop IRunLoop, mode RunLoopMode) {
	objc.CallMethod[objc.Void](p_, "scheduleInRunLoop:forMode:", runLoop, mode)
}

func (p_ Port) IsValid() bool {
	rv := objc.CallMethod[bool](p_, "isValid")
	return rv
}

func (p_ Port) ReservedSpaceLength() uint {
	rv := objc.CallMethod[uint](p_, "reservedSpaceLength")
	return rv
}
