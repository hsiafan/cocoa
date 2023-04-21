// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ButtonClass = _ButtonClass{objc.GetClass("NSButton")}

type _ButtonClass struct {
	objc.Class
}

type IButton interface {
	IControl
	SetButtonType(type_ ButtonType)
	GetPeriodicDelay_Interval(delay *float32, interval *float32)
	SetPeriodicDelay_Interval(delay float32, interval float32)
	// deprecated
	SetTitleWithMnemonic(stringWithAmpersand string)
	CompressWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions)
	MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) foundation.Size
	SetNextState()
	Highlight(flag bool)
	ContentTintColor() Color
	SetContentTintColor(value IColor)
	HasDestructiveAction() bool
	SetHasDestructiveAction(value bool)
	AlternateTitle() string
	SetAlternateTitle(value string)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	AttributedAlternateTitle() foundation.AttributedString
	SetAttributedAlternateTitle(value foundation.IAttributedString)
	Title() string
	SetTitle(value string)
	SymbolConfiguration() ImageSymbolConfiguration
	SetSymbolConfiguration(value IImageSymbolConfiguration)
	Sound() Sound
	SetSound(value ISound)
	IsSpringLoaded() bool
	SetSpringLoaded(value bool)
	MaxAcceleratorLevel() int
	SetMaxAcceleratorLevel(value int)
	Image() Image
	SetImage(value IImage)
	AlternateImage() Image
	SetAlternateImage(value IImage)
	ImagePosition() CellImagePosition
	SetImagePosition(value CellImagePosition)
	IsBordered() bool
	SetBordered(value bool)
	IsTransparent() bool
	SetTransparent(value bool)
	BezelStyle() BezelStyle
	SetBezelStyle(value BezelStyle)
	BezelColor() Color
	SetBezelColor(value IColor)
	ShowsBorderOnlyWhileMouseInside() bool
	SetShowsBorderOnlyWhileMouseInside(value bool)
	ImageHugsTitle() bool
	SetImageHugsTitle(value bool)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
	ActiveCompressionOptions() UserInterfaceCompressionOptions
	AllowsMixedState() bool
	SetAllowsMixedState(value bool)
	State() ControlStateValue
	SetState(value ControlStateValue)
	KeyEquivalent() string
	SetKeyEquivalent(value string)
	KeyEquivalentModifierMask() EventModifierFlags
	SetKeyEquivalentModifierMask(value EventModifierFlags)
}

type Button struct {
	Control
}

func MakeButton(ptr unsafe.Pointer) Button {
	return Button{
		Control: MakeControl(ptr),
	}
}

func (bc _ButtonClass) CheckboxWithTitle_Target_Action(title string, target objc.IObject, action objc.Selector) Button {
	rv := objc.CallMethod[Button](bc, "checkboxWithTitle:target:action:", title, target, action)
	return rv
}

func (bc _ButtonClass) ButtonWithImage_Target_Action(image IImage, target objc.IObject, action objc.Selector) Button {
	rv := objc.CallMethod[Button](bc, "buttonWithImage:target:action:", image, target, action)
	return rv
}

func (bc _ButtonClass) RadioButtonWithTitle_Target_Action(title string, target objc.IObject, action objc.Selector) Button {
	rv := objc.CallMethod[Button](bc, "radioButtonWithTitle:target:action:", title, target, action)
	return rv
}

func (bc _ButtonClass) ButtonWithTitle_Image_Target_Action(title string, image IImage, target objc.IObject, action objc.Selector) Button {
	rv := objc.CallMethod[Button](bc, "buttonWithTitle:image:target:action:", title, image, target, action)
	return rv
}

func (bc _ButtonClass) ButtonWithTitle_Target_Action(title string, target objc.IObject, action objc.Selector) Button {
	rv := objc.CallMethod[Button](bc, "buttonWithTitle:target:action:", title, target, action)
	return rv
}

func (b_ Button) InitWithFrame(frameRect foundation.Rect) Button {
	rv := objc.CallMethod[Button](b_, "initWithFrame:", frameRect)
	return rv
}

func (b_ Button) Init() Button {
	rv := objc.CallMethod[Button](b_, "init")
	return rv
}

func (bc _ButtonClass) Alloc() Button {
	rv := objc.CallMethod[Button](bc, "alloc")
	return rv
}

