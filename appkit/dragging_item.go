// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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

func (d_ DraggingItem) InitWithPasteboardWriter(pasteboardWriter objc.IObject) DraggingItem {
	rv := objc.CallMethod[DraggingItem](d_, objc.SEL("initWithPasteboardWriter:"), objc.ExtractPtr(pasteboardWriter))
	return rv
}

func (dc _DraggingItemClass) Alloc() DraggingItem {
	rv := objc.CallMethod[DraggingItem](dc, objc.SEL("alloc"))
	return rv
}

func (dc _DraggingItemClass) New() DraggingItem {
	rv := objc.CallMethod[DraggingItem](dc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewDraggingItem() DraggingItem {
	return DraggingItemClass.New()
}

func (d_ DraggingItem) Init() DraggingItem {
	rv := objc.CallMethod[DraggingItem](d_, objc.SEL("init"))
	return rv
}

func (d_ DraggingItem) SetDraggingFrame_Contents(frame foundation.Rect, contents objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setDraggingFrame:contents:"), frame, objc.ExtractPtr(contents))
}

func (d_ DraggingItem) DraggingFrame() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](d_, objc.SEL("draggingFrame"))
	return rv
}

func (d_ DraggingItem) SetDraggingFrame(value foundation.Rect) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setDraggingFrame:"), value)
}

func (d_ DraggingItem) ImageComponents() []DraggingImageComponent {
	rv := objc.CallMethod[[]DraggingImageComponent](d_, objc.SEL("imageComponents"))
	return rv
}

func (d_ DraggingItem) ImageComponentsProvider() func() []DraggingImageComponent {
	rv := objc.CallMethod[func() []DraggingImageComponent](d_, objc.SEL("imageComponentsProvider"))
	return rv
}

func (d_ DraggingItem) SetImageComponentsProvider(value func() []IDraggingImageComponent) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setImageComponentsProvider:"), value)
}

func (d_ DraggingItem) Item() objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.SEL("item"))
	return rv
}
