// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[TouchBar](t_, "init")
	return rv
}

func (tc _TouchBarClass) Alloc() TouchBar {
	rv := objc.CallMethod[TouchBar](tc, "alloc")
	return rv
}

func (tc _TouchBarClass) New() TouchBar {
	rv := objc.CallMethod[TouchBar](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTouchBar() TouchBar {
	return TouchBarClass.New()
}

func (t_ TouchBar) ItemForIdentifier(identifier TouchBarItemIdentifier) TouchBarItem {
	rv := objc.CallMethod[TouchBarItem](t_, "itemForIdentifier:", identifier)
	return rv
}

func (t_ TouchBar) Delegate() TouchBarDelegateWrapper {
	rv := objc.CallMethod[TouchBarDelegateWrapper](t_, "delegate")
	return rv
}

func (t_ TouchBar) SetDelegate(value TouchBarDelegate) {
	po := objc.CreateProtocol("NSTouchBarDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](t_, "setDelegate:", po)
}

func (t_ TouchBar) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, "setDelegate:", value)
}

func (t_ TouchBar) TemplateItems() foundation.Set {
	rv := objc.CallMethod[foundation.Set](t_, "templateItems")
	return rv
}

func (t_ TouchBar) SetTemplateItems(value foundation.ISet) {
	objc.CallMethod[objc.Void](t_, "setTemplateItems:", value)
}

func (t_ TouchBar) DefaultItemIdentifiers() []TouchBarItemIdentifier {
	rv := objc.CallMethod[[]TouchBarItemIdentifier](t_, "defaultItemIdentifiers")
	return rv
}

func (t_ TouchBar) SetDefaultItemIdentifiers(value []TouchBarItemIdentifier) {
	objc.CallMethod[objc.Void](t_, "setDefaultItemIdentifiers:", value)
}

func (t_ TouchBar) PrincipalItemIdentifier() TouchBarItemIdentifier {
	rv := objc.CallMethod[TouchBarItemIdentifier](t_, "principalItemIdentifier")
	return rv
}

func (t_ TouchBar) SetPrincipalItemIdentifier(value TouchBarItemIdentifier) {
	objc.CallMethod[objc.Void](t_, "setPrincipalItemIdentifier:", value)
}

func (t_ TouchBar) EscapeKeyReplacementItemIdentifier() TouchBarItemIdentifier {
	rv := objc.CallMethod[TouchBarItemIdentifier](t_, "escapeKeyReplacementItemIdentifier")
	return rv
}

func (t_ TouchBar) SetEscapeKeyReplacementItemIdentifier(value TouchBarItemIdentifier) {
	objc.CallMethod[objc.Void](t_, "setEscapeKeyReplacementItemIdentifier:", value)
}

func (t_ TouchBar) IsVisible() bool {
	rv := objc.CallMethod[bool](t_, "isVisible")
	return rv
}

func (t_ TouchBar) ItemIdentifiers() []TouchBarItemIdentifier {
	rv := objc.CallMethod[[]TouchBarItemIdentifier](t_, "itemIdentifiers")
	return rv
}

func (t_ TouchBar) CustomizationIdentifier() TouchBarCustomizationIdentifier {
	rv := objc.CallMethod[TouchBarCustomizationIdentifier](t_, "customizationIdentifier")
	return rv
}

func (t_ TouchBar) SetCustomizationIdentifier(value TouchBarCustomizationIdentifier) {
	objc.CallMethod[objc.Void](t_, "setCustomizationIdentifier:", value)
}

func (t_ TouchBar) CustomizationAllowedItemIdentifiers() []TouchBarItemIdentifier {
	rv := objc.CallMethod[[]TouchBarItemIdentifier](t_, "customizationAllowedItemIdentifiers")
	return rv
}

func (t_ TouchBar) SetCustomizationAllowedItemIdentifiers(value []TouchBarItemIdentifier) {
	objc.CallMethod[objc.Void](t_, "setCustomizationAllowedItemIdentifiers:", value)
}

func (t_ TouchBar) CustomizationRequiredItemIdentifiers() []TouchBarItemIdentifier {
	rv := objc.CallMethod[[]TouchBarItemIdentifier](t_, "customizationRequiredItemIdentifiers")
	return rv
}

func (t_ TouchBar) SetCustomizationRequiredItemIdentifiers(value []TouchBarItemIdentifier) {
	objc.CallMethod[objc.Void](t_, "setCustomizationRequiredItemIdentifiers:", value)
}

func (tc _TouchBarClass) IsAutomaticCustomizeTouchBarMenuItemEnabled() bool {
	rv := objc.CallMethod[bool](tc, "isAutomaticCustomizeTouchBarMenuItemEnabled")
	return rv
}

func (tc _TouchBarClass) SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	objc.CallMethod[objc.Void](tc, "setAutomaticCustomizeTouchBarMenuItemEnabled:", value)
}
