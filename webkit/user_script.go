// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[UserScript](u_, "initWithSource:injectionTime:forMainFrameOnly:", source, injectionTime, forMainFrameOnly)
	rv.Autorelease()
	return rv
}

func (u_ UserScript) InitWithSource_InjectionTime_ForMainFrameOnly_InContentWorld(source string, injectionTime UserScriptInjectionTime, forMainFrameOnly bool, contentWorld IContentWorld) UserScript {
	rv := ffi.CallMethod[UserScript](u_, "initWithSource:injectionTime:forMainFrameOnly:inContentWorld:", source, injectionTime, forMainFrameOnly, contentWorld)
	rv.Autorelease()
	return rv
}

func (uc _UserScriptClass) Alloc() UserScript {
	rv := ffi.CallMethod[UserScript](uc, "alloc")
	return rv
}

func (u_ UserScript) Init() UserScript {
	rv := ffi.CallMethod[UserScript](u_, "init")
	rv.Autorelease()
	return rv
}

func (uc _UserScriptClass) New() UserScript {
	rv := ffi.CallMethod[UserScript](uc, "new")
	rv.Autorelease()
	return rv
}

func NewUserScript() UserScript {
	return UserScriptClass.New()
}

func (u_ UserScript) Source() string {
	rv := ffi.CallMethod[string](u_, "source")
	return rv
}

func (u_ UserScript) InjectionTime() UserScriptInjectionTime {
	rv := ffi.CallMethod[UserScriptInjectionTime](u_, "injectionTime")
	return rv
}

func (u_ UserScript) IsForMainFrameOnly() bool {
	rv := ffi.CallMethod[bool](u_, "isForMainFrameOnly")
	return rv
}
