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
	rv := objc.CallMethod[StatusItem](sc, objc.GetSelector("alloc"))
	return rv
}

func (sc _StatusItemClass) New() StatusItem {
	rv := objc.CallMethod[StatusItem](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewStatusItem() StatusItem {
	return StatusItemClass.New()
}

func (s_ StatusItem) Init() StatusItem {
	rv := objc.CallMethod[StatusItem](s_, objc.GetSelector("init"))
	return rv
}

// deprecated
func (s_ StatusItem) SendActionOn(mask EventMask) int {
	rv := objc.CallMethod[int](s_, objc.GetSelector("sendActionOn:"), mask)
	return rv
}

// deprecated
func (s_ StatusItem) PopUpStatusItemMenu(menu IMenu) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("popUpStatusItemMenu:"), objc.ExtractPtr(menu))
}

// deprecated
func (s_ StatusItem) DrawStatusBarBackgroundInRect_WithHighlight(rect foundation.Rect, highlight bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("drawStatusBarBackgroundInRect:withHighlight:"), rect, highlight)
}

// weak property
func (s_ StatusItem) StatusBar() StatusBar {
	rv := objc.CallMethod[StatusBar](s_, objc.GetSelector("statusBar"))
	return rv
}

// weak property
func (s_ StatusItem) Behavior() StatusItemBehavior {
	rv := objc.CallMethod[StatusItemBehavior](s_, objc.GetSelector("behavior"))
	return rv
}

// weak property
func (s_ StatusItem) SetBehavior(value StatusItemBehavior) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setBehavior:"), value)
}

func (s_ StatusItem) Button() StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](s_, objc.GetSelector("button"))
	return rv
}

func (s_ StatusItem) Menu() Menu {
	rv := objc.CallMethod[Menu](s_, objc.GetSelector("menu"))
	return rv
}

func (s_ StatusItem) SetMenu(value IMenu) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setMenu:"), objc.ExtractPtr(value))
}

// weak property
func (s_ StatusItem) IsVisible() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("isVisible"))
	return rv
}

// weak property
func (s_ StatusItem) SetVisible(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setVisible:"), value)
}

func (s_ StatusItem) Length() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("length"))
	return rv
}

func (s_ StatusItem) SetLength(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setLength:"), value)
}

func (s_ StatusItem) AutosaveName() StatusItemAutosaveName {
	rv := objc.CallMethod[StatusItemAutosaveName](s_, objc.GetSelector("autosaveName"))
	return rv
}

func (s_ StatusItem) SetAutosaveName(value StatusItemAutosaveName) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAutosaveName:"), value)
}

// deprecated
func (s_ StatusItem) IsEnabled() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("isEnabled"))
	return rv
}

// deprecated
func (s_ StatusItem) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setEnabled:"), value)
}

// weak property
// deprecated
func (s_ StatusItem) Target() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.GetSelector("target"))
	return rv
}

// weak property
// deprecated
func (s_ StatusItem) SetTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setTarget:"), objc.ExtractPtr(value))
}

// deprecated
func (s_ StatusItem) Action() objc.Selector {
	rv := objc.CallMethod[objc.Selector](s_, objc.GetSelector("action"))
	return rv
}

// deprecated
func (s_ StatusItem) SetAction(value objc.Selector) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAction:"), value)
}

// deprecated
func (s_ StatusItem) DoubleAction() objc.Selector {
	rv := objc.CallMethod[objc.Selector](s_, objc.GetSelector("doubleAction"))
	return rv
}

// deprecated
func (s_ StatusItem) SetDoubleAction(value objc.Selector) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDoubleAction:"), value)
}

// deprecated
func (s_ StatusItem) Title() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("title"))
	return rv
}

// deprecated
func (s_ StatusItem) SetTitle(value string) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setTitle:"), value)
}

// deprecated
func (s_ StatusItem) AttributedTitle() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](s_, objc.GetSelector("attributedTitle"))
	return rv
}

// deprecated
func (s_ StatusItem) SetAttributedTitle(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAttributedTitle:"), objc.ExtractPtr(value))
}

// deprecated
func (s_ StatusItem) Image() Image {
	rv := objc.CallMethod[Image](s_, objc.GetSelector("image"))
	return rv
}

// deprecated
func (s_ StatusItem) SetImage(value IImage) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setImage:"), objc.ExtractPtr(value))
}

// deprecated
func (s_ StatusItem) AlternateImage() Image {
	rv := objc.CallMethod[Image](s_, objc.GetSelector("alternateImage"))
	return rv
}

// deprecated
func (s_ StatusItem) SetAlternateImage(value IImage) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAlternateImage:"), objc.ExtractPtr(value))
}

// deprecated
func (s_ StatusItem) HighlightMode() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("highlightMode"))
	return rv
}

// deprecated
func (s_ StatusItem) SetHighlightMode(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setHighlightMode:"), value)
}

// deprecated
func (s_ StatusItem) ToolTip() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("toolTip"))
	return rv
}

// deprecated
func (s_ StatusItem) SetToolTip(value string) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setToolTip:"), value)
}

// deprecated
func (s_ StatusItem) View() View {
	rv := objc.CallMethod[View](s_, objc.GetSelector("view"))
	return rv
}

// deprecated
func (s_ StatusItem) SetView(value IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setView:"), objc.ExtractPtr(value))
}
