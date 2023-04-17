// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[DateComponents](dc, "alloc")
	return rv
}

func (dc _DateComponentsClass) New() DateComponents {
	rv := ffi.CallMethod[DateComponents](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDateComponents() DateComponents {
	return DateComponentsClass.New()
}

func (d_ DateComponents) Init() DateComponents {
	rv := ffi.CallMethod[DateComponents](d_, "init")
	return rv
}

func (d_ DateComponents) IsValidDateInCalendar(calendar ICalendar) bool {
	rv := ffi.CallMethod[bool](d_, "isValidDateInCalendar:", calendar)
	return rv
}

// deprecated
func (d_ DateComponents) Week() int {
	rv := ffi.CallMethod[int](d_, "week")
	return rv
}

// deprecated
func (d_ DateComponents) SetWeek(v int) {
	ffi.CallMethod[ffi.Void](d_, "setWeek:", v)
}

func (d_ DateComponents) ValueForComponent(unit CalendarUnit) int {
	rv := ffi.CallMethod[int](d_, "valueForComponent:", unit)
	return rv
}

func (d_ DateComponents) SetValue_ForComponent(value int, unit CalendarUnit) {
	ffi.CallMethod[ffi.Void](d_, "setValue:forComponent:", value, unit)
}

func (d_ DateComponents) Calendar() Calendar {
	rv := ffi.CallMethod[Calendar](d_, "calendar")
	return rv
}

func (d_ DateComponents) SetCalendar(value ICalendar) {
	ffi.CallMethod[ffi.Void](d_, "setCalendar:", value)
}

func (d_ DateComponents) TimeZone() TimeZone {
	rv := ffi.CallMethod[TimeZone](d_, "timeZone")
	return rv
}

func (d_ DateComponents) SetTimeZone(value ITimeZone) {
	ffi.CallMethod[ffi.Void](d_, "setTimeZone:", value)
}

func (d_ DateComponents) IsValidDate() bool {
	rv := ffi.CallMethod[bool](d_, "isValidDate")
	return rv
}

func (d_ DateComponents) Date() Date {
	rv := ffi.CallMethod[Date](d_, "date")
	return rv
}

func (d_ DateComponents) Era() int {
	rv := ffi.CallMethod[int](d_, "era")
	return rv
}

func (d_ DateComponents) SetEra(value int) {
	ffi.CallMethod[ffi.Void](d_, "setEra:", value)
}

func (d_ DateComponents) Year() int {
	rv := ffi.CallMethod[int](d_, "year")
	return rv
}

func (d_ DateComponents) SetYear(value int) {
	ffi.CallMethod[ffi.Void](d_, "setYear:", value)
}

func (d_ DateComponents) YearForWeekOfYear() int {
	rv := ffi.CallMethod[int](d_, "yearForWeekOfYear")
	return rv
}

func (d_ DateComponents) SetYearForWeekOfYear(value int) {
	ffi.CallMethod[ffi.Void](d_, "setYearForWeekOfYear:", value)
}

func (d_ DateComponents) Quarter() int {
	rv := ffi.CallMethod[int](d_, "quarter")
	return rv
}

func (d_ DateComponents) SetQuarter(value int) {
	ffi.CallMethod[ffi.Void](d_, "setQuarter:", value)
}

func (d_ DateComponents) Month() int {
	rv := ffi.CallMethod[int](d_, "month")
	return rv
}

func (d_ DateComponents) SetMonth(value int) {
	ffi.CallMethod[ffi.Void](d_, "setMonth:", value)
}

func (d_ DateComponents) IsLeapMonth() bool {
	rv := ffi.CallMethod[bool](d_, "isLeapMonth")
	return rv
}

func (d_ DateComponents) SetLeapMonth(value bool) {
	ffi.CallMethod[ffi.Void](d_, "setLeapMonth:", value)
}

func (d_ DateComponents) Weekday() int {
	rv := ffi.CallMethod[int](d_, "weekday")
	return rv
}

func (d_ DateComponents) SetWeekday(value int) {
	ffi.CallMethod[ffi.Void](d_, "setWeekday:", value)
}

func (d_ DateComponents) WeekdayOrdinal() int {
	rv := ffi.CallMethod[int](d_, "weekdayOrdinal")
	return rv
}

func (d_ DateComponents) SetWeekdayOrdinal(value int) {
	ffi.CallMethod[ffi.Void](d_, "setWeekdayOrdinal:", value)
}

func (d_ DateComponents) WeekOfMonth() int {
	rv := ffi.CallMethod[int](d_, "weekOfMonth")
	return rv
}

func (d_ DateComponents) SetWeekOfMonth(value int) {
	ffi.CallMethod[ffi.Void](d_, "setWeekOfMonth:", value)
}

func (d_ DateComponents) WeekOfYear() int {
	rv := ffi.CallMethod[int](d_, "weekOfYear")
	return rv
}

func (d_ DateComponents) SetWeekOfYear(value int) {
	ffi.CallMethod[ffi.Void](d_, "setWeekOfYear:", value)
}

func (d_ DateComponents) Day() int {
	rv := ffi.CallMethod[int](d_, "day")
	return rv
}

func (d_ DateComponents) SetDay(value int) {
	ffi.CallMethod[ffi.Void](d_, "setDay:", value)
}

func (d_ DateComponents) Hour() int {
	rv := ffi.CallMethod[int](d_, "hour")
	return rv
}

func (d_ DateComponents) SetHour(value int) {
	ffi.CallMethod[ffi.Void](d_, "setHour:", value)
}

func (d_ DateComponents) Minute() int {
	rv := ffi.CallMethod[int](d_, "minute")
	return rv
}

func (d_ DateComponents) SetMinute(value int) {
	ffi.CallMethod[ffi.Void](d_, "setMinute:", value)
}

func (d_ DateComponents) Second() int {
	rv := ffi.CallMethod[int](d_, "second")
	return rv
}

func (d_ DateComponents) SetSecond(value int) {
	ffi.CallMethod[ffi.Void](d_, "setSecond:", value)
}

func (d_ DateComponents) Nanosecond() int {
	rv := ffi.CallMethod[int](d_, "nanosecond")
	return rv
}

func (d_ DateComponents) SetNanosecond(value int) {
	ffi.CallMethod[ffi.Void](d_, "setNanosecond:", value)
}
