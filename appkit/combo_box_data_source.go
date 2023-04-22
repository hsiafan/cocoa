// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type ComboBoxDataSource interface {
	ImplementsComboBox_CompletedString() bool
	// optional
	ComboBox_CompletedString(comboBox ComboBox, string_ string) string
	ImplementsComboBox_IndexOfItemWithStringValue() bool
	// optional
	ComboBox_IndexOfItemWithStringValue(comboBox ComboBox, string_ string) uint
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

func (c_ ComboBoxDataSourceWrapper) ComboBox_CompletedString(comboBox IComboBox, string_ string) string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("comboBox:completedString:"), comboBox, string_)
	return rv
}

func (c_ *ComboBoxDataSourceWrapper) ImplementsComboBox_IndexOfItemWithStringValue() bool {
	return c_.RespondsToSelector(objc.GetSelector("comboBox:indexOfItemWithStringValue:"))
}

func (c_ ComboBoxDataSourceWrapper) ComboBox_IndexOfItemWithStringValue(comboBox IComboBox, string_ string) uint {
	rv := objc.CallMethod[uint](c_, objc.GetSelector("comboBox:indexOfItemWithStringValue:"), comboBox, string_)
	return rv
}

func (c_ *ComboBoxDataSourceWrapper) ImplementsComboBox_ObjectValueForItemAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("comboBox:objectValueForItemAtIndex:"))
}

func (c_ ComboBoxDataSourceWrapper) ComboBox_ObjectValueForItemAtIndex(comboBox IComboBox, index int) objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.GetSelector("comboBox:objectValueForItemAtIndex:"), comboBox, index)
	return rv
}

func (c_ *ComboBoxDataSourceWrapper) ImplementsNumberOfItemsInComboBox() bool {
	return c_.RespondsToSelector(objc.GetSelector("numberOfItemsInComboBox:"))
}

func (c_ ComboBoxDataSourceWrapper) NumberOfItemsInComboBox(comboBox IComboBox) int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("numberOfItemsInComboBox:"), comboBox)
	return rv
}
