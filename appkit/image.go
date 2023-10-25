// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
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
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
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
	rv := objc.CallMethod[Image](ic, objc.SEL("imageWithSystemSymbolName:accessibilityDescription:"), symbolName, description)
	return rv
}

func (ic _ImageClass) ImageWithSystemSymbolName_VariableValue_AccessibilityDescription(name string, value float64, description string) Image {
	rv := objc.CallMethod[Image](ic, objc.SEL("imageWithSystemSymbolName:variableValue:accessibilityDescription:"), name, value, description)
	return rv
}

func (ic _ImageClass) ImageWithSymbolName_VariableValue(name string, value float64) Image {
	rv := objc.CallMethod[Image](ic, objc.SEL("imageWithSymbolName:variableValue:"), name, value)
	return rv
}

func (ic _ImageClass) ImageWithSize_Flipped_DrawingHandler(size foundation.Size, drawingHandlerShouldBeCalledWithFlippedContext bool, drawingHandler func(dstRect foundation.Rect) bool) Image {
	rv := objc.CallMethod[Image](ic, objc.SEL("imageWithSize:flipped:drawingHandler:"), size, drawingHandlerShouldBeCalledWithFlippedContext, drawingHandler)
	return rv
}

func (i_ Image) InitByReferencingFile(fileName string) Image {
	rv := objc.CallMethod[Image](i_, objc.SEL("initByReferencingFile:"), fileName)
	return rv
}

func (i_ Image) InitByReferencingURL(url foundation.IURL) Image {
	rv := objc.CallMethod[Image](i_, objc.SEL("initByReferencingURL:"), objc.ExtractPtr(url))
	return rv
}

func (i_ Image) InitWithContentsOfFile(fileName string) Image {
	rv := objc.CallMethod[Image](i_, objc.SEL("initWithContentsOfFile:"), fileName)
	return rv
}

func (i_ Image) InitWithContentsOfURL(url foundation.IURL) Image {
	rv := objc.CallMethod[Image](i_, objc.SEL("initWithContentsOfURL:"), objc.ExtractPtr(url))
	return rv
}

func (i_ Image) InitWithData(data []byte) Image {
	rv := objc.CallMethod[Image](i_, objc.SEL("initWithData:"), data)
	return rv
}

func (i_ Image) InitWithDataIgnoringOrientation(data []byte) Image {
	rv := objc.CallMethod[Image](i_, objc.SEL("initWithDataIgnoringOrientation:"), data)
	return rv
}

func (i_ Image) InitWithCGImage_Size(cgImage coregraphics.ImageRef, size foundation.Size) Image {
	rv := objc.CallMethod[Image](i_, objc.SEL("initWithCGImage:size:"), cgImage, size)
	return rv
}

func (i_ Image) InitWithPasteboard(pasteboard IPasteboard) Image {
	rv := objc.CallMethod[Image](i_, objc.SEL("initWithPasteboard:"), objc.ExtractPtr(pasteboard))
	return rv
}

func (i_ Image) InitWithSize(size foundation.Size) Image {
	rv := objc.CallMethod[Image](i_, objc.SEL("initWithSize:"), size)
	return rv
}

func (ic _ImageClass) Alloc() Image {
	rv := objc.CallMethod[Image](ic, objc.SEL("alloc"))
	return rv
}

func (ic _ImageClass) New() Image {
	rv := objc.CallMethod[Image](ic, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewImage() Image {
	return ImageClass.New()
}

func (i_ Image) Init() Image {
	rv := objc.CallMethod[Image](i_, objc.SEL("init"))
	return rv
}

func (ic _ImageClass) ImageNamed(name ImageName) Image {
	rv := objc.CallMethod[Image](ic, objc.SEL("imageNamed:"), name)
	return rv
}

func (i_ Image) SetName(string_ ImageName) bool {
	rv := objc.CallMethod[bool](i_, objc.SEL("setName:"), string_)
	return rv
}

func (i_ Image) Name() ImageName {
	rv := objc.CallMethod[ImageName](i_, objc.SEL("name"))
	return rv
}

func (i_ Image) ImageWithSymbolConfiguration(configuration IImageSymbolConfiguration) Image {
	rv := objc.CallMethod[Image](i_, objc.SEL("imageWithSymbolConfiguration:"), objc.ExtractPtr(configuration))
	return rv
}

func (ic _ImageClass) CanInitWithPasteboard(pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](ic, objc.SEL("canInitWithPasteboard:"), objc.ExtractPtr(pasteboard))
	return rv
}

func (i_ Image) AddRepresentation(imageRep IImageRep) {
	objc.CallMethod[objc.Void](i_, objc.SEL("addRepresentation:"), objc.ExtractPtr(imageRep))
}

func (i_ Image) AddRepresentations(imageReps []IImageRep) {
	objc.CallMethod[objc.Void](i_, objc.SEL("addRepresentations:"), imageReps)
}

func (i_ Image) RemoveRepresentation(imageRep IImageRep) {
	objc.CallMethod[objc.Void](i_, objc.SEL("removeRepresentation:"), objc.ExtractPtr(imageRep))
}

func (i_ Image) BestRepresentationForRect_Context_Hints(rect foundation.Rect, referenceContext IGraphicsContext, hints map[ImageHintKey]objc.IObject) ImageRep {
	rv := objc.CallMethod[ImageRep](i_, objc.SEL("bestRepresentationForRect:context:hints:"), rect, objc.ExtractPtr(referenceContext), hints)
	return rv
}

func (i_ Image) DrawInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](i_, objc.SEL("drawInRect:"), rect)
}

