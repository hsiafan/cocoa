// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TypesetterClass = _TypesetterClass{objc.GetClass("NSTypesetter")}

type _TypesetterClass struct {
	objc.Class
}

type ITypesetter interface {
	objc.IObject
	BaselineOffsetInLayoutManager_GlyphIndex(layoutMgr ILayoutManager, glyphIndex uint) float64
	SubstituteFontForFont(originalFont IFont) Font
	TextTabForGlyphLocation_WritingDirection_MaxLocation(glyphLocation float64, direction WritingDirection, maxLocation float64) TextTab
	SetParagraphGlyphRange_SeparatorGlyphRange(paragraphRange foundation.Range, paragraphSeparatorRange foundation.Range)
	LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) float64
	ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) float64
	ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) float64
	LayoutParagraphAtPoint(lineFragmentOrigin *foundation.Point) uint
	BeginParagraph()
	EndParagraph()
	BeginLineWithGlyphAtIndex(glyphIndex uint)
	EndLineWithGlyphRange(lineGlyphRange foundation.Range)
	LayoutCharactersInRange_ForLayoutManager_MaximumNumberOfLineFragments(characterRange foundation.Range, layoutManager ILayoutManager, maxNumLines uint) foundation.Range
	LayoutGlyphsInLayoutManager_StartingAtGlyphIndex_MaxNumberOfLineFragments_NextGlyphIndex(layoutManager ILayoutManager, startGlyphIndex uint, maxNumLines uint, nextGlyph *uint)
	BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex(glyphIndex uint, textContainer ITextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect
	GetLineFragmentRect_UsedRect_ForParagraphSeparatorGlyphRange_AtProposedOrigin(lineFragmentRect *foundation.Rect, lineFragmentUsedRect *foundation.Rect, paragraphSeparatorGlyphRange foundation.Range, lineOrigin foundation.Point)
	GetLineFragmentRect_UsedRect_RemainingRect_ForStartingGlyphAtIndex_ProposedRect_LineSpacing_ParagraphSpacingBefore_ParagraphSpacingAfter(lineFragmentRect *foundation.Rect, lineFragmentUsedRect *foundation.Rect, remainingRect *foundation.Rect, startingGlyphIndex uint, proposedRect foundation.Rect, lineSpacing float64, paragraphSpacingBefore float64, paragraphSpacingAfter float64)
	HyphenCharacterForGlyphAtIndex(glyphIndex uint) uint32
	HyphenationFactorForGlyphAtIndex(glyphIndex uint) float32
	ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(charIndex uint) bool
	ShouldBreakLineByWordBeforeCharacterAtIndex(charIndex uint) bool
	WillSetLineFragmentRect_ForGlyphRange_UsedRect_BaselineOffset(lineRect *foundation.Rect, glyphRange foundation.Range, usedRect *foundation.Rect, baselineOffset *float64)
	SetHardInvalidation_ForGlyphRange(flag bool, glyphRange foundation.Range)
	CharacterRangeForGlyphRange_ActualGlyphRange(glyphRange foundation.Range, actualGlyphRange *foundation.Range) foundation.Range
	GlyphRangeForCharacterRange_ActualCharacterRange(charRange foundation.Range, actualCharRange *foundation.Range) foundation.Range
	SetAttachmentSize_ForGlyphRange(attachmentSize foundation.Size, glyphRange foundation.Range)
	SetDrawsOutsideLineFragment_ForGlyphRange(flag bool, glyphRange foundation.Range)
	SetLineFragmentRect_ForGlyphRange_UsedRect_BaselineOffset(fragmentRect foundation.Rect, glyphRange foundation.Range, usedRect foundation.Rect, baselineOffset float64)
	SetLocation_WithAdvancements_ForStartOfGlyphRange(location foundation.Point, advancements *float64, glyphRange foundation.Range)
	SetNotShownAttribute_ForGlyphRange(flag bool, glyphRange foundation.Range)
	// deprecated
	ActionForControlCharacterAtIndex(charIndex uint) TypesetterControlCharacterAction
	// deprecated
	DeleteGlyphsInRange(glyphRange foundation.Range)
	// deprecated
	SubstituteGlyphsInRange_WithGlyphs(glyphRange foundation.Range, glyphs *Glyph)
	// deprecated
	GetGlyphsInRange_Glyphs_CharacterIndexes_GlyphInscriptions_ElasticBits_BidiLevels(glyphsRange foundation.Range, glyphBuffer *Glyph, charIndexBuffer *uint, inscribeBuffer *GlyphInscription, elasticBuffer *bool, bidiLevelBuffer *byte) uint
	// deprecated
	InsertGlyph_AtGlyphIndex_CharacterIndex(glyph Glyph, glyphIndex uint, characterIndex uint)
	LayoutManager() LayoutManager
	UsesFontLeading() bool
	SetUsesFontLeading(value bool)
	TypesetterBehavior() TypesetterBehavior
	SetTypesetterBehavior(value TypesetterBehavior)
	HyphenationFactor() float32
	SetHyphenationFactor(value float32)
	CurrentTextContainer() TextContainer
	TextContainers() []TextContainer
	LineFragmentPadding() float64
	SetLineFragmentPadding(value float64)
	BidiProcessingEnabled() bool
	SetBidiProcessingEnabled(value bool)
	CurrentParagraphStyle() ParagraphStyle
	AttributedString() foundation.AttributedString
	SetAttributedString(value foundation.IAttributedString)
	ParagraphGlyphRange() foundation.Range
	ParagraphSeparatorGlyphRange() foundation.Range
	ParagraphCharacterRange() foundation.Range
	ParagraphSeparatorCharacterRange() foundation.Range
	AttributesForExtraLineFragment() map[foundation.AttributedStringKey]objc.Object
}

