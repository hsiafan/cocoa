// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var NumberFormatterClass = _NumberFormatterClass{objc.GetClass("NSNumberFormatter")}

type _NumberFormatterClass struct {
	objc.Class
}

type INumberFormatter interface {
	IFormatter
	NumberFromString(string_ string) Number
	StringFromNumber(number INumber) string
	FormatterBehavior() NumberFormatterBehavior
	SetFormatterBehavior(value NumberFormatterBehavior)
	NumberStyle() NumberFormatterStyle
	SetNumberStyle(value NumberFormatterStyle)
	GeneratesDecimalNumbers() bool
	SetGeneratesDecimalNumbers(value bool)
	LocalizesFormat() bool
	SetLocalizesFormat(value bool)
	Locale() Locale
	SetLocale(value ILocale)
	RoundingBehavior() DecimalNumberHandler
	SetRoundingBehavior(value IDecimalNumberHandler)
	RoundingIncrement() Number
	SetRoundingIncrement(value INumber)
	RoundingMode() NumberFormatterRoundingMode
	SetRoundingMode(value NumberFormatterRoundingMode)
	MinimumIntegerDigits() uint
	SetMinimumIntegerDigits(value uint)
	MaximumIntegerDigits() uint
	SetMaximumIntegerDigits(value uint)
	MinimumFractionDigits() uint
	SetMinimumFractionDigits(value uint)
	MaximumFractionDigits() uint
	SetMaximumFractionDigits(value uint)
	UsesSignificantDigits() bool
	SetUsesSignificantDigits(value bool)
	MinimumSignificantDigits() uint
	SetMinimumSignificantDigits(value uint)
	MaximumSignificantDigits() uint
	SetMaximumSignificantDigits(value uint)
	Format() string
	SetFormat(value string)
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	FormatWidth() uint
	SetFormatWidth(value uint)
	NegativeFormat() string
	SetNegativeFormat(value string)
	PositiveFormat() string
	SetPositiveFormat(value string)
	Multiplier() Number
	SetMultiplier(value INumber)
	PercentSymbol() string
	SetPercentSymbol(value string)
	PerMillSymbol() string
	SetPerMillSymbol(value string)
	MinusSign() string
	SetMinusSign(value string)
	PlusSign() string
	SetPlusSign(value string)
	ExponentSymbol() string
	SetExponentSymbol(value string)
	ZeroSymbol() string
	SetZeroSymbol(value string)
	NilSymbol() string
	SetNilSymbol(value string)
	NotANumberSymbol() string
	SetNotANumberSymbol(value string)
	NegativeInfinitySymbol() string
	SetNegativeInfinitySymbol(value string)
	PositiveInfinitySymbol() string
	SetPositiveInfinitySymbol(value string)
	CurrencySymbol() string
	SetCurrencySymbol(value string)
	CurrencyCode() string
	SetCurrencyCode(value string)
	InternationalCurrencySymbol() string
	SetInternationalCurrencySymbol(value string)
	CurrencyGroupingSeparator() string
	SetCurrencyGroupingSeparator(value string)
	PositivePrefix() string
	SetPositivePrefix(value string)
	PositiveSuffix() string
	SetPositiveSuffix(value string)
	NegativePrefix() string
	SetNegativePrefix(value string)
	NegativeSuffix() string
	SetNegativeSuffix(value string)
	TextAttributesForNegativeValues() map[string]objc.Object
	SetTextAttributesForNegativeValues(value map[string]objc.IObject)
	TextAttributesForPositiveValues() map[string]objc.Object
	SetTextAttributesForPositiveValues(value map[string]objc.IObject)
	AttributedStringForZero() AttributedString
	SetAttributedStringForZero(value IAttributedString)
	TextAttributesForZero() map[string]objc.Object
	SetTextAttributesForZero(value map[string]objc.IObject)
	AttributedStringForNil() AttributedString
	SetAttributedStringForNil(value IAttributedString)
	TextAttributesForNil() map[string]objc.Object
	SetTextAttributesForNil(value map[string]objc.IObject)
	AttributedStringForNotANumber() AttributedString
	SetAttributedStringForNotANumber(value IAttributedString)
	TextAttributesForNotANumber() map[string]objc.Object
	SetTextAttributesForNotANumber(value map[string]objc.IObject)
	TextAttributesForPositiveInfinity() map[string]objc.Object
	SetTextAttributesForPositiveInfinity(value map[string]objc.IObject)
	TextAttributesForNegativeInfinity() map[string]objc.Object
	SetTextAttributesForNegativeInfinity(value map[string]objc.IObject)
	GroupingSeparator() string
	SetGroupingSeparator(value string)
	UsesGroupingSeparator() bool
	SetUsesGroupingSeparator(value bool)
	ThousandSeparator() string
	SetThousandSeparator(value string)
	HasThousandSeparators() bool
	SetHasThousandSeparators(value bool)
	DecimalSeparator() string
	SetDecimalSeparator(value string)
	AlwaysShowsDecimalSeparator() bool
	SetAlwaysShowsDecimalSeparator(value bool)
	CurrencyDecimalSeparator() string
	SetCurrencyDecimalSeparator(value string)
	GroupingSize() uint
	SetGroupingSize(value uint)
	SecondaryGroupingSize() uint
	SetSecondaryGroupingSize(value uint)
	PaddingCharacter() string
	SetPaddingCharacter(value string)
	PaddingPosition() NumberFormatterPadPosition
	SetPaddingPosition(value NumberFormatterPadPosition)
	AllowsFloats() bool
	SetAllowsFloats(value bool)
	Minimum() Number
	SetMinimum(value INumber)
	Maximum() Number
	SetMaximum(value INumber)
	IsLenient() bool
	SetLenient(value bool)
	IsPartialStringValidationEnabled() bool
	SetPartialStringValidationEnabled(value bool)
}

