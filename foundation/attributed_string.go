// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var AttributedStringClass = _AttributedStringClass{objc.GetClass("NSAttributedString")}

type _AttributedStringClass struct {
	objc.Class
}

type IAttributedString interface {
	objc.IObject
	AttributesAtIndex_EffectiveRange(location uint, range_ *Range) map[AttributedStringKey]objc.Object
	AttributesAtIndex_LongestEffectiveRange_InRange(location uint, range_ *Range, rangeLimit Range) map[AttributedStringKey]objc.Object
	Attribute_AtIndex_EffectiveRange(attrName AttributedStringKey, location uint, range_ *Range) objc.Object
	Attribute_AtIndex_LongestEffectiveRange_InRange(attrName AttributedStringKey, location uint, range_ *Range, rangeLimit Range) objc.Object
	IsEqualToAttributedString(other IAttributedString) bool
	AttributedSubstringFromRange(range_ Range) AttributedString
	EnumerateAttribute_InRange_Options_UsingBlock(attrName AttributedStringKey, enumerationRange Range, opts AttributedStringEnumerationOptions, block func(value objc.Object, range_ Range, stop *bool))
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
	rv := objc.CallMethod[AttributedString](a_, "initWithString:", str)
	return rv
}

func (a_ AttributedString) InitWithString_Attributes(str string, attrs map[AttributedStringKey]objc.IObject) AttributedString {
	rv := objc.CallMethod[AttributedString](a_, "initWithString:attributes:", str, attrs)
	return rv
}

func (a_ AttributedString) InitWithAttributedString(attrStr IAttributedString) AttributedString {
	rv := objc.CallMethod[AttributedString](a_, "initWithAttributedString:", attrStr)
	return rv
}

func (ac _AttributedStringClass) Alloc() AttributedString {
	rv := objc.CallMethod[AttributedString](ac, "alloc")
	return rv
}

func (ac _AttributedStringClass) New() AttributedString {
	rv := objc.CallMethod[AttributedString](ac, "new")
	rv.Autorelease()
	return rv
}

func NewAttributedString() AttributedString {
	return AttributedStringClass.New()
}

func (a_ AttributedString) Init() AttributedString {
	rv := objc.CallMethod[AttributedString](a_, "init")
	return rv
}

func (a_ AttributedString) AttributesAtIndex_EffectiveRange(location uint, range_ *Range) map[AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[AttributedStringKey]objc.Object](a_, "attributesAtIndex:effectiveRange:", location, range_)
	return rv
}

func (a_ AttributedString) AttributesAtIndex_LongestEffectiveRange_InRange(location uint, range_ *Range, rangeLimit Range) map[AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[AttributedStringKey]objc.Object](a_, "attributesAtIndex:longestEffectiveRange:inRange:", location, range_, rangeLimit)
	return rv
}

func (a_ AttributedString) Attribute_AtIndex_EffectiveRange(attrName AttributedStringKey, location uint, range_ *Range) objc.Object {
	rv := objc.CallMethod[objc.Object](a_, "attribute:atIndex:effectiveRange:", attrName, location, range_)
	return rv
}

func (a_ AttributedString) Attribute_AtIndex_LongestEffectiveRange_InRange(attrName AttributedStringKey, location uint, range_ *Range, rangeLimit Range) objc.Object {
	rv := objc.CallMethod[objc.Object](a_, "attribute:atIndex:longestEffectiveRange:inRange:", attrName, location, range_, rangeLimit)
	return rv
}

func (a_ AttributedString) IsEqualToAttributedString(other IAttributedString) bool {
	rv := objc.CallMethod[bool](a_, "isEqualToAttributedString:", other)
	return rv
}

func (a_ AttributedString) AttributedSubstringFromRange(range_ Range) AttributedString {
	rv := objc.CallMethod[AttributedString](a_, "attributedSubstringFromRange:", range_)
	return rv
}

func (a_ AttributedString) EnumerateAttribute_InRange_Options_UsingBlock(attrName AttributedStringKey, enumerationRange Range, opts AttributedStringEnumerationOptions, block func(value objc.Object, range_ Range, stop *bool)) {
	objc.CallMethod[objc.Void](a_, "enumerateAttribute:inRange:options:usingBlock:", attrName, enumerationRange, opts, block)
}

func (a_ AttributedString) AttributedStringByInflectingString() AttributedString {
	rv := objc.CallMethod[AttributedString](a_, "attributedStringByInflectingString")
	return rv
}

func (a_ AttributedString) String() string {
	rv := objc.CallMethod[string](a_, "string")
	return rv
}

func (a_ AttributedString) Length() uint {
	rv := objc.CallMethod[uint](a_, "length")
	return rv
}
