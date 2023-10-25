// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var FrameInfoClass = _FrameInfoClass{objc.GetClass("WKFrameInfo")}

type _FrameInfoClass struct {
	objc.Class
}

type IFrameInfo interface {
	objc.IObject
	IsMainFrame() bool
	Request() foundation.URLRequest
	SecurityOrigin() SecurityOrigin
	WebView() WebView
}

type FrameInfo struct {
	objc.Object
}

func MakeFrameInfo(ptr unsafe.Pointer) FrameInfo {
	return FrameInfo{
		Object: objc.MakeObject(ptr),
	}
}

func (fc _FrameInfoClass) Alloc() FrameInfo {
	rv := objc.CallMethod[FrameInfo](fc, objc.SEL("alloc"))
	return rv
}

func (fc _FrameInfoClass) New() FrameInfo {
	rv := objc.CallMethod[FrameInfo](fc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewFrameInfo() FrameInfo {
	return FrameInfoClass.New()
}

func (f_ FrameInfo) Init() FrameInfo {
	rv := objc.CallMethod[FrameInfo](f_, objc.SEL("init"))
	return rv
}

func (f_ FrameInfo) IsMainFrame() bool {
	rv := objc.CallMethod[bool](f_, objc.SEL("isMainFrame"))
	return rv
}

func (f_ FrameInfo) Request() foundation.URLRequest {
	rv := objc.CallMethod[foundation.URLRequest](f_, objc.SEL("request"))
	return rv
}

func (f_ FrameInfo) SecurityOrigin() SecurityOrigin {
	rv := objc.CallMethod[SecurityOrigin](f_, objc.SEL("securityOrigin"))
	return rv
}

// weak property
func (f_ FrameInfo) WebView() WebView {
	rv := objc.CallMethod[WebView](f_, objc.SEL("webView"))
	return rv
}
