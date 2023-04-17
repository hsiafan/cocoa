// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

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
	SetButtonType(type_ ButtonType)
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

func (b_ ButtonCell) InitTextCell(string_ string) ButtonCell {
	rv := ffi.CallMethod[ButtonCell](b_, "initTextCell:", string_)
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

func (b_ ButtonCell) SetButtonType(type_ ButtonType) {
	ffi.CallMethod[ffi.Void](b_, "setButtonType:", type_)
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
