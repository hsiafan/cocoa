// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[TokenFieldCell](t_, "initTextCell:", string_)
	return rv
}

func (t_ TokenFieldCell) InitImageCell(image IImage) TokenFieldCell {
	rv := ffi.CallMethod[TokenFieldCell](t_, "initImageCell:", image)
	return rv
}

func (t_ TokenFieldCell) Init() TokenFieldCell {
	rv := ffi.CallMethod[TokenFieldCell](t_, "init")
	return rv
}

func (tc _TokenFieldCellClass) Alloc() TokenFieldCell {
	rv := ffi.CallMethod[TokenFieldCell](tc, "alloc")
	return rv
}

func (tc _TokenFieldCellClass) New() TokenFieldCell {
	rv := ffi.CallMethod[TokenFieldCell](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTokenFieldCell() TokenFieldCell {
	return TokenFieldCellClass.New()
}

func (t_ TokenFieldCell) TokenStyle() TokenStyle {
	rv := ffi.CallMethod[TokenStyle](t_, "tokenStyle")
	return rv
}

func (t_ TokenFieldCell) SetTokenStyle(value TokenStyle) {
	ffi.CallMethod[ffi.Void](t_, "setTokenStyle:", value)
}

func (tc _TokenFieldCellClass) DefaultTokenizingCharacterSet() foundation.CharacterSet {
	rv := ffi.CallMethod[foundation.CharacterSet](tc, "defaultTokenizingCharacterSet")
	return rv
}

func (t_ TokenFieldCell) TokenizingCharacterSet() foundation.CharacterSet {
	rv := ffi.CallMethod[foundation.CharacterSet](t_, "tokenizingCharacterSet")
	return rv
}

func (t_ TokenFieldCell) SetTokenizingCharacterSet(value foundation.ICharacterSet) {
	ffi.CallMethod[ffi.Void](t_, "setTokenizingCharacterSet:", value)
}

func (t_ TokenFieldCell) CompletionDelay() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](t_, "completionDelay")
	return rv
}

func (t_ TokenFieldCell) SetCompletionDelay(value foundation.TimeInterval) {
	ffi.CallMethod[ffi.Void](t_, "setCompletionDelay:", value)
}

func (tc _TokenFieldCellClass) DefaultCompletionDelay() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](tc, "defaultCompletionDelay")
	return rv
}

func (t_ TokenFieldCell) Delegate() TokenFieldCellDelegateWrapper {
	rv := ffi.CallMethod[TokenFieldCellDelegateWrapper](t_, "delegate")
	return rv
}

func (t_ TokenFieldCell) SetDelegate(value TokenFieldCellDelegate) {
	po := ffi.CreateProtocol("NSTokenFieldCellDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](t_, "setDelegate:", po)
}

func (t_ TokenFieldCell) SetDelegate0(value objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "setDelegate:", value)
}
