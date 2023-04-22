// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[Typesetter](tc, objc.GetSelector("alloc"))
	return rv
}

func (tc _TypesetterClass) New() Typesetter {
	rv := objc.CallMethod[Typesetter](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTypesetter() Typesetter {
	return TypesetterClass.New()
}

func (t_ Typesetter) Init() Typesetter {
	rv := objc.CallMethod[Typesetter](t_, objc.GetSelector("init"))
	return rv
}

func (tc _TypesetterClass) SharedSystemTypesetterForBehavior(behavior TypesetterBehavior) objc.Object {
	rv := objc.CallMethod[objc.Object](tc, objc.GetSelector("sharedSystemTypesetterForBehavior:"), behavior)
	return rv
}

func (tc _TypesetterClass) PrintingAdjustmentInLayoutManager_ForNominallySpacedGlyphRange_PackedGlyphs_Count(layoutMgr ILayoutManager, nominallySpacedGlyphsRange foundation.Range, packedGlyphs *byte, packedGlyphsCount uint) foundation.Size {
	rv := objc.CallMethod[foundation.Size](tc, objc.GetSelector("printingAdjustmentInLayoutManager:forNominallySpacedGlyphRange:packedGlyphs:count:"), layoutMgr, nominallySpacedGlyphsRange, packedGlyphs, packedGlyphsCount)
	return rv
}

func (t_ Typesetter) BaselineOffsetInLayoutManager_GlyphIndex(layoutMgr ILayoutManager, glyphIndex uint) float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("baselineOffsetInLayoutManager:glyphIndex:"), layoutMgr, glyphIndex)
	return rv
}

func (t_ Typesetter) SubstituteFontForFont(originalFont IFont) Font {
	rv := objc.CallMethod[Font](t_, objc.GetSelector("substituteFontForFont:"), originalFont)
	return rv
}

func (t_ Typesetter) TextTabForGlyphLocation_WritingDirection_MaxLocation(glyphLocation float64, direction WritingDirection, maxLocation float64) TextTab {
	rv := objc.CallMethod[TextTab](t_, objc.GetSelector("textTabForGlyphLocation:writingDirection:maxLocation:"), glyphLocation, direction, maxLocation)
	return rv
}

func (t_ Typesetter) SetParagraphGlyphRange_SeparatorGlyphRange(paragraphRange foundation.Range, paragraphSeparatorRange foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setParagraphGlyphRange:separatorGlyphRange:"), paragraphRange, paragraphSeparatorRange)
}

func (t_ Typesetter) LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("lineSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), glyphIndex, rect)
	return rv
}

func (t_ Typesetter) ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("paragraphSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), glyphIndex, rect)
	return rv
}

func (t_ Typesetter) ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("paragraphSpacingBeforeGlyphAtIndex:withProposedLineFragmentRect:"), glyphIndex, rect)
	return rv
}

func (t_ Typesetter) LayoutParagraphAtPoint(lineFragmentOrigin *foundation.Point) uint {
	rv := objc.CallMethod[uint](t_, objc.GetSelector("layoutParagraphAtPoint:"), lineFragmentOrigin)
	return rv
}

func (t_ Typesetter) BeginParagraph() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("beginParagraph"))
}

func (t_ Typesetter) EndParagraph() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("endParagraph"))
}

func (t_ Typesetter) BeginLineWithGlyphAtIndex(glyphIndex uint) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("beginLineWithGlyphAtIndex:"), glyphIndex)
}

func (t_ Typesetter) EndLineWithGlyphRange(lineGlyphRange foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("endLineWithGlyphRange:"), lineGlyphRange)
}

func (t_ Typesetter) LayoutCharactersInRange_ForLayoutManager_MaximumNumberOfLineFragments(characterRange foundation.Range, layoutManager ILayoutManager, maxNumLines uint) foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("layoutCharactersInRange:forLayoutManager:maximumNumberOfLineFragments:"), characterRange, layoutManager, maxNumLines)
	return rv
}

func (t_ Typesetter) LayoutGlyphsInLayoutManager_StartingAtGlyphIndex_MaxNumberOfLineFragments_NextGlyphIndex(layoutManager ILayoutManager, startGlyphIndex uint, maxNumLines uint, nextGlyph *uint) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("layoutGlyphsInLayoutManager:startingAtGlyphIndex:maxNumberOfLineFragments:nextGlyphIndex:"), layoutManager, startGlyphIndex, maxNumLines, nextGlyph)
}

func (t_ Typesetter) BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex(glyphIndex uint, textContainer ITextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.GetSelector("boundingBoxForControlGlyphAtIndex:forTextContainer:proposedLineFragment:glyphPosition:characterIndex:"), glyphIndex, textContainer, proposedRect, glyphPosition, charIndex)
	return rv
}

