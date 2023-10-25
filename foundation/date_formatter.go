// auto-generated code, do not modify
package foundation

import (
	"unsafe"

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
	rv := objc.CallMethod[DateFormatter](dc, objc.SEL("alloc"))
	return rv
}

func (dc _DateFormatterClass) New() DateFormatter {
	rv := objc.CallMethod[DateFormatter](dc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewDateFormatter() DateFormatter {
	return DateFormatterClass.New()
}

func (d_ DateFormatter) Init() DateFormatter {
	rv := objc.CallMethod[DateFormatter](d_, objc.SEL("init"))
	return rv
}

func (d_ DateFormatter) DateFromString(string_ string) Date {
	rv := objc.CallMethod[Date](d_, objc.SEL("dateFromString:"), string_)
	return rv
}

func (d_ DateFormatter) StringFromDate(date IDate) string {
	rv := objc.CallMethod[string](d_, objc.SEL("stringFromDate:"), objc.ExtractPtr(date))
	return rv
}

func (dc _DateFormatterClass) LocalizedStringFromDate_DateStyle_TimeStyle(date IDate, dstyle DateFormatterStyle, tstyle DateFormatterStyle) string {
	rv := objc.CallMethod[string](dc, objc.SEL("localizedStringFromDate:dateStyle:timeStyle:"), objc.ExtractPtr(date), dstyle, tstyle)
	return rv
}

func (d_ DateFormatter) SetLocalizedDateFormatFromTemplate(dateFormatTemplate string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setLocalizedDateFormatFromTemplate:"), dateFormatTemplate)
}

func (dc _DateFormatterClass) DateFormatFromTemplate_Options_Locale(tmplate string, opts uint, locale ILocale) string {
	rv := objc.CallMethod[string](dc, objc.SEL("dateFormatFromTemplate:options:locale:"), tmplate, opts, objc.ExtractPtr(locale))
	return rv
}

// deprecated
func (d_ DateFormatter) AllowsNaturalLanguage() bool {
	rv := objc.CallMethod[bool](d_, objc.SEL("allowsNaturalLanguage"))
	return rv
}

// deprecated
func (d_ DateFormatter) InitWithDateFormat_AllowNaturalLanguage(format string, flag bool) objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.SEL("initWithDateFormat:allowNaturalLanguage:"), format, flag)
	return rv
}

func (d_ DateFormatter) DateStyle() DateFormatterStyle {
	rv := objc.CallMethod[DateFormatterStyle](d_, objc.SEL("dateStyle"))
	return rv
}

func (d_ DateFormatter) SetDateStyle(value DateFormatterStyle) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setDateStyle:"), value)
}

func (d_ DateFormatter) TimeStyle() DateFormatterStyle {
	rv := objc.CallMethod[DateFormatterStyle](d_, objc.SEL("timeStyle"))
	return rv
}

func (d_ DateFormatter) SetTimeStyle(value DateFormatterStyle) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setTimeStyle:"), value)
}

func (d_ DateFormatter) DateFormat() string {
	rv := objc.CallMethod[string](d_, objc.SEL("dateFormat"))
	return rv
}

func (d_ DateFormatter) SetDateFormat(value string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setDateFormat:"), value)
}

func (d_ DateFormatter) FormattingContext() FormattingContext {
	rv := objc.CallMethod[FormattingContext](d_, objc.SEL("formattingContext"))
	return rv
}

func (d_ DateFormatter) SetFormattingContext(value FormattingContext) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setFormattingContext:"), value)
}

func (d_ DateFormatter) Calendar() Calendar {
	rv := objc.CallMethod[Calendar](d_, objc.SEL("calendar"))
	return rv
}

func (d_ DateFormatter) SetCalendar(value ICalendar) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setCalendar:"), objc.ExtractPtr(value))
}

func (d_ DateFormatter) DefaultDate() Date {
	rv := objc.CallMethod[Date](d_, objc.SEL("defaultDate"))
	return rv
}

