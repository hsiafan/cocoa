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
	rv := objc.CallMethod[Button](bc, objc.SEL("checkboxWithTitle:target:action:"), title, objc.ExtractPtr(target), action)
	return rv
}

func (bc _ButtonClass) ButtonWithImage_Target_Action(image IImage, target objc.IObject, action objc.Selector) Button {
	rv := objc.CallMethod[Button](bc, objc.SEL("buttonWithImage:target:action:"), objc.ExtractPtr(image), objc.ExtractPtr(target), action)
	return rv
}

func (bc _ButtonClass) RadioButtonWithTitle_Target_Action(title string, target objc.IObject, action objc.Selector) Button {
	rv := objc.CallMethod[Button](bc, objc.SEL("radioButtonWithTitle:target:action:"), title, objc.ExtractPtr(target), action)
	return rv
}

func (bc _ButtonClass) ButtonWithTitle_Image_Target_Action(title string, image IImage, target objc.IObject, action objc.Selector) Button {
	rv := objc.CallMethod[Button](bc, objc.SEL("buttonWithTitle:image:target:action:"), title, objc.ExtractPtr(image), objc.ExtractPtr(target), action)
	return rv
}

func (bc _ButtonClass) ButtonWithTitle_Target_Action(title string, target objc.IObject, action objc.Selector) Button {
	rv := objc.CallMethod[Button](bc, objc.SEL("buttonWithTitle:target:action:"), title, objc.ExtractPtr(target), action)
	return rv
}

func (b_ Button) InitWithFrame(frameRect foundation.Rect) Button {
	rv := objc.CallMethod[Button](b_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (b_ Button) Init() Button {
	rv := objc.CallMethod[Button](b_, objc.SEL("init"))
	return rv
}

func (bc _ButtonClass) Alloc() Button {
	rv := objc.CallMethod[Button](bc, objc.SEL("alloc"))
	return rv
}

func (bc _ButtonClass) New() Button {
	rv := objc.CallMethod[Button](bc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewButton() Button {
	return ButtonClass.New()
}

func (b_ Button) SetButtonType(type_ ButtonType) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setButtonType:"), type_)
}

func (b_ Button) GetPeriodicDelay_Interval(delay *float32, interval *float32) {
	objc.CallMethod[objc.Void](b_, objc.SEL("getPeriodicDelay:interval:"), delay, interval)
}

func (b_ Button) SetPeriodicDelay_Interval(delay float32, interval float32) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setPeriodicDelay:interval:"), delay, interval)
}

// deprecated
func (b_ Button) SetTitleWithMnemonic(stringWithAmpersand string) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setTitleWithMnemonic:"), stringWithAmpersand)
}

func (b_ Button) CompressWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) {
	objc.CallMethod[objc.Void](b_, objc.SEL("compressWithPrioritizedCompressionOptions:"), prioritizedOptions)
}

func (b_ Button) MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) foundation.Size {
	rv := objc.CallMethod[foundation.Size](b_, objc.SEL("minimumSizeWithPrioritizedCompressionOptions:"), prioritizedOptions)
	return rv
}

func (b_ Button) SetNextState() {
	objc.CallMethod[objc.Void](b_, objc.SEL("setNextState"))
}

func (b_ Button) Highlight(flag bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("highlight:"), flag)
}

func (b_ Button) ContentTintColor() Color {
	rv := objc.CallMethod[Color](b_, objc.SEL("contentTintColor"))
	return rv
}

func (b_ Button) SetContentTintColor(value IColor) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setContentTintColor:"), objc.ExtractPtr(value))
}

func (b_ Button) HasDestructiveAction() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("hasDestructiveAction"))
	return rv
}

func (b_ Button) SetHasDestructiveAction(value bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setHasDestructiveAction:"), value)
}

func (b_ Button) AlternateTitle() string {
	rv := objc.CallMethod[string](b_, objc.SEL("alternateTitle"))
	return rv
}

func (b_ Button) SetAlternateTitle(value string) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setAlternateTitle:"), value)
}

func (b_ Button) AttributedTitle() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](b_, objc.SEL("attributedTitle"))
	return rv
}

func (b_ Button) SetAttributedTitle(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setAttributedTitle:"), objc.ExtractPtr(value))
}

func (b_ Button) AttributedAlternateTitle() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](b_, objc.SEL("attributedAlternateTitle"))
	return rv
}

func (b_ Button) SetAttributedAlternateTitle(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setAttributedAlternateTitle:"), objc.ExtractPtr(value))
}

func (b_ Button) Title() string {
	rv := objc.CallMethod[string](b_, objc.SEL("title"))
	return rv
}

