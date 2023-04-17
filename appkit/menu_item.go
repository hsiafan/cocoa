// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var MenuItemClass = _MenuItemClass{objc.GetClass("NSMenuItem")}

type _MenuItemClass struct {
	objc.Class
}

type IMenuItem interface {
	objc.IObject
	// deprecated
	SetMnemonicLocation(location uint)
	// deprecated
	MnemonicLocation() uint
	// deprecated
	SetTitleWithMnemonic(stringWithAmpersand string)
	// deprecated
	Mnemonic() string
	IsEnabled() bool
	SetEnabled(value bool)
	IsHidden() bool
	SetHidden(value bool)
	IsHiddenOrHasHiddenAncestor() bool
	Target() objc.Object
	SetTarget(value objc.IObject)
	Action() objc.Selector
	SetAction(value objc.Selector)
	Title() string
	SetTitle(value string)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	Tag() int
	SetTag(value int)
	State() ControlStateValue
	SetState(value ControlStateValue)
	Image() Image
	SetImage(value IImage)
	OnStateImage() Image
	SetOnStateImage(value IImage)
	OffStateImage() Image
	SetOffStateImage(value IImage)
	MixedStateImage() Image
	SetMixedStateImage(value IImage)
	Submenu() Menu
	SetSubmenu(value IMenu)
	HasSubmenu() bool
	ParentItem() MenuItem
	IsSeparatorItem() bool
	Menu() Menu
	SetMenu(value IMenu)
	KeyEquivalent() string
	SetKeyEquivalent(value string)
	KeyEquivalentModifierMask() EventModifierFlags
	SetKeyEquivalentModifierMask(value EventModifierFlags)
	UserKeyEquivalent() string
	AllowsAutomaticKeyEquivalentLocalization() bool
	SetAllowsAutomaticKeyEquivalentLocalization(value bool)
	AllowsAutomaticKeyEquivalentMirroring() bool
	SetAllowsAutomaticKeyEquivalentMirroring(value bool)
	AllowsKeyEquivalentWhenHidden() bool
	SetAllowsKeyEquivalentWhenHidden(value bool)
	IsAlternate() bool
	SetAlternate(value bool)
	IndentationLevel() int
	SetIndentationLevel(value int)
	ToolTip() string
	SetToolTip(value string)
	RepresentedObject() objc.Object
	SetRepresentedObject(value objc.IObject)
	View() View
	SetView(value IView)
	IsHighlighted() bool
}

type MenuItem struct {
	objc.Object
}

func MakeMenuItem(ptr unsafe.Pointer) MenuItem {
	return MenuItem{
		Object: objc.MakeObject(ptr),
	}
}

func (m_ MenuItem) InitWithTitle_Action_KeyEquivalent(string_ string, selector objc.Selector, charCode string) MenuItem {
	rv := ffi.CallMethod[MenuItem](m_, "initWithTitle:action:keyEquivalent:", string_, selector, charCode)
	return rv
}

func (mc _MenuItemClass) Alloc() MenuItem {
	rv := ffi.CallMethod[MenuItem](mc, "alloc")
	return rv
}

func (mc _MenuItemClass) New() MenuItem {
	rv := ffi.CallMethod[MenuItem](mc, "new")
	rv.Autorelease()
	return rv
}

func NewMenuItem() MenuItem {
	return MenuItemClass.New()
}

func (m_ MenuItem) Init() MenuItem {
	rv := ffi.CallMethod[MenuItem](m_, "init")
	return rv
}

func (mc _MenuItemClass) SeparatorItem() MenuItem {
	rv := ffi.CallMethod[MenuItem](mc, "separatorItem")
	return rv
}

// deprecated
func (m_ MenuItem) SetMnemonicLocation(location uint) {
	ffi.CallMethod[ffi.Void](m_, "setMnemonicLocation:", location)
}

// deprecated
func (m_ MenuItem) MnemonicLocation() uint {
	rv := ffi.CallMethod[uint](m_, "mnemonicLocation")
	return rv
}

// deprecated
func (m_ MenuItem) SetTitleWithMnemonic(stringWithAmpersand string) {
	ffi.CallMethod[ffi.Void](m_, "setTitleWithMnemonic:", stringWithAmpersand)
}

