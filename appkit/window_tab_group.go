// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var WindowTabGroupClass = _WindowTabGroupClass{objc.GetClass("NSWindowTabGroup")}

type _WindowTabGroupClass struct {
	objc.Class
}

type IWindowTabGroup interface {
	objc.IObject
	AddWindow(window IWindow)
	InsertWindow_AtIndex(window IWindow, index int)
	RemoveWindow(window IWindow)
	Identifier() WindowTabbingIdentifier
	IsOverviewVisible() bool
	SetOverviewVisible(value bool)
	IsTabBarVisible() bool
	Windows() []Window
	SelectedWindow() Window
	SetSelectedWindow(value IWindow)
}

type WindowTabGroup struct {
	objc.Object
}

func MakeWindowTabGroup(ptr unsafe.Pointer) WindowTabGroup {
	return WindowTabGroup{
		Object: objc.MakeObject(ptr),
	}
}

func (wc _WindowTabGroupClass) Alloc() WindowTabGroup {
	rv := objc.CallMethod[WindowTabGroup](wc, "alloc")
	return rv
}

func (wc _WindowTabGroupClass) New() WindowTabGroup {
	rv := objc.CallMethod[WindowTabGroup](wc, "new")
	rv.Autorelease()
	return rv
}

func NewWindowTabGroup() WindowTabGroup {
	return WindowTabGroupClass.New()
}

func (w_ WindowTabGroup) Init() WindowTabGroup {
	rv := objc.CallMethod[WindowTabGroup](w_, "init")
	return rv
}

func (w_ WindowTabGroup) AddWindow(window IWindow) {
	objc.CallMethod[objc.Void](w_, "addWindow:", window)
}

func (w_ WindowTabGroup) InsertWindow_AtIndex(window IWindow, index int) {
	objc.CallMethod[objc.Void](w_, "insertWindow:atIndex:", window, index)
}

func (w_ WindowTabGroup) RemoveWindow(window IWindow) {
	objc.CallMethod[objc.Void](w_, "removeWindow:", window)
}

func (w_ WindowTabGroup) Identifier() WindowTabbingIdentifier {
	rv := objc.CallMethod[WindowTabbingIdentifier](w_, "identifier")
	return rv
}

func (w_ WindowTabGroup) IsOverviewVisible() bool {
	rv := objc.CallMethod[bool](w_, "isOverviewVisible")
	return rv
}

func (w_ WindowTabGroup) SetOverviewVisible(value bool) {
	objc.CallMethod[objc.Void](w_, "setOverviewVisible:", value)
}

func (w_ WindowTabGroup) IsTabBarVisible() bool {
	rv := objc.CallMethod[bool](w_, "isTabBarVisible")
	return rv
}

func (w_ WindowTabGroup) Windows() []Window {
	rv := objc.CallMethod[[]Window](w_, "windows")
	return rv
}

func (w_ WindowTabGroup) SelectedWindow() Window {
	rv := objc.CallMethod[Window](w_, "selectedWindow")
	return rv
}

func (w_ WindowTabGroup) SetSelectedWindow(value IWindow) {
	objc.CallMethod[objc.Void](w_, "setSelectedWindow:", value)
}
