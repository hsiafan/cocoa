// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var MenuClass = _MenuClass{objc.GetClass("NSMenu")}

type _MenuClass struct {
	objc.Class
}

type IMenu interface {
	objc.IObject
	InsertItem_AtIndex(newItem IMenuItem, index int)
	InsertItemWithTitle_Action_KeyEquivalent_AtIndex(_string string, selector objc.Selector, charCode string, index int) MenuItem
	AddItem(newItem IMenuItem)
	AddItemWithTitle_Action_KeyEquivalent(_string string, selector objc.Selector, charCode string) MenuItem
	RemoveItem(item IMenuItem)
	RemoveItemAtIndex(index int)
	ItemChanged(item IMenuItem)
	RemoveAllItems()
	ItemWithTag(tag int) MenuItem
	ItemWithTitle(title string) MenuItem
	ItemAtIndex(index int) MenuItem
	IndexOfItem(item IMenuItem) int
	IndexOfItemWithTitle(title string) int
	IndexOfItemWithTag(tag int) int
	IndexOfItemWithTarget_AndAction(target objc.IObject, actionSelector objc.Selector) int
	IndexOfItemWithRepresentedObject(object objc.IObject) int
	IndexOfItemWithSubmenu(submenu IMenu) int
	SetSubmenu_ForItem(menu IMenu, item IMenuItem)
	SubmenuAction(sender objc.IObject)
	// deprecated
	AttachedMenu() Menu
	// deprecated
	IsAttached() bool
	// deprecated
	LocationForSubmenu(submenu IMenu) foundation.Point
	Update()
	PerformKeyEquivalent(event IEvent) bool
	PerformActionForItemAtIndex(index int)
	// deprecated
	SizeToFit()
	// deprecated
	HelpRequested(eventPtr IEvent)
	PopUpMenuPositioningItem_AtLocation_InView(item IMenuItem, location foundation.Point, view IView) bool
	CancelTracking()
	CancelTrackingWithoutAnimation()
	// deprecated
	ContextMenuRepresentation() objc.Object
	// deprecated
	SetContextMenuRepresentation(menuRep objc.IObject)
	// deprecated
	TearOffMenuRepresentation() objc.Object
	// deprecated
	SetTearOffMenuRepresentation(menuRep objc.IObject)
	// deprecated
	SetMenuRepresentation(menuRep objc.IObject)
	// deprecated
	MenuRepresentation() objc.Object
	MenuBarHeight() float64
	NumberOfItems() int
	ItemArray() []MenuItem
	SetItemArray(value []IMenuItem)
	Supermenu() Menu
	SetSupermenu(value IMenu)
	// deprecated
	IsTornOff() bool
	AutoenablesItems() bool
	SetAutoenablesItems(value bool)
	Font() Font
	SetFont(value IFont)
	Title() string
	SetTitle(value string)
	MinimumWidth() float64
	SetMinimumWidth(value float64)
	Size() foundation.Size
	PropertiesToUpdate() MenuProperties
	// deprecated
	MenuChangedMessagesEnabled() bool
	// deprecated
	SetMenuChangedMessagesEnabled(value bool)
	AllowsContextMenuPlugIns() bool
	SetAllowsContextMenuPlugIns(value bool)
	ShowsStateColumn() bool
	SetShowsStateColumn(value bool)
	HighlightedItem() MenuItem
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
	Delegate() MenuDelegateWrapper
	SetDelegate(value MenuDelegate)
}

type Menu struct {
	objc.Object
}

func MakeMenu(ptr unsafe.Pointer) Menu {
	return Menu{
		Object: objc.MakeObject(ptr),
	}
}

func (m_ Menu) InitWithTitle(title string) Menu {
	rv := ffi.CallMethod[Menu](m_, "initWithTitle:", title)
	rv.Autorelease()
	return rv
}

func (mc _MenuClass) Alloc() Menu {
	rv := ffi.CallMethod[Menu](mc, "alloc")
	return rv
}

func (m_ Menu) Init() Menu {
	rv := ffi.CallMethod[Menu](m_, "init")
	rv.Autorelease()
	return rv
}

