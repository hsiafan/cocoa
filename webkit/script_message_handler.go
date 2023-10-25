// auto-generated code, do not modify
package webkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type ScriptMessageHandler interface {
	// required
	UserContentController_DidReceiveScriptMessage(userContentController UserContentController, message ScriptMessage)
}

func WrapScriptMessageHandler(v ScriptMessageHandler) objc.Object {
	return objc.WrapAsProtocol("WKScriptMessageHandler", v)
}

type ScriptMessageHandlerBase struct {
}
