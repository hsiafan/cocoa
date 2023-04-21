// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var DownloadClass = _DownloadClass{objc.GetClass("WKDownload")}

type _DownloadClass struct {
	objc.Class
}

type IDownload interface {
	objc.IObject
	Cancel(completionHandler func(resumeData []byte))
	OriginalRequest() foundation.URLRequest
	WebView() WebView
}

type Download struct {
	objc.Object
}

func MakeDownload(ptr unsafe.Pointer) Download {
	return Download{
		Object: objc.MakeObject(ptr),
	}
}

func (dc _DownloadClass) Alloc() Download {
	rv := objc.CallMethod[Download](dc, "alloc")
	return rv
}

func (dc _DownloadClass) New() Download {
	rv := objc.CallMethod[Download](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDownload() Download {
	return DownloadClass.New()
}

func (d_ Download) Init() Download {
	rv := objc.CallMethod[Download](d_, "init")
	return rv
}

func (d_ Download) Cancel(completionHandler func(resumeData []byte)) {
	objc.CallMethod[objc.Void](d_, "cancel:", completionHandler)
}

func (d_ Download) OriginalRequest() foundation.URLRequest {
	rv := objc.CallMethod[foundation.URLRequest](d_, "originalRequest")
	return rv
}

func (d_ Download) WebView() WebView {
	rv := objc.CallMethod[WebView](d_, "webView")
	return rv
}
