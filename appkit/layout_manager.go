// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var LayoutManagerClass = _LayoutManagerClass{objc.GetClass("NSLayoutManager")}

type _LayoutManagerClass struct {
	objc.Class
}

type ILayoutManager interface {
	objc.IObject
	ReplaceTextStorage(newTextStorage ITextStorage)
	AddTextContainer(container ITextContainer)
	InsertTextContainer_AtIndex(container ITextContainer, index uint)
	RemoveTextContainerAtIndex(index uint)
	SetTextContainer_ForGlyphRange(container ITextContainer, glyphRange foundation.Range)
	TextContainerChangedGeometry(container ITextContainer)
	TextContainerChangedTextView(container ITextContainer)
	TextContainerForGlyphAtIndex_EffectiveRange(glyphIndex uint, effectiveGlyphRange *foundation.Range) TextContainer
	TextContainerForGlyphAtIndex_EffectiveRange_WithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange *foundation.Range, flag bool) TextContainer
	UsedRectForTextContainer(container ITextContainer) foundation.Rect
	InvalidateDisplayForCharacterRange(charRange foundation.Range)
	InvalidateDisplayForGlyphRange(glyphRange foundation.Range)
	InvalidateGlyphsForCharacterRange_ChangeInLength_ActualCharacterRange(charRange foundation.Range, delta int, actualCharRange *foundation.Range)
	InvalidateLayoutForCharacterRange_ActualCharacterRange(charRange foundation.Range, actualCharRange *foundation.Range)
	ProcessEditingForTextStorage_Edited_Range_ChangeInLength_InvalidatedRange(textStorage ITextStorage, editMask TextStorageEditActions, newCharRange foundation.Range, delta int, invalidatedCharRange foundation.Range)
	EnsureGlyphsForCharacterRange(charRange foundation.Range)
	EnsureGlyphsForGlyphRange(glyphRange foundation.Range)
	EnsureLayoutForBoundingRect_InTextContainer(bounds foundation.Rect, container ITextContainer)
	EnsureLayoutForCharacterRange(charRange foundation.Range)
	EnsureLayoutForGlyphRange(glyphRange foundation.Range)
	EnsureLayoutForTextContainer(container ITextContainer)
	GetGlyphsInRange_Glyphs_Properties_CharacterIndexes_BidiLevels(glyphRange foundation.Range, glyphBuffer *coregraphics.Glyph, props *GlyphProperty, charIndexBuffer *uint, bidiLevelBuffer *byte) uint
	CGGlyphAtIndex(glyphIndex uint) coregraphics.Glyph
	CGGlyphAtIndex_IsValidIndex(glyphIndex uint, isValidIndex *bool) coregraphics.Glyph
	SetGlyphs_Properties_CharacterIndexes_Font_ForGlyphRange(glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont IFont, glyphRange foundation.Range)
	CharacterIndexForGlyphAtIndex(glyphIndex uint) uint
	GlyphIndexForCharacterAtIndex(charIndex uint) uint
	IsValidGlyphIndex(glyphIndex uint) bool
	PropertyForGlyphAtIndex(glyphIndex uint) GlyphProperty
	SetAttachmentSize_ForGlyphRange(attachmentSize foundation.Size, glyphRange foundation.Range)
	SetDrawsOutsideLineFragment_ForGlyphAtIndex(flag bool, glyphIndex uint)
	SetExtraLineFragmentRect_UsedRect_TextContainer(fragmentRect foundation.Rect, usedRect foundation.Rect, container ITextContainer)
	SetLineFragmentRect_ForGlyphRange_UsedRect(fragmentRect foundation.Rect, glyphRange foundation.Range, usedRect foundation.Rect)
	SetLocation_ForStartOfGlyphRange(location foundation.Point, glyphRange foundation.Range)
	SetNotShownAttribute_ForGlyphAtIndex(flag bool, glyphIndex uint)
	AttachmentSizeForGlyphAtIndex(glyphIndex uint) foundation.Size
	DrawsOutsideLineFragmentForGlyphAtIndex(glyphIndex uint) bool
	FirstUnlaidCharacterIndex() uint
	FirstUnlaidGlyphIndex() uint
	GetFirstUnlaidCharacterIndex_GlyphIndex(charIndex *uint, glyphIndex *uint)
	LineFragmentRectForGlyphAtIndex_EffectiveRange(glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect
	LineFragmentRectForGlyphAtIndex_EffectiveRange_WithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange *foundation.Range, flag bool) foundation.Rect
	LineFragmentUsedRectForGlyphAtIndex_EffectiveRange(glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect
	LineFragmentUsedRectForGlyphAtIndex_EffectiveRange_WithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange *foundation.Range, flag bool) foundation.Rect
	LocationForGlyphAtIndex(glyphIndex uint) foundation.Point
	NotShownAttributeForGlyphAtIndex(glyphIndex uint) bool
	TruncatedGlyphRangeInLineFragmentForGlyphAtIndex(glyphIndex uint) foundation.Range
	BoundingRectForGlyphRange_InTextContainer(glyphRange foundation.Range, container ITextContainer) foundation.Rect
	CharacterIndexForPoint_InTextContainer_FractionOfDistanceBetweenInsertionPoints(point foundation.Point, container ITextContainer, partialFraction *float64) uint
	CharacterRangeForGlyphRange_ActualGlyphRange(glyphRange foundation.Range, actualGlyphRange *foundation.Range) foundation.Range
	EnumerateEnclosingRectsForGlyphRange_WithinSelectedGlyphRange_InTextContainer_UsingBlock(glyphRange foundation.Range, selectedRange foundation.Range, textContainer ITextContainer, block func(rect foundation.Rect, stop *bool))
	EnumerateLineFragmentsForGlyphRange_UsingBlock(glyphRange foundation.Range, block func(rect foundation.Rect, usedRect foundation.Rect, textContainer TextContainer, glyphRange foundation.Range, stop *bool))
	FractionOfDistanceThroughGlyphForPoint_InTextContainer(point foundation.Point, container ITextContainer) float64
	GetLineFragmentInsertionPointsForCharacterAtIndex_AlternatePositions_InDisplayOrder_Positions_CharacterIndexes(charIndex uint, aFlag bool, dFlag bool, positions *float64, charIndexes *uint) uint
	GlyphIndexForPoint_InTextContainer(point foundation.Point, container ITextContainer) uint
	GlyphIndexForPoint_InTextContainer_FractionOfDistanceThroughGlyph(point foundation.Point, container ITextContainer, partialFraction *float64) uint
	GlyphRangeForBoundingRect_InTextContainer(bounds foundation.Rect, container ITextContainer) foundation.Range
	GlyphRangeForBoundingRectWithoutAdditionalLayout_InTextContainer(bounds foundation.Rect, container ITextContainer) foundation.Range
	GlyphRangeForTextContainer(container ITextContainer) foundation.Range
	GlyphRangeForCharacterRange_ActualCharacterRange(charRange foundation.Range, actualCharRange *foundation.Range) foundation.Range
	RangeOfNominallySpacedGlyphsContainingIndex(glyphIndex uint) foundation.Range
	DrawBackgroundForGlyphRange_AtPoint(glyphsToShow foundation.Range, origin foundation.Point)
	DrawGlyphsForGlyphRange_AtPoint(glyphsToShow foundation.Range, origin foundation.Point)
	DrawStrikethroughForGlyphRange_StrikethroughType_BaselineOffset_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, strikethroughVal UnderlineStyle, baselineOffset float64, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point)
	DrawUnderlineForGlyphRange_UnderlineType_BaselineOffset_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, underlineVal UnderlineStyle, baselineOffset float64, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point)
	FillBackgroundRectArray_Count_ForCharacterRange_Color(rectArray *foundation.Rect, rectCount uint, charRange foundation.Range, color IColor)
	ShowCGGlyphs_Positions_Count_Font_TextMatrix_Attributes_InContext(glyphs *coregraphics.Glyph, positions *coregraphics.Point, glyphCount int, font IFont, textMatrix coregraphics.AffineTransform, attributes map[foundation.AttributedStringKey]objc.IObject, CGContext coregraphics.ContextRef)
	StrikethroughGlyphRange_StrikethroughType_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, strikethroughVal UnderlineStyle, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point)
	UnderlineGlyphRange_UnderlineType_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, underlineVal UnderlineStyle, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point)
	SetLayoutRect_ForTextBlock_GlyphRange(rect foundation.Rect, block ITextBlock, glyphRange foundation.Range)
	LayoutRectForTextBlock_GlyphRange(block ITextBlock, glyphRange foundation.Range) foundation.Rect
	SetBoundsRect_ForTextBlock_GlyphRange(rect foundation.Rect, block ITextBlock, glyphRange foundation.Range)
	BoundsRectForTextBlock_GlyphRange(block ITextBlock, glyphRange foundation.Range) foundation.Rect
	LayoutRectForTextBlock_AtIndex_EffectiveRange(block ITextBlock, glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect
	BoundsRectForTextBlock_AtIndex_EffectiveRange(block ITextBlock, glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect
	ShowAttachmentCell_InRect_CharacterIndex(cell ICell, rect foundation.Rect, attachmentIndex uint)
	RulerAccessoryViewForTextView_ParagraphStyle_Ruler_Enabled(view ITextView, style IParagraphStyle, ruler IRulerView, isEnabled bool) View
	RulerMarkersForTextView_ParagraphStyle_Ruler(view ITextView, style IParagraphStyle, ruler IRulerView) []RulerMarker
	LayoutManagerOwnsFirstResponderInWindow(window IWindow) bool
	DefaultLineHeightForFont(theFont IFont) float64
	DefaultBaselineOffsetForFont(theFont IFont) float64
	AddTemporaryAttributes_ForCharacterRange(attrs map[foundation.AttributedStringKey]objc.IObject, charRange foundation.Range)
	AddTemporaryAttribute_Value_ForCharacterRange(attrName foundation.AttributedStringKey, value objc.IObject, charRange foundation.Range)
	SetTemporaryAttributes_ForCharacterRange(attrs map[foundation.AttributedStringKey]objc.IObject, charRange foundation.Range)
	RemoveTemporaryAttribute_ForCharacterRange(attrName foundation.AttributedStringKey, charRange foundation.Range)
	TemporaryAttribute_AtCharacterIndex_EffectiveRange(attrName foundation.AttributedStringKey, location uint, range_ *foundation.Range) objc.Object
	TemporaryAttribute_AtCharacterIndex_LongestEffectiveRange_InRange(attrName foundation.AttributedStringKey, location uint, range_ *foundation.Range, rangeLimit foundation.Range) objc.Object
	TemporaryAttributesAtCharacterIndex_EffectiveRange(charIndex uint, effectiveCharRange *foundation.Range) map[foundation.AttributedStringKey]objc.Object
	TemporaryAttributesAtCharacterIndex_LongestEffectiveRange_InRange(location uint, range_ *foundation.Range, rangeLimit foundation.Range) map[foundation.AttributedStringKey]objc.Object
	// deprecated
	ShowCGGlyphs_Positions_Count_Font_Matrix_Attributes_InContext(glyphs *coregraphics.Glyph, positions *foundation.Point, glyphCount uint, font IFont, textMatrix foundation.IAffineTransform, attributes map[foundation.AttributedStringKey]objc.IObject, graphicsContext IGraphicsContext)
	// deprecated
	InvalidateGlyphsOnLayoutInvalidationForGlyphRange(glyphRange foundation.Range)
	// deprecated
	InvalidateLayoutForCharacterRange_IsSoft_ActualCharacterRange(charRange foundation.Range, flag bool, actualCharRange *foundation.Range)
	// deprecated
	TextStorage_Edited_Range_ChangeInLength_InvalidatedRange(str ITextStorage, editedMask TextStorageEditedOptions, newCharRange foundation.Range, delta int, invalidatedCharRange foundation.Range)
	// deprecated
	InsertGlyph_AtGlyphIndex_CharacterIndex(glyph Glyph, glyphIndex uint, charIndex uint)
	// deprecated
	InsertGlyphs_Length_ForStartingGlyphAtIndex_CharacterIndex(glyphs *Glyph, length uint, glyphIndex uint, charIndex uint)
	// deprecated
	GlyphAtIndex(glyphIndex uint) Glyph
	// deprecated
	GlyphAtIndex_IsValidIndex(glyphIndex uint, isValidIndex *bool) Glyph
	// deprecated
	ReplaceGlyphAtIndex_WithGlyph(glyphIndex uint, newGlyph Glyph)
	// deprecated
	GetGlyphs_Range(glyphArray *Glyph, glyphRange foundation.Range) uint
	// deprecated
	GetGlyphsInRange_Glyphs_CharacterIndexes_GlyphInscriptions_ElasticBits(glyphRange foundation.Range, glyphBuffer *Glyph, charIndexBuffer *uint, inscribeBuffer *GlyphInscription, elasticBuffer *bool) uint
	// deprecated
	GetGlyphsInRange_Glyphs_CharacterIndexes_GlyphInscriptions_ElasticBits_BidiLevels(glyphRange foundation.Range, glyphBuffer *Glyph, charIndexBuffer *uint, inscribeBuffer *GlyphInscription, elasticBuffer *bool, bidiLevelBuffer *byte) uint
	// deprecated
	DeleteGlyphsInRange(glyphRange foundation.Range)
	// deprecated
	SetCharacterIndex_ForGlyphAtIndex(charIndex uint, glyphIndex uint)
	// deprecated
	IntAttribute_ForGlyphAtIndex(attributeTag int, glyphIndex uint) int
	// deprecated
	SetIntAttribute_Value_ForGlyphAtIndex(attributeTag int, val int, glyphIndex uint)
	// deprecated
	SetLocations_StartingGlyphIndexes_Count_ForGlyphRange(locations *foundation.Point, glyphIndexes *uint, count uint, glyphRange foundation.Range)
	// deprecated
	RectArrayForCharacterRange_WithinSelectedCharacterRange_InTextContainer_RectCount(charRange foundation.Range, selCharRange foundation.Range, container ITextContainer, rectCount *uint) *foundation.Rect
	// deprecated
	RectArrayForGlyphRange_WithinSelectedGlyphRange_InTextContainer_RectCount(glyphRange foundation.Range, selGlyphRange foundation.Range, container ITextContainer, rectCount *uint) *foundation.Rect
	// deprecated
	SubstituteFontForFont(originalFont IFont) Font
	// deprecated
	ShowPackedGlyphs_Length_GlyphRange_AtPoint_Font_Color_PrintingAdjustment(glyphs *byte, glyphLen uint, glyphRange foundation.Range, point foundation.Point, font IFont, color IColor, printingAdjustment foundation.Size)
	Delegate() LayoutManagerDelegateWrapper
	SetDelegate(value LayoutManagerDelegate)
	SetDelegate0(value objc.IObject)
	TextStorage() TextStorage
	SetTextStorage(value ITextStorage)
	AllowsNonContiguousLayout() bool
	SetAllowsNonContiguousLayout(value bool)
	HasNonContiguousLayout() bool
	ShowsInvisibleCharacters() bool
	SetShowsInvisibleCharacters(value bool)
	ShowsControlCharacters() bool
	SetShowsControlCharacters(value bool)
	UsesFontLeading() bool
	SetUsesFontLeading(value bool)
	BackgroundLayoutEnabled() bool
	SetBackgroundLayoutEnabled(value bool)
	LimitsLayoutForSuspiciousContents() bool
	SetLimitsLayoutForSuspiciousContents(value bool)
	UsesDefaultHyphenation() bool
	SetUsesDefaultHyphenation(value bool)
	TextContainers() []TextContainer
	GlyphGenerator() GlyphGenerator
	SetGlyphGenerator(value IGlyphGenerator)
	NumberOfGlyphs() uint
	ExtraLineFragmentRect() foundation.Rect
	ExtraLineFragmentTextContainer() TextContainer
	ExtraLineFragmentUsedRect() foundation.Rect
	DefaultAttachmentScaling() ImageScaling
	SetDefaultAttachmentScaling(value ImageScaling)
	FirstTextView() TextView
	TextViewForBeginningOfSelection() TextView
	Typesetter() Typesetter
	SetTypesetter(value ITypesetter)
	TypesetterBehavior() TypesetterBehavior
	SetTypesetterBehavior(value TypesetterBehavior)
	// deprecated
	HyphenationFactor() float32
	// deprecated
	SetHyphenationFactor(value float32)
	// deprecated
	UsesScreenFonts() bool
	// deprecated
	SetUsesScreenFonts(value bool)
}

type LayoutManager struct {
	objc.Object
}

func MakeLayoutManager(ptr unsafe.Pointer) LayoutManager {
	return LayoutManager{
		Object: objc.MakeObject(ptr),
	}
}

func (l_ LayoutManager) Init() LayoutManager {
	rv := objc.CallMethod[LayoutManager](l_, objc.GetSelector("init"))
	return rv
}

func (lc _LayoutManagerClass) Alloc() LayoutManager {
	rv := objc.CallMethod[LayoutManager](lc, objc.GetSelector("alloc"))
	return rv
}

func (lc _LayoutManagerClass) New() LayoutManager {
	rv := objc.CallMethod[LayoutManager](lc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutManager() LayoutManager {
	return LayoutManagerClass.New()
}

func (l_ LayoutManager) ReplaceTextStorage(newTextStorage ITextStorage) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("replaceTextStorage:"), newTextStorage)
}

func (l_ LayoutManager) AddTextContainer(container ITextContainer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("addTextContainer:"), container)
}

