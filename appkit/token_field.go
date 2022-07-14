// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var TokenFieldClass = _TokenFieldClass{objc.GetClass("NSTokenField")}

type _TokenFieldClass struct {
	objc.Class
}

type ITokenField interface {
	ITextField
	TokenStyle() TokenStyle
	SetTokenStyle(value TokenStyle)
	TokenizingCharacterSet() foundation.CharacterSet
	SetTokenizingCharacterSet(value foundation.ICharacterSet)
	CompletionDelay() foundation.TimeInterval
	SetCompletionDelay(value foundation.TimeInterval)
}

type TokenField struct {
	TextField
}

func MakeTokenField(ptr unsafe.Pointer) TokenField {
	return TokenField{
		TextField: MakeTextField(ptr),
	}
}

func (tc _TokenFieldClass) LabelWithAttributedString(attributedStringValue foundation.IAttributedString) TokenField {
	rv := ffi.CallMethod[TokenField](tc, "labelWithAttributedString:", attributedStringValue)
	return rv
}

func (tc _TokenFieldClass) LabelWithString(stringValue string) TokenField {
	rv := ffi.CallMethod[TokenField](tc, "labelWithString:", stringValue)
	return rv
}

func (tc _TokenFieldClass) TextFieldWithString(stringValue string) TokenField {
	rv := ffi.CallMethod[TokenField](tc, "textFieldWithString:", stringValue)
	return rv
}

func (tc _TokenFieldClass) WrappingLabelWithString(stringValue string) TokenField {
	rv := ffi.CallMethod[TokenField](tc, "wrappingLabelWithString:", stringValue)
	return rv
}

func (t_ TokenField) InitWithFrame(frameRect foundation.Rect) TokenField {
	rv := ffi.CallMethod[TokenField](t_, "initWithFrame:", frameRect)
	rv.Autorelease()
	return rv
}

func (t_ TokenField) Init() TokenField {
	rv := ffi.CallMethod[TokenField](t_, "init")
	rv.Autorelease()
	return rv
}

func (tc _TokenFieldClass) Alloc() TokenField {
	rv := ffi.CallMethod[TokenField](tc, "alloc")
	return rv
}

func (tc _TokenFieldClass) New() TokenField {
	rv := ffi.CallMethod[TokenField](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTokenField() TokenField {
	return TokenFieldClass.New()
}

func (t_ TokenField) TokenStyle() TokenStyle {
	rv := ffi.CallMethod[TokenStyle](t_, "tokenStyle")
	return rv
}

func (t_ TokenField) SetTokenStyle(value TokenStyle) {
	ffi.CallMethod[ffi.Void](t_, "setTokenStyle:", value)
}

func (t_ TokenField) TokenizingCharacterSet() foundation.CharacterSet {
	rv := ffi.CallMethod[foundation.CharacterSet](t_, "tokenizingCharacterSet")
	return rv
}

func (t_ TokenField) SetTokenizingCharacterSet(value foundation.ICharacterSet) {
	ffi.CallMethod[ffi.Void](t_, "setTokenizingCharacterSet:", value)
}

func (tc _TokenFieldClass) DefaultTokenizingCharacterSet() foundation.CharacterSet {
	rv := ffi.CallMethod[foundation.CharacterSet](tc, "defaultTokenizingCharacterSet")
	return rv
}

func (t_ TokenField) CompletionDelay() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](t_, "completionDelay")
	return rv
}

func (t_ TokenField) SetCompletionDelay(value foundation.TimeInterval) {
	ffi.CallMethod[ffi.Void](t_, "setCompletionDelay:", value)
}

func (tc _TokenFieldClass) DefaultCompletionDelay() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](tc, "defaultCompletionDelay")
	return rv
}

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
	rv := ffi.CallMethod[string](t_, "tokenField:displayStringForRepresentedObject:", tokenField, representedObject)
	return rv
}

func (t_ *TokenFieldDelegateWrapper) ImplementsTokenField_StyleForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:styleForRepresentedObject:"))
}

func (t_ TokenFieldDelegateWrapper) TokenField_StyleForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) TokenStyle {
	rv := ffi.CallMethod[TokenStyle](t_, "tokenField:styleForRepresentedObject:", tokenField, representedObject)
	return rv
}

func (t_ *TokenFieldDelegateWrapper) ImplementsTokenField_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:completionsForSubstring:indexOfToken:indexOfSelectedItem:"))
}

