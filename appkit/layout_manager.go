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
	rv := ffi.CallMethod[LayoutManager](l_, "init")
	return rv
}

func (lc _LayoutManagerClass) Alloc() LayoutManager {
	rv := ffi.CallMethod[LayoutManager](lc, "alloc")
	return rv
}

func (lc _LayoutManagerClass) New() LayoutManager {
	rv := ffi.CallMethod[LayoutManager](lc, "new")
	rv.Autorelease()
	return rv
}

func NewLayoutManager() LayoutManager {
	return LayoutManagerClass.New()
}

func (l_ LayoutManager) ReplaceTextStorage(newTextStorage ITextStorage) {
	ffi.CallMethod[ffi.Void](l_, "replaceTextStorage:", newTextStorage)
}

func (l_ LayoutManager) AddTextContainer(container ITextContainer) {
	ffi.CallMethod[ffi.Void](l_, "addTextContainer:", container)
}

func (l_ LayoutManager) InsertTextContainer_AtIndex(container ITextContainer, index uint) {
	ffi.CallMethod[ffi.Void](l_, "insertTextContainer:atIndex:", container, index)
}

func (l_ LayoutManager) RemoveTextContainerAtIndex(index uint) {
	ffi.CallMethod[ffi.Void](l_, "removeTextContainerAtIndex:", index)
}

func (l_ LayoutManager) SetTextContainer_ForGlyphRange(container ITextContainer, glyphRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "setTextContainer:forGlyphRange:", container, glyphRange)
}

func (l_ LayoutManager) TextContainerChangedGeometry(container ITextContainer) {
	ffi.CallMethod[ffi.Void](l_, "textContainerChangedGeometry:", container)
}

func (l_ LayoutManager) TextContainerChangedTextView(container ITextContainer) {
	ffi.CallMethod[ffi.Void](l_, "textContainerChangedTextView:", container)
}

func (l_ LayoutManager) TextContainerForGlyphAtIndex_EffectiveRange(glyphIndex uint, effectiveGlyphRange *foundation.Range) TextContainer {
	rv := ffi.CallMethod[TextContainer](l_, "textContainerForGlyphAtIndex:effectiveRange:", glyphIndex, effectiveGlyphRange)
	return rv
}

func (l_ LayoutManager) TextContainerForGlyphAtIndex_EffectiveRange_WithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange *foundation.Range, flag bool) TextContainer {
	rv := ffi.CallMethod[TextContainer](l_, "textContainerForGlyphAtIndex:effectiveRange:withoutAdditionalLayout:", glyphIndex, effectiveGlyphRange, flag)
	return rv
}

func (l_ LayoutManager) UsedRectForTextContainer(container ITextContainer) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](l_, "usedRectForTextContainer:", container)
	return rv
}

func (l_ LayoutManager) InvalidateDisplayForCharacterRange(charRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "invalidateDisplayForCharacterRange:", charRange)
}

func (l_ LayoutManager) InvalidateDisplayForGlyphRange(glyphRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "invalidateDisplayForGlyphRange:", glyphRange)
}

func (l_ LayoutManager) InvalidateGlyphsForCharacterRange_ChangeInLength_ActualCharacterRange(charRange foundation.Range, delta int, actualCharRange *foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "invalidateGlyphsForCharacterRange:changeInLength:actualCharacterRange:", charRange, delta, actualCharRange)
}

func (l_ LayoutManager) InvalidateLayoutForCharacterRange_ActualCharacterRange(charRange foundation.Range, actualCharRange *foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "invalidateLayoutForCharacterRange:actualCharacterRange:", charRange, actualCharRange)
}

func (l_ LayoutManager) ProcessEditingForTextStorage_Edited_Range_ChangeInLength_InvalidatedRange(textStorage ITextStorage, editMask TextStorageEditActions, newCharRange foundation.Range, delta int, invalidatedCharRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "processEditingForTextStorage:edited:range:changeInLength:invalidatedRange:", textStorage, editMask, newCharRange, delta, invalidatedCharRange)
}

func (l_ LayoutManager) EnsureGlyphsForCharacterRange(charRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "ensureGlyphsForCharacterRange:", charRange)
}

func (l_ LayoutManager) EnsureGlyphsForGlyphRange(glyphRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "ensureGlyphsForGlyphRange:", glyphRange)
}

func (l_ LayoutManager) EnsureLayoutForBoundingRect_InTextContainer(bounds foundation.Rect, container ITextContainer) {
	ffi.CallMethod[ffi.Void](l_, "ensureLayoutForBoundingRect:inTextContainer:", bounds, container)
}

func (l_ LayoutManager) EnsureLayoutForCharacterRange(charRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "ensureLayoutForCharacterRange:", charRange)
}

func (l_ LayoutManager) EnsureLayoutForGlyphRange(glyphRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "ensureLayoutForGlyphRange:", glyphRange)
}

func (l_ LayoutManager) EnsureLayoutForTextContainer(container ITextContainer) {
	ffi.CallMethod[ffi.Void](l_, "ensureLayoutForTextContainer:", container)
}

