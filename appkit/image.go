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
	SetName(_string ImageName) bool
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

func (i_ Image) Init() Image {
	rv := ffi.CallMethod[Image](i_, "init")
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

func (ic _ImageClass) ImageNamed(name ImageName) Image {
	rv := ffi.CallMethod[Image](ic, "imageNamed:", name)
	return rv
}

func (i_ Image) SetName(_string ImageName) bool {
	rv := ffi.CallMethod[bool](i_, "setName:", _string)
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
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(i_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](i_, "setDelegate:", po)
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

var ImageSymbolConfigurationClass = _ImageSymbolConfigurationClass{objc.GetClass("NSImageSymbolConfiguration")}

type _ImageSymbolConfigurationClass struct {
	objc.Class
}

type IImageSymbolConfiguration interface {
	objc.IObject
}

type ImageSymbolConfiguration struct {
	objc.Object
}

func MakeImageSymbolConfiguration(ptr unsafe.Pointer) ImageSymbolConfiguration {
	return ImageSymbolConfiguration{
		Object: objc.MakeObject(ptr),
	}
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithPointSize_Weight(pointSize float64, weight FontWeight) ImageSymbolConfiguration {
	rv := ffi.CallMethod[ImageSymbolConfiguration](ic, "configurationWithPointSize:weight:", pointSize, weight)
	return rv
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithPointSize_Weight_Scale(pointSize float64, weight FontWeight, scale ImageSymbolScale) ImageSymbolConfiguration {
	rv := ffi.CallMethod[ImageSymbolConfiguration](ic, "configurationWithPointSize:weight:scale:", pointSize, weight, scale)
	return rv
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithTextStyle(style FontTextStyle) ImageSymbolConfiguration {
	rv := ffi.CallMethod[ImageSymbolConfiguration](ic, "configurationWithTextStyle:", style)
	return rv
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithTextStyle_Scale(style FontTextStyle, scale ImageSymbolScale) ImageSymbolConfiguration {
	rv := ffi.CallMethod[ImageSymbolConfiguration](ic, "configurationWithTextStyle:scale:", style, scale)
	return rv
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithScale(scale ImageSymbolScale) ImageSymbolConfiguration {
	rv := ffi.CallMethod[ImageSymbolConfiguration](ic, "configurationWithScale:", scale)
	return rv
}

func (i_ ImageSymbolConfiguration) ConfigurationByApplyingConfiguration(configuration IImageSymbolConfiguration) ImageSymbolConfiguration {
	rv := ffi.CallMethod[ImageSymbolConfiguration](i_, "configurationByApplyingConfiguration:", configuration)
	return rv
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithPaletteColors(paletteColors []IColor) ImageSymbolConfiguration {
	rv := ffi.CallMethod[ImageSymbolConfiguration](ic, "configurationWithPaletteColors:", paletteColors)
	return rv
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithHierarchicalColor(hierarchicalColor IColor) ImageSymbolConfiguration {
	rv := ffi.CallMethod[ImageSymbolConfiguration](ic, "configurationWithHierarchicalColor:", hierarchicalColor)
	return rv
}

func (ic _ImageSymbolConfigurationClass) ConfigurationPreferringMulticolor() ImageSymbolConfiguration {
	rv := ffi.CallMethod[ImageSymbolConfiguration](ic, "configurationPreferringMulticolor")
	return rv
}

func (ic _ImageSymbolConfigurationClass) ConfigurationPreferringHierarchical() ImageSymbolConfiguration {
	rv := ffi.CallMethod[ImageSymbolConfiguration](ic, "configurationPreferringHierarchical")
	return rv
}

func (ic _ImageSymbolConfigurationClass) ConfigurationPreferringMonochrome() ImageSymbolConfiguration {
	rv := ffi.CallMethod[ImageSymbolConfiguration](ic, "configurationPreferringMonochrome")
	return rv
}

func (ic _ImageSymbolConfigurationClass) Alloc() ImageSymbolConfiguration {
	rv := ffi.CallMethod[ImageSymbolConfiguration](ic, "alloc")
	return rv
}

func (i_ ImageSymbolConfiguration) Init() ImageSymbolConfiguration {
	rv := ffi.CallMethod[ImageSymbolConfiguration](i_, "init")
	return rv
}

func (ic _ImageSymbolConfigurationClass) New() ImageSymbolConfiguration {
	rv := ffi.CallMethod[ImageSymbolConfiguration](ic, "new")
	rv.Autorelease()
	return rv
}

func NewImageSymbolConfiguration() ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.New()
}

type ImageDelegate interface {
	ImplementsImageDidNotDraw_InRect() bool
	// optional
	ImageDidNotDraw_InRect(sender Image, rect foundation.Rect) IImage
	ImplementsImage_DidLoadPartOfRepresentation_WithValidRows() bool
	// optional
	Image_DidLoadPartOfRepresentation_WithValidRows(image Image, rep ImageRep, rows int)
	ImplementsImage_DidLoadRepresentation_WithStatus() bool
	// optional
	Image_DidLoadRepresentation_WithStatus(image Image, rep ImageRep, status ImageLoadStatus)
	ImplementsImage_DidLoadRepresentationHeader() bool
	// optional
	Image_DidLoadRepresentationHeader(image Image, rep ImageRep)
	ImplementsImage_WillLoadRepresentation() bool
	// optional
	Image_WillLoadRepresentation(image Image, rep ImageRep)
}

type ImageDelegateImpl struct {
	_ImageDidNotDraw_InRect                          func(sender Image, rect foundation.Rect) IImage
	_Image_DidLoadPartOfRepresentation_WithValidRows func(image Image, rep ImageRep, rows int)
	_Image_DidLoadRepresentation_WithStatus          func(image Image, rep ImageRep, status ImageLoadStatus)
	_Image_DidLoadRepresentationHeader               func(image Image, rep ImageRep)
	_Image_WillLoadRepresentation                    func(image Image, rep ImageRep)
}

func (di *ImageDelegateImpl) ImplementsImageDidNotDraw_InRect() bool {
	return di._ImageDidNotDraw_InRect != nil
}

func (di *ImageDelegateImpl) SetImageDidNotDraw_InRect(f func(sender Image, rect foundation.Rect) IImage) {
	di._ImageDidNotDraw_InRect = f
}

func (di *ImageDelegateImpl) ImageDidNotDraw_InRect(sender Image, rect foundation.Rect) IImage {
	return di._ImageDidNotDraw_InRect(sender, rect)
}
func (di *ImageDelegateImpl) ImplementsImage_DidLoadPartOfRepresentation_WithValidRows() bool {
	return di._Image_DidLoadPartOfRepresentation_WithValidRows != nil
}

func (di *ImageDelegateImpl) SetImage_DidLoadPartOfRepresentation_WithValidRows(f func(image Image, rep ImageRep, rows int)) {
	di._Image_DidLoadPartOfRepresentation_WithValidRows = f
}

func (di *ImageDelegateImpl) Image_DidLoadPartOfRepresentation_WithValidRows(image Image, rep ImageRep, rows int) {
	di._Image_DidLoadPartOfRepresentation_WithValidRows(image, rep, rows)
}
func (di *ImageDelegateImpl) ImplementsImage_DidLoadRepresentation_WithStatus() bool {
	return di._Image_DidLoadRepresentation_WithStatus != nil
}

func (di *ImageDelegateImpl) SetImage_DidLoadRepresentation_WithStatus(f func(image Image, rep ImageRep, status ImageLoadStatus)) {
	di._Image_DidLoadRepresentation_WithStatus = f
}

func (di *ImageDelegateImpl) Image_DidLoadRepresentation_WithStatus(image Image, rep ImageRep, status ImageLoadStatus) {
	di._Image_DidLoadRepresentation_WithStatus(image, rep, status)
}
func (di *ImageDelegateImpl) ImplementsImage_DidLoadRepresentationHeader() bool {
	return di._Image_DidLoadRepresentationHeader != nil
}

func (di *ImageDelegateImpl) SetImage_DidLoadRepresentationHeader(f func(image Image, rep ImageRep)) {
	di._Image_DidLoadRepresentationHeader = f
}

func (di *ImageDelegateImpl) Image_DidLoadRepresentationHeader(image Image, rep ImageRep) {
	di._Image_DidLoadRepresentationHeader(image, rep)
}
func (di *ImageDelegateImpl) ImplementsImage_WillLoadRepresentation() bool {
	return di._Image_WillLoadRepresentation != nil
}

func (di *ImageDelegateImpl) SetImage_WillLoadRepresentation(f func(image Image, rep ImageRep)) {
	di._Image_WillLoadRepresentation = f
}

func (di *ImageDelegateImpl) Image_WillLoadRepresentation(image Image, rep ImageRep) {
	di._Image_WillLoadRepresentation(image, rep)
}

type ImageDelegateWrapper struct {
	objc.Object
}

func (i_ *ImageDelegateWrapper) ImplementsImageDidNotDraw_InRect() bool {
	return i_.RespondsToSelector(objc.GetSelector("imageDidNotDraw:inRect:"))
}

func (i_ ImageDelegateWrapper) ImageDidNotDraw_InRect(sender IImage, rect foundation.Rect) Image {
	rv := ffi.CallMethod[Image](i_, "imageDidNotDraw:inRect:", sender, rect)
	return rv
}

func (i_ *ImageDelegateWrapper) ImplementsImage_DidLoadPartOfRepresentation_WithValidRows() bool {
	return i_.RespondsToSelector(objc.GetSelector("image:didLoadPartOfRepresentation:withValidRows:"))
}

func (i_ ImageDelegateWrapper) Image_DidLoadPartOfRepresentation_WithValidRows(image IImage, rep IImageRep, rows int) {
	ffi.CallMethod[ffi.Void](i_, "image:didLoadPartOfRepresentation:withValidRows:", image, rep, rows)
}

func (i_ *ImageDelegateWrapper) ImplementsImage_DidLoadRepresentation_WithStatus() bool {
	return i_.RespondsToSelector(objc.GetSelector("image:didLoadRepresentation:withStatus:"))
}

func (i_ ImageDelegateWrapper) Image_DidLoadRepresentation_WithStatus(image IImage, rep IImageRep, status ImageLoadStatus) {
	ffi.CallMethod[ffi.Void](i_, "image:didLoadRepresentation:withStatus:", image, rep, status)
}

func (i_ *ImageDelegateWrapper) ImplementsImage_DidLoadRepresentationHeader() bool {
	return i_.RespondsToSelector(objc.GetSelector("image:didLoadRepresentationHeader:"))
}

func (i_ ImageDelegateWrapper) Image_DidLoadRepresentationHeader(image IImage, rep IImageRep) {
	ffi.CallMethod[ffi.Void](i_, "image:didLoadRepresentationHeader:", image, rep)
}

func (i_ *ImageDelegateWrapper) ImplementsImage_WillLoadRepresentation() bool {
	return i_.RespondsToSelector(objc.GetSelector("image:willLoadRepresentation:"))
}

func (i_ ImageDelegateWrapper) Image_WillLoadRepresentation(image IImage, rep IImageRep) {
	ffi.CallMethod[ffi.Void](i_, "image:willLoadRepresentation:", image, rep)
}

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

var BitmapImageRepClass = _BitmapImageRepClass{objc.GetClass("NSBitmapImageRep")}

type _BitmapImageRepClass struct {
	objc.Class
}

type IBitmapImageRep interface {
	IImageRep
	ColorizeByMappingGray_ToColor_BlackMapping_WhiteMapping(midPoint float64, midPointColor IColor, shadowColor IColor, lightColor IColor)
	TIFFRepresentationUsingCompression_Factor(comp TIFFCompression, factor float32) []byte
	RepresentationUsingType_Properties(storageType BitmapImageFileType, properties map[BitmapImageRepPropertyKey]objc.IObject) []byte
	CanBeCompressedUsing(compression TIFFCompression) bool
	SetCompression_Factor(compression TIFFCompression, factor float32)
	GetCompression_Factor(compression *TIFFCompression, factor *float32)
	SetProperty_WithValue(property BitmapImageRepPropertyKey, value objc.IObject)
	ValueForProperty(property BitmapImageRepPropertyKey) objc.Object
	IncrementalLoadFromData_Complete(data []byte, complete bool) int
	SetColor_AtX_Y(color IColor, x int, y int)
	ColorAtX_Y(x int, y int) Color
	SetPixel_AtX_Y(p *uint, x int, y int)
	GetPixel_AtX_Y(p *uint, x int, y int)
	BitmapImageRepByConvertingToColorSpace_RenderingIntent(targetSpace IColorSpace, renderingIntent ColorRenderingIntent) BitmapImageRep
	BitmapImageRepByRetaggingWithColorSpace(newSpace IColorSpace) BitmapImageRep
	BitmapFormat() BitmapFormat
	BitsPerPixel() int
	BytesPerPlane() int
	BytesPerRow() int
	IsPlanar() bool
	NumberOfPlanes() int
	SamplesPerPixel() int
	BitmapData() *byte
	TIFFRepresentation() []byte
	CGImage() coregraphics.ImageRef
	ColorSpace() ColorSpace
}

type BitmapImageRep struct {
	ImageRep
}

func MakeBitmapImageRep(ptr unsafe.Pointer) BitmapImageRep {
	return BitmapImageRep{
		ImageRep: MakeImageRep(ptr),
	}
}

func (bc _BitmapImageRepClass) ImageRepWithData(data []byte) BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](bc, "imageRepWithData:", data)
	return rv
}

func (b_ BitmapImageRep) InitWithCGImage(cgImage coregraphics.ImageRef) BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](b_, "initWithCGImage:", cgImage)
	return rv
}

func (b_ BitmapImageRep) InitWithData(data []byte) BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](b_, "initWithData:", data)
	return rv
}

func (b_ BitmapImageRep) InitForIncrementalLoad() BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](b_, "initForIncrementalLoad")
	return rv
}