func (t_ TokenFieldDelegateWrapper) TokenField_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(tokenField ITokenField, substring string, tokenIndex int, selectedIndex *int) []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](t_, "tokenField:completionsForSubstring:indexOfToken:indexOfSelectedItem:", tokenField, substring, tokenIndex, selectedIndex)
	return rv
}

func (t_ *TokenFieldDelegateWrapper) ImplementsTokenField_EditingStringForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:editingStringForRepresentedObject:"))
}

func (t_ TokenFieldDelegateWrapper) TokenField_EditingStringForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) string {
	rv := ffi.CallMethod[string](t_, "tokenField:editingStringForRepresentedObject:", tokenField, representedObject)
	return rv
}

func (t_ *TokenFieldDelegateWrapper) ImplementsTokenField_RepresentedObjectForEditingString() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:representedObjectForEditingString:"))
}

func (t_ TokenFieldDelegateWrapper) TokenField_RepresentedObjectForEditingString(tokenField ITokenField, editingString string) objc.Object {
	rv := ffi.CallMethod[objc.Object](t_, "tokenField:representedObjectForEditingString:", tokenField, editingString)
	return rv
}

func (t_ *TokenFieldDelegateWrapper) ImplementsTokenField_ShouldAddObjects_AtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:shouldAddObjects:atIndex:"))
}

func (t_ TokenFieldDelegateWrapper) TokenField_ShouldAddObjects_AtIndex(tokenField ITokenField, tokens []objc.IObject, index uint) []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](t_, "tokenField:shouldAddObjects:atIndex:", tokenField, tokens, index)
	return rv
}

func (t_ *TokenFieldDelegateWrapper) ImplementsTokenField_ReadFromPasteboard() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:readFromPasteboard:"))
}

func (t_ TokenFieldDelegateWrapper) TokenField_ReadFromPasteboard(tokenField ITokenField, pboard IPasteboard) []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](t_, "tokenField:readFromPasteboard:", tokenField, pboard)
	return rv
}

func (t_ *TokenFieldDelegateWrapper) ImplementsTokenField_WriteRepresentedObjects_ToPasteboard() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:writeRepresentedObjects:toPasteboard:"))
}

func (t_ TokenFieldDelegateWrapper) TokenField_WriteRepresentedObjects_ToPasteboard(tokenField ITokenField, objects []objc.IObject, pboard IPasteboard) bool {
	rv := ffi.CallMethod[bool](t_, "tokenField:writeRepresentedObjects:toPasteboard:", tokenField, objects, pboard)
	return rv
}

func (t_ *TokenFieldDelegateWrapper) ImplementsTokenField_HasMenuForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:hasMenuForRepresentedObject:"))
}

func (t_ TokenFieldDelegateWrapper) TokenField_HasMenuForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) bool {
	rv := ffi.CallMethod[bool](t_, "tokenField:hasMenuForRepresentedObject:", tokenField, representedObject)
	return rv
}

func (t_ *TokenFieldDelegateWrapper) ImplementsTokenField_MenuForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:menuForRepresentedObject:"))
}

func (t_ TokenFieldDelegateWrapper) TokenField_MenuForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) Menu {
	rv := ffi.CallMethod[Menu](t_, "tokenField:menuForRepresentedObject:", tokenField, representedObject)
	return rv
}

var TokenFieldCellClass = _TokenFieldCellClass{objc.GetClass("NSTokenFieldCell")}

type _TokenFieldCellClass struct {
	objc.Class
}

type ITokenFieldCell interface {
	ITextFieldCell
	TokenStyle() TokenStyle
	SetTokenStyle(value TokenStyle)
	TokenizingCharacterSet() foundation.CharacterSet
	SetTokenizingCharacterSet(value foundation.ICharacterSet)
	CompletionDelay() foundation.TimeInterval
	SetCompletionDelay(value foundation.TimeInterval)
	Delegate() TokenFieldCellDelegateWrapper
	SetDelegate(value TokenFieldCellDelegate)
}

type TokenFieldCell struct {
	TextFieldCell
}

func MakeTokenFieldCell(ptr unsafe.Pointer) TokenFieldCell {
	return TokenFieldCell{
		TextFieldCell: MakeTextFieldCell(ptr),
	}
}

