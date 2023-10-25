// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var WebViewConfigurationClass = _WebViewConfigurationClass{objc.GetClass("WKWebViewConfiguration")}

type _WebViewConfigurationClass struct {
	objc.Class
}

type IWebViewConfiguration interface {
	objc.IObject
	SetURLSchemeHandler_ForURLScheme(urlSchemeHandler objc.IObject, urlScheme string)
	UrlSchemeHandlerForURLScheme(urlScheme string) objc.Object
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
	rv := objc.CallMethod[WebViewConfiguration](wc, objc.SEL("alloc"))
	return rv
}

func (wc _WebViewConfigurationClass) New() WebViewConfiguration {
	rv := objc.CallMethod[WebViewConfiguration](wc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewWebViewConfiguration() WebViewConfiguration {
	return WebViewConfigurationClass.New()
}

func (w_ WebViewConfiguration) Init() WebViewConfiguration {
	rv := objc.CallMethod[WebViewConfiguration](w_, objc.SEL("init"))
	return rv
}

func (w_ WebViewConfiguration) SetURLSchemeHandler_ForURLScheme(urlSchemeHandler objc.IObject, urlScheme string) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setURLSchemeHandler:forURLScheme:"), objc.ExtractPtr(urlSchemeHandler), urlScheme)
}

func (w_ WebViewConfiguration) UrlSchemeHandlerForURLScheme(urlScheme string) objc.Object {
	rv := objc.CallMethod[objc.Object](w_, objc.SEL("urlSchemeHandlerForURLScheme:"), urlScheme)
	return rv
}

func (w_ WebViewConfiguration) WebsiteDataStore() WebsiteDataStore {
	rv := objc.CallMethod[WebsiteDataStore](w_, objc.SEL("websiteDataStore"))
	return rv
}

func (w_ WebViewConfiguration) SetWebsiteDataStore(value IWebsiteDataStore) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setWebsiteDataStore:"), objc.ExtractPtr(value))
}

func (w_ WebViewConfiguration) UserContentController() UserContentController {
	rv := objc.CallMethod[UserContentController](w_, objc.SEL("userContentController"))
	return rv
}

func (w_ WebViewConfiguration) SetUserContentController(value IUserContentController) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setUserContentController:"), objc.ExtractPtr(value))
}

func (w_ WebViewConfiguration) ProcessPool() ProcessPool {
	rv := objc.CallMethod[ProcessPool](w_, objc.SEL("processPool"))
	return rv
}

func (w_ WebViewConfiguration) SetProcessPool(value IProcessPool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setProcessPool:"), objc.ExtractPtr(value))
}

func (w_ WebViewConfiguration) ApplicationNameForUserAgent() string {
	rv := objc.CallMethod[string](w_, objc.SEL("applicationNameForUserAgent"))
	return rv
}

func (w_ WebViewConfiguration) SetApplicationNameForUserAgent(value string) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setApplicationNameForUserAgent:"), value)
}

func (w_ WebViewConfiguration) LimitsNavigationsToAppBoundDomains() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("limitsNavigationsToAppBoundDomains"))
	return rv
}

func (w_ WebViewConfiguration) SetLimitsNavigationsToAppBoundDomains(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setLimitsNavigationsToAppBoundDomains:"), value)
}

func (w_ WebViewConfiguration) Preferences() Preferences {
	rv := objc.CallMethod[Preferences](w_, objc.SEL("preferences"))
	return rv
}

func (w_ WebViewConfiguration) SetPreferences(value IPreferences) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setPreferences:"), objc.ExtractPtr(value))
}

func (w_ WebViewConfiguration) DefaultWebpagePreferences() WebpagePreferences {
	rv := objc.CallMethod[WebpagePreferences](w_, objc.SEL("defaultWebpagePreferences"))
	return rv
}

func (w_ WebViewConfiguration) SetDefaultWebpagePreferences(value IWebpagePreferences) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setDefaultWebpagePreferences:"), objc.ExtractPtr(value))
}

func (w_ WebViewConfiguration) SuppressesIncrementalRendering() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("suppressesIncrementalRendering"))
	return rv
}

func (w_ WebViewConfiguration) SetSuppressesIncrementalRendering(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setSuppressesIncrementalRendering:"), value)
}

func (w_ WebViewConfiguration) AllowsAirPlayForMediaPlayback() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("allowsAirPlayForMediaPlayback"))
	return rv
}

func (w_ WebViewConfiguration) SetAllowsAirPlayForMediaPlayback(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setAllowsAirPlayForMediaPlayback:"), value)
}

func (w_ WebViewConfiguration) MediaTypesRequiringUserActionForPlayback() AudiovisualMediaTypes {
	rv := objc.CallMethod[AudiovisualMediaTypes](w_, objc.SEL("mediaTypesRequiringUserActionForPlayback"))
	return rv
}

func (w_ WebViewConfiguration) SetMediaTypesRequiringUserActionForPlayback(value AudiovisualMediaTypes) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setMediaTypesRequiringUserActionForPlayback:"), value)
}

func (w_ WebViewConfiguration) UserInterfaceDirectionPolicy() UserInterfaceDirectionPolicy {
	rv := objc.CallMethod[UserInterfaceDirectionPolicy](w_, objc.SEL("userInterfaceDirectionPolicy"))
	return rv
}

func (w_ WebViewConfiguration) SetUserInterfaceDirectionPolicy(value UserInterfaceDirectionPolicy) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setUserInterfaceDirectionPolicy:"), value)
}

func (w_ WebViewConfiguration) UpgradeKnownHostsToHTTPS() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("upgradeKnownHostsToHTTPS"))
	return rv
}

func (w_ WebViewConfiguration) SetUpgradeKnownHostsToHTTPS(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setUpgradeKnownHostsToHTTPS:"), value)
}
