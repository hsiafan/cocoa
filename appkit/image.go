// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var ImageClass = _ImageClass{objc.GetClass("NSImage")}

type _ImageClass struct {
	objc.Class
}

type IImage interface {
	objc.IObject
	SetName(string_ ImageName) bool
	Name() ImageName
	ImageWithSymbolConfiguration(configuration IImageSymbolConfiguration) Image
	AddRepresentation(imageRep IImageRep)
	AddRepresentations(imageReps []IImageRep)
	RemoveRepresentation(imageRep IImageRep)
	BestRepresentationForRect_Context_Hints(rect foundation.Rect, referenceContext IGraphicsContext, hints map[ImageHintKey]objc.IObject) ImageRep
	DrawInRect(rect foundation.Rect)
	DrawAtPoint_FromRect_Operation_Fraction(point foundation.Point, fromRect foundation.Rect, op CompositingOperation, delta float64)
	DrawInRect_FromRect_Operation_Fraction(rect foundation.Rect, fromRect foundation.Rect, op CompositingOperation, delta float64)
	DrawInRect_FromRect_Operation_Fraction_RespectFlipped_Hints(dstSpacePortionRect foundation.Rect, srcSpacePortionRect foundation.Rect, op CompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints map[ImageHintKey]objc.IObject)
	DrawRepresentation_InRect(imageRep IImageRep, rect foundation.Rect) bool
	Recache()
	TIFFRepresentationUsingCompression_Factor(comp TIFFCompression, factor float32) []byte
	CGImageForProposedRect_Context_Hints(proposedDestRect *foundation.Rect, referenceContext IGraphicsContext, hints map[ImageHintKey]objc.IObject) coregraphics.ImageRef
	CancelIncrementalLoad()
	HitTestRect_WithImageDestinationRect_Context_Hints_Flipped(testRectDestSpace foundation.Rect, imageRectDestSpace foundation.Rect, context IGraphicsContext, hints map[ImageHintKey]objc.IObject, flipped bool) bool
	LayerContentsForContentsScale(layerContentsScale float64) objc.Object
	RecommendedLayerContentsScale(preferredContentsScale float64) float64
	// deprecated
	LockFocus()
	// deprecated
	LockFocusFlipped(flipped bool)
	// deprecated
	UnlockFocus()
	// deprecated
	LockFocusOnRepresentation(imageRepresentation IImageRep)
	// deprecated
	CompositeToPoint_Operation(point foundation.Point, op CompositingOperation)
	// deprecated
	CompositeToPoint_FromRect_Operation(point foundation.Point, rect foundation.Rect, op CompositingOperation)
	// deprecated
	CompositeToPoint_FromRect_Operation_Fraction(point foundation.Point, rect foundation.Rect, op CompositingOperation, delta float64)
	// deprecated
	CompositeToPoint_Operation_Fraction(point foundation.Point, op CompositingOperation, delta float64)
	// deprecated
	DissolveToPoint_Fraction(point foundation.Point, fraction float64)
	// deprecated
	DissolveToPoint_FromRect_Fraction(point foundation.Point, rect foundation.Rect, fraction float64)
	// deprecated
	SetScalesWhenResized(flag bool)
	// deprecated
	ScalesWhenResized() bool
	// deprecated
	SetDataRetained(flag bool)
	// deprecated
	IsDataRetained() bool
	// deprecated
	SetCachedSeparately(flag bool)
	// deprecated
	IsCachedSeparately() bool
	// deprecated
	SetCacheDepthMatchesImageDepth(flag bool)
	// deprecated
	CacheDepthMatchesImageDepth() bool
	// deprecated
	SetFlipped(flag bool)
	// deprecated
	IsFlipped() bool
	SymbolConfiguration() ImageSymbolConfiguration
	Delegate() ImageDelegateWrapper
	SetDelegate(value ImageDelegate)
	SetDelegate0(value objc.IObject)
	Size() foundation.Size
	SetSize(value foundation.Size)
	IsTemplate() bool
	SetTemplate(value bool)
	Representations() []ImageRep
	PrefersColorMatch() bool
	SetPrefersColorMatch(value bool)
	UsesEPSOnResolutionMismatch() bool
	SetUsesEPSOnResolutionMismatch(value bool)
	MatchesOnMultipleResolution() bool
	SetMatchesOnMultipleResolution(value bool)
	IsValid() bool
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	CapInsets() foundation.EdgeInsets
	SetCapInsets(value foundation.EdgeInsets)
	ResizingMode() ImageResizingMode
	SetResizingMode(value ImageResizingMode)
	AlignmentRect() foundation.Rect
	SetAlignmentRect(value foundation.Rect)
	CacheMode() ImageCacheMode
	SetCacheMode(value ImageCacheMode)
	TIFFRepresentation() []byte
	AccessibilityDescription() string
	SetAccessibilityDescription(value string)
	MatchesOnlyOnBestFittingAxis() bool
	SetMatchesOnlyOnBestFittingAxis(value bool)
}

