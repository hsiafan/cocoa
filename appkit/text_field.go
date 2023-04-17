// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
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
	Delegate() TextFieldDelegateWrapper
	SetDelegate(value TextFieldDelegate)
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
	rv := ffi.CallMethod[TextField](tc, "labelWithAttributedString:", attributedStringValue)
	return rv
}

func (tc _TextFieldClass) LabelWithString(stringValue string) TextField {
	rv := ffi.CallMethod[TextField](tc, "labelWithString:", stringValue)
	return rv
}

func (tc _TextFieldClass) TextFieldWithString(stringValue string) TextField {
	rv := ffi.CallMethod[TextField](tc, "textFieldWithString:", stringValue)
	return rv
}

func (tc _TextFieldClass) WrappingLabelWithString(stringValue string) TextField {
	rv := ffi.CallMethod[TextField](tc, "wrappingLabelWithString:", stringValue)
	return rv
}

func (t_ TextField) InitWithFrame(frameRect foundation.Rect) TextField {
	rv := ffi.CallMethod[TextField](t_, "initWithFrame:", frameRect)
	return rv
}

func (t_ TextField) Init() TextField {
	rv := ffi.CallMethod[TextField](t_, "init")
	return rv
}

func (tc _TextFieldClass) Alloc() TextField {
	rv := ffi.CallMethod[TextField](tc, "alloc")
	return rv
}

func (tc _TextFieldClass) New() TextField {
	rv := ffi.CallMethod[TextField](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextField() TextField {
	return TextFieldClass.New()
}

func (t_ TextField) SelectText(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "selectText:", sender)
}

// deprecated
func (t_ TextField) SetTitleWithMnemonic(stringWithAmpersand string) {
	ffi.CallMethod[ffi.Void](t_, "setTitleWithMnemonic:", stringWithAmpersand)
}

func (t_ TextField) TextShouldBeginEditing(textObject IText) bool {
	rv := ffi.CallMethod[bool](t_, "textShouldBeginEditing:", textObject)
	return rv
}

func (t_ TextField) TextDidBeginEditing(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](t_, "textDidBeginEditing:", notification)
}

func (t_ TextField) TextDidChange(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](t_, "textDidChange:", notification)
}

func (t_ TextField) TextShouldEndEditing(textObject IText) bool {
	rv := ffi.CallMethod[bool](t_, "textShouldEndEditing:", textObject)
	return rv
}

func (t_ TextField) TextDidEndEditing(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](t_, "textDidEndEditing:", notification)
}

func (t_ TextField) IsSelectable() bool {
	rv := ffi.CallMethod[bool](t_, "isSelectable")
	return rv
}

func (t_ TextField) SetSelectable(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setSelectable:", value)
}

func (t_ TextField) IsEditable() bool {
	rv := ffi.CallMethod[bool](t_, "isEditable")
	return rv
}

func (t_ TextField) SetEditable(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setEditable:", value)
}

func (t_ TextField) AllowsEditingTextAttributes() bool {
	rv := ffi.CallMethod[bool](t_, "allowsEditingTextAttributes")
	return rv
}

func (t_ TextField) SetAllowsEditingTextAttributes(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAllowsEditingTextAttributes:", value)
}

func (t_ TextField) ImportsGraphics() bool {
	rv := ffi.CallMethod[bool](t_, "importsGraphics")
	return rv
}

func (t_ TextField) SetImportsGraphics(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setImportsGraphics:", value)
}

func (t_ TextField) PlaceholderString() string {
	rv := ffi.CallMethod[string](t_, "placeholderString")
	return rv
}

func (t_ TextField) SetPlaceholderString(value string) {
	ffi.CallMethod[ffi.Void](t_, "setPlaceholderString:", value)
}

func (t_ TextField) PlaceholderAttributedString() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](t_, "placeholderAttributedString")
	return rv
}

func (t_ TextField) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	ffi.CallMethod[ffi.Void](t_, "setPlaceholderAttributedString:", value)
}

