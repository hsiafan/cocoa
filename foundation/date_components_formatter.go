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
	rv := objc.CallMethod[DateComponentsFormatter](dc, objc.SEL("alloc"))
	return rv
}

func (dc _DateComponentsFormatterClass) New() DateComponentsFormatter {
	rv := objc.CallMethod[DateComponentsFormatter](dc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewDateComponentsFormatter() DateComponentsFormatter {
	return DateComponentsFormatterClass.New()
}

func (d_ DateComponentsFormatter) Init() DateComponentsFormatter {
	rv := objc.CallMethod[DateComponentsFormatter](d_, objc.SEL("init"))
	return rv
}

func (d_ DateComponentsFormatter) StringFromDateComponents(components IDateComponents) string {
	rv := objc.CallMethod[string](d_, objc.SEL("stringFromDateComponents:"), objc.ExtractPtr(components))
	return rv
}

func (d_ DateComponentsFormatter) StringFromDate_ToDate(startDate IDate, endDate IDate) string {
	rv := objc.CallMethod[string](d_, objc.SEL("stringFromDate:toDate:"), objc.ExtractPtr(startDate), objc.ExtractPtr(endDate))
	return rv
}

func (d_ DateComponentsFormatter) StringFromTimeInterval(ti TimeInterval) string {
	rv := objc.CallMethod[string](d_, objc.SEL("stringFromTimeInterval:"), ti)
	return rv
}

func (dc _DateComponentsFormatterClass) LocalizedStringFromDateComponents_UnitsStyle(components IDateComponents, unitsStyle DateComponentsFormatterUnitsStyle) string {
	rv := objc.CallMethod[string](dc, objc.SEL("localizedStringFromDateComponents:unitsStyle:"), objc.ExtractPtr(components), unitsStyle)
	return rv
}

func (d_ DateComponentsFormatter) AllowedUnits() CalendarUnit {
	rv := objc.CallMethod[CalendarUnit](d_, objc.SEL("allowedUnits"))
	return rv
}

func (d_ DateComponentsFormatter) SetAllowedUnits(value CalendarUnit) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setAllowedUnits:"), value)
}

func (d_ DateComponentsFormatter) AllowsFractionalUnits() bool {
	rv := objc.CallMethod[bool](d_, objc.SEL("allowsFractionalUnits"))
	return rv
}

func (d_ DateComponentsFormatter) SetAllowsFractionalUnits(value bool) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setAllowsFractionalUnits:"), value)
}

func (d_ DateComponentsFormatter) Calendar() Calendar {
	rv := objc.CallMethod[Calendar](d_, objc.SEL("calendar"))
	return rv
}

func (d_ DateComponentsFormatter) SetCalendar(value ICalendar) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setCalendar:"), objc.ExtractPtr(value))
}

func (d_ DateComponentsFormatter) CollapsesLargestUnit() bool {
	rv := objc.CallMethod[bool](d_, objc.SEL("collapsesLargestUnit"))
	return rv
}

func (d_ DateComponentsFormatter) SetCollapsesLargestUnit(value bool) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setCollapsesLargestUnit:"), value)
}

func (d_ DateComponentsFormatter) IncludesApproximationPhrase() bool {
	rv := objc.CallMethod[bool](d_, objc.SEL("includesApproximationPhrase"))
	return rv
}

func (d_ DateComponentsFormatter) SetIncludesApproximationPhrase(value bool) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setIncludesApproximationPhrase:"), value)
}

func (d_ DateComponentsFormatter) IncludesTimeRemainingPhrase() bool {
	rv := objc.CallMethod[bool](d_, objc.SEL("includesTimeRemainingPhrase"))
	return rv
}

func (d_ DateComponentsFormatter) SetIncludesTimeRemainingPhrase(value bool) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setIncludesTimeRemainingPhrase:"), value)
}

func (d_ DateComponentsFormatter) MaximumUnitCount() int {
	rv := objc.CallMethod[int](d_, objc.SEL("maximumUnitCount"))
	return rv
}

func (d_ DateComponentsFormatter) SetMaximumUnitCount(value int) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setMaximumUnitCount:"), value)
}

func (d_ DateComponentsFormatter) UnitsStyle() DateComponentsFormatterUnitsStyle {
	rv := objc.CallMethod[DateComponentsFormatterUnitsStyle](d_, objc.SEL("unitsStyle"))
	return rv
}

func (d_ DateComponentsFormatter) SetUnitsStyle(value DateComponentsFormatterUnitsStyle) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setUnitsStyle:"), value)
}

func (d_ DateComponentsFormatter) ZeroFormattingBehavior() DateComponentsFormatterZeroFormattingBehavior {
	rv := objc.CallMethod[DateComponentsFormatterZeroFormattingBehavior](d_, objc.SEL("zeroFormattingBehavior"))
	return rv
}

func (d_ DateComponentsFormatter) SetZeroFormattingBehavior(value DateComponentsFormatterZeroFormattingBehavior) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setZeroFormattingBehavior:"), value)
}

func (d_ DateComponentsFormatter) FormattingContext() FormattingContext {
	rv := objc.CallMethod[FormattingContext](d_, objc.SEL("formattingContext"))
	return rv
}

func (d_ DateComponentsFormatter) SetFormattingContext(value FormattingContext) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setFormattingContext:"), value)
}

func (d_ DateComponentsFormatter) ReferenceDate() Date {
	rv := objc.CallMethod[Date](d_, objc.SEL("referenceDate"))
	return rv
}

func (d_ DateComponentsFormatter) SetReferenceDate(value IDate) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setReferenceDate:"), objc.ExtractPtr(value))
}
