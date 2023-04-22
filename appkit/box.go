// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[Box](b_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func (b_ Box) Init() Box {
	rv := objc.CallMethod[Box](b_, objc.GetSelector("init"))
	return rv
}

func (bc _BoxClass) Alloc() Box {
	rv := objc.CallMethod[Box](bc, objc.GetSelector("alloc"))
	return rv
}

func (bc _BoxClass) New() Box {
	rv := objc.CallMethod[Box](bc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewBox() Box {
	return BoxClass.New()
}

// deprecated
func (b_ Box) SetTitleWithMnemonic(stringWithAmpersand string) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setTitleWithMnemonic:"), stringWithAmpersand)
}

func (b_ Box) SetFrameFromContentFrame(contentFrame foundation.Rect) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setFrameFromContentFrame:"), contentFrame)
}

func (b_ Box) SizeToFit() {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("sizeToFit"))
}

func (b_ Box) BorderRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](b_, objc.GetSelector("borderRect"))
	return rv
}

func (b_ Box) BoxType() BoxType {
	rv := objc.CallMethod[BoxType](b_, objc.GetSelector("boxType"))
	return rv
}

func (b_ Box) SetBoxType(value BoxType) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setBoxType:"), value)
}

// deprecated
func (b_ Box) BorderType() BorderType {
	rv := objc.CallMethod[BorderType](b_, objc.GetSelector("borderType"))
	return rv
}

// deprecated
func (b_ Box) SetBorderType(value BorderType) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setBorderType:"), value)
}

func (b_ Box) IsTransparent() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("isTransparent"))
	return rv
}

func (b_ Box) SetTransparent(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setTransparent:"), value)
}

func (b_ Box) Title() string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("title"))
	return rv
}

func (b_ Box) SetTitle(value string) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setTitle:"), value)
}

func (b_ Box) TitleFont() Font {
	rv := objc.CallMethod[Font](b_, objc.GetSelector("titleFont"))
	return rv
}

func (b_ Box) SetTitleFont(value IFont) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setTitleFont:"), value)
}

func (b_ Box) TitlePosition() TitlePosition {
	rv := objc.CallMethod[TitlePosition](b_, objc.GetSelector("titlePosition"))
	return rv
}

func (b_ Box) SetTitlePosition(value TitlePosition) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setTitlePosition:"), value)
}

func (b_ Box) TitleCell() objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("titleCell"))
	return rv
}

func (b_ Box) TitleRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](b_, objc.GetSelector("titleRect"))
	return rv
}

func (b_ Box) BorderColor() Color {
	rv := objc.CallMethod[Color](b_, objc.GetSelector("borderColor"))
	return rv
}

func (b_ Box) SetBorderColor(value IColor) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setBorderColor:"), value)
}

func (b_ Box) BorderWidth() float64 {
	rv := objc.CallMethod[float64](b_, objc.GetSelector("borderWidth"))
	return rv
}

func (b_ Box) SetBorderWidth(value float64) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setBorderWidth:"), value)
}

func (b_ Box) CornerRadius() float64 {
	rv := objc.CallMethod[float64](b_, objc.GetSelector("cornerRadius"))
	return rv
}

func (b_ Box) SetCornerRadius(value float64) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setCornerRadius:"), value)
}

func (b_ Box) FillColor() Color {
	rv := objc.CallMethod[Color](b_, objc.GetSelector("fillColor"))
	return rv
}

func (b_ Box) SetFillColor(value IColor) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setFillColor:"), value)
}

func (b_ Box) ContentView() View {
	rv := objc.CallMethod[View](b_, objc.GetSelector("contentView"))
	return rv
}

func (b_ Box) SetContentView(value IView) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setContentView:"), value)
}

func (b_ Box) ContentViewMargins() foundation.Size {
	rv := objc.CallMethod[foundation.Size](b_, objc.GetSelector("contentViewMargins"))
	return rv
}

func (b_ Box) SetContentViewMargins(value foundation.Size) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setContentViewMargins:"), value)
}
