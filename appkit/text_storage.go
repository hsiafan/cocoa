// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[TextStorage](t_, "initWithString:", str)
	return rv
}

func (t_ TextStorage) InitWithString_Attributes(str string, attrs map[foundation.AttributedStringKey]objc.IObject) TextStorage {
	rv := ffi.CallMethod[TextStorage](t_, "initWithString:attributes:", str, attrs)
	return rv
}

func (t_ TextStorage) InitWithAttributedString(attrStr foundation.IAttributedString) TextStorage {
	rv := ffi.CallMethod[TextStorage](t_, "initWithAttributedString:", attrStr)
	return rv
}

func (tc _TextStorageClass) Alloc() TextStorage {
	rv := ffi.CallMethod[TextStorage](tc, "alloc")
	return rv
}

func (tc _TextStorageClass) New() TextStorage {
	rv := ffi.CallMethod[TextStorage](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextStorage() TextStorage {
	return TextStorageClass.New()
}

func (t_ TextStorage) Init() TextStorage {
	rv := ffi.CallMethod[TextStorage](t_, "init")
	return rv
}

func (t_ TextStorage) AddLayoutManager(aLayoutManager ILayoutManager) {
	ffi.CallMethod[ffi.Void](t_, "addLayoutManager:", aLayoutManager)
}

func (t_ TextStorage) RemoveLayoutManager(aLayoutManager ILayoutManager) {
	ffi.CallMethod[ffi.Void](t_, "removeLayoutManager:", aLayoutManager)
}

func (t_ TextStorage) Edited_Range_ChangeInLength(editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	ffi.CallMethod[ffi.Void](t_, "edited:range:changeInLength:", editedMask, editedRange, delta)
}

func (t_ TextStorage) ProcessEditing() {
	ffi.CallMethod[ffi.Void](t_, "processEditing")
}

func (t_ TextStorage) InvalidateAttributesInRange(range_ foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "invalidateAttributesInRange:", range_)
}

func (t_ TextStorage) EnsureAttributesAreFixedInRange(range_ foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "ensureAttributesAreFixedInRange:", range_)
}

func (t_ TextStorage) Delegate() TextStorageDelegateWrapper {
	rv := ffi.CallMethod[TextStorageDelegateWrapper](t_, "delegate")
	return rv
}

func (t_ TextStorage) SetDelegate(value TextStorageDelegate) {
	po := ffi.CreateProtocol("NSTextStorageDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](t_, "setDelegate:", po)
}

func (t_ TextStorage) LayoutManagers() []LayoutManager {
	rv := ffi.CallMethod[[]LayoutManager](t_, "layoutManagers")
	return rv
}

func (t_ TextStorage) FixesAttributesLazily() bool {
	rv := ffi.CallMethod[bool](t_, "fixesAttributesLazily")
	return rv
}

func (t_ TextStorage) EditedMask() TextStorageEditActions {
	rv := ffi.CallMethod[TextStorageEditActions](t_, "editedMask")
	return rv
}

func (t_ TextStorage) EditedRange() foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "editedRange")
	return rv
}

func (t_ TextStorage) ChangeInLength() int {
	rv := ffi.CallMethod[int](t_, "changeInLength")
	return rv
}

func (t_ TextStorage) AttributeRuns() []TextStorage {
	rv := ffi.CallMethod[[]TextStorage](t_, "attributeRuns")
	return rv
}

func (t_ TextStorage) SetAttributeRuns(value []ITextStorage) {
	ffi.CallMethod[ffi.Void](t_, "setAttributeRuns:", value)
}

func (t_ TextStorage) Paragraphs() []TextStorage {
	rv := ffi.CallMethod[[]TextStorage](t_, "paragraphs")
	return rv
}

func (t_ TextStorage) SetParagraphs(value []ITextStorage) {
	ffi.CallMethod[ffi.Void](t_, "setParagraphs:", value)
}

func (t_ TextStorage) Words() []TextStorage {
	rv := ffi.CallMethod[[]TextStorage](t_, "words")
	return rv
}

func (t_ TextStorage) SetWords(value []ITextStorage) {
	ffi.CallMethod[ffi.Void](t_, "setWords:", value)
}

func (t_ TextStorage) Characters() []TextStorage {
	rv := ffi.CallMethod[[]TextStorage](t_, "characters")
	return rv
}

func (t_ TextStorage) SetCharacters(value []ITextStorage) {
	ffi.CallMethod[ffi.Void](t_, "setCharacters:", value)
}

func (t_ TextStorage) Font() Font {
	rv := ffi.CallMethod[Font](t_, "font")
	return rv
}

func (t_ TextStorage) SetFont(value IFont) {
	ffi.CallMethod[ffi.Void](t_, "setFont:", value)
}

func (t_ TextStorage) ForegroundColor() Color {
	rv := ffi.CallMethod[Color](t_, "foregroundColor")
	return rv
}

func (t_ TextStorage) SetForegroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](t_, "setForegroundColor:", value)
}
