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

func WrapNavigationDelegate(v NavigationDelegate) objc.Object {
	return objc.WrapAsProtocol("WKNavigationDelegate", v)
}

type NavigationDelegateBase struct {
}

func (p *NavigationDelegateBase) ImplementsWebView_DecidePolicyForNavigationAction_Preferences_DecisionHandler() bool {
	return false
}

func (p *NavigationDelegateBase) WebView_DecidePolicyForNavigationAction_Preferences_DecisionHandler(webView WebView, navigationAction NavigationAction, preferences WebpagePreferences, decisionHandler func(param1 NavigationActionPolicy, param2 IWebpagePreferences)) {
	panic("unimpemented protocol method")
}

func (p *NavigationDelegateBase) ImplementsWebView_DecidePolicyForNavigationAction_DecisionHandler() bool {
	return false
}

func (p *NavigationDelegateBase) WebView_DecidePolicyForNavigationAction_DecisionHandler(webView WebView, navigationAction NavigationAction, decisionHandler func(param1 NavigationActionPolicy)) {
	panic("unimpemented protocol method")
}

func (p *NavigationDelegateBase) ImplementsWebView_DecidePolicyForNavigationResponse_DecisionHandler() bool {
	return false
}

func (p *NavigationDelegateBase) WebView_DecidePolicyForNavigationResponse_DecisionHandler(webView WebView, navigationResponse NavigationResponse, decisionHandler func(param1 NavigationResponsePolicy)) {
	panic("unimpemented protocol method")
}

func (p *NavigationDelegateBase) ImplementsWebView_DidStartProvisionalNavigation() bool {
	return false
}

func (p *NavigationDelegateBase) WebView_DidStartProvisionalNavigation(webView WebView, navigation Navigation) {
	panic("unimpemented protocol method")
}

func (p *NavigationDelegateBase) ImplementsWebView_DidReceiveServerRedirectForProvisionalNavigation() bool {
	return false
}

func (p *NavigationDelegateBase) WebView_DidReceiveServerRedirectForProvisionalNavigation(webView WebView, navigation Navigation) {
	panic("unimpemented protocol method")
}

func (p *NavigationDelegateBase) ImplementsWebView_DidCommitNavigation() bool {
	return false
}

func (p *NavigationDelegateBase) WebView_DidCommitNavigation(webView WebView, navigation Navigation) {
	panic("unimpemented protocol method")
}

func (p *NavigationDelegateBase) ImplementsWebView_DidFinishNavigation() bool {
	return false
}

func (p *NavigationDelegateBase) WebView_DidFinishNavigation(webView WebView, navigation Navigation) {
	panic("unimpemented protocol method")
}

func (p *NavigationDelegateBase) ImplementsWebView_DidFailNavigation_WithError() bool {
	return false
}

func (p *NavigationDelegateBase) WebView_DidFailNavigation_WithError(webView WebView, navigation Navigation, error foundation.Error) {
	panic("unimpemented protocol method")
}

func (p *NavigationDelegateBase) ImplementsWebView_DidFailProvisionalNavigation_WithError() bool {
	return false
}

func (p *NavigationDelegateBase) WebView_DidFailProvisionalNavigation_WithError(webView WebView, navigation Navigation, error foundation.Error) {
	panic("unimpemented protocol method")
}

func (p *NavigationDelegateBase) ImplementsWebViewWebContentProcessDidTerminate() bool {
	return false
}

func (p *NavigationDelegateBase) WebViewWebContentProcessDidTerminate(webView WebView) {
	panic("unimpemented protocol method")
}

func (p *NavigationDelegateBase) ImplementsWebView_NavigationResponse_DidBecomeDownload() bool {
	return false
}

func (p *NavigationDelegateBase) WebView_NavigationResponse_DidBecomeDownload(webView WebView, navigationResponse NavigationResponse, download Download) {
	panic("unimpemented protocol method")
}

func (p *NavigationDelegateBase) ImplementsWebView_NavigationAction_DidBecomeDownload() bool {
	return false
}

func (p *NavigationDelegateBase) WebView_NavigationAction_DidBecomeDownload(webView WebView, navigationAction NavigationAction, download Download) {
	panic("unimpemented protocol method")
}

