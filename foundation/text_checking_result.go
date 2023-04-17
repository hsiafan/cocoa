// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var TextCheckingResultClass = _TextCheckingResultClass{objc.GetClass("NSTextCheckingResult")}

type _TextCheckingResultClass struct {
	objc.Class
}

type ITextCheckingResult interface {
	objc.IObject
	RangeAtIndex(idx uint) Range
	ResultByAdjustingRangesWithOffset(offset int) TextCheckingResult
	RangeWithName(name string) Range
	Range() Range
	ResultType() TextCheckingType
	NumberOfRanges() uint
	ReplacementString() string
	RegularExpression() RegularExpression
	Components() map[TextCheckingKey]string
	URL() URL
	AddressComponents() map[TextCheckingKey]string
	PhoneNumber() string
	Date() Date
	Duration() TimeInterval
	TimeZone() TimeZone
	Orthography() Orthography
	GrammarDetails() []map[string]objc.Object
	AlternativeStrings() []string
}

type TextCheckingResult struct {
	objc.Object
}

func MakeTextCheckingResult(ptr unsafe.Pointer) TextCheckingResult {
	return TextCheckingResult{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TextCheckingResultClass) Alloc() TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](tc, "alloc")
	return rv
}

func (tc _TextCheckingResultClass) New() TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextCheckingResult() TextCheckingResult {
	return TextCheckingResultClass.New()
}

func (t_ TextCheckingResult) Init() TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](t_, "init")
	return rv
}

func (t_ TextCheckingResult) RangeAtIndex(idx uint) Range {
	rv := ffi.CallMethod[Range](t_, "rangeAtIndex:", idx)
	return rv
}

func (tc _TextCheckingResultClass) ReplacementCheckingResultWithRange_ReplacementString(range_ Range, replacementString string) TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](tc, "replacementCheckingResultWithRange:replacementString:", range_, replacementString)
	return rv
}

func (tc _TextCheckingResultClass) RegularExpressionCheckingResultWithRanges_Count_RegularExpression(ranges *Range, count uint, regularExpression IRegularExpression) TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](tc, "regularExpressionCheckingResultWithRanges:count:regularExpression:", ranges, count, regularExpression)
	return rv
}

func (tc _TextCheckingResultClass) LinkCheckingResultWithRange_URL(range_ Range, url IURL) TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](tc, "linkCheckingResultWithRange:URL:", range_, url)
	return rv
}

func (tc _TextCheckingResultClass) AddressCheckingResultWithRange_Components(range_ Range, components map[TextCheckingKey]string) TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](tc, "addressCheckingResultWithRange:components:", range_, components)
	return rv
}

func (tc _TextCheckingResultClass) TransitInformationCheckingResultWithRange_Components(range_ Range, components map[TextCheckingKey]string) TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](tc, "transitInformationCheckingResultWithRange:components:", range_, components)
	return rv
}

func (tc _TextCheckingResultClass) PhoneNumberCheckingResultWithRange_PhoneNumber(range_ Range, phoneNumber string) TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](tc, "phoneNumberCheckingResultWithRange:phoneNumber:", range_, phoneNumber)
	return rv
}

func (tc _TextCheckingResultClass) DateCheckingResultWithRange_Date(range_ Range, date IDate) TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](tc, "dateCheckingResultWithRange:date:", range_, date)
	return rv
}

func (tc _TextCheckingResultClass) DateCheckingResultWithRange_Date_TimeZone_Duration(range_ Range, date IDate, timeZone ITimeZone, duration TimeInterval) TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](tc, "dateCheckingResultWithRange:date:timeZone:duration:", range_, date, timeZone, duration)
	return rv
}

func (tc _TextCheckingResultClass) DashCheckingResultWithRange_ReplacementString(range_ Range, replacementString string) TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](tc, "dashCheckingResultWithRange:replacementString:", range_, replacementString)
	return rv
}

func (tc _TextCheckingResultClass) QuoteCheckingResultWithRange_ReplacementString(range_ Range, replacementString string) TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](tc, "quoteCheckingResultWithRange:replacementString:", range_, replacementString)
	return rv
}

