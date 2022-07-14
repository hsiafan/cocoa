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
	rv.Autorelease()
	return rv
}

func (t_ Toolbar) Init() Toolbar {
	rv := ffi.CallMethod[Toolbar](t_, "init")
	rv.Autorelease()
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
	po := ffi.CreateProtocol(value)
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

type ToolbarDelegate interface {
	ImplementsToolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar() bool
	// optional
	Toolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem
	ImplementsToolbarWillAddItem() bool
	// optional
	ToolbarWillAddItem(notification foundation.Notification)
	ImplementsToolbarDidRemoveItem() bool
	// optional
	ToolbarDidRemoveItem(notification foundation.Notification)
	ImplementsToolbarAllowedItemIdentifiers() bool
	// optional
	ToolbarAllowedItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier
	ImplementsToolbarDefaultItemIdentifiers() bool
	// optional
	ToolbarDefaultItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier
	ImplementsToolbarImmovableItemIdentifiers() bool
	// optional
	ToolbarImmovableItemIdentifiers(toolbar Toolbar) foundation.ISet
	ImplementsToolbarSelectableItemIdentifiers() bool
	// optional
	ToolbarSelectableItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier
	ImplementsToolbar_ItemIdentifier_CanBeInsertedAtIndex() bool
	// optional
	Toolbar_ItemIdentifier_CanBeInsertedAtIndex(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, index int) bool
}

type ToolbarDelegateImpl struct {
	_Toolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar func(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem
	_ToolbarWillAddItem                                      func(notification foundation.Notification)
	_ToolbarDidRemoveItem                                    func(notification foundation.Notification)
	_ToolbarAllowedItemIdentifiers                           func(toolbar Toolbar) []ToolbarItemIdentifier
	_ToolbarDefaultItemIdentifiers                           func(toolbar Toolbar) []ToolbarItemIdentifier
	_ToolbarImmovableItemIdentifiers                         func(toolbar Toolbar) foundation.ISet
	_ToolbarSelectableItemIdentifiers                        func(toolbar Toolbar) []ToolbarItemIdentifier
	_Toolbar_ItemIdentifier_CanBeInsertedAtIndex             func(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, index int) bool
}

func (di *ToolbarDelegateImpl) ImplementsToolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar() bool {
	return di._Toolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar != nil
}

func (di *ToolbarDelegateImpl) SetToolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar(f func(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem) {
	di._Toolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar = f
}

func (di *ToolbarDelegateImpl) Toolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem {
	return di._Toolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar(toolbar, itemIdentifier, flag)
}
func (di *ToolbarDelegateImpl) ImplementsToolbarWillAddItem() bool {
	return di._ToolbarWillAddItem != nil
}

func (di *ToolbarDelegateImpl) SetToolbarWillAddItem(f func(notification foundation.Notification)) {
	di._ToolbarWillAddItem = f
}

func (di *ToolbarDelegateImpl) ToolbarWillAddItem(notification foundation.Notification) {
	di._ToolbarWillAddItem(notification)
}
func (di *ToolbarDelegateImpl) ImplementsToolbarDidRemoveItem() bool {
	return di._ToolbarDidRemoveItem != nil
}

func (di *ToolbarDelegateImpl) SetToolbarDidRemoveItem(f func(notification foundation.Notification)) {
	di._ToolbarDidRemoveItem = f
}

func (di *ToolbarDelegateImpl) ToolbarDidRemoveItem(notification foundation.Notification) {
	di._ToolbarDidRemoveItem(notification)
}
func (di *ToolbarDelegateImpl) ImplementsToolbarAllowedItemIdentifiers() bool {
	return di._ToolbarAllowedItemIdentifiers != nil
}

func (di *ToolbarDelegateImpl) SetToolbarAllowedItemIdentifiers(f func(toolbar Toolbar) []ToolbarItemIdentifier) {
	di._ToolbarAllowedItemIdentifiers = f
}

func (di *ToolbarDelegateImpl) ToolbarAllowedItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier {
	return di._ToolbarAllowedItemIdentifiers(toolbar)
}
func (di *ToolbarDelegateImpl) ImplementsToolbarDefaultItemIdentifiers() bool {
	return di._ToolbarDefaultItemIdentifiers != nil
}

func (di *ToolbarDelegateImpl) SetToolbarDefaultItemIdentifiers(f func(toolbar Toolbar) []ToolbarItemIdentifier) {
	di._ToolbarDefaultItemIdentifiers = f
}

func (di *ToolbarDelegateImpl) ToolbarDefaultItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier {
	return di._ToolbarDefaultItemIdentifiers(toolbar)
}
func (di *ToolbarDelegateImpl) ImplementsToolbarImmovableItemIdentifiers() bool {
	return di._ToolbarImmovableItemIdentifiers != nil
}

func (di *ToolbarDelegateImpl) SetToolbarImmovableItemIdentifiers(f func(toolbar Toolbar) foundation.ISet) {
	di._ToolbarImmovableItemIdentifiers = f
}

func (di *ToolbarDelegateImpl) ToolbarImmovableItemIdentifiers(toolbar Toolbar) foundation.ISet {
	return di._ToolbarImmovableItemIdentifiers(toolbar)
}
func (di *ToolbarDelegateImpl) ImplementsToolbarSelectableItemIdentifiers() bool {
	return di._ToolbarSelectableItemIdentifiers != nil
}

func (di *ToolbarDelegateImpl) SetToolbarSelectableItemIdentifiers(f func(toolbar Toolbar) []ToolbarItemIdentifier) {
	di._ToolbarSelectableItemIdentifiers = f
}

func (di *ToolbarDelegateImpl) ToolbarSelectableItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier {
	return di._ToolbarSelectableItemIdentifiers(toolbar)
}
func (di *ToolbarDelegateImpl) ImplementsToolbar_ItemIdentifier_CanBeInsertedAtIndex() bool {
	return di._Toolbar_ItemIdentifier_CanBeInsertedAtIndex != nil
}

func (di *ToolbarDelegateImpl) SetToolbar_ItemIdentifier_CanBeInsertedAtIndex(f func(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, index int) bool) {
	di._Toolbar_ItemIdentifier_CanBeInsertedAtIndex = f
}

func (di *ToolbarDelegateImpl) Toolbar_ItemIdentifier_CanBeInsertedAtIndex(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, index int) bool {
	return di._Toolbar_ItemIdentifier_CanBeInsertedAtIndex(toolbar, itemIdentifier, index)
}

type ToolbarDelegateWrapper struct {
	objc.Object
}

func (t_ *ToolbarDelegateWrapper) ImplementsToolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbar:itemForItemIdentifier:willBeInsertedIntoToolbar:"))
}

func (t_ ToolbarDelegateWrapper) Toolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar(toolbar IToolbar, itemIdentifier ToolbarItemIdentifier, flag bool) ToolbarItem {
	rv := ffi.CallMethod[ToolbarItem](t_, "toolbar:itemForItemIdentifier:willBeInsertedIntoToolbar:", toolbar, itemIdentifier, flag)
	return rv
}

func (t_ *ToolbarDelegateWrapper) ImplementsToolbarWillAddItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbarWillAddItem:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarWillAddItem(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](t_, "toolbarWillAddItem:", notification)
}

func (t_ *ToolbarDelegateWrapper) ImplementsToolbarDidRemoveItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbarDidRemoveItem:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarDidRemoveItem(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](t_, "toolbarDidRemoveItem:", notification)
}

