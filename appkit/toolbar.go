// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var ToolbarClass = _ToolbarClass{objc.GetClass("NSToolbar")}

type _ToolbarClass struct {
	objc.Class
}

type IToolbar interface {
	objc.IObject
	InsertItemWithItemIdentifier_AtIndex(itemIdentifier ToolbarItemIdentifier, index int)
	RemoveItemAtIndex(index int)
	SetConfigurationFromDictionary(configDict map[string]objc.IObject)
	RunCustomizationPalette(sender objc.IObject)
	ValidateVisibleItems()
	Delegate() ToolbarDelegateWrapper
	SetDelegate(value ToolbarDelegate)
	Identifier() ToolbarIdentifier
	IsVisible() bool
	SetVisible(value bool)
	DisplayMode() ToolbarDisplayMode
	SetDisplayMode(value ToolbarDisplayMode)
	ShowsBaselineSeparator() bool
	SetShowsBaselineSeparator(value bool)
	AllowsUserCustomization() bool
	SetAllowsUserCustomization(value bool)
	AllowsExtensionItems() bool
	SetAllowsExtensionItems(value bool)
	Items() []ToolbarItem
	VisibleItems() []ToolbarItem
	CenteredItemIdentifiers() foundation.Set
	SetCenteredItemIdentifiers(value foundation.ISet)
	SelectedItemIdentifier() ToolbarItemIdentifier
	SetSelectedItemIdentifier(value ToolbarItemIdentifier)
	AutosavesConfiguration() bool
	SetAutosavesConfiguration(value bool)
	ConfigurationDictionary() map[string]objc.Object
	CustomizationPaletteIsRunning() bool
	// deprecated
	CenteredItemIdentifier() ToolbarItemIdentifier
	// deprecated
	SetCenteredItemIdentifier(value ToolbarItemIdentifier)
	// deprecated
	FullScreenAccessoryView() View
	// deprecated
	SetFullScreenAccessoryView(value IView)
	// deprecated
	FullScreenAccessoryViewMinHeight() float64
	// deprecated
	SetFullScreenAccessoryViewMinHeight(value float64)
	// deprecated
	FullScreenAccessoryViewMaxHeight() float64
	// deprecated
	SetFullScreenAccessoryViewMaxHeight(value float64)
	// deprecated
	SizeMode() ToolbarSizeMode
	// deprecated
	SetSizeMode(value ToolbarSizeMode)
}

type Toolbar struct {
	objc.Object
}

func MakeToolbar(ptr unsafe.Pointer) Toolbar {
	return Toolbar{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ Toolbar) InitWithIdentifier(identifier ToolbarIdentifier) Toolbar {
	rv := ffi.CallMethod[Toolbar](t_, "initWithIdentifier:", identifier)
	return rv
}

func (t_ Toolbar) Init() Toolbar {
	rv := ffi.CallMethod[Toolbar](t_, "init")
	return rv
}

func (tc _ToolbarClass) Alloc() Toolbar {
	rv := ffi.CallMethod[Toolbar](tc, "alloc")
	return rv
}

func (tc _ToolbarClass) New() Toolbar {
	rv := ffi.CallMethod[Toolbar](tc, "new")
	rv.Autorelease()
	return rv
}

func NewToolbar() Toolbar {
	return ToolbarClass.New()
}

func (t_ Toolbar) InsertItemWithItemIdentifier_AtIndex(itemIdentifier ToolbarItemIdentifier, index int) {
	ffi.CallMethod[ffi.Void](t_, "insertItemWithItemIdentifier:atIndex:", itemIdentifier, index)
}

func (t_ Toolbar) RemoveItemAtIndex(index int) {
	ffi.CallMethod[ffi.Void](t_, "removeItemAtIndex:", index)
}

func (t_ Toolbar) SetConfigurationFromDictionary(configDict map[string]objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "setConfigurationFromDictionary:", configDict)
}

func (t_ Toolbar) RunCustomizationPalette(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "runCustomizationPalette:", sender)
}

func (t_ Toolbar) ValidateVisibleItems() {
	ffi.CallMethod[ffi.Void](t_, "validateVisibleItems")
}

func (t_ Toolbar) Delegate() ToolbarDelegateWrapper {
	rv := ffi.CallMethod[ToolbarDelegateWrapper](t_, "delegate")
	return rv
}

