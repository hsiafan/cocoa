// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TextLayoutManagerClass = _TextLayoutManagerClass{objc.GetClass("NSTextLayoutManager")}

type _TextLayoutManagerClass struct {
	objc.Class
}

type ITextLayoutManager interface {
	objc.IObject
	EnumerateTextSegmentsInRange_Type_Options_UsingBlock(textRange ITextRange, type_ TextLayoutManagerSegmentType, options TextLayoutManagerSegmentOptions, block func(textSegmentRange TextRange, textSegmentFrame coregraphics.Rect, baselinePosition float64, textContainer TextContainer) bool)
	ReplaceTextContentManager(textContentManager ITextContentManager)
	ReplaceContentsInRange_WithAttributedString(range_ ITextRange, attributedString foundation.IAttributedString)
	ReplaceContentsInRange_WithTextElements(range_ ITextRange, textElements []ITextElement)
	AddRenderingAttribute_Value_ForTextRange(renderingAttribute foundation.AttributedStringKey, value objc.IObject, textRange ITextRange)
	InvalidateRenderingAttributesForTextRange(textRange ITextRange)
	RemoveRenderingAttribute_ForTextRange(renderingAttribute foundation.AttributedStringKey, textRange ITextRange)
	SetRenderingAttributes_ForTextRange(renderingAttributes map[foundation.AttributedStringKey]objc.IObject, textRange ITextRange)
	InvalidateLayoutForRange(range_ ITextRange)
	TextLayoutFragmentForPosition(position coregraphics.Point) TextLayoutFragment
	EnsureLayoutForBounds(bounds coregraphics.Rect)
	EnsureLayoutForRange(range_ ITextRange)
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
	rv := objc.CallMethod[TextLayoutManager](t_, objc.SEL("init"))
	return rv
}

func (tc _TextLayoutManagerClass) Alloc() TextLayoutManager {
	rv := objc.CallMethod[TextLayoutManager](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TextLayoutManagerClass) New() TextLayoutManager {
	rv := objc.CallMethod[TextLayoutManager](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTextLayoutManager() TextLayoutManager {
	return TextLayoutManagerClass.New()
}

func (t_ TextLayoutManager) EnumerateTextSegmentsInRange_Type_Options_UsingBlock(textRange ITextRange, type_ TextLayoutManagerSegmentType, options TextLayoutManagerSegmentOptions, block func(textSegmentRange TextRange, textSegmentFrame coregraphics.Rect, baselinePosition float64, textContainer TextContainer) bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("enumerateTextSegmentsInRange:type:options:usingBlock:"), objc.ExtractPtr(textRange), type_, options, block)
}

func (t_ TextLayoutManager) ReplaceTextContentManager(textContentManager ITextContentManager) {
	objc.CallMethod[objc.Void](t_, objc.SEL("replaceTextContentManager:"), objc.ExtractPtr(textContentManager))
}

func (t_ TextLayoutManager) ReplaceContentsInRange_WithAttributedString(range_ ITextRange, attributedString foundation.IAttributedString) {
	objc.CallMethod[objc.Void](t_, objc.SEL("replaceContentsInRange:withAttributedString:"), objc.ExtractPtr(range_), objc.ExtractPtr(attributedString))
}

func (t_ TextLayoutManager) ReplaceContentsInRange_WithTextElements(range_ ITextRange, textElements []ITextElement) {
	objc.CallMethod[objc.Void](t_, objc.SEL("replaceContentsInRange:withTextElements:"), objc.ExtractPtr(range_), textElements)
}

func (t_ TextLayoutManager) AddRenderingAttribute_Value_ForTextRange(renderingAttribute foundation.AttributedStringKey, value objc.IObject, textRange ITextRange) {
	objc.CallMethod[objc.Void](t_, objc.SEL("addRenderingAttribute:value:forTextRange:"), renderingAttribute, objc.ExtractPtr(value), objc.ExtractPtr(textRange))
}

func (t_ TextLayoutManager) InvalidateRenderingAttributesForTextRange(textRange ITextRange) {
	objc.CallMethod[objc.Void](t_, objc.SEL("invalidateRenderingAttributesForTextRange:"), objc.ExtractPtr(textRange))
}

func (t_ TextLayoutManager) RemoveRenderingAttribute_ForTextRange(renderingAttribute foundation.AttributedStringKey, textRange ITextRange) {
	objc.CallMethod[objc.Void](t_, objc.SEL("removeRenderingAttribute:forTextRange:"), renderingAttribute, objc.ExtractPtr(textRange))
}

func (t_ TextLayoutManager) SetRenderingAttributes_ForTextRange(renderingAttributes map[foundation.AttributedStringKey]objc.IObject, textRange ITextRange) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setRenderingAttributes:forTextRange:"), renderingAttributes, objc.ExtractPtr(textRange))
}

func (t_ TextLayoutManager) InvalidateLayoutForRange(range_ ITextRange) {
	objc.CallMethod[objc.Void](t_, objc.SEL("invalidateLayoutForRange:"), objc.ExtractPtr(range_))
}

func (t_ TextLayoutManager) TextLayoutFragmentForPosition(position coregraphics.Point) TextLayoutFragment {
	rv := objc.CallMethod[TextLayoutFragment](t_, objc.SEL("textLayoutFragmentForPosition:"), position)
	return rv
}

func (t_ TextLayoutManager) EnsureLayoutForBounds(bounds coregraphics.Rect) {
	objc.CallMethod[objc.Void](t_, objc.SEL("ensureLayoutForBounds:"), bounds)
}

func (t_ TextLayoutManager) EnsureLayoutForRange(range_ ITextRange) {
	objc.CallMethod[objc.Void](t_, objc.SEL("ensureLayoutForRange:"), objc.ExtractPtr(range_))
}

func (t_ TextLayoutManager) LayoutQueue() foundation.OperationQueue {
	rv := objc.CallMethod[foundation.OperationQueue](t_, objc.SEL("layoutQueue"))
	return rv
}

func (t_ TextLayoutManager) SetLayoutQueue(value foundation.IOperationQueue) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setLayoutQueue:"), objc.ExtractPtr(value))
}

