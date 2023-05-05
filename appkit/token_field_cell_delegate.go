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

type TokenFieldCellDelegateImpl struct {
	_TokenFieldCell_DisplayStringForRepresentedObject                        func(tokenFieldCell TokenFieldCell, representedObject objc.Object) string
	_TokenFieldCell_StyleForRepresentedObject                                func(tokenFieldCell TokenFieldCell, representedObject objc.Object) TokenStyle
	_TokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem func(tokenFieldCell TokenFieldCell, substring string, tokenIndex int, selectedIndex *int) []objc.IObject
	_TokenFieldCell_EditingStringForRepresentedObject                        func(tokenFieldCell TokenFieldCell, representedObject objc.Object) string
	_TokenFieldCell_RepresentedObjectForEditingString                        func(tokenFieldCell TokenFieldCell, editingString string) objc.IObject
	_TokenFieldCell_ShouldAddObjects_AtIndex                                 func(tokenFieldCell TokenFieldCell, tokens []objc.Object, index uint) []objc.IObject
	_TokenFieldCell_ReadFromPasteboard                                       func(tokenFieldCell TokenFieldCell, pboard Pasteboard) []objc.IObject
	_TokenFieldCell_WriteRepresentedObjects_ToPasteboard                     func(tokenFieldCell TokenFieldCell, objects []objc.Object, pboard Pasteboard) bool
	_TokenFieldCell_HasMenuForRepresentedObject                              func(tokenFieldCell TokenFieldCell, representedObject objc.Object) bool
	_TokenFieldCell_MenuForRepresentedObject                                 func(tokenFieldCell TokenFieldCell, representedObject objc.Object) IMenu
}

