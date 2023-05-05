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

type TokenFieldDelegateImpl struct {
	TextFieldDelegateImpl
	_TokenField_DisplayStringForRepresentedObject                        func(tokenField TokenField, representedObject objc.Object) string
	_TokenField_StyleForRepresentedObject                                func(tokenField TokenField, representedObject objc.Object) TokenStyle
	_TokenField_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem func(tokenField TokenField, substring string, tokenIndex int, selectedIndex *int) []objc.IObject
	_TokenField_EditingStringForRepresentedObject                        func(tokenField TokenField, representedObject objc.Object) string
	_TokenField_RepresentedObjectForEditingString                        func(tokenField TokenField, editingString string) objc.IObject
	_TokenField_ShouldAddObjects_AtIndex                                 func(tokenField TokenField, tokens []objc.Object, index uint) []objc.IObject
	_TokenField_ReadFromPasteboard                                       func(tokenField TokenField, pboard Pasteboard) []objc.IObject
	_TokenField_WriteRepresentedObjects_ToPasteboard                     func(tokenField TokenField, objects []objc.Object, pboard Pasteboard) bool
	_TokenField_HasMenuForRepresentedObject                              func(tokenField TokenField, representedObject objc.Object) bool
	_TokenField_MenuForRepresentedObject                                 func(tokenField TokenField, representedObject objc.Object) IMenu
}

func (di *TokenFieldDelegateImpl) ImplementsTokenField_DisplayStringForRepresentedObject() bool {
	return di._TokenField_DisplayStringForRepresentedObject != nil
}

func (di *TokenFieldDelegateImpl) SetTokenField_DisplayStringForRepresentedObject(f func(tokenField TokenField, representedObject objc.Object) string) {
	di._TokenField_DisplayStringForRepresentedObject = f
}

func (di *TokenFieldDelegateImpl) TokenField_DisplayStringForRepresentedObject(tokenField TokenField, representedObject objc.Object) string {
	return di._TokenField_DisplayStringForRepresentedObject(tokenField, representedObject)
}
func (di *TokenFieldDelegateImpl) ImplementsTokenField_StyleForRepresentedObject() bool {
	return di._TokenField_StyleForRepresentedObject != nil
}

func (di *TokenFieldDelegateImpl) SetTokenField_StyleForRepresentedObject(f func(tokenField TokenField, representedObject objc.Object) TokenStyle) {
	di._TokenField_StyleForRepresentedObject = f
}

func (di *TokenFieldDelegateImpl) TokenField_StyleForRepresentedObject(tokenField TokenField, representedObject objc.Object) TokenStyle {
	return di._TokenField_StyleForRepresentedObject(tokenField, representedObject)
}
func (di *TokenFieldDelegateImpl) ImplementsTokenField_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem() bool {
	return di._TokenField_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem != nil
}

func (di *TokenFieldDelegateImpl) SetTokenField_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(f func(tokenField TokenField, substring string, tokenIndex int, selectedIndex *int) []objc.IObject) {
	di._TokenField_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem = f
}

func (di *TokenFieldDelegateImpl) TokenField_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(tokenField TokenField, substring string, tokenIndex int, selectedIndex *int) []objc.IObject {
	return di._TokenField_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(tokenField, substring, tokenIndex, selectedIndex)
}
func (di *TokenFieldDelegateImpl) ImplementsTokenField_EditingStringForRepresentedObject() bool {
	return di._TokenField_EditingStringForRepresentedObject != nil
}

func (di *TokenFieldDelegateImpl) SetTokenField_EditingStringForRepresentedObject(f func(tokenField TokenField, representedObject objc.Object) string) {
	di._TokenField_EditingStringForRepresentedObject = f
}

func (di *TokenFieldDelegateImpl) TokenField_EditingStringForRepresentedObject(tokenField TokenField, representedObject objc.Object) string {
	return di._TokenField_EditingStringForRepresentedObject(tokenField, representedObject)
}
func (di *TokenFieldDelegateImpl) ImplementsTokenField_RepresentedObjectForEditingString() bool {
	return di._TokenField_RepresentedObjectForEditingString != nil
}