// deprecated
func (m_ MenuItem) Mnemonic() string {
	rv := ffi.CallMethod[string](m_, "mnemonic")
	return rv
}

func (m_ MenuItem) IsEnabled() bool {
	rv := ffi.CallMethod[bool](m_, "isEnabled")
	return rv
}

func (m_ MenuItem) SetEnabled(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setEnabled:", value)
}

func (m_ MenuItem) IsHidden() bool {
	rv := ffi.CallMethod[bool](m_, "isHidden")
	return rv
}

func (m_ MenuItem) SetHidden(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setHidden:", value)
}

func (m_ MenuItem) IsHiddenOrHasHiddenAncestor() bool {
	rv := ffi.CallMethod[bool](m_, "isHiddenOrHasHiddenAncestor")
	return rv
}

func (m_ MenuItem) Target() objc.Object {
	rv := ffi.CallMethod[objc.Object](m_, "target")
	return rv
}

func (m_ MenuItem) SetTarget(value objc.IObject) {
	ffi.CallMethod[ffi.Void](m_, "setTarget:", value)
}

func (m_ MenuItem) Action() objc.Selector {
	rv := ffi.CallMethod[objc.Selector](m_, "action")
	return rv
}

func (m_ MenuItem) SetAction(value objc.Selector) {
	ffi.CallMethod[ffi.Void](m_, "setAction:", value)
}

func (m_ MenuItem) Title() string {
	rv := ffi.CallMethod[string](m_, "title")
	return rv
}

func (m_ MenuItem) SetTitle(value string) {
	ffi.CallMethod[ffi.Void](m_, "setTitle:", value)
}

func (m_ MenuItem) AttributedTitle() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](m_, "attributedTitle")
	return rv
}

func (m_ MenuItem) SetAttributedTitle(value foundation.IAttributedString) {
	ffi.CallMethod[ffi.Void](m_, "setAttributedTitle:", value)
}

func (m_ MenuItem) Tag() int {
	rv := ffi.CallMethod[int](m_, "tag")
	return rv
}

func (m_ MenuItem) SetTag(value int) {
	ffi.CallMethod[ffi.Void](m_, "setTag:", value)
}

func (m_ MenuItem) State() ControlStateValue {
	rv := ffi.CallMethod[ControlStateValue](m_, "state")
	return rv
}

func (m_ MenuItem) SetState(value ControlStateValue) {
	ffi.CallMethod[ffi.Void](m_, "setState:", value)
}

func (m_ MenuItem) Image() Image {
	rv := ffi.CallMethod[Image](m_, "image")
	return rv
}

func (m_ MenuItem) SetImage(value IImage) {
	ffi.CallMethod[ffi.Void](m_, "setImage:", value)
}

func (m_ MenuItem) OnStateImage() Image {
	rv := ffi.CallMethod[Image](m_, "onStateImage")
	return rv
}

func (m_ MenuItem) SetOnStateImage(value IImage) {
	ffi.CallMethod[ffi.Void](m_, "setOnStateImage:", value)
}

func (m_ MenuItem) OffStateImage() Image {
	rv := ffi.CallMethod[Image](m_, "offStateImage")
	return rv
}

func (m_ MenuItem) SetOffStateImage(value IImage) {
	ffi.CallMethod[ffi.Void](m_, "setOffStateImage:", value)
}

func (m_ MenuItem) MixedStateImage() Image {
	rv := ffi.CallMethod[Image](m_, "mixedStateImage")
	return rv
}

func (m_ MenuItem) SetMixedStateImage(value IImage) {
	ffi.CallMethod[ffi.Void](m_, "setMixedStateImage:", value)
}

func (m_ MenuItem) Submenu() Menu {
	rv := ffi.CallMethod[Menu](m_, "submenu")
	return rv
}

func (m_ MenuItem) SetSubmenu(value IMenu) {
	ffi.CallMethod[ffi.Void](m_, "setSubmenu:", value)
}

func (m_ MenuItem) HasSubmenu() bool {
	rv := ffi.CallMethod[bool](m_, "hasSubmenu")
	return rv
}

func (m_ MenuItem) ParentItem() MenuItem {
	rv := ffi.CallMethod[MenuItem](m_, "parentItem")
	return rv
}

