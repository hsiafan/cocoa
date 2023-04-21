// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var ExtensionContextClass = _ExtensionContextClass{objc.GetClass("NSExtensionContext")}

type _ExtensionContextClass struct {
	objc.Class
}

type IExtensionContext interface {
	objc.IObject
	CancelRequestWithError(error IError)
	CompleteRequestReturningItems_CompletionHandler(items []objc.IObject, completionHandler func(expired bool))
	OpenURL_CompletionHandler(URL IURL, completionHandler func(success bool))
	InputItems() []objc.Object
}

type ExtensionContext struct {
	objc.Object
}

func MakeExtensionContext(ptr unsafe.Pointer) ExtensionContext {
	return ExtensionContext{
		Object: objc.MakeObject(ptr),
	}
}

func (ec _ExtensionContextClass) Alloc() ExtensionContext {
	rv := objc.CallMethod[ExtensionContext](ec, "alloc")
	return rv
}

func (ec _ExtensionContextClass) New() ExtensionContext {
	rv := objc.CallMethod[ExtensionContext](ec, "new")
	rv.Autorelease()
	return rv
}

func NewExtensionContext() ExtensionContext {
	return ExtensionContextClass.New()
}

func (e_ ExtensionContext) Init() ExtensionContext {
	rv := objc.CallMethod[ExtensionContext](e_, "init")
	return rv
}

func (e_ ExtensionContext) CancelRequestWithError(error IError) {
	objc.CallMethod[objc.Void](e_, "cancelRequestWithError:", error)
}

func (e_ ExtensionContext) CompleteRequestReturningItems_CompletionHandler(items []objc.IObject, completionHandler func(expired bool)) {
	objc.CallMethod[objc.Void](e_, "completeRequestReturningItems:completionHandler:", items, completionHandler)
}

func (e_ ExtensionContext) OpenURL_CompletionHandler(URL IURL, completionHandler func(success bool)) {
	objc.CallMethod[objc.Void](e_, "openURL:completionHandler:", URL, completionHandler)
}

func (e_ ExtensionContext) InputItems() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](e_, "inputItems")
	return rv
}