func (t_ *ToolbarDelegateWrapper) ImplementsToolbarAllowedItemIdentifiers() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbarAllowedItemIdentifiers:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarAllowedItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier {
	rv := ffi.CallMethod[[]ToolbarItemIdentifier](t_, "toolbarAllowedItemIdentifiers:", toolbar)
	return rv
}

func (t_ *ToolbarDelegateWrapper) ImplementsToolbarDefaultItemIdentifiers() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbarDefaultItemIdentifiers:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarDefaultItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier {
	rv := ffi.CallMethod[[]ToolbarItemIdentifier](t_, "toolbarDefaultItemIdentifiers:", toolbar)
	return rv
}

func (t_ *ToolbarDelegateWrapper) ImplementsToolbarImmovableItemIdentifiers() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbarImmovableItemIdentifiers:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarImmovableItemIdentifiers(toolbar IToolbar) foundation.Set {
	rv := ffi.CallMethod[foundation.Set](t_, "toolbarImmovableItemIdentifiers:", toolbar)
	return rv
}

func (t_ *ToolbarDelegateWrapper) ImplementsToolbarSelectableItemIdentifiers() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbarSelectableItemIdentifiers:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarSelectableItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier {
	rv := ffi.CallMethod[[]ToolbarItemIdentifier](t_, "toolbarSelectableItemIdentifiers:", toolbar)
	return rv
}

