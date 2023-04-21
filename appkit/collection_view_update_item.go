// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionViewUpdateItemClass = _CollectionViewUpdateItemClass{objc.GetClass("NSCollectionViewUpdateItem")}

type _CollectionViewUpdateItemClass struct {
	objc.Class
}

type ICollectionViewUpdateItem interface {
	objc.IObject
	IndexPathBeforeUpdate() foundation.IndexPath
	IndexPathAfterUpdate() foundation.IndexPath
	UpdateAction() CollectionUpdateAction
}

type CollectionViewUpdateItem struct {
	objc.Object
}

func MakeCollectionViewUpdateItem(ptr unsafe.Pointer) CollectionViewUpdateItem {
	return CollectionViewUpdateItem{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionViewUpdateItemClass) Alloc() CollectionViewUpdateItem {
	rv := objc.CallMethod[CollectionViewUpdateItem](cc, "alloc")
	return rv
}

func (cc _CollectionViewUpdateItemClass) New() CollectionViewUpdateItem {
	rv := objc.CallMethod[CollectionViewUpdateItem](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewUpdateItem() CollectionViewUpdateItem {
	return CollectionViewUpdateItemClass.New()
}

func (c_ CollectionViewUpdateItem) Init() CollectionViewUpdateItem {
	rv := objc.CallMethod[CollectionViewUpdateItem](c_, "init")
	return rv
}

func (c_ CollectionViewUpdateItem) IndexPathBeforeUpdate() foundation.IndexPath {
	rv := objc.CallMethod[foundation.IndexPath](c_, "indexPathBeforeUpdate")
	return rv
}

func (c_ CollectionViewUpdateItem) IndexPathAfterUpdate() foundation.IndexPath {
	rv := objc.CallMethod[foundation.IndexPath](c_, "indexPathAfterUpdate")
	return rv
}

func (c_ CollectionViewUpdateItem) UpdateAction() CollectionUpdateAction {
	rv := objc.CallMethod[CollectionUpdateAction](c_, "updateAction")
	return rv
}
