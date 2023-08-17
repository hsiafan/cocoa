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
	rv := objc.CallMethod[MenuItem](m_, objc.GetSelector("initWithTitle:action:keyEquivalent:"), string_, selector, charCode)
	return rv
}

func (mc _MenuItemClass) Alloc() MenuItem {
	rv := objc.CallMethod[MenuItem](mc, objc.GetSelector("alloc"))
	return rv
}

func (mc _MenuItemClass) New() MenuItem {
	rv := objc.CallMethod[MenuItem](mc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewMenuItem() MenuItem {
	return MenuItemClass.New()
}

func (m_ MenuItem) Init() MenuItem {
	rv := objc.CallMethod[MenuItem](m_, objc.GetSelector("init"))
	return rv
}

func (mc _MenuItemClass) SeparatorItem() MenuItem {
	rv := objc.CallMethod[MenuItem](mc, objc.GetSelector("separatorItem"))
	return rv
}

// deprecated
func (m_ MenuItem) SetMnemonicLocation(location uint) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setMnemonicLocation:"), location)
}

// deprecated
func (m_ MenuItem) MnemonicLocation() uint {
	rv := objc.CallMethod[uint](m_, objc.GetSelector("mnemonicLocation"))
	return rv
}

// deprecated
func (m_ MenuItem) SetTitleWithMnemonic(stringWithAmpersand string) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setTitleWithMnemonic:"), stringWithAmpersand)
}

// deprecated
func (m_ MenuItem) Mnemonic() string {
	rv := objc.CallMethod[string](m_, objc.GetSelector("mnemonic"))
	return rv
}

func (m_ MenuItem) IsEnabled() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("isEnabled"))
	return rv
}

func (m_ MenuItem) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setEnabled:"), value)
}

func (m_ MenuItem) IsHidden() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("isHidden"))
	return rv
}

func (m_ MenuItem) SetHidden(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setHidden:"), value)
}

func (m_ MenuItem) IsHiddenOrHasHiddenAncestor() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("isHiddenOrHasHiddenAncestor"))
	return rv
}

// weak property
func (m_ MenuItem) Target() objc.Object {
	rv := objc.CallMethod[objc.Object](m_, objc.GetSelector("target"))
	return rv
}

// weak property
func (m_ MenuItem) SetTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setTarget:"), objc.ExtractPtr(value))
}

func (m_ MenuItem) Action() objc.Selector {
	rv := objc.CallMethod[objc.Selector](m_, objc.GetSelector("action"))
	return rv
}

func (m_ MenuItem) SetAction(value objc.Selector) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAction:"), value)
}

func (m_ MenuItem) Title() string {
	rv := objc.CallMethod[string](m_, objc.GetSelector("title"))
	return rv
}

func (m_ MenuItem) SetTitle(value string) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setTitle:"), value)
}

func (m_ MenuItem) AttributedTitle() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](m_, objc.GetSelector("attributedTitle"))
	return rv
}

func (m_ MenuItem) SetAttributedTitle(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAttributedTitle:"), objc.ExtractPtr(value))
}

func (m_ MenuItem) Tag() int {
	rv := objc.CallMethod[int](m_, objc.GetSelector("tag"))
	return rv
}

func (m_ MenuItem) SetTag(value int) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setTag:"), value)
}

func (m_ MenuItem) State() ControlStateValue {
	rv := objc.CallMethod[ControlStateValue](m_, objc.GetSelector("state"))
	return rv
}

func (m_ MenuItem) SetState(value ControlStateValue) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setState:"), value)
}

func (m_ MenuItem) Image() Image {
	rv := objc.CallMethod[Image](m_, objc.GetSelector("image"))
	return rv
}

func (m_ MenuItem) SetImage(value IImage) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setImage:"), objc.ExtractPtr(value))
}

func (m_ MenuItem) OnStateImage() Image {
	rv := objc.CallMethod[Image](m_, objc.GetSelector("onStateImage"))
	return rv
}

func (m_ MenuItem) SetOnStateImage(value IImage) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setOnStateImage:"), objc.ExtractPtr(value))
}

func (m_ MenuItem) OffStateImage() Image {
	rv := objc.CallMethod[Image](m_, objc.GetSelector("offStateImage"))
	return rv
}

func (m_ MenuItem) SetOffStateImage(value IImage) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setOffStateImage:"), objc.ExtractPtr(value))
}

func (m_ MenuItem) MixedStateImage() Image {
	rv := objc.CallMethod[Image](m_, objc.GetSelector("mixedStateImage"))
	return rv
}