func (t_ TokenFieldCell) InitTextCell(_string string) TokenFieldCell {
	rv := ffi.CallMethod[TokenFieldCell](t_, "initTextCell:", _string)
	rv.Autorelease()
	return rv
}

func (t_ TokenFieldCell) InitImageCell(image IImage) TokenFieldCell {
	rv := ffi.CallMethod[TokenFieldCell](t_, "initImageCell:", image)
	rv.Autorelease()
	return rv
}

func (t_ TokenFieldCell) Init() TokenFieldCell {
	rv := ffi.CallMethod[TokenFieldCell](t_, "init")
	rv.Autorelease()
	return rv
}

func (tc _TokenFieldCellClass) Alloc() TokenFieldCell {
	rv := ffi.CallMethod[TokenFieldCell](tc, "alloc")
	return rv
}

func (tc _TokenFieldCellClass) New() TokenFieldCell {
	rv := ffi.CallMethod[TokenFieldCell](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTokenFieldCell() TokenFieldCell {
	return TokenFieldCellClass.New()
}

func (t_ TokenFieldCell) TokenStyle() TokenStyle {
	rv := ffi.CallMethod[TokenStyle](t_, "tokenStyle")
	return rv
}

func (t_ TokenFieldCell) SetTokenStyle(value TokenStyle) {
	ffi.CallMethod[ffi.Void](t_, "setTokenStyle:", value)
}

func (tc _TokenFieldCellClass) DefaultTokenizingCharacterSet() foundation.CharacterSet {
	rv := ffi.CallMethod[foundation.CharacterSet](tc, "defaultTokenizingCharacterSet")
	return rv
}

func (t_ TokenFieldCell) TokenizingCharacterSet() foundation.CharacterSet {
	rv := ffi.CallMethod[foundation.CharacterSet](t_, "tokenizingCharacterSet")
	return rv
}

func (t_ TokenFieldCell) SetTokenizingCharacterSet(value foundation.ICharacterSet) {
	ffi.CallMethod[ffi.Void](t_, "setTokenizingCharacterSet:", value)
}

func (t_ TokenFieldCell) CompletionDelay() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](t_, "completionDelay")
	return rv
}

func (t_ TokenFieldCell) SetCompletionDelay(value foundation.TimeInterval) {
	ffi.CallMethod[ffi.Void](t_, "setCompletionDelay:", value)
}

func (tc _TokenFieldCellClass) DefaultCompletionDelay() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](tc, "defaultCompletionDelay")
	return rv
}

func (t_ TokenFieldCell) Delegate() TokenFieldCellDelegateWrapper {
	rv := ffi.CallMethod[TokenFieldCellDelegateWrapper](t_, "delegate")
	return rv
}

func (t_ TokenFieldCell) SetDelegate(value TokenFieldCellDelegate) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](t_, "setDelegate:", po)
}

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

type TokenFieldCellDelegateImpl struct {
	_TokenFieldCell_DisplayStringForRepresentedObject                        func(tokenFieldCell TokenFieldCell, representedObject objc.Object) string
	_TokenFieldCell_StyleForRepresentedObject                                func(tokenFieldCell TokenFieldCell, representedObject objc.Object) TokenStyle
	_TokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem func(tokenFieldCell TokenFieldCell, substring string, tokenIndex int, selectedIndex *int) []objc.IObject
	_TokenFieldCell_EditingStringForRepresentedObject                        func(tokenFieldCell TokenFieldCell, representedObject objc.Object) string
	_TokenFieldCell_RepresentedObjectForEditingString                        func(tokenFieldCell TokenFieldCell, editingString string) objc.IObject
	_TokenFieldCell_ShouldAddObjects_AtIndex                                 func(tokenFieldCell TokenFieldCell, tokens []objc.Object, index uint) []objc.IObject
	_TokenFieldCell_ReadFromPasteboard                                       func(tokenFieldCell TokenFieldCell, pboard Pasteboard) []objc.IObject
	_TokenFieldCell_WriteRepresentedObjects_ToPasteboard                     func(tokenFieldCell TokenFieldCell, objects []objc.Object, pboard Pasteboard) bool
	_TokenFieldCell_HasMenuForRepresentedObject                              func(tokenFieldCell TokenFieldCell, representedObject objc.Object) bool
	_TokenFieldCell_MenuForRepresentedObject                                 func(tokenFieldCell TokenFieldCell, representedObject objc.Object) IMenu
}

