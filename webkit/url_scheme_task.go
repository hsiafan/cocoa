// auto-generated code, do not modify
package webkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type URLSchemeTask struct {
	objc.Object
}

func (u_ URLSchemeTask) DidReceiveResponse(response foundation.IURLResponse) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("didReceiveResponse:"), objc.ExtractPtr(response))
}

func (u_ URLSchemeTask) DidReceiveData(data []byte) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("didReceiveData:"), data)
}

func (u_ URLSchemeTask) DidFinish() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("didFinish"))
}

func (u_ URLSchemeTask) DidFailWithError(error foundation.IError) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("didFailWithError:"), objc.ExtractPtr(error))
}

func (u_ URLSchemeTask) Request() foundation.URLRequest {
	rv := objc.CallMethod[foundation.URLRequest](u_, objc.GetSelector("request"))
	return rv
}
