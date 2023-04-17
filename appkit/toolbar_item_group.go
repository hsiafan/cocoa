// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

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
	return rv
}

func (tc _ToolbarItemGroupClass) Alloc() ToolbarItemGroup {
	rv := ffi.CallMethod[ToolbarItemGroup](tc, "alloc")
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

func (t_ ToolbarItemGroup) Init() ToolbarItemGroup {
	rv := ffi.CallMethod[ToolbarItemGroup](t_, "init")
	return rv
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
