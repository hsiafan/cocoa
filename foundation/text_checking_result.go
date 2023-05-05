// auto-generated code, do not modify
package foundation

import (
	"unsafe"

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
	rv := objc.CallMethod[TextCheckingResult](tc, objc.GetSelector("alloc"))
	return rv
}

func (tc _TextCheckingResultClass) New() TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextCheckingResult() TextCheckingResult {
	return TextCheckingResultClass.New()
}

func (t_ TextCheckingResult) Init() TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](t_, objc.GetSelector("init"))
	return rv
}

func (t_ TextCheckingResult) RangeAtIndex(idx uint) Range {
	rv := objc.CallMethod[Range](t_, objc.GetSelector("rangeAtIndex:"), idx)
	return rv
}

func (tc _TextCheckingResultClass) ReplacementCheckingResultWithRange_ReplacementString(range_ Range, replacementString string) TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](tc, objc.GetSelector("replacementCheckingResultWithRange:replacementString:"), range_, replacementString)
	return rv
}

func (tc _TextCheckingResultClass) RegularExpressionCheckingResultWithRanges_Count_RegularExpression(ranges *Range, count uint, regularExpression IRegularExpression) TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](tc, objc.GetSelector("regularExpressionCheckingResultWithRanges:count:regularExpression:"), ranges, count, objc.ExtractPtr(regularExpression))
	return rv
}

func (tc _TextCheckingResultClass) LinkCheckingResultWithRange_URL(range_ Range, url IURL) TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](tc, objc.GetSelector("linkCheckingResultWithRange:URL:"), range_, objc.ExtractPtr(url))
	return rv
}

func (tc _TextCheckingResultClass) AddressCheckingResultWithRange_Components(range_ Range, components map[TextCheckingKey]string) TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](tc, objc.GetSelector("addressCheckingResultWithRange:components:"), range_, components)
	return rv
}

func (tc _TextCheckingResultClass) TransitInformationCheckingResultWithRange_Components(range_ Range, components map[TextCheckingKey]string) TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](tc, objc.GetSelector("transitInformationCheckingResultWithRange:components:"), range_, components)
	return rv
}

func (tc _TextCheckingResultClass) PhoneNumberCheckingResultWithRange_PhoneNumber(range_ Range, phoneNumber string) TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](tc, objc.GetSelector("phoneNumberCheckingResultWithRange:phoneNumber:"), range_, phoneNumber)
	return rv
}

func (tc _TextCheckingResultClass) DateCheckingResultWithRange_Date(range_ Range, date IDate) TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](tc, objc.GetSelector("dateCheckingResultWithRange:date:"), range_, objc.ExtractPtr(date))
	return rv
}

func (tc _TextCheckingResultClass) DateCheckingResultWithRange_Date_TimeZone_Duration(range_ Range, date IDate, timeZone ITimeZone, duration TimeInterval) TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](tc, objc.GetSelector("dateCheckingResultWithRange:date:timeZone:duration:"), range_, objc.ExtractPtr(date), objc.ExtractPtr(timeZone), duration)
	return rv
}

func (tc _TextCheckingResultClass) DashCheckingResultWithRange_ReplacementString(range_ Range, replacementString string) TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](tc, objc.GetSelector("dashCheckingResultWithRange:replacementString:"), range_, replacementString)
	return rv
}

func (tc _TextCheckingResultClass) QuoteCheckingResultWithRange_ReplacementString(range_ Range, replacementString string) TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](tc, objc.GetSelector("quoteCheckingResultWithRange:replacementString:"), range_, replacementString)
	return rv
}

func (tc _TextCheckingResultClass) SpellCheckingResultWithRange(range_ Range) TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](tc, objc.GetSelector("spellCheckingResultWithRange:"), range_)
	return rv
}

func (tc _TextCheckingResultClass) CorrectionCheckingResultWithRange_ReplacementString(range_ Range, replacementString string) TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](tc, objc.GetSelector("correctionCheckingResultWithRange:replacementString:"), range_, replacementString)
	return rv
}