func (t_ TextField) LineBreakStrategy() LineBreakStrategy {
	rv := ffi.CallMethod[LineBreakStrategy](t_, "lineBreakStrategy")
	return rv
}

func (t_ TextField) SetLineBreakStrategy(value LineBreakStrategy) {
	ffi.CallMethod[ffi.Void](t_, "setLineBreakStrategy:", value)
}

func (t_ TextField) AllowsDefaultTighteningForTruncation() bool {
	rv := ffi.CallMethod[bool](t_, "allowsDefaultTighteningForTruncation")
	return rv
}

func (t_ TextField) SetAllowsDefaultTighteningForTruncation(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAllowsDefaultTighteningForTruncation:", value)
}

func (t_ TextField) MaximumNumberOfLines() int {
	rv := ffi.CallMethod[int](t_, "maximumNumberOfLines")
	return rv
}

func (t_ TextField) SetMaximumNumberOfLines(value int) {
	ffi.CallMethod[ffi.Void](t_, "setMaximumNumberOfLines:", value)
}

func (t_ TextField) PreferredMaxLayoutWidth() float64 {
	rv := ffi.CallMethod[float64](t_, "preferredMaxLayoutWidth")
	return rv
}

func (t_ TextField) SetPreferredMaxLayoutWidth(value float64) {
	ffi.CallMethod[ffi.Void](t_, "setPreferredMaxLayoutWidth:", value)
}

func (t_ TextField) TextColor() Color {
	rv := ffi.CallMethod[Color](t_, "textColor")
	return rv
}

func (t_ TextField) SetTextColor(value IColor) {
	ffi.CallMethod[ffi.Void](t_, "setTextColor:", value)
}

func (t_ TextField) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](t_, "backgroundColor")
	return rv
}

func (t_ TextField) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](t_, "setBackgroundColor:", value)
}

func (t_ TextField) DrawsBackground() bool {
	rv := ffi.CallMethod[bool](t_, "drawsBackground")
	return rv
}

func (t_ TextField) SetDrawsBackground(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setDrawsBackground:", value)
}

func (t_ TextField) IsBezeled() bool {
	rv := ffi.CallMethod[bool](t_, "isBezeled")
	return rv
}

func (t_ TextField) SetBezeled(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setBezeled:", value)
}

func (t_ TextField) BezelStyle() TextFieldBezelStyle {
	rv := ffi.CallMethod[TextFieldBezelStyle](t_, "bezelStyle")
	return rv
}

func (t_ TextField) SetBezelStyle(value TextFieldBezelStyle) {
	ffi.CallMethod[ffi.Void](t_, "setBezelStyle:", value)
}

func (t_ TextField) IsBordered() bool {
	rv := ffi.CallMethod[bool](t_, "isBordered")
	return rv
}

func (t_ TextField) SetBordered(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setBordered:", value)
}

func (t_ TextField) AllowsCharacterPickerTouchBarItem() bool {
	rv := ffi.CallMethod[bool](t_, "allowsCharacterPickerTouchBarItem")
	return rv
}

func (t_ TextField) SetAllowsCharacterPickerTouchBarItem(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAllowsCharacterPickerTouchBarItem:", value)
}

func (t_ TextField) IsAutomaticTextCompletionEnabled() bool {
	rv := ffi.CallMethod[bool](t_, "isAutomaticTextCompletionEnabled")
	return rv
}

func (t_ TextField) SetAutomaticTextCompletionEnabled(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAutomaticTextCompletionEnabled:", value)
}

func (t_ TextField) Delegate() TextFieldDelegateWrapper {
	rv := ffi.CallMethod[TextFieldDelegateWrapper](t_, "delegate")
	return rv
}

func (t_ TextField) SetDelegate(value TextFieldDelegate) {
	po := ffi.CreateProtocol("NSTextFieldDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](t_, "setDelegate:", po)
}