type Typesetter struct {
	objc.Object
}

func MakeTypesetter(ptr unsafe.Pointer) Typesetter {
	return Typesetter{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TypesetterClass) Alloc() Typesetter {
	rv := ffi.CallMethod[Typesetter](tc, "alloc")
	return rv
}

func (t_ Typesetter) Init() Typesetter {
	rv := ffi.CallMethod[Typesetter](t_, "init")
	return rv
}

func (tc _TypesetterClass) New() Typesetter {
	rv := ffi.CallMethod[Typesetter](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTypesetter() Typesetter {
	return TypesetterClass.New()
}

func (tc _TypesetterClass) SharedSystemTypesetterForBehavior(behavior TypesetterBehavior) objc.Object {
	rv := ffi.CallMethod[objc.Object](tc, "sharedSystemTypesetterForBehavior:", behavior)
	return rv
}

func (tc _TypesetterClass) PrintingAdjustmentInLayoutManager_ForNominallySpacedGlyphRange_PackedGlyphs_Count(layoutMgr ILayoutManager, nominallySpacedGlyphsRange foundation.Range, packedGlyphs *byte, packedGlyphsCount uint) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](tc, "printingAdjustmentInLayoutManager:forNominallySpacedGlyphRange:packedGlyphs:count:", layoutMgr, nominallySpacedGlyphsRange, packedGlyphs, packedGlyphsCount)
	return rv
}

func (t_ Typesetter) BaselineOffsetInLayoutManager_GlyphIndex(layoutMgr ILayoutManager, glyphIndex uint) float64 {
	rv := ffi.CallMethod[float64](t_, "baselineOffsetInLayoutManager:glyphIndex:", layoutMgr, glyphIndex)
	return rv
}

func (t_ Typesetter) SubstituteFontForFont(originalFont IFont) Font {
	rv := ffi.CallMethod[Font](t_, "substituteFontForFont:", originalFont)
	return rv
}

func (t_ Typesetter) TextTabForGlyphLocation_WritingDirection_MaxLocation(glyphLocation float64, direction WritingDirection, maxLocation float64) TextTab {
	rv := ffi.CallMethod[TextTab](t_, "textTabForGlyphLocation:writingDirection:maxLocation:", glyphLocation, direction, maxLocation)
	return rv
}

func (t_ Typesetter) SetParagraphGlyphRange_SeparatorGlyphRange(paragraphRange foundation.Range, paragraphSeparatorRange foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "setParagraphGlyphRange:separatorGlyphRange:", paragraphRange, paragraphSeparatorRange)
}

func (t_ Typesetter) LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) float64 {
	rv := ffi.CallMethod[float64](t_, "lineSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:", glyphIndex, rect)
	return rv
}

func (t_ Typesetter) ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) float64 {
	rv := ffi.CallMethod[float64](t_, "paragraphSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:", glyphIndex, rect)
	return rv
}

