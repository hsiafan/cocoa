// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type GlyphStorage interface {
	// required
	AttributedString() foundation.IAttributedString
	// required
	LayoutOptions() uint
	// required
	InsertGlyphs_Length_ForStartingGlyphAtIndex_CharacterIndex(glyphs *Glyph, length uint, glyphIndex uint, charIndex uint)
	// required
	SetIntAttribute_Value_ForGlyphAtIndex(attributeTag int, val int, glyphIndex uint)
}

func WrapGlyphStorage(v GlyphStorage) objc.Object {
	return objc.WrapAsProtocol("NSGlyphStorage", v)
}

type GlyphStorageBase struct {
}

type GlyphStorageWrapper struct {
	objc.Object
}

func (g_ GlyphStorageWrapper) AttributedString() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](g_, objc.GetSelector("attributedString"))
	return rv
}

func (g_ GlyphStorageWrapper) LayoutOptions() uint {
	rv := objc.CallMethod[uint](g_, objc.GetSelector("layoutOptions"))
	return rv
}

func (g_ GlyphStorageWrapper) InsertGlyphs_Length_ForStartingGlyphAtIndex_CharacterIndex(glyphs *Glyph, length uint, glyphIndex uint, charIndex uint) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("insertGlyphs:length:forStartingGlyphAtIndex:characterIndex:"), glyphs, length, glyphIndex, charIndex)
}

func (g_ GlyphStorageWrapper) SetIntAttribute_Value_ForGlyphAtIndex(attributeTag int, val int, glyphIndex uint) {
	objc.CallMethod[objc.Void](g_, objc.GetSelector("setIntAttribute:value:forGlyphAtIndex:"), attributeTag, val, glyphIndex)
}