func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_DisplayStringForRepresentedObject() bool {
	return di._TokenFieldCell_DisplayStringForRepresentedObject != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_DisplayStringForRepresentedObject(f func(tokenFieldCell TokenFieldCell, representedObject objc.Object) string) {
	di._TokenFieldCell_DisplayStringForRepresentedObject = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_DisplayStringForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) string {
	return di._TokenFieldCell_DisplayStringForRepresentedObject(tokenFieldCell, representedObject)
}
func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_StyleForRepresentedObject() bool {
	return di._TokenFieldCell_StyleForRepresentedObject != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_StyleForRepresentedObject(f func(tokenFieldCell TokenFieldCell, representedObject objc.Object) TokenStyle) {
	di._TokenFieldCell_StyleForRepresentedObject = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_StyleForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) TokenStyle {
	return di._TokenFieldCell_StyleForRepresentedObject(tokenFieldCell, representedObject)
}
func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem() bool {
	return di._TokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(f func(tokenFieldCell TokenFieldCell, substring string, tokenIndex int, selectedIndex *int) []objc.IObject) {
	di._TokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(tokenFieldCell TokenFieldCell, substring string, tokenIndex int, selectedIndex *int) []objc.IObject {
	return di._TokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(tokenFieldCell, substring, tokenIndex, selectedIndex)
}
func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_EditingStringForRepresentedObject() bool {
	return di._TokenFieldCell_EditingStringForRepresentedObject != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_EditingStringForRepresentedObject(f func(tokenFieldCell TokenFieldCell, representedObject objc.Object) string) {
	di._TokenFieldCell_EditingStringForRepresentedObject = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_EditingStringForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) string {
	return di._TokenFieldCell_EditingStringForRepresentedObject(tokenFieldCell, representedObject)
}
func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_RepresentedObjectForEditingString() bool {
	return di._TokenFieldCell_RepresentedObjectForEditingString != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_RepresentedObjectForEditingString(f func(tokenFieldCell TokenFieldCell, editingString string) objc.IObject) {
	di._TokenFieldCell_RepresentedObjectForEditingString = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_RepresentedObjectForEditingString(tokenFieldCell TokenFieldCell, editingString string) objc.IObject {
	return di._TokenFieldCell_RepresentedObjectForEditingString(tokenFieldCell, editingString)
}
func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_ShouldAddObjects_AtIndex() bool {
	return di._TokenFieldCell_ShouldAddObjects_AtIndex != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_ShouldAddObjects_AtIndex(f func(tokenFieldCell TokenFieldCell, tokens []objc.Object, index uint) []objc.IObject) {
	di._TokenFieldCell_ShouldAddObjects_AtIndex = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_ShouldAddObjects_AtIndex(tokenFieldCell TokenFieldCell, tokens []objc.Object, index uint) []objc.IObject {
	return di._TokenFieldCell_ShouldAddObjects_AtIndex(tokenFieldCell, tokens, index)
}
func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_ReadFromPasteboard() bool {
	return di._TokenFieldCell_ReadFromPasteboard != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_ReadFromPasteboard(f func(tokenFieldCell TokenFieldCell, pboard Pasteboard) []objc.IObject) {
	di._TokenFieldCell_ReadFromPasteboard = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_ReadFromPasteboard(tokenFieldCell TokenFieldCell, pboard Pasteboard) []objc.IObject {
	return di._TokenFieldCell_ReadFromPasteboard(tokenFieldCell, pboard)
}
func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_WriteRepresentedObjects_ToPasteboard() bool {
	return di._TokenFieldCell_WriteRepresentedObjects_ToPasteboard != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_WriteRepresentedObjects_ToPasteboard(f func(tokenFieldCell TokenFieldCell, objects []objc.Object, pboard Pasteboard) bool) {
	di._TokenFieldCell_WriteRepresentedObjects_ToPasteboard = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_WriteRepresentedObjects_ToPasteboard(tokenFieldCell TokenFieldCell, objects []objc.Object, pboard Pasteboard) bool {
	return di._TokenFieldCell_WriteRepresentedObjects_ToPasteboard(tokenFieldCell, objects, pboard)
}
func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_HasMenuForRepresentedObject() bool {
	return di._TokenFieldCell_HasMenuForRepresentedObject != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_HasMenuForRepresentedObject(f func(tokenFieldCell TokenFieldCell, representedObject objc.Object) bool) {
	di._TokenFieldCell_HasMenuForRepresentedObject = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_HasMenuForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) bool {
	return di._TokenFieldCell_HasMenuForRepresentedObject(tokenFieldCell, representedObject)
}
func (di *TokenFieldCellDelegateImpl) ImplementsTokenFieldCell_MenuForRepresentedObject() bool {
	return di._TokenFieldCell_MenuForRepresentedObject != nil
}

func (di *TokenFieldCellDelegateImpl) SetTokenFieldCell_MenuForRepresentedObject(f func(tokenFieldCell TokenFieldCell, representedObject objc.Object) IMenu) {
	di._TokenFieldCell_MenuForRepresentedObject = f
}

func (di *TokenFieldCellDelegateImpl) TokenFieldCell_MenuForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) IMenu {
	return di._TokenFieldCell_MenuForRepresentedObject(tokenFieldCell, representedObject)
}

type TokenFieldCellDelegateWrapper struct {
	objc.Object
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_DisplayStringForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:displayStringForRepresentedObject:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_DisplayStringForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) string {
	rv := ffi.CallMethod[string](t_, "tokenFieldCell:displayStringForRepresentedObject:", tokenFieldCell, representedObject)
	return rv
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_StyleForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:styleForRepresentedObject:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_StyleForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) TokenStyle {
	rv := ffi.CallMethod[TokenStyle](t_, "tokenFieldCell:styleForRepresentedObject:", tokenFieldCell, representedObject)
	return rv
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:completionsForSubstring:indexOfToken:indexOfSelectedItem:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_CompletionsForSubstring_IndexOfToken_IndexOfSelectedItem(tokenFieldCell ITokenFieldCell, substring string, tokenIndex int, selectedIndex *int) []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](t_, "tokenFieldCell:completionsForSubstring:indexOfToken:indexOfSelectedItem:", tokenFieldCell, substring, tokenIndex, selectedIndex)
	return rv
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_EditingStringForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:editingStringForRepresentedObject:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_EditingStringForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) string {
	rv := ffi.CallMethod[string](t_, "tokenFieldCell:editingStringForRepresentedObject:", tokenFieldCell, representedObject)
	return rv
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_RepresentedObjectForEditingString() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:representedObjectForEditingString:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_RepresentedObjectForEditingString(tokenFieldCell ITokenFieldCell, editingString string) objc.Object {
	rv := ffi.CallMethod[objc.Object](t_, "tokenFieldCell:representedObjectForEditingString:", tokenFieldCell, editingString)
	return rv
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_ShouldAddObjects_AtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:shouldAddObjects:atIndex:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_ShouldAddObjects_AtIndex(tokenFieldCell ITokenFieldCell, tokens []objc.IObject, index uint) []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](t_, "tokenFieldCell:shouldAddObjects:atIndex:", tokenFieldCell, tokens, index)
	return rv
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_ReadFromPasteboard() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:readFromPasteboard:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_ReadFromPasteboard(tokenFieldCell ITokenFieldCell, pboard IPasteboard) []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](t_, "tokenFieldCell:readFromPasteboard:", tokenFieldCell, pboard)
	return rv
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_WriteRepresentedObjects_ToPasteboard() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:writeRepresentedObjects:toPasteboard:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_WriteRepresentedObjects_ToPasteboard(tokenFieldCell ITokenFieldCell, objects []objc.IObject, pboard IPasteboard) bool {
	rv := ffi.CallMethod[bool](t_, "tokenFieldCell:writeRepresentedObjects:toPasteboard:", tokenFieldCell, objects, pboard)
	return rv
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_HasMenuForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:hasMenuForRepresentedObject:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_HasMenuForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) bool {
	rv := ffi.CallMethod[bool](t_, "tokenFieldCell:hasMenuForRepresentedObject:", tokenFieldCell, representedObject)
	return rv
}

func (t_ *TokenFieldCellDelegateWrapper) ImplementsTokenFieldCell_MenuForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:menuForRepresentedObject:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCell_MenuForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) Menu {
	rv := ffi.CallMethod[Menu](t_, "tokenFieldCell:menuForRepresentedObject:", tokenFieldCell, representedObject)
	return rv
}