func (l_ LayoutManager) GetGlyphsInRange_Glyphs_Properties_CharacterIndexes_BidiLevels(glyphRange foundation.Range, glyphBuffer *coregraphics.Glyph, props *GlyphProperty, charIndexBuffer *uint, bidiLevelBuffer *byte) uint {
	rv := ffi.CallMethod[uint](l_, "getGlyphsInRange:glyphs:properties:characterIndexes:bidiLevels:", glyphRange, glyphBuffer, props, charIndexBuffer, bidiLevelBuffer)
	return rv
}

func (l_ LayoutManager) CGGlyphAtIndex(glyphIndex uint) coregraphics.Glyph {
	rv := ffi.CallMethod[coregraphics.Glyph](l_, "CGGlyphAtIndex:", glyphIndex)
	return rv
}

func (l_ LayoutManager) CGGlyphAtIndex_IsValidIndex(glyphIndex uint, isValidIndex *bool) coregraphics.Glyph {
	rv := ffi.CallMethod[coregraphics.Glyph](l_, "CGGlyphAtIndex:isValidIndex:", glyphIndex, isValidIndex)
	return rv
}

func (l_ LayoutManager) SetGlyphs_Properties_CharacterIndexes_Font_ForGlyphRange(glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont IFont, glyphRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "setGlyphs:properties:characterIndexes:font:forGlyphRange:", glyphs, props, charIndexes, aFont, glyphRange)
}

func (l_ LayoutManager) CharacterIndexForGlyphAtIndex(glyphIndex uint) uint {
	rv := ffi.CallMethod[uint](l_, "characterIndexForGlyphAtIndex:", glyphIndex)
	return rv
}

func (l_ LayoutManager) GlyphIndexForCharacterAtIndex(charIndex uint) uint {
	rv := ffi.CallMethod[uint](l_, "glyphIndexForCharacterAtIndex:", charIndex)
	return rv
}

func (l_ LayoutManager) IsValidGlyphIndex(glyphIndex uint) bool {
	rv := ffi.CallMethod[bool](l_, "isValidGlyphIndex:", glyphIndex)
	return rv
}

func (l_ LayoutManager) PropertyForGlyphAtIndex(glyphIndex uint) GlyphProperty {
	rv := ffi.CallMethod[GlyphProperty](l_, "propertyForGlyphAtIndex:", glyphIndex)
	return rv
}

func (l_ LayoutManager) SetAttachmentSize_ForGlyphRange(attachmentSize foundation.Size, glyphRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "setAttachmentSize:forGlyphRange:", attachmentSize, glyphRange)
}

func (l_ LayoutManager) SetDrawsOutsideLineFragment_ForGlyphAtIndex(flag bool, glyphIndex uint) {
	ffi.CallMethod[ffi.Void](l_, "setDrawsOutsideLineFragment:forGlyphAtIndex:", flag, glyphIndex)
}

func (l_ LayoutManager) SetExtraLineFragmentRect_UsedRect_TextContainer(fragmentRect foundation.Rect, usedRect foundation.Rect, container ITextContainer) {
	ffi.CallMethod[ffi.Void](l_, "setExtraLineFragmentRect:usedRect:textContainer:", fragmentRect, usedRect, container)
}

func (l_ LayoutManager) SetLineFragmentRect_ForGlyphRange_UsedRect(fragmentRect foundation.Rect, glyphRange foundation.Range, usedRect foundation.Rect) {
	ffi.CallMethod[ffi.Void](l_, "setLineFragmentRect:forGlyphRange:usedRect:", fragmentRect, glyphRange, usedRect)
}

func (l_ LayoutManager) SetLocation_ForStartOfGlyphRange(location foundation.Point, glyphRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "setLocation:forStartOfGlyphRange:", location, glyphRange)
}

func (l_ LayoutManager) SetNotShownAttribute_ForGlyphAtIndex(flag bool, glyphIndex uint) {
	ffi.CallMethod[ffi.Void](l_, "setNotShownAttribute:forGlyphAtIndex:", flag, glyphIndex)
}

func (l_ LayoutManager) AttachmentSizeForGlyphAtIndex(glyphIndex uint) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](l_, "attachmentSizeForGlyphAtIndex:", glyphIndex)
	return rv
}

func (l_ LayoutManager) DrawsOutsideLineFragmentForGlyphAtIndex(glyphIndex uint) bool {
	rv := ffi.CallMethod[bool](l_, "drawsOutsideLineFragmentForGlyphAtIndex:", glyphIndex)
	return rv
}

func (l_ LayoutManager) FirstUnlaidCharacterIndex() uint {
	rv := ffi.CallMethod[uint](l_, "firstUnlaidCharacterIndex")
	return rv
}

func (l_ LayoutManager) FirstUnlaidGlyphIndex() uint {
	rv := ffi.CallMethod[uint](l_, "firstUnlaidGlyphIndex")
	return rv
}

func (l_ LayoutManager) GetFirstUnlaidCharacterIndex_GlyphIndex(charIndex *uint, glyphIndex *uint) {
	ffi.CallMethod[ffi.Void](l_, "getFirstUnlaidCharacterIndex:glyphIndex:", charIndex, glyphIndex)
}