type NumberFormatter struct {
	Formatter
}

func MakeNumberFormatter(ptr unsafe.Pointer) NumberFormatter {
	return NumberFormatter{
		Formatter: MakeFormatter(ptr),
	}
}

func (nc _NumberFormatterClass) Alloc() NumberFormatter {
	rv := ffi.CallMethod[NumberFormatter](nc, "alloc")
	return rv
}

func (nc _NumberFormatterClass) New() NumberFormatter {
	rv := ffi.CallMethod[NumberFormatter](nc, "new")
	rv.Autorelease()
	return rv
}

func NewNumberFormatter() NumberFormatter {
	return NumberFormatterClass.New()
}

func (n_ NumberFormatter) Init() NumberFormatter {
	rv := ffi.CallMethod[NumberFormatter](n_, "init")
	return rv
}

func (nc _NumberFormatterClass) SetDefaultFormatterBehavior(behavior NumberFormatterBehavior) {
	ffi.CallMethod[ffi.Void](nc, "setDefaultFormatterBehavior:", behavior)
}

func (nc _NumberFormatterClass) DefaultFormatterBehavior() NumberFormatterBehavior {
	rv := ffi.CallMethod[NumberFormatterBehavior](nc, "defaultFormatterBehavior")
	return rv
}

func (n_ NumberFormatter) NumberFromString(string_ string) Number {
	rv := ffi.CallMethod[Number](n_, "numberFromString:", string_)
	return rv
}

func (n_ NumberFormatter) StringFromNumber(number INumber) string {
	rv := ffi.CallMethod[string](n_, "stringFromNumber:", number)
	return rv
}

func (nc _NumberFormatterClass) LocalizedStringFromNumber_NumberStyle(num INumber, nstyle NumberFormatterStyle) string {
	rv := ffi.CallMethod[string](nc, "localizedStringFromNumber:numberStyle:", num, nstyle)
	return rv
}

func (n_ NumberFormatter) FormatterBehavior() NumberFormatterBehavior {
	rv := ffi.CallMethod[NumberFormatterBehavior](n_, "formatterBehavior")
	return rv
}

func (n_ NumberFormatter) SetFormatterBehavior(value NumberFormatterBehavior) {
	ffi.CallMethod[ffi.Void](n_, "setFormatterBehavior:", value)
}

func (n_ NumberFormatter) NumberStyle() NumberFormatterStyle {
	rv := ffi.CallMethod[NumberFormatterStyle](n_, "numberStyle")
	return rv
}

func (n_ NumberFormatter) SetNumberStyle(value NumberFormatterStyle) {
	ffi.CallMethod[ffi.Void](n_, "setNumberStyle:", value)
}

func (n_ NumberFormatter) GeneratesDecimalNumbers() bool {
	rv := ffi.CallMethod[bool](n_, "generatesDecimalNumbers")
	return rv
}

