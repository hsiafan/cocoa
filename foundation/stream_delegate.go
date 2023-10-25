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

type StreamDelegateCreator struct {
	className string
	class     objc.Class
}

func NewStreamDelegateCreator(name string) *StreamDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &StreamDelegateCreator{className: name, class: class}
}

func (c *StreamDelegateCreator) SetStream_HandleEvent(handle func(o objc.Object, aStream Stream, eventCode StreamEvent)) {
	objc.AddMethod(c.class, objc.GetSelector("stream:handleEvent:"), handle)
}

func (c *StreamDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
