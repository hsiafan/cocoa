// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var StatusBarButtonClass = _StatusBarButtonClass{objc.GetClass("NSStatusBarButton")}

type _StatusBarButtonClass struct {
	objc.Class
}

type IStatusBarButton interface {
	IButton
	AppearsDisabled() bool
	SetAppearsDisabled(value bool)
}

type StatusBarButton struct {
	Button
}

func MakeStatusBarButton(ptr unsafe.Pointer) StatusBarButton {
	return StatusBarButton{
		Button: MakeButton(ptr),
	}
}

func (sc _StatusBarButtonClass) CheckboxWithTitle_Target_Action(title string, target objc.IObject, action objc.Selector) StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](sc, objc.GetSelector("checkboxWithTitle:target:action:"), title, target, action)
	return rv
}

func (sc _StatusBarButtonClass) ButtonWithImage_Target_Action(image IImage, target objc.IObject, action objc.Selector) StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](sc, objc.GetSelector("buttonWithImage:target:action:"), image, target, action)
	return rv
}

func (sc _StatusBarButtonClass) RadioButtonWithTitle_Target_Action(title string, target objc.IObject, action objc.Selector) StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](sc, objc.GetSelector("radioButtonWithTitle:target:action:"), title, target, action)
	return rv
}

func (sc _StatusBarButtonClass) ButtonWithTitle_Image_Target_Action(title string, image IImage, target objc.IObject, action objc.Selector) StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](sc, objc.GetSelector("buttonWithTitle:image:target:action:"), title, image, target, action)
	return rv
}

func (sc _StatusBarButtonClass) ButtonWithTitle_Target_Action(title string, target objc.IObject, action objc.Selector) StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](sc, objc.GetSelector("buttonWithTitle:target:action:"), title, target, action)
	return rv
}

func (s_ StatusBarButton) InitWithFrame(frameRect foundation.Rect) StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](s_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func (s_ StatusBarButton) Init() StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](s_, objc.GetSelector("init"))
	return rv
}

func (sc _StatusBarButtonClass) Alloc() StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](sc, objc.GetSelector("alloc"))
	return rv
}

func (sc _StatusBarButtonClass) New() StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewStatusBarButton() StatusBarButton {
	return StatusBarButtonClass.New()
}

func (s_ StatusBarButton) AppearsDisabled() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("appearsDisabled"))
	return rv
}

func (s_ StatusBarButton) SetAppearsDisabled(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAppearsDisabled:"), value)
}
