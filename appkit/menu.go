// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	InsertItemWithTitle_Action_KeyEquivalent_AtIndex(string_ string, selector objc.Selector, charCode string, index int) MenuItem
	AddItem(newItem IMenuItem)
	AddItemWithTitle_Action_KeyEquivalent(string_ string, selector objc.Selector, charCode string) MenuItem
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
	SetDelegate0(value objc.IObject)
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
	rv := objc.CallMethod[Menu](m_, "initWithTitle:", title)
	return rv
}

func (mc _MenuClass) Alloc() Menu {
	rv := objc.CallMethod[Menu](mc, "alloc")
	return rv
}

func (mc _MenuClass) New() Menu {
	rv := objc.CallMethod[Menu](mc, "new")
	rv.Autorelease()
	return rv
}

func NewMenu() Menu {
	return MenuClass.New()
}

func (m_ Menu) Init() Menu {
	rv := objc.CallMethod[Menu](m_, "init")
	return rv
}

func (mc _MenuClass) MenuBarVisible() bool {
	rv := objc.CallMethod[bool](mc, "menuBarVisible")
	return rv
}

func (mc _MenuClass) SetMenuBarVisible(visible bool) {
	objc.CallMethod[objc.Void](mc, "setMenuBarVisible:", visible)
}

func (m_ Menu) InsertItem_AtIndex(newItem IMenuItem, index int) {
	objc.CallMethod[objc.Void](m_, "insertItem:atIndex:", newItem, index)
}

func (m_ Menu) InsertItemWithTitle_Action_KeyEquivalent_AtIndex(string_ string, selector objc.Selector, charCode string, index int) MenuItem {
	rv := objc.CallMethod[MenuItem](m_, "insertItemWithTitle:action:keyEquivalent:atIndex:", string_, selector, charCode, index)
	return rv
}

func (m_ Menu) AddItem(newItem IMenuItem) {
	objc.CallMethod[objc.Void](m_, "addItem:", newItem)
}

func (m_ Menu) AddItemWithTitle_Action_KeyEquivalent(string_ string, selector objc.Selector, charCode string) MenuItem {
	rv := objc.CallMethod[MenuItem](m_, "addItemWithTitle:action:keyEquivalent:", string_, selector, charCode)
	return rv
}

func (m_ Menu) RemoveItem(item IMenuItem) {
	objc.CallMethod[objc.Void](m_, "removeItem:", item)
}

func (m_ Menu) RemoveItemAtIndex(index int) {
	objc.CallMethod[objc.Void](m_, "removeItemAtIndex:", index)
}

func (m_ Menu) ItemChanged(item IMenuItem) {
	objc.CallMethod[objc.Void](m_, "itemChanged:", item)
}

func (m_ Menu) RemoveAllItems() {
	objc.CallMethod[objc.Void](m_, "removeAllItems")
}

func (m_ Menu) ItemWithTag(tag int) MenuItem {
	rv := objc.CallMethod[MenuItem](m_, "itemWithTag:", tag)
	return rv
}

func (m_ Menu) ItemWithTitle(title string) MenuItem {
	rv := objc.CallMethod[MenuItem](m_, "itemWithTitle:", title)
	return rv
}

func (m_ Menu) ItemAtIndex(index int) MenuItem {
	rv := objc.CallMethod[MenuItem](m_, "itemAtIndex:", index)
	return rv
}

func (m_ Menu) IndexOfItem(item IMenuItem) int {
	rv := objc.CallMethod[int](m_, "indexOfItem:", item)
	return rv
}

func (m_ Menu) IndexOfItemWithTitle(title string) int {
	rv := objc.CallMethod[int](m_, "indexOfItemWithTitle:", title)
	return rv
}

func (m_ Menu) IndexOfItemWithTag(tag int) int {
	rv := objc.CallMethod[int](m_, "indexOfItemWithTag:", tag)
	return rv
}

func (m_ Menu) IndexOfItemWithTarget_AndAction(target objc.IObject, actionSelector objc.Selector) int {
	rv := objc.CallMethod[int](m_, "indexOfItemWithTarget:andAction:", target, actionSelector)
	return rv
}

func (m_ Menu) IndexOfItemWithRepresentedObject(object objc.IObject) int {
	rv := objc.CallMethod[int](m_, "indexOfItemWithRepresentedObject:", object)
	return rv
}

