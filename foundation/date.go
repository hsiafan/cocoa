// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var DateClass = _DateClass{objc.GetClass("NSDate")}

type _DateClass struct {
	objc.Class
}

type IDate interface {
	objc.IObject
	IsEqualToDate(otherDate IDate) bool
	EarlierDate(anotherDate IDate) Date
	LaterDate(anotherDate IDate) Date
	Compare(other IDate) ComparisonResult
	TimeIntervalSinceDate(anotherDate IDate) TimeInterval
	DescriptionWithLocale(locale objc.IObject) string
	// deprecated
	InitWithString(description string) objc.Object
	// deprecated
	AddTimeInterval(seconds TimeInterval) objc.Object
	// deprecated
	DescriptionWithCalendarFormat_TimeZone_Locale(format string, aTimeZone ITimeZone, locale objc.IObject) string
	TimeIntervalSinceNow() TimeInterval
	TimeIntervalSinceReferenceDate() TimeInterval
	TimeIntervalSince1970() TimeInterval
	Description() string
}

type Date struct {
	objc.Object
}

func MakeDate(ptr unsafe.Pointer) Date {
	return Date{
		Object: objc.MakeObject(ptr),
	}
}

func (dc _DateClass) Date() Date {
	rv := ffi.CallMethod[Date](dc, "date")
	return rv
}

func (dc _DateClass) DateWithTimeIntervalSinceNow(secs TimeInterval) Date {
	rv := ffi.CallMethod[Date](dc, "dateWithTimeIntervalSinceNow:", secs)
	return rv
}

func (dc _DateClass) DateWithTimeInterval_SinceDate(secsToBeAdded TimeInterval, date IDate) Date {
	rv := ffi.CallMethod[Date](dc, "dateWithTimeInterval:sinceDate:", secsToBeAdded, date)
	return rv
}

func (dc _DateClass) DateWithTimeIntervalSinceReferenceDate(ti TimeInterval) Date {
	rv := ffi.CallMethod[Date](dc, "dateWithTimeIntervalSinceReferenceDate:", ti)
	return rv
}

func (dc _DateClass) DateWithTimeIntervalSince1970(secs TimeInterval) Date {
	rv := ffi.CallMethod[Date](dc, "dateWithTimeIntervalSince1970:", secs)
	return rv
}

func (d_ Date) Init() Date {
	rv := ffi.CallMethod[Date](d_, "init")
	rv.Autorelease()
	return rv
}

func (d_ Date) InitWithTimeIntervalSinceNow(secs TimeInterval) Date {
	rv := ffi.CallMethod[Date](d_, "initWithTimeIntervalSinceNow:", secs)
	rv.Autorelease()
	return rv
}

func (d_ Date) InitWithTimeInterval_SinceDate(secsToBeAdded TimeInterval, date IDate) Date {
	rv := ffi.CallMethod[Date](d_, "initWithTimeInterval:sinceDate:", secsToBeAdded, date)
	rv.Autorelease()
	return rv
}

func (d_ Date) InitWithTimeIntervalSinceReferenceDate(ti TimeInterval) Date {
	rv := ffi.CallMethod[Date](d_, "initWithTimeIntervalSinceReferenceDate:", ti)
	rv.Autorelease()
	return rv
}

func (d_ Date) InitWithTimeIntervalSince1970(secs TimeInterval) Date {
	rv := ffi.CallMethod[Date](d_, "initWithTimeIntervalSince1970:", secs)
	rv.Autorelease()
	return rv
}

func (d_ Date) DateByAddingTimeInterval(ti TimeInterval) Date {
	rv := ffi.CallMethod[Date](d_, "dateByAddingTimeInterval:", ti)
	return rv
}

func (dc _DateClass) Alloc() Date {
	rv := ffi.CallMethod[Date](dc, "alloc")
	return rv
}