func (tc _TextCheckingResultClass) SpellCheckingResultWithRange(range_ Range) TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](tc, "spellCheckingResultWithRange:", range_)
	return rv
}

func (tc _TextCheckingResultClass) CorrectionCheckingResultWithRange_ReplacementString(range_ Range, replacementString string) TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](tc, "correctionCheckingResultWithRange:replacementString:", range_, replacementString)
	return rv
}

func (tc _TextCheckingResultClass) OrthographyCheckingResultWithRange_Orthography(range_ Range, orthography IOrthography) TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](tc, "orthographyCheckingResultWithRange:orthography:", range_, orthography)
	return rv
}

func (tc _TextCheckingResultClass) GrammarCheckingResultWithRange_Details(range_ Range, details []map[string]objc.IObject) TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](tc, "grammarCheckingResultWithRange:details:", range_, details)
	return rv
}

func (t_ TextCheckingResult) ResultByAdjustingRangesWithOffset(offset int) TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](t_, "resultByAdjustingRangesWithOffset:", offset)
	return rv
}

func (t_ TextCheckingResult) RangeWithName(name string) Range {
	rv := ffi.CallMethod[Range](t_, "rangeWithName:", name)
	return rv
}

func (tc _TextCheckingResultClass) CorrectionCheckingResultWithRange_ReplacementString_AlternativeStrings(range_ Range, replacementString string, alternativeStrings []string) TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](tc, "correctionCheckingResultWithRange:replacementString:alternativeStrings:", range_, replacementString, alternativeStrings)
	return rv
}

func (t_ TextCheckingResult) Range() Range {
	rv := ffi.CallMethod[Range](t_, "range")
	return rv
}

func (t_ TextCheckingResult) ResultType() TextCheckingType {
	rv := ffi.CallMethod[TextCheckingType](t_, "resultType")
	return rv
}

func (t_ TextCheckingResult) NumberOfRanges() uint {
	rv := ffi.CallMethod[uint](t_, "numberOfRanges")
	return rv
}

func (t_ TextCheckingResult) ReplacementString() string {
	rv := ffi.CallMethod[string](t_, "replacementString")
	return rv
}

func (t_ TextCheckingResult) RegularExpression() RegularExpression {
	rv := ffi.CallMethod[RegularExpression](t_, "regularExpression")
	return rv
}

func (t_ TextCheckingResult) Components() map[TextCheckingKey]string {
	rv := ffi.CallMethod[map[TextCheckingKey]string](t_, "components")
	return rv
}

func (t_ TextCheckingResult) URL() URL {
	rv := ffi.CallMethod[URL](t_, "URL")
	return rv
}

func (t_ TextCheckingResult) AddressComponents() map[TextCheckingKey]string {
	rv := ffi.CallMethod[map[TextCheckingKey]string](t_, "addressComponents")
	return rv
}

func (t_ TextCheckingResult) PhoneNumber() string {
	rv := ffi.CallMethod[string](t_, "phoneNumber")
	return rv
}

func (t_ TextCheckingResult) Date() Date {
	rv := ffi.CallMethod[Date](t_, "date")
	return rv
}

func (t_ TextCheckingResult) Duration() TimeInterval {
	rv := ffi.CallMethod[TimeInterval](t_, "duration")
	return rv
}

func (t_ TextCheckingResult) TimeZone() TimeZone {
	rv := ffi.CallMethod[TimeZone](t_, "timeZone")
	return rv
}

func (t_ TextCheckingResult) Orthography() Orthography {
	rv := ffi.CallMethod[Orthography](t_, "orthography")
	return rv
}

func (t_ TextCheckingResult) GrammarDetails() []map[string]objc.Object {
	rv := ffi.CallMethod[[]map[string]objc.Object](t_, "grammarDetails")
	return rv
}

func (t_ TextCheckingResult) AlternativeStrings() []string {
	rv := ffi.CallMethod[[]string](t_, "alternativeStrings")
	return rv
}