func (mc _MenuClass) New() Menu {
	rv := ffi.CallMethod[Menu](mc, "new")
	rv.Autorelease()
	return rv
}

func NewMenu() Menu {
	return MenuClass.New()
}

func (mc _MenuClass) MenuBarVisible() bool {
	rv := ffi.CallMethod[bool](mc, "menuBarVisible")
	return rv
}

func (mc _MenuClass) SetMenuBarVisible(visible bool) {
	ffi.CallMethod[ffi.Void](mc, "setMenuBarVisible:", visible)
}

func (m_ Menu) InsertItem_AtIndex(newItem IMenuItem, index int) {
	ffi.CallMethod[ffi.Void](m_, "insertItem:atIndex:", newItem, index)
}

func (m_ Menu) InsertItemWithTitle_Action_KeyEquivalent_AtIndex(_string string, selector objc.Selector, charCode string, index int) MenuItem {
	rv := ffi.CallMethod[MenuItem](m_, "insertItemWithTitle:action:keyEquivalent:atIndex:", _string, selector, charCode, index)
	return rv
}

func (m_ Menu) AddItem(newItem IMenuItem) {
	ffi.CallMethod[ffi.Void](m_, "addItem:", newItem)
}

func (m_ Menu) AddItemWithTitle_Action_KeyEquivalent(_string string, selector objc.Selector, charCode string) MenuItem {
	rv := ffi.CallMethod[MenuItem](m_, "addItemWithTitle:action:keyEquivalent:", _string, selector, charCode)
	return rv
}

func (m_ Menu) RemoveItem(item IMenuItem) {
	ffi.CallMethod[ffi.Void](m_, "removeItem:", item)
}

func (m_ Menu) RemoveItemAtIndex(index int) {
	ffi.CallMethod[ffi.Void](m_, "removeItemAtIndex:", index)
}

func (m_ Menu) ItemChanged(item IMenuItem) {
	ffi.CallMethod[ffi.Void](m_, "itemChanged:", item)
}

func (m_ Menu) RemoveAllItems() {
	ffi.CallMethod[ffi.Void](m_, "removeAllItems")
}

func (m_ Menu) ItemWithTag(tag int) MenuItem {
	rv := ffi.CallMethod[MenuItem](m_, "itemWithTag:", tag)
	return rv
}

func (m_ Menu) ItemWithTitle(title string) MenuItem {
	rv := ffi.CallMethod[MenuItem](m_, "itemWithTitle:", title)
	return rv
}

func (m_ Menu) ItemAtIndex(index int) MenuItem {
	rv := ffi.CallMethod[MenuItem](m_, "itemAtIndex:", index)
	return rv
}

func (m_ Menu) IndexOfItem(item IMenuItem) int {
	rv := ffi.CallMethod[int](m_, "indexOfItem:", item)
	return rv
}

func (m_ Menu) IndexOfItemWithTitle(title string) int {
	rv := ffi.CallMethod[int](m_, "indexOfItemWithTitle:", title)
	return rv
}

func (m_ Menu) IndexOfItemWithTag(tag int) int {
	rv := ffi.CallMethod[int](m_, "indexOfItemWithTag:", tag)
	return rv
}

func (m_ Menu) IndexOfItemWithTarget_AndAction(target objc.IObject, actionSelector objc.Selector) int {
	rv := ffi.CallMethod[int](m_, "indexOfItemWithTarget:andAction:", target, actionSelector)
	return rv
}

func (m_ Menu) IndexOfItemWithRepresentedObject(object objc.IObject) int {
	rv := ffi.CallMethod[int](m_, "indexOfItemWithRepresentedObject:", object)
	return rv
}

func (m_ Menu) IndexOfItemWithSubmenu(submenu IMenu) int {
	rv := ffi.CallMethod[int](m_, "indexOfItemWithSubmenu:", submenu)
	return rv
}

func (m_ Menu) SetSubmenu_ForItem(menu IMenu, item IMenuItem) {
	ffi.CallMethod[ffi.Void](m_, "setSubmenu:forItem:", menu, item)
}