func (d_ DateFormatter) SetDefaultDate(value IDate) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setDefaultDate:"), objc.ExtractPtr(value))
}

func (d_ DateFormatter) Locale() Locale {
	rv := objc.CallMethod[Locale](d_, objc.SEL("locale"))
	return rv
}

func (d_ DateFormatter) SetLocale(value ILocale) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setLocale:"), objc.ExtractPtr(value))
}

func (d_ DateFormatter) TimeZone() TimeZone {
	rv := objc.CallMethod[TimeZone](d_, objc.SEL("timeZone"))
	return rv
}

func (d_ DateFormatter) SetTimeZone(value ITimeZone) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setTimeZone:"), objc.ExtractPtr(value))
}

func (d_ DateFormatter) TwoDigitStartDate() Date {
	rv := objc.CallMethod[Date](d_, objc.SEL("twoDigitStartDate"))
	return rv
}

func (d_ DateFormatter) SetTwoDigitStartDate(value IDate) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setTwoDigitStartDate:"), objc.ExtractPtr(value))
}

func (d_ DateFormatter) GregorianStartDate() Date {
	rv := objc.CallMethod[Date](d_, objc.SEL("gregorianStartDate"))
	return rv
}

func (d_ DateFormatter) SetGregorianStartDate(value IDate) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setGregorianStartDate:"), objc.ExtractPtr(value))
}

func (d_ DateFormatter) FormatterBehavior() DateFormatterBehavior {
	rv := objc.CallMethod[DateFormatterBehavior](d_, objc.SEL("formatterBehavior"))
	return rv
}

func (d_ DateFormatter) SetFormatterBehavior(value DateFormatterBehavior) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setFormatterBehavior:"), value)
}

func (dc _DateFormatterClass) DefaultFormatterBehavior() DateFormatterBehavior {
	rv := objc.CallMethod[DateFormatterBehavior](dc, objc.SEL("defaultFormatterBehavior"))
	return rv
}

func (dc _DateFormatterClass) SetDefaultFormatterBehavior(value DateFormatterBehavior) {
	objc.CallMethod[objc.Void](dc, objc.SEL("setDefaultFormatterBehavior:"), value)
}

func (d_ DateFormatter) IsLenient() bool {
	rv := objc.CallMethod[bool](d_, objc.SEL("isLenient"))
	return rv
}

func (d_ DateFormatter) SetLenient(value bool) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setLenient:"), value)
}

func (d_ DateFormatter) DoesRelativeDateFormatting() bool {
	rv := objc.CallMethod[bool](d_, objc.SEL("doesRelativeDateFormatting"))
	return rv
}

func (d_ DateFormatter) SetDoesRelativeDateFormatting(value bool) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setDoesRelativeDateFormatting:"), value)
}

func (d_ DateFormatter) AMSymbol() string {
	rv := objc.CallMethod[string](d_, objc.SEL("AMSymbol"))
	return rv
}

func (d_ DateFormatter) SetAMSymbol(value string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setAMSymbol:"), value)
}

func (d_ DateFormatter) PMSymbol() string {
	rv := objc.CallMethod[string](d_, objc.SEL("PMSymbol"))
	return rv
}

func (d_ DateFormatter) SetPMSymbol(value string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setPMSymbol:"), value)
}

func (d_ DateFormatter) WeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.SEL("weekdaySymbols"))
	return rv
}

func (d_ DateFormatter) SetWeekdaySymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setWeekdaySymbols:"), value)
}

func (d_ DateFormatter) ShortWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.SEL("shortWeekdaySymbols"))
	return rv
}

func (d_ DateFormatter) SetShortWeekdaySymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setShortWeekdaySymbols:"), value)
}

func (d_ DateFormatter) VeryShortWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.SEL("veryShortWeekdaySymbols"))
	return rv
}

func (d_ DateFormatter) SetVeryShortWeekdaySymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setVeryShortWeekdaySymbols:"), value)
}

