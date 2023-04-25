// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/internal"
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
	Delegate() StreamDelegateWrapper
	SetDelegate(value StreamDelegate)
	SetDelegate0(value objc.IObject)
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
	rv := objc.CallMethod[Stream](sc, objc.GetSelector("alloc"))
	return rv
}

func (sc _StreamClass) New() Stream {
	rv := objc.CallMethod[Stream](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewStream() Stream {
	return StreamClass.New()
}

func (s_ Stream) Init() Stream {
	rv := objc.CallMethod[Stream](s_, objc.GetSelector("init"))
	return rv
}

func (s_ Stream) PropertyForKey(key StreamPropertyKey) objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.GetSelector("propertyForKey:"), key)
	return rv
}

func (s_ Stream) SetProperty_ForKey(property objc.IObject, key StreamPropertyKey) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("setProperty:forKey:"), property, key)
	return rv
}

func (s_ Stream) Open() {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("open"))
}

func (s_ Stream) Close() {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("close"))
}

func (s_ Stream) ScheduleInRunLoop_ForMode(aRunLoop IRunLoop, mode RunLoopMode) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("scheduleInRunLoop:forMode:"), aRunLoop, mode)
}

func (s_ Stream) RemoveFromRunLoop_ForMode(aRunLoop IRunLoop, mode RunLoopMode) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("removeFromRunLoop:forMode:"), aRunLoop, mode)
}

func (sc _StreamClass) GetBoundStreamsWithBufferSize_InputStream_OutputStream(bufferSize uint, inputStream *InputStream, outputStream *OutputStream) {
	objc.CallMethod[objc.Void](sc, objc.GetSelector("getBoundStreamsWithBufferSize:inputStream:outputStream:"), bufferSize, unsafe.Pointer(inputStream), unsafe.Pointer(outputStream))
}

// deprecated
func (sc _StreamClass) GetStreamsToHostWithName_Port_InputStream_OutputStream(hostname string, port int, inputStream *InputStream, outputStream *OutputStream) {
	objc.CallMethod[objc.Void](sc, objc.GetSelector("getStreamsToHostWithName:port:inputStream:outputStream:"), hostname, port, unsafe.Pointer(inputStream), unsafe.Pointer(outputStream))
}

func (s_ Stream) Delegate() StreamDelegateWrapper {
	rv := objc.CallMethod[StreamDelegateWrapper](s_, objc.GetSelector("delegate"))
	return rv
}

func (s_ Stream) SetDelegate(value StreamDelegate) {
	po := objc.CreateProtocol("NSStreamDelegate", value)
	objc.SetAssociatedObject(s_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDelegate:"), po)
}

func (s_ Stream) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDelegate:"), value)
}

func (s_ Stream) StreamStatus() StreamStatus {
	rv := objc.CallMethod[StreamStatus](s_, objc.GetSelector("streamStatus"))
	return rv
}

func (s_ Stream) StreamError() Error {
	rv := objc.CallMethod[Error](s_, objc.GetSelector("streamError"))
	return rv
}
