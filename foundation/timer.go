// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var TimerClass = _TimerClass{objc.GetClass("NSTimer")}

type _TimerClass struct {
	objc.Class
}

type ITimer interface {
	objc.IObject
	Fire()
	Invalidate()
	IsValid() bool
	FireDate() Date
	SetFireDate(value IDate)
	TimeInterval() TimeInterval
	UserInfo() objc.Object
	Tolerance() TimeInterval
	SetTolerance(value TimeInterval)
}

type Timer struct {
	objc.Object
}

func MakeTimer(ptr unsafe.Pointer) Timer {
	return Timer{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ Timer) InitWithFireDate_Interval_Repeats_Block(date IDate, interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	rv := objc.CallMethod[Timer](t_, "initWithFireDate:interval:repeats:block:", date, interval, repeats, block)
	return rv
}

func (t_ Timer) InitWithFireDate_Interval_Target_Selector_UserInfo_Repeats(date IDate, ti TimeInterval, t objc.IObject, s objc.Selector, ui objc.IObject, rep bool) Timer {
	rv := objc.CallMethod[Timer](t_, "initWithFireDate:interval:target:selector:userInfo:repeats:", date, ti, t, s, ui, rep)
	return rv
}

func (tc _TimerClass) Alloc() Timer {
	rv := objc.CallMethod[Timer](tc, "alloc")
	return rv
}

func (tc _TimerClass) New() Timer {
	rv := objc.CallMethod[Timer](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTimer() Timer {
	return TimerClass.New()
}

func (t_ Timer) Init() Timer {
	rv := objc.CallMethod[Timer](t_, "init")
	return rv
}

func (tc _TimerClass) ScheduledTimerWithTimeInterval_Repeats_Block(interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	rv := objc.CallMethod[Timer](tc, "scheduledTimerWithTimeInterval:repeats:block:", interval, repeats, block)
	return rv
}

func (tc _TimerClass) ScheduledTimerWithTimeInterval_Target_Selector_UserInfo_Repeats(ti TimeInterval, aTarget objc.IObject, aSelector objc.Selector, userInfo objc.IObject, yesOrNo bool) Timer {
	rv := objc.CallMethod[Timer](tc, "scheduledTimerWithTimeInterval:target:selector:userInfo:repeats:", ti, aTarget, aSelector, userInfo, yesOrNo)
	return rv
}

func (tc _TimerClass) TimerWithTimeInterval_Repeats_Block(interval TimeInterval, repeats bool, block func(timer Timer)) Timer {
	rv := objc.CallMethod[Timer](tc, "timerWithTimeInterval:repeats:block:", interval, repeats, block)
	return rv
}

func (tc _TimerClass) TimerWithTimeInterval_Target_Selector_UserInfo_Repeats(ti TimeInterval, aTarget objc.IObject, aSelector objc.Selector, userInfo objc.IObject, yesOrNo bool) Timer {
	rv := objc.CallMethod[Timer](tc, "timerWithTimeInterval:target:selector:userInfo:repeats:", ti, aTarget, aSelector, userInfo, yesOrNo)
	return rv
}

func (t_ Timer) Fire() {
	objc.CallMethod[objc.Void](t_, "fire")
}

func (t_ Timer) Invalidate() {
	objc.CallMethod[objc.Void](t_, "invalidate")
}

func (t_ Timer) IsValid() bool {
	rv := objc.CallMethod[bool](t_, "isValid")
	return rv
}

func (t_ Timer) FireDate() Date {
	rv := objc.CallMethod[Date](t_, "fireDate")
	return rv
}

func (t_ Timer) SetFireDate(value IDate) {
	objc.CallMethod[objc.Void](t_, "setFireDate:", value)
}

func (t_ Timer) TimeInterval() TimeInterval {
	rv := objc.CallMethod[TimeInterval](t_, "timeInterval")
	return rv
}

func (t_ Timer) UserInfo() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, "userInfo")
	return rv
}

func (t_ Timer) Tolerance() TimeInterval {
	rv := objc.CallMethod[TimeInterval](t_, "tolerance")
	return rv
}

func (t_ Timer) SetTolerance(value TimeInterval) {
	objc.CallMethod[objc.Void](t_, "setTolerance:", value)
}
