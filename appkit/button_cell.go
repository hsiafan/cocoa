// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[ButtonCell](b_, objc.SEL("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func (b_ ButtonCell) InitTextCell(string_ string) ButtonCell {
	rv := objc.CallMethod[ButtonCell](b_, objc.SEL("initTextCell:"), string_)
	return rv
}

func (b_ ButtonCell) Init() ButtonCell {
	rv := objc.CallMethod[ButtonCell](b_, objc.SEL("init"))
	return rv
}

func (bc _ButtonCellClass) Alloc() ButtonCell {
	rv := objc.CallMethod[ButtonCell](bc, objc.SEL("alloc"))
	return rv
}

func (bc _ButtonCellClass) New() ButtonCell {
	rv := objc.CallMethod[ButtonCell](bc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewButtonCell() ButtonCell {
	return ButtonCellClass.New()
}

// deprecated
func (b_ ButtonCell) AlternateMnemonic() string {
	rv := objc.CallMethod[string](b_, objc.SEL("alternateMnemonic"))
	return rv
}

// deprecated
func (b_ ButtonCell) AlternateMnemonicLocation() uint {
	rv := objc.CallMethod[uint](b_, objc.SEL("alternateMnemonicLocation"))
	return rv
}

// deprecated
func (b_ ButtonCell) SetAlternateMnemonicLocation(location uint) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setAlternateMnemonicLocation:"), location)
}

// deprecated
func (b_ ButtonCell) SetAlternateTitleWithMnemonic(stringWithAmpersand string) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setAlternateTitleWithMnemonic:"), stringWithAmpersand)
}

func (b_ ButtonCell) SetPeriodicDelay_Interval(delay float32, interval float32) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setPeriodicDelay:interval:"), delay, interval)
}

// deprecated
func (b_ ButtonCell) SetKeyEquivalentFont_Size(fontName string, fontSize float64) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setKeyEquivalentFont:size:"), fontName, fontSize)
}

func (b_ ButtonCell) SetButtonType(type_ ButtonType) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setButtonType:"), type_)
}

func (b_ ButtonCell) MouseEntered(event IEvent) {
	objc.CallMethod[objc.Void](b_, objc.SEL("mouseEntered:"), objc.ExtractPtr(event))
}

func (b_ ButtonCell) MouseExited(event IEvent) {
	objc.CallMethod[objc.Void](b_, objc.SEL("mouseExited:"), objc.ExtractPtr(event))
}

func (b_ ButtonCell) DrawBezelWithFrame_InView(frame foundation.Rect, controlView IView) {
	objc.CallMethod[objc.Void](b_, objc.SEL("drawBezelWithFrame:inView:"), frame, objc.ExtractPtr(controlView))
}

func (b_ ButtonCell) DrawImage_WithFrame_InView(image IImage, frame foundation.Rect, controlView IView) {
	objc.CallMethod[objc.Void](b_, objc.SEL("drawImage:withFrame:inView:"), objc.ExtractPtr(image), frame, objc.ExtractPtr(controlView))
}

func (b_ ButtonCell) DrawTitle_WithFrame_InView(title foundation.IAttributedString, frame foundation.Rect, controlView IView) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](b_, objc.SEL("drawTitle:withFrame:inView:"), objc.ExtractPtr(title), frame, objc.ExtractPtr(controlView))
	return rv
}

func (b_ ButtonCell) AlternateTitle() string {
	rv := objc.CallMethod[string](b_, objc.SEL("alternateTitle"))
	return rv
}

func (b_ ButtonCell) SetAlternateTitle(value string) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setAlternateTitle:"), value)
}

func (b_ ButtonCell) AttributedAlternateTitle() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](b_, objc.SEL("attributedAlternateTitle"))
	return rv
}

func (b_ ButtonCell) SetAttributedAlternateTitle(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setAttributedAlternateTitle:"), objc.ExtractPtr(value))
}

func (b_ ButtonCell) AttributedTitle() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](b_, objc.SEL("attributedTitle"))
	return rv
}

func (b_ ButtonCell) SetAttributedTitle(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setAttributedTitle:"), objc.ExtractPtr(value))
}

