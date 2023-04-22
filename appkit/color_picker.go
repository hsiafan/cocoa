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
	rv := objc.CallMethod[ColorPicker](c_, objc.GetSelector("initWithPickerMask:colorPanel:"), mask, owningColorPanel)
	return rv
}

func (cc _ColorPickerClass) Alloc() ColorPicker {
	rv := objc.CallMethod[ColorPicker](cc, objc.GetSelector("alloc"))
	return rv
}

func (cc _ColorPickerClass) New() ColorPicker {
	rv := objc.CallMethod[ColorPicker](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewColorPicker() ColorPicker {
	return ColorPickerClass.New()
}

func (c_ ColorPicker) Init() ColorPicker {
	rv := objc.CallMethod[ColorPicker](c_, objc.GetSelector("init"))
	return rv
}

func (c_ ColorPicker) InsertNewButtonImage_In(newButtonImage IImage, buttonCell IButtonCell) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("insertNewButtonImage:in:"), newButtonImage, buttonCell)
}

func (c_ ColorPicker) SetMode(mode ColorPanelMode) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setMode:"), mode)
}

func (c_ ColorPicker) AttachColorList(colorList IColorList) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("attachColorList:"), colorList)
}

func (c_ ColorPicker) DetachColorList(colorList IColorList) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("detachColorList:"), colorList)
}

func (c_ ColorPicker) ViewSizeChanged(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("viewSizeChanged:"), sender)
}

func (c_ ColorPicker) ColorPanel() ColorPanel {
	rv := objc.CallMethod[ColorPanel](c_, objc.GetSelector("colorPanel"))
	return rv
}

func (c_ ColorPicker) ProvideNewButtonImage() Image {
	rv := objc.CallMethod[Image](c_, objc.GetSelector("provideNewButtonImage"))
	return rv
}

func (c_ ColorPicker) ButtonToolTip() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("buttonToolTip"))
	return rv
}

func (c_ ColorPicker) MinContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("minContentSize"))
	return rv
}
