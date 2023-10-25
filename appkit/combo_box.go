// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ComboBoxClass = _ComboBoxClass{objc.GetClass("NSComboBox")}

type _ComboBoxClass struct {
	objc.Class
}

type IComboBox interface {
	ITextField
	AddItemsWithObjectValues(objects []objc.IObject)
	AddItemWithObjectValue(object objc.IObject)
	InsertItemWithObjectValue_AtIndex(object objc.IObject, index int)
	RemoveAllItems()
	RemoveItemAtIndex(index int)
	RemoveItemWithObjectValue(object objc.IObject)
	IndexOfItemWithObjectValue(object objc.IObject) int
	ItemObjectValueAtIndex(index int) objc.Object
	NoteNumberOfItemsChanged()
	ReloadData()
	ScrollItemAtIndexToTop(index int)
	ScrollItemAtIndexToVisible(index int)
	DeselectItemAtIndex(index int)
	SelectItemAtIndex(index int)
	SelectItemWithObjectValue(object objc.IObject)
	HasVerticalScroller() bool
	SetHasVerticalScroller(value bool)
	IntercellSpacing() foundation.Size
	SetIntercellSpacing(value foundation.Size)
	IsButtonBordered() bool
	SetButtonBordered(value bool)
	ItemHeight() float64
	SetItemHeight(value float64)
	NumberOfVisibleItems() int
	SetNumberOfVisibleItems(value int)
	DataSource() objc.Object
	SetDataSource(value objc.IObject)
	UsesDataSource() bool
	SetUsesDataSource(value bool)
	ObjectValues() []objc.Object
	NumberOfItems() int
	IndexOfSelectedItem() int
	ObjectValueOfSelectedItem() objc.Object
	Completes() bool
	SetCompletes(value bool)
}

type ComboBox struct {
	TextField
}

func MakeComboBox(ptr unsafe.Pointer) ComboBox {
	return ComboBox{
		TextField: MakeTextField(ptr),
	}
}

func (cc _ComboBoxClass) LabelWithAttributedString(attributedStringValue foundation.IAttributedString) ComboBox {
	rv := objc.CallMethod[ComboBox](cc, objc.SEL("labelWithAttributedString:"), objc.ExtractPtr(attributedStringValue))
	return rv
}

func (cc _ComboBoxClass) LabelWithString(stringValue string) ComboBox {
	rv := objc.CallMethod[ComboBox](cc, objc.SEL("labelWithString:"), stringValue)
	return rv
}

func (cc _ComboBoxClass) TextFieldWithString(stringValue string) ComboBox {
	rv := objc.CallMethod[ComboBox](cc, objc.SEL("textFieldWithString:"), stringValue)
	return rv
}

func (cc _ComboBoxClass) WrappingLabelWithString(stringValue string) ComboBox {
	rv := objc.CallMethod[ComboBox](cc, objc.SEL("wrappingLabelWithString:"), stringValue)
	return rv
}

func (c_ ComboBox) InitWithFrame(frameRect foundation.Rect) ComboBox {
	rv := objc.CallMethod[ComboBox](c_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (c_ ComboBox) Init() ComboBox {
	rv := objc.CallMethod[ComboBox](c_, objc.SEL("init"))
	return rv
}

func (cc _ComboBoxClass) Alloc() ComboBox {
	rv := objc.CallMethod[ComboBox](cc, objc.SEL("alloc"))
	return rv
}

func (cc _ComboBoxClass) New() ComboBox {
	rv := objc.CallMethod[ComboBox](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewComboBox() ComboBox {
	return ComboBoxClass.New()
}

func (c_ ComboBox) AddItemsWithObjectValues(objects []objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("addItemsWithObjectValues:"), objects)
}

func (c_ ComboBox) AddItemWithObjectValue(object objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("addItemWithObjectValue:"), objc.ExtractPtr(object))
}

func (c_ ComboBox) InsertItemWithObjectValue_AtIndex(object objc.IObject, index int) {
	objc.CallMethod[objc.Void](c_, objc.SEL("insertItemWithObjectValue:atIndex:"), objc.ExtractPtr(object), index)
}

func (c_ ComboBox) RemoveAllItems() {
	objc.CallMethod[objc.Void](c_, objc.SEL("removeAllItems"))
}

func (c_ ComboBox) RemoveItemAtIndex(index int) {
	objc.CallMethod[objc.Void](c_, objc.SEL("removeItemAtIndex:"), index)
}

func (c_ ComboBox) RemoveItemWithObjectValue(object objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("removeItemWithObjectValue:"), objc.ExtractPtr(object))
}

func (c_ ComboBox) IndexOfItemWithObjectValue(object objc.IObject) int {
	rv := objc.CallMethod[int](c_, objc.SEL("indexOfItemWithObjectValue:"), objc.ExtractPtr(object))
	return rv
}

func (c_ ComboBox) ItemObjectValueAtIndex(index int) objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.SEL("itemObjectValueAtIndex:"), index)
	return rv
}

