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
