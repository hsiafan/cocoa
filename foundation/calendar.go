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
	rv := objc.CallMethod[Calendar](cc, "alloc")
	return rv
}

func (cc _CalendarClass) New() Calendar {
	rv := objc.CallMethod[Calendar](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCalendar() Calendar {
	return CalendarClass.New()
}

func (c_ Calendar) Init() Calendar {
	rv := objc.CallMethod[Calendar](c_, "init")
	return rv
}

func (cc _CalendarClass) CalendarWithIdentifier(calendarIdentifierConstant CalendarIdentifier) Calendar {
	rv := objc.CallMethod[Calendar](cc, "calendarWithIdentifier:", calendarIdentifierConstant)
	return rv
}

func (c_ Calendar) InitWithCalendarIdentifier(ident CalendarIdentifier) objc.Object {
	rv := objc.CallMethod[objc.Object](c_, "initWithCalendarIdentifier:", ident)
	return rv
}

func (c_ Calendar) Date_MatchesComponents(date IDate, components IDateComponents) bool {
	rv := objc.CallMethod[bool](c_, "date:matchesComponents:", date, components)
	return rv
}

func (c_ Calendar) Component_FromDate(unit CalendarUnit, date IDate) int {
	rv := objc.CallMethod[int](c_, "component:fromDate:", unit, date)
	return rv
}

func (c_ Calendar) Components_FromDate(unitFlags CalendarUnit, date IDate) DateComponents {
	rv := objc.CallMethod[DateComponents](c_, "components:fromDate:", unitFlags, date)
	return rv
}

func (c_ Calendar) Components_FromDate_ToDate_Options(unitFlags CalendarUnit, startingDate IDate, resultDate IDate, opts CalendarOptions) DateComponents {
	rv := objc.CallMethod[DateComponents](c_, "components:fromDate:toDate:options:", unitFlags, startingDate, resultDate, opts)
	return rv
}

func (c_ Calendar) Components_FromDateComponents_ToDateComponents_Options(unitFlags CalendarUnit, startingDateComp IDateComponents, resultDateComp IDateComponents, options CalendarOptions) DateComponents {
	rv := objc.CallMethod[DateComponents](c_, "components:fromDateComponents:toDateComponents:options:", unitFlags, startingDateComp, resultDateComp, options)
	return rv
}

func (c_ Calendar) ComponentsInTimeZone_FromDate(timezone ITimeZone, date IDate) DateComponents {
	rv := objc.CallMethod[DateComponents](c_, "componentsInTimeZone:fromDate:", timezone, date)
	return rv
}

func (c_ Calendar) MaximumRangeOfUnit(unit CalendarUnit) Range {
	rv := objc.CallMethod[Range](c_, "maximumRangeOfUnit:", unit)
	return rv
}

func (c_ Calendar) MinimumRangeOfUnit(unit CalendarUnit) Range {
	rv := objc.CallMethod[Range](c_, "minimumRangeOfUnit:", unit)
	return rv
}

func (c_ Calendar) OrdinalityOfUnit_InUnit_ForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) uint {
	rv := objc.CallMethod[uint](c_, "ordinalityOfUnit:inUnit:forDate:", smaller, larger, date)
	return rv
}

func (c_ Calendar) RangeOfUnit_InUnit_ForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) Range {
	rv := objc.CallMethod[Range](c_, "rangeOfUnit:inUnit:forDate:", smaller, larger, date)
	return rv
}

func (c_ Calendar) RangeOfUnit_StartDate_Interval_ForDate(unit CalendarUnit, datep *Date, tip *TimeInterval, date IDate) bool {
	rv := objc.CallMethod[bool](c_, "rangeOfUnit:startDate:interval:forDate:", unit, unsafe.Pointer(datep), tip, date)
	return rv
}

func (c_ Calendar) StartOfDayForDate(date IDate) Date {
	rv := objc.CallMethod[Date](c_, "startOfDayForDate:", date)
	return rv
}

func (c_ Calendar) EnumerateDatesStartingAfterDate_MatchingComponents_Options_UsingBlock(start IDate, comps IDateComponents, opts CalendarOptions, block func(date Date, exactMatch bool, stop *bool)) {
	objc.CallMethod[objc.Void](c_, "enumerateDatesStartingAfterDate:matchingComponents:options:usingBlock:", start, comps, opts, block)
}

func (c_ Calendar) NextDateAfterDate_MatchingComponents_Options(date IDate, comps IDateComponents, options CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, "nextDateAfterDate:matchingComponents:options:", date, comps, options)
	return rv
}

