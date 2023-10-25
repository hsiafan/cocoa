// auto-generated code, do not modify
package webkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type UIDelegate interface {
	ImplementsWebView_CreateWebViewWithConfiguration_ForNavigationAction_WindowFeatures() bool
	// optional
	WebView_CreateWebViewWithConfiguration_ForNavigationAction_WindowFeatures(webView WebView, configuration WebViewConfiguration, navigationAction NavigationAction, windowFeatures WindowFeatures) IWebView
	ImplementsWebViewDidClose() bool
	// optional
	WebViewDidClose(webView WebView)
	ImplementsWebView_RunJavaScriptAlertPanelWithMessage_InitiatedByFrame_CompletionHandler() bool
	// optional
	WebView_RunJavaScriptAlertPanelWithMessage_InitiatedByFrame_CompletionHandler(webView WebView, message string, frame FrameInfo, completionHandler func())
	ImplementsWebView_RunJavaScriptConfirmPanelWithMessage_InitiatedByFrame_CompletionHandler() bool
	// optional
	WebView_RunJavaScriptConfirmPanelWithMessage_InitiatedByFrame_CompletionHandler(webView WebView, message string, frame FrameInfo, completionHandler func(result bool))
	ImplementsWebView_RunJavaScriptTextInputPanelWithPrompt_DefaultText_InitiatedByFrame_CompletionHandler() bool
	// optional
	WebView_RunJavaScriptTextInputPanelWithPrompt_DefaultText_InitiatedByFrame_CompletionHandler(webView WebView, prompt string, defaultText string, frame FrameInfo, completionHandler func(result string))
	ImplementsWebView_RequestMediaCapturePermissionForOrigin_InitiatedByFrame_Type_DecisionHandler() bool
	// optional
	WebView_RequestMediaCapturePermissionForOrigin_InitiatedByFrame_Type_DecisionHandler(webView WebView, origin SecurityOrigin, frame FrameInfo, type_ MediaCaptureType, decisionHandler func(decision PermissionDecision))
}

func WrapUIDelegate(v UIDelegate) objc.Object {
	return objc.WrapAsProtocol("WKUIDelegate", v)
}

type UIDelegateBase struct {
}

func (p *UIDelegateBase) ImplementsWebView_CreateWebViewWithConfiguration_ForNavigationAction_WindowFeatures() bool {
	return false
}

func (p *UIDelegateBase) WebView_CreateWebViewWithConfiguration_ForNavigationAction_WindowFeatures(webView WebView, configuration WebViewConfiguration, navigationAction NavigationAction, windowFeatures WindowFeatures) IWebView {
	panic("unimpemented protocol method")
}

func (p *UIDelegateBase) ImplementsWebViewDidClose() bool {
	return false
}

func (p *UIDelegateBase) WebViewDidClose(webView WebView) {
	panic("unimpemented protocol method")
}

func (p *UIDelegateBase) ImplementsWebView_RunJavaScriptAlertPanelWithMessage_InitiatedByFrame_CompletionHandler() bool {
	return false
}

func (p *UIDelegateBase) WebView_RunJavaScriptAlertPanelWithMessage_InitiatedByFrame_CompletionHandler(webView WebView, message string, frame FrameInfo, completionHandler func()) {
	panic("unimpemented protocol method")
}

func (p *UIDelegateBase) ImplementsWebView_RunJavaScriptConfirmPanelWithMessage_InitiatedByFrame_CompletionHandler() bool {
	return false
}

func (p *UIDelegateBase) WebView_RunJavaScriptConfirmPanelWithMessage_InitiatedByFrame_CompletionHandler(webView WebView, message string, frame FrameInfo, completionHandler func(result bool)) {
	panic("unimpemented protocol method")
}

func (p *UIDelegateBase) ImplementsWebView_RunJavaScriptTextInputPanelWithPrompt_DefaultText_InitiatedByFrame_CompletionHandler() bool {
	return false
}

func (p *UIDelegateBase) WebView_RunJavaScriptTextInputPanelWithPrompt_DefaultText_InitiatedByFrame_CompletionHandler(webView WebView, prompt string, defaultText string, frame FrameInfo, completionHandler func(result string)) {
	panic("unimpemented protocol method")
}

func (p *UIDelegateBase) ImplementsWebView_RequestMediaCapturePermissionForOrigin_InitiatedByFrame_Type_DecisionHandler() bool {
	return false
}

func (p *UIDelegateBase) WebView_RequestMediaCapturePermissionForOrigin_InitiatedByFrame_Type_DecisionHandler(webView WebView, origin SecurityOrigin, frame FrameInfo, type_ MediaCaptureType, decisionHandler func(decision PermissionDecision)) {
	panic("unimpemented protocol method")
}