func (t_ Typesetter) ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) float64 {
	rv := ffi.CallMethod[float64](t_, "paragraphSpacingBeforeGlyphAtIndex:withProposedLineFragmentRect:", glyphIndex, rect)
	return rv
}

func (t_ Typesetter) LayoutParagraphAtPoint(lineFragmentOrigin *foundation.Point) uint {
	rv := ffi.CallMethod[uint](t_, "layoutParagraphAtPoint:", lineFragmentOrigin)
	return rv
}

func (t_ Typesetter) BeginParagraph() {
	ffi.CallMethod[ffi.Void](t_, "beginParagraph")
}

func (t_ Typesetter) EndParagraph() {
	ffi.CallMethod[ffi.Void](t_, "endParagraph")
}

func (t_ Typesetter) BeginLineWithGlyphAtIndex(glyphIndex uint) {
	ffi.CallMethod[ffi.Void](t_, "beginLineWithGlyphAtIndex:", glyphIndex)
}

func (t_ Typesetter) EndLineWithGlyphRange(lineGlyphRange foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "endLineWithGlyphRange:", lineGlyphRange)
}

func (t_ Typesetter) LayoutCharactersInRange_ForLayoutManager_MaximumNumberOfLineFragments(characterRange foundation.Range, layoutManager ILayoutManager, maxNumLines uint) foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "layoutCharactersInRange:forLayoutManager:maximumNumberOfLineFragments:", characterRange, layoutManager, maxNumLines)
	return rv
}

func (t_ Typesetter) LayoutGlyphsInLayoutManager_StartingAtGlyphIndex_MaxNumberOfLineFragments_NextGlyphIndex(layoutManager ILayoutManager, startGlyphIndex uint, maxNumLines uint, nextGlyph *uint) {
	ffi.CallMethod[ffi.Void](t_, "layoutGlyphsInLayoutManager:startingAtGlyphIndex:maxNumberOfLineFragments:nextGlyphIndex:", layoutManager, startGlyphIndex, maxNumLines, nextGlyph)
}

func (t_ Typesetter) BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex(glyphIndex uint, textContainer ITextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](t_, "boundingBoxForControlGlyphAtIndex:forTextContainer:proposedLineFragment:glyphPosition:characterIndex:", glyphIndex, textContainer, proposedRect, glyphPosition, charIndex)
	return rv
}

func (t_ Typesetter) GetLineFragmentRect_UsedRect_ForParagraphSeparatorGlyphRange_AtProposedOrigin(lineFragmentRect *foundation.Rect, lineFragmentUsedRect *foundation.Rect, paragraphSeparatorGlyphRange foundation.Range, lineOrigin foundation.Point) {
	ffi.CallMethod[ffi.Void](t_, "getLineFragmentRect:usedRect:forParagraphSeparatorGlyphRange:atProposedOrigin:", lineFragmentRect, lineFragmentUsedRect, paragraphSeparatorGlyphRange, lineOrigin)
}

func (t_ Typesetter) GetLineFragmentRect_UsedRect_RemainingRect_ForStartingGlyphAtIndex_ProposedRect_LineSpacing_ParagraphSpacingBefore_ParagraphSpacingAfter(lineFragmentRect *foundation.Rect, lineFragmentUsedRect *foundation.Rect, remainingRect *foundation.Rect, startingGlyphIndex uint, proposedRect foundation.Rect, lineSpacing float64, paragraphSpacingBefore float64, paragraphSpacingAfter float64) {
	ffi.CallMethod[ffi.Void](t_, "getLineFragmentRect:usedRect:remainingRect:forStartingGlyphAtIndex:proposedRect:lineSpacing:paragraphSpacingBefore:paragraphSpacingAfter:", lineFragmentRect, lineFragmentUsedRect, remainingRect, startingGlyphIndex, proposedRect, lineSpacing, paragraphSpacingBefore, paragraphSpacingAfter)
}

func (t_ Typesetter) HyphenCharacterForGlyphAtIndex(glyphIndex uint) uint32 {
	rv := ffi.CallMethod[uint32](t_, "hyphenCharacterForGlyphAtIndex:", glyphIndex)
	return rv
}

