// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var StatusItemClass = _StatusItemClass{objc.GetClass("NSStatusItem")}

type _StatusItemClass struct {
	objc.Class
}

type IStatusItem interface {
	objc.IObject
	// deprecated
	SendActionOn(mask EventMask) int
	// deprecated
	PopUpStatusItemMenu(menu IMenu)
	// deprecated
	DrawStatusBarBackgroundInRect_WithHighlight(rect foundation.Rect, highlight bool)
	StatusBar() StatusBar
	Behavior() StatusItemBehavior
	SetBehavior(value StatusItemBehavior)
	Button() StatusBarButton
	Menu() Menu
	SetMenu(value IMenu)
	IsVisible() bool
	SetVisible(value bool)
	Length() float64
	SetLength(value float64)
	AutosaveName() StatusItemAutosaveName
	SetAutosaveName(value StatusItemAutosaveName)
	// deprecated
	IsEnabled() bool
	// deprecated
	SetEnabled(value bool)
	// deprecated
	Target() objc.Object
	// deprecated
	SetTarget(value objc.IObject)
	// deprecated
	Action() objc.Selector
	// deprecated
	SetAction(value objc.Selector)
	// deprecated
	DoubleAction() objc.Selector
	// deprecated
	SetDoubleAction(value objc.Selector)
	// deprecated
	Title() string
	// deprecated
	SetTitle(value string)
	// deprecated
	AttributedTitle() foundation.AttributedString
	// deprecated
	SetAttributedTitle(value foundation.IAttributedString)
	// deprecated
	Image() Image
	// deprecated
	SetImage(value IImage)
	// deprecated
	AlternateImage() Image
	// deprecated
	SetAlternateImage(value IImage)
	// deprecated
	HighlightMode() bool
	// deprecated
	SetHighlightMode(value bool)
	// deprecated
	ToolTip() string
	// deprecated
	SetToolTip(value string)
	// deprecated
	View() View
	// deprecated
	SetView(value IView)
}

type StatusItem struct {
	objc.Object
}

func MakeStatusItem(ptr unsafe.Pointer) StatusItem {
	return StatusItem{
		Object: objc.MakeObject(ptr),
	}
}

func (sc _StatusItemClass) Alloc() StatusItem {
	rv := ffi.CallMethod[StatusItem](sc, "alloc")
	return rv
}

func (sc _StatusItemClass) New() StatusItem {
	rv := ffi.CallMethod[StatusItem](sc, "new")
	rv.Autorelease()
	return rv
}

func NewStatusItem() StatusItem {
	return StatusItemClass.New()
}

func (s_ StatusItem) Init() StatusItem {
	rv := ffi.CallMethod[StatusItem](s_, "init")
	return rv
}

// deprecated
func (s_ StatusItem) SendActionOn(mask EventMask) int {
	rv := ffi.CallMethod[int](s_, "sendActionOn:", mask)
	return rv
}

// deprecated
func (s_ StatusItem) PopUpStatusItemMenu(menu IMenu) {
	ffi.CallMethod[ffi.Void](s_, "popUpStatusItemMenu:", menu)
}

// deprecated
func (s_ StatusItem) DrawStatusBarBackgroundInRect_WithHighlight(rect foundation.Rect, highlight bool) {
	ffi.CallMethod[ffi.Void](s_, "drawStatusBarBackgroundInRect:withHighlight:", rect, highlight)
}

func (s_ StatusItem) StatusBar() StatusBar {
	rv := ffi.CallMethod[StatusBar](s_, "statusBar")
	return rv
}

func (s_ StatusItem) Behavior() StatusItemBehavior {
	rv := ffi.CallMethod[StatusItemBehavior](s_, "behavior")
	return rv
}

func (s_ StatusItem) SetBehavior(value StatusItemBehavior) {
	ffi.CallMethod[ffi.Void](s_, "setBehavior:", value)
}

func (s_ StatusItem) Button() StatusBarButton {
	rv := ffi.CallMethod[StatusBarButton](s_, "button")
	return rv
}

func (s_ StatusItem) Menu() Menu {
	rv := ffi.CallMethod[Menu](s_, "menu")
	return rv
}

func (s_ StatusItem) SetMenu(value IMenu) {
	ffi.CallMethod[ffi.Void](s_, "setMenu:", value)
}

