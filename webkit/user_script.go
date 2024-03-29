// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var UserScriptClass = _UserScriptClass{objc.GetClass("WKUserScript")}

type _UserScriptClass struct {
	objc.Class
}

type IUserScript interface {
	objc.IObject
	Source() string
	InjectionTime() UserScriptInjectionTime
	IsForMainFrameOnly() bool
}

type UserScript struct {
	objc.Object
}

func MakeUserScript(ptr unsafe.Pointer) UserScript {
	return UserScript{
		Object: objc.MakeObject(ptr),
	}
}

func (u_ UserScript) InitWithSource_InjectionTime_ForMainFrameOnly(source string, injectionTime UserScriptInjectionTime, forMainFrameOnly bool) UserScript {
	rv := objc.CallMethod[UserScript](u_, objc.SEL("initWithSource:injectionTime:forMainFrameOnly:"), source, injectionTime, forMainFrameOnly)
	return rv
}

func (u_ UserScript) InitWithSource_InjectionTime_ForMainFrameOnly_InContentWorld(source string, injectionTime UserScriptInjectionTime, forMainFrameOnly bool, contentWorld IContentWorld) UserScript {
	rv := objc.CallMethod[UserScript](u_, objc.SEL("initWithSource:injectionTime:forMainFrameOnly:inContentWorld:"), source, injectionTime, forMainFrameOnly, objc.ExtractPtr(contentWorld))
	return rv
}

func (uc _UserScriptClass) Alloc() UserScript {
	rv := objc.CallMethod[UserScript](uc, objc.SEL("alloc"))
	return rv
}

func (uc _UserScriptClass) New() UserScript {
	rv := objc.CallMethod[UserScript](uc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewUserScript() UserScript {
	return UserScriptClass.New()
}

func (u_ UserScript) Init() UserScript {
	rv := objc.CallMethod[UserScript](u_, objc.SEL("init"))
	return rv
}

func (u_ UserScript) Source() string {
	rv := objc.CallMethod[string](u_, objc.SEL("source"))
	return rv
}

func (u_ UserScript) InjectionTime() UserScriptInjectionTime {
	rv := objc.CallMethod[UserScriptInjectionTime](u_, objc.SEL("injectionTime"))
	return rv
}

func (u_ UserScript) IsForMainFrameOnly() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("isForMainFrameOnly"))
	return rv
}