func (t_ Typesetter) HyphenationFactorForGlyphAtIndex(glyphIndex uint) float32 {
	rv := ffi.CallMethod[float32](t_, "hyphenationFactorForGlyphAtIndex:", glyphIndex)
	return rv
}

func (t_ Typesetter) ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(charIndex uint) bool {
	rv := ffi.CallMethod[bool](t_, "shouldBreakLineByHyphenatingBeforeCharacterAtIndex:", charIndex)
	return rv
}

func (t_ Typesetter) ShouldBreakLineByWordBeforeCharacterAtIndex(charIndex uint) bool {
	rv := ffi.CallMethod[bool](t_, "shouldBreakLineByWordBeforeCharacterAtIndex:", charIndex)
	return rv
}

func (t_ Typesetter) WillSetLineFragmentRect_ForGlyphRange_UsedRect_BaselineOffset(lineRect *foundation.Rect, glyphRange foundation.Range, usedRect *foundation.Rect, baselineOffset *float64) {
	ffi.CallMethod[ffi.Void](t_, "willSetLineFragmentRect:forGlyphRange:usedRect:baselineOffset:", lineRect, glyphRange, usedRect, baselineOffset)
}

func (t_ Typesetter) SetHardInvalidation_ForGlyphRange(flag bool, glyphRange foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "setHardInvalidation:forGlyphRange:", flag, glyphRange)
}

func (t_ Typesetter) CharacterRangeForGlyphRange_ActualGlyphRange(glyphRange foundation.Range, actualGlyphRange *foundation.Range) foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "characterRangeForGlyphRange:actualGlyphRange:", glyphRange, actualGlyphRange)
	return rv
}

func (t_ Typesetter) GlyphRangeForCharacterRange_ActualCharacterRange(charRange foundation.Range, actualCharRange *foundation.Range) foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "glyphRangeForCharacterRange:actualCharacterRange:", charRange, actualCharRange)
	return rv
}

func (t_ Typesetter) SetAttachmentSize_ForGlyphRange(attachmentSize foundation.Size, glyphRange foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "setAttachmentSize:forGlyphRange:", attachmentSize, glyphRange)
}

func (t_ Typesetter) SetDrawsOutsideLineFragment_ForGlyphRange(flag bool, glyphRange foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "setDrawsOutsideLineFragment:forGlyphRange:", flag, glyphRange)
}

func (t_ Typesetter) SetLineFragmentRect_ForGlyphRange_UsedRect_BaselineOffset(fragmentRect foundation.Rect, glyphRange foundation.Range, usedRect foundation.Rect, baselineOffset float64) {
	ffi.CallMethod[ffi.Void](t_, "setLineFragmentRect:forGlyphRange:usedRect:baselineOffset:", fragmentRect, glyphRange, usedRect, baselineOffset)
}

func (t_ Typesetter) SetLocation_WithAdvancements_ForStartOfGlyphRange(location foundation.Point, advancements *float64, glyphRange foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "setLocation:withAdvancements:forStartOfGlyphRange:", location, advancements, glyphRange)
}

func (t_ Typesetter) SetNotShownAttribute_ForGlyphRange(flag bool, glyphRange foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "setNotShownAttribute:forGlyphRange:", flag, glyphRange)
}

// deprecated
func (t_ Typesetter) ActionForControlCharacterAtIndex(charIndex uint) TypesetterControlCharacterAction {
	rv := ffi.CallMethod[TypesetterControlCharacterAction](t_, "actionForControlCharacterAtIndex:", charIndex)
	return rv
}

// deprecated
func (t_ Typesetter) DeleteGlyphsInRange(glyphRange foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "deleteGlyphsInRange:", glyphRange)
}

// deprecated
func (t_ Typesetter) SubstituteGlyphsInRange_WithGlyphs(glyphRange foundation.Range, glyphs *Glyph) {
	ffi.CallMethod[ffi.Void](t_, "substituteGlyphsInRange:withGlyphs:", glyphRange, glyphs)
}

// deprecated
func (t_ Typesetter) GetGlyphsInRange_Glyphs_CharacterIndexes_GlyphInscriptions_ElasticBits_BidiLevels(glyphsRange foundation.Range, glyphBuffer *Glyph, charIndexBuffer *uint, inscribeBuffer *GlyphInscription, elasticBuffer *bool, bidiLevelBuffer *byte) uint {
	rv := ffi.CallMethod[uint](t_, "getGlyphsInRange:glyphs:characterIndexes:glyphInscriptions:elasticBits:bidiLevels:", glyphsRange, glyphBuffer, charIndexBuffer, inscribeBuffer, elasticBuffer, bidiLevelBuffer)
	return rv
}