func (t_ *ToolbarDelegateWrapper) ImplementsToolbar_ItemIdentifier_CanBeInsertedAtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbar:itemIdentifier:canBeInsertedAtIndex:"))
}

func (t_ ToolbarDelegateWrapper) Toolbar_ItemIdentifier_CanBeInsertedAtIndex(toolbar IToolbar, itemIdentifier ToolbarItemIdentifier, index int) bool {
	rv := ffi.CallMethod[bool](t_, "toolbar:itemIdentifier:canBeInsertedAtIndex:", toolbar, itemIdentifier, index)
	return rv
}

var ToolbarItemClass = _ToolbarItemClass{objc.GetClass("NSToolbarItem")}

type _ToolbarItemClass struct {
	objc.Class
}

type IToolbarItem interface {
	objc.IObject
	Validate()
	ItemIdentifier() ToolbarItemIdentifier
	PossibleLabels() foundation.Set
	SetPossibleLabels(value foundation.ISet)
	Label() string
	SetLabel(value string)
	PaletteLabel() string
	SetPaletteLabel(value string)
	Title() string
	SetTitle(value string)
	ToolTip() string
	SetToolTip(value string)
	Image() Image
	SetImage(value IImage)
	View() View
	SetView(value IView)
	Target() objc.Object
	SetTarget(value objc.IObject)
	Action() objc.Selector
	SetAction(value objc.Selector)
	MenuFormRepresentation() MenuItem
	SetMenuFormRepresentation(value IMenuItem)
	IsVisible() bool
	IsBordered() bool
	SetBordered(value bool)
	IsNavigational() bool
	SetNavigational(value bool)
	IsEnabled() bool
	SetEnabled(value bool)
	AllowsDuplicatesInToolbar() bool
	VisibilityPriority() ToolbarItemVisibilityPriority
	SetVisibilityPriority(value ToolbarItemVisibilityPriority)
	Tag() int
	SetTag(value int)
	Toolbar() Toolbar
	Autovalidates() bool
	SetAutovalidates(value bool)
	// deprecated
	MinSize() foundation.Size
	// deprecated
	SetMinSize(value foundation.Size)
	// deprecated
	MaxSize() foundation.Size
	// deprecated
	SetMaxSize(value foundation.Size)
}

type ToolbarItem struct {
	objc.Object
}

func MakeToolbarItem(ptr unsafe.Pointer) ToolbarItem {
	return ToolbarItem{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ ToolbarItem) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) ToolbarItem {
	rv := ffi.CallMethod[ToolbarItem](t_, "initWithItemIdentifier:", itemIdentifier)
	rv.Autorelease()
	return rv
}

func (tc _ToolbarItemClass) Alloc() ToolbarItem {
	rv := ffi.CallMethod[ToolbarItem](tc, "alloc")
	return rv
}

func (t_ ToolbarItem) Init() ToolbarItem {
	rv := ffi.CallMethod[ToolbarItem](t_, "init")
	rv.Autorelease()
	return rv
}

