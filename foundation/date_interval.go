// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var DateIntervalClass = _DateIntervalClass{objc.GetClass("NSDateInterval")}

type _DateIntervalClass struct {
	objc.Class
}

type IDateInterval interface {
	objc.IObject
	Compare(dateInterval IDateInterval) ComparisonResult
	IsEqualToDateInterval(dateInterval IDateInterval) bool
	IntersectsDateInterval(dateInterval IDateInterval) bool
	IntersectionWithDateInterval(dateInterval IDateInterval) DateInterval
	ContainsDate(date IDate) bool
	StartDate() Date
	EndDate() Date
	Duration() TimeInterval
}

type DateInterval struct {
	objc.Object
}

func MakeDateInterval(ptr unsafe.Pointer) DateInterval {
	return DateInterval{
		Object: objc.MakeObject(ptr),
	}
}

func (d_ DateInterval) Init() DateInterval {
	rv := ffi.CallMethod[DateInterval](d_, "init")
	return rv
}

func (d_ DateInterval) InitWithStartDate_Duration(startDate IDate, duration TimeInterval) DateInterval {
	rv := ffi.CallMethod[DateInterval](d_, "initWithStartDate:duration:", startDate, duration)
	return rv
}

func (d_ DateInterval) InitWithStartDate_EndDate(startDate IDate, endDate IDate) DateInterval {
	rv := ffi.CallMethod[DateInterval](d_, "initWithStartDate:endDate:", startDate, endDate)
	return rv
}

func (dc _DateIntervalClass) Alloc() DateInterval {
	rv := ffi.CallMethod[DateInterval](dc, "alloc")
	return rv
}

func (dc _DateIntervalClass) New() DateInterval {
	rv := ffi.CallMethod[DateInterval](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDateInterval() DateInterval {
	return DateIntervalClass.New()
}

func (d_ DateInterval) Compare(dateInterval IDateInterval) ComparisonResult {
	rv := ffi.CallMethod[ComparisonResult](d_, "compare:", dateInterval)
	return rv
}

func (d_ DateInterval) IsEqualToDateInterval(dateInterval IDateInterval) bool {
	rv := ffi.CallMethod[bool](d_, "isEqualToDateInterval:", dateInterval)
	return rv
}

func (d_ DateInterval) IntersectsDateInterval(dateInterval IDateInterval) bool {
	rv := ffi.CallMethod[bool](d_, "intersectsDateInterval:", dateInterval)
	return rv
}

func (d_ DateInterval) IntersectionWithDateInterval(dateInterval IDateInterval) DateInterval {
	rv := ffi.CallMethod[DateInterval](d_, "intersectionWithDateInterval:", dateInterval)
	return rv
}

func (d_ DateInterval) ContainsDate(date IDate) bool {
	rv := ffi.CallMethod[bool](d_, "containsDate:", date)
	return rv
}

func (d_ DateInterval) StartDate() Date {
	rv := ffi.CallMethod[Date](d_, "startDate")
	return rv
}

func (d_ DateInterval) EndDate() Date {
	rv := ffi.CallMethod[Date](d_, "endDate")
	return rv
}

func (d_ DateInterval) Duration() TimeInterval {
	rv := ffi.CallMethod[TimeInterval](d_, "duration")
	return rv
}