func (c_ ComboBox) NoteNumberOfItemsChanged() {
	objc.CallMethod[objc.Void](c_, objc.SEL("noteNumberOfItemsChanged"))
}

func (c_ ComboBox) ReloadData() {
	objc.CallMethod[objc.Void](c_, objc.SEL("reloadData"))
}

func (c_ ComboBox) ScrollItemAtIndexToTop(index int) {
	objc.CallMethod[objc.Void](c_, objc.SEL("scrollItemAtIndexToTop:"), index)
}

func (c_ ComboBox) ScrollItemAtIndexToVisible(index int) {
	objc.CallMethod[objc.Void](c_, objc.SEL("scrollItemAtIndexToVisible:"), index)
}

func (c_ ComboBox) DeselectItemAtIndex(index int) {
	objc.CallMethod[objc.Void](c_, objc.SEL("deselectItemAtIndex:"), index)
}

func (c_ ComboBox) SelectItemAtIndex(index int) {
	objc.CallMethod[objc.Void](c_, objc.SEL("selectItemAtIndex:"), index)
}

func (c_ ComboBox) SelectItemWithObjectValue(object objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("selectItemWithObjectValue:"), objc.ExtractPtr(object))
}

func (c_ ComboBox) HasVerticalScroller() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("hasVerticalScroller"))
	return rv
}

func (c_ ComboBox) SetHasVerticalScroller(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setHasVerticalScroller:"), value)
}

func (c_ ComboBox) IntercellSpacing() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.SEL("intercellSpacing"))
	return rv
}

func (c_ ComboBox) SetIntercellSpacing(value foundation.Size) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setIntercellSpacing:"), value)
}

func (c_ ComboBox) IsButtonBordered() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isButtonBordered"))
	return rv
}

func (c_ ComboBox) SetButtonBordered(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setButtonBordered:"), value)
}

func (c_ ComboBox) ItemHeight() float64 {
	rv := objc.CallMethod[float64](c_, objc.SEL("itemHeight"))
	return rv
}

func (c_ ComboBox) SetItemHeight(value float64) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setItemHeight:"), value)
}

func (c_ ComboBox) NumberOfVisibleItems() int {
	rv := objc.CallMethod[int](c_, objc.SEL("numberOfVisibleItems"))
	return rv
}

func (c_ ComboBox) SetNumberOfVisibleItems(value int) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setNumberOfVisibleItems:"), value)
}

// weak property
func (c_ ComboBox) DataSource() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.SEL("dataSource"))
	return rv
}

// weak property
func (c_ ComboBox) SetDataSource(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setDataSource:"), objc.ExtractPtr(value))
}

func (c_ ComboBox) UsesDataSource() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("usesDataSource"))
	return rv
}

func (c_ ComboBox) SetUsesDataSource(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setUsesDataSource:"), value)
}

func (c_ ComboBox) ObjectValues() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](c_, objc.SEL("objectValues"))
	return rv
}

func (c_ ComboBox) NumberOfItems() int {
	rv := objc.CallMethod[int](c_, objc.SEL("numberOfItems"))
	return rv
}

func (c_ ComboBox) IndexOfSelectedItem() int {
	rv := objc.CallMethod[int](c_, objc.SEL("indexOfSelectedItem"))
	return rv
}

func (c_ ComboBox) ObjectValueOfSelectedItem() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.SEL("objectValueOfSelectedItem"))
	return rv
}

func (c_ ComboBox) Completes() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("completes"))
	return rv
}

func (c_ ComboBox) SetCompletes(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setCompletes:"), value)
}