func (bc _ButtonClass) New() Button {
	rv := objc.CallMethod[Button](bc, "new")
	rv.Autorelease()
	return rv
}

func NewButton() Button {
	return ButtonClass.New()
}

func (b_ Button) SetButtonType(type_ ButtonType) {
	objc.CallMethod[objc.Void](b_, "setButtonType:", type_)
}

func (b_ Button) GetPeriodicDelay_Interval(delay *float32, interval *float32) {
	objc.CallMethod[objc.Void](b_, "getPeriodicDelay:interval:", delay, interval)
}

func (b_ Button) SetPeriodicDelay_Interval(delay float32, interval float32) {
	objc.CallMethod[objc.Void](b_, "setPeriodicDelay:interval:", delay, interval)
}

// deprecated
func (b_ Button) SetTitleWithMnemonic(stringWithAmpersand string) {
	objc.CallMethod[objc.Void](b_, "setTitleWithMnemonic:", stringWithAmpersand)
}

func (b_ Button) CompressWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) {
	objc.CallMethod[objc.Void](b_, "compressWithPrioritizedCompressionOptions:", prioritizedOptions)
}

func (b_ Button) MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) foundation.Size {
	rv := objc.CallMethod[foundation.Size](b_, "minimumSizeWithPrioritizedCompressionOptions:", prioritizedOptions)
	return rv
}

func (b_ Button) SetNextState() {
	objc.CallMethod[objc.Void](b_, "setNextState")
}

func (b_ Button) Highlight(flag bool) {
	objc.CallMethod[objc.Void](b_, "highlight:", flag)
}

func (b_ Button) ContentTintColor() Color {
	rv := objc.CallMethod[Color](b_, "contentTintColor")
	return rv
}

func (b_ Button) SetContentTintColor(value IColor) {
	objc.CallMethod[objc.Void](b_, "setContentTintColor:", value)
}

func (b_ Button) HasDestructiveAction() bool {
	rv := objc.CallMethod[bool](b_, "hasDestructiveAction")
	return rv
}

func (b_ Button) SetHasDestructiveAction(value bool) {
	objc.CallMethod[objc.Void](b_, "setHasDestructiveAction:", value)
}

func (b_ Button) AlternateTitle() string {
	rv := objc.CallMethod[string](b_, "alternateTitle")
	return rv
}

func (b_ Button) SetAlternateTitle(value string) {
	objc.CallMethod[objc.Void](b_, "setAlternateTitle:", value)
}

func (b_ Button) AttributedTitle() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](b_, "attributedTitle")
	return rv
}

func (b_ Button) SetAttributedTitle(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](b_, "setAttributedTitle:", value)
}

func (b_ Button) AttributedAlternateTitle() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](b_, "attributedAlternateTitle")
	return rv
}

func (b_ Button) SetAttributedAlternateTitle(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](b_, "setAttributedAlternateTitle:", value)
}

func (b_ Button) Title() string {
	rv := objc.CallMethod[string](b_, "title")
	return rv
}

func (b_ Button) SetTitle(value string) {
	objc.CallMethod[objc.Void](b_, "setTitle:", value)
}

func (b_ Button) SymbolConfiguration() ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](b_, "symbolConfiguration")
	return rv
}

func (b_ Button) SetSymbolConfiguration(value IImageSymbolConfiguration) {
	objc.CallMethod[objc.Void](b_, "setSymbolConfiguration:", value)
}

func (b_ Button) Sound() Sound {
	rv := objc.CallMethod[Sound](b_, "sound")
	return rv
}

func (b_ Button) SetSound(value ISound) {
	objc.CallMethod[objc.Void](b_, "setSound:", value)
}

func (b_ Button) IsSpringLoaded() bool {
	rv := objc.CallMethod[bool](b_, "isSpringLoaded")
	return rv
}

func (b_ Button) SetSpringLoaded(value bool) {
	objc.CallMethod[objc.Void](b_, "setSpringLoaded:", value)
}

func (b_ Button) MaxAcceleratorLevel() int {
	rv := objc.CallMethod[int](b_, "maxAcceleratorLevel")
	return rv
}

func (b_ Button) SetMaxAcceleratorLevel(value int) {
	objc.CallMethod[objc.Void](b_, "setMaxAcceleratorLevel:", value)
}

