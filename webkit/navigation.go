// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var NavigationClass = _NavigationClass{objc.GetClass("WKNavigation")}

type _NavigationClass struct {
	objc.Class
}

type INavigation interface {
	objc.IObject
	EffectiveContentMode() ContentMode
}

type Navigation struct {
	objc.Object
}

func MakeNavigation(ptr unsafe.Pointer) Navigation {
	return Navigation{
		Object: objc.MakeObject(ptr),
	}
}

func (nc _NavigationClass) Alloc() Navigation {
	rv := ffi.CallMethod[Navigation](nc, "alloc")
	return rv
}

func (n_ Navigation) Init() Navigation {
	rv := ffi.CallMethod[Navigation](n_, "init")
	rv.Autorelease()
	return rv
}

func (nc _NavigationClass) New() Navigation {
	rv := ffi.CallMethod[Navigation](nc, "new")
	rv.Autorelease()
	return rv
}

func NewNavigation() Navigation {
	return NavigationClass.New()
}

func (n_ Navigation) EffectiveContentMode() ContentMode {
	rv := ffi.CallMethod[ContentMode](n_, "effectiveContentMode")
	return rv
}

type NavigationDelegate interface {
	ImplementsWebView_DecidePolicyForNavigationAction_Preferences_DecisionHandler() bool
	// optional
	WebView_DecidePolicyForNavigationAction_Preferences_DecisionHandler(webView WebView, navigationAction NavigationAction, preferences WebpagePreferences, decisionHandler func(param1 NavigationActionPolicy, param2 IWebpagePreferences))
	ImplementsWebView_DecidePolicyForNavigationAction_DecisionHandler() bool
	// optional
	WebView_DecidePolicyForNavigationAction_DecisionHandler(webView WebView, navigationAction NavigationAction, decisionHandler func(param1 NavigationActionPolicy))
	ImplementsWebView_DecidePolicyForNavigationResponse_DecisionHandler() bool
	// optional
	WebView_DecidePolicyForNavigationResponse_DecisionHandler(webView WebView, navigationResponse NavigationResponse, decisionHandler func(param1 NavigationResponsePolicy))
	ImplementsWebView_DidStartProvisionalNavigation() bool
	// optional
	WebView_DidStartProvisionalNavigation(webView WebView, navigation Navigation)
	ImplementsWebView_DidReceiveServerRedirectForProvisionalNavigation() bool
	// optional
	WebView_DidReceiveServerRedirectForProvisionalNavigation(webView WebView, navigation Navigation)
	ImplementsWebView_DidCommitNavigation() bool
	// optional
	WebView_DidCommitNavigation(webView WebView, navigation Navigation)
	ImplementsWebView_DidFinishNavigation() bool
	// optional
	WebView_DidFinishNavigation(webView WebView, navigation Navigation)
	ImplementsWebView_DidFailNavigation_WithError() bool
	// optional
	WebView_DidFailNavigation_WithError(webView WebView, navigation Navigation, error foundation.Error)
	ImplementsWebView_DidFailProvisionalNavigation_WithError() bool
	// optional
	WebView_DidFailProvisionalNavigation_WithError(webView WebView, navigation Navigation, error foundation.Error)
	ImplementsWebViewWebContentProcessDidTerminate() bool
	// optional
	WebViewWebContentProcessDidTerminate(webView WebView)
	ImplementsWebView_NavigationResponse_DidBecomeDownload() bool
	// optional
	WebView_NavigationResponse_DidBecomeDownload(webView WebView, navigationResponse NavigationResponse, download Download)
	ImplementsWebView_NavigationAction_DidBecomeDownload() bool
	// optional
	WebView_NavigationAction_DidBecomeDownload(webView WebView, navigationAction NavigationAction, download Download)
}

type NavigationDelegateImpl struct {
	_WebView_DecidePolicyForNavigationAction_Preferences_DecisionHandler func(webView WebView, navigationAction NavigationAction, preferences WebpagePreferences, decisionHandler func(param1 NavigationActionPolicy, param2 IWebpagePreferences))
	_WebView_DecidePolicyForNavigationAction_DecisionHandler             func(webView WebView, navigationAction NavigationAction, decisionHandler func(param1 NavigationActionPolicy))
	_WebView_DecidePolicyForNavigationResponse_DecisionHandler           func(webView WebView, navigationResponse NavigationResponse, decisionHandler func(param1 NavigationResponsePolicy))
	_WebView_DidStartProvisionalNavigation                               func(webView WebView, navigation Navigation)
	_WebView_DidReceiveServerRedirectForProvisionalNavigation            func(webView WebView, navigation Navigation)
	_WebView_DidCommitNavigation                                         func(webView WebView, navigation Navigation)
	_WebView_DidFinishNavigation                                         func(webView WebView, navigation Navigation)
	_WebView_DidFailNavigation_WithError                                 func(webView WebView, navigation Navigation, error foundation.Error)
	_WebView_DidFailProvisionalNavigation_WithError                      func(webView WebView, navigation Navigation, error foundation.Error)
	_WebViewWebContentProcessDidTerminate                                func(webView WebView)
	_WebView_NavigationResponse_DidBecomeDownload                        func(webView WebView, navigationResponse NavigationResponse, download Download)
	_WebView_NavigationAction_DidBecomeDownload                          func(webView WebView, navigationAction NavigationAction, download Download)
}

