// auto-generated code, do not modify
package foundation

import (
	"unsafe"

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
	rv := objc.CallMethod[NumberFormatter](nc, objc.SEL("alloc"))
	return rv
}

func (nc _NumberFormatterClass) New() NumberFormatter {
	rv := objc.CallMethod[NumberFormatter](nc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewNumberFormatter() NumberFormatter {
	return NumberFormatterClass.New()
}

func (n_ NumberFormatter) Init() NumberFormatter {
	rv := objc.CallMethod[NumberFormatter](n_, objc.SEL("init"))
	return rv
}

func (nc _NumberFormatterClass) SetDefaultFormatterBehavior(behavior NumberFormatterBehavior) {
	objc.CallMethod[objc.Void](nc, objc.SEL("setDefaultFormatterBehavior:"), behavior)
}

func (nc _NumberFormatterClass) DefaultFormatterBehavior() NumberFormatterBehavior {
	rv := objc.CallMethod[NumberFormatterBehavior](nc, objc.SEL("defaultFormatterBehavior"))
	return rv
}

func (n_ NumberFormatter) NumberFromString(string_ string) Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("numberFromString:"), string_)
	return rv
}

func (n_ NumberFormatter) StringFromNumber(number INumber) string {
	rv := objc.CallMethod[string](n_, objc.SEL("stringFromNumber:"), objc.ExtractPtr(number))
	return rv
}

func (nc _NumberFormatterClass) LocalizedStringFromNumber_NumberStyle(num INumber, nstyle NumberFormatterStyle) string {
	rv := objc.CallMethod[string](nc, objc.SEL("localizedStringFromNumber:numberStyle:"), objc.ExtractPtr(num), nstyle)
	return rv
}

func (n_ NumberFormatter) FormatterBehavior() NumberFormatterBehavior {
	rv := objc.CallMethod[NumberFormatterBehavior](n_, objc.SEL("formatterBehavior"))
	return rv
}

func (n_ NumberFormatter) SetFormatterBehavior(value NumberFormatterBehavior) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setFormatterBehavior:"), value)
}

func (n_ NumberFormatter) NumberStyle() NumberFormatterStyle {
	rv := objc.CallMethod[NumberFormatterStyle](n_, objc.SEL("numberStyle"))
	return rv
}

func (n_ NumberFormatter) SetNumberStyle(value NumberFormatterStyle) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setNumberStyle:"), value)
}

func (n_ NumberFormatter) GeneratesDecimalNumbers() bool {
	rv := objc.CallMethod[bool](n_, objc.SEL("generatesDecimalNumbers"))
	return rv
}

func (n_ NumberFormatter) SetGeneratesDecimalNumbers(value bool) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setGeneratesDecimalNumbers:"), value)
}

func (n_ NumberFormatter) LocalizesFormat() bool {
	rv := objc.CallMethod[bool](n_, objc.SEL("localizesFormat"))
	return rv
}

func (n_ NumberFormatter) SetLocalizesFormat(value bool) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setLocalizesFormat:"), value)
}

func (n_ NumberFormatter) Locale() Locale {
	rv := objc.CallMethod[Locale](n_, objc.SEL("locale"))
	return rv
}

func (n_ NumberFormatter) SetLocale(value ILocale) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setLocale:"), objc.ExtractPtr(value))
}

func (n_ NumberFormatter) RoundingBehavior() DecimalNumberHandler {
	rv := objc.CallMethod[DecimalNumberHandler](n_, objc.SEL("roundingBehavior"))
	return rv
}

func (n_ NumberFormatter) SetRoundingBehavior(value IDecimalNumberHandler) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setRoundingBehavior:"), objc.ExtractPtr(value))
}

func (n_ NumberFormatter) RoundingIncrement() Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("roundingIncrement"))
	return rv
}

func (n_ NumberFormatter) SetRoundingIncrement(value INumber) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setRoundingIncrement:"), objc.ExtractPtr(value))
}

func (n_ NumberFormatter) RoundingMode() NumberFormatterRoundingMode {
	rv := objc.CallMethod[NumberFormatterRoundingMode](n_, objc.SEL("roundingMode"))
	return rv
}

