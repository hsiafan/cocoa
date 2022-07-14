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
	rv.Autorelease()
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
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(p_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](p_, "setDelegate:", po)
}

type PopoverDelegate interface {
	ImplementsDetachableWindowForPopover() bool
	// optional
	DetachableWindowForPopover(popover Popover) IWindow
	ImplementsPopoverShouldClose() bool
	// optional
	PopoverShouldClose(popover Popover) bool
	ImplementsPopoverWillShow() bool
	// optional
	PopoverWillShow(notification foundation.Notification)
	ImplementsPopoverDidShow() bool
	// optional
	PopoverDidShow(notification foundation.Notification)
	ImplementsPopoverWillClose() bool
	// optional
	PopoverWillClose(notification foundation.Notification)
	ImplementsPopoverDidClose() bool
	// optional
	PopoverDidClose(notification foundation.Notification)
	ImplementsPopoverDidDetach() bool
	// optional
	PopoverDidDetach(popover Popover)
	ImplementsPopoverShouldDetach() bool
	// optional
	PopoverShouldDetach(popover Popover) bool
}

type PopoverDelegateImpl struct {
	_DetachableWindowForPopover func(popover Popover) IWindow
	_PopoverShouldClose         func(popover Popover) bool
	_PopoverWillShow            func(notification foundation.Notification)
	_PopoverDidShow             func(notification foundation.Notification)
	_PopoverWillClose           func(notification foundation.Notification)
	_PopoverDidClose            func(notification foundation.Notification)
	_PopoverDidDetach           func(popover Popover)
	_PopoverShouldDetach        func(popover Popover) bool
}

func (di *PopoverDelegateImpl) ImplementsDetachableWindowForPopover() bool {
	return di._DetachableWindowForPopover != nil
}

func (di *PopoverDelegateImpl) SetDetachableWindowForPopover(f func(popover Popover) IWindow) {
	di._DetachableWindowForPopover = f
}

func (di *PopoverDelegateImpl) DetachableWindowForPopover(popover Popover) IWindow {
	return di._DetachableWindowForPopover(popover)
}
func (di *PopoverDelegateImpl) ImplementsPopoverShouldClose() bool {
	return di._PopoverShouldClose != nil
}

func (di *PopoverDelegateImpl) SetPopoverShouldClose(f func(popover Popover) bool) {
	di._PopoverShouldClose = f
}

func (di *PopoverDelegateImpl) PopoverShouldClose(popover Popover) bool {
	return di._PopoverShouldClose(popover)
}
func (di *PopoverDelegateImpl) ImplementsPopoverWillShow() bool {
	return di._PopoverWillShow != nil
}

func (di *PopoverDelegateImpl) SetPopoverWillShow(f func(notification foundation.Notification)) {
	di._PopoverWillShow = f
}

func (di *PopoverDelegateImpl) PopoverWillShow(notification foundation.Notification) {
	di._PopoverWillShow(notification)
}
func (di *PopoverDelegateImpl) ImplementsPopoverDidShow() bool {
	return di._PopoverDidShow != nil
}

func (di *PopoverDelegateImpl) SetPopoverDidShow(f func(notification foundation.Notification)) {
	di._PopoverDidShow = f
}

func (di *PopoverDelegateImpl) PopoverDidShow(notification foundation.Notification) {
	di._PopoverDidShow(notification)
}
func (di *PopoverDelegateImpl) ImplementsPopoverWillClose() bool {
	return di._PopoverWillClose != nil
}

func (di *PopoverDelegateImpl) SetPopoverWillClose(f func(notification foundation.Notification)) {
	di._PopoverWillClose = f
}

func (di *PopoverDelegateImpl) PopoverWillClose(notification foundation.Notification) {
	di._PopoverWillClose(notification)
}
func (di *PopoverDelegateImpl) ImplementsPopoverDidClose() bool {
	return di._PopoverDidClose != nil
}

func (di *PopoverDelegateImpl) SetPopoverDidClose(f func(notification foundation.Notification)) {
	di._PopoverDidClose = f
}

func (di *PopoverDelegateImpl) PopoverDidClose(notification foundation.Notification) {
	di._PopoverDidClose(notification)
}
func (di *PopoverDelegateImpl) ImplementsPopoverDidDetach() bool {
	return di._PopoverDidDetach != nil
}

func (di *PopoverDelegateImpl) SetPopoverDidDetach(f func(popover Popover)) {
	di._PopoverDidDetach = f
}

func (di *PopoverDelegateImpl) PopoverDidDetach(popover Popover) {
	di._PopoverDidDetach(popover)
}
func (di *PopoverDelegateImpl) ImplementsPopoverShouldDetach() bool {
	return di._PopoverShouldDetach != nil
}

func (di *PopoverDelegateImpl) SetPopoverShouldDetach(f func(popover Popover) bool) {
	di._PopoverShouldDetach = f
}

func (di *PopoverDelegateImpl) PopoverShouldDetach(popover Popover) bool {
	return di._PopoverShouldDetach(popover)
}

type PopoverDelegateWrapper struct {
	objc.Object
}

func (p_ *PopoverDelegateWrapper) ImplementsDetachableWindowForPopover() bool {
	return p_.RespondsToSelector(objc.GetSelector("detachableWindowForPopover:"))
}

func (p_ PopoverDelegateWrapper) DetachableWindowForPopover(popover IPopover) Window {
	rv := ffi.CallMethod[Window](p_, "detachableWindowForPopover:", popover)
	return rv
}

func (p_ *PopoverDelegateWrapper) ImplementsPopoverShouldClose() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverShouldClose:"))
}

func (p_ PopoverDelegateWrapper) PopoverShouldClose(popover IPopover) bool {
	rv := ffi.CallMethod[bool](p_, "popoverShouldClose:", popover)
	return rv
}

func (p_ *PopoverDelegateWrapper) ImplementsPopoverWillShow() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverWillShow:"))
}

func (p_ PopoverDelegateWrapper) PopoverWillShow(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](p_, "popoverWillShow:", notification)
}

func (p_ *PopoverDelegateWrapper) ImplementsPopoverDidShow() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverDidShow:"))
}

func (p_ PopoverDelegateWrapper) PopoverDidShow(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](p_, "popoverDidShow:", notification)
}

func (p_ *PopoverDelegateWrapper) ImplementsPopoverWillClose() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverWillClose:"))
}

func (p_ PopoverDelegateWrapper) PopoverWillClose(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](p_, "popoverWillClose:", notification)
}

func (p_ *PopoverDelegateWrapper) ImplementsPopoverDidClose() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverDidClose:"))
}

func (p_ PopoverDelegateWrapper) PopoverDidClose(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](p_, "popoverDidClose:", notification)
}

func (p_ *PopoverDelegateWrapper) ImplementsPopoverDidDetach() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverDidDetach:"))
}

func (p_ PopoverDelegateWrapper) PopoverDidDetach(popover IPopover) {
	ffi.CallMethod[ffi.Void](p_, "popoverDidDetach:", popover)
}

func (p_ *PopoverDelegateWrapper) ImplementsPopoverShouldDetach() bool {
	return p_.RespondsToSelector(objc.GetSelector("popoverShouldDetach:"))
}

func (p_ PopoverDelegateWrapper) PopoverShouldDetach(popover IPopover) bool {
	rv := ffi.CallMethod[bool](p_, "popoverShouldDetach:", popover)
	return rv
}
