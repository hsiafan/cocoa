// auto-generated code, do not modify
package webkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type URLSchemeHandler interface {
	// required
	WebView_StartURLSchemeTask(webView WebView, urlSchemeTask URLSchemeTask)
	// required
	WebView_StopURLSchemeTask(webView WebView, urlSchemeTask URLSchemeTask)
}

func WrapURLSchemeHandler(v URLSchemeHandler) objc.Object {
	return objc.WrapAsProtocol("WKURLSchemeHandler", v)
}

type URLSchemeHandlerBase struct {
}
