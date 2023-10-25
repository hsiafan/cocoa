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
	rv := objc.CallMethod[ColorSpace](c_, objc.SEL("initWithCGColorSpace:"), cgColorSpace)
	return rv
}

func (c_ ColorSpace) InitWithColorSyncProfile(prof unsafe.Pointer) ColorSpace {
	rv := objc.CallMethod[ColorSpace](c_, objc.SEL("initWithColorSyncProfile:"), prof)
	return rv
}

func (c_ ColorSpace) InitWithICCProfileData(iccData []byte) ColorSpace {
	rv := objc.CallMethod[ColorSpace](c_, objc.SEL("initWithICCProfileData:"), iccData)
	return rv
}

func (cc _ColorSpaceClass) Alloc() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.SEL("alloc"))
	return rv
}

func (cc _ColorSpaceClass) New() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewColorSpace() ColorSpace {
	return ColorSpaceClass.New()
}

func (c_ ColorSpace) Init() ColorSpace {
	rv := objc.CallMethod[ColorSpace](c_, objc.SEL("init"))
	return rv
}

func (cc _ColorSpaceClass) AvailableColorSpacesWithModel(model ColorSpaceModel) []ColorSpace {
	rv := objc.CallMethod[[]ColorSpace](cc, objc.SEL("availableColorSpacesWithModel:"), model)
	return rv
}

func (cc _ColorSpaceClass) DeviceRGBColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.SEL("deviceRGBColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) GenericRGBColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.SEL("genericRGBColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) DeviceCMYKColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.SEL("deviceCMYKColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) GenericCMYKColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.SEL("genericCMYKColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) DeviceGrayColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.SEL("deviceGrayColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) GenericGrayColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.SEL("genericGrayColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) SRGBColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.SEL("sRGBColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) ExtendedSRGBColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.SEL("extendedSRGBColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) DisplayP3ColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.SEL("displayP3ColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) GenericGamma22GrayColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.SEL("genericGamma22GrayColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) ExtendedGenericGamma22GrayColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.SEL("extendedGenericGamma22GrayColorSpace"))
	return rv
}

func (cc _ColorSpaceClass) AdobeRGB1998ColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](cc, objc.SEL("adobeRGB1998ColorSpace"))
	return rv
}

func (c_ ColorSpace) CGColorSpace() coregraphics.ColorSpaceRef {
	rv := objc.CallMethod[coregraphics.ColorSpaceRef](c_, objc.SEL("CGColorSpace"))
	return rv
}

func (c_ ColorSpace) ColorSpaceModel() ColorSpaceModel {
	rv := objc.CallMethod[ColorSpaceModel](c_, objc.SEL("colorSpaceModel"))
	return rv
}

func (c_ ColorSpace) ColorSyncProfile() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](c_, objc.SEL("colorSyncProfile"))
	return rv
}

func (c_ ColorSpace) ICCProfileData() []byte {
	rv := objc.CallMethod[[]byte](c_, objc.SEL("ICCProfileData"))
	return rv
}

func (c_ ColorSpace) LocalizedName() string {
	rv := objc.CallMethod[string](c_, objc.SEL("localizedName"))
	return rv
}

func (c_ ColorSpace) NumberOfColorComponents() int {
	rv := objc.CallMethod[int](c_, objc.SEL("numberOfColorComponents"))
	return rv
}
