// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[ComboBox](cc, "labelWithAttributedString:", attributedStringValue)
	return rv
}

func (cc _ComboBoxClass) LabelWithString(stringValue string) ComboBox {
	rv := ffi.CallMethod[ComboBox](cc, "labelWithString:", stringValue)
	return rv
}

func (cc _ComboBoxClass) TextFieldWithString(stringValue string) ComboBox {
	rv := ffi.CallMethod[ComboBox](cc, "textFieldWithString:", stringValue)
	return rv
}

func (cc _ComboBoxClass) WrappingLabelWithString(stringValue string) ComboBox {
	rv := ffi.CallMethod[ComboBox](cc, "wrappingLabelWithString:", stringValue)
	return rv
}

func (c_ ComboBox) InitWithFrame(frameRect foundation.Rect) ComboBox {
	rv := ffi.CallMethod[ComboBox](c_, "initWithFrame:", frameRect)
	rv.Autorelease()
	return rv
}

func (c_ ComboBox) Init() ComboBox {
	rv := ffi.CallMethod[ComboBox](c_, "init")
	rv.Autorelease()
	return rv
}

func (cc _ComboBoxClass) Alloc() ComboBox {
	rv := ffi.CallMethod[ComboBox](cc, "alloc")
	return rv
}

func (cc _ComboBoxClass) New() ComboBox {
	rv := ffi.CallMethod[ComboBox](cc, "new")
	rv.Autorelease()
	return rv
}

func NewComboBox() ComboBox {
	return ComboBoxClass.New()
}

func (c_ ComboBox) AddItemsWithObjectValues(objects []objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "addItemsWithObjectValues:", objects)
}

func (c_ ComboBox) AddItemWithObjectValue(object objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "addItemWithObjectValue:", object)
}

func (c_ ComboBox) InsertItemWithObjectValue_AtIndex(object objc.IObject, index int) {
	ffi.CallMethod[ffi.Void](c_, "insertItemWithObjectValue:atIndex:", object, index)
}

func (c_ ComboBox) RemoveAllItems() {
	ffi.CallMethod[ffi.Void](c_, "removeAllItems")
}

func (c_ ComboBox) RemoveItemAtIndex(index int) {
	ffi.CallMethod[ffi.Void](c_, "removeItemAtIndex:", index)
}

func (c_ ComboBox) RemoveItemWithObjectValue(object objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "removeItemWithObjectValue:", object)
}

func (c_ ComboBox) IndexOfItemWithObjectValue(object objc.IObject) int {
	rv := ffi.CallMethod[int](c_, "indexOfItemWithObjectValue:", object)
	return rv
}

func (c_ ComboBox) ItemObjectValueAtIndex(index int) objc.Object {
	rv := ffi.CallMethod[objc.Object](c_, "itemObjectValueAtIndex:", index)
	return rv
}

func (c_ ComboBox) NoteNumberOfItemsChanged() {
	ffi.CallMethod[ffi.Void](c_, "noteNumberOfItemsChanged")
}

func (c_ ComboBox) ReloadData() {
	ffi.CallMethod[ffi.Void](c_, "reloadData")
}

func (c_ ComboBox) ScrollItemAtIndexToTop(index int) {
	ffi.CallMethod[ffi.Void](c_, "scrollItemAtIndexToTop:", index)
}

func (c_ ComboBox) ScrollItemAtIndexToVisible(index int) {
	ffi.CallMethod[ffi.Void](c_, "scrollItemAtIndexToVisible:", index)
}

func (c_ ComboBox) DeselectItemAtIndex(index int) {
	ffi.CallMethod[ffi.Void](c_, "deselectItemAtIndex:", index)
}

func (c_ ComboBox) SelectItemAtIndex(index int) {
	ffi.CallMethod[ffi.Void](c_, "selectItemAtIndex:", index)
}

func (c_ ComboBox) SelectItemWithObjectValue(object objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "selectItemWithObjectValue:", object)
}

func (c_ ComboBox) HasVerticalScroller() bool {
	rv := ffi.CallMethod[bool](c_, "hasVerticalScroller")
	return rv
}

func (c_ ComboBox) SetHasVerticalScroller(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setHasVerticalScroller:", value)
}

func (c_ ComboBox) IntercellSpacing() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "intercellSpacing")
	return rv
}

func (c_ ComboBox) SetIntercellSpacing(value foundation.Size) {
	ffi.CallMethod[ffi.Void](c_, "setIntercellSpacing:", value)
}

func (c_ ComboBox) IsButtonBordered() bool {
	rv := ffi.CallMethod[bool](c_, "isButtonBordered")
	return rv
}

