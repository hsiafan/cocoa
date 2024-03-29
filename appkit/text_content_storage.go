// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TextContentStorageClass = _TextContentStorageClass{objc.GetClass("NSTextContentStorage")}

type _TextContentStorageClass struct {
	objc.Class
}

type ITextContentStorage interface {
	ITextContentManager
	AttributedStringForTextElement(textElement ITextElement) foundation.AttributedString
	TextElementForAttributedString(attributedString foundation.IAttributedString) TextElement
	AdjustedRangeFromRange_ForEditingTextSelection(textRange ITextRange, forEditingTextSelection bool) TextRange
	AttributedString() foundation.AttributedString
	SetAttributedString(value foundation.IAttributedString)
}

type TextContentStorage struct {
	TextContentManager
}

func MakeTextContentStorage(ptr unsafe.Pointer) TextContentStorage {
	return TextContentStorage{
		TextContentManager: MakeTextContentManager(ptr),
	}
}

func (t_ TextContentStorage) Init() TextContentStorage {
	rv := objc.CallMethod[TextContentStorage](t_, objc.SEL("init"))
	return rv
}

func (tc _TextContentStorageClass) Alloc() TextContentStorage {
	rv := objc.CallMethod[TextContentStorage](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TextContentStorageClass) New() TextContentStorage {
	rv := objc.CallMethod[TextContentStorage](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTextContentStorage() TextContentStorage {
	return TextContentStorageClass.New()
}

func (t_ TextContentStorage) AttributedStringForTextElement(textElement ITextElement) foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](t_, objc.SEL("attributedStringForTextElement:"), objc.ExtractPtr(textElement))
	return rv
}

func (t_ TextContentStorage) TextElementForAttributedString(attributedString foundation.IAttributedString) TextElement {
	rv := objc.CallMethod[TextElement](t_, objc.SEL("textElementForAttributedString:"), objc.ExtractPtr(attributedString))
	return rv
}

func (t_ TextContentStorage) AdjustedRangeFromRange_ForEditingTextSelection(textRange ITextRange, forEditingTextSelection bool) TextRange {
	rv := objc.CallMethod[TextRange](t_, objc.SEL("adjustedRangeFromRange:forEditingTextSelection:"), objc.ExtractPtr(textRange), forEditingTextSelection)
	return rv
}

func (t_ TextContentStorage) AttributedString() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](t_, objc.SEL("attributedString"))
	return rv
}

func (t_ TextContentStorage) SetAttributedString(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAttributedString:"), objc.ExtractPtr(value))
}