func (di *NavigationDelegateImpl) ImplementsWebView_DecidePolicyForNavigationAction_Preferences_DecisionHandler() bool {
	return di._WebView_DecidePolicyForNavigationAction_Preferences_DecisionHandler != nil
}

func (di *NavigationDelegateImpl) SetWebView_DecidePolicyForNavigationAction_Preferences_DecisionHandler(f func(webView WebView, navigationAction NavigationAction, preferences WebpagePreferences, decisionHandler func(param1 NavigationActionPolicy, param2 IWebpagePreferences))) {
	di._WebView_DecidePolicyForNavigationAction_Preferences_DecisionHandler = f
}

func (di *NavigationDelegateImpl) WebView_DecidePolicyForNavigationAction_Preferences_DecisionHandler(webView WebView, navigationAction NavigationAction, preferences WebpagePreferences, decisionHandler func(param1 NavigationActionPolicy, param2 IWebpagePreferences)) {
	di._WebView_DecidePolicyForNavigationAction_Preferences_DecisionHandler(webView, navigationAction, preferences, decisionHandler)
}
func (di *NavigationDelegateImpl) ImplementsWebView_DecidePolicyForNavigationAction_DecisionHandler() bool {
	return di._WebView_DecidePolicyForNavigationAction_DecisionHandler != nil
}

func (di *NavigationDelegateImpl) SetWebView_DecidePolicyForNavigationAction_DecisionHandler(f func(webView WebView, navigationAction NavigationAction, decisionHandler func(param1 NavigationActionPolicy))) {
	di._WebView_DecidePolicyForNavigationAction_DecisionHandler = f
}

func (di *NavigationDelegateImpl) WebView_DecidePolicyForNavigationAction_DecisionHandler(webView WebView, navigationAction NavigationAction, decisionHandler func(param1 NavigationActionPolicy)) {
	di._WebView_DecidePolicyForNavigationAction_DecisionHandler(webView, navigationAction, decisionHandler)
}
func (di *NavigationDelegateImpl) ImplementsWebView_DecidePolicyForNavigationResponse_DecisionHandler() bool {
	return di._WebView_DecidePolicyForNavigationResponse_DecisionHandler != nil
}

func (di *NavigationDelegateImpl) SetWebView_DecidePolicyForNavigationResponse_DecisionHandler(f func(webView WebView, navigationResponse NavigationResponse, decisionHandler func(param1 NavigationResponsePolicy))) {
	di._WebView_DecidePolicyForNavigationResponse_DecisionHandler = f
}

func (di *NavigationDelegateImpl) WebView_DecidePolicyForNavigationResponse_DecisionHandler(webView WebView, navigationResponse NavigationResponse, decisionHandler func(param1 NavigationResponsePolicy)) {
	di._WebView_DecidePolicyForNavigationResponse_DecisionHandler(webView, navigationResponse, decisionHandler)
}
func (di *NavigationDelegateImpl) ImplementsWebView_DidStartProvisionalNavigation() bool {
	return di._WebView_DidStartProvisionalNavigation != nil
}

func (di *NavigationDelegateImpl) SetWebView_DidStartProvisionalNavigation(f func(webView WebView, navigation Navigation)) {
	di._WebView_DidStartProvisionalNavigation = f
}

func (di *NavigationDelegateImpl) WebView_DidStartProvisionalNavigation(webView WebView, navigation Navigation) {
	di._WebView_DidStartProvisionalNavigation(webView, navigation)
}
func (di *NavigationDelegateImpl) ImplementsWebView_DidReceiveServerRedirectForProvisionalNavigation() bool {
	return di._WebView_DidReceiveServerRedirectForProvisionalNavigation != nil
}

func (di *NavigationDelegateImpl) SetWebView_DidReceiveServerRedirectForProvisionalNavigation(f func(webView WebView, navigation Navigation)) {
	di._WebView_DidReceiveServerRedirectForProvisionalNavigation = f
}

