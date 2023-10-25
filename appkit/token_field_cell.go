// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TokenFieldCellClass = _TokenFieldCellClass{objc.GetClass("NSTokenFieldCell")}

type _TokenFieldCellClass struct {
	objc.Class
}

type ITokenFieldCell interface {
	ITextFieldCell
	TokenStyle() TokenStyle
	SetTokenStyle(value TokenStyle)
	TokenizingCharacterSet() foundation.CharacterSet
	SetTokenizingCharacterSet(value foundation.ICharacterSet)
	CompletionDelay() foundation.TimeInterval
	SetCompletionDelay(value foundation.TimeInterval)
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
}

type TokenFieldCell struct {
	TextFieldCell
}

func MakeTokenFieldCell(ptr unsafe.Pointer) TokenFieldCell {
	return TokenFieldCell{
		TextFieldCell: MakeTextFieldCell(ptr),
	}
}

func (t_ TokenFieldCell) InitTextCell(string_ string) TokenFieldCell {
	rv := objc.CallMethod[TokenFieldCell](t_, objc.SEL("initTextCell:"), string_)
	return rv
}

func (t_ TokenFieldCell) InitImageCell(image IImage) TokenFieldCell {
	rv := objc.CallMethod[TokenFieldCell](t_, objc.SEL("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func (t_ TokenFieldCell) Init() TokenFieldCell {
	rv := objc.CallMethod[TokenFieldCell](t_, objc.SEL("init"))
	return rv
}

func (tc _TokenFieldCellClass) Alloc() TokenFieldCell {
	rv := objc.CallMethod[TokenFieldCell](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TokenFieldCellClass) New() TokenFieldCell {
	rv := objc.CallMethod[TokenFieldCell](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTokenFieldCell() TokenFieldCell {
	return TokenFieldCellClass.New()
}

func (t_ TokenFieldCell) TokenStyle() TokenStyle {
	rv := objc.CallMethod[TokenStyle](t_, objc.SEL("tokenStyle"))
	return rv
}

func (t_ TokenFieldCell) SetTokenStyle(value TokenStyle) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTokenStyle:"), value)
}

func (tc _TokenFieldCellClass) DefaultTokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.CallMethod[foundation.CharacterSet](tc, objc.SEL("defaultTokenizingCharacterSet"))
	return rv
}

func (t_ TokenFieldCell) TokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.CallMethod[foundation.CharacterSet](t_, objc.SEL("tokenizingCharacterSet"))
	return rv
}

func (t_ TokenFieldCell) SetTokenizingCharacterSet(value foundation.ICharacterSet) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTokenizingCharacterSet:"), objc.ExtractPtr(value))
}

func (t_ TokenFieldCell) CompletionDelay() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](t_, objc.SEL("completionDelay"))
	return rv
}

func (t_ TokenFieldCell) SetCompletionDelay(value foundation.TimeInterval) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setCompletionDelay:"), value)
}

func (tc _TokenFieldCellClass) DefaultCompletionDelay() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](tc, objc.SEL("defaultCompletionDelay"))
	return rv
}

// weak property
func (t_ TokenFieldCell) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.SEL("delegate"))
	return rv
}

// weak property
func (t_ TokenFieldCell) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}