func (di *TokenFieldDelegateImpl) SetTokenField_RepresentedObjectForEditingString(f func(tokenField TokenField, editingString string) objc.IObject) {
	di._TokenField_RepresentedObjectForEditingString = f
}

func (di *TokenFieldDelegateImpl) TokenField_RepresentedObjectForEditingString(tokenField TokenField, editingString string) objc.IObject {
	return di._TokenField_RepresentedObjectForEditingString(tokenField, editingString)
}
func (di *TokenFieldDelegateImpl) ImplementsTokenField_ShouldAddObjects_AtIndex() bool {
	return di._TokenField_ShouldAddObjects_AtIndex != nil
}

func (di *TokenFieldDelegateImpl) SetTokenField_ShouldAddObjects_AtIndex(f func(tokenField TokenField, tokens []objc.Object, index uint) []objc.IObject) {
	di._TokenField_ShouldAddObjects_AtIndex = f
}

func (di *TokenFieldDelegateImpl) TokenField_ShouldAddObjects_AtIndex(tokenField TokenField, tokens []objc.Object, index uint) []objc.IObject {
	return di._TokenField_ShouldAddObjects_AtIndex(tokenField, tokens, index)
}
func (di *TokenFieldDelegateImpl) ImplementsTokenField_ReadFromPasteboard() bool {
	return di._TokenField_ReadFromPasteboard != nil
}

func (di *TokenFieldDelegateImpl) SetTokenField_ReadFromPasteboard(f func(tokenField TokenField, pboard Pasteboard) []objc.IObject) {
	di._TokenField_ReadFromPasteboard = f
}

func (di *TokenFieldDelegateImpl) TokenField_ReadFromPasteboard(tokenField TokenField, pboard Pasteboard) []objc.IObject {
	return di._TokenField_ReadFromPasteboard(tokenField, pboard)
}
func (di *TokenFieldDelegateImpl) ImplementsTokenField_WriteRepresentedObjects_ToPasteboard() bool {
	return di._TokenField_WriteRepresentedObjects_ToPasteboard != nil
}

func (di *TokenFieldDelegateImpl) SetTokenField_WriteRepresentedObjects_ToPasteboard(f func(tokenField TokenField, objects []objc.Object, pboard Pasteboard) bool) {
	di._TokenField_WriteRepresentedObjects_ToPasteboard = f
}

func (di *TokenFieldDelegateImpl) TokenField_WriteRepresentedObjects_ToPasteboard(tokenField TokenField, objects []objc.Object, pboard Pasteboard) bool {
	return di._TokenField_WriteRepresentedObjects_ToPasteboard(tokenField, objects, pboard)
}
func (di *TokenFieldDelegateImpl) ImplementsTokenField_HasMenuForRepresentedObject() bool {
	return di._TokenField_HasMenuForRepresentedObject != nil
}

func (di *TokenFieldDelegateImpl) SetTokenField_HasMenuForRepresentedObject(f func(tokenField TokenField, representedObject objc.Object) bool) {
	di._TokenField_HasMenuForRepresentedObject = f
}

func (di *TokenFieldDelegateImpl) TokenField_HasMenuForRepresentedObject(tokenField TokenField, representedObject objc.Object) bool {
	return di._TokenField_HasMenuForRepresentedObject(tokenField, representedObject)
}
func (di *TokenFieldDelegateImpl) ImplementsTokenField_MenuForRepresentedObject() bool {
	return di._TokenField_MenuForRepresentedObject != nil
}

func (di *TokenFieldDelegateImpl) SetTokenField_MenuForRepresentedObject(f func(tokenField TokenField, representedObject objc.Object) IMenu) {
	di._TokenField_MenuForRepresentedObject = f
}

func (di *TokenFieldDelegateImpl) TokenField_MenuForRepresentedObject(tokenField TokenField, representedObject objc.Object) IMenu {
	return di._TokenField_MenuForRepresentedObject(tokenField, representedObject)
}

type TokenFieldDelegateWrapper struct {
	TextFieldDelegateWrapper
}

