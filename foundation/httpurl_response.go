// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var HTTPURLResponseClass = _HTTPURLResponseClass{objc.GetClass("NSHTTPURLResponse")}

type _HTTPURLResponseClass struct {
	objc.Class
}

type IHTTPURLResponse interface {
	IURLResponse
	ValueForHTTPHeaderField(field string) string
	StatusCode() int
}

type HTTPURLResponse struct {
	URLResponse
}

func MakeHTTPURLResponse(ptr unsafe.Pointer) HTTPURLResponse {
	return HTTPURLResponse{
		URLResponse: MakeURLResponse(ptr),
	}
}

func (h_ HTTPURLResponse) InitWithURL_StatusCode_HTTPVersion_HeaderFields(url IURL, statusCode int, HTTPVersion string, headerFields map[string]string) HTTPURLResponse {
	rv := objc.CallMethod[HTTPURLResponse](h_, objc.GetSelector("initWithURL:statusCode:HTTPVersion:headerFields:"), objc.ExtractPtr(url), statusCode, HTTPVersion, headerFields)
	return rv
}

func (h_ HTTPURLResponse) InitWithURL_MIMEType_ExpectedContentLength_TextEncodingName(URL IURL, MIMEType string, length int, name string) HTTPURLResponse {
	rv := objc.CallMethod[HTTPURLResponse](h_, objc.GetSelector("initWithURL:MIMEType:expectedContentLength:textEncodingName:"), objc.ExtractPtr(URL), MIMEType, length, name)
	return rv
}

func (hc _HTTPURLResponseClass) Alloc() HTTPURLResponse {
	rv := objc.CallMethod[HTTPURLResponse](hc, objc.GetSelector("alloc"))
	return rv
}

func (hc _HTTPURLResponseClass) New() HTTPURLResponse {
	rv := objc.CallMethod[HTTPURLResponse](hc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewHTTPURLResponse() HTTPURLResponse {
	return HTTPURLResponseClass.New()
}

func (h_ HTTPURLResponse) Init() HTTPURLResponse {
	rv := objc.CallMethod[HTTPURLResponse](h_, objc.GetSelector("init"))
	return rv
}

func (h_ HTTPURLResponse) ValueForHTTPHeaderField(field string) string {
	rv := objc.CallMethod[string](h_, objc.GetSelector("valueForHTTPHeaderField:"), field)
	return rv
}

func (hc _HTTPURLResponseClass) LocalizedStringForStatusCode(statusCode int) string {
	rv := objc.CallMethod[string](hc, objc.GetSelector("localizedStringForStatusCode:"), statusCode)
	return rv
}

func (h_ HTTPURLResponse) StatusCode() int {
	rv := objc.CallMethod[int](h_, objc.GetSelector("statusCode"))
	return rv
}
