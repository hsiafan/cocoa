// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var PopUpButtonClass = _PopUpButtonClass{objc.GetClass("NSPopUpButton")}

type _PopUpButtonClass struct {
	objc.Class
}

type IPopUpButton interface {
	IButton
	AddItemWithTitle(title string)
	AddItemsWithTitles(itemTitles []string)
	InsertItemWithTitle_AtIndex(title string, index int)
	RemoveAllItems()
	RemoveItemWithTitle(title string)
	RemoveItemAtIndex(index int)
	SelectItem(item IMenuItem)
	SelectItemAtIndex(index int)
	SelectItemWithTag(tag int) bool
	SelectItemWithTitle(title string)
	ItemAtIndex(index int) MenuItem
	ItemTitleAtIndex(index int) string
	ItemWithTitle(title string) MenuItem
	IndexOfItem(item IMenuItem) int
	IndexOfItemWithTag(tag int) int
	IndexOfItemWithTitle(title string) int
	IndexOfItemWithRepresentedObject(obj objc.IObject) int
	IndexOfItemWithTarget_AndAction(target objc.IObject, actionSelector objc.Selector) int
	SynchronizeTitleAndSelectedItem()
	PullsDown() bool
	SetPullsDown(value bool)
	AutoenablesItems() bool
	SetAutoenablesItems(value bool)
	SelectedItem() MenuItem
	TitleOfSelectedItem() string
	IndexOfSelectedItem() int
	NumberOfItems() int
	ItemArray() []MenuItem
	ItemTitles() []string
	LastItem() MenuItem
	PreferredEdge() foundation.RectEdge
	SetPreferredEdge(value foundation.RectEdge)
}

type PopUpButton struct {
	Button
}

func MakePopUpButton(ptr unsafe.Pointer) PopUpButton {
	return PopUpButton{
		Button: MakeButton(ptr),
	}
}

func (p_ PopUpButton) InitWithFrame_PullsDown(buttonFrame foundation.Rect, flag bool) PopUpButton {
	rv := ffi.CallMethod[PopUpButton](p_, "initWithFrame:pullsDown:", buttonFrame, flag)
	return rv
}

func (pc _PopUpButtonClass) CheckboxWithTitle_Target_Action(title string, target objc.IObject, action objc.Selector) PopUpButton {
	rv := ffi.CallMethod[PopUpButton](pc, "checkboxWithTitle:target:action:", title, target, action)
	return rv
}

func (pc _PopUpButtonClass) ButtonWithImage_Target_Action(image IImage, target objc.IObject, action objc.Selector) PopUpButton {
	rv := ffi.CallMethod[PopUpButton](pc, "buttonWithImage:target:action:", image, target, action)
	return rv
}

func (pc _PopUpButtonClass) RadioButtonWithTitle_Target_Action(title string, target objc.IObject, action objc.Selector) PopUpButton {
	rv := ffi.CallMethod[PopUpButton](pc, "radioButtonWithTitle:target:action:", title, target, action)
	return rv
}

func (pc _PopUpButtonClass) ButtonWithTitle_Image_Target_Action(title string, image IImage, target objc.IObject, action objc.Selector) PopUpButton {
	rv := ffi.CallMethod[PopUpButton](pc, "buttonWithTitle:image:target:action:", title, image, target, action)
	return rv
}

func (pc _PopUpButtonClass) ButtonWithTitle_Target_Action(title string, target objc.IObject, action objc.Selector) PopUpButton {
	rv := ffi.CallMethod[PopUpButton](pc, "buttonWithTitle:target:action:", title, target, action)
	return rv
}

func (p_ PopUpButton) InitWithFrame(frameRect foundation.Rect) PopUpButton {
	rv := ffi.CallMethod[PopUpButton](p_, "initWithFrame:", frameRect)
	return rv
}

func (p_ PopUpButton) Init() PopUpButton {
	rv := ffi.CallMethod[PopUpButton](p_, "init")
	return rv
}

func (pc _PopUpButtonClass) Alloc() PopUpButton {
	rv := ffi.CallMethod[PopUpButton](pc, "alloc")
	return rv
}

