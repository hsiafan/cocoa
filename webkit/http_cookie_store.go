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
	rv := objc.CallMethod[HTTPCookieStore](hc, objc.SEL("alloc"))
	return rv
}

func (hc _HTTPCookieStoreClass) New() HTTPCookieStore {
	rv := objc.CallMethod[HTTPCookieStore](hc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewHTTPCookieStore() HTTPCookieStore {
	return HTTPCookieStoreClass.New()
}

func (h_ HTTPCookieStore) Init() HTTPCookieStore {
	rv := objc.CallMethod[HTTPCookieStore](h_, objc.SEL("init"))
	return rv
}

func (h_ HTTPCookieStore) GetAllCookies(completionHandler func(param1 []foundation.HTTPCookie)) {
	objc.CallMethod[objc.Void](h_, objc.SEL("getAllCookies:"), completionHandler)
}

func (h_ HTTPCookieStore) SetCookie_CompletionHandler(cookie foundation.IHTTPCookie, completionHandler func()) {
	objc.CallMethod[objc.Void](h_, objc.SEL("setCookie:completionHandler:"), objc.ExtractPtr(cookie), completionHandler)
}

func (h_ HTTPCookieStore) DeleteCookie_CompletionHandler(cookie foundation.IHTTPCookie, completionHandler func()) {
	objc.CallMethod[objc.Void](h_, objc.SEL("deleteCookie:completionHandler:"), objc.ExtractPtr(cookie), completionHandler)
}