func (l_ LayoutManager) LineFragmentRectForGlyphAtIndex_EffectiveRange(glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](l_, "lineFragmentRectForGlyphAtIndex:effectiveRange:", glyphIndex, effectiveGlyphRange)
	return rv
}

func (l_ LayoutManager) LineFragmentRectForGlyphAtIndex_EffectiveRange_WithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange *foundation.Range, flag bool) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](l_, "lineFragmentRectForGlyphAtIndex:effectiveRange:withoutAdditionalLayout:", glyphIndex, effectiveGlyphRange, flag)
	return rv
}

func (l_ LayoutManager) LineFragmentUsedRectForGlyphAtIndex_EffectiveRange(glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](l_, "lineFragmentUsedRectForGlyphAtIndex:effectiveRange:", glyphIndex, effectiveGlyphRange)
	return rv
}

func (l_ LayoutManager) LineFragmentUsedRectForGlyphAtIndex_EffectiveRange_WithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange *foundation.Range, flag bool) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](l_, "lineFragmentUsedRectForGlyphAtIndex:effectiveRange:withoutAdditionalLayout:", glyphIndex, effectiveGlyphRange, flag)
	return rv
}

func (l_ LayoutManager) LocationForGlyphAtIndex(glyphIndex uint) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](l_, "locationForGlyphAtIndex:", glyphIndex)
	return rv
}

func (l_ LayoutManager) NotShownAttributeForGlyphAtIndex(glyphIndex uint) bool {
	rv := ffi.CallMethod[bool](l_, "notShownAttributeForGlyphAtIndex:", glyphIndex)
	return rv
}

func (l_ LayoutManager) TruncatedGlyphRangeInLineFragmentForGlyphAtIndex(glyphIndex uint) foundation.Range {
	rv := ffi.CallMethod[foundation.Range](l_, "truncatedGlyphRangeInLineFragmentForGlyphAtIndex:", glyphIndex)
	return rv
}

func (l_ LayoutManager) BoundingRectForGlyphRange_InTextContainer(glyphRange foundation.Range, container ITextContainer) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](l_, "boundingRectForGlyphRange:inTextContainer:", glyphRange, container)
	return rv
}

func (l_ LayoutManager) CharacterIndexForPoint_InTextContainer_FractionOfDistanceBetweenInsertionPoints(point foundation.Point, container ITextContainer, partialFraction *float64) uint {
	rv := ffi.CallMethod[uint](l_, "characterIndexForPoint:inTextContainer:fractionOfDistanceBetweenInsertionPoints:", point, container, partialFraction)
	return rv
}

func (l_ LayoutManager) CharacterRangeForGlyphRange_ActualGlyphRange(glyphRange foundation.Range, actualGlyphRange *foundation.Range) foundation.Range {
	rv := ffi.CallMethod[foundation.Range](l_, "characterRangeForGlyphRange:actualGlyphRange:", glyphRange, actualGlyphRange)
	return rv
}

func (l_ LayoutManager) EnumerateEnclosingRectsForGlyphRange_WithinSelectedGlyphRange_InTextContainer_UsingBlock(glyphRange foundation.Range, selectedRange foundation.Range, textContainer ITextContainer, block func(rect foundation.Rect, stop *bool)) {
	ffi.CallMethod[ffi.Void](l_, "enumerateEnclosingRectsForGlyphRange:withinSelectedGlyphRange:inTextContainer:usingBlock:", glyphRange, selectedRange, textContainer, block)
}

func (l_ LayoutManager) EnumerateLineFragmentsForGlyphRange_UsingBlock(glyphRange foundation.Range, block func(rect foundation.Rect, usedRect foundation.Rect, textContainer TextContainer, glyphRange foundation.Range, stop *bool)) {
	ffi.CallMethod[ffi.Void](l_, "enumerateLineFragmentsForGlyphRange:usingBlock:", glyphRange, block)
}

func (l_ LayoutManager) FractionOfDistanceThroughGlyphForPoint_InTextContainer(point foundation.Point, container ITextContainer) float64 {
	rv := ffi.CallMethod[float64](l_, "fractionOfDistanceThroughGlyphForPoint:inTextContainer:", point, container)
	return rv
}

func (l_ LayoutManager) GetLineFragmentInsertionPointsForCharacterAtIndex_AlternatePositions_InDisplayOrder_Positions_CharacterIndexes(charIndex uint, aFlag bool, dFlag bool, positions *float64, charIndexes *uint) uint {
	rv := ffi.CallMethod[uint](l_, "getLineFragmentInsertionPointsForCharacterAtIndex:alternatePositions:inDisplayOrder:positions:characterIndexes:", charIndex, aFlag, dFlag, positions, charIndexes)
	return rv
}

func (l_ LayoutManager) GlyphIndexForPoint_InTextContainer(point foundation.Point, container ITextContainer) uint {
	rv := ffi.CallMethod[uint](l_, "glyphIndexForPoint:inTextContainer:", point, container)
	return rv
}

