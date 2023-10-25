// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var HTTPCookieClass = _HTTPCookieClass{objc.GetClass("NSHTTPCookie")}

type _HTTPCookieClass struct {
	objc.Class
}

type IHTTPCookie interface {
	objc.IObject
	Domain() string
	Path() string
	PortList() []Number
	Name() string
	Value() string
	Version() uint
	ExpiresDate() Date
	IsSessionOnly() bool
	IsHTTPOnly() bool
	IsSecure() bool
	SameSitePolicy() HTTPCookieStringPolicy
	Properties() map[HTTPCookiePropertyKey]objc.Object
	Comment() string
	CommentURL() URL
}

type HTTPCookie struct {
	objc.Object
}

func MakeHTTPCookie(ptr unsafe.Pointer) HTTPCookie {
	return HTTPCookie{
		Object: objc.MakeObject(ptr),
	}
}

func (h_ HTTPCookie) InitWithProperties(properties map[HTTPCookiePropertyKey]objc.IObject) HTTPCookie {
	rv := objc.CallMethod[HTTPCookie](h_, objc.SEL("initWithProperties:"), properties)
	return rv
}

func (hc _HTTPCookieClass) Alloc() HTTPCookie {
	rv := objc.CallMethod[HTTPCookie](hc, objc.SEL("alloc"))
	return rv
}

func (hc _HTTPCookieClass) New() HTTPCookie {
	rv := objc.CallMethod[HTTPCookie](hc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewHTTPCookie() HTTPCookie {
	return HTTPCookieClass.New()
}

func (h_ HTTPCookie) Init() HTTPCookie {
	rv := objc.CallMethod[HTTPCookie](h_, objc.SEL("init"))
	return rv
}

func (hc _HTTPCookieClass) CookiesWithResponseHeaderFields_ForURL(headerFields map[string]string, URL IURL) []HTTPCookie {
	rv := objc.CallMethod[[]HTTPCookie](hc, objc.SEL("cookiesWithResponseHeaderFields:forURL:"), headerFields, objc.ExtractPtr(URL))
	return rv
}

func (hc _HTTPCookieClass) CookieWithProperties(properties map[HTTPCookiePropertyKey]objc.IObject) HTTPCookie {
	rv := objc.CallMethod[HTTPCookie](hc, objc.SEL("cookieWithProperties:"), properties)
	return rv
}

func (hc _HTTPCookieClass) RequestHeaderFieldsWithCookies(cookies []IHTTPCookie) map[string]string {
	rv := objc.CallMethod[map[string]string](hc, objc.SEL("requestHeaderFieldsWithCookies:"), cookies)
	return rv
}

func (h_ HTTPCookie) Domain() string {
	rv := objc.CallMethod[string](h_, objc.SEL("domain"))
	return rv
}

func (h_ HTTPCookie) Path() string {
	rv := objc.CallMethod[string](h_, objc.SEL("path"))
	return rv
}

func (h_ HTTPCookie) PortList() []Number {
	rv := objc.CallMethod[[]Number](h_, objc.SEL("portList"))
	return rv
}

func (h_ HTTPCookie) Name() string {
	rv := objc.CallMethod[string](h_, objc.SEL("name"))
	return rv
}

func (h_ HTTPCookie) Value() string {
	rv := objc.CallMethod[string](h_, objc.SEL("value"))
	return rv
}

func (h_ HTTPCookie) Version() uint {
	rv := objc.CallMethod[uint](h_, objc.SEL("version"))
	return rv
}

func (h_ HTTPCookie) ExpiresDate() Date {
	rv := objc.CallMethod[Date](h_, objc.SEL("expiresDate"))
	return rv
}

func (h_ HTTPCookie) IsSessionOnly() bool {
	rv := objc.CallMethod[bool](h_, objc.SEL("isSessionOnly"))
	return rv
}

func (h_ HTTPCookie) IsHTTPOnly() bool {
	rv := objc.CallMethod[bool](h_, objc.SEL("isHTTPOnly"))
	return rv
}

func (h_ HTTPCookie) IsSecure() bool {
	rv := objc.CallMethod[bool](h_, objc.SEL("isSecure"))
	return rv
}

func (h_ HTTPCookie) SameSitePolicy() HTTPCookieStringPolicy {
	rv := objc.CallMethod[HTTPCookieStringPolicy](h_, objc.SEL("sameSitePolicy"))
	return rv
}

func (h_ HTTPCookie) Properties() map[HTTPCookiePropertyKey]objc.Object {
	rv := objc.CallMethod[map[HTTPCookiePropertyKey]objc.Object](h_, objc.SEL("properties"))
	return rv
}

func (h_ HTTPCookie) Comment() string {
	rv := objc.CallMethod[string](h_, objc.SEL("comment"))
	return rv
}

func (h_ HTTPCookie) CommentURL() URL {
	rv := objc.CallMethod[URL](h_, objc.SEL("commentURL"))
	return rv
}
