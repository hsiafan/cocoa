// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[TextContentStorage](t_, "init")
	return rv
}

func (tc _TextContentStorageClass) Alloc() TextContentStorage {
	rv := ffi.CallMethod[TextContentStorage](tc, "alloc")
	return rv
}

func (tc _TextContentStorageClass) New() TextContentStorage {
	rv := ffi.CallMethod[TextContentStorage](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextContentStorage() TextContentStorage {
	return TextContentStorageClass.New()
}

func (t_ TextContentStorage) AttributedStringForTextElement(textElement ITextElement) foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](t_, "attributedStringForTextElement:", textElement)
	return rv
}

func (t_ TextContentStorage) TextElementForAttributedString(attributedString foundation.IAttributedString) TextElement {
	rv := ffi.CallMethod[TextElement](t_, "textElementForAttributedString:", attributedString)
	return rv
}

func (t_ TextContentStorage) AdjustedRangeFromRange_ForEditingTextSelection(textRange ITextRange, forEditingTextSelection bool) TextRange {
	rv := ffi.CallMethod[TextRange](t_, "adjustedRangeFromRange:forEditingTextSelection:", textRange, forEditingTextSelection)
	return rv
}

func (t_ TextContentStorage) AttributedString() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](t_, "attributedString")
	return rv
}

func (t_ TextContentStorage) SetAttributedString(value foundation.IAttributedString) {
	ffi.CallMethod[ffi.Void](t_, "setAttributedString:", value)
}