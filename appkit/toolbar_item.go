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
	rv := objc.CallMethod[ToolbarItem](t_, objc.SEL("initWithItemIdentifier:"), itemIdentifier)
	return rv
}

func (tc _ToolbarItemClass) Alloc() ToolbarItem {
	rv := objc.CallMethod[ToolbarItem](tc, objc.SEL("alloc"))
	return rv
}

func (tc _ToolbarItemClass) New() ToolbarItem {
	rv := objc.CallMethod[ToolbarItem](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewToolbarItem() ToolbarItem {
	return ToolbarItemClass.New()
}

func (t_ ToolbarItem) Init() ToolbarItem {
	rv := objc.CallMethod[ToolbarItem](t_, objc.SEL("init"))
	return rv
}

func (t_ ToolbarItem) Validate() {
	objc.CallMethod[objc.Void](t_, objc.SEL("validate"))
}

func (t_ ToolbarItem) ItemIdentifier() ToolbarItemIdentifier {
	rv := objc.CallMethod[ToolbarItemIdentifier](t_, objc.SEL("itemIdentifier"))
	return rv
}

func (t_ ToolbarItem) PossibleLabels() foundation.Set {
	rv := objc.CallMethod[foundation.Set](t_, objc.SEL("possibleLabels"))
	return rv
}

func (t_ ToolbarItem) SetPossibleLabels(value foundation.ISet) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setPossibleLabels:"), objc.ExtractPtr(value))
}

func (t_ ToolbarItem) Label() string {
	rv := objc.CallMethod[string](t_, objc.SEL("label"))
	return rv
}

func (t_ ToolbarItem) SetLabel(value string) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setLabel:"), value)
}

func (t_ ToolbarItem) PaletteLabel() string {
	rv := objc.CallMethod[string](t_, objc.SEL("paletteLabel"))
	return rv
}

func (t_ ToolbarItem) SetPaletteLabel(value string) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setPaletteLabel:"), value)
}

func (t_ ToolbarItem) Title() string {
	rv := objc.CallMethod[string](t_, objc.SEL("title"))
	return rv
}

func (t_ ToolbarItem) SetTitle(value string) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTitle:"), value)
}

func (t_ ToolbarItem) ToolTip() string {
	rv := objc.CallMethod[string](t_, objc.SEL("toolTip"))
	return rv
}

func (t_ ToolbarItem) SetToolTip(value string) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setToolTip:"), value)
}

func (t_ ToolbarItem) Image() Image {
	rv := objc.CallMethod[Image](t_, objc.SEL("image"))
	return rv
}

func (t_ ToolbarItem) SetImage(value IImage) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setImage:"), objc.ExtractPtr(value))
}

func (t_ ToolbarItem) View() View {
	rv := objc.CallMethod[View](t_, objc.SEL("view"))
	return rv
}

func (t_ ToolbarItem) SetView(value IView) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setView:"), objc.ExtractPtr(value))
}

// weak property
func (t_ ToolbarItem) Target() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.SEL("target"))
	return rv
}

// weak property
func (t_ ToolbarItem) SetTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTarget:"), objc.ExtractPtr(value))
}

func (t_ ToolbarItem) Action() objc.Selector {
	rv := objc.CallMethod[objc.Selector](t_, objc.SEL("action"))
	return rv
}

func (t_ ToolbarItem) SetAction(value objc.Selector) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAction:"), value)
}

func (t_ ToolbarItem) MenuFormRepresentation() MenuItem {
	rv := objc.CallMethod[MenuItem](t_, objc.SEL("menuFormRepresentation"))
	return rv
}

func (t_ ToolbarItem) SetMenuFormRepresentation(value IMenuItem) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setMenuFormRepresentation:"), objc.ExtractPtr(value))
}

func (t_ ToolbarItem) IsVisible() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isVisible"))
	return rv
}

func (t_ ToolbarItem) IsBordered() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isBordered"))
	return rv
}

func (t_ ToolbarItem) SetBordered(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setBordered:"), value)
}

func (t_ ToolbarItem) IsNavigational() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isNavigational"))
	return rv
}

func (t_ ToolbarItem) SetNavigational(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setNavigational:"), value)
}

func (t_ ToolbarItem) IsEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isEnabled"))
	return rv
}

func (t_ ToolbarItem) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setEnabled:"), value)
}

func (t_ ToolbarItem) AllowsDuplicatesInToolbar() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("allowsDuplicatesInToolbar"))
	return rv
}

func (t_ ToolbarItem) VisibilityPriority() ToolbarItemVisibilityPriority {
	rv := objc.CallMethod[ToolbarItemVisibilityPriority](t_, objc.SEL("visibilityPriority"))
	return rv
}

func (t_ ToolbarItem) SetVisibilityPriority(value ToolbarItemVisibilityPriority) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setVisibilityPriority:"), value)
}

func (t_ ToolbarItem) Tag() int {
	rv := objc.CallMethod[int](t_, objc.SEL("tag"))
	return rv
}

func (t_ ToolbarItem) SetTag(value int) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTag:"), value)
}

// weak property
func (t_ ToolbarItem) Toolbar() Toolbar {
	rv := objc.CallMethod[Toolbar](t_, objc.SEL("toolbar"))
	return rv
}

func (t_ ToolbarItem) Autovalidates() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("autovalidates"))
	return rv
}

func (t_ ToolbarItem) SetAutovalidates(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAutovalidates:"), value)
}

// deprecated
func (t_ ToolbarItem) MinSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.SEL("minSize"))
	return rv
}

// deprecated
func (t_ ToolbarItem) SetMinSize(value foundation.Size) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setMinSize:"), value)
}

// deprecated
func (t_ ToolbarItem) MaxSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.SEL("maxSize"))
	return rv
}

// deprecated
func (t_ ToolbarItem) SetMaxSize(value foundation.Size) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setMaxSize:"), value)
}