func (l_ LayoutManager) InsertTextContainer_AtIndex(container ITextContainer, index uint) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("insertTextContainer:atIndex:"), container, index)
}

func (l_ LayoutManager) RemoveTextContainerAtIndex(index uint) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("removeTextContainerAtIndex:"), index)
}

func (l_ LayoutManager) SetTextContainer_ForGlyphRange(container ITextContainer, glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setTextContainer:forGlyphRange:"), container, glyphRange)
}

func (l_ LayoutManager) TextContainerChangedGeometry(container ITextContainer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("textContainerChangedGeometry:"), container)
}

func (l_ LayoutManager) TextContainerChangedTextView(container ITextContainer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("textContainerChangedTextView:"), container)
}

func (l_ LayoutManager) TextContainerForGlyphAtIndex_EffectiveRange(glyphIndex uint, effectiveGlyphRange *foundation.Range) TextContainer {
	rv := objc.CallMethod[TextContainer](l_, objc.GetSelector("textContainerForGlyphAtIndex:effectiveRange:"), glyphIndex, effectiveGlyphRange)
	return rv
}

func (l_ LayoutManager) TextContainerForGlyphAtIndex_EffectiveRange_WithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange *foundation.Range, flag bool) TextContainer {
	rv := objc.CallMethod[TextContainer](l_, objc.GetSelector("textContainerForGlyphAtIndex:effectiveRange:withoutAdditionalLayout:"), glyphIndex, effectiveGlyphRange, flag)
	return rv
}

