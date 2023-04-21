// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ColorWellClass = _ColorWellClass{objc.GetClass("NSColorWell")}

type _ColorWellClass struct {
	objc.Class
}

type IColorWell interface {
	IControl
	TakeColorFrom(sender objc.IObject)
	Activate(exclusive bool)
	Deactivate()
	DrawWellInside(insideRect foundation.Rect)
	Color() Color
	SetColor(value IColor)
	ColorWellStyle() ColorWellStyle
	SetColorWellStyle(value ColorWellStyle)
	Image() Image
	SetImage(value IImage)
	// deprecated
	IsBordered() bool
	// deprecated
	SetBordered(value bool)
	IsActive() bool
	PulldownAction() objc.Selector
	SetPulldownAction(value objc.Selector)
	PulldownTarget() objc.Object
	SetPulldownTarget(value objc.IObject)
}

type ColorWell struct {
	Control
}

func MakeColorWell(ptr unsafe.Pointer) ColorWell {
	return ColorWell{
		Control: MakeControl(ptr),
	}
}

func (cc _ColorWellClass) ColorWellWithStyle(style ColorWellStyle) ColorWell {
	rv := objc.CallMethod[ColorWell](cc, "colorWellWithStyle:", style)
	return rv
}

func (c_ ColorWell) InitWithFrame(frameRect foundation.Rect) ColorWell {
	rv := objc.CallMethod[ColorWell](c_, "initWithFrame:", frameRect)
	return rv
}

func (c_ ColorWell) Init() ColorWell {
	rv := objc.CallMethod[ColorWell](c_, "init")
	return rv
}

func (cc _ColorWellClass) Alloc() ColorWell {
	rv := objc.CallMethod[ColorWell](cc, "alloc")
	return rv
}

func (cc _ColorWellClass) New() ColorWell {
	rv := objc.CallMethod[ColorWell](cc, "new")
	rv.Autorelease()
	return rv
}

func NewColorWell() ColorWell {
	return ColorWellClass.New()
}

func (c_ ColorWell) TakeColorFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, "takeColorFrom:", sender)
}

func (c_ ColorWell) Activate(exclusive bool) {
	objc.CallMethod[objc.Void](c_, "activate:", exclusive)
}

func (c_ ColorWell) Deactivate() {
	objc.CallMethod[objc.Void](c_, "deactivate")
}

func (c_ ColorWell) DrawWellInside(insideRect foundation.Rect) {
	objc.CallMethod[objc.Void](c_, "drawWellInside:", insideRect)
}

func (c_ ColorWell) Color() Color {
	rv := objc.CallMethod[Color](c_, "color")
	return rv
}

func (c_ ColorWell) SetColor(value IColor) {
	objc.CallMethod[objc.Void](c_, "setColor:", value)
}

func (c_ ColorWell) ColorWellStyle() ColorWellStyle {
	rv := objc.CallMethod[ColorWellStyle](c_, "colorWellStyle")
	return rv
}

func (c_ ColorWell) SetColorWellStyle(value ColorWellStyle) {
	objc.CallMethod[objc.Void](c_, "setColorWellStyle:", value)
}

func (c_ ColorWell) Image() Image {
	rv := objc.CallMethod[Image](c_, "image")
	return rv
}

func (c_ ColorWell) SetImage(value IImage) {
	objc.CallMethod[objc.Void](c_, "setImage:", value)
}

// deprecated
func (c_ ColorWell) IsBordered() bool {
	rv := objc.CallMethod[bool](c_, "isBordered")
	return rv
}

// deprecated
func (c_ ColorWell) SetBordered(value bool) {
	objc.CallMethod[objc.Void](c_, "setBordered:", value)
}

func (c_ ColorWell) IsActive() bool {
	rv := objc.CallMethod[bool](c_, "isActive")
	return rv
}

func (c_ ColorWell) PulldownAction() objc.Selector {
	rv := objc.CallMethod[objc.Selector](c_, "pulldownAction")
	return rv
}

func (c_ ColorWell) SetPulldownAction(value objc.Selector) {
	objc.CallMethod[objc.Void](c_, "setPulldownAction:", value)
}

func (c_ ColorWell) PulldownTarget() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, "pulldownTarget")
	return rv
}

func (c_ ColorWell) SetPulldownTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, "setPulldownTarget:", value)
}
