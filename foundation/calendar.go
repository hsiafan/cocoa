// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var CalendarClass = _CalendarClass{objc.GetClass("NSCalendar")}

type _CalendarClass struct {
	objc.Class
}

type ICalendar interface {
	objc.IObject
	InitWithCalendarIdentifier(ident CalendarIdentifier) objc.Object
	Date_MatchesComponents(date IDate, components IDateComponents) bool
	Component_FromDate(unit CalendarUnit, date IDate) int
	Components_FromDate(unitFlags CalendarUnit, date IDate) DateComponents
	Components_FromDate_ToDate_Options(unitFlags CalendarUnit, startingDate IDate, resultDate IDate, opts CalendarOptions) DateComponents
	Components_FromDateComponents_ToDateComponents_Options(unitFlags CalendarUnit, startingDateComp IDateComponents, resultDateComp IDateComponents, options CalendarOptions) DateComponents
	ComponentsInTimeZone_FromDate(timezone ITimeZone, date IDate) DateComponents
	MaximumRangeOfUnit(unit CalendarUnit) Range
	MinimumRangeOfUnit(unit CalendarUnit) Range
	OrdinalityOfUnit_InUnit_ForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) uint
	RangeOfUnit_InUnit_ForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) Range
	RangeOfUnit_StartDate_Interval_ForDate(unit CalendarUnit, datep *Date, tip *TimeInterval, date IDate) bool
	StartOfDayForDate(date IDate) Date
	EnumerateDatesStartingAfterDate_MatchingComponents_Options_UsingBlock(start IDate, comps IDateComponents, opts CalendarOptions, block func(date Date, exactMatch bool, stop *bool))
	NextDateAfterDate_MatchingComponents_Options(date IDate, comps IDateComponents, options CalendarOptions) Date
	NextDateAfterDate_MatchingHour_Minute_Second_Options(date IDate, hourValue int, minuteValue int, secondValue int, options CalendarOptions) Date
	NextDateAfterDate_MatchingUnit_Value_Options(date IDate, unit CalendarUnit, value int, options CalendarOptions) Date
	DateFromComponents(comps IDateComponents) Date
	DateByAddingComponents_ToDate_Options(comps IDateComponents, date IDate, opts CalendarOptions) Date
	DateByAddingUnit_Value_ToDate_Options(unit CalendarUnit, value int, date IDate, options CalendarOptions) Date
	DateBySettingHour_Minute_Second_OfDate_Options(h int, m int, s int, date IDate, opts CalendarOptions) Date
	DateBySettingUnit_Value_OfDate_Options(unit CalendarUnit, v int, date IDate, opts CalendarOptions) Date
	DateWithEra_Year_Month_Day_Hour_Minute_Second_Nanosecond(eraValue int, yearValue int, monthValue int, dayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date
	DateWithEra_YearForWeekOfYear_WeekOfYear_Weekday_Hour_Minute_Second_Nanosecond(eraValue int, yearValue int, weekValue int, weekdayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date
	CompareDate_ToDate_ToUnitGranularity(date1 IDate, date2 IDate, unit CalendarUnit) ComparisonResult
	IsDate_EqualToDate_ToUnitGranularity(date1 IDate, date2 IDate, unit CalendarUnit) bool
	IsDate_InSameDayAsDate(date1 IDate, date2 IDate) bool
	IsDateInToday(date IDate) bool
	IsDateInTomorrow(date IDate) bool
	IsDateInWeekend(date IDate) bool
	IsDateInYesterday(date IDate) bool
	CalendarIdentifier() CalendarIdentifier
	FirstWeekday() uint
	SetFirstWeekday(value uint)
	Locale() Locale
	SetLocale(value ILocale)
	TimeZone() TimeZone
	SetTimeZone(value ITimeZone)
	MinimumDaysInFirstWeek() uint
	SetMinimumDaysInFirstWeek(value uint)
	AMSymbol() string
	PMSymbol() string
	WeekdaySymbols() []string
	ShortWeekdaySymbols() []string
	VeryShortWeekdaySymbols() []string
	StandaloneWeekdaySymbols() []string
	ShortStandaloneWeekdaySymbols() []string
	VeryShortStandaloneWeekdaySymbols() []string
	MonthSymbols() []string
	ShortMonthSymbols() []string
	VeryShortMonthSymbols() []string
	StandaloneMonthSymbols() []string
	ShortStandaloneMonthSymbols() []string
	VeryShortStandaloneMonthSymbols() []string
	QuarterSymbols() []string
	ShortQuarterSymbols() []string
	StandaloneQuarterSymbols() []string
	ShortStandaloneQuarterSymbols() []string
	EraSymbols() []string
	LongEraSymbols() []string
}

type Calendar struct {
	objc.Object
}

func MakeCalendar(ptr unsafe.Pointer) Calendar {
	return Calendar{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CalendarClass) Alloc() Calendar {
	rv := objc.CallMethod[Calendar](cc, objc.GetSelector("alloc"))
	return rv
}

func (cc _CalendarClass) New() Calendar {
	rv := objc.CallMethod[Calendar](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCalendar() Calendar {
	return CalendarClass.New()
}

func (c_ Calendar) Init() Calendar {
	rv := objc.CallMethod[Calendar](c_, objc.GetSelector("init"))
	return rv
}

func (cc _CalendarClass) CalendarWithIdentifier(calendarIdentifierConstant CalendarIdentifier) Calendar {
	rv := objc.CallMethod[Calendar](cc, objc.GetSelector("calendarWithIdentifier:"), calendarIdentifierConstant)
	return rv
}

func (c_ Calendar) InitWithCalendarIdentifier(ident CalendarIdentifier) objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.GetSelector("initWithCalendarIdentifier:"), ident)
	return rv
}

func (c_ Calendar) Date_MatchesComponents(date IDate, components IDateComponents) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("date:matchesComponents:"), objc.ExtractPtr(date), objc.ExtractPtr(components))
	return rv
}