func (tc _ToolbarItemClass) New() ToolbarItem {
	rv := ffi.CallMethod[ToolbarItem](tc, "new")
	rv.Autorelease()
	return rv
}

func NewToolbarItem() ToolbarItem {
	return ToolbarItemClass.New()
}

func (t_ ToolbarItem) Validate() {
	ffi.CallMethod[ffi.Void](t_, "validate")
}

func (t_ ToolbarItem) ItemIdentifier() ToolbarItemIdentifier {
	rv := ffi.CallMethod[ToolbarItemIdentifier](t_, "itemIdentifier")
	return rv
}

func (t_ ToolbarItem) PossibleLabels() foundation.Set {
	rv := ffi.CallMethod[foundation.Set](t_, "possibleLabels")
	return rv
}

func (t_ ToolbarItem) SetPossibleLabels(value foundation.ISet) {
	ffi.CallMethod[ffi.Void](t_, "setPossibleLabels:", value)
}

func (t_ ToolbarItem) Label() string {
	rv := ffi.CallMethod[string](t_, "label")
	return rv
}

func (t_ ToolbarItem) SetLabel(value string) {
	ffi.CallMethod[ffi.Void](t_, "setLabel:", value)
}

func (t_ ToolbarItem) PaletteLabel() string {
	rv := ffi.CallMethod[string](t_, "paletteLabel")
	return rv
}

func (t_ ToolbarItem) SetPaletteLabel(value string) {
	ffi.CallMethod[ffi.Void](t_, "setPaletteLabel:", value)
}

func (t_ ToolbarItem) Title() string {
	rv := ffi.CallMethod[string](t_, "title")
	return rv
}

func (t_ ToolbarItem) SetTitle(value string) {
	ffi.CallMethod[ffi.Void](t_, "setTitle:", value)
}

func (t_ ToolbarItem) ToolTip() string {
	rv := ffi.CallMethod[string](t_, "toolTip")
	return rv
}

func (t_ ToolbarItem) SetToolTip(value string) {
	ffi.CallMethod[ffi.Void](t_, "setToolTip:", value)
}

func (t_ ToolbarItem) Image() Image {
	rv := ffi.CallMethod[Image](t_, "image")
	return rv
}

func (t_ ToolbarItem) SetImage(value IImage) {
	ffi.CallMethod[ffi.Void](t_, "setImage:", value)
}

func (t_ ToolbarItem) View() View {
	rv := ffi.CallMethod[View](t_, "view")
	return rv
}

func (t_ ToolbarItem) SetView(value IView) {
	ffi.CallMethod[ffi.Void](t_, "setView:", value)
}

func (t_ ToolbarItem) Target() objc.Object {
	rv := ffi.CallMethod[objc.Object](t_, "target")
	return rv
}

func (t_ ToolbarItem) SetTarget(value objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "setTarget:", value)
}

func (t_ ToolbarItem) Action() objc.Selector {
	rv := ffi.CallMethod[objc.Selector](t_, "action")
	return rv
}

func (t_ ToolbarItem) SetAction(value objc.Selector) {
	ffi.CallMethod[ffi.Void](t_, "setAction:", value)
}

func (t_ ToolbarItem) MenuFormRepresentation() MenuItem {
	rv := ffi.CallMethod[MenuItem](t_, "menuFormRepresentation")
	return rv
}

func (t_ ToolbarItem) SetMenuFormRepresentation(value IMenuItem) {
	ffi.CallMethod[ffi.Void](t_, "setMenuFormRepresentation:", value)
}

func (t_ ToolbarItem) IsVisible() bool {
	rv := ffi.CallMethod[bool](t_, "isVisible")
	return rv
}

func (t_ ToolbarItem) IsBordered() bool {
	rv := ffi.CallMethod[bool](t_, "isBordered")
	return rv
}

func (t_ ToolbarItem) SetBordered(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setBordered:", value)
}

