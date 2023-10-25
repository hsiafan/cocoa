// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var DateComponentsClass = _DateComponentsClass{objc.GetClass("NSDateComponents")}

type _DateComponentsClass struct {
	objc.Class
}

type IDateComponents interface {
	objc.IObject
	IsValidDateInCalendar(calendar ICalendar) bool
	// deprecated
	Week() int
	// deprecated
	SetWeek(v int)
	ValueForComponent(unit CalendarUnit) int
	SetValue_ForComponent(value int, unit CalendarUnit)
	Calendar() Calendar
	SetCalendar(value ICalendar)
	TimeZone() TimeZone
	SetTimeZone(value ITimeZone)
	IsValidDate() bool
	Date() Date
	Era() int
	SetEra(value int)
	Year() int
	SetYear(value int)
	YearForWeekOfYear() int
	SetYearForWeekOfYear(value int)
	Quarter() int
	SetQuarter(value int)
	Month() int
	SetMonth(value int)
	IsLeapMonth() bool
	SetLeapMonth(value bool)
	Weekday() int
	SetWeekday(value int)
	WeekdayOrdinal() int
	SetWeekdayOrdinal(value int)
	WeekOfMonth() int
	SetWeekOfMonth(value int)
	WeekOfYear() int
	SetWeekOfYear(value int)
	Day() int
	SetDay(value int)
	Hour() int
	SetHour(value int)
	Minute() int
	SetMinute(value int)
	Second() int
	SetSecond(value int)
	Nanosecond() int
	SetNanosecond(value int)
}

type DateComponents struct {
	objc.Object
}

func MakeDateComponents(ptr unsafe.Pointer) DateComponents {
	return DateComponents{
		Object: objc.MakeObject(ptr),
	}
}

func (dc _DateComponentsClass) Alloc() DateComponents {
	rv := objc.CallMethod[DateComponents](dc, objc.SEL("alloc"))
	return rv
}

func (dc _DateComponentsClass) New() DateComponents {
	rv := objc.CallMethod[DateComponents](dc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewDateComponents() DateComponents {
	return DateComponentsClass.New()
}

func (d_ DateComponents) Init() DateComponents {
	rv := objc.CallMethod[DateComponents](d_, objc.SEL("init"))
	return rv
}

func (d_ DateComponents) IsValidDateInCalendar(calendar ICalendar) bool {
	rv := objc.CallMethod[bool](d_, objc.SEL("isValidDateInCalendar:"), objc.ExtractPtr(calendar))
	return rv
}

// deprecated
func (d_ DateComponents) Week() int {
	rv := objc.CallMethod[int](d_, objc.SEL("week"))
	return rv
}

// deprecated
func (d_ DateComponents) SetWeek(v int) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setWeek:"), v)
}

func (d_ DateComponents) ValueForComponent(unit CalendarUnit) int {
	rv := objc.CallMethod[int](d_, objc.SEL("valueForComponent:"), unit)
	return rv
}

func (d_ DateComponents) SetValue_ForComponent(value int, unit CalendarUnit) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setValue:forComponent:"), value, unit)
}

func (d_ DateComponents) Calendar() Calendar {
	rv := objc.CallMethod[Calendar](d_, objc.SEL("calendar"))
	return rv
}

func (d_ DateComponents) SetCalendar(value ICalendar) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setCalendar:"), objc.ExtractPtr(value))
}

func (d_ DateComponents) TimeZone() TimeZone {
	rv := objc.CallMethod[TimeZone](d_, objc.SEL("timeZone"))
	return rv
}

func (d_ DateComponents) SetTimeZone(value ITimeZone) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setTimeZone:"), objc.ExtractPtr(value))
}

func (d_ DateComponents) IsValidDate() bool {
	rv := objc.CallMethod[bool](d_, objc.SEL("isValidDate"))
	return rv
}

func (d_ DateComponents) Date() Date {
	rv := objc.CallMethod[Date](d_, objc.SEL("date"))
	return rv
}

func (d_ DateComponents) Era() int {
	rv := objc.CallMethod[int](d_, objc.SEL("era"))
	return rv
}

func (d_ DateComponents) SetEra(value int) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setEra:"), value)
}

func (d_ DateComponents) Year() int {
	rv := objc.CallMethod[int](d_, objc.SEL("year"))
	return rv
}

func (d_ DateComponents) SetYear(value int) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setYear:"), value)
}

func (d_ DateComponents) YearForWeekOfYear() int {
	rv := objc.CallMethod[int](d_, objc.SEL("yearForWeekOfYear"))
	return rv
}

func (d_ DateComponents) SetYearForWeekOfYear(value int) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setYearForWeekOfYear:"), value)
}

func (d_ DateComponents) Quarter() int {
	rv := objc.CallMethod[int](d_, objc.SEL("quarter"))
	return rv
}

func (d_ DateComponents) SetQuarter(value int) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setQuarter:"), value)
}

func (d_ DateComponents) Month() int {
	rv := objc.CallMethod[int](d_, objc.SEL("month"))
	return rv
}

func (d_ DateComponents) SetMonth(value int) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setMonth:"), value)
}

func (d_ DateComponents) IsLeapMonth() bool {
	rv := objc.CallMethod[bool](d_, objc.SEL("isLeapMonth"))
	return rv
}

func (d_ DateComponents) SetLeapMonth(value bool) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setLeapMonth:"), value)
}

func (d_ DateComponents) Weekday() int {
	rv := objc.CallMethod[int](d_, objc.SEL("weekday"))
	return rv
}

func (d_ DateComponents) SetWeekday(value int) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setWeekday:"), value)
}

func (d_ DateComponents) WeekdayOrdinal() int {
	rv := objc.CallMethod[int](d_, objc.SEL("weekdayOrdinal"))
	return rv
}

func (d_ DateComponents) SetWeekdayOrdinal(value int) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setWeekdayOrdinal:"), value)
}

func (d_ DateComponents) WeekOfMonth() int {
	rv := objc.CallMethod[int](d_, objc.SEL("weekOfMonth"))
	return rv
}

func (d_ DateComponents) SetWeekOfMonth(value int) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setWeekOfMonth:"), value)
}

func (d_ DateComponents) WeekOfYear() int {
	rv := objc.CallMethod[int](d_, objc.SEL("weekOfYear"))
	return rv
}

func (d_ DateComponents) SetWeekOfYear(value int) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setWeekOfYear:"), value)
}

func (d_ DateComponents) Day() int {
	rv := objc.CallMethod[int](d_, objc.SEL("day"))
	return rv
}

func (d_ DateComponents) SetDay(value int) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setDay:"), value)
}

func (d_ DateComponents) Hour() int {
	rv := objc.CallMethod[int](d_, objc.SEL("hour"))
	return rv
}

func (d_ DateComponents) SetHour(value int) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setHour:"), value)
}

func (d_ DateComponents) Minute() int {
	rv := objc.CallMethod[int](d_, objc.SEL("minute"))
	return rv
}

func (d_ DateComponents) SetMinute(value int) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setMinute:"), value)
}

func (d_ DateComponents) Second() int {
	rv := objc.CallMethod[int](d_, objc.SEL("second"))
	return rv
}

func (d_ DateComponents) SetSecond(value int) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setSecond:"), value)
}

func (d_ DateComponents) Nanosecond() int {
	rv := objc.CallMethod[int](d_, objc.SEL("nanosecond"))
	return rv
}

func (d_ DateComponents) SetNanosecond(value int) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setNanosecond:"), value)
}