func (n_ NumberFormatter) SetGeneratesDecimalNumbers(value bool) {
	ffi.CallMethod[ffi.Void](n_, "setGeneratesDecimalNumbers:", value)
}

func (n_ NumberFormatter) LocalizesFormat() bool {
	rv := ffi.CallMethod[bool](n_, "localizesFormat")
	return rv
}

func (n_ NumberFormatter) SetLocalizesFormat(value bool) {
	ffi.CallMethod[ffi.Void](n_, "setLocalizesFormat:", value)
}

func (n_ NumberFormatter) Locale() Locale {
	rv := ffi.CallMethod[Locale](n_, "locale")
	return rv
}

func (n_ NumberFormatter) SetLocale(value ILocale) {
	ffi.CallMethod[ffi.Void](n_, "setLocale:", value)
}

func (n_ NumberFormatter) RoundingBehavior() DecimalNumberHandler {
	rv := ffi.CallMethod[DecimalNumberHandler](n_, "roundingBehavior")
	return rv
}

func (n_ NumberFormatter) SetRoundingBehavior(value IDecimalNumberHandler) {
	ffi.CallMethod[ffi.Void](n_, "setRoundingBehavior:", value)
}

func (n_ NumberFormatter) RoundingIncrement() Number {
	rv := ffi.CallMethod[Number](n_, "roundingIncrement")
	return rv
}

func (n_ NumberFormatter) SetRoundingIncrement(value INumber) {
	ffi.CallMethod[ffi.Void](n_, "setRoundingIncrement:", value)
}

func (n_ NumberFormatter) RoundingMode() NumberFormatterRoundingMode {
	rv := ffi.CallMethod[NumberFormatterRoundingMode](n_, "roundingMode")
	return rv
}

func (n_ NumberFormatter) SetRoundingMode(value NumberFormatterRoundingMode) {
	ffi.CallMethod[ffi.Void](n_, "setRoundingMode:", value)
}

func (n_ NumberFormatter) MinimumIntegerDigits() uint {
	rv := ffi.CallMethod[uint](n_, "minimumIntegerDigits")
	return rv
}

func (n_ NumberFormatter) SetMinimumIntegerDigits(value uint) {
	ffi.CallMethod[ffi.Void](n_, "setMinimumIntegerDigits:", value)
}

func (n_ NumberFormatter) MaximumIntegerDigits() uint {
	rv := ffi.CallMethod[uint](n_, "maximumIntegerDigits")
	return rv
}

func (n_ NumberFormatter) SetMaximumIntegerDigits(value uint) {
	ffi.CallMethod[ffi.Void](n_, "setMaximumIntegerDigits:", value)
}

func (n_ NumberFormatter) MinimumFractionDigits() uint {
	rv := ffi.CallMethod[uint](n_, "minimumFractionDigits")
	return rv
}

func (n_ NumberFormatter) SetMinimumFractionDigits(value uint) {
	ffi.CallMethod[ffi.Void](n_, "setMinimumFractionDigits:", value)
}

func (n_ NumberFormatter) MaximumFractionDigits() uint {
	rv := ffi.CallMethod[uint](n_, "maximumFractionDigits")
	return rv
}

func (n_ NumberFormatter) SetMaximumFractionDigits(value uint) {
	ffi.CallMethod[ffi.Void](n_, "setMaximumFractionDigits:", value)
}

func (n_ NumberFormatter) UsesSignificantDigits() bool {
	rv := ffi.CallMethod[bool](n_, "usesSignificantDigits")
	return rv
}

func (n_ NumberFormatter) SetUsesSignificantDigits(value bool) {
	ffi.CallMethod[ffi.Void](n_, "setUsesSignificantDigits:", value)
}

func (n_ NumberFormatter) MinimumSignificantDigits() uint {
	rv := ffi.CallMethod[uint](n_, "minimumSignificantDigits")
	return rv
}

func (n_ NumberFormatter) SetMinimumSignificantDigits(value uint) {
	ffi.CallMethod[ffi.Void](n_, "setMinimumSignificantDigits:", value)
}

func (n_ NumberFormatter) MaximumSignificantDigits() uint {
	rv := ffi.CallMethod[uint](n_, "maximumSignificantDigits")
	return rv
}