type Image struct {
	objc.Object
}

func MakeImage(ptr unsafe.Pointer) Image {
	return Image{
		Object: objc.MakeObject(ptr),
	}
}

func (ic _ImageClass) ImageWithSystemSymbolName_AccessibilityDescription(symbolName string, description string) Image {
	rv := ffi.CallMethod[Image](ic, "imageWithSystemSymbolName:accessibilityDescription:", symbolName, description)
	return rv
}

func (ic _ImageClass) ImageWithSystemSymbolName_VariableValue_AccessibilityDescription(name string, value float64, description string) Image {
	rv := ffi.CallMethod[Image](ic, "imageWithSystemSymbolName:variableValue:accessibilityDescription:", name, value, description)
	return rv
}

func (ic _ImageClass) ImageWithSymbolName_VariableValue(name string, value float64) Image {
	rv := ffi.CallMethod[Image](ic, "imageWithSymbolName:variableValue:", name, value)
	return rv
}

func (ic _ImageClass) ImageWithSize_Flipped_DrawingHandler(size foundation.Size, drawingHandlerShouldBeCalledWithFlippedContext bool, drawingHandler func(dstRect foundation.Rect) bool) Image {
	rv := ffi.CallMethod[Image](ic, "imageWithSize:flipped:drawingHandler:", size, drawingHandlerShouldBeCalledWithFlippedContext, drawingHandler)
	return rv
}

func (i_ Image) InitByReferencingFile(fileName string) Image {
	rv := ffi.CallMethod[Image](i_, "initByReferencingFile:", fileName)
	return rv
}

func (i_ Image) InitByReferencingURL(url foundation.IURL) Image {
	rv := ffi.CallMethod[Image](i_, "initByReferencingURL:", url)
	return rv
}

func (i_ Image) InitWithContentsOfFile(fileName string) Image {
	rv := ffi.CallMethod[Image](i_, "initWithContentsOfFile:", fileName)
	return rv
}

func (i_ Image) InitWithContentsOfURL(url foundation.IURL) Image {
	rv := ffi.CallMethod[Image](i_, "initWithContentsOfURL:", url)
	return rv
}

func (i_ Image) InitWithData(data []byte) Image {
	rv := ffi.CallMethod[Image](i_, "initWithData:", data)
	return rv
}

func (i_ Image) InitWithDataIgnoringOrientation(data []byte) Image {
	rv := ffi.CallMethod[Image](i_, "initWithDataIgnoringOrientation:", data)
	return rv
}

func (i_ Image) InitWithCGImage_Size(cgImage coregraphics.ImageRef, size foundation.Size) Image {
	rv := ffi.CallMethod[Image](i_, "initWithCGImage:size:", cgImage, size)
	return rv
}

