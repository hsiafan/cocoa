// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[TableCellView](t_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (t_ TableCellView) Init() TableCellView {
	rv := objc.CallMethod[TableCellView](t_, objc.SEL("init"))
	return rv
}

func (tc _TableCellViewClass) Alloc() TableCellView {
	rv := objc.CallMethod[TableCellView](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TableCellViewClass) New() TableCellView {
	rv := objc.CallMethod[TableCellView](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTableCellView() TableCellView {
	return TableCellViewClass.New()
}

func (t_ TableCellView) ObjectValue() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.SEL("objectValue"))
	return rv
}

func (t_ TableCellView) SetObjectValue(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setObjectValue:"), objc.ExtractPtr(value))
}

// weak property
func (t_ TableCellView) ImageView() ImageView {
	rv := objc.CallMethod[ImageView](t_, objc.SEL("imageView"))
	return rv
}

// weak property
func (t_ TableCellView) SetImageView(value IImageView) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setImageView:"), objc.ExtractPtr(value))
}

// weak property
func (t_ TableCellView) TextField() TextField {
	rv := objc.CallMethod[TextField](t_, objc.SEL("textField"))
	return rv
}

// weak property
func (t_ TableCellView) SetTextField(value ITextField) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTextField:"), objc.ExtractPtr(value))
}

func (t_ TableCellView) BackgroundStyle() BackgroundStyle {
	rv := objc.CallMethod[BackgroundStyle](t_, objc.SEL("backgroundStyle"))
	return rv
}

func (t_ TableCellView) SetBackgroundStyle(value BackgroundStyle) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setBackgroundStyle:"), value)
}

func (t_ TableCellView) RowSizeStyle() TableViewRowSizeStyle {
	rv := objc.CallMethod[TableViewRowSizeStyle](t_, objc.SEL("rowSizeStyle"))
	return rv
}

func (t_ TableCellView) SetRowSizeStyle(value TableViewRowSizeStyle) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setRowSizeStyle:"), value)
}

func (t_ TableCellView) DraggingImageComponents() []DraggingImageComponent {
	rv := objc.CallMethod[[]DraggingImageComponent](t_, objc.SEL("draggingImageComponents"))
	return rv
}