func (t_ ToolbarItem) IsNavigational() bool {
	rv := ffi.CallMethod[bool](t_, "isNavigational")
	return rv
}

func (t_ ToolbarItem) SetNavigational(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setNavigational:", value)
}

func (t_ ToolbarItem) IsEnabled() bool {
	rv := ffi.CallMethod[bool](t_, "isEnabled")
	return rv
}

func (t_ ToolbarItem) SetEnabled(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setEnabled:", value)
}

func (t_ ToolbarItem) AllowsDuplicatesInToolbar() bool {
	rv := ffi.CallMethod[bool](t_, "allowsDuplicatesInToolbar")
	return rv
}

func (t_ ToolbarItem) VisibilityPriority() ToolbarItemVisibilityPriority {
	rv := ffi.CallMethod[ToolbarItemVisibilityPriority](t_, "visibilityPriority")
	return rv
}

func (t_ ToolbarItem) SetVisibilityPriority(value ToolbarItemVisibilityPriority) {
	ffi.CallMethod[ffi.Void](t_, "setVisibilityPriority:", value)
}

func (t_ ToolbarItem) Tag() int {
	rv := ffi.CallMethod[int](t_, "tag")
	return rv
}

func (t_ ToolbarItem) SetTag(value int) {
	ffi.CallMethod[ffi.Void](t_, "setTag:", value)
}

func (t_ ToolbarItem) Toolbar() Toolbar {
	rv := ffi.CallMethod[Toolbar](t_, "toolbar")
	return rv
}

func (t_ ToolbarItem) Autovalidates() bool {
	rv := ffi.CallMethod[bool](t_, "autovalidates")
	return rv
}

func (t_ ToolbarItem) SetAutovalidates(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAutovalidates:", value)
}

// deprecated
func (t_ ToolbarItem) MinSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](t_, "minSize")
	return rv
}

// deprecated
func (t_ ToolbarItem) SetMinSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](t_, "setMinSize:", value)
}

// deprecated
func (t_ ToolbarItem) MaxSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](t_, "maxSize")
	return rv
}

// deprecated
func (t_ ToolbarItem) SetMaxSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](t_, "setMaxSize:", value)
}

type ToolbarItemValidation interface {
	// required
	ValidateToolbarItem(item ToolbarItem) bool
}

type ToolbarItemValidationWrapper struct {
	objc.Object
}

func (t_ ToolbarItemValidationWrapper) ValidateToolbarItem(item IToolbarItem) bool {
	rv := ffi.CallMethod[bool](t_, "validateToolbarItem:", item)
	return rv
}

var ToolbarItemGroupClass = _ToolbarItemGroupClass{objc.GetClass("NSToolbarItemGroup")}

type _ToolbarItemGroupClass struct {
	objc.Class
}

type IToolbarItemGroup interface {
	IToolbarItem
	IsSelectedAtIndex(index int) bool
	SetSelected_AtIndex(selected bool, index int)
	Subitems() []ToolbarItem
	SetSubitems(value []IToolbarItem)
	SelectedIndex() int
	SetSelectedIndex(value int)
	ControlRepresentation() ToolbarItemGroupControlRepresentation
	SetControlRepresentation(value ToolbarItemGroupControlRepresentation)
	SelectionMode() ToolbarItemGroupSelectionMode
	SetSelectionMode(value ToolbarItemGroupSelectionMode)
}

type ToolbarItemGroup struct {
	ToolbarItem
}

func MakeToolbarItemGroup(ptr unsafe.Pointer) ToolbarItemGroup {
	return ToolbarItemGroup{
		ToolbarItem: MakeToolbarItem(ptr),
	}
}

func (tc _ToolbarItemGroupClass) GroupWithItemIdentifier_Images_SelectionMode_Labels_Target_Action(itemIdentifier ToolbarItemIdentifier, images []IImage, selectionMode ToolbarItemGroupSelectionMode, labels []string, target objc.IObject, action objc.Selector) ToolbarItemGroup {
	rv := ffi.CallMethod[ToolbarItemGroup](tc, "groupWithItemIdentifier:images:selectionMode:labels:target:action:", itemIdentifier, images, selectionMode, labels, target, action)
	return rv
}

