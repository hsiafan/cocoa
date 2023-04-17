// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[ColorPicker](c_, "initWithPickerMask:colorPanel:", mask, owningColorPanel)
	return rv
}

func (cc _ColorPickerClass) Alloc() ColorPicker {
	rv := ffi.CallMethod[ColorPicker](cc, "alloc")
	return rv
}

func (cc _ColorPickerClass) New() ColorPicker {
	rv := ffi.CallMethod[ColorPicker](cc, "new")
	rv.Autorelease()
	return rv
}

func NewColorPicker() ColorPicker {
	return ColorPickerClass.New()
}

func (c_ ColorPicker) Init() ColorPicker {
	rv := ffi.CallMethod[ColorPicker](c_, "init")
	return rv
}

func (c_ ColorPicker) InsertNewButtonImage_In(newButtonImage IImage, buttonCell IButtonCell) {
	ffi.CallMethod[ffi.Void](c_, "insertNewButtonImage:in:", newButtonImage, buttonCell)
}

func (c_ ColorPicker) SetMode(mode ColorPanelMode) {
	ffi.CallMethod[ffi.Void](c_, "setMode:", mode)
}

func (c_ ColorPicker) AttachColorList(colorList IColorList) {
	ffi.CallMethod[ffi.Void](c_, "attachColorList:", colorList)
}

func (c_ ColorPicker) DetachColorList(colorList IColorList) {
	ffi.CallMethod[ffi.Void](c_, "detachColorList:", colorList)
}

func (c_ ColorPicker) ViewSizeChanged(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "viewSizeChanged:", sender)
}

func (c_ ColorPicker) ColorPanel() ColorPanel {
	rv := ffi.CallMethod[ColorPanel](c_, "colorPanel")
	return rv
}

func (c_ ColorPicker) ProvideNewButtonImage() Image {
	rv := ffi.CallMethod[Image](c_, "provideNewButtonImage")
	return rv
}

func (c_ ColorPicker) ButtonToolTip() string {
	rv := ffi.CallMethod[string](c_, "buttonToolTip")
	return rv
}

func (c_ ColorPicker) MinContentSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "minContentSize")
	return rv
}
