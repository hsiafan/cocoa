// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
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
	DataSource() ComboBoxDataSourceWrapper
	SetDataSource(value ComboBoxDataSource)
	SetDataSource0(value objc.IObject)
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
	rv := objc.CallMethod[ComboBox](cc, "labelWithAttributedString:", attributedStringValue)
	return rv
}

func (cc _ComboBoxClass) LabelWithString(stringValue string) ComboBox {
	rv := objc.CallMethod[ComboBox](cc, "labelWithString:", stringValue)
	return rv
}

func (cc _ComboBoxClass) TextFieldWithString(stringValue string) ComboBox {
	rv := objc.CallMethod[ComboBox](cc, "textFieldWithString:", stringValue)
	return rv
}

func (cc _ComboBoxClass) WrappingLabelWithString(stringValue string) ComboBox {
	rv := objc.CallMethod[ComboBox](cc, "wrappingLabelWithString:", stringValue)
	return rv
}

func (c_ ComboBox) InitWithFrame(frameRect foundation.Rect) ComboBox {
	rv := objc.CallMethod[ComboBox](c_, "initWithFrame:", frameRect)
	return rv
}

func (c_ ComboBox) Init() ComboBox {
	rv := objc.CallMethod[ComboBox](c_, "init")
	return rv
}

func (cc _ComboBoxClass) Alloc() ComboBox {
	rv := objc.CallMethod[ComboBox](cc, "alloc")
	return rv
}

func (cc _ComboBoxClass) New() ComboBox {
	rv := objc.CallMethod[ComboBox](cc, "new")
	rv.Autorelease()
	return rv
}

func NewComboBox() ComboBox {
	return ComboBoxClass.New()
}

func (c_ ComboBox) AddItemsWithObjectValues(objects []objc.IObject) {
	objc.CallMethod[objc.Void](c_, "addItemsWithObjectValues:", objects)
}

func (c_ ComboBox) AddItemWithObjectValue(object objc.IObject) {
	objc.CallMethod[objc.Void](c_, "addItemWithObjectValue:", object)
}

func (c_ ComboBox) InsertItemWithObjectValue_AtIndex(object objc.IObject, index int) {
	objc.CallMethod[objc.Void](c_, "insertItemWithObjectValue:atIndex:", object, index)
}

func (c_ ComboBox) RemoveAllItems() {
	objc.CallMethod[objc.Void](c_, "removeAllItems")
}

func (c_ ComboBox) RemoveItemAtIndex(index int) {
	objc.CallMethod[objc.Void](c_, "removeItemAtIndex:", index)
}

func (c_ ComboBox) RemoveItemWithObjectValue(object objc.IObject) {
	objc.CallMethod[objc.Void](c_, "removeItemWithObjectValue:", object)
}

func (c_ ComboBox) IndexOfItemWithObjectValue(object objc.IObject) int {
	rv := objc.CallMethod[int](c_, "indexOfItemWithObjectValue:", object)
	return rv
}

func (c_ ComboBox) ItemObjectValueAtIndex(index int) objc.Object {
	rv := objc.CallMethod[objc.Object](c_, "itemObjectValueAtIndex:", index)
	return rv
}

func (c_ ComboBox) NoteNumberOfItemsChanged() {
	objc.CallMethod[objc.Void](c_, "noteNumberOfItemsChanged")
}

func (c_ ComboBox) ReloadData() {
	objc.CallMethod[objc.Void](c_, "reloadData")
}

func (c_ ComboBox) ScrollItemAtIndexToTop(index int) {
	objc.CallMethod[objc.Void](c_, "scrollItemAtIndexToTop:", index)
}

func (c_ ComboBox) ScrollItemAtIndexToVisible(index int) {
	objc.CallMethod[objc.Void](c_, "scrollItemAtIndexToVisible:", index)
}

func (c_ ComboBox) DeselectItemAtIndex(index int) {
	objc.CallMethod[objc.Void](c_, "deselectItemAtIndex:", index)
}

func (c_ ComboBox) SelectItemAtIndex(index int) {
	objc.CallMethod[objc.Void](c_, "selectItemAtIndex:", index)
}

func (c_ ComboBox) SelectItemWithObjectValue(object objc.IObject) {
	objc.CallMethod[objc.Void](c_, "selectItemWithObjectValue:", object)
}

func (c_ ComboBox) HasVerticalScroller() bool {
	rv := objc.CallMethod[bool](c_, "hasVerticalScroller")
	return rv
}

func (c_ ComboBox) SetHasVerticalScroller(value bool) {
	objc.CallMethod[objc.Void](c_, "setHasVerticalScroller:", value)
}

func (c_ ComboBox) IntercellSpacing() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, "intercellSpacing")
	return rv
}

func (c_ ComboBox) SetIntercellSpacing(value foundation.Size) {
	objc.CallMethod[objc.Void](c_, "setIntercellSpacing:", value)
}

func (c_ ComboBox) IsButtonBordered() bool {
	rv := objc.CallMethod[bool](c_, "isButtonBordered")
	return rv
}

func (c_ ComboBox) SetButtonBordered(value bool) {
	objc.CallMethod[objc.Void](c_, "setButtonBordered:", value)
}

func (c_ ComboBox) ItemHeight() float64 {
	rv := objc.CallMethod[float64](c_, "itemHeight")
	return rv
}

func (c_ ComboBox) SetItemHeight(value float64) {
	objc.CallMethod[objc.Void](c_, "setItemHeight:", value)
}

func (c_ ComboBox) NumberOfVisibleItems() int {
	rv := objc.CallMethod[int](c_, "numberOfVisibleItems")
	return rv
}

func (c_ ComboBox) SetNumberOfVisibleItems(value int) {
	objc.CallMethod[objc.Void](c_, "setNumberOfVisibleItems:", value)
}

func (c_ ComboBox) DataSource() ComboBoxDataSourceWrapper {
	rv := objc.CallMethod[ComboBoxDataSourceWrapper](c_, "dataSource")
	return rv
}

func (c_ ComboBox) SetDataSource(value ComboBoxDataSource) {
	po := objc.CreateProtocol("NSComboBoxDataSource", value)
	defer po.Release()
	objc.SetAssociatedObject(c_, internal.AssociationKey("setDataSource"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](c_, "setDataSource:", po)
}

func (c_ ComboBox) SetDataSource0(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, "setDataSource:", value)
}

func (c_ ComboBox) UsesDataSource() bool {
	rv := objc.CallMethod[bool](c_, "usesDataSource")
	return rv
}

func (c_ ComboBox) SetUsesDataSource(value bool) {
	objc.CallMethod[objc.Void](c_, "setUsesDataSource:", value)
}

func (c_ ComboBox) ObjectValues() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](c_, "objectValues")
	return rv
}

func (c_ ComboBox) NumberOfItems() int {
	rv := objc.CallMethod[int](c_, "numberOfItems")
	return rv
}

func (c_ ComboBox) IndexOfSelectedItem() int {
	rv := objc.CallMethod[int](c_, "indexOfSelectedItem")
	return rv
}

func (c_ ComboBox) ObjectValueOfSelectedItem() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, "objectValueOfSelectedItem")
	return rv
}

func (c_ ComboBox) Completes() bool {
	rv := objc.CallMethod[bool](c_, "completes")
	return rv
}

func (c_ ComboBox) SetCompletes(value bool) {
	objc.CallMethod[objc.Void](c_, "setCompletes:", value)
}