func (t_ *TokenFieldDelegateWrapper) ImplementsTokenField_DisplayStringForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:displayStringForRepresentedObject:"))
}

func (t_ TokenFieldDelegateWrapper) TokenField_DisplayStringForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("tokenField:displayStringForRepresentedObject:"), objc.ExtractPtr(tokenField), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ *TokenFieldDelegateWrapper) ImplementsTokenField_StyleForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:styleForRepresentedObject:"))
}

func (t_ TokenFieldDelegateWrapper) TokenField_StyleForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) TokenStyle {
	rv := objc.CallMethod[TokenStyle](t_, objc.GetSelector("tokenField:styleForRepresentedObject:"), objc.ExtractPtr(tokenField), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ *TokenFieldDelegateWrapper) ImplementsTokenField_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:completionsForSubstring:indexOfToken:indexOfSelectedItem:"))
}

func (t_ TokenFieldDelegateWrapper) TokenField_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(tokenField ITokenField, substring string, tokenIndex int, selectedIndex *int) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("tokenField:completionsForSubstring:indexOfToken:indexOfSelectedItem:"), objc.ExtractPtr(tokenField), substring, tokenIndex, selectedIndex)
	return rv
}

func (t_ *TokenFieldDelegateWrapper) ImplementsTokenField_EditingStringForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:editingStringForRepresentedObject:"))
}

func (t_ TokenFieldDelegateWrapper) TokenField_EditingStringForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("tokenField:editingStringForRepresentedObject:"), objc.ExtractPtr(tokenField), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ *TokenFieldDelegateWrapper) ImplementsTokenField_RepresentedObjectForEditingString() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:representedObjectForEditingString:"))
}

func (t_ TokenFieldDelegateWrapper) TokenField_RepresentedObjectForEditingString(tokenField ITokenField, editingString string) objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.GetSelector("tokenField:representedObjectForEditingString:"), objc.ExtractPtr(tokenField), editingString)
	return rv
}

func (t_ *TokenFieldDelegateWrapper) ImplementsTokenField_ShouldAddObjects_AtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:shouldAddObjects:atIndex:"))
}

func (t_ TokenFieldDelegateWrapper) TokenField_ShouldAddObjects_AtIndex(tokenField ITokenField, tokens []objc.IObject, index uint) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("tokenField:shouldAddObjects:atIndex:"), objc.ExtractPtr(tokenField), tokens, index)
	return rv
}

func (t_ *TokenFieldDelegateWrapper) ImplementsTokenField_ReadFromPasteboard() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:readFromPasteboard:"))
}

func (t_ TokenFieldDelegateWrapper) TokenField_ReadFromPasteboard(tokenField ITokenField, pboard IPasteboard) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("tokenField:readFromPasteboard:"), objc.ExtractPtr(tokenField), objc.ExtractPtr(pboard))
	return rv
}

func (t_ *TokenFieldDelegateWrapper) ImplementsTokenField_WriteRepresentedObjects_ToPasteboard() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:writeRepresentedObjects:toPasteboard:"))
}

func (t_ TokenFieldDelegateWrapper) TokenField_WriteRepresentedObjects_ToPasteboard(tokenField ITokenField, objects []objc.IObject, pboard IPasteboard) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tokenField:writeRepresentedObjects:toPasteboard:"), objc.ExtractPtr(tokenField), objects, objc.ExtractPtr(pboard))
	return rv
}

func (t_ *TokenFieldDelegateWrapper) ImplementsTokenField_HasMenuForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:hasMenuForRepresentedObject:"))
}

func (t_ TokenFieldDelegateWrapper) TokenField_HasMenuForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tokenField:hasMenuForRepresentedObject:"), objc.ExtractPtr(tokenField), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ *TokenFieldDelegateWrapper) ImplementsTokenField_MenuForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:menuForRepresentedObject:"))
}

func (t_ TokenFieldDelegateWrapper) TokenField_MenuForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) Menu {
	rv := objc.CallMethod[Menu](t_, objc.GetSelector("tokenField:menuForRepresentedObject:"), objc.ExtractPtr(tokenField), objc.ExtractPtr(representedObject))
	return rv
}
