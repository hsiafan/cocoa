// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TextStorageClass = _TextStorageClass{objc.GetClass("NSTextStorage")}

type _TextStorageClass struct {
	objc.Class
}

type ITextStorage interface {
	foundation.IMutableAttributedString
	AddLayoutManager(aLayoutManager ILayoutManager)
	RemoveLayoutManager(aLayoutManager ILayoutManager)
	Edited_Range_ChangeInLength(editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
	ProcessEditing()
	InvalidateAttributesInRange(range_ foundation.Range)
	EnsureAttributesAreFixedInRange(range_ foundation.Range)
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
	LayoutManagers() []LayoutManager
	FixesAttributesLazily() bool
	EditedMask() TextStorageEditActions
	EditedRange() foundation.Range
	ChangeInLength() int
	AttributeRuns() []TextStorage
	SetAttributeRuns(value []ITextStorage)
	Paragraphs() []TextStorage
	SetParagraphs(value []ITextStorage)
	Words() []TextStorage
	SetWords(value []ITextStorage)
	Characters() []TextStorage
	SetCharacters(value []ITextStorage)
	Font() Font
	SetFont(value IFont)
	ForegroundColor() Color
	SetForegroundColor(value IColor)
}

type TextStorage struct {
	foundation.MutableAttributedString
}

func MakeTextStorage(ptr unsafe.Pointer) TextStorage {
	return TextStorage{
		MutableAttributedString: foundation.MakeMutableAttributedString(ptr),
	}
}

func (t_ TextStorage) InitWithString(str string) TextStorage {
	rv := objc.CallMethod[TextStorage](t_, objc.SEL("initWithString:"), str)
	return rv
}

func (t_ TextStorage) InitWithString_Attributes(str string, attrs map[foundation.AttributedStringKey]objc.IObject) TextStorage {
	rv := objc.CallMethod[TextStorage](t_, objc.SEL("initWithString:attributes:"), str, attrs)
	return rv
}

func (t_ TextStorage) InitWithAttributedString(attrStr foundation.IAttributedString) TextStorage {
	rv := objc.CallMethod[TextStorage](t_, objc.SEL("initWithAttributedString:"), objc.ExtractPtr(attrStr))
	return rv
}

func (tc _TextStorageClass) Alloc() TextStorage {
	rv := objc.CallMethod[TextStorage](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TextStorageClass) New() TextStorage {
	rv := objc.CallMethod[TextStorage](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTextStorage() TextStorage {
	return TextStorageClass.New()
}

func (t_ TextStorage) Init() TextStorage {
	rv := objc.CallMethod[TextStorage](t_, objc.SEL("init"))
	return rv
}

func (t_ TextStorage) AddLayoutManager(aLayoutManager ILayoutManager) {
	objc.CallMethod[objc.Void](t_, objc.SEL("addLayoutManager:"), objc.ExtractPtr(aLayoutManager))
}

func (t_ TextStorage) RemoveLayoutManager(aLayoutManager ILayoutManager) {
	objc.CallMethod[objc.Void](t_, objc.SEL("removeLayoutManager:"), objc.ExtractPtr(aLayoutManager))
}

func (t_ TextStorage) Edited_Range_ChangeInLength(editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	objc.CallMethod[objc.Void](t_, objc.SEL("edited:range:changeInLength:"), editedMask, editedRange, delta)
}

func (t_ TextStorage) ProcessEditing() {
	objc.CallMethod[objc.Void](t_, objc.SEL("processEditing"))
}

func (t_ TextStorage) InvalidateAttributesInRange(range_ foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.SEL("invalidateAttributesInRange:"), range_)
}

func (t_ TextStorage) EnsureAttributesAreFixedInRange(range_ foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.SEL("ensureAttributesAreFixedInRange:"), range_)
}

// weak property
func (t_ TextStorage) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.SEL("delegate"))
	return rv
}

// weak property
func (t_ TextStorage) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}

func (t_ TextStorage) LayoutManagers() []LayoutManager {
	rv := objc.CallMethod[[]LayoutManager](t_, objc.SEL("layoutManagers"))
	return rv
}

func (t_ TextStorage) FixesAttributesLazily() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("fixesAttributesLazily"))
	return rv
}

func (t_ TextStorage) EditedMask() TextStorageEditActions {
	rv := objc.CallMethod[TextStorageEditActions](t_, objc.SEL("editedMask"))
	return rv
}

func (t_ TextStorage) EditedRange() foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.SEL("editedRange"))
	return rv
}

func (t_ TextStorage) ChangeInLength() int {
	rv := objc.CallMethod[int](t_, objc.SEL("changeInLength"))
	return rv
}

func (t_ TextStorage) AttributeRuns() []TextStorage {
	rv := objc.CallMethod[[]TextStorage](t_, objc.SEL("attributeRuns"))
	return rv
}

func (t_ TextStorage) SetAttributeRuns(value []ITextStorage) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAttributeRuns:"), value)
}

func (t_ TextStorage) Paragraphs() []TextStorage {
	rv := objc.CallMethod[[]TextStorage](t_, objc.SEL("paragraphs"))
	return rv
}

func (t_ TextStorage) SetParagraphs(value []ITextStorage) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setParagraphs:"), value)
}

func (t_ TextStorage) Words() []TextStorage {
	rv := objc.CallMethod[[]TextStorage](t_, objc.SEL("words"))
	return rv
}

func (t_ TextStorage) SetWords(value []ITextStorage) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setWords:"), value)
}

func (t_ TextStorage) Characters() []TextStorage {
	rv := objc.CallMethod[[]TextStorage](t_, objc.SEL("characters"))
	return rv
}

func (t_ TextStorage) SetCharacters(value []ITextStorage) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setCharacters:"), value)
}

func (t_ TextStorage) Font() Font {
	rv := objc.CallMethod[Font](t_, objc.SEL("font"))
	return rv
}

func (t_ TextStorage) SetFont(value IFont) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setFont:"), objc.ExtractPtr(value))
}

func (t_ TextStorage) ForegroundColor() Color {
	rv := objc.CallMethod[Color](t_, objc.SEL("foregroundColor"))
	return rv
}

func (t_ TextStorage) SetForegroundColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setForegroundColor:"), objc.ExtractPtr(value))
}
