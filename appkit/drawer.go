// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
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
	SetDelegate(value DrawerDelegate)
	// deprecated
	SetDelegate0(value objc.IObject)
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
	rv := ffi.CallMethod[Drawer](d_, "initWithContentSize:preferredEdge:", contentSize, edge)
	return rv
}

func (d_ Drawer) Init() Drawer {
	rv := ffi.CallMethod[Drawer](d_, "init")
	return rv
}

func (dc _DrawerClass) Alloc() Drawer {
	rv := ffi.CallMethod[Drawer](dc, "alloc")
	return rv
}

func (dc _DrawerClass) New() Drawer {
	rv := ffi.CallMethod[Drawer](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDrawer() Drawer {
	return DrawerClass.New()
}

// deprecated
func (d_ Drawer) Close() {
	ffi.CallMethod[ffi.Void](d_, "close")
}

// deprecated
func (d_ Drawer) Close1(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "close:", sender)
}

// deprecated
func (d_ Drawer) Open() {
	ffi.CallMethod[ffi.Void](d_, "open")
}

// deprecated
func (d_ Drawer) Open1(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "open:", sender)
}

// deprecated
func (d_ Drawer) OpenOnEdge(edge foundation.RectEdge) {
	ffi.CallMethod[ffi.Void](d_, "openOnEdge:", edge)
}

// deprecated
func (d_ Drawer) Toggle(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "toggle:", sender)
}

// deprecated
func (d_ Drawer) Delegate() DrawerDelegateWrapper {
	rv := ffi.CallMethod[DrawerDelegateWrapper](d_, "delegate")
	return rv
}

// deprecated
func (d_ Drawer) SetDelegate(value DrawerDelegate) {
	po := ffi.CreateProtocol("NSDrawerDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(d_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](d_, "setDelegate:", po)
}

// deprecated
func (d_ Drawer) SetDelegate0(value objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "setDelegate:", value)
}

// deprecated
func (d_ Drawer) State() int {
	rv := ffi.CallMethod[int](d_, "state")
	return rv
}

// deprecated
func (d_ Drawer) ContentSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](d_, "contentSize")
	return rv
}

// deprecated
func (d_ Drawer) SetContentSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](d_, "setContentSize:", value)
}

// deprecated
func (d_ Drawer) LeadingOffset() float64 {
	rv := ffi.CallMethod[float64](d_, "leadingOffset")
	return rv
}

// deprecated
func (d_ Drawer) SetLeadingOffset(value float64) {
	ffi.CallMethod[ffi.Void](d_, "setLeadingOffset:", value)
}

// deprecated
func (d_ Drawer) MaxContentSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](d_, "maxContentSize")
	return rv
}

// deprecated
func (d_ Drawer) SetMaxContentSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](d_, "setMaxContentSize:", value)
}

// deprecated
func (d_ Drawer) MinContentSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](d_, "minContentSize")
	return rv
}

// deprecated
func (d_ Drawer) SetMinContentSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](d_, "setMinContentSize:", value)
}

// deprecated
func (d_ Drawer) TrailingOffset() float64 {
	rv := ffi.CallMethod[float64](d_, "trailingOffset")
	return rv
}

// deprecated
func (d_ Drawer) SetTrailingOffset(value float64) {
	ffi.CallMethod[ffi.Void](d_, "setTrailingOffset:", value)
}

// deprecated
func (d_ Drawer) Edge() foundation.RectEdge {
	rv := ffi.CallMethod[foundation.RectEdge](d_, "edge")
	return rv
}

// deprecated
func (d_ Drawer) PreferredEdge() foundation.RectEdge {
	rv := ffi.CallMethod[foundation.RectEdge](d_, "preferredEdge")
	return rv
}

// deprecated
func (d_ Drawer) SetPreferredEdge(value foundation.RectEdge) {
	ffi.CallMethod[ffi.Void](d_, "setPreferredEdge:", value)
}

// deprecated
func (d_ Drawer) ContentView() View {
	rv := ffi.CallMethod[View](d_, "contentView")
	return rv
}

// deprecated
func (d_ Drawer) SetContentView(value IView) {
	ffi.CallMethod[ffi.Void](d_, "setContentView:", value)
}

// deprecated
func (d_ Drawer) ParentWindow() Window {
	rv := ffi.CallMethod[Window](d_, "parentWindow")
	return rv
}

// deprecated
func (d_ Drawer) SetParentWindow(value IWindow) {
	ffi.CallMethod[ffi.Void](d_, "setParentWindow:", value)
}