func (b_ Button) Image() Image {
	rv := objc.CallMethod[Image](b_, "image")
	return rv
}

func (b_ Button) SetImage(value IImage) {
	objc.CallMethod[objc.Void](b_, "setImage:", value)
}

func (b_ Button) AlternateImage() Image {
	rv := objc.CallMethod[Image](b_, "alternateImage")
	return rv
}

func (b_ Button) SetAlternateImage(value IImage) {
	objc.CallMethod[objc.Void](b_, "setAlternateImage:", value)
}

func (b_ Button) ImagePosition() CellImagePosition {
	rv := objc.CallMethod[CellImagePosition](b_, "imagePosition")
	return rv
}

func (b_ Button) SetImagePosition(value CellImagePosition) {
	objc.CallMethod[objc.Void](b_, "setImagePosition:", value)
}

func (b_ Button) IsBordered() bool {
	rv := objc.CallMethod[bool](b_, "isBordered")
	return rv
}

func (b_ Button) SetBordered(value bool) {
	objc.CallMethod[objc.Void](b_, "setBordered:", value)
}

func (b_ Button) IsTransparent() bool {
	rv := objc.CallMethod[bool](b_, "isTransparent")
	return rv
}

func (b_ Button) SetTransparent(value bool) {
	objc.CallMethod[objc.Void](b_, "setTransparent:", value)
}

func (b_ Button) BezelStyle() BezelStyle {
	rv := objc.CallMethod[BezelStyle](b_, "bezelStyle")
	return rv
}

func (b_ Button) SetBezelStyle(value BezelStyle) {
	objc.CallMethod[objc.Void](b_, "setBezelStyle:", value)
}

func (b_ Button) BezelColor() Color {
	rv := objc.CallMethod[Color](b_, "bezelColor")
	return rv
}

func (b_ Button) SetBezelColor(value IColor) {
	objc.CallMethod[objc.Void](b_, "setBezelColor:", value)
}

func (b_ Button) ShowsBorderOnlyWhileMouseInside() bool {
	rv := objc.CallMethod[bool](b_, "showsBorderOnlyWhileMouseInside")
	return rv
}

func (b_ Button) SetShowsBorderOnlyWhileMouseInside(value bool) {
	objc.CallMethod[objc.Void](b_, "setShowsBorderOnlyWhileMouseInside:", value)
}

func (b_ Button) ImageHugsTitle() bool {
	rv := objc.CallMethod[bool](b_, "imageHugsTitle")
	return rv
}

func (b_ Button) SetImageHugsTitle(value bool) {
	objc.CallMethod[objc.Void](b_, "setImageHugsTitle:", value)
}

func (b_ Button) ImageScaling() ImageScaling {
	rv := objc.CallMethod[ImageScaling](b_, "imageScaling")
	return rv
}

func (b_ Button) SetImageScaling(value ImageScaling) {
	objc.CallMethod[objc.Void](b_, "setImageScaling:", value)
}

func (b_ Button) ActiveCompressionOptions() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](b_, "activeCompressionOptions")
	return rv
}

func (b_ Button) AllowsMixedState() bool {
	rv := objc.CallMethod[bool](b_, "allowsMixedState")
	return rv
}

func (b_ Button) SetAllowsMixedState(value bool) {
	objc.CallMethod[objc.Void](b_, "setAllowsMixedState:", value)
}

func (b_ Button) State() ControlStateValue {
	rv := objc.CallMethod[ControlStateValue](b_, "state")
	return rv
}

func (b_ Button) SetState(value ControlStateValue) {
	objc.CallMethod[objc.Void](b_, "setState:", value)
}

func (b_ Button) KeyEquivalent() string {
	rv := objc.CallMethod[string](b_, "keyEquivalent")
	return rv
}

func (b_ Button) SetKeyEquivalent(value string) {
	objc.CallMethod[objc.Void](b_, "setKeyEquivalent:", value)
}

func (b_ Button) KeyEquivalentModifierMask() EventModifierFlags {
	rv := objc.CallMethod[EventModifierFlags](b_, "keyEquivalentModifierMask")
	return rv
}

func (b_ Button) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	objc.CallMethod[objc.Void](b_, "setKeyEquivalentModifierMask:", value)
}