func (tc _ToolbarItemGroupClass) GroupWithItemIdentifier_Titles_SelectionMode_Labels_Target_Action(itemIdentifier ToolbarItemIdentifier, titles []string, selectionMode ToolbarItemGroupSelectionMode, labels []string, target objc.IObject, action objc.Selector) ToolbarItemGroup {
	rv := ffi.CallMethod[ToolbarItemGroup](tc, "groupWithItemIdentifier:titles:selectionMode:labels:target:action:", itemIdentifier, titles, selectionMode, labels, target, action)
	return rv
}

func (t_ ToolbarItemGroup) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) ToolbarItemGroup {
	rv := ffi.CallMethod[ToolbarItemGroup](t_, "initWithItemIdentifier:", itemIdentifier)
	rv.Autorelease()
	return rv
}

func (tc _ToolbarItemGroupClass) Alloc() ToolbarItemGroup {
	rv := ffi.CallMethod[ToolbarItemGroup](tc, "alloc")
	return rv
}

func (t_ ToolbarItemGroup) Init() ToolbarItemGroup {
	rv := ffi.CallMethod[ToolbarItemGroup](t_, "init")
	rv.Autorelease()
	return rv
}

func (tc _ToolbarItemGroupClass) New() ToolbarItemGroup {
	rv := ffi.CallMethod[ToolbarItemGroup](tc, "new")
	rv.Autorelease()
	return rv
}

func NewToolbarItemGroup() ToolbarItemGroup {
	return ToolbarItemGroupClass.New()
}

func (t_ ToolbarItemGroup) IsSelectedAtIndex(index int) bool {
	rv := ffi.CallMethod[bool](t_, "isSelectedAtIndex:", index)
	return rv
}

func (t_ ToolbarItemGroup) SetSelected_AtIndex(selected bool, index int) {
	ffi.CallMethod[ffi.Void](t_, "setSelected:atIndex:", selected, index)
}

func (t_ ToolbarItemGroup) Subitems() []ToolbarItem {
	rv := ffi.CallMethod[[]ToolbarItem](t_, "subitems")
	return rv
}

func (t_ ToolbarItemGroup) SetSubitems(value []IToolbarItem) {
	ffi.CallMethod[ffi.Void](t_, "setSubitems:", value)
}

func (t_ ToolbarItemGroup) SelectedIndex() int {
	rv := ffi.CallMethod[int](t_, "selectedIndex")
	return rv
}

func (t_ ToolbarItemGroup) SetSelectedIndex(value int) {
	ffi.CallMethod[ffi.Void](t_, "setSelectedIndex:", value)
}

func (t_ ToolbarItemGroup) ControlRepresentation() ToolbarItemGroupControlRepresentation {
	rv := ffi.CallMethod[ToolbarItemGroupControlRepresentation](t_, "controlRepresentation")
	return rv
}

func (t_ ToolbarItemGroup) SetControlRepresentation(value ToolbarItemGroupControlRepresentation) {
	ffi.CallMethod[ffi.Void](t_, "setControlRepresentation:", value)
}

func (t_ ToolbarItemGroup) SelectionMode() ToolbarItemGroupSelectionMode {
	rv := ffi.CallMethod[ToolbarItemGroupSelectionMode](t_, "selectionMode")
	return rv
}

func (t_ ToolbarItemGroup) SetSelectionMode(value ToolbarItemGroupSelectionMode) {
	ffi.CallMethod[ffi.Void](t_, "setSelectionMode:", value)
}

var MenuToolbarItemClass = _MenuToolbarItemClass{objc.GetClass("NSMenuToolbarItem")}

type _MenuToolbarItemClass struct {
	objc.Class
}

type IMenuToolbarItem interface {
	IToolbarItem
	ShowsIndicator() bool
	SetShowsIndicator(value bool)
	Menu() Menu
	SetMenu(value IMenu)
}

