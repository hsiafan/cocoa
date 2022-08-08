// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[Font](fc, "alloc")
	return rv
}

func (f_ Font) Init() Font {
	rv := ffi.CallMethod[Font](f_, "init")
	return rv
}

func (fc _FontClass) New() Font {
	rv := ffi.CallMethod[Font](fc, "new")
	rv.Autorelease()
	return rv
}

func NewFont() Font {
	return FontClass.New()
}

func (fc _FontClass) FontWithName_Size(fontName string, fontSize float64) Font {
	rv := ffi.CallMethod[Font](fc, "fontWithName:size:", fontName, fontSize)
	return rv
}

func (fc _FontClass) FontWithDescriptor_Size(fontDescriptor IFontDescriptor, fontSize float64) Font {
	rv := ffi.CallMethod[Font](fc, "fontWithDescriptor:size:", fontDescriptor, fontSize)
	return rv
}

func (fc _FontClass) FontWithDescriptor_TextTransform(fontDescriptor IFontDescriptor, textTransform foundation.IAffineTransform) Font {
	rv := ffi.CallMethod[Font](fc, "fontWithDescriptor:textTransform:", fontDescriptor, textTransform)
	return rv
}

func (fc _FontClass) FontWithName_Matrix(fontName string, fontMatrix *float64) Font {
	rv := ffi.CallMethod[Font](fc, "fontWithName:matrix:", fontName, fontMatrix)
	return rv
}

