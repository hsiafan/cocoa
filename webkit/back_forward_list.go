// auto-generated code, do not modify
package webkit

import (
	"unsafe"

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
	rv := objc.CallMethod[BackForwardList](bc, objc.SEL("alloc"))
	return rv
}

func (bc _BackForwardListClass) New() BackForwardList {
	rv := objc.CallMethod[BackForwardList](bc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewBackForwardList() BackForwardList {
	return BackForwardListClass.New()
}

func (b_ BackForwardList) Init() BackForwardList {
	rv := objc.CallMethod[BackForwardList](b_, objc.SEL("init"))
	return rv
}

func (b_ BackForwardList) ItemAtIndex(index int) BackForwardListItem {
	rv := objc.CallMethod[BackForwardListItem](b_, objc.SEL("itemAtIndex:"), index)
	return rv
}

func (b_ BackForwardList) BackItem() BackForwardListItem {
	rv := objc.CallMethod[BackForwardListItem](b_, objc.SEL("backItem"))
	return rv
}

func (b_ BackForwardList) CurrentItem() BackForwardListItem {
	rv := objc.CallMethod[BackForwardListItem](b_, objc.SEL("currentItem"))
	return rv
}

func (b_ BackForwardList) ForwardItem() BackForwardListItem {
	rv := objc.CallMethod[BackForwardListItem](b_, objc.SEL("forwardItem"))
	return rv
}

func (b_ BackForwardList) BackList() []BackForwardListItem {
	rv := objc.CallMethod[[]BackForwardListItem](b_, objc.SEL("backList"))
	return rv
}

func (b_ BackForwardList) ForwardList() []BackForwardListItem {
	rv := objc.CallMethod[[]BackForwardListItem](b_, objc.SEL("forwardList"))
	return rv
}
