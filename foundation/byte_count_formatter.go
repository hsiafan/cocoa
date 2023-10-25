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
	rv := objc.CallMethod[ByteCountFormatter](bc, objc.SEL("alloc"))
	return rv
}

func (bc _ByteCountFormatterClass) New() ByteCountFormatter {
	rv := objc.CallMethod[ByteCountFormatter](bc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewByteCountFormatter() ByteCountFormatter {
	return ByteCountFormatterClass.New()
}

func (b_ ByteCountFormatter) Init() ByteCountFormatter {
	rv := objc.CallMethod[ByteCountFormatter](b_, objc.SEL("init"))
	return rv
}

func (bc _ByteCountFormatterClass) StringFromByteCount_CountStyle(byteCount int64, countStyle ByteCountFormatterCountStyle) string {
	rv := objc.CallMethod[string](bc, objc.SEL("stringFromByteCount:countStyle:"), byteCount, countStyle)
	return rv
}

func (b_ ByteCountFormatter) StringFromByteCount(byteCount int64) string {
	rv := objc.CallMethod[string](b_, objc.SEL("stringFromByteCount:"), byteCount)
	return rv
}

func (b_ ByteCountFormatter) StringFromMeasurement(measurement IMeasurement) string {
	rv := objc.CallMethod[string](b_, objc.SEL("stringFromMeasurement:"), objc.ExtractPtr(measurement))
	return rv
}

func (bc _ByteCountFormatterClass) StringFromMeasurement_CountStyle(measurement IMeasurement, countStyle ByteCountFormatterCountStyle) string {
	rv := objc.CallMethod[string](bc, objc.SEL("stringFromMeasurement:countStyle:"), objc.ExtractPtr(measurement), countStyle)
	return rv
}

func (b_ ByteCountFormatter) FormattingContext() FormattingContext {
	rv := objc.CallMethod[FormattingContext](b_, objc.SEL("formattingContext"))
	return rv
}

func (b_ ByteCountFormatter) SetFormattingContext(value FormattingContext) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setFormattingContext:"), value)
}

func (b_ ByteCountFormatter) CountStyle() ByteCountFormatterCountStyle {
	rv := objc.CallMethod[ByteCountFormatterCountStyle](b_, objc.SEL("countStyle"))
	return rv
}

func (b_ ByteCountFormatter) SetCountStyle(value ByteCountFormatterCountStyle) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setCountStyle:"), value)
}

func (b_ ByteCountFormatter) AllowsNonnumericFormatting() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("allowsNonnumericFormatting"))
	return rv
}

func (b_ ByteCountFormatter) SetAllowsNonnumericFormatting(value bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setAllowsNonnumericFormatting:"), value)
}

func (b_ ByteCountFormatter) IncludesActualByteCount() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("includesActualByteCount"))
	return rv
}

func (b_ ByteCountFormatter) SetIncludesActualByteCount(value bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setIncludesActualByteCount:"), value)
}

func (b_ ByteCountFormatter) IsAdaptive() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("isAdaptive"))
	return rv
}

func (b_ ByteCountFormatter) SetAdaptive(value bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setAdaptive:"), value)
}

func (b_ ByteCountFormatter) AllowedUnits() ByteCountFormatterUnits {
	rv := objc.CallMethod[ByteCountFormatterUnits](b_, objc.SEL("allowedUnits"))
	return rv
}

func (b_ ByteCountFormatter) SetAllowedUnits(value ByteCountFormatterUnits) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setAllowedUnits:"), value)
}

func (b_ ByteCountFormatter) IncludesCount() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("includesCount"))
	return rv
}

func (b_ ByteCountFormatter) SetIncludesCount(value bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setIncludesCount:"), value)
}

func (b_ ByteCountFormatter) IncludesUnit() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("includesUnit"))
	return rv
}

func (b_ ByteCountFormatter) SetIncludesUnit(value bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setIncludesUnit:"), value)
}

func (b_ ByteCountFormatter) ZeroPadsFractionDigits() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("zeroPadsFractionDigits"))
	return rv
}

func (b_ ByteCountFormatter) SetZeroPadsFractionDigits(value bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setZeroPadsFractionDigits:"), value)
}
