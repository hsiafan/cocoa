// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var NavigationResponseClass = _NavigationResponseClass{objc.GetClass("WKNavigationResponse")}

type _NavigationResponseClass struct {
	objc.Class
}

type INavigationResponse interface {
	objc.IObject
	Response() foundation.URLResponse
	CanShowMIMEType() bool
	IsForMainFrame() bool
}

type NavigationResponse struct {
	objc.Object
}

func MakeNavigationResponse(ptr unsafe.Pointer) NavigationResponse {
	return NavigationResponse{
		Object: objc.MakeObject(ptr),
	}
}

func (nc _NavigationResponseClass) Alloc() NavigationResponse {
	rv := ffi.CallMethod[NavigationResponse](nc, "alloc")
	return rv
}

func (nc _NavigationResponseClass) New() NavigationResponse {
	rv := ffi.CallMethod[NavigationResponse](nc, "new")
	rv.Autorelease()
	return rv
}

func NewNavigationResponse() NavigationResponse {
	return NavigationResponseClass.New()
}

func (n_ NavigationResponse) Init() NavigationResponse {
	rv := ffi.CallMethod[NavigationResponse](n_, "init")
	return rv
}

func (n_ NavigationResponse) Response() foundation.URLResponse {
	rv := ffi.CallMethod[foundation.URLResponse](n_, "response")
	return rv
}

func (n_ NavigationResponse) CanShowMIMEType() bool {
	rv := ffi.CallMethod[bool](n_, "canShowMIMEType")
	return rv
}

func (n_ NavigationResponse) IsForMainFrame() bool {
	rv := ffi.CallMethod[bool](n_, "isForMainFrame")
	return rv
}