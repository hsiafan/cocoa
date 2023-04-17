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

func (t_ TextLayoutManager) EnumerateTextSegmentsInRange_Type_Options_UsingBlock(textRange ITextRange, type_ TextLayoutManagerSegmentType, options TextLayoutManagerSegmentOptions, block func(textSegmentRange TextRange, textSegmentFrame coregraphics.Rect, baselinePosition float64, textContainer TextContainer) bool) {
	ffi.CallMethod[ffi.Void](t_, "enumerateTextSegmentsInRange:type:options:usingBlock:", textRange, type_, options, block)
}

func (t_ TextLayoutManager) ReplaceTextContentManager(textContentManager ITextContentManager) {
	ffi.CallMethod[ffi.Void](t_, "replaceTextContentManager:", textContentManager)
}

func (t_ TextLayoutManager) ReplaceContentsInRange_WithAttributedString(range_ ITextRange, attributedString foundation.IAttributedString) {
	ffi.CallMethod[ffi.Void](t_, "replaceContentsInRange:withAttributedString:", range_, attributedString)
}

func (t_ TextLayoutManager) ReplaceContentsInRange_WithTextElements(range_ ITextRange, textElements []ITextElement) {
	ffi.CallMethod[ffi.Void](t_, "replaceContentsInRange:withTextElements:", range_, textElements)
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

func (t_ TextLayoutManager) InvalidateLayoutForRange(range_ ITextRange) {
	ffi.CallMethod[ffi.Void](t_, "invalidateLayoutForRange:", range_)
}

func (t_ TextLayoutManager) TextLayoutFragmentForPosition(position coregraphics.Point) TextLayoutFragment {
	rv := ffi.CallMethod[TextLayoutFragment](t_, "textLayoutFragmentForPosition:", position)
	return rv
}

func (t_ TextLayoutManager) EnsureLayoutForBounds(bounds coregraphics.Rect) {
	ffi.CallMethod[ffi.Void](t_, "ensureLayoutForBounds:", bounds)
}

func (t_ TextLayoutManager) EnsureLayoutForRange(range_ ITextRange) {
	ffi.CallMethod[ffi.Void](t_, "ensureLayoutForRange:", range_)
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
