// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
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

type TokenFieldDelegateWrapper struct {
	objc.Object
}

func (t_ TokenFieldDelegateWrapper) TokenField_DisplayStringForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("tokenField:displayStringForRepresentedObject:"), objc.ExtractPtr(tokenField), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ TokenFieldDelegateWrapper) TokenField_StyleForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) TokenStyle {
	rv := objc.CallMethod[TokenStyle](t_, objc.GetSelector("tokenField:styleForRepresentedObject:"), objc.ExtractPtr(tokenField), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ TokenFieldDelegateWrapper) TokenField_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(tokenField ITokenField, substring string, tokenIndex int, selectedIndex *int) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("tokenField:completionsForSubstring:indexOfToken:indexOfSelectedItem:"), objc.ExtractPtr(tokenField), substring, tokenIndex, selectedIndex)
	return rv
}

func (t_ TokenFieldDelegateWrapper) TokenField_EditingStringForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("tokenField:editingStringForRepresentedObject:"), objc.ExtractPtr(tokenField), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ TokenFieldDelegateWrapper) TokenField_RepresentedObjectForEditingString(tokenField ITokenField, editingString string) objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.GetSelector("tokenField:representedObjectForEditingString:"), objc.ExtractPtr(tokenField), editingString)
	return rv
}

func (t_ TokenFieldDelegateWrapper) TokenField_ShouldAddObjects_AtIndex(tokenField ITokenField, tokens []objc.IObject, index uint) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("tokenField:shouldAddObjects:atIndex:"), objc.ExtractPtr(tokenField), tokens, index)
	return rv
}

func (t_ TokenFieldDelegateWrapper) TokenField_ReadFromPasteboard(tokenField ITokenField, pboard IPasteboard) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("tokenField:readFromPasteboard:"), objc.ExtractPtr(tokenField), objc.ExtractPtr(pboard))
	return rv
}

func (t_ TokenFieldDelegateWrapper) TokenField_WriteRepresentedObjects_ToPasteboard(tokenField ITokenField, objects []objc.IObject, pboard IPasteboard) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tokenField:writeRepresentedObjects:toPasteboard:"), objc.ExtractPtr(tokenField), objects, objc.ExtractPtr(pboard))
	return rv
}

func (t_ TokenFieldDelegateWrapper) TokenField_HasMenuForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tokenField:hasMenuForRepresentedObject:"), objc.ExtractPtr(tokenField), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ TokenFieldDelegateWrapper) TokenField_MenuForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) Menu {
	rv := objc.CallMethod[Menu](t_, objc.GetSelector("tokenField:menuForRepresentedObject:"), objc.ExtractPtr(tokenField), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ TokenFieldDelegateWrapper) TextField_TextView_Candidates_ForSelectedRange(textField ITextField, textView ITextView, candidates []foundation.ITextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult {
	rv := objc.CallMethod[[]foundation.TextCheckingResult](t_, objc.GetSelector("textField:textView:candidates:forSelectedRange:"), objc.ExtractPtr(textField), objc.ExtractPtr(textView), candidates, selectedRange)
	return rv
}

func (t_ TokenFieldDelegateWrapper) TextField_TextView_CandidatesForSelectedRange(textField ITextField, textView ITextView, selectedRange foundation.Range) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("textField:textView:candidatesForSelectedRange:"), objc.ExtractPtr(textField), objc.ExtractPtr(textView), selectedRange)
	return rv
}

func (t_ TokenFieldDelegateWrapper) TextField_TextView_ShouldSelectCandidateAtIndex(textField ITextField, textView ITextView, index uint) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textField:textView:shouldSelectCandidateAtIndex:"), objc.ExtractPtr(textField), objc.ExtractPtr(textView), index)
	return rv
}

func (t_ TokenFieldDelegateWrapper) Control_IsValidObject(control IControl, obj objc.IObject) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("control:isValidObject:"), objc.ExtractPtr(control), objc.ExtractPtr(obj))
	return rv
}

func (t_ TokenFieldDelegateWrapper) Control_DidFailToValidatePartialString_ErrorDescription(control IControl, string_ string, error string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("control:didFailToValidatePartialString:errorDescription:"), objc.ExtractPtr(control), string_, error)
}

func (t_ TokenFieldDelegateWrapper) Control_DidFailToFormatString_ErrorDescription(control IControl, string_ string, error string) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("control:didFailToFormatString:errorDescription:"), objc.ExtractPtr(control), string_, error)
	return rv
}

func (t_ TokenFieldDelegateWrapper) Control_TextShouldBeginEditing(control IControl, fieldEditor IText) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("control:textShouldBeginEditing:"), objc.ExtractPtr(control), objc.ExtractPtr(fieldEditor))
	return rv
}

func (t_ TokenFieldDelegateWrapper) Control_TextShouldEndEditing(control IControl, fieldEditor IText) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("control:textShouldEndEditing:"), objc.ExtractPtr(control), objc.ExtractPtr(fieldEditor))
	return rv
}

func (t_ TokenFieldDelegateWrapper) Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(control IControl, textView ITextView, words []string, charRange foundation.Range, index *int) []string {
	rv := objc.CallMethod[[]string](t_, objc.GetSelector("control:textView:completions:forPartialWordRange:indexOfSelectedItem:"), objc.ExtractPtr(control), objc.ExtractPtr(textView), words, charRange, index)
	return rv
}

func (t_ TokenFieldDelegateWrapper) Control_TextView_DoCommandBySelector(control IControl, textView ITextView, commandSelector objc.Selector) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("control:textView:doCommandBySelector:"), objc.ExtractPtr(control), objc.ExtractPtr(textView), commandSelector)
	return rv
}

func (t_ TokenFieldDelegateWrapper) ControlTextDidBeginEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("controlTextDidBeginEditing:"), objc.ExtractPtr(obj))
}

func (t_ TokenFieldDelegateWrapper) ControlTextDidChange(obj foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("controlTextDidChange:"), objc.ExtractPtr(obj))
}

func (t_ TokenFieldDelegateWrapper) ControlTextDidEndEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("controlTextDidEndEditing:"), objc.ExtractPtr(obj))
}
