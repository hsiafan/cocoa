// auto-generated code, do not modify
package webkit

import (
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

type ScriptMessageHandlerWithReply interface {
	// required
	UserContentController_DidReceiveScriptMessage_ReplyHandler(userContentController UserContentController, message ScriptMessage, replyHandler func(reply objc.IObject, errorMessage string))
}

type ScriptMessageHandlerWithReplyWrapper struct {
	objc.Object
}

func (s_ ScriptMessageHandlerWithReplyWrapper) UserContentController_DidReceiveScriptMessage_ReplyHandler(userContentController IUserContentController, message IScriptMessage, replyHandler func(reply objc.Object, errorMessage string)) {
	ffi.CallMethod[ffi.Void](s_, "userContentController:didReceiveScriptMessage:replyHandler:", userContentController, message, replyHandler)
}
