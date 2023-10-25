// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type TextViewDelegate interface {
	TextDelegate
	ImplementsUndoManagerForTextView() bool
	// optional
	UndoManagerForTextView(view TextView) foundation.IUndoManager
	ImplementsTextView_WillDisplayToolTip_ForCharacterAtIndex() bool
	// optional
	TextView_WillDisplayToolTip_ForCharacterAtIndex(textView TextView, tooltip string, characterIndex uint) string
	ImplementsTextView_URLForContentsOfTextAttachment_AtIndex() bool
	// optional
	TextView_URLForContentsOfTextAttachment_AtIndex(textView TextView, textAttachment TextAttachment, charIndex uint) foundation.IURL
	ImplementsTextView_WillChangeSelectionFromCharacterRange_ToCharacterRange() bool
	// optional
	TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange(textView TextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range
	ImplementsTextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges() bool
	// optional
	TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges(textView TextView, oldSelectedCharRanges []foundation.Value, newSelectedCharRanges []foundation.Value) []foundation.IValue
	ImplementsTextViewDidChangeSelection() bool
	// optional
	TextViewDidChangeSelection(notification foundation.Notification)
	ImplementsTextView_Candidates_ForSelectedRange() bool
	// optional
	TextView_Candidates_ForSelectedRange(textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult
	ImplementsTextView_CandidatesForSelectedRange() bool
	// optional
	TextView_CandidatesForSelectedRange(textView TextView, selectedRange foundation.Range) []objc.IObject
	ImplementsTextView_ShouldSelectCandidateAtIndex() bool
	// optional
	TextView_ShouldSelectCandidateAtIndex(textView TextView, index uint) bool
	ImplementsTextView_ShouldUpdateTouchBarItemIdentifiers() bool
	// optional
	TextView_ShouldUpdateTouchBarItemIdentifiers(textView TextView, identifiers []TouchBarItemIdentifier) []TouchBarItemIdentifier
	ImplementsTextView_ShouldChangeTextInRange_ReplacementString() bool
	// optional
	TextView_ShouldChangeTextInRange_ReplacementString(textView TextView, affectedCharRange foundation.Range, replacementString string) bool
	ImplementsTextView_ShouldChangeTextInRanges_ReplacementStrings() bool
	// optional
	TextView_ShouldChangeTextInRanges_ReplacementStrings(textView TextView, affectedRanges []foundation.Value, replacementStrings []string) bool
	ImplementsTextView_ShouldChangeTypingAttributes_ToAttributes() bool
	// optional
	TextView_ShouldChangeTypingAttributes_ToAttributes(textView TextView, oldTypingAttributes map[string]objc.Object, newTypingAttributes map[foundation.AttributedStringKey]objc.Object) map[foundation.AttributedStringKey]objc.IObject
	ImplementsTextViewDidChangeTypingAttributes() bool
	// optional
	TextViewDidChangeTypingAttributes(notification foundation.Notification)
	ImplementsTextView_ClickedOnLink_AtIndex() bool
	// optional
	TextView_ClickedOnLink_AtIndex(textView TextView, link objc.Object, charIndex uint) bool
	ImplementsTextView_ShouldSetSpellingState_Range() bool
	// optional
	TextView_ShouldSetSpellingState_Range(textView TextView, value int, affectedCharRange foundation.Range) int
	ImplementsTextView_WillCheckTextInRange_Options_Types() bool
	// optional
	TextView_WillCheckTextInRange_Options_Types(view TextView, range_ foundation.Range, options map[TextCheckingOptionKey]objc.Object, checkingTypes *foundation.TextCheckingTypes) map[TextCheckingOptionKey]objc.IObject
	ImplementsTextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount() bool
	// optional
	TextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount(view TextView, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.Object, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.ITextCheckingResult
	ImplementsTextView_Completions_ForPartialWordRange_IndexOfSelectedItem() bool
	// optional
	TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(textView TextView, words []string, charRange foundation.Range, index *int) []string
	ImplementsTextView_WillShowSharingServicePicker_ForItems() bool
	// optional
	TextView_WillShowSharingServicePicker_ForItems(textView TextView, servicePicker SharingServicePicker, items []objc.Object) ISharingServicePicker
	ImplementsTextView_DoCommandBySelector() bool
	// optional
	TextView_DoCommandBySelector(textView TextView, commandSelector objc.Selector) bool
	ImplementsTextView_Menu_ForEvent_AtIndex() bool
	// optional
	TextView_Menu_ForEvent_AtIndex(view TextView, menu Menu, event Event, charIndex uint) IMenu
	ImplementsTextView_ClickedOnLink() bool
	// optional
	// deprecated
	TextView_ClickedOnLink(textView TextView, link objc.Object) bool
}

func WrapTextViewDelegate(v TextViewDelegate) objc.Object {
	return objc.WrapAsProtocol("NSTextViewDelegate", v)
}

type TextViewDelegateBase struct {
	TextDelegateBase
}

func (p *TextViewDelegateBase) ImplementsUndoManagerForTextView() bool {
	return false
}

func (p *TextViewDelegateBase) UndoManagerForTextView(view TextView) foundation.IUndoManager {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_WillDisplayToolTip_ForCharacterAtIndex() bool {
	return false
}

func (p *TextViewDelegateBase) TextView_WillDisplayToolTip_ForCharacterAtIndex(textView TextView, tooltip string, characterIndex uint) string {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_URLForContentsOfTextAttachment_AtIndex() bool {
	return false
}

func (p *TextViewDelegateBase) TextView_URLForContentsOfTextAttachment_AtIndex(textView TextView, textAttachment TextAttachment, charIndex uint) foundation.IURL {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_WillChangeSelectionFromCharacterRange_ToCharacterRange() bool {
	return false
}

func (p *TextViewDelegateBase) TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange(textView TextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges() bool {
	return false
}

func (p *TextViewDelegateBase) TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges(textView TextView, oldSelectedCharRanges []foundation.Value, newSelectedCharRanges []foundation.Value) []foundation.IValue {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextViewDidChangeSelection() bool {
	return false
}

func (p *TextViewDelegateBase) TextViewDidChangeSelection(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_Candidates_ForSelectedRange() bool {
	return false
}

func (p *TextViewDelegateBase) TextView_Candidates_ForSelectedRange(textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_CandidatesForSelectedRange() bool {
	return false
}

func (p *TextViewDelegateBase) TextView_CandidatesForSelectedRange(textView TextView, selectedRange foundation.Range) []objc.IObject {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_ShouldSelectCandidateAtIndex() bool {
	return false
}

func (p *TextViewDelegateBase) TextView_ShouldSelectCandidateAtIndex(textView TextView, index uint) bool {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_ShouldUpdateTouchBarItemIdentifiers() bool {
	return false
}

func (p *TextViewDelegateBase) TextView_ShouldUpdateTouchBarItemIdentifiers(textView TextView, identifiers []TouchBarItemIdentifier) []TouchBarItemIdentifier {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_ShouldChangeTextInRange_ReplacementString() bool {
	return false
}

func (p *TextViewDelegateBase) TextView_ShouldChangeTextInRange_ReplacementString(textView TextView, affectedCharRange foundation.Range, replacementString string) bool {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_ShouldChangeTextInRanges_ReplacementStrings() bool {
	return false
}

func (p *TextViewDelegateBase) TextView_ShouldChangeTextInRanges_ReplacementStrings(textView TextView, affectedRanges []foundation.Value, replacementStrings []string) bool {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_ShouldChangeTypingAttributes_ToAttributes() bool {
	return false
}

func (p *TextViewDelegateBase) TextView_ShouldChangeTypingAttributes_ToAttributes(textView TextView, oldTypingAttributes map[string]objc.Object, newTypingAttributes map[foundation.AttributedStringKey]objc.Object) map[foundation.AttributedStringKey]objc.IObject {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextViewDidChangeTypingAttributes() bool {
	return false
}

func (p *TextViewDelegateBase) TextViewDidChangeTypingAttributes(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_ClickedOnLink_AtIndex() bool {
	return false
}

func (p *TextViewDelegateBase) TextView_ClickedOnLink_AtIndex(textView TextView, link objc.Object, charIndex uint) bool {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_ShouldSetSpellingState_Range() bool {
	return false
}

func (p *TextViewDelegateBase) TextView_ShouldSetSpellingState_Range(textView TextView, value int, affectedCharRange foundation.Range) int {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_WillCheckTextInRange_Options_Types() bool {
	return false
}

func (p *TextViewDelegateBase) TextView_WillCheckTextInRange_Options_Types(view TextView, range_ foundation.Range, options map[TextCheckingOptionKey]objc.Object, checkingTypes *foundation.TextCheckingTypes) map[TextCheckingOptionKey]objc.IObject {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount() bool {
	return false
}

func (p *TextViewDelegateBase) TextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount(view TextView, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.Object, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.ITextCheckingResult {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_Completions_ForPartialWordRange_IndexOfSelectedItem() bool {
	return false
}

func (p *TextViewDelegateBase) TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(textView TextView, words []string, charRange foundation.Range, index *int) []string {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_WillShowSharingServicePicker_ForItems() bool {
	return false
}

func (p *TextViewDelegateBase) TextView_WillShowSharingServicePicker_ForItems(textView TextView, servicePicker SharingServicePicker, items []objc.Object) ISharingServicePicker {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_DoCommandBySelector() bool {
	return false
}

func (p *TextViewDelegateBase) TextView_DoCommandBySelector(textView TextView, commandSelector objc.Selector) bool {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_Menu_ForEvent_AtIndex() bool {
	return false
}

func (p *TextViewDelegateBase) TextView_Menu_ForEvent_AtIndex(view TextView, menu Menu, event Event, charIndex uint) IMenu {
	panic("unimpemented protocol method")
}

func (p *TextViewDelegateBase) ImplementsTextView_ClickedOnLink() bool {
	return false
}

// deprecated
func (p *TextViewDelegateBase) TextView_ClickedOnLink(textView TextView, link objc.Object) bool {
	panic("unimpemented protocol method")
}

type TextViewDelegateCreator struct {
	className string
	class     objc.Class
}

func NewTextViewDelegateCreator(name string) *TextViewDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("ProtocolBase"), name, 0)
	objc.RegisterClassPair(class)
	return &TextViewDelegateCreator{className: name, class: class}
}

func (c *TextViewDelegateCreator) SetUndoManagerForTextView(handle func(o objc.ProtocolBase, view TextView) foundation.IUndoManager) {
	objc.AddMethod(c.class, objc.SEL("undoManagerForTextView:"), handle)
}

func (c *TextViewDelegateCreator) SetTextView_WillDisplayToolTip_ForCharacterAtIndex(handle func(o objc.ProtocolBase, textView TextView, tooltip string, characterIndex uint) string) {
	objc.AddMethod(c.class, objc.SEL("textView:willDisplayToolTip:forCharacterAtIndex:"), handle)
}

func (c *TextViewDelegateCreator) SetTextView_URLForContentsOfTextAttachment_AtIndex(handle func(o objc.ProtocolBase, textView TextView, textAttachment TextAttachment, charIndex uint) foundation.IURL) {
	objc.AddMethod(c.class, objc.SEL("textView:URLForContentsOfTextAttachment:atIndex:"), handle)
}

func (c *TextViewDelegateCreator) SetTextView_WillChangeSelectionFromCharacterRange_ToCharacterRange(handle func(o objc.ProtocolBase, textView TextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range) {
	objc.AddMethod(c.class, objc.SEL("textView:willChangeSelectionFromCharacterRange:toCharacterRange:"), handle)
}

func (c *TextViewDelegateCreator) SetTextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges(handle func(o objc.ProtocolBase, textView TextView, oldSelectedCharRanges []foundation.Value, newSelectedCharRanges []foundation.Value) []foundation.IValue) {
	objc.AddMethod(c.class, objc.SEL("textView:willChangeSelectionFromCharacterRanges:toCharacterRanges:"), handle)
}

func (c *TextViewDelegateCreator) SetTextViewDidChangeSelection(handle func(o objc.ProtocolBase, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("textViewDidChangeSelection:"), handle)
}

func (c *TextViewDelegateCreator) SetTextView_Candidates_ForSelectedRange(handle func(o objc.ProtocolBase, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult) {
	objc.AddMethod(c.class, objc.SEL("textView:candidates:forSelectedRange:"), handle)
}

func (c *TextViewDelegateCreator) SetTextView_CandidatesForSelectedRange(handle func(o objc.ProtocolBase, textView TextView, selectedRange foundation.Range) []objc.IObject) {
	objc.AddMethod(c.class, objc.SEL("textView:candidatesForSelectedRange:"), handle)
}

func (c *TextViewDelegateCreator) SetTextView_ShouldSelectCandidateAtIndex(handle func(o objc.ProtocolBase, textView TextView, index uint) bool) {
	objc.AddMethod(c.class, objc.SEL("textView:shouldSelectCandidateAtIndex:"), handle)
}

func (c *TextViewDelegateCreator) SetTextView_ShouldUpdateTouchBarItemIdentifiers(handle func(o objc.ProtocolBase, textView TextView, identifiers []TouchBarItemIdentifier) []TouchBarItemIdentifier) {
	objc.AddMethod(c.class, objc.SEL("textView:shouldUpdateTouchBarItemIdentifiers:"), handle)
}

func (c *TextViewDelegateCreator) SetTextView_ShouldChangeTextInRange_ReplacementString(handle func(o objc.ProtocolBase, textView TextView, affectedCharRange foundation.Range, replacementString string) bool) {
	objc.AddMethod(c.class, objc.SEL("textView:shouldChangeTextInRange:replacementString:"), handle)
}

func (c *TextViewDelegateCreator) SetTextView_ShouldChangeTextInRanges_ReplacementStrings(handle func(o objc.ProtocolBase, textView TextView, affectedRanges []foundation.Value, replacementStrings []string) bool) {
	objc.AddMethod(c.class, objc.SEL("textView:shouldChangeTextInRanges:replacementStrings:"), handle)
}

func (c *TextViewDelegateCreator) SetTextView_ShouldChangeTypingAttributes_ToAttributes(handle func(o objc.ProtocolBase, textView TextView, oldTypingAttributes map[string]objc.Object, newTypingAttributes map[foundation.AttributedStringKey]objc.Object) map[foundation.AttributedStringKey]objc.IObject) {
	objc.AddMethod(c.class, objc.SEL("textView:shouldChangeTypingAttributes:toAttributes:"), handle)
}

func (c *TextViewDelegateCreator) SetTextViewDidChangeTypingAttributes(handle func(o objc.ProtocolBase, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("textViewDidChangeTypingAttributes:"), handle)
}

func (c *TextViewDelegateCreator) SetTextView_ClickedOnLink_AtIndex(handle func(o objc.ProtocolBase, textView TextView, link objc.Object, charIndex uint) bool) {
	objc.AddMethod(c.class, objc.SEL("textView:clickedOnLink:atIndex:"), handle)
}

func (c *TextViewDelegateCreator) SetTextView_ShouldSetSpellingState_Range(handle func(o objc.ProtocolBase, textView TextView, value int, affectedCharRange foundation.Range) int) {
	objc.AddMethod(c.class, objc.SEL("textView:shouldSetSpellingState:range:"), handle)
}

func (c *TextViewDelegateCreator) SetTextView_WillCheckTextInRange_Options_Types(handle func(o objc.ProtocolBase, view TextView, range_ foundation.Range, options map[TextCheckingOptionKey]objc.Object, checkingTypes *foundation.TextCheckingTypes) map[TextCheckingOptionKey]objc.IObject) {
	objc.AddMethod(c.class, objc.SEL("textView:willCheckTextInRange:options:types:"), handle)
}

func (c *TextViewDelegateCreator) SetTextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount(handle func(o objc.ProtocolBase, view TextView, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.Object, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.ITextCheckingResult) {
	objc.AddMethod(c.class, objc.SEL("textView:didCheckTextInRange:types:options:results:orthography:wordCount:"), handle)
}

func (c *TextViewDelegateCreator) SetTextView_Completions_ForPartialWordRange_IndexOfSelectedItem(handle func(o objc.ProtocolBase, textView TextView, words []string, charRange foundation.Range, index *int) []string) {
	objc.AddMethod(c.class, objc.SEL("textView:completions:forPartialWordRange:indexOfSelectedItem:"), handle)
}

func (c *TextViewDelegateCreator) SetTextView_WillShowSharingServicePicker_ForItems(handle func(o objc.ProtocolBase, textView TextView, servicePicker SharingServicePicker, items []objc.Object) ISharingServicePicker) {
	objc.AddMethod(c.class, objc.SEL("textView:willShowSharingServicePicker:forItems:"), handle)
}

func (c *TextViewDelegateCreator) SetTextView_DoCommandBySelector(handle func(o objc.ProtocolBase, textView TextView, commandSelector objc.Selector) bool) {
	objc.AddMethod(c.class, objc.SEL("textView:doCommandBySelector:"), handle)
}

func (c *TextViewDelegateCreator) SetTextView_Menu_ForEvent_AtIndex(handle func(o objc.ProtocolBase, view TextView, menu Menu, event Event, charIndex uint) IMenu) {
	objc.AddMethod(c.class, objc.SEL("textView:menu:forEvent:atIndex:"), handle)
}

// deprecated
func (c *TextViewDelegateCreator) SetTextView_ClickedOnLink(handle func(o objc.ProtocolBase, textView TextView, link objc.Object) bool) {
	objc.AddMethod(c.class, objc.SEL("textView:clickedOnLink:"), handle)
}

func (c *TextViewDelegateCreator) Create() objc.ProtocolBase {
	return objc.ProtocolBase{Object: c.class.CreateInstance(0)}
}