func (t_ Typesetter) GetLineFragmentRect_UsedRect_ForParagraphSeparatorGlyphRange_AtProposedOrigin(lineFragmentRect *foundation.Rect, lineFragmentUsedRect *foundation.Rect, paragraphSeparatorGlyphRange foundation.Range, lineOrigin foundation.Point) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("getLineFragmentRect:usedRect:forParagraphSeparatorGlyphRange:atProposedOrigin:"), lineFragmentRect, lineFragmentUsedRect, paragraphSeparatorGlyphRange, lineOrigin)
}

func (t_ Typesetter) GetLineFragmentRect_UsedRect_RemainingRect_ForStartingGlyphAtIndex_ProposedRect_LineSpacing_ParagraphSpacingBefore_ParagraphSpacingAfter(lineFragmentRect *foundation.Rect, lineFragmentUsedRect *foundation.Rect, remainingRect *foundation.Rect, startingGlyphIndex uint, proposedRect foundation.Rect, lineSpacing float64, paragraphSpacingBefore float64, paragraphSpacingAfter float64) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("getLineFragmentRect:usedRect:remainingRect:forStartingGlyphAtIndex:proposedRect:lineSpacing:paragraphSpacingBefore:paragraphSpacingAfter:"), lineFragmentRect, lineFragmentUsedRect, remainingRect, startingGlyphIndex, proposedRect, lineSpacing, paragraphSpacingBefore, paragraphSpacingAfter)
}

func (t_ Typesetter) HyphenCharacterForGlyphAtIndex(glyphIndex uint) uint32 {
	rv := objc.CallMethod[uint32](t_, objc.GetSelector("hyphenCharacterForGlyphAtIndex:"), glyphIndex)
	return rv
}

func (t_ Typesetter) HyphenationFactorForGlyphAtIndex(glyphIndex uint) float32 {
	rv := objc.CallMethod[float32](t_, objc.GetSelector("hyphenationFactorForGlyphAtIndex:"), glyphIndex)
	return rv
}

func (t_ Typesetter) ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(charIndex uint) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("shouldBreakLineByHyphenatingBeforeCharacterAtIndex:"), charIndex)
	return rv
}

func (t_ Typesetter) ShouldBreakLineByWordBeforeCharacterAtIndex(charIndex uint) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("shouldBreakLineByWordBeforeCharacterAtIndex:"), charIndex)
	return rv
}

func (t_ Typesetter) WillSetLineFragmentRect_ForGlyphRange_UsedRect_BaselineOffset(lineRect *foundation.Rect, glyphRange foundation.Range, usedRect *foundation.Rect, baselineOffset *float64) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("willSetLineFragmentRect:forGlyphRange:usedRect:baselineOffset:"), lineRect, glyphRange, usedRect, baselineOffset)
}

func (t_ Typesetter) SetHardInvalidation_ForGlyphRange(flag bool, glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setHardInvalidation:forGlyphRange:"), flag, glyphRange)
}

func (t_ Typesetter) CharacterRangeForGlyphRange_ActualGlyphRange(glyphRange foundation.Range, actualGlyphRange *foundation.Range) foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("characterRangeForGlyphRange:actualGlyphRange:"), glyphRange, actualGlyphRange)
	return rv
}

func (t_ Typesetter) GlyphRangeForCharacterRange_ActualCharacterRange(charRange foundation.Range, actualCharRange *foundation.Range) foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("glyphRangeForCharacterRange:actualCharacterRange:"), charRange, actualCharRange)
	return rv
}

func (t_ Typesetter) SetAttachmentSize_ForGlyphRange(attachmentSize foundation.Size, glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAttachmentSize:forGlyphRange:"), attachmentSize, glyphRange)
}

func (t_ Typesetter) SetDrawsOutsideLineFragment_ForGlyphRange(flag bool, glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDrawsOutsideLineFragment:forGlyphRange:"), flag, glyphRange)
}

func (t_ Typesetter) SetLineFragmentRect_ForGlyphRange_UsedRect_BaselineOffset(fragmentRect foundation.Rect, glyphRange foundation.Range, usedRect foundation.Rect, baselineOffset float64) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setLineFragmentRect:forGlyphRange:usedRect:baselineOffset:"), fragmentRect, glyphRange, usedRect, baselineOffset)
}

func (t_ Typesetter) SetLocation_WithAdvancements_ForStartOfGlyphRange(location foundation.Point, advancements *float64, glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setLocation:withAdvancements:forStartOfGlyphRange:"), location, advancements, glyphRange)
}

func (t_ Typesetter) SetNotShownAttribute_ForGlyphRange(flag bool, glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setNotShownAttribute:forGlyphRange:"), flag, glyphRange)
}

// deprecated
func (t_ Typesetter) ActionForControlCharacterAtIndex(charIndex uint) TypesetterControlCharacterAction {
	rv := objc.CallMethod[TypesetterControlCharacterAction](t_, objc.GetSelector("actionForControlCharacterAtIndex:"), charIndex)
	return rv
}