func (b_ BitmapImageRep) InitWithFocusedViewRect(rect foundation.Rect) BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](b_, "initWithFocusedViewRect:", rect)
	return rv
}

func (b_ BitmapImageRep) Init() BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](b_, "init")
	return rv
}

func (bc _BitmapImageRepClass) Alloc() BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](bc, "alloc")
	return rv
}

func (bc _BitmapImageRepClass) New() BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](bc, "new")
	rv.Autorelease()
	return rv
}

func NewBitmapImageRep() BitmapImageRep {
	return BitmapImageRepClass.New()
}

func (bc _BitmapImageRepClass) ImageRepsWithData(data []byte) []ImageRep {
	rv := ffi.CallMethod[[]ImageRep](bc, "imageRepsWithData:", data)
	return rv
}

func (b_ BitmapImageRep) ColorizeByMappingGray_ToColor_BlackMapping_WhiteMapping(midPoint float64, midPointColor IColor, shadowColor IColor, lightColor IColor) {
	ffi.CallMethod[ffi.Void](b_, "colorizeByMappingGray:toColor:blackMapping:whiteMapping:", midPoint, midPointColor, shadowColor, lightColor)
}

func (bc _BitmapImageRepClass) TIFFRepresentationOfImageRepsInArray(array []IImageRep) []byte {
	rv := ffi.CallMethod[[]byte](bc, "TIFFRepresentationOfImageRepsInArray:", array)
	return rv
}

