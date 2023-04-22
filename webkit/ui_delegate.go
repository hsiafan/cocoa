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

type UIDelegateImpl struct {
	_WebView_CreateWebViewWithConfiguration_ForNavigationAction_WindowFeatures                    func(webView WebView, configuration WebViewConfiguration, navigationAction NavigationAction, windowFeatures WindowFeatures) IWebView
	_WebViewDidClose                                                                              func(webView WebView)
	_WebView_RunJavaScriptAlertPanelWithMessage_InitiatedByFrame_CompletionHandler                func(webView WebView, message string, frame FrameInfo, completionHandler func())
	_WebView_RunJavaScriptConfirmPanelWithMessage_InitiatedByFrame_CompletionHandler              func(webView WebView, message string, frame FrameInfo, completionHandler func(result bool))
	_WebView_RunJavaScriptTextInputPanelWithPrompt_DefaultText_InitiatedByFrame_CompletionHandler func(webView WebView, prompt string, defaultText string, frame FrameInfo, completionHandler func(result string))
	_WebView_RequestMediaCapturePermissionForOrigin_InitiatedByFrame_Type_DecisionHandler         func(webView WebView, origin SecurityOrigin, frame FrameInfo, type_ MediaCaptureType, decisionHandler func(decision PermissionDecision))
}

func (di *UIDelegateImpl) ImplementsWebView_CreateWebViewWithConfiguration_ForNavigationAction_WindowFeatures() bool {
	return di._WebView_CreateWebViewWithConfiguration_ForNavigationAction_WindowFeatures != nil
}

func (di *UIDelegateImpl) SetWebView_CreateWebViewWithConfiguration_ForNavigationAction_WindowFeatures(f func(webView WebView, configuration WebViewConfiguration, navigationAction NavigationAction, windowFeatures WindowFeatures) IWebView) {
	di._WebView_CreateWebViewWithConfiguration_ForNavigationAction_WindowFeatures = f
}

func (di *UIDelegateImpl) WebView_CreateWebViewWithConfiguration_ForNavigationAction_WindowFeatures(webView WebView, configuration WebViewConfiguration, navigationAction NavigationAction, windowFeatures WindowFeatures) IWebView {
	return di._WebView_CreateWebViewWithConfiguration_ForNavigationAction_WindowFeatures(webView, configuration, navigationAction, windowFeatures)
}
func (di *UIDelegateImpl) ImplementsWebViewDidClose() bool {
	return di._WebViewDidClose != nil
}

func (di *UIDelegateImpl) SetWebViewDidClose(f func(webView WebView)) {
	di._WebViewDidClose = f
}

func (di *UIDelegateImpl) WebViewDidClose(webView WebView) {
	di._WebViewDidClose(webView)
}
func (di *UIDelegateImpl) ImplementsWebView_RunJavaScriptAlertPanelWithMessage_InitiatedByFrame_CompletionHandler() bool {
	return di._WebView_RunJavaScriptAlertPanelWithMessage_InitiatedByFrame_CompletionHandler != nil
}

func (di *UIDelegateImpl) SetWebView_RunJavaScriptAlertPanelWithMessage_InitiatedByFrame_CompletionHandler(f func(webView WebView, message string, frame FrameInfo, completionHandler func())) {
	di._WebView_RunJavaScriptAlertPanelWithMessage_InitiatedByFrame_CompletionHandler = f
}

func (di *UIDelegateImpl) WebView_RunJavaScriptAlertPanelWithMessage_InitiatedByFrame_CompletionHandler(webView WebView, message string, frame FrameInfo, completionHandler func()) {
	di._WebView_RunJavaScriptAlertPanelWithMessage_InitiatedByFrame_CompletionHandler(webView, message, frame, completionHandler)
}
func (di *UIDelegateImpl) ImplementsWebView_RunJavaScriptConfirmPanelWithMessage_InitiatedByFrame_CompletionHandler() bool {
	return di._WebView_RunJavaScriptConfirmPanelWithMessage_InitiatedByFrame_CompletionHandler != nil
}

func (di *UIDelegateImpl) SetWebView_RunJavaScriptConfirmPanelWithMessage_InitiatedByFrame_CompletionHandler(f func(webView WebView, message string, frame FrameInfo, completionHandler func(result bool))) {
	di._WebView_RunJavaScriptConfirmPanelWithMessage_InitiatedByFrame_CompletionHandler = f
}

func (di *UIDelegateImpl) WebView_RunJavaScriptConfirmPanelWithMessage_InitiatedByFrame_CompletionHandler(webView WebView, message string, frame FrameInfo, completionHandler func(result bool)) {
	di._WebView_RunJavaScriptConfirmPanelWithMessage_InitiatedByFrame_CompletionHandler(webView, message, frame, completionHandler)
}
func (di *UIDelegateImpl) ImplementsWebView_RunJavaScriptTextInputPanelWithPrompt_DefaultText_InitiatedByFrame_CompletionHandler() bool {
	return di._WebView_RunJavaScriptTextInputPanelWithPrompt_DefaultText_InitiatedByFrame_CompletionHandler != nil
}

func (di *UIDelegateImpl) SetWebView_RunJavaScriptTextInputPanelWithPrompt_DefaultText_InitiatedByFrame_CompletionHandler(f func(webView WebView, prompt string, defaultText string, frame FrameInfo, completionHandler func(result string))) {
	di._WebView_RunJavaScriptTextInputPanelWithPrompt_DefaultText_InitiatedByFrame_CompletionHandler = f
}

