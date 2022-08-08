// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var PersonNameComponentsClass = _PersonNameComponentsClass{objc.GetClass("NSPersonNameComponents")}

type _PersonNameComponentsClass struct {
	objc.Class
}

type IPersonNameComponents interface {
	objc.IObject
	NamePrefix() string
	SetNamePrefix(value string)
	GivenName() string
	SetGivenName(value string)
	MiddleName() string
	SetMiddleName(value string)
	FamilyName() string
	SetFamilyName(value string)
	NameSuffix() string
	SetNameSuffix(value string)
	Nickname() string
	SetNickname(value string)
	PhoneticRepresentation() PersonNameComponents
	SetPhoneticRepresentation(value IPersonNameComponents)
}

type PersonNameComponents struct {
	objc.Object
}

func MakePersonNameComponents(ptr unsafe.Pointer) PersonNameComponents {
	return PersonNameComponents{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PersonNameComponentsClass) Alloc() PersonNameComponents {
	rv := ffi.CallMethod[PersonNameComponents](pc, "alloc")
	return rv
}

func (p_ PersonNameComponents) Init() PersonNameComponents {
	rv := ffi.CallMethod[PersonNameComponents](p_, "init")
	return rv
}

func (pc _PersonNameComponentsClass) New() PersonNameComponents {
	rv := ffi.CallMethod[PersonNameComponents](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPersonNameComponents() PersonNameComponents {
	return PersonNameComponentsClass.New()
}

func (p_ PersonNameComponents) NamePrefix() string {
	rv := ffi.CallMethod[string](p_, "namePrefix")
	return rv
}

func (p_ PersonNameComponents) SetNamePrefix(value string) {
	ffi.CallMethod[ffi.Void](p_, "setNamePrefix:", value)
}

func (p_ PersonNameComponents) GivenName() string {
	rv := ffi.CallMethod[string](p_, "givenName")
	return rv
}

func (p_ PersonNameComponents) SetGivenName(value string) {
	ffi.CallMethod[ffi.Void](p_, "setGivenName:", value)
}

func (p_ PersonNameComponents) MiddleName() string {
	rv := ffi.CallMethod[string](p_, "middleName")
	return rv
}

func (p_ PersonNameComponents) SetMiddleName(value string) {
	ffi.CallMethod[ffi.Void](p_, "setMiddleName:", value)
}

func (p_ PersonNameComponents) FamilyName() string {
	rv := ffi.CallMethod[string](p_, "familyName")
	return rv
}

func (p_ PersonNameComponents) SetFamilyName(value string) {
	ffi.CallMethod[ffi.Void](p_, "setFamilyName:", value)
}

func (p_ PersonNameComponents) NameSuffix() string {
	rv := ffi.CallMethod[string](p_, "nameSuffix")
	return rv
}

func (p_ PersonNameComponents) SetNameSuffix(value string) {
	ffi.CallMethod[ffi.Void](p_, "setNameSuffix:", value)
}

func (p_ PersonNameComponents) Nickname() string {
	rv := ffi.CallMethod[string](p_, "nickname")
	return rv
}

func (p_ PersonNameComponents) SetNickname(value string) {
	ffi.CallMethod[ffi.Void](p_, "setNickname:", value)
}

func (p_ PersonNameComponents) PhoneticRepresentation() PersonNameComponents {
	rv := ffi.CallMethod[PersonNameComponents](p_, "phoneticRepresentation")
	return rv
}

func (p_ PersonNameComponents) SetPhoneticRepresentation(value IPersonNameComponents) {
	ffi.CallMethod[ffi.Void](p_, "setPhoneticRepresentation:", value)
}

var PersonNameComponentsFormatterClass = _PersonNameComponentsFormatterClass{objc.GetClass("NSPersonNameComponentsFormatter")}

type _PersonNameComponentsFormatterClass struct {
	objc.Class
}

type IPersonNameComponentsFormatter interface {
	IFormatter
	StringFromPersonNameComponents(components IPersonNameComponents) string
	AnnotatedStringFromPersonNameComponents(components IPersonNameComponents) AttributedString
	PersonNameComponentsFromString(_string string) PersonNameComponents
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

func (p_ PersonNameComponentsFormatter) Init() PersonNameComponentsFormatter {
	rv := ffi.CallMethod[PersonNameComponentsFormatter](p_, "init")
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

func (p_ PersonNameComponentsFormatter) PersonNameComponentsFromString(_string string) PersonNameComponents {
	rv := ffi.CallMethod[PersonNameComponents](p_, "personNameComponentsFromString:", _string)
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