// deprecated
func (t_ Typesetter) DeleteGlyphsInRange(glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("deleteGlyphsInRange:"), glyphRange)
}

// deprecated
func (t_ Typesetter) SubstituteGlyphsInRange_WithGlyphs(glyphRange foundation.Range, glyphs *Glyph) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("substituteGlyphsInRange:withGlyphs:"), glyphRange, glyphs)
}

// deprecated
func (t_ Typesetter) GetGlyphsInRange_Glyphs_CharacterIndexes_GlyphInscriptions_ElasticBits_BidiLevels(glyphsRange foundation.Range, glyphBuffer *Glyph, charIndexBuffer *uint, inscribeBuffer *GlyphInscription, elasticBuffer *bool, bidiLevelBuffer *byte) uint {
	rv := objc.CallMethod[uint](t_, objc.GetSelector("getGlyphsInRange:glyphs:characterIndexes:glyphInscriptions:elasticBits:bidiLevels:"), glyphsRange, glyphBuffer, charIndexBuffer, inscribeBuffer, elasticBuffer, bidiLevelBuffer)
	return rv
}

// deprecated
func (t_ Typesetter) InsertGlyph_AtGlyphIndex_CharacterIndex(glyph Glyph, glyphIndex uint, characterIndex uint) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("insertGlyph:atGlyphIndex:characterIndex:"), glyph, glyphIndex, characterIndex)
}

func (tc _TypesetterClass) SharedSystemTypesetter() Typesetter {
	rv := objc.CallMethod[Typesetter](tc, objc.GetSelector("sharedSystemTypesetter"))
	return rv
}

func (tc _TypesetterClass) DefaultTypesetterBehavior() TypesetterBehavior {
	rv := objc.CallMethod[TypesetterBehavior](tc, objc.GetSelector("defaultTypesetterBehavior"))
	return rv
}

func (t_ Typesetter) LayoutManager() LayoutManager {
	rv := objc.CallMethod[LayoutManager](t_, objc.GetSelector("layoutManager"))
	return rv
}

func (t_ Typesetter) UsesFontLeading() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("usesFontLeading"))
	return rv
}

func (t_ Typesetter) SetUsesFontLeading(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setUsesFontLeading:"), value)
}

func (t_ Typesetter) TypesetterBehavior() TypesetterBehavior {
	rv := objc.CallMethod[TypesetterBehavior](t_, objc.GetSelector("typesetterBehavior"))
	return rv
}

func (t_ Typesetter) SetTypesetterBehavior(value TypesetterBehavior) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTypesetterBehavior:"), value)
}

func (t_ Typesetter) HyphenationFactor() float32 {
	rv := objc.CallMethod[float32](t_, objc.GetSelector("hyphenationFactor"))
	return rv
}

func (t_ Typesetter) SetHyphenationFactor(value float32) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setHyphenationFactor:"), value)
}

func (t_ Typesetter) CurrentTextContainer() TextContainer {
	rv := objc.CallMethod[TextContainer](t_, objc.GetSelector("currentTextContainer"))
	return rv
}

func (t_ Typesetter) TextContainers() []TextContainer {
	rv := objc.CallMethod[[]TextContainer](t_, objc.GetSelector("textContainers"))
	return rv
}

func (t_ Typesetter) LineFragmentPadding() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("lineFragmentPadding"))
	return rv
}

func (t_ Typesetter) SetLineFragmentPadding(value float64) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setLineFragmentPadding:"), value)
}

func (t_ Typesetter) BidiProcessingEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("bidiProcessingEnabled"))
	return rv
}

func (t_ Typesetter) SetBidiProcessingEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setBidiProcessingEnabled:"), value)
}

func (t_ Typesetter) CurrentParagraphStyle() ParagraphStyle {
	rv := objc.CallMethod[ParagraphStyle](t_, objc.GetSelector("currentParagraphStyle"))
	return rv
}

func (t_ Typesetter) AttributedString() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](t_, objc.GetSelector("attributedString"))
	return rv
}

func (t_ Typesetter) SetAttributedString(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAttributedString:"), value)
}

func (t_ Typesetter) ParagraphGlyphRange() foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("paragraphGlyphRange"))
	return rv
}

func (t_ Typesetter) ParagraphSeparatorGlyphRange() foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("paragraphSeparatorGlyphRange"))
	return rv
}

func (t_ Typesetter) ParagraphCharacterRange() foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("paragraphCharacterRange"))
	return rv
}

func (t_ Typesetter) ParagraphSeparatorCharacterRange() foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("paragraphSeparatorCharacterRange"))
	return rv
}

func (t_ Typesetter) AttributesForExtraLineFragment() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, objc.GetSelector("attributesForExtraLineFragment"))
	return rv
}
