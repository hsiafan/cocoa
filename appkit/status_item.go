// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[StatusItem](sc, "alloc")
	return rv
}

func (sc _StatusItemClass) New() StatusItem {
	rv := objc.CallMethod[StatusItem](sc, "new")
	rv.Autorelease()
	return rv
}

func NewStatusItem() StatusItem {
	return StatusItemClass.New()
}

func (s_ StatusItem) Init() StatusItem {
	rv := objc.CallMethod[StatusItem](s_, "init")
	return rv
}

// deprecated
func (s_ StatusItem) SendActionOn(mask EventMask) int {
	rv := objc.CallMethod[int](s_, "sendActionOn:", mask)
	return rv
}

// deprecated
func (s_ StatusItem) PopUpStatusItemMenu(menu IMenu) {
	objc.CallMethod[objc.Void](s_, "popUpStatusItemMenu:", menu)
}

// deprecated
func (s_ StatusItem) DrawStatusBarBackgroundInRect_WithHighlight(rect foundation.Rect, highlight bool) {
	objc.CallMethod[objc.Void](s_, "drawStatusBarBackgroundInRect:withHighlight:", rect, highlight)
}

func (s_ StatusItem) StatusBar() StatusBar {
	rv := objc.CallMethod[StatusBar](s_, "statusBar")
	return rv
}

func (s_ StatusItem) Behavior() StatusItemBehavior {
	rv := objc.CallMethod[StatusItemBehavior](s_, "behavior")
	return rv
}

func (s_ StatusItem) SetBehavior(value StatusItemBehavior) {
	objc.CallMethod[objc.Void](s_, "setBehavior:", value)
}

func (s_ StatusItem) Button() StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](s_, "button")
	return rv
}

func (s_ StatusItem) Menu() Menu {
	rv := objc.CallMethod[Menu](s_, "menu")
	return rv
}

func (s_ StatusItem) SetMenu(value IMenu) {
	objc.CallMethod[objc.Void](s_, "setMenu:", value)
}

func (s_ StatusItem) IsVisible() bool {
	rv := objc.CallMethod[bool](s_, "isVisible")
	return rv
}

func (s_ StatusItem) SetVisible(value bool) {
	objc.CallMethod[objc.Void](s_, "setVisible:", value)
}

func (s_ StatusItem) Length() float64 {
	rv := objc.CallMethod[float64](s_, "length")
	return rv
}

func (s_ StatusItem) SetLength(value float64) {
	objc.CallMethod[objc.Void](s_, "setLength:", value)
}

func (s_ StatusItem) AutosaveName() StatusItemAutosaveName {
	rv := objc.CallMethod[StatusItemAutosaveName](s_, "autosaveName")
	return rv
}

func (s_ StatusItem) SetAutosaveName(value StatusItemAutosaveName) {
	objc.CallMethod[objc.Void](s_, "setAutosaveName:", value)
}

// deprecated
func (s_ StatusItem) IsEnabled() bool {
	rv := objc.CallMethod[bool](s_, "isEnabled")
	return rv
}

// deprecated
func (s_ StatusItem) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](s_, "setEnabled:", value)
}

// deprecated
func (s_ StatusItem) Target() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, "target")
	return rv
}

// deprecated
func (s_ StatusItem) SetTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, "setTarget:", value)
}

// deprecated
func (s_ StatusItem) Action() objc.Selector {
	rv := objc.CallMethod[objc.Selector](s_, "action")
	return rv
}

// deprecated
func (s_ StatusItem) SetAction(value objc.Selector) {
	objc.CallMethod[objc.Void](s_, "setAction:", value)
}

// deprecated
func (s_ StatusItem) DoubleAction() objc.Selector {
	rv := objc.CallMethod[objc.Selector](s_, "doubleAction")
	return rv
}

// deprecated
func (s_ StatusItem) SetDoubleAction(value objc.Selector) {
	objc.CallMethod[objc.Void](s_, "setDoubleAction:", value)
}

// deprecated
func (s_ StatusItem) Title() string {
	rv := objc.CallMethod[string](s_, "title")
	return rv
}

// deprecated
func (s_ StatusItem) SetTitle(value string) {
	objc.CallMethod[objc.Void](s_, "setTitle:", value)
}

// deprecated
func (s_ StatusItem) AttributedTitle() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](s_, "attributedTitle")
	return rv
}

// deprecated
func (s_ StatusItem) SetAttributedTitle(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](s_, "setAttributedTitle:", value)
}

// deprecated
func (s_ StatusItem) Image() Image {
	rv := objc.CallMethod[Image](s_, "image")
	return rv
}

// deprecated
func (s_ StatusItem) SetImage(value IImage) {
	objc.CallMethod[objc.Void](s_, "setImage:", value)
}

// deprecated
func (s_ StatusItem) AlternateImage() Image {
	rv := objc.CallMethod[Image](s_, "alternateImage")
	return rv
}

// deprecated
func (s_ StatusItem) SetAlternateImage(value IImage) {
	objc.CallMethod[objc.Void](s_, "setAlternateImage:", value)
}

// deprecated
func (s_ StatusItem) HighlightMode() bool {
	rv := objc.CallMethod[bool](s_, "highlightMode")
	return rv
}

// deprecated
func (s_ StatusItem) SetHighlightMode(value bool) {
	objc.CallMethod[objc.Void](s_, "setHighlightMode:", value)
}

// deprecated
func (s_ StatusItem) ToolTip() string {
	rv := objc.CallMethod[string](s_, "toolTip")
	return rv
}

// deprecated
func (s_ StatusItem) SetToolTip(value string) {
	objc.CallMethod[objc.Void](s_, "setToolTip:", value)
}

// deprecated
func (s_ StatusItem) View() View {
	rv := objc.CallMethod[View](s_, "view")
	return rv
}

// deprecated
func (s_ StatusItem) SetView(value IView) {
	objc.CallMethod[objc.Void](s_, "setView:", value)
}
