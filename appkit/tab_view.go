// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
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
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
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
	rv := objc.CallMethod[TabView](t_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (t_ TabView) Init() TabView {
	rv := objc.CallMethod[TabView](t_, objc.SEL("init"))
	return rv
}

func (tc _TabViewClass) Alloc() TabView {
	rv := objc.CallMethod[TabView](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TabViewClass) New() TabView {
	rv := objc.CallMethod[TabView](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTabView() TabView {
	return TabViewClass.New()
}

func (t_ TabView) AddTabViewItem(tabViewItem ITabViewItem) {
	objc.CallMethod[objc.Void](t_, objc.SEL("addTabViewItem:"), objc.ExtractPtr(tabViewItem))
}

func (t_ TabView) InsertTabViewItem_AtIndex(tabViewItem ITabViewItem, index int) {
	objc.CallMethod[objc.Void](t_, objc.SEL("insertTabViewItem:atIndex:"), objc.ExtractPtr(tabViewItem), index)
}

func (t_ TabView) RemoveTabViewItem(tabViewItem ITabViewItem) {
	objc.CallMethod[objc.Void](t_, objc.SEL("removeTabViewItem:"), objc.ExtractPtr(tabViewItem))
}

func (t_ TabView) IndexOfTabViewItem(tabViewItem ITabViewItem) int {
	rv := objc.CallMethod[int](t_, objc.SEL("indexOfTabViewItem:"), objc.ExtractPtr(tabViewItem))
	return rv
}

func (t_ TabView) IndexOfTabViewItemWithIdentifier(identifier objc.IObject) int {
	rv := objc.CallMethod[int](t_, objc.SEL("indexOfTabViewItemWithIdentifier:"), objc.ExtractPtr(identifier))
	return rv
}

func (t_ TabView) TabViewItemAtIndex(index int) TabViewItem {
	rv := objc.CallMethod[TabViewItem](t_, objc.SEL("tabViewItemAtIndex:"), index)
	return rv
}

func (t_ TabView) SelectFirstTabViewItem(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("selectFirstTabViewItem:"), objc.ExtractPtr(sender))
}

func (t_ TabView) SelectLastTabViewItem(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("selectLastTabViewItem:"), objc.ExtractPtr(sender))
}

func (t_ TabView) SelectNextTabViewItem(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("selectNextTabViewItem:"), objc.ExtractPtr(sender))
}

func (t_ TabView) SelectPreviousTabViewItem(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("selectPreviousTabViewItem:"), objc.ExtractPtr(sender))
}

func (t_ TabView) SelectTabViewItem(tabViewItem ITabViewItem) {
	objc.CallMethod[objc.Void](t_, objc.SEL("selectTabViewItem:"), objc.ExtractPtr(tabViewItem))
}

func (t_ TabView) SelectTabViewItemAtIndex(index int) {
	objc.CallMethod[objc.Void](t_, objc.SEL("selectTabViewItemAtIndex:"), index)
}

func (t_ TabView) SelectTabViewItemWithIdentifier(identifier objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("selectTabViewItemWithIdentifier:"), objc.ExtractPtr(identifier))
}

func (t_ TabView) TakeSelectedTabViewItemFromSender(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("takeSelectedTabViewItemFromSender:"), objc.ExtractPtr(sender))
}

func (t_ TabView) TabViewItemAtPoint(point foundation.Point) TabViewItem {
	rv := objc.CallMethod[TabViewItem](t_, objc.SEL("tabViewItemAtPoint:"), point)
	return rv
}

// weak property
func (t_ TabView) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.SEL("delegate"))
	return rv
}

// weak property
func (t_ TabView) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}

func (t_ TabView) NumberOfTabViewItems() int {
	rv := objc.CallMethod[int](t_, objc.SEL("numberOfTabViewItems"))
	return rv
}

func (t_ TabView) TabViewItems() []TabViewItem {
	rv := objc.CallMethod[[]TabViewItem](t_, objc.SEL("tabViewItems"))
	return rv
}

func (t_ TabView) SetTabViewItems(value []ITabViewItem) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTabViewItems:"), value)
}

func (t_ TabView) TabViewType() TabViewType {
	rv := objc.CallMethod[TabViewType](t_, objc.SEL("tabViewType"))
	return rv
}

func (t_ TabView) SetTabViewType(value TabViewType) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTabViewType:"), value)
}

func (t_ TabView) TabPosition() TabPosition {
	rv := objc.CallMethod[TabPosition](t_, objc.SEL("tabPosition"))
	return rv
}

func (t_ TabView) SetTabPosition(value TabPosition) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTabPosition:"), value)
}

func (t_ TabView) TabViewBorderType() TabViewBorderType {
	rv := objc.CallMethod[TabViewBorderType](t_, objc.SEL("tabViewBorderType"))
	return rv
}

func (t_ TabView) SetTabViewBorderType(value TabViewBorderType) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTabViewBorderType:"), value)
}

func (t_ TabView) SelectedTabViewItem() TabViewItem {
	rv := objc.CallMethod[TabViewItem](t_, objc.SEL("selectedTabViewItem"))
	return rv
}

func (t_ TabView) Font() Font {
	rv := objc.CallMethod[Font](t_, objc.SEL("font"))
	return rv
}

func (t_ TabView) SetFont(value IFont) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setFont:"), objc.ExtractPtr(value))
}

// deprecated
func (t_ TabView) ControlTint() ControlTint {
	rv := objc.CallMethod[ControlTint](t_, objc.SEL("controlTint"))
	return rv
}

// deprecated
func (t_ TabView) SetControlTint(value ControlTint) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setControlTint:"), value)
}

func (t_ TabView) DrawsBackground() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("drawsBackground"))
	return rv
}

func (t_ TabView) SetDrawsBackground(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDrawsBackground:"), value)
}

func (t_ TabView) MinimumSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.SEL("minimumSize"))
	return rv
}

func (t_ TabView) ContentRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.SEL("contentRect"))
	return rv
}

func (t_ TabView) ControlSize() ControlSize {
	rv := objc.CallMethod[ControlSize](t_, objc.SEL("controlSize"))
	return rv
}

func (t_ TabView) SetControlSize(value ControlSize) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setControlSize:"), value)
}

func (t_ TabView) AllowsTruncatedLabels() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("allowsTruncatedLabels"))
	return rv
}

func (t_ TabView) SetAllowsTruncatedLabels(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAllowsTruncatedLabels:"), value)
}