func (b_ ButtonCell) AlternateImage() Image {
	rv := objc.CallMethod[Image](b_, objc.SEL("alternateImage"))
	return rv
}

func (b_ ButtonCell) SetAlternateImage(value IImage) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setAlternateImage:"), objc.ExtractPtr(value))
}

func (b_ ButtonCell) ImagePosition() CellImagePosition {
	rv := objc.CallMethod[CellImagePosition](b_, objc.SEL("imagePosition"))
	return rv
}

func (b_ ButtonCell) SetImagePosition(value CellImagePosition) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setImagePosition:"), value)
}

func (b_ ButtonCell) ImageScaling() ImageScaling {
	rv := objc.CallMethod[ImageScaling](b_, objc.SEL("imageScaling"))
	return rv
}

func (b_ ButtonCell) SetImageScaling(value ImageScaling) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setImageScaling:"), value)
}

func (b_ ButtonCell) SetKeyEquivalent(value string) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setKeyEquivalent:"), value)
}

// deprecated
func (b_ ButtonCell) KeyEquivalentFont() Font {
	rv := objc.CallMethod[Font](b_, objc.SEL("keyEquivalentFont"))
	return rv
}

// deprecated
func (b_ ButtonCell) SetKeyEquivalentFont(value IFont) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setKeyEquivalentFont:"), objc.ExtractPtr(value))
}

func (b_ ButtonCell) KeyEquivalentModifierMask() EventModifierFlags {
	rv := objc.CallMethod[EventModifierFlags](b_, objc.SEL("keyEquivalentModifierMask"))
	return rv
}

func (b_ ButtonCell) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setKeyEquivalentModifierMask:"), value)
}

func (b_ ButtonCell) BackgroundColor() Color {
	rv := objc.CallMethod[Color](b_, objc.SEL("backgroundColor"))
	return rv
}

func (b_ ButtonCell) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (b_ ButtonCell) BezelStyle() BezelStyle {
	rv := objc.CallMethod[BezelStyle](b_, objc.SEL("bezelStyle"))
	return rv
}

func (b_ ButtonCell) SetBezelStyle(value BezelStyle) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setBezelStyle:"), value)
}

// deprecated
func (b_ ButtonCell) GradientType() GradientType {
	rv := objc.CallMethod[GradientType](b_, objc.SEL("gradientType"))
	return rv
}

// deprecated
func (b_ ButtonCell) SetGradientType(value GradientType) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setGradientType:"), value)
}

func (b_ ButtonCell) ImageDimsWhenDisabled() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("imageDimsWhenDisabled"))
	return rv
}

func (b_ ButtonCell) SetImageDimsWhenDisabled(value bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setImageDimsWhenDisabled:"), value)
}

func (b_ ButtonCell) IsTransparent() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("isTransparent"))
	return rv
}

func (b_ ButtonCell) SetTransparent(value bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setTransparent:"), value)
}

func (b_ ButtonCell) ShowsBorderOnlyWhileMouseInside() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("showsBorderOnlyWhileMouseInside"))
	return rv
}

func (b_ ButtonCell) SetShowsBorderOnlyWhileMouseInside(value bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setShowsBorderOnlyWhileMouseInside:"), value)
}

func (b_ ButtonCell) HighlightsBy() CellStyleMask {
	rv := objc.CallMethod[CellStyleMask](b_, objc.SEL("highlightsBy"))
	return rv
}

func (b_ ButtonCell) SetHighlightsBy(value CellStyleMask) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setHighlightsBy:"), value)
}

func (b_ ButtonCell) ShowsStateBy() CellStyleMask {
	rv := objc.CallMethod[CellStyleMask](b_, objc.SEL("showsStateBy"))
	return rv
}

func (b_ ButtonCell) SetShowsStateBy(value CellStyleMask) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setShowsStateBy:"), value)
}

func (b_ ButtonCell) Sound() Sound {
	rv := objc.CallMethod[Sound](b_, objc.SEL("sound"))
	return rv
}

func (b_ ButtonCell) SetSound(value ISound) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setSound:"), objc.ExtractPtr(value))
}