func (pc _PopUpButtonClass) New() PopUpButton {
	rv := ffi.CallMethod[PopUpButton](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPopUpButton() PopUpButton {
	return PopUpButtonClass.New()
}

func (p_ PopUpButton) AddItemWithTitle(title string) {
	ffi.CallMethod[ffi.Void](p_, "addItemWithTitle:", title)
}

func (p_ PopUpButton) AddItemsWithTitles(itemTitles []string) {
	ffi.CallMethod[ffi.Void](p_, "addItemsWithTitles:", itemTitles)
}

func (p_ PopUpButton) InsertItemWithTitle_AtIndex(title string, index int) {
	ffi.CallMethod[ffi.Void](p_, "insertItemWithTitle:atIndex:", title, index)
}

func (p_ PopUpButton) RemoveAllItems() {
	ffi.CallMethod[ffi.Void](p_, "removeAllItems")
}

func (p_ PopUpButton) RemoveItemWithTitle(title string) {
	ffi.CallMethod[ffi.Void](p_, "removeItemWithTitle:", title)
}

func (p_ PopUpButton) RemoveItemAtIndex(index int) {
	ffi.CallMethod[ffi.Void](p_, "removeItemAtIndex:", index)
}

func (p_ PopUpButton) SelectItem(item IMenuItem) {
	ffi.CallMethod[ffi.Void](p_, "selectItem:", item)
}

func (p_ PopUpButton) SelectItemAtIndex(index int) {
	ffi.CallMethod[ffi.Void](p_, "selectItemAtIndex:", index)
}

func (p_ PopUpButton) SelectItemWithTag(tag int) bool {
	rv := ffi.CallMethod[bool](p_, "selectItemWithTag:", tag)
	return rv
}

func (p_ PopUpButton) SelectItemWithTitle(title string) {
	ffi.CallMethod[ffi.Void](p_, "selectItemWithTitle:", title)
}

func (p_ PopUpButton) ItemAtIndex(index int) MenuItem {
	rv := ffi.CallMethod[MenuItem](p_, "itemAtIndex:", index)
	return rv
}

func (p_ PopUpButton) ItemTitleAtIndex(index int) string {
	rv := ffi.CallMethod[string](p_, "itemTitleAtIndex:", index)
	return rv
}

func (p_ PopUpButton) ItemWithTitle(title string) MenuItem {
	rv := ffi.CallMethod[MenuItem](p_, "itemWithTitle:", title)
	return rv
}

func (p_ PopUpButton) IndexOfItem(item IMenuItem) int {
	rv := ffi.CallMethod[int](p_, "indexOfItem:", item)
	return rv
}

func (p_ PopUpButton) IndexOfItemWithTag(tag int) int {
	rv := ffi.CallMethod[int](p_, "indexOfItemWithTag:", tag)
	return rv
}

func (p_ PopUpButton) IndexOfItemWithTitle(title string) int {
	rv := ffi.CallMethod[int](p_, "indexOfItemWithTitle:", title)
	return rv
}

func (p_ PopUpButton) IndexOfItemWithRepresentedObject(obj objc.IObject) int {
	rv := ffi.CallMethod[int](p_, "indexOfItemWithRepresentedObject:", obj)
	return rv
}

func (p_ PopUpButton) IndexOfItemWithTarget_AndAction(target objc.IObject, actionSelector objc.Selector) int {
	rv := ffi.CallMethod[int](p_, "indexOfItemWithTarget:andAction:", target, actionSelector)
	return rv
}

func (p_ PopUpButton) SynchronizeTitleAndSelectedItem() {
	ffi.CallMethod[ffi.Void](p_, "synchronizeTitleAndSelectedItem")
}

func (p_ PopUpButton) PullsDown() bool {
	rv := ffi.CallMethod[bool](p_, "pullsDown")
	return rv
}

func (p_ PopUpButton) SetPullsDown(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setPullsDown:", value)
}

func (p_ PopUpButton) AutoenablesItems() bool {
	rv := ffi.CallMethod[bool](p_, "autoenablesItems")
	return rv
}

func (p_ PopUpButton) SetAutoenablesItems(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setAutoenablesItems:", value)
}

func (p_ PopUpButton) SelectedItem() MenuItem {
	rv := ffi.CallMethod[MenuItem](p_, "selectedItem")
	return rv
}

func (p_ PopUpButton) TitleOfSelectedItem() string {
	rv := ffi.CallMethod[string](p_, "titleOfSelectedItem")
	return rv
}

func (p_ PopUpButton) IndexOfSelectedItem() int {
	rv := ffi.CallMethod[int](p_, "indexOfSelectedItem")
	return rv
}

func (p_ PopUpButton) NumberOfItems() int {
	rv := ffi.CallMethod[int](p_, "numberOfItems")
	return rv
}

func (p_ PopUpButton) ItemArray() []MenuItem {
	rv := ffi.CallMethod[[]MenuItem](p_, "itemArray")
	return rv
}

func (p_ PopUpButton) ItemTitles() []string {
	rv := ffi.CallMethod[[]string](p_, "itemTitles")
	return rv
}

func (p_ PopUpButton) LastItem() MenuItem {
	rv := ffi.CallMethod[MenuItem](p_, "lastItem")
	return rv
}

func (p_ PopUpButton) PreferredEdge() foundation.RectEdge {
	rv := ffi.CallMethod[foundation.RectEdge](p_, "preferredEdge")
	return rv
}

func (p_ PopUpButton) SetPreferredEdge(value foundation.RectEdge) {
	ffi.CallMethod[ffi.Void](p_, "setPreferredEdge:", value)
}
