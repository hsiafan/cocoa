// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type PopoverDelegate interface {
	ImplementsDetachableWindowForPopover() bool
	// optional
	DetachableWindowForPopover(popover Popover) IWindow
	ImplementsPopoverShouldClose() bool
	// optional
	PopoverShouldClose(popover Popover) bool
	ImplementsPopoverWillShow() bool
	// optional
	PopoverWillShow(notification foundation.Notification)
	ImplementsPopoverDidShow() bool
	// optional
	PopoverDidShow(notification foundation.Notification)
	ImplementsPopoverWillClose() bool
	// optional
	PopoverWillClose(notification foundation.Notification)
	ImplementsPopoverDidClose() bool
	// optional
	PopoverDidClose(notification foundation.Notification)
	ImplementsPopoverDidDetach() bool
	// optional
	PopoverDidDetach(popover Popover)
	ImplementsPopoverShouldDetach() bool
	// optional
	PopoverShouldDetach(popover Popover) bool
}

type PopoverDelegateImpl struct {
	_DetachableWindowForPopover func(popover Popover) IWindow
	_PopoverShouldClose         func(popover Popover) bool
	_PopoverWillShow            func(notification foundation.Notification)
	_PopoverDidShow             func(notification foundation.Notification)
	_PopoverWillClose           func(notification foundation.Notification)
	_PopoverDidClose            func(notification foundation.Notification)
	_PopoverDidDetach           func(popover Popover)
	_PopoverShouldDetach        func(popover Popover) bool
}

func (di *PopoverDelegateImpl) ImplementsDetachableWindowForPopover() bool {
	return di._DetachableWindowForPopover != nil
}

func (di *PopoverDelegateImpl) SetDetachableWindowForPopover(f func(popover Popover) IWindow) {
	di._DetachableWindowForPopover = f
}

func (di *PopoverDelegateImpl) DetachableWindowForPopover(popover Popover) IWindow {
	return di._DetachableWindowForPopover(popover)
}
func (di *PopoverDelegateImpl) ImplementsPopoverShouldClose() bool {
	return di._PopoverShouldClose != nil
}

func (di *PopoverDelegateImpl) SetPopoverShouldClose(f func(popover Popover) bool) {
	di._PopoverShouldClose = f
}

func (di *PopoverDelegateImpl) PopoverShouldClose(popover Popover) bool {
	return di._PopoverShouldClose(popover)
}
func (di *PopoverDelegateImpl) ImplementsPopoverWillShow() bool {
	return di._PopoverWillShow != nil
}

func (di *PopoverDelegateImpl) SetPopoverWillShow(f func(notification foundation.Notification)) {
	di._PopoverWillShow = f
}

func (di *PopoverDelegateImpl) PopoverWillShow(notification foundation.Notification) {
	di._PopoverWillShow(notification)
}
func (di *PopoverDelegateImpl) ImplementsPopoverDidShow() bool {
	return di._PopoverDidShow != nil
}

func (di *PopoverDelegateImpl) SetPopoverDidShow(f func(notification foundation.Notification)) {
	di._PopoverDidShow = f
}

func (di *PopoverDelegateImpl) PopoverDidShow(notification foundation.Notification) {
	di._PopoverDidShow(notification)
}
func (di *PopoverDelegateImpl) ImplementsPopoverWillClose() bool {
	return di._PopoverWillClose != nil
}

func (di *PopoverDelegateImpl) SetPopoverWillClose(f func(notification foundation.Notification)) {
	di._PopoverWillClose = f
}

func (di *PopoverDelegateImpl) PopoverWillClose(notification foundation.Notification) {
	di._PopoverWillClose(notification)
}
func (di *PopoverDelegateImpl) ImplementsPopoverDidClose() bool {
	return di._PopoverDidClose != nil
}

func (di *PopoverDelegateImpl) SetPopoverDidClose(f func(notification foundation.Notification)) {
	di._PopoverDidClose = f
}

func (di *PopoverDelegateImpl) PopoverDidClose(notification foundation.Notification) {
	di._PopoverDidClose(notification)
}
func (di *PopoverDelegateImpl) ImplementsPopoverDidDetach() bool {
	return di._PopoverDidDetach != nil
}

func (di *PopoverDelegateImpl) SetPopoverDidDetach(f func(popover Popover)) {
	di._PopoverDidDetach = f
}

func (di *PopoverDelegateImpl) PopoverDidDetach(popover Popover) {
	di._PopoverDidDetach(popover)
}
func (di *PopoverDelegateImpl) ImplementsPopoverShouldDetach() bool {
	return di._PopoverShouldDetach != nil
}

func (di *PopoverDelegateImpl) SetPopoverShouldDetach(f func(popover Popover) bool) {
	di._PopoverShouldDetach = f
}

func (di *PopoverDelegateImpl) PopoverShouldDetach(popover Popover) bool {
	return di._PopoverShouldDetach(popover)
}

type PopoverDelegateWrapper struct {
	objc.Object
}

func (p_ *PopoverDelegateWrapper) ImplementsDetachableWindowForPopover() bool {
	return p_.RespondsToSelector(objc.GetSelector("detachableWindowForPopover:"))
}

func (p_ PopoverDelegateWrapper) DetachableWindowForPopover(popover IPopover) Window {
	rv := objc.CallMethod[Window](p_, "detachableWindowForPopover:", popover)
	return rv
}

func (p_ *PopoverDelegateWrapper) ImplementsPopoverShouldClose() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverShouldClose:"))
}

func (p_ PopoverDelegateWrapper) PopoverShouldClose(popover IPopover) bool {
	rv := objc.CallMethod[bool](p_, "popoverShouldClose:", popover)
	return rv
}

func (p_ *PopoverDelegateWrapper) ImplementsPopoverWillShow() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverWillShow:"))
}

func (p_ PopoverDelegateWrapper) PopoverWillShow(notification foundation.INotification) {
	objc.CallMethod[objc.Void](p_, "popoverWillShow:", notification)
}

func (p_ *PopoverDelegateWrapper) ImplementsPopoverDidShow() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverDidShow:"))
}

func (p_ PopoverDelegateWrapper) PopoverDidShow(notification foundation.INotification) {
	objc.CallMethod[objc.Void](p_, "popoverDidShow:", notification)
}

func (p_ *PopoverDelegateWrapper) ImplementsPopoverWillClose() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverWillClose:"))
}

func (p_ PopoverDelegateWrapper) PopoverWillClose(notification foundation.INotification) {
	objc.CallMethod[objc.Void](p_, "popoverWillClose:", notification)
}

func (p_ *PopoverDelegateWrapper) ImplementsPopoverDidClose() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverDidClose:"))
}

func (p_ PopoverDelegateWrapper) PopoverDidClose(notification foundation.INotification) {
	objc.CallMethod[objc.Void](p_, "popoverDidClose:", notification)
}

func (p_ *PopoverDelegateWrapper) ImplementsPopoverDidDetach() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverDidDetach:"))
}

func (p_ PopoverDelegateWrapper) PopoverDidDetach(popover IPopover) {
	objc.CallMethod[objc.Void](p_, "popoverDidDetach:", popover)
}

func (p_ *PopoverDelegateWrapper) ImplementsPopoverShouldDetach() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverShouldDetach:"))
}

func (p_ PopoverDelegateWrapper) PopoverShouldDetach(popover IPopover) bool {
	rv := objc.CallMethod[bool](p_, "popoverShouldDetach:", popover)
	return rv
}