func (i_ Image) InitWithPasteboard(pasteboard IPasteboard) Image {
	rv := ffi.CallMethod[Image](i_, "initWithPasteboard:", pasteboard)
	return rv
}

func (i_ Image) InitWithSize(size foundation.Size) Image {
	rv := ffi.CallMethod[Image](i_, "initWithSize:", size)
	return rv
}

func (ic _ImageClass) Alloc() Image {
	rv := ffi.CallMethod[Image](ic, "alloc")
	return rv
}

func (ic _ImageClass) New() Image {
	rv := ffi.CallMethod[Image](ic, "new")
	rv.Autorelease()
	return rv
}

func NewImage() Image {
	return ImageClass.New()
}

func (i_ Image) Init() Image {
	rv := ffi.CallMethod[Image](i_, "init")
	return rv
}

func (ic _ImageClass) ImageNamed(name ImageName) Image {
	rv := ffi.CallMethod[Image](ic, "imageNamed:", name)
	return rv
}

func (i_ Image) SetName(string_ ImageName) bool {
	rv := ffi.CallMethod[bool](i_, "setName:", string_)
	return rv
}

func (i_ Image) Name() ImageName {
	rv := ffi.CallMethod[ImageName](i_, "name")
	return rv
}

func (i_ Image) ImageWithSymbolConfiguration(configuration IImageSymbolConfiguration) Image {
	rv := ffi.CallMethod[Image](i_, "imageWithSymbolConfiguration:", configuration)
	return rv
}

func (ic _ImageClass) CanInitWithPasteboard(pasteboard IPasteboard) bool {
	rv := ffi.CallMethod[bool](ic, "canInitWithPasteboard:", pasteboard)
	return rv
}

func (i_ Image) AddRepresentation(imageRep IImageRep) {
	ffi.CallMethod[ffi.Void](i_, "addRepresentation:", imageRep)
}

func (i_ Image) AddRepresentations(imageReps []IImageRep) {
	ffi.CallMethod[ffi.Void](i_, "addRepresentations:", imageReps)
}

func (i_ Image) RemoveRepresentation(imageRep IImageRep) {
	ffi.CallMethod[ffi.Void](i_, "removeRepresentation:", imageRep)
}

func (i_ Image) BestRepresentationForRect_Context_Hints(rect foundation.Rect, referenceContext IGraphicsContext, hints map[ImageHintKey]objc.IObject) ImageRep {
	rv := ffi.CallMethod[ImageRep](i_, "bestRepresentationForRect:context:hints:", rect, referenceContext, hints)
	return rv
}

func (i_ Image) DrawInRect(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](i_, "drawInRect:", rect)
}

func (i_ Image) DrawAtPoint_FromRect_Operation_Fraction(point foundation.Point, fromRect foundation.Rect, op CompositingOperation, delta float64) {
	ffi.CallMethod[ffi.Void](i_, "drawAtPoint:fromRect:operation:fraction:", point, fromRect, op, delta)
}

func (i_ Image) DrawInRect_FromRect_Operation_Fraction(rect foundation.Rect, fromRect foundation.Rect, op CompositingOperation, delta float64) {
	ffi.CallMethod[ffi.Void](i_, "drawInRect:fromRect:operation:fraction:", rect, fromRect, op, delta)
}

func (i_ Image) DrawInRect_FromRect_Operation_Fraction_RespectFlipped_Hints(dstSpacePortionRect foundation.Rect, srcSpacePortionRect foundation.Rect, op CompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints map[ImageHintKey]objc.IObject) {
	ffi.CallMethod[ffi.Void](i_, "drawInRect:fromRect:operation:fraction:respectFlipped:hints:", dstSpacePortionRect, srcSpacePortionRect, op, requestedAlpha, respectContextIsFlipped, hints)
}

func (i_ Image) DrawRepresentation_InRect(imageRep IImageRep, rect foundation.Rect) bool {
	rv := ffi.CallMethod[bool](i_, "drawRepresentation:inRect:", imageRep, rect)
	return rv
}

