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

type TokenFieldCellDelegateWrapper struct {
	objc.Object
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_DisplayStringForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("tokenFieldCell:displayStringForRepresentedObject:"), objc.ExtractPtr(tokenFieldCell), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_StyleForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) TokenStyle {
	rv := objc.CallMethod[TokenStyle](t_, objc.GetSelector("tokenFieldCell:styleForRepresentedObject:"), objc.ExtractPtr(tokenFieldCell), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(tokenFieldCell ITokenFieldCell, substring string, tokenIndex int, selectedIndex *int) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("tokenFieldCell:completionsForSubstring:indexOfToken:indexOfSelectedItem:"), objc.ExtractPtr(tokenFieldCell), substring, tokenIndex, selectedIndex)
	return rv
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_EditingStringForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("tokenFieldCell:editingStringForRepresentedObject:"), objc.ExtractPtr(tokenFieldCell), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_RepresentedObjectForEditingString(tokenFieldCell ITokenFieldCell, editingString string) objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.GetSelector("tokenFieldCell:representedObjectForEditingString:"), objc.ExtractPtr(tokenFieldCell), editingString)
	return rv
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_ShouldAddObjects_AtIndex(tokenFieldCell ITokenFieldCell, tokens []objc.IObject, index uint) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("tokenFieldCell:shouldAddObjects:atIndex:"), objc.ExtractPtr(tokenFieldCell), tokens, index)
	return rv
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_ReadFromPasteboard(tokenFieldCell ITokenFieldCell, pboard IPasteboard) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("tokenFieldCell:readFromPasteboard:"), objc.ExtractPtr(tokenFieldCell), objc.ExtractPtr(pboard))
	return rv
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_WriteRepresentedObjects_ToPasteboard(tokenFieldCell ITokenFieldCell, objects []objc.IObject, pboard IPasteboard) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tokenFieldCell:writeRepresentedObjects:toPasteboard:"), objc.ExtractPtr(tokenFieldCell), objects, objc.ExtractPtr(pboard))
	return rv
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_HasMenuForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tokenFieldCell:hasMenuForRepresentedObject:"), objc.ExtractPtr(tokenFieldCell), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_MenuForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) Menu {
	rv := objc.CallMethod[Menu](t_, objc.GetSelector("tokenFieldCell:menuForRepresentedObject:"), objc.ExtractPtr(tokenFieldCell), objc.ExtractPtr(representedObject))
	return rv
}
