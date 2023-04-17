// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var TableViewRowActionClass = _TableViewRowActionClass{objc.GetClass("NSTableViewRowAction")}

type _TableViewRowActionClass struct {
	objc.Class
}

type ITableViewRowAction interface {
	objc.IObject
	Style() TableViewRowActionStyle
	Title() string
	SetTitle(value string)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	Image() Image
	SetImage(value IImage)
}

type TableViewRowAction struct {
	objc.Object
}

func MakeTableViewRowAction(ptr unsafe.Pointer) TableViewRowAction {
	return TableViewRowAction{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TableViewRowActionClass) RowActionWithStyle_Title_Handler(style TableViewRowActionStyle, title string, handler func(action TableViewRowAction, row int)) TableViewRowAction {
	rv := ffi.CallMethod[TableViewRowAction](tc, "rowActionWithStyle:title:handler:", style, title, handler)
	return rv
}

func (tc _TableViewRowActionClass) Alloc() TableViewRowAction {
	rv := ffi.CallMethod[TableViewRowAction](tc, "alloc")
	return rv
}

func (tc _TableViewRowActionClass) New() TableViewRowAction {
	rv := ffi.CallMethod[TableViewRowAction](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTableViewRowAction() TableViewRowAction {
	return TableViewRowActionClass.New()
}

func (t_ TableViewRowAction) Init() TableViewRowAction {
	rv := ffi.CallMethod[TableViewRowAction](t_, "init")
	return rv
}

func (t_ TableViewRowAction) Style() TableViewRowActionStyle {
	rv := ffi.CallMethod[TableViewRowActionStyle](t_, "style")
	return rv
}

func (t_ TableViewRowAction) Title() string {
	rv := ffi.CallMethod[string](t_, "title")
	return rv
}

func (t_ TableViewRowAction) SetTitle(value string) {
	ffi.CallMethod[ffi.Void](t_, "setTitle:", value)
}

func (t_ TableViewRowAction) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](t_, "backgroundColor")
	return rv
}

func (t_ TableViewRowAction) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](t_, "setBackgroundColor:", value)
}

func (t_ TableViewRowAction) Image() Image {
	rv := ffi.CallMethod[Image](t_, "image")
	return rv
}

func (t_ TableViewRowAction) SetImage(value IImage) {
	ffi.CallMethod[ffi.Void](t_, "setImage:", value)
}
