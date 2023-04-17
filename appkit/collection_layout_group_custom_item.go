// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionLayoutGroupCustomItemClass = _CollectionLayoutGroupCustomItemClass{objc.GetClass("NSCollectionLayoutGroupCustomItem")}

type _CollectionLayoutGroupCustomItemClass struct {
	objc.Class
}

type ICollectionLayoutGroupCustomItem interface {
	objc.IObject
	Frame() foundation.Rect
	ZIndex() int
}

type CollectionLayoutGroupCustomItem struct {
	objc.Object
}

func MakeCollectionLayoutGroupCustomItem(ptr unsafe.Pointer) CollectionLayoutGroupCustomItem {
	return CollectionLayoutGroupCustomItem{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutGroupCustomItemClass) CustomItemWithFrame(frame foundation.Rect) CollectionLayoutGroupCustomItem {
	rv := ffi.CallMethod[CollectionLayoutGroupCustomItem](cc, "customItemWithFrame:", frame)
	return rv
}

func (cc _CollectionLayoutGroupCustomItemClass) CustomItemWithFrame_ZIndex(frame foundation.Rect, zIndex int) CollectionLayoutGroupCustomItem {
	rv := ffi.CallMethod[CollectionLayoutGroupCustomItem](cc, "customItemWithFrame:zIndex:", frame, zIndex)
	return rv
}

func (cc _CollectionLayoutGroupCustomItemClass) Alloc() CollectionLayoutGroupCustomItem {
	rv := ffi.CallMethod[CollectionLayoutGroupCustomItem](cc, "alloc")
	return rv
}

func (cc _CollectionLayoutGroupCustomItemClass) New() CollectionLayoutGroupCustomItem {
	rv := ffi.CallMethod[CollectionLayoutGroupCustomItem](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutGroupCustomItem() CollectionLayoutGroupCustomItem {
	return CollectionLayoutGroupCustomItemClass.New()
}

func (c_ CollectionLayoutGroupCustomItem) Init() CollectionLayoutGroupCustomItem {
	rv := ffi.CallMethod[CollectionLayoutGroupCustomItem](c_, "init")
	return rv
}

func (c_ CollectionLayoutGroupCustomItem) Frame() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](c_, "frame")
	return rv
}

func (c_ CollectionLayoutGroupCustomItem) ZIndex() int {
	rv := ffi.CallMethod[int](c_, "zIndex")
	return rv
}