type MenuToolbarItem struct {
	ToolbarItem
}

func MakeMenuToolbarItem(ptr unsafe.Pointer) MenuToolbarItem {
	return MenuToolbarItem{
		ToolbarItem: MakeToolbarItem(ptr),
	}
}

func (m_ MenuToolbarItem) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) MenuToolbarItem {
	rv := ffi.CallMethod[MenuToolbarItem](m_, "initWithItemIdentifier:", itemIdentifier)
	rv.Autorelease()
	return rv
}

func (mc _MenuToolbarItemClass) Alloc() MenuToolbarItem {
	rv := ffi.CallMethod[MenuToolbarItem](mc, "alloc")
	return rv
}

func (m_ MenuToolbarItem) Init() MenuToolbarItem {
	rv := ffi.CallMethod[MenuToolbarItem](m_, "init")
	rv.Autorelease()
	return rv
}

func (mc _MenuToolbarItemClass) New() MenuToolbarItem {
	rv := ffi.CallMethod[MenuToolbarItem](mc, "new")
	rv.Autorelease()
	return rv
}

func NewMenuToolbarItem() MenuToolbarItem {
	return MenuToolbarItemClass.New()
}

func (m_ MenuToolbarItem) ShowsIndicator() bool {
	rv := ffi.CallMethod[bool](m_, "showsIndicator")
	return rv
}

func (m_ MenuToolbarItem) SetShowsIndicator(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setShowsIndicator:", value)
}

func (m_ MenuToolbarItem) Menu() Menu {
	rv := ffi.CallMethod[Menu](m_, "menu")
	return rv
}

func (m_ MenuToolbarItem) SetMenu(value IMenu) {
	ffi.CallMethod[ffi.Void](m_, "setMenu:", value)
}

var SearchToolbarItemClass = _SearchToolbarItemClass{objc.GetClass("NSSearchToolbarItem")}

type _SearchToolbarItemClass struct {
	objc.Class
}

type ISearchToolbarItem interface {
	IToolbarItem
	BeginSearchInteraction()
	EndSearchInteraction()
	PreferredWidthForSearchField() float64
	SetPreferredWidthForSearchField(value float64)
	ResignsFirstResponderWithCancel() bool
	SetResignsFirstResponderWithCancel(value bool)
	SearchField() SearchField
	SetSearchField(value ISearchField)
}

type SearchToolbarItem struct {
	ToolbarItem
}

func MakeSearchToolbarItem(ptr unsafe.Pointer) SearchToolbarItem {
	return SearchToolbarItem{
		ToolbarItem: MakeToolbarItem(ptr),
	}
}

func (s_ SearchToolbarItem) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) SearchToolbarItem {
	rv := ffi.CallMethod[SearchToolbarItem](s_, "initWithItemIdentifier:", itemIdentifier)
	rv.Autorelease()
	return rv
}

func (sc _SearchToolbarItemClass) Alloc() SearchToolbarItem {
	rv := ffi.CallMethod[SearchToolbarItem](sc, "alloc")
	return rv
}

func (s_ SearchToolbarItem) Init() SearchToolbarItem {
	rv := ffi.CallMethod[SearchToolbarItem](s_, "init")
	rv.Autorelease()
	return rv
}

