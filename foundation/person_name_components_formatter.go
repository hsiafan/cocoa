// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[PersonNameComponentsFormatter](pc, "alloc")
	return rv
}

func (pc _PersonNameComponentsFormatterClass) New() PersonNameComponentsFormatter {
	rv := ffi.CallMethod[PersonNameComponentsFormatter](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPersonNameComponentsFormatter() PersonNameComponentsFormatter {
	return PersonNameComponentsFormatterClass.New()
}

func (p_ PersonNameComponentsFormatter) Init() PersonNameComponentsFormatter {
	rv := ffi.CallMethod[PersonNameComponentsFormatter](p_, "init")
	return rv
}

func (pc _PersonNameComponentsFormatterClass) LocalizedStringFromPersonNameComponents_Style_Options(components IPersonNameComponents, nameFormatStyle PersonNameComponentsFormatterStyle, nameOptions PersonNameComponentsFormatterOptions) string {
	rv := ffi.CallMethod[string](pc, "localizedStringFromPersonNameComponents:style:options:", components, nameFormatStyle, nameOptions)
	return rv
}

func (p_ PersonNameComponentsFormatter) StringFromPersonNameComponents(components IPersonNameComponents) string {
	rv := ffi.CallMethod[string](p_, "stringFromPersonNameComponents:", components)
	return rv
}

func (p_ PersonNameComponentsFormatter) AnnotatedStringFromPersonNameComponents(components IPersonNameComponents) AttributedString {
	rv := ffi.CallMethod[AttributedString](p_, "annotatedStringFromPersonNameComponents:", components)
	return rv
}

func (p_ PersonNameComponentsFormatter) PersonNameComponentsFromString(string_ string) PersonNameComponents {
	rv := ffi.CallMethod[PersonNameComponents](p_, "personNameComponentsFromString:", string_)
	return rv
}

func (p_ PersonNameComponentsFormatter) Style() PersonNameComponentsFormatterStyle {
	rv := ffi.CallMethod[PersonNameComponentsFormatterStyle](p_, "style")
	return rv
}

func (p_ PersonNameComponentsFormatter) SetStyle(value PersonNameComponentsFormatterStyle) {
	ffi.CallMethod[ffi.Void](p_, "setStyle:", value)
}

func (p_ PersonNameComponentsFormatter) IsPhonetic() bool {
	rv := ffi.CallMethod[bool](p_, "isPhonetic")
	return rv
}

func (p_ PersonNameComponentsFormatter) SetPhonetic(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setPhonetic:", value)
}

func (p_ PersonNameComponentsFormatter) Locale() Locale {
	rv := ffi.CallMethod[Locale](p_, "locale")
	return rv
}

func (p_ PersonNameComponentsFormatter) SetLocale(value ILocale) {
	ffi.CallMethod[ffi.Void](p_, "setLocale:", value)
}
