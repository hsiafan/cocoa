// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var UserContentControllerClass = _UserContentControllerClass{objc.GetClass("WKUserContentController")}

type _UserContentControllerClass struct {
	objc.Class
}

type IUserContentController interface {
	objc.IObject
	AddUserScript(userScript IUserScript)
	RemoveAllUserScripts()
	AddScriptMessageHandler_Name(scriptMessageHandler ScriptMessageHandler, name string)
	AddScriptMessageHandler0_Name(scriptMessageHandler objc.IObject, name string)
	AddScriptMessageHandler_ContentWorld_Name(scriptMessageHandler ScriptMessageHandler, world IContentWorld, name string)
	AddScriptMessageHandler0_ContentWorld_Name(scriptMessageHandler objc.IObject, world IContentWorld, name string)
	AddScriptMessageHandlerWithReply_ContentWorld_Name(scriptMessageHandlerWithReply ScriptMessageHandlerWithReply, contentWorld IContentWorld, name string)
	AddScriptMessageHandlerWithReply0_ContentWorld_Name(scriptMessageHandlerWithReply objc.IObject, contentWorld IContentWorld, name string)
	RemoveScriptMessageHandlerForName(name string)
	RemoveScriptMessageHandlerForName_ContentWorld(name string, contentWorld IContentWorld)
	RemoveAllScriptMessageHandlersFromContentWorld(contentWorld IContentWorld)
	RemoveAllScriptMessageHandlers()
	AddContentRuleList(contentRuleList IContentRuleList)
	RemoveContentRuleList(contentRuleList IContentRuleList)
	RemoveAllContentRuleLists()
	UserScripts() []UserScript
}

type UserContentController struct {
	objc.Object
}

func MakeUserContentController(ptr unsafe.Pointer) UserContentController {
	return UserContentController{
		Object: objc.MakeObject(ptr),
	}
}

func (uc _UserContentControllerClass) Alloc() UserContentController {
	rv := objc.CallMethod[UserContentController](uc, objc.GetSelector("alloc"))
	return rv
}

func (uc _UserContentControllerClass) New() UserContentController {
	rv := objc.CallMethod[UserContentController](uc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewUserContentController() UserContentController {
	return UserContentControllerClass.New()
}

func (u_ UserContentController) Init() UserContentController {
	rv := objc.CallMethod[UserContentController](u_, objc.GetSelector("init"))
	return rv
}

func (u_ UserContentController) AddUserScript(userScript IUserScript) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("addUserScript:"), userScript)
}

func (u_ UserContentController) RemoveAllUserScripts() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeAllUserScripts"))
}

func (u_ UserContentController) AddScriptMessageHandler_Name(scriptMessageHandler ScriptMessageHandler, name string) {
	po := objc.CreateProtocol("WKScriptMessageHandler", scriptMessageHandler)
	objc.CallMethod[objc.Void](u_, objc.GetSelector("addScriptMessageHandler:name:"), po, name)
}

func (u_ UserContentController) AddScriptMessageHandler0_Name(scriptMessageHandler objc.IObject, name string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("addScriptMessageHandler:name:"), scriptMessageHandler, name)
}

func (u_ UserContentController) AddScriptMessageHandler_ContentWorld_Name(scriptMessageHandler ScriptMessageHandler, world IContentWorld, name string) {
	po := objc.CreateProtocol("WKScriptMessageHandler", scriptMessageHandler)
	objc.CallMethod[objc.Void](u_, objc.GetSelector("addScriptMessageHandler:contentWorld:name:"), po, world, name)
}

func (u_ UserContentController) AddScriptMessageHandler0_ContentWorld_Name(scriptMessageHandler objc.IObject, world IContentWorld, name string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("addScriptMessageHandler:contentWorld:name:"), scriptMessageHandler, world, name)
}

func (u_ UserContentController) AddScriptMessageHandlerWithReply_ContentWorld_Name(scriptMessageHandlerWithReply ScriptMessageHandlerWithReply, contentWorld IContentWorld, name string) {
	po := objc.CreateProtocol("WKScriptMessageHandlerWithReply", scriptMessageHandlerWithReply)
	objc.CallMethod[objc.Void](u_, objc.GetSelector("addScriptMessageHandlerWithReply:contentWorld:name:"), po, contentWorld, name)
}

func (u_ UserContentController) AddScriptMessageHandlerWithReply0_ContentWorld_Name(scriptMessageHandlerWithReply objc.IObject, contentWorld IContentWorld, name string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("addScriptMessageHandlerWithReply:contentWorld:name:"), scriptMessageHandlerWithReply, contentWorld, name)
}

func (u_ UserContentController) RemoveScriptMessageHandlerForName(name string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeScriptMessageHandlerForName:"), name)
}

func (u_ UserContentController) RemoveScriptMessageHandlerForName_ContentWorld(name string, contentWorld IContentWorld) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeScriptMessageHandlerForName:contentWorld:"), name, contentWorld)
}

func (u_ UserContentController) RemoveAllScriptMessageHandlersFromContentWorld(contentWorld IContentWorld) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeAllScriptMessageHandlersFromContentWorld:"), contentWorld)
}

func (u_ UserContentController) RemoveAllScriptMessageHandlers() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeAllScriptMessageHandlers"))
}

func (u_ UserContentController) AddContentRuleList(contentRuleList IContentRuleList) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("addContentRuleList:"), contentRuleList)
}

func (u_ UserContentController) RemoveContentRuleList(contentRuleList IContentRuleList) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeContentRuleList:"), contentRuleList)
}

func (u_ UserContentController) RemoveAllContentRuleLists() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeAllContentRuleLists"))
}

func (u_ UserContentController) UserScripts() []UserScript {
	rv := objc.CallMethod[[]UserScript](u_, objc.GetSelector("userScripts"))
	return rv
}
