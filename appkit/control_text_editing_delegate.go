// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type ControlTextEditingDelegate interface {
	ImplementsControl_IsValidObject() bool
	// optional
	Control_IsValidObject(control Control, obj objc.Object) bool
	ImplementsControl_DidFailToValidatePartialString_ErrorDescription() bool
	// optional
	Control_DidFailToValidatePartialString_ErrorDescription(control Control, string_ string, error string)
	ImplementsControl_DidFailToFormatString_ErrorDescription() bool
	// optional
	Control_DidFailToFormatString_ErrorDescription(control Control, string_ string, error string) bool
	ImplementsControl_TextShouldBeginEditing() bool
	// optional
	Control_TextShouldBeginEditing(control Control, fieldEditor Text) bool
	ImplementsControl_TextShouldEndEditing() bool
	// optional
	Control_TextShouldEndEditing(control Control, fieldEditor Text) bool
	ImplementsControl_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem() bool
	// optional
	Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(control Control, textView TextView, words []string, charRange foundation.Range, index *int) []string
	ImplementsControl_TextView_DoCommandBySelector() bool
	// optional
	Control_TextView_DoCommandBySelector(control Control, textView TextView, commandSelector objc.Selector) bool
	ImplementsControlTextDidBeginEditing() bool
	// optional
	ControlTextDidBeginEditing(obj foundation.Notification)
	ImplementsControlTextDidChange() bool
	// optional
	ControlTextDidChange(obj foundation.Notification)
	ImplementsControlTextDidEndEditing() bool
	// optional
	ControlTextDidEndEditing(obj foundation.Notification)
}

type ControlTextEditingDelegateImpl struct {
	_Control_IsValidObject                                                func(control Control, obj objc.Object) bool
	_Control_DidFailToValidatePartialString_ErrorDescription              func(control Control, string_ string, error string)
	_Control_DidFailToFormatString_ErrorDescription                       func(control Control, string_ string, error string) bool
	_Control_TextShouldBeginEditing                                       func(control Control, fieldEditor Text) bool
	_Control_TextShouldEndEditing                                         func(control Control, fieldEditor Text) bool
	_Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem func(control Control, textView TextView, words []string, charRange foundation.Range, index *int) []string
	_Control_TextView_DoCommandBySelector                                 func(control Control, textView TextView, commandSelector objc.Selector) bool
	_ControlTextDidBeginEditing                                           func(obj foundation.Notification)
	_ControlTextDidChange                                                 func(obj foundation.Notification)
	_ControlTextDidEndEditing                                             func(obj foundation.Notification)
}

func (di *ControlTextEditingDelegateImpl) ImplementsControl_IsValidObject() bool {
	return di._Control_IsValidObject != nil
}

func (di *ControlTextEditingDelegateImpl) SetControl_IsValidObject(f func(control Control, obj objc.Object) bool) {
	di._Control_IsValidObject = f
}

func (di *ControlTextEditingDelegateImpl) Control_IsValidObject(control Control, obj objc.Object) bool {
	return di._Control_IsValidObject(control, obj)
}
func (di *ControlTextEditingDelegateImpl) ImplementsControl_DidFailToValidatePartialString_ErrorDescription() bool {
	return di._Control_DidFailToValidatePartialString_ErrorDescription != nil
}

func (di *ControlTextEditingDelegateImpl) SetControl_DidFailToValidatePartialString_ErrorDescription(f func(control Control, string_ string, error string)) {
	di._Control_DidFailToValidatePartialString_ErrorDescription = f
}

func (di *ControlTextEditingDelegateImpl) Control_DidFailToValidatePartialString_ErrorDescription(control Control, string_ string, error string) {
	di._Control_DidFailToValidatePartialString_ErrorDescription(control, string_, error)
}
func (di *ControlTextEditingDelegateImpl) ImplementsControl_DidFailToFormatString_ErrorDescription() bool {
	return di._Control_DidFailToFormatString_ErrorDescription != nil
}

