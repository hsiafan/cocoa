// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type TokenFieldDelegate interface {
	TextFieldDelegate
	ImplementsTokenField_DisplayStringForRepresentedObject() bool
	// optional
	TokenField_DisplayStringForRepresentedObject(tokenField TokenField, representedObject objc.Object) string
	ImplementsTokenField_StyleForRepresentedObject() bool
	// optional
	TokenField_StyleForRepresentedObject(tokenField TokenField, representedObject objc.Object) TokenStyle
	ImplementsTokenField_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem() bool
	// optional
	TokenField_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(tokenField TokenField, substring string, tokenIndex int, selectedIndex *int) []objc.IObject
	ImplementsTokenField_EditingStringForRepresentedObject() bool
	// optional
	TokenField_EditingStringForRepresentedObject(tokenField TokenField, representedObject objc.Object) string
	ImplementsTokenField_RepresentedObjectForEditingString() bool
	// optional
	TokenField_RepresentedObjectForEditingString(tokenField TokenField, editingString string) objc.IObject
	ImplementsTokenField_ShouldAddObjects_AtIndex() bool
	// optional
	TokenField_ShouldAddObjects_AtIndex(tokenField TokenField, tokens []objc.Object, index uint) []objc.IObject
	ImplementsTokenField_ReadFromPasteboard() bool
	// optional
	TokenField_ReadFromPasteboard(tokenField TokenField, pboard Pasteboard) []objc.IObject
	ImplementsTokenField_WriteRepresentedObjects_ToPasteboard() bool
	// optional
	TokenField_WriteRepresentedObjects_ToPasteboard(tokenField TokenField, objects []objc.Object, pboard Pasteboard) bool
	ImplementsTokenField_HasMenuForRepresentedObject() bool
	// optional
	TokenField_HasMenuForRepresentedObject(tokenField TokenField, representedObject objc.Object) bool
	ImplementsTokenField_MenuForRepresentedObject() bool
	// optional
	TokenField_MenuForRepresentedObject(tokenField TokenField, representedObject objc.Object) IMenu
}

func WrapTokenFieldDelegate(v TokenFieldDelegate) objc.Object {
	return objc.WrapAsProtocol("NSTokenFieldDelegate", v)
}

type TokenFieldDelegateBase struct {
	TextFieldDelegateBase
}

func (p *TokenFieldDelegateBase) ImplementsTokenField_DisplayStringForRepresentedObject() bool {
	return false
}

func (p *TokenFieldDelegateBase) TokenField_DisplayStringForRepresentedObject(tokenField TokenField, representedObject objc.Object) string {
	panic("unimpemented protocol method")
}

func (p *TokenFieldDelegateBase) ImplementsTokenField_StyleForRepresentedObject() bool {
	return false
}

func (p *TokenFieldDelegateBase) TokenField_StyleForRepresentedObject(tokenField TokenField, representedObject objc.Object) TokenStyle {
	panic("unimpemented protocol method")
}

func (p *TokenFieldDelegateBase) ImplementsTokenField_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem() bool {
	return false
}

func (p *TokenFieldDelegateBase) TokenField_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(tokenField TokenField, substring string, tokenIndex int, selectedIndex *int) []objc.IObject {
	panic("unimpemented protocol method")
}

func (p *TokenFieldDelegateBase) ImplementsTokenField_EditingStringForRepresentedObject() bool {
	return false
}

func (p *TokenFieldDelegateBase) TokenField_EditingStringForRepresentedObject(tokenField TokenField, representedObject objc.Object) string {
	panic("unimpemented protocol method")
}

func (p *TokenFieldDelegateBase) ImplementsTokenField_RepresentedObjectForEditingString() bool {
	return false
}

func (p *TokenFieldDelegateBase) TokenField_RepresentedObjectForEditingString(tokenField TokenField, editingString string) objc.IObject {
	panic("unimpemented protocol method")
}

func (p *TokenFieldDelegateBase) ImplementsTokenField_ShouldAddObjects_AtIndex() bool {
	return false
}

func (p *TokenFieldDelegateBase) TokenField_ShouldAddObjects_AtIndex(tokenField TokenField, tokens []objc.Object, index uint) []objc.IObject {
	panic("unimpemented protocol method")
}

func (p *TokenFieldDelegateBase) ImplementsTokenField_ReadFromPasteboard() bool {
	return false
}

func (p *TokenFieldDelegateBase) TokenField_ReadFromPasteboard(tokenField TokenField, pboard Pasteboard) []objc.IObject {
	panic("unimpemented protocol method")
}

func (p *TokenFieldDelegateBase) ImplementsTokenField_WriteRepresentedObjects_ToPasteboard() bool {
	return false
}

func (p *TokenFieldDelegateBase) TokenField_WriteRepresentedObjects_ToPasteboard(tokenField TokenField, objects []objc.Object, pboard Pasteboard) bool {
	panic("unimpemented protocol method")
}

func (p *TokenFieldDelegateBase) ImplementsTokenField_HasMenuForRepresentedObject() bool {
	return false
}

func (p *TokenFieldDelegateBase) TokenField_HasMenuForRepresentedObject(tokenField TokenField, representedObject objc.Object) bool {
	panic("unimpemented protocol method")
}

func (p *TokenFieldDelegateBase) ImplementsTokenField_MenuForRepresentedObject() bool {
	return false
}

func (p *TokenFieldDelegateBase) TokenField_MenuForRepresentedObject(tokenField TokenField, representedObject objc.Object) IMenu {
	panic("unimpemented protocol method")
}
