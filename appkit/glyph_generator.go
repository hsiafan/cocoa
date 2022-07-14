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

func (g_ GlyphGenerator) Init() GlyphGenerator {
	rv := ffi.CallMethod[GlyphGenerator](g_, "init")
	rv.Autorelease()
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

func (gc _GlyphGeneratorClass) SharedGlyphGenerator() GlyphGenerator {
	rv := ffi.CallMethod[GlyphGenerator](gc, "sharedGlyphGenerator")
	return rv
}
