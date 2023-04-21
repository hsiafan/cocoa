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

type TextViewDelegateImpl struct {
	TextDelegateImpl
	_UndoManagerForTextView                                                   func(view TextView) foundation.IUndoManager
	_TextView_WillDisplayToolTip_ForCharacterAtIndex                          func(textView TextView, tooltip string, characterIndex uint) string
	_TextView_URLForContentsOfTextAttachment_AtIndex                          func(textView TextView, textAttachment TextAttachment, charIndex uint) foundation.IURL
	_TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange          func(textView TextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range
	_TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges        func(textView TextView, oldSelectedCharRanges []foundation.Value, newSelectedCharRanges []foundation.Value) []foundation.IValue
	_TextViewDidChangeSelection                                               func(notification foundation.Notification)
	_TextView_Candidates_ForSelectedRange                                     func(textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult
	_TextView_CandidatesForSelectedRange                                      func(textView TextView, selectedRange foundation.Range) []objc.IObject
	_TextView_ShouldSelectCandidateAtIndex                                    func(textView TextView, index uint) bool
	_TextView_ShouldUpdateTouchBarItemIdentifiers                             func(textView TextView, identifiers []TouchBarItemIdentifier) []TouchBarItemIdentifier
	_TextView_ShouldChangeTextInRange_ReplacementString                       func(textView TextView, affectedCharRange foundation.Range, replacementString string) bool
	_TextView_ShouldChangeTextInRanges_ReplacementStrings                     func(textView TextView, affectedRanges []foundation.Value, replacementStrings []string) bool
	_TextView_ShouldChangeTypingAttributes_ToAttributes                       func(textView TextView, oldTypingAttributes map[string]objc.Object, newTypingAttributes map[foundation.AttributedStringKey]objc.Object) map[foundation.AttributedStringKey]objc.IObject
	_TextViewDidChangeTypingAttributes                                        func(notification foundation.Notification)
	_TextView_ClickedOnLink_AtIndex                                           func(textView TextView, link objc.Object, charIndex uint) bool
	_TextView_ShouldSetSpellingState_Range                                    func(textView TextView, value int, affectedCharRange foundation.Range) int
	_TextView_WillCheckTextInRange_Options_Types                              func(view TextView, range_ foundation.Range, options map[TextCheckingOptionKey]objc.Object, checkingTypes *foundation.TextCheckingTypes) map[TextCheckingOptionKey]objc.IObject
	_TextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount func(view TextView, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.Object, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.ITextCheckingResult
	_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem             func(textView TextView, words []string, charRange foundation.Range, index *int) []string
	_TextView_WillShowSharingServicePicker_ForItems                           func(textView TextView, servicePicker SharingServicePicker, items []objc.Object) ISharingServicePicker
	_TextView_DoCommandBySelector                                             func(textView TextView, commandSelector objc.Selector) bool
	_TextView_Menu_ForEvent_AtIndex                                           func(view TextView, menu Menu, event Event, charIndex uint) IMenu
	_TextView_ClickedOnLink                                                   func(textView TextView, link objc.Object) bool
}

func (di *TextViewDelegateImpl) ImplementsUndoManagerForTextView() bool {
	return di._UndoManagerForTextView != nil
}

func (di *TextViewDelegateImpl) SetUndoManagerForTextView(f func(view TextView) foundation.IUndoManager) {
	di._UndoManagerForTextView = f
}

func (di *TextViewDelegateImpl) UndoManagerForTextView(view TextView) foundation.IUndoManager {
	return di._UndoManagerForTextView(view)
}
func (di *TextViewDelegateImpl) ImplementsTextView_WillDisplayToolTip_ForCharacterAtIndex() bool {
	return di._TextView_WillDisplayToolTip_ForCharacterAtIndex != nil
}

func (di *TextViewDelegateImpl) SetTextView_WillDisplayToolTip_ForCharacterAtIndex(f func(textView TextView, tooltip string, characterIndex uint) string) {
	di._TextView_WillDisplayToolTip_ForCharacterAtIndex = f
}

func (di *TextViewDelegateImpl) TextView_WillDisplayToolTip_ForCharacterAtIndex(textView TextView, tooltip string, characterIndex uint) string {
	return di._TextView_WillDisplayToolTip_ForCharacterAtIndex(textView, tooltip, characterIndex)
}
func (di *TextViewDelegateImpl) ImplementsTextView_URLForContentsOfTextAttachment_AtIndex() bool {
	return di._TextView_URLForContentsOfTextAttachment_AtIndex != nil
}

func (di *TextViewDelegateImpl) SetTextView_URLForContentsOfTextAttachment_AtIndex(f func(textView TextView, textAttachment TextAttachment, charIndex uint) foundation.IURL) {
	di._TextView_URLForContentsOfTextAttachment_AtIndex = f
}

func (di *TextViewDelegateImpl) TextView_URLForContentsOfTextAttachment_AtIndex(textView TextView, textAttachment TextAttachment, charIndex uint) foundation.IURL {
	return di._TextView_URLForContentsOfTextAttachment_AtIndex(textView, textAttachment, charIndex)
}
func (di *TextViewDelegateImpl) ImplementsTextView_WillChangeSelectionFromCharacterRange_ToCharacterRange() bool {
	return di._TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange != nil
}

func (di *TextViewDelegateImpl) SetTextView_WillChangeSelectionFromCharacterRange_ToCharacterRange(f func(textView TextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range) {
	di._TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange = f
}

func (di *TextViewDelegateImpl) TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange(textView TextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range {
	return di._TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange(textView, oldSelectedCharRange, newSelectedCharRange)
}
func (di *TextViewDelegateImpl) ImplementsTextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges() bool {
	return di._TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges != nil
}

func (di *TextViewDelegateImpl) SetTextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges(f func(textView TextView, oldSelectedCharRanges []foundation.Value, newSelectedCharRanges []foundation.Value) []foundation.IValue) {
	di._TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges = f
}

func (di *TextViewDelegateImpl) TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges(textView TextView, oldSelectedCharRanges []foundation.Value, newSelectedCharRanges []foundation.Value) []foundation.IValue {
	return di._TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges(textView, oldSelectedCharRanges, newSelectedCharRanges)
}
func (di *TextViewDelegateImpl) ImplementsTextViewDidChangeSelection() bool {
	return di._TextViewDidChangeSelection != nil
}

func (di *TextViewDelegateImpl) SetTextViewDidChangeSelection(f func(notification foundation.Notification)) {
	di._TextViewDidChangeSelection = f
}

func (di *TextViewDelegateImpl) TextViewDidChangeSelection(notification foundation.Notification) {
	di._TextViewDidChangeSelection(notification)
}
func (di *TextViewDelegateImpl) ImplementsTextView_Candidates_ForSelectedRange() bool {
	return di._TextView_Candidates_ForSelectedRange != nil
}

func (di *TextViewDelegateImpl) SetTextView_Candidates_ForSelectedRange(f func(textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult) {
	di._TextView_Candidates_ForSelectedRange = f
}

func (di *TextViewDelegateImpl) TextView_Candidates_ForSelectedRange(textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult {
	return di._TextView_Candidates_ForSelectedRange(textView, candidates, selectedRange)
}
func (di *TextViewDelegateImpl) ImplementsTextView_CandidatesForSelectedRange() bool {
	return di._TextView_CandidatesForSelectedRange != nil
}

func (di *TextViewDelegateImpl) SetTextView_CandidatesForSelectedRange(f func(textView TextView, selectedRange foundation.Range) []objc.IObject) {
	di._TextView_CandidatesForSelectedRange = f
}

func (di *TextViewDelegateImpl) TextView_CandidatesForSelectedRange(textView TextView, selectedRange foundation.Range) []objc.IObject {
	return di._TextView_CandidatesForSelectedRange(textView, selectedRange)
}
func (di *TextViewDelegateImpl) ImplementsTextView_ShouldSelectCandidateAtIndex() bool {
	return di._TextView_ShouldSelectCandidateAtIndex != nil
}

func (di *TextViewDelegateImpl) SetTextView_ShouldSelectCandidateAtIndex(f func(textView TextView, index uint) bool) {
	di._TextView_ShouldSelectCandidateAtIndex = f
}

func (di *TextViewDelegateImpl) TextView_ShouldSelectCandidateAtIndex(textView TextView, index uint) bool {
	return di._TextView_ShouldSelectCandidateAtIndex(textView, index)
}
func (di *TextViewDelegateImpl) ImplementsTextView_ShouldUpdateTouchBarItemIdentifiers() bool {
	return di._TextView_ShouldUpdateTouchBarItemIdentifiers != nil
}

func (di *TextViewDelegateImpl) SetTextView_ShouldUpdateTouchBarItemIdentifiers(f func(textView TextView, identifiers []TouchBarItemIdentifier) []TouchBarItemIdentifier) {
	di._TextView_ShouldUpdateTouchBarItemIdentifiers = f
}

func (di *TextViewDelegateImpl) TextView_ShouldUpdateTouchBarItemIdentifiers(textView TextView, identifiers []TouchBarItemIdentifier) []TouchBarItemIdentifier {
	return di._TextView_ShouldUpdateTouchBarItemIdentifiers(textView, identifiers)
}
func (di *TextViewDelegateImpl) ImplementsTextView_ShouldChangeTextInRange_ReplacementString() bool {
	return di._TextView_ShouldChangeTextInRange_ReplacementString != nil
}

func (di *TextViewDelegateImpl) SetTextView_ShouldChangeTextInRange_ReplacementString(f func(textView TextView, affectedCharRange foundation.Range, replacementString string) bool) {
	di._TextView_ShouldChangeTextInRange_ReplacementString = f
}

func (di *TextViewDelegateImpl) TextView_ShouldChangeTextInRange_ReplacementString(textView TextView, affectedCharRange foundation.Range, replacementString string) bool {
	return di._TextView_ShouldChangeTextInRange_ReplacementString(textView, affectedCharRange, replacementString)
}
func (di *TextViewDelegateImpl) ImplementsTextView_ShouldChangeTextInRanges_ReplacementStrings() bool {
	return di._TextView_ShouldChangeTextInRanges_ReplacementStrings != nil
}

func (di *TextViewDelegateImpl) SetTextView_ShouldChangeTextInRanges_ReplacementStrings(f func(textView TextView, affectedRanges []foundation.Value, replacementStrings []string) bool) {
	di._TextView_ShouldChangeTextInRanges_ReplacementStrings = f
}

func (di *TextViewDelegateImpl) TextView_ShouldChangeTextInRanges_ReplacementStrings(textView TextView, affectedRanges []foundation.Value, replacementStrings []string) bool {
	return di._TextView_ShouldChangeTextInRanges_ReplacementStrings(textView, affectedRanges, replacementStrings)
}
func (di *TextViewDelegateImpl) ImplementsTextView_ShouldChangeTypingAttributes_ToAttributes() bool {
	return di._TextView_ShouldChangeTypingAttributes_ToAttributes != nil
}

func (di *TextViewDelegateImpl) SetTextView_ShouldChangeTypingAttributes_ToAttributes(f func(textView TextView, oldTypingAttributes map[string]objc.Object, newTypingAttributes map[foundation.AttributedStringKey]objc.Object) map[foundation.AttributedStringKey]objc.IObject) {
	di._TextView_ShouldChangeTypingAttributes_ToAttributes = f
}

func (di *TextViewDelegateImpl) TextView_ShouldChangeTypingAttributes_ToAttributes(textView TextView, oldTypingAttributes map[string]objc.Object, newTypingAttributes map[foundation.AttributedStringKey]objc.Object) map[foundation.AttributedStringKey]objc.IObject {
	return di._TextView_ShouldChangeTypingAttributes_ToAttributes(textView, oldTypingAttributes, newTypingAttributes)
}
func (di *TextViewDelegateImpl) ImplementsTextViewDidChangeTypingAttributes() bool {
	return di._TextViewDidChangeTypingAttributes != nil
}

func (di *TextViewDelegateImpl) SetTextViewDidChangeTypingAttributes(f func(notification foundation.Notification)) {
	di._TextViewDidChangeTypingAttributes = f
}

func (di *TextViewDelegateImpl) TextViewDidChangeTypingAttributes(notification foundation.Notification) {
	di._TextViewDidChangeTypingAttributes(notification)
}
func (di *TextViewDelegateImpl) ImplementsTextView_ClickedOnLink_AtIndex() bool {
	return di._TextView_ClickedOnLink_AtIndex != nil
}

func (di *TextViewDelegateImpl) SetTextView_ClickedOnLink_AtIndex(f func(textView TextView, link objc.Object, charIndex uint) bool) {
	di._TextView_ClickedOnLink_AtIndex = f
}

func (di *TextViewDelegateImpl) TextView_ClickedOnLink_AtIndex(textView TextView, link objc.Object, charIndex uint) bool {
	return di._TextView_ClickedOnLink_AtIndex(textView, link, charIndex)
}
func (di *TextViewDelegateImpl) ImplementsTextView_ShouldSetSpellingState_Range() bool {
	return di._TextView_ShouldSetSpellingState_Range != nil
}

func (di *TextViewDelegateImpl) SetTextView_ShouldSetSpellingState_Range(f func(textView TextView, value int, affectedCharRange foundation.Range) int) {
	di._TextView_ShouldSetSpellingState_Range = f
}

func (di *TextViewDelegateImpl) TextView_ShouldSetSpellingState_Range(textView TextView, value int, affectedCharRange foundation.Range) int {
	return di._TextView_ShouldSetSpellingState_Range(textView, value, affectedCharRange)
}
func (di *TextViewDelegateImpl) ImplementsTextView_WillCheckTextInRange_Options_Types() bool {
	return di._TextView_WillCheckTextInRange_Options_Types != nil
}

func (di *TextViewDelegateImpl) SetTextView_WillCheckTextInRange_Options_Types(f func(view TextView, range_ foundation.Range, options map[TextCheckingOptionKey]objc.Object, checkingTypes *foundation.TextCheckingTypes) map[TextCheckingOptionKey]objc.IObject) {
	di._TextView_WillCheckTextInRange_Options_Types = f
}

func (di *TextViewDelegateImpl) TextView_WillCheckTextInRange_Options_Types(view TextView, range_ foundation.Range, options map[TextCheckingOptionKey]objc.Object, checkingTypes *foundation.TextCheckingTypes) map[TextCheckingOptionKey]objc.IObject {
	return di._TextView_WillCheckTextInRange_Options_Types(view, range_, options, checkingTypes)
}
func (di *TextViewDelegateImpl) ImplementsTextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount() bool {
	return di._TextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount != nil
}

func (di *TextViewDelegateImpl) SetTextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount(f func(view TextView, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.Object, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.ITextCheckingResult) {
	di._TextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount = f
}

func (di *TextViewDelegateImpl) TextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount(view TextView, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.Object, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.ITextCheckingResult {
	return di._TextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount(view, range_, checkingTypes, options, results, orthography, wordCount)
}
func (di *TextViewDelegateImpl) ImplementsTextView_Completions_ForPartialWordRange_IndexOfSelectedItem() bool {
	return di._TextView_Completions_ForPartialWordRange_IndexOfSelectedItem != nil
}

func (di *TextViewDelegateImpl) SetTextView_Completions_ForPartialWordRange_IndexOfSelectedItem(f func(textView TextView, words []string, charRange foundation.Range, index *int) []string) {
	di._TextView_Completions_ForPartialWordRange_IndexOfSelectedItem = f
}

func (di *TextViewDelegateImpl) TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(textView TextView, words []string, charRange foundation.Range, index *int) []string {
	return di._TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(textView, words, charRange, index)
}
func (di *TextViewDelegateImpl) ImplementsTextView_WillShowSharingServicePicker_ForItems() bool {
	return di._TextView_WillShowSharingServicePicker_ForItems != nil
}

func (di *TextViewDelegateImpl) SetTextView_WillShowSharingServicePicker_ForItems(f func(textView TextView, servicePicker SharingServicePicker, items []objc.Object) ISharingServicePicker) {
	di._TextView_WillShowSharingServicePicker_ForItems = f
}

func (di *TextViewDelegateImpl) TextView_WillShowSharingServicePicker_ForItems(textView TextView, servicePicker SharingServicePicker, items []objc.Object) ISharingServicePicker {
	return di._TextView_WillShowSharingServicePicker_ForItems(textView, servicePicker, items)
}
func (di *TextViewDelegateImpl) ImplementsTextView_DoCommandBySelector() bool {
	return di._TextView_DoCommandBySelector != nil
}

func (di *TextViewDelegateImpl) SetTextView_DoCommandBySelector(f func(textView TextView, commandSelector objc.Selector) bool) {
	di._TextView_DoCommandBySelector = f
}

func (di *TextViewDelegateImpl) TextView_DoCommandBySelector(textView TextView, commandSelector objc.Selector) bool {
	return di._TextView_DoCommandBySelector(textView, commandSelector)
}
func (di *TextViewDelegateImpl) ImplementsTextView_Menu_ForEvent_AtIndex() bool {
	return di._TextView_Menu_ForEvent_AtIndex != nil
}

func (di *TextViewDelegateImpl) SetTextView_Menu_ForEvent_AtIndex(f func(view TextView, menu Menu, event Event, charIndex uint) IMenu) {
	di._TextView_Menu_ForEvent_AtIndex = f
}

func (di *TextViewDelegateImpl) TextView_Menu_ForEvent_AtIndex(view TextView, menu Menu, event Event, charIndex uint) IMenu {
	return di._TextView_Menu_ForEvent_AtIndex(view, menu, event, charIndex)
}
func (di *TextViewDelegateImpl) ImplementsTextView_ClickedOnLink() bool {
	return di._TextView_ClickedOnLink != nil
}

// deprecated
func (di *TextViewDelegateImpl) SetTextView_ClickedOnLink(f func(textView TextView, link objc.Object) bool) {
	di._TextView_ClickedOnLink = f
}

// deprecated
func (di *TextViewDelegateImpl) TextView_ClickedOnLink(textView TextView, link objc.Object) bool {
	return di._TextView_ClickedOnLink(textView, link)
}

type TextViewDelegateWrapper struct {
	TextDelegateWrapper
}

func (t_ *TextViewDelegateWrapper) ImplementsUndoManagerForTextView() bool {
	return t_.RespondsToSelector(objc.GetSelector("undoManagerForTextView:"))
}

func (t_ TextViewDelegateWrapper) UndoManagerForTextView(view ITextView) foundation.UndoManager {
	rv := objc.CallMethod[foundation.UndoManager](t_, "undoManagerForTextView:", view)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_WillDisplayToolTip_ForCharacterAtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:willDisplayToolTip:forCharacterAtIndex:"))
}

func (t_ TextViewDelegateWrapper) TextView_WillDisplayToolTip_ForCharacterAtIndex(textView ITextView, tooltip string, characterIndex uint) string {
	rv := objc.CallMethod[string](t_, "textView:willDisplayToolTip:forCharacterAtIndex:", textView, tooltip, characterIndex)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_URLForContentsOfTextAttachment_AtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:URLForContentsOfTextAttachment:atIndex:"))
}

func (t_ TextViewDelegateWrapper) TextView_URLForContentsOfTextAttachment_AtIndex(textView ITextView, textAttachment ITextAttachment, charIndex uint) foundation.URL {
	rv := objc.CallMethod[foundation.URL](t_, "textView:URLForContentsOfTextAttachment:atIndex:", textView, textAttachment, charIndex)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_WillChangeSelectionFromCharacterRange_ToCharacterRange() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:willChangeSelectionFromCharacterRange:toCharacterRange:"))
}

func (t_ TextViewDelegateWrapper) TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange(textView ITextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, "textView:willChangeSelectionFromCharacterRange:toCharacterRange:", textView, oldSelectedCharRange, newSelectedCharRange)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:willChangeSelectionFromCharacterRanges:toCharacterRanges:"))
}

func (t_ TextViewDelegateWrapper) TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges(textView ITextView, oldSelectedCharRanges []foundation.IValue, newSelectedCharRanges []foundation.IValue) []foundation.Value {
	rv := objc.CallMethod[[]foundation.Value](t_, "textView:willChangeSelectionFromCharacterRanges:toCharacterRanges:", textView, oldSelectedCharRanges, newSelectedCharRanges)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextViewDidChangeSelection() bool {
	return t_.RespondsToSelector(objc.GetSelector("textViewDidChangeSelection:"))
}

func (t_ TextViewDelegateWrapper) TextViewDidChangeSelection(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, "textViewDidChangeSelection:", notification)
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_Candidates_ForSelectedRange() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:candidates:forSelectedRange:"))
}

func (t_ TextViewDelegateWrapper) TextView_Candidates_ForSelectedRange(textView ITextView, candidates []foundation.ITextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult {
	rv := objc.CallMethod[[]foundation.TextCheckingResult](t_, "textView:candidates:forSelectedRange:", textView, candidates, selectedRange)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_CandidatesForSelectedRange() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:candidatesForSelectedRange:"))
}

func (t_ TextViewDelegateWrapper) TextView_CandidatesForSelectedRange(textView ITextView, selectedRange foundation.Range) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, "textView:candidatesForSelectedRange:", textView, selectedRange)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_ShouldSelectCandidateAtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:shouldSelectCandidateAtIndex:"))
}

func (t_ TextViewDelegateWrapper) TextView_ShouldSelectCandidateAtIndex(textView ITextView, index uint) bool {
	rv := objc.CallMethod[bool](t_, "textView:shouldSelectCandidateAtIndex:", textView, index)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_ShouldUpdateTouchBarItemIdentifiers() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:shouldUpdateTouchBarItemIdentifiers:"))
}

func (t_ TextViewDelegateWrapper) TextView_ShouldUpdateTouchBarItemIdentifiers(textView ITextView, identifiers []TouchBarItemIdentifier) []TouchBarItemIdentifier {
	rv := objc.CallMethod[[]TouchBarItemIdentifier](t_, "textView:shouldUpdateTouchBarItemIdentifiers:", textView, identifiers)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_ShouldChangeTextInRange_ReplacementString() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:shouldChangeTextInRange:replacementString:"))
}

func (t_ TextViewDelegateWrapper) TextView_ShouldChangeTextInRange_ReplacementString(textView ITextView, affectedCharRange foundation.Range, replacementString string) bool {
	rv := objc.CallMethod[bool](t_, "textView:shouldChangeTextInRange:replacementString:", textView, affectedCharRange, replacementString)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_ShouldChangeTextInRanges_ReplacementStrings() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:shouldChangeTextInRanges:replacementStrings:"))
}

func (t_ TextViewDelegateWrapper) TextView_ShouldChangeTextInRanges_ReplacementStrings(textView ITextView, affectedRanges []foundation.IValue, replacementStrings []string) bool {
	rv := objc.CallMethod[bool](t_, "textView:shouldChangeTextInRanges:replacementStrings:", textView, affectedRanges, replacementStrings)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_ShouldChangeTypingAttributes_ToAttributes() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:shouldChangeTypingAttributes:toAttributes:"))
}

func (t_ TextViewDelegateWrapper) TextView_ShouldChangeTypingAttributes_ToAttributes(textView ITextView, oldTypingAttributes map[string]objc.IObject, newTypingAttributes map[foundation.AttributedStringKey]objc.IObject) map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, "textView:shouldChangeTypingAttributes:toAttributes:", textView, oldTypingAttributes, newTypingAttributes)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextViewDidChangeTypingAttributes() bool {
	return t_.RespondsToSelector(objc.GetSelector("textViewDidChangeTypingAttributes:"))
}

func (t_ TextViewDelegateWrapper) TextViewDidChangeTypingAttributes(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, "textViewDidChangeTypingAttributes:", notification)
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_ClickedOnLink_AtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:clickedOnLink:atIndex:"))
}

func (t_ TextViewDelegateWrapper) TextView_ClickedOnLink_AtIndex(textView ITextView, link objc.IObject, charIndex uint) bool {
	rv := objc.CallMethod[bool](t_, "textView:clickedOnLink:atIndex:", textView, link, charIndex)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_ShouldSetSpellingState_Range() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:shouldSetSpellingState:range:"))
}

func (t_ TextViewDelegateWrapper) TextView_ShouldSetSpellingState_Range(textView ITextView, value int, affectedCharRange foundation.Range) int {
	rv := objc.CallMethod[int](t_, "textView:shouldSetSpellingState:range:", textView, value, affectedCharRange)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_WillCheckTextInRange_Options_Types() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:willCheckTextInRange:options:types:"))
}

func (t_ TextViewDelegateWrapper) TextView_WillCheckTextInRange_Options_Types(view ITextView, range_ foundation.Range, options map[TextCheckingOptionKey]objc.IObject, checkingTypes *foundation.TextCheckingTypes) map[TextCheckingOptionKey]objc.Object {
	rv := objc.CallMethod[map[TextCheckingOptionKey]objc.Object](t_, "textView:willCheckTextInRange:options:types:", view, range_, options, checkingTypes)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:didCheckTextInRange:types:options:results:orthography:wordCount:"))
}

func (t_ TextViewDelegateWrapper) TextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount(view ITextView, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject, results []foundation.ITextCheckingResult, orthography foundation.IOrthography, wordCount int) []foundation.TextCheckingResult {
	rv := objc.CallMethod[[]foundation.TextCheckingResult](t_, "textView:didCheckTextInRange:types:options:results:orthography:wordCount:", view, range_, checkingTypes, options, results, orthography, wordCount)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_Completions_ForPartialWordRange_IndexOfSelectedItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:completions:forPartialWordRange:indexOfSelectedItem:"))
}

func (t_ TextViewDelegateWrapper) TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(textView ITextView, words []string, charRange foundation.Range, index *int) []string {
	rv := objc.CallMethod[[]string](t_, "textView:completions:forPartialWordRange:indexOfSelectedItem:", textView, words, charRange, index)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_WillShowSharingServicePicker_ForItems() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:willShowSharingServicePicker:forItems:"))
}

func (t_ TextViewDelegateWrapper) TextView_WillShowSharingServicePicker_ForItems(textView ITextView, servicePicker ISharingServicePicker, items []objc.IObject) SharingServicePicker {
	rv := objc.CallMethod[SharingServicePicker](t_, "textView:willShowSharingServicePicker:forItems:", textView, servicePicker, items)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_DoCommandBySelector() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:doCommandBySelector:"))
}

func (t_ TextViewDelegateWrapper) TextView_DoCommandBySelector(textView ITextView, commandSelector objc.Selector) bool {
	rv := objc.CallMethod[bool](t_, "textView:doCommandBySelector:", textView, commandSelector)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_Menu_ForEvent_AtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:menu:forEvent:atIndex:"))
}

func (t_ TextViewDelegateWrapper) TextView_Menu_ForEvent_AtIndex(view ITextView, menu IMenu, event IEvent, charIndex uint) Menu {
	rv := objc.CallMethod[Menu](t_, "textView:menu:forEvent:atIndex:", view, menu, event, charIndex)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_ClickedOnLink() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:clickedOnLink:"))
}

// deprecated
func (t_ TextViewDelegateWrapper) TextView_ClickedOnLink(textView ITextView, link objc.IObject) bool {
	rv := objc.CallMethod[bool](t_, "textView:clickedOnLink:", textView, link)
	return rv
}