func (m_ Menu) SubmenuAction(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](m_, "submenuAction:", sender)
}

// deprecated
func (m_ Menu) AttachedMenu() Menu {
	rv := ffi.CallMethod[Menu](m_, "attachedMenu")
	return rv
}

// deprecated
func (m_ Menu) IsAttached() bool {
	rv := ffi.CallMethod[bool](m_, "isAttached")
	return rv
}

// deprecated
func (m_ Menu) LocationForSubmenu(submenu IMenu) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](m_, "locationForSubmenu:", submenu)
	return rv
}

func (m_ Menu) Update() {
	ffi.CallMethod[ffi.Void](m_, "update")
}

func (m_ Menu) PerformKeyEquivalent(event IEvent) bool {
	rv := ffi.CallMethod[bool](m_, "performKeyEquivalent:", event)
	return rv
}

func (m_ Menu) PerformActionForItemAtIndex(index int) {
	ffi.CallMethod[ffi.Void](m_, "performActionForItemAtIndex:", index)
}

// deprecated
func (m_ Menu) SizeToFit() {
	ffi.CallMethod[ffi.Void](m_, "sizeToFit")
}

func (mc _MenuClass) PopUpContextMenu_WithEvent_ForView(menu IMenu, event IEvent, view IView) {
	ffi.CallMethod[ffi.Void](mc, "popUpContextMenu:withEvent:forView:", menu, event, view)
}

func (mc _MenuClass) PopUpContextMenu_WithEvent_ForView_WithFont(menu IMenu, event IEvent, view IView, font IFont) {
	ffi.CallMethod[ffi.Void](mc, "popUpContextMenu:withEvent:forView:withFont:", menu, event, view, font)
}

// deprecated
func (m_ Menu) HelpRequested(eventPtr IEvent) {
	ffi.CallMethod[ffi.Void](m_, "helpRequested:", eventPtr)
}

func (m_ Menu) PopUpMenuPositioningItem_AtLocation_InView(item IMenuItem, location foundation.Point, view IView) bool {
	rv := ffi.CallMethod[bool](m_, "popUpMenuPositioningItem:atLocation:inView:", item, location, view)
	return rv
}

func (m_ Menu) CancelTracking() {
	ffi.CallMethod[ffi.Void](m_, "cancelTracking")
}

func (m_ Menu) CancelTrackingWithoutAnimation() {
	ffi.CallMethod[ffi.Void](m_, "cancelTrackingWithoutAnimation")
}

// deprecated
func (m_ Menu) ContextMenuRepresentation() objc.Object {
	rv := ffi.CallMethod[objc.Object](m_, "contextMenuRepresentation")
	return rv
}

// deprecated
func (m_ Menu) SetContextMenuRepresentation(menuRep objc.IObject) {
	ffi.CallMethod[ffi.Void](m_, "setContextMenuRepresentation:", menuRep)
}

// deprecated
func (m_ Menu) TearOffMenuRepresentation() objc.Object {
	rv := ffi.CallMethod[objc.Object](m_, "tearOffMenuRepresentation")
	return rv
}

// deprecated
func (m_ Menu) SetTearOffMenuRepresentation(menuRep objc.IObject) {
	ffi.CallMethod[ffi.Void](m_, "setTearOffMenuRepresentation:", menuRep)
}

// deprecated
func (m_ Menu) SetMenuRepresentation(menuRep objc.IObject) {
	ffi.CallMethod[ffi.Void](m_, "setMenuRepresentation:", menuRep)
}

// deprecated
func (m_ Menu) MenuRepresentation() objc.Object {
	rv := ffi.CallMethod[objc.Object](m_, "menuRepresentation")
	return rv
}

func (m_ Menu) MenuBarHeight() float64 {
	rv := ffi.CallMethod[float64](m_, "menuBarHeight")
	return rv
}

func (m_ Menu) NumberOfItems() int {
	rv := ffi.CallMethod[int](m_, "numberOfItems")
	return rv
}

func (m_ Menu) ItemArray() []MenuItem {
	rv := ffi.CallMethod[[]MenuItem](m_, "itemArray")
	return rv
}

