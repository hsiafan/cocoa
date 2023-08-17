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

type ToolbarDelegateWrapper struct {
	objc.Object
}

func (t_ ToolbarDelegateWrapper) Toolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar(toolbar IToolbar, itemIdentifier ToolbarItemIdentifier, flag bool) ToolbarItem {
	rv := objc.CallMethod[ToolbarItem](t_, objc.GetSelector("toolbar:itemForItemIdentifier:willBeInsertedIntoToolbar:"), objc.ExtractPtr(toolbar), itemIdentifier, flag)
	return rv
}

func (t_ ToolbarDelegateWrapper) ToolbarWillAddItem(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("toolbarWillAddItem:"), objc.ExtractPtr(notification))
}

func (t_ ToolbarDelegateWrapper) ToolbarDidRemoveItem(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("toolbarDidRemoveItem:"), objc.ExtractPtr(notification))
}

func (t_ ToolbarDelegateWrapper) ToolbarAllowedItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier {
	rv := objc.CallMethod[[]ToolbarItemIdentifier](t_, objc.GetSelector("toolbarAllowedItemIdentifiers:"), objc.ExtractPtr(toolbar))
	return rv
}

func (t_ ToolbarDelegateWrapper) ToolbarDefaultItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier {
	rv := objc.CallMethod[[]ToolbarItemIdentifier](t_, objc.GetSelector("toolbarDefaultItemIdentifiers:"), objc.ExtractPtr(toolbar))
	return rv
}

func (t_ ToolbarDelegateWrapper) ToolbarImmovableItemIdentifiers(toolbar IToolbar) foundation.Set {
	rv := objc.CallMethod[foundation.Set](t_, objc.GetSelector("toolbarImmovableItemIdentifiers:"), objc.ExtractPtr(toolbar))
	return rv
}

func (t_ ToolbarDelegateWrapper) ToolbarSelectableItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier {
	rv := objc.CallMethod[[]ToolbarItemIdentifier](t_, objc.GetSelector("toolbarSelectableItemIdentifiers:"), objc.ExtractPtr(toolbar))
	return rv
}

func (t_ ToolbarDelegateWrapper) Toolbar_ItemIdentifier_CanBeInsertedAtIndex(toolbar IToolbar, itemIdentifier ToolbarItemIdentifier, index int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("toolbar:itemIdentifier:canBeInsertedAtIndex:"), objc.ExtractPtr(toolbar), itemIdentifier, index)
	return rv
}
