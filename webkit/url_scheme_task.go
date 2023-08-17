// auto-generated code, do not modify
package webkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type URLSchemeTask interface {
	// required
	DidReceiveResponse(response foundation.URLResponse)
	// required
	DidReceiveData(data []byte)
	// required
	DidFinish()
	// required
	DidFailWithError(error foundation.Error)
	ImplementsRequest() bool
	// optional
	Request() foundation.IURLRequest
}

func WrapURLSchemeTask(v URLSchemeTask) objc.Object {
	return objc.WrapAsProtocol("WKURLSchemeTask", v)
}

type URLSchemeTaskBase struct {
}

func (p *URLSchemeTaskBase) ImplementsRequest() bool {
	return false
}

func (p *URLSchemeTaskBase) Request() foundation.IURLRequest {
	panic("unimpemented protocol method")
}

type URLSchemeTaskWrapper struct {
	objc.Object
}

func (u_ URLSchemeTaskWrapper) DidReceiveResponse(response foundation.IURLResponse) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("didReceiveResponse:"), objc.ExtractPtr(response))
}

func (u_ URLSchemeTaskWrapper) DidReceiveData(data []byte) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("didReceiveData:"), data)
}

func (u_ URLSchemeTaskWrapper) DidFinish() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("didFinish"))
}

func (u_ URLSchemeTaskWrapper) DidFailWithError(error foundation.IError) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("didFailWithError:"), objc.ExtractPtr(error))
}

func (u_ URLSchemeTaskWrapper) Request() foundation.URLRequest {
	rv := objc.CallMethod[foundation.URLRequest](u_, objc.GetSelector("request"))
	return rv
}
