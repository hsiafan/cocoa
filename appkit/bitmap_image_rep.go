// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var BitmapImageRepClass = _BitmapImageRepClass{objc.GetClass("NSBitmapImageRep")}

type _BitmapImageRepClass struct {
	objc.Class
}

type IBitmapImageRep interface {
	IImageRep
	ColorizeByMappingGray_ToColor_BlackMapping_WhiteMapping(midPoint float64, midPointColor IColor, shadowColor IColor, lightColor IColor)
	TIFFRepresentationUsingCompression_Factor(comp TIFFCompression, factor float32) []byte
	RepresentationUsingType_Properties(storageType BitmapImageFileType, properties map[BitmapImageRepPropertyKey]objc.IObject) []byte
	CanBeCompressedUsing(compression TIFFCompression) bool
	SetCompression_Factor(compression TIFFCompression, factor float32)
	GetCompression_Factor(compression *TIFFCompression, factor *float32)
	SetProperty_WithValue(property BitmapImageRepPropertyKey, value objc.IObject)
	ValueForProperty(property BitmapImageRepPropertyKey) objc.Object
	IncrementalLoadFromData_Complete(data []byte, complete bool) int
	SetColor_AtX_Y(color IColor, x int, y int)
	ColorAtX_Y(x int, y int) Color
	SetPixel_AtX_Y(p *uint, x int, y int)
	GetPixel_AtX_Y(p *uint, x int, y int)
	BitmapImageRepByConvertingToColorSpace_RenderingIntent(targetSpace IColorSpace, renderingIntent ColorRenderingIntent) BitmapImageRep
	BitmapImageRepByRetaggingWithColorSpace(newSpace IColorSpace) BitmapImageRep
	BitmapFormat() BitmapFormat
	BitsPerPixel() int
	BytesPerPlane() int
	BytesPerRow() int
	IsPlanar() bool
	NumberOfPlanes() int
	SamplesPerPixel() int
	BitmapData() *byte
	TIFFRepresentation() []byte
	CGImage() coregraphics.ImageRef
	ColorSpace() ColorSpace
}

type BitmapImageRep struct {
	ImageRep
}

func MakeBitmapImageRep(ptr unsafe.Pointer) BitmapImageRep {
	return BitmapImageRep{
		ImageRep: MakeImageRep(ptr),
	}
}

func (bc _BitmapImageRepClass) ImageRepWithData(data []byte) BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](bc, "imageRepWithData:", data)
	return rv
}

func (b_ BitmapImageRep) InitWithCGImage(cgImage coregraphics.ImageRef) BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](b_, "initWithCGImage:", cgImage)
	return rv
}

func (b_ BitmapImageRep) InitWithData(data []byte) BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](b_, "initWithData:", data)
	return rv
}

func (b_ BitmapImageRep) InitForIncrementalLoad() BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](b_, "initForIncrementalLoad")
	return rv
}

func (b_ BitmapImageRep) InitWithFocusedViewRect(rect foundation.Rect) BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](b_, "initWithFocusedViewRect:", rect)
	return rv
}

func (b_ BitmapImageRep) Init() BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](b_, "init")
	return rv
}

func (bc _BitmapImageRepClass) Alloc() BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](bc, "alloc")
	return rv
}

func (bc _BitmapImageRepClass) New() BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](bc, "new")
	rv.Autorelease()
	return rv
}

func NewBitmapImageRep() BitmapImageRep {
	return BitmapImageRepClass.New()
}

func (bc _BitmapImageRepClass) ImageRepsWithData(data []byte) []ImageRep {
	rv := ffi.CallMethod[[]ImageRep](bc, "imageRepsWithData:", data)
	return rv
}

func (b_ BitmapImageRep) ColorizeByMappingGray_ToColor_BlackMapping_WhiteMapping(midPoint float64, midPointColor IColor, shadowColor IColor, lightColor IColor) {
	ffi.CallMethod[ffi.Void](b_, "colorizeByMappingGray:toColor:blackMapping:whiteMapping:", midPoint, midPointColor, shadowColor, lightColor)
}

func (bc _BitmapImageRepClass) TIFFRepresentationOfImageRepsInArray(array []IImageRep) []byte {
	rv := ffi.CallMethod[[]byte](bc, "TIFFRepresentationOfImageRepsInArray:", array)
	return rv
}

func (bc _BitmapImageRepClass) TIFFRepresentationOfImageRepsInArray_UsingCompression_Factor(array []IImageRep, comp TIFFCompression, factor float32) []byte {
	rv := ffi.CallMethod[[]byte](bc, "TIFFRepresentationOfImageRepsInArray:usingCompression:factor:", array, comp, factor)
	return rv
}

func (b_ BitmapImageRep) TIFFRepresentationUsingCompression_Factor(comp TIFFCompression, factor float32) []byte {
	rv := ffi.CallMethod[[]byte](b_, "TIFFRepresentationUsingCompression:factor:", comp, factor)
	return rv
}

