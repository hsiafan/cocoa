// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var BackForwardListClass = _BackForwardListClass{objc.GetClass("WKBackForwardList")}

type _BackForwardListClass struct {
	objc.Class
}

type IBackForwardList interface {
	objc.IObject
	ItemAtIndex(index int) BackForwardListItem
	BackItem() BackForwardListItem
	CurrentItem() BackForwardListItem
	ForwardItem() BackForwardListItem
	BackList() []BackForwardListItem
	ForwardList() []BackForwardListItem
}

type BackForwardList struct {
	objc.Object
}

func MakeBackForwardList(ptr unsafe.Pointer) BackForwardList {
	return BackForwardList{
		Object: objc.MakeObject(ptr),
	}
}

func (bc _BackForwardListClass) Alloc() BackForwardList {
	rv := ffi.CallMethod[BackForwardList](bc, "alloc")
	return rv
}

func (bc _BackForwardListClass) New() BackForwardList {
	rv := ffi.CallMethod[BackForwardList](bc, "new")
	rv.Autorelease()
	return rv
}

func NewBackForwardList() BackForwardList {
	return BackForwardListClass.New()
}

func (b_ BackForwardList) Init() BackForwardList {
	rv := ffi.CallMethod[BackForwardList](b_, "init")
	return rv
}

func (b_ BackForwardList) ItemAtIndex(index int) BackForwardListItem {
	rv := ffi.CallMethod[BackForwardListItem](b_, "itemAtIndex:", index)
	return rv
}

func (b_ BackForwardList) BackItem() BackForwardListItem {
	rv := ffi.CallMethod[BackForwardListItem](b_, "backItem")
	return rv
}

func (b_ BackForwardList) CurrentItem() BackForwardListItem {
	rv := ffi.CallMethod[BackForwardListItem](b_, "currentItem")
	return rv
}

func (b_ BackForwardList) ForwardItem() BackForwardListItem {
	rv := ffi.CallMethod[BackForwardListItem](b_, "forwardItem")
	return rv
}

func (b_ BackForwardList) BackList() []BackForwardListItem {
	rv := ffi.CallMethod[[]BackForwardListItem](b_, "backList")
	return rv
}

func (b_ BackForwardList) ForwardList() []BackForwardListItem {
	rv := ffi.CallMethod[[]BackForwardListItem](b_, "forwardList")
	return rv
}
