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
	objc.CallMethod[objc.Void](u_, objc.SEL("didReceiveResponse:"), objc.ExtractPtr(response))
}

func (u_ URLSchemeTask) DidReceiveData(data []byte) {
	objc.CallMethod[objc.Void](u_, objc.SEL("didReceiveData:"), data)
}

func (u_ URLSchemeTask) DidFinish() {
	objc.CallMethod[objc.Void](u_, objc.SEL("didFinish"))
}

func (u_ URLSchemeTask) DidFailWithError(error foundation.IError) {
	objc.CallMethod[objc.Void](u_, objc.SEL("didFailWithError:"), objc.ExtractPtr(error))
}

func (u_ URLSchemeTask) Request() foundation.URLRequest {
	rv := objc.CallMethod[foundation.URLRequest](u_, objc.SEL("request"))
	return rv
}
