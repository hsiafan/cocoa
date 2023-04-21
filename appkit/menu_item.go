// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[MenuItem](m_, "initWithTitle:action:keyEquivalent:", string_, selector, charCode)
	return rv
}

func (mc _MenuItemClass) Alloc() MenuItem {
	rv := objc.CallMethod[MenuItem](mc, "alloc")
	return rv
}

func (mc _MenuItemClass) New() MenuItem {
	rv := objc.CallMethod[MenuItem](mc, "new")
	rv.Autorelease()
	return rv
}

func NewMenuItem() MenuItem {
	return MenuItemClass.New()
}

func (m_ MenuItem) Init() MenuItem {
	rv := objc.CallMethod[MenuItem](m_, "init")
	return rv
}

func (mc _MenuItemClass) SeparatorItem() MenuItem {
	rv := objc.CallMethod[MenuItem](mc, "separatorItem")
	return rv
}

// deprecated
func (m_ MenuItem) SetMnemonicLocation(location uint) {
	objc.CallMethod[objc.Void](m_, "setMnemonicLocation:", location)
}

// deprecated
func (m_ MenuItem) MnemonicLocation() uint {
	rv := objc.CallMethod[uint](m_, "mnemonicLocation")
	return rv
}

// deprecated
func (m_ MenuItem) SetTitleWithMnemonic(stringWithAmpersand string) {
	objc.CallMethod[objc.Void](m_, "setTitleWithMnemonic:", stringWithAmpersand)
}

// deprecated
func (m_ MenuItem) Mnemonic() string {
	rv := objc.CallMethod[string](m_, "mnemonic")
	return rv
}

func (m_ MenuItem) IsEnabled() bool {
	rv := objc.CallMethod[bool](m_, "isEnabled")
	return rv
}

func (m_ MenuItem) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](m_, "setEnabled:", value)
}

func (m_ MenuItem) IsHidden() bool {
	rv := objc.CallMethod[bool](m_, "isHidden")
	return rv
}

func (m_ MenuItem) SetHidden(value bool) {
	objc.CallMethod[objc.Void](m_, "setHidden:", value)
}

func (m_ MenuItem) IsHiddenOrHasHiddenAncestor() bool {
	rv := objc.CallMethod[bool](m_, "isHiddenOrHasHiddenAncestor")
	return rv
}

func (m_ MenuItem) Target() objc.Object {
	rv := objc.CallMethod[objc.Object](m_, "target")
	return rv
}

func (m_ MenuItem) SetTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](m_, "setTarget:", value)
}

func (m_ MenuItem) Action() objc.Selector {
	rv := objc.CallMethod[objc.Selector](m_, "action")
	return rv
}

func (m_ MenuItem) SetAction(value objc.Selector) {
	objc.CallMethod[objc.Void](m_, "setAction:", value)
}

func (m_ MenuItem) Title() string {
	rv := objc.CallMethod[string](m_, "title")
	return rv
}

func (m_ MenuItem) SetTitle(value string) {
	objc.CallMethod[objc.Void](m_, "setTitle:", value)
}

func (m_ MenuItem) AttributedTitle() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](m_, "attributedTitle")
	return rv
}

func (m_ MenuItem) SetAttributedTitle(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](m_, "setAttributedTitle:", value)
}

func (m_ MenuItem) Tag() int {
	rv := objc.CallMethod[int](m_, "tag")
	return rv
}

func (m_ MenuItem) SetTag(value int) {
	objc.CallMethod[objc.Void](m_, "setTag:", value)
}

func (m_ MenuItem) State() ControlStateValue {
	rv := objc.CallMethod[ControlStateValue](m_, "state")
	return rv
}

func (m_ MenuItem) SetState(value ControlStateValue) {
	objc.CallMethod[objc.Void](m_, "setState:", value)
}

func (m_ MenuItem) Image() Image {
	rv := objc.CallMethod[Image](m_, "image")
	return rv
}

func (m_ MenuItem) SetImage(value IImage) {
	objc.CallMethod[objc.Void](m_, "setImage:", value)
}

func (m_ MenuItem) OnStateImage() Image {
	rv := objc.CallMethod[Image](m_, "onStateImage")
	return rv
}

func (m_ MenuItem) SetOnStateImage(value IImage) {
	objc.CallMethod[objc.Void](m_, "setOnStateImage:", value)
}

func (m_ MenuItem) OffStateImage() Image {
	rv := objc.CallMethod[Image](m_, "offStateImage")
	return rv
}

func (m_ MenuItem) SetOffStateImage(value IImage) {
	objc.CallMethod[objc.Void](m_, "setOffStateImage:", value)
}

func (m_ MenuItem) MixedStateImage() Image {
	rv := objc.CallMethod[Image](m_, "mixedStateImage")
	return rv
}

func (m_ MenuItem) SetMixedStateImage(value IImage) {
	objc.CallMethod[objc.Void](m_, "setMixedStateImage:", value)
}

