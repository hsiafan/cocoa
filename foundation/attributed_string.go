// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var AttributedStringClass = _AttributedStringClass{objc.GetClass("NSAttributedString")}

type _AttributedStringClass struct {
	objc.Class
}

type IAttributedString interface {
	objc.IObject
	AttributesAtIndex_EffectiveRange(location uint, _range *Range) map[AttributedStringKey]objc.Object
	AttributesAtIndex_LongestEffectiveRange_InRange(location uint, _range *Range, rangeLimit Range) map[AttributedStringKey]objc.Object
	Attribute_AtIndex_EffectiveRange(attrName AttributedStringKey, location uint, _range *Range) objc.Object
	Attribute_AtIndex_LongestEffectiveRange_InRange(attrName AttributedStringKey, location uint, _range *Range, rangeLimit Range) objc.Object
	IsEqualToAttributedString(other IAttributedString) bool
	AttributedSubstringFromRange(_range Range) AttributedString
	EnumerateAttribute_InRange_Options_UsingBlock(attrName AttributedStringKey, enumerationRange Range, opts AttributedStringEnumerationOptions, block func(value objc.Object, _range Range, stop *bool))
	EnumerateAttributesInRange_Options_UsingBlock(enumerationRange Range, opts AttributedStringEnumerationOptions, block func(attrs map[AttributedStringKey]objc.Object, _range Range, stop *bool))
	AttributedStringByInflectingString() AttributedString
	String() string
	Length() uint
}

type AttributedString struct {
	objc.Object
}

func MakeAttributedString(ptr unsafe.Pointer) AttributedString {
	return AttributedString{
		Object: objc.MakeObject(ptr),
	}
}

func (a_ AttributedString) InitWithString(str string) AttributedString {
	rv := ffi.CallMethod[AttributedString](a_, "initWithString:", str)
	rv.Autorelease()
	return rv
}

func (a_ AttributedString) InitWithString_Attributes(str string, attrs map[AttributedStringKey]objc.IObject) AttributedString {
	rv := ffi.CallMethod[AttributedString](a_, "initWithString:attributes:", str, attrs)
	rv.Autorelease()
	return rv
}

func (a_ AttributedString) InitWithAttributedString(attrStr IAttributedString) AttributedString {
	rv := ffi.CallMethod[AttributedString](a_, "initWithAttributedString:", attrStr)
	rv.Autorelease()
	return rv
}

func (ac _AttributedStringClass) Alloc() AttributedString {
	rv := ffi.CallMethod[AttributedString](ac, "alloc")
	return rv
}

func (a_ AttributedString) Init() AttributedString {
	rv := ffi.CallMethod[AttributedString](a_, "init")
	rv.Autorelease()
	return rv
}

func (ac _AttributedStringClass) New() AttributedString {
	rv := ffi.CallMethod[AttributedString](ac, "new")
	rv.Autorelease()
	return rv
}

func NewAttributedString() AttributedString {
	return AttributedStringClass.New()
}

func (a_ AttributedString) AttributesAtIndex_EffectiveRange(location uint, _range *Range) map[AttributedStringKey]objc.Object {
	rv := ffi.CallMethod[map[AttributedStringKey]objc.Object](a_, "attributesAtIndex:effectiveRange:", location, _range)
	return rv
}

func (a_ AttributedString) AttributesAtIndex_LongestEffectiveRange_InRange(location uint, _range *Range, rangeLimit Range) map[AttributedStringKey]objc.Object {
	rv := ffi.CallMethod[map[AttributedStringKey]objc.Object](a_, "attributesAtIndex:longestEffectiveRange:inRange:", location, _range, rangeLimit)
	return rv
}

func (a_ AttributedString) Attribute_AtIndex_EffectiveRange(attrName AttributedStringKey, location uint, _range *Range) objc.Object {
	rv := ffi.CallMethod[objc.Object](a_, "attribute:atIndex:effectiveRange:", attrName, location, _range)
	return rv
}

func (a_ AttributedString) Attribute_AtIndex_LongestEffectiveRange_InRange(attrName AttributedStringKey, location uint, _range *Range, rangeLimit Range) objc.Object {
	rv := ffi.CallMethod[objc.Object](a_, "attribute:atIndex:longestEffectiveRange:inRange:", attrName, location, _range, rangeLimit)
	return rv
}

func (a_ AttributedString) IsEqualToAttributedString(other IAttributedString) bool {
	rv := ffi.CallMethod[bool](a_, "isEqualToAttributedString:", other)
	return rv
}

func (a_ AttributedString) AttributedSubstringFromRange(_range Range) AttributedString {
	rv := ffi.CallMethod[AttributedString](a_, "attributedSubstringFromRange:", _range)
	return rv
}

func (a_ AttributedString) EnumerateAttribute_InRange_Options_UsingBlock(attrName AttributedStringKey, enumerationRange Range, opts AttributedStringEnumerationOptions, block func(value objc.Object, _range Range, stop *bool)) {
	ffi.CallMethod[ffi.Void](a_, "enumerateAttribute:inRange:options:usingBlock:", attrName, enumerationRange, opts, block)
}

func (a_ AttributedString) EnumerateAttributesInRange_Options_UsingBlock(enumerationRange Range, opts AttributedStringEnumerationOptions, block func(attrs map[AttributedStringKey]objc.Object, _range Range, stop *bool)) {
	ffi.CallMethod[ffi.Void](a_, "enumerateAttributesInRange:options:usingBlock:", enumerationRange, opts, block)
}

func (a_ AttributedString) AttributedStringByInflectingString() AttributedString {
	rv := ffi.CallMethod[AttributedString](a_, "attributedStringByInflectingString")
	return rv
}

