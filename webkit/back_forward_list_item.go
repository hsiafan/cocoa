// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[BackForwardListItem](bc, "alloc")
	return rv
}

func (bc _BackForwardListItemClass) New() BackForwardListItem {
	rv := ffi.CallMethod[BackForwardListItem](bc, "new")
	rv.Autorelease()
	return rv
}

func NewBackForwardListItem() BackForwardListItem {
	return BackForwardListItemClass.New()
}

func (b_ BackForwardListItem) Init() BackForwardListItem {
	rv := ffi.CallMethod[BackForwardListItem](b_, "init")
	return rv
}

func (b_ BackForwardListItem) Title() string {
	rv := ffi.CallMethod[string](b_, "title")
	return rv
}

func (b_ BackForwardListItem) URL() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](b_, "URL")
	return rv
}

func (b_ BackForwardListItem) InitialURL() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](b_, "initialURL")
	return rv
}
