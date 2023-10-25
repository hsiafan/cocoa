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

type LayoutManagerDelegateCreator struct {
	className string
	class     objc.Class
}

func NewLayoutManagerDelegateCreator(name string) *LayoutManagerDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &LayoutManagerDelegateCreator{className: name, class: class}
}

func (c *LayoutManagerDelegateCreator) SetLayoutManagerDidInvalidateLayout(handle func(o objc.Object, sender LayoutManager)) {
	objc.AddMethod(c.class, objc.GetSelector("layoutManagerDidInvalidateLayout:"), handle)
}

func (c *LayoutManagerDelegateCreator) SetLayoutManager_ShouldGenerateGlyphs_Properties_CharacterIndexes_Font_ForGlyphRange(handle func(o objc.Object, layoutManager LayoutManager, glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont Font, glyphRange foundation.Range) uint) {
	objc.AddMethod(c.class, objc.GetSelector("layoutManager:shouldGenerateGlyphs:properties:characterIndexes:font:forGlyphRange:"), handle)
}

func (c *LayoutManagerDelegateCreator) SetLayoutManager_ShouldUseAction_ForControlCharacterAtIndex(handle func(o objc.Object, layoutManager LayoutManager, action ControlCharacterAction, charIndex uint) ControlCharacterAction) {
	objc.AddMethod(c.class, objc.GetSelector("layoutManager:shouldUseAction:forControlCharacterAtIndex:"), handle)
}

func (c *LayoutManagerDelegateCreator) SetLayoutManager_DidCompleteLayoutForTextContainer_AtEnd(handle func(o objc.Object, layoutManager LayoutManager, textContainer TextContainer, layoutFinishedFlag bool)) {
	objc.AddMethod(c.class, objc.GetSelector("layoutManager:didCompleteLayoutForTextContainer:atEnd:"), handle)
}

func (c *LayoutManagerDelegateCreator) SetLayoutManager_TextContainer_DidChangeGeometryFromSize(handle func(o objc.Object, layoutManager LayoutManager, textContainer TextContainer, oldSize foundation.Size)) {
	objc.AddMethod(c.class, objc.GetSelector("layoutManager:textContainer:didChangeGeometryFromSize:"), handle)
}

func (c *LayoutManagerDelegateCreator) SetLayoutManager_ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(handle func(o objc.Object, layoutManager LayoutManager, charIndex uint) bool) {
	objc.AddMethod(c.class, objc.GetSelector("layoutManager:shouldBreakLineByHyphenatingBeforeCharacterAtIndex:"), handle)
}

func (c *LayoutManagerDelegateCreator) SetLayoutManager_ShouldBreakLineByWordBeforeCharacterAtIndex(handle func(o objc.Object, layoutManager LayoutManager, charIndex uint) bool) {
	objc.AddMethod(c.class, objc.GetSelector("layoutManager:shouldBreakLineByWordBeforeCharacterAtIndex:"), handle)
}

func (c *LayoutManagerDelegateCreator) SetLayoutManager_LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(handle func(o objc.Object, layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64) {
	objc.AddMethod(c.class, objc.GetSelector("layoutManager:lineSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), handle)
}

func (c *LayoutManagerDelegateCreator) SetLayoutManager_ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(handle func(o objc.Object, layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64) {
	objc.AddMethod(c.class, objc.GetSelector("layoutManager:paragraphSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), handle)
}

func (c *LayoutManagerDelegateCreator) SetLayoutManager_ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect(handle func(o objc.Object, layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64) {
	objc.AddMethod(c.class, objc.GetSelector("layoutManager:paragraphSpacingBeforeGlyphAtIndex:withProposedLineFragmentRect:"), handle)
}

func (c *LayoutManagerDelegateCreator) SetLayoutManager_BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex(handle func(o objc.Object, layoutManager LayoutManager, glyphIndex uint, textContainer TextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect) {
	objc.AddMethod(c.class, objc.GetSelector("layoutManager:boundingBoxForControlGlyphAtIndex:forTextContainer:proposedLineFragment:glyphPosition:characterIndex:"), handle)
}

func (c *LayoutManagerDelegateCreator) SetLayoutManager_ShouldUseTemporaryAttributes_ForDrawingToScreen_AtCharacterIndex_EffectiveRange(handle func(o objc.Object, layoutManager LayoutManager, attrs map[foundation.AttributedStringKey]objc.Object, toScreen bool, charIndex uint, effectiveCharRange *foundation.Range) map[foundation.AttributedStringKey]objc.IObject) {
	objc.AddMethod(c.class, objc.GetSelector("layoutManager:shouldUseTemporaryAttributes:forDrawingToScreen:atCharacterIndex:effectiveRange:"), handle)
}

func (c *LayoutManagerDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