func (bc _BitmapImageRepClass) RepresentationOfImageRepsInArray_UsingType_Properties(imageReps []IImageRep, storageType BitmapImageFileType, properties map[BitmapImageRepPropertyKey]objc.IObject) []byte {
	rv := ffi.CallMethod[[]byte](bc, "representationOfImageRepsInArray:usingType:properties:", imageReps, storageType, properties)
	return rv
}

func (b_ BitmapImageRep) RepresentationUsingType_Properties(storageType BitmapImageFileType, properties map[BitmapImageRepPropertyKey]objc.IObject) []byte {
	rv := ffi.CallMethod[[]byte](b_, "representationUsingType:properties:", storageType, properties)
	return rv
}

func (bc _BitmapImageRepClass) LocalizedNameForTIFFCompressionType(compression TIFFCompression) string {
	rv := ffi.CallMethod[string](bc, "localizedNameForTIFFCompressionType:", compression)
	return rv
}

func (b_ BitmapImageRep) CanBeCompressedUsing(compression TIFFCompression) bool {
	rv := ffi.CallMethod[bool](b_, "canBeCompressedUsing:", compression)
	return rv
}

func (b_ BitmapImageRep) SetCompression_Factor(compression TIFFCompression, factor float32) {
	ffi.CallMethod[ffi.Void](b_, "setCompression:factor:", compression, factor)
}

func (b_ BitmapImageRep) GetCompression_Factor(compression *TIFFCompression, factor *float32) {
	ffi.CallMethod[ffi.Void](b_, "getCompression:factor:", compression, factor)
}

func (b_ BitmapImageRep) SetProperty_WithValue(property BitmapImageRepPropertyKey, value objc.IObject) {
	ffi.CallMethod[ffi.Void](b_, "setProperty:withValue:", property, value)
}

func (b_ BitmapImageRep) ValueForProperty(property BitmapImageRepPropertyKey) objc.Object {
	rv := ffi.CallMethod[objc.Object](b_, "valueForProperty:", property)
	return rv
}

func (b_ BitmapImageRep) IncrementalLoadFromData_Complete(data []byte, complete bool) int {
	rv := ffi.CallMethod[int](b_, "incrementalLoadFromData:complete:", data, complete)
	return rv
}

func (b_ BitmapImageRep) SetColor_AtX_Y(color IColor, x int, y int) {
	ffi.CallMethod[ffi.Void](b_, "setColor:atX:y:", color, x, y)
}

func (b_ BitmapImageRep) ColorAtX_Y(x int, y int) Color {
	rv := ffi.CallMethod[Color](b_, "colorAtX:y:", x, y)
	return rv
}

func (b_ BitmapImageRep) SetPixel_AtX_Y(p *uint, x int, y int) {
	ffi.CallMethod[ffi.Void](b_, "setPixel:atX:y:", p, x, y)
}

func (b_ BitmapImageRep) GetPixel_AtX_Y(p *uint, x int, y int) {
	ffi.CallMethod[ffi.Void](b_, "getPixel:atX:y:", p, x, y)
}

func (b_ BitmapImageRep) BitmapImageRepByConvertingToColorSpace_RenderingIntent(targetSpace IColorSpace, renderingIntent ColorRenderingIntent) BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](b_, "bitmapImageRepByConvertingToColorSpace:renderingIntent:", targetSpace, renderingIntent)
	return rv
}

func (b_ BitmapImageRep) BitmapImageRepByRetaggingWithColorSpace(newSpace IColorSpace) BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](b_, "bitmapImageRepByRetaggingWithColorSpace:", newSpace)
	return rv
}

func (b_ BitmapImageRep) BitmapFormat() BitmapFormat {
	rv := ffi.CallMethod[BitmapFormat](b_, "bitmapFormat")
	return rv
}

func (b_ BitmapImageRep) BitsPerPixel() int {
	rv := ffi.CallMethod[int](b_, "bitsPerPixel")
	return rv
}

func (b_ BitmapImageRep) BytesPerPlane() int {
	rv := ffi.CallMethod[int](b_, "bytesPerPlane")
	return rv
}

func (b_ BitmapImageRep) BytesPerRow() int {
	rv := ffi.CallMethod[int](b_, "bytesPerRow")
	return rv
}

func (b_ BitmapImageRep) IsPlanar() bool {
	rv := ffi.CallMethod[bool](b_, "isPlanar")
	return rv
}

func (b_ BitmapImageRep) NumberOfPlanes() int {
	rv := ffi.CallMethod[int](b_, "numberOfPlanes")
	return rv
}

func (b_ BitmapImageRep) SamplesPerPixel() int {
	rv := ffi.CallMethod[int](b_, "samplesPerPixel")
	return rv
}

func (b_ BitmapImageRep) BitmapData() *byte {
	rv := ffi.CallMethod[*byte](b_, "bitmapData")
	return rv
}

func (b_ BitmapImageRep) TIFFRepresentation() []byte {
	rv := ffi.CallMethod[[]byte](b_, "TIFFRepresentation")
	return rv
}

func (b_ BitmapImageRep) CGImage() coregraphics.ImageRef {
	rv := ffi.CallMethod[coregraphics.ImageRef](b_, "CGImage")
	return rv
}

func (b_ BitmapImageRep) ColorSpace() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](b_, "colorSpace")
	return rv
}