func (l_ LayoutManager) UsedRectForTextContainer(container ITextContainer) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("usedRectForTextContainer:"), container)
	return rv
}

func (l_ LayoutManager) InvalidateDisplayForCharacterRange(charRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("invalidateDisplayForCharacterRange:"), charRange)
}

func (l_ LayoutManager) InvalidateDisplayForGlyphRange(glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("invalidateDisplayForGlyphRange:"), glyphRange)
}

func (l_ LayoutManager) InvalidateGlyphsForCharacterRange_ChangeInLength_ActualCharacterRange(charRange foundation.Range, delta int, actualCharRange *foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("invalidateGlyphsForCharacterRange:changeInLength:actualCharacterRange:"), charRange, delta, actualCharRange)
}

func (l_ LayoutManager) InvalidateLayoutForCharacterRange_ActualCharacterRange(charRange foundation.Range, actualCharRange *foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("invalidateLayoutForCharacterRange:actualCharacterRange:"), charRange, actualCharRange)
}

func (l_ LayoutManager) ProcessEditingForTextStorage_Edited_Range_ChangeInLength_InvalidatedRange(textStorage ITextStorage, editMask TextStorageEditActions, newCharRange foundation.Range, delta int, invalidatedCharRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("processEditingForTextStorage:edited:range:changeInLength:invalidatedRange:"), textStorage, editMask, newCharRange, delta, invalidatedCharRange)
}

