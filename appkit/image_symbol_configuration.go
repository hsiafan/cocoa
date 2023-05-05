// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var ImageSymbolConfigurationClass = _ImageSymbolConfigurationClass{objc.GetClass("NSImageSymbolConfiguration")}

type _ImageSymbolConfigurationClass struct {
	objc.Class
}

type IImageSymbolConfiguration interface {
	objc.IObject
}

type ImageSymbolConfiguration struct {
	objc.Object
}

func MakeImageSymbolConfiguration(ptr unsafe.Pointer) ImageSymbolConfiguration {
	return ImageSymbolConfiguration{
		Object: objc.MakeObject(ptr),
	}
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithPointSize_Weight(pointSize float64, weight FontWeight) ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationWithPointSize:weight:"), pointSize, weight)
	return rv
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithPointSize_Weight_Scale(pointSize float64, weight FontWeight, scale ImageSymbolScale) ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationWithPointSize:weight:scale:"), pointSize, weight, scale)
	return rv
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithTextStyle(style FontTextStyle) ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationWithTextStyle:"), style)
	return rv
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithTextStyle_Scale(style FontTextStyle, scale ImageSymbolScale) ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationWithTextStyle:scale:"), style, scale)
	return rv
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithScale(scale ImageSymbolScale) ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationWithScale:"), scale)
	return rv
}

func (i_ ImageSymbolConfiguration) ConfigurationByApplyingConfiguration(configuration IImageSymbolConfiguration) ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](i_, objc.GetSelector("configurationByApplyingConfiguration:"), objc.ExtractPtr(configuration))
	return rv
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithPaletteColors(paletteColors []IColor) ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationWithPaletteColors:"), paletteColors)
	return rv
}

func (ic _ImageSymbolConfigurationClass) ConfigurationWithHierarchicalColor(hierarchicalColor IColor) ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationWithHierarchicalColor:"), objc.ExtractPtr(hierarchicalColor))
	return rv
}

func (ic _ImageSymbolConfigurationClass) ConfigurationPreferringMulticolor() ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationPreferringMulticolor"))
	return rv
}

func (ic _ImageSymbolConfigurationClass) ConfigurationPreferringHierarchical() ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationPreferringHierarchical"))
	return rv
}

func (ic _ImageSymbolConfigurationClass) ConfigurationPreferringMonochrome() ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("configurationPreferringMonochrome"))
	return rv
}

func (ic _ImageSymbolConfigurationClass) Alloc() ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("alloc"))
	return rv
}

func (ic _ImageSymbolConfigurationClass) New() ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](ic, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewImageSymbolConfiguration() ImageSymbolConfiguration {
	return ImageSymbolConfigurationClass.New()
}

func (i_ ImageSymbolConfiguration) Init() ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](i_, objc.GetSelector("init"))
	return rv
}
