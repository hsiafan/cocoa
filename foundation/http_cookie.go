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
	rv := objc.CallMethod[HTTPCookie](h_, "initWithProperties:", properties)
	return rv
}

func (hc _HTTPCookieClass) Alloc() HTTPCookie {
	rv := objc.CallMethod[HTTPCookie](hc, "alloc")
	return rv
}

func (hc _HTTPCookieClass) New() HTTPCookie {
	rv := objc.CallMethod[HTTPCookie](hc, "new")
	rv.Autorelease()
	return rv
}

func NewHTTPCookie() HTTPCookie {
	return HTTPCookieClass.New()
}

func (h_ HTTPCookie) Init() HTTPCookie {
	rv := objc.CallMethod[HTTPCookie](h_, "init")
	return rv
}

func (hc _HTTPCookieClass) CookiesWithResponseHeaderFields_ForURL(headerFields map[string]string, URL IURL) []HTTPCookie {
	rv := objc.CallMethod[[]HTTPCookie](hc, "cookiesWithResponseHeaderFields:forURL:", headerFields, URL)
	return rv
}

func (hc _HTTPCookieClass) CookieWithProperties(properties map[HTTPCookiePropertyKey]objc.IObject) HTTPCookie {
	rv := objc.CallMethod[HTTPCookie](hc, "cookieWithProperties:", properties)
	return rv
}

func (hc _HTTPCookieClass) RequestHeaderFieldsWithCookies(cookies []IHTTPCookie) map[string]string {
	rv := objc.CallMethod[map[string]string](hc, "requestHeaderFieldsWithCookies:", cookies)
	return rv
}

func (h_ HTTPCookie) Domain() string {
	rv := objc.CallMethod[string](h_, "domain")
	return rv
}

func (h_ HTTPCookie) Path() string {
	rv := objc.CallMethod[string](h_, "path")
	return rv
}

func (h_ HTTPCookie) PortList() []Number {
	rv := objc.CallMethod[[]Number](h_, "portList")
	return rv
}

func (h_ HTTPCookie) Name() string {
	rv := objc.CallMethod[string](h_, "name")
	return rv
}

func (h_ HTTPCookie) Value() string {
	rv := objc.CallMethod[string](h_, "value")
	return rv
}

func (h_ HTTPCookie) Version() uint {
	rv := objc.CallMethod[uint](h_, "version")
	return rv
}

func (h_ HTTPCookie) ExpiresDate() Date {
	rv := objc.CallMethod[Date](h_, "expiresDate")
	return rv
}

func (h_ HTTPCookie) IsSessionOnly() bool {
	rv := objc.CallMethod[bool](h_, "isSessionOnly")
	return rv
}

func (h_ HTTPCookie) IsHTTPOnly() bool {
	rv := objc.CallMethod[bool](h_, "isHTTPOnly")
	return rv
}

func (h_ HTTPCookie) IsSecure() bool {
	rv := objc.CallMethod[bool](h_, "isSecure")
	return rv
}

func (h_ HTTPCookie) SameSitePolicy() HTTPCookieStringPolicy {
	rv := objc.CallMethod[HTTPCookieStringPolicy](h_, "sameSitePolicy")
	return rv
}

func (h_ HTTPCookie) Properties() map[HTTPCookiePropertyKey]objc.Object {
	rv := objc.CallMethod[map[HTTPCookiePropertyKey]objc.Object](h_, "properties")
	return rv
}

func (h_ HTTPCookie) Comment() string {
	rv := objc.CallMethod[string](h_, "comment")
	return rv
}

func (h_ HTTPCookie) CommentURL() URL {
	rv := objc.CallMethod[URL](h_, "commentURL")
	return rv
}
