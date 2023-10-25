// auto-generated code, do not modify
package foundation

import (
	"unsafe"

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
	rv := objc.CallMethod[DateInterval](d_, objc.SEL("init"))
	return rv
}

func (d_ DateInterval) InitWithStartDate_Duration(startDate IDate, duration TimeInterval) DateInterval {
	rv := objc.CallMethod[DateInterval](d_, objc.SEL("initWithStartDate:duration:"), objc.ExtractPtr(startDate), duration)
	return rv
}

func (d_ DateInterval) InitWithStartDate_EndDate(startDate IDate, endDate IDate) DateInterval {
	rv := objc.CallMethod[DateInterval](d_, objc.SEL("initWithStartDate:endDate:"), objc.ExtractPtr(startDate), objc.ExtractPtr(endDate))
	return rv
}

func (dc _DateIntervalClass) Alloc() DateInterval {
	rv := objc.CallMethod[DateInterval](dc, objc.SEL("alloc"))
	return rv
}

func (dc _DateIntervalClass) New() DateInterval {
	rv := objc.CallMethod[DateInterval](dc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewDateInterval() DateInterval {
	return DateIntervalClass.New()
}

func (d_ DateInterval) Compare(dateInterval IDateInterval) ComparisonResult {
	rv := objc.CallMethod[ComparisonResult](d_, objc.SEL("compare:"), objc.ExtractPtr(dateInterval))
	return rv
}

func (d_ DateInterval) IsEqualToDateInterval(dateInterval IDateInterval) bool {
	rv := objc.CallMethod[bool](d_, objc.SEL("isEqualToDateInterval:"), objc.ExtractPtr(dateInterval))
	return rv
}

func (d_ DateInterval) IntersectsDateInterval(dateInterval IDateInterval) bool {
	rv := objc.CallMethod[bool](d_, objc.SEL("intersectsDateInterval:"), objc.ExtractPtr(dateInterval))
	return rv
}

func (d_ DateInterval) IntersectionWithDateInterval(dateInterval IDateInterval) DateInterval {
	rv := objc.CallMethod[DateInterval](d_, objc.SEL("intersectionWithDateInterval:"), objc.ExtractPtr(dateInterval))
	return rv
}

func (d_ DateInterval) ContainsDate(date IDate) bool {
	rv := objc.CallMethod[bool](d_, objc.SEL("containsDate:"), objc.ExtractPtr(date))
	return rv
}

func (d_ DateInterval) StartDate() Date {
	rv := objc.CallMethod[Date](d_, objc.SEL("startDate"))
	return rv
}

func (d_ DateInterval) EndDate() Date {
	rv := objc.CallMethod[Date](d_, objc.SEL("endDate"))
	return rv
}

func (d_ DateInterval) Duration() TimeInterval {
	rv := objc.CallMethod[TimeInterval](d_, objc.SEL("duration"))
	return rv
}
