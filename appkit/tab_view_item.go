// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

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
	rv := objc.CallMethod[TabViewItem](t_, "initWithIdentifier:", identifier)
	return rv
}

func (tc _TabViewItemClass) TabViewItemWithViewController(viewController IViewController) TabViewItem {
	rv := objc.CallMethod[TabViewItem](tc, "tabViewItemWithViewController:", viewController)
	return rv
}

func (tc _TabViewItemClass) Alloc() TabViewItem {
	rv := objc.CallMethod[TabViewItem](tc, "alloc")
	return rv
}

func (tc _TabViewItemClass) New() TabViewItem {
	rv := objc.CallMethod[TabViewItem](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTabViewItem() TabViewItem {
	return TabViewItemClass.New()
}

func (t_ TabViewItem) Init() TabViewItem {
	rv := objc.CallMethod[TabViewItem](t_, "init")
	return rv
}

func (t_ TabViewItem) DrawLabel_InRect(shouldTruncateLabel bool, labelRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, "drawLabel:inRect:", shouldTruncateLabel, labelRect)
}

func (t_ TabViewItem) SizeOfLabel(computeMin bool) foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, "sizeOfLabel:", computeMin)
	return rv
}

func (t_ TabViewItem) Label() string {
	rv := objc.CallMethod[string](t_, "label")
	return rv
}

func (t_ TabViewItem) SetLabel(value string) {
	objc.CallMethod[objc.Void](t_, "setLabel:", value)
}

func (t_ TabViewItem) TabState() TabState {
	rv := objc.CallMethod[TabState](t_, "tabState")
	return rv
}

func (t_ TabViewItem) Identifier() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, "identifier")
	return rv
}

func (t_ TabViewItem) SetIdentifier(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, "setIdentifier:", value)
}

func (t_ TabViewItem) Color() Color {
	rv := objc.CallMethod[Color](t_, "color")
	return rv
}

func (t_ TabViewItem) SetColor(value IColor) {
	objc.CallMethod[objc.Void](t_, "setColor:", value)
}

func (t_ TabViewItem) View() View {
	rv := objc.CallMethod[View](t_, "view")
	return rv
}

func (t_ TabViewItem) SetView(value IView) {
	objc.CallMethod[objc.Void](t_, "setView:", value)
}

func (t_ TabViewItem) InitialFirstResponder() View {
	rv := objc.CallMethod[View](t_, "initialFirstResponder")
	return rv
}

func (t_ TabViewItem) SetInitialFirstResponder(value IView) {
	objc.CallMethod[objc.Void](t_, "setInitialFirstResponder:", value)
}

func (t_ TabViewItem) TabView() TabView {
	rv := objc.CallMethod[TabView](t_, "tabView")
	return rv
}

func (t_ TabViewItem) ToolTip() string {
	rv := objc.CallMethod[string](t_, "toolTip")
	return rv
}

func (t_ TabViewItem) SetToolTip(value string) {
	objc.CallMethod[objc.Void](t_, "setToolTip:", value)
}

func (t_ TabViewItem) Image() Image {
	rv := objc.CallMethod[Image](t_, "image")
	return rv
}

func (t_ TabViewItem) SetImage(value IImage) {
	objc.CallMethod[objc.Void](t_, "setImage:", value)
}

func (t_ TabViewItem) ViewController() ViewController {
	rv := objc.CallMethod[ViewController](t_, "viewController")
	return rv
}

func (t_ TabViewItem) SetViewController(value IViewController) {
	objc.CallMethod[objc.Void](t_, "setViewController:", value)
}