func (n_ NumberFormatter) SetMaximumSignificantDigits(value uint) {
	ffi.CallMethod[ffi.Void](n_, "setMaximumSignificantDigits:", value)
}

func (n_ NumberFormatter) Format() string {
	rv := ffi.CallMethod[string](n_, "format")
	return rv
}

func (n_ NumberFormatter) SetFormat(value string) {
	ffi.CallMethod[ffi.Void](n_, "setFormat:", value)
}

func (n_ NumberFormatter) FormattingContext() FormattingContext {
	rv := ffi.CallMethod[FormattingContext](n_, "formattingContext")
	return rv
}

func (n_ NumberFormatter) SetFormattingContext(value FormattingContext) {
	ffi.CallMethod[ffi.Void](n_, "setFormattingContext:", value)
}

func (n_ NumberFormatter) FormatWidth() uint {
	rv := ffi.CallMethod[uint](n_, "formatWidth")
	return rv
}

func (n_ NumberFormatter) SetFormatWidth(value uint) {
	ffi.CallMethod[ffi.Void](n_, "setFormatWidth:", value)
}

func (n_ NumberFormatter) NegativeFormat() string {
	rv := ffi.CallMethod[string](n_, "negativeFormat")
	return rv
}

func (n_ NumberFormatter) SetNegativeFormat(value string) {
	ffi.CallMethod[ffi.Void](n_, "setNegativeFormat:", value)
}

func (n_ NumberFormatter) PositiveFormat() string {
	rv := ffi.CallMethod[string](n_, "positiveFormat")
	return rv
}

func (n_ NumberFormatter) SetPositiveFormat(value string) {
	ffi.CallMethod[ffi.Void](n_, "setPositiveFormat:", value)
}

func (n_ NumberFormatter) Multiplier() Number {
	rv := ffi.CallMethod[Number](n_, "multiplier")
	return rv
}

func (n_ NumberFormatter) SetMultiplier(value INumber) {
	ffi.CallMethod[ffi.Void](n_, "setMultiplier:", value)
}

func (n_ NumberFormatter) PercentSymbol() string {
	rv := ffi.CallMethod[string](n_, "percentSymbol")
	return rv
}

func (n_ NumberFormatter) SetPercentSymbol(value string) {
	ffi.CallMethod[ffi.Void](n_, "setPercentSymbol:", value)
}

func (n_ NumberFormatter) PerMillSymbol() string {
	rv := ffi.CallMethod[string](n_, "perMillSymbol")
	return rv
}

func (n_ NumberFormatter) SetPerMillSymbol(value string) {
	ffi.CallMethod[ffi.Void](n_, "setPerMillSymbol:", value)
}

func (n_ NumberFormatter) MinusSign() string {
	rv := ffi.CallMethod[string](n_, "minusSign")
	return rv
}

func (n_ NumberFormatter) SetMinusSign(value string) {
	ffi.CallMethod[ffi.Void](n_, "setMinusSign:", value)
}

func (n_ NumberFormatter) PlusSign() string {
	rv := ffi.CallMethod[string](n_, "plusSign")
	return rv
}

func (n_ NumberFormatter) SetPlusSign(value string) {
	ffi.CallMethod[ffi.Void](n_, "setPlusSign:", value)
}

func (n_ NumberFormatter) ExponentSymbol() string {
	rv := ffi.CallMethod[string](n_, "exponentSymbol")
	return rv
}

func (n_ NumberFormatter) SetExponentSymbol(value string) {
	ffi.CallMethod[ffi.Void](n_, "setExponentSymbol:", value)
}

func (n_ NumberFormatter) ZeroSymbol() string {
	rv := ffi.CallMethod[string](n_, "zeroSymbol")
	return rv
}

func (n_ NumberFormatter) SetZeroSymbol(value string) {
	ffi.CallMethod[ffi.Void](n_, "setZeroSymbol:", value)
}

func (n_ NumberFormatter) NilSymbol() string {
	rv := ffi.CallMethod[string](n_, "nilSymbol")
	return rv
}

func (n_ NumberFormatter) SetNilSymbol(value string) {
	ffi.CallMethod[ffi.Void](n_, "setNilSymbol:", value)
}

func (n_ NumberFormatter) NotANumberSymbol() string {
	rv := ffi.CallMethod[string](n_, "notANumberSymbol")
	return rv
}