func (i_ Image) DrawAtPoint_FromRect_Operation_Fraction(point foundation.Point, fromRect foundation.Rect, op CompositingOperation, delta float64) {
	objc.CallMethod[objc.Void](i_, objc.SEL("drawAtPoint:fromRect:operation:fraction:"), point, fromRect, op, delta)
}

func (i_ Image) DrawInRect_FromRect_Operation_Fraction(rect foundation.Rect, fromRect foundation.Rect, op CompositingOperation, delta float64) {
	objc.CallMethod[objc.Void](i_, objc.SEL("drawInRect:fromRect:operation:fraction:"), rect, fromRect, op, delta)
}

func (i_ Image) DrawInRect_FromRect_Operation_Fraction_RespectFlipped_Hints(dstSpacePortionRect foundation.Rect, srcSpacePortionRect foundation.Rect, op CompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints map[ImageHintKey]objc.IObject) {
	objc.CallMethod[objc.Void](i_, objc.SEL("drawInRect:fromRect:operation:fraction:respectFlipped:hints:"), dstSpacePortionRect, srcSpacePortionRect, op, requestedAlpha, respectContextIsFlipped, hints)
}

func (i_ Image) DrawRepresentation_InRect(imageRep IImageRep, rect foundation.Rect) bool {
	rv := objc.CallMethod[bool](i_, objc.SEL("drawRepresentation:inRect:"), objc.ExtractPtr(imageRep), rect)
	return rv
}

func (i_ Image) Recache() {
	objc.CallMethod[objc.Void](i_, objc.SEL("recache"))
}

func (i_ Image) TIFFRepresentationUsingCompression_Factor(comp TIFFCompression, factor float32) []byte {
	rv := objc.CallMethod[[]byte](i_, objc.SEL("TIFFRepresentationUsingCompression:factor:"), comp, factor)
	return rv
}

func (i_ Image) CGImageForProposedRect_Context_Hints(proposedDestRect *foundation.Rect, referenceContext IGraphicsContext, hints map[ImageHintKey]objc.IObject) coregraphics.ImageRef {
	rv := objc.CallMethod[coregraphics.ImageRef](i_, objc.SEL("CGImageForProposedRect:context:hints:"), proposedDestRect, objc.ExtractPtr(referenceContext), hints)
	return rv
}

func (i_ Image) CancelIncrementalLoad() {
	objc.CallMethod[objc.Void](i_, objc.SEL("cancelIncrementalLoad"))
}

func (i_ Image) HitTestRect_WithImageDestinationRect_Context_Hints_Flipped(testRectDestSpace foundation.Rect, imageRectDestSpace foundation.Rect, context IGraphicsContext, hints map[ImageHintKey]objc.IObject, flipped bool) bool {
	rv := objc.CallMethod[bool](i_, objc.SEL("hitTestRect:withImageDestinationRect:context:hints:flipped:"), testRectDestSpace, imageRectDestSpace, objc.ExtractPtr(context), hints, flipped)
	return rv
}

func (i_ Image) LayerContentsForContentsScale(layerContentsScale float64) objc.Object {
	rv := objc.CallMethod[objc.Object](i_, objc.SEL("layerContentsForContentsScale:"), layerContentsScale)
	return rv
}

func (i_ Image) RecommendedLayerContentsScale(preferredContentsScale float64) float64 {
	rv := objc.CallMethod[float64](i_, objc.SEL("recommendedLayerContentsScale:"), preferredContentsScale)
	return rv
}

