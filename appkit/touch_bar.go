// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var TouchBarClass = _TouchBarClass{objc.GetClass("NSTouchBar")}

type _TouchBarClass struct {
	objc.Class
}

type ITouchBar interface {
	objc.IObject
	ItemForIdentifier(identifier TouchBarItemIdentifier) TouchBarItem
	Delegate() TouchBarDelegateWrapper
	SetDelegate(value TouchBarDelegate)
	TemplateItems() foundation.Set
	SetTemplateItems(value foundation.ISet)
	DefaultItemIdentifiers() []TouchBarItemIdentifier
	SetDefaultItemIdentifiers(value []TouchBarItemIdentifier)
	PrincipalItemIdentifier() TouchBarItemIdentifier
	SetPrincipalItemIdentifier(value TouchBarItemIdentifier)
	EscapeKeyReplacementItemIdentifier() TouchBarItemIdentifier
	SetEscapeKeyReplacementItemIdentifier(value TouchBarItemIdentifier)
	IsVisible() bool
	ItemIdentifiers() []TouchBarItemIdentifier
	CustomizationIdentifier() TouchBarCustomizationIdentifier
	SetCustomizationIdentifier(value TouchBarCustomizationIdentifier)
	CustomizationAllowedItemIdentifiers() []TouchBarItemIdentifier
	SetCustomizationAllowedItemIdentifiers(value []TouchBarItemIdentifier)
	CustomizationRequiredItemIdentifiers() []TouchBarItemIdentifier
	SetCustomizationRequiredItemIdentifiers(value []TouchBarItemIdentifier)
}

type TouchBar struct {
	objc.Object
}

func MakeTouchBar(ptr unsafe.Pointer) TouchBar {
	return TouchBar{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TouchBar) Init() TouchBar {
	rv := ffi.CallMethod[TouchBar](t_, "init")
	return rv
}

func (tc _TouchBarClass) Alloc() TouchBar {
	rv := ffi.CallMethod[TouchBar](tc, "alloc")
	return rv
}

func (tc _TouchBarClass) New() TouchBar {
	rv := ffi.CallMethod[TouchBar](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTouchBar() TouchBar {
	return TouchBarClass.New()
}

func (t_ TouchBar) ItemForIdentifier(identifier TouchBarItemIdentifier) TouchBarItem {
	rv := ffi.CallMethod[TouchBarItem](t_, "itemForIdentifier:", identifier)
	return rv
}

func (t_ TouchBar) Delegate() TouchBarDelegateWrapper {
	rv := ffi.CallMethod[TouchBarDelegateWrapper](t_, "delegate")
	return rv
}

func (t_ TouchBar) SetDelegate(value TouchBarDelegate) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](t_, "setDelegate:", po)
}

func (t_ TouchBar) TemplateItems() foundation.Set {
	rv := ffi.CallMethod[foundation.Set](t_, "templateItems")
	return rv
}

func (t_ TouchBar) SetTemplateItems(value foundation.ISet) {
	ffi.CallMethod[ffi.Void](t_, "setTemplateItems:", value)
}

func (t_ TouchBar) DefaultItemIdentifiers() []TouchBarItemIdentifier {
	rv := ffi.CallMethod[[]TouchBarItemIdentifier](t_, "defaultItemIdentifiers")
	return rv
}

func (t_ TouchBar) SetDefaultItemIdentifiers(value []TouchBarItemIdentifier) {
	ffi.CallMethod[ffi.Void](t_, "setDefaultItemIdentifiers:", value)
}

func (t_ TouchBar) PrincipalItemIdentifier() TouchBarItemIdentifier {
	rv := ffi.CallMethod[TouchBarItemIdentifier](t_, "principalItemIdentifier")
	return rv
}

func (t_ TouchBar) SetPrincipalItemIdentifier(value TouchBarItemIdentifier) {
	ffi.CallMethod[ffi.Void](t_, "setPrincipalItemIdentifier:", value)
}

func (t_ TouchBar) EscapeKeyReplacementItemIdentifier() TouchBarItemIdentifier {
	rv := ffi.CallMethod[TouchBarItemIdentifier](t_, "escapeKeyReplacementItemIdentifier")
	return rv
}

func (t_ TouchBar) SetEscapeKeyReplacementItemIdentifier(value TouchBarItemIdentifier) {
	ffi.CallMethod[ffi.Void](t_, "setEscapeKeyReplacementItemIdentifier:", value)
}

func (t_ TouchBar) IsVisible() bool {
	rv := ffi.CallMethod[bool](t_, "isVisible")
	return rv
}

func (t_ TouchBar) ItemIdentifiers() []TouchBarItemIdentifier {
	rv := ffi.CallMethod[[]TouchBarItemIdentifier](t_, "itemIdentifiers")
	return rv
}

func (t_ TouchBar) CustomizationIdentifier() TouchBarCustomizationIdentifier {
	rv := ffi.CallMethod[TouchBarCustomizationIdentifier](t_, "customizationIdentifier")
	return rv
}

func (t_ TouchBar) SetCustomizationIdentifier(value TouchBarCustomizationIdentifier) {
	ffi.CallMethod[ffi.Void](t_, "setCustomizationIdentifier:", value)
}

func (t_ TouchBar) CustomizationAllowedItemIdentifiers() []TouchBarItemIdentifier {
	rv := ffi.CallMethod[[]TouchBarItemIdentifier](t_, "customizationAllowedItemIdentifiers")
	return rv
}

func (t_ TouchBar) SetCustomizationAllowedItemIdentifiers(value []TouchBarItemIdentifier) {
	ffi.CallMethod[ffi.Void](t_, "setCustomizationAllowedItemIdentifiers:", value)
}

func (t_ TouchBar) CustomizationRequiredItemIdentifiers() []TouchBarItemIdentifier {
	rv := ffi.CallMethod[[]TouchBarItemIdentifier](t_, "customizationRequiredItemIdentifiers")
	return rv
}

func (t_ TouchBar) SetCustomizationRequiredItemIdentifiers(value []TouchBarItemIdentifier) {
	ffi.CallMethod[ffi.Void](t_, "setCustomizationRequiredItemIdentifiers:", value)
}

func (tc _TouchBarClass) AutomaticCustomizeTouchBarMenuItemEnabled() bool {
	rv := ffi.CallMethod[bool](tc, "automaticCustomizeTouchBarMenuItemEnabled")
	return rv
}

func (tc _TouchBarClass) SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	ffi.CallMethod[ffi.Void](tc, "setAutomaticCustomizeTouchBarMenuItemEnabled:", value)
}

