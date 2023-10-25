// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[PopUpButton](p_, objc.SEL("initWithFrame:pullsDown:"), buttonFrame, flag)
	return rv
}

func (pc _PopUpButtonClass) CheckboxWithTitle_Target_Action(title string, target objc.IObject, action objc.Selector) PopUpButton {
	rv := objc.CallMethod[PopUpButton](pc, objc.SEL("checkboxWithTitle:target:action:"), title, objc.ExtractPtr(target), action)
	return rv
}

func (pc _PopUpButtonClass) ButtonWithImage_Target_Action(image IImage, target objc.IObject, action objc.Selector) PopUpButton {
	rv := objc.CallMethod[PopUpButton](pc, objc.SEL("buttonWithImage:target:action:"), objc.ExtractPtr(image), objc.ExtractPtr(target), action)
	return rv
}

func (pc _PopUpButtonClass) RadioButtonWithTitle_Target_Action(title string, target objc.IObject, action objc.Selector) PopUpButton {
	rv := objc.CallMethod[PopUpButton](pc, objc.SEL("radioButtonWithTitle:target:action:"), title, objc.ExtractPtr(target), action)
	return rv
}

func (pc _PopUpButtonClass) ButtonWithTitle_Image_Target_Action(title string, image IImage, target objc.IObject, action objc.Selector) PopUpButton {
	rv := objc.CallMethod[PopUpButton](pc, objc.SEL("buttonWithTitle:image:target:action:"), title, objc.ExtractPtr(image), objc.ExtractPtr(target), action)
	return rv
}

func (pc _PopUpButtonClass) ButtonWithTitle_Target_Action(title string, target objc.IObject, action objc.Selector) PopUpButton {
	rv := objc.CallMethod[PopUpButton](pc, objc.SEL("buttonWithTitle:target:action:"), title, objc.ExtractPtr(target), action)
	return rv
}

