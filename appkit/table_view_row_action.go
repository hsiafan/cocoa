// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[TableViewRowAction](tc, objc.SEL("rowActionWithStyle:title:handler:"), style, title, handler)
	return rv
}

func (tc _TableViewRowActionClass) Alloc() TableViewRowAction {
	rv := objc.CallMethod[TableViewRowAction](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TableViewRowActionClass) New() TableViewRowAction {
	rv := objc.CallMethod[TableViewRowAction](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTableViewRowAction() TableViewRowAction {
	return TableViewRowActionClass.New()
}

func (t_ TableViewRowAction) Init() TableViewRowAction {
	rv := objc.CallMethod[TableViewRowAction](t_, objc.SEL("init"))
	return rv
}

func (t_ TableViewRowAction) Style() TableViewRowActionStyle {
	rv := objc.CallMethod[TableViewRowActionStyle](t_, objc.SEL("style"))
	return rv
}

func (t_ TableViewRowAction) Title() string {
	rv := objc.CallMethod[string](t_, objc.SEL("title"))
	return rv
}

func (t_ TableViewRowAction) SetTitle(value string) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTitle:"), value)
}

func (t_ TableViewRowAction) BackgroundColor() Color {
	rv := objc.CallMethod[Color](t_, objc.SEL("backgroundColor"))
	return rv
}

func (t_ TableViewRowAction) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (t_ TableViewRowAction) Image() Image {
	rv := objc.CallMethod[Image](t_, objc.SEL("image"))
	return rv
}

func (t_ TableViewRowAction) SetImage(value IImage) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setImage:"), objc.ExtractPtr(value))
}
