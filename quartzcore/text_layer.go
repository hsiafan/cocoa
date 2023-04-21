// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
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
	rv := objc.CallMethod[TextLayer](tc, "layer")
	return rv
}

func (t_ TextLayer) Init() TextLayer {
	rv := objc.CallMethod[TextLayer](t_, "init")
	return rv
}

func (t_ TextLayer) InitWithLayer(layer objc.IObject) TextLayer {
	rv := objc.CallMethod[TextLayer](t_, "initWithLayer:", layer)
	return rv
}

func (t_ TextLayer) PresentationLayer() TextLayer {
	rv := objc.CallMethod[TextLayer](t_, "presentationLayer")
	return rv
}

func (t_ TextLayer) ModelLayer() TextLayer {
	rv := objc.CallMethod[TextLayer](t_, "modelLayer")
	return rv
}

func (tc _TextLayerClass) Alloc() TextLayer {
	rv := objc.CallMethod[TextLayer](tc, "alloc")
	return rv
}

func (tc _TextLayerClass) New() TextLayer {
	rv := objc.CallMethod[TextLayer](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextLayer() TextLayer {
	return TextLayerClass.New()
}

func (t_ TextLayer) String() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, "string")
	return rv
}

func (t_ TextLayer) SetString(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, "setString:", value)
}

func (t_ TextLayer) FontSize() float64 {
	rv := objc.CallMethod[float64](t_, "fontSize")
	return rv
}

func (t_ TextLayer) SetFontSize(value float64) {
	objc.CallMethod[objc.Void](t_, "setFontSize:", value)
}

func (t_ TextLayer) ForegroundColor() coregraphics.ColorRef {
	rv := objc.CallMethod[coregraphics.ColorRef](t_, "foregroundColor")
	return rv
}

func (t_ TextLayer) SetForegroundColor(value coregraphics.ColorRef) {
	objc.CallMethod[objc.Void](t_, "setForegroundColor:", value)
}

func (t_ TextLayer) AllowsFontSubpixelQuantization() bool {
	rv := objc.CallMethod[bool](t_, "allowsFontSubpixelQuantization")
	return rv
}

func (t_ TextLayer) SetAllowsFontSubpixelQuantization(value bool) {
	objc.CallMethod[objc.Void](t_, "setAllowsFontSubpixelQuantization:", value)
}

func (t_ TextLayer) IsWrapped() bool {
	rv := objc.CallMethod[bool](t_, "isWrapped")
	return rv
}

func (t_ TextLayer) SetWrapped(value bool) {
	objc.CallMethod[objc.Void](t_, "setWrapped:", value)
}

func (t_ TextLayer) AlignmentMode() TextLayerAlignmentMode {
	rv := objc.CallMethod[TextLayerAlignmentMode](t_, "alignmentMode")
	return rv
}

func (t_ TextLayer) SetAlignmentMode(value TextLayerAlignmentMode) {
	objc.CallMethod[objc.Void](t_, "setAlignmentMode:", value)
}

func (t_ TextLayer) TruncationMode() TextLayerTruncationMode {
	rv := objc.CallMethod[TextLayerTruncationMode](t_, "truncationMode")
	return rv
}

func (t_ TextLayer) SetTruncationMode(value TextLayerTruncationMode) {
	objc.CallMethod[objc.Void](t_, "setTruncationMode:", value)
}
