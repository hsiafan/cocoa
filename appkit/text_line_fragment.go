// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TextLineFragmentClass = _TextLineFragmentClass{objc.GetClass("NSTextLineFragment")}

type _TextLineFragmentClass struct {
	objc.Class
}

type ITextLineFragment interface {
	objc.IObject
	CharacterIndexForPoint(point coregraphics.Point) int
	FractionOfDistanceThroughGlyphForPoint(point coregraphics.Point) float64
	LocationForCharacterAtIndex(index int) coregraphics.Point
	DrawAtPoint_InContext(point coregraphics.Point, context coregraphics.ContextRef)
	AttributedString() foundation.AttributedString
	CharacterRange() foundation.Range
	GlyphOrigin() coregraphics.Point
	TypographicBounds() coregraphics.Rect
}

type TextLineFragment struct {
	objc.Object
}

func MakeTextLineFragment(ptr unsafe.Pointer) TextLineFragment {
	return TextLineFragment{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextLineFragment) InitWithAttributedString_Range(attributedString foundation.IAttributedString, range_ foundation.Range) TextLineFragment {
	rv := ffi.CallMethod[TextLineFragment](t_, "initWithAttributedString:range:", attributedString, range_)
	return rv
}

func (t_ TextLineFragment) InitWithString_Attributes_Range(string_ string, attributes map[foundation.AttributedStringKey]objc.IObject, range_ foundation.Range) TextLineFragment {
	rv := ffi.CallMethod[TextLineFragment](t_, "initWithString:attributes:range:", string_, attributes, range_)
	return rv
}

func (tc _TextLineFragmentClass) Alloc() TextLineFragment {
	rv := ffi.CallMethod[TextLineFragment](tc, "alloc")
	return rv
}

func (tc _TextLineFragmentClass) New() TextLineFragment {
	rv := ffi.CallMethod[TextLineFragment](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextLineFragment() TextLineFragment {
	return TextLineFragmentClass.New()
}

func (t_ TextLineFragment) Init() TextLineFragment {
	rv := ffi.CallMethod[TextLineFragment](t_, "init")
	return rv
}

func (t_ TextLineFragment) CharacterIndexForPoint(point coregraphics.Point) int {
	rv := ffi.CallMethod[int](t_, "characterIndexForPoint:", point)
	return rv
}

func (t_ TextLineFragment) FractionOfDistanceThroughGlyphForPoint(point coregraphics.Point) float64 {
	rv := ffi.CallMethod[float64](t_, "fractionOfDistanceThroughGlyphForPoint:", point)
	return rv
}

func (t_ TextLineFragment) LocationForCharacterAtIndex(index int) coregraphics.Point {
	rv := ffi.CallMethod[coregraphics.Point](t_, "locationForCharacterAtIndex:", index)
	return rv
}

func (t_ TextLineFragment) DrawAtPoint_InContext(point coregraphics.Point, context coregraphics.ContextRef) {
	ffi.CallMethod[ffi.Void](t_, "drawAtPoint:inContext:", point, context)
}

func (t_ TextLineFragment) AttributedString() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](t_, "attributedString")
	return rv
}

func (t_ TextLineFragment) CharacterRange() foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "characterRange")
	return rv
}

func (t_ TextLineFragment) GlyphOrigin() coregraphics.Point {
	rv := ffi.CallMethod[coregraphics.Point](t_, "glyphOrigin")
	return rv
}

func (t_ TextLineFragment) TypographicBounds() coregraphics.Rect {
	rv := ffi.CallMethod[coregraphics.Rect](t_, "typographicBounds")
	return rv
}
