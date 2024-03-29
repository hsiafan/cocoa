// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var GlyphGeneratorClass = _GlyphGeneratorClass{objc.GetClass("NSGlyphGenerator")}

type _GlyphGeneratorClass struct {
	objc.Class
}

type IGlyphGenerator interface {
	objc.IObject
	GenerateGlyphsForGlyphStorage_DesiredNumberOfCharacters_GlyphIndex_CharacterIndex(glyphStorage objc.IObject, nChars uint, glyphIndex *uint, charIndex *uint)
}

type GlyphGenerator struct {
	objc.Object
}

func MakeGlyphGenerator(ptr unsafe.Pointer) GlyphGenerator {
	return GlyphGenerator{
		Object: objc.MakeObject(ptr),
	}
}

func (gc _GlyphGeneratorClass) Alloc() GlyphGenerator {
	rv := objc.CallMethod[GlyphGenerator](gc, objc.SEL("alloc"))
	return rv
}

func (gc _GlyphGeneratorClass) New() GlyphGenerator {
	rv := objc.CallMethod[GlyphGenerator](gc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewGlyphGenerator() GlyphGenerator {
	return GlyphGeneratorClass.New()
}

func (g_ GlyphGenerator) Init() GlyphGenerator {
	rv := objc.CallMethod[GlyphGenerator](g_, objc.SEL("init"))
	return rv
}

func (g_ GlyphGenerator) GenerateGlyphsForGlyphStorage_DesiredNumberOfCharacters_GlyphIndex_CharacterIndex(glyphStorage objc.IObject, nChars uint, glyphIndex *uint, charIndex *uint) {
	objc.CallMethod[objc.Void](g_, objc.SEL("generateGlyphsForGlyphStorage:desiredNumberOfCharacters:glyphIndex:characterIndex:"), objc.ExtractPtr(glyphStorage), nChars, glyphIndex, charIndex)
}

func (gc _GlyphGeneratorClass) SharedGlyphGenerator() GlyphGenerator {
	rv := objc.CallMethod[GlyphGenerator](gc, objc.SEL("sharedGlyphGenerator"))
	return rv
}
