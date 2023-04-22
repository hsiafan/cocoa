// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
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
	rv := objc.CallMethod[ColorSpace](c_, objc.GetSelector("initWithCGColorSpace:"), cgColorSpace)
	return rv
}

func (c_ ColorSpace) InitWithColorSyncProfile(prof unsafe.Pointer) ColorSpace {
	rv := objc.CallMethod[ColorSpace](c_, objc.GetSelector("initWithColorSyncProfile:"), prof)
	return rv
}

func (c_ ColorSpace) InitWithICCProfileData(iccData []byte) ColorSpace {
	rv := objc.CallMethod[ColorSpace](c_, objc.GetSelector("initWithICCProfileData:"), iccData)
	return rv
}

func (cc _ColorSpaceClass) Alloc() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("alloc"))
	return rv
}

func (cc _ColorSpaceClass) New() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewColorSpace() ColorSpace {
	return ColorSpaceClass.New()
}

func (c_ ColorSpace) Init() ColorSpace {
	rv := objc.CallMethod[ColorSpace](c_, objc.GetSelector("init"))
	return rv
}

func (cc _ColorSpaceClass) AvailableColorSpacesWithModel(model ColorSpaceModel) []ColorSpace {
	rv := objc.CallMethod[[]ColorSpace](cc, objc.GetSelector("availableColorSpacesWithModel:"), model)
	return rv
}

func (cc _ColorSpaceClass) DeviceRGBColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("deviceRGBColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) GenericRGBColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("genericRGBColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) DeviceCMYKColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("deviceCMYKColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) GenericCMYKColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("genericCMYKColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) DeviceGrayColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("deviceGrayColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) GenericGrayColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("genericGrayColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) SRGBColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("sRGBColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) ExtendedSRGBColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("extendedSRGBColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) DisplayP3ColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("displayP3ColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) GenericGamma22GrayColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("genericGamma22GrayColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) ExtendedGenericGamma22GrayColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("extendedGenericGamma22GrayColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) AdobeRGB1998ColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.GetSelector("adobeRGB1998ColorSpace"))
	return rv
}

func (c_ ColorSpace) CGColorSpace() coregraphics.ColorSpaceRef {
	rv := objc.CallMethod[coregraphics.ColorSpaceRef](c_, objc.GetSelector("CGColorSpace"))
	return rv
}

func (c_ ColorSpace) ColorSpaceModel() ColorSpaceModel {
	rv := objc.CallMethod[ColorSpaceModel](c_, objc.GetSelector("colorSpaceModel"))
	return rv
}

func (c_ ColorSpace) ColorSyncProfile() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](c_, objc.GetSelector("colorSyncProfile"))
	return rv
}

func (c_ ColorSpace) ICCProfileData() []byte {
	rv := objc.CallMethod[[]byte](c_, objc.GetSelector("ICCProfileData"))
	return rv
}

func (c_ ColorSpace) LocalizedName() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("localizedName"))
	return rv
}

func (c_ ColorSpace) NumberOfColorComponents() int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("numberOfColorComponents"))
	return rv
}