func (di *ControlTextEditingDelegateImpl) SetControl_DidFailToFormatString_ErrorDescription(f func(control Control, string_ string, error string) bool) {
	di._Control_DidFailToFormatString_ErrorDescription = f
}

func (di *ControlTextEditingDelegateImpl) Control_DidFailToFormatString_ErrorDescription(control Control, string_ string, error string) bool {
	return di._Control_DidFailToFormatString_ErrorDescription(control, string_, error)
}
func (di *ControlTextEditingDelegateImpl) ImplementsControl_TextShouldBeginEditing() bool {
	return di._Control_TextShouldBeginEditing != nil
}

func (di *ControlTextEditingDelegateImpl) SetControl_TextShouldBeginEditing(f func(control Control, fieldEditor Text) bool) {
	di._Control_TextShouldBeginEditing = f
}

func (di *ControlTextEditingDelegateImpl) Control_TextShouldBeginEditing(control Control, fieldEditor Text) bool {
	return di._Control_TextShouldBeginEditing(control, fieldEditor)
}
func (di *ControlTextEditingDelegateImpl) ImplementsControl_TextShouldEndEditing() bool {
	return di._Control_TextShouldEndEditing != nil
}

func (di *ControlTextEditingDelegateImpl) SetControl_TextShouldEndEditing(f func(control Control, fieldEditor Text) bool) {
	di._Control_TextShouldEndEditing = f
}

func (di *ControlTextEditingDelegateImpl) Control_TextShouldEndEditing(control Control, fieldEditor Text) bool {
	return di._Control_TextShouldEndEditing(control, fieldEditor)
}
func (di *ControlTextEditingDelegateImpl) ImplementsControl_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem() bool {
	return di._Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem != nil
}

func (di *ControlTextEditingDelegateImpl) SetControl_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(f func(control Control, textView TextView, words []string, charRange foundation.Range, index *int) []string) {
	di._Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem = f
}

func (di *ControlTextEditingDelegateImpl) Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(control Control, textView TextView, words []string, charRange foundation.Range, index *int) []string {
	return di._Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(control, textView, words, charRange, index)
}
func (di *ControlTextEditingDelegateImpl) ImplementsControl_TextView_DoCommandBySelector() bool {
	return di._Control_TextView_DoCommandBySelector != nil
}

func (di *ControlTextEditingDelegateImpl) SetControl_TextView_DoCommandBySelector(f func(control Control, textView TextView, commandSelector objc.Selector) bool) {
	di._Control_TextView_DoCommandBySelector = f
}

func (di *ControlTextEditingDelegateImpl) Control_TextView_DoCommandBySelector(control Control, textView TextView, commandSelector objc.Selector) bool {
	return di._Control_TextView_DoCommandBySelector(control, textView, commandSelector)
}
func (di *ControlTextEditingDelegateImpl) ImplementsControlTextDidBeginEditing() bool {
	return di._ControlTextDidBeginEditing != nil
}

func (di *ControlTextEditingDelegateImpl) SetControlTextDidBeginEditing(f func(obj foundation.Notification)) {
	di._ControlTextDidBeginEditing = f
}

func (di *ControlTextEditingDelegateImpl) ControlTextDidBeginEditing(obj foundation.Notification) {
	di._ControlTextDidBeginEditing(obj)
}
func (di *ControlTextEditingDelegateImpl) ImplementsControlTextDidChange() bool {
	return di._ControlTextDidChange != nil
}

func (di *ControlTextEditingDelegateImpl) SetControlTextDidChange(f func(obj foundation.Notification)) {
	di._ControlTextDidChange = f
}

func (di *ControlTextEditingDelegateImpl) ControlTextDidChange(obj foundation.Notification) {
	di._ControlTextDidChange(obj)
}
func (di *ControlTextEditingDelegateImpl) ImplementsControlTextDidEndEditing() bool {
	return di._ControlTextDidEndEditing != nil
}

