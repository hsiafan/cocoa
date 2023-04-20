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
	SetDelegate0(value objc.IObject)
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
	po := ffi.CreateProtocol("NSTouchBarDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](t_, "setDelegate:", po)
}

func (t_ TouchBar) SetDelegate0(value objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "setDelegate:", value)
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

func (tc _TouchBarClass) IsAutomaticCustomizeTouchBarMenuItemEnabled() bool {
	rv := ffi.CallMethod[bool](tc, "isAutomaticCustomizeTouchBarMenuItemEnabled")
	return rv
}

func (tc _TouchBarClass) SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	ffi.CallMethod[ffi.Void](tc, "setAutomaticCustomizeTouchBarMenuItemEnabled:", value)
}