func (l_ LayoutManager) GlyphIndexForPoint_InTextContainer_FractionOfDistanceThroughGlyph(point foundation.Point, container ITextContainer, partialFraction *float64) uint {
	rv := ffi.CallMethod[uint](l_, "glyphIndexForPoint:inTextContainer:fractionOfDistanceThroughGlyph:", point, container, partialFraction)
	return rv
}

func (l_ LayoutManager) GlyphRangeForBoundingRect_InTextContainer(bounds foundation.Rect, container ITextContainer) foundation.Range {
	rv := ffi.CallMethod[foundation.Range](l_, "glyphRangeForBoundingRect:inTextContainer:", bounds, container)
	return rv
}

func (l_ LayoutManager) GlyphRangeForBoundingRectWithoutAdditionalLayout_InTextContainer(bounds foundation.Rect, container ITextContainer) foundation.Range {
	rv := ffi.CallMethod[foundation.Range](l_, "glyphRangeForBoundingRectWithoutAdditionalLayout:inTextContainer:", bounds, container)
	return rv
}

func (l_ LayoutManager) GlyphRangeForTextContainer(container ITextContainer) foundation.Range {
	rv := ffi.CallMethod[foundation.Range](l_, "glyphRangeForTextContainer:", container)
	return rv
}

func (l_ LayoutManager) GlyphRangeForCharacterRange_ActualCharacterRange(charRange foundation.Range, actualCharRange *foundation.Range) foundation.Range {
	rv := ffi.CallMethod[foundation.Range](l_, "glyphRangeForCharacterRange:actualCharacterRange:", charRange, actualCharRange)
	return rv
}

func (l_ LayoutManager) RangeOfNominallySpacedGlyphsContainingIndex(glyphIndex uint) foundation.Range {
	rv := ffi.CallMethod[foundation.Range](l_, "rangeOfNominallySpacedGlyphsContainingIndex:", glyphIndex)
	return rv
}

func (l_ LayoutManager) DrawBackgroundForGlyphRange_AtPoint(glyphsToShow foundation.Range, origin foundation.Point) {
	ffi.CallMethod[ffi.Void](l_, "drawBackgroundForGlyphRange:atPoint:", glyphsToShow, origin)
}

func (l_ LayoutManager) DrawGlyphsForGlyphRange_AtPoint(glyphsToShow foundation.Range, origin foundation.Point) {
	ffi.CallMethod[ffi.Void](l_, "drawGlyphsForGlyphRange:atPoint:", glyphsToShow, origin)
}

func (l_ LayoutManager) DrawStrikethroughForGlyphRange_StrikethroughType_BaselineOffset_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, strikethroughVal UnderlineStyle, baselineOffset float64, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point) {
	ffi.CallMethod[ffi.Void](l_, "drawStrikethroughForGlyphRange:strikethroughType:baselineOffset:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:", glyphRange, strikethroughVal, baselineOffset, lineRect, lineGlyphRange, containerOrigin)
}

func (l_ LayoutManager) DrawUnderlineForGlyphRange_UnderlineType_BaselineOffset_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, underlineVal UnderlineStyle, baselineOffset float64, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point) {
	ffi.CallMethod[ffi.Void](l_, "drawUnderlineForGlyphRange:underlineType:baselineOffset:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:", glyphRange, underlineVal, baselineOffset, lineRect, lineGlyphRange, containerOrigin)
}

func (l_ LayoutManager) FillBackgroundRectArray_Count_ForCharacterRange_Color(rectArray *foundation.Rect, rectCount uint, charRange foundation.Range, color IColor) {
	ffi.CallMethod[ffi.Void](l_, "fillBackgroundRectArray:count:forCharacterRange:color:", rectArray, rectCount, charRange, color)
}

func (l_ LayoutManager) ShowCGGlyphs_Positions_Count_Font_TextMatrix_Attributes_InContext(glyphs *coregraphics.Glyph, positions *coregraphics.Point, glyphCount int, font IFont, textMatrix coregraphics.AffineTransform, attributes map[foundation.AttributedStringKey]objc.IObject, CGContext coregraphics.ContextRef) {
	ffi.CallMethod[ffi.Void](l_, "showCGGlyphs:positions:count:font:textMatrix:attributes:inContext:", glyphs, positions, glyphCount, font, textMatrix, attributes, CGContext)
}

func (l_ LayoutManager) StrikethroughGlyphRange_StrikethroughType_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, strikethroughVal UnderlineStyle, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point) {
	ffi.CallMethod[ffi.Void](l_, "strikethroughGlyphRange:strikethroughType:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:", glyphRange, strikethroughVal, lineRect, lineGlyphRange, containerOrigin)
}

func (l_ LayoutManager) UnderlineGlyphRange_UnderlineType_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, underlineVal UnderlineStyle, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point) {
	ffi.CallMethod[ffi.Void](l_, "underlineGlyphRange:underlineType:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:", glyphRange, underlineVal, lineRect, lineGlyphRange, containerOrigin)
}

func (l_ LayoutManager) SetLayoutRect_ForTextBlock_GlyphRange(rect foundation.Rect, block ITextBlock, glyphRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "setLayoutRect:forTextBlock:glyphRange:", rect, block, glyphRange)
}

