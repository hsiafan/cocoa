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

type TabViewDelegateCreator struct {
	className string
	class     objc.Class
}

func NewTabViewDelegateCreator(name string) *TabViewDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &TabViewDelegateCreator{className: name, class: class}
}

func (c *TabViewDelegateCreator) SetTabViewDidChangeNumberOfTabViewItems(handle func(o objc.Object, tabView TabView)) {
	objc.AddMethod(c.class, objc.GetSelector("tabViewDidChangeNumberOfTabViewItems:"), handle)
}

func (c *TabViewDelegateCreator) SetTabView_ShouldSelectTabViewItem(handle func(o objc.Object, tabView TabView, tabViewItem TabViewItem) bool) {
	objc.AddMethod(c.class, objc.GetSelector("tabView:shouldSelectTabViewItem:"), handle)
}

func (c *TabViewDelegateCreator) SetTabView_WillSelectTabViewItem(handle func(o objc.Object, tabView TabView, tabViewItem TabViewItem)) {
	objc.AddMethod(c.class, objc.GetSelector("tabView:willSelectTabViewItem:"), handle)
}

func (c *TabViewDelegateCreator) SetTabView_DidSelectTabViewItem(handle func(o objc.Object, tabView TabView, tabViewItem TabViewItem)) {
	objc.AddMethod(c.class, objc.GetSelector("tabView:didSelectTabViewItem:"), handle)
}

func (c *TabViewDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