func (bc _BitmapImageRepClass) TIFFRepresentationOfImageRepsInArray_UsingCompression_Factor(array []IImageRep, comp TIFFCompression, factor float32) []byte {
	rv := ffi.CallMethod[[]byte](bc, "TIFFRepresentationOfImageRepsInArray:usingCompression:factor:", array, comp, factor)
	return rv
}

func (b_ BitmapImageRep) TIFFRepresentationUsingCompression_Factor(comp TIFFCompression, factor float32) []byte {
	rv := ffi.CallMethod[[]byte](b_, "TIFFRepresentationUsingCompression:factor:", comp, factor)
	return rv
}

func (bc _BitmapImageRepClass) RepresentationOfImageRepsInArray_UsingType_Properties(imageReps []IImageRep, storageType BitmapImageFileType, properties map[BitmapImageRepPropertyKey]objc.IObject) []byte {
	rv := ffi.CallMethod[[]byte](bc, "representationOfImageRepsInArray:usingType:properties:", imageReps, storageType, properties)
	return rv
}

func (b_ BitmapImageRep) RepresentationUsingType_Properties(storageType BitmapImageFileType, properties map[BitmapImageRepPropertyKey]objc.IObject) []byte {
	rv := ffi.CallMethod[[]byte](b_, "representationUsingType:properties:", storageType, properties)
	return rv
}

