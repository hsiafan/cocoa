// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var FormatterClass = _FormatterClass{objc.GetClass("NSFormatter")}

type _FormatterClass struct {
	objc.Class
}

type IFormatter interface {
	objc.IObject
	StringForObjectValue(obj objc.IObject) string
	AttributedStringForObjectValue_WithDefaultAttributes(obj objc.IObject, attrs map[AttributedStringKey]objc.IObject) AttributedString
	EditingStringForObjectValue(obj objc.IObject) string
	IsPartialStringValid_NewEditingString_ErrorDescription(partialString string, newString *String, error *String) bool
	IsPartialStringValid_ProposedSelectedRange_OriginalString_OriginalSelectedRange_ErrorDescription(partialStringPtr *String, proposedSelRangePtr *Range, origString string, origSelRange Range, error *String) bool
}

type Formatter struct {
	objc.Object
}

func MakeFormatter(ptr unsafe.Pointer) Formatter {
	return Formatter{
		Object: objc.MakeObject(ptr),
	}
}

func (fc _FormatterClass) Alloc() Formatter {
	rv := ffi.CallMethod[Formatter](fc, "alloc")
	return rv
}

func (f_ Formatter) Init() Formatter {
	rv := ffi.CallMethod[Formatter](f_, "init")
	return rv
}

func (fc _FormatterClass) New() Formatter {
	rv := ffi.CallMethod[Formatter](fc, "new")
	rv.Autorelease()
	return rv
}

func NewFormatter() Formatter {
	return FormatterClass.New()
}

func (f_ Formatter) StringForObjectValue(obj objc.IObject) string {
	rv := ffi.CallMethod[string](f_, "stringForObjectValue:", obj)
	return rv
}

func (f_ Formatter) AttributedStringForObjectValue_WithDefaultAttributes(obj objc.IObject, attrs map[AttributedStringKey]objc.IObject) AttributedString {
	rv := ffi.CallMethod[AttributedString](f_, "attributedStringForObjectValue:withDefaultAttributes:", obj, attrs)
	return rv
}

func (f_ Formatter) EditingStringForObjectValue(obj objc.IObject) string {
	rv := ffi.CallMethod[string](f_, "editingStringForObjectValue:", obj)
	return rv
}

func (f_ Formatter) IsPartialStringValid_NewEditingString_ErrorDescription(partialString string, newString *String, error *String) bool {
	rv := ffi.CallMethod[bool](f_, "isPartialStringValid:newEditingString:errorDescription:", partialString, newString, error)
	return rv
}

func (f_ Formatter) IsPartialStringValid_ProposedSelectedRange_OriginalString_OriginalSelectedRange_ErrorDescription(partialStringPtr *String, proposedSelRangePtr *Range, origString string, origSelRange Range, error *String) bool {
	rv := ffi.CallMethod[bool](f_, "isPartialStringValid:proposedSelectedRange:originalString:originalSelectedRange:errorDescription:", partialStringPtr, proposedSelRangePtr, origString, origSelRange, error)
	return rv
}

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

func (b_ ByteCountFormatter) Init() ByteCountFormatter {
	rv := ffi.CallMethod[ByteCountFormatter](b_, "init")
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
