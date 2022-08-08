// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var BoxClass = _BoxClass{objc.GetClass("NSBox")}

type _BoxClass struct {
	objc.Class
}

type IBox interface {
	IView
	// deprecated
	SetTitleWithMnemonic(stringWithAmpersand string)
	SetFrameFromContentFrame(contentFrame foundation.Rect)
	SizeToFit()
	BorderRect() foundation.Rect
	BoxType() BoxType
	SetBoxType(value BoxType)
	// deprecated
	BorderType() BorderType
	// deprecated
	SetBorderType(value BorderType)
	IsTransparent() bool
	SetTransparent(value bool)
	Title() string
	SetTitle(value string)
	TitleFont() Font
	SetTitleFont(value IFont)
	TitlePosition() TitlePosition
	SetTitlePosition(value TitlePosition)
	TitleCell() objc.Object
	TitleRect() foundation.Rect
	BorderColor() Color
	SetBorderColor(value IColor)
	BorderWidth() float64
	SetBorderWidth(value float64)
	CornerRadius() float64
	SetCornerRadius(value float64)
	FillColor() Color
	SetFillColor(value IColor)
	ContentView() View
	SetContentView(value IView)
	ContentViewMargins() foundation.Size
	SetContentViewMargins(value foundation.Size)
}

type Box struct {
	View
}

func MakeBox(ptr unsafe.Pointer) Box {
	return Box{
		View: MakeView(ptr),
	}
}

func (b_ Box) InitWithFrame(frameRect foundation.Rect) Box {
	rv := ffi.CallMethod[Box](b_, "initWithFrame:", frameRect)
	return rv
}

func (b_ Box) Init() Box {
	rv := ffi.CallMethod[Box](b_, "init")
	return rv
}

func (bc _BoxClass) Alloc() Box {
	rv := ffi.CallMethod[Box](bc, "alloc")
	return rv
}

func (bc _BoxClass) New() Box {
	rv := ffi.CallMethod[Box](bc, "new")
	rv.Autorelease()
	return rv
}

func NewBox() Box {
	return BoxClass.New()
}

// deprecated
func (b_ Box) SetTitleWithMnemonic(stringWithAmpersand string) {
	ffi.CallMethod[ffi.Void](b_, "setTitleWithMnemonic:", stringWithAmpersand)
}

func (b_ Box) SetFrameFromContentFrame(contentFrame foundation.Rect) {
	ffi.CallMethod[ffi.Void](b_, "setFrameFromContentFrame:", contentFrame)
}

func (b_ Box) SizeToFit() {
	ffi.CallMethod[ffi.Void](b_, "sizeToFit")
}

func (b_ Box) BorderRect() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](b_, "borderRect")
	return rv
}

func (b_ Box) BoxType() BoxType {
	rv := ffi.CallMethod[BoxType](b_, "boxType")
	return rv
}

func (b_ Box) SetBoxType(value BoxType) {
	ffi.CallMethod[ffi.Void](b_, "setBoxType:", value)
}

// deprecated
func (b_ Box) BorderType() BorderType {
	rv := ffi.CallMethod[BorderType](b_, "borderType")
	return rv
}

// deprecated
func (b_ Box) SetBorderType(value BorderType) {
	ffi.CallMethod[ffi.Void](b_, "setBorderType:", value)
}

func (b_ Box) IsTransparent() bool {
	rv := ffi.CallMethod[bool](b_, "isTransparent")
	return rv
}

func (b_ Box) SetTransparent(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setTransparent:", value)
}

func (b_ Box) Title() string {
	rv := ffi.CallMethod[string](b_, "title")
	return rv
}

func (b_ Box) SetTitle(value string) {
	ffi.CallMethod[ffi.Void](b_, "setTitle:", value)
}

func (b_ Box) TitleFont() Font {
	rv := ffi.CallMethod[Font](b_, "titleFont")
	return rv
}

func (b_ Box) SetTitleFont(value IFont) {
	ffi.CallMethod[ffi.Void](b_, "setTitleFont:", value)
}

func (b_ Box) TitlePosition() TitlePosition {
	rv := ffi.CallMethod[TitlePosition](b_, "titlePosition")
	return rv
}

func (b_ Box) SetTitlePosition(value TitlePosition) {
	ffi.CallMethod[ffi.Void](b_, "setTitlePosition:", value)
}

func (b_ Box) TitleCell() objc.Object {
	rv := ffi.CallMethod[objc.Object](b_, "titleCell")
	return rv
}

func (b_ Box) TitleRect() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](b_, "titleRect")
	return rv
}

func (b_ Box) BorderColor() Color {
	rv := ffi.CallMethod[Color](b_, "borderColor")
	return rv
}

func (b_ Box) SetBorderColor(value IColor) {
	ffi.CallMethod[ffi.Void](b_, "setBorderColor:", value)
}

func (b_ Box) BorderWidth() float64 {
	rv := ffi.CallMethod[float64](b_, "borderWidth")
	return rv
}

func (b_ Box) SetBorderWidth(value float64) {
	ffi.CallMethod[ffi.Void](b_, "setBorderWidth:", value)
}

func (b_ Box) CornerRadius() float64 {
	rv := ffi.CallMethod[float64](b_, "cornerRadius")
	return rv
}

func (b_ Box) SetCornerRadius(value float64) {
	ffi.CallMethod[ffi.Void](b_, "setCornerRadius:", value)
}

func (b_ Box) FillColor() Color {
	rv := ffi.CallMethod[Color](b_, "fillColor")
	return rv
}

func (b_ Box) SetFillColor(value IColor) {
	ffi.CallMethod[ffi.Void](b_, "setFillColor:", value)
}

func (b_ Box) ContentView() View {
	rv := ffi.CallMethod[View](b_, "contentView")
	return rv
}

func (b_ Box) SetContentView(value IView) {
	ffi.CallMethod[ffi.Void](b_, "setContentView:", value)
}

func (b_ Box) ContentViewMargins() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](b_, "contentViewMargins")
	return rv
}

func (b_ Box) SetContentViewMargins(value foundation.Size) {
	ffi.CallMethod[ffi.Void](b_, "setContentViewMargins:", value)
}
