// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var MeasurementFormatterClass = _MeasurementFormatterClass{objc.GetClass("NSMeasurementFormatter")}

type _MeasurementFormatterClass struct {
	objc.Class
}

type IMeasurementFormatter interface {
	IFormatter
	StringFromMeasurement(measurement IMeasurement) string
	StringFromUnit(unit IUnit) string
	UnitOptions() MeasurementFormatterUnitOptions
	SetUnitOptions(value MeasurementFormatterUnitOptions)
	UnitStyle() FormattingUnitStyle
	SetUnitStyle(value FormattingUnitStyle)
	Locale() Locale
	SetLocale(value ILocale)
	NumberFormatter() NumberFormatter
	SetNumberFormatter(value INumberFormatter)
}

type MeasurementFormatter struct {
	Formatter
}

func MakeMeasurementFormatter(ptr unsafe.Pointer) MeasurementFormatter {
	return MeasurementFormatter{
		Formatter: MakeFormatter(ptr),
	}
}

func (mc _MeasurementFormatterClass) Alloc() MeasurementFormatter {
	rv := objc.CallMethod[MeasurementFormatter](mc, objc.SEL("alloc"))
	return rv
}

func (mc _MeasurementFormatterClass) New() MeasurementFormatter {
	rv := objc.CallMethod[MeasurementFormatter](mc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewMeasurementFormatter() MeasurementFormatter {
	return MeasurementFormatterClass.New()
}

func (m_ MeasurementFormatter) Init() MeasurementFormatter {
	rv := objc.CallMethod[MeasurementFormatter](m_, objc.SEL("init"))
	return rv
}

func (m_ MeasurementFormatter) StringFromMeasurement(measurement IMeasurement) string {
	rv := objc.CallMethod[string](m_, objc.SEL("stringFromMeasurement:"), objc.ExtractPtr(measurement))
	return rv
}

func (m_ MeasurementFormatter) StringFromUnit(unit IUnit) string {
	rv := objc.CallMethod[string](m_, objc.SEL("stringFromUnit:"), objc.ExtractPtr(unit))
	return rv
}

func (m_ MeasurementFormatter) UnitOptions() MeasurementFormatterUnitOptions {
	rv := objc.CallMethod[MeasurementFormatterUnitOptions](m_, objc.SEL("unitOptions"))
	return rv
}

func (m_ MeasurementFormatter) SetUnitOptions(value MeasurementFormatterUnitOptions) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setUnitOptions:"), value)
}

func (m_ MeasurementFormatter) UnitStyle() FormattingUnitStyle {
	rv := objc.CallMethod[FormattingUnitStyle](m_, objc.SEL("unitStyle"))
	return rv
}

func (m_ MeasurementFormatter) SetUnitStyle(value FormattingUnitStyle) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setUnitStyle:"), value)
}

func (m_ MeasurementFormatter) Locale() Locale {
	rv := objc.CallMethod[Locale](m_, objc.SEL("locale"))
	return rv
}

func (m_ MeasurementFormatter) SetLocale(value ILocale) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setLocale:"), objc.ExtractPtr(value))
}

func (m_ MeasurementFormatter) NumberFormatter() NumberFormatter {
	rv := objc.CallMethod[NumberFormatter](m_, objc.SEL("numberFormatter"))
	return rv
}

func (m_ MeasurementFormatter) SetNumberFormatter(value INumberFormatter) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setNumberFormatter:"), objc.ExtractPtr(value))
}