func (tc _TextCheckingResultClass) OrthographyCheckingResultWithRange_Orthography(range_ Range, orthography IOrthography) TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](tc, objc.GetSelector("orthographyCheckingResultWithRange:orthography:"), range_, objc.ExtractPtr(orthography))
	return rv
}

func (tc _TextCheckingResultClass) GrammarCheckingResultWithRange_Details(range_ Range, details []map[string]objc.IObject) TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](tc, objc.GetSelector("grammarCheckingResultWithRange:details:"), range_, details)
	return rv
}

func (t_ TextCheckingResult) ResultByAdjustingRangesWithOffset(offset int) TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](t_, objc.GetSelector("resultByAdjustingRangesWithOffset:"), offset)
	return rv
}

func (t_ TextCheckingResult) RangeWithName(name string) Range {
	rv := objc.CallMethod[Range](t_, objc.GetSelector("rangeWithName:"), name)
	return rv
}

func (tc _TextCheckingResultClass) CorrectionCheckingResultWithRange_ReplacementString_AlternativeStrings(range_ Range, replacementString string, alternativeStrings []string) TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](tc, objc.GetSelector("correctionCheckingResultWithRange:replacementString:alternativeStrings:"), range_, replacementString, alternativeStrings)
	return rv
}

func (t_ TextCheckingResult) Range() Range {
	rv := objc.CallMethod[Range](t_, objc.GetSelector("range"))
	return rv
}

func (t_ TextCheckingResult) ResultType() TextCheckingType {
	rv := objc.CallMethod[TextCheckingType](t_, objc.GetSelector("resultType"))
	return rv
}

func (t_ TextCheckingResult) NumberOfRanges() uint {
	rv := objc.CallMethod[uint](t_, objc.GetSelector("numberOfRanges"))
	return rv
}

func (t_ TextCheckingResult) ReplacementString() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("replacementString"))
	return rv
}

func (t_ TextCheckingResult) RegularExpression() RegularExpression {
	rv := objc.CallMethod[RegularExpression](t_, objc.GetSelector("regularExpression"))
	return rv
}

func (t_ TextCheckingResult) Components() map[TextCheckingKey]string {
	rv := objc.CallMethod[map[TextCheckingKey]string](t_, objc.GetSelector("components"))
	return rv
}

func (t_ TextCheckingResult) URL() URL {
	rv := objc.CallMethod[URL](t_, objc.GetSelector("URL"))
	return rv
}

func (t_ TextCheckingResult) AddressComponents() map[TextCheckingKey]string {
	rv := objc.CallMethod[map[TextCheckingKey]string](t_, objc.GetSelector("addressComponents"))
	return rv
}

func (t_ TextCheckingResult) PhoneNumber() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("phoneNumber"))
	return rv
}

func (t_ TextCheckingResult) Date() Date {
	rv := objc.CallMethod[Date](t_, objc.GetSelector("date"))
	return rv
}

func (t_ TextCheckingResult) Duration() TimeInterval {
	rv := objc.CallMethod[TimeInterval](t_, objc.GetSelector("duration"))
	return rv
}

func (t_ TextCheckingResult) TimeZone() TimeZone {
	rv := objc.CallMethod[TimeZone](t_, objc.GetSelector("timeZone"))
	return rv
}

func (t_ TextCheckingResult) Orthography() Orthography {
	rv := objc.CallMethod[Orthography](t_, objc.GetSelector("orthography"))
	return rv
}

func (t_ TextCheckingResult) GrammarDetails() []map[string]objc.Object {
	rv := objc.CallMethod[[]map[string]objc.Object](t_, objc.GetSelector("grammarDetails"))
	return rv
}

func (t_ TextCheckingResult) AlternativeStrings() []string {
	rv := objc.CallMethod[[]string](t_, objc.GetSelector("alternativeStrings"))
	return rv
}