func (n_ NumberFormatter) SetRoundingMode(value NumberFormatterRoundingMode) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setRoundingMode:"), value)
}

func (n_ NumberFormatter) MinimumIntegerDigits() uint {
	rv := objc.CallMethod[uint](n_, objc.SEL("minimumIntegerDigits"))
	return rv
}

func (n_ NumberFormatter) SetMinimumIntegerDigits(value uint) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setMinimumIntegerDigits:"), value)
}

func (n_ NumberFormatter) MaximumIntegerDigits() uint {
	rv := objc.CallMethod[uint](n_, objc.SEL("maximumIntegerDigits"))
	return rv
}

func (n_ NumberFormatter) SetMaximumIntegerDigits(value uint) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setMaximumIntegerDigits:"), value)
}

func (n_ NumberFormatter) MinimumFractionDigits() uint {
	rv := objc.CallMethod[uint](n_, objc.SEL("minimumFractionDigits"))
	return rv
}

func (n_ NumberFormatter) SetMinimumFractionDigits(value uint) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setMinimumFractionDigits:"), value)
}

func (n_ NumberFormatter) MaximumFractionDigits() uint {
	rv := objc.CallMethod[uint](n_, objc.SEL("maximumFractionDigits"))
	return rv
}

func (n_ NumberFormatter) SetMaximumFractionDigits(value uint) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setMaximumFractionDigits:"), value)
}

func (n_ NumberFormatter) UsesSignificantDigits() bool {
	rv := objc.CallMethod[bool](n_, objc.SEL("usesSignificantDigits"))
	return rv
}

func (n_ NumberFormatter) SetUsesSignificantDigits(value bool) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setUsesSignificantDigits:"), value)
}

func (n_ NumberFormatter) MinimumSignificantDigits() uint {
	rv := objc.CallMethod[uint](n_, objc.SEL("minimumSignificantDigits"))
	return rv
}

func (n_ NumberFormatter) SetMinimumSignificantDigits(value uint) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setMinimumSignificantDigits:"), value)
}

func (n_ NumberFormatter) MaximumSignificantDigits() uint {
	rv := objc.CallMethod[uint](n_, objc.SEL("maximumSignificantDigits"))
	return rv
}

func (n_ NumberFormatter) SetMaximumSignificantDigits(value uint) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setMaximumSignificantDigits:"), value)
}

func (n_ NumberFormatter) Format() string {
	rv := objc.CallMethod[string](n_, objc.SEL("format"))
	return rv
}

func (n_ NumberFormatter) SetFormat(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setFormat:"), value)
}

func (n_ NumberFormatter) FormattingContext() FormattingContext {
	rv := objc.CallMethod[FormattingContext](n_, objc.SEL("formattingContext"))
	return rv
}

func (n_ NumberFormatter) SetFormattingContext(value FormattingContext) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setFormattingContext:"), value)
}

func (n_ NumberFormatter) FormatWidth() uint {
	rv := objc.CallMethod[uint](n_, objc.SEL("formatWidth"))
	return rv
}

func (n_ NumberFormatter) SetFormatWidth(value uint) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setFormatWidth:"), value)
}

func (n_ NumberFormatter) NegativeFormat() string {
	rv := objc.CallMethod[string](n_, objc.SEL("negativeFormat"))
	return rv
}

func (n_ NumberFormatter) SetNegativeFormat(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setNegativeFormat:"), value)
}

func (n_ NumberFormatter) PositiveFormat() string {
	rv := objc.CallMethod[string](n_, objc.SEL("positiveFormat"))
	return rv
}

func (n_ NumberFormatter) SetPositiveFormat(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setPositiveFormat:"), value)
}

func (n_ NumberFormatter) Multiplier() Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("multiplier"))
	return rv
}

func (n_ NumberFormatter) SetMultiplier(value INumber) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setMultiplier:"), objc.ExtractPtr(value))
}

func (n_ NumberFormatter) PercentSymbol() string {
	rv := objc.CallMethod[string](n_, objc.SEL("percentSymbol"))
	return rv
}

func (n_ NumberFormatter) SetPercentSymbol(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setPercentSymbol:"), value)
}

func (n_ NumberFormatter) PerMillSymbol() string {
	rv := objc.CallMethod[string](n_, objc.SEL("perMillSymbol"))
	return rv
}

