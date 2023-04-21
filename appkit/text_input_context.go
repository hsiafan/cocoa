// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var TextInputContextClass = _TextInputContextClass{objc.GetClass("NSTextInputContext")}

type _TextInputContextClass struct {
	objc.Class
}

type ITextInputContext interface {
	objc.IObject
	Activate()
	Deactivate()
	HandleEvent(event IEvent) bool
	DiscardMarkedText()
	InvalidateCharacterCoordinates()
	AcceptsGlyphInfo() bool
	SetAcceptsGlyphInfo(value bool)
	AllowedInputSourceLocales() []string
	SetAllowedInputSourceLocales(value []string)
	KeyboardInputSources() []TextInputSourceIdentifier
	SelectedKeyboardInputSource() TextInputSourceIdentifier
	SetSelectedKeyboardInputSource(value TextInputSourceIdentifier)
}

type TextInputContext struct {
	objc.Object
}

func MakeTextInputContext(ptr unsafe.Pointer) TextInputContext {
	return TextInputContext{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TextInputContextClass) Alloc() TextInputContext {
	rv := objc.CallMethod[TextInputContext](tc, "alloc")
	return rv
}

func (tc _TextInputContextClass) New() TextInputContext {
	rv := objc.CallMethod[TextInputContext](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextInputContext() TextInputContext {
	return TextInputContextClass.New()
}

func (t_ TextInputContext) Init() TextInputContext {
	rv := objc.CallMethod[TextInputContext](t_, "init")
	return rv
}

func (t_ TextInputContext) Activate() {
	objc.CallMethod[objc.Void](t_, "activate")
}

func (t_ TextInputContext) Deactivate() {
	objc.CallMethod[objc.Void](t_, "deactivate")
}

func (t_ TextInputContext) HandleEvent(event IEvent) bool {
	rv := objc.CallMethod[bool](t_, "handleEvent:", event)
	return rv
}

func (t_ TextInputContext) DiscardMarkedText() {
	objc.CallMethod[objc.Void](t_, "discardMarkedText")
}

func (t_ TextInputContext) InvalidateCharacterCoordinates() {
	objc.CallMethod[objc.Void](t_, "invalidateCharacterCoordinates")
}

func (tc _TextInputContextClass) LocalizedNameForInputSource(inputSourceIdentifier TextInputSourceIdentifier) string {
	rv := objc.CallMethod[string](tc, "localizedNameForInputSource:", inputSourceIdentifier)
	return rv
}

func (tc _TextInputContextClass) CurrentInputContext() TextInputContext {
	rv := objc.CallMethod[TextInputContext](tc, "currentInputContext")
	return rv
}

func (t_ TextInputContext) AcceptsGlyphInfo() bool {
	rv := objc.CallMethod[bool](t_, "acceptsGlyphInfo")
	return rv
}

func (t_ TextInputContext) SetAcceptsGlyphInfo(value bool) {
	objc.CallMethod[objc.Void](t_, "setAcceptsGlyphInfo:", value)
}

func (t_ TextInputContext) AllowedInputSourceLocales() []string {
	rv := objc.CallMethod[[]string](t_, "allowedInputSourceLocales")
	return rv
}

func (t_ TextInputContext) SetAllowedInputSourceLocales(value []string) {
	objc.CallMethod[objc.Void](t_, "setAllowedInputSourceLocales:", value)
}

func (t_ TextInputContext) KeyboardInputSources() []TextInputSourceIdentifier {
	rv := objc.CallMethod[[]TextInputSourceIdentifier](t_, "keyboardInputSources")
	return rv
}

func (t_ TextInputContext) SelectedKeyboardInputSource() TextInputSourceIdentifier {
	rv := objc.CallMethod[TextInputSourceIdentifier](t_, "selectedKeyboardInputSource")
	return rv
}

func (t_ TextInputContext) SetSelectedKeyboardInputSource(value TextInputSourceIdentifier) {
	objc.CallMethod[objc.Void](t_, "setSelectedKeyboardInputSource:", value)
}