func (bc _BitmapImageRepClass) LocalizedNameForTIFFCompressionType(compression TIFFCompression) string {
	rv := ffi.CallMethod[string](bc, "localizedNameForTIFFCompressionType:", compression)
	return rv
}

func (b_ BitmapImageRep) CanBeCompressedUsing(compression TIFFCompression) bool {
	rv := ffi.CallMethod[bool](b_, "canBeCompressedUsing:", compression)
	return rv
}

func (b_ BitmapImageRep) SetCompression_Factor(compression TIFFCompression, factor float32) {
	ffi.CallMethod[ffi.Void](b_, "setCompression:factor:", compression, factor)
}

func (b_ BitmapImageRep) GetCompression_Factor(compression *TIFFCompression, factor *float32) {
	ffi.CallMethod[ffi.Void](b_, "getCompression:factor:", compression, factor)
}

func (b_ BitmapImageRep) SetProperty_WithValue(property BitmapImageRepPropertyKey, value objc.IObject) {
	ffi.CallMethod[ffi.Void](b_, "setProperty:withValue:", property, value)
}

func (b_ BitmapImageRep) ValueForProperty(property BitmapImageRepPropertyKey) objc.Object {
	rv := ffi.CallMethod[objc.Object](b_, "valueForProperty:", property)
	return rv
}