func (c_ Calendar) Component_FromDate(unit CalendarUnit, date IDate) int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("component:fromDate:"), unit, objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) Components_FromDate(unitFlags CalendarUnit, date IDate) DateComponents {
	rv := objc.CallMethod[DateComponents](c_, objc.GetSelector("components:fromDate:"), unitFlags, objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) Components_FromDate_ToDate_Options(unitFlags CalendarUnit, startingDate IDate, resultDate IDate, opts CalendarOptions) DateComponents {
	rv := objc.CallMethod[DateComponents](c_, objc.GetSelector("components:fromDate:toDate:options:"), unitFlags, objc.ExtractPtr(startingDate), objc.ExtractPtr(resultDate), opts)
	return rv
}

func (c_ Calendar) Components_FromDateComponents_ToDateComponents_Options(unitFlags CalendarUnit, startingDateComp IDateComponents, resultDateComp IDateComponents, options CalendarOptions) DateComponents {
	rv := objc.CallMethod[DateComponents](c_, objc.GetSelector("components:fromDateComponents:toDateComponents:options:"), unitFlags, objc.ExtractPtr(startingDateComp), objc.ExtractPtr(resultDateComp), options)
	return rv
}

func (c_ Calendar) ComponentsInTimeZone_FromDate(timezone ITimeZone, date IDate) DateComponents {
	rv := objc.CallMethod[DateComponents](c_, objc.GetSelector("componentsInTimeZone:fromDate:"), objc.ExtractPtr(timezone), objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) MaximumRangeOfUnit(unit CalendarUnit) Range {
	rv := objc.CallMethod[Range](c_, objc.GetSelector("maximumRangeOfUnit:"), unit)
	return rv
}

func (c_ Calendar) MinimumRangeOfUnit(unit CalendarUnit) Range {
	rv := objc.CallMethod[Range](c_, objc.GetSelector("minimumRangeOfUnit:"), unit)
	return rv
}

func (c_ Calendar) OrdinalityOfUnit_InUnit_ForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) uint {
	rv := objc.CallMethod[uint](c_, objc.GetSelector("ordinalityOfUnit:inUnit:forDate:"), smaller, larger, objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) RangeOfUnit_InUnit_ForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) Range {
	rv := objc.CallMethod[Range](c_, objc.GetSelector("rangeOfUnit:inUnit:forDate:"), smaller, larger, objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) RangeOfUnit_StartDate_Interval_ForDate(unit CalendarUnit, datep *Date, tip *TimeInterval, date IDate) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("rangeOfUnit:startDate:interval:forDate:"), unit, unsafe.Pointer(datep), tip, objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) StartOfDayForDate(date IDate) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("startOfDayForDate:"), objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) EnumerateDatesStartingAfterDate_MatchingComponents_Options_UsingBlock(start IDate, comps IDateComponents, opts CalendarOptions, block func(date Date, exactMatch bool, stop *bool)) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("enumerateDatesStartingAfterDate:matchingComponents:options:usingBlock:"), objc.ExtractPtr(start), objc.ExtractPtr(comps), opts, block)
}

