// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ComboButtonClass = _ComboButtonClass{objc.GetClass("NSComboButton")}

type _ComboButtonClass struct {
	objc.Class
}

type IComboButton interface {
	IControl
	Style() ComboButtonStyle
	SetStyle(value ComboButtonStyle)
	Title() string
	SetTitle(value string)
	Image() Image
	SetImage(value IImage)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
}

type ComboButton struct {
	Control
}

func MakeComboButton(ptr unsafe.Pointer) ComboButton {
	return ComboButton{
		Control: MakeControl(ptr),
	}
}

func (cc _ComboButtonClass) ComboButtonWithTitle_Image_Menu_Target_Action(title string, image IImage, menu IMenu, target objc.IObject, action objc.Selector) ComboButton {
	rv := objc.CallMethod[ComboButton](cc, "comboButtonWithTitle:image:menu:target:action:", title, image, menu, target, action)
	return rv
}

func (cc _ComboButtonClass) ComboButtonWithTitle_Menu_Target_Action(title string, menu IMenu, target objc.IObject, action objc.Selector) ComboButton {
	rv := objc.CallMethod[ComboButton](cc, "comboButtonWithTitle:menu:target:action:", title, menu, target, action)
	return rv
}

func (cc _ComboButtonClass) ComboButtonWithImage_Menu_Target_Action(image IImage, menu IMenu, target objc.IObject, action objc.Selector) ComboButton {
	rv := objc.CallMethod[ComboButton](cc, "comboButtonWithImage:menu:target:action:", image, menu, target, action)
	return rv
}

func (c_ ComboButton) InitWithFrame(frameRect foundation.Rect) ComboButton {
	rv := objc.CallMethod[ComboButton](c_, "initWithFrame:", frameRect)
	return rv
}

func (c_ ComboButton) Init() ComboButton {
	rv := objc.CallMethod[ComboButton](c_, "init")
	return rv
}

func (cc _ComboButtonClass) Alloc() ComboButton {
	rv := objc.CallMethod[ComboButton](cc, "alloc")
	return rv
}

func (cc _ComboButtonClass) New() ComboButton {
	rv := objc.CallMethod[ComboButton](cc, "new")
	rv.Autorelease()
	return rv
}

func NewComboButton() ComboButton {
	return ComboButtonClass.New()
}

func (c_ ComboButton) Style() ComboButtonStyle {
	rv := objc.CallMethod[ComboButtonStyle](c_, "style")
	return rv
}

func (c_ ComboButton) SetStyle(value ComboButtonStyle) {
	objc.CallMethod[objc.Void](c_, "setStyle:", value)
}

func (c_ ComboButton) Title() string {
	rv := objc.CallMethod[string](c_, "title")
	return rv
}

func (c_ ComboButton) SetTitle(value string) {
	objc.CallMethod[objc.Void](c_, "setTitle:", value)
}

func (c_ ComboButton) Image() Image {
	rv := objc.CallMethod[Image](c_, "image")
	return rv
}

func (c_ ComboButton) SetImage(value IImage) {
	objc.CallMethod[objc.Void](c_, "setImage:", value)
}

func (c_ ComboButton) ImageScaling() ImageScaling {
	rv := objc.CallMethod[ImageScaling](c_, "imageScaling")
	return rv
}

func (c_ ComboButton) SetImageScaling(value ImageScaling) {
	objc.CallMethod[objc.Void](c_, "setImageScaling:", value)
}