func (t_ TextLayoutManager) RenderingAttributesValidator() func(textLayoutManager ITextLayoutManager, textLayoutFragment ITextLayoutFragment) {
	rv := objc.CallMethod[func(textLayoutManager ITextLayoutManager, textLayoutFragment ITextLayoutFragment)](t_, objc.SEL("renderingAttributesValidator"))
	return rv
}

func (t_ TextLayoutManager) SetRenderingAttributesValidator(value func(textLayoutManager TextLayoutManager, textLayoutFragment TextLayoutFragment)) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setRenderingAttributesValidator:"), value)
}

func (t_ TextLayoutManager) UsesFontLeading() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("usesFontLeading"))
	return rv
}

func (t_ TextLayoutManager) SetUsesFontLeading(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setUsesFontLeading:"), value)
}

func (t_ TextLayoutManager) UsesHyphenation() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("usesHyphenation"))
	return rv
}

func (t_ TextLayoutManager) SetUsesHyphenation(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setUsesHyphenation:"), value)
}

func (t_ TextLayoutManager) LimitsLayoutForSuspiciousContents() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("limitsLayoutForSuspiciousContents"))
	return rv
}

func (t_ TextLayoutManager) SetLimitsLayoutForSuspiciousContents(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setLimitsLayoutForSuspiciousContents:"), value)
}

// weak property
func (t_ TextLayoutManager) TextContentManager() TextContentManager {
	rv := objc.CallMethod[TextContentManager](t_, objc.SEL("textContentManager"))
	return rv
}

func (t_ TextLayoutManager) TextContainer() TextContainer {
	rv := objc.CallMethod[TextContainer](t_, objc.SEL("textContainer"))
	return rv
}

func (t_ TextLayoutManager) SetTextContainer(value ITextContainer) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTextContainer:"), objc.ExtractPtr(value))
}

func (t_ TextLayoutManager) TextSelectionNavigation() TextSelectionNavigation {
	rv := objc.CallMethod[TextSelectionNavigation](t_, objc.SEL("textSelectionNavigation"))
	return rv
}

func (t_ TextLayoutManager) SetTextSelectionNavigation(value ITextSelectionNavigation) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTextSelectionNavigation:"), objc.ExtractPtr(value))
}

func (t_ TextLayoutManager) TextSelections() []TextSelection {
	rv := objc.CallMethod[[]TextSelection](t_, objc.SEL("textSelections"))
	return rv
}

func (t_ TextLayoutManager) SetTextSelections(value []ITextSelection) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTextSelections:"), value)
}

func (t_ TextLayoutManager) UsageBoundsForTextContainer() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](t_, objc.SEL("usageBoundsForTextContainer"))
	return rv
}

func (tc _TextLayoutManagerClass) LinkRenderingAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](tc, objc.SEL("linkRenderingAttributes"))
	return rv
}

func (t_ TextLayoutManager) TextViewportLayoutController() TextViewportLayoutController {
	rv := objc.CallMethod[TextViewportLayoutController](t_, objc.SEL("textViewportLayoutController"))
	return rv
}
