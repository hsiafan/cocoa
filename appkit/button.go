// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ButtonClass = _ButtonClass{objc.GetClass("NSButton")}

type _ButtonClass struct {
	objc.Class
}

type IButton interface {
	IControl
	SetButtonType(_type ButtonType)
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
	rv := ffi.CallMethod[Button](bc, "checkboxWithTitle:target:action:", title, target, action)
	return rv
}

func (bc _ButtonClass) ButtonWithImage_Target_Action(image IImage, target objc.IObject, action objc.Selector) Button {
	rv := ffi.CallMethod[Button](bc, "buttonWithImage:target:action:", image, target, action)
	return rv
}

func (bc _ButtonClass) RadioButtonWithTitle_Target_Action(title string, target objc.IObject, action objc.Selector) Button {
	rv := ffi.CallMethod[Button](bc, "radioButtonWithTitle:target:action:", title, target, action)
	return rv
}

func (bc _ButtonClass) ButtonWithTitle_Image_Target_Action(title string, image IImage, target objc.IObject, action objc.Selector) Button {
	rv := ffi.CallMethod[Button](bc, "buttonWithTitle:image:target:action:", title, image, target, action)
	return rv
}

func (bc _ButtonClass) ButtonWithTitle_Target_Action(title string, target objc.IObject, action objc.Selector) Button {
	rv := ffi.CallMethod[Button](bc, "buttonWithTitle:target:action:", title, target, action)
	return rv
}

func (b_ Button) InitWithFrame(frameRect foundation.Rect) Button {
	rv := ffi.CallMethod[Button](b_, "initWithFrame:", frameRect)
	return rv
}

func (b_ Button) Init() Button {
	rv := ffi.CallMethod[Button](b_, "init")
	return rv
}

func (bc _ButtonClass) Alloc() Button {
	rv := ffi.CallMethod[Button](bc, "alloc")
	return rv
}

func (bc _ButtonClass) New() Button {
	rv := ffi.CallMethod[Button](bc, "new")
	rv.Autorelease()
	return rv
}

func NewButton() Button {
	return ButtonClass.New()
}

func (b_ Button) SetButtonType(_type ButtonType) {
	ffi.CallMethod[ffi.Void](b_, "setButtonType:", _type)
}

func (b_ Button) GetPeriodicDelay_Interval(delay *float32, interval *float32) {
	ffi.CallMethod[ffi.Void](b_, "getPeriodicDelay:interval:", delay, interval)
}

func (b_ Button) SetPeriodicDelay_Interval(delay float32, interval float32) {
	ffi.CallMethod[ffi.Void](b_, "setPeriodicDelay:interval:", delay, interval)
}

// deprecated
func (b_ Button) SetTitleWithMnemonic(stringWithAmpersand string) {
	ffi.CallMethod[ffi.Void](b_, "setTitleWithMnemonic:", stringWithAmpersand)
}

func (b_ Button) CompressWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) {
	ffi.CallMethod[ffi.Void](b_, "compressWithPrioritizedCompressionOptions:", prioritizedOptions)
}

func (b_ Button) MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](b_, "minimumSizeWithPrioritizedCompressionOptions:", prioritizedOptions)
	return rv
}

func (b_ Button) SetNextState() {
	ffi.CallMethod[ffi.Void](b_, "setNextState")
}

func (b_ Button) Highlight(flag bool) {
	ffi.CallMethod[ffi.Void](b_, "highlight:", flag)
}

func (b_ Button) ContentTintColor() Color {
	rv := ffi.CallMethod[Color](b_, "contentTintColor")
	return rv
}

func (b_ Button) SetContentTintColor(value IColor) {
	ffi.CallMethod[ffi.Void](b_, "setContentTintColor:", value)
}

func (b_ Button) HasDestructiveAction() bool {
	rv := ffi.CallMethod[bool](b_, "hasDestructiveAction")
	return rv
}

func (b_ Button) SetHasDestructiveAction(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setHasDestructiveAction:", value)
}

func (b_ Button) AlternateTitle() string {
	rv := ffi.CallMethod[string](b_, "alternateTitle")
	return rv
}

func (b_ Button) SetAlternateTitle(value string) {
	ffi.CallMethod[ffi.Void](b_, "setAlternateTitle:", value)
}

func (b_ Button) AttributedTitle() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](b_, "attributedTitle")
	return rv
}