func (t_ Toolbar) SetDelegate(value ToolbarDelegate) {
	po := ffi.CreateProtocol("NSToolbarDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](t_, "setDelegate:", po)
}

func (t_ Toolbar) Identifier() ToolbarIdentifier {
	rv := ffi.CallMethod[ToolbarIdentifier](t_, "identifier")
	return rv
}

func (t_ Toolbar) IsVisible() bool {
	rv := ffi.CallMethod[bool](t_, "isVisible")
	return rv
}

func (t_ Toolbar) SetVisible(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setVisible:", value)
}

func (t_ Toolbar) DisplayMode() ToolbarDisplayMode {
	rv := ffi.CallMethod[ToolbarDisplayMode](t_, "displayMode")
	return rv
}

func (t_ Toolbar) SetDisplayMode(value ToolbarDisplayMode) {
	ffi.CallMethod[ffi.Void](t_, "setDisplayMode:", value)
}

func (t_ Toolbar) ShowsBaselineSeparator() bool {
	rv := ffi.CallMethod[bool](t_, "showsBaselineSeparator")
	return rv
}

func (t_ Toolbar) SetShowsBaselineSeparator(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setShowsBaselineSeparator:", value)
}

func (t_ Toolbar) AllowsUserCustomization() bool {
	rv := ffi.CallMethod[bool](t_, "allowsUserCustomization")
	return rv
}

func (t_ Toolbar) SetAllowsUserCustomization(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAllowsUserCustomization:", value)
}

func (t_ Toolbar) AllowsExtensionItems() bool {
	rv := ffi.CallMethod[bool](t_, "allowsExtensionItems")
	return rv
}

func (t_ Toolbar) SetAllowsExtensionItems(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAllowsExtensionItems:", value)
}

func (t_ Toolbar) Items() []ToolbarItem {
	rv := ffi.CallMethod[[]ToolbarItem](t_, "items")
	return rv
}

func (t_ Toolbar) VisibleItems() []ToolbarItem {
	rv := ffi.CallMethod[[]ToolbarItem](t_, "visibleItems")
	return rv
}

func (t_ Toolbar) CenteredItemIdentifiers() foundation.Set {
	rv := ffi.CallMethod[foundation.Set](t_, "centeredItemIdentifiers")
	return rv
}

func (t_ Toolbar) SetCenteredItemIdentifiers(value foundation.ISet) {
	ffi.CallMethod[ffi.Void](t_, "setCenteredItemIdentifiers:", value)
}

func (t_ Toolbar) SelectedItemIdentifier() ToolbarItemIdentifier {
	rv := ffi.CallMethod[ToolbarItemIdentifier](t_, "selectedItemIdentifier")
	return rv
}

func (t_ Toolbar) SetSelectedItemIdentifier(value ToolbarItemIdentifier) {
	ffi.CallMethod[ffi.Void](t_, "setSelectedItemIdentifier:", value)
}

func (t_ Toolbar) AutosavesConfiguration() bool {
	rv := ffi.CallMethod[bool](t_, "autosavesConfiguration")
	return rv
}

func (t_ Toolbar) SetAutosavesConfiguration(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAutosavesConfiguration:", value)
}

func (t_ Toolbar) ConfigurationDictionary() map[string]objc.Object {
	rv := ffi.CallMethod[map[string]objc.Object](t_, "configurationDictionary")
	return rv
}

func (t_ Toolbar) CustomizationPaletteIsRunning() bool {
	rv := ffi.CallMethod[bool](t_, "customizationPaletteIsRunning")
	return rv
}

// deprecated
func (t_ Toolbar) CenteredItemIdentifier() ToolbarItemIdentifier {
	rv := ffi.CallMethod[ToolbarItemIdentifier](t_, "centeredItemIdentifier")
	return rv
}

// deprecated
func (t_ Toolbar) SetCenteredItemIdentifier(value ToolbarItemIdentifier) {
	ffi.CallMethod[ffi.Void](t_, "setCenteredItemIdentifier:", value)
}

// deprecated
func (t_ Toolbar) FullScreenAccessoryView() View {
	rv := ffi.CallMethod[View](t_, "fullScreenAccessoryView")
	return rv
}

// deprecated
func (t_ Toolbar) SetFullScreenAccessoryView(value IView) {
	ffi.CallMethod[ffi.Void](t_, "setFullScreenAccessoryView:", value)
}

// deprecated
func (t_ Toolbar) FullScreenAccessoryViewMinHeight() float64 {
	rv := ffi.CallMethod[float64](t_, "fullScreenAccessoryViewMinHeight")
	return rv
}

// deprecated
func (t_ Toolbar) SetFullScreenAccessoryViewMinHeight(value float64) {
	ffi.CallMethod[ffi.Void](t_, "setFullScreenAccessoryViewMinHeight:", value)
}

// deprecated
func (t_ Toolbar) FullScreenAccessoryViewMaxHeight() float64 {
	rv := ffi.CallMethod[float64](t_, "fullScreenAccessoryViewMaxHeight")
	return rv
}

// deprecated
func (t_ Toolbar) SetFullScreenAccessoryViewMaxHeight(value float64) {
	ffi.CallMethod[ffi.Void](t_, "setFullScreenAccessoryViewMaxHeight:", value)
}

// deprecated
func (t_ Toolbar) SizeMode() ToolbarSizeMode {
	rv := ffi.CallMethod[ToolbarSizeMode](t_, "sizeMode")
	return rv
}

// deprecated
func (t_ Toolbar) SetSizeMode(value ToolbarSizeMode) {
	ffi.CallMethod[ffi.Void](t_, "setSizeMode:", value)
}
