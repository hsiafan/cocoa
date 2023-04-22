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

type ToolbarDelegateImpl struct {
	_Toolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar func(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem
	_ToolbarWillAddItem                                      func(notification foundation.Notification)
	_ToolbarDidRemoveItem                                    func(notification foundation.Notification)
	_ToolbarAllowedItemIdentifiers                           func(toolbar Toolbar) []ToolbarItemIdentifier
	_ToolbarDefaultItemIdentifiers                           func(toolbar Toolbar) []ToolbarItemIdentifier
	_ToolbarImmovableItemIdentifiers                         func(toolbar Toolbar) foundation.ISet
	_ToolbarSelectableItemIdentifiers                        func(toolbar Toolbar) []ToolbarItemIdentifier
	_Toolbar_ItemIdentifier_CanBeInsertedAtIndex             func(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, index int) bool
}

func (di *ToolbarDelegateImpl) ImplementsToolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar() bool {
	return di._Toolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar != nil
}

func (di *ToolbarDelegateImpl) SetToolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar(f func(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem) {
	di._Toolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar = f
}

func (di *ToolbarDelegateImpl) Toolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, flag bool) IToolbarItem {
	return di._Toolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar(toolbar, itemIdentifier, flag)
}
func (di *ToolbarDelegateImpl) ImplementsToolbarWillAddItem() bool {
	return di._ToolbarWillAddItem != nil
}

func (di *ToolbarDelegateImpl) SetToolbarWillAddItem(f func(notification foundation.Notification)) {
	di._ToolbarWillAddItem = f
}

func (di *ToolbarDelegateImpl) ToolbarWillAddItem(notification foundation.Notification) {
	di._ToolbarWillAddItem(notification)
}
func (di *ToolbarDelegateImpl) ImplementsToolbarDidRemoveItem() bool {
	return di._ToolbarDidRemoveItem != nil
}

func (di *ToolbarDelegateImpl) SetToolbarDidRemoveItem(f func(notification foundation.Notification)) {
	di._ToolbarDidRemoveItem = f
}

func (di *ToolbarDelegateImpl) ToolbarDidRemoveItem(notification foundation.Notification) {
	di._ToolbarDidRemoveItem(notification)
}
func (di *ToolbarDelegateImpl) ImplementsToolbarAllowedItemIdentifiers() bool {
	return di._ToolbarAllowedItemIdentifiers != nil
}

func (di *ToolbarDelegateImpl) SetToolbarAllowedItemIdentifiers(f func(toolbar Toolbar) []ToolbarItemIdentifier) {
	di._ToolbarAllowedItemIdentifiers = f
}

func (di *ToolbarDelegateImpl) ToolbarAllowedItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier {
	return di._ToolbarAllowedItemIdentifiers(toolbar)
}
func (di *ToolbarDelegateImpl) ImplementsToolbarDefaultItemIdentifiers() bool {
	return di._ToolbarDefaultItemIdentifiers != nil
}

func (di *ToolbarDelegateImpl) SetToolbarDefaultItemIdentifiers(f func(toolbar Toolbar) []ToolbarItemIdentifier) {
	di._ToolbarDefaultItemIdentifiers = f
}

func (di *ToolbarDelegateImpl) ToolbarDefaultItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier {
	return di._ToolbarDefaultItemIdentifiers(toolbar)
}
func (di *ToolbarDelegateImpl) ImplementsToolbarImmovableItemIdentifiers() bool {
	return di._ToolbarImmovableItemIdentifiers != nil
}

func (di *ToolbarDelegateImpl) SetToolbarImmovableItemIdentifiers(f func(toolbar Toolbar) foundation.ISet) {
	di._ToolbarImmovableItemIdentifiers = f
}

func (di *ToolbarDelegateImpl) ToolbarImmovableItemIdentifiers(toolbar Toolbar) foundation.ISet {
	return di._ToolbarImmovableItemIdentifiers(toolbar)
}
func (di *ToolbarDelegateImpl) ImplementsToolbarSelectableItemIdentifiers() bool {
	return di._ToolbarSelectableItemIdentifiers != nil
}

func (di *ToolbarDelegateImpl) SetToolbarSelectableItemIdentifiers(f func(toolbar Toolbar) []ToolbarItemIdentifier) {
	di._ToolbarSelectableItemIdentifiers = f
}

func (di *ToolbarDelegateImpl) ToolbarSelectableItemIdentifiers(toolbar Toolbar) []ToolbarItemIdentifier {
	return di._ToolbarSelectableItemIdentifiers(toolbar)
}
func (di *ToolbarDelegateImpl) ImplementsToolbar_ItemIdentifier_CanBeInsertedAtIndex() bool {
	return di._Toolbar_ItemIdentifier_CanBeInsertedAtIndex != nil
}

func (di *ToolbarDelegateImpl) SetToolbar_ItemIdentifier_CanBeInsertedAtIndex(f func(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, index int) bool) {
	di._Toolbar_ItemIdentifier_CanBeInsertedAtIndex = f
}

func (di *ToolbarDelegateImpl) Toolbar_ItemIdentifier_CanBeInsertedAtIndex(toolbar Toolbar, itemIdentifier ToolbarItemIdentifier, index int) bool {
	return di._Toolbar_ItemIdentifier_CanBeInsertedAtIndex(toolbar, itemIdentifier, index)
}

type ToolbarDelegateWrapper struct {
	objc.Object
}

func (t_ *ToolbarDelegateWrapper) ImplementsToolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbar:itemForItemIdentifier:willBeInsertedIntoToolbar:"))
}