func (l_ LayoutManager) EnsureGlyphsForCharacterRange(charRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("ensureGlyphsForCharacterRange:"), charRange)
}

func (l_ LayoutManager) EnsureGlyphsForGlyphRange(glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("ensureGlyphsForGlyphRange:"), glyphRange)
}

func (l_ LayoutManager) EnsureLayoutForBoundingRect_InTextContainer(bounds foundation.Rect, container ITextContainer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("ensureLayoutForBoundingRect:inTextContainer:"), bounds, container)
}

func (l_ LayoutManager) EnsureLayoutForCharacterRange(charRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("ensureLayoutForCharacterRange:"), charRange)
}

func (l_ LayoutManager) EnsureLayoutForGlyphRange(glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("ensureLayoutForGlyphRange:"), glyphRange)
}

func (l_ LayoutManager) EnsureLayoutForTextContainer(container ITextContainer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("ensureLayoutForTextContainer:"), container)
}

func (l_ LayoutManager) GetGlyphsInRange_Glyphs_Properties_CharacterIndexes_BidiLevels(glyphRange foundation.Range, glyphBuffer *coregraphics.Glyph, props *GlyphProperty, charIndexBuffer *uint, bidiLevelBuffer *byte) uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("getGlyphsInRange:glyphs:properties:characterIndexes:bidiLevels:"), glyphRange, glyphBuffer, props, charIndexBuffer, bidiLevelBuffer)
	return rv
}

func (l_ LayoutManager) CGGlyphAtIndex(glyphIndex uint) coregraphics.Glyph {
	rv := objc.CallMethod[coregraphics.Glyph](l_, objc.GetSelector("CGGlyphAtIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) CGGlyphAtIndex_IsValidIndex(glyphIndex uint, isValidIndex *bool) coregraphics.Glyph {
	rv := objc.CallMethod[coregraphics.Glyph](l_, objc.GetSelector("CGGlyphAtIndex:isValidIndex:"), glyphIndex, isValidIndex)
	return rv
}

func (l_ LayoutManager) SetGlyphs_Properties_CharacterIndexes_Font_ForGlyphRange(glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont IFont, glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setGlyphs:properties:characterIndexes:font:forGlyphRange:"), glyphs, props, charIndexes, aFont, glyphRange)
}

func (l_ LayoutManager) CharacterIndexForGlyphAtIndex(glyphIndex uint) uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("characterIndexForGlyphAtIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) GlyphIndexForCharacterAtIndex(charIndex uint) uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("glyphIndexForCharacterAtIndex:"), charIndex)
	return rv
}

func (l_ LayoutManager) IsValidGlyphIndex(glyphIndex uint) bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("isValidGlyphIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) PropertyForGlyphAtIndex(glyphIndex uint) GlyphProperty {
	rv := objc.CallMethod[GlyphProperty](l_, objc.GetSelector("propertyForGlyphAtIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) SetAttachmentSize_ForGlyphRange(attachmentSize foundation.Size, glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setAttachmentSize:forGlyphRange:"), attachmentSize, glyphRange)
}

func (l_ LayoutManager) SetDrawsOutsideLineFragment_ForGlyphAtIndex(flag bool, glyphIndex uint) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setDrawsOutsideLineFragment:forGlyphAtIndex:"), flag, glyphIndex)
}

func (l_ LayoutManager) SetExtraLineFragmentRect_UsedRect_TextContainer(fragmentRect foundation.Rect, usedRect foundation.Rect, container ITextContainer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setExtraLineFragmentRect:usedRect:textContainer:"), fragmentRect, usedRect, container)
}

func (l_ LayoutManager) SetLineFragmentRect_ForGlyphRange_UsedRect(fragmentRect foundation.Rect, glyphRange foundation.Range, usedRect foundation.Rect) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setLineFragmentRect:forGlyphRange:usedRect:"), fragmentRect, glyphRange, usedRect)
}

func (l_ LayoutManager) SetLocation_ForStartOfGlyphRange(location foundation.Point, glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setLocation:forStartOfGlyphRange:"), location, glyphRange)
}

func (l_ LayoutManager) SetNotShownAttribute_ForGlyphAtIndex(flag bool, glyphIndex uint) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setNotShownAttribute:forGlyphAtIndex:"), flag, glyphIndex)
}

func (l_ LayoutManager) AttachmentSizeForGlyphAtIndex(glyphIndex uint) foundation.Size {
	rv := objc.CallMethod[foundation.Size](l_, objc.GetSelector("attachmentSizeForGlyphAtIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) DrawsOutsideLineFragmentForGlyphAtIndex(glyphIndex uint) bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("drawsOutsideLineFragmentForGlyphAtIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) FirstUnlaidCharacterIndex() uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("firstUnlaidCharacterIndex"))
	return rv
}

func (l_ LayoutManager) FirstUnlaidGlyphIndex() uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("firstUnlaidGlyphIndex"))
	return rv
}

func (l_ LayoutManager) GetFirstUnlaidCharacterIndex_GlyphIndex(charIndex *uint, glyphIndex *uint) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("getFirstUnlaidCharacterIndex:glyphIndex:"), charIndex, glyphIndex)
}

func (l_ LayoutManager) LineFragmentRectForGlyphAtIndex_EffectiveRange(glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("lineFragmentRectForGlyphAtIndex:effectiveRange:"), glyphIndex, effectiveGlyphRange)
	return rv
}

func (l_ LayoutManager) LineFragmentRectForGlyphAtIndex_EffectiveRange_WithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange *foundation.Range, flag bool) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("lineFragmentRectForGlyphAtIndex:effectiveRange:withoutAdditionalLayout:"), glyphIndex, effectiveGlyphRange, flag)
	return rv
}

