// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type SearchFieldDelegate interface {
	TextFieldDelegate
	ImplementsSearchFieldDidStartSearching() bool
	// optional
	SearchFieldDidStartSearching(sender SearchField)
	ImplementsSearchFieldDidEndSearching() bool
	// optional
	SearchFieldDidEndSearching(sender SearchField)
}

func WrapSearchFieldDelegate(v SearchFieldDelegate) objc.Object {
	return objc.WrapAsProtocol("NSSearchFieldDelegate", v)
}

type SearchFieldDelegateBase struct {
	TextFieldDelegateBase
}

func (p *SearchFieldDelegateBase) ImplementsSearchFieldDidStartSearching() bool {
	return false
}

func (p *SearchFieldDelegateBase) SearchFieldDidStartSearching(sender SearchField) {
	panic("unimpemented protocol method")
}

func (p *SearchFieldDelegateBase) ImplementsSearchFieldDidEndSearching() bool {
	return false
}

func (p *SearchFieldDelegateBase) SearchFieldDidEndSearching(sender SearchField) {
	panic("unimpemented protocol method")
}

type SearchFieldDelegateWrapper struct {
	objc.Object
}

func (s_ SearchFieldDelegateWrapper) SearchFieldDidStartSearching(sender ISearchField) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("searchFieldDidStartSearching:"), objc.ExtractPtr(sender))
}

func (s_ SearchFieldDelegateWrapper) SearchFieldDidEndSearching(sender ISearchField) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("searchFieldDidEndSearching:"), objc.ExtractPtr(sender))
}

func (s_ SearchFieldDelegateWrapper) TextField_TextView_Candidates_ForSelectedRange(textField ITextField, textView ITextView, candidates []foundation.ITextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult {
	rv := objc.CallMethod[[]foundation.TextCheckingResult](s_, objc.GetSelector("textField:textView:candidates:forSelectedRange:"), objc.ExtractPtr(textField), objc.ExtractPtr(textView), candidates, selectedRange)
	return rv
}

func (s_ SearchFieldDelegateWrapper) TextField_TextView_CandidatesForSelectedRange(textField ITextField, textView ITextView, selectedRange foundation.Range) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](s_, objc.GetSelector("textField:textView:candidatesForSelectedRange:"), objc.ExtractPtr(textField), objc.ExtractPtr(textView), selectedRange)
	return rv
}

func (s_ SearchFieldDelegateWrapper) TextField_TextView_ShouldSelectCandidateAtIndex(textField ITextField, textView ITextView, index uint) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("textField:textView:shouldSelectCandidateAtIndex:"), objc.ExtractPtr(textField), objc.ExtractPtr(textView), index)
	return rv
}

func (s_ SearchFieldDelegateWrapper) Control_IsValidObject(control IControl, obj objc.IObject) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("control:isValidObject:"), objc.ExtractPtr(control), objc.ExtractPtr(obj))
	return rv
}

func (s_ SearchFieldDelegateWrapper) Control_DidFailToValidatePartialString_ErrorDescription(control IControl, string_ string, error string) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("control:didFailToValidatePartialString:errorDescription:"), objc.ExtractPtr(control), string_, error)
}

func (s_ SearchFieldDelegateWrapper) Control_DidFailToFormatString_ErrorDescription(control IControl, string_ string, error string) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("control:didFailToFormatString:errorDescription:"), objc.ExtractPtr(control), string_, error)
	return rv
}

func (s_ SearchFieldDelegateWrapper) Control_TextShouldBeginEditing(control IControl, fieldEditor IText) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("control:textShouldBeginEditing:"), objc.ExtractPtr(control), objc.ExtractPtr(fieldEditor))
	return rv
}

func (s_ SearchFieldDelegateWrapper) Control_TextShouldEndEditing(control IControl, fieldEditor IText) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("control:textShouldEndEditing:"), objc.ExtractPtr(control), objc.ExtractPtr(fieldEditor))
	return rv
}

func (s_ SearchFieldDelegateWrapper) Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(control IControl, textView ITextView, words []string, charRange foundation.Range, index *int) []string {
	rv := objc.CallMethod[[]string](s_, objc.GetSelector("control:textView:completions:forPartialWordRange:indexOfSelectedItem:"), objc.ExtractPtr(control), objc.ExtractPtr(textView), words, charRange, index)
	return rv
}

func (s_ SearchFieldDelegateWrapper) Control_TextView_DoCommandBySelector(control IControl, textView ITextView, commandSelector objc.Selector) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("control:textView:doCommandBySelector:"), objc.ExtractPtr(control), objc.ExtractPtr(textView), commandSelector)
	return rv
}

func (s_ SearchFieldDelegateWrapper) ControlTextDidBeginEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("controlTextDidBeginEditing:"), objc.ExtractPtr(obj))
}

func (s_ SearchFieldDelegateWrapper) ControlTextDidChange(obj foundation.INotification) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("controlTextDidChange:"), objc.ExtractPtr(obj))
}

func (s_ SearchFieldDelegateWrapper) ControlTextDidEndEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("controlTextDidEndEditing:"), objc.ExtractPtr(obj))
}
