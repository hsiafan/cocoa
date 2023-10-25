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

func WrapTextStorageDelegate(v TextStorageDelegate) objc.Object {
	return objc.WrapAsProtocol("NSTextStorageDelegate", v)
}

type TextStorageDelegateBase struct {
}

func (p *TextStorageDelegateBase) ImplementsTextStorage_WillProcessEditing_Range_ChangeInLength() bool {
	return false
}

func (p *TextStorageDelegateBase) TextStorage_WillProcessEditing_Range_ChangeInLength(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	panic("unimpemented protocol method")
}

func (p *TextStorageDelegateBase) ImplementsTextStorage_DidProcessEditing_Range_ChangeInLength() bool {
	return false
}

func (p *TextStorageDelegateBase) TextStorage_DidProcessEditing_Range_ChangeInLength(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	panic("unimpemented protocol method")
}
