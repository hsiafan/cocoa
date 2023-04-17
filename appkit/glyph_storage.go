// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
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

type GlyphStorageWrapper struct {
	objc.Object
}

func (g_ GlyphStorageWrapper) AttributedString() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](g_, "attributedString")
	return rv
}

func (g_ GlyphStorageWrapper) LayoutOptions() uint {
	rv := ffi.CallMethod[uint](g_, "layoutOptions")
	return rv
}

func (g_ GlyphStorageWrapper) InsertGlyphs_Length_ForStartingGlyphAtIndex_CharacterIndex(glyphs *Glyph, length uint, glyphIndex uint, charIndex uint) {
	ffi.CallMethod[ffi.Void](g_, "insertGlyphs:length:forStartingGlyphAtIndex:characterIndex:", glyphs, length, glyphIndex, charIndex)
}

func (g_ GlyphStorageWrapper) SetIntAttribute_Value_ForGlyphAtIndex(attributeTag int, val int, glyphIndex uint) {
	ffi.CallMethod[ffi.Void](g_, "setIntAttribute:value:forGlyphAtIndex:", attributeTag, val, glyphIndex)
}
