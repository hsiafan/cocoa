// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
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
	rv := objc.CallMethod[BitmapImageRep](bc, objc.GetSelector("imageRepWithData:"), data)
	return rv
}

func (b_ BitmapImageRep) InitWithCGImage(cgImage coregraphics.ImageRef) BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](b_, objc.GetSelector("initWithCGImage:"), cgImage)
	return rv
}

func (b_ BitmapImageRep) InitWithData(data []byte) BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](b_, objc.GetSelector("initWithData:"), data)
	return rv
}

func (b_ BitmapImageRep) InitForIncrementalLoad() BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](b_, objc.GetSelector("initForIncrementalLoad"))
	return rv
}

func (b_ BitmapImageRep) InitWithFocusedViewRect(rect foundation.Rect) BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](b_, objc.GetSelector("initWithFocusedViewRect:"), rect)
	return rv
}

func (b_ BitmapImageRep) Init() BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](b_, objc.GetSelector("init"))
	return rv
}

func (bc _BitmapImageRepClass) Alloc() BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](bc, objc.GetSelector("alloc"))
	return rv
}

func (bc _BitmapImageRepClass) New() BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](bc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewBitmapImageRep() BitmapImageRep {
	return BitmapImageRepClass.New()
}

func (bc _BitmapImageRepClass) ImageRepsWithData(data []byte) []ImageRep {
	rv := objc.CallMethod[[]ImageRep](bc, objc.GetSelector("imageRepsWithData:"), data)
	return rv
}

func (b_ BitmapImageRep) ColorizeByMappingGray_ToColor_BlackMapping_WhiteMapping(midPoint float64, midPointColor IColor, shadowColor IColor, lightColor IColor) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("colorizeByMappingGray:toColor:blackMapping:whiteMapping:"), midPoint, objc.ExtractPtr(midPointColor), objc.ExtractPtr(shadowColor), objc.ExtractPtr(lightColor))
}

func (bc _BitmapImageRepClass) TIFFRepresentationOfImageRepsInArray(array []IImageRep) []byte {
	rv := objc.CallMethod[[]byte](bc, objc.GetSelector("TIFFRepresentationOfImageRepsInArray:"), array)
	return rv
}

func (bc _BitmapImageRepClass) TIFFRepresentationOfImageRepsInArray_UsingCompression_Factor(array []IImageRep, comp TIFFCompression, factor float32) []byte {
	rv := objc.CallMethod[[]byte](bc, objc.GetSelector("TIFFRepresentationOfImageRepsInArray:usingCompression:factor:"), array, comp, factor)
	return rv
}

func (b_ BitmapImageRep) TIFFRepresentationUsingCompression_Factor(comp TIFFCompression, factor float32) []byte {
	rv := objc.CallMethod[[]byte](b_, objc.GetSelector("TIFFRepresentationUsingCompression:factor:"), comp, factor)
	return rv
}

func (bc _BitmapImageRepClass) RepresentationOfImageRepsInArray_UsingType_Properties(imageReps []IImageRep, storageType BitmapImageFileType, properties map[BitmapImageRepPropertyKey]objc.IObject) []byte {
	rv := objc.CallMethod[[]byte](bc, objc.GetSelector("representationOfImageRepsInArray:usingType:properties:"), imageReps, storageType, properties)
	return rv
}

func (b_ BitmapImageRep) RepresentationUsingType_Properties(storageType BitmapImageFileType, properties map[BitmapImageRepPropertyKey]objc.IObject) []byte {
	rv := objc.CallMethod[[]byte](b_, objc.GetSelector("representationUsingType:properties:"), storageType, properties)
	return rv
}

func (bc _BitmapImageRepClass) LocalizedNameForTIFFCompressionType(compression TIFFCompression) string {
	rv := objc.CallMethod[string](bc, objc.GetSelector("localizedNameForTIFFCompressionType:"), compression)
	return rv
}