func (m_ Menu) IndexOfItemWithSubmenu(submenu IMenu) int {
	rv := objc.CallMethod[int](m_, "indexOfItemWithSubmenu:", submenu)
	return rv
}

func (m_ Menu) SetSubmenu_ForItem(menu IMenu, item IMenuItem) {
	objc.CallMethod[objc.Void](m_, "setSubmenu:forItem:", menu, item)
}

func (m_ Menu) SubmenuAction(sender objc.IObject) {
	objc.CallMethod[objc.Void](m_, "submenuAction:", sender)
}

// deprecated
func (m_ Menu) AttachedMenu() Menu {
	rv := objc.CallMethod[Menu](m_, "attachedMenu")
	return rv
}

// deprecated
func (m_ Menu) IsAttached() bool {
	rv := objc.CallMethod[bool](m_, "isAttached")
	return rv
}

// deprecated
func (m_ Menu) LocationForSubmenu(submenu IMenu) foundation.Point {
	rv := objc.CallMethod[foundation.Point](m_, "locationForSubmenu:", submenu)
	return rv
}

func (m_ Menu) Update() {
	objc.CallMethod[objc.Void](m_, "update")
}

func (m_ Menu) PerformKeyEquivalent(event IEvent) bool {
	rv := objc.CallMethod[bool](m_, "performKeyEquivalent:", event)
	return rv
}

func (m_ Menu) PerformActionForItemAtIndex(index int) {
	objc.CallMethod[objc.Void](m_, "performActionForItemAtIndex:", index)
}

// deprecated
func (m_ Menu) SizeToFit() {
	objc.CallMethod[objc.Void](m_, "sizeToFit")
}

func (mc _MenuClass) PopUpContextMenu_WithEvent_ForView(menu IMenu, event IEvent, view IView) {
	objc.CallMethod[objc.Void](mc, "popUpContextMenu:withEvent:forView:", menu, event, view)
}

func (mc _MenuClass) PopUpContextMenu_WithEvent_ForView_WithFont(menu IMenu, event IEvent, view IView, font IFont) {
	objc.CallMethod[objc.Void](mc, "popUpContextMenu:withEvent:forView:withFont:", menu, event, view, font)
}

// deprecated
func (m_ Menu) HelpRequested(eventPtr IEvent) {
	objc.CallMethod[objc.Void](m_, "helpRequested:", eventPtr)
}

func (m_ Menu) PopUpMenuPositioningItem_AtLocation_InView(item IMenuItem, location foundation.Point, view IView) bool {
	rv := objc.CallMethod[bool](m_, "popUpMenuPositioningItem:atLocation:inView:", item, location, view)
	return rv
}

func (m_ Menu) CancelTracking() {
	objc.CallMethod[objc.Void](m_, "cancelTracking")
}

func (m_ Menu) CancelTrackingWithoutAnimation() {
	objc.CallMethod[objc.Void](m_, "cancelTrackingWithoutAnimation")
}

// deprecated
func (m_ Menu) ContextMenuRepresentation() objc.Object {
	rv := objc.CallMethod[objc.Object](m_, "contextMenuRepresentation")
	return rv
}

// deprecated
func (m_ Menu) SetContextMenuRepresentation(menuRep objc.IObject) {
	objc.CallMethod[objc.Void](m_, "setContextMenuRepresentation:", menuRep)
}

// deprecated
func (m_ Menu) TearOffMenuRepresentation() objc.Object {
	rv := objc.CallMethod[objc.Object](m_, "tearOffMenuRepresentation")
	return rv
}

// deprecated
func (m_ Menu) SetTearOffMenuRepresentation(menuRep objc.IObject) {
	objc.CallMethod[objc.Void](m_, "setTearOffMenuRepresentation:", menuRep)
}

// deprecated
func (m_ Menu) SetMenuRepresentation(menuRep objc.IObject) {
	objc.CallMethod[objc.Void](m_, "setMenuRepresentation:", menuRep)
}

// deprecated
func (m_ Menu) MenuRepresentation() objc.Object {
	rv := objc.CallMethod[objc.Object](m_, "menuRepresentation")
	return rv
}

