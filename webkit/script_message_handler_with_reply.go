// auto-generated code, do not modify
package webkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type ScriptMessageHandlerWithReply interface {
	// required
	UserContentController_DidReceiveScriptMessage_ReplyHandler(userContentController UserContentController, message ScriptMessage, replyHandler func(reply objc.IObject, errorMessage foundation.String))
}

func WrapScriptMessageHandlerWithReply(v ScriptMessageHandlerWithReply) objc.Object {
	return objc.WrapAsProtocol("WKScriptMessageHandlerWithReply", v)
}

type ScriptMessageHandlerWithReplyBase struct {
}