func (di *UIDelegateImpl) WebView_RunJavaScriptTextInputPanelWithPrompt_DefaultText_InitiatedByFrame_CompletionHandler(webView WebView, prompt string, defaultText string, frame FrameInfo, completionHandler func(result string)) {
	di._WebView_RunJavaScriptTextInputPanelWithPrompt_DefaultText_InitiatedByFrame_CompletionHandler(webView, prompt, defaultText, frame, completionHandler)
}
func (di *UIDelegateImpl) ImplementsWebView_RequestMediaCapturePermissionForOrigin_InitiatedByFrame_Type_DecisionHandler() bool {
	return di._WebView_RequestMediaCapturePermissionForOrigin_InitiatedByFrame_Type_DecisionHandler != nil
}

func (di *UIDelegateImpl) SetWebView_RequestMediaCapturePermissionForOrigin_InitiatedByFrame_Type_DecisionHandler(f func(webView WebView, origin SecurityOrigin, frame FrameInfo, type_ MediaCaptureType, decisionHandler func(decision PermissionDecision))) {
	di._WebView_RequestMediaCapturePermissionForOrigin_InitiatedByFrame_Type_DecisionHandler = f
}

func (di *UIDelegateImpl) WebView_RequestMediaCapturePermissionForOrigin_InitiatedByFrame_Type_DecisionHandler(webView WebView, origin SecurityOrigin, frame FrameInfo, type_ MediaCaptureType, decisionHandler func(decision PermissionDecision)) {
	di._WebView_RequestMediaCapturePermissionForOrigin_InitiatedByFrame_Type_DecisionHandler(webView, origin, frame, type_, decisionHandler)
}

type UIDelegateWrapper struct {
	objc.Object
}

func (u_ *UIDelegateWrapper) ImplementsWebView_CreateWebViewWithConfiguration_ForNavigationAction_WindowFeatures() bool {
	return u_.RespondsToSelector(objc.GetSelector("webView:createWebViewWithConfiguration:forNavigationAction:windowFeatures:"))
}

func (u_ UIDelegateWrapper) WebView_CreateWebViewWithConfiguration_ForNavigationAction_WindowFeatures(webView IWebView, configuration IWebViewConfiguration, navigationAction INavigationAction, windowFeatures IWindowFeatures) WebView {
	rv := objc.CallMethod[WebView](u_, objc.GetSelector("webView:createWebViewWithConfiguration:forNavigationAction:windowFeatures:"), webView, configuration, navigationAction, windowFeatures)
	return rv
}

func (u_ *UIDelegateWrapper) ImplementsWebViewDidClose() bool {
	return u_.RespondsToSelector(objc.GetSelector("webViewDidClose:"))
}

func (u_ UIDelegateWrapper) WebViewDidClose(webView IWebView) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("webViewDidClose:"), webView)
}

func (u_ *UIDelegateWrapper) ImplementsWebView_RunJavaScriptAlertPanelWithMessage_InitiatedByFrame_CompletionHandler() bool {
	return u_.RespondsToSelector(objc.GetSelector("webView:runJavaScriptAlertPanelWithMessage:initiatedByFrame:completionHandler:"))
}

func (u_ UIDelegateWrapper) WebView_RunJavaScriptAlertPanelWithMessage_InitiatedByFrame_CompletionHandler(webView IWebView, message string, frame IFrameInfo, completionHandler func()) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("webView:runJavaScriptAlertPanelWithMessage:initiatedByFrame:completionHandler:"), webView, message, frame, completionHandler)
}

func (u_ *UIDelegateWrapper) ImplementsWebView_RunJavaScriptConfirmPanelWithMessage_InitiatedByFrame_CompletionHandler() bool {
	return u_.RespondsToSelector(objc.GetSelector("webView:runJavaScriptConfirmPanelWithMessage:initiatedByFrame:completionHandler:"))
}

func (u_ UIDelegateWrapper) WebView_RunJavaScriptConfirmPanelWithMessage_InitiatedByFrame_CompletionHandler(webView IWebView, message string, frame IFrameInfo, completionHandler func(result bool)) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("webView:runJavaScriptConfirmPanelWithMessage:initiatedByFrame:completionHandler:"), webView, message, frame, completionHandler)
}

func (u_ *UIDelegateWrapper) ImplementsWebView_RunJavaScriptTextInputPanelWithPrompt_DefaultText_InitiatedByFrame_CompletionHandler() bool {
	return u_.RespondsToSelector(objc.GetSelector("webView:runJavaScriptTextInputPanelWithPrompt:defaultText:initiatedByFrame:completionHandler:"))
}

func (u_ UIDelegateWrapper) WebView_RunJavaScriptTextInputPanelWithPrompt_DefaultText_InitiatedByFrame_CompletionHandler(webView IWebView, prompt string, defaultText string, frame IFrameInfo, completionHandler func(result string)) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("webView:runJavaScriptTextInputPanelWithPrompt:defaultText:initiatedByFrame:completionHandler:"), webView, prompt, defaultText, frame, completionHandler)
}

func (u_ *UIDelegateWrapper) ImplementsWebView_RequestMediaCapturePermissionForOrigin_InitiatedByFrame_Type_DecisionHandler() bool {
	return u_.RespondsToSelector(objc.GetSelector("webView:requestMediaCapturePermissionForOrigin:initiatedByFrame:type:decisionHandler:"))
}

func (u_ UIDelegateWrapper) WebView_RequestMediaCapturePermissionForOrigin_InitiatedByFrame_Type_DecisionHandler(webView IWebView, origin ISecurityOrigin, frame IFrameInfo, type_ MediaCaptureType, decisionHandler func(decision PermissionDecision)) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("webView:requestMediaCapturePermissionForOrigin:initiatedByFrame:type:decisionHandler:"), webView, origin, frame, type_, decisionHandler)
}
