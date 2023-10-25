// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var ScriptMessageClass = _ScriptMessageClass{objc.GetClass("WKScriptMessage")}

type _ScriptMessageClass struct {
	objc.Class
}

type IScriptMessage interface {
	objc.IObject
	Body() objc.Object
	FrameInfo() FrameInfo
	WebView() WebView
	World() ContentWorld
	Name() string
}

type ScriptMessage struct {
	objc.Object
}

func MakeScriptMessage(ptr unsafe.Pointer) ScriptMessage {
	return ScriptMessage{
		Object: objc.MakeObject(ptr),
	}
}

func (sc _ScriptMessageClass) Alloc() ScriptMessage {
	rv := objc.CallMethod[ScriptMessage](sc, objc.SEL("alloc"))
	return rv
}

func (sc _ScriptMessageClass) New() ScriptMessage {
	rv := objc.CallMethod[ScriptMessage](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewScriptMessage() ScriptMessage {
	return ScriptMessageClass.New()
}

func (s_ ScriptMessage) Init() ScriptMessage {
	rv := objc.CallMethod[ScriptMessage](s_, objc.SEL("init"))
	return rv
}

func (s_ ScriptMessage) Body() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.SEL("body"))
	return rv
}

func (s_ ScriptMessage) FrameInfo() FrameInfo {
	rv := objc.CallMethod[FrameInfo](s_, objc.SEL("frameInfo"))
	return rv
}

// weak property
func (s_ ScriptMessage) WebView() WebView {
	rv := objc.CallMethod[WebView](s_, objc.SEL("webView"))
	return rv
}

func (s_ ScriptMessage) World() ContentWorld {
	rv := objc.CallMethod[ContentWorld](s_, objc.SEL("world"))
	return rv
}

func (s_ ScriptMessage) Name() string {
	rv := objc.CallMethod[string](s_, objc.SEL("name"))
	return rv
}