func (l_ LayoutManager) LayoutRectForTextBlock_GlyphRange(block ITextBlock, glyphRange foundation.Range) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](l_, "layoutRectForTextBlock:glyphRange:", block, glyphRange)
	return rv
}

func (l_ LayoutManager) SetBoundsRect_ForTextBlock_GlyphRange(rect foundation.Rect, block ITextBlock, glyphRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "setBoundsRect:forTextBlock:glyphRange:", rect, block, glyphRange)
}

func (l_ LayoutManager) BoundsRectForTextBlock_GlyphRange(block ITextBlock, glyphRange foundation.Range) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](l_, "boundsRectForTextBlock:glyphRange:", block, glyphRange)
	return rv
}

func (l_ LayoutManager) LayoutRectForTextBlock_AtIndex_EffectiveRange(block ITextBlock, glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](l_, "layoutRectForTextBlock:atIndex:effectiveRange:", block, glyphIndex, effectiveGlyphRange)
	return rv
}

func (l_ LayoutManager) BoundsRectForTextBlock_AtIndex_EffectiveRange(block ITextBlock, glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](l_, "boundsRectForTextBlock:atIndex:effectiveRange:", block, glyphIndex, effectiveGlyphRange)
	return rv
}

func (l_ LayoutManager) ShowAttachmentCell_InRect_CharacterIndex(cell ICell, rect foundation.Rect, attachmentIndex uint) {
	ffi.CallMethod[ffi.Void](l_, "showAttachmentCell:inRect:characterIndex:", cell, rect, attachmentIndex)
}

func (l_ LayoutManager) RulerAccessoryViewForTextView_ParagraphStyle_Ruler_Enabled(view ITextView, style IParagraphStyle, ruler IRulerView, isEnabled bool) View {
	rv := ffi.CallMethod[View](l_, "rulerAccessoryViewForTextView:paragraphStyle:ruler:enabled:", view, style, ruler, isEnabled)
	return rv
}

func (l_ LayoutManager) RulerMarkersForTextView_ParagraphStyle_Ruler(view ITextView, style IParagraphStyle, ruler IRulerView) []RulerMarker {
	rv := ffi.CallMethod[[]RulerMarker](l_, "rulerMarkersForTextView:paragraphStyle:ruler:", view, style, ruler)
	return rv
}

func (l_ LayoutManager) LayoutManagerOwnsFirstResponderInWindow(window IWindow) bool {
	rv := ffi.CallMethod[bool](l_, "layoutManagerOwnsFirstResponderInWindow:", window)
	return rv
}

func (l_ LayoutManager) DefaultLineHeightForFont(theFont IFont) float64 {
	rv := ffi.CallMethod[float64](l_, "defaultLineHeightForFont:", theFont)
	return rv
}

func (l_ LayoutManager) DefaultBaselineOffsetForFont(theFont IFont) float64 {
	rv := ffi.CallMethod[float64](l_, "defaultBaselineOffsetForFont:", theFont)
	return rv
}

func (l_ LayoutManager) AddTemporaryAttributes_ForCharacterRange(attrs map[foundation.AttributedStringKey]objc.IObject, charRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "addTemporaryAttributes:forCharacterRange:", attrs, charRange)
}

func (l_ LayoutManager) AddTemporaryAttribute_Value_ForCharacterRange(attrName foundation.AttributedStringKey, value objc.IObject, charRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "addTemporaryAttribute:value:forCharacterRange:", attrName, value, charRange)
}

func (l_ LayoutManager) SetTemporaryAttributes_ForCharacterRange(attrs map[foundation.AttributedStringKey]objc.IObject, charRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "setTemporaryAttributes:forCharacterRange:", attrs, charRange)
}

func (l_ LayoutManager) RemoveTemporaryAttribute_ForCharacterRange(attrName foundation.AttributedStringKey, charRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "removeTemporaryAttribute:forCharacterRange:", attrName, charRange)
}

func (l_ LayoutManager) TemporaryAttribute_AtCharacterIndex_EffectiveRange(attrName foundation.AttributedStringKey, location uint, range_ *foundation.Range) objc.Object {
	rv := ffi.CallMethod[objc.Object](l_, "temporaryAttribute:atCharacterIndex:effectiveRange:", attrName, location, range_)
	return rv
}

func (l_ LayoutManager) TemporaryAttribute_AtCharacterIndex_LongestEffectiveRange_InRange(attrName foundation.AttributedStringKey, location uint, range_ *foundation.Range, rangeLimit foundation.Range) objc.Object {
	rv := ffi.CallMethod[objc.Object](l_, "temporaryAttribute:atCharacterIndex:longestEffectiveRange:inRange:", attrName, location, range_, rangeLimit)
	return rv
}

func (l_ LayoutManager) TemporaryAttributesAtCharacterIndex_EffectiveRange(charIndex uint, effectiveCharRange *foundation.Range) map[foundation.AttributedStringKey]objc.Object {
	rv := ffi.CallMethod[map[foundation.AttributedStringKey]objc.Object](l_, "temporaryAttributesAtCharacterIndex:effectiveRange:", charIndex, effectiveCharRange)
	return rv
}

