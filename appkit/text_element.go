// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[TextElement](t_, objc.GetSelector("initWithTextContentManager:"), textContentManager)
	return rv
}

func (tc _TextElementClass) Alloc() TextElement {
	rv := objc.CallMethod[TextElement](tc, objc.GetSelector("alloc"))
	return rv
}

func (tc _TextElementClass) New() TextElement {
	rv := objc.CallMethod[TextElement](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextElement() TextElement {
	return TextElementClass.New()
}

func (t_ TextElement) Init() TextElement {
	rv := objc.CallMethod[TextElement](t_, objc.GetSelector("init"))
	return rv
}

func (t_ TextElement) TextContentManager() TextContentManager {
	rv := objc.CallMethod[TextContentManager](t_, objc.GetSelector("textContentManager"))
	return rv
}

func (t_ TextElement) SetTextContentManager(value ITextContentManager) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTextContentManager:"), value)
}

func (t_ TextElement) ElementRange() TextRange {
	rv := objc.CallMethod[TextRange](t_, objc.GetSelector("elementRange"))
	return rv
}

func (t_ TextElement) SetElementRange(value ITextRange) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setElementRange:"), value)
}

func (t_ TextElement) IsRepresentedElement() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isRepresentedElement"))
	return rv
}

func (t_ TextElement) ParentElement() TextElement {
	rv := objc.CallMethod[TextElement](t_, objc.GetSelector("parentElement"))
	return rv
}

func (t_ TextElement) ChildElements() []TextElement {
	rv := objc.CallMethod[[]TextElement](t_, objc.GetSelector("childElements"))
	return rv
}
