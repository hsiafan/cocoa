// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var TouchBarItemClass = _TouchBarItemClass{objc.GetClass("NSTouchBarItem")}

type _TouchBarItemClass struct {
	objc.Class
}

type ITouchBarItem interface {
	objc.IObject
	Identifier() TouchBarItemIdentifier
	VisibilityPriority() TouchBarItemPriority
	SetVisibilityPriority(value TouchBarItemPriority)
	IsVisible() bool
	CustomizationLabel() string
	ViewController() ViewController
	View() View
}

type TouchBarItem struct {
	objc.Object
}

func MakeTouchBarItem(ptr unsafe.Pointer) TouchBarItem {
	return TouchBarItem{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) TouchBarItem {
	rv := objc.CallMethod[TouchBarItem](t_, "initWithIdentifier:", identifier)
	return rv
}

func (tc _TouchBarItemClass) Alloc() TouchBarItem {
	rv := objc.CallMethod[TouchBarItem](tc, "alloc")
	return rv
}

func (tc _TouchBarItemClass) New() TouchBarItem {
	rv := objc.CallMethod[TouchBarItem](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTouchBarItem() TouchBarItem {
	return TouchBarItemClass.New()
}

func (t_ TouchBarItem) Init() TouchBarItem {
	rv := objc.CallMethod[TouchBarItem](t_, "init")
	return rv
}

func (t_ TouchBarItem) Identifier() TouchBarItemIdentifier {
	rv := objc.CallMethod[TouchBarItemIdentifier](t_, "identifier")
	return rv
}

func (t_ TouchBarItem) VisibilityPriority() TouchBarItemPriority {
	rv := objc.CallMethod[TouchBarItemPriority](t_, "visibilityPriority")
	return rv
}

func (t_ TouchBarItem) SetVisibilityPriority(value TouchBarItemPriority) {
	objc.CallMethod[objc.Void](t_, "setVisibilityPriority:", value)
}

func (t_ TouchBarItem) IsVisible() bool {
	rv := objc.CallMethod[bool](t_, "isVisible")
	return rv
}

func (t_ TouchBarItem) CustomizationLabel() string {
	rv := objc.CallMethod[string](t_, "customizationLabel")
	return rv
}

func (t_ TouchBarItem) ViewController() ViewController {
	rv := objc.CallMethod[ViewController](t_, "viewController")
	return rv
}

func (t_ TouchBarItem) View() View {
	rv := objc.CallMethod[View](t_, "view")
	return rv
}
