// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type TextStorageDelegate interface {
	ImplementsTextStorage_WillProcessEditing_Range_ChangeInLength() bool
	// optional
	TextStorage_WillProcessEditing_Range_ChangeInLength(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
	ImplementsTextStorage_DidProcessEditing_Range_ChangeInLength() bool
	// optional
	TextStorage_DidProcessEditing_Range_ChangeInLength(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
}

type TextStorageDelegateImpl struct {
	_TextStorage_WillProcessEditing_Range_ChangeInLength func(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
	_TextStorage_DidProcessEditing_Range_ChangeInLength  func(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
}

func (di *TextStorageDelegateImpl) ImplementsTextStorage_WillProcessEditing_Range_ChangeInLength() bool {
	return di._TextStorage_WillProcessEditing_Range_ChangeInLength != nil
}

func (di *TextStorageDelegateImpl) SetTextStorage_WillProcessEditing_Range_ChangeInLength(f func(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)) {
	di._TextStorage_WillProcessEditing_Range_ChangeInLength = f
}

func (di *TextStorageDelegateImpl) TextStorage_WillProcessEditing_Range_ChangeInLength(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	di._TextStorage_WillProcessEditing_Range_ChangeInLength(textStorage, editedMask, editedRange, delta)
}
func (di *TextStorageDelegateImpl) ImplementsTextStorage_DidProcessEditing_Range_ChangeInLength() bool {
	return di._TextStorage_DidProcessEditing_Range_ChangeInLength != nil
}

func (di *TextStorageDelegateImpl) SetTextStorage_DidProcessEditing_Range_ChangeInLength(f func(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)) {
	di._TextStorage_DidProcessEditing_Range_ChangeInLength = f
}

func (di *TextStorageDelegateImpl) TextStorage_DidProcessEditing_Range_ChangeInLength(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	di._TextStorage_DidProcessEditing_Range_ChangeInLength(textStorage, editedMask, editedRange, delta)
}

type TextStorageDelegateWrapper struct {
	objc.Object
}

func (t_ *TextStorageDelegateWrapper) ImplementsTextStorage_WillProcessEditing_Range_ChangeInLength() bool {
	return t_.RespondsToSelector(objc.GetSelector("textStorage:willProcessEditing:range:changeInLength:"))
}

func (t_ TextStorageDelegateWrapper) TextStorage_WillProcessEditing_Range_ChangeInLength(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textStorage:willProcessEditing:range:changeInLength:"), objc.ExtractPtr(textStorage), editedMask, editedRange, delta)
}

func (t_ *TextStorageDelegateWrapper) ImplementsTextStorage_DidProcessEditing_Range_ChangeInLength() bool {
	return t_.RespondsToSelector(objc.GetSelector("textStorage:didProcessEditing:range:changeInLength:"))
}

func (t_ TextStorageDelegateWrapper) TextStorage_DidProcessEditing_Range_ChangeInLength(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textStorage:didProcessEditing:range:changeInLength:"), objc.ExtractPtr(textStorage), editedMask, editedRange, delta)
}