type NavigationDelegateCreator struct {
	className string
	class     objc.Class
}

func NewNavigationDelegateCreator(name string) *NavigationDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &NavigationDelegateCreator{className: name, class: class}
}

func (c *NavigationDelegateCreator) SetWebView_DecidePolicyForNavigationAction_Preferences_DecisionHandler(handle func(o objc.Object, webView WebView, navigationAction NavigationAction, preferences WebpagePreferences, decisionHandler func(param1 NavigationActionPolicy, param2 IWebpagePreferences))) {
	objc.AddMethod(c.class, objc.SEL("webView:decidePolicyForNavigationAction:preferences:decisionHandler:"), handle)
}

func (c *NavigationDelegateCreator) SetWebView_DecidePolicyForNavigationAction_DecisionHandler(handle func(o objc.Object, webView WebView, navigationAction NavigationAction, decisionHandler func(param1 NavigationActionPolicy))) {
	objc.AddMethod(c.class, objc.SEL("webView:decidePolicyForNavigationAction:decisionHandler:"), handle)
}

func (c *NavigationDelegateCreator) SetWebView_DecidePolicyForNavigationResponse_DecisionHandler(handle func(o objc.Object, webView WebView, navigationResponse NavigationResponse, decisionHandler func(param1 NavigationResponsePolicy))) {
	objc.AddMethod(c.class, objc.SEL("webView:decidePolicyForNavigationResponse:decisionHandler:"), handle)
}

func (c *NavigationDelegateCreator) SetWebView_DidStartProvisionalNavigation(handle func(o objc.Object, webView WebView, navigation Navigation)) {
	objc.AddMethod(c.class, objc.SEL("webView:didStartProvisionalNavigation:"), handle)
}

func (c *NavigationDelegateCreator) SetWebView_DidReceiveServerRedirectForProvisionalNavigation(handle func(o objc.Object, webView WebView, navigation Navigation)) {
	objc.AddMethod(c.class, objc.SEL("webView:didReceiveServerRedirectForProvisionalNavigation:"), handle)
}

func (c *NavigationDelegateCreator) SetWebView_DidCommitNavigation(handle func(o objc.Object, webView WebView, navigation Navigation)) {
	objc.AddMethod(c.class, objc.SEL("webView:didCommitNavigation:"), handle)
}

func (c *NavigationDelegateCreator) SetWebView_DidFinishNavigation(handle func(o objc.Object, webView WebView, navigation Navigation)) {
	objc.AddMethod(c.class, objc.SEL("webView:didFinishNavigation:"), handle)
}

func (c *NavigationDelegateCreator) SetWebView_DidFailNavigation_WithError(handle func(o objc.Object, webView WebView, navigation Navigation, error foundation.Error)) {
	objc.AddMethod(c.class, objc.SEL("webView:didFailNavigation:withError:"), handle)
}

func (c *NavigationDelegateCreator) SetWebView_DidFailProvisionalNavigation_WithError(handle func(o objc.Object, webView WebView, navigation Navigation, error foundation.Error)) {
	objc.AddMethod(c.class, objc.SEL("webView:didFailProvisionalNavigation:withError:"), handle)
}

func (c *NavigationDelegateCreator) SetWebViewWebContentProcessDidTerminate(handle func(o objc.Object, webView WebView)) {
	objc.AddMethod(c.class, objc.SEL("webViewWebContentProcessDidTerminate:"), handle)
}

func (c *NavigationDelegateCreator) SetWebView_NavigationResponse_DidBecomeDownload(handle func(o objc.Object, webView WebView, navigationResponse NavigationResponse, download Download)) {
	objc.AddMethod(c.class, objc.SEL("webView:navigationResponse:didBecomeDownload:"), handle)
}

func (c *NavigationDelegateCreator) SetWebView_NavigationAction_DidBecomeDownload(handle func(o objc.Object, webView WebView, navigationAction NavigationAction, download Download)) {
	objc.AddMethod(c.class, objc.SEL("webView:navigationAction:didBecomeDownload:"), handle)
}

func (c *NavigationDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
