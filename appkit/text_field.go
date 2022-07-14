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
	rv.Autorelease()
	return rv
}

func (t_ TextField) Init() TextField {
	rv := ffi.CallMethod[TextField](t_, "init")
	rv.Autorelease()
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
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](t_, "setDelegate:", po)
}

var SecureTextFieldClass = _SecureTextFieldClass{objc.GetClass("NSSecureTextField")}

type _SecureTextFieldClass struct {
	objc.Class
}

type ISecureTextField interface {
	ITextField
}

type SecureTextField struct {
	TextField
}

func MakeSecureTextField(ptr unsafe.Pointer) SecureTextField {
	return SecureTextField{
		TextField: MakeTextField(ptr),
	}
}

func (sc _SecureTextFieldClass) LabelWithAttributedString(attributedStringValue foundation.IAttributedString) SecureTextField {
	rv := ffi.CallMethod[SecureTextField](sc, "labelWithAttributedString:", attributedStringValue)
	return rv
}

func (sc _SecureTextFieldClass) LabelWithString(stringValue string) SecureTextField {
	rv := ffi.CallMethod[SecureTextField](sc, "labelWithString:", stringValue)
	return rv
}

func (sc _SecureTextFieldClass) TextFieldWithString(stringValue string) SecureTextField {
	rv := ffi.CallMethod[SecureTextField](sc, "textFieldWithString:", stringValue)
	return rv
}

func (sc _SecureTextFieldClass) WrappingLabelWithString(stringValue string) SecureTextField {
	rv := ffi.CallMethod[SecureTextField](sc, "wrappingLabelWithString:", stringValue)
	return rv
}

func (s_ SecureTextField) InitWithFrame(frameRect foundation.Rect) SecureTextField {
	rv := ffi.CallMethod[SecureTextField](s_, "initWithFrame:", frameRect)
	rv.Autorelease()
	return rv
}

func (s_ SecureTextField) Init() SecureTextField {
	rv := ffi.CallMethod[SecureTextField](s_, "init")
	rv.Autorelease()
	return rv
}

func (sc _SecureTextFieldClass) Alloc() SecureTextField {
	rv := ffi.CallMethod[SecureTextField](sc, "alloc")
	return rv
}

