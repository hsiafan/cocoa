// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var ByteCountFormatterClass = _ByteCountFormatterClass{objc.GetClass("NSByteCountFormatter")}

type _ByteCountFormatterClass struct {
	objc.Class
}

type IByteCountFormatter interface {
	IFormatter
	StringFromByteCount(byteCount int64) string
	StringFromMeasurement(measurement IMeasurement) string
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	CountStyle() ByteCountFormatterCountStyle
	SetCountStyle(value ByteCountFormatterCountStyle)
	AllowsNonnumericFormatting() bool
	SetAllowsNonnumericFormatting(value bool)
	IncludesActualByteCount() bool
	SetIncludesActualByteCount(value bool)
	IsAdaptive() bool
	SetAdaptive(value bool)
	AllowedUnits() ByteCountFormatterUnits
	SetAllowedUnits(value ByteCountFormatterUnits)
	IncludesCount() bool
	SetIncludesCount(value bool)
	IncludesUnit() bool
	SetIncludesUnit(value bool)
	ZeroPadsFractionDigits() bool
	SetZeroPadsFractionDigits(value bool)
}

type ByteCountFormatter struct {
	Formatter
}

func MakeByteCountFormatter(ptr unsafe.Pointer) ByteCountFormatter {
	return ByteCountFormatter{
		Formatter: MakeFormatter(ptr),
	}
}

func (bc _ByteCountFormatterClass) Alloc() ByteCountFormatter {
	rv := ffi.CallMethod[ByteCountFormatter](bc, "alloc")
	return rv
}

func (bc _ByteCountFormatterClass) New() ByteCountFormatter {
	rv := ffi.CallMethod[ByteCountFormatter](bc, "new")
	rv.Autorelease()
	return rv
}

func NewByteCountFormatter() ByteCountFormatter {
	return ByteCountFormatterClass.New()
}

func (b_ ByteCountFormatter) Init() ByteCountFormatter {
	rv := ffi.CallMethod[ByteCountFormatter](b_, "init")
	return rv
}

func (bc _ByteCountFormatterClass) StringFromByteCount_CountStyle(byteCount int64, countStyle ByteCountFormatterCountStyle) string {
	rv := ffi.CallMethod[string](bc, "stringFromByteCount:countStyle:", byteCount, countStyle)
	return rv
}

func (b_ ByteCountFormatter) StringFromByteCount(byteCount int64) string {
	rv := ffi.CallMethod[string](b_, "stringFromByteCount:", byteCount)
	return rv
}

func (b_ ByteCountFormatter) StringFromMeasurement(measurement IMeasurement) string {
	rv := ffi.CallMethod[string](b_, "stringFromMeasurement:", measurement)
	return rv
}

func (bc _ByteCountFormatterClass) StringFromMeasurement_CountStyle(measurement IMeasurement, countStyle ByteCountFormatterCountStyle) string {
	rv := ffi.CallMethod[string](bc, "stringFromMeasurement:countStyle:", measurement, countStyle)
	return rv
}

func (b_ ByteCountFormatter) FormattingContext() FormattingContext {
	rv := ffi.CallMethod[FormattingContext](b_, "formattingContext")
	return rv
}

func (b_ ByteCountFormatter) SetFormattingContext(value FormattingContext) {
	ffi.CallMethod[ffi.Void](b_, "setFormattingContext:", value)
}

func (b_ ByteCountFormatter) CountStyle() ByteCountFormatterCountStyle {
	rv := ffi.CallMethod[ByteCountFormatterCountStyle](b_, "countStyle")
	return rv
}

func (b_ ByteCountFormatter) SetCountStyle(value ByteCountFormatterCountStyle) {
	ffi.CallMethod[ffi.Void](b_, "setCountStyle:", value)
}

func (b_ ByteCountFormatter) AllowsNonnumericFormatting() bool {
	rv := ffi.CallMethod[bool](b_, "allowsNonnumericFormatting")
	return rv
}

func (b_ ByteCountFormatter) SetAllowsNonnumericFormatting(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setAllowsNonnumericFormatting:", value)
}

func (b_ ByteCountFormatter) IncludesActualByteCount() bool {
	rv := ffi.CallMethod[bool](b_, "includesActualByteCount")
	return rv
}

func (b_ ByteCountFormatter) SetIncludesActualByteCount(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setIncludesActualByteCount:", value)
}

func (b_ ByteCountFormatter) IsAdaptive() bool {
	rv := ffi.CallMethod[bool](b_, "isAdaptive")
	return rv
}

func (b_ ByteCountFormatter) SetAdaptive(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setAdaptive:", value)
}

func (b_ ByteCountFormatter) AllowedUnits() ByteCountFormatterUnits {
	rv := ffi.CallMethod[ByteCountFormatterUnits](b_, "allowedUnits")
	return rv
}

func (b_ ByteCountFormatter) SetAllowedUnits(value ByteCountFormatterUnits) {
	ffi.CallMethod[ffi.Void](b_, "setAllowedUnits:", value)
}

func (b_ ByteCountFormatter) IncludesCount() bool {
	rv := ffi.CallMethod[bool](b_, "includesCount")
	return rv
}

func (b_ ByteCountFormatter) SetIncludesCount(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setIncludesCount:", value)
}

func (b_ ByteCountFormatter) IncludesUnit() bool {
	rv := ffi.CallMethod[bool](b_, "includesUnit")
	return rv
}

func (b_ ByteCountFormatter) SetIncludesUnit(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setIncludesUnit:", value)
}

func (b_ ByteCountFormatter) ZeroPadsFractionDigits() bool {
	rv := ffi.CallMethod[bool](b_, "zeroPadsFractionDigits")
	return rv
}

func (b_ ByteCountFormatter) SetZeroPadsFractionDigits(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setZeroPadsFractionDigits:", value)
}