func (m_ Menu) MenuBarHeight() float64 {
	rv := objc.CallMethod[float64](m_, "menuBarHeight")
	return rv
}

func (m_ Menu) NumberOfItems() int {
	rv := objc.CallMethod[int](m_, "numberOfItems")
	return rv
}

func (m_ Menu) ItemArray() []MenuItem {
	rv := objc.CallMethod[[]MenuItem](m_, "itemArray")
	return rv
}

func (m_ Menu) SetItemArray(value []IMenuItem) {
	objc.CallMethod[objc.Void](m_, "setItemArray:", value)
}

func (m_ Menu) Supermenu() Menu {
	rv := objc.CallMethod[Menu](m_, "supermenu")
	return rv
}

func (m_ Menu) SetSupermenu(value IMenu) {
	objc.CallMethod[objc.Void](m_, "setSupermenu:", value)
}

// deprecated
func (m_ Menu) IsTornOff() bool {
	rv := objc.CallMethod[bool](m_, "isTornOff")
	return rv
}

func (m_ Menu) AutoenablesItems() bool {
	rv := objc.CallMethod[bool](m_, "autoenablesItems")
	return rv
}

func (m_ Menu) SetAutoenablesItems(value bool) {
	objc.CallMethod[objc.Void](m_, "setAutoenablesItems:", value)
}

func (m_ Menu) Font() Font {
	rv := objc.CallMethod[Font](m_, "font")
	return rv
}

func (m_ Menu) SetFont(value IFont) {
	objc.CallMethod[objc.Void](m_, "setFont:", value)
}

func (m_ Menu) Title() string {
	rv := objc.CallMethod[string](m_, "title")
	return rv
}

func (m_ Menu) SetTitle(value string) {
	objc.CallMethod[objc.Void](m_, "setTitle:", value)
}

func (m_ Menu) MinimumWidth() float64 {
	rv := objc.CallMethod[float64](m_, "minimumWidth")
	return rv
}

func (m_ Menu) SetMinimumWidth(value float64) {
	objc.CallMethod[objc.Void](m_, "setMinimumWidth:", value)
}

func (m_ Menu) Size() foundation.Size {
	rv := objc.CallMethod[foundation.Size](m_, "size")
	return rv
}

func (m_ Menu) PropertiesToUpdate() MenuProperties {
	rv := objc.CallMethod[MenuProperties](m_, "propertiesToUpdate")
	return rv
}

// deprecated
func (m_ Menu) MenuChangedMessagesEnabled() bool {
	rv := objc.CallMethod[bool](m_, "menuChangedMessagesEnabled")
	return rv
}

// deprecated
func (m_ Menu) SetMenuChangedMessagesEnabled(value bool) {
	objc.CallMethod[objc.Void](m_, "setMenuChangedMessagesEnabled:", value)
}

func (m_ Menu) AllowsContextMenuPlugIns() bool {
	rv := objc.CallMethod[bool](m_, "allowsContextMenuPlugIns")
	return rv
}

func (m_ Menu) SetAllowsContextMenuPlugIns(value bool) {
	objc.CallMethod[objc.Void](m_, "setAllowsContextMenuPlugIns:", value)
}

func (m_ Menu) ShowsStateColumn() bool {
	rv := objc.CallMethod[bool](m_, "showsStateColumn")
	return rv
}

func (m_ Menu) SetShowsStateColumn(value bool) {
	objc.CallMethod[objc.Void](m_, "setShowsStateColumn:", value)
}

func (m_ Menu) HighlightedItem() MenuItem {
	rv := objc.CallMethod[MenuItem](m_, "highlightedItem")
	return rv
}

func (m_ Menu) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.CallMethod[UserInterfaceLayoutDirection](m_, "userInterfaceLayoutDirection")
	return rv
}

func (m_ Menu) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.CallMethod[objc.Void](m_, "setUserInterfaceLayoutDirection:", value)
}

func (m_ Menu) Delegate() MenuDelegateWrapper {
	rv := objc.CallMethod[MenuDelegateWrapper](m_, "delegate")
	return rv
}

func (m_ Menu) SetDelegate(value MenuDelegate) {
	po := objc.CreateProtocol("NSMenuDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(m_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](m_, "setDelegate:", po)
}

func (m_ Menu) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](m_, "setDelegate:", value)
}
