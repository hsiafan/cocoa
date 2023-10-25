// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[ToolbarItemGroup](tc, objc.SEL("groupWithItemIdentifier:images:selectionMode:labels:target:action:"), itemIdentifier, images, selectionMode, labels, objc.ExtractPtr(target), action)
	return rv
}

func (tc _ToolbarItemGroupClass) GroupWithItemIdentifier_Titles_SelectionMode_Labels_Target_Action(itemIdentifier ToolbarItemIdentifier, titles []string, selectionMode ToolbarItemGroupSelectionMode, labels []string, target objc.IObject, action objc.Selector) ToolbarItemGroup {
	rv := objc.CallMethod[ToolbarItemGroup](tc, objc.SEL("groupWithItemIdentifier:titles:selectionMode:labels:target:action:"), itemIdentifier, titles, selectionMode, labels, objc.ExtractPtr(target), action)
	return rv
}

func (t_ ToolbarItemGroup) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) ToolbarItemGroup {
	rv := objc.CallMethod[ToolbarItemGroup](t_, objc.SEL("initWithItemIdentifier:"), itemIdentifier)
	return rv
}

func (tc _ToolbarItemGroupClass) Alloc() ToolbarItemGroup {
	rv := objc.CallMethod[ToolbarItemGroup](tc, objc.SEL("alloc"))
	return rv
}

func (tc _ToolbarItemGroupClass) New() ToolbarItemGroup {
	rv := objc.CallMethod[ToolbarItemGroup](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewToolbarItemGroup() ToolbarItemGroup {
	return ToolbarItemGroupClass.New()
}

func (t_ ToolbarItemGroup) Init() ToolbarItemGroup {
	rv := objc.CallMethod[ToolbarItemGroup](t_, objc.SEL("init"))
	return rv
}

func (t_ ToolbarItemGroup) IsSelectedAtIndex(index int) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isSelectedAtIndex:"), index)
	return rv
}

func (t_ ToolbarItemGroup) SetSelected_AtIndex(selected bool, index int) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSelected:atIndex:"), selected, index)
}

func (t_ ToolbarItemGroup) Subitems() []ToolbarItem {
	rv := objc.CallMethod[[]ToolbarItem](t_, objc.SEL("subitems"))
	return rv
}

func (t_ ToolbarItemGroup) SetSubitems(value []IToolbarItem) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSubitems:"), value)
}

func (t_ ToolbarItemGroup) SelectedIndex() int {
	rv := objc.CallMethod[int](t_, objc.SEL("selectedIndex"))
	return rv
}

func (t_ ToolbarItemGroup) SetSelectedIndex(value int) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSelectedIndex:"), value)
}

func (t_ ToolbarItemGroup) ControlRepresentation() ToolbarItemGroupControlRepresentation {
	rv := objc.CallMethod[ToolbarItemGroupControlRepresentation](t_, objc.SEL("controlRepresentation"))
	return rv
}

func (t_ ToolbarItemGroup) SetControlRepresentation(value ToolbarItemGroupControlRepresentation) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setControlRepresentation:"), value)
}

func (t_ ToolbarItemGroup) SelectionMode() ToolbarItemGroupSelectionMode {
	rv := objc.CallMethod[ToolbarItemGroupSelectionMode](t_, objc.SEL("selectionMode"))
	return rv
}

func (t_ ToolbarItemGroup) SetSelectionMode(value ToolbarItemGroupSelectionMode) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSelectionMode:"), value)
}