func (fc _FontClass) UserFontOfSize(fontSize float64) Font {
	rv := ffi.CallMethod[Font](fc, "userFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) UserFixedPitchFontOfSize(fontSize float64) Font {
	rv := ffi.CallMethod[Font](fc, "userFixedPitchFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) PreferredFontForTextStyle_Options(style FontTextStyle, options map[FontTextStyleOptionKey]objc.IObject) Font {
	rv := ffi.CallMethod[Font](fc, "preferredFontForTextStyle:options:", style, options)
	return rv
}

func (fc _FontClass) SystemFontOfSize(fontSize float64) Font {
	rv := ffi.CallMethod[Font](fc, "systemFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) SystemFontOfSize_Weight(fontSize float64, weight FontWeight) Font {
	rv := ffi.CallMethod[Font](fc, "systemFontOfSize:weight:", fontSize, weight)
	return rv
}

func (fc _FontClass) BoldSystemFontOfSize(fontSize float64) Font {
	rv := ffi.CallMethod[Font](fc, "boldSystemFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) MonospacedSystemFontOfSize_Weight(fontSize float64, weight FontWeight) Font {
	rv := ffi.CallMethod[Font](fc, "monospacedSystemFontOfSize:weight:", fontSize, weight)
	return rv
}

func (fc _FontClass) MonospacedDigitSystemFontOfSize_Weight(fontSize float64, weight FontWeight) Font {
	rv := ffi.CallMethod[Font](fc, "monospacedDigitSystemFontOfSize:weight:", fontSize, weight)
	return rv
}

func (fc _FontClass) LabelFontOfSize(fontSize float64) Font {
	rv := ffi.CallMethod[Font](fc, "labelFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) MessageFontOfSize(fontSize float64) Font {
	rv := ffi.CallMethod[Font](fc, "messageFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) MenuBarFontOfSize(fontSize float64) Font {
	rv := ffi.CallMethod[Font](fc, "menuBarFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) MenuFontOfSize(fontSize float64) Font {
	rv := ffi.CallMethod[Font](fc, "menuFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) ControlContentFontOfSize(fontSize float64) Font {
	rv := ffi.CallMethod[Font](fc, "controlContentFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) TitleBarFontOfSize(fontSize float64) Font {
	rv := ffi.CallMethod[Font](fc, "titleBarFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) PaletteFontOfSize(fontSize float64) Font {
	rv := ffi.CallMethod[Font](fc, "paletteFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) ToolTipsFontOfSize(fontSize float64) Font {
	rv := ffi.CallMethod[Font](fc, "toolTipsFontOfSize:", fontSize)
	return rv
}

func (fc _FontClass) SystemFontSizeForControlSize(controlSize ControlSize) float64 {
	rv := ffi.CallMethod[float64](fc, "systemFontSizeForControlSize:", controlSize)
	return rv
}

func (f_ Font) Set() {
	ffi.CallMethod[ffi.Void](f_, "set")
}

func (f_ Font) SetInContext(graphicsContext IGraphicsContext) {
	ffi.CallMethod[ffi.Void](f_, "setInContext:", graphicsContext)
}

func (fc _FontClass) SetUserFont(font IFont) {
	ffi.CallMethod[ffi.Void](fc, "setUserFont:", font)
}

func (fc _FontClass) SetUserFixedPitchFont(font IFont) {
	ffi.CallMethod[ffi.Void](fc, "setUserFixedPitchFont:", font)
}

func (f_ Font) FontWithSize(fontSize float64) Font {
	rv := ffi.CallMethod[Font](f_, "fontWithSize:", fontSize)
	return rv
}

func (fc _FontClass) SystemFontOfSize_Weight_Width(fontSize float64, weight FontWeight, width FontWidth) Font {
	rv := ffi.CallMethod[Font](fc, "systemFontOfSize:weight:width:", fontSize, weight, width)
	return rv
}

func (f_ Font) BoundingRectForCGGlyph(glyph coregraphics.Glyph) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](f_, "boundingRectForCGGlyph:", glyph)
	return rv
}

func (f_ Font) GetBoundingRects_ForCGGlyphs_Count(bounds *foundation.Rect, glyphs *coregraphics.Glyph, glyphCount uint) {
	ffi.CallMethod[ffi.Void](f_, "getBoundingRects:forCGGlyphs:count:", bounds, glyphs, glyphCount)
}

func (f_ Font) AdvancementForCGGlyph(glyph coregraphics.Glyph) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](f_, "advancementForCGGlyph:", glyph)
	return rv
}

func (f_ Font) GetAdvancements_ForCGGlyphs_Count(advancements *foundation.Size, glyphs *coregraphics.Glyph, glyphCount uint) {
	ffi.CallMethod[ffi.Void](f_, "getAdvancements:forCGGlyphs:count:", advancements, glyphs, glyphCount)
}

// deprecated
func (f_ Font) GetAdvancements_ForGlyphs_Count(advancements *foundation.Size, glyphs *Glyph, glyphCount uint) {
	ffi.CallMethod[ffi.Void](f_, "getAdvancements:forGlyphs:count:", advancements, glyphs, glyphCount)
}

// deprecated
func (f_ Font) GetAdvancements_ForPackedGlyphs_Length(advancements *foundation.Size, packedGlyphs unsafe.Pointer, length uint) {
	ffi.CallMethod[ffi.Void](f_, "getAdvancements:forPackedGlyphs:length:", advancements, packedGlyphs, length)
}

// deprecated
func (f_ Font) AdvancementForGlyph(glyph Glyph) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](f_, "advancementForGlyph:", glyph)
	return rv
}

// deprecated
func (f_ Font) BoundingRectForGlyph(glyph Glyph) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](f_, "boundingRectForGlyph:", glyph)
	return rv
}

// deprecated
func (f_ Font) GetBoundingRects_ForGlyphs_Count(bounds *foundation.Rect, glyphs *Glyph, glyphCount uint) {
	ffi.CallMethod[ffi.Void](f_, "getBoundingRects:forGlyphs:count:", bounds, glyphs, glyphCount)
}

// deprecated
func (f_ Font) GlyphWithName(name string) Glyph {
	rv := ffi.CallMethod[Glyph](f_, "glyphWithName:", name)
	return rv
}

// deprecated
func (f_ Font) ScreenFontWithRenderingMode(renderingMode FontRenderingMode) Font {
	rv := ffi.CallMethod[Font](f_, "screenFontWithRenderingMode:", renderingMode)
	return rv
}

func (fc _FontClass) SystemFontSize() float64 {
	rv := ffi.CallMethod[float64](fc, "systemFontSize")
	return rv
}

func (fc _FontClass) SmallSystemFontSize() float64 {
	rv := ffi.CallMethod[float64](fc, "smallSystemFontSize")
	return rv
}

func (fc _FontClass) LabelFontSize() float64 {
	rv := ffi.CallMethod[float64](fc, "labelFontSize")
	return rv
}

func (f_ Font) PointSize() float64 {
	rv := ffi.CallMethod[float64](f_, "pointSize")
	return rv
}

func (f_ Font) CoveredCharacterSet() foundation.CharacterSet {
	rv := ffi.CallMethod[foundation.CharacterSet](f_, "coveredCharacterSet")
	return rv
}

func (f_ Font) FontDescriptor() FontDescriptor {
	rv := ffi.CallMethod[FontDescriptor](f_, "fontDescriptor")
	return rv
}

func (f_ Font) IsFixedPitch() bool {
	rv := ffi.CallMethod[bool](f_, "isFixedPitch")
	return rv
}

func (f_ Font) MostCompatibleStringEncoding() foundation.StringEncoding {
	rv := ffi.CallMethod[foundation.StringEncoding](f_, "mostCompatibleStringEncoding")
	return rv
}

func (f_ Font) NumberOfGlyphs() uint {
	rv := ffi.CallMethod[uint](f_, "numberOfGlyphs")
	return rv
}

func (f_ Font) DisplayName() string {
	rv := ffi.CallMethod[string](f_, "displayName")
	return rv
}

func (f_ Font) FamilyName() string {
	rv := ffi.CallMethod[string](f_, "familyName")
	return rv
}

func (f_ Font) FontName() string {
	rv := ffi.CallMethod[string](f_, "fontName")
	return rv
}

func (f_ Font) IsVertical() bool {
	rv := ffi.CallMethod[bool](f_, "isVertical")
	return rv
}

func (f_ Font) VerticalFont() Font {
	rv := ffi.CallMethod[Font](f_, "verticalFont")
	return rv
}

func (f_ Font) Ascender() float64 {
	rv := ffi.CallMethod[float64](f_, "ascender")
	return rv
}

func (f_ Font) Descender() float64 {
	rv := ffi.CallMethod[float64](f_, "descender")
	return rv
}

func (f_ Font) CapHeight() float64 {
	rv := ffi.CallMethod[float64](f_, "capHeight")
	return rv
}

func (f_ Font) Leading() float64 {
	rv := ffi.CallMethod[float64](f_, "leading")
	return rv
}

func (f_ Font) XHeight() float64 {
	rv := ffi.CallMethod[float64](f_, "xHeight")
	return rv
}

func (f_ Font) ItalicAngle() float64 {
	rv := ffi.CallMethod[float64](f_, "italicAngle")
	return rv
}

func (f_ Font) UnderlinePosition() float64 {
	rv := ffi.CallMethod[float64](f_, "underlinePosition")
	return rv
}

func (f_ Font) UnderlineThickness() float64 {
	rv := ffi.CallMethod[float64](f_, "underlineThickness")
	return rv
}

func (f_ Font) BoundingRectForFont() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](f_, "boundingRectForFont")
	return rv
}

func (f_ Font) MaximumAdvancement() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](f_, "maximumAdvancement")
	return rv
}

