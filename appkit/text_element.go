// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var TextElementClass = _TextElementClass{objc.GetClass("NSTextElement")}

type _TextElementClass struct {
	objc.Class
}

type ITextElement interface {
	objc.IObject
	TextContentManager() TextContentManager
	SetTextContentManager(value ITextContentManager)
	ElementRange() TextRange
	SetElementRange(value ITextRange)
	IsRepresentedElement() bool
	ParentElement() TextElement
	ChildElements() []TextElement
}

type TextElement struct {
	objc.Object
}

func MakeTextElement(ptr unsafe.Pointer) TextElement {
	return TextElement{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextElement) InitWithTextContentManager(textContentManager ITextContentManager) TextElement {
	rv := ffi.CallMethod[TextElement](t_, "initWithTextContentManager:", textContentManager)
	return rv
}

func (tc _TextElementClass) Alloc() TextElement {
	rv := ffi.CallMethod[TextElement](tc, "alloc")
	return rv
}

func (tc _TextElementClass) New() TextElement {
	rv := ffi.CallMethod[TextElement](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextElement() TextElement {
	return TextElementClass.New()
}

func (t_ TextElement) Init() TextElement {
	rv := ffi.CallMethod[TextElement](t_, "init")
	return rv
}

func (t_ TextElement) TextContentManager() TextContentManager {
	rv := ffi.CallMethod[TextContentManager](t_, "textContentManager")
	return rv
}

func (t_ TextElement) SetTextContentManager(value ITextContentManager) {
	ffi.CallMethod[ffi.Void](t_, "setTextContentManager:", value)
}

func (t_ TextElement) ElementRange() TextRange {
	rv := ffi.CallMethod[TextRange](t_, "elementRange")
	return rv
}

func (t_ TextElement) SetElementRange(value ITextRange) {
	ffi.CallMethod[ffi.Void](t_, "setElementRange:", value)
}

func (t_ TextElement) IsRepresentedElement() bool {
	rv := ffi.CallMethod[bool](t_, "isRepresentedElement")
	return rv
}

func (t_ TextElement) ParentElement() TextElement {
	rv := ffi.CallMethod[TextElement](t_, "parentElement")
	return rv
}

func (t_ TextElement) ChildElements() []TextElement {
	rv := ffi.CallMethod[[]TextElement](t_, "childElements")
	return rv
}
