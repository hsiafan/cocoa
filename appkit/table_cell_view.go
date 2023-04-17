// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TableCellViewClass = _TableCellViewClass{objc.GetClass("NSTableCellView")}

type _TableCellViewClass struct {
	objc.Class
}

type ITableCellView interface {
	IView
	ObjectValue() objc.Object
	SetObjectValue(value objc.IObject)
	ImageView() ImageView
	SetImageView(value IImageView)
	TextField() TextField
	SetTextField(value ITextField)
	BackgroundStyle() BackgroundStyle
	SetBackgroundStyle(value BackgroundStyle)
	RowSizeStyle() TableViewRowSizeStyle
	SetRowSizeStyle(value TableViewRowSizeStyle)
	DraggingImageComponents() []DraggingImageComponent
}

type TableCellView struct {
	View
}

func MakeTableCellView(ptr unsafe.Pointer) TableCellView {
	return TableCellView{
		View: MakeView(ptr),
	}
}

func (t_ TableCellView) InitWithFrame(frameRect foundation.Rect) TableCellView {
	rv := ffi.CallMethod[TableCellView](t_, "initWithFrame:", frameRect)
	return rv
}

func (t_ TableCellView) Init() TableCellView {
	rv := ffi.CallMethod[TableCellView](t_, "init")
	return rv
}

func (tc _TableCellViewClass) Alloc() TableCellView {
	rv := ffi.CallMethod[TableCellView](tc, "alloc")
	return rv
}

func (tc _TableCellViewClass) New() TableCellView {
	rv := ffi.CallMethod[TableCellView](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTableCellView() TableCellView {
	return TableCellViewClass.New()
}

func (t_ TableCellView) ObjectValue() objc.Object {
	rv := ffi.CallMethod[objc.Object](t_, "objectValue")
	return rv
}

func (t_ TableCellView) SetObjectValue(value objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "setObjectValue:", value)
}

func (t_ TableCellView) ImageView() ImageView {
	rv := ffi.CallMethod[ImageView](t_, "imageView")
	return rv
}

func (t_ TableCellView) SetImageView(value IImageView) {
	ffi.CallMethod[ffi.Void](t_, "setImageView:", value)
}

func (t_ TableCellView) TextField() TextField {
	rv := ffi.CallMethod[TextField](t_, "textField")
	return rv
}

func (t_ TableCellView) SetTextField(value ITextField) {
	ffi.CallMethod[ffi.Void](t_, "setTextField:", value)
}

func (t_ TableCellView) BackgroundStyle() BackgroundStyle {
	rv := ffi.CallMethod[BackgroundStyle](t_, "backgroundStyle")
	return rv
}

func (t_ TableCellView) SetBackgroundStyle(value BackgroundStyle) {
	ffi.CallMethod[ffi.Void](t_, "setBackgroundStyle:", value)
}

func (t_ TableCellView) RowSizeStyle() TableViewRowSizeStyle {
	rv := ffi.CallMethod[TableViewRowSizeStyle](t_, "rowSizeStyle")
	return rv
}

func (t_ TableCellView) SetRowSizeStyle(value TableViewRowSizeStyle) {
	ffi.CallMethod[ffi.Void](t_, "setRowSizeStyle:", value)
}

func (t_ TableCellView) DraggingImageComponents() []DraggingImageComponent {
	rv := ffi.CallMethod[[]DraggingImageComponent](t_, "draggingImageComponents")
	return rv
}