func (f_ Font) Matrix() *float64 {
	rv := ffi.CallMethod[*float64](f_, "matrix")
	return rv
}

func (f_ Font) TextTransform() foundation.AffineTransform {
	rv := ffi.CallMethod[foundation.AffineTransform](f_, "textTransform")
	return rv
}

// deprecated
func (f_ Font) RenderingMode() FontRenderingMode {
	rv := ffi.CallMethod[FontRenderingMode](f_, "renderingMode")
	return rv
}

// deprecated
func (f_ Font) PrinterFont() Font {
	rv := ffi.CallMethod[Font](f_, "printerFont")
	return rv
}

// deprecated
func (f_ Font) ScreenFont() Font {
	rv := ffi.CallMethod[Font](f_, "screenFont")
	return rv
}

var FontDescriptorClass = _FontDescriptorClass{objc.GetClass("NSFontDescriptor")}

type _FontDescriptorClass struct {
	objc.Class
}

type IFontDescriptor interface {
	objc.IObject
	FontDescriptorByAddingAttributes(attributes map[FontDescriptorAttributeName]objc.IObject) FontDescriptor
	FontDescriptorWithFace(newFace string) FontDescriptor
	FontDescriptorWithFamily(newFamily string) FontDescriptor
	FontDescriptorWithMatrix(matrix foundation.IAffineTransform) FontDescriptor
	FontDescriptorWithSize(newPointSize float64) FontDescriptor
	FontDescriptorWithSymbolicTraits(symbolicTraits FontDescriptorSymbolicTraits) FontDescriptor
	MatchingFontDescriptorsWithMandatoryKeys(mandatoryKeys foundation.ISet) []FontDescriptor
	MatchingFontDescriptorWithMandatoryKeys(mandatoryKeys foundation.ISet) FontDescriptor
	ObjectForKey(attribute FontDescriptorAttributeName) objc.Object
	FontAttributes() map[FontDescriptorAttributeName]objc.Object
	Matrix() foundation.AffineTransform
	PointSize() float64
	PostscriptName() string
	SymbolicTraits() FontDescriptorSymbolicTraits
	RequiresFontAssetRequest() bool
}