func (n_ NumberFormatter) SetNotANumberSymbol(value string) {
	ffi.CallMethod[ffi.Void](n_, "setNotANumberSymbol:", value)
}

func (n_ NumberFormatter) NegativeInfinitySymbol() string {
	rv := ffi.CallMethod[string](n_, "negativeInfinitySymbol")
	return rv
}

func (n_ NumberFormatter) SetNegativeInfinitySymbol(value string) {
	ffi.CallMethod[ffi.Void](n_, "setNegativeInfinitySymbol:", value)
}

func (n_ NumberFormatter) PositiveInfinitySymbol() string {
	rv := ffi.CallMethod[string](n_, "positiveInfinitySymbol")
	return rv
}

func (n_ NumberFormatter) SetPositiveInfinitySymbol(value string) {
	ffi.CallMethod[ffi.Void](n_, "setPositiveInfinitySymbol:", value)
}

func (n_ NumberFormatter) CurrencySymbol() string {
	rv := ffi.CallMethod[string](n_, "currencySymbol")
	return rv
}

func (n_ NumberFormatter) SetCurrencySymbol(value string) {
	ffi.CallMethod[ffi.Void](n_, "setCurrencySymbol:", value)
}

func (n_ NumberFormatter) CurrencyCode() string {
	rv := ffi.CallMethod[string](n_, "currencyCode")
	return rv
}

func (n_ NumberFormatter) SetCurrencyCode(value string) {
	ffi.CallMethod[ffi.Void](n_, "setCurrencyCode:", value)
}

func (n_ NumberFormatter) InternationalCurrencySymbol() string {
	rv := ffi.CallMethod[string](n_, "internationalCurrencySymbol")
	return rv
}

func (n_ NumberFormatter) SetInternationalCurrencySymbol(value string) {
	ffi.CallMethod[ffi.Void](n_, "setInternationalCurrencySymbol:", value)
}

func (n_ NumberFormatter) CurrencyGroupingSeparator() string {
	rv := ffi.CallMethod[string](n_, "currencyGroupingSeparator")
	return rv
}

func (n_ NumberFormatter) SetCurrencyGroupingSeparator(value string) {
	ffi.CallMethod[ffi.Void](n_, "setCurrencyGroupingSeparator:", value)
}

func (n_ NumberFormatter) PositivePrefix() string {
	rv := ffi.CallMethod[string](n_, "positivePrefix")
	return rv
}

func (n_ NumberFormatter) SetPositivePrefix(value string) {
	ffi.CallMethod[ffi.Void](n_, "setPositivePrefix:", value)
}

func (n_ NumberFormatter) PositiveSuffix() string {
	rv := ffi.CallMethod[string](n_, "positiveSuffix")
	return rv
}

func (n_ NumberFormatter) SetPositiveSuffix(value string) {
	ffi.CallMethod[ffi.Void](n_, "setPositiveSuffix:", value)
}

func (n_ NumberFormatter) NegativePrefix() string {
	rv := ffi.CallMethod[string](n_, "negativePrefix")
	return rv
}

func (n_ NumberFormatter) SetNegativePrefix(value string) {
	ffi.CallMethod[ffi.Void](n_, "setNegativePrefix:", value)
}

func (n_ NumberFormatter) NegativeSuffix() string {
	rv := ffi.CallMethod[string](n_, "negativeSuffix")
	return rv
}

func (n_ NumberFormatter) SetNegativeSuffix(value string) {
	ffi.CallMethod[ffi.Void](n_, "setNegativeSuffix:", value)
}

func (n_ NumberFormatter) TextAttributesForNegativeValues() map[string]objc.Object {
	rv := ffi.CallMethod[map[string]objc.Object](n_, "textAttributesForNegativeValues")
	return rv
}

func (n_ NumberFormatter) SetTextAttributesForNegativeValues(value map[string]objc.IObject) {
	ffi.CallMethod[ffi.Void](n_, "setTextAttributesForNegativeValues:", value)
}

func (n_ NumberFormatter) TextAttributesForPositiveValues() map[string]objc.Object {
	rv := ffi.CallMethod[map[string]objc.Object](n_, "textAttributesForPositiveValues")
	return rv
}

func (n_ NumberFormatter) SetTextAttributesForPositiveValues(value map[string]objc.IObject) {
	ffi.CallMethod[ffi.Void](n_, "setTextAttributesForPositiveValues:", value)
}