func (m_ Menu) SetItemArray(value []IMenuItem) {
	ffi.CallMethod[ffi.Void](m_, "setItemArray:", value)
}

func (m_ Menu) Supermenu() Menu {
	rv := ffi.CallMethod[Menu](m_, "supermenu")
	return rv
}

func (m_ Menu) SetSupermenu(value IMenu) {
	ffi.CallMethod[ffi.Void](m_, "setSupermenu:", value)
}

// deprecated
func (m_ Menu) IsTornOff() bool {
	rv := ffi.CallMethod[bool](m_, "isTornOff")
	return rv
}

func (m_ Menu) AutoenablesItems() bool {
	rv := ffi.CallMethod[bool](m_, "autoenablesItems")
	return rv
}

func (m_ Menu) SetAutoenablesItems(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setAutoenablesItems:", value)
}

func (m_ Menu) Font() Font {
	rv := ffi.CallMethod[Font](m_, "font")
	return rv
}

func (m_ Menu) SetFont(value IFont) {
	ffi.CallMethod[ffi.Void](m_, "setFont:", value)
}

func (m_ Menu) Title() string {
	rv := ffi.CallMethod[string](m_, "title")
	return rv
}

func (m_ Menu) SetTitle(value string) {
	ffi.CallMethod[ffi.Void](m_, "setTitle:", value)
}

func (m_ Menu) MinimumWidth() float64 {
	rv := ffi.CallMethod[float64](m_, "minimumWidth")
	return rv
}

func (m_ Menu) SetMinimumWidth(value float64) {
	ffi.CallMethod[ffi.Void](m_, "setMinimumWidth:", value)
}

func (m_ Menu) Size() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](m_, "size")
	return rv
}

func (m_ Menu) PropertiesToUpdate() MenuProperties {
	rv := ffi.CallMethod[MenuProperties](m_, "propertiesToUpdate")
	return rv
}

// deprecated
func (m_ Menu) MenuChangedMessagesEnabled() bool {
	rv := ffi.CallMethod[bool](m_, "menuChangedMessagesEnabled")
	return rv
}

// deprecated
func (m_ Menu) SetMenuChangedMessagesEnabled(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setMenuChangedMessagesEnabled:", value)
}

func (m_ Menu) AllowsContextMenuPlugIns() bool {
	rv := ffi.CallMethod[bool](m_, "allowsContextMenuPlugIns")
	return rv
}

func (m_ Menu) SetAllowsContextMenuPlugIns(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setAllowsContextMenuPlugIns:", value)
}

func (m_ Menu) ShowsStateColumn() bool {
	rv := ffi.CallMethod[bool](m_, "showsStateColumn")
	return rv
}

func (m_ Menu) SetShowsStateColumn(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setShowsStateColumn:", value)
}

func (m_ Menu) HighlightedItem() MenuItem {
	rv := ffi.CallMethod[MenuItem](m_, "highlightedItem")
	return rv
}

func (m_ Menu) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := ffi.CallMethod[UserInterfaceLayoutDirection](m_, "userInterfaceLayoutDirection")
	return rv
}

func (m_ Menu) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	ffi.CallMethod[ffi.Void](m_, "setUserInterfaceLayoutDirection:", value)
}

func (m_ Menu) Delegate() MenuDelegateWrapper {
	rv := ffi.CallMethod[MenuDelegateWrapper](m_, "delegate")
	return rv
}

func (m_ Menu) SetDelegate(value MenuDelegate) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(m_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](m_, "setDelegate:", po)
}

type MenuDelegate interface {
	ImplementsMenu_UpdateItem_AtIndex_ShouldCancel() bool
	// optional
	Menu_UpdateItem_AtIndex_ShouldCancel(menu Menu, item MenuItem, index int, shouldCancel bool) bool
	ImplementsConfinementRectForMenu_OnScreen() bool
	// optional
	ConfinementRectForMenu_OnScreen(menu Menu, screen Screen) foundation.Rect
	ImplementsMenu_WillHighlightItem() bool
	// optional
	Menu_WillHighlightItem(menu Menu, item MenuItem)
	ImplementsMenuWillOpen() bool
	// optional
	MenuWillOpen(menu Menu)
	ImplementsMenuDidClose() bool
	// optional
	MenuDidClose(menu Menu)
	ImplementsNumberOfItemsInMenu() bool
	// optional
	NumberOfItemsInMenu(menu Menu) int
	ImplementsMenuNeedsUpdate() bool
	// optional
	MenuNeedsUpdate(menu Menu)
}