func (n_ NumberFormatter) SetPerMillSymbol(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setPerMillSymbol:"), value)
}

func (n_ NumberFormatter) MinusSign() string {
	rv := objc.CallMethod[string](n_, objc.SEL("minusSign"))
	return rv
}

func (n_ NumberFormatter) SetMinusSign(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setMinusSign:"), value)
}

func (n_ NumberFormatter) PlusSign() string {
	rv := objc.CallMethod[string](n_, objc.SEL("plusSign"))
	return rv
}

func (n_ NumberFormatter) SetPlusSign(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setPlusSign:"), value)
}

func (n_ NumberFormatter) ExponentSymbol() string {
	rv := objc.CallMethod[string](n_, objc.SEL("exponentSymbol"))
	return rv
}

func (n_ NumberFormatter) SetExponentSymbol(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setExponentSymbol:"), value)
}

func (n_ NumberFormatter) ZeroSymbol() string {
	rv := objc.CallMethod[string](n_, objc.SEL("zeroSymbol"))
	return rv
}

func (n_ NumberFormatter) SetZeroSymbol(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setZeroSymbol:"), value)
}

func (n_ NumberFormatter) NilSymbol() string {
	rv := objc.CallMethod[string](n_, objc.SEL("nilSymbol"))
	return rv
}

func (n_ NumberFormatter) SetNilSymbol(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setNilSymbol:"), value)
}

func (n_ NumberFormatter) NotANumberSymbol() string {
	rv := objc.CallMethod[string](n_, objc.SEL("notANumberSymbol"))
	return rv
}

func (n_ NumberFormatter) SetNotANumberSymbol(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setNotANumberSymbol:"), value)
}

func (n_ NumberFormatter) NegativeInfinitySymbol() string {
	rv := objc.CallMethod[string](n_, objc.SEL("negativeInfinitySymbol"))
	return rv
}

func (n_ NumberFormatter) SetNegativeInfinitySymbol(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setNegativeInfinitySymbol:"), value)
}

func (n_ NumberFormatter) PositiveInfinitySymbol() string {
	rv := objc.CallMethod[string](n_, objc.SEL("positiveInfinitySymbol"))
	return rv
}

func (n_ NumberFormatter) SetPositiveInfinitySymbol(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setPositiveInfinitySymbol:"), value)
}

func (n_ NumberFormatter) CurrencySymbol() string {
	rv := objc.CallMethod[string](n_, objc.SEL("currencySymbol"))
	return rv
}

func (n_ NumberFormatter) SetCurrencySymbol(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setCurrencySymbol:"), value)
}

func (n_ NumberFormatter) CurrencyCode() string {
	rv := objc.CallMethod[string](n_, objc.SEL("currencyCode"))
	return rv
}

func (n_ NumberFormatter) SetCurrencyCode(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setCurrencyCode:"), value)
}

func (n_ NumberFormatter) InternationalCurrencySymbol() string {
	rv := objc.CallMethod[string](n_, objc.SEL("internationalCurrencySymbol"))
	return rv
}

func (n_ NumberFormatter) SetInternationalCurrencySymbol(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setInternationalCurrencySymbol:"), value)
}

func (n_ NumberFormatter) CurrencyGroupingSeparator() string {
	rv := objc.CallMethod[string](n_, objc.SEL("currencyGroupingSeparator"))
	return rv
}

func (n_ NumberFormatter) SetCurrencyGroupingSeparator(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setCurrencyGroupingSeparator:"), value)
}

func (n_ NumberFormatter) PositivePrefix() string {
	rv := objc.CallMethod[string](n_, objc.SEL("positivePrefix"))
	return rv
}

func (n_ NumberFormatter) SetPositivePrefix(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setPositivePrefix:"), value)
}

func (n_ NumberFormatter) PositiveSuffix() string {
	rv := objc.CallMethod[string](n_, objc.SEL("positiveSuffix"))
	return rv
}

func (n_ NumberFormatter) SetPositiveSuffix(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setPositiveSuffix:"), value)
}

func (n_ NumberFormatter) NegativePrefix() string {
	rv := objc.CallMethod[string](n_, objc.SEL("negativePrefix"))
	return rv
}

