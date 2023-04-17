// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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

func (fc _FontDescriptorClass) New() FontDescriptor {
	rv := ffi.CallMethod[FontDescriptor](fc, "new")
	rv.Autorelease()
	return rv
}

func NewFontDescriptor() FontDescriptor {
	return FontDescriptorClass.New()
}

func (f_ FontDescriptor) Init() FontDescriptor {
	rv := ffi.CallMethod[FontDescriptor](f_, "init")
	return rv
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
