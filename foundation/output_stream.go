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
	rv := objc.CallMethod[OutputStream](oc, "outputStreamToMemory")
	return rv
}

func (oc _OutputStreamClass) OutputStreamToFileAtPath_Append(path string, shouldAppend bool) OutputStream {
	rv := objc.CallMethod[OutputStream](oc, "outputStreamToFileAtPath:append:", path, shouldAppend)
	return rv
}

func (oc _OutputStreamClass) OutputStreamWithURL_Append(url IURL, shouldAppend bool) OutputStream {
	rv := objc.CallMethod[OutputStream](oc, "outputStreamWithURL:append:", url, shouldAppend)
	return rv
}

func (o_ OutputStream) InitToMemory() OutputStream {
	rv := objc.CallMethod[OutputStream](o_, "initToMemory")
	return rv
}

func (o_ OutputStream) InitToFileAtPath_Append(path string, shouldAppend bool) OutputStream {
	rv := objc.CallMethod[OutputStream](o_, "initToFileAtPath:append:", path, shouldAppend)
	return rv
}

func (o_ OutputStream) InitWithURL_Append(url IURL, shouldAppend bool) OutputStream {
	rv := objc.CallMethod[OutputStream](o_, "initWithURL:append:", url, shouldAppend)
	return rv
}

func (oc _OutputStreamClass) Alloc() OutputStream {
	rv := objc.CallMethod[OutputStream](oc, "alloc")
	return rv
}

func (oc _OutputStreamClass) New() OutputStream {
	rv := objc.CallMethod[OutputStream](oc, "new")
	rv.Autorelease()
	return rv
}

func NewOutputStream() OutputStream {
	return OutputStreamClass.New()
}

func (o_ OutputStream) Init() OutputStream {
	rv := objc.CallMethod[OutputStream](o_, "init")
	return rv
}

func (o_ OutputStream) HasSpaceAvailable() bool {
	rv := objc.CallMethod[bool](o_, "hasSpaceAvailable")
	return rv
}
