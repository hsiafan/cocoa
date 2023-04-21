// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var TextRangeClass = _TextRangeClass{objc.GetClass("NSTextRange")}

type _TextRangeClass struct {
	objc.Class
}

type ITextRange interface {
	objc.IObject
	IntersectsWithTextRange(textRange ITextRange) bool
	IsEqualToTextRange(textRange ITextRange) bool
	ContainsRange(textRange ITextRange) bool
	IsEmpty() bool
}

type TextRange struct {
	objc.Object
}

func MakeTextRange(ptr unsafe.Pointer) TextRange {
	return TextRange{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextRange) TextRangeByIntersectingWithTextRange(textRange ITextRange) TextRange {
	rv := objc.CallMethod[TextRange](t_, "textRangeByIntersectingWithTextRange:", textRange)
	return rv
}

func (t_ TextRange) TextRangeByFormingUnionWithTextRange(textRange ITextRange) TextRange {
	rv := objc.CallMethod[TextRange](t_, "textRangeByFormingUnionWithTextRange:", textRange)
	return rv
}

func (tc _TextRangeClass) Alloc() TextRange {
	rv := objc.CallMethod[TextRange](tc, "alloc")
	return rv
}

func (tc _TextRangeClass) New() TextRange {
	rv := objc.CallMethod[TextRange](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextRange() TextRange {
	return TextRangeClass.New()
}

func (t_ TextRange) Init() TextRange {
	rv := objc.CallMethod[TextRange](t_, "init")
	return rv
}

func (t_ TextRange) IntersectsWithTextRange(textRange ITextRange) bool {
	rv := objc.CallMethod[bool](t_, "intersectsWithTextRange:", textRange)
	return rv
}

func (t_ TextRange) IsEqualToTextRange(textRange ITextRange) bool {
	rv := objc.CallMethod[bool](t_, "isEqualToTextRange:", textRange)
	return rv
}

func (t_ TextRange) ContainsRange(textRange ITextRange) bool {
	rv := objc.CallMethod[bool](t_, "containsRange:", textRange)
	return rv
}

func (t_ TextRange) IsEmpty() bool {
	rv := objc.CallMethod[bool](t_, "isEmpty")
	return rv
}