func (b_ BitmapImageRep) IncrementalLoadFromData_Complete(data []byte, complete bool) int {
	rv := ffi.CallMethod[int](b_, "incrementalLoadFromData:complete:", data, complete)
	return rv
}

func (b_ BitmapImageRep) SetColor_AtX_Y(color IColor, x int, y int) {
	ffi.CallMethod[ffi.Void](b_, "setColor:atX:y:", color, x, y)
}

func (b_ BitmapImageRep) ColorAtX_Y(x int, y int) Color {
	rv := ffi.CallMethod[Color](b_, "colorAtX:y:", x, y)
	return rv
}

func (b_ BitmapImageRep) SetPixel_AtX_Y(p *uint, x int, y int) {
	ffi.CallMethod[ffi.Void](b_, "setPixel:atX:y:", p, x, y)
}

func (b_ BitmapImageRep) GetPixel_AtX_Y(p *uint, x int, y int) {
	ffi.CallMethod[ffi.Void](b_, "getPixel:atX:y:", p, x, y)
}

func (b_ BitmapImageRep) BitmapImageRepByConvertingToColorSpace_RenderingIntent(targetSpace IColorSpace, renderingIntent ColorRenderingIntent) BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](b_, "bitmapImageRepByConvertingToColorSpace:renderingIntent:", targetSpace, renderingIntent)
	return rv
}

func (b_ BitmapImageRep) BitmapImageRepByRetaggingWithColorSpace(newSpace IColorSpace) BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](b_, "bitmapImageRepByRetaggingWithColorSpace:", newSpace)
	return rv
}

func (b_ BitmapImageRep) BitmapFormat() BitmapFormat {
	rv := ffi.CallMethod[BitmapFormat](b_, "bitmapFormat")
	return rv
}

func (b_ BitmapImageRep) BitsPerPixel() int {
	rv := ffi.CallMethod[int](b_, "bitsPerPixel")
	return rv
}

func (b_ BitmapImageRep) BytesPerPlane() int {
	rv := ffi.CallMethod[int](b_, "bytesPerPlane")
	return rv
}

func (b_ BitmapImageRep) BytesPerRow() int {
	rv := ffi.CallMethod[int](b_, "bytesPerRow")
	return rv
}

func (b_ BitmapImageRep) IsPlanar() bool {
	rv := ffi.CallMethod[bool](b_, "isPlanar")
	return rv
}

func (b_ BitmapImageRep) NumberOfPlanes() int {
	rv := ffi.CallMethod[int](b_, "numberOfPlanes")
	return rv
}

func (b_ BitmapImageRep) SamplesPerPixel() int {
	rv := ffi.CallMethod[int](b_, "samplesPerPixel")
	return rv
}

func (b_ BitmapImageRep) BitmapData() *byte {
	rv := ffi.CallMethod[*byte](b_, "bitmapData")
	return rv
}

func (b_ BitmapImageRep) TIFFRepresentation() []byte {
	rv := ffi.CallMethod[[]byte](b_, "TIFFRepresentation")
	return rv
}

func (b_ BitmapImageRep) CGImage() coregraphics.ImageRef {
	rv := ffi.CallMethod[coregraphics.ImageRef](b_, "CGImage")
	return rv
}

func (b_ BitmapImageRep) ColorSpace() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](b_, "colorSpace")
	return rv
}
