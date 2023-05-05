// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TokenFieldClass = _TokenFieldClass{objc.GetClass("NSTokenField")}

type _TokenFieldClass struct {
	objc.Class
}

type ITokenField interface {
	ITextField
	TokenStyle() TokenStyle
	SetTokenStyle(value TokenStyle)
	TokenizingCharacterSet() foundation.CharacterSet
	SetTokenizingCharacterSet(value foundation.ICharacterSet)
	CompletionDelay() foundation.TimeInterval
	SetCompletionDelay(value foundation.TimeInterval)
}

type TokenField struct {
	TextField
}

func MakeTokenField(ptr unsafe.Pointer) TokenField {
	return TokenField{
		TextField: MakeTextField(ptr),
	}
}

func (tc _TokenFieldClass) LabelWithAttributedString(attributedStringValue foundation.IAttributedString) TokenField {
	rv := objc.CallMethod[TokenField](tc, objc.GetSelector("labelWithAttributedString:"), objc.ExtractPtr(attributedStringValue))
	return rv
}

func (tc _TokenFieldClass) LabelWithString(stringValue string) TokenField {
	rv := objc.CallMethod[TokenField](tc, objc.GetSelector("labelWithString:"), stringValue)
	return rv
}

func (tc _TokenFieldClass) TextFieldWithString(stringValue string) TokenField {
	rv := objc.CallMethod[TokenField](tc, objc.GetSelector("textFieldWithString:"), stringValue)
	return rv
}

func (tc _TokenFieldClass) WrappingLabelWithString(stringValue string) TokenField {
	rv := objc.CallMethod[TokenField](tc, objc.GetSelector("wrappingLabelWithString:"), stringValue)
	return rv
}

func (t_ TokenField) InitWithFrame(frameRect foundation.Rect) TokenField {
	rv := objc.CallMethod[TokenField](t_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func (t_ TokenField) Init() TokenField {
	rv := objc.CallMethod[TokenField](t_, objc.GetSelector("init"))
	return rv
}

func (tc _TokenFieldClass) Alloc() TokenField {
	rv := objc.CallMethod[TokenField](tc, objc.GetSelector("alloc"))
	return rv
}

func (tc _TokenFieldClass) New() TokenField {
	rv := objc.CallMethod[TokenField](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTokenField() TokenField {
	return TokenFieldClass.New()
}

func (t_ TokenField) TokenStyle() TokenStyle {
	rv := objc.CallMethod[TokenStyle](t_, objc.GetSelector("tokenStyle"))
	return rv
}

func (t_ TokenField) SetTokenStyle(value TokenStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTokenStyle:"), value)
}

func (t_ TokenField) TokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.CallMethod[foundation.CharacterSet](t_, objc.GetSelector("tokenizingCharacterSet"))
	return rv
}

func (t_ TokenField) SetTokenizingCharacterSet(value foundation.ICharacterSet) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTokenizingCharacterSet:"), objc.ExtractPtr(value))
}

func (tc _TokenFieldClass) DefaultTokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.CallMethod[foundation.CharacterSet](tc, objc.GetSelector("defaultTokenizingCharacterSet"))
	return rv
}

func (t_ TokenField) CompletionDelay() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](t_, objc.GetSelector("completionDelay"))
	return rv
}

func (t_ TokenField) SetCompletionDelay(value foundation.TimeInterval) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setCompletionDelay:"), value)
}

func (tc _TokenFieldClass) DefaultCompletionDelay() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](tc, objc.GetSelector("defaultCompletionDelay"))
	return rv
}
