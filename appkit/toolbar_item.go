// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ToolbarItemClass = _ToolbarItemClass{objc.GetClass("NSToolbarItem")}

type _ToolbarItemClass struct {
	objc.Class
}

type IToolbarItem interface {
	objc.IObject
	Validate()
	ItemIdentifier() ToolbarItemIdentifier
	PossibleLabels() foundation.Set
	SetPossibleLabels(value foundation.ISet)
	Label() string
	SetLabel(value string)
	PaletteLabel() string
	SetPaletteLabel(value string)
	Title() string
	SetTitle(value string)
	ToolTip() string
	SetToolTip(value string)
	Image() Image
	SetImage(value IImage)
	View() View
	SetView(value IView)
	Target() objc.Object
	SetTarget(value objc.IObject)
	Action() objc.Selector
	SetAction(value objc.Selector)
	MenuFormRepresentation() MenuItem
	SetMenuFormRepresentation(value IMenuItem)
	IsVisible() bool
	IsBordered() bool
	SetBordered(value bool)
	IsNavigational() bool
	SetNavigational(value bool)
	IsEnabled() bool
	SetEnabled(value bool)
	AllowsDuplicatesInToolbar() bool
	VisibilityPriority() ToolbarItemVisibilityPriority
	SetVisibilityPriority(value ToolbarItemVisibilityPriority)
	Tag() int
	SetTag(value int)
	Toolbar() Toolbar
	Autovalidates() bool
	SetAutovalidates(value bool)
	// deprecated
	MinSize() foundation.Size
	// deprecated
	SetMinSize(value foundation.Size)
	// deprecated
	MaxSize() foundation.Size
	// deprecated
	SetMaxSize(value foundation.Size)
}

type ToolbarItem struct {
	objc.Object
}

func MakeToolbarItem(ptr unsafe.Pointer) ToolbarItem {
	return ToolbarItem{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ ToolbarItem) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) ToolbarItem {
	rv := objc.CallMethod[ToolbarItem](t_, objc.GetSelector("initWithItemIdentifier:"), itemIdentifier)
	return rv
}

func (tc _ToolbarItemClass) Alloc() ToolbarItem {
	rv := objc.CallMethod[ToolbarItem](tc, objc.GetSelector("alloc"))
	return rv
}

func (tc _ToolbarItemClass) New() ToolbarItem {
	rv := objc.CallMethod[ToolbarItem](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewToolbarItem() ToolbarItem {
	return ToolbarItemClass.New()
}

func (t_ ToolbarItem) Init() ToolbarItem {
	rv := objc.CallMethod[ToolbarItem](t_, objc.GetSelector("init"))
	return rv
}

func (t_ ToolbarItem) Validate() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("validate"))
}

func (t_ ToolbarItem) ItemIdentifier() ToolbarItemIdentifier {
	rv := objc.CallMethod[ToolbarItemIdentifier](t_, objc.GetSelector("itemIdentifier"))
	return rv
}

func (t_ ToolbarItem) PossibleLabels() foundation.Set {
	rv := objc.CallMethod[foundation.Set](t_, objc.GetSelector("possibleLabels"))
	return rv
}

func (t_ ToolbarItem) SetPossibleLabels(value foundation.ISet) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setPossibleLabels:"), objc.ExtractPtr(value))
}

func (t_ ToolbarItem) Label() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("label"))
	return rv
}

func (t_ ToolbarItem) SetLabel(value string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setLabel:"), value)
}

func (t_ ToolbarItem) PaletteLabel() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("paletteLabel"))
	return rv
}

func (t_ ToolbarItem) SetPaletteLabel(value string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setPaletteLabel:"), value)
}

func (t_ ToolbarItem) Title() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("title"))
	return rv
}

func (t_ ToolbarItem) SetTitle(value string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTitle:"), value)
}

func (t_ ToolbarItem) ToolTip() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("toolTip"))
	return rv
}

func (t_ ToolbarItem) SetToolTip(value string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setToolTip:"), value)
}

func (t_ ToolbarItem) Image() Image {
	rv := objc.CallMethod[Image](t_, objc.GetSelector("image"))
	return rv
}

func (t_ ToolbarItem) SetImage(value IImage) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setImage:"), objc.ExtractPtr(value))
}

func (t_ ToolbarItem) View() View {
	rv := objc.CallMethod[View](t_, objc.GetSelector("view"))
	return rv
}

func (t_ ToolbarItem) SetView(value IView) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setView:"), objc.ExtractPtr(value))
}

func (t_ ToolbarItem) Target() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.GetSelector("target"))
	return rv
}

func (t_ ToolbarItem) SetTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTarget:"), objc.ExtractPtr(value))
}

func (t_ ToolbarItem) Action() objc.Selector {
	rv := objc.CallMethod[objc.Selector](t_, objc.GetSelector("action"))
	return rv
}

func (t_ ToolbarItem) SetAction(value objc.Selector) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAction:"), value)
}

func (t_ ToolbarItem) MenuFormRepresentation() MenuItem {
	rv := objc.CallMethod[MenuItem](t_, objc.GetSelector("menuFormRepresentation"))
	return rv
}

func (t_ ToolbarItem) SetMenuFormRepresentation(value IMenuItem) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setMenuFormRepresentation:"), objc.ExtractPtr(value))
}

func (t_ ToolbarItem) IsVisible() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isVisible"))
	return rv
}

func (t_ ToolbarItem) IsBordered() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isBordered"))
	return rv
}

func (t_ ToolbarItem) SetBordered(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setBordered:"), value)
}

func (t_ ToolbarItem) IsNavigational() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isNavigational"))
	return rv
}

func (t_ ToolbarItem) SetNavigational(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setNavigational:"), value)
}

func (t_ ToolbarItem) IsEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isEnabled"))
	return rv
}

func (t_ ToolbarItem) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setEnabled:"), value)
}

func (t_ ToolbarItem) AllowsDuplicatesInToolbar() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsDuplicatesInToolbar"))
	return rv
}

func (t_ ToolbarItem) VisibilityPriority() ToolbarItemVisibilityPriority {
	rv := objc.CallMethod[ToolbarItemVisibilityPriority](t_, objc.GetSelector("visibilityPriority"))
	return rv
}

func (t_ ToolbarItem) SetVisibilityPriority(value ToolbarItemVisibilityPriority) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setVisibilityPriority:"), value)
}

func (t_ ToolbarItem) Tag() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("tag"))
	return rv
}

func (t_ ToolbarItem) SetTag(value int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTag:"), value)
}

func (t_ ToolbarItem) Toolbar() Toolbar {
	rv := objc.CallMethod[Toolbar](t_, objc.GetSelector("toolbar"))
	return rv
}

func (t_ ToolbarItem) Autovalidates() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("autovalidates"))
	return rv
}

func (t_ ToolbarItem) SetAutovalidates(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAutovalidates:"), value)
}

// deprecated
func (t_ ToolbarItem) MinSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.GetSelector("minSize"))
	return rv
}

// deprecated
func (t_ ToolbarItem) SetMinSize(value foundation.Size) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setMinSize:"), value)
}

// deprecated
func (t_ ToolbarItem) MaxSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.GetSelector("maxSize"))
	return rv
}

// deprecated
func (t_ ToolbarItem) SetMaxSize(value foundation.Size) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setMaxSize:"), value)
}