type MenuDelegateImpl struct {
	_Menu_UpdateItem_AtIndex_ShouldCancel func(menu Menu, item MenuItem, index int, shouldCancel bool) bool
	_ConfinementRectForMenu_OnScreen      func(menu Menu, screen Screen) foundation.Rect
	_Menu_WillHighlightItem               func(menu Menu, item MenuItem)
	_MenuWillOpen                         func(menu Menu)
	_MenuDidClose                         func(menu Menu)
	_NumberOfItemsInMenu                  func(menu Menu) int
	_MenuNeedsUpdate                      func(menu Menu)
}

func (di *MenuDelegateImpl) ImplementsMenu_UpdateItem_AtIndex_ShouldCancel() bool {
	return di._Menu_UpdateItem_AtIndex_ShouldCancel != nil
}

func (di *MenuDelegateImpl) SetMenu_UpdateItem_AtIndex_ShouldCancel(f func(menu Menu, item MenuItem, index int, shouldCancel bool) bool) {
	di._Menu_UpdateItem_AtIndex_ShouldCancel = f
}

func (di *MenuDelegateImpl) Menu_UpdateItem_AtIndex_ShouldCancel(menu Menu, item MenuItem, index int, shouldCancel bool) bool {
	return di._Menu_UpdateItem_AtIndex_ShouldCancel(menu, item, index, shouldCancel)
}
func (di *MenuDelegateImpl) ImplementsConfinementRectForMenu_OnScreen() bool {
	return di._ConfinementRectForMenu_OnScreen != nil
}

func (di *MenuDelegateImpl) SetConfinementRectForMenu_OnScreen(f func(menu Menu, screen Screen) foundation.Rect) {
	di._ConfinementRectForMenu_OnScreen = f
}

func (di *MenuDelegateImpl) ConfinementRectForMenu_OnScreen(menu Menu, screen Screen) foundation.Rect {
	return di._ConfinementRectForMenu_OnScreen(menu, screen)
}
func (di *MenuDelegateImpl) ImplementsMenu_WillHighlightItem() bool {
	return di._Menu_WillHighlightItem != nil
}

func (di *MenuDelegateImpl) SetMenu_WillHighlightItem(f func(menu Menu, item MenuItem)) {
	di._Menu_WillHighlightItem = f
}

func (di *MenuDelegateImpl) Menu_WillHighlightItem(menu Menu, item MenuItem) {
	di._Menu_WillHighlightItem(menu, item)
}
func (di *MenuDelegateImpl) ImplementsMenuWillOpen() bool {
	return di._MenuWillOpen != nil
}

func (di *MenuDelegateImpl) SetMenuWillOpen(f func(menu Menu)) {
	di._MenuWillOpen = f
}

func (di *MenuDelegateImpl) MenuWillOpen(menu Menu) {
	di._MenuWillOpen(menu)
}
func (di *MenuDelegateImpl) ImplementsMenuDidClose() bool {
	return di._MenuDidClose != nil
}

func (di *MenuDelegateImpl) SetMenuDidClose(f func(menu Menu)) {
	di._MenuDidClose = f
}

func (di *MenuDelegateImpl) MenuDidClose(menu Menu) {
	di._MenuDidClose(menu)
}
func (di *MenuDelegateImpl) ImplementsNumberOfItemsInMenu() bool {
	return di._NumberOfItemsInMenu != nil
}

func (di *MenuDelegateImpl) SetNumberOfItemsInMenu(f func(menu Menu) int) {
	di._NumberOfItemsInMenu = f
}

func (di *MenuDelegateImpl) NumberOfItemsInMenu(menu Menu) int {
	return di._NumberOfItemsInMenu(menu)
}
func (di *MenuDelegateImpl) ImplementsMenuNeedsUpdate() bool {
	return di._MenuNeedsUpdate != nil
}

