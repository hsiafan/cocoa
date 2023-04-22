// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ImageRepClass = _ImageRepClass{objc.GetClass("NSImageRep")}

type _ImageRepClass struct {
	objc.Class
}

type IImageRep interface {
	objc.IObject
	CGImageForProposedRect_Context_Hints(proposedDestRect *foundation.Rect, context IGraphicsContext, hints map[ImageHintKey]objc.IObject) coregraphics.ImageRef
	Draw() bool
	DrawAtPoint(point foundation.Point) bool
	DrawInRect(rect foundation.Rect) bool
	DrawInRect_FromRect_Operation_Fraction_RespectFlipped_Hints(dstSpacePortionRect foundation.Rect, srcSpacePortionRect foundation.Rect, op CompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints map[ImageHintKey]objc.IObject) bool
	Size() foundation.Size
	SetSize(value foundation.Size)
	BitsPerSample() int
	SetBitsPerSample(value int)
	ColorSpaceName() ColorSpaceName
	SetColorSpaceName(value ColorSpaceName)
	HasAlpha() bool
	SetAlpha(value bool)
	IsOpaque() bool
	SetOpaque(value bool)
	PixelsHigh() int
	SetPixelsHigh(value int)
	PixelsWide() int
	SetPixelsWide(value int)
	LayoutDirection() ImageLayoutDirection
	SetLayoutDirection(value ImageLayoutDirection)
}

type ImageRep struct {
	objc.Object
}

func MakeImageRep(ptr unsafe.Pointer) ImageRep {
	return ImageRep{
		Object: objc.MakeObject(ptr),
	}
}

func (i_ ImageRep) Init() ImageRep {
	rv := objc.CallMethod[ImageRep](i_, objc.GetSelector("init"))
	return rv
}

func (ic _ImageRepClass) Alloc() ImageRep {
	rv := objc.CallMethod[ImageRep](ic, objc.GetSelector("alloc"))
	return rv
}

func (ic _ImageRepClass) New() ImageRep {
	rv := objc.CallMethod[ImageRep](ic, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewImageRep() ImageRep {
	return ImageRepClass.New()
}

func (ic _ImageRepClass) ImageRepsWithContentsOfFile(filename string) []ImageRep {
	rv := objc.CallMethod[[]ImageRep](ic, objc.GetSelector("imageRepsWithContentsOfFile:"), filename)
	return rv
}

func (ic _ImageRepClass) ImageRepsWithPasteboard(pasteboard IPasteboard) []ImageRep {
	rv := objc.CallMethod[[]ImageRep](ic, objc.GetSelector("imageRepsWithPasteboard:"), pasteboard)
	return rv
}

func (ic _ImageRepClass) ImageRepsWithContentsOfURL(url foundation.IURL) []ImageRep {
	rv := objc.CallMethod[[]ImageRep](ic, objc.GetSelector("imageRepsWithContentsOfURL:"), url)
	return rv
}

func (ic _ImageRepClass) ImageRepWithContentsOfFile(filename string) ImageRep {
	rv := objc.CallMethod[ImageRep](ic, objc.GetSelector("imageRepWithContentsOfFile:"), filename)
	return rv
}

func (ic _ImageRepClass) ImageRepWithPasteboard(pasteboard IPasteboard) ImageRep {
	rv := objc.CallMethod[ImageRep](ic, objc.GetSelector("imageRepWithPasteboard:"), pasteboard)
	return rv
}

func (ic _ImageRepClass) ImageRepWithContentsOfURL(url foundation.IURL) ImageRep {
	rv := objc.CallMethod[ImageRep](ic, objc.GetSelector("imageRepWithContentsOfURL:"), url)
	return rv
}

func (ic _ImageRepClass) CanInitWithData(data []byte) bool {
	rv := objc.CallMethod[bool](ic, objc.GetSelector("canInitWithData:"), data)
	return rv
}

func (ic _ImageRepClass) CanInitWithPasteboard(pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](ic, objc.GetSelector("canInitWithPasteboard:"), pasteboard)
	return rv
}

// deprecated
func (ic _ImageRepClass) ImageFileTypes() []string {
	rv := objc.CallMethod[[]string](ic, objc.GetSelector("imageFileTypes"))
	return rv
}

// deprecated
func (ic _ImageRepClass) ImagePasteboardTypes() []PasteboardType {
	rv := objc.CallMethod[[]PasteboardType](ic, objc.GetSelector("imagePasteboardTypes"))
	return rv
}

// deprecated
func (ic _ImageRepClass) ImageUnfilteredFileTypes() []string {
	rv := objc.CallMethod[[]string](ic, objc.GetSelector("imageUnfilteredFileTypes"))
	return rv
}

// deprecated
func (ic _ImageRepClass) ImageUnfilteredPasteboardTypes() []PasteboardType {
	rv := objc.CallMethod[[]PasteboardType](ic, objc.GetSelector("imageUnfilteredPasteboardTypes"))
	return rv
}

func (i_ ImageRep) CGImageForProposedRect_Context_Hints(proposedDestRect *foundation.Rect, context IGraphicsContext, hints map[ImageHintKey]objc.IObject) coregraphics.ImageRef {
	rv := objc.CallMethod[coregraphics.ImageRef](i_, objc.GetSelector("CGImageForProposedRect:context:hints:"), proposedDestRect, context, hints)
	return rv
}

func (i_ ImageRep) Draw() bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("draw"))
	return rv
}

