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

type TextViewDelegateWrapper struct {
	objc.Object
}

func (t_ TextViewDelegateWrapper) UndoManagerForTextView(view ITextView) foundation.UndoManager {
	rv := objc.CallMethod[foundation.UndoManager](t_, objc.GetSelector("undoManagerForTextView:"), objc.ExtractPtr(view))
	return rv
}

func (t_ TextViewDelegateWrapper) TextView_WillDisplayToolTip_ForCharacterAtIndex(textView ITextView, tooltip string, characterIndex uint) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("textView:willDisplayToolTip:forCharacterAtIndex:"), objc.ExtractPtr(textView), tooltip, characterIndex)
	return rv
}

func (t_ TextViewDelegateWrapper) TextView_URLForContentsOfTextAttachment_AtIndex(textView ITextView, textAttachment ITextAttachment, charIndex uint) foundation.URL {
	rv := objc.CallMethod[foundation.URL](t_, objc.GetSelector("textView:URLForContentsOfTextAttachment:atIndex:"), objc.ExtractPtr(textView), objc.ExtractPtr(textAttachment), charIndex)
	return rv
}

func (t_ TextViewDelegateWrapper) TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange(textView ITextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("textView:willChangeSelectionFromCharacterRange:toCharacterRange:"), objc.ExtractPtr(textView), oldSelectedCharRange, newSelectedCharRange)
	return rv
}

func (t_ TextViewDelegateWrapper) TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges(textView ITextView, oldSelectedCharRanges []foundation.IValue, newSelectedCharRanges []foundation.IValue) []foundation.Value {
	rv := objc.CallMethod[[]foundation.Value](t_, objc.GetSelector("textView:willChangeSelectionFromCharacterRanges:toCharacterRanges:"), objc.ExtractPtr(textView), oldSelectedCharRanges, newSelectedCharRanges)
	return rv
}

func (t_ TextViewDelegateWrapper) TextViewDidChangeSelection(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textViewDidChangeSelection:"), objc.ExtractPtr(notification))
}

func (t_ TextViewDelegateWrapper) TextView_Candidates_ForSelectedRange(textView ITextView, candidates []foundation.ITextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult {
	rv := objc.CallMethod[[]foundation.TextCheckingResult](t_, objc.GetSelector("textView:candidates:forSelectedRange:"), objc.ExtractPtr(textView), candidates, selectedRange)
	return rv
}

func (t_ TextViewDelegateWrapper) TextView_CandidatesForSelectedRange(textView ITextView, selectedRange foundation.Range) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("textView:candidatesForSelectedRange:"), objc.ExtractPtr(textView), selectedRange)
	return rv
}

func (t_ TextViewDelegateWrapper) TextView_ShouldSelectCandidateAtIndex(textView ITextView, index uint) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textView:shouldSelectCandidateAtIndex:"), objc.ExtractPtr(textView), index)
	return rv
}

func (t_ TextViewDelegateWrapper) TextView_ShouldUpdateTouchBarItemIdentifiers(textView ITextView, identifiers []TouchBarItemIdentifier) []TouchBarItemIdentifier {
	rv := objc.CallMethod[[]TouchBarItemIdentifier](t_, objc.GetSelector("textView:shouldUpdateTouchBarItemIdentifiers:"), objc.ExtractPtr(textView), identifiers)
	return rv
}

func (t_ TextViewDelegateWrapper) TextView_ShouldChangeTextInRange_ReplacementString(textView ITextView, affectedCharRange foundation.Range, replacementString string) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textView:shouldChangeTextInRange:replacementString:"), objc.ExtractPtr(textView), affectedCharRange, replacementString)
	return rv
}

func (t_ TextViewDelegateWrapper) TextView_ShouldChangeTextInRanges_ReplacementStrings(textView ITextView, affectedRanges []foundation.IValue, replacementStrings []string) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textView:shouldChangeTextInRanges:replacementStrings:"), objc.ExtractPtr(textView), affectedRanges, replacementStrings)
	return rv
}

func (t_ TextViewDelegateWrapper) TextView_ShouldChangeTypingAttributes_ToAttributes(textView ITextView, oldTypingAttributes map[string]objc.IObject, newTypingAttributes map[foundation.AttributedStringKey]objc.IObject) map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, objc.GetSelector("textView:shouldChangeTypingAttributes:toAttributes:"), objc.ExtractPtr(textView), oldTypingAttributes, newTypingAttributes)
	return rv
}

