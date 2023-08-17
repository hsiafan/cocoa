// auto-generated code, do not modify
package appkit

import (
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

func WrapTabViewDelegate(v TabViewDelegate) objc.Object {
	return objc.WrapAsProtocol("NSTabViewDelegate", v)
}

type TabViewDelegateBase struct {
}

func (p *TabViewDelegateBase) ImplementsTabViewDidChangeNumberOfTabViewItems() bool {
	return false
}

func (p *TabViewDelegateBase) TabViewDidChangeNumberOfTabViewItems(tabView TabView) {
	panic("unimpemented protocol method")
}

func (p *TabViewDelegateBase) ImplementsTabView_ShouldSelectTabViewItem() bool {
	return false
}

func (p *TabViewDelegateBase) TabView_ShouldSelectTabViewItem(tabView TabView, tabViewItem TabViewItem) bool {
	panic("unimpemented protocol method")
}

func (p *TabViewDelegateBase) ImplementsTabView_WillSelectTabViewItem() bool {
	return false
}

func (p *TabViewDelegateBase) TabView_WillSelectTabViewItem(tabView TabView, tabViewItem TabViewItem) {
	panic("unimpemented protocol method")
}

func (p *TabViewDelegateBase) ImplementsTabView_DidSelectTabViewItem() bool {
	return false
}

func (p *TabViewDelegateBase) TabView_DidSelectTabViewItem(tabView TabView, tabViewItem TabViewItem) {
	panic("unimpemented protocol method")
}

type TabViewDelegateWrapper struct {
	objc.Object
}

func (t_ TabViewDelegateWrapper) TabViewDidChangeNumberOfTabViewItems(tabView ITabView) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tabViewDidChangeNumberOfTabViewItems:"), objc.ExtractPtr(tabView))
}

func (t_ TabViewDelegateWrapper) TabView_ShouldSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tabView:shouldSelectTabViewItem:"), objc.ExtractPtr(tabView), objc.ExtractPtr(tabViewItem))
	return rv
}

func (t_ TabViewDelegateWrapper) TabView_WillSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tabView:willSelectTabViewItem:"), objc.ExtractPtr(tabView), objc.ExtractPtr(tabViewItem))
}

func (t_ TabViewDelegateWrapper) TabView_DidSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tabView:didSelectTabViewItem:"), objc.ExtractPtr(tabView), objc.ExtractPtr(tabViewItem))
}
