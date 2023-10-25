// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[CollectionLayoutGroupCustomItem](cc, objc.SEL("customItemWithFrame:"), frame)
	return rv
}

func (cc _CollectionLayoutGroupCustomItemClass) CustomItemWithFrame_ZIndex(frame foundation.Rect, zIndex int) CollectionLayoutGroupCustomItem {
	rv := objc.CallMethod[CollectionLayoutGroupCustomItem](cc, objc.SEL("customItemWithFrame:zIndex:"), frame, zIndex)
	return rv
}

func (cc _CollectionLayoutGroupCustomItemClass) Alloc() CollectionLayoutGroupCustomItem {
	rv := objc.CallMethod[CollectionLayoutGroupCustomItem](cc, objc.SEL("alloc"))
	return rv
}

func (cc _CollectionLayoutGroupCustomItemClass) New() CollectionLayoutGroupCustomItem {
	rv := objc.CallMethod[CollectionLayoutGroupCustomItem](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutGroupCustomItem() CollectionLayoutGroupCustomItem {
	return CollectionLayoutGroupCustomItemClass.New()
}

func (c_ CollectionLayoutGroupCustomItem) Init() CollectionLayoutGroupCustomItem {
	rv := objc.CallMethod[CollectionLayoutGroupCustomItem](c_, objc.SEL("init"))
	return rv
}

func (c_ CollectionLayoutGroupCustomItem) Frame() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.SEL("frame"))
	return rv
}

func (c_ CollectionLayoutGroupCustomItem) ZIndex() int {
	rv := objc.CallMethod[int](c_, objc.SEL("zIndex"))
	return rv
}