func (l_ LayoutManager) LineFragmentUsedRectForGlyphAtIndex_EffectiveRange(glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("lineFragmentUsedRectForGlyphAtIndex:effectiveRange:"), glyphIndex, effectiveGlyphRange)
	return rv
}

func (l_ LayoutManager) LineFragmentUsedRectForGlyphAtIndex_EffectiveRange_WithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange *foundation.Range, flag bool) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("lineFragmentUsedRectForGlyphAtIndex:effectiveRange:withoutAdditionalLayout:"), glyphIndex, effectiveGlyphRange, flag)
	return rv
}

func (l_ LayoutManager) LocationForGlyphAtIndex(glyphIndex uint) foundation.Point {
	rv := objc.CallMethod[foundation.Point](l_, objc.GetSelector("locationForGlyphAtIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) NotShownAttributeForGlyphAtIndex(glyphIndex uint) bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("notShownAttributeForGlyphAtIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) TruncatedGlyphRangeInLineFragmentForGlyphAtIndex(glyphIndex uint) foundation.Range {
	rv := objc.CallMethod[foundation.Range](l_, objc.GetSelector("truncatedGlyphRangeInLineFragmentForGlyphAtIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) BoundingRectForGlyphRange_InTextContainer(glyphRange foundation.Range, container ITextContainer) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("boundingRectForGlyphRange:inTextContainer:"), glyphRange, container)
	return rv
}

func (l_ LayoutManager) CharacterIndexForPoint_InTextContainer_FractionOfDistanceBetweenInsertionPoints(point foundation.Point, container ITextContainer, partialFraction *float64) uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("characterIndexForPoint:inTextContainer:fractionOfDistanceBetweenInsertionPoints:"), point, container, partialFraction)
	return rv
}

func (l_ LayoutManager) CharacterRangeForGlyphRange_ActualGlyphRange(glyphRange foundation.Range, actualGlyphRange *foundation.Range) foundation.Range {
	rv := objc.CallMethod[foundation.Range](l_, objc.GetSelector("characterRangeForGlyphRange:actualGlyphRange:"), glyphRange, actualGlyphRange)
	return rv
}

func (l_ LayoutManager) EnumerateEnclosingRectsForGlyphRange_WithinSelectedGlyphRange_InTextContainer_UsingBlock(glyphRange foundation.Range, selectedRange foundation.Range, textContainer ITextContainer, block func(rect foundation.Rect, stop *bool)) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("enumerateEnclosingRectsForGlyphRange:withinSelectedGlyphRange:inTextContainer:usingBlock:"), glyphRange, selectedRange, textContainer, block)
}

func (l_ LayoutManager) EnumerateLineFragmentsForGlyphRange_UsingBlock(glyphRange foundation.Range, block func(rect foundation.Rect, usedRect foundation.Rect, textContainer TextContainer, glyphRange foundation.Range, stop *bool)) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("enumerateLineFragmentsForGlyphRange:usingBlock:"), glyphRange, block)
}

func (l_ LayoutManager) FractionOfDistanceThroughGlyphForPoint_InTextContainer(point foundation.Point, container ITextContainer) float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("fractionOfDistanceThroughGlyphForPoint:inTextContainer:"), point, container)
	return rv
}

func (l_ LayoutManager) GetLineFragmentInsertionPointsForCharacterAtIndex_AlternatePositions_InDisplayOrder_Positions_CharacterIndexes(charIndex uint, aFlag bool, dFlag bool, positions *float64, charIndexes *uint) uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("getLineFragmentInsertionPointsForCharacterAtIndex:alternatePositions:inDisplayOrder:positions:characterIndexes:"), charIndex, aFlag, dFlag, positions, charIndexes)
	return rv
}

func (l_ LayoutManager) GlyphIndexForPoint_InTextContainer(point foundation.Point, container ITextContainer) uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("glyphIndexForPoint:inTextContainer:"), point, container)
	return rv
}

func (l_ LayoutManager) GlyphIndexForPoint_InTextContainer_FractionOfDistanceThroughGlyph(point foundation.Point, container ITextContainer, partialFraction *float64) uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("glyphIndexForPoint:inTextContainer:fractionOfDistanceThroughGlyph:"), point, container, partialFraction)
	return rv
}

func (l_ LayoutManager) GlyphRangeForBoundingRect_InTextContainer(bounds foundation.Rect, container ITextContainer) foundation.Range {
	rv := objc.CallMethod[foundation.Range](l_, objc.GetSelector("glyphRangeForBoundingRect:inTextContainer:"), bounds, container)
	return rv
}

func (l_ LayoutManager) GlyphRangeForBoundingRectWithoutAdditionalLayout_InTextContainer(bounds foundation.Rect, container ITextContainer) foundation.Range {
	rv := objc.CallMethod[foundation.Range](l_, objc.GetSelector("glyphRangeForBoundingRectWithoutAdditionalLayout:inTextContainer:"), bounds, container)
	return rv
}

func (l_ LayoutManager) GlyphRangeForTextContainer(container ITextContainer) foundation.Range {
	rv := objc.CallMethod[foundation.Range](l_, objc.GetSelector("glyphRangeForTextContainer:"), container)
	return rv
}

func (l_ LayoutManager) GlyphRangeForCharacterRange_ActualCharacterRange(charRange foundation.Range, actualCharRange *foundation.Range) foundation.Range {
	rv := objc.CallMethod[foundation.Range](l_, objc.GetSelector("glyphRangeForCharacterRange:actualCharacterRange:"), charRange, actualCharRange)
	return rv
}

func (l_ LayoutManager) RangeOfNominallySpacedGlyphsContainingIndex(glyphIndex uint) foundation.Range {
	rv := objc.CallMethod[foundation.Range](l_, objc.GetSelector("rangeOfNominallySpacedGlyphsContainingIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) DrawBackgroundForGlyphRange_AtPoint(glyphsToShow foundation.Range, origin foundation.Point) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("drawBackgroundForGlyphRange:atPoint:"), glyphsToShow, origin)
}

func (l_ LayoutManager) DrawGlyphsForGlyphRange_AtPoint(glyphsToShow foundation.Range, origin foundation.Point) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("drawGlyphsForGlyphRange:atPoint:"), glyphsToShow, origin)
}

func (l_ LayoutManager) DrawStrikethroughForGlyphRange_StrikethroughType_BaselineOffset_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, strikethroughVal UnderlineStyle, baselineOffset float64, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("drawStrikethroughForGlyphRange:strikethroughType:baselineOffset:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, strikethroughVal, baselineOffset, lineRect, lineGlyphRange, containerOrigin)
}