func (l_ LayoutManager) TemporaryAttributesAtCharacterIndex_LongestEffectiveRange_InRange(location uint, range_ *foundation.Range, rangeLimit foundation.Range) map[foundation.AttributedStringKey]objc.Object {
	rv := ffi.CallMethod[map[foundation.AttributedStringKey]objc.Object](l_, "temporaryAttributesAtCharacterIndex:longestEffectiveRange:inRange:", location, range_, rangeLimit)
	return rv
}

// deprecated
func (l_ LayoutManager) ShowCGGlyphs_Positions_Count_Font_Matrix_Attributes_InContext(glyphs *coregraphics.Glyph, positions *foundation.Point, glyphCount uint, font IFont, textMatrix foundation.IAffineTransform, attributes map[foundation.AttributedStringKey]objc.IObject, graphicsContext IGraphicsContext) {
	ffi.CallMethod[ffi.Void](l_, "showCGGlyphs:positions:count:font:matrix:attributes:inContext:", glyphs, positions, glyphCount, font, textMatrix, attributes, graphicsContext)
}

// deprecated
func (l_ LayoutManager) InvalidateGlyphsOnLayoutInvalidationForGlyphRange(glyphRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "invalidateGlyphsOnLayoutInvalidationForGlyphRange:", glyphRange)
}

// deprecated
func (l_ LayoutManager) InvalidateLayoutForCharacterRange_IsSoft_ActualCharacterRange(charRange foundation.Range, flag bool, actualCharRange *foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "invalidateLayoutForCharacterRange:isSoft:actualCharacterRange:", charRange, flag, actualCharRange)
}

// deprecated
func (l_ LayoutManager) TextStorage_Edited_Range_ChangeInLength_InvalidatedRange(str ITextStorage, editedMask TextStorageEditedOptions, newCharRange foundation.Range, delta int, invalidatedCharRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "textStorage:edited:range:changeInLength:invalidatedRange:", str, editedMask, newCharRange, delta, invalidatedCharRange)
}

// deprecated
func (l_ LayoutManager) InsertGlyph_AtGlyphIndex_CharacterIndex(glyph Glyph, glyphIndex uint, charIndex uint) {
	ffi.CallMethod[ffi.Void](l_, "insertGlyph:atGlyphIndex:characterIndex:", glyph, glyphIndex, charIndex)
}

// deprecated
func (l_ LayoutManager) InsertGlyphs_Length_ForStartingGlyphAtIndex_CharacterIndex(glyphs *Glyph, length uint, glyphIndex uint, charIndex uint) {
	ffi.CallMethod[ffi.Void](l_, "insertGlyphs:length:forStartingGlyphAtIndex:characterIndex:", glyphs, length, glyphIndex, charIndex)
}

// deprecated
func (l_ LayoutManager) GlyphAtIndex(glyphIndex uint) Glyph {
	rv := ffi.CallMethod[Glyph](l_, "glyphAtIndex:", glyphIndex)
	return rv
}

// deprecated
func (l_ LayoutManager) GlyphAtIndex_IsValidIndex(glyphIndex uint, isValidIndex *bool) Glyph {
	rv := ffi.CallMethod[Glyph](l_, "glyphAtIndex:isValidIndex:", glyphIndex, isValidIndex)
	return rv
}

// deprecated
func (l_ LayoutManager) ReplaceGlyphAtIndex_WithGlyph(glyphIndex uint, newGlyph Glyph) {
	ffi.CallMethod[ffi.Void](l_, "replaceGlyphAtIndex:withGlyph:", glyphIndex, newGlyph)
}

// deprecated
func (l_ LayoutManager) GetGlyphs_Range(glyphArray *Glyph, glyphRange foundation.Range) uint {
	rv := ffi.CallMethod[uint](l_, "getGlyphs:range:", glyphArray, glyphRange)
	return rv
}

// deprecated
func (l_ LayoutManager) GetGlyphsInRange_Glyphs_CharacterIndexes_GlyphInscriptions_ElasticBits(glyphRange foundation.Range, glyphBuffer *Glyph, charIndexBuffer *uint, inscribeBuffer *GlyphInscription, elasticBuffer *bool) uint {
	rv := ffi.CallMethod[uint](l_, "getGlyphsInRange:glyphs:characterIndexes:glyphInscriptions:elasticBits:", glyphRange, glyphBuffer, charIndexBuffer, inscribeBuffer, elasticBuffer)
	return rv
}

// deprecated
func (l_ LayoutManager) GetGlyphsInRange_Glyphs_CharacterIndexes_GlyphInscriptions_ElasticBits_BidiLevels(glyphRange foundation.Range, glyphBuffer *Glyph, charIndexBuffer *uint, inscribeBuffer *GlyphInscription, elasticBuffer *bool, bidiLevelBuffer *byte) uint {
	rv := ffi.CallMethod[uint](l_, "getGlyphsInRange:glyphs:characterIndexes:glyphInscriptions:elasticBits:bidiLevels:", glyphRange, glyphBuffer, charIndexBuffer, inscribeBuffer, elasticBuffer, bidiLevelBuffer)
	return rv
}