func (b_ Button) SetAttributedTitle(value foundation.IAttributedString) {
	ffi.CallMethod[ffi.Void](b_, "setAttributedTitle:", value)
}

func (b_ Button) AttributedAlternateTitle() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](b_, "attributedAlternateTitle")
	return rv
}

func (b_ Button) SetAttributedAlternateTitle(value foundation.IAttributedString) {
	ffi.CallMethod[ffi.Void](b_, "setAttributedAlternateTitle:", value)
}

func (b_ Button) Title() string {
	rv := ffi.CallMethod[string](b_, "title")
	return rv
}

func (b_ Button) SetTitle(value string) {
	ffi.CallMethod[ffi.Void](b_, "setTitle:", value)
}

func (b_ Button) SymbolConfiguration() ImageSymbolConfiguration {
	rv := ffi.CallMethod[ImageSymbolConfiguration](b_, "symbolConfiguration")
	return rv
}

func (b_ Button) SetSymbolConfiguration(value IImageSymbolConfiguration) {
	ffi.CallMethod[ffi.Void](b_, "setSymbolConfiguration:", value)
}

func (b_ Button) Sound() Sound {
	rv := ffi.CallMethod[Sound](b_, "sound")
	return rv
}

func (b_ Button) SetSound(value ISound) {
	ffi.CallMethod[ffi.Void](b_, "setSound:", value)
}

func (b_ Button) IsSpringLoaded() bool {
	rv := ffi.CallMethod[bool](b_, "isSpringLoaded")
	return rv
}

func (b_ Button) SetSpringLoaded(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setSpringLoaded:", value)
}

func (b_ Button) MaxAcceleratorLevel() int {
	rv := ffi.CallMethod[int](b_, "maxAcceleratorLevel")
	return rv
}

func (b_ Button) SetMaxAcceleratorLevel(value int) {
	ffi.CallMethod[ffi.Void](b_, "setMaxAcceleratorLevel:", value)
}

func (b_ Button) Image() Image {
	rv := ffi.CallMethod[Image](b_, "image")
	return rv
}

func (b_ Button) SetImage(value IImage) {
	ffi.CallMethod[ffi.Void](b_, "setImage:", value)
}

func (b_ Button) AlternateImage() Image {
	rv := ffi.CallMethod[Image](b_, "alternateImage")
	return rv
}

func (b_ Button) SetAlternateImage(value IImage) {
	ffi.CallMethod[ffi.Void](b_, "setAlternateImage:", value)
}

func (b_ Button) ImagePosition() CellImagePosition {
	rv := ffi.CallMethod[CellImagePosition](b_, "imagePosition")
	return rv
}

func (b_ Button) SetImagePosition(value CellImagePosition) {
	ffi.CallMethod[ffi.Void](b_, "setImagePosition:", value)
}

func (b_ Button) IsBordered() bool {
	rv := ffi.CallMethod[bool](b_, "isBordered")
	return rv
}

func (b_ Button) SetBordered(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setBordered:", value)
}

func (b_ Button) IsTransparent() bool {
	rv := ffi.CallMethod[bool](b_, "isTransparent")
	return rv
}

func (b_ Button) SetTransparent(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setTransparent:", value)
}

func (b_ Button) BezelStyle() BezelStyle {
	rv := ffi.CallMethod[BezelStyle](b_, "bezelStyle")
	return rv
}

func (b_ Button) SetBezelStyle(value BezelStyle) {
	ffi.CallMethod[ffi.Void](b_, "setBezelStyle:", value)
}

func (b_ Button) BezelColor() Color {
	rv := ffi.CallMethod[Color](b_, "bezelColor")
	return rv
}

func (b_ Button) SetBezelColor(value IColor) {
	ffi.CallMethod[ffi.Void](b_, "setBezelColor:", value)
}

func (b_ Button) ShowsBorderOnlyWhileMouseInside() bool {
	rv := ffi.CallMethod[bool](b_, "showsBorderOnlyWhileMouseInside")
	return rv
}

func (b_ Button) SetShowsBorderOnlyWhileMouseInside(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setShowsBorderOnlyWhileMouseInside:", value)
}

func (b_ Button) ImageHugsTitle() bool {
	rv := ffi.CallMethod[bool](b_, "imageHugsTitle")
	return rv
}

func (b_ Button) SetImageHugsTitle(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setImageHugsTitle:", value)
}

