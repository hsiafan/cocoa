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
	rv := objc.CallMethod[TextElement](t_, objc.SEL("initWithTextContentManager:"), objc.ExtractPtr(textContentManager))
	return rv
}

func (tc _TextElementClass) Alloc() TextElement {
	rv := objc.CallMethod[TextElement](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TextElementClass) New() TextElement {
	rv := objc.CallMethod[TextElement](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTextElement() TextElement {
	return TextElementClass.New()
}

func (t_ TextElement) Init() TextElement {
	rv := objc.CallMethod[TextElement](t_, objc.SEL("init"))
	return rv
}

// weak property
func (t_ TextElement) TextContentManager() TextContentManager {
	rv := objc.CallMethod[TextContentManager](t_, objc.SEL("textContentManager"))
	return rv
}

// weak property
func (t_ TextElement) SetTextContentManager(value ITextContentManager) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTextContentManager:"), objc.ExtractPtr(value))
}

func (t_ TextElement) ElementRange() TextRange {
	rv := objc.CallMethod[TextRange](t_, objc.SEL("elementRange"))
	return rv
}

func (t_ TextElement) SetElementRange(value ITextRange) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setElementRange:"), objc.ExtractPtr(value))
}

func (t_ TextElement) IsRepresentedElement() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isRepresentedElement"))
	return rv
}

// weak property
func (t_ TextElement) ParentElement() TextElement {
	rv := objc.CallMethod[TextElement](t_, objc.SEL("parentElement"))
	return rv
}

func (t_ TextElement) ChildElements() []TextElement {
	rv := objc.CallMethod[[]TextElement](t_, objc.SEL("childElements"))
	return rv
}
