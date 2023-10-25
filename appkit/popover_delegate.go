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

type PopoverDelegateCreator struct {
	className string
	class     objc.Class
}

func NewPopoverDelegateCreator(name string) *PopoverDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("ProtocolBase"), name, 0)
	objc.RegisterClassPair(class)
	return &PopoverDelegateCreator{className: name, class: class}
}

func (c *PopoverDelegateCreator) SetDetachableWindowForPopover(handle func(o objc.ProtocolBase, popover Popover) IWindow) {
	objc.AddMethod(c.class, objc.SEL("detachableWindowForPopover:"), handle)
}

func (c *PopoverDelegateCreator) SetPopoverShouldClose(handle func(o objc.ProtocolBase, popover Popover) bool) {
	objc.AddMethod(c.class, objc.SEL("popoverShouldClose:"), handle)
}

func (c *PopoverDelegateCreator) SetPopoverWillShow(handle func(o objc.ProtocolBase, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("popoverWillShow:"), handle)
}

func (c *PopoverDelegateCreator) SetPopoverDidShow(handle func(o objc.ProtocolBase, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("popoverDidShow:"), handle)
}

func (c *PopoverDelegateCreator) SetPopoverWillClose(handle func(o objc.ProtocolBase, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("popoverWillClose:"), handle)
}

func (c *PopoverDelegateCreator) SetPopoverDidClose(handle func(o objc.ProtocolBase, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("popoverDidClose:"), handle)
}

func (c *PopoverDelegateCreator) SetPopoverDidDetach(handle func(o objc.ProtocolBase, popover Popover)) {
	objc.AddMethod(c.class, objc.SEL("popoverDidDetach:"), handle)
}

func (c *PopoverDelegateCreator) SetPopoverShouldDetach(handle func(o objc.ProtocolBase, popover Popover) bool) {
	objc.AddMethod(c.class, objc.SEL("popoverShouldDetach:"), handle)
}

func (c *PopoverDelegateCreator) Create() objc.ProtocolBase {
	return objc.ProtocolBase{Object: c.class.CreateInstance(0)}
}