type TouchBarDelegate interface {
	ImplementsTouchBar_MakeItemForIdentifier() bool
	// optional
	TouchBar_MakeItemForIdentifier(touchBar TouchBar, identifier TouchBarItemIdentifier) ITouchBarItem
}

type TouchBarDelegateImpl struct {
	_TouchBar_MakeItemForIdentifier func(touchBar TouchBar, identifier TouchBarItemIdentifier) ITouchBarItem
}

func (di *TouchBarDelegateImpl) ImplementsTouchBar_MakeItemForIdentifier() bool {
	return di._TouchBar_MakeItemForIdentifier != nil
}

func (di *TouchBarDelegateImpl) SetTouchBar_MakeItemForIdentifier(f func(touchBar TouchBar, identifier TouchBarItemIdentifier) ITouchBarItem) {
	di._TouchBar_MakeItemForIdentifier = f
}

func (di *TouchBarDelegateImpl) TouchBar_MakeItemForIdentifier(touchBar TouchBar, identifier TouchBarItemIdentifier) ITouchBarItem {
	return di._TouchBar_MakeItemForIdentifier(touchBar, identifier)
}

type TouchBarDelegateWrapper struct {
	objc.Object
}

func (t_ *TouchBarDelegateWrapper) ImplementsTouchBar_MakeItemForIdentifier() bool {
	return t_.RespondsToSelector(objc.GetSelector("touchBar:makeItemForIdentifier:"))
}

func (t_ TouchBarDelegateWrapper) TouchBar_MakeItemForIdentifier(touchBar ITouchBar, identifier TouchBarItemIdentifier) TouchBarItem {
	rv := ffi.CallMethod[TouchBarItem](t_, "touchBar:makeItemForIdentifier:", touchBar, identifier)
	return rv
}

var TouchBarItemClass = _TouchBarItemClass{objc.GetClass("NSTouchBarItem")}

type _TouchBarItemClass struct {
	objc.Class
}

type ITouchBarItem interface {
	objc.IObject
	Identifier() TouchBarItemIdentifier
	VisibilityPriority() TouchBarItemPriority
	SetVisibilityPriority(value TouchBarItemPriority)
	IsVisible() bool
	CustomizationLabel() string
	View() View
}

type TouchBarItem struct {
	objc.Object
}

func MakeTouchBarItem(ptr unsafe.Pointer) TouchBarItem {
	return TouchBarItem{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) TouchBarItem {
	rv := ffi.CallMethod[TouchBarItem](t_, "initWithIdentifier:", identifier)
	return rv
}

func (tc _TouchBarItemClass) Alloc() TouchBarItem {
	rv := ffi.CallMethod[TouchBarItem](tc, "alloc")
	return rv
}

func (t_ TouchBarItem) Init() TouchBarItem {
	rv := ffi.CallMethod[TouchBarItem](t_, "init")
	return rv
}

func (tc _TouchBarItemClass) New() TouchBarItem {
	rv := ffi.CallMethod[TouchBarItem](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTouchBarItem() TouchBarItem {
	return TouchBarItemClass.New()
}

func (t_ TouchBarItem) Identifier() TouchBarItemIdentifier {
	rv := ffi.CallMethod[TouchBarItemIdentifier](t_, "identifier")
	return rv
}

func (t_ TouchBarItem) VisibilityPriority() TouchBarItemPriority {
	rv := ffi.CallMethod[TouchBarItemPriority](t_, "visibilityPriority")
	return rv
}

func (t_ TouchBarItem) SetVisibilityPriority(value TouchBarItemPriority) {
	ffi.CallMethod[ffi.Void](t_, "setVisibilityPriority:", value)
}

func (t_ TouchBarItem) IsVisible() bool {
	rv := ffi.CallMethod[bool](t_, "isVisible")
	return rv
}

func (t_ TouchBarItem) CustomizationLabel() string {
	rv := ffi.CallMethod[string](t_, "customizationLabel")
	return rv
}

func (t_ TouchBarItem) View() View {
	rv := ffi.CallMethod[View](t_, "view")
	return rv
}
