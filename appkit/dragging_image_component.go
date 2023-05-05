// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var DraggingImageComponentClass = _DraggingImageComponentClass{objc.GetClass("NSDraggingImageComponent")}

type _DraggingImageComponentClass struct {
	objc.Class
}

type IDraggingImageComponent interface {
	objc.IObject
	Key() DraggingImageComponentKey
	SetKey(value DraggingImageComponentKey)
	Contents() objc.Object
	SetContents(value objc.IObject)
	Frame() foundation.Rect
	SetFrame(value foundation.Rect)
}

type DraggingImageComponent struct {
	objc.Object
}

func MakeDraggingImageComponent(ptr unsafe.Pointer) DraggingImageComponent {
	return DraggingImageComponent{
		Object: objc.MakeObject(ptr),
	}
}

func (d_ DraggingImageComponent) InitWithKey(key DraggingImageComponentKey) DraggingImageComponent {
	rv := objc.CallMethod[DraggingImageComponent](d_, objc.GetSelector("initWithKey:"), key)
	return rv
}

func (dc _DraggingImageComponentClass) Alloc() DraggingImageComponent {
	rv := objc.CallMethod[DraggingImageComponent](dc, objc.GetSelector("alloc"))
	return rv
}

func (dc _DraggingImageComponentClass) New() DraggingImageComponent {
	rv := objc.CallMethod[DraggingImageComponent](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDraggingImageComponent() DraggingImageComponent {
	return DraggingImageComponentClass.New()
}

func (d_ DraggingImageComponent) Init() DraggingImageComponent {
	rv := objc.CallMethod[DraggingImageComponent](d_, objc.GetSelector("init"))
	return rv
}

func (dc _DraggingImageComponentClass) DraggingImageComponentWithKey(key DraggingImageComponentKey) DraggingImageComponent {
	rv := objc.CallMethod[DraggingImageComponent](dc, objc.GetSelector("draggingImageComponentWithKey:"), key)
	return rv
}

func (d_ DraggingImageComponent) Key() DraggingImageComponentKey {
	rv := objc.CallMethod[DraggingImageComponentKey](d_, objc.GetSelector("key"))
	return rv
}

func (d_ DraggingImageComponent) SetKey(value DraggingImageComponentKey) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setKey:"), value)
}

func (d_ DraggingImageComponent) Contents() objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.GetSelector("contents"))
	return rv
}

func (d_ DraggingImageComponent) SetContents(value objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setContents:"), objc.ExtractPtr(value))
}

func (d_ DraggingImageComponent) Frame() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](d_, objc.GetSelector("frame"))
	return rv
}

func (d_ DraggingImageComponent) SetFrame(value foundation.Rect) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setFrame:"), value)
}
