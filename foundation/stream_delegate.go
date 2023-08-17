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

func WrapStreamDelegate(v StreamDelegate) objc.Object {
	return objc.WrapAsProtocol("NSStreamDelegate", v)
}

type StreamDelegateBase struct {
}

func (p *StreamDelegateBase) ImplementsStream_HandleEvent() bool {
	return false
}

func (p *StreamDelegateBase) Stream_HandleEvent(aStream Stream, eventCode StreamEvent) {
	panic("unimpemented protocol method")
}

type StreamDelegateWrapper struct {
	objc.Object
}

func (s_ StreamDelegateWrapper) Stream_HandleEvent(aStream IStream, eventCode StreamEvent) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("stream:handleEvent:"), objc.ExtractPtr(aStream), eventCode)
}
