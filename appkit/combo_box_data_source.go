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

func WrapComboBoxDataSource(v ComboBoxDataSource) objc.Object {
	return objc.WrapAsProtocol("NSComboBoxDataSource", v)
}

type ComboBoxDataSourceBase struct {
}

func (p *ComboBoxDataSourceBase) ImplementsComboBox_CompletedString() bool {
	return false
}

func (p *ComboBoxDataSourceBase) ComboBox_CompletedString(comboBox ComboBox, string_ string) string {
	panic("unimpemented protocol method")
}

func (p *ComboBoxDataSourceBase) ImplementsComboBox_IndexOfItemWithStringValue() bool {
	return false
}

func (p *ComboBoxDataSourceBase) ComboBox_IndexOfItemWithStringValue(comboBox ComboBox, string_ string) uint {
	panic("unimpemented protocol method")
}

func (p *ComboBoxDataSourceBase) ImplementsComboBox_ObjectValueForItemAtIndex() bool {
	return false
}

func (p *ComboBoxDataSourceBase) ComboBox_ObjectValueForItemAtIndex(comboBox ComboBox, index int) objc.IObject {
	panic("unimpemented protocol method")
}

func (p *ComboBoxDataSourceBase) ImplementsNumberOfItemsInComboBox() bool {
	return false
}

func (p *ComboBoxDataSourceBase) NumberOfItemsInComboBox(comboBox ComboBox) int {
	panic("unimpemented protocol method")
}

type ComboBoxDataSourceWrapper struct {
	objc.Object
}

func (c_ ComboBoxDataSourceWrapper) ComboBox_CompletedString(comboBox IComboBox, string_ string) string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("comboBox:completedString:"), objc.ExtractPtr(comboBox), string_)
	return rv
}

func (c_ ComboBoxDataSourceWrapper) ComboBox_IndexOfItemWithStringValue(comboBox IComboBox, string_ string) uint {
	rv := objc.CallMethod[uint](c_, objc.GetSelector("comboBox:indexOfItemWithStringValue:"), objc.ExtractPtr(comboBox), string_)
	return rv
}

func (c_ ComboBoxDataSourceWrapper) ComboBox_ObjectValueForItemAtIndex(comboBox IComboBox, index int) objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.GetSelector("comboBox:objectValueForItemAtIndex:"), objc.ExtractPtr(comboBox), index)
	return rv
}

func (c_ ComboBoxDataSourceWrapper) NumberOfItemsInComboBox(comboBox IComboBox) int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("numberOfItemsInComboBox:"), objc.ExtractPtr(comboBox))
	return rv
}
