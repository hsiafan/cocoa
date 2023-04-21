// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var DateComponentsFormatterClass = _DateComponentsFormatterClass{objc.GetClass("NSDateComponentsFormatter")}

type _DateComponentsFormatterClass struct {
	objc.Class
}

type IDateComponentsFormatter interface {
	IFormatter
	StringFromDateComponents(components IDateComponents) string
	StringFromDate_ToDate(startDate IDate, endDate IDate) string
	StringFromTimeInterval(ti TimeInterval) string
	AllowedUnits() CalendarUnit
	SetAllowedUnits(value CalendarUnit)
	AllowsFractionalUnits() bool
	SetAllowsFractionalUnits(value bool)
	Calendar() Calendar
	SetCalendar(value ICalendar)
	CollapsesLargestUnit() bool
	SetCollapsesLargestUnit(value bool)
	IncludesApproximationPhrase() bool
	SetIncludesApproximationPhrase(value bool)
	IncludesTimeRemainingPhrase() bool
	SetIncludesTimeRemainingPhrase(value bool)
	MaximumUnitCount() int
	SetMaximumUnitCount(value int)
	UnitsStyle() DateComponentsFormatterUnitsStyle
	SetUnitsStyle(value DateComponentsFormatterUnitsStyle)
	ZeroFormattingBehavior() DateComponentsFormatterZeroFormattingBehavior
	SetZeroFormattingBehavior(value DateComponentsFormatterZeroFormattingBehavior)
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	ReferenceDate() Date
	SetReferenceDate(value IDate)
}

type DateComponentsFormatter struct {
	Formatter
}

func MakeDateComponentsFormatter(ptr unsafe.Pointer) DateComponentsFormatter {
	return DateComponentsFormatter{
		Formatter: MakeFormatter(ptr),
	}
}

func (dc _DateComponentsFormatterClass) Alloc() DateComponentsFormatter {
	rv := objc.CallMethod[DateComponentsFormatter](dc, "alloc")
	return rv
}

func (dc _DateComponentsFormatterClass) New() DateComponentsFormatter {
	rv := objc.CallMethod[DateComponentsFormatter](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDateComponentsFormatter() DateComponentsFormatter {
	return DateComponentsFormatterClass.New()
}

func (d_ DateComponentsFormatter) Init() DateComponentsFormatter {
	rv := objc.CallMethod[DateComponentsFormatter](d_, "init")
	return rv
}

func (d_ DateComponentsFormatter) StringFromDateComponents(components IDateComponents) string {
	rv := objc.CallMethod[string](d_, "stringFromDateComponents:", components)
	return rv
}

func (d_ DateComponentsFormatter) StringFromDate_ToDate(startDate IDate, endDate IDate) string {
	rv := objc.CallMethod[string](d_, "stringFromDate:toDate:", startDate, endDate)
	return rv
}

func (d_ DateComponentsFormatter) StringFromTimeInterval(ti TimeInterval) string {
	rv := objc.CallMethod[string](d_, "stringFromTimeInterval:", ti)
	return rv
}

func (dc _DateComponentsFormatterClass) LocalizedStringFromDateComponents_UnitsStyle(components IDateComponents, unitsStyle DateComponentsFormatterUnitsStyle) string {
	rv := objc.CallMethod[string](dc, "localizedStringFromDateComponents:unitsStyle:", components, unitsStyle)
	return rv
}

func (d_ DateComponentsFormatter) AllowedUnits() CalendarUnit {
	rv := objc.CallMethod[CalendarUnit](d_, "allowedUnits")
	return rv
}

func (d_ DateComponentsFormatter) SetAllowedUnits(value CalendarUnit) {
	objc.CallMethod[objc.Void](d_, "setAllowedUnits:", value)
}

func (d_ DateComponentsFormatter) AllowsFractionalUnits() bool {
	rv := objc.CallMethod[bool](d_, "allowsFractionalUnits")
	return rv
}

func (d_ DateComponentsFormatter) SetAllowsFractionalUnits(value bool) {
	objc.CallMethod[objc.Void](d_, "setAllowsFractionalUnits:", value)
}

func (d_ DateComponentsFormatter) Calendar() Calendar {
	rv := objc.CallMethod[Calendar](d_, "calendar")
	return rv
}

func (d_ DateComponentsFormatter) SetCalendar(value ICalendar) {
	objc.CallMethod[objc.Void](d_, "setCalendar:", value)
}

func (d_ DateComponentsFormatter) CollapsesLargestUnit() bool {
	rv := objc.CallMethod[bool](d_, "collapsesLargestUnit")
	return rv
}

func (d_ DateComponentsFormatter) SetCollapsesLargestUnit(value bool) {
	objc.CallMethod[objc.Void](d_, "setCollapsesLargestUnit:", value)
}

func (d_ DateComponentsFormatter) IncludesApproximationPhrase() bool {
	rv := objc.CallMethod[bool](d_, "includesApproximationPhrase")
	return rv
}

func (d_ DateComponentsFormatter) SetIncludesApproximationPhrase(value bool) {
	objc.CallMethod[objc.Void](d_, "setIncludesApproximationPhrase:", value)
}

func (d_ DateComponentsFormatter) IncludesTimeRemainingPhrase() bool {
	rv := objc.CallMethod[bool](d_, "includesTimeRemainingPhrase")
	return rv
}

func (d_ DateComponentsFormatter) SetIncludesTimeRemainingPhrase(value bool) {
	objc.CallMethod[objc.Void](d_, "setIncludesTimeRemainingPhrase:", value)
}

func (d_ DateComponentsFormatter) MaximumUnitCount() int {
	rv := objc.CallMethod[int](d_, "maximumUnitCount")
	return rv
}

func (d_ DateComponentsFormatter) SetMaximumUnitCount(value int) {
	objc.CallMethod[objc.Void](d_, "setMaximumUnitCount:", value)
}

func (d_ DateComponentsFormatter) UnitsStyle() DateComponentsFormatterUnitsStyle {
	rv := objc.CallMethod[DateComponentsFormatterUnitsStyle](d_, "unitsStyle")
	return rv
}

func (d_ DateComponentsFormatter) SetUnitsStyle(value DateComponentsFormatterUnitsStyle) {
	objc.CallMethod[objc.Void](d_, "setUnitsStyle:", value)
}

func (d_ DateComponentsFormatter) ZeroFormattingBehavior() DateComponentsFormatterZeroFormattingBehavior {
	rv := objc.CallMethod[DateComponentsFormatterZeroFormattingBehavior](d_, "zeroFormattingBehavior")
	return rv
}

func (d_ DateComponentsFormatter) SetZeroFormattingBehavior(value DateComponentsFormatterZeroFormattingBehavior) {
	objc.CallMethod[objc.Void](d_, "setZeroFormattingBehavior:", value)
}

func (d_ DateComponentsFormatter) FormattingContext() FormattingContext {
	rv := objc.CallMethod[FormattingContext](d_, "formattingContext")
	return rv
}

func (d_ DateComponentsFormatter) SetFormattingContext(value FormattingContext) {
	objc.CallMethod[objc.Void](d_, "setFormattingContext:", value)
}

func (d_ DateComponentsFormatter) ReferenceDate() Date {
	rv := objc.CallMethod[Date](d_, "referenceDate")
	return rv
}

func (d_ DateComponentsFormatter) SetReferenceDate(value IDate) {
	objc.CallMethod[objc.Void](d_, "setReferenceDate:", value)
}
