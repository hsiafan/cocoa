// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var StreamClass = _StreamClass{objc.GetClass("NSStream")}

type _StreamClass struct {
	objc.Class
}

type IStream interface {
	objc.IObject
	PropertyForKey(key StreamPropertyKey) objc.Object
	SetProperty_ForKey(property objc.IObject, key StreamPropertyKey) bool
	Open()
	Close()
	ScheduleInRunLoop_ForMode(aRunLoop IRunLoop, mode RunLoopMode)
	RemoveFromRunLoop_ForMode(aRunLoop IRunLoop, mode RunLoopMode)
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
	StreamStatus() StreamStatus
	StreamError() Error
}

type Stream struct {
	objc.Object
}

func MakeStream(ptr unsafe.Pointer) Stream {
	return Stream{
		Object: objc.MakeObject(ptr),
	}
}

func (sc _StreamClass) Alloc() Stream {
	rv := objc.CallMethod[Stream](sc, objc.SEL("alloc"))
	return rv
}

func (sc _StreamClass) New() Stream {
	rv := objc.CallMethod[Stream](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewStream() Stream {
	return StreamClass.New()
}

func (s_ Stream) Init() Stream {
	rv := objc.CallMethod[Stream](s_, objc.SEL("init"))
	return rv
}

func (s_ Stream) PropertyForKey(key StreamPropertyKey) objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.SEL("propertyForKey:"), key)
	return rv
}

func (s_ Stream) SetProperty_ForKey(property objc.IObject, key StreamPropertyKey) bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("setProperty:forKey:"), objc.ExtractPtr(property), key)
	return rv
}

func (s_ Stream) Open() {
	objc.CallMethod[objc.Void](s_, objc.SEL("open"))
}

func (s_ Stream) Close() {
	objc.CallMethod[objc.Void](s_, objc.SEL("close"))
}

func (s_ Stream) ScheduleInRunLoop_ForMode(aRunLoop IRunLoop, mode RunLoopMode) {
	objc.CallMethod[objc.Void](s_, objc.SEL("scheduleInRunLoop:forMode:"), objc.ExtractPtr(aRunLoop), mode)
}

func (s_ Stream) RemoveFromRunLoop_ForMode(aRunLoop IRunLoop, mode RunLoopMode) {
	objc.CallMethod[objc.Void](s_, objc.SEL("removeFromRunLoop:forMode:"), objc.ExtractPtr(aRunLoop), mode)
}

func (sc _StreamClass) GetBoundStreamsWithBufferSize_InputStream_OutputStream(bufferSize uint, inputStream *InputStream, outputStream *OutputStream) {
	objc.CallMethod[objc.Void](sc, objc.SEL("getBoundStreamsWithBufferSize:inputStream:outputStream:"), bufferSize, unsafe.Pointer(inputStream), unsafe.Pointer(outputStream))
}

// deprecated
func (sc _StreamClass) GetStreamsToHostWithName_Port_InputStream_OutputStream(hostname string, port int, inputStream *InputStream, outputStream *OutputStream) {
	objc.CallMethod[objc.Void](sc, objc.SEL("getStreamsToHostWithName:port:inputStream:outputStream:"), hostname, port, unsafe.Pointer(inputStream), unsafe.Pointer(outputStream))
}

// weak property
func (s_ Stream) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.SEL("delegate"))
	return rv
}

// weak property
func (s_ Stream) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}

func (s_ Stream) StreamStatus() StreamStatus {
	rv := objc.CallMethod[StreamStatus](s_, objc.SEL("streamStatus"))
	return rv
}

func (s_ Stream) StreamError() Error {
	rv := objc.CallMethod[Error](s_, objc.SEL("streamError"))
	return rv
}
