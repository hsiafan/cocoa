// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[FrameInfo](fc, "alloc")
	return rv
}

func (fc _FrameInfoClass) New() FrameInfo {
	rv := ffi.CallMethod[FrameInfo](fc, "new")
	rv.Autorelease()
	return rv
}

func NewFrameInfo() FrameInfo {
	return FrameInfoClass.New()
}

func (f_ FrameInfo) Init() FrameInfo {
	rv := ffi.CallMethod[FrameInfo](f_, "init")
	return rv
}

func (f_ FrameInfo) IsMainFrame() bool {
	rv := ffi.CallMethod[bool](f_, "isMainFrame")
	return rv
}

func (f_ FrameInfo) Request() foundation.URLRequest {
	rv := ffi.CallMethod[foundation.URLRequest](f_, "request")
	return rv
}

func (f_ FrameInfo) SecurityOrigin() SecurityOrigin {
	rv := ffi.CallMethod[SecurityOrigin](f_, "securityOrigin")
	return rv
}

func (f_ FrameInfo) WebView() WebView {
	rv := ffi.CallMethod[WebView](f_, "webView")
	return rv
}
