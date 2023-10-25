// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TouchBarClass = _TouchBarClass{objc.GetClass("NSTouchBar")}

type _TouchBarClass struct {
	objc.Class
}

type ITouchBar interface {
	objc.IObject
	ItemForIdentifier(identifier TouchBarItemIdentifier) TouchBarItem
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
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
	rv := objc.CallMethod[TouchBar](t_, objc.SEL("init"))
	return rv
}

func (tc _TouchBarClass) Alloc() TouchBar {
	rv := objc.CallMethod[TouchBar](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TouchBarClass) New() TouchBar {
	rv := objc.CallMethod[TouchBar](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTouchBar() TouchBar {
	return TouchBarClass.New()
}

func (t_ TouchBar) ItemForIdentifier(identifier TouchBarItemIdentifier) TouchBarItem {
	rv := objc.CallMethod[TouchBarItem](t_, objc.SEL("itemForIdentifier:"), identifier)
	return rv
}

// weak property
func (t_ TouchBar) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.SEL("delegate"))
	return rv
}

// weak property
func (t_ TouchBar) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}

func (t_ TouchBar) TemplateItems() foundation.Set {
	rv := objc.CallMethod[foundation.Set](t_, objc.SEL("templateItems"))
	return rv
}

func (t_ TouchBar) SetTemplateItems(value foundation.ISet) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTemplateItems:"), objc.ExtractPtr(value))
}

func (t_ TouchBar) DefaultItemIdentifiers() []TouchBarItemIdentifier {
	rv := objc.CallMethod[[]TouchBarItemIdentifier](t_, objc.SEL("defaultItemIdentifiers"))
	return rv
}

func (t_ TouchBar) SetDefaultItemIdentifiers(value []TouchBarItemIdentifier) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDefaultItemIdentifiers:"), value)
}

func (t_ TouchBar) PrincipalItemIdentifier() TouchBarItemIdentifier {
	rv := objc.CallMethod[TouchBarItemIdentifier](t_, objc.SEL("principalItemIdentifier"))
	return rv
}

func (t_ TouchBar) SetPrincipalItemIdentifier(value TouchBarItemIdentifier) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setPrincipalItemIdentifier:"), value)
}

func (t_ TouchBar) EscapeKeyReplacementItemIdentifier() TouchBarItemIdentifier {
	rv := objc.CallMethod[TouchBarItemIdentifier](t_, objc.SEL("escapeKeyReplacementItemIdentifier"))
	return rv
}

func (t_ TouchBar) SetEscapeKeyReplacementItemIdentifier(value TouchBarItemIdentifier) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setEscapeKeyReplacementItemIdentifier:"), value)
}

func (t_ TouchBar) IsVisible() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isVisible"))
	return rv
}

func (t_ TouchBar) ItemIdentifiers() []TouchBarItemIdentifier {
	rv := objc.CallMethod[[]TouchBarItemIdentifier](t_, objc.SEL("itemIdentifiers"))
	return rv
}

func (t_ TouchBar) CustomizationIdentifier() TouchBarCustomizationIdentifier {
	rv := objc.CallMethod[TouchBarCustomizationIdentifier](t_, objc.SEL("customizationIdentifier"))
	return rv
}

func (t_ TouchBar) SetCustomizationIdentifier(value TouchBarCustomizationIdentifier) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setCustomizationIdentifier:"), value)
}

func (t_ TouchBar) CustomizationAllowedItemIdentifiers() []TouchBarItemIdentifier {
	rv := objc.CallMethod[[]TouchBarItemIdentifier](t_, objc.SEL("customizationAllowedItemIdentifiers"))
	return rv
}

func (t_ TouchBar) SetCustomizationAllowedItemIdentifiers(value []TouchBarItemIdentifier) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setCustomizationAllowedItemIdentifiers:"), value)
}

func (t_ TouchBar) CustomizationRequiredItemIdentifiers() []TouchBarItemIdentifier {
	rv := objc.CallMethod[[]TouchBarItemIdentifier](t_, objc.SEL("customizationRequiredItemIdentifiers"))
	return rv
}

func (t_ TouchBar) SetCustomizationRequiredItemIdentifiers(value []TouchBarItemIdentifier) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setCustomizationRequiredItemIdentifiers:"), value)
}

func (tc _TouchBarClass) IsAutomaticCustomizeTouchBarMenuItemEnabled() bool {
	rv := objc.CallMethod[bool](tc, objc.SEL("isAutomaticCustomizeTouchBarMenuItemEnabled"))
	return rv
}

func (tc _TouchBarClass) SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	objc.CallMethod[objc.Void](tc, objc.SEL("setAutomaticCustomizeTouchBarMenuItemEnabled:"), value)
}
