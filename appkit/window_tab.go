// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var WindowTabClass = _WindowTabClass{objc.GetClass("NSWindowTab")}

type _WindowTabClass struct {
	objc.Class
}

type IWindowTab interface {
	objc.IObject
	Title() string
	SetTitle(value string)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	ToolTip() string
	SetToolTip(value string)
	AccessoryView() View
	SetAccessoryView(value IView)
}

type WindowTab struct {
	objc.Object
}

func MakeWindowTab(ptr unsafe.Pointer) WindowTab {
	return WindowTab{
		Object: objc.MakeObject(ptr),
	}
}

func (wc _WindowTabClass) Alloc() WindowTab {
	rv := ffi.CallMethod[WindowTab](wc, "alloc")
	return rv
}

func (w_ WindowTab) Init() WindowTab {
	rv := ffi.CallMethod[WindowTab](w_, "init")
	return rv
}

func (wc _WindowTabClass) New() WindowTab {
	rv := ffi.CallMethod[WindowTab](wc, "new")
	rv.Autorelease()
	return rv
}

func NewWindowTab() WindowTab {
	return WindowTabClass.New()
}

func (w_ WindowTab) Title() string {
	rv := ffi.CallMethod[string](w_, "title")
	return rv
}

func (w_ WindowTab) SetTitle(value string) {
	ffi.CallMethod[ffi.Void](w_, "setTitle:", value)
}

func (w_ WindowTab) AttributedTitle() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](w_, "attributedTitle")
	return rv
}

func (w_ WindowTab) SetAttributedTitle(value foundation.IAttributedString) {
	ffi.CallMethod[ffi.Void](w_, "setAttributedTitle:", value)
}

func (w_ WindowTab) ToolTip() string {
	rv := ffi.CallMethod[string](w_, "toolTip")
	return rv
}

func (w_ WindowTab) SetToolTip(value string) {
	ffi.CallMethod[ffi.Void](w_, "setToolTip:", value)
}

func (w_ WindowTab) AccessoryView() View {
	rv := ffi.CallMethod[View](w_, "accessoryView")
	return rv
}

func (w_ WindowTab) SetAccessoryView(value IView) {
	ffi.CallMethod[ffi.Void](w_, "setAccessoryView:", value)
}

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
	rv := ffi.CallMethod[WindowTabGroup](wc, "alloc")
	return rv
}

func (w_ WindowTabGroup) Init() WindowTabGroup {
	rv := ffi.CallMethod[WindowTabGroup](w_, "init")
	return rv
}

func (wc _WindowTabGroupClass) New() WindowTabGroup {
	rv := ffi.CallMethod[WindowTabGroup](wc, "new")
	rv.Autorelease()
	return rv
}

func NewWindowTabGroup() WindowTabGroup {
	return WindowTabGroupClass.New()
}

func (w_ WindowTabGroup) AddWindow(window IWindow) {
	ffi.CallMethod[ffi.Void](w_, "addWindow:", window)
}

func (w_ WindowTabGroup) InsertWindow_AtIndex(window IWindow, index int) {
	ffi.CallMethod[ffi.Void](w_, "insertWindow:atIndex:", window, index)
}

func (w_ WindowTabGroup) RemoveWindow(window IWindow) {
	ffi.CallMethod[ffi.Void](w_, "removeWindow:", window)
}

func (w_ WindowTabGroup) Identifier() WindowTabbingIdentifier {
	rv := ffi.CallMethod[WindowTabbingIdentifier](w_, "identifier")
	return rv
}

func (w_ WindowTabGroup) IsOverviewVisible() bool {
	rv := ffi.CallMethod[bool](w_, "isOverviewVisible")
	return rv
}

func (w_ WindowTabGroup) SetOverviewVisible(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setOverviewVisible:", value)
}

func (w_ WindowTabGroup) IsTabBarVisible() bool {
	rv := ffi.CallMethod[bool](w_, "isTabBarVisible")
	return rv
}

func (w_ WindowTabGroup) Windows() []Window {
	rv := ffi.CallMethod[[]Window](w_, "windows")
	return rv
}

func (w_ WindowTabGroup) SelectedWindow() Window {
	rv := ffi.CallMethod[Window](w_, "selectedWindow")
	return rv
}

func (w_ WindowTabGroup) SetSelectedWindow(value IWindow) {
	ffi.CallMethod[ffi.Void](w_, "setSelectedWindow:", value)
}
