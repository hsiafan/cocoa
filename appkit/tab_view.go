// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[TabView](t_, "initWithFrame:", frameRect)
	return rv
}

func (t_ TabView) Init() TabView {
	rv := ffi.CallMethod[TabView](t_, "init")
	return rv
}

func (tc _TabViewClass) Alloc() TabView {
	rv := ffi.CallMethod[TabView](tc, "alloc")
	return rv
}

func (tc _TabViewClass) New() TabView {
	rv := ffi.CallMethod[TabView](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTabView() TabView {
	return TabViewClass.New()
}

func (t_ TabView) AddTabViewItem(tabViewItem ITabViewItem) {
	ffi.CallMethod[ffi.Void](t_, "addTabViewItem:", tabViewItem)
}

func (t_ TabView) InsertTabViewItem_AtIndex(tabViewItem ITabViewItem, index int) {
	ffi.CallMethod[ffi.Void](t_, "insertTabViewItem:atIndex:", tabViewItem, index)
}

func (t_ TabView) RemoveTabViewItem(tabViewItem ITabViewItem) {
	ffi.CallMethod[ffi.Void](t_, "removeTabViewItem:", tabViewItem)
}

func (t_ TabView) IndexOfTabViewItem(tabViewItem ITabViewItem) int {
	rv := ffi.CallMethod[int](t_, "indexOfTabViewItem:", tabViewItem)
	return rv
}

func (t_ TabView) IndexOfTabViewItemWithIdentifier(identifier objc.IObject) int {
	rv := ffi.CallMethod[int](t_, "indexOfTabViewItemWithIdentifier:", identifier)
	return rv
}

func (t_ TabView) TabViewItemAtIndex(index int) TabViewItem {
	rv := ffi.CallMethod[TabViewItem](t_, "tabViewItemAtIndex:", index)
	return rv
}

func (t_ TabView) SelectFirstTabViewItem(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "selectFirstTabViewItem:", sender)
}

func (t_ TabView) SelectLastTabViewItem(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "selectLastTabViewItem:", sender)
}

func (t_ TabView) SelectNextTabViewItem(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "selectNextTabViewItem:", sender)
}

func (t_ TabView) SelectPreviousTabViewItem(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "selectPreviousTabViewItem:", sender)
}

func (t_ TabView) SelectTabViewItem(tabViewItem ITabViewItem) {
	ffi.CallMethod[ffi.Void](t_, "selectTabViewItem:", tabViewItem)
}

func (t_ TabView) SelectTabViewItemAtIndex(index int) {
	ffi.CallMethod[ffi.Void](t_, "selectTabViewItemAtIndex:", index)
}

func (t_ TabView) SelectTabViewItemWithIdentifier(identifier objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "selectTabViewItemWithIdentifier:", identifier)
}

func (t_ TabView) TakeSelectedTabViewItemFromSender(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "takeSelectedTabViewItemFromSender:", sender)
}

func (t_ TabView) TabViewItemAtPoint(point foundation.Point) TabViewItem {
	rv := ffi.CallMethod[TabViewItem](t_, "tabViewItemAtPoint:", point)
	return rv
}

func (t_ TabView) Delegate() TabViewDelegateWrapper {
	rv := ffi.CallMethod[TabViewDelegateWrapper](t_, "delegate")
	return rv
}

func (t_ TabView) SetDelegate(value TabViewDelegate) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](t_, "setDelegate:", po)
}

func (t_ TabView) NumberOfTabViewItems() int {
	rv := ffi.CallMethod[int](t_, "numberOfTabViewItems")
	return rv
}

func (t_ TabView) TabViewItems() []TabViewItem {
	rv := ffi.CallMethod[[]TabViewItem](t_, "tabViewItems")
	return rv
}

func (t_ TabView) SetTabViewItems(value []ITabViewItem) {
	ffi.CallMethod[ffi.Void](t_, "setTabViewItems:", value)
}

func (t_ TabView) TabViewType() TabViewType {
	rv := ffi.CallMethod[TabViewType](t_, "tabViewType")
	return rv
}

func (t_ TabView) SetTabViewType(value TabViewType) {
	ffi.CallMethod[ffi.Void](t_, "setTabViewType:", value)
}

func (t_ TabView) TabPosition() TabPosition {
	rv := ffi.CallMethod[TabPosition](t_, "tabPosition")
	return rv
}

func (t_ TabView) SetTabPosition(value TabPosition) {
	ffi.CallMethod[ffi.Void](t_, "setTabPosition:", value)
}

func (t_ TabView) TabViewBorderType() TabViewBorderType {
	rv := ffi.CallMethod[TabViewBorderType](t_, "tabViewBorderType")
	return rv
}

func (t_ TabView) SetTabViewBorderType(value TabViewBorderType) {
	ffi.CallMethod[ffi.Void](t_, "setTabViewBorderType:", value)
}

func (t_ TabView) SelectedTabViewItem() TabViewItem {
	rv := ffi.CallMethod[TabViewItem](t_, "selectedTabViewItem")
	return rv
}

func (t_ TabView) Font() Font {
	rv := ffi.CallMethod[Font](t_, "font")
	return rv
}

func (t_ TabView) SetFont(value IFont) {
	ffi.CallMethod[ffi.Void](t_, "setFont:", value)
}

// deprecated
func (t_ TabView) ControlTint() ControlTint {
	rv := ffi.CallMethod[ControlTint](t_, "controlTint")
	return rv
}

// deprecated
func (t_ TabView) SetControlTint(value ControlTint) {
	ffi.CallMethod[ffi.Void](t_, "setControlTint:", value)
}

func (t_ TabView) DrawsBackground() bool {
	rv := ffi.CallMethod[bool](t_, "drawsBackground")
	return rv
}

