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

func WrapTokenFieldCellDelegate(v TokenFieldCellDelegate) objc.Object {
	return objc.WrapAsProtocol("NSTokenFieldCellDelegate", v)
}

type TokenFieldCellDelegateBase struct {
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_DisplayStringForRepresentedObject() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_DisplayStringForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) string {
	panic("unimpemented protocol method")
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_StyleForRepresentedObject() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_StyleForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) TokenStyle {
	panic("unimpemented protocol method")
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(tokenFieldCell TokenFieldCell, substring string, tokenIndex int, selectedIndex *int) []objc.IObject {
	panic("unimpemented protocol method")
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_EditingStringForRepresentedObject() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_EditingStringForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) string {
	panic("unimpemented protocol method")
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_RepresentedObjectForEditingString() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_RepresentedObjectForEditingString(tokenFieldCell TokenFieldCell, editingString string) objc.IObject {
	panic("unimpemented protocol method")
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_ShouldAddObjects_AtIndex() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_ShouldAddObjects_AtIndex(tokenFieldCell TokenFieldCell, tokens []objc.Object, index uint) []objc.IObject {
	panic("unimpemented protocol method")
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_ReadFromPasteboard() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_ReadFromPasteboard(tokenFieldCell TokenFieldCell, pboard Pasteboard) []objc.IObject {
	panic("unimpemented protocol method")
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_WriteRepresentedObjects_ToPasteboard() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_WriteRepresentedObjects_ToPasteboard(tokenFieldCell TokenFieldCell, objects []objc.Object, pboard Pasteboard) bool {
	panic("unimpemented protocol method")
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_HasMenuForRepresentedObject() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_HasMenuForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) bool {
	panic("unimpemented protocol method")
}

func (p *TokenFieldCellDelegateBase) ImplementsTokenFieldCell_MenuForRepresentedObject() bool {
	return false
}

func (p *TokenFieldCellDelegateBase) TokenFieldCell_MenuForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) IMenu {
	panic("unimpemented protocol method")
}

type TokenFieldCellDelegateCreator struct {
	className string
	class     objc.Class
}

func NewTokenFieldCellDelegateCreator(name string) *TokenFieldCellDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &TokenFieldCellDelegateCreator{className: name, class: class}
}

func (c *TokenFieldCellDelegateCreator) SetTokenFieldCell_DisplayStringForRepresentedObject(handle func(o objc.Object, tokenFieldCell TokenFieldCell, representedObject objc.Object) string) {
	objc.AddMethod(c.class, objc.GetSelector("tokenFieldCell:displayStringForRepresentedObject:"), handle)
}

func (c *TokenFieldCellDelegateCreator) SetTokenFieldCell_StyleForRepresentedObject(handle func(o objc.Object, tokenFieldCell TokenFieldCell, representedObject objc.Object) TokenStyle) {
	objc.AddMethod(c.class, objc.GetSelector("tokenFieldCell:styleForRepresentedObject:"), handle)
}

func (c *TokenFieldCellDelegateCreator) SetTokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(handle func(o objc.Object, tokenFieldCell TokenFieldCell, substring string, tokenIndex int, selectedIndex *int) []objc.IObject) {
	objc.AddMethod(c.class, objc.GetSelector("tokenFieldCell:completionsForSubstring:indexOfToken:indexOfSelectedItem:"), handle)
}

func (c *TokenFieldCellDelegateCreator) SetTokenFieldCell_EditingStringForRepresentedObject(handle func(o objc.Object, tokenFieldCell TokenFieldCell, representedObject objc.Object) string) {
	objc.AddMethod(c.class, objc.GetSelector("tokenFieldCell:editingStringForRepresentedObject:"), handle)
}

func (c *TokenFieldCellDelegateCreator) SetTokenFieldCell_RepresentedObjectForEditingString(handle func(o objc.Object, tokenFieldCell TokenFieldCell, editingString string) objc.IObject) {
	objc.AddMethod(c.class, objc.GetSelector("tokenFieldCell:representedObjectForEditingString:"), handle)
}

func (c *TokenFieldCellDelegateCreator) SetTokenFieldCell_ShouldAddObjects_AtIndex(handle func(o objc.Object, tokenFieldCell TokenFieldCell, tokens []objc.Object, index uint) []objc.IObject) {
	objc.AddMethod(c.class, objc.GetSelector("tokenFieldCell:shouldAddObjects:atIndex:"), handle)
}

func (c *TokenFieldCellDelegateCreator) SetTokenFieldCell_ReadFromPasteboard(handle func(o objc.Object, tokenFieldCell TokenFieldCell, pboard Pasteboard) []objc.IObject) {
	objc.AddMethod(c.class, objc.GetSelector("tokenFieldCell:readFromPasteboard:"), handle)
}

func (c *TokenFieldCellDelegateCreator) SetTokenFieldCell_WriteRepresentedObjects_ToPasteboard(handle func(o objc.Object, tokenFieldCell TokenFieldCell, objects []objc.Object, pboard Pasteboard) bool) {
	objc.AddMethod(c.class, objc.GetSelector("tokenFieldCell:writeRepresentedObjects:toPasteboard:"), handle)
}

func (c *TokenFieldCellDelegateCreator) SetTokenFieldCell_HasMenuForRepresentedObject(handle func(o objc.Object, tokenFieldCell TokenFieldCell, representedObject objc.Object) bool) {
	objc.AddMethod(c.class, objc.GetSelector("tokenFieldCell:hasMenuForRepresentedObject:"), handle)
}

func (c *TokenFieldCellDelegateCreator) SetTokenFieldCell_MenuForRepresentedObject(handle func(o objc.Object, tokenFieldCell TokenFieldCell, representedObject objc.Object) IMenu) {
	objc.AddMethod(c.class, objc.GetSelector("tokenFieldCell:menuForRepresentedObject:"), handle)
}

func (c *TokenFieldCellDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
