// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

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
	return rv
}

func (mc _MenuToolbarItemClass) Alloc() MenuToolbarItem {
	rv := ffi.CallMethod[MenuToolbarItem](mc, "alloc")
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

func (m_ MenuToolbarItem) Init() MenuToolbarItem {
	rv := ffi.CallMethod[MenuToolbarItem](m_, "init")
	return rv
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
