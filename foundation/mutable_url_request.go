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
	rv := objc.CallMethod[MutableURLRequest](mc, objc.GetSelector("requestWithURL:"), URL)
	return rv
}

func (m_ MutableURLRequest) InitWithURL(URL IURL) MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](m_, objc.GetSelector("initWithURL:"), URL)
	return rv
}

func (mc _MutableURLRequestClass) RequestWithURL_CachePolicy_TimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](mc, objc.GetSelector("requestWithURL:cachePolicy:timeoutInterval:"), URL, cachePolicy, timeoutInterval)
	return rv
}

func (m_ MutableURLRequest) InitWithURL_CachePolicy_TimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](m_, objc.GetSelector("initWithURL:cachePolicy:timeoutInterval:"), URL, cachePolicy, timeoutInterval)
	return rv
}

func (mc _MutableURLRequestClass) Alloc() MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](mc, objc.GetSelector("alloc"))
	return rv
}

func (mc _MutableURLRequestClass) New() MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](mc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewMutableURLRequest() MutableURLRequest {
	return MutableURLRequestClass.New()
}

func (m_ MutableURLRequest) Init() MutableURLRequest {
	rv := objc.CallMethod[MutableURLRequest](m_, objc.GetSelector("init"))
	return rv
}

func (m_ MutableURLRequest) AddValue_ForHTTPHeaderField(value string, field string) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("addValue:forHTTPHeaderField:"), value, field)
}

func (m_ MutableURLRequest) SetValue_ForHTTPHeaderField(value string, field string) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setValue:forHTTPHeaderField:"), value, field)
}

func (m_ MutableURLRequest) SetCachePolicy(value URLRequestCachePolicy) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setCachePolicy:"), value)
}

func (m_ MutableURLRequest) SetHTTPMethod(value string) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setHTTPMethod:"), value)
}

func (m_ MutableURLRequest) SetURL(value IURL) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setURL:"), value)
}

func (m_ MutableURLRequest) SetHTTPBody(value []byte) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setHTTPBody:"), value)
}

func (m_ MutableURLRequest) SetHTTPBodyStream(value IInputStream) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setHTTPBodyStream:"), value)
}

func (m_ MutableURLRequest) SetMainDocumentURL(value IURL) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setMainDocumentURL:"), value)
}

func (m_ MutableURLRequest) SetAllHTTPHeaderFields(value map[string]string) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAllHTTPHeaderFields:"), value)
}

func (m_ MutableURLRequest) SetTimeoutInterval(value TimeInterval) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setTimeoutInterval:"), value)
}

func (m_ MutableURLRequest) SetHTTPShouldHandleCookies(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setHTTPShouldHandleCookies:"), value)
}

func (m_ MutableURLRequest) SetHTTPShouldUsePipelining(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setHTTPShouldUsePipelining:"), value)
}

func (m_ MutableURLRequest) SetAllowsCellularAccess(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAllowsCellularAccess:"), value)
}

func (m_ MutableURLRequest) SetAllowsConstrainedNetworkAccess(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAllowsConstrainedNetworkAccess:"), value)
}

func (m_ MutableURLRequest) SetAllowsExpensiveNetworkAccess(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAllowsExpensiveNetworkAccess:"), value)
}

func (m_ MutableURLRequest) SetNetworkServiceType(value URLRequestNetworkServiceType) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setNetworkServiceType:"), value)
}

func (m_ MutableURLRequest) SetAttribution(value URLRequestAttribution) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAttribution:"), value)
}

func (m_ MutableURLRequest) SetAssumesHTTP3Capable(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAssumesHTTP3Capable:"), value)
}

func (m_ MutableURLRequest) RequiresDNSSECValidation() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("requiresDNSSECValidation"))
	return rv
}

func (m_ MutableURLRequest) SetRequiresDNSSECValidation(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setRequiresDNSSECValidation:"), value)
}