func (di *NavigationDelegateImpl) WebView_DidReceiveServerRedirectForProvisionalNavigation(webView WebView, navigation Navigation) {
	di._WebView_DidReceiveServerRedirectForProvisionalNavigation(webView, navigation)
}
func (di *NavigationDelegateImpl) ImplementsWebView_DidCommitNavigation() bool {
	return di._WebView_DidCommitNavigation != nil
}

func (di *NavigationDelegateImpl) SetWebView_DidCommitNavigation(f func(webView WebView, navigation Navigation)) {
	di._WebView_DidCommitNavigation = f
}

func (di *NavigationDelegateImpl) WebView_DidCommitNavigation(webView WebView, navigation Navigation) {
	di._WebView_DidCommitNavigation(webView, navigation)
}
func (di *NavigationDelegateImpl) ImplementsWebView_DidFinishNavigation() bool {
	return di._WebView_DidFinishNavigation != nil
}

func (di *NavigationDelegateImpl) SetWebView_DidFinishNavigation(f func(webView WebView, navigation Navigation)) {
	di._WebView_DidFinishNavigation = f
}

func (di *NavigationDelegateImpl) WebView_DidFinishNavigation(webView WebView, navigation Navigation) {
	di._WebView_DidFinishNavigation(webView, navigation)
}
func (di *NavigationDelegateImpl) ImplementsWebView_DidFailNavigation_WithError() bool {
	return di._WebView_DidFailNavigation_WithError != nil
}

func (di *NavigationDelegateImpl) SetWebView_DidFailNavigation_WithError(f func(webView WebView, navigation Navigation, error foundation.Error)) {
	di._WebView_DidFailNavigation_WithError = f
}

func (di *NavigationDelegateImpl) WebView_DidFailNavigation_WithError(webView WebView, navigation Navigation, error foundation.Error) {
	di._WebView_DidFailNavigation_WithError(webView, navigation, error)
}
func (di *NavigationDelegateImpl) ImplementsWebView_DidFailProvisionalNavigation_WithError() bool {
	return di._WebView_DidFailProvisionalNavigation_WithError != nil
}

func (di *NavigationDelegateImpl) SetWebView_DidFailProvisionalNavigation_WithError(f func(webView WebView, navigation Navigation, error foundation.Error)) {
	di._WebView_DidFailProvisionalNavigation_WithError = f
}

func (di *NavigationDelegateImpl) WebView_DidFailProvisionalNavigation_WithError(webView WebView, navigation Navigation, error foundation.Error) {
	di._WebView_DidFailProvisionalNavigation_WithError(webView, navigation, error)
}
func (di *NavigationDelegateImpl) ImplementsWebViewWebContentProcessDidTerminate() bool {
	return di._WebViewWebContentProcessDidTerminate != nil
}

func (di *NavigationDelegateImpl) SetWebViewWebContentProcessDidTerminate(f func(webView WebView)) {
	di._WebViewWebContentProcessDidTerminate = f
}

func (di *NavigationDelegateImpl) WebViewWebContentProcessDidTerminate(webView WebView) {
	di._WebViewWebContentProcessDidTerminate(webView)
}
func (di *NavigationDelegateImpl) ImplementsWebView_NavigationResponse_DidBecomeDownload() bool {
	return di._WebView_NavigationResponse_DidBecomeDownload != nil
}

func (di *NavigationDelegateImpl) SetWebView_NavigationResponse_DidBecomeDownload(f func(webView WebView, navigationResponse NavigationResponse, download Download)) {
	di._WebView_NavigationResponse_DidBecomeDownload = f
}

func (di *NavigationDelegateImpl) WebView_NavigationResponse_DidBecomeDownload(webView WebView, navigationResponse NavigationResponse, download Download) {
	di._WebView_NavigationResponse_DidBecomeDownload(webView, navigationResponse, download)
}
func (di *NavigationDelegateImpl) ImplementsWebView_NavigationAction_DidBecomeDownload() bool {
	return di._WebView_NavigationAction_DidBecomeDownload != nil
}

func (di *NavigationDelegateImpl) SetWebView_NavigationAction_DidBecomeDownload(f func(webView WebView, navigationAction NavigationAction, download Download)) {
	di._WebView_NavigationAction_DidBecomeDownload = f
}

func (di *NavigationDelegateImpl) WebView_NavigationAction_DidBecomeDownload(webView WebView, navigationAction NavigationAction, download Download) {
	di._WebView_NavigationAction_DidBecomeDownload(webView, navigationAction, download)
}