func (p_ PopUpButton) InitWithFrame(frameRect foundation.Rect) PopUpButton {
	rv := objc.CallMethod[PopUpButton](p_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (p_ PopUpButton) Init() PopUpButton {
	rv := objc.CallMethod[PopUpButton](p_, objc.SEL("init"))
	return rv
}

func (pc _PopUpButtonClass) Alloc() PopUpButton {
	rv := objc.CallMethod[PopUpButton](pc, objc.SEL("alloc"))
	return rv
}

func (pc _PopUpButtonClass) New() PopUpButton {
	rv := objc.CallMethod[PopUpButton](pc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewPopUpButton() PopUpButton {
	return PopUpButtonClass.New()
}

func (p_ PopUpButton) AddItemWithTitle(title string) {
	objc.CallMethod[objc.Void](p_, objc.SEL("addItemWithTitle:"), title)
}

func (p_ PopUpButton) AddItemsWithTitles(itemTitles []string) {
	objc.CallMethod[objc.Void](p_, objc.SEL("addItemsWithTitles:"), itemTitles)
}

func (p_ PopUpButton) InsertItemWithTitle_AtIndex(title string, index int) {
	objc.CallMethod[objc.Void](p_, objc.SEL("insertItemWithTitle:atIndex:"), title, index)
}

func (p_ PopUpButton) RemoveAllItems() {
	objc.CallMethod[objc.Void](p_, objc.SEL("removeAllItems"))
}

func (p_ PopUpButton) RemoveItemWithTitle(title string) {
	objc.CallMethod[objc.Void](p_, objc.SEL("removeItemWithTitle:"), title)
}

func (p_ PopUpButton) RemoveItemAtIndex(index int) {
	objc.CallMethod[objc.Void](p_, objc.SEL("removeItemAtIndex:"), index)
}

func (p_ PopUpButton) SelectItem(item IMenuItem) {
	objc.CallMethod[objc.Void](p_, objc.SEL("selectItem:"), objc.ExtractPtr(item))
}

func (p_ PopUpButton) SelectItemAtIndex(index int) {
	objc.CallMethod[objc.Void](p_, objc.SEL("selectItemAtIndex:"), index)
}

func (p_ PopUpButton) SelectItemWithTag(tag int) bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("selectItemWithTag:"), tag)
	return rv
}

func (p_ PopUpButton) SelectItemWithTitle(title string) {
	objc.CallMethod[objc.Void](p_, objc.SEL("selectItemWithTitle:"), title)
}

func (p_ PopUpButton) ItemAtIndex(index int) MenuItem {
	rv := objc.CallMethod[MenuItem](p_, objc.SEL("itemAtIndex:"), index)
	return rv
}

func (p_ PopUpButton) ItemTitleAtIndex(index int) string {
	rv := objc.CallMethod[string](p_, objc.SEL("itemTitleAtIndex:"), index)
	return rv
}

func (p_ PopUpButton) ItemWithTitle(title string) MenuItem {
	rv := objc.CallMethod[MenuItem](p_, objc.SEL("itemWithTitle:"), title)
	return rv
}

func (p_ PopUpButton) IndexOfItem(item IMenuItem) int {
	rv := objc.CallMethod[int](p_, objc.SEL("indexOfItem:"), objc.ExtractPtr(item))
	return rv
}

func (p_ PopUpButton) IndexOfItemWithTag(tag int) int {
	rv := objc.CallMethod[int](p_, objc.SEL("indexOfItemWithTag:"), tag)
	return rv
}

func (p_ PopUpButton) IndexOfItemWithTitle(title string) int {
	rv := objc.CallMethod[int](p_, objc.SEL("indexOfItemWithTitle:"), title)
	return rv
}

func (p_ PopUpButton) IndexOfItemWithRepresentedObject(obj objc.IObject) int {
	rv := objc.CallMethod[int](p_, objc.SEL("indexOfItemWithRepresentedObject:"), objc.ExtractPtr(obj))
	return rv
}

func (p_ PopUpButton) IndexOfItemWithTarget_AndAction(target objc.IObject, actionSelector objc.Selector) int {
	rv := objc.CallMethod[int](p_, objc.SEL("indexOfItemWithTarget:andAction:"), objc.ExtractPtr(target), actionSelector)
	return rv
}

func (p_ PopUpButton) SynchronizeTitleAndSelectedItem() {
	objc.CallMethod[objc.Void](p_, objc.SEL("synchronizeTitleAndSelectedItem"))
}

func (p_ PopUpButton) PullsDown() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("pullsDown"))
	return rv
}

func (p_ PopUpButton) SetPullsDown(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setPullsDown:"), value)
}

func (p_ PopUpButton) AutoenablesItems() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("autoenablesItems"))
	return rv
}

func (p_ PopUpButton) SetAutoenablesItems(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setAutoenablesItems:"), value)
}

func (p_ PopUpButton) SelectedItem() MenuItem {
	rv := objc.CallMethod[MenuItem](p_, objc.SEL("selectedItem"))
	return rv
}

func (p_ PopUpButton) TitleOfSelectedItem() string {
	rv := objc.CallMethod[string](p_, objc.SEL("titleOfSelectedItem"))
	return rv
}

func (p_ PopUpButton) IndexOfSelectedItem() int {
	rv := objc.CallMethod[int](p_, objc.SEL("indexOfSelectedItem"))
	return rv
}

func (p_ PopUpButton) NumberOfItems() int {
	rv := objc.CallMethod[int](p_, objc.SEL("numberOfItems"))
	return rv
}

func (p_ PopUpButton) ItemArray() []MenuItem {
	rv := objc.CallMethod[[]MenuItem](p_, objc.SEL("itemArray"))
	return rv
}

func (p_ PopUpButton) ItemTitles() []string {
	rv := objc.CallMethod[[]string](p_, objc.SEL("itemTitles"))
	return rv
}

func (p_ PopUpButton) LastItem() MenuItem {
	rv := objc.CallMethod[MenuItem](p_, objc.SEL("lastItem"))
	return rv
}

func (p_ PopUpButton) PreferredEdge() foundation.RectEdge {
	rv := objc.CallMethod[foundation.RectEdge](p_, objc.SEL("preferredEdge"))
	return rv
}

func (p_ PopUpButton) SetPreferredEdge(value foundation.RectEdge) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setPreferredEdge:"), value)
}
