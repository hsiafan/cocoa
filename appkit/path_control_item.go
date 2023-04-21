// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var PathControlItemClass = _PathControlItemClass{objc.GetClass("NSPathControlItem")}

type _PathControlItemClass struct {
	objc.Class
}

type IPathControlItem interface {
	objc.IObject
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	Image() Image
	SetImage(value IImage)
	Title() string
	SetTitle(value string)
	URL() foundation.URL
}

type PathControlItem struct {
	objc.Object
}

func MakePathControlItem(ptr unsafe.Pointer) PathControlItem {
	return PathControlItem{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PathControlItemClass) Alloc() PathControlItem {
	rv := objc.CallMethod[PathControlItem](pc, "alloc")
	return rv
}

func (pc _PathControlItemClass) New() PathControlItem {
	rv := objc.CallMethod[PathControlItem](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPathControlItem() PathControlItem {
	return PathControlItemClass.New()
}

func (p_ PathControlItem) Init() PathControlItem {
	rv := objc.CallMethod[PathControlItem](p_, "init")
	return rv
}

func (p_ PathControlItem) AttributedTitle() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](p_, "attributedTitle")
	return rv
}

func (p_ PathControlItem) SetAttributedTitle(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](p_, "setAttributedTitle:", value)
}

func (p_ PathControlItem) Image() Image {
	rv := objc.CallMethod[Image](p_, "image")
	return rv
}

func (p_ PathControlItem) SetImage(value IImage) {
	objc.CallMethod[objc.Void](p_, "setImage:", value)
}

func (p_ PathControlItem) Title() string {
	rv := objc.CallMethod[string](p_, "title")
	return rv
}

func (p_ PathControlItem) SetTitle(value string) {
	objc.CallMethod[objc.Void](p_, "setTitle:", value)
}

func (p_ PathControlItem) URL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](p_, "URL")
	return rv
}
