// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[DraggingImageComponent](d_, "initWithKey:", key)
	return rv
}

func (dc _DraggingImageComponentClass) Alloc() DraggingImageComponent {
	rv := ffi.CallMethod[DraggingImageComponent](dc, "alloc")
	return rv
}

func (dc _DraggingImageComponentClass) New() DraggingImageComponent {
	rv := ffi.CallMethod[DraggingImageComponent](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDraggingImageComponent() DraggingImageComponent {
	return DraggingImageComponentClass.New()
}

func (d_ DraggingImageComponent) Init() DraggingImageComponent {
	rv := ffi.CallMethod[DraggingImageComponent](d_, "init")
	return rv
}

func (dc _DraggingImageComponentClass) DraggingImageComponentWithKey(key DraggingImageComponentKey) DraggingImageComponent {
	rv := ffi.CallMethod[DraggingImageComponent](dc, "draggingImageComponentWithKey:", key)
	return rv
}

func (d_ DraggingImageComponent) Key() DraggingImageComponentKey {
	rv := ffi.CallMethod[DraggingImageComponentKey](d_, "key")
	return rv
}

func (d_ DraggingImageComponent) SetKey(value DraggingImageComponentKey) {
	ffi.CallMethod[ffi.Void](d_, "setKey:", value)
}

func (d_ DraggingImageComponent) Contents() objc.Object {
	rv := ffi.CallMethod[objc.Object](d_, "contents")
	return rv
}

func (d_ DraggingImageComponent) SetContents(value objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "setContents:", value)
}

func (d_ DraggingImageComponent) Frame() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](d_, "frame")
	return rv
}

func (d_ DraggingImageComponent) SetFrame(value foundation.Rect) {
	ffi.CallMethod[ffi.Void](d_, "setFrame:", value)
}