// deprecated
func (ic _ImageClass) ImageFileTypes() []string {
	rv := objc.CallMethod[[]string](ic, objc.SEL("imageFileTypes"))
	return rv
}

// deprecated
func (ic _ImageClass) ImageUnfilteredFileTypes() []string {
	rv := objc.CallMethod[[]string](ic, objc.SEL("imageUnfilteredFileTypes"))
	return rv
}

// deprecated
func (ic _ImageClass) ImagePasteboardTypes() []PasteboardType {
	rv := objc.CallMethod[[]PasteboardType](ic, objc.SEL("imagePasteboardTypes"))
	return rv
}

// deprecated
func (ic _ImageClass) ImageUnfilteredPasteboardTypes() []PasteboardType {
	rv := objc.CallMethod[[]PasteboardType](ic, objc.SEL("imageUnfilteredPasteboardTypes"))
	return rv
}

// deprecated
func (i_ Image) LockFocus() {
	objc.CallMethod[objc.Void](i_, objc.SEL("lockFocus"))
}

// deprecated
func (i_ Image) LockFocusFlipped(flipped bool) {
	objc.CallMethod[objc.Void](i_, objc.SEL("lockFocusFlipped:"), flipped)
}

// deprecated
func (i_ Image) UnlockFocus() {
	objc.CallMethod[objc.Void](i_, objc.SEL("unlockFocus"))
}

// deprecated
func (i_ Image) LockFocusOnRepresentation(imageRepresentation IImageRep) {
	objc.CallMethod[objc.Void](i_, objc.SEL("lockFocusOnRepresentation:"), objc.ExtractPtr(imageRepresentation))
}

// deprecated
func (i_ Image) CompositeToPoint_Operation(point foundation.Point, op CompositingOperation) {
	objc.CallMethod[objc.Void](i_, objc.SEL("compositeToPoint:operation:"), point, op)
}

// deprecated
func (i_ Image) CompositeToPoint_FromRect_Operation(point foundation.Point, rect foundation.Rect, op CompositingOperation) {
	objc.CallMethod[objc.Void](i_, objc.SEL("compositeToPoint:fromRect:operation:"), point, rect, op)
}

// deprecated
func (i_ Image) CompositeToPoint_FromRect_Operation_Fraction(point foundation.Point, rect foundation.Rect, op CompositingOperation, delta float64) {
	objc.CallMethod[objc.Void](i_, objc.SEL("compositeToPoint:fromRect:operation:fraction:"), point, rect, op, delta)
}

// deprecated
func (i_ Image) CompositeToPoint_Operation_Fraction(point foundation.Point, op CompositingOperation, delta float64) {
	objc.CallMethod[objc.Void](i_, objc.SEL("compositeToPoint:operation:fraction:"), point, op, delta)
}

// deprecated
func (i_ Image) DissolveToPoint_Fraction(point foundation.Point, fraction float64) {
	objc.CallMethod[objc.Void](i_, objc.SEL("dissolveToPoint:fraction:"), point, fraction)
}

// deprecated
func (i_ Image) DissolveToPoint_FromRect_Fraction(point foundation.Point, rect foundation.Rect, fraction float64) {
	objc.CallMethod[objc.Void](i_, objc.SEL("dissolveToPoint:fromRect:fraction:"), point, rect, fraction)
}

// deprecated
func (i_ Image) SetScalesWhenResized(flag bool) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setScalesWhenResized:"), flag)
}

// deprecated
func (i_ Image) ScalesWhenResized() bool {
	rv := objc.CallMethod[bool](i_, objc.SEL("scalesWhenResized"))
	return rv
}

// deprecated
func (i_ Image) SetDataRetained(flag bool) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setDataRetained:"), flag)
}

// deprecated
func (i_ Image) IsDataRetained() bool {
	rv := objc.CallMethod[bool](i_, objc.SEL("isDataRetained"))
	return rv
}

// deprecated
func (i_ Image) SetCachedSeparately(flag bool) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setCachedSeparately:"), flag)
}

// deprecated
func (i_ Image) IsCachedSeparately() bool {
	rv := objc.CallMethod[bool](i_, objc.SEL("isCachedSeparately"))
	return rv
}

// deprecated
func (i_ Image) SetCacheDepthMatchesImageDepth(flag bool) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setCacheDepthMatchesImageDepth:"), flag)
}

// deprecated
func (i_ Image) CacheDepthMatchesImageDepth() bool {
	rv := objc.CallMethod[bool](i_, objc.SEL("cacheDepthMatchesImageDepth"))
	return rv
}

// deprecated
func (i_ Image) SetFlipped(flag bool) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setFlipped:"), flag)
}

