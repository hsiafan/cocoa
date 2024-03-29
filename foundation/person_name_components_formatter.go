// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var PersonNameComponentsFormatterClass = _PersonNameComponentsFormatterClass{objc.GetClass("NSPersonNameComponentsFormatter")}

type _PersonNameComponentsFormatterClass struct {
	objc.Class
}

type IPersonNameComponentsFormatter interface {
	IFormatter
	StringFromPersonNameComponents(components IPersonNameComponents) string
	AnnotatedStringFromPersonNameComponents(components IPersonNameComponents) AttributedString
	PersonNameComponentsFromString(string_ string) PersonNameComponents
	Style() PersonNameComponentsFormatterStyle
	SetStyle(value PersonNameComponentsFormatterStyle)
	IsPhonetic() bool
	SetPhonetic(value bool)
	Locale() Locale
	SetLocale(value ILocale)
}

type PersonNameComponentsFormatter struct {
	Formatter
}

func MakePersonNameComponentsFormatter(ptr unsafe.Pointer) PersonNameComponentsFormatter {
	return PersonNameComponentsFormatter{
		Formatter: MakeFormatter(ptr),
	}
}

func (pc _PersonNameComponentsFormatterClass) Alloc() PersonNameComponentsFormatter {
	rv := objc.CallMethod[PersonNameComponentsFormatter](pc, objc.SEL("alloc"))
	return rv
}

func (pc _PersonNameComponentsFormatterClass) New() PersonNameComponentsFormatter {
	rv := objc.CallMethod[PersonNameComponentsFormatter](pc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewPersonNameComponentsFormatter() PersonNameComponentsFormatter {
	return PersonNameComponentsFormatterClass.New()
}

func (p_ PersonNameComponentsFormatter) Init() PersonNameComponentsFormatter {
	rv := objc.CallMethod[PersonNameComponentsFormatter](p_, objc.SEL("init"))
	return rv
}

func (pc _PersonNameComponentsFormatterClass) LocalizedStringFromPersonNameComponents_Style_Options(components IPersonNameComponents, nameFormatStyle PersonNameComponentsFormatterStyle, nameOptions PersonNameComponentsFormatterOptions) string {
	rv := objc.CallMethod[string](pc, objc.SEL("localizedStringFromPersonNameComponents:style:options:"), objc.ExtractPtr(components), nameFormatStyle, nameOptions)
	return rv
}

func (p_ PersonNameComponentsFormatter) StringFromPersonNameComponents(components IPersonNameComponents) string {
	rv := objc.CallMethod[string](p_, objc.SEL("stringFromPersonNameComponents:"), objc.ExtractPtr(components))
	return rv
}

func (p_ PersonNameComponentsFormatter) AnnotatedStringFromPersonNameComponents(components IPersonNameComponents) AttributedString {
	rv := objc.CallMethod[AttributedString](p_, objc.SEL("annotatedStringFromPersonNameComponents:"), objc.ExtractPtr(components))
	return rv
}

func (p_ PersonNameComponentsFormatter) PersonNameComponentsFromString(string_ string) PersonNameComponents {
	rv := objc.CallMethod[PersonNameComponents](p_, objc.SEL("personNameComponentsFromString:"), string_)
	return rv
}

func (p_ PersonNameComponentsFormatter) Style() PersonNameComponentsFormatterStyle {
	rv := objc.CallMethod[PersonNameComponentsFormatterStyle](p_, objc.SEL("style"))
	return rv
}

func (p_ PersonNameComponentsFormatter) SetStyle(value PersonNameComponentsFormatterStyle) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setStyle:"), value)
}

func (p_ PersonNameComponentsFormatter) IsPhonetic() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isPhonetic"))
	return rv
}

func (p_ PersonNameComponentsFormatter) SetPhonetic(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setPhonetic:"), value)
}

func (p_ PersonNameComponentsFormatter) Locale() Locale {
	rv := objc.CallMethod[Locale](p_, objc.SEL("locale"))
	return rv
}

func (p_ PersonNameComponentsFormatter) SetLocale(value ILocale) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setLocale:"), objc.ExtractPtr(value))
}
