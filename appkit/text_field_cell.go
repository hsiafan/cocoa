// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[TextFieldCell](t_, objc.SEL("initTextCell:"), string_)
	return rv
}

func (t_ TextFieldCell) InitImageCell(image IImage) TextFieldCell {
	rv := objc.CallMethod[TextFieldCell](t_, objc.SEL("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func (t_ TextFieldCell) Init() TextFieldCell {
	rv := objc.CallMethod[TextFieldCell](t_, objc.SEL("init"))
	return rv
}

func (tc _TextFieldCellClass) Alloc() TextFieldCell {
	rv := objc.CallMethod[TextFieldCell](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TextFieldCellClass) New() TextFieldCell {
	rv := objc.CallMethod[TextFieldCell](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTextFieldCell() TextFieldCell {
	return TextFieldCellClass.New()
}

func (t_ TextFieldCell) SetWantsNotificationForMarkedText(flag bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setWantsNotificationForMarkedText:"), flag)
}

func (t_ TextFieldCell) TextColor() Color {
	rv := objc.CallMethod[Color](t_, objc.SEL("textColor"))
	return rv
}

func (t_ TextFieldCell) SetTextColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTextColor:"), objc.ExtractPtr(value))
}

func (t_ TextFieldCell) BezelStyle() TextFieldBezelStyle {
	rv := objc.CallMethod[TextFieldBezelStyle](t_, objc.SEL("bezelStyle"))
	return rv
}

func (t_ TextFieldCell) SetBezelStyle(value TextFieldBezelStyle) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setBezelStyle:"), value)
}

func (t_ TextFieldCell) BackgroundColor() Color {
	rv := objc.CallMethod[Color](t_, objc.SEL("backgroundColor"))
	return rv
}

func (t_ TextFieldCell) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (t_ TextFieldCell) DrawsBackground() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("drawsBackground"))
	return rv
}

func (t_ TextFieldCell) SetDrawsBackground(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDrawsBackground:"), value)
}

func (t_ TextFieldCell) PlaceholderString() string {
	rv := objc.CallMethod[string](t_, objc.SEL("placeholderString"))
	return rv
}

func (t_ TextFieldCell) SetPlaceholderString(value string) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setPlaceholderString:"), value)
}

func (t_ TextFieldCell) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](t_, objc.SEL("placeholderAttributedString"))
	return rv
}

func (t_ TextFieldCell) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setPlaceholderAttributedString:"), objc.ExtractPtr(value))
}

func (t_ TextFieldCell) AllowedInputSourceLocales() []string {
	rv := objc.CallMethod[[]string](t_, objc.SEL("allowedInputSourceLocales"))
	return rv
}

func (t_ TextFieldCell) SetAllowedInputSourceLocales(value []string) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAllowedInputSourceLocales:"), value)
}
