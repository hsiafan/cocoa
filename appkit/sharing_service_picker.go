// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var SharingServicePickerClass = _SharingServicePickerClass{objc.GetClass("NSSharingServicePicker")}

type _SharingServicePickerClass struct {
	objc.Class
}

type ISharingServicePicker interface {
	objc.IObject
	ShowRelativeToRect_OfView_PreferredEdge(rect foundation.Rect, view IView, preferredEdge foundation.RectEdge)
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
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
	rv := objc.CallMethod[SharingServicePicker](s_, objc.SEL("initWithItems:"), items)
	return rv
}

func (sc _SharingServicePickerClass) Alloc() SharingServicePicker {
	rv := objc.CallMethod[SharingServicePicker](sc, objc.SEL("alloc"))
	return rv
}

func (sc _SharingServicePickerClass) New() SharingServicePicker {
	rv := objc.CallMethod[SharingServicePicker](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewSharingServicePicker() SharingServicePicker {
	return SharingServicePickerClass.New()
}

func (s_ SharingServicePicker) Init() SharingServicePicker {
	rv := objc.CallMethod[SharingServicePicker](s_, objc.SEL("init"))
	return rv
}

func (s_ SharingServicePicker) ShowRelativeToRect_OfView_PreferredEdge(rect foundation.Rect, view IView, preferredEdge foundation.RectEdge) {
	objc.CallMethod[objc.Void](s_, objc.SEL("showRelativeToRect:ofView:preferredEdge:"), rect, objc.ExtractPtr(view), preferredEdge)
}

// weak property
func (s_ SharingServicePicker) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.SEL("delegate"))
	return rv
}

// weak property
func (s_ SharingServicePicker) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}

func (s_ SharingServicePicker) StandardShareMenuItem() MenuItem {
	rv := objc.CallMethod[MenuItem](s_, objc.SEL("standardShareMenuItem"))
	return rv
}
