// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
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
	Delegate() TokenFieldCellDelegateWrapper
	SetDelegate(value TokenFieldCellDelegate)
	SetDelegate0(value objc.IObject)
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
	rv := objc.CallMethod[TokenFieldCell](t_, objc.GetSelector("initTextCell:"), string_)
	return rv
}

func (t_ TokenFieldCell) InitImageCell(image IImage) TokenFieldCell {
	rv := objc.CallMethod[TokenFieldCell](t_, objc.GetSelector("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func (t_ TokenFieldCell) Init() TokenFieldCell {
	rv := objc.CallMethod[TokenFieldCell](t_, objc.GetSelector("init"))
	return rv
}

func (tc _TokenFieldCellClass) Alloc() TokenFieldCell {
	rv := objc.CallMethod[TokenFieldCell](tc, objc.GetSelector("alloc"))
	return rv
}

func (tc _TokenFieldCellClass) New() TokenFieldCell {
	rv := objc.CallMethod[TokenFieldCell](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTokenFieldCell() TokenFieldCell {
	return TokenFieldCellClass.New()
}

func (t_ TokenFieldCell) TokenStyle() TokenStyle {
	rv := objc.CallMethod[TokenStyle](t_, objc.GetSelector("tokenStyle"))
	return rv
}

func (t_ TokenFieldCell) SetTokenStyle(value TokenStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTokenStyle:"), value)
}

func (tc _TokenFieldCellClass) DefaultTokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.CallMethod[foundation.CharacterSet](tc, objc.GetSelector("defaultTokenizingCharacterSet"))
	return rv
}

func (t_ TokenFieldCell) TokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.CallMethod[foundation.CharacterSet](t_, objc.GetSelector("tokenizingCharacterSet"))
	return rv
}

func (t_ TokenFieldCell) SetTokenizingCharacterSet(value foundation.ICharacterSet) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTokenizingCharacterSet:"), objc.ExtractPtr(value))
}

func (t_ TokenFieldCell) CompletionDelay() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](t_, objc.GetSelector("completionDelay"))
	return rv
}

func (t_ TokenFieldCell) SetCompletionDelay(value foundation.TimeInterval) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setCompletionDelay:"), value)
}

func (tc _TokenFieldCellClass) DefaultCompletionDelay() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](tc, objc.GetSelector("defaultCompletionDelay"))
	return rv
}

func (t_ TokenFieldCell) Delegate() TokenFieldCellDelegateWrapper {
	rv := objc.CallMethod[TokenFieldCellDelegateWrapper](t_, objc.GetSelector("delegate"))
	return rv
}

func (t_ TokenFieldCell) SetDelegate(value TokenFieldCellDelegate) {
	po := objc.CreateProtocol("NSTokenFieldCellDelegate", value)
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDelegate:"), po)
}

func (t_ TokenFieldCell) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}