func (i_ ImageRep) DrawAtPoint(point foundation.Point) bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("drawAtPoint:"), point)
	return rv
}

func (i_ ImageRep) DrawInRect(rect foundation.Rect) bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("drawInRect:"), rect)
	return rv
}

func (i_ ImageRep) DrawInRect_FromRect_Operation_Fraction_RespectFlipped_Hints(dstSpacePortionRect foundation.Rect, srcSpacePortionRect foundation.Rect, op CompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints map[ImageHintKey]objc.IObject) bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("drawInRect:fromRect:operation:fraction:respectFlipped:hints:"), dstSpacePortionRect, srcSpacePortionRect, op, requestedAlpha, respectContextIsFlipped, hints)
	return rv
}

func (ic _ImageRepClass) ImageRepClassForType(type_ string) objc.Class {
	rv := objc.CallMethod[objc.Class](ic, objc.GetSelector("imageRepClassForType:"), type_)
	return rv
}

func (ic _ImageRepClass) ImageRepClassForData(data []byte) objc.Class {
	rv := objc.CallMethod[objc.Class](ic, objc.GetSelector("imageRepClassForData:"), data)
	return rv
}

func (ic _ImageRepClass) RegisterImageRepClass(imageRepClass objc.IClass) {
	objc.CallMethod[objc.Void](ic, objc.GetSelector("registerImageRepClass:"), imageRepClass)
}

func (ic _ImageRepClass) UnregisterImageRepClass(imageRepClass objc.IClass) {
	objc.CallMethod[objc.Void](ic, objc.GetSelector("unregisterImageRepClass:"), imageRepClass)
}

// deprecated
func (ic _ImageRepClass) ImageRepClassForFileType(type_ string) objc.Class {
	rv := objc.CallMethod[objc.Class](ic, objc.GetSelector("imageRepClassForFileType:"), type_)
	return rv
}

// deprecated
func (ic _ImageRepClass) ImageRepClassForPasteboardType(type_ PasteboardType) objc.Class {
	rv := objc.CallMethod[objc.Class](ic, objc.GetSelector("imageRepClassForPasteboardType:"), type_)
	return rv
}

func (ic _ImageRepClass) ImageTypes() []string {
	rv := objc.CallMethod[[]string](ic, objc.GetSelector("imageTypes"))
	return rv
}

func (ic _ImageRepClass) ImageUnfilteredTypes() []string {
	rv := objc.CallMethod[[]string](ic, objc.GetSelector("imageUnfilteredTypes"))
	return rv
}

func (i_ ImageRep) Size() foundation.Size {
	rv := objc.CallMethod[foundation.Size](i_, objc.GetSelector("size"))
	return rv
}

func (i_ ImageRep) SetSize(value foundation.Size) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setSize:"), value)
}

func (i_ ImageRep) BitsPerSample() int {
	rv := objc.CallMethod[int](i_, objc.GetSelector("bitsPerSample"))
	return rv
}

func (i_ ImageRep) SetBitsPerSample(value int) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setBitsPerSample:"), value)
}

func (i_ ImageRep) ColorSpaceName() ColorSpaceName {
	rv := objc.CallMethod[ColorSpaceName](i_, objc.GetSelector("colorSpaceName"))
	return rv
}

func (i_ ImageRep) SetColorSpaceName(value ColorSpaceName) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setColorSpaceName:"), value)
}

func (i_ ImageRep) HasAlpha() bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("hasAlpha"))
	return rv
}

func (i_ ImageRep) SetAlpha(value bool) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setAlpha:"), value)
}

func (i_ ImageRep) IsOpaque() bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("isOpaque"))
	return rv
}

func (i_ ImageRep) SetOpaque(value bool) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setOpaque:"), value)
}

func (i_ ImageRep) PixelsHigh() int {
	rv := objc.CallMethod[int](i_, objc.GetSelector("pixelsHigh"))
	return rv
}

func (i_ ImageRep) SetPixelsHigh(value int) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setPixelsHigh:"), value)
}

func (i_ ImageRep) PixelsWide() int {
	rv := objc.CallMethod[int](i_, objc.GetSelector("pixelsWide"))
	return rv
}

func (i_ ImageRep) SetPixelsWide(value int) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setPixelsWide:"), value)
}

func (i_ ImageRep) LayoutDirection() ImageLayoutDirection {
	rv := objc.CallMethod[ImageLayoutDirection](i_, objc.GetSelector("layoutDirection"))
	return rv
}

func (i_ ImageRep) SetLayoutDirection(value ImageLayoutDirection) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("setLayoutDirection:"), value)
}

func (ic _ImageRepClass) RegisteredImageRepClasses() []objc.Class {
	rv := objc.CallMethod[[]objc.Class](ic, objc.GetSelector("registeredImageRepClasses"))
	return rv
}