func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_DisplayStringForRepresentedObject() bool {
	return di._TokenFieldCell_DisplayStringForRepresentedObject != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_DisplayStringForRepresentedObject(f func(tokenFieldCell TokenFieldCell, representedObject objc.Object) string) {
	di._TokenFieldCell_DisplayStringForRepresentedObject = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_DisplayStringForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) string {
	return di._TokenFieldCell_DisplayStringForRepresentedObject(tokenFieldCell, representedObject)
}
func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_StyleForRepresentedObject() bool {
	return di._TokenFieldCell_StyleForRepresentedObject != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_StyleForRepresentedObject(f func(tokenFieldCell TokenFieldCell, representedObject objc.Object) TokenStyle) {
	di._TokenFieldCell_StyleForRepresentedObject = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_StyleForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) TokenStyle {
	return di._TokenFieldCell_StyleForRepresentedObject(tokenFieldCell, representedObject)
}
func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem() bool {
	return di._TokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(f func(tokenFieldCell TokenFieldCell, substring string, tokenIndex int, selectedIndex *int) []objc.IObject) {
	di._TokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(tokenFieldCell TokenFieldCell, substring string, tokenIndex int, selectedIndex *int) []objc.IObject {
	return di._TokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(tokenFieldCell, substring, tokenIndex, selectedIndex)
}
func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_EditingStringForRepresentedObject() bool {
	return di._TokenFieldCell_EditingStringForRepresentedObject != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_EditingStringForRepresentedObject(f func(tokenFieldCell TokenFieldCell, representedObject objc.Object) string) {
	di._TokenFieldCell_EditingStringForRepresentedObject = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_EditingStringForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) string {
	return di._TokenFieldCell_EditingStringForRepresentedObject(tokenFieldCell, representedObject)
}
func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_RepresentedObjectForEditingString() bool {
	return di._TokenFieldCell_RepresentedObjectForEditingString != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_RepresentedObjectForEditingString(f func(tokenFieldCell TokenFieldCell, editingString string) objc.IObject) {
	di._TokenFieldCell_RepresentedObjectForEditingString = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_RepresentedObjectForEditingString(tokenFieldCell TokenFieldCell, editingString string) objc.IObject {
	return di._TokenFieldCell_RepresentedObjectForEditingString(tokenFieldCell, editingString)
}
func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_ShouldAddObjects_AtIndex() bool {
	return di._TokenFieldCell_ShouldAddObjects_AtIndex != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_ShouldAddObjects_AtIndex(f func(tokenFieldCell TokenFieldCell, tokens []objc.Object, index uint) []objc.IObject) {
	di._TokenFieldCell_ShouldAddObjects_AtIndex = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_ShouldAddObjects_AtIndex(tokenFieldCell TokenFieldCell, tokens []objc.Object, index uint) []objc.IObject {
	return di._TokenFieldCell_ShouldAddObjects_AtIndex(tokenFieldCell, tokens, index)
}
func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_ReadFromPasteboard() bool {
	return di._TokenFieldCell_ReadFromPasteboard != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_ReadFromPasteboard(f func(tokenFieldCell TokenFieldCell, pboard Pasteboard) []objc.IObject) {
	di._TokenFieldCell_ReadFromPasteboard = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_ReadFromPasteboard(tokenFieldCell TokenFieldCell, pboard Pasteboard) []objc.IObject {
	return di._TokenFieldCell_ReadFromPasteboard(tokenFieldCell, pboard)
}
func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_WriteRepresentedObjects_ToPasteboard() bool {
	return di._TokenFieldCell_WriteRepresentedObjects_ToPasteboard != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_WriteRepresentedObjects_ToPasteboard(f func(tokenFieldCell TokenFieldCell, objects []objc.Object, pboard Pasteboard) bool) {
	di._TokenFieldCell_WriteRepresentedObjects_ToPasteboard = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_WriteRepresentedObjects_ToPasteboard(tokenFieldCell TokenFieldCell, objects []objc.Object, pboard Pasteboard) bool {
	return di._TokenFieldCell_WriteRepresentedObjects_ToPasteboard(tokenFieldCell, objects, pboard)
}
func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_HasMenuForRepresentedObject() bool {
	return di._TokenFieldCell_HasMenuForRepresentedObject != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_HasMenuForRepresentedObject(f func(tokenFieldCell TokenFieldCell, representedObject objc.Object) bool) {
	di._TokenFieldCell_HasMenuForRepresentedObject = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_HasMenuForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) bool {
	return di._TokenFieldCell_HasMenuForRepresentedObject(tokenFieldCell, representedObject)
}
func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_MenuForRepresentedObject() bool {
	return di._TokenFieldCell_MenuForRepresentedObject != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_MenuForRepresentedObject(f func(tokenFieldCell TokenFieldCell, representedObject objc.Object) IMenu) {
	di._TokenFieldCell_MenuForRepresentedObject = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_MenuForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) IMenu {
	return di._TokenFieldCell_MenuForRepresentedObject(tokenFieldCell, representedObject)
}

type TokenFieldCellDelegateWrapper struct {
	objc.Object
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_DisplayStringForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:displayStringForRepresentedObject:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_DisplayStringForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("tokenFieldCell:displayStringForRepresentedObject:"), objc.ExtractPtr(tokenFieldCell), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_StyleForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:styleForRepresentedObject:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_StyleForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) TokenStyle {
	rv := objc.CallMethod[TokenStyle](t_, objc.GetSelector("tokenFieldCell:styleForRepresentedObject:"), objc.ExtractPtr(tokenFieldCell), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:completionsForSubstring:indexOfToken:indexOfSelectedItem:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(tokenFieldCell ITokenFieldCell, substring string, tokenIndex int, selectedIndex *int) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("tokenFieldCell:completionsForSubstring:indexOfToken:indexOfSelectedItem:"), objc.ExtractPtr(tokenFieldCell), substring, tokenIndex, selectedIndex)
	return rv
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_EditingStringForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:editingStringForRepresentedObject:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_EditingStringForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("tokenFieldCell:editingStringForRepresentedObject:"), objc.ExtractPtr(tokenFieldCell), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_RepresentedObjectForEditingString() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:representedObjectForEditingString:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_RepresentedObjectForEditingString(tokenFieldCell ITokenFieldCell, editingString string) objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.GetSelector("tokenFieldCell:representedObjectForEditingString:"), objc.ExtractPtr(tokenFieldCell), editingString)
	return rv
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_ShouldAddObjects_AtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:shouldAddObjects:atIndex:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_ShouldAddObjects_AtIndex(tokenFieldCell ITokenFieldCell, tokens []objc.IObject, index uint) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("tokenFieldCell:shouldAddObjects:atIndex:"), objc.ExtractPtr(tokenFieldCell), tokens, index)
	return rv
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_ReadFromPasteboard() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:readFromPasteboard:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_ReadFromPasteboard(tokenFieldCell ITokenFieldCell, pboard IPasteboard) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("tokenFieldCell:readFromPasteboard:"), objc.ExtractPtr(tokenFieldCell), objc.ExtractPtr(pboard))
	return rv
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_WriteRepresentedObjects_ToPasteboard() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:writeRepresentedObjects:toPasteboard:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_WriteRepresentedObjects_ToPasteboard(tokenFieldCell ITokenFieldCell, objects []objc.IObject, pboard IPasteboard) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tokenFieldCell:writeRepresentedObjects:toPasteboard:"), objc.ExtractPtr(tokenFieldCell), objects, objc.ExtractPtr(pboard))
	return rv
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_HasMenuForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:hasMenuForRepresentedObject:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_HasMenuForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tokenFieldCell:hasMenuForRepresentedObject:"), objc.ExtractPtr(tokenFieldCell), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_MenuForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:menuForRepresentedObject:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_MenuForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) Menu {
	rv := objc.CallMethod[Menu](t_, objc.GetSelector("tokenFieldCell:menuForRepresentedObject:"), objc.ExtractPtr(tokenFieldCell), objc.ExtractPtr(representedObject))
	return rv
}
