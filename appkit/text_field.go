// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TextFieldClass = _TextFieldClass{objc.GetClass("NSTextField")}

type _TextFieldClass struct {
	objc.Class
}

type ITextField interface {
	IControl
	SelectText(sender objc.IObject)
	// deprecated
	SetTitleWithMnemonic(stringWithAmpersand string)
	TextShouldBeginEditing(textObject IText) bool
	TextDidBeginEditing(notification foundation.INotification)
	TextDidChange(notification foundation.INotification)
	TextShouldEndEditing(textObject IText) bool
	TextDidEndEditing(notification foundation.INotification)
	IsSelectable() bool
	SetSelectable(value bool)
	IsEditable() bool
	SetEditable(value bool)
	AllowsEditingTextAttributes() bool
	SetAllowsEditingTextAttributes(value bool)
	ImportsGraphics() bool
	SetImportsGraphics(value bool)
	PlaceholderString() string
	SetPlaceholderString(value string)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.IAttributedString)
	LineBreakStrategy() LineBreakStrategy
	SetLineBreakStrategy(value LineBreakStrategy)
	AllowsDefaultTighteningForTruncation() bool
	SetAllowsDefaultTighteningForTruncation(value bool)
	MaximumNumberOfLines() int
	SetMaximumNumberOfLines(value int)
	PreferredMaxLayoutWidth() float64
	SetPreferredMaxLayoutWidth(value float64)
	TextColor() Color
	SetTextColor(value IColor)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	IsBezeled() bool
	SetBezeled(value bool)
	BezelStyle() TextFieldBezelStyle
	SetBezelStyle(value TextFieldBezelStyle)
	IsBordered() bool
	SetBordered(value bool)
	AllowsCharacterPickerTouchBarItem() bool
	SetAllowsCharacterPickerTouchBarItem(value bool)
	IsAutomaticTextCompletionEnabled() bool
	SetAutomaticTextCompletionEnabled(value bool)
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
}

type TextField struct {
	Control
}

func MakeTextField(ptr unsafe.Pointer) TextField {
	return TextField{
		Control: MakeControl(ptr),
	}
}

func (tc _TextFieldClass) LabelWithAttributedString(attributedStringValue foundation.IAttributedString) TextField {
	rv := objc.CallMethod[TextField](tc, objc.SEL("labelWithAttributedString:"), objc.ExtractPtr(attributedStringValue))
	return rv
}

func (tc _TextFieldClass) LabelWithString(stringValue string) TextField {
	rv := objc.CallMethod[TextField](tc, objc.SEL("labelWithString:"), stringValue)
	return rv
}

func (tc _TextFieldClass) TextFieldWithString(stringValue string) TextField {
	rv := objc.CallMethod[TextField](tc, objc.SEL("textFieldWithString:"), stringValue)
	return rv
}

func (tc _TextFieldClass) WrappingLabelWithString(stringValue string) TextField {
	rv := objc.CallMethod[TextField](tc, objc.SEL("wrappingLabelWithString:"), stringValue)
	return rv
}

func (t_ TextField) InitWithFrame(frameRect foundation.Rect) TextField {
	rv := objc.CallMethod[TextField](t_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (t_ TextField) Init() TextField {
	rv := objc.CallMethod[TextField](t_, objc.SEL("init"))
	return rv
}

func (tc _TextFieldClass) Alloc() TextField {
	rv := objc.CallMethod[TextField](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TextFieldClass) New() TextField {
	rv := objc.CallMethod[TextField](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTextField() TextField {
	return TextFieldClass.New()
}

func (t_ TextField) SelectText(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("selectText:"), objc.ExtractPtr(sender))
}

// deprecated
func (t_ TextField) SetTitleWithMnemonic(stringWithAmpersand string) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTitleWithMnemonic:"), stringWithAmpersand)
}

func (t_ TextField) TextShouldBeginEditing(textObject IText) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("textShouldBeginEditing:"), objc.ExtractPtr(textObject))
	return rv
}

func (t_ TextField) TextDidBeginEditing(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.SEL("textDidBeginEditing:"), objc.ExtractPtr(notification))
}

func (t_ TextField) TextDidChange(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.SEL("textDidChange:"), objc.ExtractPtr(notification))
}

func (t_ TextField) TextShouldEndEditing(textObject IText) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("textShouldEndEditing:"), objc.ExtractPtr(textObject))
	return rv
}

