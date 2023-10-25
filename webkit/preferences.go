// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var PreferencesClass = _PreferencesClass{objc.GetClass("WKPreferences")}

type _PreferencesClass struct {
	objc.Class
}

type IPreferences interface {
	objc.IObject
	MinimumFontSize() float64
	SetMinimumFontSize(value float64)
	TabFocusesLinks() bool
	SetTabFocusesLinks(value bool)
	JavaScriptCanOpenWindowsAutomatically() bool
	SetJavaScriptCanOpenWindowsAutomatically(value bool)
	IsFraudulentWebsiteWarningEnabled() bool
	SetFraudulentWebsiteWarningEnabled(value bool)
	// deprecated
	JavaEnabled() bool
	// deprecated
	SetJavaEnabled(value bool)
	// deprecated
	JavaScriptEnabled() bool
	// deprecated
	SetJavaScriptEnabled(value bool)
	// deprecated
	PlugInsEnabled() bool
	// deprecated
	SetPlugInsEnabled(value bool)
	IsElementFullscreenEnabled() bool
	SetElementFullscreenEnabled(value bool)
	IsSiteSpecificQuirksModeEnabled() bool
	SetSiteSpecificQuirksModeEnabled(value bool)
	IsTextInteractionEnabled() bool
	SetTextInteractionEnabled(value bool)
}

type Preferences struct {
	objc.Object
}

func MakePreferences(ptr unsafe.Pointer) Preferences {
	return Preferences{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PreferencesClass) Alloc() Preferences {
	rv := objc.CallMethod[Preferences](pc, objc.SEL("alloc"))
	return rv
}

func (pc _PreferencesClass) New() Preferences {
	rv := objc.CallMethod[Preferences](pc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewPreferences() Preferences {
	return PreferencesClass.New()
}

func (p_ Preferences) Init() Preferences {
	rv := objc.CallMethod[Preferences](p_, objc.SEL("init"))
	return rv
}

func (p_ Preferences) MinimumFontSize() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("minimumFontSize"))
	return rv
}

func (p_ Preferences) SetMinimumFontSize(value float64) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setMinimumFontSize:"), value)
}

func (p_ Preferences) TabFocusesLinks() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("tabFocusesLinks"))
	return rv
}

func (p_ Preferences) SetTabFocusesLinks(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setTabFocusesLinks:"), value)
}

func (p_ Preferences) JavaScriptCanOpenWindowsAutomatically() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("javaScriptCanOpenWindowsAutomatically"))
	return rv
}

func (p_ Preferences) SetJavaScriptCanOpenWindowsAutomatically(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setJavaScriptCanOpenWindowsAutomatically:"), value)
}

func (p_ Preferences) IsFraudulentWebsiteWarningEnabled() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isFraudulentWebsiteWarningEnabled"))
	return rv
}

func (p_ Preferences) SetFraudulentWebsiteWarningEnabled(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setFraudulentWebsiteWarningEnabled:"), value)
}

// deprecated
func (p_ Preferences) JavaEnabled() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("javaEnabled"))
	return rv
}

// deprecated
func (p_ Preferences) SetJavaEnabled(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setJavaEnabled:"), value)
}

// deprecated
func (p_ Preferences) JavaScriptEnabled() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("javaScriptEnabled"))
	return rv
}

// deprecated
func (p_ Preferences) SetJavaScriptEnabled(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setJavaScriptEnabled:"), value)
}

// deprecated
func (p_ Preferences) PlugInsEnabled() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("plugInsEnabled"))
	return rv
}

// deprecated
func (p_ Preferences) SetPlugInsEnabled(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setPlugInsEnabled:"), value)
}

func (p_ Preferences) IsElementFullscreenEnabled() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isElementFullscreenEnabled"))
	return rv
}

func (p_ Preferences) SetElementFullscreenEnabled(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setElementFullscreenEnabled:"), value)
}

func (p_ Preferences) IsSiteSpecificQuirksModeEnabled() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isSiteSpecificQuirksModeEnabled"))
	return rv
}

func (p_ Preferences) SetSiteSpecificQuirksModeEnabled(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setSiteSpecificQuirksModeEnabled:"), value)
}

func (p_ Preferences) IsTextInteractionEnabled() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("isTextInteractionEnabled"))
	return rv
}

func (p_ Preferences) SetTextInteractionEnabled(value bool) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setTextInteractionEnabled:"), value)
}
