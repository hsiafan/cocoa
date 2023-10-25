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