func (l_ LayoutManager) DrawUnderlineForGlyphRange_UnderlineType_BaselineOffset_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, underlineVal UnderlineStyle, baselineOffset float64, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("drawUnderlineForGlyphRange:underlineType:baselineOffset:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, underlineVal, baselineOffset, lineRect, lineGlyphRange, containerOrigin)
}

func (l_ LayoutManager) FillBackgroundRectArray_Count_ForCharacterRange_Color(rectArray *foundation.Rect, rectCount uint, charRange foundation.Range, color IColor) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("fillBackgroundRectArray:count:forCharacterRange:color:"), rectArray, rectCount, charRange, color)
}

func (l_ LayoutManager) ShowCGGlyphs_Positions_Count_Font_TextMatrix_Attributes_InContext(glyphs *coregraphics.Glyph, positions *coregraphics.Point, glyphCount int, font IFont, textMatrix coregraphics.AffineTransform, attributes map[foundation.AttributedStringKey]objc.IObject, CGContext coregraphics.ContextRef) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("showCGGlyphs:positions:count:font:textMatrix:attributes:inContext:"), glyphs, positions, glyphCount, font, textMatrix, attributes, CGContext)
}

func (l_ LayoutManager) StrikethroughGlyphRange_StrikethroughType_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, strikethroughVal UnderlineStyle, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("strikethroughGlyphRange:strikethroughType:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, strikethroughVal, lineRect, lineGlyphRange, containerOrigin)
}

func (l_ LayoutManager) UnderlineGlyphRange_UnderlineType_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, underlineVal UnderlineStyle, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("underlineGlyphRange:underlineType:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, underlineVal, lineRect, lineGlyphRange, containerOrigin)
}

func (l_ LayoutManager) SetLayoutRect_ForTextBlock_GlyphRange(rect foundation.Rect, block ITextBlock, glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setLayoutRect:forTextBlock:glyphRange:"), rect, block, glyphRange)
}

func (l_ LayoutManager) LayoutRectForTextBlock_GlyphRange(block ITextBlock, glyphRange foundation.Range) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("layoutRectForTextBlock:glyphRange:"), block, glyphRange)
	return rv
}

func (l_ LayoutManager) SetBoundsRect_ForTextBlock_GlyphRange(rect foundation.Rect, block ITextBlock, glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setBoundsRect:forTextBlock:glyphRange:"), rect, block, glyphRange)
}

func (l_ LayoutManager) BoundsRectForTextBlock_GlyphRange(block ITextBlock, glyphRange foundation.Range) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("boundsRectForTextBlock:glyphRange:"), block, glyphRange)
	return rv
}

func (l_ LayoutManager) LayoutRectForTextBlock_AtIndex_EffectiveRange(block ITextBlock, glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("layoutRectForTextBlock:atIndex:effectiveRange:"), block, glyphIndex, effectiveGlyphRange)
	return rv
}

func (l_ LayoutManager) BoundsRectForTextBlock_AtIndex_EffectiveRange(block ITextBlock, glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("boundsRectForTextBlock:atIndex:effectiveRange:"), block, glyphIndex, effectiveGlyphRange)
	return rv
}

func (l_ LayoutManager) ShowAttachmentCell_InRect_CharacterIndex(cell ICell, rect foundation.Rect, attachmentIndex uint) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("showAttachmentCell:inRect:characterIndex:"), cell, rect, attachmentIndex)
}

func (l_ LayoutManager) RulerAccessoryViewForTextView_ParagraphStyle_Ruler_Enabled(view ITextView, style IParagraphStyle, ruler IRulerView, isEnabled bool) View {
	rv := objc.CallMethod[View](l_, objc.GetSelector("rulerAccessoryViewForTextView:paragraphStyle:ruler:enabled:"), view, style, ruler, isEnabled)
	return rv
}

func (l_ LayoutManager) RulerMarkersForTextView_ParagraphStyle_Ruler(view ITextView, style IParagraphStyle, ruler IRulerView) []RulerMarker {
	rv := objc.CallMethod[[]RulerMarker](l_, objc.GetSelector("rulerMarkersForTextView:paragraphStyle:ruler:"), view, style, ruler)
	return rv
}

func (l_ LayoutManager) LayoutManagerOwnsFirstResponderInWindow(window IWindow) bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("layoutManagerOwnsFirstResponderInWindow:"), window)
	return rv
}

func (l_ LayoutManager) DefaultLineHeightForFont(theFont IFont) float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("defaultLineHeightForFont:"), theFont)
	return rv
}

func (l_ LayoutManager) DefaultBaselineOffsetForFont(theFont IFont) float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("defaultBaselineOffsetForFont:"), theFont)
	return rv
}

func (l_ LayoutManager) AddTemporaryAttributes_ForCharacterRange(attrs map[foundation.AttributedStringKey]objc.IObject, charRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("addTemporaryAttributes:forCharacterRange:"), attrs, charRange)
}

func (l_ LayoutManager) AddTemporaryAttribute_Value_ForCharacterRange(attrName foundation.AttributedStringKey, value objc.IObject, charRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("addTemporaryAttribute:value:forCharacterRange:"), attrName, value, charRange)
}

func (l_ LayoutManager) SetTemporaryAttributes_ForCharacterRange(attrs map[foundation.AttributedStringKey]objc.IObject, charRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setTemporaryAttributes:forCharacterRange:"), attrs, charRange)
}

func (l_ LayoutManager) RemoveTemporaryAttribute_ForCharacterRange(attrName foundation.AttributedStringKey, charRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("removeTemporaryAttribute:forCharacterRange:"), attrName, charRange)
}

func (l_ LayoutManager) TemporaryAttribute_AtCharacterIndex_EffectiveRange(attrName foundation.AttributedStringKey, location uint, range_ *foundation.Range) objc.Object {
	rv := objc.CallMethod[objc.Object](l_, objc.GetSelector("temporaryAttribute:atCharacterIndex:effectiveRange:"), attrName, location, range_)
	return rv
}

func (l_ LayoutManager) TemporaryAttribute_AtCharacterIndex_LongestEffectiveRange_InRange(attrName foundation.AttributedStringKey, location uint, range_ *foundation.Range, rangeLimit foundation.Range) objc.Object {
	rv := objc.CallMethod[objc.Object](l_, objc.GetSelector("temporaryAttribute:atCharacterIndex:longestEffectiveRange:inRange:"), attrName, location, range_, rangeLimit)
	return rv
}

func (l_ LayoutManager) TemporaryAttributesAtCharacterIndex_EffectiveRange(charIndex uint, effectiveCharRange *foundation.Range) map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](l_, objc.GetSelector("temporaryAttributesAtCharacterIndex:effectiveRange:"), charIndex, effectiveCharRange)
	return rv
}

