// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[Stream](sc, "alloc")
	return rv
}

func (sc _StreamClass) New() Stream {
	rv := ffi.CallMethod[Stream](sc, "new")
	rv.Autorelease()
	return rv
}

func NewStream() Stream {
	return StreamClass.New()
}

func (s_ Stream) Init() Stream {
	rv := ffi.CallMethod[Stream](s_, "init")
	return rv
}

func (s_ Stream) PropertyForKey(key StreamPropertyKey) objc.Object {
	rv := ffi.CallMethod[objc.Object](s_, "propertyForKey:", key)
	return rv
}

func (s_ Stream) SetProperty_ForKey(property objc.IObject, key StreamPropertyKey) bool {
	rv := ffi.CallMethod[bool](s_, "setProperty:forKey:", property, key)
	return rv
}

func (s_ Stream) Open() {
	ffi.CallMethod[ffi.Void](s_, "open")
}

func (s_ Stream) Close() {
	ffi.CallMethod[ffi.Void](s_, "close")
}

func (s_ Stream) ScheduleInRunLoop_ForMode(aRunLoop IRunLoop, mode RunLoopMode) {
	ffi.CallMethod[ffi.Void](s_, "scheduleInRunLoop:forMode:", aRunLoop, mode)
}

func (s_ Stream) RemoveFromRunLoop_ForMode(aRunLoop IRunLoop, mode RunLoopMode) {
	ffi.CallMethod[ffi.Void](s_, "removeFromRunLoop:forMode:", aRunLoop, mode)
}

func (sc _StreamClass) GetBoundStreamsWithBufferSize_InputStream_OutputStream(bufferSize uint, inputStream *InputStream, outputStream *OutputStream) {
	ffi.CallMethod[ffi.Void](sc, "getBoundStreamsWithBufferSize:inputStream:outputStream:", bufferSize, unsafe.Pointer(inputStream), unsafe.Pointer(outputStream))
}

// deprecated
func (sc _StreamClass) GetStreamsToHostWithName_Port_InputStream_OutputStream(hostname string, port int, inputStream *InputStream, outputStream *OutputStream) {
	ffi.CallMethod[ffi.Void](sc, "getStreamsToHostWithName:port:inputStream:outputStream:", hostname, port, unsafe.Pointer(inputStream), unsafe.Pointer(outputStream))
}

func (s_ Stream) Delegate() StreamDelegateWrapper {
	rv := ffi.CallMethod[StreamDelegateWrapper](s_, "delegate")
	return rv
}

func (s_ Stream) SetDelegate(value StreamDelegate) {
	po := ffi.CreateProtocol("NSStreamDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(s_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](s_, "setDelegate:", po)
}

func (s_ Stream) StreamStatus() StreamStatus {
	rv := ffi.CallMethod[StreamStatus](s_, "streamStatus")
	return rv
}

func (s_ Stream) StreamError() Error {
	rv := ffi.CallMethod[Error](s_, "streamError")
	return rv
}
