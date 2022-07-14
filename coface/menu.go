package coface

import (
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/objc"
)

// NewMenuItem create a new menu item, with selector
func NewMenuItem(title string, charCode string, selector objc.Selector) appkit.MenuItem {
	return appkit.MenuItemClass.Alloc().InitWithTitle_Action_KeyEquivalent(title, selector, charCode)
}

// NewMenuItemWithAction create a new menu item with action
func NewMenuItemWithAction(title string, charCode string, handler objc.ActionHandler) appkit.MenuItem {
	item := appkit.MenuItemClass.Alloc().InitWithTitle_Action_KeyEquivalent(title, objc.Selector{}, charCode)
	objc.SetAction(item, handler)
	return item
}

// NewSubMenuItem create a menu item that hold a sub menu
func NewSubMenuItem(menu appkit.IMenu) appkit.MenuItem {
	item := appkit.MenuItemClass.Alloc().InitWithTitle_Action_KeyEquivalent("", objc.Selector{}, "")
	item.SetSubmenu(menu)
	return item
}