func (n_ NumberFormatter) AttributedStringForZero() AttributedString {
	rv := ffi.CallMethod[AttributedString](n_, "attributedStringForZero")
	return rv
}

func (n_ NumberFormatter) SetAttributedStringForZero(value IAttributedString) {
	ffi.CallMethod[ffi.Void](n_, "setAttributedStringForZero:", value)
}

func (n_ NumberFormatter) TextAttributesForZero() map[string]objc.Object {
	rv := ffi.CallMethod[map[string]objc.Object](n_, "textAttributesForZero")
	return rv
}

func (n_ NumberFormatter) SetTextAttributesForZero(value map[string]objc.IObject) {
	ffi.CallMethod[ffi.Void](n_, "setTextAttributesForZero:", value)
}

func (n_ NumberFormatter) AttributedStringForNil() AttributedString {
	rv := ffi.CallMethod[AttributedString](n_, "attributedStringForNil")
	return rv
}

func (n_ NumberFormatter) SetAttributedStringForNil(value IAttributedString) {
	ffi.CallMethod[ffi.Void](n_, "setAttributedStringForNil:", value)
}

func (n_ NumberFormatter) TextAttributesForNil() map[string]objc.Object {
	rv := ffi.CallMethod[map[string]objc.Object](n_, "textAttributesForNil")
	return rv
}

func (n_ NumberFormatter) SetTextAttributesForNil(value map[string]objc.IObject) {
	ffi.CallMethod[ffi.Void](n_, "setTextAttributesForNil:", value)
}

func (n_ NumberFormatter) AttributedStringForNotANumber() AttributedString {
	rv := ffi.CallMethod[AttributedString](n_, "attributedStringForNotANumber")
	return rv
}

func (n_ NumberFormatter) SetAttributedStringForNotANumber(value IAttributedString) {
	ffi.CallMethod[ffi.Void](n_, "setAttributedStringForNotANumber:", value)
}

func (n_ NumberFormatter) TextAttributesForNotANumber() map[string]objc.Object {
	rv := ffi.CallMethod[map[string]objc.Object](n_, "textAttributesForNotANumber")
	return rv
}

func (n_ NumberFormatter) SetTextAttributesForNotANumber(value map[string]objc.IObject) {
	ffi.CallMethod[ffi.Void](n_, "setTextAttributesForNotANumber:", value)
}

func (n_ NumberFormatter) TextAttributesForPositiveInfinity() map[string]objc.Object {
	rv := ffi.CallMethod[map[string]objc.Object](n_, "textAttributesForPositiveInfinity")
	return rv
}

func (n_ NumberFormatter) SetTextAttributesForPositiveInfinity(value map[string]objc.IObject) {
	ffi.CallMethod[ffi.Void](n_, "setTextAttributesForPositiveInfinity:", value)
}

func (n_ NumberFormatter) TextAttributesForNegativeInfinity() map[string]objc.Object {
	rv := ffi.CallMethod[map[string]objc.Object](n_, "textAttributesForNegativeInfinity")
	return rv
}

func (n_ NumberFormatter) SetTextAttributesForNegativeInfinity(value map[string]objc.IObject) {
	ffi.CallMethod[ffi.Void](n_, "setTextAttributesForNegativeInfinity:", value)
}

func (n_ NumberFormatter) GroupingSeparator() string {
	rv := ffi.CallMethod[string](n_, "groupingSeparator")
	return rv
}

func (n_ NumberFormatter) SetGroupingSeparator(value string) {
	ffi.CallMethod[ffi.Void](n_, "setGroupingSeparator:", value)
}

func (n_ NumberFormatter) UsesGroupingSeparator() bool {
	rv := ffi.CallMethod[bool](n_, "usesGroupingSeparator")
	return rv
}

func (n_ NumberFormatter) SetUsesGroupingSeparator(value bool) {
	ffi.CallMethod[ffi.Void](n_, "setUsesGroupingSeparator:", value)
}

func (n_ NumberFormatter) ThousandSeparator() string {
	rv := ffi.CallMethod[string](n_, "thousandSeparator")
	return rv
}

func (n_ NumberFormatter) SetThousandSeparator(value string) {
	ffi.CallMethod[ffi.Void](n_, "setThousandSeparator:", value)
}

