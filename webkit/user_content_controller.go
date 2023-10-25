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
	AddScriptMessageHandler_Name(scriptMessageHandler objc.IObject, name string)
	AddScriptMessageHandler_ContentWorld_Name(scriptMessageHandler objc.IObject, world IContentWorld, name string)
	AddScriptMessageHandlerWithReply_ContentWorld_Name(scriptMessageHandlerWithReply objc.IObject, contentWorld IContentWorld, name string)
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
	rv := objc.CallMethod[UserContentController](uc, objc.SEL("alloc"))
	return rv
}

func (uc _UserContentControllerClass) New() UserContentController {
	rv := objc.CallMethod[UserContentController](uc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewUserContentController() UserContentController {
	return UserContentControllerClass.New()
}

func (u_ UserContentController) Init() UserContentController {
	rv := objc.CallMethod[UserContentController](u_, objc.SEL("init"))
	return rv
}

func (u_ UserContentController) AddUserScript(userScript IUserScript) {
	objc.CallMethod[objc.Void](u_, objc.SEL("addUserScript:"), objc.ExtractPtr(userScript))
}

func (u_ UserContentController) RemoveAllUserScripts() {
	objc.CallMethod[objc.Void](u_, objc.SEL("removeAllUserScripts"))
}

func (u_ UserContentController) AddScriptMessageHandler_Name(scriptMessageHandler objc.IObject, name string) {
	objc.CallMethod[objc.Void](u_, objc.SEL("addScriptMessageHandler:name:"), objc.ExtractPtr(scriptMessageHandler), name)
}

func (u_ UserContentController) AddScriptMessageHandler_ContentWorld_Name(scriptMessageHandler objc.IObject, world IContentWorld, name string) {
	objc.CallMethod[objc.Void](u_, objc.SEL("addScriptMessageHandler:contentWorld:name:"), objc.ExtractPtr(scriptMessageHandler), objc.ExtractPtr(world), name)
}

func (u_ UserContentController) AddScriptMessageHandlerWithReply_ContentWorld_Name(scriptMessageHandlerWithReply objc.IObject, contentWorld IContentWorld, name string) {
	objc.CallMethod[objc.Void](u_, objc.SEL("addScriptMessageHandlerWithReply:contentWorld:name:"), objc.ExtractPtr(scriptMessageHandlerWithReply), objc.ExtractPtr(contentWorld), name)
}

func (u_ UserContentController) RemoveScriptMessageHandlerForName(name string) {
	objc.CallMethod[objc.Void](u_, objc.SEL("removeScriptMessageHandlerForName:"), name)
}

func (u_ UserContentController) RemoveScriptMessageHandlerForName_ContentWorld(name string, contentWorld IContentWorld) {
	objc.CallMethod[objc.Void](u_, objc.SEL("removeScriptMessageHandlerForName:contentWorld:"), name, objc.ExtractPtr(contentWorld))
}

func (u_ UserContentController) RemoveAllScriptMessageHandlersFromContentWorld(contentWorld IContentWorld) {
	objc.CallMethod[objc.Void](u_, objc.SEL("removeAllScriptMessageHandlersFromContentWorld:"), objc.ExtractPtr(contentWorld))
}

func (u_ UserContentController) RemoveAllScriptMessageHandlers() {
	objc.CallMethod[objc.Void](u_, objc.SEL("removeAllScriptMessageHandlers"))
}

func (u_ UserContentController) AddContentRuleList(contentRuleList IContentRuleList) {
	objc.CallMethod[objc.Void](u_, objc.SEL("addContentRuleList:"), objc.ExtractPtr(contentRuleList))
}

func (u_ UserContentController) RemoveContentRuleList(contentRuleList IContentRuleList) {
	objc.CallMethod[objc.Void](u_, objc.SEL("removeContentRuleList:"), objc.ExtractPtr(contentRuleList))
}

func (u_ UserContentController) RemoveAllContentRuleLists() {
	objc.CallMethod[objc.Void](u_, objc.SEL("removeAllContentRuleLists"))
}

func (u_ UserContentController) UserScripts() []UserScript {
	rv := objc.CallMethod[[]UserScript](u_, objc.SEL("userScripts"))
	return rv
}
