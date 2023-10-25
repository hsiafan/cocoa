// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ColorPickerClass = _ColorPickerClass{objc.GetClass("NSColorPicker")}

type _ColorPickerClass struct {
	objc.Class
}

type IColorPicker interface {
	objc.IObject
	InsertNewButtonImage_In(newButtonImage IImage, buttonCell IButtonCell)
	SetMode(mode ColorPanelMode)
	AttachColorList(colorList IColorList)
	DetachColorList(colorList IColorList)
	ViewSizeChanged(sender objc.IObject)
	ColorPanel() ColorPanel
	ProvideNewButtonImage() Image
	ButtonToolTip() string
	MinContentSize() foundation.Size
}

type ColorPicker struct {
	objc.Object
}

func MakeColorPicker(ptr unsafe.Pointer) ColorPicker {
	return ColorPicker{
		Object: objc.MakeObject(ptr),
	}
}

func (c_ ColorPicker) InitWithPickerMask_ColorPanel(mask uint, owningColorPanel IColorPanel) ColorPicker {
	rv := objc.CallMethod[ColorPicker](c_, objc.SEL("initWithPickerMask:colorPanel:"), mask, objc.ExtractPtr(owningColorPanel))
	return rv
}

func (cc _ColorPickerClass) Alloc() ColorPicker {
	rv := objc.CallMethod[ColorPicker](cc, objc.SEL("alloc"))
	return rv
}

func (cc _ColorPickerClass) New() ColorPicker {
	rv := objc.CallMethod[ColorPicker](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewColorPicker() ColorPicker {
	return ColorPickerClass.New()
}

func (c_ ColorPicker) Init() ColorPicker {
	rv := objc.CallMethod[ColorPicker](c_, objc.SEL("init"))
	return rv
}

func (c_ ColorPicker) InsertNewButtonImage_In(newButtonImage IImage, buttonCell IButtonCell) {
	objc.CallMethod[objc.Void](c_, objc.SEL("insertNewButtonImage:in:"), objc.ExtractPtr(newButtonImage), objc.ExtractPtr(buttonCell))
}

func (c_ ColorPicker) SetMode(mode ColorPanelMode) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setMode:"), mode)
}

func (c_ ColorPicker) AttachColorList(colorList IColorList) {
	objc.CallMethod[objc.Void](c_, objc.SEL("attachColorList:"), objc.ExtractPtr(colorList))
}

func (c_ ColorPicker) DetachColorList(colorList IColorList) {
	objc.CallMethod[objc.Void](c_, objc.SEL("detachColorList:"), objc.ExtractPtr(colorList))
}

func (c_ ColorPicker) ViewSizeChanged(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("viewSizeChanged:"), objc.ExtractPtr(sender))
}

func (c_ ColorPicker) ColorPanel() ColorPanel {
	rv := objc.CallMethod[ColorPanel](c_, objc.SEL("colorPanel"))
	return rv
}

func (c_ ColorPicker) ProvideNewButtonImage() Image {
	rv := objc.CallMethod[Image](c_, objc.SEL("provideNewButtonImage"))
	return rv
}

func (c_ ColorPicker) ButtonToolTip() string {
	rv := objc.CallMethod[string](c_, objc.SEL("buttonToolTip"))
	return rv
}

func (c_ ColorPicker) MinContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.SEL("minContentSize"))
	return rv
}
