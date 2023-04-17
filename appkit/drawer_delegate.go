// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
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

type DrawerDelegateImpl struct {
	_DrawerShouldOpen                func(sender Drawer) bool
	_DrawerWillOpen                  func(notification foundation.Notification)
	_DrawerDidOpen                   func(notification foundation.Notification)
	_DrawerShouldClose               func(sender Drawer) bool
	_DrawerWillClose                 func(notification foundation.Notification)
	_DrawerDidClose                  func(notification foundation.Notification)
	_DrawerWillResizeContents_ToSize func(sender Drawer, contentSize foundation.Size) foundation.Size
}

func (di *DrawerDelegateImpl) ImplementsDrawerShouldOpen() bool {
	return di._DrawerShouldOpen != nil
}

// deprecated
func (di *DrawerDelegateImpl) SetDrawerShouldOpen(f func(sender Drawer) bool) {
	di._DrawerShouldOpen = f
}

// deprecated
func (di *DrawerDelegateImpl) DrawerShouldOpen(sender Drawer) bool {
	return di._DrawerShouldOpen(sender)
}
func (di *DrawerDelegateImpl) ImplementsDrawerWillOpen() bool {
	return di._DrawerWillOpen != nil
}

// deprecated
func (di *DrawerDelegateImpl) SetDrawerWillOpen(f func(notification foundation.Notification)) {
	di._DrawerWillOpen = f
}

// deprecated
func (di *DrawerDelegateImpl) DrawerWillOpen(notification foundation.Notification) {
	di._DrawerWillOpen(notification)
}
func (di *DrawerDelegateImpl) ImplementsDrawerDidOpen() bool {
	return di._DrawerDidOpen != nil
}

// deprecated
func (di *DrawerDelegateImpl) SetDrawerDidOpen(f func(notification foundation.Notification)) {
	di._DrawerDidOpen = f
}

// deprecated
func (di *DrawerDelegateImpl) DrawerDidOpen(notification foundation.Notification) {
	di._DrawerDidOpen(notification)
}
func (di *DrawerDelegateImpl) ImplementsDrawerShouldClose() bool {
	return di._DrawerShouldClose != nil
}

// deprecated
func (di *DrawerDelegateImpl) SetDrawerShouldClose(f func(sender Drawer) bool) {
	di._DrawerShouldClose = f
}

// deprecated
func (di *DrawerDelegateImpl) DrawerShouldClose(sender Drawer) bool {
	return di._DrawerShouldClose(sender)
}
func (di *DrawerDelegateImpl) ImplementsDrawerWillClose() bool {
	return di._DrawerWillClose != nil
}

// deprecated
func (di *DrawerDelegateImpl) SetDrawerWillClose(f func(notification foundation.Notification)) {
	di._DrawerWillClose = f
}

// deprecated
func (di *DrawerDelegateImpl) DrawerWillClose(notification foundation.Notification) {
	di._DrawerWillClose(notification)
}
func (di *DrawerDelegateImpl) ImplementsDrawerDidClose() bool {
	return di._DrawerDidClose != nil
}

// deprecated
func (di *DrawerDelegateImpl) SetDrawerDidClose(f func(notification foundation.Notification)) {
	di._DrawerDidClose = f
}

// deprecated
func (di *DrawerDelegateImpl) DrawerDidClose(notification foundation.Notification) {
	di._DrawerDidClose(notification)
}
func (di *DrawerDelegateImpl) ImplementsDrawerWillResizeContents_ToSize() bool {
	return di._DrawerWillResizeContents_ToSize != nil
}

// deprecated
func (di *DrawerDelegateImpl) SetDrawerWillResizeContents_ToSize(f func(sender Drawer, contentSize foundation.Size) foundation.Size) {
	di._DrawerWillResizeContents_ToSize = f
}

// deprecated
func (di *DrawerDelegateImpl) DrawerWillResizeContents_ToSize(sender Drawer, contentSize foundation.Size) foundation.Size {
	return di._DrawerWillResizeContents_ToSize(sender, contentSize)
}

type DrawerDelegateWrapper struct {
	objc.Object
}

func (d_ *DrawerDelegateWrapper) ImplementsDrawerShouldOpen() bool {
	return d_.RespondsToSelector(objc.GetSelector("drawerShouldOpen:"))
}

// deprecated
func (d_ DrawerDelegateWrapper) DrawerShouldOpen(sender IDrawer) bool {
	rv := ffi.CallMethod[bool](d_, "drawerShouldOpen:", sender)
	return rv
}

func (d_ *DrawerDelegateWrapper) ImplementsDrawerWillOpen() bool {
	return d_.RespondsToSelector(objc.GetSelector("drawerWillOpen:"))
}

// deprecated
func (d_ DrawerDelegateWrapper) DrawerWillOpen(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](d_, "drawerWillOpen:", notification)
}

func (d_ *DrawerDelegateWrapper) ImplementsDrawerDidOpen() bool {
	return d_.RespondsToSelector(objc.GetSelector("drawerDidOpen:"))
}

// deprecated
func (d_ DrawerDelegateWrapper) DrawerDidOpen(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](d_, "drawerDidOpen:", notification)
}

func (d_ *DrawerDelegateWrapper) ImplementsDrawerShouldClose() bool {
	return d_.RespondsToSelector(objc.GetSelector("drawerShouldClose:"))
}

// deprecated
func (d_ DrawerDelegateWrapper) DrawerShouldClose(sender IDrawer) bool {
	rv := ffi.CallMethod[bool](d_, "drawerShouldClose:", sender)
	return rv
}

func (d_ *DrawerDelegateWrapper) ImplementsDrawerWillClose() bool {
	return d_.RespondsToSelector(objc.GetSelector("drawerWillClose:"))
}

// deprecated
func (d_ DrawerDelegateWrapper) DrawerWillClose(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](d_, "drawerWillClose:", notification)
}

func (d_ *DrawerDelegateWrapper) ImplementsDrawerDidClose() bool {
	return d_.RespondsToSelector(objc.GetSelector("drawerDidClose:"))
}

// deprecated
func (d_ DrawerDelegateWrapper) DrawerDidClose(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](d_, "drawerDidClose:", notification)
}

func (d_ *DrawerDelegateWrapper) ImplementsDrawerWillResizeContents_ToSize() bool {
	return d_.RespondsToSelector(objc.GetSelector("drawerWillResizeContents:toSize:"))
}

// deprecated
func (d_ DrawerDelegateWrapper) DrawerWillResizeContents_ToSize(sender IDrawer, contentSize foundation.Size) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](d_, "drawerWillResizeContents:toSize:", sender, contentSize)
	return rv
}
