// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type TokenFieldCellDelegate interface {
	ImplementsTokenFieldCell_DisplayStringForRepresentedObject() bool
	// optional
	TokenFieldCell_DisplayStringForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) string
	ImplementsTokenFieldCell_StyleForRepresentedObject() bool
	// optional
	TokenFieldCell_StyleForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) TokenStyle
	ImplementsTokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem() bool
	// optional
	TokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(tokenFieldCell TokenFieldCell, substring string, tokenIndex int, selectedIndex *int) []objc.IObject
	ImplementsTokenFieldCell_EditingStringForRepresentedObject() bool
	// optional
	TokenFieldCell_EditingStringForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) string
	ImplementsTokenFieldCell_RepresentedObjectForEditingString() bool
	// optional
	TokenFieldCell_RepresentedObjectForEditingString(tokenFieldCell TokenFieldCell, editingString string) objc.IObject
	ImplementsTokenFieldCell_ShouldAddObjects_AtIndex() bool
	// optional
	TokenFieldCell_ShouldAddObjects_AtIndex(tokenFieldCell TokenFieldCell, tokens []objc.Object, index uint) []objc.IObject
	ImplementsTokenFieldCell_ReadFromPasteboard() bool
	// optional
	TokenFieldCell_ReadFromPasteboard(tokenFieldCell TokenFieldCell, pboard Pasteboard) []objc.IObject
	ImplementsTokenFieldCell_WriteRepresentedObjects_ToPasteboard() bool
	// optional
	TokenFieldCell_WriteRepresentedObjects_ToPasteboard(tokenFieldCell TokenFieldCell, objects []objc.Object, pboard Pasteboard) bool
	ImplementsTokenFieldCell_HasMenuForRepresentedObject() bool
	// optional
	TokenFieldCell_HasMenuForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) bool
	ImplementsTokenFieldCell_MenuForRepresentedObject() bool
	// optional
	TokenFieldCell_MenuForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) IMenu
}

func WrapTokenFieldCellDelegate(v TokenFieldCellDelegate) objc.Object {
	return objc.WrapAsProtocol("NSTokenFieldCellDelegate", v)
}

type TokenFieldCellDelegateBase struct {
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_DisplayStringForRepresentedObject() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_DisplayStringForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) string {
	panic("unimpemented protocol method")
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_StyleForRepresentedObject() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_StyleForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) TokenStyle {
	panic("unimpemented protocol method")
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(tokenFieldCell TokenFieldCell, substring string, tokenIndex int, selectedIndex *int) []objc.IObject {
	panic("unimpemented protocol method")
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_EditingStringForRepresentedObject() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_EditingStringForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) string {
	panic("unimpemented protocol method")
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_RepresentedObjectForEditingString() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_RepresentedObjectForEditingString(tokenFieldCell TokenFieldCell, editingString string) objc.IObject {
	panic("unimpemented protocol method")
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_ShouldAddObjects_AtIndex() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_ShouldAddObjects_AtIndex(tokenFieldCell TokenFieldCell, tokens []objc.Object, index uint) []objc.IObject {
	panic("unimpemented protocol method")
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_ReadFromPasteboard() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_ReadFromPasteboard(tokenFieldCell TokenFieldCell, pboard Pasteboard) []objc.IObject {
	panic("unimpemented protocol method")
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_WriteRepresentedObjects_ToPasteboard() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_WriteRepresentedObjects_ToPasteboard(tokenFieldCell TokenFieldCell, objects []objc.Object, pboard Pasteboard) bool {
	panic("unimpemented protocol method")
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_HasMenuForRepresentedObject() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_HasMenuForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) bool {
	panic("unimpemented protocol method")
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_MenuForRepresentedObject() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_MenuForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) IMenu {
	panic("unimpemented protocol method")
}
