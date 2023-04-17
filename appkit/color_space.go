// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var ColorSpaceClass = _ColorSpaceClass{objc.GetClass("NSColorSpace")}

type _ColorSpaceClass struct {
	objc.Class
}

type IColorSpace interface {
	objc.IObject
	CGColorSpace() coregraphics.ColorSpaceRef
	ColorSpaceModel() ColorSpaceModel
	ColorSyncProfile() unsafe.Pointer
	ICCProfileData() []byte
	LocalizedName() string
	NumberOfColorComponents() int
}

type ColorSpace struct {
	objc.Object
}

func MakeColorSpace(ptr unsafe.Pointer) ColorSpace {
	return ColorSpace{
		Object: objc.MakeObject(ptr),
	}
}

func (c_ ColorSpace) InitWithCGColorSpace(cgColorSpace coregraphics.ColorSpaceRef) ColorSpace {
	rv := ffi.CallMethod[ColorSpace](c_, "initWithCGColorSpace:", cgColorSpace)
	return rv
}

func (c_ ColorSpace) InitWithColorSyncProfile(prof unsafe.Pointer) ColorSpace {
	rv := ffi.CallMethod[ColorSpace](c_, "initWithColorSyncProfile:", prof)
	return rv
}

func (c_ ColorSpace) InitWithICCProfileData(iccData []byte) ColorSpace {
	rv := ffi.CallMethod[ColorSpace](c_, "initWithICCProfileData:", iccData)
	return rv
}

func (cc _ColorSpaceClass) Alloc() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](cc, "alloc")
	return rv
}

func (cc _ColorSpaceClass) New() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](cc, "new")
	rv.Autorelease()
	return rv
}

func NewColorSpace() ColorSpace {
	return ColorSpaceClass.New()
}

func (c_ ColorSpace) Init() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](c_, "init")
	return rv
}

func (cc _ColorSpaceClass) AvailableColorSpacesWithModel(model ColorSpaceModel) []ColorSpace {
	rv := ffi.CallMethod[[]ColorSpace](cc, "availableColorSpacesWithModel:", model)
	return rv
}

func (cc _ColorSpaceClass) DeviceRGBColorSpace() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](cc, "deviceRGBColorSpace")
	return rv
}

func (cc _ColorSpaceClass) GenericRGBColorSpace() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](cc, "genericRGBColorSpace")
	return rv
}

func (cc _ColorSpaceClass) DeviceCMYKColorSpace() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](cc, "deviceCMYKColorSpace")
	return rv
}

func (cc _ColorSpaceClass) GenericCMYKColorSpace() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](cc, "genericCMYKColorSpace")
	return rv
}

func (cc _ColorSpaceClass) DeviceGrayColorSpace() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](cc, "deviceGrayColorSpace")
	return rv
}

func (cc _ColorSpaceClass) GenericGrayColorSpace() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](cc, "genericGrayColorSpace")
	return rv
}

func (cc _ColorSpaceClass) SRGBColorSpace() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](cc, "sRGBColorSpace")
	return rv
}

func (cc _ColorSpaceClass) ExtendedSRGBColorSpace() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](cc, "extendedSRGBColorSpace")
	return rv
}

func (cc _ColorSpaceClass) DisplayP3ColorSpace() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](cc, "displayP3ColorSpace")
	return rv
}

func (cc _ColorSpaceClass) GenericGamma22GrayColorSpace() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](cc, "genericGamma22GrayColorSpace")
	return rv
}

func (cc _ColorSpaceClass) ExtendedGenericGamma22GrayColorSpace() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](cc, "extendedGenericGamma22GrayColorSpace")
	return rv
}

func (cc _ColorSpaceClass) AdobeRGB1998ColorSpace() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](cc, "adobeRGB1998ColorSpace")
	return rv
}

func (c_ ColorSpace) CGColorSpace() coregraphics.ColorSpaceRef {
	rv := ffi.CallMethod[coregraphics.ColorSpaceRef](c_, "CGColorSpace")
	return rv
}

func (c_ ColorSpace) ColorSpaceModel() ColorSpaceModel {
	rv := ffi.CallMethod[ColorSpaceModel](c_, "colorSpaceModel")
	return rv
}

func (c_ ColorSpace) ColorSyncProfile() unsafe.Pointer {
	rv := ffi.CallMethod[unsafe.Pointer](c_, "colorSyncProfile")
	return rv
}

func (c_ ColorSpace) ICCProfileData() []byte {
	rv := ffi.CallMethod[[]byte](c_, "ICCProfileData")
	return rv
}

func (c_ ColorSpace) LocalizedName() string {
	rv := ffi.CallMethod[string](c_, "localizedName")
	return rv
}

func (c_ ColorSpace) NumberOfColorComponents() int {
	rv := ffi.CallMethod[int](c_, "numberOfColorComponents")
	return rv
}