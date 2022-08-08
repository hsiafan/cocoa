// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[ExtensionContext](ec, "alloc")
	return rv
}

func (e_ ExtensionContext) Init() ExtensionContext {
	rv := ffi.CallMethod[ExtensionContext](e_, "init")
	return rv
}

func (ec _ExtensionContextClass) New() ExtensionContext {
	rv := ffi.CallMethod[ExtensionContext](ec, "new")
	rv.Autorelease()
	return rv
}

func NewExtensionContext() ExtensionContext {
	return ExtensionContextClass.New()
}

func (e_ ExtensionContext) CancelRequestWithError(error IError) {
	ffi.CallMethod[ffi.Void](e_, "cancelRequestWithError:", error)
}

func (e_ ExtensionContext) CompleteRequestReturningItems_CompletionHandler(items []objc.IObject, completionHandler func(expired bool)) {
	ffi.CallMethod[ffi.Void](e_, "completeRequestReturningItems:completionHandler:", items, completionHandler)
}

func (e_ ExtensionContext) OpenURL_CompletionHandler(URL IURL, completionHandler func(success bool)) {
	ffi.CallMethod[ffi.Void](e_, "openURL:completionHandler:", URL, completionHandler)
}

func (e_ ExtensionContext) InputItems() []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](e_, "inputItems")
	return rv
}
