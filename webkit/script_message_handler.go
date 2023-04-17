// auto-generated code, do not modify
package webkit

import (
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

type ScriptMessageHandler interface {
	// required
	UserContentController_DidReceiveScriptMessage(userContentController UserContentController, message ScriptMessage)
}

type ScriptMessageHandlerWrapper struct {
	objc.Object
}

func (s_ ScriptMessageHandlerWrapper) UserContentController_DidReceiveScriptMessage(userContentController IUserContentController, message IScriptMessage) {
	ffi.CallMethod[ffi.Void](s_, "userContentController:didReceiveScriptMessage:", userContentController, message)
}
