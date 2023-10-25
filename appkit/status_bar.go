// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[StatusBar](sc, objc.SEL("alloc"))
	return rv
}

func (sc _StatusBarClass) New() StatusBar {
	rv := objc.CallMethod[StatusBar](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewStatusBar() StatusBar {
	return StatusBarClass.New()
}

func (s_ StatusBar) Init() StatusBar {
	rv := objc.CallMethod[StatusBar](s_, objc.SEL("init"))
	return rv
}

func (s_ StatusBar) StatusItemWithLength(length float64) StatusItem {
	rv := objc.CallMethod[StatusItem](s_, objc.SEL("statusItemWithLength:"), length)
	return rv
}

func (s_ StatusBar) RemoveStatusItem(item IStatusItem) {
	objc.CallMethod[objc.Void](s_, objc.SEL("removeStatusItem:"), objc.ExtractPtr(item))
}

func (sc _StatusBarClass) SystemStatusBar() StatusBar {
	rv := objc.CallMethod[StatusBar](sc, objc.SEL("systemStatusBar"))
	return rv
}

func (s_ StatusBar) IsVertical() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("isVertical"))
	return rv
}

func (s_ StatusBar) Thickness() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("thickness"))
	return rv
}
