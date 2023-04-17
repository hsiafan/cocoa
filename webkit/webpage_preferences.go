// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

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

func (wc _WebpagePreferencesClass) New() WebpagePreferences {
	rv := ffi.CallMethod[WebpagePreferences](wc, "new")
	rv.Autorelease()
	return rv
}

func NewWebpagePreferences() WebpagePreferences {
	return WebpagePreferencesClass.New()
}

func (w_ WebpagePreferences) Init() WebpagePreferences {
	rv := ffi.CallMethod[WebpagePreferences](w_, "init")
	return rv
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
