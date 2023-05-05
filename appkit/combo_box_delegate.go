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
	objc.CallMethod[objc.Void](c_, objc.GetSelector("comboBoxSelectionDidChange:"), objc.ExtractPtr(notification))
}

func (c_ *ComboBoxDelegateWrapper) ImplementsComboBoxSelectionIsChanging() bool {
	return c_.RespondsToSelector(objc.GetSelector("comboBoxSelectionIsChanging:"))
}

func (c_ ComboBoxDelegateWrapper) ComboBoxSelectionIsChanging(notification foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("comboBoxSelectionIsChanging:"), objc.ExtractPtr(notification))
}

func (c_ *ComboBoxDelegateWrapper) ImplementsComboBoxWillDismiss() bool {
	return c_.RespondsToSelector(objc.GetSelector("comboBoxWillDismiss:"))
}

func (c_ ComboBoxDelegateWrapper) ComboBoxWillDismiss(notification foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("comboBoxWillDismiss:"), objc.ExtractPtr(notification))
}

func (c_ *ComboBoxDelegateWrapper) ImplementsComboBoxWillPopUp() bool {
	return c_.RespondsToSelector(objc.GetSelector("comboBoxWillPopUp:"))
}

func (c_ ComboBoxDelegateWrapper) ComboBoxWillPopUp(notification foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("comboBoxWillPopUp:"), objc.ExtractPtr(notification))
}
