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

func (mc _MeasurementClass) New() Measurement {
	rv := ffi.CallMethod[Measurement](mc, "new")
	rv.Autorelease()
	return rv
}

func NewMeasurement() Measurement {
	return MeasurementClass.New()
}

func (m_ Measurement) Init() Measurement {
	rv := ffi.CallMethod[Measurement](m_, "init")
	return rv
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