func (i_ Image) Recache() {
	ffi.CallMethod[ffi.Void](i_, "recache")
}

func (i_ Image) TIFFRepresentationUsingCompression_Factor(comp TIFFCompression, factor float32) []byte {
	rv := ffi.CallMethod[[]byte](i_, "TIFFRepresentationUsingCompression:factor:", comp, factor)
	return rv
}

func (i_ Image) CGImageForProposedRect_Context_Hints(proposedDestRect *foundation.Rect, referenceContext IGraphicsContext, hints map[ImageHintKey]objc.IObject) coregraphics.ImageRef {
	rv := ffi.CallMethod[coregraphics.ImageRef](i_, "CGImageForProposedRect:context:hints:", proposedDestRect, referenceContext, hints)
	return rv
}

func (i_ Image) CancelIncrementalLoad() {
	ffi.CallMethod[ffi.Void](i_, "cancelIncrementalLoad")
}

func (i_ Image) HitTestRect_WithImageDestinationRect_Context_Hints_Flipped(testRectDestSpace foundation.Rect, imageRectDestSpace foundation.Rect, context IGraphicsContext, hints map[ImageHintKey]objc.IObject, flipped bool) bool {
	rv := ffi.CallMethod[bool](i_, "hitTestRect:withImageDestinationRect:context:hints:flipped:", testRectDestSpace, imageRectDestSpace, context, hints, flipped)
	return rv
}

func (i_ Image) LayerContentsForContentsScale(layerContentsScale float64) objc.Object {
	rv := ffi.CallMethod[objc.Object](i_, "layerContentsForContentsScale:", layerContentsScale)
	return rv
}

func (i_ Image) RecommendedLayerContentsScale(preferredContentsScale float64) float64 {
	rv := ffi.CallMethod[float64](i_, "recommendedLayerContentsScale:", preferredContentsScale)
	return rv
}

// deprecated
func (ic _ImageClass) ImageFileTypes() []string {
	rv := ffi.CallMethod[[]string](ic, "imageFileTypes")
	return rv
}

// deprecated
func (ic _ImageClass) ImageUnfilteredFileTypes() []string {
	rv := ffi.CallMethod[[]string](ic, "imageUnfilteredFileTypes")
	return rv
}

// deprecated
func (ic _ImageClass) ImagePasteboardTypes() []PasteboardType {
	rv := ffi.CallMethod[[]PasteboardType](ic, "imagePasteboardTypes")
	return rv
}

// deprecated
func (ic _ImageClass) ImageUnfilteredPasteboardTypes() []PasteboardType {
	rv := ffi.CallMethod[[]PasteboardType](ic, "imageUnfilteredPasteboardTypes")
	return rv
}

// deprecated
func (i_ Image) LockFocus() {
	ffi.CallMethod[ffi.Void](i_, "lockFocus")
}

// deprecated
func (i_ Image) LockFocusFlipped(flipped bool) {
	ffi.CallMethod[ffi.Void](i_, "lockFocusFlipped:", flipped)
}

// deprecated
func (i_ Image) UnlockFocus() {
	ffi.CallMethod[ffi.Void](i_, "unlockFocus")
}

// deprecated
func (i_ Image) LockFocusOnRepresentation(imageRepresentation IImageRep) {
	ffi.CallMethod[ffi.Void](i_, "lockFocusOnRepresentation:", imageRepresentation)
}

// deprecated
func (i_ Image) CompositeToPoint_Operation(point foundation.Point, op CompositingOperation) {
	ffi.CallMethod[ffi.Void](i_, "compositeToPoint:operation:", point, op)
}

// deprecated
func (i_ Image) CompositeToPoint_FromRect_Operation(point foundation.Point, rect foundation.Rect, op CompositingOperation) {
	ffi.CallMethod[ffi.Void](i_, "compositeToPoint:fromRect:operation:", point, rect, op)
}