func (t_ TextViewDelegateWrapper) TextViewDidChangeTypingAttributes(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textViewDidChangeTypingAttributes:"), objc.ExtractPtr(notification))
}

func (t_ TextViewDelegateWrapper) TextView_ClickedOnLink_AtIndex(textView ITextView, link objc.IObject, charIndex uint) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textView:clickedOnLink:atIndex:"), objc.ExtractPtr(textView), objc.ExtractPtr(link), charIndex)
	return rv
}

func (t_ TextViewDelegateWrapper) TextView_ShouldSetSpellingState_Range(textView ITextView, value int, affectedCharRange foundation.Range) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("textView:shouldSetSpellingState:range:"), objc.ExtractPtr(textView), value, affectedCharRange)
	return rv
}

func (t_ TextViewDelegateWrapper) TextView_WillCheckTextInRange_Options_Types(view ITextView, range_ foundation.Range, options map[TextCheckingOptionKey]objc.IObject, checkingTypes *foundation.TextCheckingTypes) map[TextCheckingOptionKey]objc.Object {
	rv := objc.CallMethod[map[TextCheckingOptionKey]objc.Object](t_, objc.GetSelector("textView:willCheckTextInRange:options:types:"), objc.ExtractPtr(view), range_, options, checkingTypes)
	return rv
}

func (t_ TextViewDelegateWrapper) TextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount(view ITextView, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject, results []foundation.ITextCheckingResult, orthography foundation.IOrthography, wordCount int) []foundation.TextCheckingResult {
	rv := objc.CallMethod[[]foundation.TextCheckingResult](t_, objc.GetSelector("textView:didCheckTextInRange:types:options:results:orthography:wordCount:"), objc.ExtractPtr(view), range_, checkingTypes, options, results, objc.ExtractPtr(orthography), wordCount)
	return rv
}

func (t_ TextViewDelegateWrapper) TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(textView ITextView, words []string, charRange foundation.Range, index *int) []string {
	rv := objc.CallMethod[[]string](t_, objc.GetSelector("textView:completions:forPartialWordRange:indexOfSelectedItem:"), objc.ExtractPtr(textView), words, charRange, index)
	return rv
}

func (t_ TextViewDelegateWrapper) TextView_WillShowSharingServicePicker_ForItems(textView ITextView, servicePicker ISharingServicePicker, items []objc.IObject) SharingServicePicker {
	rv := objc.CallMethod[SharingServicePicker](t_, objc.GetSelector("textView:willShowSharingServicePicker:forItems:"), objc.ExtractPtr(textView), objc.ExtractPtr(servicePicker), items)
	return rv
}

func (t_ TextViewDelegateWrapper) TextView_DoCommandBySelector(textView ITextView, commandSelector objc.Selector) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textView:doCommandBySelector:"), objc.ExtractPtr(textView), commandSelector)
	return rv
}

func (t_ TextViewDelegateWrapper) TextView_Menu_ForEvent_AtIndex(view ITextView, menu IMenu, event IEvent, charIndex uint) Menu {
	rv := objc.CallMethod[Menu](t_, objc.GetSelector("textView:menu:forEvent:atIndex:"), objc.ExtractPtr(view), objc.ExtractPtr(menu), objc.ExtractPtr(event), charIndex)
	return rv
}

// deprecated
func (t_ TextViewDelegateWrapper) TextView_ClickedOnLink(textView ITextView, link objc.IObject) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textView:clickedOnLink:"), objc.ExtractPtr(textView), objc.ExtractPtr(link))
	return rv
}

func (t_ TextViewDelegateWrapper) TextDidChange(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textDidChange:"), objc.ExtractPtr(notification))
}

func (t_ TextViewDelegateWrapper) TextShouldBeginEditing(textObject IText) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textShouldBeginEditing:"), objc.ExtractPtr(textObject))
	return rv
}

func (t_ TextViewDelegateWrapper) TextDidBeginEditing(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textDidBeginEditing:"), objc.ExtractPtr(notification))
}

func (t_ TextViewDelegateWrapper) TextShouldEndEditing(textObject IText) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textShouldEndEditing:"), objc.ExtractPtr(textObject))
	return rv
}

func (t_ TextViewDelegateWrapper) TextDidEndEditing(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textDidEndEditing:"), objc.ExtractPtr(notification))
}
