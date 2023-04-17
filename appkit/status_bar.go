// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var StatusBarClass = _StatusBarClass{objc.GetClass("NSStatusBar")}

type _StatusBarClass struct {
	objc.Class
}

type IStatusBar interface {
	objc.IObject
	StatusItemWithLength(length float64) StatusItem
	RemoveStatusItem(item IStatusItem)
	IsVertical() bool
	Thickness() float64
}

type StatusBar struct {
	objc.Object
}

func MakeStatusBar(ptr unsafe.Pointer) StatusBar {
	return StatusBar{
		Object: objc.MakeObject(ptr),
	}
}

func (sc _StatusBarClass) Alloc() StatusBar {
	rv := ffi.CallMethod[StatusBar](sc, "alloc")
	return rv
}

func (sc _StatusBarClass) New() StatusBar {
	rv := ffi.CallMethod[StatusBar](sc, "new")
	rv.Autorelease()
	return rv
}

func NewStatusBar() StatusBar {
	return StatusBarClass.New()
}

func (s_ StatusBar) Init() StatusBar {
	rv := ffi.CallMethod[StatusBar](s_, "init")
	return rv
}

func (s_ StatusBar) StatusItemWithLength(length float64) StatusItem {
	rv := ffi.CallMethod[StatusItem](s_, "statusItemWithLength:", length)
	return rv
}

func (s_ StatusBar) RemoveStatusItem(item IStatusItem) {
	ffi.CallMethod[ffi.Void](s_, "removeStatusItem:", item)
}

func (sc _StatusBarClass) SystemStatusBar() StatusBar {
	rv := ffi.CallMethod[StatusBar](sc, "systemStatusBar")
	return rv
}

func (s_ StatusBar) IsVertical() bool {
	rv := ffi.CallMethod[bool](s_, "isVertical")
	return rv
}

func (s_ StatusBar) Thickness() float64 {
	rv := ffi.CallMethod[float64](s_, "thickness")
	return rv
}