type FontDescriptor struct {
	objc.Object
}

func MakeFontDescriptor(ptr unsafe.Pointer) FontDescriptor {
	return FontDescriptor{
		Object: objc.MakeObject(ptr),
	}
}

func (f_ FontDescriptor) InitWithFontAttributes(attributes map[FontDescriptorAttributeName]objc.IObject) FontDescriptor {
	rv := ffi.CallMethod[FontDescriptor](f_, "initWithFontAttributes:", attributes)
	return rv
}

func (f_ FontDescriptor) FontDescriptorWithDesign(design FontDescriptorSystemDesign) FontDescriptor {
	rv := ffi.CallMethod[FontDescriptor](f_, "fontDescriptorWithDesign:", design)
	return rv
}

func (fc _FontDescriptorClass) Alloc() FontDescriptor {
	rv := ffi.CallMethod[FontDescriptor](fc, "alloc")
	return rv
}

func (f_ FontDescriptor) Init() FontDescriptor {
	rv := ffi.CallMethod[FontDescriptor](f_, "init")
	return rv
}

func (fc _FontDescriptorClass) New() FontDescriptor {
	rv := ffi.CallMethod[FontDescriptor](fc, "new")
	rv.Autorelease()
	return rv
}

func NewFontDescriptor() FontDescriptor {
	return FontDescriptorClass.New()
}

func (fc _FontDescriptorClass) PreferredFontDescriptorForTextStyle_Options(style FontTextStyle, options map[FontTextStyleOptionKey]objc.IObject) FontDescriptor {
	rv := ffi.CallMethod[FontDescriptor](fc, "preferredFontDescriptorForTextStyle:options:", style, options)
	return rv
}

func (fc _FontDescriptorClass) FontDescriptorWithFontAttributes(attributes map[FontDescriptorAttributeName]objc.IObject) FontDescriptor {
	rv := ffi.CallMethod[FontDescriptor](fc, "fontDescriptorWithFontAttributes:", attributes)
	return rv
}

func (fc _FontDescriptorClass) FontDescriptorWithName_Matrix(fontName string, matrix foundation.IAffineTransform) FontDescriptor {
	rv := ffi.CallMethod[FontDescriptor](fc, "fontDescriptorWithName:matrix:", fontName, matrix)
	return rv
}

func (fc _FontDescriptorClass) FontDescriptorWithName_Size(fontName string, size float64) FontDescriptor {
	rv := ffi.CallMethod[FontDescriptor](fc, "fontDescriptorWithName:size:", fontName, size)
	return rv
}

func (f_ FontDescriptor) FontDescriptorByAddingAttributes(attributes map[FontDescriptorAttributeName]objc.IObject) FontDescriptor {
	rv := ffi.CallMethod[FontDescriptor](f_, "fontDescriptorByAddingAttributes:", attributes)
	return rv
}

