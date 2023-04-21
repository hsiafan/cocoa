// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

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
	rv := objc.CallMethod[bool](m_, "menu:updateItem:atIndex:shouldCancel:", menu, item, index, shouldCancel)
	return rv
}

func (m_ *MenuDelegateWrapper) ImplementsConfinementRectForMenu_OnScreen() bool {
	return m_.RespondsToSelector(objc.GetSelector("confinementRectForMenu:onScreen:"))
}

func (m_ MenuDelegateWrapper) ConfinementRectForMenu_OnScreen(menu IMenu, screen IScreen) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](m_, "confinementRectForMenu:onScreen:", menu, screen)
	return rv
}

func (m_ *MenuDelegateWrapper) ImplementsMenu_WillHighlightItem() bool {
	return m_.RespondsToSelector(objc.GetSelector("menu:willHighlightItem:"))
}

func (m_ MenuDelegateWrapper) Menu_WillHighlightItem(menu IMenu, item IMenuItem) {
	objc.CallMethod[objc.Void](m_, "menu:willHighlightItem:", menu, item)
}

func (m_ *MenuDelegateWrapper) ImplementsMenuWillOpen() bool {
	return m_.RespondsToSelector(objc.GetSelector("menuWillOpen:"))
}

func (m_ MenuDelegateWrapper) MenuWillOpen(menu IMenu) {
	objc.CallMethod[objc.Void](m_, "menuWillOpen:", menu)
}

func (m_ *MenuDelegateWrapper) ImplementsMenuDidClose() bool {
	return m_.RespondsToSelector(objc.GetSelector("menuDidClose:"))
}

func (m_ MenuDelegateWrapper) MenuDidClose(menu IMenu) {
	objc.CallMethod[objc.Void](m_, "menuDidClose:", menu)
}

func (m_ *MenuDelegateWrapper) ImplementsNumberOfItemsInMenu() bool {
	return m_.RespondsToSelector(objc.GetSelector("numberOfItemsInMenu:"))
}

func (m_ MenuDelegateWrapper) NumberOfItemsInMenu(menu IMenu) int {
	rv := objc.CallMethod[int](m_, "numberOfItemsInMenu:", menu)
	return rv
}

func (m_ *MenuDelegateWrapper) ImplementsMenuNeedsUpdate() bool {
	return m_.RespondsToSelector(objc.GetSelector("menuNeedsUpdate:"))
}

func (m_ MenuDelegateWrapper) MenuNeedsUpdate(menu IMenu) {
	objc.CallMethod[objc.Void](m_, "menuNeedsUpdate:", menu)
}
