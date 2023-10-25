// auto-generated code, do not modify
package foundation

import (
	"unsafe"

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
	rv := objc.CallMethod[URLRequest](uc, objc.SEL("requestWithURL:"), objc.ExtractPtr(URL))
	return rv
}

func (u_ URLRequest) InitWithURL(URL IURL) URLRequest {
	rv := objc.CallMethod[URLRequest](u_, objc.SEL("initWithURL:"), objc.ExtractPtr(URL))
	return rv
}

func (uc _URLRequestClass) RequestWithURL_CachePolicy_TimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) URLRequest {
	rv := objc.CallMethod[URLRequest](uc, objc.SEL("requestWithURL:cachePolicy:timeoutInterval:"), objc.ExtractPtr(URL), cachePolicy, timeoutInterval)
	return rv
}

func (u_ URLRequest) InitWithURL_CachePolicy_TimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) URLRequest {
	rv := objc.CallMethod[URLRequest](u_, objc.SEL("initWithURL:cachePolicy:timeoutInterval:"), objc.ExtractPtr(URL), cachePolicy, timeoutInterval)
	return rv
}

func (uc _URLRequestClass) Alloc() URLRequest {
	rv := objc.CallMethod[URLRequest](uc, objc.SEL("alloc"))
	return rv
}

func (uc _URLRequestClass) New() URLRequest {
	rv := objc.CallMethod[URLRequest](uc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewURLRequest() URLRequest {
	return URLRequestClass.New()
}

func (u_ URLRequest) Init() URLRequest {
	rv := objc.CallMethod[URLRequest](u_, objc.SEL("init"))
	return rv
}

func (u_ URLRequest) ValueForHTTPHeaderField(field string) string {
	rv := objc.CallMethod[string](u_, objc.SEL("valueForHTTPHeaderField:"), field)
	return rv
}

func (u_ URLRequest) CachePolicy() URLRequestCachePolicy {
	rv := objc.CallMethod[URLRequestCachePolicy](u_, objc.SEL("cachePolicy"))
	return rv
}

func (u_ URLRequest) HTTPMethod() string {
	rv := objc.CallMethod[string](u_, objc.SEL("HTTPMethod"))
	return rv
}

func (u_ URLRequest) URL() URL {
	rv := objc.CallMethod[URL](u_, objc.SEL("URL"))
	return rv
}

func (u_ URLRequest) HTTPBody() []byte {
	rv := objc.CallMethod[[]byte](u_, objc.SEL("HTTPBody"))
	return rv
}

func (u_ URLRequest) HTTPBodyStream() InputStream {
	rv := objc.CallMethod[InputStream](u_, objc.SEL("HTTPBodyStream"))
	return rv
}

func (u_ URLRequest) MainDocumentURL() URL {
	rv := objc.CallMethod[URL](u_, objc.SEL("mainDocumentURL"))
	return rv
}

func (u_ URLRequest) AllHTTPHeaderFields() map[string]string {
	rv := objc.CallMethod[map[string]string](u_, objc.SEL("allHTTPHeaderFields"))
	return rv
}

func (u_ URLRequest) TimeoutInterval() TimeInterval {
	rv := objc.CallMethod[TimeInterval](u_, objc.SEL("timeoutInterval"))
	return rv
}

func (u_ URLRequest) HTTPShouldHandleCookies() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("HTTPShouldHandleCookies"))
	return rv
}

func (u_ URLRequest) HTTPShouldUsePipelining() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("HTTPShouldUsePipelining"))
	return rv
}

func (u_ URLRequest) AllowsCellularAccess() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("allowsCellularAccess"))
	return rv
}

func (u_ URLRequest) AllowsConstrainedNetworkAccess() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("allowsConstrainedNetworkAccess"))
	return rv
}

func (u_ URLRequest) AllowsExpensiveNetworkAccess() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("allowsExpensiveNetworkAccess"))
	return rv
}

func (u_ URLRequest) NetworkServiceType() URLRequestNetworkServiceType {
	rv := objc.CallMethod[URLRequestNetworkServiceType](u_, objc.SEL("networkServiceType"))
	return rv
}

func (uc _URLRequestClass) SupportsSecureCoding() bool {
	rv := objc.CallMethod[bool](uc, objc.SEL("supportsSecureCoding"))
	return rv
}

func (u_ URLRequest) Attribution() URLRequestAttribution {
	rv := objc.CallMethod[URLRequestAttribution](u_, objc.SEL("attribution"))
	return rv
}

func (u_ URLRequest) AssumesHTTP3Capable() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("assumesHTTP3Capable"))
	return rv
}
