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

func WrapMenuDelegate(v MenuDelegate) objc.Object {
	return objc.WrapAsProtocol("NSMenuDelegate", v)
}

type MenuDelegateBase struct {
}

func (p *MenuDelegateBase) ImplementsMenu_UpdateItem_AtIndex_ShouldCancel() bool {
	return false
}

func (p *MenuDelegateBase) Menu_UpdateItem_AtIndex_ShouldCancel(menu Menu, item MenuItem, index int, shouldCancel bool) bool {
	panic("unimpemented protocol method")
}

func (p *MenuDelegateBase) ImplementsConfinementRectForMenu_OnScreen() bool {
	return false
}

func (p *MenuDelegateBase) ConfinementRectForMenu_OnScreen(menu Menu, screen Screen) foundation.Rect {
	panic("unimpemented protocol method")
}

func (p *MenuDelegateBase) ImplementsMenu_WillHighlightItem() bool {
	return false
}

func (p *MenuDelegateBase) Menu_WillHighlightItem(menu Menu, item MenuItem) {
	panic("unimpemented protocol method")
}

func (p *MenuDelegateBase) ImplementsMenuWillOpen() bool {
	return false
}

func (p *MenuDelegateBase) MenuWillOpen(menu Menu) {
	panic("unimpemented protocol method")
}

func (p *MenuDelegateBase) ImplementsMenuDidClose() bool {
	return false
}

func (p *MenuDelegateBase) MenuDidClose(menu Menu) {
	panic("unimpemented protocol method")
}

func (p *MenuDelegateBase) ImplementsNumberOfItemsInMenu() bool {
	return false
}

func (p *MenuDelegateBase) NumberOfItemsInMenu(menu Menu) int {
	panic("unimpemented protocol method")
}

func (p *MenuDelegateBase) ImplementsMenuNeedsUpdate() bool {
	return false
}

func (p *MenuDelegateBase) MenuNeedsUpdate(menu Menu) {
	panic("unimpemented protocol method")
}