func (t_ ToolbarDelegateWrapper) Toolbar_ItemForItemIdentifier_WillBeInsertedIntoToolbar(toolbar IToolbar, itemIdentifier ToolbarItemIdentifier, flag bool) ToolbarItem {
	rv := objc.CallMethod[ToolbarItem](t_, objc.GetSelector("toolbar:itemForItemIdentifier:willBeInsertedIntoToolbar:"), toolbar, itemIdentifier, flag)
	return rv
}

func (t_ *ToolbarDelegateWrapper) ImplementsToolbarWillAddItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbarWillAddItem:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarWillAddItem(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("toolbarWillAddItem:"), notification)
}

func (t_ *ToolbarDelegateWrapper) ImplementsToolbarDidRemoveItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbarDidRemoveItem:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarDidRemoveItem(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("toolbarDidRemoveItem:"), notification)
}

func (t_ *ToolbarDelegateWrapper) ImplementsToolbarAllowedItemIdentifiers() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbarAllowedItemIdentifiers:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarAllowedItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier {
	rv := objc.CallMethod[[]ToolbarItemIdentifier](t_, objc.GetSelector("toolbarAllowedItemIdentifiers:"), toolbar)
	return rv
}

func (t_ *ToolbarDelegateWrapper) ImplementsToolbarDefaultItemIdentifiers() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbarDefaultItemIdentifiers:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarDefaultItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier {
	rv := objc.CallMethod[[]ToolbarItemIdentifier](t_, objc.GetSelector("toolbarDefaultItemIdentifiers:"), toolbar)
	return rv
}

func (t_ *ToolbarDelegateWrapper) ImplementsToolbarImmovableItemIdentifiers() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbarImmovableItemIdentifiers:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarImmovableItemIdentifiers(toolbar IToolbar) foundation.Set {
	rv := objc.CallMethod[foundation.Set](t_, objc.GetSelector("toolbarImmovableItemIdentifiers:"), toolbar)
	return rv
}

func (t_ *ToolbarDelegateWrapper) ImplementsToolbarSelectableItemIdentifiers() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbarSelectableItemIdentifiers:"))
}

func (t_ ToolbarDelegateWrapper) ToolbarSelectableItemIdentifiers(toolbar IToolbar) []ToolbarItemIdentifier {
	rv := objc.CallMethod[[]ToolbarItemIdentifier](t_, objc.GetSelector("toolbarSelectableItemIdentifiers:"), toolbar)
	return rv
}

func (t_ *ToolbarDelegateWrapper) ImplementsToolbar_ItemIdentifier_CanBeInsertedAtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("toolbar:itemIdentifier:canBeInsertedAtIndex:"))
}

func (t_ ToolbarDelegateWrapper) Toolbar_ItemIdentifier_CanBeInsertedAtIndex(toolbar IToolbar, itemIdentifier ToolbarItemIdentifier, index int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("toolbar:itemIdentifier:canBeInsertedAtIndex:"), toolbar, itemIdentifier, index)
	return rv
}
