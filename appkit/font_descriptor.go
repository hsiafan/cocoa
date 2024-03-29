// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

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
	rv := objc.CallMethod[FontDescriptor](f_, objc.SEL("initWithFontAttributes:"), attributes)
	return rv
}

func (f_ FontDescriptor) FontDescriptorWithDesign(design FontDescriptorSystemDesign) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.SEL("fontDescriptorWithDesign:"), design)
	return rv
}

func (fc _FontDescriptorClass) Alloc() FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](fc, objc.SEL("alloc"))
	return rv
}

func (fc _FontDescriptorClass) New() FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](fc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewFontDescriptor() FontDescriptor {
	return FontDescriptorClass.New()
}

func (f_ FontDescriptor) Init() FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.SEL("init"))
	return rv
}

func (fc _FontDescriptorClass) PreferredFontDescriptorForTextStyle_Options(style FontTextStyle, options map[FontTextStyleOptionKey]objc.IObject) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](fc, objc.SEL("preferredFontDescriptorForTextStyle:options:"), style, options)
	return rv
}

func (fc _FontDescriptorClass) FontDescriptorWithFontAttributes(attributes map[FontDescriptorAttributeName]objc.IObject) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](fc, objc.SEL("fontDescriptorWithFontAttributes:"), attributes)
	return rv
}

func (fc _FontDescriptorClass) FontDescriptorWithName_Matrix(fontName string, matrix foundation.IAffineTransform) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](fc, objc.SEL("fontDescriptorWithName:matrix:"), fontName, objc.ExtractPtr(matrix))
	return rv
}

func (fc _FontDescriptorClass) FontDescriptorWithName_Size(fontName string, size float64) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](fc, objc.SEL("fontDescriptorWithName:size:"), fontName, size)
	return rv
}

func (f_ FontDescriptor) FontDescriptorByAddingAttributes(attributes map[FontDescriptorAttributeName]objc.IObject) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.SEL("fontDescriptorByAddingAttributes:"), attributes)
	return rv
}

func (f_ FontDescriptor) FontDescriptorWithFace(newFace string) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.SEL("fontDescriptorWithFace:"), newFace)
	return rv
}

func (f_ FontDescriptor) FontDescriptorWithFamily(newFamily string) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.SEL("fontDescriptorWithFamily:"), newFamily)
	return rv
}

func (f_ FontDescriptor) FontDescriptorWithMatrix(matrix foundation.IAffineTransform) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.SEL("fontDescriptorWithMatrix:"), objc.ExtractPtr(matrix))
	return rv
}

func (f_ FontDescriptor) FontDescriptorWithSize(newPointSize float64) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.SEL("fontDescriptorWithSize:"), newPointSize)
	return rv
}

func (f_ FontDescriptor) FontDescriptorWithSymbolicTraits(symbolicTraits FontDescriptorSymbolicTraits) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.SEL("fontDescriptorWithSymbolicTraits:"), symbolicTraits)
	return rv
}

func (f_ FontDescriptor) MatchingFontDescriptorsWithMandatoryKeys(mandatoryKeys foundation.ISet) []FontDescriptor {
	rv := objc.CallMethod[[]FontDescriptor](f_, objc.SEL("matchingFontDescriptorsWithMandatoryKeys:"), objc.ExtractPtr(mandatoryKeys))
	return rv
}

func (f_ FontDescriptor) MatchingFontDescriptorWithMandatoryKeys(mandatoryKeys foundation.ISet) FontDescriptor {
	rv := objc.CallMethod[FontDescriptor](f_, objc.SEL("matchingFontDescriptorWithMandatoryKeys:"), objc.ExtractPtr(mandatoryKeys))
	return rv
}

func (f_ FontDescriptor) ObjectForKey(attribute FontDescriptorAttributeName) objc.Object {
	rv := objc.CallMethod[objc.Object](f_, objc.SEL("objectForKey:"), attribute)
	return rv
}

func (f_ FontDescriptor) FontAttributes() map[FontDescriptorAttributeName]objc.Object {
	rv := objc.CallMethod[map[FontDescriptorAttributeName]objc.Object](f_, objc.SEL("fontAttributes"))
	return rv
}

func (f_ FontDescriptor) Matrix() foundation.AffineTransform {
	rv := objc.CallMethod[foundation.AffineTransform](f_, objc.SEL("matrix"))
	return rv
}

func (f_ FontDescriptor) PointSize() float64 {
	rv := objc.CallMethod[float64](f_, objc.SEL("pointSize"))
	return rv
}

func (f_ FontDescriptor) PostscriptName() string {
	rv := objc.CallMethod[string](f_, objc.SEL("postscriptName"))
	return rv
}

func (f_ FontDescriptor) SymbolicTraits() FontDescriptorSymbolicTraits {
	rv := objc.CallMethod[FontDescriptorSymbolicTraits](f_, objc.SEL("symbolicTraits"))
	return rv
}

func (f_ FontDescriptor) RequiresFontAssetRequest() bool {
	rv := objc.CallMethod[bool](f_, objc.SEL("requiresFontAssetRequest"))
	return rv
}