type NavigationDelegateWrapper struct {
	objc.Object
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_DecidePolicyForNavigationAction_Preferences_DecisionHandler() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:decidePolicyForNavigationAction:preferences:decisionHandler:"))
}

func (n_ NavigationDelegateWrapper) WebView_DecidePolicyForNavigationAction_Preferences_DecisionHandler(webView IWebView, navigationAction INavigationAction, preferences IWebpagePreferences, decisionHandler func(param1 NavigationActionPolicy, param2 WebpagePreferences)) {
	ffi.CallMethod[ffi.Void](n_, "webView:decidePolicyForNavigationAction:preferences:decisionHandler:", webView, navigationAction, preferences, decisionHandler)
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_DecidePolicyForNavigationAction_DecisionHandler() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:decidePolicyForNavigationAction:decisionHandler:"))
}

func (n_ NavigationDelegateWrapper) WebView_DecidePolicyForNavigationAction_DecisionHandler(webView IWebView, navigationAction INavigationAction, decisionHandler func(param1 NavigationActionPolicy)) {
	ffi.CallMethod[ffi.Void](n_, "webView:decidePolicyForNavigationAction:decisionHandler:", webView, navigationAction, decisionHandler)
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_DecidePolicyForNavigationResponse_DecisionHandler() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:decidePolicyForNavigationResponse:decisionHandler:"))
}

func (n_ NavigationDelegateWrapper) WebView_DecidePolicyForNavigationResponse_DecisionHandler(webView IWebView, navigationResponse INavigationResponse, decisionHandler func(param1 NavigationResponsePolicy)) {
	ffi.CallMethod[ffi.Void](n_, "webView:decidePolicyForNavigationResponse:decisionHandler:", webView, navigationResponse, decisionHandler)
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_DidStartProvisionalNavigation() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:didStartProvisionalNavigation:"))
}

func (n_ NavigationDelegateWrapper) WebView_DidStartProvisionalNavigation(webView IWebView, navigation INavigation) {
	ffi.CallMethod[ffi.Void](n_, "webView:didStartProvisionalNavigation:", webView, navigation)
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_DidReceiveServerRedirectForProvisionalNavigation() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:didReceiveServerRedirectForProvisionalNavigation:"))
}

func (n_ NavigationDelegateWrapper) WebView_DidReceiveServerRedirectForProvisionalNavigation(webView IWebView, navigation INavigation) {
	ffi.CallMethod[ffi.Void](n_, "webView:didReceiveServerRedirectForProvisionalNavigation:", webView, navigation)
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_DidCommitNavigation() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:didCommitNavigation:"))
}

func (n_ NavigationDelegateWrapper) WebView_DidCommitNavigation(webView IWebView, navigation INavigation) {
	ffi.CallMethod[ffi.Void](n_, "webView:didCommitNavigation:", webView, navigation)
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_DidFinishNavigation() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:didFinishNavigation:"))
}

func (n_ NavigationDelegateWrapper) WebView_DidFinishNavigation(webView IWebView, navigation INavigation) {
	ffi.CallMethod[ffi.Void](n_, "webView:didFinishNavigation:", webView, navigation)
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_DidFailNavigation_WithError() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:didFailNavigation:withError:"))
}

func (n_ NavigationDelegateWrapper) WebView_DidFailNavigation_WithError(webView IWebView, navigation INavigation, error foundation.IError) {
	ffi.CallMethod[ffi.Void](n_, "webView:didFailNavigation:withError:", webView, navigation, error)
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_DidFailProvisionalNavigation_WithError() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:didFailProvisionalNavigation:withError:"))
}

func (n_ NavigationDelegateWrapper) WebView_DidFailProvisionalNavigation_WithError(webView IWebView, navigation INavigation, error foundation.IError) {
	ffi.CallMethod[ffi.Void](n_, "webView:didFailProvisionalNavigation:withError:", webView, navigation, error)
}

func (n_ *NavigationDelegateWrapper) ImplementsWebViewWebContentProcessDidTerminate() bool {
	return n_.RespondsToSelector(objc.GetSelector("webViewWebContentProcessDidTerminate:"))
}

func (n_ NavigationDelegateWrapper) WebViewWebContentProcessDidTerminate(webView IWebView) {
	ffi.CallMethod[ffi.Void](n_, "webViewWebContentProcessDidTerminate:", webView)
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_NavigationResponse_DidBecomeDownload() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:navigationResponse:didBecomeDownload:"))
}