func (di *MenuDelegateImpl) SetMenuNeedsUpdate(f func(menu Menu)) {
	di._MenuNeedsUpdate = f
}

func (di *MenuDelegateImpl) MenuNeedsUpdate(menu Menu) {
	di._MenuNeedsUpdate(menu)
}

type MenuDelegateWrapper struct {
	objc.Object
}

func (m_ *MenuDelegateWrapper) ImplementsMenu_UpdateItem_AtIndex_ShouldCancel() bool {
	return m_.RespondsToSelector(objc.GetSelector("menu:updateItem:atIndex:shouldCancel:"))
}

func (m_ MenuDelegateWrapper) Menu_UpdateItem_AtIndex_ShouldCancel(menu IMenu, item IMenuItem, index int, shouldCancel bool) bool {
	rv := ffi.CallMethod[bool](m_, "menu:updateItem:atIndex:shouldCancel:", menu, item, index, shouldCancel)
	return rv
}

func (m_ *MenuDelegateWrapper) ImplementsConfinementRectForMenu_OnScreen() bool {
	return m_.RespondsToSelector(objc.GetSelector("confinementRectForMenu:onScreen:"))
}

func (m_ MenuDelegateWrapper) ConfinementRectForMenu_OnScreen(menu IMenu, screen IScreen) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](m_, "confinementRectForMenu:onScreen:", menu, screen)
	return rv
}

func (m_ *MenuDelegateWrapper) ImplementsMenu_WillHighlightItem() bool {
	return m_.RespondsToSelector(objc.GetSelector("menu:willHighlightItem:"))
}

func (m_ MenuDelegateWrapper) Menu_WillHighlightItem(menu IMenu, item IMenuItem) {
	ffi.CallMethod[ffi.Void](m_, "menu:willHighlightItem:", menu, item)
}

func (m_ *MenuDelegateWrapper) ImplementsMenuWillOpen() bool {
	return m_.RespondsToSelector(objc.GetSelector("menuWillOpen:"))
}

func (m_ MenuDelegateWrapper) MenuWillOpen(menu IMenu) {
	ffi.CallMethod[ffi.Void](m_, "menuWillOpen:", menu)
}

func (m_ *MenuDelegateWrapper) ImplementsMenuDidClose() bool {
	return m_.RespondsToSelector(objc.GetSelector("menuDidClose:"))
}

func (m_ MenuDelegateWrapper) MenuDidClose(menu IMenu) {
	ffi.CallMethod[ffi.Void](m_, "menuDidClose:", menu)
}

func (m_ *MenuDelegateWrapper) ImplementsNumberOfItemsInMenu() bool {
	return m_.RespondsToSelector(objc.GetSelector("numberOfItemsInMenu:"))
}

func (m_ MenuDelegateWrapper) NumberOfItemsInMenu(menu IMenu) int {
	rv := ffi.CallMethod[int](m_, "numberOfItemsInMenu:", menu)
	return rv
}

func (m_ *MenuDelegateWrapper) ImplementsMenuNeedsUpdate() bool {
	return m_.RespondsToSelector(objc.GetSelector("menuNeedsUpdate:"))
}

func (m_ MenuDelegateWrapper) MenuNeedsUpdate(menu IMenu) {
	ffi.CallMethod[ffi.Void](m_, "menuNeedsUpdate:", menu)
}

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

func (m_ MenuItem) InitWithTitle_Action_KeyEquivalent(_string string, selector objc.Selector, charCode string) MenuItem {
	rv := ffi.CallMethod[MenuItem](m_, "initWithTitle:action:keyEquivalent:", _string, selector, charCode)
	rv.Autorelease()
	return rv
}

func (mc _MenuItemClass) Alloc() MenuItem {
	rv := ffi.CallMethod[MenuItem](mc, "alloc")
	return rv
}

func (m_ MenuItem) Init() MenuItem {
	rv := ffi.CallMethod[MenuItem](m_, "init")
	rv.Autorelease()
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