func (m_ MenuItem) IsSeparatorItem() bool {
	rv := ffi.CallMethod[bool](m_, "isSeparatorItem")
	return rv
}

func (m_ MenuItem) Menu() Menu {
	rv := ffi.CallMethod[Menu](m_, "menu")
	return rv
}

func (m_ MenuItem) SetMenu(value IMenu) {
	ffi.CallMethod[ffi.Void](m_, "setMenu:", value)
}

func (m_ MenuItem) KeyEquivalent() string {
	rv := ffi.CallMethod[string](m_, "keyEquivalent")
	return rv
}

func (m_ MenuItem) SetKeyEquivalent(value string) {
	ffi.CallMethod[ffi.Void](m_, "setKeyEquivalent:", value)
}

func (m_ MenuItem) KeyEquivalentModifierMask() EventModifierFlags {
	rv := ffi.CallMethod[EventModifierFlags](m_, "keyEquivalentModifierMask")
	return rv
}

func (m_ MenuItem) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	ffi.CallMethod[ffi.Void](m_, "setKeyEquivalentModifierMask:", value)
}

func (mc _MenuItemClass) UsesUserKeyEquivalents() bool {
	rv := ffi.CallMethod[bool](mc, "usesUserKeyEquivalents")
	return rv
}

func (mc _MenuItemClass) SetUsesUserKeyEquivalents(value bool) {
	ffi.CallMethod[ffi.Void](mc, "setUsesUserKeyEquivalents:", value)
}

func (m_ MenuItem) UserKeyEquivalent() string {
	rv := ffi.CallMethod[string](m_, "userKeyEquivalent")
	return rv
}

func (m_ MenuItem) AllowsAutomaticKeyEquivalentLocalization() bool {
	rv := ffi.CallMethod[bool](m_, "allowsAutomaticKeyEquivalentLocalization")
	return rv
}

func (m_ MenuItem) SetAllowsAutomaticKeyEquivalentLocalization(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setAllowsAutomaticKeyEquivalentLocalization:", value)
}

func (m_ MenuItem) AllowsAutomaticKeyEquivalentMirroring() bool {
	rv := ffi.CallMethod[bool](m_, "allowsAutomaticKeyEquivalentMirroring")
	return rv
}

func (m_ MenuItem) SetAllowsAutomaticKeyEquivalentMirroring(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setAllowsAutomaticKeyEquivalentMirroring:", value)
}

func (m_ MenuItem) AllowsKeyEquivalentWhenHidden() bool {
	rv := ffi.CallMethod[bool](m_, "allowsKeyEquivalentWhenHidden")
	return rv
}

func (m_ MenuItem) SetAllowsKeyEquivalentWhenHidden(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setAllowsKeyEquivalentWhenHidden:", value)
}

func (m_ MenuItem) IsAlternate() bool {
	rv := ffi.CallMethod[bool](m_, "isAlternate")
	return rv
}

func (m_ MenuItem) SetAlternate(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setAlternate:", value)
}

func (m_ MenuItem) IndentationLevel() int {
	rv := ffi.CallMethod[int](m_, "indentationLevel")
	return rv
}

func (m_ MenuItem) SetIndentationLevel(value int) {
	ffi.CallMethod[ffi.Void](m_, "setIndentationLevel:", value)
}

func (m_ MenuItem) ToolTip() string {
	rv := ffi.CallMethod[string](m_, "toolTip")
	return rv
}

func (m_ MenuItem) SetToolTip(value string) {
	ffi.CallMethod[ffi.Void](m_, "setToolTip:", value)
}

func (m_ MenuItem) RepresentedObject() objc.Object {
	rv := ffi.CallMethod[objc.Object](m_, "representedObject")
	return rv
}

func (m_ MenuItem) SetRepresentedObject(value objc.IObject) {
	ffi.CallMethod[ffi.Void](m_, "setRepresentedObject:", value)
}

func (m_ MenuItem) View() View {
	rv := ffi.CallMethod[View](m_, "view")
	return rv
}

func (m_ MenuItem) SetView(value IView) {
	ffi.CallMethod[ffi.Void](m_, "setView:", value)
}

func (m_ MenuItem) IsHighlighted() bool {
	rv := ffi.CallMethod[bool](m_, "isHighlighted")
	return rv
}
