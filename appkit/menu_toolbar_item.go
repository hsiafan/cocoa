// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[MenuToolbarItem](m_, objc.SEL("initWithItemIdentifier:"), itemIdentifier)
	return rv
}

func (mc _MenuToolbarItemClass) Alloc() MenuToolbarItem {
	rv := objc.CallMethod[MenuToolbarItem](mc, objc.SEL("alloc"))
	return rv
}

func (mc _MenuToolbarItemClass) New() MenuToolbarItem {
	rv := objc.CallMethod[MenuToolbarItem](mc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewMenuToolbarItem() MenuToolbarItem {
	return MenuToolbarItemClass.New()
}

func (m_ MenuToolbarItem) Init() MenuToolbarItem {
	rv := objc.CallMethod[MenuToolbarItem](m_, objc.SEL("init"))
	return rv
}

func (m_ MenuToolbarItem) ShowsIndicator() bool {
	rv := objc.CallMethod[bool](m_, objc.SEL("showsIndicator"))
	return rv
}

func (m_ MenuToolbarItem) SetShowsIndicator(value bool) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setShowsIndicator:"), value)
}

func (m_ MenuToolbarItem) Menu() Menu {
	rv := objc.CallMethod[Menu](m_, objc.SEL("menu"))
	return rv
}

func (m_ MenuToolbarItem) SetMenu(value IMenu) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setMenu:"), objc.ExtractPtr(value))
}
