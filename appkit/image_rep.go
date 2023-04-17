// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[ImageRep](i_, "init")
	return rv
}

func (ic _ImageRepClass) Alloc() ImageRep {
	rv := ffi.CallMethod[ImageRep](ic, "alloc")
	return rv
}

func (ic _ImageRepClass) New() ImageRep {
	rv := ffi.CallMethod[ImageRep](ic, "new")
	rv.Autorelease()
	return rv
}

func NewImageRep() ImageRep {
	return ImageRepClass.New()
}

func (ic _ImageRepClass) ImageRepsWithContentsOfFile(filename string) []ImageRep {
	rv := ffi.CallMethod[[]ImageRep](ic, "imageRepsWithContentsOfFile:", filename)
	return rv
}

func (ic _ImageRepClass) ImageRepsWithPasteboard(pasteboard IPasteboard) []ImageRep {
	rv := ffi.CallMethod[[]ImageRep](ic, "imageRepsWithPasteboard:", pasteboard)
	return rv
}

func (ic _ImageRepClass) ImageRepsWithContentsOfURL(url foundation.IURL) []ImageRep {
	rv := ffi.CallMethod[[]ImageRep](ic, "imageRepsWithContentsOfURL:", url)
	return rv
}

func (ic _ImageRepClass) ImageRepWithContentsOfFile(filename string) ImageRep {
	rv := ffi.CallMethod[ImageRep](ic, "imageRepWithContentsOfFile:", filename)
	return rv
}

func (ic _ImageRepClass) ImageRepWithPasteboard(pasteboard IPasteboard) ImageRep {
	rv := ffi.CallMethod[ImageRep](ic, "imageRepWithPasteboard:", pasteboard)
	return rv
}

func (ic _ImageRepClass) ImageRepWithContentsOfURL(url foundation.IURL) ImageRep {
	rv := ffi.CallMethod[ImageRep](ic, "imageRepWithContentsOfURL:", url)
	return rv
}

func (ic _ImageRepClass) CanInitWithData(data []byte) bool {
	rv := ffi.CallMethod[bool](ic, "canInitWithData:", data)
	return rv
}

func (ic _ImageRepClass) CanInitWithPasteboard(pasteboard IPasteboard) bool {
	rv := ffi.CallMethod[bool](ic, "canInitWithPasteboard:", pasteboard)
	return rv
}

// deprecated
func (ic _ImageRepClass) ImageFileTypes() []string {
	rv := ffi.CallMethod[[]string](ic, "imageFileTypes")
	return rv
}

// deprecated
func (ic _ImageRepClass) ImagePasteboardTypes() []PasteboardType {
	rv := ffi.CallMethod[[]PasteboardType](ic, "imagePasteboardTypes")
	return rv
}

// deprecated
func (ic _ImageRepClass) ImageUnfilteredFileTypes() []string {
	rv := ffi.CallMethod[[]string](ic, "imageUnfilteredFileTypes")
	return rv
}

// deprecated
func (ic _ImageRepClass) ImageUnfilteredPasteboardTypes() []PasteboardType {
	rv := ffi.CallMethod[[]PasteboardType](ic, "imageUnfilteredPasteboardTypes")
	return rv
}

func (i_ ImageRep) CGImageForProposedRect_Context_Hints(proposedDestRect *foundation.Rect, context IGraphicsContext, hints map[ImageHintKey]objc.IObject) coregraphics.ImageRef {
	rv := ffi.CallMethod[coregraphics.ImageRef](i_, "CGImageForProposedRect:context:hints:", proposedDestRect, context, hints)
	return rv
}

func (i_ ImageRep) Draw() bool {
	rv := ffi.CallMethod[bool](i_, "draw")
	return rv
}

func (i_ ImageRep) DrawAtPoint(point foundation.Point) bool {
	rv := ffi.CallMethod[bool](i_, "drawAtPoint:", point)
	return rv
}

func (i_ ImageRep) DrawInRect(rect foundation.Rect) bool {
	rv := ffi.CallMethod[bool](i_, "drawInRect:", rect)
	return rv
}

func (i_ ImageRep) DrawInRect_FromRect_Operation_Fraction_RespectFlipped_Hints(dstSpacePortionRect foundation.Rect, srcSpacePortionRect foundation.Rect, op CompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints map[ImageHintKey]objc.IObject) bool {
	rv := ffi.CallMethod[bool](i_, "drawInRect:fromRect:operation:fraction:respectFlipped:hints:", dstSpacePortionRect, srcSpacePortionRect, op, requestedAlpha, respectContextIsFlipped, hints)
	return rv
}

