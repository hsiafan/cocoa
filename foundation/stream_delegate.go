// auto-generated code, do not modify
package foundation

import (
	"github.com/hsiafan/cocoa/objc"
)

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
	objc.CallMethod[objc.Void](s_, objc.GetSelector("stream:handleEvent:"), objc.ExtractPtr(aStream), eventCode)
}
