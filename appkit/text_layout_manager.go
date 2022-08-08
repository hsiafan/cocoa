// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TextLayoutManagerClass = _TextLayoutManagerClass{objc.GetClass("NSTextLayoutManager")}

type _TextLayoutManagerClass struct {
	objc.Class
}

type ITextLayoutManager interface {
	objc.IObject
	EnumerateTextSegmentsInRange_Type_Options_UsingBlock(textRange ITextRange, _type TextLayoutManagerSegmentType, options TextLayoutManagerSegmentOptions, block func(textSegmentRange TextRange, textSegmentFrame coregraphics.Rect, baselinePosition float64, textContainer TextContainer) bool)
	ReplaceTextContentManager(textContentManager ITextContentManager)
	ReplaceContentsInRange_WithAttributedString(_range ITextRange, attributedString foundation.IAttributedString)
	ReplaceContentsInRange_WithTextElements(_range ITextRange, textElements []ITextElement)
	AddRenderingAttribute_Value_ForTextRange(renderingAttribute foundation.AttributedStringKey, value objc.IObject, textRange ITextRange)
	InvalidateRenderingAttributesForTextRange(textRange ITextRange)
	RemoveRenderingAttribute_ForTextRange(renderingAttribute foundation.AttributedStringKey, textRange ITextRange)
	SetRenderingAttributes_ForTextRange(renderingAttributes map[foundation.AttributedStringKey]objc.IObject, textRange ITextRange)
	InvalidateLayoutForRange(_range ITextRange)
	TextLayoutFragmentForPosition(position coregraphics.Point) TextLayoutFragment
	EnsureLayoutForBounds(bounds coregraphics.Rect)
	EnsureLayoutForRange(_range ITextRange)
	LayoutQueue() foundation.OperationQueue
	SetLayoutQueue(value foundation.IOperationQueue)
	RenderingAttributesValidator() func(textLayoutManager ITextLayoutManager, textLayoutFragment ITextLayoutFragment)
	SetRenderingAttributesValidator(value func(textLayoutManager TextLayoutManager, textLayoutFragment TextLayoutFragment))
	UsesFontLeading() bool
	SetUsesFontLeading(value bool)
	UsesHyphenation() bool
	SetUsesHyphenation(value bool)
	LimitsLayoutForSuspiciousContents() bool
	SetLimitsLayoutForSuspiciousContents(value bool)
	TextContentManager() TextContentManager
	TextContainer() TextContainer
	SetTextContainer(value ITextContainer)
	TextSelectionNavigation() TextSelectionNavigation
	SetTextSelectionNavigation(value ITextSelectionNavigation)
	TextSelections() []TextSelection
	SetTextSelections(value []ITextSelection)
	UsageBoundsForTextContainer() coregraphics.Rect
	TextViewportLayoutController() TextViewportLayoutController
}

type TextLayoutManager struct {
	objc.Object
}