func (f_ FontDescriptor) FontDescriptorWithFace(newFace string) FontDescriptor {
	rv := ffi.CallMethod[FontDescriptor](f_, "fontDescriptorWithFace:", newFace)
	return rv
}

func (f_ FontDescriptor) FontDescriptorWithFamily(newFamily string) FontDescriptor {
	rv := ffi.CallMethod[FontDescriptor](f_, "fontDescriptorWithFamily:", newFamily)
	return rv
}

func (f_ FontDescriptor) FontDescriptorWithMatrix(matrix foundation.IAffineTransform) FontDescriptor {
	rv := ffi.CallMethod[FontDescriptor](f_, "fontDescriptorWithMatrix:", matrix)
	return rv
}

func (f_ FontDescriptor) FontDescriptorWithSize(newPointSize float64) FontDescriptor {
	rv := ffi.CallMethod[FontDescriptor](f_, "fontDescriptorWithSize:", newPointSize)
	return rv
}

func (f_ FontDescriptor) FontDescriptorWithSymbolicTraits(symbolicTraits FontDescriptorSymbolicTraits) FontDescriptor {
	rv := ffi.CallMethod[FontDescriptor](f_, "fontDescriptorWithSymbolicTraits:", symbolicTraits)
	return rv
}

func (f_ FontDescriptor) MatchingFontDescriptorsWithMandatoryKeys(mandatoryKeys foundation.ISet) []FontDescriptor {
	rv := ffi.CallMethod[[]FontDescriptor](f_, "matchingFontDescriptorsWithMandatoryKeys:", mandatoryKeys)
	return rv
}

func (f_ FontDescriptor) MatchingFontDescriptorWithMandatoryKeys(mandatoryKeys foundation.ISet) FontDescriptor {
	rv := ffi.CallMethod[FontDescriptor](f_, "matchingFontDescriptorWithMandatoryKeys:", mandatoryKeys)
	return rv
}

func (f_ FontDescriptor) ObjectForKey(attribute FontDescriptorAttributeName) objc.Object {
	rv := ffi.CallMethod[objc.Object](f_, "objectForKey:", attribute)
	return rv
}

func (f_ FontDescriptor) FontAttributes() map[FontDescriptorAttributeName]objc.Object {
	rv := ffi.CallMethod[map[FontDescriptorAttributeName]objc.Object](f_, "fontAttributes")
	return rv
}

func (f_ FontDescriptor) Matrix() foundation.AffineTransform {
	rv := ffi.CallMethod[foundation.AffineTransform](f_, "matrix")
	return rv
}

func (f_ FontDescriptor) PointSize() float64 {
	rv := ffi.CallMethod[float64](f_, "pointSize")
	return rv
}

func (f_ FontDescriptor) PostscriptName() string {
	rv := ffi.CallMethod[string](f_, "postscriptName")
	return rv
}

func (f_ FontDescriptor) SymbolicTraits() FontDescriptorSymbolicTraits {
	rv := ffi.CallMethod[FontDescriptorSymbolicTraits](f_, "symbolicTraits")
	return rv
}

func (f_ FontDescriptor) RequiresFontAssetRequest() bool {
	rv := ffi.CallMethod[bool](f_, "requiresFontAssetRequest")
	return rv
}

var FontManagerClass = _FontManagerClass{objc.GetClass("NSFontManager")}

type _FontManagerClass struct {
	objc.Class
}