// deprecated
func (l_ LayoutManager) DeleteGlyphsInRange(glyphRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "deleteGlyphsInRange:", glyphRange)
}

// deprecated
func (l_ LayoutManager) SetCharacterIndex_ForGlyphAtIndex(charIndex uint, glyphIndex uint) {
	ffi.CallMethod[ffi.Void](l_, "setCharacterIndex:forGlyphAtIndex:", charIndex, glyphIndex)
}

// deprecated
func (l_ LayoutManager) IntAttribute_ForGlyphAtIndex(attributeTag int, glyphIndex uint) int {
	rv := ffi.CallMethod[int](l_, "intAttribute:forGlyphAtIndex:", attributeTag, glyphIndex)
	return rv
}

// deprecated
func (l_ LayoutManager) SetIntAttribute_Value_ForGlyphAtIndex(attributeTag int, val int, glyphIndex uint) {
	ffi.CallMethod[ffi.Void](l_, "setIntAttribute:value:forGlyphAtIndex:", attributeTag, val, glyphIndex)
}

// deprecated
func (l_ LayoutManager) SetLocations_StartingGlyphIndexes_Count_ForGlyphRange(locations *foundation.Point, glyphIndexes *uint, count uint, glyphRange foundation.Range) {
	ffi.CallMethod[ffi.Void](l_, "setLocations:startingGlyphIndexes:count:forGlyphRange:", locations, glyphIndexes, count, glyphRange)
}

// deprecated
func (l_ LayoutManager) RectArrayForCharacterRange_WithinSelectedCharacterRange_InTextContainer_RectCount(charRange foundation.Range, selCharRange foundation.Range, container ITextContainer, rectCount *uint) *foundation.Rect {
	rv := ffi.CallMethod[*foundation.Rect](l_, "rectArrayForCharacterRange:withinSelectedCharacterRange:inTextContainer:rectCount:", charRange, selCharRange, container, rectCount)
	return rv
}

// deprecated
func (l_ LayoutManager) RectArrayForGlyphRange_WithinSelectedGlyphRange_InTextContainer_RectCount(glyphRange foundation.Range, selGlyphRange foundation.Range, container ITextContainer, rectCount *uint) *foundation.Rect {
	rv := ffi.CallMethod[*foundation.Rect](l_, "rectArrayForGlyphRange:withinSelectedGlyphRange:inTextContainer:rectCount:", glyphRange, selGlyphRange, container, rectCount)
	return rv
}

// deprecated
func (l_ LayoutManager) SubstituteFontForFont(originalFont IFont) Font {
	rv := ffi.CallMethod[Font](l_, "substituteFontForFont:", originalFont)
	return rv
}

// deprecated
func (l_ LayoutManager) ShowPackedGlyphs_Length_GlyphRange_AtPoint_Font_Color_PrintingAdjustment(glyphs *byte, glyphLen uint, glyphRange foundation.Range, point foundation.Point, font IFont, color IColor, printingAdjustment foundation.Size) {
	ffi.CallMethod[ffi.Void](l_, "showPackedGlyphs:length:glyphRange:atPoint:font:color:printingAdjustment:", glyphs, glyphLen, glyphRange, point, font, color, printingAdjustment)
}

func (l_ LayoutManager) Delegate() LayoutManagerDelegateWrapper {
	rv := ffi.CallMethod[LayoutManagerDelegateWrapper](l_, "delegate")
	return rv
}

