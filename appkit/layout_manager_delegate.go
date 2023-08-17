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

type LayoutManagerDelegateWrapper struct {
	objc.Object
}

func (l_ LayoutManagerDelegateWrapper) LayoutManagerDidInvalidateLayout(sender ILayoutManager) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("layoutManagerDidInvalidateLayout:"), objc.ExtractPtr(sender))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_ShouldGenerateGlyphs_Properties_CharacterIndexes_Font_ForGlyphRange(layoutManager ILayoutManager, glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont IFont, glyphRange foundation.Range) uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("layoutManager:shouldGenerateGlyphs:properties:characterIndexes:font:forGlyphRange:"), objc.ExtractPtr(layoutManager), glyphs, props, charIndexes, objc.ExtractPtr(aFont), glyphRange)
	return rv
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_ShouldUseAction_ForControlCharacterAtIndex(layoutManager ILayoutManager, action ControlCharacterAction, charIndex uint) ControlCharacterAction {
	rv := objc.CallMethod[ControlCharacterAction](l_, objc.GetSelector("layoutManager:shouldUseAction:forControlCharacterAtIndex:"), objc.ExtractPtr(layoutManager), action, charIndex)
	return rv
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_DidCompleteLayoutForTextContainer_AtEnd(layoutManager ILayoutManager, textContainer ITextContainer, layoutFinishedFlag bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("layoutManager:didCompleteLayoutForTextContainer:atEnd:"), objc.ExtractPtr(layoutManager), objc.ExtractPtr(textContainer), layoutFinishedFlag)
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_TextContainer_DidChangeGeometryFromSize(layoutManager ILayoutManager, textContainer ITextContainer, oldSize foundation.Size) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("layoutManager:textContainer:didChangeGeometryFromSize:"), objc.ExtractPtr(layoutManager), objc.ExtractPtr(textContainer), oldSize)
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(layoutManager ILayoutManager, charIndex uint) bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("layoutManager:shouldBreakLineByHyphenatingBeforeCharacterAtIndex:"), objc.ExtractPtr(layoutManager), charIndex)
	return rv
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_ShouldBreakLineByWordBeforeCharacterAtIndex(layoutManager ILayoutManager, charIndex uint) bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("layoutManager:shouldBreakLineByWordBeforeCharacterAtIndex:"), objc.ExtractPtr(layoutManager), charIndex)
	return rv
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(layoutManager ILayoutManager, glyphIndex uint, rect foundation.Rect) float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("layoutManager:lineSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), objc.ExtractPtr(layoutManager), glyphIndex, rect)
	return rv
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(layoutManager ILayoutManager, glyphIndex uint, rect foundation.Rect) float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("layoutManager:paragraphSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), objc.ExtractPtr(layoutManager), glyphIndex, rect)
	return rv
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect(layoutManager ILayoutManager, glyphIndex uint, rect foundation.Rect) float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("layoutManager:paragraphSpacingBeforeGlyphAtIndex:withProposedLineFragmentRect:"), objc.ExtractPtr(layoutManager), glyphIndex, rect)
	return rv
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex(layoutManager ILayoutManager, glyphIndex uint, textContainer ITextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("layoutManager:boundingBoxForControlGlyphAtIndex:forTextContainer:proposedLineFragment:glyphPosition:characterIndex:"), objc.ExtractPtr(layoutManager), glyphIndex, objc.ExtractPtr(textContainer), proposedRect, glyphPosition, charIndex)
	return rv
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_ShouldUseTemporaryAttributes_ForDrawingToScreen_AtCharacterIndex_EffectiveRange(layoutManager ILayoutManager, attrs map[foundation.AttributedStringKey]objc.IObject, toScreen bool, charIndex uint, effectiveCharRange *foundation.Range) map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](l_, objc.GetSelector("layoutManager:shouldUseTemporaryAttributes:forDrawingToScreen:atCharacterIndex:effectiveRange:"), objc.ExtractPtr(layoutManager), attrs, toScreen, charIndex, effectiveCharRange)
	return rv
}