func (c_ ComboBox) SetButtonBordered(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setButtonBordered:", value)
}

func (c_ ComboBox) ItemHeight() float64 {
	rv := ffi.CallMethod[float64](c_, "itemHeight")
	return rv
}

func (c_ ComboBox) SetItemHeight(value float64) {
	ffi.CallMethod[ffi.Void](c_, "setItemHeight:", value)
}

func (c_ ComboBox) NumberOfVisibleItems() int {
	rv := ffi.CallMethod[int](c_, "numberOfVisibleItems")
	return rv
}

func (c_ ComboBox) SetNumberOfVisibleItems(value int) {
	ffi.CallMethod[ffi.Void](c_, "setNumberOfVisibleItems:", value)
}

func (c_ ComboBox) DataSource() ComboBoxDataSourceWrapper {
	rv := ffi.CallMethod[ComboBoxDataSourceWrapper](c_, "dataSource")
	return rv
}

func (c_ ComboBox) SetDataSource(value ComboBoxDataSource) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(c_, internal.AssociationKey("setDataSource"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](c_, "setDataSource:", po)
}

func (c_ ComboBox) UsesDataSource() bool {
	rv := ffi.CallMethod[bool](c_, "usesDataSource")
	return rv
}

func (c_ ComboBox) SetUsesDataSource(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setUsesDataSource:", value)
}

func (c_ ComboBox) ObjectValues() []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](c_, "objectValues")
	return rv
}

func (c_ ComboBox) NumberOfItems() int {
	rv := ffi.CallMethod[int](c_, "numberOfItems")
	return rv
}

func (c_ ComboBox) IndexOfSelectedItem() int {
	rv := ffi.CallMethod[int](c_, "indexOfSelectedItem")
	return rv
}

func (c_ ComboBox) ObjectValueOfSelectedItem() objc.Object {
	rv := ffi.CallMethod[objc.Object](c_, "objectValueOfSelectedItem")
	return rv
}

func (c_ ComboBox) Completes() bool {
	rv := ffi.CallMethod[bool](c_, "completes")
	return rv
}

func (c_ ComboBox) SetCompletes(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setCompletes:", value)
}

type ComboBoxDataSource interface {
	ImplementsComboBox_CompletedString() bool
	// optional
	ComboBox_CompletedString(comboBox ComboBox, _string string) string
	ImplementsComboBox_IndexOfItemWithStringValue() bool
	// optional
	ComboBox_IndexOfItemWithStringValue(comboBox ComboBox, _string string) uint
	ImplementsComboBox_ObjectValueForItemAtIndex() bool
	// optional
	ComboBox_ObjectValueForItemAtIndex(comboBox ComboBox, index int) objc.IObject
	ImplementsNumberOfItemsInComboBox() bool
	// optional
	NumberOfItemsInComboBox(comboBox ComboBox) int
}

type ComboBoxDataSourceWrapper struct {
	objc.Object
}

func (c_ *ComboBoxDataSourceWrapper) ImplementsComboBox_CompletedString() bool {
	return c_.RespondsToSelector(objc.GetSelector("comboBox:completedString:"))
}

func (c_ ComboBoxDataSourceWrapper) ComboBox_CompletedString(comboBox IComboBox, _string string) string {
	rv := ffi.CallMethod[string](c_, "comboBox:completedString:", comboBox, _string)
	return rv
}

func (c_ *ComboBoxDataSourceWrapper) ImplementsComboBox_IndexOfItemWithStringValue() bool {
	return c_.RespondsToSelector(objc.GetSelector("comboBox:indexOfItemWithStringValue:"))
}

func (c_ ComboBoxDataSourceWrapper) ComboBox_IndexOfItemWithStringValue(comboBox IComboBox, _string string) uint {
	rv := ffi.CallMethod[uint](c_, "comboBox:indexOfItemWithStringValue:", comboBox, _string)
	return rv
}

func (c_ *ComboBoxDataSourceWrapper) ImplementsComboBox_ObjectValueForItemAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("comboBox:objectValueForItemAtIndex:"))
}

func (c_ ComboBoxDataSourceWrapper) ComboBox_ObjectValueForItemAtIndex(comboBox IComboBox, index int) objc.Object {
	rv := ffi.CallMethod[objc.Object](c_, "comboBox:objectValueForItemAtIndex:", comboBox, index)
	return rv
}

func (c_ *ComboBoxDataSourceWrapper) ImplementsNumberOfItemsInComboBox() bool {
	return c_.RespondsToSelector(objc.GetSelector("numberOfItemsInComboBox:"))
}

func (c_ ComboBoxDataSourceWrapper) NumberOfItemsInComboBox(comboBox IComboBox) int {
	rv := ffi.CallMethod[int](c_, "numberOfItemsInComboBox:", comboBox)
	return rv
}