type IFontManager interface {
	objc.IObject
	AvailableFontNamesWithTraits(someTraits FontTraitMask) []string
	AvailableMembersOfFontFamily(fam string) [][]objc.Object
	SetSelectedFont_IsMultiple(fontObj IFont, flag bool)
	SendAction() bool
	LocalizedNameForFamily_Face(family string, faceKey string) string
	AddFontTrait(sender objc.IObject)
	RemoveFontTrait(sender objc.IObject)
	ModifyFont(sender objc.IObject)
	ModifyFontViaPanel(sender objc.IObject)
	OrderFrontStylesPanel(sender objc.IObject)
	OrderFrontFontPanel(sender objc.IObject)
	ConvertFont(fontObj IFont) Font
	// deprecated
	ChangeFont(sender objc.IObject)
	ConvertFont_ToFace(fontObj IFont, typeface string) Font
	ConvertFont_ToFamily(fontObj IFont, family string) Font
	ConvertFont_ToHaveTrait(fontObj IFont, trait FontTraitMask) Font
	ConvertFont_ToNotHaveTrait(fontObj IFont, trait FontTraitMask) Font
	ConvertFont_ToSize(fontObj IFont, size float64) Font
	ConvertWeight_OfFont(upFlag bool, fontObj IFont) Font
	ConvertFontTraits(traits FontTraitMask) FontTraitMask
	FontWithFamily_Traits_Weight_Size(family string, traits FontTraitMask, weight int, size float64) Font
	TraitsOfFont(fontObj IFont) FontTraitMask
	FontNamed_HasTraits(fName string, someTraits FontTraitMask) bool
	WeightOfFont(fontObj IFont) int
	FontPanel(create bool) FontPanel
	SetFontMenu(newMenu IMenu)
	FontMenu(create bool) Menu
	SetSelectedAttributes_IsMultiple(attributes map[string]objc.IObject, flag bool)
	ConvertAttributes(attributes map[string]objc.IObject) map[string]objc.Object
	// deprecated
	AvailableFontNamesMatchingFontDescriptor(descriptor IFontDescriptor) []objc.Object
	// deprecated
	FontDescriptorsInCollection(collectionNames string) []objc.Object
	// deprecated
	AddCollection_Options(collectionName string, collectionOptions FontCollectionOptions) bool
	// deprecated
	RemoveCollection(collectionName string) bool
	// deprecated
	AddFontDescriptors_ToCollection(descriptors []objc.IObject, collectionName string)
	// deprecated
	RemoveFontDescriptor_FromCollection(descriptor IFontDescriptor, collection string)
	// deprecated
	FontManager_WillIncludeFont(sender objc.IObject, fontName string) bool
	AvailableFonts() []string
	AvailableFontFamilies() []string
	SelectedFont() Font
	IsMultiple() bool
	CurrentFontAction() FontAction
	IsEnabled() bool
	SetEnabled(value bool)
	Action() objc.Selector
	SetAction(value objc.Selector)
	Target() objc.Object
	SetTarget(value objc.IObject)
	// deprecated
	CollectionNames() []objc.Object
	// deprecated
	Delegate() objc.Object
	// deprecated
	SetDelegate(value objc.IObject)
}

type FontManager struct {
	objc.Object
}

func MakeFontManager(ptr unsafe.Pointer) FontManager {
	return FontManager{
		Object: objc.MakeObject(ptr),
	}
}

func (fc _FontManagerClass) Alloc() FontManager {
	rv := ffi.CallMethod[FontManager](fc, "alloc")
	return rv
}

func (f_ FontManager) Init() FontManager {
	rv := ffi.CallMethod[FontManager](f_, "init")
	return rv
}

func (fc _FontManagerClass) New() FontManager {
	rv := ffi.CallMethod[FontManager](fc, "new")
	rv.Autorelease()
	return rv
}

func NewFontManager() FontManager {
	return FontManagerClass.New()
}

func (f_ FontManager) AvailableFontNamesWithTraits(someTraits FontTraitMask) []string {
	rv := ffi.CallMethod[[]string](f_, "availableFontNamesWithTraits:", someTraits)
	return rv
}

func (f_ FontManager) AvailableMembersOfFontFamily(fam string) [][]objc.Object {
	rv := ffi.CallMethod[[][]objc.Object](f_, "availableMembersOfFontFamily:", fam)
	return rv
}

func (f_ FontManager) SetSelectedFont_IsMultiple(fontObj IFont, flag bool) {
	ffi.CallMethod[ffi.Void](f_, "setSelectedFont:isMultiple:", fontObj, flag)
}

func (f_ FontManager) SendAction() bool {
	rv := ffi.CallMethod[bool](f_, "sendAction")
	return rv
}

func (f_ FontManager) LocalizedNameForFamily_Face(family string, faceKey string) string {
	rv := ffi.CallMethod[string](f_, "localizedNameForFamily:face:", family, faceKey)
	return rv
}

