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
	rv := objc.CallMethod[ColorWell](cc, objc.SEL("colorWellWithStyle:"), style)
	return rv
}

func (c_ ColorWell) InitWithFrame(frameRect foundation.Rect) ColorWell {
	rv := objc.CallMethod[ColorWell](c_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (c_ ColorWell) Init() ColorWell {
	rv := objc.CallMethod[ColorWell](c_, objc.SEL("init"))
	return rv
}

func (cc _ColorWellClass) Alloc() ColorWell {
	rv := objc.CallMethod[ColorWell](cc, objc.SEL("alloc"))
	return rv
}

func (cc _ColorWellClass) New() ColorWell {
	rv := objc.CallMethod[ColorWell](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewColorWell() ColorWell {
	return ColorWellClass.New()
}

func (c_ ColorWell) TakeColorFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("takeColorFrom:"), objc.ExtractPtr(sender))
}

func (c_ ColorWell) Activate(exclusive bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("activate:"), exclusive)
}

func (c_ ColorWell) Deactivate() {
	objc.CallMethod[objc.Void](c_, objc.SEL("deactivate"))
}

func (c_ ColorWell) DrawWellInside(insideRect foundation.Rect) {
	objc.CallMethod[objc.Void](c_, objc.SEL("drawWellInside:"), insideRect)
}

func (c_ ColorWell) Color() Color {
	rv := objc.CallMethod[Color](c_, objc.SEL("color"))
	return rv
}

func (c_ ColorWell) SetColor(value IColor) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setColor:"), objc.ExtractPtr(value))
}

// weak property
func (c_ ColorWell) ColorWellStyle() ColorWellStyle {
	rv := objc.CallMethod[ColorWellStyle](c_, objc.SEL("colorWellStyle"))
	return rv
}

// weak property
func (c_ ColorWell) SetColorWellStyle(value ColorWellStyle) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setColorWellStyle:"), value)
}

func (c_ ColorWell) Image() Image {
	rv := objc.CallMethod[Image](c_, objc.SEL("image"))
	return rv
}

func (c_ ColorWell) SetImage(value IImage) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setImage:"), objc.ExtractPtr(value))
}

// deprecated
func (c_ ColorWell) IsBordered() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isBordered"))
	return rv
}

// deprecated
func (c_ ColorWell) SetBordered(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setBordered:"), value)
}

func (c_ ColorWell) IsActive() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isActive"))
	return rv
}

func (c_ ColorWell) PulldownAction() objc.Selector {
	rv := objc.CallMethod[objc.Selector](c_, objc.SEL("pulldownAction"))
	return rv
}

func (c_ ColorWell) SetPulldownAction(value objc.Selector) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setPulldownAction:"), value)
}

// weak property
func (c_ ColorWell) PulldownTarget() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.SEL("pulldownTarget"))
	return rv
}

// weak property
func (c_ ColorWell) SetPulldownTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setPulldownTarget:"), objc.ExtractPtr(value))
}
