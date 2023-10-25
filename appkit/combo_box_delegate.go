// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

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

func WrapComboBoxDelegate(v ComboBoxDelegate) objc.Object {
	return objc.WrapAsProtocol("NSComboBoxDelegate", v)
}

type ComboBoxDelegateBase struct {
	TextFieldDelegateBase
}

func (p *ComboBoxDelegateBase) ImplementsComboBoxSelectionDidChange() bool {
	return false
}

func (p *ComboBoxDelegateBase) ComboBoxSelectionDidChange(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ComboBoxDelegateBase) ImplementsComboBoxSelectionIsChanging() bool {
	return false
}

func (p *ComboBoxDelegateBase) ComboBoxSelectionIsChanging(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ComboBoxDelegateBase) ImplementsComboBoxWillDismiss() bool {
	return false
}

func (p *ComboBoxDelegateBase) ComboBoxWillDismiss(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ComboBoxDelegateBase) ImplementsComboBoxWillPopUp() bool {
	return false
}

func (p *ComboBoxDelegateBase) ComboBoxWillPopUp(notification foundation.Notification) {
	panic("unimpemented protocol method")
}
