// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var BackForwardListItemClass = _BackForwardListItemClass{objc.GetClass("WKBackForwardListItem")}

type _BackForwardListItemClass struct {
	objc.Class
}

type IBackForwardListItem interface {
	objc.IObject
	Title() string
	URL() foundation.URL
	InitialURL() foundation.URL
}

type BackForwardListItem struct {
	objc.Object
}

func MakeBackForwardListItem(ptr unsafe.Pointer) BackForwardListItem {
	return BackForwardListItem{
		Object: objc.MakeObject(ptr),
	}
}

func (bc _BackForwardListItemClass) Alloc() BackForwardListItem {
	rv := objc.CallMethod[BackForwardListItem](bc, objc.GetSelector("alloc"))
	return rv
}

func (bc _BackForwardListItemClass) New() BackForwardListItem {
	rv := objc.CallMethod[BackForwardListItem](bc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewBackForwardListItem() BackForwardListItem {
	return BackForwardListItemClass.New()
}

func (b_ BackForwardListItem) Init() BackForwardListItem {
	rv := objc.CallMethod[BackForwardListItem](b_, objc.GetSelector("init"))
	return rv
}

func (b_ BackForwardListItem) Title() string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("title"))
	return rv
}

func (b_ BackForwardListItem) URL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](b_, objc.GetSelector("URL"))
	return rv
}

func (b_ BackForwardListItem) InitialURL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](b_, objc.GetSelector("initialURL"))
	return rv
}