func (d_ DateFormatter) StandaloneWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.SEL("standaloneWeekdaySymbols"))
	return rv
}

func (d_ DateFormatter) SetStandaloneWeekdaySymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setStandaloneWeekdaySymbols:"), value)
}

func (d_ DateFormatter) ShortStandaloneWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.SEL("shortStandaloneWeekdaySymbols"))
	return rv
}

func (d_ DateFormatter) SetShortStandaloneWeekdaySymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setShortStandaloneWeekdaySymbols:"), value)
}

func (d_ DateFormatter) VeryShortStandaloneWeekdaySymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.SEL("veryShortStandaloneWeekdaySymbols"))
	return rv
}

func (d_ DateFormatter) SetVeryShortStandaloneWeekdaySymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setVeryShortStandaloneWeekdaySymbols:"), value)
}

func (d_ DateFormatter) MonthSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.SEL("monthSymbols"))
	return rv
}

func (d_ DateFormatter) SetMonthSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setMonthSymbols:"), value)
}

func (d_ DateFormatter) ShortMonthSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.SEL("shortMonthSymbols"))
	return rv
}

func (d_ DateFormatter) SetShortMonthSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setShortMonthSymbols:"), value)
}

func (d_ DateFormatter) VeryShortMonthSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.SEL("veryShortMonthSymbols"))
	return rv
}

func (d_ DateFormatter) SetVeryShortMonthSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setVeryShortMonthSymbols:"), value)
}

func (d_ DateFormatter) StandaloneMonthSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.SEL("standaloneMonthSymbols"))
	return rv
}

func (d_ DateFormatter) SetStandaloneMonthSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setStandaloneMonthSymbols:"), value)
}

func (d_ DateFormatter) ShortStandaloneMonthSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.SEL("shortStandaloneMonthSymbols"))
	return rv
}

func (d_ DateFormatter) SetShortStandaloneMonthSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setShortStandaloneMonthSymbols:"), value)
}

func (d_ DateFormatter) VeryShortStandaloneMonthSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.SEL("veryShortStandaloneMonthSymbols"))
	return rv
}

func (d_ DateFormatter) SetVeryShortStandaloneMonthSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setVeryShortStandaloneMonthSymbols:"), value)
}

func (d_ DateFormatter) QuarterSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.SEL("quarterSymbols"))
	return rv
}

func (d_ DateFormatter) SetQuarterSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setQuarterSymbols:"), value)
}

func (d_ DateFormatter) ShortQuarterSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.SEL("shortQuarterSymbols"))
	return rv
}

func (d_ DateFormatter) SetShortQuarterSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setShortQuarterSymbols:"), value)
}

func (d_ DateFormatter) StandaloneQuarterSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.SEL("standaloneQuarterSymbols"))
	return rv
}

func (d_ DateFormatter) SetStandaloneQuarterSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setStandaloneQuarterSymbols:"), value)
}

func (d_ DateFormatter) ShortStandaloneQuarterSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.SEL("shortStandaloneQuarterSymbols"))
	return rv
}

func (d_ DateFormatter) SetShortStandaloneQuarterSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setShortStandaloneQuarterSymbols:"), value)
}

func (d_ DateFormatter) EraSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.SEL("eraSymbols"))
	return rv
}

func (d_ DateFormatter) SetEraSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setEraSymbols:"), value)
}

func (d_ DateFormatter) LongEraSymbols() []string {
	rv := objc.CallMethod[[]string](d_, objc.SEL("longEraSymbols"))
	return rv
}

func (d_ DateFormatter) SetLongEraSymbols(value []string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setLongEraSymbols:"), value)
}

func (d_ DateFormatter) GeneratesCalendarDates() bool {
	rv := objc.CallMethod[bool](d_, objc.SEL("generatesCalendarDates"))
	return rv
}

func (d_ DateFormatter) SetGeneratesCalendarDates(value bool) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setGeneratesCalendarDates:"), value)
}
