// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var SharingServicePickerClass = _SharingServicePickerClass{objc.GetClass("NSSharingServicePicker")}

type _SharingServicePickerClass struct {
	objc.Class
}

type ISharingServicePicker interface {
	objc.IObject
	ShowRelativeToRect_OfView_PreferredEdge(rect foundation.Rect, view IView, preferredEdge foundation.RectEdge)
	Delegate() SharingServicePickerDelegateWrapper
	SetDelegate(value SharingServicePickerDelegate)
	SetDelegate0(value objc.IObject)
	StandardShareMenuItem() MenuItem
}

type SharingServicePicker struct {
	objc.Object
}

func MakeSharingServicePicker(ptr unsafe.Pointer) SharingServicePicker {
	return SharingServicePicker{
		Object: objc.MakeObject(ptr),
	}
}

func (s_ SharingServicePicker) InitWithItems(items []objc.IObject) SharingServicePicker {
	rv := objc.CallMethod[SharingServicePicker](s_, "initWithItems:", items)
	return rv
}

func (sc _SharingServicePickerClass) Alloc() SharingServicePicker {
	rv := objc.CallMethod[SharingServicePicker](sc, "alloc")
	return rv
}

func (sc _SharingServicePickerClass) New() SharingServicePicker {
	rv := objc.CallMethod[SharingServicePicker](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSharingServicePicker() SharingServicePicker {
	return SharingServicePickerClass.New()
}

func (s_ SharingServicePicker) Init() SharingServicePicker {
	rv := objc.CallMethod[SharingServicePicker](s_, "init")
	return rv
}

func (s_ SharingServicePicker) ShowRelativeToRect_OfView_PreferredEdge(rect foundation.Rect, view IView, preferredEdge foundation.RectEdge) {
	objc.CallMethod[objc.Void](s_, "showRelativeToRect:ofView:preferredEdge:", rect, view, preferredEdge)
}

func (s_ SharingServicePicker) Delegate() SharingServicePickerDelegateWrapper {
	rv := objc.CallMethod[SharingServicePickerDelegateWrapper](s_, "delegate")
	return rv
}

func (s_ SharingServicePicker) SetDelegate(value SharingServicePickerDelegate) {
	po := objc.CreateProtocol("NSSharingServicePickerDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(s_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](s_, "setDelegate:", po)
}

func (s_ SharingServicePicker) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, "setDelegate:", value)
}

func (s_ SharingServicePicker) StandardShareMenuItem() MenuItem {
	rv := objc.CallMethod[MenuItem](s_, "standardShareMenuItem")
	return rv
}