// deprecated
func (t_ Typesetter) InsertGlyph_AtGlyphIndex_CharacterIndex(glyph Glyph, glyphIndex uint, characterIndex uint) {
	ffi.CallMethod[ffi.Void](t_, "insertGlyph:atGlyphIndex:characterIndex:", glyph, glyphIndex, characterIndex)
}

func (tc _TypesetterClass) SharedSystemTypesetter() Typesetter {
	rv := ffi.CallMethod[Typesetter](tc, "sharedSystemTypesetter")
	return rv
}

func (tc _TypesetterClass) DefaultTypesetterBehavior() TypesetterBehavior {
	rv := ffi.CallMethod[TypesetterBehavior](tc, "defaultTypesetterBehavior")
	return rv
}

func (t_ Typesetter) LayoutManager() LayoutManager {
	rv := ffi.CallMethod[LayoutManager](t_, "layoutManager")
	return rv
}

func (t_ Typesetter) UsesFontLeading() bool {
	rv := ffi.CallMethod[bool](t_, "usesFontLeading")
	return rv
}

func (t_ Typesetter) SetUsesFontLeading(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setUsesFontLeading:", value)
}

func (t_ Typesetter) TypesetterBehavior() TypesetterBehavior {
	rv := ffi.CallMethod[TypesetterBehavior](t_, "typesetterBehavior")
	return rv
}

func (t_ Typesetter) SetTypesetterBehavior(value TypesetterBehavior) {
	ffi.CallMethod[ffi.Void](t_, "setTypesetterBehavior:", value)
}

func (t_ Typesetter) HyphenationFactor() float32 {
	rv := ffi.CallMethod[float32](t_, "hyphenationFactor")
	return rv
}

func (t_ Typesetter) SetHyphenationFactor(value float32) {
	ffi.CallMethod[ffi.Void](t_, "setHyphenationFactor:", value)
}

func (t_ Typesetter) CurrentTextContainer() TextContainer {
	rv := ffi.CallMethod[TextContainer](t_, "currentTextContainer")
	return rv
}

func (t_ Typesetter) TextContainers() []TextContainer {
	rv := ffi.CallMethod[[]TextContainer](t_, "textContainers")
	return rv
}

func (t_ Typesetter) LineFragmentPadding() float64 {
	rv := ffi.CallMethod[float64](t_, "lineFragmentPadding")
	return rv
}

func (t_ Typesetter) SetLineFragmentPadding(value float64) {
	ffi.CallMethod[ffi.Void](t_, "setLineFragmentPadding:", value)
}

func (t_ Typesetter) BidiProcessingEnabled() bool {
	rv := ffi.CallMethod[bool](t_, "bidiProcessingEnabled")
	return rv
}

func (t_ Typesetter) SetBidiProcessingEnabled(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setBidiProcessingEnabled:", value)
}

func (t_ Typesetter) CurrentParagraphStyle() ParagraphStyle {
	rv := ffi.CallMethod[ParagraphStyle](t_, "currentParagraphStyle")
	return rv
}

func (t_ Typesetter) AttributedString() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](t_, "attributedString")
	return rv
}

func (t_ Typesetter) SetAttributedString(value foundation.IAttributedString) {
	ffi.CallMethod[ffi.Void](t_, "setAttributedString:", value)
}

func (t_ Typesetter) ParagraphGlyphRange() foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "paragraphGlyphRange")
	return rv
}

func (t_ Typesetter) ParagraphSeparatorGlyphRange() foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "paragraphSeparatorGlyphRange")
	return rv
}

func (t_ Typesetter) ParagraphCharacterRange() foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "paragraphCharacterRange")
	return rv
}

func (t_ Typesetter) ParagraphSeparatorCharacterRange() foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "paragraphSeparatorCharacterRange")
	return rv
}

func (t_ Typesetter) AttributesForExtraLineFragment() map[foundation.AttributedStringKey]objc.Object {
	rv := ffi.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, "attributesForExtraLineFragment")
	return rv
}
