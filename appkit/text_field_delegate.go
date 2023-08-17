// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type TextFieldDelegate interface {
	ControlTextEditingDelegate
	ImplementsTextField_TextView_Candidates_ForSelectedRange() bool
	// optional
	TextField_TextView_Candidates_ForSelectedRange(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult
	ImplementsTextField_TextView_CandidatesForSelectedRange() bool
	// optional
	TextField_TextView_CandidatesForSelectedRange(textField TextField, textView TextView, selectedRange foundation.Range) []objc.IObject
	ImplementsTextField_TextView_ShouldSelectCandidateAtIndex() bool
	// optional
	TextField_TextView_ShouldSelectCandidateAtIndex(textField TextField, textView TextView, index uint) bool
}

func WrapTextFieldDelegate(v TextFieldDelegate) objc.Object {
	return objc.WrapAsProtocol("NSTextFieldDelegate", v)
}

type TextFieldDelegateBase struct {
	ControlTextEditingDelegateBase
}

func (p *TextFieldDelegateBase) ImplementsTextField_TextView_Candidates_ForSelectedRange() bool {
	return false
}

func (p *TextFieldDelegateBase) TextField_TextView_Candidates_ForSelectedRange(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult {
	panic("unimpemented protocol method")
}

func (p *TextFieldDelegateBase) ImplementsTextField_TextView_CandidatesForSelectedRange() bool {
	return false
}

func (p *TextFieldDelegateBase) TextField_TextView_CandidatesForSelectedRange(textField TextField, textView TextView, selectedRange foundation.Range) []objc.IObject {
	panic("unimpemented protocol method")
}

func (p *TextFieldDelegateBase) ImplementsTextField_TextView_ShouldSelectCandidateAtIndex() bool {
	return false
}

func (p *TextFieldDelegateBase) TextField_TextView_ShouldSelectCandidateAtIndex(textField TextField, textView TextView, index uint) bool {
	panic("unimpemented protocol method")
}

type TextFieldDelegateWrapper struct {
	objc.Object
}

func (t_ TextFieldDelegateWrapper) TextField_TextView_Candidates_ForSelectedRange(textField ITextField, textView ITextView, candidates []foundation.ITextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult {
	rv := objc.CallMethod[[]foundation.TextCheckingResult](t_, objc.GetSelector("textField:textView:candidates:forSelectedRange:"), objc.ExtractPtr(textField), objc.ExtractPtr(textView), candidates, selectedRange)
	return rv
}

func (t_ TextFieldDelegateWrapper) TextField_TextView_CandidatesForSelectedRange(textField ITextField, textView ITextView, selectedRange foundation.Range) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("textField:textView:candidatesForSelectedRange:"), objc.ExtractPtr(textField), objc.ExtractPtr(textView), selectedRange)
	return rv
}

func (t_ TextFieldDelegateWrapper) TextField_TextView_ShouldSelectCandidateAtIndex(textField ITextField, textView ITextView, index uint) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textField:textView:shouldSelectCandidateAtIndex:"), objc.ExtractPtr(textField), objc.ExtractPtr(textView), index)
	return rv
}

func (t_ TextFieldDelegateWrapper) Control_IsValidObject(control IControl, obj objc.IObject) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("control:isValidObject:"), objc.ExtractPtr(control), objc.ExtractPtr(obj))
	return rv
}

func (t_ TextFieldDelegateWrapper) Control_DidFailToValidatePartialString_ErrorDescription(control IControl, string_ string, error string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("control:didFailToValidatePartialString:errorDescription:"), objc.ExtractPtr(control), string_, error)
}

func (t_ TextFieldDelegateWrapper) Control_DidFailToFormatString_ErrorDescription(control IControl, string_ string, error string) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("control:didFailToFormatString:errorDescription:"), objc.ExtractPtr(control), string_, error)
	return rv
}

func (t_ TextFieldDelegateWrapper) Control_TextShouldBeginEditing(control IControl, fieldEditor IText) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("control:textShouldBeginEditing:"), objc.ExtractPtr(control), objc.ExtractPtr(fieldEditor))
	return rv
}

func (t_ TextFieldDelegateWrapper) Control_TextShouldEndEditing(control IControl, fieldEditor IText) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("control:textShouldEndEditing:"), objc.ExtractPtr(control), objc.ExtractPtr(fieldEditor))
	return rv
}

func (t_ TextFieldDelegateWrapper) Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(control IControl, textView ITextView, words []string, charRange foundation.Range, index *int) []string {
	rv := objc.CallMethod[[]string](t_, objc.GetSelector("control:textView:completions:forPartialWordRange:indexOfSelectedItem:"), objc.ExtractPtr(control), objc.ExtractPtr(textView), words, charRange, index)
	return rv
}

func (t_ TextFieldDelegateWrapper) Control_TextView_DoCommandBySelector(control IControl, textView ITextView, commandSelector objc.Selector) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("control:textView:doCommandBySelector:"), objc.ExtractPtr(control), objc.ExtractPtr(textView), commandSelector)
	return rv
}

func (t_ TextFieldDelegateWrapper) ControlTextDidBeginEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("controlTextDidBeginEditing:"), objc.ExtractPtr(obj))
}

func (t_ TextFieldDelegateWrapper) ControlTextDidChange(obj foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("controlTextDidChange:"), objc.ExtractPtr(obj))
}

func (t_ TextFieldDelegateWrapper) ControlTextDidEndEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("controlTextDidEndEditing:"), objc.ExtractPtr(obj))
}
