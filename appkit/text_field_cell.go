// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TextFieldCellClass = _TextFieldCellClass{objc.GetClass("NSTextFieldCell")}

type _TextFieldCellClass struct {
	objc.Class
}

type ITextFieldCell interface {
	IActionCell
	SetWantsNotificationForMarkedText(flag bool)
	TextColor() Color
	SetTextColor(value IColor)
	BezelStyle() TextFieldBezelStyle
	SetBezelStyle(value TextFieldBezelStyle)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	PlaceholderString() string
	SetPlaceholderString(value string)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.IAttributedString)
	AllowedInputSourceLocales() []string
	SetAllowedInputSourceLocales(value []string)
}

type TextFieldCell struct {
	ActionCell
}

func MakeTextFieldCell(ptr unsafe.Pointer) TextFieldCell {
	return TextFieldCell{
		ActionCell: MakeActionCell(ptr),
	}
}

func (t_ TextFieldCell) InitTextCell(string_ string) TextFieldCell {
	rv := ffi.CallMethod[TextFieldCell](t_, "initTextCell:", string_)
	return rv
}

func (t_ TextFieldCell) InitImageCell(image IImage) TextFieldCell {
	rv := ffi.CallMethod[TextFieldCell](t_, "initImageCell:", image)
	return rv
}

func (t_ TextFieldCell) Init() TextFieldCell {
	rv := ffi.CallMethod[TextFieldCell](t_, "init")
	return rv
}

func (tc _TextFieldCellClass) Alloc() TextFieldCell {
	rv := ffi.CallMethod[TextFieldCell](tc, "alloc")
	return rv
}

func (tc _TextFieldCellClass) New() TextFieldCell {
	rv := ffi.CallMethod[TextFieldCell](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextFieldCell() TextFieldCell {
	return TextFieldCellClass.New()
}

func (t_ TextFieldCell) SetWantsNotificationForMarkedText(flag bool) {
	ffi.CallMethod[ffi.Void](t_, "setWantsNotificationForMarkedText:", flag)
}

func (t_ TextFieldCell) TextColor() Color {
	rv := ffi.CallMethod[Color](t_, "textColor")
	return rv
}

func (t_ TextFieldCell) SetTextColor(value IColor) {
	ffi.CallMethod[ffi.Void](t_, "setTextColor:", value)
}

func (t_ TextFieldCell) BezelStyle() TextFieldBezelStyle {
	rv := ffi.CallMethod[TextFieldBezelStyle](t_, "bezelStyle")
	return rv
}

func (t_ TextFieldCell) SetBezelStyle(value TextFieldBezelStyle) {
	ffi.CallMethod[ffi.Void](t_, "setBezelStyle:", value)
}

func (t_ TextFieldCell) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](t_, "backgroundColor")
	return rv
}

func (t_ TextFieldCell) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](t_, "setBackgroundColor:", value)
}

func (t_ TextFieldCell) DrawsBackground() bool {
	rv := ffi.CallMethod[bool](t_, "drawsBackground")
	return rv
}

func (t_ TextFieldCell) SetDrawsBackground(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setDrawsBackground:", value)
}

func (t_ TextFieldCell) PlaceholderString() string {
	rv := ffi.CallMethod[string](t_, "placeholderString")
	return rv
}

func (t_ TextFieldCell) SetPlaceholderString(value string) {
	ffi.CallMethod[ffi.Void](t_, "setPlaceholderString:", value)
}

func (t_ TextFieldCell) PlaceholderAttributedString() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](t_, "placeholderAttributedString")
	return rv
}

func (t_ TextFieldCell) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	ffi.CallMethod[ffi.Void](t_, "setPlaceholderAttributedString:", value)
}

func (t_ TextFieldCell) AllowedInputSourceLocales() []string {
	rv := ffi.CallMethod[[]string](t_, "allowedInputSourceLocales")
	return rv
}

func (t_ TextFieldCell) SetAllowedInputSourceLocales(value []string) {
	ffi.CallMethod[ffi.Void](t_, "setAllowedInputSourceLocales:", value)
}