func (n_ NumberFormatter) SetNegativePrefix(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setNegativePrefix:"), value)
}

func (n_ NumberFormatter) NegativeSuffix() string {
	rv := objc.CallMethod[string](n_, objc.SEL("negativeSuffix"))
	return rv
}

func (n_ NumberFormatter) SetNegativeSuffix(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setNegativeSuffix:"), value)
}

func (n_ NumberFormatter) TextAttributesForNegativeValues() map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](n_, objc.SEL("textAttributesForNegativeValues"))
	return rv
}

func (n_ NumberFormatter) SetTextAttributesForNegativeValues(value map[string]objc.IObject) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setTextAttributesForNegativeValues:"), value)
}

func (n_ NumberFormatter) TextAttributesForPositiveValues() map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](n_, objc.SEL("textAttributesForPositiveValues"))
	return rv
}

func (n_ NumberFormatter) SetTextAttributesForPositiveValues(value map[string]objc.IObject) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setTextAttributesForPositiveValues:"), value)
}

func (n_ NumberFormatter) AttributedStringForZero() AttributedString {
	rv := objc.CallMethod[AttributedString](n_, objc.SEL("attributedStringForZero"))
	return rv
}

func (n_ NumberFormatter) SetAttributedStringForZero(value IAttributedString) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setAttributedStringForZero:"), objc.ExtractPtr(value))
}

func (n_ NumberFormatter) TextAttributesForZero() map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](n_, objc.SEL("textAttributesForZero"))
	return rv
}

func (n_ NumberFormatter) SetTextAttributesForZero(value map[string]objc.IObject) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setTextAttributesForZero:"), value)
}

func (n_ NumberFormatter) AttributedStringForNil() AttributedString {
	rv := objc.CallMethod[AttributedString](n_, objc.SEL("attributedStringForNil"))
	return rv
}

func (n_ NumberFormatter) SetAttributedStringForNil(value IAttributedString) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setAttributedStringForNil:"), objc.ExtractPtr(value))
}

func (n_ NumberFormatter) TextAttributesForNil() map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](n_, objc.SEL("textAttributesForNil"))
	return rv
}

func (n_ NumberFormatter) SetTextAttributesForNil(value map[string]objc.IObject) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setTextAttributesForNil:"), value)
}

func (n_ NumberFormatter) AttributedStringForNotANumber() AttributedString {
	rv := objc.CallMethod[AttributedString](n_, objc.SEL("attributedStringForNotANumber"))
	return rv
}

func (n_ NumberFormatter) SetAttributedStringForNotANumber(value IAttributedString) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setAttributedStringForNotANumber:"), objc.ExtractPtr(value))
}

func (n_ NumberFormatter) TextAttributesForNotANumber() map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](n_, objc.SEL("textAttributesForNotANumber"))
	return rv
}

func (n_ NumberFormatter) SetTextAttributesForNotANumber(value map[string]objc.IObject) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setTextAttributesForNotANumber:"), value)
}

func (n_ NumberFormatter) TextAttributesForPositiveInfinity() map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](n_, objc.SEL("textAttributesForPositiveInfinity"))
	return rv
}

func (n_ NumberFormatter) SetTextAttributesForPositiveInfinity(value map[string]objc.IObject) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setTextAttributesForPositiveInfinity:"), value)
}

func (n_ NumberFormatter) TextAttributesForNegativeInfinity() map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](n_, objc.SEL("textAttributesForNegativeInfinity"))
	return rv
}

func (n_ NumberFormatter) SetTextAttributesForNegativeInfinity(value map[string]objc.IObject) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setTextAttributesForNegativeInfinity:"), value)
}

func (n_ NumberFormatter) GroupingSeparator() string {
	rv := objc.CallMethod[string](n_, objc.SEL("groupingSeparator"))
	return rv
}

func (n_ NumberFormatter) SetGroupingSeparator(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setGroupingSeparator:"), value)
}

func (n_ NumberFormatter) UsesGroupingSeparator() bool {
	rv := objc.CallMethod[bool](n_, objc.SEL("usesGroupingSeparator"))
	return rv
}

