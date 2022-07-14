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
	InvalidateAttributesInRange(_range foundation.Range)
	EnsureAttributesAreFixedInRange(_range foundation.Range)
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
	rv.Autorelease()
	return rv
}

func (t_ TextStorage) InitWithString_Attributes(str string, attrs map[foundation.AttributedStringKey]objc.IObject) TextStorage {
	rv := ffi.CallMethod[TextStorage](t_, "initWithString:attributes:", str, attrs)
	rv.Autorelease()
	return rv
}

func (t_ TextStorage) InitWithAttributedString(attrStr foundation.IAttributedString) TextStorage {
	rv := ffi.CallMethod[TextStorage](t_, "initWithAttributedString:", attrStr)
	rv.Autorelease()
	return rv
}

func (tc _TextStorageClass) Alloc() TextStorage {
	rv := ffi.CallMethod[TextStorage](tc, "alloc")
	return rv
}

func (t_ TextStorage) Init() TextStorage {
	rv := ffi.CallMethod[TextStorage](t_, "init")
	rv.Autorelease()
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

func (t_ TextStorage) InvalidateAttributesInRange(_range foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "invalidateAttributesInRange:", _range)
}

func (t_ TextStorage) EnsureAttributesAreFixedInRange(_range foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "ensureAttributesAreFixedInRange:", _range)
}

func (t_ TextStorage) Delegate() TextStorageDelegateWrapper {
	rv := ffi.CallMethod[TextStorageDelegateWrapper](t_, "delegate")
	return rv
}

func (t_ TextStorage) SetDelegate(value TextStorageDelegate) {
	po := ffi.CreateProtocol(value)
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

type TextStorageDelegate interface {
	ImplementsTextStorage_WillProcessEditing_Range_ChangeInLength() bool
	// optional
	TextStorage_WillProcessEditing_Range_ChangeInLength(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
	ImplementsTextStorage_DidProcessEditing_Range_ChangeInLength() bool
	// optional
	TextStorage_DidProcessEditing_Range_ChangeInLength(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
}

type TextStorageDelegateImpl struct {
	_TextStorage_WillProcessEditing_Range_ChangeInLength func(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
	_TextStorage_DidProcessEditing_Range_ChangeInLength  func(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
}

func (di *TextStorageDelegateImpl) ImplementsTextStorage_WillProcessEditing_Range_ChangeInLength() bool {
	return di._TextStorage_WillProcessEditing_Range_ChangeInLength != nil
}

func (di *TextStorageDelegateImpl) SetTextStorage_WillProcessEditing_Range_ChangeInLength(f func(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)) {
	di._TextStorage_WillProcessEditing_Range_ChangeInLength = f
}

func (di *TextStorageDelegateImpl) TextStorage_WillProcessEditing_Range_ChangeInLength(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	di._TextStorage_WillProcessEditing_Range_ChangeInLength(textStorage, editedMask, editedRange, delta)
}
func (di *TextStorageDelegateImpl) ImplementsTextStorage_DidProcessEditing_Range_ChangeInLength() bool {
	return di._TextStorage_DidProcessEditing_Range_ChangeInLength != nil
}

func (di *TextStorageDelegateImpl) SetTextStorage_DidProcessEditing_Range_ChangeInLength(f func(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int)) {
	di._TextStorage_DidProcessEditing_Range_ChangeInLength = f
}

func (di *TextStorageDelegateImpl) TextStorage_DidProcessEditing_Range_ChangeInLength(textStorage TextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	di._TextStorage_DidProcessEditing_Range_ChangeInLength(textStorage, editedMask, editedRange, delta)
}

type TextStorageDelegateWrapper struct {
	objc.Object
}

func (t_ *TextStorageDelegateWrapper) ImplementsTextStorage_WillProcessEditing_Range_ChangeInLength() bool {
	return t_.RespondsToSelector(objc.GetSelector("textStorage:willProcessEditing:range:changeInLength:"))
}

func (t_ TextStorageDelegateWrapper) TextStorage_WillProcessEditing_Range_ChangeInLength(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	ffi.CallMethod[ffi.Void](t_, "textStorage:willProcessEditing:range:changeInLength:", textStorage, editedMask, editedRange, delta)
}

func (t_ *TextStorageDelegateWrapper) ImplementsTextStorage_DidProcessEditing_Range_ChangeInLength() bool {
	return t_.RespondsToSelector(objc.GetSelector("textStorage:didProcessEditing:range:changeInLength:"))
}

func (t_ TextStorageDelegateWrapper) TextStorage_DidProcessEditing_Range_ChangeInLength(textStorage ITextStorage, editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	ffi.CallMethod[ffi.Void](t_, "textStorage:didProcessEditing:range:changeInLength:", textStorage, editedMask, editedRange, delta)
}

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
	rv.Autorelease()
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