func (f_ FontManager) AddFontTrait(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](f_, "addFontTrait:", sender)
}

func (f_ FontManager) RemoveFontTrait(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](f_, "removeFontTrait:", sender)
}

func (f_ FontManager) ModifyFont(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](f_, "modifyFont:", sender)
}

func (f_ FontManager) ModifyFontViaPanel(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](f_, "modifyFontViaPanel:", sender)
}

func (f_ FontManager) OrderFrontStylesPanel(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](f_, "orderFrontStylesPanel:", sender)
}

func (f_ FontManager) OrderFrontFontPanel(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](f_, "orderFrontFontPanel:", sender)
}

func (f_ FontManager) ConvertFont(fontObj IFont) Font {
	rv := ffi.CallMethod[Font](f_, "convertFont:", fontObj)
	return rv
}

// deprecated
func (f_ FontManager) ChangeFont(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](f_, "changeFont:", sender)
}

func (f_ FontManager) ConvertFont_ToFace(fontObj IFont, typeface string) Font {
	rv := ffi.CallMethod[Font](f_, "convertFont:toFace:", fontObj, typeface)
	return rv
}

func (f_ FontManager) ConvertFont_ToFamily(fontObj IFont, family string) Font {
	rv := ffi.CallMethod[Font](f_, "convertFont:toFamily:", fontObj, family)
	return rv
}

func (f_ FontManager) ConvertFont_ToHaveTrait(fontObj IFont, trait FontTraitMask) Font {
	rv := ffi.CallMethod[Font](f_, "convertFont:toHaveTrait:", fontObj, trait)
	return rv
}

func (f_ FontManager) ConvertFont_ToNotHaveTrait(fontObj IFont, trait FontTraitMask) Font {
	rv := ffi.CallMethod[Font](f_, "convertFont:toNotHaveTrait:", fontObj, trait)
	return rv
}

func (f_ FontManager) ConvertFont_ToSize(fontObj IFont, size float64) Font {
	rv := ffi.CallMethod[Font](f_, "convertFont:toSize:", fontObj, size)
	return rv
}

func (f_ FontManager) ConvertWeight_OfFont(upFlag bool, fontObj IFont) Font {
	rv := ffi.CallMethod[Font](f_, "convertWeight:ofFont:", upFlag, fontObj)
	return rv
}

func (f_ FontManager) ConvertFontTraits(traits FontTraitMask) FontTraitMask {
	rv := ffi.CallMethod[FontTraitMask](f_, "convertFontTraits:", traits)
	return rv
}

func (f_ FontManager) FontWithFamily_Traits_Weight_Size(family string, traits FontTraitMask, weight int, size float64) Font {
	rv := ffi.CallMethod[Font](f_, "fontWithFamily:traits:weight:size:", family, traits, weight, size)
	return rv
}

func (f_ FontManager) TraitsOfFont(fontObj IFont) FontTraitMask {
	rv := ffi.CallMethod[FontTraitMask](f_, "traitsOfFont:", fontObj)
	return rv
}

func (f_ FontManager) FontNamed_HasTraits(fName string, someTraits FontTraitMask) bool {
	rv := ffi.CallMethod[bool](f_, "fontNamed:hasTraits:", fName, someTraits)
	return rv
}

func (f_ FontManager) WeightOfFont(fontObj IFont) int {
	rv := ffi.CallMethod[int](f_, "weightOfFont:", fontObj)
	return rv
}

func (f_ FontManager) FontPanel(create bool) FontPanel {
	rv := ffi.CallMethod[FontPanel](f_, "fontPanel:", create)
	return rv
}

func (f_ FontManager) SetFontMenu(newMenu IMenu) {
	ffi.CallMethod[ffi.Void](f_, "setFontMenu:", newMenu)
}

func (f_ FontManager) FontMenu(create bool) Menu {
	rv := ffi.CallMethod[Menu](f_, "fontMenu:", create)
	return rv
}

func (f_ FontManager) SetSelectedAttributes_IsMultiple(attributes map[string]objc.IObject, flag bool) {
	ffi.CallMethod[ffi.Void](f_, "setSelectedAttributes:isMultiple:", attributes, flag)
}

