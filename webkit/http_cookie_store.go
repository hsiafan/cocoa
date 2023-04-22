// auto-generated code, do not modify
package webkit

import (
	"unsafe"

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
	rv := objc.CallMethod[HTTPCookieStore](hc, objc.GetSelector("alloc"))
	return rv
}

func (hc _HTTPCookieStoreClass) New() HTTPCookieStore {
	rv := objc.CallMethod[HTTPCookieStore](hc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewHTTPCookieStore() HTTPCookieStore {
	return HTTPCookieStoreClass.New()
}

func (h_ HTTPCookieStore) Init() HTTPCookieStore {
	rv := objc.CallMethod[HTTPCookieStore](h_, objc.GetSelector("init"))
	return rv
}

func (h_ HTTPCookieStore) GetAllCookies(completionHandler func(param1 []foundation.HTTPCookie)) {
	objc.CallMethod[objc.Void](h_, objc.GetSelector("getAllCookies:"), completionHandler)
}

func (h_ HTTPCookieStore) SetCookie_CompletionHandler(cookie foundation.IHTTPCookie, completionHandler func()) {
	objc.CallMethod[objc.Void](h_, objc.GetSelector("setCookie:completionHandler:"), cookie, completionHandler)
}

func (h_ HTTPCookieStore) DeleteCookie_CompletionHandler(cookie foundation.IHTTPCookie, completionHandler func()) {
	objc.CallMethod[objc.Void](h_, objc.GetSelector("deleteCookie:completionHandler:"), cookie, completionHandler)
}
