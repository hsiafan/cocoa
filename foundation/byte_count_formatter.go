// auto-generated code, do not modify
package foundation

import (
	"unsafe"

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
	rv := objc.CallMethod[ByteCountFormatter](bc, objc.GetSelector("alloc"))
	return rv
}

func (bc _ByteCountFormatterClass) New() ByteCountFormatter {
	rv := objc.CallMethod[ByteCountFormatter](bc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewByteCountFormatter() ByteCountFormatter {
	return ByteCountFormatterClass.New()
}

func (b_ ByteCountFormatter) Init() ByteCountFormatter {
	rv := objc.CallMethod[ByteCountFormatter](b_, objc.GetSelector("init"))
	return rv
}

func (bc _ByteCountFormatterClass) StringFromByteCount_CountStyle(byteCount int64, countStyle ByteCountFormatterCountStyle) string {
	rv := objc.CallMethod[string](bc, objc.GetSelector("stringFromByteCount:countStyle:"), byteCount, countStyle)
	return rv
}

func (b_ ByteCountFormatter) StringFromByteCount(byteCount int64) string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("stringFromByteCount:"), byteCount)
	return rv
}

func (b_ ByteCountFormatter) StringFromMeasurement(measurement IMeasurement) string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("stringFromMeasurement:"), measurement)
	return rv
}

func (bc _ByteCountFormatterClass) StringFromMeasurement_CountStyle(measurement IMeasurement, countStyle ByteCountFormatterCountStyle) string {
	rv := objc.CallMethod[string](bc, objc.GetSelector("stringFromMeasurement:countStyle:"), measurement, countStyle)
	return rv
}

func (b_ ByteCountFormatter) FormattingContext() FormattingContext {
	rv := objc.CallMethod[FormattingContext](b_, objc.GetSelector("formattingContext"))
	return rv
}

func (b_ ByteCountFormatter) SetFormattingContext(value FormattingContext) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setFormattingContext:"), value)
}

func (b_ ByteCountFormatter) CountStyle() ByteCountFormatterCountStyle {
	rv := objc.CallMethod[ByteCountFormatterCountStyle](b_, objc.GetSelector("countStyle"))
	return rv
}

func (b_ ByteCountFormatter) SetCountStyle(value ByteCountFormatterCountStyle) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setCountStyle:"), value)
}

func (b_ ByteCountFormatter) AllowsNonnumericFormatting() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("allowsNonnumericFormatting"))
	return rv
}

func (b_ ByteCountFormatter) SetAllowsNonnumericFormatting(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setAllowsNonnumericFormatting:"), value)
}

func (b_ ByteCountFormatter) IncludesActualByteCount() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("includesActualByteCount"))
	return rv
}

func (b_ ByteCountFormatter) SetIncludesActualByteCount(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setIncludesActualByteCount:"), value)
}

func (b_ ByteCountFormatter) IsAdaptive() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("isAdaptive"))
	return rv
}

func (b_ ByteCountFormatter) SetAdaptive(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setAdaptive:"), value)
}

func (b_ ByteCountFormatter) AllowedUnits() ByteCountFormatterUnits {
	rv := objc.CallMethod[ByteCountFormatterUnits](b_, objc.GetSelector("allowedUnits"))
	return rv
}

func (b_ ByteCountFormatter) SetAllowedUnits(value ByteCountFormatterUnits) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setAllowedUnits:"), value)
}

func (b_ ByteCountFormatter) IncludesCount() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("includesCount"))
	return rv
}

func (b_ ByteCountFormatter) SetIncludesCount(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setIncludesCount:"), value)
}

func (b_ ByteCountFormatter) IncludesUnit() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("includesUnit"))
	return rv
}

func (b_ ByteCountFormatter) SetIncludesUnit(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setIncludesUnit:"), value)
}

func (b_ ByteCountFormatter) ZeroPadsFractionDigits() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("zeroPadsFractionDigits"))
	return rv
}

func (b_ ByteCountFormatter) SetZeroPadsFractionDigits(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setZeroPadsFractionDigits:"), value)
}
