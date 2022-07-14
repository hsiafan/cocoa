// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var MeasurementClass = _MeasurementClass{objc.GetClass("NSMeasurement")}

type _MeasurementClass struct {
	objc.Class
}

type IMeasurement interface {
	objc.IObject
	CanBeConvertedToUnit(unit IUnit) bool
	MeasurementByConvertingToUnit(unit IUnit) Measurement
	MeasurementByAddingMeasurement(measurement IMeasurement) Measurement
	MeasurementBySubtractingMeasurement(measurement IMeasurement) Measurement
	DoubleValue() float64
}

type Measurement struct {
	objc.Object
}

func MakeMeasurement(ptr unsafe.Pointer) Measurement {
	return Measurement{
		Object: objc.MakeObject(ptr),
	}
}

func (mc _MeasurementClass) Alloc() Measurement {
	rv := ffi.CallMethod[Measurement](mc, "alloc")
	return rv
}

func (m_ Measurement) Init() Measurement {
	rv := ffi.CallMethod[Measurement](m_, "init")
	rv.Autorelease()
	return rv
}

func (mc _MeasurementClass) New() Measurement {
	rv := ffi.CallMethod[Measurement](mc, "new")
	rv.Autorelease()
	return rv
}

func NewMeasurement() Measurement {
	return MeasurementClass.New()
}

func (m_ Measurement) CanBeConvertedToUnit(unit IUnit) bool {
	rv := ffi.CallMethod[bool](m_, "canBeConvertedToUnit:", unit)
	return rv
}

func (m_ Measurement) MeasurementByConvertingToUnit(unit IUnit) Measurement {
	rv := ffi.CallMethod[Measurement](m_, "measurementByConvertingToUnit:", unit)
	return rv
}

func (m_ Measurement) MeasurementByAddingMeasurement(measurement IMeasurement) Measurement {
	rv := ffi.CallMethod[Measurement](m_, "measurementByAddingMeasurement:", measurement)
	return rv
}

func (m_ Measurement) MeasurementBySubtractingMeasurement(measurement IMeasurement) Measurement {
	rv := ffi.CallMethod[Measurement](m_, "measurementBySubtractingMeasurement:", measurement)
	return rv
}

func (m_ Measurement) DoubleValue() float64 {
	rv := ffi.CallMethod[float64](m_, "doubleValue")
	return rv
}

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
	rv := ffi.CallMethod[MeasurementFormatter](mc, "alloc")
	return rv
}

func (m_ MeasurementFormatter) Init() MeasurementFormatter {
	rv := ffi.CallMethod[MeasurementFormatter](m_, "init")
	rv.Autorelease()
	return rv
}

func (mc _MeasurementFormatterClass) New() MeasurementFormatter {
	rv := ffi.CallMethod[MeasurementFormatter](mc, "new")
	rv.Autorelease()
	return rv
}

func NewMeasurementFormatter() MeasurementFormatter {
	return MeasurementFormatterClass.New()
}

func (m_ MeasurementFormatter) StringFromMeasurement(measurement IMeasurement) string {
	rv := ffi.CallMethod[string](m_, "stringFromMeasurement:", measurement)
	return rv
}

func (m_ MeasurementFormatter) StringFromUnit(unit IUnit) string {
	rv := ffi.CallMethod[string](m_, "stringFromUnit:", unit)
	return rv
}

func (m_ MeasurementFormatter) UnitOptions() MeasurementFormatterUnitOptions {
	rv := ffi.CallMethod[MeasurementFormatterUnitOptions](m_, "unitOptions")
	return rv
}

func (m_ MeasurementFormatter) SetUnitOptions(value MeasurementFormatterUnitOptions) {
	ffi.CallMethod[ffi.Void](m_, "setUnitOptions:", value)
}

func (m_ MeasurementFormatter) UnitStyle() FormattingUnitStyle {
	rv := ffi.CallMethod[FormattingUnitStyle](m_, "unitStyle")
	return rv
}

func (m_ MeasurementFormatter) SetUnitStyle(value FormattingUnitStyle) {
	ffi.CallMethod[ffi.Void](m_, "setUnitStyle:", value)
}

func (m_ MeasurementFormatter) Locale() Locale {
	rv := ffi.CallMethod[Locale](m_, "locale")
	return rv
}

func (m_ MeasurementFormatter) SetLocale(value ILocale) {
	ffi.CallMethod[ffi.Void](m_, "setLocale:", value)
}

func (m_ MeasurementFormatter) NumberFormatter() NumberFormatter {
	rv := ffi.CallMethod[NumberFormatter](m_, "numberFormatter")
	return rv
}

func (m_ MeasurementFormatter) SetNumberFormatter(value INumberFormatter) {
	ffi.CallMethod[ffi.Void](m_, "setNumberFormatter:", value)
}
