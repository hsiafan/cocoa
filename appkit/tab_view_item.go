// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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

func (tc _TabViewItemClass) New() TabViewItem {
	rv := ffi.CallMethod[TabViewItem](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTabViewItem() TabViewItem {
	return TabViewItemClass.New()
}

func (t_ TabViewItem) Init() TabViewItem {
	rv := ffi.CallMethod[TabViewItem](t_, "init")
	return rv
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
