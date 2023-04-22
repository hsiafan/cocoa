// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var InputStreamClass = _InputStreamClass{objc.GetClass("NSInputStream")}

type _InputStreamClass struct {
	objc.Class
}

type IInputStream interface {
	IStream
	HasBytesAvailable() bool
}

type InputStream struct {
	Stream
}

func MakeInputStream(ptr unsafe.Pointer) InputStream {
	return InputStream{
		Stream: MakeStream(ptr),
	}
}

func (ic _InputStreamClass) InputStreamWithData(data []byte) InputStream {
	rv := objc.CallMethod[InputStream](ic, objc.GetSelector("inputStreamWithData:"), data)
	return rv
}

func (ic _InputStreamClass) InputStreamWithFileAtPath(path string) InputStream {
	rv := objc.CallMethod[InputStream](ic, objc.GetSelector("inputStreamWithFileAtPath:"), path)
	return rv
}

func (ic _InputStreamClass) InputStreamWithURL(url IURL) InputStream {
	rv := objc.CallMethod[InputStream](ic, objc.GetSelector("inputStreamWithURL:"), url)
	return rv
}

func (i_ InputStream) InitWithData(data []byte) InputStream {
	rv := objc.CallMethod[InputStream](i_, objc.GetSelector("initWithData:"), data)
	return rv
}

func (i_ InputStream) InitWithFileAtPath(path string) InputStream {
	rv := objc.CallMethod[InputStream](i_, objc.GetSelector("initWithFileAtPath:"), path)
	return rv
}

func (i_ InputStream) InitWithURL(url IURL) InputStream {
	rv := objc.CallMethod[InputStream](i_, objc.GetSelector("initWithURL:"), url)
	return rv
}

func (ic _InputStreamClass) Alloc() InputStream {
	rv := objc.CallMethod[InputStream](ic, objc.GetSelector("alloc"))
	return rv
}

func (ic _InputStreamClass) New() InputStream {
	rv := objc.CallMethod[InputStream](ic, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewInputStream() InputStream {
	return InputStreamClass.New()
}

func (i_ InputStream) Init() InputStream {
	rv := objc.CallMethod[InputStream](i_, objc.GetSelector("init"))
	return rv
}

func (i_ InputStream) HasBytesAvailable() bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("hasBytesAvailable"))
	return rv
}