// deprecated
func (i_ Image) IsFlipped() bool {
	rv := objc.CallMethod[bool](i_, objc.SEL("isFlipped"))
	return rv
}

func (i_ Image) SymbolConfiguration() ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](i_, objc.SEL("symbolConfiguration"))
	return rv
}

// weak property
func (i_ Image) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](i_, objc.SEL("delegate"))
	return rv
}

// weak property
func (i_ Image) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}

func (i_ Image) Size() foundation.Size {
	rv := objc.CallMethod[foundation.Size](i_, objc.SEL("size"))
	return rv
}

func (i_ Image) SetSize(value foundation.Size) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setSize:"), value)
}

func (i_ Image) IsTemplate() bool {
	rv := objc.CallMethod[bool](i_, objc.SEL("isTemplate"))
	return rv
}

func (i_ Image) SetTemplate(value bool) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setTemplate:"), value)
}

func (ic _ImageClass) ImageTypes() []string {
	rv := objc.CallMethod[[]string](ic, objc.SEL("imageTypes"))
	return rv
}

func (ic _ImageClass) ImageUnfilteredTypes() []string {
	rv := objc.CallMethod[[]string](ic, objc.SEL("imageUnfilteredTypes"))
	return rv
}

func (i_ Image) Representations() []ImageRep {
	rv := objc.CallMethod[[]ImageRep](i_, objc.SEL("representations"))
	return rv
}

func (i_ Image) PrefersColorMatch() bool {
	rv := objc.CallMethod[bool](i_, objc.SEL("prefersColorMatch"))
	return rv
}

func (i_ Image) SetPrefersColorMatch(value bool) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setPrefersColorMatch:"), value)
}

func (i_ Image) UsesEPSOnResolutionMismatch() bool {
	rv := objc.CallMethod[bool](i_, objc.SEL("usesEPSOnResolutionMismatch"))
	return rv
}

func (i_ Image) SetUsesEPSOnResolutionMismatch(value bool) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setUsesEPSOnResolutionMismatch:"), value)
}

func (i_ Image) MatchesOnMultipleResolution() bool {
	rv := objc.CallMethod[bool](i_, objc.SEL("matchesOnMultipleResolution"))
	return rv
}

func (i_ Image) SetMatchesOnMultipleResolution(value bool) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setMatchesOnMultipleResolution:"), value)
}

func (i_ Image) IsValid() bool {
	rv := objc.CallMethod[bool](i_, objc.SEL("isValid"))
	return rv
}

func (i_ Image) BackgroundColor() Color {
	rv := objc.CallMethod[Color](i_, objc.SEL("backgroundColor"))
	return rv
}

func (i_ Image) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (i_ Image) CapInsets() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](i_, objc.SEL("capInsets"))
	return rv
}

func (i_ Image) SetCapInsets(value foundation.EdgeInsets) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setCapInsets:"), value)
}

func (i_ Image) ResizingMode() ImageResizingMode {
	rv := objc.CallMethod[ImageResizingMode](i_, objc.SEL("resizingMode"))
	return rv
}

func (i_ Image) SetResizingMode(value ImageResizingMode) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setResizingMode:"), value)
}

func (i_ Image) AlignmentRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](i_, objc.SEL("alignmentRect"))
	return rv
}

func (i_ Image) SetAlignmentRect(value foundation.Rect) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setAlignmentRect:"), value)
}

func (i_ Image) CacheMode() ImageCacheMode {
	rv := objc.CallMethod[ImageCacheMode](i_, objc.SEL("cacheMode"))
	return rv
}

func (i_ Image) SetCacheMode(value ImageCacheMode) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setCacheMode:"), value)
}

func (i_ Image) TIFFRepresentation() []byte {
	rv := objc.CallMethod[[]byte](i_, objc.SEL("TIFFRepresentation"))
	return rv
}

func (i_ Image) AccessibilityDescription() string {
	rv := objc.CallMethod[string](i_, objc.SEL("accessibilityDescription"))
	return rv
}

func (i_ Image) SetAccessibilityDescription(value string) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setAccessibilityDescription:"), value)
}

func (i_ Image) MatchesOnlyOnBestFittingAxis() bool {
	rv := objc.CallMethod[bool](i_, objc.SEL("matchesOnlyOnBestFittingAxis"))
	return rv
}

func (i_ Image) SetMatchesOnlyOnBestFittingAxis(value bool) {
	objc.CallMethod[objc.Void](i_, objc.SEL("setMatchesOnlyOnBestFittingAxis:"), value)
}
