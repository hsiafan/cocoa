// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var URLRequestClass = _URLRequestClass{objc.GetClass("NSURLRequest")}

type _URLRequestClass struct {
	objc.Class
}

type IURLRequest interface {
	objc.IObject
	ValueForHTTPHeaderField(field string) string
	CachePolicy() URLRequestCachePolicy
	HTTPMethod() string
	URL() URL
	HTTPBody() []byte
	HTTPBodyStream() InputStream
	MainDocumentURL() URL
	AllHTTPHeaderFields() map[string]string
	TimeoutInterval() TimeInterval
	HTTPShouldHandleCookies() bool
	HTTPShouldUsePipelining() bool
	AllowsCellularAccess() bool
	AllowsConstrainedNetworkAccess() bool
	AllowsExpensiveNetworkAccess() bool
	NetworkServiceType() URLRequestNetworkServiceType
	Attribution() URLRequestAttribution
	AssumesHTTP3Capable() bool
}

type URLRequest struct {
	objc.Object
}

func MakeURLRequest(ptr unsafe.Pointer) URLRequest {
	return URLRequest{
		Object: objc.MakeObject(ptr),
	}
}

func (uc _URLRequestClass) RequestWithURL(URL IURL) URLRequest {
	rv := ffi.CallMethod[URLRequest](uc, "requestWithURL:", URL)
	return rv
}

func (u_ URLRequest) InitWithURL(URL IURL) URLRequest {
	rv := ffi.CallMethod[URLRequest](u_, "initWithURL:", URL)
	return rv
}

func (uc _URLRequestClass) RequestWithURL_CachePolicy_TimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) URLRequest {
	rv := ffi.CallMethod[URLRequest](uc, "requestWithURL:cachePolicy:timeoutInterval:", URL, cachePolicy, timeoutInterval)
	return rv
}

func (u_ URLRequest) InitWithURL_CachePolicy_TimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) URLRequest {
	rv := ffi.CallMethod[URLRequest](u_, "initWithURL:cachePolicy:timeoutInterval:", URL, cachePolicy, timeoutInterval)
	return rv
}

func (uc _URLRequestClass) Alloc() URLRequest {
	rv := ffi.CallMethod[URLRequest](uc, "alloc")
	return rv
}

func (uc _URLRequestClass) New() URLRequest {
	rv := ffi.CallMethod[URLRequest](uc, "new")
	rv.Autorelease()
	return rv
}

func NewURLRequest() URLRequest {
	return URLRequestClass.New()
}

func (u_ URLRequest) Init() URLRequest {
	rv := ffi.CallMethod[URLRequest](u_, "init")
	return rv
}

func (u_ URLRequest) ValueForHTTPHeaderField(field string) string {
	rv := ffi.CallMethod[string](u_, "valueForHTTPHeaderField:", field)
	return rv
}

func (u_ URLRequest) CachePolicy() URLRequestCachePolicy {
	rv := ffi.CallMethod[URLRequestCachePolicy](u_, "cachePolicy")
	return rv
}

func (u_ URLRequest) HTTPMethod() string {
	rv := ffi.CallMethod[string](u_, "HTTPMethod")
	return rv
}

func (u_ URLRequest) URL() URL {
	rv := ffi.CallMethod[URL](u_, "URL")
	return rv
}

func (u_ URLRequest) HTTPBody() []byte {
	rv := ffi.CallMethod[[]byte](u_, "HTTPBody")
	return rv
}

func (u_ URLRequest) HTTPBodyStream() InputStream {
	rv := ffi.CallMethod[InputStream](u_, "HTTPBodyStream")
	return rv
}

func (u_ URLRequest) MainDocumentURL() URL {
	rv := ffi.CallMethod[URL](u_, "mainDocumentURL")
	return rv
}

func (u_ URLRequest) AllHTTPHeaderFields() map[string]string {
	rv := ffi.CallMethod[map[string]string](u_, "allHTTPHeaderFields")
	return rv
}

func (u_ URLRequest) TimeoutInterval() TimeInterval {
	rv := ffi.CallMethod[TimeInterval](u_, "timeoutInterval")
	return rv
}

func (u_ URLRequest) HTTPShouldHandleCookies() bool {
	rv := ffi.CallMethod[bool](u_, "HTTPShouldHandleCookies")
	return rv
}

func (u_ URLRequest) HTTPShouldUsePipelining() bool {
	rv := ffi.CallMethod[bool](u_, "HTTPShouldUsePipelining")
	return rv
}

func (u_ URLRequest) AllowsCellularAccess() bool {
	rv := ffi.CallMethod[bool](u_, "allowsCellularAccess")
	return rv
}

func (u_ URLRequest) AllowsConstrainedNetworkAccess() bool {
	rv := ffi.CallMethod[bool](u_, "allowsConstrainedNetworkAccess")
	return rv
}

func (u_ URLRequest) AllowsExpensiveNetworkAccess() bool {
	rv := ffi.CallMethod[bool](u_, "allowsExpensiveNetworkAccess")
	return rv
}

func (u_ URLRequest) NetworkServiceType() URLRequestNetworkServiceType {
	rv := ffi.CallMethod[URLRequestNetworkServiceType](u_, "networkServiceType")
	return rv
}

func (uc _URLRequestClass) SupportsSecureCoding() bool {
	rv := ffi.CallMethod[bool](uc, "supportsSecureCoding")
	return rv
}

func (u_ URLRequest) Attribution() URLRequestAttribution {
	rv := ffi.CallMethod[URLRequestAttribution](u_, "attribution")
	return rv
}

func (u_ URLRequest) AssumesHTTP3Capable() bool {
	rv := ffi.CallMethod[bool](u_, "assumesHTTP3Capable")
	return rv
}