func MakeTextLayoutManager(ptr unsafe.Pointer) TextLayoutManager {
	return TextLayoutManager{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextLayoutManager) Init() TextLayoutManager {
	rv := ffi.CallMethod[TextLayoutManager](t_, "init")
	return rv
}

func (tc _TextLayoutManagerClass) Alloc() TextLayoutManager {
	rv := ffi.CallMethod[TextLayoutManager](tc, "alloc")
	return rv
}

func (tc _TextLayoutManagerClass) New() TextLayoutManager {
	rv := ffi.CallMethod[TextLayoutManager](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextLayoutManager() TextLayoutManager {
	return TextLayoutManagerClass.New()
}

func (t_ TextLayoutManager) EnumerateTextSegmentsInRange_Type_Options_UsingBlock(textRange ITextRange, _type TextLayoutManagerSegmentType, options TextLayoutManagerSegmentOptions, block func(textSegmentRange TextRange, textSegmentFrame coregraphics.Rect, baselinePosition float64, textContainer TextContainer) bool) {
	ffi.CallMethod[ffi.Void](t_, "enumerateTextSegmentsInRange:type:options:usingBlock:", textRange, _type, options, block)
}

func (t_ TextLayoutManager) ReplaceTextContentManager(textContentManager ITextContentManager) {
	ffi.CallMethod[ffi.Void](t_, "replaceTextContentManager:", textContentManager)
}

func (t_ TextLayoutManager) ReplaceContentsInRange_WithAttributedString(_range ITextRange, attributedString foundation.IAttributedString) {
	ffi.CallMethod[ffi.Void](t_, "replaceContentsInRange:withAttributedString:", _range, attributedString)
}

func (t_ TextLayoutManager) ReplaceContentsInRange_WithTextElements(_range ITextRange, textElements []ITextElement) {
	ffi.CallMethod[ffi.Void](t_, "replaceContentsInRange:withTextElements:", _range, textElements)
}

func (t_ TextLayoutManager) AddRenderingAttribute_Value_ForTextRange(renderingAttribute foundation.AttributedStringKey, value objc.IObject, textRange ITextRange) {
	ffi.CallMethod[ffi.Void](t_, "addRenderingAttribute:value:forTextRange:", renderingAttribute, value, textRange)
}

func (t_ TextLayoutManager) InvalidateRenderingAttributesForTextRange(textRange ITextRange) {
	ffi.CallMethod[ffi.Void](t_, "invalidateRenderingAttributesForTextRange:", textRange)
}

func (t_ TextLayoutManager) RemoveRenderingAttribute_ForTextRange(renderingAttribute foundation.AttributedStringKey, textRange ITextRange) {
	ffi.CallMethod[ffi.Void](t_, "removeRenderingAttribute:forTextRange:", renderingAttribute, textRange)
}

func (t_ TextLayoutManager) SetRenderingAttributes_ForTextRange(renderingAttributes map[foundation.AttributedStringKey]objc.IObject, textRange ITextRange) {
	ffi.CallMethod[ffi.Void](t_, "setRenderingAttributes:forTextRange:", renderingAttributes, textRange)
}

func (t_ TextLayoutManager) InvalidateLayoutForRange(_range ITextRange) {
	ffi.CallMethod[ffi.Void](t_, "invalidateLayoutForRange:", _range)
}

func (t_ TextLayoutManager) TextLayoutFragmentForPosition(position coregraphics.Point) TextLayoutFragment {
	rv := ffi.CallMethod[TextLayoutFragment](t_, "textLayoutFragmentForPosition:", position)
	return rv
}

func (t_ TextLayoutManager) EnsureLayoutForBounds(bounds coregraphics.Rect) {
	ffi.CallMethod[ffi.Void](t_, "ensureLayoutForBounds:", bounds)
}

func (t_ TextLayoutManager) EnsureLayoutForRange(_range ITextRange) {
	ffi.CallMethod[ffi.Void](t_, "ensureLayoutForRange:", _range)
}

func (t_ TextLayoutManager) LayoutQueue() foundation.OperationQueue {
	rv := ffi.CallMethod[foundation.OperationQueue](t_, "layoutQueue")
	return rv
}

func (t_ TextLayoutManager) SetLayoutQueue(value foundation.IOperationQueue) {
	ffi.CallMethod[ffi.Void](t_, "setLayoutQueue:", value)
}

func (t_ TextLayoutManager) RenderingAttributesValidator() func(textLayoutManager ITextLayoutManager, textLayoutFragment ITextLayoutFragment) {
	rv := ffi.CallMethod[func(textLayoutManager ITextLayoutManager, textLayoutFragment ITextLayoutFragment)](t_, "renderingAttributesValidator")
	return rv
}

func (t_ TextLayoutManager) SetRenderingAttributesValidator(value func(textLayoutManager TextLayoutManager, textLayoutFragment TextLayoutFragment)) {
	ffi.CallMethod[ffi.Void](t_, "setRenderingAttributesValidator:", value)
}

func (t_ TextLayoutManager) UsesFontLeading() bool {
	rv := ffi.CallMethod[bool](t_, "usesFontLeading")
	return rv
}

func (t_ TextLayoutManager) SetUsesFontLeading(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setUsesFontLeading:", value)
}

func (t_ TextLayoutManager) UsesHyphenation() bool {
	rv := ffi.CallMethod[bool](t_, "usesHyphenation")
	return rv
}

func (t_ TextLayoutManager) SetUsesHyphenation(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setUsesHyphenation:", value)
}

func (t_ TextLayoutManager) LimitsLayoutForSuspiciousContents() bool {
	rv := ffi.CallMethod[bool](t_, "limitsLayoutForSuspiciousContents")
	return rv
}

func (t_ TextLayoutManager) SetLimitsLayoutForSuspiciousContents(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setLimitsLayoutForSuspiciousContents:", value)
}

func (t_ TextLayoutManager) TextContentManager() TextContentManager {
	rv := ffi.CallMethod[TextContentManager](t_, "textContentManager")
	return rv
}

func (t_ TextLayoutManager) TextContainer() TextContainer {
	rv := ffi.CallMethod[TextContainer](t_, "textContainer")
	return rv
}

func (t_ TextLayoutManager) SetTextContainer(value ITextContainer) {
	ffi.CallMethod[ffi.Void](t_, "setTextContainer:", value)
}

func (t_ TextLayoutManager) TextSelectionNavigation() TextSelectionNavigation {
	rv := ffi.CallMethod[TextSelectionNavigation](t_, "textSelectionNavigation")
	return rv
}

func (t_ TextLayoutManager) SetTextSelectionNavigation(value ITextSelectionNavigation) {
	ffi.CallMethod[ffi.Void](t_, "setTextSelectionNavigation:", value)
}

func (t_ TextLayoutManager) TextSelections() []TextSelection {
	rv := ffi.CallMethod[[]TextSelection](t_, "textSelections")
	return rv
}

func (t_ TextLayoutManager) SetTextSelections(value []ITextSelection) {
	ffi.CallMethod[ffi.Void](t_, "setTextSelections:", value)
}

func (t_ TextLayoutManager) UsageBoundsForTextContainer() coregraphics.Rect {
	rv := ffi.CallMethod[coregraphics.Rect](t_, "usageBoundsForTextContainer")
	return rv
}

func (tc _TextLayoutManagerClass) LinkRenderingAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := ffi.CallMethod[map[foundation.AttributedStringKey]objc.Object](tc, "linkRenderingAttributes")
	return rv
}

func (t_ TextLayoutManager) TextViewportLayoutController() TextViewportLayoutController {
	rv := ffi.CallMethod[TextViewportLayoutController](t_, "textViewportLayoutController")
	return rv
}

var TextLayoutFragmentClass = _TextLayoutFragmentClass{objc.GetClass("NSTextLayoutFragment")}

type _TextLayoutFragmentClass struct {
	objc.Class
}

type ITextLayoutFragment interface {
	objc.IObject
	DrawAtPoint_InContext(point coregraphics.Point, context coregraphics.ContextRef)
	InvalidateLayout()
	BottomMargin() float64
	LeadingPadding() float64
	TopMargin() float64
	TrailingPadding() float64
	TextLayoutManager() TextLayoutManager
	LayoutQueue() foundation.OperationQueue
	SetLayoutQueue(value foundation.IOperationQueue)
	LayoutFragmentFrame() coregraphics.Rect
	RenderingSurfaceBounds() coregraphics.Rect
	TextAttachmentViewProviders() []TextAttachmentViewProvider
	State() TextLayoutFragmentState
	RangeInElement() TextRange
	TextElement() TextElement
	TextLineFragments() []TextLineFragment
}

type TextLayoutFragment struct {
	objc.Object
}

func MakeTextLayoutFragment(ptr unsafe.Pointer) TextLayoutFragment {
	return TextLayoutFragment{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextLayoutFragment) InitWithTextElement_Range(textElement ITextElement, rangeInElement ITextRange) TextLayoutFragment {
	rv := ffi.CallMethod[TextLayoutFragment](t_, "initWithTextElement:range:", textElement, rangeInElement)
	return rv
}

func (tc _TextLayoutFragmentClass) Alloc() TextLayoutFragment {
	rv := ffi.CallMethod[TextLayoutFragment](tc, "alloc")
	return rv
}

func (t_ TextLayoutFragment) Init() TextLayoutFragment {
	rv := ffi.CallMethod[TextLayoutFragment](t_, "init")
	return rv
}

func (tc _TextLayoutFragmentClass) New() TextLayoutFragment {
	rv := ffi.CallMethod[TextLayoutFragment](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextLayoutFragment() TextLayoutFragment {
	return TextLayoutFragmentClass.New()
}

func (t_ TextLayoutFragment) DrawAtPoint_InContext(point coregraphics.Point, context coregraphics.ContextRef) {
	ffi.CallMethod[ffi.Void](t_, "drawAtPoint:inContext:", point, context)
}

func (t_ TextLayoutFragment) InvalidateLayout() {
	ffi.CallMethod[ffi.Void](t_, "invalidateLayout")
}

func (t_ TextLayoutFragment) BottomMargin() float64 {
	rv := ffi.CallMethod[float64](t_, "bottomMargin")
	return rv
}

func (t_ TextLayoutFragment) LeadingPadding() float64 {
	rv := ffi.CallMethod[float64](t_, "leadingPadding")
	return rv
}

func (t_ TextLayoutFragment) TopMargin() float64 {
	rv := ffi.CallMethod[float64](t_, "topMargin")
	return rv
}

func (t_ TextLayoutFragment) TrailingPadding() float64 {
	rv := ffi.CallMethod[float64](t_, "trailingPadding")
	return rv
}

func (t_ TextLayoutFragment) TextLayoutManager() TextLayoutManager {
	rv := ffi.CallMethod[TextLayoutManager](t_, "textLayoutManager")
	return rv
}

func (t_ TextLayoutFragment) LayoutQueue() foundation.OperationQueue {
	rv := ffi.CallMethod[foundation.OperationQueue](t_, "layoutQueue")
	return rv
}

func (t_ TextLayoutFragment) SetLayoutQueue(value foundation.IOperationQueue) {
	ffi.CallMethod[ffi.Void](t_, "setLayoutQueue:", value)
}

func (t_ TextLayoutFragment) LayoutFragmentFrame() coregraphics.Rect {
	rv := ffi.CallMethod[coregraphics.Rect](t_, "layoutFragmentFrame")
	return rv
}

func (t_ TextLayoutFragment) RenderingSurfaceBounds() coregraphics.Rect {
	rv := ffi.CallMethod[coregraphics.Rect](t_, "renderingSurfaceBounds")
	return rv
}

func (t_ TextLayoutFragment) TextAttachmentViewProviders() []TextAttachmentViewProvider {
	rv := ffi.CallMethod[[]TextAttachmentViewProvider](t_, "textAttachmentViewProviders")
	return rv
}

func (t_ TextLayoutFragment) State() TextLayoutFragmentState {
	rv := ffi.CallMethod[TextLayoutFragmentState](t_, "state")
	return rv
}

func (t_ TextLayoutFragment) RangeInElement() TextRange {
	rv := ffi.CallMethod[TextRange](t_, "rangeInElement")
	return rv
}

func (t_ TextLayoutFragment) TextElement() TextElement {
	rv := ffi.CallMethod[TextElement](t_, "textElement")
	return rv
}

func (t_ TextLayoutFragment) TextLineFragments() []TextLineFragment {
	rv := ffi.CallMethod[[]TextLineFragment](t_, "textLineFragments")
	return rv
}

var TextLineFragmentClass = _TextLineFragmentClass{objc.GetClass("NSTextLineFragment")}

type _TextLineFragmentClass struct {
	objc.Class
}

type ITextLineFragment interface {
	objc.IObject
	CharacterIndexForPoint(point coregraphics.Point) int
	FractionOfDistanceThroughGlyphForPoint(point coregraphics.Point) float64
	LocationForCharacterAtIndex(index int) coregraphics.Point
	DrawAtPoint_InContext(point coregraphics.Point, context coregraphics.ContextRef)
	AttributedString() foundation.AttributedString
	CharacterRange() foundation.Range
	GlyphOrigin() coregraphics.Point
	TypographicBounds() coregraphics.Rect
}

type TextLineFragment struct {
	objc.Object
}

func MakeTextLineFragment(ptr unsafe.Pointer) TextLineFragment {
	return TextLineFragment{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextLineFragment) InitWithAttributedString_Range(attributedString foundation.IAttributedString, _range foundation.Range) TextLineFragment {
	rv := ffi.CallMethod[TextLineFragment](t_, "initWithAttributedString:range:", attributedString, _range)
	return rv
}

func (t_ TextLineFragment) InitWithString_Attributes_Range(_string string, attributes map[foundation.AttributedStringKey]objc.IObject, _range foundation.Range) TextLineFragment {
	rv := ffi.CallMethod[TextLineFragment](t_, "initWithString:attributes:range:", _string, attributes, _range)
	return rv
}

func (tc _TextLineFragmentClass) Alloc() TextLineFragment {
	rv := ffi.CallMethod[TextLineFragment](tc, "alloc")
	return rv
}

func (t_ TextLineFragment) Init() TextLineFragment {
	rv := ffi.CallMethod[TextLineFragment](t_, "init")
	return rv
}

func (tc _TextLineFragmentClass) New() TextLineFragment {
	rv := ffi.CallMethod[TextLineFragment](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextLineFragment() TextLineFragment {
	return TextLineFragmentClass.New()
}

func (t_ TextLineFragment) CharacterIndexForPoint(point coregraphics.Point) int {
	rv := ffi.CallMethod[int](t_, "characterIndexForPoint:", point)
	return rv
}

func (t_ TextLineFragment) FractionOfDistanceThroughGlyphForPoint(point coregraphics.Point) float64 {
	rv := ffi.CallMethod[float64](t_, "fractionOfDistanceThroughGlyphForPoint:", point)
	return rv
}

func (t_ TextLineFragment) LocationForCharacterAtIndex(index int) coregraphics.Point {
	rv := ffi.CallMethod[coregraphics.Point](t_, "locationForCharacterAtIndex:", index)
	return rv
}

func (t_ TextLineFragment) DrawAtPoint_InContext(point coregraphics.Point, context coregraphics.ContextRef) {
	ffi.CallMethod[ffi.Void](t_, "drawAtPoint:inContext:", point, context)
}

func (t_ TextLineFragment) AttributedString() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](t_, "attributedString")
	return rv
}

func (t_ TextLineFragment) CharacterRange() foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "characterRange")
	return rv
}

func (t_ TextLineFragment) GlyphOrigin() coregraphics.Point {
	rv := ffi.CallMethod[coregraphics.Point](t_, "glyphOrigin")
	return rv
}

func (t_ TextLineFragment) TypographicBounds() coregraphics.Rect {
	rv := ffi.CallMethod[coregraphics.Rect](t_, "typographicBounds")
	return rv
}

var TextTableClass = _TextTableClass{objc.GetClass("NSTextTable")}

type _TextTableClass struct {
	objc.Class
}

type ITextTable interface {
	ITextBlock
	NumberOfColumns() uint
	SetNumberOfColumns(value uint)
	LayoutAlgorithm() TextTableLayoutAlgorithm
	SetLayoutAlgorithm(value TextTableLayoutAlgorithm)
	CollapsesBorders() bool
	SetCollapsesBorders(value bool)
	HidesEmptyCells() bool
	SetHidesEmptyCells(value bool)
}

type TextTable struct {
	TextBlock
}

func MakeTextTable(ptr unsafe.Pointer) TextTable {
	return TextTable{
		TextBlock: MakeTextBlock(ptr),
	}
}

func (t_ TextTable) Init() TextTable {
	rv := ffi.CallMethod[TextTable](t_, "init")
	return rv
}

func (tc _TextTableClass) Alloc() TextTable {
	rv := ffi.CallMethod[TextTable](tc, "alloc")
	return rv
}

func (tc _TextTableClass) New() TextTable {
	rv := ffi.CallMethod[TextTable](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextTable() TextTable {
	return TextTableClass.New()
}

func (t_ TextTable) NumberOfColumns() uint {
	rv := ffi.CallMethod[uint](t_, "numberOfColumns")
	return rv
}

func (t_ TextTable) SetNumberOfColumns(value uint) {
	ffi.CallMethod[ffi.Void](t_, "setNumberOfColumns:", value)
}

func (t_ TextTable) LayoutAlgorithm() TextTableLayoutAlgorithm {
	rv := ffi.CallMethod[TextTableLayoutAlgorithm](t_, "layoutAlgorithm")
	return rv
}

func (t_ TextTable) SetLayoutAlgorithm(value TextTableLayoutAlgorithm) {
	ffi.CallMethod[ffi.Void](t_, "setLayoutAlgorithm:", value)
}

func (t_ TextTable) CollapsesBorders() bool {
	rv := ffi.CallMethod[bool](t_, "collapsesBorders")
	return rv
}

func (t_ TextTable) SetCollapsesBorders(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setCollapsesBorders:", value)
}

func (t_ TextTable) HidesEmptyCells() bool {
	rv := ffi.CallMethod[bool](t_, "hidesEmptyCells")
	return rv
}

func (t_ TextTable) SetHidesEmptyCells(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setHidesEmptyCells:", value)
}

var TextBlockClass = _TextBlockClass{objc.GetClass("NSTextBlock")}

type _TextBlockClass struct {
	objc.Class
}

type ITextBlock interface {
	objc.IObject
	SetValue_Type_ForDimension(val float64, _type TextBlockValueType, dimension TextBlockDimension)
	ValueForDimension(dimension TextBlockDimension) float64
	ValueTypeForDimension(dimension TextBlockDimension) TextBlockValueType
	SetContentWidth_Type(val float64, _type TextBlockValueType)
	SetWidth_Type_ForLayer_Edge(val float64, _type TextBlockValueType, layer TextBlockLayer, edge foundation.RectEdge)
	SetWidth_Type_ForLayer(val float64, _type TextBlockValueType, layer TextBlockLayer)
	WidthForLayer_Edge(layer TextBlockLayer, edge foundation.RectEdge) float64
	WidthValueTypeForLayer_Edge(layer TextBlockLayer, edge foundation.RectEdge) TextBlockValueType
	SetBorderColor_ForEdge(color IColor, edge foundation.RectEdge)
	SetBorderColor(color IColor)
	BorderColorForEdge(edge foundation.RectEdge) Color
	RectForLayoutAtPoint_InRect_TextContainer_CharacterRange(startingPoint foundation.Point, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect
	BoundsRectForContentRect_InRect_TextContainer_CharacterRange(contentRect foundation.Rect, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect
	DrawBackgroundWithFrame_InView_CharacterRange_LayoutManager(frameRect foundation.Rect, controlView IView, charRange foundation.Range, layoutManager ILayoutManager)
	ContentWidth() float64
	ContentWidthValueType() TextBlockValueType
	VerticalAlignment() TextBlockVerticalAlignment
	SetVerticalAlignment(value TextBlockVerticalAlignment)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
}

type TextBlock struct {
	objc.Object
}

func MakeTextBlock(ptr unsafe.Pointer) TextBlock {
	return TextBlock{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextBlock) Init() TextBlock {
	rv := ffi.CallMethod[TextBlock](t_, "init")
	return rv
}

func (tc _TextBlockClass) Alloc() TextBlock {
	rv := ffi.CallMethod[TextBlock](tc, "alloc")
	return rv
}

func (tc _TextBlockClass) New() TextBlock {
	rv := ffi.CallMethod[TextBlock](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextBlock() TextBlock {
	return TextBlockClass.New()
}

func (t_ TextBlock) SetValue_Type_ForDimension(val float64, _type TextBlockValueType, dimension TextBlockDimension) {
	ffi.CallMethod[ffi.Void](t_, "setValue:type:forDimension:", val, _type, dimension)
}

func (t_ TextBlock) ValueForDimension(dimension TextBlockDimension) float64 {
	rv := ffi.CallMethod[float64](t_, "valueForDimension:", dimension)
	return rv
}

func (t_ TextBlock) ValueTypeForDimension(dimension TextBlockDimension) TextBlockValueType {
	rv := ffi.CallMethod[TextBlockValueType](t_, "valueTypeForDimension:", dimension)
	return rv
}

func (t_ TextBlock) SetContentWidth_Type(val float64, _type TextBlockValueType) {
	ffi.CallMethod[ffi.Void](t_, "setContentWidth:type:", val, _type)
}

func (t_ TextBlock) SetWidth_Type_ForLayer_Edge(val float64, _type TextBlockValueType, layer TextBlockLayer, edge foundation.RectEdge) {
	ffi.CallMethod[ffi.Void](t_, "setWidth:type:forLayer:edge:", val, _type, layer, edge)
}

func (t_ TextBlock) SetWidth_Type_ForLayer(val float64, _type TextBlockValueType, layer TextBlockLayer) {
	ffi.CallMethod[ffi.Void](t_, "setWidth:type:forLayer:", val, _type, layer)
}

func (t_ TextBlock) WidthForLayer_Edge(layer TextBlockLayer, edge foundation.RectEdge) float64 {
	rv := ffi.CallMethod[float64](t_, "widthForLayer:edge:", layer, edge)
	return rv
}

func (t_ TextBlock) WidthValueTypeForLayer_Edge(layer TextBlockLayer, edge foundation.RectEdge) TextBlockValueType {
	rv := ffi.CallMethod[TextBlockValueType](t_, "widthValueTypeForLayer:edge:", layer, edge)
	return rv
}

func (t_ TextBlock) SetBorderColor_ForEdge(color IColor, edge foundation.RectEdge) {
	ffi.CallMethod[ffi.Void](t_, "setBorderColor:forEdge:", color, edge)
}

func (t_ TextBlock) SetBorderColor(color IColor) {
	ffi.CallMethod[ffi.Void](t_, "setBorderColor:", color)
}

func (t_ TextBlock) BorderColorForEdge(edge foundation.RectEdge) Color {
	rv := ffi.CallMethod[Color](t_, "borderColorForEdge:", edge)
	return rv
}

func (t_ TextBlock) RectForLayoutAtPoint_InRect_TextContainer_CharacterRange(startingPoint foundation.Point, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](t_, "rectForLayoutAtPoint:inRect:textContainer:characterRange:", startingPoint, rect, textContainer, charRange)
	return rv
}

func (t_ TextBlock) BoundsRectForContentRect_InRect_TextContainer_CharacterRange(contentRect foundation.Rect, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](t_, "boundsRectForContentRect:inRect:textContainer:characterRange:", contentRect, rect, textContainer, charRange)
	return rv
}

func (t_ TextBlock) DrawBackgroundWithFrame_InView_CharacterRange_LayoutManager(frameRect foundation.Rect, controlView IView, charRange foundation.Range, layoutManager ILayoutManager) {
	ffi.CallMethod[ffi.Void](t_, "drawBackgroundWithFrame:inView:characterRange:layoutManager:", frameRect, controlView, charRange, layoutManager)
}

func (t_ TextBlock) ContentWidth() float64 {
	rv := ffi.CallMethod[float64](t_, "contentWidth")
	return rv
}

func (t_ TextBlock) ContentWidthValueType() TextBlockValueType {
	rv := ffi.CallMethod[TextBlockValueType](t_, "contentWidthValueType")
	return rv
}

func (t_ TextBlock) VerticalAlignment() TextBlockVerticalAlignment {
	rv := ffi.CallMethod[TextBlockVerticalAlignment](t_, "verticalAlignment")
	return rv
}

func (t_ TextBlock) SetVerticalAlignment(value TextBlockVerticalAlignment) {
	ffi.CallMethod[ffi.Void](t_, "setVerticalAlignment:", value)
}

func (t_ TextBlock) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](t_, "backgroundColor")
	return rv
}

func (t_ TextBlock) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](t_, "setBackgroundColor:", value)
}