func (a_ AttributedString) String() string {
	rv := ffi.CallMethod[string](a_, "string")
	return rv
}

func (a_ AttributedString) Length() uint {
	rv := ffi.CallMethod[uint](a_, "length")
	return rv
}

var MutableAttributedStringClass = _MutableAttributedStringClass{objc.GetClass("NSMutableAttributedString")}

type _MutableAttributedStringClass struct {
	objc.Class
}

type IMutableAttributedString interface {
	IAttributedString
	ReplaceCharactersInRange_WithString(_range Range, str string)
	DeleteCharactersInRange(_range Range)
	SetAttributes_Range(attrs map[AttributedStringKey]objc.IObject, _range Range)
	AddAttribute_Value_Range(name AttributedStringKey, value objc.IObject, _range Range)
	AddAttributes_Range(attrs map[AttributedStringKey]objc.IObject, _range Range)
	RemoveAttribute_Range(name AttributedStringKey, _range Range)
	AppendAttributedString(attrString IAttributedString)
	InsertAttributedString_AtIndex(attrString IAttributedString, loc uint)
	ReplaceCharactersInRange_WithAttributedString(_range Range, attrString IAttributedString)
	SetAttributedString(attrString IAttributedString)
	BeginEditing()
	EndEditing()
	MutableString() MutableString
}

type MutableAttributedString struct {
	AttributedString
}

func MakeMutableAttributedString(ptr unsafe.Pointer) MutableAttributedString {
	return MutableAttributedString{
		AttributedString: MakeAttributedString(ptr),
	}
}

func (m_ MutableAttributedString) InitWithString(str string) MutableAttributedString {
	rv := ffi.CallMethod[MutableAttributedString](m_, "initWithString:", str)
	rv.Autorelease()
	return rv
}

func (m_ MutableAttributedString) InitWithString_Attributes(str string, attrs map[AttributedStringKey]objc.IObject) MutableAttributedString {
	rv := ffi.CallMethod[MutableAttributedString](m_, "initWithString:attributes:", str, attrs)
	rv.Autorelease()
	return rv
}

func (m_ MutableAttributedString) InitWithAttributedString(attrStr IAttributedString) MutableAttributedString {
	rv := ffi.CallMethod[MutableAttributedString](m_, "initWithAttributedString:", attrStr)
	rv.Autorelease()
	return rv
}

func (mc _MutableAttributedStringClass) Alloc() MutableAttributedString {
	rv := ffi.CallMethod[MutableAttributedString](mc, "alloc")
	return rv
}

func (m_ MutableAttributedString) Init() MutableAttributedString {
	rv := ffi.CallMethod[MutableAttributedString](m_, "init")
	rv.Autorelease()
	return rv
}

func (mc _MutableAttributedStringClass) New() MutableAttributedString {
	rv := ffi.CallMethod[MutableAttributedString](mc, "new")
	rv.Autorelease()
	return rv
}

func NewMutableAttributedString() MutableAttributedString {
	return MutableAttributedStringClass.New()
}

func (m_ MutableAttributedString) ReplaceCharactersInRange_WithString(_range Range, str string) {
	ffi.CallMethod[ffi.Void](m_, "replaceCharactersInRange:withString:", _range, str)
}

func (m_ MutableAttributedString) DeleteCharactersInRange(_range Range) {
	ffi.CallMethod[ffi.Void](m_, "deleteCharactersInRange:", _range)
}

func (m_ MutableAttributedString) SetAttributes_Range(attrs map[AttributedStringKey]objc.IObject, _range Range) {
	ffi.CallMethod[ffi.Void](m_, "setAttributes:range:", attrs, _range)
}

func (m_ MutableAttributedString) AddAttribute_Value_Range(name AttributedStringKey, value objc.IObject, _range Range) {
	ffi.CallMethod[ffi.Void](m_, "addAttribute:value:range:", name, value, _range)
}

func (m_ MutableAttributedString) AddAttributes_Range(attrs map[AttributedStringKey]objc.IObject, _range Range) {
	ffi.CallMethod[ffi.Void](m_, "addAttributes:range:", attrs, _range)
}

func (m_ MutableAttributedString) RemoveAttribute_Range(name AttributedStringKey, _range Range) {
	ffi.CallMethod[ffi.Void](m_, "removeAttribute:range:", name, _range)
}

func (m_ MutableAttributedString) AppendAttributedString(attrString IAttributedString) {
	ffi.CallMethod[ffi.Void](m_, "appendAttributedString:", attrString)
}

func (m_ MutableAttributedString) InsertAttributedString_AtIndex(attrString IAttributedString, loc uint) {
	ffi.CallMethod[ffi.Void](m_, "insertAttributedString:atIndex:", attrString, loc)
}

func (m_ MutableAttributedString) ReplaceCharactersInRange_WithAttributedString(_range Range, attrString IAttributedString) {
	ffi.CallMethod[ffi.Void](m_, "replaceCharactersInRange:withAttributedString:", _range, attrString)
}

func (m_ MutableAttributedString) SetAttributedString(attrString IAttributedString) {
	ffi.CallMethod[ffi.Void](m_, "setAttributedString:", attrString)
}

func (m_ MutableAttributedString) BeginEditing() {
	ffi.CallMethod[ffi.Void](m_, "beginEditing")
}

func (m_ MutableAttributedString) EndEditing() {
	ffi.CallMethod[ffi.Void](m_, "endEditing")
}

func (m_ MutableAttributedString) MutableString() MutableString {
	rv := ffi.CallMethod[MutableString](m_, "mutableString")
	return rv
}