func (dc _DateClass) New() Date {
	rv := ffi.CallMethod[Date](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDate() Date {
	return DateClass.New()
}

func (d_ Date) IsEqualToDate(otherDate IDate) bool {
	rv := ffi.CallMethod[bool](d_, "isEqualToDate:", otherDate)
	return rv
}

func (d_ Date) EarlierDate(anotherDate IDate) Date {
	rv := ffi.CallMethod[Date](d_, "earlierDate:", anotherDate)
	return rv
}

func (d_ Date) LaterDate(anotherDate IDate) Date {
	rv := ffi.CallMethod[Date](d_, "laterDate:", anotherDate)
	return rv
}

func (d_ Date) Compare(other IDate) ComparisonResult {
	rv := ffi.CallMethod[ComparisonResult](d_, "compare:", other)
	return rv
}

func (d_ Date) TimeIntervalSinceDate(anotherDate IDate) TimeInterval {
	rv := ffi.CallMethod[TimeInterval](d_, "timeIntervalSinceDate:", anotherDate)
	return rv
}

func (d_ Date) DescriptionWithLocale(locale objc.IObject) string {
	rv := ffi.CallMethod[string](d_, "descriptionWithLocale:", locale)
	return rv
}

// deprecated
func (dc _DateClass) DateWithNaturalLanguageString(_string string) objc.Object {
	rv := ffi.CallMethod[objc.Object](dc, "dateWithNaturalLanguageString:", _string)
	return rv
}

// deprecated
func (dc _DateClass) DateWithNaturalLanguageString_Locale(_string string, locale objc.IObject) objc.Object {
	rv := ffi.CallMethod[objc.Object](dc, "dateWithNaturalLanguageString:locale:", _string, locale)
	return rv
}

// deprecated
func (dc _DateClass) DateWithString(aString string) objc.Object {
	rv := ffi.CallMethod[objc.Object](dc, "dateWithString:", aString)
	return rv
}

// deprecated
func (d_ Date) InitWithString(description string) objc.Object {
	rv := ffi.CallMethod[objc.Object](d_, "initWithString:", description)
	rv.Autorelease()
	return rv
}

// deprecated
func (d_ Date) AddTimeInterval(seconds TimeInterval) objc.Object {
	rv := ffi.CallMethod[objc.Object](d_, "addTimeInterval:", seconds)
	return rv
}

// deprecated
func (d_ Date) DescriptionWithCalendarFormat_TimeZone_Locale(format string, aTimeZone ITimeZone, locale objc.IObject) string {
	rv := ffi.CallMethod[string](d_, "descriptionWithCalendarFormat:timeZone:locale:", format, aTimeZone, locale)
	return rv
}

func (dc _DateClass) DistantFuture() Date {
	rv := ffi.CallMethod[Date](dc, "distantFuture")
	return rv
}

func (dc _DateClass) DistantPast() Date {
	rv := ffi.CallMethod[Date](dc, "distantPast")
	return rv
}

func (dc _DateClass) Now() Date {
	rv := ffi.CallMethod[Date](dc, "now")
	return rv
}

func (d_ Date) TimeIntervalSinceNow() TimeInterval {
	rv := ffi.CallMethod[TimeInterval](d_, "timeIntervalSinceNow")
	return rv
}

func (d_ Date) TimeIntervalSinceReferenceDate() TimeInterval {
	rv := ffi.CallMethod[TimeInterval](d_, "timeIntervalSinceReferenceDate")
	return rv
}

func (d_ Date) TimeIntervalSince1970() TimeInterval {
	rv := ffi.CallMethod[TimeInterval](d_, "timeIntervalSince1970")
	return rv
}

func (d_ Date) Description() string {
	rv := ffi.CallMethod[string](d_, "description")
	return rv
}

var DateFormatterClass = _DateFormatterClass{objc.GetClass("NSDateFormatter")}

type _DateFormatterClass struct {
	objc.Class
}

type IDateFormatter interface {
	IFormatter
	DateFromString(_string string) Date
	StringFromDate(date IDate) string
	SetLocalizedDateFormatFromTemplate(dateFormatTemplate string)
	// deprecated
	AllowsNaturalLanguage() bool
	// deprecated
	InitWithDateFormat_AllowNaturalLanguage(format string, flag bool) objc.Object
	DateStyle() DateFormatterStyle
	SetDateStyle(value DateFormatterStyle)
	TimeStyle() DateFormatterStyle
	SetTimeStyle(value DateFormatterStyle)
	DateFormat() string
	SetDateFormat(value string)
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	Calendar() Calendar
	SetCalendar(value ICalendar)
	DefaultDate() Date
	SetDefaultDate(value IDate)
	Locale() Locale
	SetLocale(value ILocale)
	TimeZone() TimeZone
	SetTimeZone(value ITimeZone)
	TwoDigitStartDate() Date
	SetTwoDigitStartDate(value IDate)
	GregorianStartDate() Date
	SetGregorianStartDate(value IDate)
	FormatterBehavior() DateFormatterBehavior
	SetFormatterBehavior(value DateFormatterBehavior)
	IsLenient() bool
	SetLenient(value bool)
	DoesRelativeDateFormatting() bool
	SetDoesRelativeDateFormatting(value bool)
	AMSymbol() string
	SetAMSymbol(value string)
	PMSymbol() string
	SetPMSymbol(value string)
	WeekdaySymbols() []string
	SetWeekdaySymbols(value []string)
	ShortWeekdaySymbols() []string
	SetShortWeekdaySymbols(value []string)
	VeryShortWeekdaySymbols() []string
	SetVeryShortWeekdaySymbols(value []string)
	StandaloneWeekdaySymbols() []string
	SetStandaloneWeekdaySymbols(value []string)
	ShortStandaloneWeekdaySymbols() []string
	SetShortStandaloneWeekdaySymbols(value []string)
	VeryShortStandaloneWeekdaySymbols() []string
	SetVeryShortStandaloneWeekdaySymbols(value []string)
	MonthSymbols() []string
	SetMonthSymbols(value []string)
	ShortMonthSymbols() []string
	SetShortMonthSymbols(value []string)
	VeryShortMonthSymbols() []string
	SetVeryShortMonthSymbols(value []string)
	StandaloneMonthSymbols() []string
	SetStandaloneMonthSymbols(value []string)
	ShortStandaloneMonthSymbols() []string
	SetShortStandaloneMonthSymbols(value []string)
	VeryShortStandaloneMonthSymbols() []string
	SetVeryShortStandaloneMonthSymbols(value []string)
	QuarterSymbols() []string
	SetQuarterSymbols(value []string)
	ShortQuarterSymbols() []string
	SetShortQuarterSymbols(value []string)
	StandaloneQuarterSymbols() []string
	SetStandaloneQuarterSymbols(value []string)
	ShortStandaloneQuarterSymbols() []string
	SetShortStandaloneQuarterSymbols(value []string)
	EraSymbols() []string
	SetEraSymbols(value []string)
	LongEraSymbols() []string
	SetLongEraSymbols(value []string)
	GeneratesCalendarDates() bool
	SetGeneratesCalendarDates(value bool)
}

type DateFormatter struct {
	Formatter
}

func MakeDateFormatter(ptr unsafe.Pointer) DateFormatter {
	return DateFormatter{
		Formatter: MakeFormatter(ptr),
	}
}

func (dc _DateFormatterClass) Alloc() DateFormatter {
	rv := ffi.CallMethod[DateFormatter](dc, "alloc")
	return rv
}

func (d_ DateFormatter) Init() DateFormatter {
	rv := ffi.CallMethod[DateFormatter](d_, "init")
	rv.Autorelease()
	return rv
}

func (dc _DateFormatterClass) New() DateFormatter {
	rv := ffi.CallMethod[DateFormatter](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDateFormatter() DateFormatter {
	return DateFormatterClass.New()
}

func (d_ DateFormatter) DateFromString(_string string) Date {
	rv := ffi.CallMethod[Date](d_, "dateFromString:", _string)
	return rv
}

func (d_ DateFormatter) StringFromDate(date IDate) string {
	rv := ffi.CallMethod[string](d_, "stringFromDate:", date)
	return rv
}

func (dc _DateFormatterClass) LocalizedStringFromDate_DateStyle_TimeStyle(date IDate, dstyle DateFormatterStyle, tstyle DateFormatterStyle) string {
	rv := ffi.CallMethod[string](dc, "localizedStringFromDate:dateStyle:timeStyle:", date, dstyle, tstyle)
	return rv
}

func (d_ DateFormatter) SetLocalizedDateFormatFromTemplate(dateFormatTemplate string) {
	ffi.CallMethod[ffi.Void](d_, "setLocalizedDateFormatFromTemplate:", dateFormatTemplate)
}

func (dc _DateFormatterClass) DateFormatFromTemplate_Options_Locale(tmplate string, opts uint, locale ILocale) string {
	rv := ffi.CallMethod[string](dc, "dateFormatFromTemplate:options:locale:", tmplate, opts, locale)
	return rv
}

// deprecated
func (d_ DateFormatter) AllowsNaturalLanguage() bool {
	rv := ffi.CallMethod[bool](d_, "allowsNaturalLanguage")
	return rv
}

// deprecated
func (d_ DateFormatter) InitWithDateFormat_AllowNaturalLanguage(format string, flag bool) objc.Object {
	rv := ffi.CallMethod[objc.Object](d_, "initWithDateFormat:allowNaturalLanguage:", format, flag)
	rv.Autorelease()
	return rv
}

func (d_ DateFormatter) DateStyle() DateFormatterStyle {
	rv := ffi.CallMethod[DateFormatterStyle](d_, "dateStyle")
	return rv
}

func (d_ DateFormatter) SetDateStyle(value DateFormatterStyle) {
	ffi.CallMethod[ffi.Void](d_, "setDateStyle:", value)
}

func (d_ DateFormatter) TimeStyle() DateFormatterStyle {
	rv := ffi.CallMethod[DateFormatterStyle](d_, "timeStyle")
	return rv
}

func (d_ DateFormatter) SetTimeStyle(value DateFormatterStyle) {
	ffi.CallMethod[ffi.Void](d_, "setTimeStyle:", value)
}

func (d_ DateFormatter) DateFormat() string {
	rv := ffi.CallMethod[string](d_, "dateFormat")
	return rv
}

func (d_ DateFormatter) SetDateFormat(value string) {
	ffi.CallMethod[ffi.Void](d_, "setDateFormat:", value)
}

func (d_ DateFormatter) FormattingContext() FormattingContext {
	rv := ffi.CallMethod[FormattingContext](d_, "formattingContext")
	return rv
}

func (d_ DateFormatter) SetFormattingContext(value FormattingContext) {
	ffi.CallMethod[ffi.Void](d_, "setFormattingContext:", value)
}

func (d_ DateFormatter) Calendar() Calendar {
	rv := ffi.CallMethod[Calendar](d_, "calendar")
	return rv
}

func (d_ DateFormatter) SetCalendar(value ICalendar) {
	ffi.CallMethod[ffi.Void](d_, "setCalendar:", value)
}

func (d_ DateFormatter) DefaultDate() Date {
	rv := ffi.CallMethod[Date](d_, "defaultDate")
	return rv
}

func (d_ DateFormatter) SetDefaultDate(value IDate) {
	ffi.CallMethod[ffi.Void](d_, "setDefaultDate:", value)
}

func (d_ DateFormatter) Locale() Locale {
	rv := ffi.CallMethod[Locale](d_, "locale")
	return rv
}

func (d_ DateFormatter) SetLocale(value ILocale) {
	ffi.CallMethod[ffi.Void](d_, "setLocale:", value)
}

func (d_ DateFormatter) TimeZone() TimeZone {
	rv := ffi.CallMethod[TimeZone](d_, "timeZone")
	return rv
}

func (d_ DateFormatter) SetTimeZone(value ITimeZone) {
	ffi.CallMethod[ffi.Void](d_, "setTimeZone:", value)
}

func (d_ DateFormatter) TwoDigitStartDate() Date {
	rv := ffi.CallMethod[Date](d_, "twoDigitStartDate")
	return rv
}

func (d_ DateFormatter) SetTwoDigitStartDate(value IDate) {
	ffi.CallMethod[ffi.Void](d_, "setTwoDigitStartDate:", value)
}

func (d_ DateFormatter) GregorianStartDate() Date {
	rv := ffi.CallMethod[Date](d_, "gregorianStartDate")
	return rv
}

func (d_ DateFormatter) SetGregorianStartDate(value IDate) {
	ffi.CallMethod[ffi.Void](d_, "setGregorianStartDate:", value)
}

func (d_ DateFormatter) FormatterBehavior() DateFormatterBehavior {
	rv := ffi.CallMethod[DateFormatterBehavior](d_, "formatterBehavior")
	return rv
}

func (d_ DateFormatter) SetFormatterBehavior(value DateFormatterBehavior) {
	ffi.CallMethod[ffi.Void](d_, "setFormatterBehavior:", value)
}

func (dc _DateFormatterClass) DefaultFormatterBehavior() DateFormatterBehavior {
	rv := ffi.CallMethod[DateFormatterBehavior](dc, "defaultFormatterBehavior")
	return rv
}

func (dc _DateFormatterClass) SetDefaultFormatterBehavior(value DateFormatterBehavior) {
	ffi.CallMethod[ffi.Void](dc, "setDefaultFormatterBehavior:", value)
}

func (d_ DateFormatter) IsLenient() bool {
	rv := ffi.CallMethod[bool](d_, "isLenient")
	return rv
}

func (d_ DateFormatter) SetLenient(value bool) {
	ffi.CallMethod[ffi.Void](d_, "setLenient:", value)
}

func (d_ DateFormatter) DoesRelativeDateFormatting() bool {
	rv := ffi.CallMethod[bool](d_, "doesRelativeDateFormatting")
	return rv
}

func (d_ DateFormatter) SetDoesRelativeDateFormatting(value bool) {
	ffi.CallMethod[ffi.Void](d_, "setDoesRelativeDateFormatting:", value)
}

func (d_ DateFormatter) AMSymbol() string {
	rv := ffi.CallMethod[string](d_, "AMSymbol")
	return rv
}

func (d_ DateFormatter) SetAMSymbol(value string) {
	ffi.CallMethod[ffi.Void](d_, "setAMSymbol:", value)
}

func (d_ DateFormatter) PMSymbol() string {
	rv := ffi.CallMethod[string](d_, "PMSymbol")
	return rv
}

func (d_ DateFormatter) SetPMSymbol(value string) {
	ffi.CallMethod[ffi.Void](d_, "setPMSymbol:", value)
}

func (d_ DateFormatter) WeekdaySymbols() []string {
	rv := ffi.CallMethod[[]string](d_, "weekdaySymbols")
	return rv
}

func (d_ DateFormatter) SetWeekdaySymbols(value []string) {
	ffi.CallMethod[ffi.Void](d_, "setWeekdaySymbols:", value)
}

func (d_ DateFormatter) ShortWeekdaySymbols() []string {
	rv := ffi.CallMethod[[]string](d_, "shortWeekdaySymbols")
	return rv
}

func (d_ DateFormatter) SetShortWeekdaySymbols(value []string) {
	ffi.CallMethod[ffi.Void](d_, "setShortWeekdaySymbols:", value)
}

func (d_ DateFormatter) VeryShortWeekdaySymbols() []string {
	rv := ffi.CallMethod[[]string](d_, "veryShortWeekdaySymbols")
	return rv
}

func (d_ DateFormatter) SetVeryShortWeekdaySymbols(value []string) {
	ffi.CallMethod[ffi.Void](d_, "setVeryShortWeekdaySymbols:", value)
}

func (d_ DateFormatter) StandaloneWeekdaySymbols() []string {
	rv := ffi.CallMethod[[]string](d_, "standaloneWeekdaySymbols")
	return rv
}

func (d_ DateFormatter) SetStandaloneWeekdaySymbols(value []string) {
	ffi.CallMethod[ffi.Void](d_, "setStandaloneWeekdaySymbols:", value)
}

func (d_ DateFormatter) ShortStandaloneWeekdaySymbols() []string {
	rv := ffi.CallMethod[[]string](d_, "shortStandaloneWeekdaySymbols")
	return rv
}

func (d_ DateFormatter) SetShortStandaloneWeekdaySymbols(value []string) {
	ffi.CallMethod[ffi.Void](d_, "setShortStandaloneWeekdaySymbols:", value)
}

func (d_ DateFormatter) VeryShortStandaloneWeekdaySymbols() []string {
	rv := ffi.CallMethod[[]string](d_, "veryShortStandaloneWeekdaySymbols")
	return rv
}

func (d_ DateFormatter) SetVeryShortStandaloneWeekdaySymbols(value []string) {
	ffi.CallMethod[ffi.Void](d_, "setVeryShortStandaloneWeekdaySymbols:", value)
}

func (d_ DateFormatter) MonthSymbols() []string {
	rv := ffi.CallMethod[[]string](d_, "monthSymbols")
	return rv
}

func (d_ DateFormatter) SetMonthSymbols(value []string) {
	ffi.CallMethod[ffi.Void](d_, "setMonthSymbols:", value)
}

func (d_ DateFormatter) ShortMonthSymbols() []string {
	rv := ffi.CallMethod[[]string](d_, "shortMonthSymbols")
	return rv
}

func (d_ DateFormatter) SetShortMonthSymbols(value []string) {
	ffi.CallMethod[ffi.Void](d_, "setShortMonthSymbols:", value)
}

func (d_ DateFormatter) VeryShortMonthSymbols() []string {
	rv := ffi.CallMethod[[]string](d_, "veryShortMonthSymbols")
	return rv
}

func (d_ DateFormatter) SetVeryShortMonthSymbols(value []string) {
	ffi.CallMethod[ffi.Void](d_, "setVeryShortMonthSymbols:", value)
}

func (d_ DateFormatter) StandaloneMonthSymbols() []string {
	rv := ffi.CallMethod[[]string](d_, "standaloneMonthSymbols")
	return rv
}

func (d_ DateFormatter) SetStandaloneMonthSymbols(value []string) {
	ffi.CallMethod[ffi.Void](d_, "setStandaloneMonthSymbols:", value)
}

func (d_ DateFormatter) ShortStandaloneMonthSymbols() []string {
	rv := ffi.CallMethod[[]string](d_, "shortStandaloneMonthSymbols")
	return rv
}

func (d_ DateFormatter) SetShortStandaloneMonthSymbols(value []string) {
	ffi.CallMethod[ffi.Void](d_, "setShortStandaloneMonthSymbols:", value)
}

func (d_ DateFormatter) VeryShortStandaloneMonthSymbols() []string {
	rv := ffi.CallMethod[[]string](d_, "veryShortStandaloneMonthSymbols")
	return rv
}

func (d_ DateFormatter) SetVeryShortStandaloneMonthSymbols(value []string) {
	ffi.CallMethod[ffi.Void](d_, "setVeryShortStandaloneMonthSymbols:", value)
}

func (d_ DateFormatter) QuarterSymbols() []string {
	rv := ffi.CallMethod[[]string](d_, "quarterSymbols")
	return rv
}

func (d_ DateFormatter) SetQuarterSymbols(value []string) {
	ffi.CallMethod[ffi.Void](d_, "setQuarterSymbols:", value)
}

func (d_ DateFormatter) ShortQuarterSymbols() []string {
	rv := ffi.CallMethod[[]string](d_, "shortQuarterSymbols")
	return rv
}

func (d_ DateFormatter) SetShortQuarterSymbols(value []string) {
	ffi.CallMethod[ffi.Void](d_, "setShortQuarterSymbols:", value)
}

func (d_ DateFormatter) StandaloneQuarterSymbols() []string {
	rv := ffi.CallMethod[[]string](d_, "standaloneQuarterSymbols")
	return rv
}

func (d_ DateFormatter) SetStandaloneQuarterSymbols(value []string) {
	ffi.CallMethod[ffi.Void](d_, "setStandaloneQuarterSymbols:", value)
}

func (d_ DateFormatter) ShortStandaloneQuarterSymbols() []string {
	rv := ffi.CallMethod[[]string](d_, "shortStandaloneQuarterSymbols")
	return rv
}

func (d_ DateFormatter) SetShortStandaloneQuarterSymbols(value []string) {
	ffi.CallMethod[ffi.Void](d_, "setShortStandaloneQuarterSymbols:", value)
}

func (d_ DateFormatter) EraSymbols() []string {
	rv := ffi.CallMethod[[]string](d_, "eraSymbols")
	return rv
}

func (d_ DateFormatter) SetEraSymbols(value []string) {
	ffi.CallMethod[ffi.Void](d_, "setEraSymbols:", value)
}

func (d_ DateFormatter) LongEraSymbols() []string {
	rv := ffi.CallMethod[[]string](d_, "longEraSymbols")
	return rv
}

func (d_ DateFormatter) SetLongEraSymbols(value []string) {
	ffi.CallMethod[ffi.Void](d_, "setLongEraSymbols:", value)
}

func (d_ DateFormatter) GeneratesCalendarDates() bool {
	rv := ffi.CallMethod[bool](d_, "generatesCalendarDates")
	return rv
}

func (d_ DateFormatter) SetGeneratesCalendarDates(value bool) {
	ffi.CallMethod[ffi.Void](d_, "setGeneratesCalendarDates:", value)
}

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

func (d_ DateComponents) Init() DateComponents {
	rv := ffi.CallMethod[DateComponents](d_, "init")
	rv.Autorelease()
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
	rv := ffi.CallMethod[DateComponentsFormatter](dc, "alloc")
	return rv
}

func (d_ DateComponentsFormatter) Init() DateComponentsFormatter {
	rv := ffi.CallMethod[DateComponentsFormatter](d_, "init")
	rv.Autorelease()
	return rv
}

func (dc _DateComponentsFormatterClass) New() DateComponentsFormatter {
	rv := ffi.CallMethod[DateComponentsFormatter](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDateComponentsFormatter() DateComponentsFormatter {
	return DateComponentsFormatterClass.New()
}

func (d_ DateComponentsFormatter) StringFromDateComponents(components IDateComponents) string {
	rv := ffi.CallMethod[string](d_, "stringFromDateComponents:", components)
	return rv
}

func (d_ DateComponentsFormatter) StringFromDate_ToDate(startDate IDate, endDate IDate) string {
	rv := ffi.CallMethod[string](d_, "stringFromDate:toDate:", startDate, endDate)
	return rv
}

func (d_ DateComponentsFormatter) StringFromTimeInterval(ti TimeInterval) string {
	rv := ffi.CallMethod[string](d_, "stringFromTimeInterval:", ti)
	return rv
}

func (dc _DateComponentsFormatterClass) LocalizedStringFromDateComponents_UnitsStyle(components IDateComponents, unitsStyle DateComponentsFormatterUnitsStyle) string {
	rv := ffi.CallMethod[string](dc, "localizedStringFromDateComponents:unitsStyle:", components, unitsStyle)
	return rv
}

func (d_ DateComponentsFormatter) AllowedUnits() CalendarUnit {
	rv := ffi.CallMethod[CalendarUnit](d_, "allowedUnits")
	return rv
}

func (d_ DateComponentsFormatter) SetAllowedUnits(value CalendarUnit) {
	ffi.CallMethod[ffi.Void](d_, "setAllowedUnits:", value)
}

func (d_ DateComponentsFormatter) AllowsFractionalUnits() bool {
	rv := ffi.CallMethod[bool](d_, "allowsFractionalUnits")
	return rv
}

func (d_ DateComponentsFormatter) SetAllowsFractionalUnits(value bool) {
	ffi.CallMethod[ffi.Void](d_, "setAllowsFractionalUnits:", value)
}

func (d_ DateComponentsFormatter) Calendar() Calendar {
	rv := ffi.CallMethod[Calendar](d_, "calendar")
	return rv
}

func (d_ DateComponentsFormatter) SetCalendar(value ICalendar) {
	ffi.CallMethod[ffi.Void](d_, "setCalendar:", value)
}

func (d_ DateComponentsFormatter) CollapsesLargestUnit() bool {
	rv := ffi.CallMethod[bool](d_, "collapsesLargestUnit")
	return rv
}

func (d_ DateComponentsFormatter) SetCollapsesLargestUnit(value bool) {
	ffi.CallMethod[ffi.Void](d_, "setCollapsesLargestUnit:", value)
}

func (d_ DateComponentsFormatter) IncludesApproximationPhrase() bool {
	rv := ffi.CallMethod[bool](d_, "includesApproximationPhrase")
	return rv
}

func (d_ DateComponentsFormatter) SetIncludesApproximationPhrase(value bool) {
	ffi.CallMethod[ffi.Void](d_, "setIncludesApproximationPhrase:", value)
}

func (d_ DateComponentsFormatter) IncludesTimeRemainingPhrase() bool {
	rv := ffi.CallMethod[bool](d_, "includesTimeRemainingPhrase")
	return rv
}

func (d_ DateComponentsFormatter) SetIncludesTimeRemainingPhrase(value bool) {
	ffi.CallMethod[ffi.Void](d_, "setIncludesTimeRemainingPhrase:", value)
}

func (d_ DateComponentsFormatter) MaximumUnitCount() int {
	rv := ffi.CallMethod[int](d_, "maximumUnitCount")
	return rv
}

func (d_ DateComponentsFormatter) SetMaximumUnitCount(value int) {
	ffi.CallMethod[ffi.Void](d_, "setMaximumUnitCount:", value)
}

func (d_ DateComponentsFormatter) UnitsStyle() DateComponentsFormatterUnitsStyle {
	rv := ffi.CallMethod[DateComponentsFormatterUnitsStyle](d_, "unitsStyle")
	return rv
}

func (d_ DateComponentsFormatter) SetUnitsStyle(value DateComponentsFormatterUnitsStyle) {
	ffi.CallMethod[ffi.Void](d_, "setUnitsStyle:", value)
}

func (d_ DateComponentsFormatter) ZeroFormattingBehavior() DateComponentsFormatterZeroFormattingBehavior {
	rv := ffi.CallMethod[DateComponentsFormatterZeroFormattingBehavior](d_, "zeroFormattingBehavior")
	return rv
}

func (d_ DateComponentsFormatter) SetZeroFormattingBehavior(value DateComponentsFormatterZeroFormattingBehavior) {
	ffi.CallMethod[ffi.Void](d_, "setZeroFormattingBehavior:", value)
}

func (d_ DateComponentsFormatter) FormattingContext() FormattingContext {
	rv := ffi.CallMethod[FormattingContext](d_, "formattingContext")
	return rv
}

func (d_ DateComponentsFormatter) SetFormattingContext(value FormattingContext) {
	ffi.CallMethod[ffi.Void](d_, "setFormattingContext:", value)
}

func (d_ DateComponentsFormatter) ReferenceDate() Date {
	rv := ffi.CallMethod[Date](d_, "referenceDate")
	return rv
}

func (d_ DateComponentsFormatter) SetReferenceDate(value IDate) {
	ffi.CallMethod[ffi.Void](d_, "setReferenceDate:", value)
}

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
	rv.Autorelease()
	return rv
}

func (d_ DateInterval) InitWithStartDate_Duration(startDate IDate, duration TimeInterval) DateInterval {
	rv := ffi.CallMethod[DateInterval](d_, "initWithStartDate:duration:", startDate, duration)
	rv.Autorelease()
	return rv
}

func (d_ DateInterval) InitWithStartDate_EndDate(startDate IDate, endDate IDate) DateInterval {
	rv := ffi.CallMethod[DateInterval](d_, "initWithStartDate:endDate:", startDate, endDate)
	rv.Autorelease()
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
	rv := ffi.CallMethod[DateIntervalFormatter](dc, "alloc")
	return rv
}

func (d_ DateIntervalFormatter) Init() DateIntervalFormatter {
	rv := ffi.CallMethod[DateIntervalFormatter](d_, "init")
	rv.Autorelease()
	return rv
}

func (dc _DateIntervalFormatterClass) New() DateIntervalFormatter {
	rv := ffi.CallMethod[DateIntervalFormatter](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDateIntervalFormatter() DateIntervalFormatter {
	return DateIntervalFormatterClass.New()
}

func (d_ DateIntervalFormatter) StringFromDate_ToDate(fromDate IDate, toDate IDate) string {
	rv := ffi.CallMethod[string](d_, "stringFromDate:toDate:", fromDate, toDate)
	return rv
}

func (d_ DateIntervalFormatter) StringFromDateInterval(dateInterval IDateInterval) string {
	rv := ffi.CallMethod[string](d_, "stringFromDateInterval:", dateInterval)
	return rv
}

func (d_ DateIntervalFormatter) DateStyle() DateIntervalFormatterStyle {
	rv := ffi.CallMethod[DateIntervalFormatterStyle](d_, "dateStyle")
	return rv
}

func (d_ DateIntervalFormatter) SetDateStyle(value DateIntervalFormatterStyle) {
	ffi.CallMethod[ffi.Void](d_, "setDateStyle:", value)
}

func (d_ DateIntervalFormatter) TimeStyle() DateIntervalFormatterStyle {
	rv := ffi.CallMethod[DateIntervalFormatterStyle](d_, "timeStyle")
	return rv
}

func (d_ DateIntervalFormatter) SetTimeStyle(value DateIntervalFormatterStyle) {
	ffi.CallMethod[ffi.Void](d_, "setTimeStyle:", value)
}

func (d_ DateIntervalFormatter) DateTemplate() string {
	rv := ffi.CallMethod[string](d_, "dateTemplate")
	return rv
}

func (d_ DateIntervalFormatter) SetDateTemplate(value string) {
	ffi.CallMethod[ffi.Void](d_, "setDateTemplate:", value)
}

func (d_ DateIntervalFormatter) Calendar() Calendar {
	rv := ffi.CallMethod[Calendar](d_, "calendar")
	return rv
}

func (d_ DateIntervalFormatter) SetCalendar(value ICalendar) {
	ffi.CallMethod[ffi.Void](d_, "setCalendar:", value)
}

func (d_ DateIntervalFormatter) Locale() Locale {
	rv := ffi.CallMethod[Locale](d_, "locale")
	return rv
}

func (d_ DateIntervalFormatter) SetLocale(value ILocale) {
	ffi.CallMethod[ffi.Void](d_, "setLocale:", value)
}

func (d_ DateIntervalFormatter) TimeZone() TimeZone {
	rv := ffi.CallMethod[TimeZone](d_, "timeZone")
	return rv
}

func (d_ DateIntervalFormatter) SetTimeZone(value ITimeZone) {
	ffi.CallMethod[ffi.Void](d_, "setTimeZone:", value)
}

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
	rv := ffi.CallMethod[Calendar](cc, "alloc")
	return rv
}

func (c_ Calendar) Init() Calendar {
	rv := ffi.CallMethod[Calendar](c_, "init")
	rv.Autorelease()
	return rv
}

func (cc _CalendarClass) New() Calendar {
	rv := ffi.CallMethod[Calendar](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCalendar() Calendar {
	return CalendarClass.New()
}

func (cc _CalendarClass) CalendarWithIdentifier(calendarIdentifierConstant CalendarIdentifier) Calendar {
	rv := ffi.CallMethod[Calendar](cc, "calendarWithIdentifier:", calendarIdentifierConstant)
	return rv
}

func (c_ Calendar) InitWithCalendarIdentifier(ident CalendarIdentifier) objc.Object {
	rv := ffi.CallMethod[objc.Object](c_, "initWithCalendarIdentifier:", ident)
	rv.Autorelease()
	return rv
}

func (c_ Calendar) Date_MatchesComponents(date IDate, components IDateComponents) bool {
	rv := ffi.CallMethod[bool](c_, "date:matchesComponents:", date, components)
	return rv
}

func (c_ Calendar) Component_FromDate(unit CalendarUnit, date IDate) int {
	rv := ffi.CallMethod[int](c_, "component:fromDate:", unit, date)
	return rv
}

func (c_ Calendar) Components_FromDate(unitFlags CalendarUnit, date IDate) DateComponents {
	rv := ffi.CallMethod[DateComponents](c_, "components:fromDate:", unitFlags, date)
	return rv
}

func (c_ Calendar) Components_FromDate_ToDate_Options(unitFlags CalendarUnit, startingDate IDate, resultDate IDate, opts CalendarOptions) DateComponents {
	rv := ffi.CallMethod[DateComponents](c_, "components:fromDate:toDate:options:", unitFlags, startingDate, resultDate, opts)
	return rv
}

func (c_ Calendar) Components_FromDateComponents_ToDateComponents_Options(unitFlags CalendarUnit, startingDateComp IDateComponents, resultDateComp IDateComponents, options CalendarOptions) DateComponents {
	rv := ffi.CallMethod[DateComponents](c_, "components:fromDateComponents:toDateComponents:options:", unitFlags, startingDateComp, resultDateComp, options)
	return rv
}

func (c_ Calendar) ComponentsInTimeZone_FromDate(timezone ITimeZone, date IDate) DateComponents {
	rv := ffi.CallMethod[DateComponents](c_, "componentsInTimeZone:fromDate:", timezone, date)
	return rv
}

func (c_ Calendar) MaximumRangeOfUnit(unit CalendarUnit) Range {
	rv := ffi.CallMethod[Range](c_, "maximumRangeOfUnit:", unit)
	return rv
}

func (c_ Calendar) MinimumRangeOfUnit(unit CalendarUnit) Range {
	rv := ffi.CallMethod[Range](c_, "minimumRangeOfUnit:", unit)
	return rv
}

func (c_ Calendar) OrdinalityOfUnit_InUnit_ForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) uint {
	rv := ffi.CallMethod[uint](c_, "ordinalityOfUnit:inUnit:forDate:", smaller, larger, date)
	return rv
}

func (c_ Calendar) RangeOfUnit_InUnit_ForDate(smaller CalendarUnit, larger CalendarUnit, date IDate) Range {
	rv := ffi.CallMethod[Range](c_, "rangeOfUnit:inUnit:forDate:", smaller, larger, date)
	return rv
}

func (c_ Calendar) RangeOfUnit_StartDate_Interval_ForDate(unit CalendarUnit, datep *Date, tip *TimeInterval, date IDate) bool {
	rv := ffi.CallMethod[bool](c_, "rangeOfUnit:startDate:interval:forDate:", unit, unsafe.Pointer(datep), tip, date)
	return rv
}

func (c_ Calendar) StartOfDayForDate(date IDate) Date {
	rv := ffi.CallMethod[Date](c_, "startOfDayForDate:", date)
	return rv
}

func (c_ Calendar) EnumerateDatesStartingAfterDate_MatchingComponents_Options_UsingBlock(start IDate, comps IDateComponents, opts CalendarOptions, block func(date Date, exactMatch bool, stop *bool)) {
	ffi.CallMethod[ffi.Void](c_, "enumerateDatesStartingAfterDate:matchingComponents:options:usingBlock:", start, comps, opts, block)
}

func (c_ Calendar) NextDateAfterDate_MatchingComponents_Options(date IDate, comps IDateComponents, options CalendarOptions) Date {
	rv := ffi.CallMethod[Date](c_, "nextDateAfterDate:matchingComponents:options:", date, comps, options)
	return rv
}

func (c_ Calendar) NextDateAfterDate_MatchingHour_Minute_Second_Options(date IDate, hourValue int, minuteValue int, secondValue int, options CalendarOptions) Date {
	rv := ffi.CallMethod[Date](c_, "nextDateAfterDate:matchingHour:minute:second:options:", date, hourValue, minuteValue, secondValue, options)
	return rv
}

func (c_ Calendar) NextDateAfterDate_MatchingUnit_Value_Options(date IDate, unit CalendarUnit, value int, options CalendarOptions) Date {
	rv := ffi.CallMethod[Date](c_, "nextDateAfterDate:matchingUnit:value:options:", date, unit, value, options)
	return rv
}

func (c_ Calendar) DateFromComponents(comps IDateComponents) Date {
	rv := ffi.CallMethod[Date](c_, "dateFromComponents:", comps)
	return rv
}

func (c_ Calendar) DateByAddingComponents_ToDate_Options(comps IDateComponents, date IDate, opts CalendarOptions) Date {
	rv := ffi.CallMethod[Date](c_, "dateByAddingComponents:toDate:options:", comps, date, opts)
	return rv
}

func (c_ Calendar) DateByAddingUnit_Value_ToDate_Options(unit CalendarUnit, value int, date IDate, options CalendarOptions) Date {
	rv := ffi.CallMethod[Date](c_, "dateByAddingUnit:value:toDate:options:", unit, value, date, options)
	return rv
}

func (c_ Calendar) DateBySettingHour_Minute_Second_OfDate_Options(h int, m int, s int, date IDate, opts CalendarOptions) Date {
	rv := ffi.CallMethod[Date](c_, "dateBySettingHour:minute:second:ofDate:options:", h, m, s, date, opts)
	return rv
}

func (c_ Calendar) DateBySettingUnit_Value_OfDate_Options(unit CalendarUnit, v int, date IDate, opts CalendarOptions) Date {
	rv := ffi.CallMethod[Date](c_, "dateBySettingUnit:value:ofDate:options:", unit, v, date, opts)
	return rv
}

func (c_ Calendar) DateWithEra_Year_Month_Day_Hour_Minute_Second_Nanosecond(eraValue int, yearValue int, monthValue int, dayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date {
	rv := ffi.CallMethod[Date](c_, "dateWithEra:year:month:day:hour:minute:second:nanosecond:", eraValue, yearValue, monthValue, dayValue, hourValue, minuteValue, secondValue, nanosecondValue)
	return rv
}

func (c_ Calendar) DateWithEra_YearForWeekOfYear_WeekOfYear_Weekday_Hour_Minute_Second_Nanosecond(eraValue int, yearValue int, weekValue int, weekdayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) Date {
	rv := ffi.CallMethod[Date](c_, "dateWithEra:yearForWeekOfYear:weekOfYear:weekday:hour:minute:second:nanosecond:", eraValue, yearValue, weekValue, weekdayValue, hourValue, minuteValue, secondValue, nanosecondValue)
	return rv
}

func (c_ Calendar) CompareDate_ToDate_ToUnitGranularity(date1 IDate, date2 IDate, unit CalendarUnit) ComparisonResult {
	rv := ffi.CallMethod[ComparisonResult](c_, "compareDate:toDate:toUnitGranularity:", date1, date2, unit)
	return rv
}

func (c_ Calendar) IsDate_EqualToDate_ToUnitGranularity(date1 IDate, date2 IDate, unit CalendarUnit) bool {
	rv := ffi.CallMethod[bool](c_, "isDate:equalToDate:toUnitGranularity:", date1, date2, unit)
	return rv
}

func (c_ Calendar) IsDate_InSameDayAsDate(date1 IDate, date2 IDate) bool {
	rv := ffi.CallMethod[bool](c_, "isDate:inSameDayAsDate:", date1, date2)
	return rv
}

func (c_ Calendar) IsDateInToday(date IDate) bool {
	rv := ffi.CallMethod[bool](c_, "isDateInToday:", date)
	return rv
}

func (c_ Calendar) IsDateInTomorrow(date IDate) bool {
	rv := ffi.CallMethod[bool](c_, "isDateInTomorrow:", date)
	return rv
}

func (c_ Calendar) IsDateInWeekend(date IDate) bool {
	rv := ffi.CallMethod[bool](c_, "isDateInWeekend:", date)
	return rv
}

func (c_ Calendar) IsDateInYesterday(date IDate) bool {
	rv := ffi.CallMethod[bool](c_, "isDateInYesterday:", date)
	return rv
}

func (cc _CalendarClass) CurrentCalendar() Calendar {
	rv := ffi.CallMethod[Calendar](cc, "currentCalendar")
	return rv
}

func (cc _CalendarClass) AutoupdatingCurrentCalendar() Calendar {
	rv := ffi.CallMethod[Calendar](cc, "autoupdatingCurrentCalendar")
	return rv
}

func (c_ Calendar) CalendarIdentifier() CalendarIdentifier {
	rv := ffi.CallMethod[CalendarIdentifier](c_, "calendarIdentifier")
	return rv
}

func (c_ Calendar) FirstWeekday() uint {
	rv := ffi.CallMethod[uint](c_, "firstWeekday")
	return rv
}

func (c_ Calendar) SetFirstWeekday(value uint) {
	ffi.CallMethod[ffi.Void](c_, "setFirstWeekday:", value)
}

func (c_ Calendar) Locale() Locale {
	rv := ffi.CallMethod[Locale](c_, "locale")
	return rv
}

func (c_ Calendar) SetLocale(value ILocale) {
	ffi.CallMethod[ffi.Void](c_, "setLocale:", value)
}

func (c_ Calendar) TimeZone() TimeZone {
	rv := ffi.CallMethod[TimeZone](c_, "timeZone")
	return rv
}

func (c_ Calendar) SetTimeZone(value ITimeZone) {
	ffi.CallMethod[ffi.Void](c_, "setTimeZone:", value)
}

func (c_ Calendar) MinimumDaysInFirstWeek() uint {
	rv := ffi.CallMethod[uint](c_, "minimumDaysInFirstWeek")
	return rv
}

func (c_ Calendar) SetMinimumDaysInFirstWeek(value uint) {
	ffi.CallMethod[ffi.Void](c_, "setMinimumDaysInFirstWeek:", value)
}

func (c_ Calendar) AMSymbol() string {
	rv := ffi.CallMethod[string](c_, "AMSymbol")
	return rv
}

func (c_ Calendar) PMSymbol() string {
	rv := ffi.CallMethod[string](c_, "PMSymbol")
	return rv
}

func (c_ Calendar) WeekdaySymbols() []string {
	rv := ffi.CallMethod[[]string](c_, "weekdaySymbols")
	return rv
}

func (c_ Calendar) ShortWeekdaySymbols() []string {
	rv := ffi.CallMethod[[]string](c_, "shortWeekdaySymbols")
	return rv
}

func (c_ Calendar) VeryShortWeekdaySymbols() []string {
	rv := ffi.CallMethod[[]string](c_, "veryShortWeekdaySymbols")
	return rv
}

func (c_ Calendar) StandaloneWeekdaySymbols() []string {
	rv := ffi.CallMethod[[]string](c_, "standaloneWeekdaySymbols")
	return rv
}

func (c_ Calendar) ShortStandaloneWeekdaySymbols() []string {
	rv := ffi.CallMethod[[]string](c_, "shortStandaloneWeekdaySymbols")
	return rv
}

func (c_ Calendar) VeryShortStandaloneWeekdaySymbols() []string {
	rv := ffi.CallMethod[[]string](c_, "veryShortStandaloneWeekdaySymbols")
	return rv
}

func (c_ Calendar) MonthSymbols() []string {
	rv := ffi.CallMethod[[]string](c_, "monthSymbols")
	return rv
}

func (c_ Calendar) ShortMonthSymbols() []string {
	rv := ffi.CallMethod[[]string](c_, "shortMonthSymbols")
	return rv
}

func (c_ Calendar) VeryShortMonthSymbols() []string {
	rv := ffi.CallMethod[[]string](c_, "veryShortMonthSymbols")
	return rv
}

func (c_ Calendar) StandaloneMonthSymbols() []string {
	rv := ffi.CallMethod[[]string](c_, "standaloneMonthSymbols")
	return rv
}

func (c_ Calendar) ShortStandaloneMonthSymbols() []string {
	rv := ffi.CallMethod[[]string](c_, "shortStandaloneMonthSymbols")
	return rv
}

func (c_ Calendar) VeryShortStandaloneMonthSymbols() []string {
	rv := ffi.CallMethod[[]string](c_, "veryShortStandaloneMonthSymbols")
	return rv
}

func (c_ Calendar) QuarterSymbols() []string {
	rv := ffi.CallMethod[[]string](c_, "quarterSymbols")
	return rv
}

func (c_ Calendar) ShortQuarterSymbols() []string {
	rv := ffi.CallMethod[[]string](c_, "shortQuarterSymbols")
	return rv
}

func (c_ Calendar) StandaloneQuarterSymbols() []string {
	rv := ffi.CallMethod[[]string](c_, "standaloneQuarterSymbols")
	return rv
}

func (c_ Calendar) ShortStandaloneQuarterSymbols() []string {
	rv := ffi.CallMethod[[]string](c_, "shortStandaloneQuarterSymbols")
	return rv
}

func (c_ Calendar) EraSymbols() []string {
	rv := ffi.CallMethod[[]string](c_, "eraSymbols")
	return rv
}

func (c_ Calendar) LongEraSymbols() []string {
	rv := ffi.CallMethod[[]string](c_, "longEraSymbols")
	return rv
}

var TimeZoneClass = _TimeZoneClass{objc.GetClass("NSTimeZone")}

type _TimeZoneClass struct {
	objc.Class
}

type ITimeZone interface {
	objc.IObject
	AbbreviationForDate(aDate IDate) string
	SecondsFromGMTForDate(aDate IDate) int
	IsDaylightSavingTimeForDate(aDate IDate) bool
	DaylightSavingTimeOffsetForDate(aDate IDate) TimeInterval
	NextDaylightSavingTimeTransitionAfterDate(aDate IDate) Date
	IsEqualToTimeZone(aTimeZone ITimeZone) bool
	LocalizedName_Locale(style TimeZoneNameStyle, locale ILocale) string
	Name() string
	Abbreviation() string
	SecondsFromGMT() int
	Data() []byte
	IsDaylightSavingTime() bool
	DaylightSavingTimeOffset() TimeInterval
	NextDaylightSavingTimeTransition() Date
	Description() string
}

type TimeZone struct {
	objc.Object
}

func MakeTimeZone(ptr unsafe.Pointer) TimeZone {
	return TimeZone{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TimeZoneClass) TimeZoneWithName(tzName string) TimeZone {
	rv := ffi.CallMethod[TimeZone](tc, "timeZoneWithName:", tzName)
	return rv
}

func (tc _TimeZoneClass) TimeZoneWithName_Data(tzName string, aData []byte) TimeZone {
	rv := ffi.CallMethod[TimeZone](tc, "timeZoneWithName:data:", tzName, aData)
	return rv
}

func (t_ TimeZone) InitWithName(tzName string) TimeZone {
	rv := ffi.CallMethod[TimeZone](t_, "initWithName:", tzName)
	rv.Autorelease()
	return rv
}

func (t_ TimeZone) InitWithName_Data(tzName string, aData []byte) TimeZone {
	rv := ffi.CallMethod[TimeZone](t_, "initWithName:data:", tzName, aData)
	rv.Autorelease()
	return rv
}

func (tc _TimeZoneClass) TimeZoneWithAbbreviation(abbreviation string) TimeZone {
	rv := ffi.CallMethod[TimeZone](tc, "timeZoneWithAbbreviation:", abbreviation)
	return rv
}

func (tc _TimeZoneClass) TimeZoneForSecondsFromGMT(seconds int) TimeZone {
	rv := ffi.CallMethod[TimeZone](tc, "timeZoneForSecondsFromGMT:", seconds)
	return rv
}

func (tc _TimeZoneClass) Alloc() TimeZone {
	rv := ffi.CallMethod[TimeZone](tc, "alloc")
	return rv
}

func (t_ TimeZone) Init() TimeZone {
	rv := ffi.CallMethod[TimeZone](t_, "init")
	rv.Autorelease()
	return rv
}

func (tc _TimeZoneClass) New() TimeZone {
	rv := ffi.CallMethod[TimeZone](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTimeZone() TimeZone {
	return TimeZoneClass.New()
}

func (tc _TimeZoneClass) ResetSystemTimeZone() {
	ffi.CallMethod[ffi.Void](tc, "resetSystemTimeZone")
}

func (t_ TimeZone) AbbreviationForDate(aDate IDate) string {
	rv := ffi.CallMethod[string](t_, "abbreviationForDate:", aDate)
	return rv
}

func (t_ TimeZone) SecondsFromGMTForDate(aDate IDate) int {
	rv := ffi.CallMethod[int](t_, "secondsFromGMTForDate:", aDate)
	return rv
}

func (t_ TimeZone) IsDaylightSavingTimeForDate(aDate IDate) bool {
	rv := ffi.CallMethod[bool](t_, "isDaylightSavingTimeForDate:", aDate)
	return rv
}

func (t_ TimeZone) DaylightSavingTimeOffsetForDate(aDate IDate) TimeInterval {
	rv := ffi.CallMethod[TimeInterval](t_, "daylightSavingTimeOffsetForDate:", aDate)
	return rv
}

func (t_ TimeZone) NextDaylightSavingTimeTransitionAfterDate(aDate IDate) Date {
	rv := ffi.CallMethod[Date](t_, "nextDaylightSavingTimeTransitionAfterDate:", aDate)
	return rv
}

func (t_ TimeZone) IsEqualToTimeZone(aTimeZone ITimeZone) bool {
	rv := ffi.CallMethod[bool](t_, "isEqualToTimeZone:", aTimeZone)
	return rv
}

func (t_ TimeZone) LocalizedName_Locale(style TimeZoneNameStyle, locale ILocale) string {
	rv := ffi.CallMethod[string](t_, "localizedName:locale:", style, locale)
	return rv
}

func (tc _TimeZoneClass) LocalTimeZone() TimeZone {
	rv := ffi.CallMethod[TimeZone](tc, "localTimeZone")
	return rv
}

func (tc _TimeZoneClass) SystemTimeZone() TimeZone {
	rv := ffi.CallMethod[TimeZone](tc, "systemTimeZone")
	return rv
}

func (tc _TimeZoneClass) DefaultTimeZone() TimeZone {
	rv := ffi.CallMethod[TimeZone](tc, "defaultTimeZone")
	return rv
}

func (tc _TimeZoneClass) SetDefaultTimeZone(value ITimeZone) {
	ffi.CallMethod[ffi.Void](tc, "setDefaultTimeZone:", value)
}

func (tc _TimeZoneClass) KnownTimeZoneNames() []string {
	rv := ffi.CallMethod[[]string](tc, "knownTimeZoneNames")
	return rv
}

func (tc _TimeZoneClass) AbbreviationDictionary() map[string]string {
	rv := ffi.CallMethod[map[string]string](tc, "abbreviationDictionary")
	return rv
}

func (tc _TimeZoneClass) SetAbbreviationDictionary(value map[string]string) {
	ffi.CallMethod[ffi.Void](tc, "setAbbreviationDictionary:", value)
}

func (t_ TimeZone) Name() string {
	rv := ffi.CallMethod[string](t_, "name")
	return rv
}

func (t_ TimeZone) Abbreviation() string {
	rv := ffi.CallMethod[string](t_, "abbreviation")
	return rv
}

func (t_ TimeZone) SecondsFromGMT() int {
	rv := ffi.CallMethod[int](t_, "secondsFromGMT")
	return rv
}

func (t_ TimeZone) Data() []byte {
	rv := ffi.CallMethod[[]byte](t_, "data")
	return rv
}

func (tc _TimeZoneClass) TimeZoneDataVersion() string {
	rv := ffi.CallMethod[string](tc, "timeZoneDataVersion")
	return rv
}

func (t_ TimeZone) IsDaylightSavingTime() bool {
	rv := ffi.CallMethod[bool](t_, "isDaylightSavingTime")
	return rv
}

func (t_ TimeZone) DaylightSavingTimeOffset() TimeInterval {
	rv := ffi.CallMethod[TimeInterval](t_, "daylightSavingTimeOffset")
	return rv
}

func (t_ TimeZone) NextDaylightSavingTimeTransition() Date {
	rv := ffi.CallMethod[Date](t_, "nextDaylightSavingTimeTransition")
	return rv
}

func (t_ TimeZone) Description() string {
	rv := ffi.CallMethod[string](t_, "description")
	return rv
}
