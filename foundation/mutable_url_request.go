// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[MutableURLRequest](mc, "requestWithURL:", URL)
	return rv
}

func (m_ MutableURLRequest) InitWithURL(URL IURL) MutableURLRequest {
	rv := ffi.CallMethod[MutableURLRequest](m_, "initWithURL:", URL)
	return rv
}

func (mc _MutableURLRequestClass) RequestWithURL_CachePolicy_TimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) MutableURLRequest {
	rv := ffi.CallMethod[MutableURLRequest](mc, "requestWithURL:cachePolicy:timeoutInterval:", URL, cachePolicy, timeoutInterval)
	return rv
}

func (m_ MutableURLRequest) InitWithURL_CachePolicy_TimeoutInterval(URL IURL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) MutableURLRequest {
	rv := ffi.CallMethod[MutableURLRequest](m_, "initWithURL:cachePolicy:timeoutInterval:", URL, cachePolicy, timeoutInterval)
	return rv
}

func (mc _MutableURLRequestClass) Alloc() MutableURLRequest {
	rv := ffi.CallMethod[MutableURLRequest](mc, "alloc")
	return rv
}

func (mc _MutableURLRequestClass) New() MutableURLRequest {
	rv := ffi.CallMethod[MutableURLRequest](mc, "new")
	rv.Autorelease()
	return rv
}

func NewMutableURLRequest() MutableURLRequest {
	return MutableURLRequestClass.New()
}

func (m_ MutableURLRequest) Init() MutableURLRequest {
	rv := ffi.CallMethod[MutableURLRequest](m_, "init")
	return rv
}

func (m_ MutableURLRequest) AddValue_ForHTTPHeaderField(value string, field string) {
	ffi.CallMethod[ffi.Void](m_, "addValue:forHTTPHeaderField:", value, field)
}

func (m_ MutableURLRequest) SetValue_ForHTTPHeaderField(value string, field string) {
	ffi.CallMethod[ffi.Void](m_, "setValue:forHTTPHeaderField:", value, field)
}

func (m_ MutableURLRequest) SetCachePolicy(value URLRequestCachePolicy) {
	ffi.CallMethod[ffi.Void](m_, "setCachePolicy:", value)
}

func (m_ MutableURLRequest) SetHTTPMethod(value string) {
	ffi.CallMethod[ffi.Void](m_, "setHTTPMethod:", value)
}

func (m_ MutableURLRequest) SetURL(value IURL) {
	ffi.CallMethod[ffi.Void](m_, "setURL:", value)
}

func (m_ MutableURLRequest) SetHTTPBody(value []byte) {
	ffi.CallMethod[ffi.Void](m_, "setHTTPBody:", value)
}

func (m_ MutableURLRequest) SetHTTPBodyStream(value IInputStream) {
	ffi.CallMethod[ffi.Void](m_, "setHTTPBodyStream:", value)
}

func (m_ MutableURLRequest) SetMainDocumentURL(value IURL) {
	ffi.CallMethod[ffi.Void](m_, "setMainDocumentURL:", value)
}

func (m_ MutableURLRequest) SetAllHTTPHeaderFields(value map[string]string) {
	ffi.CallMethod[ffi.Void](m_, "setAllHTTPHeaderFields:", value)
}

func (m_ MutableURLRequest) SetTimeoutInterval(value TimeInterval) {
	ffi.CallMethod[ffi.Void](m_, "setTimeoutInterval:", value)
}

func (m_ MutableURLRequest) SetHTTPShouldHandleCookies(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setHTTPShouldHandleCookies:", value)
}

func (m_ MutableURLRequest) SetHTTPShouldUsePipelining(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setHTTPShouldUsePipelining:", value)
}

func (m_ MutableURLRequest) SetAllowsCellularAccess(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setAllowsCellularAccess:", value)
}

func (m_ MutableURLRequest) SetAllowsConstrainedNetworkAccess(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setAllowsConstrainedNetworkAccess:", value)
}

func (m_ MutableURLRequest) SetAllowsExpensiveNetworkAccess(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setAllowsExpensiveNetworkAccess:", value)
}

func (m_ MutableURLRequest) SetNetworkServiceType(value URLRequestNetworkServiceType) {
	ffi.CallMethod[ffi.Void](m_, "setNetworkServiceType:", value)
}

func (m_ MutableURLRequest) SetAttribution(value URLRequestAttribution) {
	ffi.CallMethod[ffi.Void](m_, "setAttribution:", value)
}

func (m_ MutableURLRequest) SetAssumesHTTP3Capable(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setAssumesHTTP3Capable:", value)
}

func (m_ MutableURLRequest) RequiresDNSSECValidation() bool {
	rv := ffi.CallMethod[bool](m_, "requiresDNSSECValidation")
	return rv
}

func (m_ MutableURLRequest) SetRequiresDNSSECValidation(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setRequiresDNSSECValidation:", value)
}
