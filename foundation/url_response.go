// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var URLResponseClass = _URLResponseClass{objc.GetClass("NSURLResponse")}

type _URLResponseClass struct {
	objc.Class
}

type IURLResponse interface {
	objc.IObject
	ExpectedContentLength() int64
	SuggestedFilename() string
	MIMEType() string
	TextEncodingName() string
	URL() URL
}

type URLResponse struct {
	objc.Object
}

func MakeURLResponse(ptr unsafe.Pointer) URLResponse {
	return URLResponse{
		Object: objc.MakeObject(ptr),
	}
}

func (u_ URLResponse) InitWithURL_MIMEType_ExpectedContentLength_TextEncodingName(URL IURL, MIMEType string, length int, name string) URLResponse {
	rv := objc.CallMethod[URLResponse](u_, "initWithURL:MIMEType:expectedContentLength:textEncodingName:", URL, MIMEType, length, name)
	return rv
}

func (uc _URLResponseClass) Alloc() URLResponse {
	rv := objc.CallMethod[URLResponse](uc, "alloc")
	return rv
}

func (uc _URLResponseClass) New() URLResponse {
	rv := objc.CallMethod[URLResponse](uc, "new")
	rv.Autorelease()
	return rv
}

func NewURLResponse() URLResponse {
	return URLResponseClass.New()
}

func (u_ URLResponse) Init() URLResponse {
	rv := objc.CallMethod[URLResponse](u_, "init")
	return rv
}

func (u_ URLResponse) ExpectedContentLength() int64 {
	rv := objc.CallMethod[int64](u_, "expectedContentLength")
	return rv
}

func (u_ URLResponse) SuggestedFilename() string {
	rv := objc.CallMethod[string](u_, "suggestedFilename")
	return rv
}

func (u_ URLResponse) MIMEType() string {
	rv := objc.CallMethod[string](u_, "MIMEType")
	return rv
}

func (u_ URLResponse) TextEncodingName() string {
	rv := objc.CallMethod[string](u_, "textEncodingName")
	return rv
}

func (u_ URLResponse) URL() URL {
	rv := objc.CallMethod[URL](u_, "URL")
	return rv
}