func (di *ControlTextEditingDelegateImpl) SetControlTextDidEndEditing(f func(obj foundation.Notification)) {
	di._ControlTextDidEndEditing = f
}

func (di *ControlTextEditingDelegateImpl) ControlTextDidEndEditing(obj foundation.Notification) {
	di._ControlTextDidEndEditing(obj)
}

type ControlTextEditingDelegateWrapper struct {
	objc.Object
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControl_IsValidObject() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:isValidObject:"))
}

func (c_ ControlTextEditingDelegateWrapper) Control_IsValidObject(control IControl, obj objc.IObject) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:isValidObject:"), objc.ExtractPtr(control), objc.ExtractPtr(obj))
	return rv
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControl_DidFailToValidatePartialString_ErrorDescription() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:didFailToValidatePartialString:errorDescription:"))
}

func (c_ ControlTextEditingDelegateWrapper) Control_DidFailToValidatePartialString_ErrorDescription(control IControl, string_ string, error string) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("control:didFailToValidatePartialString:errorDescription:"), objc.ExtractPtr(control), string_, error)
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControl_DidFailToFormatString_ErrorDescription() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:didFailToFormatString:errorDescription:"))
}

func (c_ ControlTextEditingDelegateWrapper) Control_DidFailToFormatString_ErrorDescription(control IControl, string_ string, error string) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:didFailToFormatString:errorDescription:"), objc.ExtractPtr(control), string_, error)
	return rv
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControl_TextShouldBeginEditing() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:textShouldBeginEditing:"))
}

func (c_ ControlTextEditingDelegateWrapper) Control_TextShouldBeginEditing(control IControl, fieldEditor IText) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:textShouldBeginEditing:"), objc.ExtractPtr(control), objc.ExtractPtr(fieldEditor))
	return rv
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControl_TextShouldEndEditing() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:textShouldEndEditing:"))
}

func (c_ ControlTextEditingDelegateWrapper) Control_TextShouldEndEditing(control IControl, fieldEditor IText) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:textShouldEndEditing:"), objc.ExtractPtr(control), objc.ExtractPtr(fieldEditor))
	return rv
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControl_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:textView:completions:forPartialWordRange:indexOfSelectedItem:"))
}

func (c_ ControlTextEditingDelegateWrapper) Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(control IControl, textView ITextView, words []string, charRange foundation.Range, index *int) []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("control:textView:completions:forPartialWordRange:indexOfSelectedItem:"), objc.ExtractPtr(control), objc.ExtractPtr(textView), words, charRange, index)
	return rv
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControl_TextView_DoCommandBySelector() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:textView:doCommandBySelector:"))
}

func (c_ ControlTextEditingDelegateWrapper) Control_TextView_DoCommandBySelector(control IControl, textView ITextView, commandSelector objc.Selector) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:textView:doCommandBySelector:"), objc.ExtractPtr(control), objc.ExtractPtr(textView), commandSelector)
	return rv
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControlTextDidBeginEditing() bool {
	return c_.RespondsToSelector(objc.GetSelector("controlTextDidBeginEditing:"))
}

func (c_ ControlTextEditingDelegateWrapper) ControlTextDidBeginEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("controlTextDidBeginEditing:"), objc.ExtractPtr(obj))
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControlTextDidChange() bool {
	return c_.RespondsToSelector(objc.GetSelector("controlTextDidChange:"))
}

func (c_ ControlTextEditingDelegateWrapper) ControlTextDidChange(obj foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("controlTextDidChange:"), objc.ExtractPtr(obj))
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControlTextDidEndEditing() bool {
	return c_.RespondsToSelector(objc.GetSelector("controlTextDidEndEditing:"))
}

func (c_ ControlTextEditingDelegateWrapper) ControlTextDidEndEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("controlTextDidEndEditing:"), objc.ExtractPtr(obj))
}
