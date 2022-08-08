// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var TextLayerClass = _TextLayerClass{objc.GetClass("CATextLayer")}

type _TextLayerClass struct {
	objc.Class
}

type ITextLayer interface {
	ILayer
	String() objc.Object
	SetString(value objc.IObject)
	FontSize() float64
	SetFontSize(value float64)
	ForegroundColor() coregraphics.ColorRef
	SetForegroundColor(value coregraphics.ColorRef)
	AllowsFontSubpixelQuantization() bool
	SetAllowsFontSubpixelQuantization(value bool)
	IsWrapped() bool
	SetWrapped(value bool)
	AlignmentMode() TextLayerAlignmentMode
	SetAlignmentMode(value TextLayerAlignmentMode)
	TruncationMode() TextLayerTruncationMode
	SetTruncationMode(value TextLayerTruncationMode)
}

type TextLayer struct {
	Layer
}

func MakeTextLayer(ptr unsafe.Pointer) TextLayer {
	return TextLayer{
		Layer: MakeLayer(ptr),
	}
}

func (tc _TextLayerClass) Layer() TextLayer {
	rv := ffi.CallMethod[TextLayer](tc, "layer")
	return rv
}

func (t_ TextLayer) Init() TextLayer {
	rv := ffi.CallMethod[TextLayer](t_, "init")
	return rv
}

func (t_ TextLayer) InitWithLayer(layer objc.IObject) TextLayer {
	rv := ffi.CallMethod[TextLayer](t_, "initWithLayer:", layer)
	return rv
}

func (t_ TextLayer) PresentationLayer() TextLayer {
	rv := ffi.CallMethod[TextLayer](t_, "presentationLayer")
	return rv
}

func (t_ TextLayer) ModelLayer() TextLayer {
	rv := ffi.CallMethod[TextLayer](t_, "modelLayer")
	return rv
}

func (tc _TextLayerClass) Alloc() TextLayer {
	rv := ffi.CallMethod[TextLayer](tc, "alloc")
	return rv
}

func (tc _TextLayerClass) New() TextLayer {
	rv := ffi.CallMethod[TextLayer](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextLayer() TextLayer {
	return TextLayerClass.New()
}

func (t_ TextLayer) String() objc.Object {
	rv := ffi.CallMethod[objc.Object](t_, "string")
	return rv
}

func (t_ TextLayer) SetString(value objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "setString:", value)
}

func (t_ TextLayer) FontSize() float64 {
	rv := ffi.CallMethod[float64](t_, "fontSize")
	return rv
}

func (t_ TextLayer) SetFontSize(value float64) {
	ffi.CallMethod[ffi.Void](t_, "setFontSize:", value)
}

func (t_ TextLayer) ForegroundColor() coregraphics.ColorRef {
	rv := ffi.CallMethod[coregraphics.ColorRef](t_, "foregroundColor")
	return rv
}

func (t_ TextLayer) SetForegroundColor(value coregraphics.ColorRef) {
	ffi.CallMethod[ffi.Void](t_, "setForegroundColor:", value)
}

func (t_ TextLayer) AllowsFontSubpixelQuantization() bool {
	rv := ffi.CallMethod[bool](t_, "allowsFontSubpixelQuantization")
	return rv
}

func (t_ TextLayer) SetAllowsFontSubpixelQuantization(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAllowsFontSubpixelQuantization:", value)
}

func (t_ TextLayer) IsWrapped() bool {
	rv := ffi.CallMethod[bool](t_, "isWrapped")
	return rv
}

func (t_ TextLayer) SetWrapped(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setWrapped:", value)
}

func (t_ TextLayer) AlignmentMode() TextLayerAlignmentMode {
	rv := ffi.CallMethod[TextLayerAlignmentMode](t_, "alignmentMode")
	return rv
}

func (t_ TextLayer) SetAlignmentMode(value TextLayerAlignmentMode) {
	ffi.CallMethod[ffi.Void](t_, "setAlignmentMode:", value)
}

func (t_ TextLayer) TruncationMode() TextLayerTruncationMode {
	rv := ffi.CallMethod[TextLayerTruncationMode](t_, "truncationMode")
	return rv
}

func (t_ TextLayer) SetTruncationMode(value TextLayerTruncationMode) {
	ffi.CallMethod[ffi.Void](t_, "setTruncationMode:", value)
}