// deprecated
func (i_ Image) CompositeToPoint_FromRect_Operation_Fraction(point foundation.Point, rect foundation.Rect, op CompositingOperation, delta float64) {
	ffi.CallMethod[ffi.Void](i_, "compositeToPoint:fromRect:operation:fraction:", point, rect, op, delta)
}

// deprecated
func (i_ Image) CompositeToPoint_Operation_Fraction(point foundation.Point, op CompositingOperation, delta float64) {
	ffi.CallMethod[ffi.Void](i_, "compositeToPoint:operation:fraction:", point, op, delta)
}

// deprecated
func (i_ Image) DissolveToPoint_Fraction(point foundation.Point, fraction float64) {
	ffi.CallMethod[ffi.Void](i_, "dissolveToPoint:fraction:", point, fraction)
}

// deprecated
func (i_ Image) DissolveToPoint_FromRect_Fraction(point foundation.Point, rect foundation.Rect, fraction float64) {
	ffi.CallMethod[ffi.Void](i_, "dissolveToPoint:fromRect:fraction:", point, rect, fraction)
}

// deprecated
func (i_ Image) SetScalesWhenResized(flag bool) {
	ffi.CallMethod[ffi.Void](i_, "setScalesWhenResized:", flag)
}

// deprecated
func (i_ Image) ScalesWhenResized() bool {
	rv := ffi.CallMethod[bool](i_, "scalesWhenResized")
	return rv
}

// deprecated
func (i_ Image) SetDataRetained(flag bool) {
	ffi.CallMethod[ffi.Void](i_, "setDataRetained:", flag)
}

// deprecated
func (i_ Image) IsDataRetained() bool {
	rv := ffi.CallMethod[bool](i_, "isDataRetained")
	return rv
}

// deprecated
func (i_ Image) SetCachedSeparately(flag bool) {
	ffi.CallMethod[ffi.Void](i_, "setCachedSeparately:", flag)
}

// deprecated
func (i_ Image) IsCachedSeparately() bool {
	rv := ffi.CallMethod[bool](i_, "isCachedSeparately")
	return rv
}

// deprecated
func (i_ Image) SetCacheDepthMatchesImageDepth(flag bool) {
	ffi.CallMethod[ffi.Void](i_, "setCacheDepthMatchesImageDepth:", flag)
}

// deprecated
func (i_ Image) CacheDepthMatchesImageDepth() bool {
	rv := ffi.CallMethod[bool](i_, "cacheDepthMatchesImageDepth")
	return rv
}

// deprecated
func (i_ Image) SetFlipped(flag bool) {
	ffi.CallMethod[ffi.Void](i_, "setFlipped:", flag)
}

// deprecated
func (i_ Image) IsFlipped() bool {
	rv := ffi.CallMethod[bool](i_, "isFlipped")
	return rv
}

func (i_ Image) SymbolConfiguration() ImageSymbolConfiguration {
	rv := ffi.CallMethod[ImageSymbolConfiguration](i_, "symbolConfiguration")
	return rv
}

func (i_ Image) Delegate() ImageDelegateWrapper {
	rv := ffi.CallMethod[ImageDelegateWrapper](i_, "delegate")
	return rv
}

func (i_ Image) SetDelegate(value ImageDelegate) {
	po := ffi.CreateProtocol("NSImageDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(i_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](i_, "setDelegate:", po)
}

func (i_ Image) SetDelegate0(value objc.IObject) {
	ffi.CallMethod[ffi.Void](i_, "setDelegate:", value)
}

func (i_ Image) Size() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](i_, "size")
	return rv
}

func (i_ Image) SetSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](i_, "setSize:", value)
}

func (i_ Image) IsTemplate() bool {
	rv := ffi.CallMethod[bool](i_, "isTemplate")
	return rv
}

func (i_ Image) SetTemplate(value bool) {
	ffi.CallMethod[ffi.Void](i_, "setTemplate:", value)
}

