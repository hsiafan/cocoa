// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var WindowTabClass = _WindowTabClass{objc.GetClass("NSWindowTab")}

type _WindowTabClass struct {
	objc.Class
}

type IWindowTab interface {
	objc.IObject
	Title() string
	SetTitle(value string)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	ToolTip() string
	SetToolTip(value string)
	AccessoryView() View
	SetAccessoryView(value IView)
}

type WindowTab struct {
	objc.Object
}

func MakeWindowTab(ptr unsafe.Pointer) WindowTab {
	return WindowTab{
		Object: objc.MakeObject(ptr),
	}
}

func (wc _WindowTabClass) Alloc() WindowTab {
	rv := objc.CallMethod[WindowTab](wc, "alloc")
	return rv
}

func (wc _WindowTabClass) New() WindowTab {
	rv := objc.CallMethod[WindowTab](wc, "new")
	rv.Autorelease()
	return rv
}

func NewWindowTab() WindowTab {
	return WindowTabClass.New()
}

func (w_ WindowTab) Init() WindowTab {
	rv := objc.CallMethod[WindowTab](w_, "init")
	return rv
}

func (w_ WindowTab) Title() string {
	rv := objc.CallMethod[string](w_, "title")
	return rv
}

func (w_ WindowTab) SetTitle(value string) {
	objc.CallMethod[objc.Void](w_, "setTitle:", value)
}

func (w_ WindowTab) AttributedTitle() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](w_, "attributedTitle")
	return rv
}

func (w_ WindowTab) SetAttributedTitle(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](w_, "setAttributedTitle:", value)
}

func (w_ WindowTab) ToolTip() string {
	rv := objc.CallMethod[string](w_, "toolTip")
	return rv
}

func (w_ WindowTab) SetToolTip(value string) {
	objc.CallMethod[objc.Void](w_, "setToolTip:", value)
}

func (w_ WindowTab) AccessoryView() View {
	rv := objc.CallMethod[View](w_, "accessoryView")
	return rv
}

func (w_ WindowTab) SetAccessoryView(value IView) {
	objc.CallMethod[objc.Void](w_, "setAccessoryView:", value)
}