type ComboBoxDelegate interface {
	TextFieldDelegate
	ImplementsComboBoxSelectionDidChange() bool
	// optional
	ComboBoxSelectionDidChange(notification foundation.Notification)
	ImplementsComboBoxSelectionIsChanging() bool
	// optional
	ComboBoxSelectionIsChanging(notification foundation.Notification)
	ImplementsComboBoxWillDismiss() bool
	// optional
	ComboBoxWillDismiss(notification foundation.Notification)
	ImplementsComboBoxWillPopUp() bool
	// optional
	ComboBoxWillPopUp(notification foundation.Notification)
}

type ComboBoxDelegateImpl struct {
	TextFieldDelegateImpl
	_ComboBoxSelectionDidChange  func(notification foundation.Notification)
	_ComboBoxSelectionIsChanging func(notification foundation.Notification)
	_ComboBoxWillDismiss         func(notification foundation.Notification)
	_ComboBoxWillPopUp           func(notification foundation.Notification)
}

func (di *ComboBoxDelegateImpl) ImplementsComboBoxSelectionDidChange() bool {
	return di._ComboBoxSelectionDidChange != nil
}

func (di *ComboBoxDelegateImpl) SetComboBoxSelectionDidChange(f func(notification foundation.Notification)) {
	di._ComboBoxSelectionDidChange = f
}

func (di *ComboBoxDelegateImpl) ComboBoxSelectionDidChange(notification foundation.Notification) {
	di._ComboBoxSelectionDidChange(notification)
}
func (di *ComboBoxDelegateImpl) ImplementsComboBoxSelectionIsChanging() bool {
	return di._ComboBoxSelectionIsChanging != nil
}

func (di *ComboBoxDelegateImpl) SetComboBoxSelectionIsChanging(f func(notification foundation.Notification)) {
	di._ComboBoxSelectionIsChanging = f
}

func (di *ComboBoxDelegateImpl) ComboBoxSelectionIsChanging(notification foundation.Notification) {
	di._ComboBoxSelectionIsChanging(notification)
}
func (di *ComboBoxDelegateImpl) ImplementsComboBoxWillDismiss() bool {
	return di._ComboBoxWillDismiss != nil
}

func (di *ComboBoxDelegateImpl) SetComboBoxWillDismiss(f func(notification foundation.Notification)) {
	di._ComboBoxWillDismiss = f
}

func (di *ComboBoxDelegateImpl) ComboBoxWillDismiss(notification foundation.Notification) {
	di._ComboBoxWillDismiss(notification)
}
func (di *ComboBoxDelegateImpl) ImplementsComboBoxWillPopUp() bool {
	return di._ComboBoxWillPopUp != nil
}

func (di *ComboBoxDelegateImpl) SetComboBoxWillPopUp(f func(notification foundation.Notification)) {
	di._ComboBoxWillPopUp = f
}

func (di *ComboBoxDelegateImpl) ComboBoxWillPopUp(notification foundation.Notification) {
	di._ComboBoxWillPopUp(notification)
}

type ComboBoxDelegateWrapper struct {
	TextFieldDelegateWrapper
}

func (c_ *ComboBoxDelegateWrapper) ImplementsComboBoxSelectionDidChange() bool {
	return c_.RespondsToSelector(objc.GetSelector("comboBoxSelectionDidChange:"))
}

func (c_ ComboBoxDelegateWrapper) ComboBoxSelectionDidChange(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](c_, "comboBoxSelectionDidChange:", notification)
}

func (c_ *ComboBoxDelegateWrapper) ImplementsComboBoxSelectionIsChanging() bool {
	return c_.RespondsToSelector(objc.GetSelector("comboBoxSelectionIsChanging:"))
}

func (c_ ComboBoxDelegateWrapper) ComboBoxSelectionIsChanging(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](c_, "comboBoxSelectionIsChanging:", notification)
}

func (c_ *ComboBoxDelegateWrapper) ImplementsComboBoxWillDismiss() bool {
	return c_.RespondsToSelector(objc.GetSelector("comboBoxWillDismiss:"))
}

func (c_ ComboBoxDelegateWrapper) ComboBoxWillDismiss(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](c_, "comboBoxWillDismiss:", notification)
}

func (c_ *ComboBoxDelegateWrapper) ImplementsComboBoxWillPopUp() bool {
	return c_.RespondsToSelector(objc.GetSelector("comboBoxWillPopUp:"))
}

func (c_ ComboBoxDelegateWrapper) ComboBoxWillPopUp(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](c_, "comboBoxWillPopUp:", notification)
}