func (sc _SearchToolbarItemClass) New() SearchToolbarItem {
	rv := ffi.CallMethod[SearchToolbarItem](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSearchToolbarItem() SearchToolbarItem {
	return SearchToolbarItemClass.New()
}

func (s_ SearchToolbarItem) BeginSearchInteraction() {
	ffi.CallMethod[ffi.Void](s_, "beginSearchInteraction")
}

func (s_ SearchToolbarItem) EndSearchInteraction() {
	ffi.CallMethod[ffi.Void](s_, "endSearchInteraction")
}

func (s_ SearchToolbarItem) PreferredWidthForSearchField() float64 {
	rv := ffi.CallMethod[float64](s_, "preferredWidthForSearchField")
	return rv
}

func (s_ SearchToolbarItem) SetPreferredWidthForSearchField(value float64) {
	ffi.CallMethod[ffi.Void](s_, "setPreferredWidthForSearchField:", value)
}

func (s_ SearchToolbarItem) ResignsFirstResponderWithCancel() bool {
	rv := ffi.CallMethod[bool](s_, "resignsFirstResponderWithCancel")
	return rv
}

func (s_ SearchToolbarItem) SetResignsFirstResponderWithCancel(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setResignsFirstResponderWithCancel:", value)
}

func (s_ SearchToolbarItem) SearchField() SearchField {
	rv := ffi.CallMethod[SearchField](s_, "searchField")
	return rv
}

func (s_ SearchToolbarItem) SetSearchField(value ISearchField) {
	ffi.CallMethod[ffi.Void](s_, "setSearchField:", value)
}

var TrackingSeparatorToolbarItemClass = _TrackingSeparatorToolbarItemClass{objc.GetClass("NSTrackingSeparatorToolbarItem")}

type _TrackingSeparatorToolbarItemClass struct {
	objc.Class
}

type ITrackingSeparatorToolbarItem interface {
	IToolbarItem
	DividerIndex() int
	SetDividerIndex(value int)
	SplitView() SplitView
	SetSplitView(value ISplitView)
}

type TrackingSeparatorToolbarItem struct {
	ToolbarItem
}

func MakeTrackingSeparatorToolbarItem(ptr unsafe.Pointer) TrackingSeparatorToolbarItem {
	return TrackingSeparatorToolbarItem{
		ToolbarItem: MakeToolbarItem(ptr),
	}
}

func (tc _TrackingSeparatorToolbarItemClass) TrackingSeparatorToolbarItemWithIdentifier_SplitView_DividerIndex(identifier ToolbarItemIdentifier, splitView ISplitView, dividerIndex int) TrackingSeparatorToolbarItem {
	rv := ffi.CallMethod[TrackingSeparatorToolbarItem](tc, "trackingSeparatorToolbarItemWithIdentifier:splitView:dividerIndex:", identifier, splitView, dividerIndex)
	return rv
}

func (t_ TrackingSeparatorToolbarItem) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) TrackingSeparatorToolbarItem {
	rv := ffi.CallMethod[TrackingSeparatorToolbarItem](t_, "initWithItemIdentifier:", itemIdentifier)
	rv.Autorelease()
	return rv
}

func (tc _TrackingSeparatorToolbarItemClass) Alloc() TrackingSeparatorToolbarItem {
	rv := ffi.CallMethod[TrackingSeparatorToolbarItem](tc, "alloc")
	return rv
}

func (t_ TrackingSeparatorToolbarItem) Init() TrackingSeparatorToolbarItem {
	rv := ffi.CallMethod[TrackingSeparatorToolbarItem](t_, "init")
	rv.Autorelease()
	return rv
}

func (tc _TrackingSeparatorToolbarItemClass) New() TrackingSeparatorToolbarItem {
	rv := ffi.CallMethod[TrackingSeparatorToolbarItem](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTrackingSeparatorToolbarItem() TrackingSeparatorToolbarItem {
	return TrackingSeparatorToolbarItemClass.New()
}

func (t_ TrackingSeparatorToolbarItem) DividerIndex() int {
	rv := ffi.CallMethod[int](t_, "dividerIndex")
	return rv
}

func (t_ TrackingSeparatorToolbarItem) SetDividerIndex(value int) {
	ffi.CallMethod[ffi.Void](t_, "setDividerIndex:", value)
}

func (t_ TrackingSeparatorToolbarItem) SplitView() SplitView {
	rv := ffi.CallMethod[SplitView](t_, "splitView")
	return rv
}

func (t_ TrackingSeparatorToolbarItem) SetSplitView(value ISplitView) {
	ffi.CallMethod[ffi.Void](t_, "setSplitView:", value)
}
