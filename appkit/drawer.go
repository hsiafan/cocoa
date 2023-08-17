// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var DrawerClass = _DrawerClass{objc.GetClass("NSDrawer")}

type _DrawerClass struct {
	objc.Class
}

type IDrawer interface {
	IResponder
	// deprecated
	Close()
	// deprecated
	Close1(sender objc.IObject)
	// deprecated
	Open()
	// deprecated
	Open1(sender objc.IObject)
	// deprecated
	OpenOnEdge(edge foundation.RectEdge)
	// deprecated
	Toggle(sender objc.IObject)
	// deprecated
	Delegate() DrawerDelegateWrapper
	// deprecated
	SetDelegate(value objc.IObject)
	// deprecated
	State() int
	// deprecated
	ContentSize() foundation.Size
	// deprecated
	SetContentSize(value foundation.Size)
	// deprecated
	LeadingOffset() float64
	// deprecated
	SetLeadingOffset(value float64)
	// deprecated
	MaxContentSize() foundation.Size
	// deprecated
	SetMaxContentSize(value foundation.Size)
	// deprecated
	MinContentSize() foundation.Size
	// deprecated
	SetMinContentSize(value foundation.Size)
	// deprecated
	TrailingOffset() float64
	// deprecated
	SetTrailingOffset(value float64)
	// deprecated
	Edge() foundation.RectEdge
	// deprecated
	PreferredEdge() foundation.RectEdge
	// deprecated
	SetPreferredEdge(value foundation.RectEdge)
	// deprecated
	ContentView() View
	// deprecated
	SetContentView(value IView)
	// deprecated
	ParentWindow() Window
	// deprecated
	SetParentWindow(value IWindow)
}

type Drawer struct {
	Responder
}

func MakeDrawer(ptr unsafe.Pointer) Drawer {
	return Drawer{
		Responder: MakeResponder(ptr),
	}
}

func (d_ Drawer) InitWithContentSize_PreferredEdge(contentSize foundation.Size, edge foundation.RectEdge) Drawer {
	rv := objc.CallMethod[Drawer](d_, objc.GetSelector("initWithContentSize:preferredEdge:"), contentSize, edge)
	return rv
}

func (d_ Drawer) Init() Drawer {
	rv := objc.CallMethod[Drawer](d_, objc.GetSelector("init"))
	return rv
}

func (dc _DrawerClass) Alloc() Drawer {
	rv := objc.CallMethod[Drawer](dc, objc.GetSelector("alloc"))
	return rv
}

func (dc _DrawerClass) New() Drawer {
	rv := objc.CallMethod[Drawer](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDrawer() Drawer {
	return DrawerClass.New()
}

// deprecated
func (d_ Drawer) Close() {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("close"))
}

// deprecated
func (d_ Drawer) Close1(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("close:"), objc.ExtractPtr(sender))
}

// deprecated
func (d_ Drawer) Open() {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("open"))
}

// deprecated
func (d_ Drawer) Open1(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("open:"), objc.ExtractPtr(sender))
}

// deprecated
func (d_ Drawer) OpenOnEdge(edge foundation.RectEdge) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("openOnEdge:"), edge)
}

// deprecated
func (d_ Drawer) Toggle(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("toggle:"), objc.ExtractPtr(sender))
}

// weak property
// deprecated
func (d_ Drawer) Delegate() DrawerDelegateWrapper {
	rv := objc.CallMethod[DrawerDelegateWrapper](d_, objc.GetSelector("delegate"))
	return rv
}

// weak property
// deprecated
func (d_ Drawer) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

// deprecated
func (d_ Drawer) State() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("state"))
	return rv
}

// deprecated
func (d_ Drawer) ContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](d_, objc.GetSelector("contentSize"))
	return rv
}

// deprecated
func (d_ Drawer) SetContentSize(value foundation.Size) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setContentSize:"), value)
}

// deprecated
func (d_ Drawer) LeadingOffset() float64 {
	rv := objc.CallMethod[float64](d_, objc.GetSelector("leadingOffset"))
	return rv
}

// deprecated
func (d_ Drawer) SetLeadingOffset(value float64) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setLeadingOffset:"), value)
}

// deprecated
func (d_ Drawer) MaxContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](d_, objc.GetSelector("maxContentSize"))
	return rv
}

// deprecated
func (d_ Drawer) SetMaxContentSize(value foundation.Size) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setMaxContentSize:"), value)
}

// deprecated
func (d_ Drawer) MinContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](d_, objc.GetSelector("minContentSize"))
	return rv
}

// deprecated
func (d_ Drawer) SetMinContentSize(value foundation.Size) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setMinContentSize:"), value)
}

// deprecated
func (d_ Drawer) TrailingOffset() float64 {
	rv := objc.CallMethod[float64](d_, objc.GetSelector("trailingOffset"))
	return rv
}

// deprecated
func (d_ Drawer) SetTrailingOffset(value float64) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setTrailingOffset:"), value)
}

// deprecated
func (d_ Drawer) Edge() foundation.RectEdge {
	rv := objc.CallMethod[foundation.RectEdge](d_, objc.GetSelector("edge"))
	return rv
}

// deprecated
func (d_ Drawer) PreferredEdge() foundation.RectEdge {
	rv := objc.CallMethod[foundation.RectEdge](d_, objc.GetSelector("preferredEdge"))
	return rv
}

// deprecated
func (d_ Drawer) SetPreferredEdge(value foundation.RectEdge) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setPreferredEdge:"), value)
}

// deprecated
func (d_ Drawer) ContentView() View {
	rv := objc.CallMethod[View](d_, objc.GetSelector("contentView"))
	return rv
}

// deprecated
func (d_ Drawer) SetContentView(value IView) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setContentView:"), objc.ExtractPtr(value))
}

// weak property
// deprecated
func (d_ Drawer) ParentWindow() Window {
	rv := objc.CallMethod[Window](d_, objc.GetSelector("parentWindow"))
	return rv
}

// weak property
// deprecated
func (d_ Drawer) SetParentWindow(value IWindow) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setParentWindow:"), objc.ExtractPtr(value))
}
