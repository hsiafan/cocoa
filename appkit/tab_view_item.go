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
	rv := objc.CallMethod[TabViewItem](t_, objc.SEL("initWithIdentifier:"), objc.ExtractPtr(identifier))
	return rv
}

func (tc _TabViewItemClass) TabViewItemWithViewController(viewController IViewController) TabViewItem {
	rv := objc.CallMethod[TabViewItem](tc, objc.SEL("tabViewItemWithViewController:"), objc.ExtractPtr(viewController))
	return rv
}

func (tc _TabViewItemClass) Alloc() TabViewItem {
	rv := objc.CallMethod[TabViewItem](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TabViewItemClass) New() TabViewItem {
	rv := objc.CallMethod[TabViewItem](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTabViewItem() TabViewItem {
	return TabViewItemClass.New()
}

func (t_ TabViewItem) Init() TabViewItem {
	rv := objc.CallMethod[TabViewItem](t_, objc.SEL("init"))
	return rv
}

func (t_ TabViewItem) DrawLabel_InRect(shouldTruncateLabel bool, labelRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.SEL("drawLabel:inRect:"), shouldTruncateLabel, labelRect)
}

func (t_ TabViewItem) SizeOfLabel(computeMin bool) foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.SEL("sizeOfLabel:"), computeMin)
	return rv
}

func (t_ TabViewItem) Label() string {
	rv := objc.CallMethod[string](t_, objc.SEL("label"))
	return rv
}

func (t_ TabViewItem) SetLabel(value string) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setLabel:"), value)
}

func (t_ TabViewItem) TabState() TabState {
	rv := objc.CallMethod[TabState](t_, objc.SEL("tabState"))
	return rv
}

func (t_ TabViewItem) Identifier() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.SEL("identifier"))
	return rv
}

func (t_ TabViewItem) SetIdentifier(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setIdentifier:"), objc.ExtractPtr(value))
}

func (t_ TabViewItem) Color() Color {
	rv := objc.CallMethod[Color](t_, objc.SEL("color"))
	return rv
}

func (t_ TabViewItem) SetColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setColor:"), objc.ExtractPtr(value))
}

func (t_ TabViewItem) View() View {
	rv := objc.CallMethod[View](t_, objc.SEL("view"))
	return rv
}

func (t_ TabViewItem) SetView(value IView) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setView:"), objc.ExtractPtr(value))
}

// weak property
func (t_ TabViewItem) InitialFirstResponder() View {
	rv := objc.CallMethod[View](t_, objc.SEL("initialFirstResponder"))
	return rv
}

// weak property
func (t_ TabViewItem) SetInitialFirstResponder(value IView) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setInitialFirstResponder:"), objc.ExtractPtr(value))
}

func (t_ TabViewItem) TabView() TabView {
	rv := objc.CallMethod[TabView](t_, objc.SEL("tabView"))
	return rv
}

func (t_ TabViewItem) ToolTip() string {
	rv := objc.CallMethod[string](t_, objc.SEL("toolTip"))
	return rv
}

func (t_ TabViewItem) SetToolTip(value string) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setToolTip:"), value)
}

func (t_ TabViewItem) Image() Image {
	rv := objc.CallMethod[Image](t_, objc.SEL("image"))
	return rv
}

func (t_ TabViewItem) SetImage(value IImage) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setImage:"), objc.ExtractPtr(value))
}

func (t_ TabViewItem) ViewController() ViewController {
	rv := objc.CallMethod[ViewController](t_, objc.SEL("viewController"))
	return rv
}

func (t_ TabViewItem) SetViewController(value IViewController) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setViewController:"), objc.ExtractPtr(value))
}
