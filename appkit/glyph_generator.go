// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var GlyphGeneratorClass = _GlyphGeneratorClass{objc.GetClass("NSGlyphGenerator")}

type _GlyphGeneratorClass struct {
	objc.Class
}

type IGlyphGenerator interface {
	objc.IObject
	GenerateGlyphsForGlyphStorage_DesiredNumberOfCharacters_GlyphIndex_CharacterIndex(glyphStorage GlyphStorage, nChars uint, glyphIndex *uint, charIndex *uint)
	GenerateGlyphsForGlyphStorage0_DesiredNumberOfCharacters_GlyphIndex_CharacterIndex(glyphStorage objc.IObject, nChars uint, glyphIndex *uint, charIndex *uint)
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
	rv := ffi.CallMethod[GlyphGenerator](gc, "alloc")
	return rv
}

func (gc _GlyphGeneratorClass) New() GlyphGenerator {
	rv := ffi.CallMethod[GlyphGenerator](gc, "new")
	rv.Autorelease()
	return rv
}

func NewGlyphGenerator() GlyphGenerator {
	return GlyphGeneratorClass.New()
}

func (g_ GlyphGenerator) Init() GlyphGenerator {
	rv := ffi.CallMethod[GlyphGenerator](g_, "init")
	return rv
}

func (g_ GlyphGenerator) GenerateGlyphsForGlyphStorage_DesiredNumberOfCharacters_GlyphIndex_CharacterIndex(glyphStorage GlyphStorage, nChars uint, glyphIndex *uint, charIndex *uint) {
	po := ffi.CreateProtocol("NSGlyphStorage", glyphStorage)
	defer po.Release()
	ffi.CallMethod[ffi.Void](g_, "generateGlyphsForGlyphStorage:desiredNumberOfCharacters:glyphIndex:characterIndex:", po, nChars, glyphIndex, charIndex)
}

func (g_ GlyphGenerator) GenerateGlyphsForGlyphStorage0_DesiredNumberOfCharacters_GlyphIndex_CharacterIndex(glyphStorage objc.IObject, nChars uint, glyphIndex *uint, charIndex *uint) {
	ffi.CallMethod[ffi.Void](g_, "generateGlyphsForGlyphStorage:desiredNumberOfCharacters:glyphIndex:characterIndex:", glyphStorage, nChars, glyphIndex, charIndex)
}

func (gc _GlyphGeneratorClass) SharedGlyphGenerator() GlyphGenerator {
	rv := ffi.CallMethod[GlyphGenerator](gc, "sharedGlyphGenerator")
	return rv
}
