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

type TokenFieldDelegateCreator struct {
	className string
	class     objc.Class
}

func NewTokenFieldDelegateCreator(name string) *TokenFieldDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &TokenFieldDelegateCreator{className: name, class: class}
}

func (c *TokenFieldDelegateCreator) SetTokenField_DisplayStringForRepresentedObject(handle func(o objc.Object, tokenField TokenField, representedObject objc.Object) string) {
	objc.AddMethod(c.class, objc.GetSelector("tokenField:displayStringForRepresentedObject:"), handle)
}

func (c *TokenFieldDelegateCreator) SetTokenField_StyleForRepresentedObject(handle func(o objc.Object, tokenField TokenField, representedObject objc.Object) TokenStyle) {
	objc.AddMethod(c.class, objc.GetSelector("tokenField:styleForRepresentedObject:"), handle)
}

func (c *TokenFieldDelegateCreator) SetTokenField_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(handle func(o objc.Object, tokenField TokenField, substring string, tokenIndex int, selectedIndex *int) []objc.IObject) {
	objc.AddMethod(c.class, objc.GetSelector("tokenField:completionsForSubstring:indexOfToken:indexOfSelectedItem:"), handle)
}

func (c *TokenFieldDelegateCreator) SetTokenField_EditingStringForRepresentedObject(handle func(o objc.Object, tokenField TokenField, representedObject objc.Object) string) {
	objc.AddMethod(c.class, objc.GetSelector("tokenField:editingStringForRepresentedObject:"), handle)
}

func (c *TokenFieldDelegateCreator) SetTokenField_RepresentedObjectForEditingString(handle func(o objc.Object, tokenField TokenField, editingString string) objc.IObject) {
	objc.AddMethod(c.class, objc.GetSelector("tokenField:representedObjectForEditingString:"), handle)
}

func (c *TokenFieldDelegateCreator) SetTokenField_ShouldAddObjects_AtIndex(handle func(o objc.Object, tokenField TokenField, tokens []objc.Object, index uint) []objc.IObject) {
	objc.AddMethod(c.class, objc.GetSelector("tokenField:shouldAddObjects:atIndex:"), handle)
}

func (c *TokenFieldDelegateCreator) SetTokenField_ReadFromPasteboard(handle func(o objc.Object, tokenField TokenField, pboard Pasteboard) []objc.IObject) {
	objc.AddMethod(c.class, objc.GetSelector("tokenField:readFromPasteboard:"), handle)
}

func (c *TokenFieldDelegateCreator) SetTokenField_WriteRepresentedObjects_ToPasteboard(handle func(o objc.Object, tokenField TokenField, objects []objc.Object, pboard Pasteboard) bool) {
	objc.AddMethod(c.class, objc.GetSelector("tokenField:writeRepresentedObjects:toPasteboard:"), handle)
}

func (c *TokenFieldDelegateCreator) SetTokenField_HasMenuForRepresentedObject(handle func(o objc.Object, tokenField TokenField, representedObject objc.Object) bool) {
	objc.AddMethod(c.class, objc.GetSelector("tokenField:hasMenuForRepresentedObject:"), handle)
}

func (c *TokenFieldDelegateCreator) SetTokenField_MenuForRepresentedObject(handle func(o objc.Object, tokenField TokenField, representedObject objc.Object) IMenu) {
	objc.AddMethod(c.class, objc.GetSelector("tokenField:menuForRepresentedObject:"), handle)
}

func (c *TokenFieldDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
