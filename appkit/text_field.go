// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	SetDelegate0(value objc.IObject)
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
	rv := objc.CallMethod[TextField](tc, "labelWithAttributedString:", attributedStringValue)
	return rv
}

func (tc _TextFieldClass) LabelWithString(stringValue string) TextField {
	rv := objc.CallMethod[TextField](tc, "labelWithString:", stringValue)
	return rv
}

func (tc _TextFieldClass) TextFieldWithString(stringValue string) TextField {
	rv := objc.CallMethod[TextField](tc, "textFieldWithString:", stringValue)
	return rv
}

func (tc _TextFieldClass) WrappingLabelWithString(stringValue string) TextField {
	rv := objc.CallMethod[TextField](tc, "wrappingLabelWithString:", stringValue)
	return rv
}

func (t_ TextField) InitWithFrame(frameRect foundation.Rect) TextField {
	rv := objc.CallMethod[TextField](t_, "initWithFrame:", frameRect)
	return rv
}

func (t_ TextField) Init() TextField {
	rv := objc.CallMethod[TextField](t_, "init")
	return rv
}

func (tc _TextFieldClass) Alloc() TextField {
	rv := objc.CallMethod[TextField](tc, "alloc")
	return rv
}

func (tc _TextFieldClass) New() TextField {
	rv := objc.CallMethod[TextField](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextField() TextField {
	return TextFieldClass.New()
}

func (t_ TextField) SelectText(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, "selectText:", sender)
}

// deprecated
func (t_ TextField) SetTitleWithMnemonic(stringWithAmpersand string) {
	objc.CallMethod[objc.Void](t_, "setTitleWithMnemonic:", stringWithAmpersand)
}

func (t_ TextField) TextShouldBeginEditing(textObject IText) bool {
	rv := objc.CallMethod[bool](t_, "textShouldBeginEditing:", textObject)
	return rv
}

func (t_ TextField) TextDidBeginEditing(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, "textDidBeginEditing:", notification)
}

func (t_ TextField) TextDidChange(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, "textDidChange:", notification)
}

func (t_ TextField) TextShouldEndEditing(textObject IText) bool {
	rv := objc.CallMethod[bool](t_, "textShouldEndEditing:", textObject)
	return rv
}

func (t_ TextField) TextDidEndEditing(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, "textDidEndEditing:", notification)
}

func (t_ TextField) IsSelectable() bool {
	rv := objc.CallMethod[bool](t_, "isSelectable")
	return rv
}

func (t_ TextField) SetSelectable(value bool) {
	objc.CallMethod[objc.Void](t_, "setSelectable:", value)
}

func (t_ TextField) IsEditable() bool {
	rv := objc.CallMethod[bool](t_, "isEditable")
	return rv
}

func (t_ TextField) SetEditable(value bool) {
	objc.CallMethod[objc.Void](t_, "setEditable:", value)
}

func (t_ TextField) AllowsEditingTextAttributes() bool {
	rv := objc.CallMethod[bool](t_, "allowsEditingTextAttributes")
	return rv
}

func (t_ TextField) SetAllowsEditingTextAttributes(value bool) {
	objc.CallMethod[objc.Void](t_, "setAllowsEditingTextAttributes:", value)
}

func (t_ TextField) ImportsGraphics() bool {
	rv := objc.CallMethod[bool](t_, "importsGraphics")
	return rv
}

func (t_ TextField) SetImportsGraphics(value bool) {
	objc.CallMethod[objc.Void](t_, "setImportsGraphics:", value)
}

func (t_ TextField) PlaceholderString() string {
	rv := objc.CallMethod[string](t_, "placeholderString")
	return rv
}

func (t_ TextField) SetPlaceholderString(value string) {
	objc.CallMethod[objc.Void](t_, "setPlaceholderString:", value)
}

func (t_ TextField) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](t_, "placeholderAttributedString")
	return rv
}

func (t_ TextField) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](t_, "setPlaceholderAttributedString:", value)
}

func (t_ TextField) LineBreakStrategy() LineBreakStrategy {
	rv := objc.CallMethod[LineBreakStrategy](t_, "lineBreakStrategy")
	return rv
}

