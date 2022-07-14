// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var RegularExpressionClass = _RegularExpressionClass{objc.GetClass("NSRegularExpression")}

type _RegularExpressionClass struct {
	objc.Class
}

type IRegularExpression interface {
	objc.IObject
	NumberOfMatchesInString_Options_Range(_string string, options MatchingOptions, _range Range) uint
	EnumerateMatchesInString_Options_Range_UsingBlock(_string string, options MatchingOptions, _range Range, block func(result TextCheckingResult, flags MatchingFlags, stop *bool))
	MatchesInString_Options_Range(_string string, options MatchingOptions, _range Range) []TextCheckingResult
	FirstMatchInString_Options_Range(_string string, options MatchingOptions, _range Range) TextCheckingResult
	RangeOfFirstMatchInString_Options_Range(_string string, options MatchingOptions, _range Range) Range
	ReplaceMatchesInString_Options_Range_WithTemplate(_string IMutableString, options MatchingOptions, _range Range, templ string) uint
	StringByReplacingMatchesInString_Options_Range_WithTemplate(_string string, options MatchingOptions, _range Range, templ string) string
	ReplacementStringForResult_InString_Offset_Template(result ITextCheckingResult, _string string, offset int, templ string) string
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
	rv := ffi.CallMethod[RegularExpression](r_, "initWithPattern:options:error:", pattern, options, error)
	rv.Autorelease()
	return rv
}

func (rc _RegularExpressionClass) Alloc() RegularExpression {
	rv := ffi.CallMethod[RegularExpression](rc, "alloc")
	return rv
}

func (r_ RegularExpression) Init() RegularExpression {
	rv := ffi.CallMethod[RegularExpression](r_, "init")
	rv.Autorelease()
	return rv
}

func (rc _RegularExpressionClass) New() RegularExpression {
	rv := ffi.CallMethod[RegularExpression](rc, "new")
	rv.Autorelease()
	return rv
}

func NewRegularExpression() RegularExpression {
	return RegularExpressionClass.New()
}

func (rc _RegularExpressionClass) RegularExpressionWithPattern_Options_Error(pattern string, options RegularExpressionOptions, error *Error) RegularExpression {
	rv := ffi.CallMethod[RegularExpression](rc, "regularExpressionWithPattern:options:error:", pattern, options, error)
	return rv
}

func (r_ RegularExpression) NumberOfMatchesInString_Options_Range(_string string, options MatchingOptions, _range Range) uint {
	rv := ffi.CallMethod[uint](r_, "numberOfMatchesInString:options:range:", _string, options, _range)
	return rv
}

func (r_ RegularExpression) EnumerateMatchesInString_Options_Range_UsingBlock(_string string, options MatchingOptions, _range Range, block func(result TextCheckingResult, flags MatchingFlags, stop *bool)) {
	ffi.CallMethod[ffi.Void](r_, "enumerateMatchesInString:options:range:usingBlock:", _string, options, _range, block)
}

func (r_ RegularExpression) MatchesInString_Options_Range(_string string, options MatchingOptions, _range Range) []TextCheckingResult {
	rv := ffi.CallMethod[[]TextCheckingResult](r_, "matchesInString:options:range:", _string, options, _range)
	return rv
}

func (r_ RegularExpression) FirstMatchInString_Options_Range(_string string, options MatchingOptions, _range Range) TextCheckingResult {
	rv := ffi.CallMethod[TextCheckingResult](r_, "firstMatchInString:options:range:", _string, options, _range)
	return rv
}

func (r_ RegularExpression) RangeOfFirstMatchInString_Options_Range(_string string, options MatchingOptions, _range Range) Range {
	rv := ffi.CallMethod[Range](r_, "rangeOfFirstMatchInString:options:range:", _string, options, _range)
	return rv
}

func (r_ RegularExpression) ReplaceMatchesInString_Options_Range_WithTemplate(_string IMutableString, options MatchingOptions, _range Range, templ string) uint {
	rv := ffi.CallMethod[uint](r_, "replaceMatchesInString:options:range:withTemplate:", _string, options, _range, templ)
	return rv
}

func (r_ RegularExpression) StringByReplacingMatchesInString_Options_Range_WithTemplate(_string string, options MatchingOptions, _range Range, templ string) string {
	rv := ffi.CallMethod[string](r_, "stringByReplacingMatchesInString:options:range:withTemplate:", _string, options, _range, templ)
	return rv
}

func (rc _RegularExpressionClass) EscapedTemplateForString(_string string) string {
	rv := ffi.CallMethod[string](rc, "escapedTemplateForString:", _string)
	return rv
}

func (rc _RegularExpressionClass) EscapedPatternForString(_string string) string {
	rv := ffi.CallMethod[string](rc, "escapedPatternForString:", _string)
	return rv
}

func (r_ RegularExpression) ReplacementStringForResult_InString_Offset_Template(result ITextCheckingResult, _string string, offset int, templ string) string {
	rv := ffi.CallMethod[string](r_, "replacementStringForResult:inString:offset:template:", result, _string, offset, templ)
	return rv
}

func (r_ RegularExpression) Pattern() string {
	rv := ffi.CallMethod[string](r_, "pattern")
	return rv
}

func (r_ RegularExpression) Options() RegularExpressionOptions {
	rv := ffi.CallMethod[RegularExpressionOptions](r_, "options")
	return rv
}

func (r_ RegularExpression) NumberOfCaptureGroups() uint {
	rv := ffi.CallMethod[uint](r_, "numberOfCaptureGroups")
	return rv
}