func (sc _SecureTextFieldClass) New() SecureTextField {
	rv := ffi.CallMethod[SecureTextField](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSecureTextField() SecureTextField {
	return SecureTextFieldClass.New()
}

type TextFieldDelegate interface {
	ControlTextEditingDelegate
	ImplementsTextField_TextView_Candidates_ForSelectedRange() bool
	// optional
	TextField_TextView_Candidates_ForSelectedRange(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult
	ImplementsTextField_TextView_CandidatesForSelectedRange() bool
	// optional
	TextField_TextView_CandidatesForSelectedRange(textField TextField, textView TextView, selectedRange foundation.Range) []objc.IObject
	ImplementsTextField_TextView_ShouldSelectCandidateAtIndex() bool
	// optional
	TextField_TextView_ShouldSelectCandidateAtIndex(textField TextField, textView TextView, index uint) bool
}

type TextFieldDelegateImpl struct {
	ControlTextEditingDelegateImpl
	_TextField_TextView_Candidates_ForSelectedRange  func(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult
	_TextField_TextView_CandidatesForSelectedRange   func(textField TextField, textView TextView, selectedRange foundation.Range) []objc.IObject
	_TextField_TextView_ShouldSelectCandidateAtIndex func(textField TextField, textView TextView, index uint) bool
}

func (di *TextFieldDelegateImpl) ImplementsTextField_TextView_Candidates_ForSelectedRange() bool {
	return di._TextField_TextView_Candidates_ForSelectedRange != nil
}

func (di *TextFieldDelegateImpl) SetTextField_TextView_Candidates_ForSelectedRange(f func(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult) {
	di._TextField_TextView_Candidates_ForSelectedRange = f
}

func (di *TextFieldDelegateImpl) TextField_TextView_Candidates_ForSelectedRange(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult {
	return di._TextField_TextView_Candidates_ForSelectedRange(textField, textView, candidates, selectedRange)
}
func (di *TextFieldDelegateImpl) ImplementsTextField_TextView_CandidatesForSelectedRange() bool {
	return di._TextField_TextView_CandidatesForSelectedRange != nil
}

func (di *TextFieldDelegateImpl) SetTextField_TextView_CandidatesForSelectedRange(f func(textField TextField, textView TextView, selectedRange foundation.Range) []objc.IObject) {
	di._TextField_TextView_CandidatesForSelectedRange = f
}

func (di *TextFieldDelegateImpl) TextField_TextView_CandidatesForSelectedRange(textField TextField, textView TextView, selectedRange foundation.Range) []objc.IObject {
	return di._TextField_TextView_CandidatesForSelectedRange(textField, textView, selectedRange)
}
func (di *TextFieldDelegateImpl) ImplementsTextField_TextView_ShouldSelectCandidateAtIndex() bool {
	return di._TextField_TextView_ShouldSelectCandidateAtIndex != nil
}

func (di *TextFieldDelegateImpl) SetTextField_TextView_ShouldSelectCandidateAtIndex(f func(textField TextField, textView TextView, index uint) bool) {
	di._TextField_TextView_ShouldSelectCandidateAtIndex = f
}

func (di *TextFieldDelegateImpl) TextField_TextView_ShouldSelectCandidateAtIndex(textField TextField, textView TextView, index uint) bool {
	return di._TextField_TextView_ShouldSelectCandidateAtIndex(textField, textView, index)
}

type TextFieldDelegateWrapper struct {
	ControlTextEditingDelegateWrapper
}

func (t_ *TextFieldDelegateWrapper) ImplementsTextField_TextView_Candidates_ForSelectedRange() bool {
	return t_.RespondsToSelector(objc.GetSelector("textField:textView:candidates:forSelectedRange:"))
}

func (t_ TextFieldDelegateWrapper) TextField_TextView_Candidates_ForSelectedRange(textField ITextField, textView ITextView, candidates []foundation.ITextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult {
	rv := ffi.CallMethod[[]foundation.TextCheckingResult](t_, "textField:textView:candidates:forSelectedRange:", textField, textView, candidates, selectedRange)
	return rv
}

func (t_ *TextFieldDelegateWrapper) ImplementsTextField_TextView_CandidatesForSelectedRange() bool {
	return t_.RespondsToSelector(objc.GetSelector("textField:textView:candidatesForSelectedRange:"))
}

func (t_ TextFieldDelegateWrapper) TextField_TextView_CandidatesForSelectedRange(textField ITextField, textView ITextView, selectedRange foundation.Range) []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](t_, "textField:textView:candidatesForSelectedRange:", textField, textView, selectedRange)
	return rv
}

func (t_ *TextFieldDelegateWrapper) ImplementsTextField_TextView_ShouldSelectCandidateAtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("textField:textView:shouldSelectCandidateAtIndex:"))
}

func (t_ TextFieldDelegateWrapper) TextField_TextView_ShouldSelectCandidateAtIndex(textField ITextField, textView ITextView, index uint) bool {
	rv := ffi.CallMethod[bool](t_, "textField:textView:shouldSelectCandidateAtIndex:", textField, textView, index)
	return rv
}

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

func (t_ TextFieldCell) InitTextCell(_string string) TextFieldCell {
	rv := ffi.CallMethod[TextFieldCell](t_, "initTextCell:", _string)
	rv.Autorelease()
	return rv
}

func (t_ TextFieldCell) InitImageCell(image IImage) TextFieldCell {
	rv := ffi.CallMethod[TextFieldCell](t_, "initImageCell:", image)
	rv.Autorelease()
	return rv
}

func (t_ TextFieldCell) Init() TextFieldCell {
	rv := ffi.CallMethod[TextFieldCell](t_, "init")
	rv.Autorelease()
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
