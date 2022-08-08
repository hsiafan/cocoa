// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[Preferences](pc, "alloc")
	return rv
}

func (p_ Preferences) Init() Preferences {
	rv := ffi.CallMethod[Preferences](p_, "init")
	return rv
}

func (pc _PreferencesClass) New() Preferences {
	rv := ffi.CallMethod[Preferences](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPreferences() Preferences {
	return PreferencesClass.New()
}

func (p_ Preferences) MinimumFontSize() float64 {
	rv := ffi.CallMethod[float64](p_, "minimumFontSize")
	return rv
}

func (p_ Preferences) SetMinimumFontSize(value float64) {
	ffi.CallMethod[ffi.Void](p_, "setMinimumFontSize:", value)
}

func (p_ Preferences) TabFocusesLinks() bool {
	rv := ffi.CallMethod[bool](p_, "tabFocusesLinks")
	return rv
}

func (p_ Preferences) SetTabFocusesLinks(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setTabFocusesLinks:", value)
}

func (p_ Preferences) JavaScriptCanOpenWindowsAutomatically() bool {
	rv := ffi.CallMethod[bool](p_, "javaScriptCanOpenWindowsAutomatically")
	return rv
}

func (p_ Preferences) SetJavaScriptCanOpenWindowsAutomatically(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setJavaScriptCanOpenWindowsAutomatically:", value)
}

func (p_ Preferences) IsFraudulentWebsiteWarningEnabled() bool {
	rv := ffi.CallMethod[bool](p_, "isFraudulentWebsiteWarningEnabled")
	return rv
}

func (p_ Preferences) SetFraudulentWebsiteWarningEnabled(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setFraudulentWebsiteWarningEnabled:", value)
}

// deprecated
func (p_ Preferences) JavaEnabled() bool {
	rv := ffi.CallMethod[bool](p_, "javaEnabled")
	return rv
}

// deprecated
func (p_ Preferences) SetJavaEnabled(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setJavaEnabled:", value)
}

// deprecated
func (p_ Preferences) JavaScriptEnabled() bool {
	rv := ffi.CallMethod[bool](p_, "javaScriptEnabled")
	return rv
}

// deprecated
func (p_ Preferences) SetJavaScriptEnabled(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setJavaScriptEnabled:", value)
}

// deprecated
func (p_ Preferences) PlugInsEnabled() bool {
	rv := ffi.CallMethod[bool](p_, "plugInsEnabled")
	return rv
}

// deprecated
func (p_ Preferences) SetPlugInsEnabled(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setPlugInsEnabled:", value)
}

func (p_ Preferences) IsElementFullscreenEnabled() bool {
	rv := ffi.CallMethod[bool](p_, "isElementFullscreenEnabled")
	return rv
}

func (p_ Preferences) SetElementFullscreenEnabled(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setElementFullscreenEnabled:", value)
}

func (p_ Preferences) IsSiteSpecificQuirksModeEnabled() bool {
	rv := ffi.CallMethod[bool](p_, "isSiteSpecificQuirksModeEnabled")
	return rv
}

func (p_ Preferences) SetSiteSpecificQuirksModeEnabled(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setSiteSpecificQuirksModeEnabled:", value)
}

func (p_ Preferences) IsTextInteractionEnabled() bool {
	rv := ffi.CallMethod[bool](p_, "isTextInteractionEnabled")
	return rv
}

func (p_ Preferences) SetTextInteractionEnabled(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setTextInteractionEnabled:", value)
}

var WebpagePreferencesClass = _WebpagePreferencesClass{objc.GetClass("WKWebpagePreferences")}

type _WebpagePreferencesClass struct {
	objc.Class
}

type IWebpagePreferences interface {
	objc.IObject
	AllowsContentJavaScript() bool
	SetAllowsContentJavaScript(value bool)
	PreferredContentMode() ContentMode
	SetPreferredContentMode(value ContentMode)
	IsLockdownModeEnabled() bool
	SetLockdownModeEnabled(value bool)
}

type WebpagePreferences struct {
	objc.Object
}

func MakeWebpagePreferences(ptr unsafe.Pointer) WebpagePreferences {
	return WebpagePreferences{
		Object: objc.MakeObject(ptr),
	}
}

func (wc _WebpagePreferencesClass) Alloc() WebpagePreferences {
	rv := ffi.CallMethod[WebpagePreferences](wc, "alloc")
	return rv
}

func (w_ WebpagePreferences) Init() WebpagePreferences {
	rv := ffi.CallMethod[WebpagePreferences](w_, "init")
	return rv
}

func (wc _WebpagePreferencesClass) New() WebpagePreferences {
	rv := ffi.CallMethod[WebpagePreferences](wc, "new")
	rv.Autorelease()
	return rv
}

func NewWebpagePreferences() WebpagePreferences {
	return WebpagePreferencesClass.New()
}

func (w_ WebpagePreferences) AllowsContentJavaScript() bool {
	rv := ffi.CallMethod[bool](w_, "allowsContentJavaScript")
	return rv
}

func (w_ WebpagePreferences) SetAllowsContentJavaScript(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setAllowsContentJavaScript:", value)
}

func (w_ WebpagePreferences) PreferredContentMode() ContentMode {
	rv := ffi.CallMethod[ContentMode](w_, "preferredContentMode")
	return rv
}

func (w_ WebpagePreferences) SetPreferredContentMode(value ContentMode) {
	ffi.CallMethod[ffi.Void](w_, "setPreferredContentMode:", value)
}

func (w_ WebpagePreferences) IsLockdownModeEnabled() bool {
	rv := ffi.CallMethod[bool](w_, "isLockdownModeEnabled")
	return rv
}

func (w_ WebpagePreferences) SetLockdownModeEnabled(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setLockdownModeEnabled:", value)
}
