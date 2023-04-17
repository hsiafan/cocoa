// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var HTTPCookieStoreClass = _HTTPCookieStoreClass{objc.GetClass("WKHTTPCookieStore")}

type _HTTPCookieStoreClass struct {
	objc.Class
}

type IHTTPCookieStore interface {
	objc.IObject
	GetAllCookies(completionHandler func(param1 []foundation.HTTPCookie))
	SetCookie_CompletionHandler(cookie foundation.IHTTPCookie, completionHandler func())
	DeleteCookie_CompletionHandler(cookie foundation.IHTTPCookie, completionHandler func())
}

type HTTPCookieStore struct {
	objc.Object
}

func MakeHTTPCookieStore(ptr unsafe.Pointer) HTTPCookieStore {
	return HTTPCookieStore{
		Object: objc.MakeObject(ptr),
	}
}

func (hc _HTTPCookieStoreClass) Alloc() HTTPCookieStore {
	rv := ffi.CallMethod[HTTPCookieStore](hc, "alloc")
	return rv
}

func (hc _HTTPCookieStoreClass) New() HTTPCookieStore {
	rv := ffi.CallMethod[HTTPCookieStore](hc, "new")
	rv.Autorelease()
	return rv
}

func NewHTTPCookieStore() HTTPCookieStore {
	return HTTPCookieStoreClass.New()
}

func (h_ HTTPCookieStore) Init() HTTPCookieStore {
	rv := ffi.CallMethod[HTTPCookieStore](h_, "init")
	return rv
}

func (h_ HTTPCookieStore) GetAllCookies(completionHandler func(param1 []foundation.HTTPCookie)) {
	ffi.CallMethod[ffi.Void](h_, "getAllCookies:", completionHandler)
}

func (h_ HTTPCookieStore) SetCookie_CompletionHandler(cookie foundation.IHTTPCookie, completionHandler func()) {
	ffi.CallMethod[ffi.Void](h_, "setCookie:completionHandler:", cookie, completionHandler)
}

func (h_ HTTPCookieStore) DeleteCookie_CompletionHandler(cookie foundation.IHTTPCookie, completionHandler func()) {
	ffi.CallMethod[ffi.Void](h_, "deleteCookie:completionHandler:", cookie, completionHandler)
}