func (c_ Calendar) NextDateAfterDate_MatchingHour_Minute_Second_Options(date IDate, hourValue int, minuteValue int, secondValue int, options CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, "nextDateAfterDate:matchingHour:minute:second:options:", date, hourValue, minuteValue, secondValue, options)
	return rv
}

func (c_ Calendar) NextDateAfterDate_MatchingUnit_Value_Options(date IDate, unit CalendarUnit, value int, options CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, "nextDateAfterDate:matchingUnit:value:options:", date, unit, value, options)
	return rv
}

func (c_ Calendar) DateFromComponents(comps IDateComponents) Date {
	rv := objc.CallMethod[Date](c_, "dateFromComponents:", comps)
	return rv
}

func (c_ Calendar) DateByAddingComponents_ToDate_Options(comps IDateComponents, date IDate, opts CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, "dateByAddingComponents:toDate:options:", comps, date, opts)
	return rv
}

func (c_ Calendar) DateByAddingUnit_Value_ToDate_Options(unit CalendarUnit, value int, date IDate, options CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, "dateByAddingUnit:value:toDate:options:", unit, value, date, options)
	return rv
}

func (c_ Calendar) DateBySettingHour_Minute_Second_OfDate_Options(h int, m int, s int, date IDate, opts CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, "dateBySettingHour:minute:second:ofDate:options:", h, m, s, date, opts)
	return rv
}

func (c_ Calendar) DateBySettingUnit_Value_OfDate_Options(unit CalendarUnit, v int, date IDate, opts CalendarOptions) Date {
	rv := objc.CallMethod[Date](c_, "dateBySettingUnit:value:ofDate:options:", unit, v, date, opts)
	return rv
}

func (c_ Calendar) DateWithEra_Year_Month_Day_Hour_Minute_Second_Nanosecond(eraValue int, yearValue int, monthValue int, dayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date {
	rv := objc.CallMethod[Date](c_, "dateWithEra:year:month:day:hour:minute:second:nanosecond:", eraValue, yearValue, monthValue, dayValue, hourValue, minuteValue, secondValue, nanosecondValue)
	return rv
}

func (c_ Calendar) DateWithEra_YearForWeekOfYear_WeekOfYear_Weekday_Hour_Minute_Second_Nanosecond(eraValue int, yearValue int, weekValue int, weekdayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date {
	rv := objc.CallMethod[Date](c_, "dateWithEra:yearForWeekOfYear:weekOfYear:weekday:hour:minute:second:nanosecond:", eraValue, yearValue, weekValue, weekdayValue, hourValue, minuteValue, secondValue, nanosecondValue)
	return rv
}

func (c_ Calendar) CompareDate_ToDate_ToUnitGranularity(date1 IDate, date2 IDate, unit CalendarUnit) ComparisonResult {
	rv := objc.CallMethod[ComparisonResult](c_, "compareDate:toDate:toUnitGranularity:", date1, date2, unit)
	return rv
}

func (c_ Calendar) IsDate_EqualToDate_ToUnitGranularity(date1 IDate, date2 IDate, unit CalendarUnit) bool {
	rv := objc.CallMethod[bool](c_, "isDate:equalToDate:toUnitGranularity:", date1, date2, unit)
	return rv
}

func (c_ Calendar) IsDate_InSameDayAsDate(date1 IDate, date2 IDate) bool {
	rv := objc.CallMethod[bool](c_, "isDate:inSameDayAsDate:", date1, date2)
	return rv
}

func (c_ Calendar) IsDateInToday(date IDate) bool {
	rv := objc.CallMethod[bool](c_, "isDateInToday:", date)
	return rv
}

func (c_ Calendar) IsDateInTomorrow(date IDate) bool {
	rv := objc.CallMethod[bool](c_, "isDateInTomorrow:", date)
	return rv
}

func (c_ Calendar) IsDateInWeekend(date IDate) bool {
	rv := objc.CallMethod[bool](c_, "isDateInWeekend:", date)
	return rv
}

func (c_ Calendar) IsDateInYesterday(date IDate) bool {
	rv := objc.CallMethod[bool](c_, "isDateInYesterday:", date)
	return rv
}

func (cc _CalendarClass) CurrentCalendar() Calendar {
	rv := objc.CallMethod[Calendar](cc, "currentCalendar")
	return rv
}

func (cc _CalendarClass) AutoupdatingCurrentCalendar() Calendar {
	rv := objc.CallMethod[Calendar](cc, "autoupdatingCurrentCalendar")
	return rv
}

func (c_ Calendar) CalendarIdentifier() CalendarIdentifier {
	rv := objc.CallMethod[CalendarIdentifier](c_, "calendarIdentifier")
	return rv
}

func (c_ Calendar) FirstWeekday() uint {
	rv := objc.CallMethod[uint](c_, "firstWeekday")
	return rv
}

func (c_ Calendar) SetFirstWeekday(value uint) {
	objc.CallMethod[objc.Void](c_, "setFirstWeekday:", value)
}

func (c_ Calendar) Locale() Locale {
	rv := objc.CallMethod[Locale](c_, "locale")
	return rv
}

func (c_ Calendar) SetLocale(value ILocale) {
	objc.CallMethod[objc.Void](c_, "setLocale:", value)
}

func (c_ Calendar) TimeZone() TimeZone {
	rv := objc.CallMethod[TimeZone](c_, "timeZone")
	return rv
}

func (c_ Calendar) SetTimeZone(value ITimeZone) {
	objc.CallMethod[objc.Void](c_, "setTimeZone:", value)
}

func (c_ Calendar) MinimumDaysInFirstWeek() uint {
	rv := objc.CallMethod[uint](c_, "minimumDaysInFirstWeek")
	return rv
}

func (c_ Calendar) SetMinimumDaysInFirstWeek(value uint) {
	objc.CallMethod[objc.Void](c_, "setMinimumDaysInFirstWeek:", value)
}

func (c_ Calendar) AMSymbol() string {
	rv := objc.CallMethod[string](c_, "AMSymbol")
	return rv
}

func (c_ Calendar) PMSymbol() string {
	rv := objc.CallMethod[string](c_, "PMSymbol")
	return rv
}

func (c_ Calendar) WeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](c_, "weekdaySymbols")
	return rv
}

