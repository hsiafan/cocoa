// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type LayoutManagerDelegate interface {
	ImplementsLayoutManagerDidInvalidateLayout() bool
	// optional
	LayoutManagerDidInvalidateLayout(sender LayoutManager)
	ImplementsLayoutManager_ShouldGenerateGlyphs_Properties_CharacterIndexes_Font_ForGlyphRange() bool
	// optional
	LayoutManager_ShouldGenerateGlyphs_Properties_CharacterIndexes_Font_ForGlyphRange(layoutManager LayoutManager, glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont Font, glyphRange foundation.Range) uint
	ImplementsLayoutManager_ShouldUseAction_ForControlCharacterAtIndex() bool
	// optional
	LayoutManager_ShouldUseAction_ForControlCharacterAtIndex(layoutManager LayoutManager, action ControlCharacterAction, charIndex uint) ControlCharacterAction
	ImplementsLayoutManager_DidCompleteLayoutForTextContainer_AtEnd() bool
	// optional
	LayoutManager_DidCompleteLayoutForTextContainer_AtEnd(layoutManager LayoutManager, textContainer TextContainer, layoutFinishedFlag bool)
	ImplementsLayoutManager_TextContainer_DidChangeGeometryFromSize() bool
	// optional
	LayoutManager_TextContainer_DidChangeGeometryFromSize(layoutManager LayoutManager, textContainer TextContainer, oldSize foundation.Size)
	ImplementsLayoutManager_ShouldBreakLineByHyphenatingBeforeCharacterAtIndex() bool
	// optional
	LayoutManager_ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(layoutManager LayoutManager, charIndex uint) bool
	ImplementsLayoutManager_ShouldBreakLineByWordBeforeCharacterAtIndex() bool
	// optional
	LayoutManager_ShouldBreakLineByWordBeforeCharacterAtIndex(layoutManager LayoutManager, charIndex uint) bool
	ImplementsLayoutManager_LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect() bool
	// optional
	LayoutManager_LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64
	ImplementsLayoutManager_ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect() bool
	// optional
	LayoutManager_ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64
	ImplementsLayoutManager_ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect() bool
	// optional
	LayoutManager_ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64
	ImplementsLayoutManager_BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex() bool
	// optional
	LayoutManager_BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex(layoutManager LayoutManager, glyphIndex uint, textContainer TextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect
	ImplementsLayoutManager_ShouldUseTemporaryAttributes_ForDrawingToScreen_AtCharacterIndex_EffectiveRange() bool
	// optional
	LayoutManager_ShouldUseTemporaryAttributes_ForDrawingToScreen_AtCharacterIndex_EffectiveRange(layoutManager LayoutManager, attrs map[foundation.AttributedStringKey]objc.Object, toScreen bool, charIndex uint, effectiveCharRange *foundation.Range) map[foundation.AttributedStringKey]objc.IObject
}

func WrapLayoutManagerDelegate(v LayoutManagerDelegate) objc.Object {
	return objc.WrapAsProtocol("NSLayoutManagerDelegate", v)
}

type LayoutManagerDelegateBase struct {
}

func (p *LayoutManagerDelegateBase) ImplementsLayoutManagerDidInvalidateLayout() bool {
	return false
}

func (p *LayoutManagerDelegateBase) LayoutManagerDidInvalidateLayout(sender LayoutManager) {
	panic("unimpemented protocol method")
}

func (p *LayoutManagerDelegateBase) ImplementsLayoutManager_ShouldGenerateGlyphs_Properties_CharacterIndexes_Font_ForGlyphRange() bool {
	return false
}

func (p *LayoutManagerDelegateBase) LayoutManager_ShouldGenerateGlyphs_Properties_CharacterIndexes_Font_ForGlyphRange(layoutManager LayoutManager, glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont Font, glyphRange foundation.Range) uint {
	panic("unimpemented protocol method")
}

func (p *LayoutManagerDelegateBase) ImplementsLayoutManager_ShouldUseAction_ForControlCharacterAtIndex() bool {
	return false
}

func (p *LayoutManagerDelegateBase) LayoutManager_ShouldUseAction_ForControlCharacterAtIndex(layoutManager LayoutManager, action ControlCharacterAction, charIndex uint) ControlCharacterAction {
	panic("unimpemented protocol method")
}

func (p *LayoutManagerDelegateBase) ImplementsLayoutManager_DidCompleteLayoutForTextContainer_AtEnd() bool {
	return false
}

func (p *LayoutManagerDelegateBase) LayoutManager_DidCompleteLayoutForTextContainer_AtEnd(layoutManager LayoutManager, textContainer TextContainer, layoutFinishedFlag bool) {
	panic("unimpemented protocol method")
}

func (p *LayoutManagerDelegateBase) ImplementsLayoutManager_TextContainer_DidChangeGeometryFromSize() bool {
	return false
}

func (p *LayoutManagerDelegateBase) LayoutManager_TextContainer_DidChangeGeometryFromSize(layoutManager LayoutManager, textContainer TextContainer, oldSize foundation.Size) {
	panic("unimpemented protocol method")
}

func (p *LayoutManagerDelegateBase) ImplementsLayoutManager_ShouldBreakLineByHyphenatingBeforeCharacterAtIndex() bool {
	return false
}

func (p *LayoutManagerDelegateBase) LayoutManager_ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(layoutManager LayoutManager, charIndex uint) bool {
	panic("unimpemented protocol method")
}

func (p *LayoutManagerDelegateBase) ImplementsLayoutManager_ShouldBreakLineByWordBeforeCharacterAtIndex() bool {
	return false
}

func (p *LayoutManagerDelegateBase) LayoutManager_ShouldBreakLineByWordBeforeCharacterAtIndex(layoutManager LayoutManager, charIndex uint) bool {
	panic("unimpemented protocol method")
}

func (p *LayoutManagerDelegateBase) ImplementsLayoutManager_LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect() bool {
	return false
}

func (p *LayoutManagerDelegateBase) LayoutManager_LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64 {
	panic("unimpemented protocol method")
}

func (p *LayoutManagerDelegateBase) ImplementsLayoutManager_ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect() bool {
	return false
}

func (p *LayoutManagerDelegateBase) LayoutManager_ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64 {
	panic("unimpemented protocol method")
}

func (p *LayoutManagerDelegateBase) ImplementsLayoutManager_ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect() bool {
	return false
}

func (p *LayoutManagerDelegateBase) LayoutManager_ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64 {
	panic("unimpemented protocol method")
}

func (p *LayoutManagerDelegateBase) ImplementsLayoutManager_BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex() bool {
	return false
}

func (p *LayoutManagerDelegateBase) LayoutManager_BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex(layoutManager LayoutManager, glyphIndex uint, textContainer TextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect {
	panic("unimpemented protocol method")
}

func (p *LayoutManagerDelegateBase) ImplementsLayoutManager_ShouldUseTemporaryAttributes_ForDrawingToScreen_AtCharacterIndex_EffectiveRange() bool {
	return false
}

func (p *LayoutManagerDelegateBase) LayoutManager_ShouldUseTemporaryAttributes_ForDrawingToScreen_AtCharacterIndex_EffectiveRange(layoutManager LayoutManager, attrs map[foundation.AttributedStringKey]objc.Object, toScreen bool, charIndex uint, effectiveCharRange *foundation.Range) map[foundation.AttributedStringKey]objc.IObject {
	panic("unimpemented protocol method")
}