func (ic _ImageClass) ImageTypes() []string {
	rv := ffi.CallMethod[[]string](ic, "imageTypes")
	return rv
}

func (ic _ImageClass) ImageUnfilteredTypes() []string {
	rv := ffi.CallMethod[[]string](ic, "imageUnfilteredTypes")
	return rv
}

func (i_ Image) Representations() []ImageRep {
	rv := ffi.CallMethod[[]ImageRep](i_, "representations")
	return rv
}

func (i_ Image) PrefersColorMatch() bool {
	rv := ffi.CallMethod[bool](i_, "prefersColorMatch")
	return rv
}

func (i_ Image) SetPrefersColorMatch(value bool) {
	ffi.CallMethod[ffi.Void](i_, "setPrefersColorMatch:", value)
}

func (i_ Image) UsesEPSOnResolutionMismatch() bool {
	rv := ffi.CallMethod[bool](i_, "usesEPSOnResolutionMismatch")
	return rv
}

func (i_ Image) SetUsesEPSOnResolutionMismatch(value bool) {
	ffi.CallMethod[ffi.Void](i_, "setUsesEPSOnResolutionMismatch:", value)
}

func (i_ Image) MatchesOnMultipleResolution() bool {
	rv := ffi.CallMethod[bool](i_, "matchesOnMultipleResolution")
	return rv
}

func (i_ Image) SetMatchesOnMultipleResolution(value bool) {
	ffi.CallMethod[ffi.Void](i_, "setMatchesOnMultipleResolution:", value)
}

func (i_ Image) IsValid() bool {
	rv := ffi.CallMethod[bool](i_, "isValid")
	return rv
}

func (i_ Image) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](i_, "backgroundColor")
	return rv
}

func (i_ Image) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](i_, "setBackgroundColor:", value)
}

func (i_ Image) CapInsets() foundation.EdgeInsets {
	rv := ffi.CallMethod[foundation.EdgeInsets](i_, "capInsets")
	return rv
}

func (i_ Image) SetCapInsets(value foundation.EdgeInsets) {
	ffi.CallMethod[ffi.Void](i_, "setCapInsets:", value)
}

func (i_ Image) ResizingMode() ImageResizingMode {
	rv := ffi.CallMethod[ImageResizingMode](i_, "resizingMode")
	return rv
}

func (i_ Image) SetResizingMode(value ImageResizingMode) {
	ffi.CallMethod[ffi.Void](i_, "setResizingMode:", value)
}

func (i_ Image) AlignmentRect() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](i_, "alignmentRect")
	return rv
}

func (i_ Image) SetAlignmentRect(value foundation.Rect) {
	ffi.CallMethod[ffi.Void](i_, "setAlignmentRect:", value)
}

func (i_ Image) CacheMode() ImageCacheMode {
	rv := ffi.CallMethod[ImageCacheMode](i_, "cacheMode")
	return rv
}

func (i_ Image) SetCacheMode(value ImageCacheMode) {
	ffi.CallMethod[ffi.Void](i_, "setCacheMode:", value)
}

func (i_ Image) TIFFRepresentation() []byte {
	rv := ffi.CallMethod[[]byte](i_, "TIFFRepresentation")
	return rv
}

func (i_ Image) AccessibilityDescription() string {
	rv := ffi.CallMethod[string](i_, "accessibilityDescription")
	return rv
}

func (i_ Image) SetAccessibilityDescription(value string) {
	ffi.CallMethod[ffi.Void](i_, "setAccessibilityDescription:", value)
}

func (i_ Image) MatchesOnlyOnBestFittingAxis() bool {
	rv := ffi.CallMethod[bool](i_, "matchesOnlyOnBestFittingAxis")
	return rv
}

func (i_ Image) SetMatchesOnlyOnBestFittingAxis(value bool) {
	ffi.CallMethod[ffi.Void](i_, "setMatchesOnlyOnBestFittingAxis:", value)
}
