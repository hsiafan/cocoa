// auto-generated code, do not modify
package webkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

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
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:decidePolicyForNavigationAction:preferences:decisionHandler:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigationAction), objc.ExtractPtr(preferences), decisionHandler)
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_DecidePolicyForNavigationAction_DecisionHandler() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:decidePolicyForNavigationAction:decisionHandler:"))
}

func (n_ NavigationDelegateWrapper) WebView_DecidePolicyForNavigationAction_DecisionHandler(webView IWebView, navigationAction INavigationAction, decisionHandler func(param1 NavigationActionPolicy)) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:decidePolicyForNavigationAction:decisionHandler:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigationAction), decisionHandler)
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_DecidePolicyForNavigationResponse_DecisionHandler() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:decidePolicyForNavigationResponse:decisionHandler:"))
}

func (n_ NavigationDelegateWrapper) WebView_DecidePolicyForNavigationResponse_DecisionHandler(webView IWebView, navigationResponse INavigationResponse, decisionHandler func(param1 NavigationResponsePolicy)) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:decidePolicyForNavigationResponse:decisionHandler:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigationResponse), decisionHandler)
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_DidStartProvisionalNavigation() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:didStartProvisionalNavigation:"))
}

func (n_ NavigationDelegateWrapper) WebView_DidStartProvisionalNavigation(webView IWebView, navigation INavigation) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:didStartProvisionalNavigation:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigation))
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_DidReceiveServerRedirectForProvisionalNavigation() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:didReceiveServerRedirectForProvisionalNavigation:"))
}

func (n_ NavigationDelegateWrapper) WebView_DidReceiveServerRedirectForProvisionalNavigation(webView IWebView, navigation INavigation) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:didReceiveServerRedirectForProvisionalNavigation:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigation))
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_DidCommitNavigation() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:didCommitNavigation:"))
}

func (n_ NavigationDelegateWrapper) WebView_DidCommitNavigation(webView IWebView, navigation INavigation) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:didCommitNavigation:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigation))
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_DidFinishNavigation() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:didFinishNavigation:"))
}

func (n_ NavigationDelegateWrapper) WebView_DidFinishNavigation(webView IWebView, navigation INavigation) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:didFinishNavigation:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigation))
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_DidFailNavigation_WithError() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:didFailNavigation:withError:"))
}

func (n_ NavigationDelegateWrapper) WebView_DidFailNavigation_WithError(webView IWebView, navigation INavigation, error foundation.IError) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:didFailNavigation:withError:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigation), objc.ExtractPtr(error))
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_DidFailProvisionalNavigation_WithError() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:didFailProvisionalNavigation:withError:"))
}

func (n_ NavigationDelegateWrapper) WebView_DidFailProvisionalNavigation_WithError(webView IWebView, navigation INavigation, error foundation.IError) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:didFailProvisionalNavigation:withError:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigation), objc.ExtractPtr(error))
}

func (n_ *NavigationDelegateWrapper) ImplementsWebViewWebContentProcessDidTerminate() bool {
	return n_.RespondsToSelector(objc.GetSelector("webViewWebContentProcessDidTerminate:"))
}

func (n_ NavigationDelegateWrapper) WebViewWebContentProcessDidTerminate(webView IWebView) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webViewWebContentProcessDidTerminate:"), objc.ExtractPtr(webView))
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_NavigationResponse_DidBecomeDownload() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:navigationResponse:didBecomeDownload:"))
}

func (n_ NavigationDelegateWrapper) WebView_NavigationResponse_DidBecomeDownload(webView IWebView, navigationResponse INavigationResponse, download IDownload) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:navigationResponse:didBecomeDownload:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigationResponse), objc.ExtractPtr(download))
}

func (n_ *NavigationDelegateWrapper) ImplementsWebView_NavigationAction_DidBecomeDownload() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:navigationAction:didBecomeDownload:"))
}

func (n_ NavigationDelegateWrapper) WebView_NavigationAction_DidBecomeDownload(webView IWebView, navigationAction INavigationAction, download IDownload) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:navigationAction:didBecomeDownload:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigationAction), objc.ExtractPtr(download))
}