func (l_ LayoutManager) TemporaryAttributesAtCharacterIndex_LongestEffectiveRange_InRange(location uint, range_ *foundation.Range, rangeLimit foundation.Range) map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](l_, objc.GetSelector("temporaryAttributesAtCharacterIndex:longestEffectiveRange:inRange:"), location, range_, rangeLimit)
	return rv
}

// deprecated
func (l_ LayoutManager) ShowCGGlyphs_Positions_Count_Font_Matrix_Attributes_InContext(glyphs *coregraphics.Glyph, positions *foundation.Point, glyphCount uint, font IFont, textMatrix foundation.IAffineTransform, attributes map[foundation.AttributedStringKey]objc.IObject, graphicsContext IGraphicsContext) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("showCGGlyphs:positions:count:font:matrix:attributes:inContext:"), glyphs, positions, glyphCount, font, textMatrix, attributes, graphicsContext)
}

// deprecated
func (l_ LayoutManager) InvalidateGlyphsOnLayoutInvalidationForGlyphRange(glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("invalidateGlyphsOnLayoutInvalidationForGlyphRange:"), glyphRange)
}

// deprecated
func (l_ LayoutManager) InvalidateLayoutForCharacterRange_IsSoft_ActualCharacterRange(charRange foundation.Range, flag bool, actualCharRange *foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("invalidateLayoutForCharacterRange:isSoft:actualCharacterRange:"), charRange, flag, actualCharRange)
}

// deprecated
func (l_ LayoutManager) TextStorage_Edited_Range_ChangeInLength_InvalidatedRange(str ITextStorage, editedMask TextStorageEditedOptions, newCharRange foundation.Range, delta int, invalidatedCharRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("textStorage:edited:range:changeInLength:invalidatedRange:"), str, editedMask, newCharRange, delta, invalidatedCharRange)
}

// deprecated
func (l_ LayoutManager) InsertGlyph_AtGlyphIndex_CharacterIndex(glyph Glyph, glyphIndex uint, charIndex uint) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("insertGlyph:atGlyphIndex:characterIndex:"), glyph, glyphIndex, charIndex)
}

// deprecated
func (l_ LayoutManager) InsertGlyphs_Length_ForStartingGlyphAtIndex_CharacterIndex(glyphs *Glyph, length uint, glyphIndex uint, charIndex uint) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("insertGlyphs:length:forStartingGlyphAtIndex:characterIndex:"), glyphs, length, glyphIndex, charIndex)
}

// deprecated
func (l_ LayoutManager) GlyphAtIndex(glyphIndex uint) Glyph {
	rv := objc.CallMethod[Glyph](l_, objc.GetSelector("glyphAtIndex:"), glyphIndex)
	return rv
}

// deprecated
func (l_ LayoutManager) GlyphAtIndex_IsValidIndex(glyphIndex uint, isValidIndex *bool) Glyph {
	rv := objc.CallMethod[Glyph](l_, objc.GetSelector("glyphAtIndex:isValidIndex:"), glyphIndex, isValidIndex)
	return rv
}

// deprecated
func (l_ LayoutManager) ReplaceGlyphAtIndex_WithGlyph(glyphIndex uint, newGlyph Glyph) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("replaceGlyphAtIndex:withGlyph:"), glyphIndex, newGlyph)
}

// deprecated
func (l_ LayoutManager) GetGlyphs_Range(glyphArray *Glyph, glyphRange foundation.Range) uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("getGlyphs:range:"), glyphArray, glyphRange)
	return rv
}

// deprecated
func (l_ LayoutManager) GetGlyphsInRange_Glyphs_CharacterIndexes_GlyphInscriptions_ElasticBits(glyphRange foundation.Range, glyphBuffer *Glyph, charIndexBuffer *uint, inscribeBuffer *GlyphInscription, elasticBuffer *bool) uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("getGlyphsInRange:glyphs:characterIndexes:glyphInscriptions:elasticBits:"), glyphRange, glyphBuffer, charIndexBuffer, inscribeBuffer, elasticBuffer)
	return rv
}

// deprecated
func (l_ LayoutManager) GetGlyphsInRange_Glyphs_CharacterIndexes_GlyphInscriptions_ElasticBits_BidiLevels(glyphRange foundation.Range, glyphBuffer *Glyph, charIndexBuffer *uint, inscribeBuffer *GlyphInscription, elasticBuffer *bool, bidiLevelBuffer *byte) uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("getGlyphsInRange:glyphs:characterIndexes:glyphInscriptions:elasticBits:bidiLevels:"), glyphRange, glyphBuffer, charIndexBuffer, inscribeBuffer, elasticBuffer, bidiLevelBuffer)
	return rv
}

// deprecated
func (l_ LayoutManager) DeleteGlyphsInRange(glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("deleteGlyphsInRange:"), glyphRange)
}

// deprecated
func (l_ LayoutManager) SetCharacterIndex_ForGlyphAtIndex(charIndex uint, glyphIndex uint) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setCharacterIndex:forGlyphAtIndex:"), charIndex, glyphIndex)
}

// deprecated
func (l_ LayoutManager) IntAttribute_ForGlyphAtIndex(attributeTag int, glyphIndex uint) int {
	rv := objc.CallMethod[int](l_, objc.GetSelector("intAttribute:forGlyphAtIndex:"), attributeTag, glyphIndex)
	return rv
}

// deprecated
func (l_ LayoutManager) SetIntAttribute_Value_ForGlyphAtIndex(attributeTag int, val int, glyphIndex uint) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setIntAttribute:value:forGlyphAtIndex:"), attributeTag, val, glyphIndex)
}

// deprecated
func (l_ LayoutManager) SetLocations_StartingGlyphIndexes_Count_ForGlyphRange(locations *foundation.Point, glyphIndexes *uint, count uint, glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setLocations:startingGlyphIndexes:count:forGlyphRange:"), locations, glyphIndexes, count, glyphRange)
}

// deprecated
func (l_ LayoutManager) RectArrayForCharacterRange_WithinSelectedCharacterRange_InTextContainer_RectCount(charRange foundation.Range, selCharRange foundation.Range, container ITextContainer, rectCount *uint) *foundation.Rect {
	rv := objc.CallMethod[*foundation.Rect](l_, objc.GetSelector("rectArrayForCharacterRange:withinSelectedCharacterRange:inTextContainer:rectCount:"), charRange, selCharRange, container, rectCount)
	return rv
}