func (m_ MenuItem) Submenu() Menu {
	rv := objc.CallMethod[Menu](m_, "submenu")
	return rv
}

func (m_ MenuItem) SetSubmenu(value IMenu) {
	objc.CallMethod[objc.Void](m_, "setSubmenu:", value)
}

func (m_ MenuItem) HasSubmenu() bool {
	rv := objc.CallMethod[bool](m_, "hasSubmenu")
	return rv
}

func (m_ MenuItem) ParentItem() MenuItem {
	rv := objc.CallMethod[MenuItem](m_, "parentItem")
	return rv
}

func (m_ MenuItem) IsSeparatorItem() bool {
	rv := objc.CallMethod[bool](m_, "isSeparatorItem")
	return rv
}

func (m_ MenuItem) Menu() Menu {
	rv := objc.CallMethod[Menu](m_, "menu")
	return rv
}

func (m_ MenuItem) SetMenu(value IMenu) {
	objc.CallMethod[objc.Void](m_, "setMenu:", value)
}

func (m_ MenuItem) KeyEquivalent() string {
	rv := objc.CallMethod[string](m_, "keyEquivalent")
	return rv
}

func (m_ MenuItem) SetKeyEquivalent(value string) {
	objc.CallMethod[objc.Void](m_, "setKeyEquivalent:", value)
}

func (m_ MenuItem) KeyEquivalentModifierMask() EventModifierFlags {
	rv := objc.CallMethod[EventModifierFlags](m_, "keyEquivalentModifierMask")
	return rv
}

func (m_ MenuItem) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	objc.CallMethod[objc.Void](m_, "setKeyEquivalentModifierMask:", value)
}

func (mc _MenuItemClass) UsesUserKeyEquivalents() bool {
	rv := objc.CallMethod[bool](mc, "usesUserKeyEquivalents")
	return rv
}

func (mc _MenuItemClass) SetUsesUserKeyEquivalents(value bool) {
	objc.CallMethod[objc.Void](mc, "setUsesUserKeyEquivalents:", value)
}

func (m_ MenuItem) UserKeyEquivalent() string {
	rv := objc.CallMethod[string](m_, "userKeyEquivalent")
	return rv
}

func (m_ MenuItem) AllowsAutomaticKeyEquivalentLocalization() bool {
	rv := objc.CallMethod[bool](m_, "allowsAutomaticKeyEquivalentLocalization")
	return rv
}

func (m_ MenuItem) SetAllowsAutomaticKeyEquivalentLocalization(value bool) {
	objc.CallMethod[objc.Void](m_, "setAllowsAutomaticKeyEquivalentLocalization:", value)
}

func (m_ MenuItem) AllowsAutomaticKeyEquivalentMirroring() bool {
	rv := objc.CallMethod[bool](m_, "allowsAutomaticKeyEquivalentMirroring")
	return rv
}

func (m_ MenuItem) SetAllowsAutomaticKeyEquivalentMirroring(value bool) {
	objc.CallMethod[objc.Void](m_, "setAllowsAutomaticKeyEquivalentMirroring:", value)
}

func (m_ MenuItem) AllowsKeyEquivalentWhenHidden() bool {
	rv := objc.CallMethod[bool](m_, "allowsKeyEquivalentWhenHidden")
	return rv
}

func (m_ MenuItem) SetAllowsKeyEquivalentWhenHidden(value bool) {
	objc.CallMethod[objc.Void](m_, "setAllowsKeyEquivalentWhenHidden:", value)
}

func (m_ MenuItem) IsAlternate() bool {
	rv := objc.CallMethod[bool](m_, "isAlternate")
	return rv
}

func (m_ MenuItem) SetAlternate(value bool) {
	objc.CallMethod[objc.Void](m_, "setAlternate:", value)
}

func (m_ MenuItem) IndentationLevel() int {
	rv := objc.CallMethod[int](m_, "indentationLevel")
	return rv
}

func (m_ MenuItem) SetIndentationLevel(value int) {
	objc.CallMethod[objc.Void](m_, "setIndentationLevel:", value)
}

func (m_ MenuItem) ToolTip() string {
	rv := objc.CallMethod[string](m_, "toolTip")
	return rv
}

func (m_ MenuItem) SetToolTip(value string) {
	objc.CallMethod[objc.Void](m_, "setToolTip:", value)
}

func (m_ MenuItem) RepresentedObject() objc.Object {
	rv := objc.CallMethod[objc.Object](m_, "representedObject")
	return rv
}

func (m_ MenuItem) SetRepresentedObject(value objc.IObject) {
	objc.CallMethod[objc.Void](m_, "setRepresentedObject:", value)
}

func (m_ MenuItem) View() View {
	rv := objc.CallMethod[View](m_, "view")
	return rv
}

func (m_ MenuItem) SetView(value IView) {
	objc.CallMethod[objc.Void](m_, "setView:", value)
}

func (m_ MenuItem) IsHighlighted() bool {
	rv := objc.CallMethod[bool](m_, "isHighlighted")
	return rv
}
