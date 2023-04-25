// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
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
	Delegate() TextStorageDelegateWrapper
	SetDelegate(value TextStorageDelegate)
	SetDelegate0(value objc.IObject)
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
	rv := objc.CallMethod[TextStorage](t_, objc.GetSelector("initWithString:"), str)
	return rv
}

func (t_ TextStorage) InitWithString_Attributes(str string, attrs map[foundation.AttributedStringKey]objc.IObject) TextStorage {
	rv := objc.CallMethod[TextStorage](t_, objc.GetSelector("initWithString:attributes:"), str, attrs)
	return rv
}

func (t_ TextStorage) InitWithAttributedString(attrStr foundation.IAttributedString) TextStorage {
	rv := objc.CallMethod[TextStorage](t_, objc.GetSelector("initWithAttributedString:"), attrStr)
	return rv
}

func (tc _TextStorageClass) Alloc() TextStorage {
	rv := objc.CallMethod[TextStorage](tc, objc.GetSelector("alloc"))
	return rv
}

func (tc _TextStorageClass) New() TextStorage {
	rv := objc.CallMethod[TextStorage](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextStorage() TextStorage {
	return TextStorageClass.New()
}

func (t_ TextStorage) Init() TextStorage {
	rv := objc.CallMethod[TextStorage](t_, objc.GetSelector("init"))
	return rv
}

func (t_ TextStorage) AddLayoutManager(aLayoutManager ILayoutManager) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("addLayoutManager:"), aLayoutManager)
}

func (t_ TextStorage) RemoveLayoutManager(aLayoutManager ILayoutManager) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("removeLayoutManager:"), aLayoutManager)
}

func (t_ TextStorage) Edited_Range_ChangeInLength(editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("edited:range:changeInLength:"), editedMask, editedRange, delta)
}

func (t_ TextStorage) ProcessEditing() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("processEditing"))
}

func (t_ TextStorage) InvalidateAttributesInRange(range_ foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("invalidateAttributesInRange:"), range_)
}

func (t_ TextStorage) EnsureAttributesAreFixedInRange(range_ foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("ensureAttributesAreFixedInRange:"), range_)
}

func (t_ TextStorage) Delegate() TextStorageDelegateWrapper {
	rv := objc.CallMethod[TextStorageDelegateWrapper](t_, objc.GetSelector("delegate"))
	return rv
}

func (t_ TextStorage) SetDelegate(value TextStorageDelegate) {
	po := objc.CreateProtocol("NSTextStorageDelegate", value)
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDelegate:"), po)
}

func (t_ TextStorage) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDelegate:"), value)
}

func (t_ TextStorage) LayoutManagers() []LayoutManager {
	rv := objc.CallMethod[[]LayoutManager](t_, objc.GetSelector("layoutManagers"))
	return rv
}

func (t_ TextStorage) FixesAttributesLazily() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("fixesAttributesLazily"))
	return rv
}

func (t_ TextStorage) EditedMask() TextStorageEditActions {
	rv := objc.CallMethod[TextStorageEditActions](t_, objc.GetSelector("editedMask"))
	return rv
}

func (t_ TextStorage) EditedRange() foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("editedRange"))
	return rv
}

func (t_ TextStorage) ChangeInLength() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("changeInLength"))
	return rv
}

func (t_ TextStorage) AttributeRuns() []TextStorage {
	rv := objc.CallMethod[[]TextStorage](t_, objc.GetSelector("attributeRuns"))
	return rv
}

func (t_ TextStorage) SetAttributeRuns(value []ITextStorage) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAttributeRuns:"), value)
}

func (t_ TextStorage) Paragraphs() []TextStorage {
	rv := objc.CallMethod[[]TextStorage](t_, objc.GetSelector("paragraphs"))
	return rv
}

func (t_ TextStorage) SetParagraphs(value []ITextStorage) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setParagraphs:"), value)
}

func (t_ TextStorage) Words() []TextStorage {
	rv := objc.CallMethod[[]TextStorage](t_, objc.GetSelector("words"))
	return rv
}

func (t_ TextStorage) SetWords(value []ITextStorage) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setWords:"), value)
}

func (t_ TextStorage) Characters() []TextStorage {
	rv := objc.CallMethod[[]TextStorage](t_, objc.GetSelector("characters"))
	return rv
}

func (t_ TextStorage) SetCharacters(value []ITextStorage) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setCharacters:"), value)
}

func (t_ TextStorage) Font() Font {
	rv := objc.CallMethod[Font](t_, objc.GetSelector("font"))
	return rv
}

func (t_ TextStorage) SetFont(value IFont) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setFont:"), value)
}

func (t_ TextStorage) ForegroundColor() Color {
	rv := objc.CallMethod[Color](t_, objc.GetSelector("foregroundColor"))
	return rv
}

func (t_ TextStorage) SetForegroundColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setForegroundColor:"), value)
}