func (b_ Button) ImageScaling() ImageScaling {
	rv := ffi.CallMethod[ImageScaling](b_, "imageScaling")
	return rv
}

func (b_ Button) SetImageScaling(value ImageScaling) {
	ffi.CallMethod[ffi.Void](b_, "setImageScaling:", value)
}

func (b_ Button) ActiveCompressionOptions() UserInterfaceCompressionOptions {
	rv := ffi.CallMethod[UserInterfaceCompressionOptions](b_, "activeCompressionOptions")
	return rv
}

func (b_ Button) AllowsMixedState() bool {
	rv := ffi.CallMethod[bool](b_, "allowsMixedState")
	return rv
}

func (b_ Button) SetAllowsMixedState(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setAllowsMixedState:", value)
}

func (b_ Button) State() ControlStateValue {
	rv := ffi.CallMethod[ControlStateValue](b_, "state")
	return rv
}

func (b_ Button) SetState(value ControlStateValue) {
	ffi.CallMethod[ffi.Void](b_, "setState:", value)
}

func (b_ Button) KeyEquivalent() string {
	rv := ffi.CallMethod[string](b_, "keyEquivalent")
	return rv
}

func (b_ Button) SetKeyEquivalent(value string) {
	ffi.CallMethod[ffi.Void](b_, "setKeyEquivalent:", value)
}

func (b_ Button) KeyEquivalentModifierMask() EventModifierFlags {
	rv := ffi.CallMethod[EventModifierFlags](b_, "keyEquivalentModifierMask")
	return rv
}

func (b_ Button) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	ffi.CallMethod[ffi.Void](b_, "setKeyEquivalentModifierMask:", value)
}

var ButtonCellClass = _ButtonCellClass{objc.GetClass("NSButtonCell")}

type _ButtonCellClass struct {
	objc.Class
}

type IButtonCell interface {
	IActionCell
	// deprecated
	AlternateMnemonic() string
	// deprecated
	AlternateMnemonicLocation() uint
	// deprecated
	SetAlternateMnemonicLocation(location uint)
	// deprecated
	SetAlternateTitleWithMnemonic(stringWithAmpersand string)
	SetPeriodicDelay_Interval(delay float32, interval float32)
	// deprecated
	SetKeyEquivalentFont_Size(fontName string, fontSize float64)
	SetButtonType(_type ButtonType)
	MouseEntered(event IEvent)
	MouseExited(event IEvent)
	DrawBezelWithFrame_InView(frame foundation.Rect, controlView IView)
	DrawImage_WithFrame_InView(image IImage, frame foundation.Rect, controlView IView)
	DrawTitle_WithFrame_InView(title foundation.IAttributedString, frame foundation.Rect, controlView IView) foundation.Rect
	AlternateTitle() string
	SetAlternateTitle(value string)
	AttributedAlternateTitle() foundation.AttributedString
	SetAttributedAlternateTitle(value foundation.IAttributedString)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	AlternateImage() Image
	SetAlternateImage(value IImage)
	ImagePosition() CellImagePosition
	SetImagePosition(value CellImagePosition)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
	SetKeyEquivalent(value string)
	// deprecated
	KeyEquivalentFont() Font
	// deprecated
	SetKeyEquivalentFont(value IFont)
	KeyEquivalentModifierMask() EventModifierFlags
	SetKeyEquivalentModifierMask(value EventModifierFlags)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	BezelStyle() BezelStyle
	SetBezelStyle(value BezelStyle)
	// deprecated
	GradientType() GradientType
	// deprecated
	SetGradientType(value GradientType)
	ImageDimsWhenDisabled() bool
	SetImageDimsWhenDisabled(value bool)
	IsTransparent() bool
	SetTransparent(value bool)
	ShowsBorderOnlyWhileMouseInside() bool
	SetShowsBorderOnlyWhileMouseInside(value bool)
	HighlightsBy() CellStyleMask
	SetHighlightsBy(value CellStyleMask)
	ShowsStateBy() CellStyleMask
	SetShowsStateBy(value CellStyleMask)
	Sound() Sound
	SetSound(value ISound)
}

type ButtonCell struct {
	ActionCell
}

func MakeButtonCell(ptr unsafe.Pointer) ButtonCell {
	return ButtonCell{
		ActionCell: MakeActionCell(ptr),
	}
}

func (b_ ButtonCell) InitImageCell(image IImage) ButtonCell {
	rv := ffi.CallMethod[ButtonCell](b_, "initImageCell:", image)
	return rv
}