func (c_ Calendar) NextDateAfterDate_MatchingComponents_Options(date IDate, comps IDateComponents, options CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("nextDateAfterDate:matchingComponents:options:"), objc.ExtractPtr(date), objc.ExtractPtr(comps), options)
	return rv
}

func (c_ Calendar) NextDateAfterDate_MatchingHour_Minute_Second_Options(date IDate, hourValue int, minuteValue int, secondValue int, options CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("nextDateAfterDate:matchingHour:minute:second:options:"), objc.ExtractPtr(date), hourValue, minuteValue, secondValue, options)
	return rv
}

func (c_ Calendar) NextDateAfterDate_MatchingUnit_Value_Options(date IDate, unit CalendarUnit, value int, options CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("nextDateAfterDate:matchingUnit:value:options:"), objc.ExtractPtr(date), unit, value, options)
	return rv
}

func (c_ Calendar) DateFromComponents(comps IDateComponents) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("dateFromComponents:"), objc.ExtractPtr(comps))
	return rv
}

func (c_ Calendar) DateByAddingComponents_ToDate_Options(comps IDateComponents, date IDate, opts CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("dateByAddingComponents:toDate:options:"), objc.ExtractPtr(comps), objc.ExtractPtr(date), opts)
	return rv
}

func (c_ Calendar) DateByAddingUnit_Value_ToDate_Options(unit CalendarUnit, value int, date IDate, options CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("dateByAddingUnit:value:toDate:options:"), unit, value, objc.ExtractPtr(date), options)
	return rv
}

func (c_ Calendar) DateBySettingHour_Minute_Second_OfDate_Options(h int, m int, s int, date IDate, opts CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("dateBySettingHour:minute:second:ofDate:options:"), h, m, s, objc.ExtractPtr(date), opts)
	return rv
}

func (c_ Calendar) DateBySettingUnit_Value_OfDate_Options(unit CalendarUnit, v int, date IDate, opts CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("dateBySettingUnit:value:ofDate:options:"), unit, v, objc.ExtractPtr(date), opts)
	return rv
}

func (c_ Calendar) DateWithEra_Year_Month_Day_Hour_Minute_Second_Nanosecond(eraValue int, yearValue int, monthValue int, dayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("dateWithEra:year:month:day:hour:minute:second:nanosecond:"), eraValue, yearValue, monthValue, dayValue, hourValue, minuteValue, secondValue, nanosecondValue)
	return rv
}

func (c_ Calendar) DateWithEra_YearForWeekOfYear_WeekOfYear_Weekday_Hour_Minute_Second_Nanosecond(eraValue int, yearValue int, weekValue int, weekdayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date {
	rv := objc.CallMethod[Date](c_, objc.GetSelector("dateWithEra:yearForWeekOfYear:weekOfYear:weekday:hour:minute:second:nanosecond:"), eraValue, yearValue, weekValue, weekdayValue, hourValue, minuteValue, secondValue, nanosecondValue)
	return rv
}

func (c_ Calendar) CompareDate_ToDate_ToUnitGranularity(date1 IDate, date2 IDate, unit CalendarUnit) ComparisonResult {
	rv := objc.CallMethod[ComparisonResult](c_, objc.GetSelector("compareDate:toDate:toUnitGranularity:"), objc.ExtractPtr(date1), objc.ExtractPtr(date2), unit)
	return rv
}

func (c_ Calendar) IsDate_EqualToDate_ToUnitGranularity(date1 IDate, date2 IDate, unit CalendarUnit) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isDate:equalToDate:toUnitGranularity:"), objc.ExtractPtr(date1), objc.ExtractPtr(date2), unit)
	return rv
}

func (c_ Calendar) IsDate_InSameDayAsDate(date1 IDate, date2 IDate) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isDate:inSameDayAsDate:"), objc.ExtractPtr(date1), objc.ExtractPtr(date2))
	return rv
}

