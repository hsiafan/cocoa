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
	rv := objc.CallMethod[TextStorage](t_, "initWithString:", str)
	return rv
}

func (t_ TextStorage) InitWithString_Attributes(str string, attrs map[foundation.AttributedStringKey]objc.IObject) TextStorage {
	rv := objc.CallMethod[TextStorage](t_, "initWithString:attributes:", str, attrs)
	return rv
}

func (t_ TextStorage) InitWithAttributedString(attrStr foundation.IAttributedString) TextStorage {
	rv := objc.CallMethod[TextStorage](t_, "initWithAttributedString:", attrStr)
	return rv
}

func (tc _TextStorageClass) Alloc() TextStorage {
	rv := objc.CallMethod[TextStorage](tc, "alloc")
	return rv
}

func (tc _TextStorageClass) New() TextStorage {
	rv := objc.CallMethod[TextStorage](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextStorage() TextStorage {
	return TextStorageClass.New()
}

func (t_ TextStorage) Init() TextStorage {
	rv := objc.CallMethod[TextStorage](t_, "init")
	return rv
}

func (t_ TextStorage) AddLayoutManager(aLayoutManager ILayoutManager) {
	objc.CallMethod[objc.Void](t_, "addLayoutManager:", aLayoutManager)
}

func (t_ TextStorage) RemoveLayoutManager(aLayoutManager ILayoutManager) {
	objc.CallMethod[objc.Void](t_, "removeLayoutManager:", aLayoutManager)
}

func (t_ TextStorage) Edited_Range_ChangeInLength(editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	objc.CallMethod[objc.Void](t_, "edited:range:changeInLength:", editedMask, editedRange, delta)
}

func (t_ TextStorage) ProcessEditing() {
	objc.CallMethod[objc.Void](t_, "processEditing")
}

func (t_ TextStorage) InvalidateAttributesInRange(range_ foundation.Range) {
	objc.CallMethod[objc.Void](t_, "invalidateAttributesInRange:", range_)
}

func (t_ TextStorage) EnsureAttributesAreFixedInRange(range_ foundation.Range) {
	objc.CallMethod[objc.Void](t_, "ensureAttributesAreFixedInRange:", range_)
}

func (t_ TextStorage) Delegate() TextStorageDelegateWrapper {
	rv := objc.CallMethod[TextStorageDelegateWrapper](t_, "delegate")
	return rv
}

func (t_ TextStorage) SetDelegate(value TextStorageDelegate) {
	po := objc.CreateProtocol("NSTextStorageDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](t_, "setDelegate:", po)
}

func (t_ TextStorage) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, "setDelegate:", value)
}

func (t_ TextStorage) LayoutManagers() []LayoutManager {
	rv := objc.CallMethod[[]LayoutManager](t_, "layoutManagers")
	return rv
}

func (t_ TextStorage) FixesAttributesLazily() bool {
	rv := objc.CallMethod[bool](t_, "fixesAttributesLazily")
	return rv
}

func (t_ TextStorage) EditedMask() TextStorageEditActions {
	rv := objc.CallMethod[TextStorageEditActions](t_, "editedMask")
	return rv
}

func (t_ TextStorage) EditedRange() foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, "editedRange")
	return rv
}

func (t_ TextStorage) ChangeInLength() int {
	rv := objc.CallMethod[int](t_, "changeInLength")
	return rv
}

func (t_ TextStorage) AttributeRuns() []TextStorage {
	rv := objc.CallMethod[[]TextStorage](t_, "attributeRuns")
	return rv
}

func (t_ TextStorage) SetAttributeRuns(value []ITextStorage) {
	objc.CallMethod[objc.Void](t_, "setAttributeRuns:", value)
}

func (t_ TextStorage) Paragraphs() []TextStorage {
	rv := objc.CallMethod[[]TextStorage](t_, "paragraphs")
	return rv
}

func (t_ TextStorage) SetParagraphs(value []ITextStorage) {
	objc.CallMethod[objc.Void](t_, "setParagraphs:", value)
}

func (t_ TextStorage) Words() []TextStorage {
	rv := objc.CallMethod[[]TextStorage](t_, "words")
	return rv
}

func (t_ TextStorage) SetWords(value []ITextStorage) {
	objc.CallMethod[objc.Void](t_, "setWords:", value)
}

func (t_ TextStorage) Characters() []TextStorage {
	rv := objc.CallMethod[[]TextStorage](t_, "characters")
	return rv
}

func (t_ TextStorage) SetCharacters(value []ITextStorage) {
	objc.CallMethod[objc.Void](t_, "setCharacters:", value)
}

func (t_ TextStorage) Font() Font {
	rv := objc.CallMethod[Font](t_, "font")
	return rv
}

func (t_ TextStorage) SetFont(value IFont) {
	objc.CallMethod[objc.Void](t_, "setFont:", value)
}

func (t_ TextStorage) ForegroundColor() Color {
	rv := objc.CallMethod[Color](t_, "foregroundColor")
	return rv
}

func (t_ TextStorage) SetForegroundColor(value IColor) {
	objc.CallMethod[objc.Void](t_, "setForegroundColor:", value)
}
