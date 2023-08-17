// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type ComboBoxDelegate interface {
	TextFieldDelegate
	ImplementsComboBoxSelectionDidChange() bool
	// optional
	ComboBoxSelectionDidChange(notification foundation.Notification)
	ImplementsComboBoxSelectionIsChanging() bool
	// optional
	ComboBoxSelectionIsChanging(notification foundation.Notification)
	ImplementsComboBoxWillDismiss() bool
	// optional
	ComboBoxWillDismiss(notification foundation.Notification)
	ImplementsComboBoxWillPopUp() bool
	// optional
	ComboBoxWillPopUp(notification foundation.Notification)
}

func WrapComboBoxDelegate(v ComboBoxDelegate) objc.Object {
	return objc.WrapAsProtocol("NSComboBoxDelegate", v)
}

type ComboBoxDelegateBase struct {
	TextFieldDelegateBase
}

func (p *ComboBoxDelegateBase) ImplementsComboBoxSelectionDidChange() bool {
	return false
}

func (p *ComboBoxDelegateBase) ComboBoxSelectionDidChange(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ComboBoxDelegateBase) ImplementsComboBoxSelectionIsChanging() bool {
	return false
}

func (p *ComboBoxDelegateBase) ComboBoxSelectionIsChanging(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ComboBoxDelegateBase) ImplementsComboBoxWillDismiss() bool {
	return false
}

func (p *ComboBoxDelegateBase) ComboBoxWillDismiss(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ComboBoxDelegateBase) ImplementsComboBoxWillPopUp() bool {
	return false
}

func (p *ComboBoxDelegateBase) ComboBoxWillPopUp(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

type ComboBoxDelegateWrapper struct {
	objc.Object
}

func (c_ ComboBoxDelegateWrapper) ComboBoxSelectionDidChange(notification foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("comboBoxSelectionDidChange:"), objc.ExtractPtr(notification))
}

func (c_ ComboBoxDelegateWrapper) ComboBoxSelectionIsChanging(notification foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("comboBoxSelectionIsChanging:"), objc.ExtractPtr(notification))
}

func (c_ ComboBoxDelegateWrapper) ComboBoxWillDismiss(notification foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("comboBoxWillDismiss:"), objc.ExtractPtr(notification))
}

func (c_ ComboBoxDelegateWrapper) ComboBoxWillPopUp(notification foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("comboBoxWillPopUp:"), objc.ExtractPtr(notification))
}

func (c_ ComboBoxDelegateWrapper) TextField_TextView_Candidates_ForSelectedRange(textField ITextField, textView ITextView, candidates []foundation.ITextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult {
	rv := objc.CallMethod[[]foundation.TextCheckingResult](c_, objc.GetSelector("textField:textView:candidates:forSelectedRange:"), objc.ExtractPtr(textField), objc.ExtractPtr(textView), candidates, selectedRange)
	return rv
}

func (c_ ComboBoxDelegateWrapper) TextField_TextView_CandidatesForSelectedRange(textField ITextField, textView ITextView, selectedRange foundation.Range) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](c_, objc.GetSelector("textField:textView:candidatesForSelectedRange:"), objc.ExtractPtr(textField), objc.ExtractPtr(textView), selectedRange)
	return rv
}

func (c_ ComboBoxDelegateWrapper) TextField_TextView_ShouldSelectCandidateAtIndex(textField ITextField, textView ITextView, index uint) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("textField:textView:shouldSelectCandidateAtIndex:"), objc.ExtractPtr(textField), objc.ExtractPtr(textView), index)
	return rv
}

func (c_ ComboBoxDelegateWrapper) Control_IsValidObject(control IControl, obj objc.IObject) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:isValidObject:"), objc.ExtractPtr(control), objc.ExtractPtr(obj))
	return rv
}

func (c_ ComboBoxDelegateWrapper) Control_DidFailToValidatePartialString_ErrorDescription(control IControl, string_ string, error string) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("control:didFailToValidatePartialString:errorDescription:"), objc.ExtractPtr(control), string_, error)
}

func (c_ ComboBoxDelegateWrapper) Control_DidFailToFormatString_ErrorDescription(control IControl, string_ string, error string) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:didFailToFormatString:errorDescription:"), objc.ExtractPtr(control), string_, error)
	return rv
}

func (c_ ComboBoxDelegateWrapper) Control_TextShouldBeginEditing(control IControl, fieldEditor IText) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:textShouldBeginEditing:"), objc.ExtractPtr(control), objc.ExtractPtr(fieldEditor))
	return rv
}

func (c_ ComboBoxDelegateWrapper) Control_TextShouldEndEditing(control IControl, fieldEditor IText) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:textShouldEndEditing:"), objc.ExtractPtr(control), objc.ExtractPtr(fieldEditor))
	return rv
}

func (c_ ComboBoxDelegateWrapper) Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(control IControl, textView ITextView, words []string, charRange foundation.Range, index *int) []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("control:textView:completions:forPartialWordRange:indexOfSelectedItem:"), objc.ExtractPtr(control), objc.ExtractPtr(textView), words, charRange, index)
	return rv
}

func (c_ ComboBoxDelegateWrapper) Control_TextView_DoCommandBySelector(control IControl, textView ITextView, commandSelector objc.Selector) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:textView:doCommandBySelector:"), objc.ExtractPtr(control), objc.ExtractPtr(textView), commandSelector)
	return rv
}

func (c_ ComboBoxDelegateWrapper) ControlTextDidBeginEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("controlTextDidBeginEditing:"), objc.ExtractPtr(obj))
}

func (c_ ComboBoxDelegateWrapper) ControlTextDidChange(obj foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("controlTextDidChange:"), objc.ExtractPtr(obj))
}

func (c_ ComboBoxDelegateWrapper) ControlTextDidEndEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("controlTextDidEndEditing:"), objc.ExtractPtr(obj))
}
