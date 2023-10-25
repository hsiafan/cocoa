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

type ControlTextEditingDelegateCreator struct {
	className string
	class     objc.Class
}

func NewControlTextEditingDelegateCreator(name string) *ControlTextEditingDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &ControlTextEditingDelegateCreator{className: name, class: class}
}

func (c *ControlTextEditingDelegateCreator) SetControl_IsValidObject(handle func(o objc.Object, control Control, obj objc.Object) bool) {
	objc.AddMethod(c.class, objc.SEL("control:isValidObject:"), handle)
}

func (c *ControlTextEditingDelegateCreator) SetControl_DidFailToValidatePartialString_ErrorDescription(handle func(o objc.Object, control Control, string_ string, error string)) {
	objc.AddMethod(c.class, objc.SEL("control:didFailToValidatePartialString:errorDescription:"), handle)
}

func (c *ControlTextEditingDelegateCreator) SetControl_DidFailToFormatString_ErrorDescription(handle func(o objc.Object, control Control, string_ string, error string) bool) {
	objc.AddMethod(c.class, objc.SEL("control:didFailToFormatString:errorDescription:"), handle)
}

func (c *ControlTextEditingDelegateCreator) SetControl_TextShouldBeginEditing(handle func(o objc.Object, control Control, fieldEditor Text) bool) {
	objc.AddMethod(c.class, objc.SEL("control:textShouldBeginEditing:"), handle)
}

func (c *ControlTextEditingDelegateCreator) SetControl_TextShouldEndEditing(handle func(o objc.Object, control Control, fieldEditor Text) bool) {
	objc.AddMethod(c.class, objc.SEL("control:textShouldEndEditing:"), handle)
}

func (c *ControlTextEditingDelegateCreator) SetControl_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(handle func(o objc.Object, control Control, textView TextView, words []string, charRange foundation.Range, index *int) []string) {
	objc.AddMethod(c.class, objc.SEL("control:textView:completions:forPartialWordRange:indexOfSelectedItem:"), handle)
}

func (c *ControlTextEditingDelegateCreator) SetControl_TextView_DoCommandBySelector(handle func(o objc.Object, control Control, textView TextView, commandSelector objc.Selector) bool) {
	objc.AddMethod(c.class, objc.SEL("control:textView:doCommandBySelector:"), handle)
}

func (c *ControlTextEditingDelegateCreator) SetControlTextDidBeginEditing(handle func(o objc.Object, obj foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("controlTextDidBeginEditing:"), handle)
}

func (c *ControlTextEditingDelegateCreator) SetControlTextDidChange(handle func(o objc.Object, obj foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("controlTextDidChange:"), handle)
}

func (c *ControlTextEditingDelegateCreator) SetControlTextDidEndEditing(handle func(o objc.Object, obj foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("controlTextDidEndEditing:"), handle)
}

func (c *ControlTextEditingDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
