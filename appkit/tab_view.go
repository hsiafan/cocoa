// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var TabViewClass = _TabViewClass{objc.GetClass("NSTabView")}

type _TabViewClass struct {
	objc.Class
}

type ITabView interface {
	IView
	AddTabViewItem(tabViewItem ITabViewItem)
	InsertTabViewItem_AtIndex(tabViewItem ITabViewItem, index int)
	RemoveTabViewItem(tabViewItem ITabViewItem)
	IndexOfTabViewItem(tabViewItem ITabViewItem) int
	IndexOfTabViewItemWithIdentifier(identifier objc.IObject) int
	TabViewItemAtIndex(index int) TabViewItem
	SelectFirstTabViewItem(sender objc.IObject)
	SelectLastTabViewItem(sender objc.IObject)
	SelectNextTabViewItem(sender objc.IObject)
	SelectPreviousTabViewItem(sender objc.IObject)
	SelectTabViewItem(tabViewItem ITabViewItem)
	SelectTabViewItemAtIndex(index int)
	SelectTabViewItemWithIdentifier(identifier objc.IObject)
	TakeSelectedTabViewItemFromSender(sender objc.IObject)
	TabViewItemAtPoint(point foundation.Point) TabViewItem
	Delegate() TabViewDelegateWrapper
	SetDelegate(value TabViewDelegate)
	SetDelegate0(value objc.IObject)
	NumberOfTabViewItems() int
	TabViewItems() []TabViewItem
	SetTabViewItems(value []ITabViewItem)
	TabViewType() TabViewType
	SetTabViewType(value TabViewType)
	TabPosition() TabPosition
	SetTabPosition(value TabPosition)
	TabViewBorderType() TabViewBorderType
	SetTabViewBorderType(value TabViewBorderType)
	SelectedTabViewItem() TabViewItem
	Font() Font
	SetFont(value IFont)
	// deprecated
	ControlTint() ControlTint
	// deprecated
	SetControlTint(value ControlTint)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	MinimumSize() foundation.Size
	ContentRect() foundation.Rect
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	AllowsTruncatedLabels() bool
	SetAllowsTruncatedLabels(value bool)
}

type TabView struct {
	View
}

func MakeTabView(ptr unsafe.Pointer) TabView {
	return TabView{
		View: MakeView(ptr),
	}
}

func (t_ TabView) InitWithFrame(frameRect foundation.Rect) TabView {
	rv := objc.CallMethod[TabView](t_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func (t_ TabView) Init() TabView {
	rv := objc.CallMethod[TabView](t_, objc.GetSelector("init"))
	return rv
}

func (tc _TabViewClass) Alloc() TabView {
	rv := objc.CallMethod[TabView](tc, objc.GetSelector("alloc"))
	return rv
}

func (tc _TabViewClass) New() TabView {
	rv := objc.CallMethod[TabView](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTabView() TabView {
	return TabViewClass.New()
}

func (t_ TabView) AddTabViewItem(tabViewItem ITabViewItem) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("addTabViewItem:"), tabViewItem)
}

func (t_ TabView) InsertTabViewItem_AtIndex(tabViewItem ITabViewItem, index int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("insertTabViewItem:atIndex:"), tabViewItem, index)
}

func (t_ TabView) RemoveTabViewItem(tabViewItem ITabViewItem) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("removeTabViewItem:"), tabViewItem)
}

func (t_ TabView) IndexOfTabViewItem(tabViewItem ITabViewItem) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("indexOfTabViewItem:"), tabViewItem)
	return rv
}

func (t_ TabView) IndexOfTabViewItemWithIdentifier(identifier objc.IObject) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("indexOfTabViewItemWithIdentifier:"), identifier)
	return rv
}

func (t_ TabView) TabViewItemAtIndex(index int) TabViewItem {
	rv := objc.CallMethod[TabViewItem](t_, objc.GetSelector("tabViewItemAtIndex:"), index)
	return rv
}

func (t_ TabView) SelectFirstTabViewItem(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("selectFirstTabViewItem:"), sender)
}

func (t_ TabView) SelectLastTabViewItem(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("selectLastTabViewItem:"), sender)
}

