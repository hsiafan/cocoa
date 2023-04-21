// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var FontClass = _FontClass{objc.GetClass("NSFont")}

type _FontClass struct {
	objc.Class
}

type IFont interface {
	objc.IObject
	Set()
	SetInContext(graphicsContext IGraphicsContext)
	FontWithSize(fontSize float64) Font
	BoundingRectForCGGlyph(glyph coregraphics.Glyph) foundation.Rect
	GetBoundingRects_ForCGGlyphs_Count(bounds *foundation.Rect, glyphs *coregraphics.Glyph, glyphCount uint)
	AdvancementForCGGlyph(glyph coregraphics.Glyph) foundation.Size
	GetAdvancements_ForCGGlyphs_Count(advancements *foundation.Size, glyphs *coregraphics.Glyph, glyphCount uint)
	// deprecated
	GetAdvancements_ForGlyphs_Count(advancements *foundation.Size, glyphs *Glyph, glyphCount uint)
	// deprecated
	GetAdvancements_ForPackedGlyphs_Length(advancements *foundation.Size, packedGlyphs unsafe.Pointer, length uint)
	// deprecated
	AdvancementForGlyph(glyph Glyph) foundation.Size
	// deprecated
	BoundingRectForGlyph(glyph Glyph) foundation.Rect
	// deprecated
	GetBoundingRects_ForGlyphs_Count(bounds *foundation.Rect, glyphs *Glyph, glyphCount uint)
	// deprecated
	GlyphWithName(name string) Glyph
	// deprecated
	ScreenFontWithRenderingMode(renderingMode FontRenderingMode) Font
	PointSize() float64
	CoveredCharacterSet() foundation.CharacterSet
	FontDescriptor() FontDescriptor
	IsFixedPitch() bool
	MostCompatibleStringEncoding() foundation.StringEncoding
	NumberOfGlyphs() uint
	DisplayName() string
	FamilyName() string
	FontName() string
	IsVertical() bool
	VerticalFont() Font
	Ascender() float64
	Descender() float64
	CapHeight() float64
	Leading() float64
	XHeight() float64
	ItalicAngle() float64
	UnderlinePosition() float64
	UnderlineThickness() float64
	BoundingRectForFont() foundation.Rect
	MaximumAdvancement() foundation.Size
	Matrix() *float64
	TextTransform() foundation.AffineTransform
	// deprecated
	RenderingMode() FontRenderingMode
	// deprecated
	PrinterFont() Font
	// deprecated
	ScreenFont() Font
}

type Font struct {
	objc.Object
}

func MakeFont(ptr unsafe.Pointer) Font {
	return Font{
		Object: objc.MakeObject(ptr),
	}
}

func (fc _FontClass) Alloc() Font {
	rv := objc.CallMethod[Font](fc, "alloc")
	return rv
}

func (fc _FontClass) New() Font {
	rv := objc.CallMethod[Font](fc, "new")
	rv.Autorelease()
	return rv
}

func NewFont() Font {
	return FontClass.New()
}

func (f_ Font) Init() Font {
	rv := objc.CallMethod[Font](f_, "init")
	return rv
}

func (fc _FontClass) FontWithName_Size(fontName string, fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, "fontWithName:size:", fontName, fontSize)
	return rv
}

func (fc _FontClass) FontWithDescriptor_Size(fontDescriptor IFontDescriptor, fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, "fontWithDescriptor:size:", fontDescriptor, fontSize)
	return rv
}

func (fc _FontClass) FontWithDescriptor_TextTransform(fontDescriptor IFontDescriptor, textTransform foundation.IAffineTransform) Font {
	rv := objc.CallMethod[Font](fc, "fontWithDescriptor:textTransform:", fontDescriptor, textTransform)
	return rv
}

func (fc _FontClass) FontWithName_Matrix(fontName string, fontMatrix *float64) Font {
	rv := objc.CallMethod[Font](fc, "fontWithName:matrix:", fontName, fontMatrix)
	return rv
}

func (fc _FontClass) UserFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, "userFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) UserFixedPitchFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, "userFixedPitchFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) PreferredFontForTextStyle_Options(style FontTextStyle, options map[FontTextStyleOptionKey]objc.IObject) Font {
	rv := objc.CallMethod[Font](fc, "preferredFontForTextStyle:options:", style, options)
	return rv
}

func (fc _FontClass) SystemFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, "systemFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) SystemFontOfSize_Weight(fontSize float64, weight FontWeight) Font {
	rv := objc.CallMethod[Font](fc, "systemFontOfSize:weight:", fontSize, weight)
	return rv
}

