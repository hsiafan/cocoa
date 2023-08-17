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

func WrapControlTextEditingDelegate(v ControlTextEditingDelegate) objc.Object {
	return objc.WrapAsProtocol("NSControlTextEditingDelegate", v)
}

type ControlTextEditingDelegateBase struct {
}

func (p *ControlTextEditingDelegateBase) ImplementsControl_IsValidObject() bool {
	return false
}

func (p *ControlTextEditingDelegateBase) Control_IsValidObject(control Control, obj objc.Object) bool {
	panic("unimpemented protocol method")
}

func (p *ControlTextEditingDelegateBase) ImplementsControl_DidFailToValidatePartialString_ErrorDescription() bool {
	return false
}

func (p *ControlTextEditingDelegateBase) Control_DidFailToValidatePartialString_ErrorDescription(control Control, string_ string, error string) {
	panic("unimpemented protocol method")
}

func (p *ControlTextEditingDelegateBase) ImplementsControl_DidFailToFormatString_ErrorDescription() bool {
	return false
}

func (p *ControlTextEditingDelegateBase) Control_DidFailToFormatString_ErrorDescription(control Control, string_ string, error string) bool {
	panic("unimpemented protocol method")
}

func (p *ControlTextEditingDelegateBase) ImplementsControl_TextShouldBeginEditing() bool {
	return false
}

func (p *ControlTextEditingDelegateBase) Control_TextShouldBeginEditing(control Control, fieldEditor Text) bool {
	panic("unimpemented protocol method")
}

func (p *ControlTextEditingDelegateBase) ImplementsControl_TextShouldEndEditing() bool {
	return false
}

func (p *ControlTextEditingDelegateBase) Control_TextShouldEndEditing(control Control, fieldEditor Text) bool {
	panic("unimpemented protocol method")
}

func (p *ControlTextEditingDelegateBase) ImplementsControl_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem() bool {
	return false
}

func (p *ControlTextEditingDelegateBase) Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(control Control, textView TextView, words []string, charRange foundation.Range, index *int) []string {
	panic("unimpemented protocol method")
}

func (p *ControlTextEditingDelegateBase) ImplementsControl_TextView_DoCommandBySelector() bool {
	return false
}

func (p *ControlTextEditingDelegateBase) Control_TextView_DoCommandBySelector(control Control, textView TextView, commandSelector objc.Selector) bool {
	panic("unimpemented protocol method")
}

func (p *ControlTextEditingDelegateBase) ImplementsControlTextDidBeginEditing() bool {
	return false
}

func (p *ControlTextEditingDelegateBase) ControlTextDidBeginEditing(obj foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ControlTextEditingDelegateBase) ImplementsControlTextDidChange() bool {
	return false
}

func (p *ControlTextEditingDelegateBase) ControlTextDidChange(obj foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ControlTextEditingDelegateBase) ImplementsControlTextDidEndEditing() bool {
	return false
}

func (p *ControlTextEditingDelegateBase) ControlTextDidEndEditing(obj foundation.Notification) {
	panic("unimpemented protocol method")
}

type ControlTextEditingDelegateWrapper struct {
	objc.Object
}

func (c_ ControlTextEditingDelegateWrapper) Control_IsValidObject(control IControl, obj objc.IObject) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:isValidObject:"), objc.ExtractPtr(control), objc.ExtractPtr(obj))
	return rv
}

func (c_ ControlTextEditingDelegateWrapper) Control_DidFailToValidatePartialString_ErrorDescription(control IControl, string_ string, error string) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("control:didFailToValidatePartialString:errorDescription:"), objc.ExtractPtr(control), string_, error)
}

func (c_ ControlTextEditingDelegateWrapper) Control_DidFailToFormatString_ErrorDescription(control IControl, string_ string, error string) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:didFailToFormatString:errorDescription:"), objc.ExtractPtr(control), string_, error)
	return rv
}

func (c_ ControlTextEditingDelegateWrapper) Control_TextShouldBeginEditing(control IControl, fieldEditor IText) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:textShouldBeginEditing:"), objc.ExtractPtr(control), objc.ExtractPtr(fieldEditor))
	return rv
}

func (c_ ControlTextEditingDelegateWrapper) Control_TextShouldEndEditing(control IControl, fieldEditor IText) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:textShouldEndEditing:"), objc.ExtractPtr(control), objc.ExtractPtr(fieldEditor))
	return rv
}

func (c_ ControlTextEditingDelegateWrapper) Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(control IControl, textView ITextView, words []string, charRange foundation.Range, index *int) []string {
	rv := objc.CallMethod[[]string](c_, objc.GetSelector("control:textView:completions:forPartialWordRange:indexOfSelectedItem:"), objc.ExtractPtr(control), objc.ExtractPtr(textView), words, charRange, index)
	return rv
}

func (c_ ControlTextEditingDelegateWrapper) Control_TextView_DoCommandBySelector(control IControl, textView ITextView, commandSelector objc.Selector) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("control:textView:doCommandBySelector:"), objc.ExtractPtr(control), objc.ExtractPtr(textView), commandSelector)
	return rv
}

func (c_ ControlTextEditingDelegateWrapper) ControlTextDidBeginEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("controlTextDidBeginEditing:"), objc.ExtractPtr(obj))
}

func (c_ ControlTextEditingDelegateWrapper) ControlTextDidChange(obj foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("controlTextDidChange:"), objc.ExtractPtr(obj))
}

func (c_ ControlTextEditingDelegateWrapper) ControlTextDidEndEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("controlTextDidEndEditing:"), objc.ExtractPtr(obj))
}