func (b_ ButtonCell) InitTextCell(_string string) ButtonCell {
	rv := ffi.CallMethod[ButtonCell](b_, "initTextCell:", _string)
	return rv
}

func (b_ ButtonCell) Init() ButtonCell {
	rv := ffi.CallMethod[ButtonCell](b_, "init")
	return rv
}

func (bc _ButtonCellClass) Alloc() ButtonCell {
	rv := ffi.CallMethod[ButtonCell](bc, "alloc")
	return rv
}

func (bc _ButtonCellClass) New() ButtonCell {
	rv := ffi.CallMethod[ButtonCell](bc, "new")
	rv.Autorelease()
	return rv
}

func NewButtonCell() ButtonCell {
	return ButtonCellClass.New()
}

// deprecated
func (b_ ButtonCell) AlternateMnemonic() string {
	rv := ffi.CallMethod[string](b_, "alternateMnemonic")
	return rv
}

// deprecated
func (b_ ButtonCell) AlternateMnemonicLocation() uint {
	rv := ffi.CallMethod[uint](b_, "alternateMnemonicLocation")
	return rv
}

// deprecated
func (b_ ButtonCell) SetAlternateMnemonicLocation(location uint) {
	ffi.CallMethod[ffi.Void](b_, "setAlternateMnemonicLocation:", location)
}

// deprecated
func (b_ ButtonCell) SetAlternateTitleWithMnemonic(stringWithAmpersand string) {
	ffi.CallMethod[ffi.Void](b_, "setAlternateTitleWithMnemonic:", stringWithAmpersand)
}

func (b_ ButtonCell) SetPeriodicDelay_Interval(delay float32, interval float32) {
	ffi.CallMethod[ffi.Void](b_, "setPeriodicDelay:interval:", delay, interval)
}

// deprecated
func (b_ ButtonCell) SetKeyEquivalentFont_Size(fontName string, fontSize float64) {
	ffi.CallMethod[ffi.Void](b_, "setKeyEquivalentFont:size:", fontName, fontSize)
}

func (b_ ButtonCell) SetButtonType(_type ButtonType) {
	ffi.CallMethod[ffi.Void](b_, "setButtonType:", _type)
}

func (b_ ButtonCell) MouseEntered(event IEvent) {
	ffi.CallMethod[ffi.Void](b_, "mouseEntered:", event)
}

func (b_ ButtonCell) MouseExited(event IEvent) {
	ffi.CallMethod[ffi.Void](b_, "mouseExited:", event)
}

func (b_ ButtonCell) DrawBezelWithFrame_InView(frame foundation.Rect, controlView IView) {
	ffi.CallMethod[ffi.Void](b_, "drawBezelWithFrame:inView:", frame, controlView)
}

func (b_ ButtonCell) DrawImage_WithFrame_InView(image IImage, frame foundation.Rect, controlView IView) {
	ffi.CallMethod[ffi.Void](b_, "drawImage:withFrame:inView:", image, frame, controlView)
}

func (b_ ButtonCell) DrawTitle_WithFrame_InView(title foundation.IAttributedString, frame foundation.Rect, controlView IView) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](b_, "drawTitle:withFrame:inView:", title, frame, controlView)
	return rv
}

func (b_ ButtonCell) AlternateTitle() string {
	rv := ffi.CallMethod[string](b_, "alternateTitle")
	return rv
}

func (b_ ButtonCell) SetAlternateTitle(value string) {
	ffi.CallMethod[ffi.Void](b_, "setAlternateTitle:", value)
}

func (b_ ButtonCell) AttributedAlternateTitle() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](b_, "attributedAlternateTitle")
	return rv
}

func (b_ ButtonCell) SetAttributedAlternateTitle(value foundation.IAttributedString) {
	ffi.CallMethod[ffi.Void](b_, "setAttributedAlternateTitle:", value)
}

func (b_ ButtonCell) AttributedTitle() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](b_, "attributedTitle")
	return rv
}

func (b_ ButtonCell) SetAttributedTitle(value foundation.IAttributedString) {
	ffi.CallMethod[ffi.Void](b_, "setAttributedTitle:", value)
}

func (b_ ButtonCell) AlternateImage() Image {
	rv := ffi.CallMethod[Image](b_, "alternateImage")
	return rv
}