func (b_ BitmapImageRep) CanBeCompressedUsing(compression TIFFCompression) bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("canBeCompressedUsing:"), compression)
	return rv
}

func (b_ BitmapImageRep) SetCompression_Factor(compression TIFFCompression, factor float32) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setCompression:factor:"), compression, factor)
}

func (b_ BitmapImageRep) GetCompression_Factor(compression *TIFFCompression, factor *float32) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("getCompression:factor:"), compression, factor)
}

func (b_ BitmapImageRep) SetProperty_WithValue(property BitmapImageRepPropertyKey, value objc.IObject) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setProperty:withValue:"), property, objc.ExtractPtr(value))
}

func (b_ BitmapImageRep) ValueForProperty(property BitmapImageRepPropertyKey) objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("valueForProperty:"), property)
	return rv
}

func (b_ BitmapImageRep) IncrementalLoadFromData_Complete(data []byte, complete bool) int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("incrementalLoadFromData:complete:"), data, complete)
	return rv
}

func (b_ BitmapImageRep) SetColor_AtX_Y(color IColor, x int, y int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setColor:atX:y:"), objc.ExtractPtr(color), x, y)
}

func (b_ BitmapImageRep) ColorAtX_Y(x int, y int) Color {
	rv := objc.CallMethod[Color](b_, objc.GetSelector("colorAtX:y:"), x, y)
	return rv
}

func (b_ BitmapImageRep) SetPixel_AtX_Y(p *uint, x int, y int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setPixel:atX:y:"), p, x, y)
}

func (b_ BitmapImageRep) GetPixel_AtX_Y(p *uint, x int, y int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("getPixel:atX:y:"), p, x, y)
}

func (b_ BitmapImageRep) BitmapImageRepByConvertingToColorSpace_RenderingIntent(targetSpace IColorSpace, renderingIntent ColorRenderingIntent) BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](b_, objc.GetSelector("bitmapImageRepByConvertingToColorSpace:renderingIntent:"), objc.ExtractPtr(targetSpace), renderingIntent)
	return rv
}

func (b_ BitmapImageRep) BitmapImageRepByRetaggingWithColorSpace(newSpace IColorSpace) BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](b_, objc.GetSelector("bitmapImageRepByRetaggingWithColorSpace:"), objc.ExtractPtr(newSpace))
	return rv
}

func (b_ BitmapImageRep) BitmapFormat() BitmapFormat {
	rv := objc.CallMethod[BitmapFormat](b_, objc.GetSelector("bitmapFormat"))
	return rv
}

func (b_ BitmapImageRep) BitsPerPixel() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("bitsPerPixel"))
	return rv
}

func (b_ BitmapImageRep) BytesPerPlane() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("bytesPerPlane"))
	return rv
}

func (b_ BitmapImageRep) BytesPerRow() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("bytesPerRow"))
	return rv
}

func (b_ BitmapImageRep) IsPlanar() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("isPlanar"))
	return rv
}

func (b_ BitmapImageRep) NumberOfPlanes() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("numberOfPlanes"))
	return rv
}

func (b_ BitmapImageRep) SamplesPerPixel() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("samplesPerPixel"))
	return rv
}

func (b_ BitmapImageRep) BitmapData() *byte {
	rv := objc.CallMethod[*byte](b_, objc.GetSelector("bitmapData"))
	return rv
}

func (b_ BitmapImageRep) TIFFRepresentation() []byte {
	rv := objc.CallMethod[[]byte](b_, objc.GetSelector("TIFFRepresentation"))
	return rv
}

func (b_ BitmapImageRep) CGImage() coregraphics.ImageRef {
	rv := objc.CallMethod[coregraphics.ImageRef](b_, objc.GetSelector("CGImage"))
	return rv
}

func (b_ BitmapImageRep) ColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](b_, objc.GetSelector("colorSpace"))
	return rv
}
