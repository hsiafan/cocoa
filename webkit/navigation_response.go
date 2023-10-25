// auto-generated code, do not modify
package webkit

import (
	"unsafe"

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
	rv := objc.CallMethod[NavigationResponse](nc, objc.SEL("alloc"))
	return rv
}

func (nc _NavigationResponseClass) New() NavigationResponse {
	rv := objc.CallMethod[NavigationResponse](nc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewNavigationResponse() NavigationResponse {
	return NavigationResponseClass.New()
}

func (n_ NavigationResponse) Init() NavigationResponse {
	rv := objc.CallMethod[NavigationResponse](n_, objc.SEL("init"))
	return rv
}

func (n_ NavigationResponse) Response() foundation.URLResponse {
	rv := objc.CallMethod[foundation.URLResponse](n_, objc.SEL("response"))
	return rv
}

func (n_ NavigationResponse) CanShowMIMEType() bool {
	rv := objc.CallMethod[bool](n_, objc.SEL("canShowMIMEType"))
	return rv
}

func (n_ NavigationResponse) IsForMainFrame() bool {
	rv := objc.CallMethod[bool](n_, objc.SEL("isForMainFrame"))
	return rv
}
