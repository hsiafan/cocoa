// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var DateFormatterClass = _DateFormatterClass{objc.GetClass("NSDateFormatter")}

type _DateFormatterClass struct {
	objc.Class
}

type IDateFormatter interface {
	IFormatter
	DateFromString(string_ string) Date
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

func (dc _DateFormatterClass) New() DateFormatter {
	rv := ffi.CallMethod[DateFormatter](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDateFormatter() DateFormatter {
	return DateFormatterClass.New()
}

func (d_ DateFormatter) Init() DateFormatter {
	rv := ffi.CallMethod[DateFormatter](d_, "init")
	return rv
}

func (d_ DateFormatter) DateFromString(string_ string) Date {
	rv := ffi.CallMethod[Date](d_, "dateFromString:", string_)
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