func (ic _ImageRepClass) ImageRepClassForType(type_ string) objc.Class {
	rv := ffi.CallMethod[objc.Class](ic, "imageRepClassForType:", type_)
	return rv
}

func (ic _ImageRepClass) ImageRepClassForData(data []byte) objc.Class {
	rv := ffi.CallMethod[objc.Class](ic, "imageRepClassForData:", data)
	return rv
}

func (ic _ImageRepClass) RegisterImageRepClass(imageRepClass objc.IClass) {
	ffi.CallMethod[ffi.Void](ic, "registerImageRepClass:", imageRepClass)
}

func (ic _ImageRepClass) UnregisterImageRepClass(imageRepClass objc.IClass) {
	ffi.CallMethod[ffi.Void](ic, "unregisterImageRepClass:", imageRepClass)
}

// deprecated
func (ic _ImageRepClass) ImageRepClassForFileType(type_ string) objc.Class {
	rv := ffi.CallMethod[objc.Class](ic, "imageRepClassForFileType:", type_)
	return rv
}

// deprecated
func (ic _ImageRepClass) ImageRepClassForPasteboardType(type_ PasteboardType) objc.Class {
	rv := ffi.CallMethod[objc.Class](ic, "imageRepClassForPasteboardType:", type_)
	return rv
}

func (ic _ImageRepClass) ImageTypes() []string {
	rv := ffi.CallMethod[[]string](ic, "imageTypes")
	return rv
}

func (ic _ImageRepClass) ImageUnfilteredTypes() []string {
	rv := ffi.CallMethod[[]string](ic, "imageUnfilteredTypes")
	return rv
}

func (i_ ImageRep) Size() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](i_, "size")
	return rv
}

func (i_ ImageRep) SetSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](i_, "setSize:", value)
}

func (i_ ImageRep) BitsPerSample() int {
	rv := ffi.CallMethod[int](i_, "bitsPerSample")
	return rv
}

func (i_ ImageRep) SetBitsPerSample(value int) {
	ffi.CallMethod[ffi.Void](i_, "setBitsPerSample:", value)
}

func (i_ ImageRep) ColorSpaceName() ColorSpaceName {
	rv := ffi.CallMethod[ColorSpaceName](i_, "colorSpaceName")
	return rv
}

func (i_ ImageRep) SetColorSpaceName(value ColorSpaceName) {
	ffi.CallMethod[ffi.Void](i_, "setColorSpaceName:", value)
}

func (i_ ImageRep) HasAlpha() bool {
	rv := ffi.CallMethod[bool](i_, "hasAlpha")
	return rv
}

func (i_ ImageRep) SetAlpha(value bool) {
	ffi.CallMethod[ffi.Void](i_, "setAlpha:", value)
}

func (i_ ImageRep) IsOpaque() bool {
	rv := ffi.CallMethod[bool](i_, "isOpaque")
	return rv
}

func (i_ ImageRep) SetOpaque(value bool) {
	ffi.CallMethod[ffi.Void](i_, "setOpaque:", value)
}

func (i_ ImageRep) PixelsHigh() int {
	rv := ffi.CallMethod[int](i_, "pixelsHigh")
	return rv
}

func (i_ ImageRep) SetPixelsHigh(value int) {
	ffi.CallMethod[ffi.Void](i_, "setPixelsHigh:", value)
}

func (i_ ImageRep) PixelsWide() int {
	rv := ffi.CallMethod[int](i_, "pixelsWide")
	return rv
}

func (i_ ImageRep) SetPixelsWide(value int) {
	ffi.CallMethod[ffi.Void](i_, "setPixelsWide:", value)
}

func (i_ ImageRep) LayoutDirection() ImageLayoutDirection {
	rv := ffi.CallMethod[ImageLayoutDirection](i_, "layoutDirection")
	return rv
}

func (i_ ImageRep) SetLayoutDirection(value ImageLayoutDirection) {
	ffi.CallMethod[ffi.Void](i_, "setLayoutDirection:", value)
}

func (ic _ImageRepClass) RegisteredImageRepClasses() []objc.Class {
	rv := ffi.CallMethod[[]objc.Class](ic, "registeredImageRepClasses")
	return rv
}