func (n_ NumberFormatter) SetUsesGroupingSeparator(value bool) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setUsesGroupingSeparator:"), value)
}

func (n_ NumberFormatter) ThousandSeparator() string {
	rv := objc.CallMethod[string](n_, objc.SEL("thousandSeparator"))
	return rv
}

func (n_ NumberFormatter) SetThousandSeparator(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setThousandSeparator:"), value)
}

func (n_ NumberFormatter) HasThousandSeparators() bool {
	rv := objc.CallMethod[bool](n_, objc.SEL("hasThousandSeparators"))
	return rv
}

func (n_ NumberFormatter) SetHasThousandSeparators(value bool) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setHasThousandSeparators:"), value)
}

func (n_ NumberFormatter) DecimalSeparator() string {
	rv := objc.CallMethod[string](n_, objc.SEL("decimalSeparator"))
	return rv
}

func (n_ NumberFormatter) SetDecimalSeparator(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setDecimalSeparator:"), value)
}

func (n_ NumberFormatter) AlwaysShowsDecimalSeparator() bool {
	rv := objc.CallMethod[bool](n_, objc.SEL("alwaysShowsDecimalSeparator"))
	return rv
}

func (n_ NumberFormatter) SetAlwaysShowsDecimalSeparator(value bool) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setAlwaysShowsDecimalSeparator:"), value)
}

func (n_ NumberFormatter) CurrencyDecimalSeparator() string {
	rv := objc.CallMethod[string](n_, objc.SEL("currencyDecimalSeparator"))
	return rv
}

func (n_ NumberFormatter) SetCurrencyDecimalSeparator(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setCurrencyDecimalSeparator:"), value)
}

func (n_ NumberFormatter) GroupingSize() uint {
	rv := objc.CallMethod[uint](n_, objc.SEL("groupingSize"))
	return rv
}

func (n_ NumberFormatter) SetGroupingSize(value uint) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setGroupingSize:"), value)
}

func (n_ NumberFormatter) SecondaryGroupingSize() uint {
	rv := objc.CallMethod[uint](n_, objc.SEL("secondaryGroupingSize"))
	return rv
}

func (n_ NumberFormatter) SetSecondaryGroupingSize(value uint) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setSecondaryGroupingSize:"), value)
}

func (n_ NumberFormatter) PaddingCharacter() string {
	rv := objc.CallMethod[string](n_, objc.SEL("paddingCharacter"))
	return rv
}

func (n_ NumberFormatter) SetPaddingCharacter(value string) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setPaddingCharacter:"), value)
}

func (n_ NumberFormatter) PaddingPosition() NumberFormatterPadPosition {
	rv := objc.CallMethod[NumberFormatterPadPosition](n_, objc.SEL("paddingPosition"))
	return rv
}

func (n_ NumberFormatter) SetPaddingPosition(value NumberFormatterPadPosition) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setPaddingPosition:"), value)
}

func (n_ NumberFormatter) AllowsFloats() bool {
	rv := objc.CallMethod[bool](n_, objc.SEL("allowsFloats"))
	return rv
}

func (n_ NumberFormatter) SetAllowsFloats(value bool) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setAllowsFloats:"), value)
}

func (n_ NumberFormatter) Minimum() Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("minimum"))
	return rv
}

func (n_ NumberFormatter) SetMinimum(value INumber) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setMinimum:"), objc.ExtractPtr(value))
}

func (n_ NumberFormatter) Maximum() Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("maximum"))
	return rv
}

func (n_ NumberFormatter) SetMaximum(value INumber) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setMaximum:"), objc.ExtractPtr(value))
}

func (n_ NumberFormatter) IsLenient() bool {
	rv := objc.CallMethod[bool](n_, objc.SEL("isLenient"))
	return rv
}

func (n_ NumberFormatter) SetLenient(value bool) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setLenient:"), value)
}

func (n_ NumberFormatter) IsPartialStringValidationEnabled() bool {
	rv := objc.CallMethod[bool](n_, objc.SEL("isPartialStringValidationEnabled"))
	return rv
}

func (n_ NumberFormatter) SetPartialStringValidationEnabled(value bool) {
	objc.CallMethod[objc.Void](n_, objc.SEL("setPartialStringValidationEnabled:"), value)
}
