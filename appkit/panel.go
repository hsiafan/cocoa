// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var PanelClass = _PanelClass{objc.GetClass("NSPanel")}

type _PanelClass struct {
	objc.Class
}

type IPanel interface {
	IWindow
	SetFloatingPanel(value bool)
	BecomesKeyOnlyIfNeeded() bool
	SetBecomesKeyOnlyIfNeeded(value bool)
	SetWorksWhenModal(value bool)
}

type Panel struct {
	Window
}

func MakePanel(ptr unsafe.Pointer) Panel {
	return Panel{
		Window: MakeWindow(ptr),
	}
}

func (pc _PanelClass) WindowWithContentViewController(contentViewController IViewController) Panel {
	rv := objc.CallMethod[Panel](pc, objc.GetSelector("windowWithContentViewController:"), contentViewController)
	return rv
}

func (p_ Panel) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) Panel {
	rv := objc.CallMethod[Panel](p_, objc.GetSelector("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return rv
}

func (p_ Panel) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) Panel {
	rv := objc.CallMethod[Panel](p_, objc.GetSelector("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, screen)
	return rv
}

func (p_ Panel) Init() Panel {
	rv := objc.CallMethod[Panel](p_, objc.GetSelector("init"))
	return rv
}

func (pc _PanelClass) Alloc() Panel {
	rv := objc.CallMethod[Panel](pc, objc.GetSelector("alloc"))
	return rv
}

func (pc _PanelClass) New() Panel {
	rv := objc.CallMethod[Panel](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPanel() Panel {
	return PanelClass.New()
}

func (p_ Panel) SetFloatingPanel(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setFloatingPanel:"), value)
}

func (p_ Panel) BecomesKeyOnlyIfNeeded() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("becomesKeyOnlyIfNeeded"))
	return rv
}

func (p_ Panel) SetBecomesKeyOnlyIfNeeded(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setBecomesKeyOnlyIfNeeded:"), value)
}

func (p_ Panel) SetWorksWhenModal(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setWorksWhenModal:"), value)
}
