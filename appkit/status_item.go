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
	rv := objc.CallMethod[StatusItem](sc, objc.SEL("alloc"))
	return rv
}

func (sc _StatusItemClass) New() StatusItem {
	rv := objc.CallMethod[StatusItem](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewStatusItem() StatusItem {
	return StatusItemClass.New()
}

func (s_ StatusItem) Init() StatusItem {
	rv := objc.CallMethod[StatusItem](s_, objc.SEL("init"))
	return rv
}

// deprecated
func (s_ StatusItem) SendActionOn(mask EventMask) int {
	rv := objc.CallMethod[int](s_, objc.SEL("sendActionOn:"), mask)
	return rv
}

// deprecated
func (s_ StatusItem) PopUpStatusItemMenu(menu IMenu) {
	objc.CallMethod[objc.Void](s_, objc.SEL("popUpStatusItemMenu:"), objc.ExtractPtr(menu))
}

// deprecated
func (s_ StatusItem) DrawStatusBarBackgroundInRect_WithHighlight(rect foundation.Rect, highlight bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("drawStatusBarBackgroundInRect:withHighlight:"), rect, highlight)
}

// weak property
func (s_ StatusItem) StatusBar() StatusBar {
	rv := objc.CallMethod[StatusBar](s_, objc.SEL("statusBar"))
	return rv
}

// weak property
func (s_ StatusItem) Behavior() StatusItemBehavior {
	rv := objc.CallMethod[StatusItemBehavior](s_, objc.SEL("behavior"))
	return rv
}

// weak property
func (s_ StatusItem) SetBehavior(value StatusItemBehavior) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setBehavior:"), value)
}

func (s_ StatusItem) Button() StatusBarButton {
	rv := objc.CallMethod[StatusBarButton](s_, objc.SEL("button"))
	return rv
}

func (s_ StatusItem) Menu() Menu {
	rv := objc.CallMethod[Menu](s_, objc.SEL("menu"))
	return rv
}

func (s_ StatusItem) SetMenu(value IMenu) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setMenu:"), objc.ExtractPtr(value))
}

// weak property
func (s_ StatusItem) IsVisible() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("isVisible"))
	return rv
}

// weak property
func (s_ StatusItem) SetVisible(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setVisible:"), value)
}

func (s_ StatusItem) Length() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("length"))
	return rv
}

func (s_ StatusItem) SetLength(value float64) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setLength:"), value)
}

func (s_ StatusItem) AutosaveName() StatusItemAutosaveName {
	rv := objc.CallMethod[StatusItemAutosaveName](s_, objc.SEL("autosaveName"))
	return rv
}

func (s_ StatusItem) SetAutosaveName(value StatusItemAutosaveName) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setAutosaveName:"), value)
}

// deprecated
func (s_ StatusItem) IsEnabled() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("isEnabled"))
	return rv
}

// deprecated
func (s_ StatusItem) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setEnabled:"), value)
}

// weak property
// deprecated
func (s_ StatusItem) Target() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.SEL("target"))
	return rv
}

// weak property
// deprecated
func (s_ StatusItem) SetTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setTarget:"), objc.ExtractPtr(value))
}

// deprecated
func (s_ StatusItem) Action() objc.Selector {
	rv := objc.CallMethod[objc.Selector](s_, objc.SEL("action"))
	return rv
}

// deprecated
func (s_ StatusItem) SetAction(value objc.Selector) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setAction:"), value)
}

// deprecated
func (s_ StatusItem) DoubleAction() objc.Selector {
	rv := objc.CallMethod[objc.Selector](s_, objc.SEL("doubleAction"))
	return rv
}

// deprecated
func (s_ StatusItem) SetDoubleAction(value objc.Selector) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setDoubleAction:"), value)
}

// deprecated
func (s_ StatusItem) Title() string {
	rv := objc.CallMethod[string](s_, objc.SEL("title"))
	return rv
}

// deprecated
func (s_ StatusItem) SetTitle(value string) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setTitle:"), value)
}

// deprecated
func (s_ StatusItem) AttributedTitle() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](s_, objc.SEL("attributedTitle"))
	return rv
}

// deprecated
func (s_ StatusItem) SetAttributedTitle(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setAttributedTitle:"), objc.ExtractPtr(value))
}

// deprecated
func (s_ StatusItem) Image() Image {
	rv := objc.CallMethod[Image](s_, objc.SEL("image"))
	return rv
}

// deprecated
func (s_ StatusItem) SetImage(value IImage) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setImage:"), objc.ExtractPtr(value))
}

// deprecated
func (s_ StatusItem) AlternateImage() Image {
	rv := objc.CallMethod[Image](s_, objc.SEL("alternateImage"))
	return rv
}

// deprecated
func (s_ StatusItem) SetAlternateImage(value IImage) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setAlternateImage:"), objc.ExtractPtr(value))
}

// deprecated
func (s_ StatusItem) HighlightMode() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("highlightMode"))
	return rv
}

// deprecated
func (s_ StatusItem) SetHighlightMode(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setHighlightMode:"), value)
}

// deprecated
func (s_ StatusItem) ToolTip() string {
	rv := objc.CallMethod[string](s_, objc.SEL("toolTip"))
	return rv
}

// deprecated
func (s_ StatusItem) SetToolTip(value string) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setToolTip:"), value)
}

// deprecated
func (s_ StatusItem) View() View {
	rv := objc.CallMethod[View](s_, objc.SEL("view"))
	return rv
}

// deprecated
func (s_ StatusItem) SetView(value IView) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setView:"), objc.ExtractPtr(value))
}
