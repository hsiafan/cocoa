// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var ErrorClass = _ErrorClass{objc.GetClass("NSError")}

type _ErrorClass struct {
	objc.Class
}

type IError interface {
	objc.IObject
	AttemptRecoveryFromError_OptionIndex_Delegate_DidRecoverSelector_ContextInfo(error IError, recoveryOptionIndex uint, delegate objc.IObject, didRecoverSelector objc.Selector, contextInfo unsafe.Pointer)
	AttemptRecoveryFromError_OptionIndex(error IError, recoveryOptionIndex uint) bool
	Code() int
	Domain() ErrorDomain
	UserInfo() map[ErrorUserInfoKey]objc.Object
	LocalizedDescription() string
	LocalizedRecoveryOptions() []string
	LocalizedRecoverySuggestion() string
	LocalizedFailureReason() string
	RecoveryAttempter() objc.Object
	HelpAnchor() string
	UnderlyingErrors() []Error
}

type Error struct {
	objc.Object
}

func MakeError(ptr unsafe.Pointer) Error {
	return Error{
		Object: objc.MakeObject(ptr),
	}
}

func (ec _ErrorClass) ErrorWithDomain_Code_UserInfo(domain ErrorDomain, code int, dict map[ErrorUserInfoKey]objc.IObject) Error {
	rv := objc.CallMethod[Error](ec, objc.SEL("errorWithDomain:code:userInfo:"), domain, code, dict)
	return rv
}

func (e_ Error) InitWithDomain_Code_UserInfo(domain ErrorDomain, code int, dict map[ErrorUserInfoKey]objc.IObject) Error {
	rv := objc.CallMethod[Error](e_, objc.SEL("initWithDomain:code:userInfo:"), domain, code, dict)
	return rv
}

func (ec _ErrorClass) Alloc() Error {
	rv := objc.CallMethod[Error](ec, objc.SEL("alloc"))
	return rv
}

func (ec _ErrorClass) New() Error {
	rv := objc.CallMethod[Error](ec, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewError() Error {
	return ErrorClass.New()
}

func (e_ Error) Init() Error {
	rv := objc.CallMethod[Error](e_, objc.SEL("init"))
	return rv
}

func (ec _ErrorClass) SetUserInfoValueProviderForDomain_Provider(errorDomain ErrorDomain, provider func(err Error, userInfoKey ErrorUserInfoKey) objc.IObject) {
	objc.CallMethod[objc.Void](ec, objc.SEL("setUserInfoValueProviderForDomain:provider:"), errorDomain, provider)
}

func (ec _ErrorClass) UserInfoValueProviderForDomain(errorDomain ErrorDomain) func(param1 IError, param2 ErrorUserInfoKey) objc.Object {
	rv := objc.CallMethod[func(param1 IError, param2 ErrorUserInfoKey) objc.Object](ec, objc.SEL("userInfoValueProviderForDomain:"), errorDomain)
	return rv
}

func (e_ Error) AttemptRecoveryFromError_OptionIndex_Delegate_DidRecoverSelector_ContextInfo(error IError, recoveryOptionIndex uint, delegate objc.IObject, didRecoverSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](e_, objc.SEL("attemptRecoveryFromError:optionIndex:delegate:didRecoverSelector:contextInfo:"), objc.ExtractPtr(error), recoveryOptionIndex, objc.ExtractPtr(delegate), didRecoverSelector, contextInfo)
}

func (e_ Error) AttemptRecoveryFromError_OptionIndex(error IError, recoveryOptionIndex uint) bool {
	rv := objc.CallMethod[bool](e_, objc.SEL("attemptRecoveryFromError:optionIndex:"), objc.ExtractPtr(error), recoveryOptionIndex)
	return rv
}

func (e_ Error) Code() int {
	rv := objc.CallMethod[int](e_, objc.SEL("code"))
	return rv
}

func (e_ Error) Domain() ErrorDomain {
	rv := objc.CallMethod[ErrorDomain](e_, objc.SEL("domain"))
	return rv
}

func (e_ Error) UserInfo() map[ErrorUserInfoKey]objc.Object {
	rv := objc.CallMethod[map[ErrorUserInfoKey]objc.Object](e_, objc.SEL("userInfo"))
	return rv
}

func (e_ Error) LocalizedDescription() string {
	rv := objc.CallMethod[string](e_, objc.SEL("localizedDescription"))
	return rv
}

func (e_ Error) LocalizedRecoveryOptions() []string {
	rv := objc.CallMethod[[]string](e_, objc.SEL("localizedRecoveryOptions"))
	return rv
}

func (e_ Error) LocalizedRecoverySuggestion() string {
	rv := objc.CallMethod[string](e_, objc.SEL("localizedRecoverySuggestion"))
	return rv
}

func (e_ Error) LocalizedFailureReason() string {
	rv := objc.CallMethod[string](e_, objc.SEL("localizedFailureReason"))
	return rv
}

func (e_ Error) RecoveryAttempter() objc.Object {
	rv := objc.CallMethod[objc.Object](e_, objc.SEL("recoveryAttempter"))
	return rv
}

func (e_ Error) HelpAnchor() string {
	rv := objc.CallMethod[string](e_, objc.SEL("helpAnchor"))
	return rv
}

func (e_ Error) UnderlyingErrors() []Error {
	rv := objc.CallMethod[[]Error](e_, objc.SEL("underlyingErrors"))
	return rv
}