func (c_ Calendar) ShortWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](c_, "shortWeekdaySymbols")
	return rv
}

func (c_ Calendar) VeryShortWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](c_, "veryShortWeekdaySymbols")
	return rv
}

func (c_ Calendar) StandaloneWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](c_, "standaloneWeekdaySymbols")
	return rv
}

func (c_ Calendar) ShortStandaloneWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](c_, "shortStandaloneWeekdaySymbols")
	return rv
}

func (c_ Calendar) VeryShortStandaloneWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](c_, "veryShortStandaloneWeekdaySymbols")
	return rv
}

func (c_ Calendar) MonthSymbols() []string {
	rv := objc.CallMethod[[]string](c_, "monthSymbols")
	return rv
}

func (c_ Calendar) ShortMonthSymbols() []string {
	rv := objc.CallMethod[[]string](c_, "shortMonthSymbols")
	return rv
}

func (c_ Calendar) VeryShortMonthSymbols() []string {
	rv := objc.CallMethod[[]string](c_, "veryShortMonthSymbols")
	return rv
}

func (c_ Calendar) StandaloneMonthSymbols() []string {
	rv := objc.CallMethod[[]string](c_, "standaloneMonthSymbols")
	return rv
}

func (c_ Calendar) ShortStandaloneMonthSymbols() []string {
	rv := objc.CallMethod[[]string](c_, "shortStandaloneMonthSymbols")
	return rv
}

func (c_ Calendar) VeryShortStandaloneMonthSymbols() []string {
	rv := objc.CallMethod[[]string](c_, "veryShortStandaloneMonthSymbols")
	return rv
}

func (c_ Calendar) QuarterSymbols() []string {
	rv := objc.CallMethod[[]string](c_, "quarterSymbols")
	return rv
}

func (c_ Calendar) ShortQuarterSymbols() []string {
	rv := objc.CallMethod[[]string](c_, "shortQuarterSymbols")
	return rv
}

func (c_ Calendar) StandaloneQuarterSymbols() []string {
	rv := objc.CallMethod[[]string](c_, "standaloneQuarterSymbols")
	return rv
}

func (c_ Calendar) ShortStandaloneQuarterSymbols() []string {
	rv := objc.CallMethod[[]string](c_, "shortStandaloneQuarterSymbols")
	return rv
}

func (c_ Calendar) EraSymbols() []string {
	rv := objc.CallMethod[[]string](c_, "eraSymbols")
	return rv
}

func (c_ Calendar) LongEraSymbols() []string {
	rv := objc.CallMethod[[]string](c_, "longEraSymbols")
	return rv
}