func (n_ NavigationDelegateWrapper) WebView_NavigationResponse_DidBecomeDownload(webView IWebView, navigationResponse INavigationResponse, download IDownload) {
	ffi.CallMethod[ffi.Void](n_, "webView:navigationResponse:didBecomeDownload:", webView, navigationResponse, download)
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_NavigationAction_DidBecomeDownload() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:navigationAction:didBecomeDownload:"))
}

func (n_ NavigationDelegateWrapper) WebView_NavigationAction_DidBecomeDownload(webView IWebView, navigationAction INavigationAction, download IDownload) {
	ffi.CallMethod[ffi.Void](n_, "webView:navigationAction:didBecomeDownload:", webView, navigationAction, download)
}

var NavigationActionClass = _NavigationActionClass{objc.GetClass("WKNavigationAction")}

type _NavigationActionClass struct {
	objc.Class
}

type INavigationAction interface {
	objc.IObject
	NavigationType() NavigationType
	Request() foundation.URLRequest
	SourceFrame() FrameInfo
	TargetFrame() FrameInfo
	ButtonNumber() int
	ModifierFlags() appkit.EventModifierFlags
	ShouldPerformDownload() bool
}

type NavigationAction struct {
	objc.Object
}

func MakeNavigationAction(ptr unsafe.Pointer) NavigationAction {
	return NavigationAction{
		Object: objc.MakeObject(ptr),
	}
}

func (nc _NavigationActionClass) Alloc() NavigationAction {
	rv := ffi.CallMethod[NavigationAction](nc, "alloc")
	return rv
}

func (n_ NavigationAction) Init() NavigationAction {
	rv := ffi.CallMethod[NavigationAction](n_, "init")
	rv.Autorelease()
	return rv
}

func (nc _NavigationActionClass) New() NavigationAction {
	rv := ffi.CallMethod[NavigationAction](nc, "new")
	rv.Autorelease()
	return rv
}

func NewNavigationAction() NavigationAction {
	return NavigationActionClass.New()
}

func (n_ NavigationAction) NavigationType() NavigationType {
	rv := ffi.CallMethod[NavigationType](n_, "navigationType")
	return rv
}

func (n_ NavigationAction) Request() foundation.URLRequest {
	rv := ffi.CallMethod[foundation.URLRequest](n_, "request")
	return rv
}

func (n_ NavigationAction) SourceFrame() FrameInfo {
	rv := ffi.CallMethod[FrameInfo](n_, "sourceFrame")
	return rv
}

func (n_ NavigationAction) TargetFrame() FrameInfo {
	rv := ffi.CallMethod[FrameInfo](n_, "targetFrame")
	return rv
}

func (n_ NavigationAction) ButtonNumber() int {
	rv := ffi.CallMethod[int](n_, "buttonNumber")
	return rv
}

func (n_ NavigationAction) ModifierFlags() appkit.EventModifierFlags {
	rv := ffi.CallMethod[appkit.EventModifierFlags](n_, "modifierFlags")
	return rv
}

func (n_ NavigationAction) ShouldPerformDownload() bool {
	rv := ffi.CallMethod[bool](n_, "shouldPerformDownload")
	return rv
}

var NavigationResponseClass = _NavigationResponseClass{objc.GetClass("WKNavigationResponse")}

type _NavigationResponseClass struct {
	objc.Class
}

type INavigationResponse interface {
	objc.IObject
	Response() foundation.URLResponse
	CanShowMIMEType() bool
	IsForMainFrame() bool
}

type NavigationResponse struct {
	objc.Object
}

func MakeNavigationResponse(ptr unsafe.Pointer) NavigationResponse {
	return NavigationResponse{
		Object: objc.MakeObject(ptr),
	}
}

func (nc _NavigationResponseClass) Alloc() NavigationResponse {
	rv := ffi.CallMethod[NavigationResponse](nc, "alloc")
	return rv
}

func (n_ NavigationResponse) Init() NavigationResponse {
	rv := ffi.CallMethod[NavigationResponse](n_, "init")
	rv.Autorelease()
	return rv
}

func (nc _NavigationResponseClass) New() NavigationResponse {
	rv := ffi.CallMethod[NavigationResponse](nc, "new")
	rv.Autorelease()
	return rv
}

func NewNavigationResponse() NavigationResponse {
	return NavigationResponseClass.New()
}

func (n_ NavigationResponse) Response() foundation.URLResponse {
	rv := ffi.CallMethod[foundation.URLResponse](n_, "response")
	return rv
}

func (n_ NavigationResponse) CanShowMIMEType() bool {
	rv := ffi.CallMethod[bool](n_, "canShowMIMEType")
	return rv
}

func (n_ NavigationResponse) IsForMainFrame() bool {
	rv := ffi.CallMethod[bool](n_, "isForMainFrame")
	return rv
}