func (b_ Button) SetTitle(value string) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setTitle:"), value)
}

func (b_ Button) SymbolConfiguration() ImageSymbolConfiguration {
	rv := objc.CallMethod[ImageSymbolConfiguration](b_, objc.SEL("symbolConfiguration"))
	return rv
}

func (b_ Button) SetSymbolConfiguration(value IImageSymbolConfiguration) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setSymbolConfiguration:"), objc.ExtractPtr(value))
}

func (b_ Button) Sound() Sound {
	rv := objc.CallMethod[Sound](b_, objc.SEL("sound"))
	return rv
}

func (b_ Button) SetSound(value ISound) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setSound:"), objc.ExtractPtr(value))
}

func (b_ Button) IsSpringLoaded() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("isSpringLoaded"))
	return rv
}

func (b_ Button) SetSpringLoaded(value bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setSpringLoaded:"), value)
}

func (b_ Button) MaxAcceleratorLevel() int {
	rv := objc.CallMethod[int](b_, objc.SEL("maxAcceleratorLevel"))
	return rv
}

func (b_ Button) SetMaxAcceleratorLevel(value int) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setMaxAcceleratorLevel:"), value)
}

func (b_ Button) Image() Image {
	rv := objc.CallMethod[Image](b_, objc.SEL("image"))
	return rv
}

func (b_ Button) SetImage(value IImage) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setImage:"), objc.ExtractPtr(value))
}

func (b_ Button) AlternateImage() Image {
	rv := objc.CallMethod[Image](b_, objc.SEL("alternateImage"))
	return rv
}

func (b_ Button) SetAlternateImage(value IImage) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setAlternateImage:"), objc.ExtractPtr(value))
}

func (b_ Button) ImagePosition() CellImagePosition {
	rv := objc.CallMethod[CellImagePosition](b_, objc.SEL("imagePosition"))
	return rv
}

func (b_ Button) SetImagePosition(value CellImagePosition) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setImagePosition:"), value)
}

func (b_ Button) IsBordered() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("isBordered"))
	return rv
}

func (b_ Button) SetBordered(value bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setBordered:"), value)
}

func (b_ Button) IsTransparent() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("isTransparent"))
	return rv
}

func (b_ Button) SetTransparent(value bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setTransparent:"), value)
}

func (b_ Button) BezelStyle() BezelStyle {
	rv := objc.CallMethod[BezelStyle](b_, objc.SEL("bezelStyle"))
	return rv
}

func (b_ Button) SetBezelStyle(value BezelStyle) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setBezelStyle:"), value)
}

func (b_ Button) BezelColor() Color {
	rv := objc.CallMethod[Color](b_, objc.SEL("bezelColor"))
	return rv
}

func (b_ Button) SetBezelColor(value IColor) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setBezelColor:"), objc.ExtractPtr(value))
}

func (b_ Button) ShowsBorderOnlyWhileMouseInside() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("showsBorderOnlyWhileMouseInside"))
	return rv
}

func (b_ Button) SetShowsBorderOnlyWhileMouseInside(value bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setShowsBorderOnlyWhileMouseInside:"), value)
}

func (b_ Button) ImageHugsTitle() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("imageHugsTitle"))
	return rv
}

func (b_ Button) SetImageHugsTitle(value bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setImageHugsTitle:"), value)
}

func (b_ Button) ImageScaling() ImageScaling {
	rv := objc.CallMethod[ImageScaling](b_, objc.SEL("imageScaling"))
	return rv
}

func (b_ Button) SetImageScaling(value ImageScaling) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setImageScaling:"), value)
}

func (b_ Button) ActiveCompressionOptions() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](b_, objc.SEL("activeCompressionOptions"))
	return rv
}

func (b_ Button) AllowsMixedState() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("allowsMixedState"))
	return rv
}

func (b_ Button) SetAllowsMixedState(value bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setAllowsMixedState:"), value)
}

func (b_ Button) State() ControlStateValue {
	rv := objc.CallMethod[ControlStateValue](b_, objc.SEL("state"))
	return rv
}

func (b_ Button) SetState(value ControlStateValue) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setState:"), value)
}

func (b_ Button) KeyEquivalent() string {
	rv := objc.CallMethod[string](b_, objc.SEL("keyEquivalent"))
	return rv
}

func (b_ Button) SetKeyEquivalent(value string) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setKeyEquivalent:"), value)
}

func (b_ Button) KeyEquivalentModifierMask() EventModifierFlags {
	rv := objc.CallMethod[EventModifierFlags](b_, objc.SEL("keyEquivalentModifierMask"))
	return rv
}

func (b_ Button) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setKeyEquivalentModifierMask:"), value)
}