func (s_ StatusItem) IsVisible() bool {
	rv := ffi.CallMethod[bool](s_, "isVisible")
	return rv
}

func (s_ StatusItem) SetVisible(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setVisible:", value)
}

func (s_ StatusItem) Length() float64 {
	rv := ffi.CallMethod[float64](s_, "length")
	return rv
}

func (s_ StatusItem) SetLength(value float64) {
	ffi.CallMethod[ffi.Void](s_, "setLength:", value)
}

func (s_ StatusItem) AutosaveName() StatusItemAutosaveName {
	rv := ffi.CallMethod[StatusItemAutosaveName](s_, "autosaveName")
	return rv
}

func (s_ StatusItem) SetAutosaveName(value StatusItemAutosaveName) {
	ffi.CallMethod[ffi.Void](s_, "setAutosaveName:", value)
}

// deprecated
func (s_ StatusItem) IsEnabled() bool {
	rv := ffi.CallMethod[bool](s_, "isEnabled")
	return rv
}

// deprecated
func (s_ StatusItem) SetEnabled(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setEnabled:", value)
}

// deprecated
func (s_ StatusItem) Target() objc.Object {
	rv := ffi.CallMethod[objc.Object](s_, "target")
	return rv
}

// deprecated
func (s_ StatusItem) SetTarget(value objc.IObject) {
	ffi.CallMethod[ffi.Void](s_, "setTarget:", value)
}

// deprecated
func (s_ StatusItem) Action() objc.Selector {
	rv := ffi.CallMethod[objc.Selector](s_, "action")
	return rv
}

// deprecated
func (s_ StatusItem) SetAction(value objc.Selector) {
	ffi.CallMethod[ffi.Void](s_, "setAction:", value)
}

// deprecated
func (s_ StatusItem) DoubleAction() objc.Selector {
	rv := ffi.CallMethod[objc.Selector](s_, "doubleAction")
	return rv
}

// deprecated
func (s_ StatusItem) SetDoubleAction(value objc.Selector) {
	ffi.CallMethod[ffi.Void](s_, "setDoubleAction:", value)
}

// deprecated
func (s_ StatusItem) Title() string {
	rv := ffi.CallMethod[string](s_, "title")
	return rv
}

// deprecated
func (s_ StatusItem) SetTitle(value string) {
	ffi.CallMethod[ffi.Void](s_, "setTitle:", value)
}

// deprecated
func (s_ StatusItem) AttributedTitle() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](s_, "attributedTitle")
	return rv
}

// deprecated
func (s_ StatusItem) SetAttributedTitle(value foundation.IAttributedString) {
	ffi.CallMethod[ffi.Void](s_, "setAttributedTitle:", value)
}

// deprecated
func (s_ StatusItem) Image() Image {
	rv := ffi.CallMethod[Image](s_, "image")
	return rv
}

// deprecated
func (s_ StatusItem) SetImage(value IImage) {
	ffi.CallMethod[ffi.Void](s_, "setImage:", value)
}

// deprecated
func (s_ StatusItem) AlternateImage() Image {
	rv := ffi.CallMethod[Image](s_, "alternateImage")
	return rv
}

// deprecated
func (s_ StatusItem) SetAlternateImage(value IImage) {
	ffi.CallMethod[ffi.Void](s_, "setAlternateImage:", value)
}

// deprecated
func (s_ StatusItem) HighlightMode() bool {
	rv := ffi.CallMethod[bool](s_, "highlightMode")
	return rv
}

// deprecated
func (s_ StatusItem) SetHighlightMode(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setHighlightMode:", value)
}

// deprecated
func (s_ StatusItem) ToolTip() string {
	rv := ffi.CallMethod[string](s_, "toolTip")
	return rv
}

// deprecated
func (s_ StatusItem) SetToolTip(value string) {
	ffi.CallMethod[ffi.Void](s_, "setToolTip:", value)
}

// deprecated
func (s_ StatusItem) View() View {
	rv := ffi.CallMethod[View](s_, "view")
	return rv
}

// deprecated
func (s_ StatusItem) SetView(value IView) {
	ffi.CallMethod[ffi.Void](s_, "setView:", value)
}
