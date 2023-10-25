// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ImageViewClass = _ImageViewClass{objc.GetClass("NSImageView")}

type _ImageViewClass struct {
	objc.Class
}

type IImageView interface {
	IControl
	Image() Image
	SetImage(value IImage)
	ImageFrameStyle() ImageFrameStyle
	SetImageFrameStyle(value ImageFrameStyle)
	ImageAlignment() ImageAlignment
	SetImageAlignment(value ImageAlignment)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
	Animates() bool
	SetAnimates(value bool)
	IsEditable() bool
	SetEditable(value bool)
	AllowsCutCopyPaste() bool
	SetAllowsCutCopyPaste(value bool)
	ContentTintColor() Color
	SetContentTintColor(value IColor)
	SymbolConfiguration() ImageSymbolConfiguration
	SetSymbolConfiguration(value IImageSymbolConfiguration)
}

type ImageView struct {
	Control
}

func MakeImageView(ptr unsafe.Pointer) ImageView {
	return ImageView{
		Control: MakeControl(ptr),
	}
}

func (ic _ImageViewClass) ImageViewWithImage(image IImage) ImageView {
	rv := objc.CallMethod[ImageView](ic, objc.SEL("imageViewWithImage:"), objc.ExtractPtr(image))
	return rv
}

func (i_ ImageView) InitWithFrame(frameRect foundation.Rect) ImageView {
	rv := objc.CallMethod[ImageView](i_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (i_ ImageView) Init() ImageView {
	rv := objc.CallMethod[ImageView](i_, objc.SEL("init"))
	return rv
}

func (ic _ImageViewClass) Alloc() ImageView {
	rv := objc.CallMethod[ImageView](ic, objc.SEL("alloc"))
	return rv
}

func (ic _ImageViewClass) New() ImageView {
	rv := objc.CallMethod[ImageView](ic, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewImageView() ImageView {
	return ImageViewClass.New()
}

func (i_ ImageView) Image() Image {
	rv := objc.CallMethod[Image](i_, objc.SEL("image"))
	return rv
}

func (i_ ImageView) SetImage(value IImage) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setImage:"), objc.ExtractPtr(value))
}

func (i_ ImageView) ImageFrameStyle() ImageFrameStyle {
	rv := objc.CallMethod[ImageFrameStyle](i_, objc.SEL("imageFrameStyle"))
	return rv
}

func (i_ ImageView) SetImageFrameStyle(value ImageFrameStyle) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setImageFrameStyle:"), value)
}

func (i_ ImageView) ImageAlignment() ImageAlignment {
	rv := objc.CallMethod[ImageAlignment](i_, objc.SEL("imageAlignment"))
	return rv
}

func (i_ ImageView) SetImageAlignment(value ImageAlignment) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setImageAlignment:"), value)
}

func (i_ ImageView) ImageScaling() ImageScaling {
	rv := objc.CallMethod[ImageScaling](i_, objc.SEL("imageScaling"))
	return rv
}

func (i_ ImageView) SetImageScaling(value ImageScaling) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setImageScaling:"), value)
}

func (i_ ImageView) Animates() bool {
	rv := objc.CallMethod[bool](i_, objc.SEL("animates"))
	return rv
}

func (i_ ImageView) SetAnimates(value bool) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setAnimates:"), value)
}

func (i_ ImageView) IsEditable() bool {
	rv := objc.CallMethod[bool](i_, objc.SEL("isEditable"))
	return rv
}

func (i_ ImageView) SetEditable(value bool) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setEditable:"), value)
}

func (i_ ImageView) AllowsCutCopyPaste() bool {
	rv := objc.CallMethod[bool](i_, objc.SEL("allowsCutCopyPaste"))
	return rv
}

func (i_ ImageView) SetAllowsCutCopyPaste(value bool) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setAllowsCutCopyPaste:"), value)
}

func (i_ ImageView) ContentTintColor() Color {
	rv := objc.CallMethod[Color](i_, objc.SEL("contentTintColor"))
	return rv
}

func (i_ ImageView) SetContentTintColor(value IColor) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setContentTintColor:"), objc.ExtractPtr(value))
}

func (i_ ImageView) SymbolConfiguration() ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](i_, objc.SEL("symbolConfiguration"))
	return rv
}

func (i_ ImageView) SetSymbolConfiguration(value IImageSymbolConfiguration) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setSymbolConfiguration:"), objc.ExtractPtr(value))
}
