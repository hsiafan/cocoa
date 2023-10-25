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
