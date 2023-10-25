// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var RegularExpressionClass = _RegularExpressionClass{objc.GetClass("NSRegularExpression")}

type _RegularExpressionClass struct {
	objc.Class
}

type IRegularExpression interface {
	objc.IObject
	NumberOfMatchesInString_Options_Range(string_ string, options MatchingOptions, range_ Range) uint
	EnumerateMatchesInString_Options_Range_UsingBlock(string_ string, options MatchingOptions, range_ Range, block func(result TextCheckingResult, flags MatchingFlags, stop *bool))
	MatchesInString_Options_Range(string_ string, options MatchingOptions, range_ Range) []TextCheckingResult
	FirstMatchInString_Options_Range(string_ string, options MatchingOptions, range_ Range) TextCheckingResult
	RangeOfFirstMatchInString_Options_Range(string_ string, options MatchingOptions, range_ Range) Range
	StringByReplacingMatchesInString_Options_Range_WithTemplate(string_ string, options MatchingOptions, range_ Range, templ string) string
	ReplacementStringForResult_InString_Offset_Template(result ITextCheckingResult, string_ string, offset int, templ string) string
	Pattern() string
	Options() RegularExpressionOptions
	NumberOfCaptureGroups() uint
}

type RegularExpression struct {
	objc.Object
}

func MakeRegularExpression(ptr unsafe.Pointer) RegularExpression {
	return RegularExpression{
		Object: objc.MakeObject(ptr),
	}
}

func (r_ RegularExpression) InitWithPattern_Options_Error(pattern string, options RegularExpressionOptions, error *Error) RegularExpression {
	rv := objc.CallMethod[RegularExpression](r_, objc.SEL("initWithPattern:options:error:"), pattern, options, unsafe.Pointer(error))
	return rv
}

func (rc _RegularExpressionClass) Alloc() RegularExpression {
	rv := objc.CallMethod[RegularExpression](rc, objc.SEL("alloc"))
	return rv
}

func (rc _RegularExpressionClass) New() RegularExpression {
	rv := objc.CallMethod[RegularExpression](rc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewRegularExpression() RegularExpression {
	return RegularExpressionClass.New()
}

func (r_ RegularExpression) Init() RegularExpression {
	rv := objc.CallMethod[RegularExpression](r_, objc.SEL("init"))
	return rv
}

func (rc _RegularExpressionClass) RegularExpressionWithPattern_Options_Error(pattern string, options RegularExpressionOptions, error *Error) RegularExpression {
	rv := objc.CallMethod[RegularExpression](rc, objc.SEL("regularExpressionWithPattern:options:error:"), pattern, options, unsafe.Pointer(error))
	return rv
}

func (r_ RegularExpression) NumberOfMatchesInString_Options_Range(string_ string, options MatchingOptions, range_ Range) uint {
	rv := objc.CallMethod[uint](r_, objc.SEL("numberOfMatchesInString:options:range:"), string_, options, range_)
	return rv
}

func (r_ RegularExpression) EnumerateMatchesInString_Options_Range_UsingBlock(string_ string, options MatchingOptions, range_ Range, block func(result TextCheckingResult, flags MatchingFlags, stop *bool)) {
	objc.CallMethod[objc.Void](r_, objc.SEL("enumerateMatchesInString:options:range:usingBlock:"), string_, options, range_, block)
}

func (r_ RegularExpression) MatchesInString_Options_Range(string_ string, options MatchingOptions, range_ Range) []TextCheckingResult {
	rv := objc.CallMethod[[]TextCheckingResult](r_, objc.SEL("matchesInString:options:range:"), string_, options, range_)
	return rv
}

func (r_ RegularExpression) FirstMatchInString_Options_Range(string_ string, options MatchingOptions, range_ Range) TextCheckingResult {
	rv := objc.CallMethod[TextCheckingResult](r_, objc.SEL("firstMatchInString:options:range:"), string_, options, range_)
	return rv
}

func (r_ RegularExpression) RangeOfFirstMatchInString_Options_Range(string_ string, options MatchingOptions, range_ Range) Range {
	rv := objc.CallMethod[Range](r_, objc.SEL("rangeOfFirstMatchInString:options:range:"), string_, options, range_)
	return rv
}

func (r_ RegularExpression) StringByReplacingMatchesInString_Options_Range_WithTemplate(string_ string, options MatchingOptions, range_ Range, templ string) string {
	rv := objc.CallMethod[string](r_, objc.SEL("stringByReplacingMatchesInString:options:range:withTemplate:"), string_, options, range_, templ)
	return rv
}

func (rc _RegularExpressionClass) EscapedTemplateForString(string_ string) string {
	rv := objc.CallMethod[string](rc, objc.SEL("escapedTemplateForString:"), string_)
	return rv
}

func (rc _RegularExpressionClass) EscapedPatternForString(string_ string) string {
	rv := objc.CallMethod[string](rc, objc.SEL("escapedPatternForString:"), string_)
	return rv
}

func (r_ RegularExpression) ReplacementStringForResult_InString_Offset_Template(result ITextCheckingResult, string_ string, offset int, templ string) string {
	rv := objc.CallMethod[string](r_, objc.SEL("replacementStringForResult:inString:offset:template:"), objc.ExtractPtr(result), string_, offset, templ)
	return rv
}

func (r_ RegularExpression) Pattern() string {
	rv := objc.CallMethod[string](r_, objc.SEL("pattern"))
	return rv
}

func (r_ RegularExpression) Options() RegularExpressionOptions {
	rv := objc.CallMethod[RegularExpressionOptions](r_, objc.SEL("options"))
	return rv
}

func (r_ RegularExpression) NumberOfCaptureGroups() uint {
	rv := objc.CallMethod[uint](r_, objc.SEL("numberOfCaptureGroups"))
	return rv
}