func (t_ TextField) SetLineBreakStrategy(value LineBreakStrategy) {
	objc.CallMethod[objc.Void](t_, "setLineBreakStrategy:", value)
}

func (t_ TextField) AllowsDefaultTighteningForTruncation() bool {
	rv := objc.CallMethod[bool](t_, "allowsDefaultTighteningForTruncation")
	return rv
}

func (t_ TextField) SetAllowsDefaultTighteningForTruncation(value bool) {
	objc.CallMethod[objc.Void](t_, "setAllowsDefaultTighteningForTruncation:", value)
}

func (t_ TextField) MaximumNumberOfLines() int {
	rv := objc.CallMethod[int](t_, "maximumNumberOfLines")
	return rv
}

func (t_ TextField) SetMaximumNumberOfLines(value int) {
	objc.CallMethod[objc.Void](t_, "setMaximumNumberOfLines:", value)
}

func (t_ TextField) PreferredMaxLayoutWidth() float64 {
	rv := objc.CallMethod[float64](t_, "preferredMaxLayoutWidth")
	return rv
}

func (t_ TextField) SetPreferredMaxLayoutWidth(value float64) {
	objc.CallMethod[objc.Void](t_, "setPreferredMaxLayoutWidth:", value)
}

func (t_ TextField) TextColor() Color {
	rv := objc.CallMethod[Color](t_, "textColor")
	return rv
}

func (t_ TextField) SetTextColor(value IColor) {
	objc.CallMethod[objc.Void](t_, "setTextColor:", value)
}

func (t_ TextField) BackgroundColor() Color {
	rv := objc.CallMethod[Color](t_, "backgroundColor")
	return rv
}

func (t_ TextField) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](t_, "setBackgroundColor:", value)
}

func (t_ TextField) DrawsBackground() bool {
	rv := objc.CallMethod[bool](t_, "drawsBackground")
	return rv
}

func (t_ TextField) SetDrawsBackground(value bool) {
	objc.CallMethod[objc.Void](t_, "setDrawsBackground:", value)
}

func (t_ TextField) IsBezeled() bool {
	rv := objc.CallMethod[bool](t_, "isBezeled")
	return rv
}

func (t_ TextField) SetBezeled(value bool) {
	objc.CallMethod[objc.Void](t_, "setBezeled:", value)
}

func (t_ TextField) BezelStyle() TextFieldBezelStyle {
	rv := objc.CallMethod[TextFieldBezelStyle](t_, "bezelStyle")
	return rv
}

func (t_ TextField) SetBezelStyle(value TextFieldBezelStyle) {
	objc.CallMethod[objc.Void](t_, "setBezelStyle:", value)
}

func (t_ TextField) IsBordered() bool {
	rv := objc.CallMethod[bool](t_, "isBordered")
	return rv
}

func (t_ TextField) SetBordered(value bool) {
	objc.CallMethod[objc.Void](t_, "setBordered:", value)
}

func (t_ TextField) AllowsCharacterPickerTouchBarItem() bool {
	rv := objc.CallMethod[bool](t_, "allowsCharacterPickerTouchBarItem")
	return rv
}

func (t_ TextField) SetAllowsCharacterPickerTouchBarItem(value bool) {
	objc.CallMethod[objc.Void](t_, "setAllowsCharacterPickerTouchBarItem:", value)
}

func (t_ TextField) IsAutomaticTextCompletionEnabled() bool {
	rv := objc.CallMethod[bool](t_, "isAutomaticTextCompletionEnabled")
	return rv
}

func (t_ TextField) SetAutomaticTextCompletionEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, "setAutomaticTextCompletionEnabled:", value)
}

func (t_ TextField) Delegate() TextFieldDelegateWrapper {
	rv := objc.CallMethod[TextFieldDelegateWrapper](t_, "delegate")
	return rv
}

func (t_ TextField) SetDelegate(value TextFieldDelegate) {
	po := objc.CreateProtocol("NSTextFieldDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](t_, "setDelegate:", po)
}

func (t_ TextField) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, "setDelegate:", value)
}