func (f_ FontManager) ConvertAttributes(attributes map[string]objc.IObject) map[string]objc.Object {
	rv := ffi.CallMethod[map[string]objc.Object](f_, "convertAttributes:", attributes)
	return rv
}

// deprecated
func (f_ FontManager) AvailableFontNamesMatchingFontDescriptor(descriptor IFontDescriptor) []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](f_, "availableFontNamesMatchingFontDescriptor:", descriptor)
	return rv
}

// deprecated
func (f_ FontManager) FontDescriptorsInCollection(collectionNames string) []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](f_, "fontDescriptorsInCollection:", collectionNames)
	return rv
}

// deprecated
func (f_ FontManager) AddCollection_Options(collectionName string, collectionOptions FontCollectionOptions) bool {
	rv := ffi.CallMethod[bool](f_, "addCollection:options:", collectionName, collectionOptions)
	return rv
}

// deprecated
func (f_ FontManager) RemoveCollection(collectionName string) bool {
	rv := ffi.CallMethod[bool](f_, "removeCollection:", collectionName)
	return rv
}

// deprecated
func (f_ FontManager) AddFontDescriptors_ToCollection(descriptors []objc.IObject, collectionName string) {
	ffi.CallMethod[ffi.Void](f_, "addFontDescriptors:toCollection:", descriptors, collectionName)
}

// deprecated
func (f_ FontManager) RemoveFontDescriptor_FromCollection(descriptor IFontDescriptor, collection string) {
	ffi.CallMethod[ffi.Void](f_, "removeFontDescriptor:fromCollection:", descriptor, collection)
}

// deprecated
func (f_ FontManager) FontManager_WillIncludeFont(sender objc.IObject, fontName string) bool {
	rv := ffi.CallMethod[bool](f_, "fontManager:willIncludeFont:", sender, fontName)
	return rv
}

func (fc _FontManagerClass) SharedFontManager() FontManager {
	rv := ffi.CallMethod[FontManager](fc, "sharedFontManager")
	return rv
}

func (f_ FontManager) AvailableFonts() []string {
	rv := ffi.CallMethod[[]string](f_, "availableFonts")
	return rv
}

func (f_ FontManager) AvailableFontFamilies() []string {
	rv := ffi.CallMethod[[]string](f_, "availableFontFamilies")
	return rv
}

func (f_ FontManager) SelectedFont() Font {
	rv := ffi.CallMethod[Font](f_, "selectedFont")
	return rv
}

func (f_ FontManager) IsMultiple() bool {
	rv := ffi.CallMethod[bool](f_, "isMultiple")
	return rv
}

func (f_ FontManager) CurrentFontAction() FontAction {
	rv := ffi.CallMethod[FontAction](f_, "currentFontAction")
	return rv
}

func (f_ FontManager) IsEnabled() bool {
	rv := ffi.CallMethod[bool](f_, "isEnabled")
	return rv
}

func (f_ FontManager) SetEnabled(value bool) {
	ffi.CallMethod[ffi.Void](f_, "setEnabled:", value)
}

func (f_ FontManager) Action() objc.Selector {
	rv := ffi.CallMethod[objc.Selector](f_, "action")
	return rv
}

func (f_ FontManager) SetAction(value objc.Selector) {
	ffi.CallMethod[ffi.Void](f_, "setAction:", value)
}

func (f_ FontManager) Target() objc.Object {
	rv := ffi.CallMethod[objc.Object](f_, "target")
	return rv
}

func (f_ FontManager) SetTarget(value objc.IObject) {
	ffi.CallMethod[ffi.Void](f_, "setTarget:", value)
}

// deprecated
func (f_ FontManager) CollectionNames() []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](f_, "collectionNames")
	return rv
}

// deprecated
func (f_ FontManager) Delegate() objc.Object {
	rv := ffi.CallMethod[objc.Object](f_, "delegate")
	return rv
}

// deprecated
func (f_ FontManager) SetDelegate(value objc.IObject) {
	ffi.CallMethod[ffi.Void](f_, "setDelegate:", value)
}