func (fc _FontClass) BoldSystemFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, "boldSystemFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) MonospacedSystemFontOfSize_Weight(fontSize float64, weight FontWeight) Font {
	rv := objc.CallMethod[Font](fc, "monospacedSystemFontOfSize:weight:", fontSize, weight)
	return rv
}

func (fc _FontClass) MonospacedDigitSystemFontOfSize_Weight(fontSize float64, weight FontWeight) Font {
	rv := objc.CallMethod[Font](fc, "monospacedDigitSystemFontOfSize:weight:", fontSize, weight)
	return rv
}

func (fc _FontClass) LabelFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, "labelFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) MessageFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, "messageFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) MenuBarFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, "menuBarFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) MenuFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, "menuFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) ControlContentFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, "controlContentFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) TitleBarFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, "titleBarFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) PaletteFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, "paletteFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) ToolTipsFontOfSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](fc, "toolTipsFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) SystemFontSizeForControlSize(controlSize ControlSize) float64 {
	rv := objc.CallMethod[float64](fc, "systemFontSizeForControlSize:", controlSize)
	return rv
}

func (f_ Font) Set() {
	objc.CallMethod[objc.Void](f_, "set")
}

func (f_ Font) SetInContext(graphicsContext IGraphicsContext) {
	objc.CallMethod[objc.Void](f_, "setInContext:", graphicsContext)
}

func (fc _FontClass) SetUserFont(font IFont) {
	objc.CallMethod[objc.Void](fc, "setUserFont:", font)
}

func (fc _FontClass) SetUserFixedPitchFont(font IFont) {
	objc.CallMethod[objc.Void](fc, "setUserFixedPitchFont:", font)
}

func (f_ Font) FontWithSize(fontSize float64) Font {
	rv := objc.CallMethod[Font](f_, "fontWithSize:", fontSize)
	return rv
}

func (fc _FontClass) SystemFontOfSize_Weight_Width(fontSize float64, weight FontWeight, width FontWidth) Font {
	rv := objc.CallMethod[Font](fc, "systemFontOfSize:weight:width:", fontSize, weight, width)
	return rv
}

func (f_ Font) BoundingRectForCGGlyph(glyph coregraphics.Glyph) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](f_, "boundingRectForCGGlyph:", glyph)
	return rv
}

func (f_ Font) GetBoundingRects_ForCGGlyphs_Count(bounds *foundation.Rect, glyphs *coregraphics.Glyph, glyphCount uint) {
	objc.CallMethod[objc.Void](f_, "getBoundingRects:forCGGlyphs:count:", bounds, glyphs, glyphCount)
}

func (f_ Font) AdvancementForCGGlyph(glyph coregraphics.Glyph) foundation.Size {
	rv := objc.CallMethod[foundation.Size](f_, "advancementForCGGlyph:", glyph)
	return rv
}

func (f_ Font) GetAdvancements_ForCGGlyphs_Count(advancements *foundation.Size, glyphs *coregraphics.Glyph, glyphCount uint) {
	objc.CallMethod[objc.Void](f_, "getAdvancements:forCGGlyphs:count:", advancements, glyphs, glyphCount)
}

// deprecated
func (f_ Font) GetAdvancements_ForGlyphs_Count(advancements *foundation.Size, glyphs *Glyph, glyphCount uint) {
	objc.CallMethod[objc.Void](f_, "getAdvancements:forGlyphs:count:", advancements, glyphs, glyphCount)
}

// deprecated
func (f_ Font) GetAdvancements_ForPackedGlyphs_Length(advancements *foundation.Size, packedGlyphs unsafe.Pointer, length uint) {
	objc.CallMethod[objc.Void](f_, "getAdvancements:forPackedGlyphs:length:", advancements, packedGlyphs, length)
}

// deprecated
func (f_ Font) AdvancementForGlyph(glyph Glyph) foundation.Size {
	rv := objc.CallMethod[foundation.Size](f_, "advancementForGlyph:", glyph)
	return rv
}

// deprecated
func (f_ Font) BoundingRectForGlyph(glyph Glyph) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](f_, "boundingRectForGlyph:", glyph)
	return rv
}

// deprecated
func (f_ Font) GetBoundingRects_ForGlyphs_Count(bounds *foundation.Rect, glyphs *Glyph, glyphCount uint) {
	objc.CallMethod[objc.Void](f_, "getBoundingRects:forGlyphs:count:", bounds, glyphs, glyphCount)
}

