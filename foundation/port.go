// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[Port](pc, "alloc")
	return rv
}

func (pc _PortClass) New() Port {
	rv := ffi.CallMethod[Port](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPort() Port {
	return PortClass.New()
}

func (p_ Port) Init() Port {
	rv := ffi.CallMethod[Port](p_, "init")
	return rv
}

func (pc _PortClass) Port() Port {
	rv := ffi.CallMethod[Port](pc, "port")
	return rv
}

func (p_ Port) Invalidate() {
	ffi.CallMethod[ffi.Void](p_, "invalidate")
}

func (p_ Port) SetDelegate(anObject PortDelegate) {
	po := ffi.CreateProtocol("NSPortDelegate", anObject)
	defer po.Release()
	ffi.CallMethod[ffi.Void](p_, "setDelegate:", po)
}

func (p_ Port) SetDelegate0(anObject objc.IObject) {
	ffi.CallMethod[ffi.Void](p_, "setDelegate:", anObject)
}

func (p_ Port) Delegate() PortDelegateWrapper {
	rv := ffi.CallMethod[PortDelegateWrapper](p_, "delegate")
	return rv
}

func (p_ Port) RemoveFromRunLoop_ForMode(runLoop IRunLoop, mode RunLoopMode) {
	ffi.CallMethod[ffi.Void](p_, "removeFromRunLoop:forMode:", runLoop, mode)
}

func (p_ Port) ScheduleInRunLoop_ForMode(runLoop IRunLoop, mode RunLoopMode) {
	ffi.CallMethod[ffi.Void](p_, "scheduleInRunLoop:forMode:", runLoop, mode)
}

func (p_ Port) IsValid() bool {
	rv := ffi.CallMethod[bool](p_, "isValid")
	return rv
}

func (p_ Port) ReservedSpaceLength() uint {
	rv := ffi.CallMethod[uint](p_, "reservedSpaceLength")
	return rv
}
