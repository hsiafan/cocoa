// auto-generated code, do not modify
package webkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type URLSchemeHandler interface {
	// required
	WebView_StartURLSchemeTask(webView WebView, urlSchemeTask URLSchemeTaskWrapper)
	// required
	WebView_StopURLSchemeTask(webView WebView, urlSchemeTask URLSchemeTaskWrapper)
}

type URLSchemeHandlerWrapper struct {
	objc.Object
}

func (u_ URLSchemeHandlerWrapper) WebView_StartURLSchemeTask(webView IWebView, urlSchemeTask URLSchemeTask) {
	po := objc.WrapAsProtocol("WKURLSchemeTask", urlSchemeTask)
	objc.CallMethod[objc.Void](u_, objc.GetSelector("webView:startURLSchemeTask:"), objc.ExtractPtr(webView), po)
}

func (u_ URLSchemeHandlerWrapper) WebView_StopURLSchemeTask(webView IWebView, urlSchemeTask URLSchemeTask) {
	po := objc.WrapAsProtocol("WKURLSchemeTask", urlSchemeTask)
	objc.CallMethod[objc.Void](u_, objc.GetSelector("webView:stopURLSchemeTask:"), objc.ExtractPtr(webView), po)
}
