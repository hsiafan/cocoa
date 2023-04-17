// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[Panel](pc, "windowWithContentViewController:", contentViewController)
	return rv
}

func (p_ Panel) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) Panel {
	rv := ffi.CallMethod[Panel](p_, "initWithContentRect:styleMask:backing:defer:", contentRect, style, backingStoreType, flag)
	return rv
}

func (p_ Panel) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) Panel {
	rv := ffi.CallMethod[Panel](p_, "initWithContentRect:styleMask:backing:defer:screen:", contentRect, style, backingStoreType, flag, screen)
	return rv
}

func (p_ Panel) Init() Panel {
	rv := ffi.CallMethod[Panel](p_, "init")
	return rv
}

func (pc _PanelClass) Alloc() Panel {
	rv := ffi.CallMethod[Panel](pc, "alloc")
	return rv
}

func (pc _PanelClass) New() Panel {
	rv := ffi.CallMethod[Panel](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPanel() Panel {
	return PanelClass.New()
}

func (p_ Panel) SetFloatingPanel(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setFloatingPanel:", value)
}

func (p_ Panel) BecomesKeyOnlyIfNeeded() bool {
	rv := ffi.CallMethod[bool](p_, "becomesKeyOnlyIfNeeded")
	return rv
}

func (p_ Panel) SetBecomesKeyOnlyIfNeeded(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setBecomesKeyOnlyIfNeeded:", value)
}

func (p_ Panel) SetWorksWhenModal(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setWorksWhenModal:", value)
}
