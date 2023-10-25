// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type ToolbarDelegate interface {
	ImplementsToolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar() bool
	// optional
	Toolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem
	ImplementsToolbarWillAddItem() bool
	// optional
	ToolbarWillAddItem(notification foundation.Notification)
	ImplementsToolbarDidRemoveItem() bool
	// optional
	ToolbarDidRemoveItem(notification foundation.Notification)
	ImplementsToolbarAllowedItemIdentifiers() bool
	// optional
	ToolbarAllowedItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier
	ImplementsToolbarDefaultItemIdentifiers() bool
	// optional
	ToolbarDefaultItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier
	ImplementsToolbarImmovableItemIdentifiers() bool
	// optional
	ToolbarImmovableItemIdentifiers(toolbar Toolbar) foundation.ISet
	ImplementsToolbarSelectableItemIdentifiers() bool
	// optional
	ToolbarSelectableItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier
	ImplementsToolbar_ItemIdentifier_CanBeInsertedAtIndex() bool
	// optional
	Toolbar_ItemIdentifier_CanBeInsertedAtIndex(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, index int) bool
}

func WrapToolbarDelegate(v ToolbarDelegate) objc.Object {
	return objc.WrapAsProtocol("NSToolbarDelegate", v)
}

type ToolbarDelegateBase struct {
}

func (p *ToolbarDelegateBase) ImplementsToolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar() bool {
	return false
}

func (p *ToolbarDelegateBase) Toolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem {
	panic("unimpemented protocol method")
}

func (p *ToolbarDelegateBase) ImplementsToolbarWillAddItem() bool {
	return false
}

func (p *ToolbarDelegateBase) ToolbarWillAddItem(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ToolbarDelegateBase) ImplementsToolbarDidRemoveItem() bool {
	return false
}

func (p *ToolbarDelegateBase) ToolbarDidRemoveItem(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ToolbarDelegateBase) ImplementsToolbarAllowedItemIdentifiers() bool {
	return false
}

func (p *ToolbarDelegateBase) ToolbarAllowedItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier {
	panic("unimpemented protocol method")
}

func (p *ToolbarDelegateBase) ImplementsToolbarDefaultItemIdentifiers() bool {
	return false
}

func (p *ToolbarDelegateBase) ToolbarDefaultItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier {
	panic("unimpemented protocol method")
}

func (p *ToolbarDelegateBase) ImplementsToolbarImmovableItemIdentifiers() bool {
	return false
}

func (p *ToolbarDelegateBase) ToolbarImmovableItemIdentifiers(toolbar Toolbar) foundation.ISet {
	panic("unimpemented protocol method")
}

func (p *ToolbarDelegateBase) ImplementsToolbarSelectableItemIdentifiers() bool {
	return false
}

func (p *ToolbarDelegateBase) ToolbarSelectableItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier {
	panic("unimpemented protocol method")
}

func (p *ToolbarDelegateBase) ImplementsToolbar_ItemIdentifier_CanBeInsertedAtIndex() bool {
	return false
}

func (p *ToolbarDelegateBase) Toolbar_ItemIdentifier_CanBeInsertedAtIndex(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, index int) bool {
	panic("unimpemented protocol method")
}

type ToolbarDelegateCreator struct {
	className string
	class     objc.Class
}

func NewToolbarDelegateCreator(name string) *ToolbarDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &ToolbarDelegateCreator{className: name, class: class}
}

func (c *ToolbarDelegateCreator) SetToolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar(handle func(o objc.Object, toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem) {
	objc.AddMethod(c.class, objc.GetSelector("toolbar:itemForItemIdentifier:willBeInsertedIntoToolbar:"), handle)
}

func (c *ToolbarDelegateCreator) SetToolbarWillAddItem(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.GetSelector("toolbarWillAddItem:"), handle)
}

func (c *ToolbarDelegateCreator) SetToolbarDidRemoveItem(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.GetSelector("toolbarDidRemoveItem:"), handle)
}

func (c *ToolbarDelegateCreator) SetToolbarAllowedItemIdentifiers(handle func(o objc.Object, toolbar Toolbar) []ToolbarItemIdentifier) {
	objc.AddMethod(c.class, objc.GetSelector("toolbarAllowedItemIdentifiers:"), handle)
}

func (c *ToolbarDelegateCreator) SetToolbarDefaultItemIdentifiers(handle func(o objc.Object, toolbar Toolbar) []ToolbarItemIdentifier) {
	objc.AddMethod(c.class, objc.GetSelector("toolbarDefaultItemIdentifiers:"), handle)
}

func (c *ToolbarDelegateCreator) SetToolbarImmovableItemIdentifiers(handle func(o objc.Object, toolbar Toolbar) foundation.ISet) {
	objc.AddMethod(c.class, objc.GetSelector("toolbarImmovableItemIdentifiers:"), handle)
}

func (c *ToolbarDelegateCreator) SetToolbarSelectableItemIdentifiers(handle func(o objc.Object, toolbar Toolbar) []ToolbarItemIdentifier) {
	objc.AddMethod(c.class, objc.GetSelector("toolbarSelectableItemIdentifiers:"), handle)
}

func (c *ToolbarDelegateCreator) SetToolbar_ItemIdentifier_CanBeInsertedAtIndex(handle func(o objc.Object, toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, index int) bool) {
	objc.AddMethod(c.class, objc.GetSelector("toolbar:itemIdentifier:canBeInsertedAtIndex:"), handle)
}

func (c *ToolbarDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
