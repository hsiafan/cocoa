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

type DrawerDelegateCreator struct {
	className string
	class     objc.Class
}

func NewDrawerDelegateCreator(name string) *DrawerDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("ProtocolBase"), name, 0)
	objc.RegisterClassPair(class)
	return &DrawerDelegateCreator{className: name, class: class}
}

// deprecated
func (c *DrawerDelegateCreator) SetDrawerShouldOpen(handle func(o objc.ProtocolBase, sender Drawer) bool) {
	objc.AddMethod(c.class, objc.SEL("drawerShouldOpen:"), handle)
}

// deprecated
func (c *DrawerDelegateCreator) SetDrawerWillOpen(handle func(o objc.ProtocolBase, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("drawerWillOpen:"), handle)
}

// deprecated
func (c *DrawerDelegateCreator) SetDrawerDidOpen(handle func(o objc.ProtocolBase, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("drawerDidOpen:"), handle)
}

// deprecated
func (c *DrawerDelegateCreator) SetDrawerShouldClose(handle func(o objc.ProtocolBase, sender Drawer) bool) {
	objc.AddMethod(c.class, objc.SEL("drawerShouldClose:"), handle)
}

// deprecated
func (c *DrawerDelegateCreator) SetDrawerWillClose(handle func(o objc.ProtocolBase, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("drawerWillClose:"), handle)
}

// deprecated
func (c *DrawerDelegateCreator) SetDrawerDidClose(handle func(o objc.ProtocolBase, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("drawerDidClose:"), handle)
}

// deprecated
func (c *DrawerDelegateCreator) SetDrawerWillResizeContents_ToSize(handle func(o objc.ProtocolBase, sender Drawer, contentSize foundation.Size) foundation.Size) {
	objc.AddMethod(c.class, objc.SEL("drawerWillResizeContents:toSize:"), handle)
}

func (c *DrawerDelegateCreator) Create() objc.ProtocolBase {
	return objc.ProtocolBase{Object: c.class.CreateInstance(0)}
}
