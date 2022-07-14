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

func (s_ Stream) Init() Stream {
	rv := ffi.CallMethod[Stream](s_, "init")
	rv.Autorelease()
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
	ffi.CallMethod[ffi.Void](sc, "getBoundStreamsWithBufferSize:inputStream:outputStream:", bufferSize, inputStream, outputStream)
}

// deprecated
func (sc _StreamClass) GetStreamsToHostWithName_Port_InputStream_OutputStream(hostname string, port int, inputStream *InputStream, outputStream *OutputStream) {
	ffi.CallMethod[ffi.Void](sc, "getStreamsToHostWithName:port:inputStream:outputStream:", hostname, port, inputStream, outputStream)
}

func (s_ Stream) Delegate() StreamDelegateWrapper {
	rv := ffi.CallMethod[StreamDelegateWrapper](s_, "delegate")
	return rv
}

func (s_ Stream) SetDelegate(value StreamDelegate) {
	po := ffi.CreateProtocol(value)
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
	rv := ffi.CallMethod[InputStream](ic, "inputStreamWithData:", data)
	return rv
}

func (ic _InputStreamClass) InputStreamWithFileAtPath(path string) InputStream {
	rv := ffi.CallMethod[InputStream](ic, "inputStreamWithFileAtPath:", path)
	return rv
}

func (ic _InputStreamClass) InputStreamWithURL(url IURL) InputStream {
	rv := ffi.CallMethod[InputStream](ic, "inputStreamWithURL:", url)
	return rv
}

func (i_ InputStream) InitWithData(data []byte) InputStream {
	rv := ffi.CallMethod[InputStream](i_, "initWithData:", data)
	rv.Autorelease()
	return rv
}

func (i_ InputStream) InitWithFileAtPath(path string) InputStream {
	rv := ffi.CallMethod[InputStream](i_, "initWithFileAtPath:", path)
	rv.Autorelease()
	return rv
}

func (i_ InputStream) InitWithURL(url IURL) InputStream {
	rv := ffi.CallMethod[InputStream](i_, "initWithURL:", url)
	rv.Autorelease()
	return rv
}

func (ic _InputStreamClass) Alloc() InputStream {
	rv := ffi.CallMethod[InputStream](ic, "alloc")
	return rv
}

func (i_ InputStream) Init() InputStream {
	rv := ffi.CallMethod[InputStream](i_, "init")
	rv.Autorelease()
	return rv
}

func (ic _InputStreamClass) New() InputStream {
	rv := ffi.CallMethod[InputStream](ic, "new")
	rv.Autorelease()
	return rv
}

func NewInputStream() InputStream {
	return InputStreamClass.New()
}

func (i_ InputStream) HasBytesAvailable() bool {
	rv := ffi.CallMethod[bool](i_, "hasBytesAvailable")
	return rv
}

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
	rv := ffi.CallMethod[OutputStream](oc, "outputStreamToMemory")
	return rv
}

func (oc _OutputStreamClass) OutputStreamToFileAtPath_Append(path string, shouldAppend bool) OutputStream {
	rv := ffi.CallMethod[OutputStream](oc, "outputStreamToFileAtPath:append:", path, shouldAppend)
	return rv
}

func (oc _OutputStreamClass) OutputStreamWithURL_Append(url IURL, shouldAppend bool) OutputStream {
	rv := ffi.CallMethod[OutputStream](oc, "outputStreamWithURL:append:", url, shouldAppend)
	return rv
}

func (o_ OutputStream) InitToMemory() OutputStream {
	rv := ffi.CallMethod[OutputStream](o_, "initToMemory")
	rv.Autorelease()
	return rv
}

func (o_ OutputStream) InitToFileAtPath_Append(path string, shouldAppend bool) OutputStream {
	rv := ffi.CallMethod[OutputStream](o_, "initToFileAtPath:append:", path, shouldAppend)
	rv.Autorelease()
	return rv
}

func (o_ OutputStream) InitWithURL_Append(url IURL, shouldAppend bool) OutputStream {
	rv := ffi.CallMethod[OutputStream](o_, "initWithURL:append:", url, shouldAppend)
	rv.Autorelease()
	return rv
}

func (oc _OutputStreamClass) Alloc() OutputStream {
	rv := ffi.CallMethod[OutputStream](oc, "alloc")
	return rv
}

func (o_ OutputStream) Init() OutputStream {
	rv := ffi.CallMethod[OutputStream](o_, "init")
	rv.Autorelease()
	return rv
}

func (oc _OutputStreamClass) New() OutputStream {
	rv := ffi.CallMethod[OutputStream](oc, "new")
	rv.Autorelease()
	return rv
}

func NewOutputStream() OutputStream {
	return OutputStreamClass.New()
}

func (o_ OutputStream) HasSpaceAvailable() bool {
	rv := ffi.CallMethod[bool](o_, "hasSpaceAvailable")
	return rv
}

type StreamDelegate interface {
	ImplementsStream_HandleEvent() bool
	// optional
	Stream_HandleEvent(aStream Stream, eventCode StreamEvent)
}

type StreamDelegateImpl struct {
	_Stream_HandleEvent func(aStream Stream, eventCode StreamEvent)
}

func (di *StreamDelegateImpl) ImplementsStream_HandleEvent() bool {
	return di._Stream_HandleEvent != nil
}

func (di *StreamDelegateImpl) SetStream_HandleEvent(f func(aStream Stream, eventCode StreamEvent)) {
	di._Stream_HandleEvent = f
}

func (di *StreamDelegateImpl) Stream_HandleEvent(aStream Stream, eventCode StreamEvent) {
	di._Stream_HandleEvent(aStream, eventCode)
}

type StreamDelegateWrapper struct {
	objc.Object
}

func (s_ *StreamDelegateWrapper) ImplementsStream_HandleEvent() bool {
	return s_.RespondsToSelector(objc.GetSelector("stream:handleEvent:"))
}

func (s_ StreamDelegateWrapper) Stream_HandleEvent(aStream IStream, eventCode StreamEvent) {
	ffi.CallMethod[ffi.Void](s_, "stream:handleEvent:", aStream, eventCode)
}
