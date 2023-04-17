// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
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

type LayoutManagerDelegateImpl struct {
	_LayoutManagerDidInvalidateLayout                                                                                   func(sender LayoutManager)
	_LayoutManager_ShouldGenerateGlyphs_Properties_CharacterIndexes_Font_ForGlyphRange                                  func(layoutManager LayoutManager, glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont Font, glyphRange foundation.Range) uint
	_LayoutManager_ShouldUseAction_ForControlCharacterAtIndex                                                           func(layoutManager LayoutManager, action ControlCharacterAction, charIndex uint) ControlCharacterAction
	_LayoutManager_DidCompleteLayoutForTextContainer_AtEnd                                                              func(layoutManager LayoutManager, textContainer TextContainer, layoutFinishedFlag bool)
	_LayoutManager_TextContainer_DidChangeGeometryFromSize                                                              func(layoutManager LayoutManager, textContainer TextContainer, oldSize foundation.Size)
	_LayoutManager_ShouldBreakLineByHyphenatingBeforeCharacterAtIndex                                                   func(layoutManager LayoutManager, charIndex uint) bool
	_LayoutManager_ShouldBreakLineByWordBeforeCharacterAtIndex                                                          func(layoutManager LayoutManager, charIndex uint) bool
	_LayoutManager_LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect                                            func(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64
	_LayoutManager_ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect                                       func(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64
	_LayoutManager_ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect                                      func(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64
	_LayoutManager_BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex func(layoutManager LayoutManager, glyphIndex uint, textContainer TextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect
	_LayoutManager_ShouldUseTemporaryAttributes_ForDrawingToScreen_AtCharacterIndex_EffectiveRange                      func(layoutManager LayoutManager, attrs map[foundation.AttributedStringKey]objc.Object, toScreen bool, charIndex uint, effectiveCharRange *foundation.Range) map[foundation.AttributedStringKey]objc.IObject
}

func (di *LayoutManagerDelegateImpl) ImplementsLayoutManagerDidInvalidateLayout() bool {
	return di._LayoutManagerDidInvalidateLayout != nil
}

func (di *LayoutManagerDelegateImpl) SetLayoutManagerDidInvalidateLayout(f func(sender LayoutManager)) {
	di._LayoutManagerDidInvalidateLayout = f
}

func (di *LayoutManagerDelegateImpl) LayoutManagerDidInvalidateLayout(sender LayoutManager) {
	di._LayoutManagerDidInvalidateLayout(sender)
}
func (di *LayoutManagerDelegateImpl) ImplementsLayoutManager_ShouldGenerateGlyphs_Properties_CharacterIndexes_Font_ForGlyphRange() bool {
	return di._LayoutManager_ShouldGenerateGlyphs_Properties_CharacterIndexes_Font_ForGlyphRange != nil
}

func (di *LayoutManagerDelegateImpl) SetLayoutManager_ShouldGenerateGlyphs_Properties_CharacterIndexes_Font_ForGlyphRange(f func(layoutManager LayoutManager, glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont Font, glyphRange foundation.Range) uint) {
	di._LayoutManager_ShouldGenerateGlyphs_Properties_CharacterIndexes_Font_ForGlyphRange = f
}

func (di *LayoutManagerDelegateImpl) LayoutManager_ShouldGenerateGlyphs_Properties_CharacterIndexes_Font_ForGlyphRange(layoutManager LayoutManager, glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont Font, glyphRange foundation.Range) uint {
	return di._LayoutManager_ShouldGenerateGlyphs_Properties_CharacterIndexes_Font_ForGlyphRange(layoutManager, glyphs, props, charIndexes, aFont, glyphRange)
}
func (di *LayoutManagerDelegateImpl) ImplementsLayoutManager_ShouldUseAction_ForControlCharacterAtIndex() bool {
	return di._LayoutManager_ShouldUseAction_ForControlCharacterAtIndex != nil
}

func (di *LayoutManagerDelegateImpl) SetLayoutManager_ShouldUseAction_ForControlCharacterAtIndex(f func(layoutManager LayoutManager, action ControlCharacterAction, charIndex uint) ControlCharacterAction) {
	di._LayoutManager_ShouldUseAction_ForControlCharacterAtIndex = f
}

func (di *LayoutManagerDelegateImpl) LayoutManager_ShouldUseAction_ForControlCharacterAtIndex(layoutManager LayoutManager, action ControlCharacterAction, charIndex uint) ControlCharacterAction {
	return di._LayoutManager_ShouldUseAction_ForControlCharacterAtIndex(layoutManager, action, charIndex)
}
func (di *LayoutManagerDelegateImpl) ImplementsLayoutManager_DidCompleteLayoutForTextContainer_AtEnd() bool {
	return di._LayoutManager_DidCompleteLayoutForTextContainer_AtEnd != nil
}

func (di *LayoutManagerDelegateImpl) SetLayoutManager_DidCompleteLayoutForTextContainer_AtEnd(f func(layoutManager LayoutManager, textContainer TextContainer, layoutFinishedFlag bool)) {
	di._LayoutManager_DidCompleteLayoutForTextContainer_AtEnd = f
}

func (di *LayoutManagerDelegateImpl) LayoutManager_DidCompleteLayoutForTextContainer_AtEnd(layoutManager LayoutManager, textContainer TextContainer, layoutFinishedFlag bool) {
	di._LayoutManager_DidCompleteLayoutForTextContainer_AtEnd(layoutManager, textContainer, layoutFinishedFlag)
}
func (di *LayoutManagerDelegateImpl) ImplementsLayoutManager_TextContainer_DidChangeGeometryFromSize() bool {
	return di._LayoutManager_TextContainer_DidChangeGeometryFromSize != nil
}

func (di *LayoutManagerDelegateImpl) SetLayoutManager_TextContainer_DidChangeGeometryFromSize(f func(layoutManager LayoutManager, textContainer TextContainer, oldSize foundation.Size)) {
	di._LayoutManager_TextContainer_DidChangeGeometryFromSize = f
}

func (di *LayoutManagerDelegateImpl) LayoutManager_TextContainer_DidChangeGeometryFromSize(layoutManager LayoutManager, textContainer TextContainer, oldSize foundation.Size) {
	di._LayoutManager_TextContainer_DidChangeGeometryFromSize(layoutManager, textContainer, oldSize)
}
func (di *LayoutManagerDelegateImpl) ImplementsLayoutManager_ShouldBreakLineByHyphenatingBeforeCharacterAtIndex() bool {
	return di._LayoutManager_ShouldBreakLineByHyphenatingBeforeCharacterAtIndex != nil
}

func (di *LayoutManagerDelegateImpl) SetLayoutManager_ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(f func(layoutManager LayoutManager, charIndex uint) bool) {
	di._LayoutManager_ShouldBreakLineByHyphenatingBeforeCharacterAtIndex = f
}

func (di *LayoutManagerDelegateImpl) LayoutManager_ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(layoutManager LayoutManager, charIndex uint) bool {
	return di._LayoutManager_ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(layoutManager, charIndex)
}
func (di *LayoutManagerDelegateImpl) ImplementsLayoutManager_ShouldBreakLineByWordBeforeCharacterAtIndex() bool {
	return di._LayoutManager_ShouldBreakLineByWordBeforeCharacterAtIndex != nil
}

func (di *LayoutManagerDelegateImpl) SetLayoutManager_ShouldBreakLineByWordBeforeCharacterAtIndex(f func(layoutManager LayoutManager, charIndex uint) bool) {
	di._LayoutManager_ShouldBreakLineByWordBeforeCharacterAtIndex = f
}

func (di *LayoutManagerDelegateImpl) LayoutManager_ShouldBreakLineByWordBeforeCharacterAtIndex(layoutManager LayoutManager, charIndex uint) bool {
	return di._LayoutManager_ShouldBreakLineByWordBeforeCharacterAtIndex(layoutManager, charIndex)
}
func (di *LayoutManagerDelegateImpl) ImplementsLayoutManager_LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect() bool {
	return di._LayoutManager_LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect != nil
}

func (di *LayoutManagerDelegateImpl) SetLayoutManager_LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(f func(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64) {
	di._LayoutManager_LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect = f
}

func (di *LayoutManagerDelegateImpl) LayoutManager_LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64 {
	return di._LayoutManager_LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(layoutManager, glyphIndex, rect)
}
func (di *LayoutManagerDelegateImpl) ImplementsLayoutManager_ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect() bool {
	return di._LayoutManager_ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect != nil
}

func (di *LayoutManagerDelegateImpl) SetLayoutManager_ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(f func(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64) {
	di._LayoutManager_ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect = f
}

func (di *LayoutManagerDelegateImpl) LayoutManager_ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64 {
	return di._LayoutManager_ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(layoutManager, glyphIndex, rect)
}
func (di *LayoutManagerDelegateImpl) ImplementsLayoutManager_ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect() bool {
	return di._LayoutManager_ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect != nil
}

func (di *LayoutManagerDelegateImpl) SetLayoutManager_ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect(f func(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64) {
	di._LayoutManager_ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect = f
}

func (di *LayoutManagerDelegateImpl) LayoutManager_ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64 {
	return di._LayoutManager_ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect(layoutManager, glyphIndex, rect)
}
func (di *LayoutManagerDelegateImpl) ImplementsLayoutManager_BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex() bool {
	return di._LayoutManager_BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex != nil
}

func (di *LayoutManagerDelegateImpl) SetLayoutManager_BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex(f func(layoutManager LayoutManager, glyphIndex uint, textContainer TextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect) {
	di._LayoutManager_BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex = f
}

func (di *LayoutManagerDelegateImpl) LayoutManager_BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex(layoutManager LayoutManager, glyphIndex uint, textContainer TextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect {
	return di._LayoutManager_BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex(layoutManager, glyphIndex, textContainer, proposedRect, glyphPosition, charIndex)
}
func (di *LayoutManagerDelegateImpl) ImplementsLayoutManager_ShouldUseTemporaryAttributes_ForDrawingToScreen_AtCharacterIndex_EffectiveRange() bool {
	return di._LayoutManager_ShouldUseTemporaryAttributes_ForDrawingToScreen_AtCharacterIndex_EffectiveRange != nil
}

func (di *LayoutManagerDelegateImpl) SetLayoutManager_ShouldUseTemporaryAttributes_ForDrawingToScreen_AtCharacterIndex_EffectiveRange(f func(layoutManager LayoutManager, attrs map[foundation.AttributedStringKey]objc.Object, toScreen bool, charIndex uint, effectiveCharRange *foundation.Range) map[foundation.AttributedStringKey]objc.IObject) {
	di._LayoutManager_ShouldUseTemporaryAttributes_ForDrawingToScreen_AtCharacterIndex_EffectiveRange = f
}

func (di *LayoutManagerDelegateImpl) LayoutManager_ShouldUseTemporaryAttributes_ForDrawingToScreen_AtCharacterIndex_EffectiveRange(layoutManager LayoutManager, attrs map[foundation.AttributedStringKey]objc.Object, toScreen bool, charIndex uint, effectiveCharRange *foundation.Range) map[foundation.AttributedStringKey]objc.IObject {
	return di._LayoutManager_ShouldUseTemporaryAttributes_ForDrawingToScreen_AtCharacterIndex_EffectiveRange(layoutManager, attrs, toScreen, charIndex, effectiveCharRange)
}

type LayoutManagerDelegateWrapper struct {
	objc.Object
}

func (l_ *LayoutManagerDelegateWrapper) ImplementsLayoutManagerDidInvalidateLayout() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManagerDidInvalidateLayout:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManagerDidInvalidateLayout(sender ILayoutManager) {
	ffi.CallMethod[ffi.Void](l_, "layoutManagerDidInvalidateLayout:", sender)
}

func (l_ *LayoutManagerDelegateWrapper) ImplementsLayoutManager_ShouldGenerateGlyphs_Properties_CharacterIndexes_Font_ForGlyphRange() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:shouldGenerateGlyphs:properties:characterIndexes:font:forGlyphRange:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_ShouldGenerateGlyphs_Properties_CharacterIndexes_Font_ForGlyphRange(layoutManager ILayoutManager, glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont IFont, glyphRange foundation.Range) uint {
	rv := ffi.CallMethod[uint](l_, "layoutManager:shouldGenerateGlyphs:properties:characterIndexes:font:forGlyphRange:", layoutManager, glyphs, props, charIndexes, aFont, glyphRange)
	return rv
}

func (l_ *LayoutManagerDelegateWrapper) ImplementsLayoutManager_ShouldUseAction_ForControlCharacterAtIndex() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:shouldUseAction:forControlCharacterAtIndex:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_ShouldUseAction_ForControlCharacterAtIndex(layoutManager ILayoutManager, action ControlCharacterAction, charIndex uint) ControlCharacterAction {
	rv := ffi.CallMethod[ControlCharacterAction](l_, "layoutManager:shouldUseAction:forControlCharacterAtIndex:", layoutManager, action, charIndex)
	return rv
}

func (l_ *LayoutManagerDelegateWrapper) ImplementsLayoutManager_DidCompleteLayoutForTextContainer_AtEnd() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:didCompleteLayoutForTextContainer:atEnd:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_DidCompleteLayoutForTextContainer_AtEnd(layoutManager ILayoutManager, textContainer ITextContainer, layoutFinishedFlag bool) {
	ffi.CallMethod[ffi.Void](l_, "layoutManager:didCompleteLayoutForTextContainer:atEnd:", layoutManager, textContainer, layoutFinishedFlag)
}

func (l_ *LayoutManagerDelegateWrapper) ImplementsLayoutManager_TextContainer_DidChangeGeometryFromSize() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:textContainer:didChangeGeometryFromSize:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_TextContainer_DidChangeGeometryFromSize(layoutManager ILayoutManager, textContainer ITextContainer, oldSize foundation.Size) {
	ffi.CallMethod[ffi.Void](l_, "layoutManager:textContainer:didChangeGeometryFromSize:", layoutManager, textContainer, oldSize)
}

func (l_ *LayoutManagerDelegateWrapper) ImplementsLayoutManager_ShouldBreakLineByHyphenatingBeforeCharacterAtIndex() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:shouldBreakLineByHyphenatingBeforeCharacterAtIndex:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(layoutManager ILayoutManager, charIndex uint) bool {
	rv := ffi.CallMethod[bool](l_, "layoutManager:shouldBreakLineByHyphenatingBeforeCharacterAtIndex:", layoutManager, charIndex)
	return rv
}

func (l_ *LayoutManagerDelegateWrapper) ImplementsLayoutManager_ShouldBreakLineByWordBeforeCharacterAtIndex() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:shouldBreakLineByWordBeforeCharacterAtIndex:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_ShouldBreakLineByWordBeforeCharacterAtIndex(layoutManager ILayoutManager, charIndex uint) bool {
	rv := ffi.CallMethod[bool](l_, "layoutManager:shouldBreakLineByWordBeforeCharacterAtIndex:", layoutManager, charIndex)
	return rv
}

func (l_ *LayoutManagerDelegateWrapper) ImplementsLayoutManager_LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:lineSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(layoutManager ILayoutManager, glyphIndex uint, rect foundation.Rect) float64 {
	rv := ffi.CallMethod[float64](l_, "layoutManager:lineSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:", layoutManager, glyphIndex, rect)
	return rv
}

func (l_ *LayoutManagerDelegateWrapper) ImplementsLayoutManager_ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:paragraphSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(layoutManager ILayoutManager, glyphIndex uint, rect foundation.Rect) float64 {
	rv := ffi.CallMethod[float64](l_, "layoutManager:paragraphSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:", layoutManager, glyphIndex, rect)
	return rv
}

func (l_ *LayoutManagerDelegateWrapper) ImplementsLayoutManager_ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:paragraphSpacingBeforeGlyphAtIndex:withProposedLineFragmentRect:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect(layoutManager ILayoutManager, glyphIndex uint, rect foundation.Rect) float64 {
	rv := ffi.CallMethod[float64](l_, "layoutManager:paragraphSpacingBeforeGlyphAtIndex:withProposedLineFragmentRect:", layoutManager, glyphIndex, rect)
	return rv
}

func (l_ *LayoutManagerDelegateWrapper) ImplementsLayoutManager_BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:boundingBoxForControlGlyphAtIndex:forTextContainer:proposedLineFragment:glyphPosition:characterIndex:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex(layoutManager ILayoutManager, glyphIndex uint, textContainer ITextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](l_, "layoutManager:boundingBoxForControlGlyphAtIndex:forTextContainer:proposedLineFragment:glyphPosition:characterIndex:", layoutManager, glyphIndex, textContainer, proposedRect, glyphPosition, charIndex)
	return rv
}

func (l_ *LayoutManagerDelegateWrapper) ImplementsLayoutManager_ShouldUseTemporaryAttributes_ForDrawingToScreen_AtCharacterIndex_EffectiveRange() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:shouldUseTemporaryAttributes:forDrawingToScreen:atCharacterIndex:effectiveRange:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManager_ShouldUseTemporaryAttributes_ForDrawingToScreen_AtCharacterIndex_EffectiveRange(layoutManager ILayoutManager, attrs map[foundation.AttributedStringKey]objc.IObject, toScreen bool, charIndex uint, effectiveCharRange *foundation.Range) map[foundation.AttributedStringKey]objc.Object {
	rv := ffi.CallMethod[map[foundation.AttributedStringKey]objc.Object](l_, "layoutManager:shouldUseTemporaryAttributes:forDrawingToScreen:atCharacterIndex:effectiveRange:", layoutManager, attrs, toScreen, charIndex, effectiveCharRange)
	return rv
}