// deprecated
func (l_ LayoutManager) RectArrayForGlyphRange_WithinSelectedGlyphRange_InTextContainer_RectCount(glyphRange foundation.Range, selGlyphRange foundation.Range, container ITextContainer, rectCount *uint) *foundation.Rect {
	rv := objc.CallMethod[*foundation.Rect](l_, objc.GetSelector("rectArrayForGlyphRange:withinSelectedGlyphRange:inTextContainer:rectCount:"), glyphRange, selGlyphRange, container, rectCount)
	return rv
}

// deprecated
func (l_ LayoutManager) SubstituteFontForFont(originalFont IFont) Font {
	rv := objc.CallMethod[Font](l_, objc.GetSelector("substituteFontForFont:"), originalFont)
	return rv
}

// deprecated
func (l_ LayoutManager) ShowPackedGlyphs_Length_GlyphRange_AtPoint_Font_Color_PrintingAdjustment(glyphs *byte, glyphLen uint, glyphRange foundation.Range, point foundation.Point, font IFont, color IColor, printingAdjustment foundation.Size) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("showPackedGlyphs:length:glyphRange:atPoint:font:color:printingAdjustment:"), glyphs, glyphLen, glyphRange, point, font, color, printingAdjustment)
}

func (l_ LayoutManager) Delegate() LayoutManagerDelegateWrapper {
	rv := objc.CallMethod[LayoutManagerDelegateWrapper](l_, objc.GetSelector("delegate"))
	return rv
}

func (l_ LayoutManager) SetDelegate(value LayoutManagerDelegate) {
	po := objc.CreateProtocol("NSLayoutManagerDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(l_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setDelegate:"), po)
}

func (l_ LayoutManager) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setDelegate:"), value)
}

func (l_ LayoutManager) TextStorage() TextStorage {
	rv := objc.CallMethod[TextStorage](l_, objc.GetSelector("textStorage"))
	return rv
}

func (l_ LayoutManager) SetTextStorage(value ITextStorage) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setTextStorage:"), value)
}

func (l_ LayoutManager) AllowsNonContiguousLayout() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("allowsNonContiguousLayout"))
	return rv
}

func (l_ LayoutManager) SetAllowsNonContiguousLayout(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setAllowsNonContiguousLayout:"), value)
}

func (l_ LayoutManager) HasNonContiguousLayout() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("hasNonContiguousLayout"))
	return rv
}

func (l_ LayoutManager) ShowsInvisibleCharacters() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("showsInvisibleCharacters"))
	return rv
}

func (l_ LayoutManager) SetShowsInvisibleCharacters(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setShowsInvisibleCharacters:"), value)
}

func (l_ LayoutManager) ShowsControlCharacters() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("showsControlCharacters"))
	return rv
}

func (l_ LayoutManager) SetShowsControlCharacters(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setShowsControlCharacters:"), value)
}

func (l_ LayoutManager) UsesFontLeading() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("usesFontLeading"))
	return rv
}

func (l_ LayoutManager) SetUsesFontLeading(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setUsesFontLeading:"), value)
}

func (l_ LayoutManager) BackgroundLayoutEnabled() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("backgroundLayoutEnabled"))
	return rv
}

func (l_ LayoutManager) SetBackgroundLayoutEnabled(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setBackgroundLayoutEnabled:"), value)
}

func (l_ LayoutManager) LimitsLayoutForSuspiciousContents() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("limitsLayoutForSuspiciousContents"))
	return rv
}

func (l_ LayoutManager) SetLimitsLayoutForSuspiciousContents(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setLimitsLayoutForSuspiciousContents:"), value)
}

func (l_ LayoutManager) UsesDefaultHyphenation() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("usesDefaultHyphenation"))
	return rv
}

func (l_ LayoutManager) SetUsesDefaultHyphenation(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setUsesDefaultHyphenation:"), value)
}

func (l_ LayoutManager) TextContainers() []TextContainer {
	rv := objc.CallMethod[[]TextContainer](l_, objc.GetSelector("textContainers"))
	return rv
}

func (l_ LayoutManager) GlyphGenerator() GlyphGenerator {
	rv := objc.CallMethod[GlyphGenerator](l_, objc.GetSelector("glyphGenerator"))
	return rv
}

func (l_ LayoutManager) SetGlyphGenerator(value IGlyphGenerator) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setGlyphGenerator:"), value)
}

func (l_ LayoutManager) NumberOfGlyphs() uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("numberOfGlyphs"))
	return rv
}

func (l_ LayoutManager) ExtraLineFragmentRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("extraLineFragmentRect"))
	return rv
}

func (l_ LayoutManager) ExtraLineFragmentTextContainer() TextContainer {
	rv := objc.CallMethod[TextContainer](l_, objc.GetSelector("extraLineFragmentTextContainer"))
	return rv
}

func (l_ LayoutManager) ExtraLineFragmentUsedRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("extraLineFragmentUsedRect"))
	return rv
}

func (l_ LayoutManager) DefaultAttachmentScaling() ImageScaling {
	rv := objc.CallMethod[ImageScaling](l_, objc.GetSelector("defaultAttachmentScaling"))
	return rv
}

func (l_ LayoutManager) SetDefaultAttachmentScaling(value ImageScaling) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setDefaultAttachmentScaling:"), value)
}

func (l_ LayoutManager) FirstTextView() TextView {
	rv := objc.CallMethod[TextView](l_, objc.GetSelector("firstTextView"))
	return rv
}

func (l_ LayoutManager) TextViewForBeginningOfSelection() TextView {
	rv := objc.CallMethod[TextView](l_, objc.GetSelector("textViewForBeginningOfSelection"))
	return rv
}

func (l_ LayoutManager) Typesetter() Typesetter {
	rv := objc.CallMethod[Typesetter](l_, objc.GetSelector("typesetter"))
	return rv
}

func (l_ LayoutManager) SetTypesetter(value ITypesetter) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setTypesetter:"), value)
}

func (l_ LayoutManager) TypesetterBehavior() TypesetterBehavior {
	rv := objc.CallMethod[TypesetterBehavior](l_, objc.GetSelector("typesetterBehavior"))
	return rv
}

func (l_ LayoutManager) SetTypesetterBehavior(value TypesetterBehavior) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setTypesetterBehavior:"), value)
}

// deprecated
func (l_ LayoutManager) HyphenationFactor() float32 {
	rv := objc.CallMethod[float32](l_, objc.GetSelector("hyphenationFactor"))
	return rv
}

// deprecated
func (l_ LayoutManager) SetHyphenationFactor(value float32) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setHyphenationFactor:"), value)
}

// deprecated
func (l_ LayoutManager) UsesScreenFonts() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("usesScreenFonts"))
	return rv
}

// deprecated
func (l_ LayoutManager) SetUsesScreenFonts(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setUsesScreenFonts:"), value)
}
