// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var MutableURLRequestClass = _MutableURLRequestClass{objc.GetClass("NSMutableURLRequest")}

type _MutableURLRequestClass struct {
	objc.Class
}

type IMutableURLRequest interface {
	IURLRequest
	AddValue_ForHTTPHeaderField(value string, field string)
	SetValue_ForHTTPHeaderField(value string, field string)
	SetCachePolicy(value URLRequestCachePolicy)
	SetHTTPMethod(value string)
	SetURL(value IURL)
	SetHTTPBody(value []byte)
	SetHTTPBodyStream(value IInputStream)
	SetMainDocumentURL(value IURL)
	SetAllHTTPHeaderFields(value map[string]string)
	SetTimeoutInterval(value TimeInterval)
	SetHTTPShouldHandleCookies(value bool)
	SetHTTPShouldUsePipelining(value bool)
	SetAllowsCellularAccess(value bool)
	SetAllowsConstrainedNetworkAccess(value bool)
	SetAllowsExpensiveNetworkAccess(value bool)
	SetNetworkServiceType(value URLRequestNetworkServiceType)
	SetAttribution(value URLRequestAttribution)
	SetAssumesHTTP3Capable(value bool)
	RequiresDNSSECValidation() bool
	SetRequiresDNSSECValidation(value bool)
}

type MutableURLRequest struct {
	URLRequest
}

func MakeMutableURLRequest(ptr unsafe.Pointer) MutableURLRequest {
	return MutableURLRequest{
		URLRequest: MakeURLRequest(ptr),
	}
}

func (mc _MutableURLRequestClass) RequestWithURL(URL IURL) MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](mc, objc.SEL("requestWithURL:"), objc.ExtractPtr(URL))
	return rv
}

func (m_ MutableURLRequest) InitWithURL(URL IURL) MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](m_, objc.SEL("initWithURL:"), objc.ExtractPtr(URL))
	return rv
}

func (mc _MutableURLRequestClass) RequestWithURL_CachePolicy_TimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](mc, objc.SEL("requestWithURL:cachePolicy:timeoutInterval:"), objc.ExtractPtr(URL), cachePolicy, timeoutInterval)
	return rv
}

func (m_ MutableURLRequest) InitWithURL_CachePolicy_TimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](m_, objc.SEL("initWithURL:cachePolicy:timeoutInterval:"), objc.ExtractPtr(URL), cachePolicy, timeoutInterval)
	return rv
}

func (mc _MutableURLRequestClass) Alloc() MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](mc, objc.SEL("alloc"))
	return rv
}

func (mc _MutableURLRequestClass) New() MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](mc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewMutableURLRequest() MutableURLRequest {
	return MutableURLRequestClass.New()
}

func (m_ MutableURLRequest) Init() MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](m_, objc.SEL("init"))
	return rv
}

func (m_ MutableURLRequest) AddValue_ForHTTPHeaderField(value string, field string) {
	objc.CallMethod[objc.Void](m_, objc.SEL("addValue:forHTTPHeaderField:"), value, field)
}

func (m_ MutableURLRequest) SetValue_ForHTTPHeaderField(value string, field string) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setValue:forHTTPHeaderField:"), value, field)
}

func (m_ MutableURLRequest) SetCachePolicy(value URLRequestCachePolicy) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setCachePolicy:"), value)
}

func (m_ MutableURLRequest) SetHTTPMethod(value string) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setHTTPMethod:"), value)
}

func (m_ MutableURLRequest) SetURL(value IURL) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setURL:"), objc.ExtractPtr(value))
}

func (m_ MutableURLRequest) SetHTTPBody(value []byte) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setHTTPBody:"), value)
}

func (m_ MutableURLRequest) SetHTTPBodyStream(value IInputStream) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setHTTPBodyStream:"), objc.ExtractPtr(value))
}

func (m_ MutableURLRequest) SetMainDocumentURL(value IURL) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setMainDocumentURL:"), objc.ExtractPtr(value))
}

func (m_ MutableURLRequest) SetAllHTTPHeaderFields(value map[string]string) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setAllHTTPHeaderFields:"), value)
}

func (m_ MutableURLRequest) SetTimeoutInterval(value TimeInterval) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setTimeoutInterval:"), value)
}

func (m_ MutableURLRequest) SetHTTPShouldHandleCookies(value bool) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setHTTPShouldHandleCookies:"), value)
}

func (m_ MutableURLRequest) SetHTTPShouldUsePipelining(value bool) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setHTTPShouldUsePipelining:"), value)
}

func (m_ MutableURLRequest) SetAllowsCellularAccess(value bool) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setAllowsCellularAccess:"), value)
}

func (m_ MutableURLRequest) SetAllowsConstrainedNetworkAccess(value bool) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setAllowsConstrainedNetworkAccess:"), value)
}

func (m_ MutableURLRequest) SetAllowsExpensiveNetworkAccess(value bool) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setAllowsExpensiveNetworkAccess:"), value)
}

func (m_ MutableURLRequest) SetNetworkServiceType(value URLRequestNetworkServiceType) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setNetworkServiceType:"), value)
}

func (m_ MutableURLRequest) SetAttribution(value URLRequestAttribution) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setAttribution:"), value)
}

func (m_ MutableURLRequest) SetAssumesHTTP3Capable(value bool) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setAssumesHTTP3Capable:"), value)
}

func (m_ MutableURLRequest) RequiresDNSSECValidation() bool {
	rv := objc.CallMethod[bool](m_, objc.SEL("requiresDNSSECValidation"))
	return rv
}

func (m_ MutableURLRequest) SetRequiresDNSSECValidation(value bool) {
	objc.CallMethod[objc.Void](m_, objc.SEL("setRequiresDNSSECValidation:"), value)
}