func (n_ NumberFormatter) HasThousandSeparators() bool {
	rv := ffi.CallMethod[bool](n_, "hasThousandSeparators")
	return rv
}

func (n_ NumberFormatter) SetHasThousandSeparators(value bool) {
	ffi.CallMethod[ffi.Void](n_, "setHasThousandSeparators:", value)
}

func (n_ NumberFormatter) DecimalSeparator() string {
	rv := ffi.CallMethod[string](n_, "decimalSeparator")
	return rv
}

func (n_ NumberFormatter) SetDecimalSeparator(value string) {
	ffi.CallMethod[ffi.Void](n_, "setDecimalSeparator:", value)
}

func (n_ NumberFormatter) AlwaysShowsDecimalSeparator() bool {
	rv := ffi.CallMethod[bool](n_, "alwaysShowsDecimalSeparator")
	return rv
}

func (n_ NumberFormatter) SetAlwaysShowsDecimalSeparator(value bool) {
	ffi.CallMethod[ffi.Void](n_, "setAlwaysShowsDecimalSeparator:", value)
}

func (n_ NumberFormatter) CurrencyDecimalSeparator() string {
	rv := ffi.CallMethod[string](n_, "currencyDecimalSeparator")
	return rv
}

func (n_ NumberFormatter) SetCurrencyDecimalSeparator(value string) {
	ffi.CallMethod[ffi.Void](n_, "setCurrencyDecimalSeparator:", value)
}

func (n_ NumberFormatter) GroupingSize() uint {
	rv := ffi.CallMethod[uint](n_, "groupingSize")
	return rv
}

func (n_ NumberFormatter) SetGroupingSize(value uint) {
	ffi.CallMethod[ffi.Void](n_, "setGroupingSize:", value)
}

func (n_ NumberFormatter) SecondaryGroupingSize() uint {
	rv := ffi.CallMethod[uint](n_, "secondaryGroupingSize")
	return rv
}

func (n_ NumberFormatter) SetSecondaryGroupingSize(value uint) {
	ffi.CallMethod[ffi.Void](n_, "setSecondaryGroupingSize:", value)
}

func (n_ NumberFormatter) PaddingCharacter() string {
	rv := ffi.CallMethod[string](n_, "paddingCharacter")
	return rv
}

func (n_ NumberFormatter) SetPaddingCharacter(value string) {
	ffi.CallMethod[ffi.Void](n_, "setPaddingCharacter:", value)
}

func (n_ NumberFormatter) PaddingPosition() NumberFormatterPadPosition {
	rv := ffi.CallMethod[NumberFormatterPadPosition](n_, "paddingPosition")
	return rv
}

func (n_ NumberFormatter) SetPaddingPosition(value NumberFormatterPadPosition) {
	ffi.CallMethod[ffi.Void](n_, "setPaddingPosition:", value)
}

func (n_ NumberFormatter) AllowsFloats() bool {
	rv := ffi.CallMethod[bool](n_, "allowsFloats")
	return rv
}

func (n_ NumberFormatter) SetAllowsFloats(value bool) {
	ffi.CallMethod[ffi.Void](n_, "setAllowsFloats:", value)
}

func (n_ NumberFormatter) Minimum() Number {
	rv := ffi.CallMethod[Number](n_, "minimum")
	return rv
}

func (n_ NumberFormatter) SetMinimum(value INumber) {
	ffi.CallMethod[ffi.Void](n_, "setMinimum:", value)
}

func (n_ NumberFormatter) Maximum() Number {
	rv := ffi.CallMethod[Number](n_, "maximum")
	return rv
}

func (n_ NumberFormatter) SetMaximum(value INumber) {
	ffi.CallMethod[ffi.Void](n_, "setMaximum:", value)
}

func (n_ NumberFormatter) IsLenient() bool {
	rv := ffi.CallMethod[bool](n_, "isLenient")
	return rv
}

func (n_ NumberFormatter) SetLenient(value bool) {
	ffi.CallMethod[ffi.Void](n_, "setLenient:", value)
}

func (n_ NumberFormatter) IsPartialStringValidationEnabled() bool {
	rv := ffi.CallMethod[bool](n_, "isPartialStringValidationEnabled")
	return rv
}

func (n_ NumberFormatter) SetPartialStringValidationEnabled(value bool) {
	ffi.CallMethod[ffi.Void](n_, "setPartialStringValidationEnabled:", value)
}
