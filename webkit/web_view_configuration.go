// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var WebViewConfigurationClass = _WebViewConfigurationClass{objc.GetClass("WKWebViewConfiguration")}

type _WebViewConfigurationClass struct {
	objc.Class
}

type IWebViewConfiguration interface {
	objc.IObject
	SetURLSchemeHandler_ForURLScheme(urlSchemeHandler URLSchemeHandler, urlScheme string)
	UrlSchemeHandlerForURLScheme(urlScheme string) URLSchemeHandlerWrapper
	WebsiteDataStore() WebsiteDataStore
	SetWebsiteDataStore(value IWebsiteDataStore)
	UserContentController() UserContentController
	SetUserContentController(value IUserContentController)
	ProcessPool() ProcessPool
	SetProcessPool(value IProcessPool)
	ApplicationNameForUserAgent() string
	SetApplicationNameForUserAgent(value string)
	LimitsNavigationsToAppBoundDomains() bool
	SetLimitsNavigationsToAppBoundDomains(value bool)
	Preferences() Preferences
	SetPreferences(value IPreferences)
	DefaultWebpagePreferences() WebpagePreferences
	SetDefaultWebpagePreferences(value IWebpagePreferences)
	SuppressesIncrementalRendering() bool
	SetSuppressesIncrementalRendering(value bool)
	AllowsAirPlayForMediaPlayback() bool
	SetAllowsAirPlayForMediaPlayback(value bool)
	MediaTypesRequiringUserActionForPlayback() AudiovisualMediaTypes
	SetMediaTypesRequiringUserActionForPlayback(value AudiovisualMediaTypes)
	UserInterfaceDirectionPolicy() UserInterfaceDirectionPolicy
	SetUserInterfaceDirectionPolicy(value UserInterfaceDirectionPolicy)
	UpgradeKnownHostsToHTTPS() bool
	SetUpgradeKnownHostsToHTTPS(value bool)
}

type WebViewConfiguration struct {
	objc.Object
}

func MakeWebViewConfiguration(ptr unsafe.Pointer) WebViewConfiguration {
	return WebViewConfiguration{
		Object: objc.MakeObject(ptr),
	}
}

func (wc _WebViewConfigurationClass) Alloc() WebViewConfiguration {
	rv := ffi.CallMethod[WebViewConfiguration](wc, "alloc")
	return rv
}

func (w_ WebViewConfiguration) Init() WebViewConfiguration {
	rv := ffi.CallMethod[WebViewConfiguration](w_, "init")
	return rv
}

func (wc _WebViewConfigurationClass) New() WebViewConfiguration {
	rv := ffi.CallMethod[WebViewConfiguration](wc, "new")
	rv.Autorelease()
	return rv
}

func NewWebViewConfiguration() WebViewConfiguration {
	return WebViewConfigurationClass.New()
}

func (w_ WebViewConfiguration) SetURLSchemeHandler_ForURLScheme(urlSchemeHandler URLSchemeHandler, urlScheme string) {
	po := ffi.CreateProtocol(urlSchemeHandler)
	defer po.Release()
	ffi.CallMethod[ffi.Void](w_, "setURLSchemeHandler:forURLScheme:", po, urlScheme)
}

func (w_ WebViewConfiguration) UrlSchemeHandlerForURLScheme(urlScheme string) URLSchemeHandlerWrapper {
	rv := ffi.CallMethod[URLSchemeHandlerWrapper](w_, "urlSchemeHandlerForURLScheme:", urlScheme)
	return rv
}

func (w_ WebViewConfiguration) WebsiteDataStore() WebsiteDataStore {
	rv := ffi.CallMethod[WebsiteDataStore](w_, "websiteDataStore")
	return rv
}

func (w_ WebViewConfiguration) SetWebsiteDataStore(value IWebsiteDataStore) {
	ffi.CallMethod[ffi.Void](w_, "setWebsiteDataStore:", value)
}

func (w_ WebViewConfiguration) UserContentController() UserContentController {
	rv := ffi.CallMethod[UserContentController](w_, "userContentController")
	return rv
}

