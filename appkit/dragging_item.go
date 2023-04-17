// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var DraggingItemClass = _DraggingItemClass{objc.GetClass("NSDraggingItem")}

type _DraggingItemClass struct {
	objc.Class
}

type IDraggingItem interface {
	objc.IObject
	SetDraggingFrame_Contents(frame foundation.Rect, contents objc.IObject)
	DraggingFrame() foundation.Rect
	SetDraggingFrame(value foundation.Rect)
	ImageComponents() []DraggingImageComponent
	ImageComponentsProvider() func() []DraggingImageComponent
	SetImageComponentsProvider(value func() []IDraggingImageComponent)
	Item() objc.Object
}

type DraggingItem struct {
	objc.Object
}

func MakeDraggingItem(ptr unsafe.Pointer) DraggingItem {
	return DraggingItem{
		Object: objc.MakeObject(ptr),
	}
}

func (d_ DraggingItem) InitWithPasteboardWriter(pasteboardWriter PasteboardWriting) DraggingItem {
	po := ffi.CreateProtocol("NSPasteboardWriting", pasteboardWriter)
	defer po.Release()
	rv := ffi.CallMethod[DraggingItem](d_, "initWithPasteboardWriter:", po)
	return rv
}

func (dc _DraggingItemClass) Alloc() DraggingItem {
	rv := ffi.CallMethod[DraggingItem](dc, "alloc")
	return rv
}

func (dc _DraggingItemClass) New() DraggingItem {
	rv := ffi.CallMethod[DraggingItem](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDraggingItem() DraggingItem {
	return DraggingItemClass.New()
}

func (d_ DraggingItem) Init() DraggingItem {
	rv := ffi.CallMethod[DraggingItem](d_, "init")
	return rv
}

func (d_ DraggingItem) SetDraggingFrame_Contents(frame foundation.Rect, contents objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "setDraggingFrame:contents:", frame, contents)
}

func (d_ DraggingItem) DraggingFrame() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](d_, "draggingFrame")
	return rv
}

func (d_ DraggingItem) SetDraggingFrame(value foundation.Rect) {
	ffi.CallMethod[ffi.Void](d_, "setDraggingFrame:", value)
}

func (d_ DraggingItem) ImageComponents() []DraggingImageComponent {
	rv := ffi.CallMethod[[]DraggingImageComponent](d_, "imageComponents")
	return rv
}

func (d_ DraggingItem) ImageComponentsProvider() func() []DraggingImageComponent {
	rv := ffi.CallMethod[func() []DraggingImageComponent](d_, "imageComponentsProvider")
	return rv
}

func (d_ DraggingItem) SetImageComponentsProvider(value func() []IDraggingImageComponent) {
	ffi.CallMethod[ffi.Void](d_, "setImageComponentsProvider:", value)
}

func (d_ DraggingItem) Item() objc.Object {
	rv := ffi.CallMethod[objc.Object](d_, "item")
	return rv
}