func (t_ TabView) SetDrawsBackground(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setDrawsBackground:", value)
}

func (t_ TabView) MinimumSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](t_, "minimumSize")
	return rv
}

func (t_ TabView) ContentRect() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](t_, "contentRect")
	return rv
}

func (t_ TabView) ControlSize() ControlSize {
	rv := ffi.CallMethod[ControlSize](t_, "controlSize")
	return rv
}

func (t_ TabView) SetControlSize(value ControlSize) {
	ffi.CallMethod[ffi.Void](t_, "setControlSize:", value)
}

func (t_ TabView) AllowsTruncatedLabels() bool {
	rv := ffi.CallMethod[bool](t_, "allowsTruncatedLabels")
	return rv
}

func (t_ TabView) SetAllowsTruncatedLabels(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAllowsTruncatedLabels:", value)
}

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

var TabViewItemClass = _TabViewItemClass{objc.GetClass("NSTabViewItem")}

type _TabViewItemClass struct {
	objc.Class
}

type ITabViewItem interface {
	objc.IObject
	DrawLabel_InRect(shouldTruncateLabel bool, labelRect foundation.Rect)
	SizeOfLabel(computeMin bool) foundation.Size
	Label() string
	SetLabel(value string)
	TabState() TabState
	Identifier() objc.Object
	SetIdentifier(value objc.IObject)
	Color() Color
	SetColor(value IColor)
	View() View
	SetView(value IView)
	InitialFirstResponder() View
	SetInitialFirstResponder(value IView)
	TabView() TabView
	ToolTip() string
	SetToolTip(value string)
	Image() Image
	SetImage(value IImage)
	ViewController() ViewController
	SetViewController(value IViewController)
}

type TabViewItem struct {
	objc.Object
}

func MakeTabViewItem(ptr unsafe.Pointer) TabViewItem {
	return TabViewItem{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TabViewItem) InitWithIdentifier(identifier objc.IObject) TabViewItem {
	rv := ffi.CallMethod[TabViewItem](t_, "initWithIdentifier:", identifier)
	return rv
}

func (tc _TabViewItemClass) TabViewItemWithViewController(viewController IViewController) TabViewItem {
	rv := ffi.CallMethod[TabViewItem](tc, "tabViewItemWithViewController:", viewController)
	return rv
}

func (tc _TabViewItemClass) Alloc() TabViewItem {
	rv := ffi.CallMethod[TabViewItem](tc, "alloc")
	return rv
}

func (t_ TabViewItem) Init() TabViewItem {
	rv := ffi.CallMethod[TabViewItem](t_, "init")
	return rv
}

func (tc _TabViewItemClass) New() TabViewItem {
	rv := ffi.CallMethod[TabViewItem](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTabViewItem() TabViewItem {
	return TabViewItemClass.New()
}

func (t_ TabViewItem) DrawLabel_InRect(shouldTruncateLabel bool, labelRect foundation.Rect) {
	ffi.CallMethod[ffi.Void](t_, "drawLabel:inRect:", shouldTruncateLabel, labelRect)
}

func (t_ TabViewItem) SizeOfLabel(computeMin bool) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](t_, "sizeOfLabel:", computeMin)
	return rv
}

func (t_ TabViewItem) Label() string {
	rv := ffi.CallMethod[string](t_, "label")
	return rv
}

func (t_ TabViewItem) SetLabel(value string) {
	ffi.CallMethod[ffi.Void](t_, "setLabel:", value)
}

func (t_ TabViewItem) TabState() TabState {
	rv := ffi.CallMethod[TabState](t_, "tabState")
	return rv
}

func (t_ TabViewItem) Identifier() objc.Object {
	rv := ffi.CallMethod[objc.Object](t_, "identifier")
	return rv
}

func (t_ TabViewItem) SetIdentifier(value objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "setIdentifier:", value)
}

func (t_ TabViewItem) Color() Color {
	rv := ffi.CallMethod[Color](t_, "color")
	return rv
}

func (t_ TabViewItem) SetColor(value IColor) {
	ffi.CallMethod[ffi.Void](t_, "setColor:", value)
}

func (t_ TabViewItem) View() View {
	rv := ffi.CallMethod[View](t_, "view")
	return rv
}

func (t_ TabViewItem) SetView(value IView) {
	ffi.CallMethod[ffi.Void](t_, "setView:", value)
}

func (t_ TabViewItem) InitialFirstResponder() View {
	rv := ffi.CallMethod[View](t_, "initialFirstResponder")
	return rv
}

func (t_ TabViewItem) SetInitialFirstResponder(value IView) {
	ffi.CallMethod[ffi.Void](t_, "setInitialFirstResponder:", value)
}

func (t_ TabViewItem) TabView() TabView {
	rv := ffi.CallMethod[TabView](t_, "tabView")
	return rv
}

func (t_ TabViewItem) ToolTip() string {
	rv := ffi.CallMethod[string](t_, "toolTip")
	return rv
}

func (t_ TabViewItem) SetToolTip(value string) {
	ffi.CallMethod[ffi.Void](t_, "setToolTip:", value)
}

func (t_ TabViewItem) Image() Image {
	rv := ffi.CallMethod[Image](t_, "image")
	return rv
}

func (t_ TabViewItem) SetImage(value IImage) {
	ffi.CallMethod[ffi.Void](t_, "setImage:", value)
}

func (t_ TabViewItem) ViewController() ViewController {
	rv := ffi.CallMethod[ViewController](t_, "viewController")
	return rv
}

func (t_ TabViewItem) SetViewController(value IViewController) {
	ffi.CallMethod[ffi.Void](t_, "setViewController:", value)
}