func (m_ MenuItem) SetMixedStateImage(value IImage) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setMixedStateImage:"), objc.ExtractPtr(value))
}

func (m_ MenuItem) Submenu() Menu {
	rv := objc.CallMethod[Menu](m_, objc.GetSelector("submenu"))
	return rv
}

func (m_ MenuItem) SetSubmenu(value IMenu) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setSubmenu:"), objc.ExtractPtr(value))
}

func (m_ MenuItem) HasSubmenu() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("hasSubmenu"))
	return rv
}

// weak property
func (m_ MenuItem) ParentItem() MenuItem {
	rv := objc.CallMethod[MenuItem](m_, objc.GetSelector("parentItem"))
	return rv
}

func (m_ MenuItem) IsSeparatorItem() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("isSeparatorItem"))
	return rv
}

// weak property
func (m_ MenuItem) Menu() Menu {
	rv := objc.CallMethod[Menu](m_, objc.GetSelector("menu"))
	return rv
}

// weak property
func (m_ MenuItem) SetMenu(value IMenu) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setMenu:"), objc.ExtractPtr(value))
}

func (m_ MenuItem) KeyEquivalent() string {
	rv := objc.CallMethod[string](m_, objc.GetSelector("keyEquivalent"))
	return rv
}

func (m_ MenuItem) SetKeyEquivalent(value string) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setKeyEquivalent:"), value)
}

func (m_ MenuItem) KeyEquivalentModifierMask() EventModifierFlags {
	rv := objc.CallMethod[EventModifierFlags](m_, objc.GetSelector("keyEquivalentModifierMask"))
	return rv
}

func (m_ MenuItem) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setKeyEquivalentModifierMask:"), value)
}

func (mc _MenuItemClass) UsesUserKeyEquivalents() bool {
	rv := objc.CallMethod[bool](mc, objc.GetSelector("usesUserKeyEquivalents"))
	return rv
}

func (mc _MenuItemClass) SetUsesUserKeyEquivalents(value bool) {
	objc.CallMethod[objc.Void](mc, objc.GetSelector("setUsesUserKeyEquivalents:"), value)
}

func (m_ MenuItem) UserKeyEquivalent() string {
	rv := objc.CallMethod[string](m_, objc.GetSelector("userKeyEquivalent"))
	return rv
}

func (m_ MenuItem) AllowsAutomaticKeyEquivalentLocalization() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("allowsAutomaticKeyEquivalentLocalization"))
	return rv
}

func (m_ MenuItem) SetAllowsAutomaticKeyEquivalentLocalization(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAllowsAutomaticKeyEquivalentLocalization:"), value)
}

func (m_ MenuItem) AllowsAutomaticKeyEquivalentMirroring() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("allowsAutomaticKeyEquivalentMirroring"))
	return rv
}

func (m_ MenuItem) SetAllowsAutomaticKeyEquivalentMirroring(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAllowsAutomaticKeyEquivalentMirroring:"), value)
}

func (m_ MenuItem) AllowsKeyEquivalentWhenHidden() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("allowsKeyEquivalentWhenHidden"))
	return rv
}

func (m_ MenuItem) SetAllowsKeyEquivalentWhenHidden(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAllowsKeyEquivalentWhenHidden:"), value)
}

func (m_ MenuItem) IsAlternate() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("isAlternate"))
	return rv
}

func (m_ MenuItem) SetAlternate(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAlternate:"), value)
}

func (m_ MenuItem) IndentationLevel() int {
	rv := objc.CallMethod[int](m_, objc.GetSelector("indentationLevel"))
	return rv
}

func (m_ MenuItem) SetIndentationLevel(value int) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setIndentationLevel:"), value)
}

func (m_ MenuItem) ToolTip() string {
	rv := objc.CallMethod[string](m_, objc.GetSelector("toolTip"))
	return rv
}

func (m_ MenuItem) SetToolTip(value string) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setToolTip:"), value)
}

func (m_ MenuItem) RepresentedObject() objc.Object {
	rv := objc.CallMethod[objc.Object](m_, objc.GetSelector("representedObject"))
	return rv
}

func (m_ MenuItem) SetRepresentedObject(value objc.IObject) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setRepresentedObject:"), objc.ExtractPtr(value))
}

func (m_ MenuItem) View() View {
	rv := objc.CallMethod[View](m_, objc.GetSelector("view"))
	return rv
}

func (m_ MenuItem) SetView(value IView) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setView:"), objc.ExtractPtr(value))
}

func (m_ MenuItem) IsHighlighted() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("isHighlighted"))
	return rv
}
