// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
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
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
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
	rv := objc.CallMethod[Popover](p_, objc.SEL("init"))
	return rv
}

func (pc _PopoverClass) Alloc() Popover {
	rv := objc.CallMethod[Popover](pc, objc.SEL("alloc"))
	return rv
}

func (pc _PopoverClass) New() Popover {
	rv := objc.CallMethod[Popover](pc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewPopover() Popover {
	return PopoverClass.New()
}

func (p_ Popover) ShowRelativeToRect_OfView_PreferredEdge(positioningRect foundation.Rect, positioningView IView, preferredEdge foundation.RectEdge) {
	objc.CallMethod[objc.Void](p_, objc.SEL("showRelativeToRect:ofView:preferredEdge:"), positioningRect, objc.ExtractPtr(positioningView), preferredEdge)
}

func (p_ Popover) PerformClose(sender objc.IObject) {
	objc.CallMethod[objc.Void](p_, objc.SEL("performClose:"), objc.ExtractPtr(sender))
}

func (p_ Popover) Close() {
	objc.CallMethod[objc.Void](p_, objc.SEL("close"))
}

func (p_ Popover) ContentViewController() ViewController {
	rv := objc.CallMethod[ViewController](p_, objc.SEL("contentViewController"))
	return rv
}

func (p_ Popover) SetContentViewController(value IViewController) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setContentViewController:"), objc.ExtractPtr(value))
}

func (p_ Popover) Behavior() PopoverBehavior {
	rv := objc.CallMethod[PopoverBehavior](p_, objc.SEL("behavior"))
	return rv
}

func (p_ Popover) SetBehavior(value PopoverBehavior) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setBehavior:"), value)
}

func (p_ Popover) PositioningRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](p_, objc.SEL("positioningRect"))
	return rv
}

func (p_ Popover) SetPositioningRect(value foundation.Rect) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setPositioningRect:"), value)
}

func (p_ Popover) Appearance() Appearance {
	rv := objc.CallMethod[Appearance](p_, objc.SEL("appearance"))
	return rv
}

func (p_ Popover) SetAppearance(value IAppearance) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setAppearance:"), objc.ExtractPtr(value))
}

func (p_ Popover) EffectiveAppearance() Appearance {
	rv := objc.CallMethod[Appearance](p_, objc.SEL("effectiveAppearance"))
	return rv
}

func (p_ Popover) Animates() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("animates"))
	return rv
}

func (p_ Popover) SetAnimates(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setAnimates:"), value)
}

func (p_ Popover) ContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](p_, objc.SEL("contentSize"))
	return rv
}

func (p_ Popover) SetContentSize(value foundation.Size) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setContentSize:"), value)
}

func (p_ Popover) IsShown() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isShown"))
	return rv
}

func (p_ Popover) IsDetached() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isDetached"))
	return rv
}

// weak property
func (p_ Popover) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](p_, objc.SEL("delegate"))
	return rv
}

// weak property
func (p_ Popover) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}
