// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

type TabViewDelegate interface {
	ImplementsTabViewDidChangeNumberOfTabViewItems() bool
	// optional
	TabViewDidChangeNumberOfTabViewItems(tabView TabView)
	ImplementsTabView_ShouldSelectTabViewItem() bool
	// optional
	TabView_ShouldSelectTabViewItem(tabView TabView, tabViewItem TabViewItem) bool
	ImplementsTabView_WillSelectTabViewItem() bool
	// optional
	TabView_WillSelectTabViewItem(tabView TabView, tabViewItem TabViewItem)
	ImplementsTabView_DidSelectTabViewItem() bool
	// optional
	TabView_DidSelectTabViewItem(tabView TabView, tabViewItem TabViewItem)
}

type TabViewDelegateImpl struct {
	_TabViewDidChangeNumberOfTabViewItems func(tabView TabView)
	_TabView_ShouldSelectTabViewItem      func(tabView TabView, tabViewItem TabViewItem) bool
	_TabView_WillSelectTabViewItem        func(tabView TabView, tabViewItem TabViewItem)
	_TabView_DidSelectTabViewItem         func(tabView TabView, tabViewItem TabViewItem)
}

func (di *TabViewDelegateImpl) ImplementsTabViewDidChangeNumberOfTabViewItems() bool {
	return di._TabViewDidChangeNumberOfTabViewItems != nil
}

func (di *TabViewDelegateImpl) SetTabViewDidChangeNumberOfTabViewItems(f func(tabView TabView)) {
	di._TabViewDidChangeNumberOfTabViewItems = f
}

func (di *TabViewDelegateImpl) TabViewDidChangeNumberOfTabViewItems(tabView TabView) {
	di._TabViewDidChangeNumberOfTabViewItems(tabView)
}
func (di *TabViewDelegateImpl) ImplementsTabView_ShouldSelectTabViewItem() bool {
	return di._TabView_ShouldSelectTabViewItem != nil
}

func (di *TabViewDelegateImpl) SetTabView_ShouldSelectTabViewItem(f func(tabView TabView, tabViewItem TabViewItem) bool) {
	di._TabView_ShouldSelectTabViewItem = f
}

func (di *TabViewDelegateImpl) TabView_ShouldSelectTabViewItem(tabView TabView, tabViewItem TabViewItem) bool {
	return di._TabView_ShouldSelectTabViewItem(tabView, tabViewItem)
}
func (di *TabViewDelegateImpl) ImplementsTabView_WillSelectTabViewItem() bool {
	return di._TabView_WillSelectTabViewItem != nil
}

func (di *TabViewDelegateImpl) SetTabView_WillSelectTabViewItem(f func(tabView TabView, tabViewItem TabViewItem)) {
	di._TabView_WillSelectTabViewItem = f
}

func (di *TabViewDelegateImpl) TabView_WillSelectTabViewItem(tabView TabView, tabViewItem TabViewItem) {
	di._TabView_WillSelectTabViewItem(tabView, tabViewItem)
}
func (di *TabViewDelegateImpl) ImplementsTabView_DidSelectTabViewItem() bool {
	return di._TabView_DidSelectTabViewItem != nil
}

func (di *TabViewDelegateImpl) SetTabView_DidSelectTabViewItem(f func(tabView TabView, tabViewItem TabViewItem)) {
	di._TabView_DidSelectTabViewItem = f
}

func (di *TabViewDelegateImpl) TabView_DidSelectTabViewItem(tabView TabView, tabViewItem TabViewItem) {
	di._TabView_DidSelectTabViewItem(tabView, tabViewItem)
}

type TabViewDelegateWrapper struct {
	objc.Object
}

func (t_ *TabViewDelegateWrapper) ImplementsTabViewDidChangeNumberOfTabViewItems() bool {
	return t_.RespondsToSelector(objc.GetSelector("tabViewDidChangeNumberOfTabViewItems:"))
}

func (t_ TabViewDelegateWrapper) TabViewDidChangeNumberOfTabViewItems(tabView ITabView) {
	ffi.CallMethod[ffi.Void](t_, "tabViewDidChangeNumberOfTabViewItems:", tabView)
}

func (t_ *TabViewDelegateWrapper) ImplementsTabView_ShouldSelectTabViewItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("tabView:shouldSelectTabViewItem:"))
}

func (t_ TabViewDelegateWrapper) TabView_ShouldSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) bool {
	rv := ffi.CallMethod[bool](t_, "tabView:shouldSelectTabViewItem:", tabView, tabViewItem)
	return rv
}

func (t_ *TabViewDelegateWrapper) ImplementsTabView_WillSelectTabViewItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("tabView:willSelectTabViewItem:"))
}

func (t_ TabViewDelegateWrapper) TabView_WillSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) {
	ffi.CallMethod[ffi.Void](t_, "tabView:willSelectTabViewItem:", tabView, tabViewItem)
}

func (t_ *TabViewDelegateWrapper) ImplementsTabView_DidSelectTabViewItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("tabView:didSelectTabViewItem:"))
}

func (t_ TabViewDelegateWrapper) TabView_DidSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) {
	ffi.CallMethod[ffi.Void](t_, "tabView:didSelectTabViewItem:", tabView, tabViewItem)
}