func (c_ Calendar) IsDateInToday(date IDate) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isDateInToday:"), objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) IsDateInTomorrow(date IDate) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isDateInTomorrow:"), objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) IsDateInWeekend(date IDate) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isDateInWeekend:"), objc.ExtractPtr(date))
	return rv
}

func (c_ Calendar) IsDateInYesterday(date IDate) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isDateInYesterday:"), objc.ExtractPtr(date))
	return rv
}

func (cc _CalendarClass) CurrentCalendar() Calendar {
	rv := objc.CallMethod[Calendar](cc, objc.GetSelector("currentCalendar"))
	return rv
}

func (cc _CalendarClass) AutoupdatingCurrentCalendar() Calendar {
	rv := objc.CallMethod[Calendar](cc, objc.GetSelector("autoupdatingCurrentCalendar"))
	return rv
}

func (c_ Calendar) CalendarIdentifier() CalendarIdentifier {
	rv := objc.CallMethod[CalendarIdentifier](c_, objc.GetSelector("calendarIdentifier"))
	return rv
}

func (c_ Calendar) FirstWeekday() uint {
	rv := objc.CallMethod[uint](c_, objc.GetSelector("firstWeekday"))
	return rv
}

func (c_ Calendar) SetFirstWeekday(value uint) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setFirstWeekday:"), value)
}

func (c_ Calendar) Locale() Locale {
	rv := objc.CallMethod[Locale](c_, objc.GetSelector("locale"))
	return rv
}

func (c_ Calendar) SetLocale(value ILocale) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setLocale:"), objc.ExtractPtr(value))
}

func (c_ Calendar) TimeZone() TimeZone {
	rv := objc.CallMethod[TimeZone](c_, objc.GetSelector("timeZone"))
	return rv
}

func (c_ Calendar) SetTimeZone(value ITimeZone) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setTimeZone:"), objc.ExtractPtr(value))
}

func (c_ Calendar) MinimumDaysInFirstWeek() uint {
	rv := objc.CallMethod[uint](c_, objc.GetSelector("minimumDaysInFirstWeek"))
	return rv
}

func (c_ Calendar) SetMinimumDaysInFirstWeek(value uint) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setMinimumDaysInFirstWeek:"), value)
}

func (c_ Calendar) AMSymbol() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("AMSymbol"))
	return rv
}

func (c_ Calendar) PMSymbol() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("PMSymbol"))
	return rv
}

func (c_ Calendar) WeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("weekdaySymbols"))
	return rv
}

func (c_ Calendar) ShortWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("shortWeekdaySymbols"))
	return rv
}

func (c_ Calendar) VeryShortWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("veryShortWeekdaySymbols"))
	return rv
}

func (c_ Calendar) StandaloneWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("standaloneWeekdaySymbols"))
	return rv
}

func (c_ Calendar) ShortStandaloneWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("shortStandaloneWeekdaySymbols"))
	return rv
}

func (c_ Calendar) VeryShortStandaloneWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("veryShortStandaloneWeekdaySymbols"))
	return rv
}

func (c_ Calendar) MonthSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("monthSymbols"))
	return rv
}

func (c_ Calendar) ShortMonthSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("shortMonthSymbols"))
	return rv
}

func (c_ Calendar) VeryShortMonthSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("veryShortMonthSymbols"))
	return rv
}

func (c_ Calendar) StandaloneMonthSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("standaloneMonthSymbols"))
	return rv
}

func (c_ Calendar) ShortStandaloneMonthSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("shortStandaloneMonthSymbols"))
	return rv
}

func (c_ Calendar) VeryShortStandaloneMonthSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("veryShortStandaloneMonthSymbols"))
	return rv
}

func (c_ Calendar) QuarterSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("quarterSymbols"))
	return rv
}

func (c_ Calendar) ShortQuarterSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("shortQuarterSymbols"))
	return rv
}

func (c_ Calendar) StandaloneQuarterSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("standaloneQuarterSymbols"))
	return rv
}

func (c_ Calendar) ShortStandaloneQuarterSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("shortStandaloneQuarterSymbols"))
	return rv
}

func (c_ Calendar) EraSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("eraSymbols"))
	return rv
}

func (c_ Calendar) LongEraSymbols() []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("longEraSymbols"))
	return rv
}