func (w_ WebViewConfiguration) SetUserContentController(value IUserContentController) {
	ffi.CallMethod[ffi.Void](w_, "setUserContentController:", value)
}

func (w_ WebViewConfiguration) ProcessPool() ProcessPool {
	rv := ffi.CallMethod[ProcessPool](w_, "processPool")
	return rv
}

func (w_ WebViewConfiguration) SetProcessPool(value IProcessPool) {
	ffi.CallMethod[ffi.Void](w_, "setProcessPool:", value)
}

func (w_ WebViewConfiguration) ApplicationNameForUserAgent() string {
	rv := ffi.CallMethod[string](w_, "applicationNameForUserAgent")
	return rv
}

func (w_ WebViewConfiguration) SetApplicationNameForUserAgent(value string) {
	ffi.CallMethod[ffi.Void](w_, "setApplicationNameForUserAgent:", value)
}

func (w_ WebViewConfiguration) LimitsNavigationsToAppBoundDomains() bool {
	rv := ffi.CallMethod[bool](w_, "limitsNavigationsToAppBoundDomains")
	return rv
}

func (w_ WebViewConfiguration) SetLimitsNavigationsToAppBoundDomains(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setLimitsNavigationsToAppBoundDomains:", value)
}

func (w_ WebViewConfiguration) Preferences() Preferences {
	rv := ffi.CallMethod[Preferences](w_, "preferences")
	return rv
}

func (w_ WebViewConfiguration) SetPreferences(value IPreferences) {
	ffi.CallMethod[ffi.Void](w_, "setPreferences:", value)
}

func (w_ WebViewConfiguration) DefaultWebpagePreferences() WebpagePreferences {
	rv := ffi.CallMethod[WebpagePreferences](w_, "defaultWebpagePreferences")
	return rv
}

func (w_ WebViewConfiguration) SetDefaultWebpagePreferences(value IWebpagePreferences) {
	ffi.CallMethod[ffi.Void](w_, "setDefaultWebpagePreferences:", value)
}

func (w_ WebViewConfiguration) SuppressesIncrementalRendering() bool {
	rv := ffi.CallMethod[bool](w_, "suppressesIncrementalRendering")
	return rv
}

func (w_ WebViewConfiguration) SetSuppressesIncrementalRendering(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setSuppressesIncrementalRendering:", value)
}

func (w_ WebViewConfiguration) AllowsAirPlayForMediaPlayback() bool {
	rv := ffi.CallMethod[bool](w_, "allowsAirPlayForMediaPlayback")
	return rv
}

func (w_ WebViewConfiguration) SetAllowsAirPlayForMediaPlayback(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setAllowsAirPlayForMediaPlayback:", value)
}

func (w_ WebViewConfiguration) MediaTypesRequiringUserActionForPlayback() AudiovisualMediaTypes {
	rv := ffi.CallMethod[AudiovisualMediaTypes](w_, "mediaTypesRequiringUserActionForPlayback")
	return rv
}

func (w_ WebViewConfiguration) SetMediaTypesRequiringUserActionForPlayback(value AudiovisualMediaTypes) {
	ffi.CallMethod[ffi.Void](w_, "setMediaTypesRequiringUserActionForPlayback:", value)
}

func (w_ WebViewConfiguration) UserInterfaceDirectionPolicy() UserInterfaceDirectionPolicy {
	rv := ffi.CallMethod[UserInterfaceDirectionPolicy](w_, "userInterfaceDirectionPolicy")
	return rv
}

func (w_ WebViewConfiguration) SetUserInterfaceDirectionPolicy(value UserInterfaceDirectionPolicy) {
	ffi.CallMethod[ffi.Void](w_, "setUserInterfaceDirectionPolicy:", value)
}

func (w_ WebViewConfiguration) UpgradeKnownHostsToHTTPS() bool {
	rv := ffi.CallMethod[bool](w_, "upgradeKnownHostsToHTTPS")
	return rv
}

func (w_ WebViewConfiguration) SetUpgradeKnownHostsToHTTPS(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setUpgradeKnownHostsToHTTPS:", value)
}
