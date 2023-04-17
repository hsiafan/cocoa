// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[ToolbarItem](t_, "initWithItemIdentifier:", itemIdentifier)
	return rv
}

func (tc _ToolbarItemClass) Alloc() ToolbarItem {
	rv := ffi.CallMethod[ToolbarItem](tc, "alloc")
	return rv
}

func (tc _ToolbarItemClass) New() ToolbarItem {
	rv := ffi.CallMethod[ToolbarItem](tc, "new")
	rv.Autorelease()
	return rv
}

func NewToolbarItem() ToolbarItem {
	return ToolbarItemClass.New()
}

func (t_ ToolbarItem) Init() ToolbarItem {
	rv := ffi.CallMethod[ToolbarItem](t_, "init")
	return rv
}

func (t_ ToolbarItem) Validate() {
	ffi.CallMethod[ffi.Void](t_, "validate")
}

func (t_ ToolbarItem) ItemIdentifier() ToolbarItemIdentifier {
	rv := ffi.CallMethod[ToolbarItemIdentifier](t_, "itemIdentifier")
	return rv
}

func (t_ ToolbarItem) PossibleLabels() foundation.Set {
	rv := ffi.CallMethod[foundation.Set](t_, "possibleLabels")
	return rv
}

func (t_ ToolbarItem) SetPossibleLabels(value foundation.ISet) {
	ffi.CallMethod[ffi.Void](t_, "setPossibleLabels:", value)
}

func (t_ ToolbarItem) Label() string {
	rv := ffi.CallMethod[string](t_, "label")
	return rv
}

func (t_ ToolbarItem) SetLabel(value string) {
	ffi.CallMethod[ffi.Void](t_, "setLabel:", value)
}

func (t_ ToolbarItem) PaletteLabel() string {
	rv := ffi.CallMethod[string](t_, "paletteLabel")
	return rv
}

func (t_ ToolbarItem) SetPaletteLabel(value string) {
	ffi.CallMethod[ffi.Void](t_, "setPaletteLabel:", value)
}

func (t_ ToolbarItem) Title() string {
	rv := ffi.CallMethod[string](t_, "title")
	return rv
}

func (t_ ToolbarItem) SetTitle(value string) {
	ffi.CallMethod[ffi.Void](t_, "setTitle:", value)
}

func (t_ ToolbarItem) ToolTip() string {
	rv := ffi.CallMethod[string](t_, "toolTip")
	return rv
}

func (t_ ToolbarItem) SetToolTip(value string) {
	ffi.CallMethod[ffi.Void](t_, "setToolTip:", value)
}

func (t_ ToolbarItem) Image() Image {
	rv := ffi.CallMethod[Image](t_, "image")
	return rv
}

func (t_ ToolbarItem) SetImage(value IImage) {
	ffi.CallMethod[ffi.Void](t_, "setImage:", value)
}

func (t_ ToolbarItem) View() View {
	rv := ffi.CallMethod[View](t_, "view")
	return rv
}

func (t_ ToolbarItem) SetView(value IView) {
	ffi.CallMethod[ffi.Void](t_, "setView:", value)
}

func (t_ ToolbarItem) Target() objc.Object {
	rv := ffi.CallMethod[objc.Object](t_, "target")
	return rv
}

func (t_ ToolbarItem) SetTarget(value objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "setTarget:", value)
}

func (t_ ToolbarItem) Action() objc.Selector {
	rv := ffi.CallMethod[objc.Selector](t_, "action")
	return rv
}

func (t_ ToolbarItem) SetAction(value objc.Selector) {
	ffi.CallMethod[ffi.Void](t_, "setAction:", value)
}

func (t_ ToolbarItem) MenuFormRepresentation() MenuItem {
	rv := ffi.CallMethod[MenuItem](t_, "menuFormRepresentation")
	return rv
}

func (t_ ToolbarItem) SetMenuFormRepresentation(value IMenuItem) {
	ffi.CallMethod[ffi.Void](t_, "setMenuFormRepresentation:", value)
}

func (t_ ToolbarItem) IsVisible() bool {
	rv := ffi.CallMethod[bool](t_, "isVisible")
	return rv
}

func (t_ ToolbarItem) IsBordered() bool {
	rv := ffi.CallMethod[bool](t_, "isBordered")
	return rv
}

func (t_ ToolbarItem) SetBordered(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setBordered:", value)
}

func (t_ ToolbarItem) IsNavigational() bool {
	rv := ffi.CallMethod[bool](t_, "isNavigational")
	return rv
}

func (t_ ToolbarItem) SetNavigational(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setNavigational:", value)
}

func (t_ ToolbarItem) IsEnabled() bool {
	rv := ffi.CallMethod[bool](t_, "isEnabled")
	return rv
}

func (t_ ToolbarItem) SetEnabled(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setEnabled:", value)
}

func (t_ ToolbarItem) AllowsDuplicatesInToolbar() bool {
	rv := ffi.CallMethod[bool](t_, "allowsDuplicatesInToolbar")
	return rv
}

func (t_ ToolbarItem) VisibilityPriority() ToolbarItemVisibilityPriority {
	rv := ffi.CallMethod[ToolbarItemVisibilityPriority](t_, "visibilityPriority")
	return rv
}

func (t_ ToolbarItem) SetVisibilityPriority(value ToolbarItemVisibilityPriority) {
	ffi.CallMethod[ffi.Void](t_, "setVisibilityPriority:", value)
}

func (t_ ToolbarItem) Tag() int {
	rv := ffi.CallMethod[int](t_, "tag")
	return rv
}

func (t_ ToolbarItem) SetTag(value int) {
	ffi.CallMethod[ffi.Void](t_, "setTag:", value)
}

func (t_ ToolbarItem) Toolbar() Toolbar {
	rv := ffi.CallMethod[Toolbar](t_, "toolbar")
	return rv
}

func (t_ ToolbarItem) Autovalidates() bool {
	rv := ffi.CallMethod[bool](t_, "autovalidates")
	return rv
}

func (t_ ToolbarItem) SetAutovalidates(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAutovalidates:", value)
}

// deprecated
func (t_ ToolbarItem) MinSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](t_, "minSize")
	return rv
}

// deprecated
func (t_ ToolbarItem) SetMinSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](t_, "setMinSize:", value)
}

// deprecated
func (t_ ToolbarItem) MaxSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](t_, "maxSize")
	return rv
}

// deprecated
func (t_ ToolbarItem) SetMaxSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](t_, "setMaxSize:", value)
}
