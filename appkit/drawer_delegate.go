// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type DrawerDelegate interface {
	ImplementsDrawerShouldOpen() bool
	// optional
	// deprecated
	DrawerShouldOpen(sender Drawer) bool
	ImplementsDrawerWillOpen() bool
	// optional
	// deprecated
	DrawerWillOpen(notification foundation.Notification)
	ImplementsDrawerDidOpen() bool
	// optional
	// deprecated
	DrawerDidOpen(notification foundation.Notification)
	ImplementsDrawerShouldClose() bool
	// optional
	// deprecated
	DrawerShouldClose(sender Drawer) bool
	ImplementsDrawerWillClose() bool
	// optional
	// deprecated
	DrawerWillClose(notification foundation.Notification)
	ImplementsDrawerDidClose() bool
	// optional
	// deprecated
	DrawerDidClose(notification foundation.Notification)
	ImplementsDrawerWillResizeContents_ToSize() bool
	// optional
	// deprecated
	DrawerWillResizeContents_ToSize(sender Drawer, contentSize foundation.Size) foundation.Size
}

func WrapDrawerDelegate(v DrawerDelegate) objc.Object {
	return objc.WrapAsProtocol("NSDrawerDelegate", v)
}

type DrawerDelegateBase struct {
}

func (p *DrawerDelegateBase) ImplementsDrawerShouldOpen() bool {
	return false
}

// deprecated
func (p *DrawerDelegateBase) DrawerShouldOpen(sender Drawer) bool {
	panic("unimpemented protocol method")
}

func (p *DrawerDelegateBase) ImplementsDrawerWillOpen() bool {
	return false
}

// deprecated
func (p *DrawerDelegateBase) DrawerWillOpen(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *DrawerDelegateBase) ImplementsDrawerDidOpen() bool {
	return false
}

// deprecated
func (p *DrawerDelegateBase) DrawerDidOpen(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *DrawerDelegateBase) ImplementsDrawerShouldClose() bool {
	return false
}

// deprecated
func (p *DrawerDelegateBase) DrawerShouldClose(sender Drawer) bool {
	panic("unimpemented protocol method")
}

func (p *DrawerDelegateBase) ImplementsDrawerWillClose() bool {
	return false
}

// deprecated
func (p *DrawerDelegateBase) DrawerWillClose(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *DrawerDelegateBase) ImplementsDrawerDidClose() bool {
	return false
}

// deprecated
func (p *DrawerDelegateBase) DrawerDidClose(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *DrawerDelegateBase) ImplementsDrawerWillResizeContents_ToSize() bool {
	return false
}

// deprecated
func (p *DrawerDelegateBase) DrawerWillResizeContents_ToSize(sender Drawer, contentSize foundation.Size) foundation.Size {
	panic("unimpemented protocol method")
}