// deprecated
func (f_ Font) GlyphWithName(name string) Glyph {
	rv := objc.CallMethod[Glyph](f_, "glyphWithName:", name)
	return rv
}

// deprecated
func (f_ Font) ScreenFontWithRenderingMode(renderingMode FontRenderingMode) Font {
	rv := objc.CallMethod[Font](f_, "screenFontWithRenderingMode:", renderingMode)
	return rv
}

func (fc _FontClass) SystemFontSize() float64 {
	rv := objc.CallMethod[float64](fc, "systemFontSize")
	return rv
}

func (fc _FontClass) SmallSystemFontSize() float64 {
	rv := objc.CallMethod[float64](fc, "smallSystemFontSize")
	return rv
}

func (fc _FontClass) LabelFontSize() float64 {
	rv := objc.CallMethod[float64](fc, "labelFontSize")
	return rv
}

func (f_ Font) PointSize() float64 {
	rv := objc.CallMethod[float64](f_, "pointSize")
	return rv
}

func (f_ Font) CoveredCharacterSet() foundation.CharacterSet {
	rv := objc.CallMethod[foundation.CharacterSet](f_, "coveredCharacterSet")
	return rv
}

func (f_ Font) FontDescriptor() FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, "fontDescriptor")
	return rv
}

func (f_ Font) IsFixedPitch() bool {
	rv := objc.CallMethod[bool](f_, "isFixedPitch")
	return rv
}

func (f_ Font) MostCompatibleStringEncoding() foundation.StringEncoding {
	rv := objc.CallMethod[foundation.StringEncoding](f_, "mostCompatibleStringEncoding")
	return rv
}

func (f_ Font) NumberOfGlyphs() uint {
	rv := objc.CallMethod[uint](f_, "numberOfGlyphs")
	return rv
}

func (f_ Font) DisplayName() string {
	rv := objc.CallMethod[string](f_, "displayName")
	return rv
}

func (f_ Font) FamilyName() string {
	rv := objc.CallMethod[string](f_, "familyName")
	return rv
}

func (f_ Font) FontName() string {
	rv := objc.CallMethod[string](f_, "fontName")
	return rv
}

func (f_ Font) IsVertical() bool {
	rv := objc.CallMethod[bool](f_, "isVertical")
	return rv
}

func (f_ Font) VerticalFont() Font {
	rv := objc.CallMethod[Font](f_, "verticalFont")
	return rv
}

func (f_ Font) Ascender() float64 {
	rv := objc.CallMethod[float64](f_, "ascender")
	return rv
}

func (f_ Font) Descender() float64 {
	rv := objc.CallMethod[float64](f_, "descender")
	return rv
}

func (f_ Font) CapHeight() float64 {
	rv := objc.CallMethod[float64](f_, "capHeight")
	return rv
}

func (f_ Font) Leading() float64 {
	rv := objc.CallMethod[float64](f_, "leading")
	return rv
}

func (f_ Font) XHeight() float64 {
	rv := objc.CallMethod[float64](f_, "xHeight")
	return rv
}

func (f_ Font) ItalicAngle() float64 {
	rv := objc.CallMethod[float64](f_, "italicAngle")
	return rv
}

func (f_ Font) UnderlinePosition() float64 {
	rv := objc.CallMethod[float64](f_, "underlinePosition")
	return rv
}

func (f_ Font) UnderlineThickness() float64 {
	rv := objc.CallMethod[float64](f_, "underlineThickness")
	return rv
}

func (f_ Font) BoundingRectForFont() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](f_, "boundingRectForFont")
	return rv
}

func (f_ Font) MaximumAdvancement() foundation.Size {
	rv := objc.CallMethod[foundation.Size](f_, "maximumAdvancement")
	return rv
}

func (f_ Font) Matrix() *float64 {
	rv := objc.CallMethod[*float64](f_, "matrix")
	return rv
}

func (f_ Font) TextTransform() foundation.AffineTransform {
	rv := objc.CallMethod[foundation.AffineTransform](f_, "textTransform")
	return rv
}

// deprecated
func (f_ Font) RenderingMode() FontRenderingMode {
	rv := objc.CallMethod[FontRenderingMode](f_, "renderingMode")
	return rv
}

// deprecated
func (f_ Font) PrinterFont() Font {
	rv := objc.CallMethod[Font](f_, "printerFont")
	return rv
}

// deprecated
func (f_ Font) ScreenFont() Font {
	rv := objc.CallMethod[Font](f_, "screenFont")
	return rv
}
