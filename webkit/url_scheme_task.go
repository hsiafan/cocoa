// auto-generated code, do not modify
package webkit

import (
	"github.com/hsiafan/cocoa/ffi"
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

type URLSchemeTaskWrapper struct {
	objc.Object
}

func (u_ URLSchemeTaskWrapper) DidReceiveResponse(response foundation.IURLResponse) {
	ffi.CallMethod[ffi.Void](u_, "didReceiveResponse:", response)
}

func (u_ URLSchemeTaskWrapper) DidReceiveData(data []byte) {
	ffi.CallMethod[ffi.Void](u_, "didReceiveData:", data)
}

func (u_ URLSchemeTaskWrapper) DidFinish() {
	ffi.CallMethod[ffi.Void](u_, "didFinish")
}

func (u_ URLSchemeTaskWrapper) DidFailWithError(error foundation.IError) {
	ffi.CallMethod[ffi.Void](u_, "didFailWithError:", error)
}

func (u_ *URLSchemeTaskWrapper) ImplementsRequest() bool {
	return u_.RespondsToSelector(objc.GetSelector("request"))
}

func (u_ URLSchemeTaskWrapper) Request() foundation.URLRequest {
	rv := ffi.CallMethod[foundation.URLRequest](u_, "request")
	return rv
}