func (l_ LayoutManager) SetDelegate(value LayoutManagerDelegate) {
	po := ffi.CreateProtocol("NSLayoutManagerDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(l_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](l_, "setDelegate:", po)
}

func (l_ LayoutManager) SetDelegate0(value objc.IObject) {
	ffi.CallMethod[ffi.Void](l_, "setDelegate:", value)
}

func (l_ LayoutManager) TextStorage() TextStorage {
	rv := ffi.CallMethod[TextStorage](l_, "textStorage")
	return rv
}

func (l_ LayoutManager) SetTextStorage(value ITextStorage) {
	ffi.CallMethod[ffi.Void](l_, "setTextStorage:", value)
}

func (l_ LayoutManager) AllowsNonContiguousLayout() bool {
	rv := ffi.CallMethod[bool](l_, "allowsNonContiguousLayout")
	return rv
}

func (l_ LayoutManager) SetAllowsNonContiguousLayout(value bool) {
	ffi.CallMethod[ffi.Void](l_, "setAllowsNonContiguousLayout:", value)
}

func (l_ LayoutManager) HasNonContiguousLayout() bool {
	rv := ffi.CallMethod[bool](l_, "hasNonContiguousLayout")
	return rv
}

func (l_ LayoutManager) ShowsInvisibleCharacters() bool {
	rv := ffi.CallMethod[bool](l_, "showsInvisibleCharacters")
	return rv
}

func (l_ LayoutManager) SetShowsInvisibleCharacters(value bool) {
	ffi.CallMethod[ffi.Void](l_, "setShowsInvisibleCharacters:", value)
}

func (l_ LayoutManager) ShowsControlCharacters() bool {
	rv := ffi.CallMethod[bool](l_, "showsControlCharacters")
	return rv
}

func (l_ LayoutManager) SetShowsControlCharacters(value bool) {
	ffi.CallMethod[ffi.Void](l_, "setShowsControlCharacters:", value)
}

func (l_ LayoutManager) UsesFontLeading() bool {
	rv := ffi.CallMethod[bool](l_, "usesFontLeading")
	return rv
}

func (l_ LayoutManager) SetUsesFontLeading(value bool) {
	ffi.CallMethod[ffi.Void](l_, "setUsesFontLeading:", value)
}

func (l_ LayoutManager) BackgroundLayoutEnabled() bool {
	rv := ffi.CallMethod[bool](l_, "backgroundLayoutEnabled")
	return rv
}

func (l_ LayoutManager) SetBackgroundLayoutEnabled(value bool) {
	ffi.CallMethod[ffi.Void](l_, "setBackgroundLayoutEnabled:", value)
}

func (l_ LayoutManager) LimitsLayoutForSuspiciousContents() bool {
	rv := ffi.CallMethod[bool](l_, "limitsLayoutForSuspiciousContents")
	return rv
}

func (l_ LayoutManager) SetLimitsLayoutForSuspiciousContents(value bool) {
	ffi.CallMethod[ffi.Void](l_, "setLimitsLayoutForSuspiciousContents:", value)
}

func (l_ LayoutManager) UsesDefaultHyphenation() bool {
	rv := ffi.CallMethod[bool](l_, "usesDefaultHyphenation")
	return rv
}

func (l_ LayoutManager) SetUsesDefaultHyphenation(value bool) {
	ffi.CallMethod[ffi.Void](l_, "setUsesDefaultHyphenation:", value)
}

func (l_ LayoutManager) TextContainers() []TextContainer {
	rv := ffi.CallMethod[[]TextContainer](l_, "textContainers")
	return rv
}

func (l_ LayoutManager) GlyphGenerator() GlyphGenerator {
	rv := ffi.CallMethod[GlyphGenerator](l_, "glyphGenerator")
	return rv
}

func (l_ LayoutManager) SetGlyphGenerator(value IGlyphGenerator) {
	ffi.CallMethod[ffi.Void](l_, "setGlyphGenerator:", value)
}

func (l_ LayoutManager) NumberOfGlyphs() uint {
	rv := ffi.CallMethod[uint](l_, "numberOfGlyphs")
	return rv
}

func (l_ LayoutManager) ExtraLineFragmentRect() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](l_, "extraLineFragmentRect")
	return rv
}

func (l_ LayoutManager) ExtraLineFragmentTextContainer() TextContainer {
	rv := ffi.CallMethod[TextContainer](l_, "extraLineFragmentTextContainer")
	return rv
}

func (l_ LayoutManager) ExtraLineFragmentUsedRect() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](l_, "extraLineFragmentUsedRect")
	return rv
}

func (l_ LayoutManager) DefaultAttachmentScaling() ImageScaling {
	rv := ffi.CallMethod[ImageScaling](l_, "defaultAttachmentScaling")
	return rv
}

func (l_ LayoutManager) SetDefaultAttachmentScaling(value ImageScaling) {
	ffi.CallMethod[ffi.Void](l_, "setDefaultAttachmentScaling:", value)
}

func (l_ LayoutManager) FirstTextView() TextView {
	rv := ffi.CallMethod[TextView](l_, "firstTextView")
	return rv
}

func (l_ LayoutManager) TextViewForBeginningOfSelection() TextView {
	rv := ffi.CallMethod[TextView](l_, "textViewForBeginningOfSelection")
	return rv
}

func (l_ LayoutManager) Typesetter() Typesetter {
	rv := ffi.CallMethod[Typesetter](l_, "typesetter")
	return rv
}

func (l_ LayoutManager) SetTypesetter(value ITypesetter) {
	ffi.CallMethod[ffi.Void](l_, "setTypesetter:", value)
}

func (l_ LayoutManager) TypesetterBehavior() TypesetterBehavior {
	rv := ffi.CallMethod[TypesetterBehavior](l_, "typesetterBehavior")
	return rv
}

func (l_ LayoutManager) SetTypesetterBehavior(value TypesetterBehavior) {
	ffi.CallMethod[ffi.Void](l_, "setTypesetterBehavior:", value)
}

// deprecated
func (l_ LayoutManager) HyphenationFactor() float32 {
	rv := ffi.CallMethod[float32](l_, "hyphenationFactor")
	return rv
}

// deprecated
func (l_ LayoutManager) SetHyphenationFactor(value float32) {
	ffi.CallMethod[ffi.Void](l_, "setHyphenationFactor:", value)
}

// deprecated
func (l_ LayoutManager) UsesScreenFonts() bool {
	rv := ffi.CallMethod[bool](l_, "usesScreenFonts")
	return rv
}

// deprecated
func (l_ LayoutManager) SetUsesScreenFonts(value bool) {
	ffi.CallMethod[ffi.Void](l_, "setUsesScreenFonts:", value)
}
