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

func WrapPopoverDelegate(v PopoverDelegate) objc.Object {
	return objc.WrapAsProtocol("NSPopoverDelegate", v)
}

type PopoverDelegateBase struct {
}

func (p *PopoverDelegateBase) ImplementsDetachableWindowForPopover() bool {
	return false
}

func (p *PopoverDelegateBase) DetachableWindowForPopover(popover Popover) IWindow {
	panic("unimpemented protocol method")
}

func (p *PopoverDelegateBase) ImplementsPopoverShouldClose() bool {
	return false
}

func (p *PopoverDelegateBase) PopoverShouldClose(popover Popover) bool {
	panic("unimpemented protocol method")
}

func (p *PopoverDelegateBase) ImplementsPopoverWillShow() bool {
	return false
}

func (p *PopoverDelegateBase) PopoverWillShow(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *PopoverDelegateBase) ImplementsPopoverDidShow() bool {
	return false
}

func (p *PopoverDelegateBase) PopoverDidShow(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *PopoverDelegateBase) ImplementsPopoverWillClose() bool {
	return false
}

func (p *PopoverDelegateBase) PopoverWillClose(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *PopoverDelegateBase) ImplementsPopoverDidClose() bool {
	return false
}

func (p *PopoverDelegateBase) PopoverDidClose(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *PopoverDelegateBase) ImplementsPopoverDidDetach() bool {
	return false
}

func (p *PopoverDelegateBase) PopoverDidDetach(popover Popover) {
	panic("unimpemented protocol method")
}

func (p *PopoverDelegateBase) ImplementsPopoverShouldDetach() bool {
	return false
}

func (p *PopoverDelegateBase) PopoverShouldDetach(popover Popover) bool {
	panic("unimpemented protocol method")
}
