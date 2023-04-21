// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var DateIntervalFormatterClass = _DateIntervalFormatterClass{objc.GetClass("NSDateIntervalFormatter")}

type _DateIntervalFormatterClass struct {
	objc.Class
}

type IDateIntervalFormatter interface {
	IFormatter
	StringFromDate_ToDate(fromDate IDate, toDate IDate) string
	StringFromDateInterval(dateInterval IDateInterval) string
	DateStyle() DateIntervalFormatterStyle
	SetDateStyle(value DateIntervalFormatterStyle)
	TimeStyle() DateIntervalFormatterStyle
	SetTimeStyle(value DateIntervalFormatterStyle)
	DateTemplate() string
	SetDateTemplate(value string)
	Calendar() Calendar
	SetCalendar(value ICalendar)
	Locale() Locale
	SetLocale(value ILocale)
	TimeZone() TimeZone
	SetTimeZone(value ITimeZone)
}

type DateIntervalFormatter struct {
	Formatter
}

func MakeDateIntervalFormatter(ptr unsafe.Pointer) DateIntervalFormatter {
	return DateIntervalFormatter{
		Formatter: MakeFormatter(ptr),
	}
}

func (dc _DateIntervalFormatterClass) Alloc() DateIntervalFormatter {
	rv := objc.CallMethod[DateIntervalFormatter](dc, "alloc")
	return rv
}

func (dc _DateIntervalFormatterClass) New() DateIntervalFormatter {
	rv := objc.CallMethod[DateIntervalFormatter](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDateIntervalFormatter() DateIntervalFormatter {
	return DateIntervalFormatterClass.New()
}

func (d_ DateIntervalFormatter) Init() DateIntervalFormatter {
	rv := objc.CallMethod[DateIntervalFormatter](d_, "init")
	return rv
}

func (d_ DateIntervalFormatter) StringFromDate_ToDate(fromDate IDate, toDate IDate) string {
	rv := objc.CallMethod[string](d_, "stringFromDate:toDate:", fromDate, toDate)
	return rv
}

func (d_ DateIntervalFormatter) StringFromDateInterval(dateInterval IDateInterval) string {
	rv := objc.CallMethod[string](d_, "stringFromDateInterval:", dateInterval)
	return rv
}

func (d_ DateIntervalFormatter) DateStyle() DateIntervalFormatterStyle {
	rv := objc.CallMethod[DateIntervalFormatterStyle](d_, "dateStyle")
	return rv
}

func (d_ DateIntervalFormatter) SetDateStyle(value DateIntervalFormatterStyle) {
	objc.CallMethod[objc.Void](d_, "setDateStyle:", value)
}

func (d_ DateIntervalFormatter) TimeStyle() DateIntervalFormatterStyle {
	rv := objc.CallMethod[DateIntervalFormatterStyle](d_, "timeStyle")
	return rv
}

func (d_ DateIntervalFormatter) SetTimeStyle(value DateIntervalFormatterStyle) {
	objc.CallMethod[objc.Void](d_, "setTimeStyle:", value)
}

func (d_ DateIntervalFormatter) DateTemplate() string {
	rv := objc.CallMethod[string](d_, "dateTemplate")
	return rv
}

func (d_ DateIntervalFormatter) SetDateTemplate(value string) {
	objc.CallMethod[objc.Void](d_, "setDateTemplate:", value)
}

func (d_ DateIntervalFormatter) Calendar() Calendar {
	rv := objc.CallMethod[Calendar](d_, "calendar")
	return rv
}

func (d_ DateIntervalFormatter) SetCalendar(value ICalendar) {
	objc.CallMethod[objc.Void](d_, "setCalendar:", value)
}

func (d_ DateIntervalFormatter) Locale() Locale {
	rv := objc.CallMethod[Locale](d_, "locale")
	return rv
}

func (d_ DateIntervalFormatter) SetLocale(value ILocale) {
	objc.CallMethod[objc.Void](d_, "setLocale:", value)
}

func (d_ DateIntervalFormatter) TimeZone() TimeZone {
	rv := objc.CallMethod[TimeZone](d_, "timeZone")
	return rv
}

func (d_ DateIntervalFormatter) SetTimeZone(value ITimeZone) {
	objc.CallMethod[objc.Void](d_, "setTimeZone:", value)
}
