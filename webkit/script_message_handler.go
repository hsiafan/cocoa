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

type ScriptMessageHandlerWrapper struct {
	objc.Object
}

func (s_ ScriptMessageHandlerWrapper) UserContentController_DidReceiveScriptMessage(userContentController IUserContentController, message IScriptMessage) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("userContentController:didReceiveScriptMessage:"), objc.ExtractPtr(userContentController), objc.ExtractPtr(message))
}
