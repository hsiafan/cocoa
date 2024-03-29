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
	SetDelegate(anObject objc.IObject)
	Delegate() objc.Object
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
	rv := objc.CallMethod[Port](pc, objc.SEL("alloc"))
	return rv
}

func (pc _PortClass) New() Port {
	rv := objc.CallMethod[Port](pc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewPort() Port {
	return PortClass.New()
}

func (p_ Port) Init() Port {
	rv := objc.CallMethod[Port](p_, objc.SEL("init"))
	return rv
}

func (pc _PortClass) Port() Port {
	rv := objc.CallMethod[Port](pc, objc.SEL("port"))
	return rv
}

func (p_ Port) Invalidate() {
	objc.CallMethod[objc.Void](p_, objc.SEL("invalidate"))
}

func (p_ Port) SetDelegate(anObject objc.IObject) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setDelegate:"), objc.ExtractPtr(anObject))
}

func (p_ Port) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](p_, objc.SEL("delegate"))
	return rv
}

func (p_ Port) RemoveFromRunLoop_ForMode(runLoop IRunLoop, mode RunLoopMode) {
	objc.CallMethod[objc.Void](p_, objc.SEL("removeFromRunLoop:forMode:"), objc.ExtractPtr(runLoop), mode)
}

func (p_ Port) ScheduleInRunLoop_ForMode(runLoop IRunLoop, mode RunLoopMode) {
	objc.CallMethod[objc.Void](p_, objc.SEL("scheduleInRunLoop:forMode:"), objc.ExtractPtr(runLoop), mode)
}

func (p_ Port) IsValid() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isValid"))
	return rv
}

func (p_ Port) ReservedSpaceLength() uint {
	rv := objc.CallMethod[uint](p_, objc.SEL("reservedSpaceLength"))
	return rv
}