func (t_ TabView) SelectNextTabViewItem(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("selectNextTabViewItem:"), sender)
}

func (t_ TabView) SelectPreviousTabViewItem(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("selectPreviousTabViewItem:"), sender)
}

func (t_ TabView) SelectTabViewItem(tabViewItem ITabViewItem) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("selectTabViewItem:"), tabViewItem)
}

func (t_ TabView) SelectTabViewItemAtIndex(index int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("selectTabViewItemAtIndex:"), index)
}

func (t_ TabView) SelectTabViewItemWithIdentifier(identifier objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("selectTabViewItemWithIdentifier:"), identifier)
}

func (t_ TabView) TakeSelectedTabViewItemFromSender(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("takeSelectedTabViewItemFromSender:"), sender)
}

func (t_ TabView) TabViewItemAtPoint(point foundation.Point) TabViewItem {
	rv := objc.CallMethod[TabViewItem](t_, objc.GetSelector("tabViewItemAtPoint:"), point)
	return rv
}

func (t_ TabView) Delegate() TabViewDelegateWrapper {
	rv := objc.CallMethod[TabViewDelegateWrapper](t_, objc.GetSelector("delegate"))
	return rv
}

func (t_ TabView) SetDelegate(value TabViewDelegate) {
	po := objc.CreateProtocol("NSTabViewDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDelegate:"), po)
}

func (t_ TabView) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDelegate:"), value)
}

func (t_ TabView) NumberOfTabViewItems() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("numberOfTabViewItems"))
	return rv
}

func (t_ TabView) TabViewItems() []TabViewItem {
	rv := objc.CallMethod[[]TabViewItem](t_, objc.GetSelector("tabViewItems"))
	return rv
}

func (t_ TabView) SetTabViewItems(value []ITabViewItem) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTabViewItems:"), value)
}

func (t_ TabView) TabViewType() TabViewType {
	rv := objc.CallMethod[TabViewType](t_, objc.GetSelector("tabViewType"))
	return rv
}

func (t_ TabView) SetTabViewType(value TabViewType) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTabViewType:"), value)
}

func (t_ TabView) TabPosition() TabPosition {
	rv := objc.CallMethod[TabPosition](t_, objc.GetSelector("tabPosition"))
	return rv
}

func (t_ TabView) SetTabPosition(value TabPosition) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTabPosition:"), value)
}

func (t_ TabView) TabViewBorderType() TabViewBorderType {
	rv := objc.CallMethod[TabViewBorderType](t_, objc.GetSelector("tabViewBorderType"))
	return rv
}

func (t_ TabView) SetTabViewBorderType(value TabViewBorderType) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTabViewBorderType:"), value)
}

func (t_ TabView) SelectedTabViewItem() TabViewItem {
	rv := objc.CallMethod[TabViewItem](t_, objc.GetSelector("selectedTabViewItem"))
	return rv
}

func (t_ TabView) Font() Font {
	rv := objc.CallMethod[Font](t_, objc.GetSelector("font"))
	return rv
}

func (t_ TabView) SetFont(value IFont) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setFont:"), value)
}

// deprecated
func (t_ TabView) ControlTint() ControlTint {
	rv := objc.CallMethod[ControlTint](t_, objc.GetSelector("controlTint"))
	return rv
}

// deprecated
func (t_ TabView) SetControlTint(value ControlTint) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setControlTint:"), value)
}

func (t_ TabView) DrawsBackground() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("drawsBackground"))
	return rv
}

func (t_ TabView) SetDrawsBackground(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDrawsBackground:"), value)
}

func (t_ TabView) MinimumSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.GetSelector("minimumSize"))
	return rv
}

func (t_ TabView) ContentRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.GetSelector("contentRect"))
	return rv
}

func (t_ TabView) ControlSize() ControlSize {
	rv := objc.CallMethod[ControlSize](t_, objc.GetSelector("controlSize"))
	return rv
}

func (t_ TabView) SetControlSize(value ControlSize) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setControlSize:"), value)
}

func (t_ TabView) AllowsTruncatedLabels() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsTruncatedLabels"))
	return rv
}

func (t_ TabView) SetAllowsTruncatedLabels(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsTruncatedLabels:"), value)
}
