// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var PopoverClass = _PopoverClass{objc.GetClass("NSPopover")}

type _PopoverClass struct {
	objc.Class
}

type IPopover interface {
	IResponder
	ShowRelativeToRect_OfView_PreferredEdge(positioningRect foundation.Rect, positioningView IView, preferredEdge foundation.RectEdge)
	PerformClose(sender objc.IObject)
	Close()
	ContentViewController() ViewController
	SetContentViewController(value IViewController)
	Behavior() PopoverBehavior
	SetBehavior(value PopoverBehavior)
	PositioningRect() foundation.Rect
	SetPositioningRect(value foundation.Rect)
	Appearance() Appearance
	SetAppearance(value IAppearance)
	EffectiveAppearance() Appearance
	Animates() bool
	SetAnimates(value bool)
	ContentSize() foundation.Size
	SetContentSize(value foundation.Size)
	IsShown() bool
	IsDetached() bool
	Delegate() PopoverDelegateWrapper
	SetDelegate(value PopoverDelegate)
}

type Popover struct {
	Responder
}

func MakePopover(ptr unsafe.Pointer) Popover {
	return Popover{
		Responder: MakeResponder(ptr),
	}
}

func (p_ Popover) Init() Popover {
	rv := ffi.CallMethod[Popover](p_, "init")
	return rv
}

func (pc _PopoverClass) Alloc() Popover {
	rv := ffi.CallMethod[Popover](pc, "alloc")
	return rv
}

func (pc _PopoverClass) New() Popover {
	rv := ffi.CallMethod[Popover](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPopover() Popover {
	return PopoverClass.New()
}

func (p_ Popover) ShowRelativeToRect_OfView_PreferredEdge(positioningRect foundation.Rect, positioningView IView, preferredEdge foundation.RectEdge) {
	ffi.CallMethod[ffi.Void](p_, "showRelativeToRect:ofView:preferredEdge:", positioningRect, positioningView, preferredEdge)
}

func (p_ Popover) PerformClose(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](p_, "performClose:", sender)
}

func (p_ Popover) Close() {
	ffi.CallMethod[ffi.Void](p_, "close")
}

func (p_ Popover) ContentViewController() ViewController {
	rv := ffi.CallMethod[ViewController](p_, "contentViewController")
	return rv
}

func (p_ Popover) SetContentViewController(value IViewController) {
	ffi.CallMethod[ffi.Void](p_, "setContentViewController:", value)
}

func (p_ Popover) Behavior() PopoverBehavior {
	rv := ffi.CallMethod[PopoverBehavior](p_, "behavior")
	return rv
}

func (p_ Popover) SetBehavior(value PopoverBehavior) {
	ffi.CallMethod[ffi.Void](p_, "setBehavior:", value)
}

func (p_ Popover) PositioningRect() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](p_, "positioningRect")
	return rv
}

func (p_ Popover) SetPositioningRect(value foundation.Rect) {
	ffi.CallMethod[ffi.Void](p_, "setPositioningRect:", value)
}

func (p_ Popover) Appearance() Appearance {
	rv := ffi.CallMethod[Appearance](p_, "appearance")
	return rv
}

func (p_ Popover) SetAppearance(value IAppearance) {
	ffi.CallMethod[ffi.Void](p_, "setAppearance:", value)
}

func (p_ Popover) EffectiveAppearance() Appearance {
	rv := ffi.CallMethod[Appearance](p_, "effectiveAppearance")
	return rv
}

func (p_ Popover) Animates() bool {
	rv := ffi.CallMethod[bool](p_, "animates")
	return rv
}

func (p_ Popover) SetAnimates(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setAnimates:", value)
}

func (p_ Popover) ContentSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](p_, "contentSize")
	return rv
}

func (p_ Popover) SetContentSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](p_, "setContentSize:", value)
}

func (p_ Popover) IsShown() bool {
	rv := ffi.CallMethod[bool](p_, "isShown")
	return rv
}

func (p_ Popover) IsDetached() bool {
	rv := ffi.CallMethod[bool](p_, "isDetached")
	return rv
}

func (p_ Popover) Delegate() PopoverDelegateWrapper {
	rv := ffi.CallMethod[PopoverDelegateWrapper](p_, "delegate")
	return rv
}

func (p_ Popover) SetDelegate(value PopoverDelegate) {
	po := ffi.CreateProtocol("NSPopoverDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(p_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](p_, "setDelegate:", po)
}
