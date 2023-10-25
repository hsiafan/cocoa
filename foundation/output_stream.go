// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var OutputStreamClass = _OutputStreamClass{objc.GetClass("NSOutputStream")}

type _OutputStreamClass struct {
	objc.Class
}

type IOutputStream interface {
	IStream
	HasSpaceAvailable() bool
}

type OutputStream struct {
	Stream
}

func MakeOutputStream(ptr unsafe.Pointer) OutputStream {
	return OutputStream{
		Stream: MakeStream(ptr),
	}
}

func (oc _OutputStreamClass) OutputStreamToMemory() OutputStream {
	rv := objc.CallMethod[OutputStream](oc, objc.SEL("outputStreamToMemory"))
	return rv
}

func (oc _OutputStreamClass) OutputStreamToFileAtPath_Append(path string, shouldAppend bool) OutputStream {
	rv := objc.CallMethod[OutputStream](oc, objc.SEL("outputStreamToFileAtPath:append:"), path, shouldAppend)
	return rv
}

func (oc _OutputStreamClass) OutputStreamWithURL_Append(url IURL, shouldAppend bool) OutputStream {
	rv := objc.CallMethod[OutputStream](oc, objc.SEL("outputStreamWithURL:append:"), objc.ExtractPtr(url), shouldAppend)
	return rv
}

func (o_ OutputStream) InitToMemory() OutputStream {
	rv := objc.CallMethod[OutputStream](o_, objc.SEL("initToMemory"))
	return rv
}

func (o_ OutputStream) InitToFileAtPath_Append(path string, shouldAppend bool) OutputStream {
	rv := objc.CallMethod[OutputStream](o_, objc.SEL("initToFileAtPath:append:"), path, shouldAppend)
	return rv
}

func (o_ OutputStream) InitWithURL_Append(url IURL, shouldAppend bool) OutputStream {
	rv := objc.CallMethod[OutputStream](o_, objc.SEL("initWithURL:append:"), objc.ExtractPtr(url), shouldAppend)
	return rv
}

func (oc _OutputStreamClass) Alloc() OutputStream {
	rv := objc.CallMethod[OutputStream](oc, objc.SEL("alloc"))
	return rv
}

func (oc _OutputStreamClass) New() OutputStream {
	rv := objc.CallMethod[OutputStream](oc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewOutputStream() OutputStream {
	return OutputStreamClass.New()
}

func (o_ OutputStream) Init() OutputStream {
	rv := objc.CallMethod[OutputStream](o_, objc.SEL("init"))
	return rv
}

func (o_ OutputStream) HasSpaceAvailable() bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("hasSpaceAvailable"))
	return rv
}
