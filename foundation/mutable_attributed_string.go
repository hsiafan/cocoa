// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var MutableAttributedStringClass = _MutableAttributedStringClass{objc.GetClass("NSMutableAttributedString")}

type _MutableAttributedStringClass struct {
	objc.Class
}

type IMutableAttributedString interface {
	IAttributedString
	ReplaceCharactersInRange_WithString(range_ Range, str string)
	DeleteCharactersInRange(range_ Range)
	SetAttributes_Range(attrs map[AttributedStringKey]objc.IObject, range_ Range)
	AddAttribute_Value_Range(name AttributedStringKey, value objc.IObject, range_ Range)
	AddAttributes_Range(attrs map[AttributedStringKey]objc.IObject, range_ Range)
	RemoveAttribute_Range(name AttributedStringKey, range_ Range)
	AppendAttributedString(attrString IAttributedString)
	InsertAttributedString_AtIndex(attrString IAttributedString, loc uint)
	ReplaceCharactersInRange_WithAttributedString(range_ Range, attrString IAttributedString)
	SetAttributedString(attrString IAttributedString)
	BeginEditing()
	EndEditing()
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
	rv := objc.CallMethod[MutableAttributedString](m_, objc.SEL("initWithString:"), str)
	return rv
}

func (m_ MutableAttributedString) InitWithString_Attributes(str string, attrs map[AttributedStringKey]objc.IObject) MutableAttributedString {
	rv := objc.CallMethod[MutableAttributedString](m_, objc.SEL("initWithString:attributes:"), str, attrs)
	return rv
}

func (m_ MutableAttributedString) InitWithAttributedString(attrStr IAttributedString) MutableAttributedString {
	rv := objc.CallMethod[MutableAttributedString](m_, objc.SEL("initWithAttributedString:"), objc.ExtractPtr(attrStr))
	return rv
}

func (mc _MutableAttributedStringClass) Alloc() MutableAttributedString {
	rv := objc.CallMethod[MutableAttributedString](mc, objc.SEL("alloc"))
	return rv
}

func (mc _MutableAttributedStringClass) New() MutableAttributedString {
	rv := objc.CallMethod[MutableAttributedString](mc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewMutableAttributedString() MutableAttributedString {
	return MutableAttributedStringClass.New()
}

func (m_ MutableAttributedString) Init() MutableAttributedString {
	rv := objc.CallMethod[MutableAttributedString](m_, objc.SEL("init"))
	return rv
}

func (m_ MutableAttributedString) ReplaceCharactersInRange_WithString(range_ Range, str string) {
	objc.CallMethod[objc.Void](m_, objc.SEL("replaceCharactersInRange:withString:"), range_, str)
}

func (m_ MutableAttributedString) DeleteCharactersInRange(range_ Range) {
	objc.CallMethod[objc.Void](m_, objc.SEL("deleteCharactersInRange:"), range_)
}

func (m_ MutableAttributedString) SetAttributes_Range(attrs map[AttributedStringKey]objc.IObject, range_ Range) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setAttributes:range:"), attrs, range_)
}

func (m_ MutableAttributedString) AddAttribute_Value_Range(name AttributedStringKey, value objc.IObject, range_ Range) {
	objc.CallMethod[objc.Void](m_, objc.SEL("addAttribute:value:range:"), name, objc.ExtractPtr(value), range_)
}

func (m_ MutableAttributedString) AddAttributes_Range(attrs map[AttributedStringKey]objc.IObject, range_ Range) {
	objc.CallMethod[objc.Void](m_, objc.SEL("addAttributes:range:"), attrs, range_)
}

func (m_ MutableAttributedString) RemoveAttribute_Range(name AttributedStringKey, range_ Range) {
	objc.CallMethod[objc.Void](m_, objc.SEL("removeAttribute:range:"), name, range_)
}

func (m_ MutableAttributedString) AppendAttributedString(attrString IAttributedString) {
	objc.CallMethod[objc.Void](m_, objc.SEL("appendAttributedString:"), objc.ExtractPtr(attrString))
}

func (m_ MutableAttributedString) InsertAttributedString_AtIndex(attrString IAttributedString, loc uint) {
	objc.CallMethod[objc.Void](m_, objc.SEL("insertAttributedString:atIndex:"), objc.ExtractPtr(attrString), loc)
}

func (m_ MutableAttributedString) ReplaceCharactersInRange_WithAttributedString(range_ Range, attrString IAttributedString) {
	objc.CallMethod[objc.Void](m_, objc.SEL("replaceCharactersInRange:withAttributedString:"), range_, objc.ExtractPtr(attrString))
}

func (m_ MutableAttributedString) SetAttributedString(attrString IAttributedString) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setAttributedString:"), objc.ExtractPtr(attrString))
}

func (m_ MutableAttributedString) BeginEditing() {
	objc.CallMethod[objc.Void](m_, objc.SEL("beginEditing"))
}

func (m_ MutableAttributedString) EndEditing() {
	objc.CallMethod[objc.Void](m_, objc.SEL("endEditing"))
}