func (t_ TextField) TextDidEndEditing(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.SEL("textDidEndEditing:"), objc.ExtractPtr(notification))
}

func (t_ TextField) IsSelectable() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isSelectable"))
	return rv
}

func (t_ TextField) SetSelectable(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSelectable:"), value)
}

func (t_ TextField) IsEditable() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isEditable"))
	return rv
}

func (t_ TextField) SetEditable(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setEditable:"), value)
}

func (t_ TextField) AllowsEditingTextAttributes() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("allowsEditingTextAttributes"))
	return rv
}

func (t_ TextField) SetAllowsEditingTextAttributes(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAllowsEditingTextAttributes:"), value)
}

func (t_ TextField) ImportsGraphics() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("importsGraphics"))
	return rv
}

func (t_ TextField) SetImportsGraphics(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setImportsGraphics:"), value)
}

func (t_ TextField) PlaceholderString() string {
	rv := objc.CallMethod[string](t_, objc.SEL("placeholderString"))
	return rv
}

func (t_ TextField) SetPlaceholderString(value string) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setPlaceholderString:"), value)
}

func (t_ TextField) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](t_, objc.SEL("placeholderAttributedString"))
	return rv
}

func (t_ TextField) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setPlaceholderAttributedString:"), objc.ExtractPtr(value))
}

func (t_ TextField) LineBreakStrategy() LineBreakStrategy {
	rv := objc.CallMethod[LineBreakStrategy](t_, objc.SEL("lineBreakStrategy"))
	return rv
}

func (t_ TextField) SetLineBreakStrategy(value LineBreakStrategy) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setLineBreakStrategy:"), value)
}

func (t_ TextField) AllowsDefaultTighteningForTruncation() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("allowsDefaultTighteningForTruncation"))
	return rv
}

func (t_ TextField) SetAllowsDefaultTighteningForTruncation(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAllowsDefaultTighteningForTruncation:"), value)
}

func (t_ TextField) MaximumNumberOfLines() int {
	rv := objc.CallMethod[int](t_, objc.SEL("maximumNumberOfLines"))
	return rv
}

func (t_ TextField) SetMaximumNumberOfLines(value int) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setMaximumNumberOfLines:"), value)
}

func (t_ TextField) PreferredMaxLayoutWidth() float64 {
	rv := objc.CallMethod[float64](t_, objc.SEL("preferredMaxLayoutWidth"))
	return rv
}

func (t_ TextField) SetPreferredMaxLayoutWidth(value float64) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setPreferredMaxLayoutWidth:"), value)
}

func (t_ TextField) TextColor() Color {
	rv := objc.CallMethod[Color](t_, objc.SEL("textColor"))
	return rv
}

func (t_ TextField) SetTextColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTextColor:"), objc.ExtractPtr(value))
}

func (t_ TextField) BackgroundColor() Color {
	rv := objc.CallMethod[Color](t_, objc.SEL("backgroundColor"))
	return rv
}

func (t_ TextField) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (t_ TextField) DrawsBackground() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("drawsBackground"))
	return rv
}

func (t_ TextField) SetDrawsBackground(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDrawsBackground:"), value)
}

func (t_ TextField) IsBezeled() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isBezeled"))
	return rv
}

func (t_ TextField) SetBezeled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setBezeled:"), value)
}

func (t_ TextField) BezelStyle() TextFieldBezelStyle {
	rv := objc.CallMethod[TextFieldBezelStyle](t_, objc.SEL("bezelStyle"))
	return rv
}

func (t_ TextField) SetBezelStyle(value TextFieldBezelStyle) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setBezelStyle:"), value)
}

func (t_ TextField) IsBordered() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isBordered"))
	return rv
}

func (t_ TextField) SetBordered(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setBordered:"), value)
}

func (t_ TextField) AllowsCharacterPickerTouchBarItem() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("allowsCharacterPickerTouchBarItem"))
	return rv
}

func (t_ TextField) SetAllowsCharacterPickerTouchBarItem(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAllowsCharacterPickerTouchBarItem:"), value)
}

func (t_ TextField) IsAutomaticTextCompletionEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isAutomaticTextCompletionEnabled"))
	return rv
}

func (t_ TextField) SetAutomaticTextCompletionEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAutomaticTextCompletionEnabled:"), value)
}

// weak property
func (t_ TextField) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.SEL("delegate"))
	return rv
}

// weak property
func (t_ TextField) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}