func (b_ ButtonCell) SetAlternateImage(value IImage) {
	ffi.CallMethod[ffi.Void](b_, "setAlternateImage:", value)
}

func (b_ ButtonCell) ImagePosition() CellImagePosition {
	rv := ffi.CallMethod[CellImagePosition](b_, "imagePosition")
	return rv
}

func (b_ ButtonCell) SetImagePosition(value CellImagePosition) {
	ffi.CallMethod[ffi.Void](b_, "setImagePosition:", value)
}

func (b_ ButtonCell) ImageScaling() ImageScaling {
	rv := ffi.CallMethod[ImageScaling](b_, "imageScaling")
	return rv
}

func (b_ ButtonCell) SetImageScaling(value ImageScaling) {
	ffi.CallMethod[ffi.Void](b_, "setImageScaling:", value)
}

func (b_ ButtonCell) SetKeyEquivalent(value string) {
	ffi.CallMethod[ffi.Void](b_, "setKeyEquivalent:", value)
}

// deprecated
func (b_ ButtonCell) KeyEquivalentFont() Font {
	rv := ffi.CallMethod[Font](b_, "keyEquivalentFont")
	return rv
}

// deprecated
func (b_ ButtonCell) SetKeyEquivalentFont(value IFont) {
	ffi.CallMethod[ffi.Void](b_, "setKeyEquivalentFont:", value)
}

func (b_ ButtonCell) KeyEquivalentModifierMask() EventModifierFlags {
	rv := ffi.CallMethod[EventModifierFlags](b_, "keyEquivalentModifierMask")
	return rv
}

func (b_ ButtonCell) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	ffi.CallMethod[ffi.Void](b_, "setKeyEquivalentModifierMask:", value)
}

func (b_ ButtonCell) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](b_, "backgroundColor")
	return rv
}

func (b_ ButtonCell) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](b_, "setBackgroundColor:", value)
}

func (b_ ButtonCell) BezelStyle() BezelStyle {
	rv := ffi.CallMethod[BezelStyle](b_, "bezelStyle")
	return rv
}

func (b_ ButtonCell) SetBezelStyle(value BezelStyle) {
	ffi.CallMethod[ffi.Void](b_, "setBezelStyle:", value)
}

// deprecated
func (b_ ButtonCell) GradientType() GradientType {
	rv := ffi.CallMethod[GradientType](b_, "gradientType")
	return rv
}

// deprecated
func (b_ ButtonCell) SetGradientType(value GradientType) {
	ffi.CallMethod[ffi.Void](b_, "setGradientType:", value)
}

func (b_ ButtonCell) ImageDimsWhenDisabled() bool {
	rv := ffi.CallMethod[bool](b_, "imageDimsWhenDisabled")
	return rv
}

func (b_ ButtonCell) SetImageDimsWhenDisabled(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setImageDimsWhenDisabled:", value)
}

func (b_ ButtonCell) IsTransparent() bool {
	rv := ffi.CallMethod[bool](b_, "isTransparent")
	return rv
}

func (b_ ButtonCell) SetTransparent(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setTransparent:", value)
}

func (b_ ButtonCell) ShowsBorderOnlyWhileMouseInside() bool {
	rv := ffi.CallMethod[bool](b_, "showsBorderOnlyWhileMouseInside")
	return rv
}

func (b_ ButtonCell) SetShowsBorderOnlyWhileMouseInside(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setShowsBorderOnlyWhileMouseInside:", value)
}

func (b_ ButtonCell) HighlightsBy() CellStyleMask {
	rv := ffi.CallMethod[CellStyleMask](b_, "highlightsBy")
	return rv
}

func (b_ ButtonCell) SetHighlightsBy(value CellStyleMask) {
	ffi.CallMethod[ffi.Void](b_, "setHighlightsBy:", value)
}

func (b_ ButtonCell) ShowsStateBy() CellStyleMask {
	rv := ffi.CallMethod[CellStyleMask](b_, "showsStateBy")
	return rv
}

func (b_ ButtonCell) SetShowsStateBy(value CellStyleMask) {
	ffi.CallMethod[ffi.Void](b_, "setShowsStateBy:", value)
}

func (b_ ButtonCell) Sound() Sound {
	rv := ffi.CallMethod[Sound](b_, "sound")
	return rv
}

func (b_ ButtonCell) SetSound(value ISound) {
	ffi.CallMethod[ffi.Void](b_, "setSound:", value)
}
