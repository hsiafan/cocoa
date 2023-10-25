// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TextViewClass = _TextViewClass{objc.GetClass("NSTextView")}

type _TextViewClass struct {
	objc.Class
}

type ITextView interface {
	IText
	ReplaceTextContainer(newContainer ITextContainer)
	InvalidateTextContainerOrigin()
	ChangeDocumentBackgroundColor(sender objc.IObject)
	SetNeedsDisplayInRect_AvoidAdditionalLayout(rect foundation.Rect, flag bool)
	DrawInsertionPointInRect_Color_TurnedOn(rect foundation.Rect, color IColor, flag bool)
	DrawViewBackgroundInRect(rect foundation.Rect)
	SetConstrainedFrameSize(desiredSize foundation.Size)
	CleanUpAfterDragOperation()
	ShowFindIndicatorForRange(charRange foundation.Range)
	// deprecated
	InsertText(insertString objc.IObject)
	SetBaseWritingDirection_Range(writingDirection WritingDirection, range_ foundation.Range)
	// deprecated
	ToggleBaseWritingDirection(sender objc.IObject)
	Outline(sender objc.IObject)
	ToggleAutomaticQuoteSubstitution(sender objc.IObject)
	ToggleAutomaticLinkDetection(sender objc.IObject)
	ToggleAutomaticTextCompletion(sender objc.IObject)
	SetSelectedRange_Affinity_StillSelecting(charRange foundation.Range, affinity SelectionAffinity, stillSelectingFlag bool)
	SetSelectedRanges_Affinity_StillSelecting(ranges []foundation.IValue, affinity SelectionAffinity, stillSelectingFlag bool)
	UpdateInsertionPointStateAndRestartTimer(restartFlag bool)
	CharacterIndexForInsertionAtPoint(point foundation.Point) uint
	UpdateCandidates()
	PreferredPasteboardTypeFromArray_RestrictedToTypesFromArray(availableTypes []PasteboardType, allowedTypes []PasteboardType) PasteboardType
	ReadSelectionFromPasteboard(pboard IPasteboard) bool
	ReadSelectionFromPasteboard_Type(pboard IPasteboard, type_ PasteboardType) bool
	WriteSelectionToPasteboard_Type(pboard IPasteboard, type_ PasteboardType) bool
	WriteSelectionToPasteboard_Types(pboard IPasteboard, types []PasteboardType) bool
	AlignJustified(sender objc.IObject)
	ChangeAttributes(sender objc.IObject)
	ChangeColor(sender objc.IObject)
	SetAlignment_Range(alignment TextAlignment, range_ foundation.Range)
	UseStandardKerning(sender objc.IObject)
	LowerBaseline(sender objc.IObject)
	RaiseBaseline(sender objc.IObject)
	TurnOffKerning(sender objc.IObject)
	LoosenKerning(sender objc.IObject)
	TightenKerning(sender objc.IObject)
	UseStandardLigatures(sender objc.IObject)
	TurnOffLigatures(sender objc.IObject)
	UseAllLigatures(sender objc.IObject)
	// deprecated
	ToggleTraditionalCharacterShape(sender objc.IObject)
	ClickedOnLink_AtIndex(link objc.IObject, charIndex uint)
	PasteAsPlainText(sender objc.IObject)
	PasteAsRichText(sender objc.IObject)
	BreakUndoCoalescing()
	UpdateFontPanel()
	UpdateRuler()
	UpdateDragTypeRegistration()
	SelectionRangeForProposedRange_Granularity(proposedCharRange foundation.Range, granularity SelectionGranularity) foundation.Range
	ShouldChangeTextInRange_ReplacementString(affectedCharRange foundation.Range, replacementString string) bool
	ShouldChangeTextInRanges_ReplacementStrings(affectedRanges []foundation.IValue, replacementStrings []string) bool
	DidChangeText()
	SmartDeleteRangeForProposedRange(proposedCharRange foundation.Range) foundation.Range
	SmartInsertAfterStringForString_ReplacingRange(pasteString string, charRangeToReplace foundation.Range) string
	SmartInsertBeforeStringForString_ReplacingRange(pasteString string, charRangeToReplace foundation.Range) string
	SmartInsertForString_ReplacingRange_BeforeString_AfterString(pasteString string, charRangeToReplace foundation.Range, beforeString *foundation.String, afterString *foundation.String)
	ToggleSmartInsertDelete(sender objc.IObject)
	ToggleContinuousSpellChecking(sender objc.IObject)
	ToggleGrammarChecking(sender objc.IObject)
	SetSpellingState_Range(value int, charRange foundation.Range)
	OrderFrontSharingServicePicker(sender objc.IObject)
	DragImageForSelectionWithEvent_Origin(event IEvent, origin *foundation.Point) Image
	DragOperationForDraggingInfo_Type(dragInfo objc.IObject, type_ PasteboardType) DragOperation
	DragSelectionWithEvent_Offset_SlideBack(event IEvent, mouseOffset foundation.Size, slideBack bool) bool
	StartSpeaking(sender objc.IObject)
	StopSpeaking(sender objc.IObject)
	PerformFindPanelAction(sender objc.IObject)
	OrderFrontLinkPanel(sender objc.IObject)
	OrderFrontListPanel(sender objc.IObject)
	OrderFrontSpacingPanel(sender objc.IObject)
	OrderFrontTablePanel(sender objc.IObject)
	OrderFrontSubstitutionsPanel(sender objc.IObject)
	Complete(sender objc.IObject)
	CompletionsForPartialWordRange_IndexOfSelectedItem(charRange foundation.Range, index *int) []string
	InsertCompletion_ForPartialWordRange_Movement_IsFinal(word string, charRange foundation.Range, movement int, flag bool)
	CheckTextInDocument(sender objc.IObject)
	CheckTextInSelection(sender objc.IObject)
	CheckTextInRange_Types_Options(range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject)
	HandleTextCheckingResults_ForRange_Types_Options_Orthography_WordCount(results []foundation.ITextCheckingResult, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject, orthography foundation.IOrthography, wordCount int)
	ToggleAutomaticDashSubstitution(sender objc.IObject)
	ToggleAutomaticDataDetection(sender objc.IObject)
	ToggleAutomaticSpellingCorrection(sender objc.IObject)
	ToggleAutomaticTextReplacement(sender objc.IObject)
	PerformValidatedReplacementInRange_WithAttributedString(range_ foundation.Range, attributedString foundation.IAttributedString) bool
	UpdateQuickLookPreviewPanel()
	ToggleQuickLookPreviewPanel(sender objc.IObject)
	ChangeLayoutOrientation(sender objc.IObject)
	SetLayoutOrientation(orientation TextLayoutOrientation)
	UpdateTextTouchBarItems()
	UpdateTouchBarItemIdentifiers()
	TextContainer() TextContainer
	SetTextContainer(value ITextContainer)
	TextContainerInset() foundation.Size
	SetTextContainerInset(value foundation.Size)
	TextContainerOrigin() foundation.Point
	TextLayoutManager() TextLayoutManager
	LayoutManager() LayoutManager
	TextContentStorage() TextContentStorage
	TextStorage() TextStorage
	AllowsDocumentBackgroundColorChange() bool
	SetAllowsDocumentBackgroundColorChange(value bool)
	ShouldDrawInsertionPoint() bool
	AllowedInputSourceLocales() []string
	SetAllowedInputSourceLocales(value []string)
	AllowsUndo() bool
	SetAllowsUndo(value bool)
	DefaultParagraphStyle() ParagraphStyle
	SetDefaultParagraphStyle(value IParagraphStyle)
	AllowsImageEditing() bool
	SetAllowsImageEditing(value bool)
	IsAutomaticQuoteSubstitutionEnabled() bool
	SetAutomaticQuoteSubstitutionEnabled(value bool)
	IsAutomaticLinkDetectionEnabled() bool
	SetAutomaticLinkDetectionEnabled(value bool)
	DisplaysLinkToolTips() bool
	SetDisplaysLinkToolTips(value bool)
	IsAutomaticTextCompletionEnabled() bool
	SetAutomaticTextCompletionEnabled(value bool)
	UsesAdaptiveColorMappingForDarkAppearance() bool
	SetUsesAdaptiveColorMappingForDarkAppearance(value bool)
	UsesRolloverButtonForSelection() bool
	SetUsesRolloverButtonForSelection(value bool)
	UsesRuler() bool
	SetUsesRuler(value bool)
	SetRulerVisible(value bool)
	UsesInspectorBar() bool
	SetUsesInspectorBar(value bool)
	SelectedRanges() []foundation.Value
	SetSelectedRanges(value []foundation.IValue)
	SelectionAffinity() SelectionAffinity
	SelectionGranularity() SelectionGranularity
	SetSelectionGranularity(value SelectionGranularity)
	InsertionPointColor() Color
	SetInsertionPointColor(value IColor)
	SelectedTextAttributes() map[foundation.AttributedStringKey]objc.Object
	SetSelectedTextAttributes(value map[foundation.AttributedStringKey]objc.IObject)
	MarkedTextAttributes() map[foundation.AttributedStringKey]objc.Object
	SetMarkedTextAttributes(value map[foundation.AttributedStringKey]objc.IObject)
	LinkTextAttributes() map[foundation.AttributedStringKey]objc.Object
	SetLinkTextAttributes(value map[foundation.AttributedStringKey]objc.IObject)
	ReadablePasteboardTypes() []PasteboardType
	WritablePasteboardTypes() []PasteboardType
	TypingAttributes() map[foundation.AttributedStringKey]objc.Object
	SetTypingAttributes(value map[foundation.AttributedStringKey]objc.IObject)
	IsCoalescingUndo() bool
	AcceptableDragTypes() []PasteboardType
	RangeForUserCharacterAttributeChange() foundation.Range
	RangesForUserCharacterAttributeChange() []foundation.Value
	RangeForUserParagraphAttributeChange() foundation.Range
	RangesForUserParagraphAttributeChange() []foundation.Value
	RangeForUserTextChange() foundation.Range
	RangesForUserTextChange() []foundation.Value
	SmartInsertDeleteEnabled() bool
	SetSmartInsertDeleteEnabled(value bool)
	IsContinuousSpellCheckingEnabled() bool
	SetContinuousSpellCheckingEnabled(value bool)
	SpellCheckerDocumentTag() int
	IsGrammarCheckingEnabled() bool
	SetGrammarCheckingEnabled(value bool)
	AcceptsGlyphInfo() bool
	SetAcceptsGlyphInfo(value bool)
	UsesFindPanel() bool
	SetUsesFindPanel(value bool)
	RangeForUserCompletion() foundation.Range
	EnabledTextCheckingTypes() foundation.TextCheckingTypes
	SetEnabledTextCheckingTypes(value foundation.TextCheckingTypes)
	IsAutomaticDashSubstitutionEnabled() bool
	SetAutomaticDashSubstitutionEnabled(value bool)
	IsAutomaticDataDetectionEnabled() bool
	SetAutomaticDataDetectionEnabled(value bool)
	IsAutomaticSpellingCorrectionEnabled() bool
	SetAutomaticSpellingCorrectionEnabled(value bool)
	IsAutomaticTextReplacementEnabled() bool
	SetAutomaticTextReplacementEnabled(value bool)
	UsesFindBar() bool
	SetUsesFindBar(value bool)
	IsIncrementalSearchingEnabled() bool
	SetIncrementalSearchingEnabled(value bool)
	AllowsCharacterPickerTouchBarItem() bool
	SetAllowsCharacterPickerTouchBarItem(value bool)
}

type TextView struct {
	Text
}

func MakeTextView(ptr unsafe.Pointer) TextView {
	return TextView{
		Text: MakeText(ptr),
	}
}

func (t_ TextView) InitWithFrame_TextContainer(frameRect foundation.Rect, container ITextContainer) TextView {
	rv := objc.CallMethod[TextView](t_, objc.SEL("initWithFrame:textContainer:"), frameRect, objc.ExtractPtr(container))
	return rv
}

func (t_ TextView) InitWithFrame(frameRect foundation.Rect) TextView {
	rv := objc.CallMethod[TextView](t_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (tc _TextViewClass) FieldEditor() TextView {
	rv := objc.CallMethod[TextView](tc, objc.SEL("fieldEditor"))
	return rv
}

func (t_ TextView) Init() TextView {
	rv := objc.CallMethod[TextView](t_, objc.SEL("init"))
	return rv
}

func (tc _TextViewClass) Alloc() TextView {
	rv := objc.CallMethod[TextView](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TextViewClass) New() TextView {
	rv := objc.CallMethod[TextView](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTextView() TextView {
	return TextViewClass.New()
}

func (tc _TextViewClass) RegisterForServices() {
	objc.CallMethod[objc.Void](tc, objc.SEL("registerForServices"))
}

func (t_ TextView) ReplaceTextContainer(newContainer ITextContainer) {
	objc.CallMethod[objc.Void](t_, objc.SEL("replaceTextContainer:"), objc.ExtractPtr(newContainer))
}

func (t_ TextView) InvalidateTextContainerOrigin() {
	objc.CallMethod[objc.Void](t_, objc.SEL("invalidateTextContainerOrigin"))
}

func (t_ TextView) ChangeDocumentBackgroundColor(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("changeDocumentBackgroundColor:"), objc.ExtractPtr(sender))
}

func (t_ TextView) SetNeedsDisplayInRect_AvoidAdditionalLayout(rect foundation.Rect, flag bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setNeedsDisplayInRect:avoidAdditionalLayout:"), rect, flag)
}

func (t_ TextView) DrawInsertionPointInRect_Color_TurnedOn(rect foundation.Rect, color IColor, flag bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("drawInsertionPointInRect:color:turnedOn:"), rect, objc.ExtractPtr(color), flag)
}

func (t_ TextView) DrawViewBackgroundInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.SEL("drawViewBackgroundInRect:"), rect)
}

func (t_ TextView) SetConstrainedFrameSize(desiredSize foundation.Size) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setConstrainedFrameSize:"), desiredSize)
}

func (t_ TextView) CleanUpAfterDragOperation() {
	objc.CallMethod[objc.Void](t_, objc.SEL("cleanUpAfterDragOperation"))
}

func (t_ TextView) ShowFindIndicatorForRange(charRange foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.SEL("showFindIndicatorForRange:"), charRange)
}

func (tc _TextViewClass) ScrollableDocumentContentTextView() ScrollView {
	rv := objc.CallMethod[ScrollView](tc, objc.SEL("scrollableDocumentContentTextView"))
	return rv
}

func (tc _TextViewClass) ScrollablePlainDocumentContentTextView() ScrollView {
	rv := objc.CallMethod[ScrollView](tc, objc.SEL("scrollablePlainDocumentContentTextView"))
	return rv
}

func (tc _TextViewClass) ScrollableTextView() ScrollView {
	rv := objc.CallMethod[ScrollView](tc, objc.SEL("scrollableTextView"))
	return rv
}

// deprecated
func (t_ TextView) InsertText(insertString objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("insertText:"), objc.ExtractPtr(insertString))
}

func (t_ TextView) SetBaseWritingDirection_Range(writingDirection WritingDirection, range_ foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setBaseWritingDirection:range:"), writingDirection, range_)
}

// deprecated
func (t_ TextView) ToggleBaseWritingDirection(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("toggleBaseWritingDirection:"), objc.ExtractPtr(sender))
}

func (t_ TextView) Outline(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("outline:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ToggleAutomaticQuoteSubstitution(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("toggleAutomaticQuoteSubstitution:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ToggleAutomaticLinkDetection(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("toggleAutomaticLinkDetection:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ToggleAutomaticTextCompletion(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("toggleAutomaticTextCompletion:"), objc.ExtractPtr(sender))
}

func (t_ TextView) SetSelectedRange_Affinity_StillSelecting(charRange foundation.Range, affinity SelectionAffinity, stillSelectingFlag bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSelectedRange:affinity:stillSelecting:"), charRange, affinity, stillSelectingFlag)
}

func (t_ TextView) SetSelectedRanges_Affinity_StillSelecting(ranges []foundation.IValue, affinity SelectionAffinity, stillSelectingFlag bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSelectedRanges:affinity:stillSelecting:"), ranges, affinity, stillSelectingFlag)
}

func (t_ TextView) UpdateInsertionPointStateAndRestartTimer(restartFlag bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("updateInsertionPointStateAndRestartTimer:"), restartFlag)
}

func (t_ TextView) CharacterIndexForInsertionAtPoint(point foundation.Point) uint {
	rv := objc.CallMethod[uint](t_, objc.SEL("characterIndexForInsertionAtPoint:"), point)
	return rv
}

func (t_ TextView) UpdateCandidates() {
	objc.CallMethod[objc.Void](t_, objc.SEL("updateCandidates"))
}

func (t_ TextView) PreferredPasteboardTypeFromArray_RestrictedToTypesFromArray(availableTypes []PasteboardType, allowedTypes []PasteboardType) PasteboardType {
	rv := objc.CallMethod[PasteboardType](t_, objc.SEL("preferredPasteboardTypeFromArray:restrictedToTypesFromArray:"), availableTypes, allowedTypes)
	return rv
}

func (t_ TextView) ReadSelectionFromPasteboard(pboard IPasteboard) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("readSelectionFromPasteboard:"), objc.ExtractPtr(pboard))
	return rv
}

func (t_ TextView) ReadSelectionFromPasteboard_Type(pboard IPasteboard, type_ PasteboardType) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("readSelectionFromPasteboard:type:"), objc.ExtractPtr(pboard), type_)
	return rv
}

func (t_ TextView) WriteSelectionToPasteboard_Type(pboard IPasteboard, type_ PasteboardType) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("writeSelectionToPasteboard:type:"), objc.ExtractPtr(pboard), type_)
	return rv
}

func (t_ TextView) WriteSelectionToPasteboard_Types(pboard IPasteboard, types []PasteboardType) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("writeSelectionToPasteboard:types:"), objc.ExtractPtr(pboard), types)
	return rv
}

func (t_ TextView) AlignJustified(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("alignJustified:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ChangeAttributes(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("changeAttributes:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ChangeColor(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("changeColor:"), objc.ExtractPtr(sender))
}

func (t_ TextView) SetAlignment_Range(alignment TextAlignment, range_ foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAlignment:range:"), alignment, range_)
}

func (t_ TextView) UseStandardKerning(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("useStandardKerning:"), objc.ExtractPtr(sender))
}

func (t_ TextView) LowerBaseline(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("lowerBaseline:"), objc.ExtractPtr(sender))
}

func (t_ TextView) RaiseBaseline(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("raiseBaseline:"), objc.ExtractPtr(sender))
}

func (t_ TextView) TurnOffKerning(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("turnOffKerning:"), objc.ExtractPtr(sender))
}

func (t_ TextView) LoosenKerning(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("loosenKerning:"), objc.ExtractPtr(sender))
}

func (t_ TextView) TightenKerning(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("tightenKerning:"), objc.ExtractPtr(sender))
}

func (t_ TextView) UseStandardLigatures(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("useStandardLigatures:"), objc.ExtractPtr(sender))
}

func (t_ TextView) TurnOffLigatures(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("turnOffLigatures:"), objc.ExtractPtr(sender))
}

func (t_ TextView) UseAllLigatures(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("useAllLigatures:"), objc.ExtractPtr(sender))
}

// deprecated
func (t_ TextView) ToggleTraditionalCharacterShape(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("toggleTraditionalCharacterShape:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ClickedOnLink_AtIndex(link objc.IObject, charIndex uint) {
	objc.CallMethod[objc.Void](t_, objc.SEL("clickedOnLink:atIndex:"), objc.ExtractPtr(link), charIndex)
}

func (t_ TextView) PasteAsPlainText(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("pasteAsPlainText:"), objc.ExtractPtr(sender))
}

func (t_ TextView) PasteAsRichText(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("pasteAsRichText:"), objc.ExtractPtr(sender))
}

func (t_ TextView) BreakUndoCoalescing() {
	objc.CallMethod[objc.Void](t_, objc.SEL("breakUndoCoalescing"))
}

func (t_ TextView) UpdateFontPanel() {
	objc.CallMethod[objc.Void](t_, objc.SEL("updateFontPanel"))
}

func (t_ TextView) UpdateRuler() {
	objc.CallMethod[objc.Void](t_, objc.SEL("updateRuler"))
}

func (t_ TextView) UpdateDragTypeRegistration() {
	objc.CallMethod[objc.Void](t_, objc.SEL("updateDragTypeRegistration"))
}

func (t_ TextView) SelectionRangeForProposedRange_Granularity(proposedCharRange foundation.Range, granularity SelectionGranularity) foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.SEL("selectionRangeForProposedRange:granularity:"), proposedCharRange, granularity)
	return rv
}

func (t_ TextView) ShouldChangeTextInRange_ReplacementString(affectedCharRange foundation.Range, replacementString string) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("shouldChangeTextInRange:replacementString:"), affectedCharRange, replacementString)
	return rv
}

func (t_ TextView) ShouldChangeTextInRanges_ReplacementStrings(affectedRanges []foundation.IValue, replacementStrings []string) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("shouldChangeTextInRanges:replacementStrings:"), affectedRanges, replacementStrings)
	return rv
}

func (t_ TextView) DidChangeText() {
	objc.CallMethod[objc.Void](t_, objc.SEL("didChangeText"))
}

func (t_ TextView) SmartDeleteRangeForProposedRange(proposedCharRange foundation.Range) foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.SEL("smartDeleteRangeForProposedRange:"), proposedCharRange)
	return rv
}

func (t_ TextView) SmartInsertAfterStringForString_ReplacingRange(pasteString string, charRangeToReplace foundation.Range) string {
	rv := objc.CallMethod[string](t_, objc.SEL("smartInsertAfterStringForString:replacingRange:"), pasteString, charRangeToReplace)
	return rv
}

func (t_ TextView) SmartInsertBeforeStringForString_ReplacingRange(pasteString string, charRangeToReplace foundation.Range) string {
	rv := objc.CallMethod[string](t_, objc.SEL("smartInsertBeforeStringForString:replacingRange:"), pasteString, charRangeToReplace)
	return rv
}

func (t_ TextView) SmartInsertForString_ReplacingRange_BeforeString_AfterString(pasteString string, charRangeToReplace foundation.Range, beforeString *foundation.String, afterString *foundation.String) {
	objc.CallMethod[objc.Void](t_, objc.SEL("smartInsertForString:replacingRange:beforeString:afterString:"), pasteString, charRangeToReplace, beforeString, afterString)
}

func (t_ TextView) ToggleSmartInsertDelete(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("toggleSmartInsertDelete:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ToggleContinuousSpellChecking(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("toggleContinuousSpellChecking:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ToggleGrammarChecking(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("toggleGrammarChecking:"), objc.ExtractPtr(sender))
}

func (t_ TextView) SetSpellingState_Range(value int, charRange foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSpellingState:range:"), value, charRange)
}

func (t_ TextView) OrderFrontSharingServicePicker(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("orderFrontSharingServicePicker:"), objc.ExtractPtr(sender))
}

func (t_ TextView) DragImageForSelectionWithEvent_Origin(event IEvent, origin *foundation.Point) Image {
	rv := objc.CallMethod[Image](t_, objc.SEL("dragImageForSelectionWithEvent:origin:"), objc.ExtractPtr(event), origin)
	return rv
}

func (t_ TextView) DragOperationForDraggingInfo_Type(dragInfo objc.IObject, type_ PasteboardType) DragOperation {
	rv := objc.CallMethod[DragOperation](t_, objc.SEL("dragOperationForDraggingInfo:type:"), objc.ExtractPtr(dragInfo), type_)
	return rv
}

func (t_ TextView) DragSelectionWithEvent_Offset_SlideBack(event IEvent, mouseOffset foundation.Size, slideBack bool) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("dragSelectionWithEvent:offset:slideBack:"), objc.ExtractPtr(event), mouseOffset, slideBack)
	return rv
}

func (t_ TextView) StartSpeaking(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("startSpeaking:"), objc.ExtractPtr(sender))
}

func (t_ TextView) StopSpeaking(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("stopSpeaking:"), objc.ExtractPtr(sender))
}

func (t_ TextView) PerformFindPanelAction(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("performFindPanelAction:"), objc.ExtractPtr(sender))
}

func (t_ TextView) OrderFrontLinkPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("orderFrontLinkPanel:"), objc.ExtractPtr(sender))
}

func (t_ TextView) OrderFrontListPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("orderFrontListPanel:"), objc.ExtractPtr(sender))
}

func (t_ TextView) OrderFrontSpacingPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("orderFrontSpacingPanel:"), objc.ExtractPtr(sender))
}

func (t_ TextView) OrderFrontTablePanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("orderFrontTablePanel:"), objc.ExtractPtr(sender))
}

func (t_ TextView) OrderFrontSubstitutionsPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("orderFrontSubstitutionsPanel:"), objc.ExtractPtr(sender))
}

func (t_ TextView) Complete(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("complete:"), objc.ExtractPtr(sender))
}

func (t_ TextView) CompletionsForPartialWordRange_IndexOfSelectedItem(charRange foundation.Range, index *int) []string {
	rv := objc.CallMethod[[]string](t_, objc.SEL("completionsForPartialWordRange:indexOfSelectedItem:"), charRange, index)
	return rv
}

func (t_ TextView) InsertCompletion_ForPartialWordRange_Movement_IsFinal(word string, charRange foundation.Range, movement int, flag bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("insertCompletion:forPartialWordRange:movement:isFinal:"), word, charRange, movement, flag)
}

func (t_ TextView) CheckTextInDocument(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("checkTextInDocument:"), objc.ExtractPtr(sender))
}

func (t_ TextView) CheckTextInSelection(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("checkTextInSelection:"), objc.ExtractPtr(sender))
}

func (t_ TextView) CheckTextInRange_Types_Options(range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("checkTextInRange:types:options:"), range_, checkingTypes, options)
}

func (t_ TextView) HandleTextCheckingResults_ForRange_Types_Options_Orthography_WordCount(results []foundation.ITextCheckingResult, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject, orthography foundation.IOrthography, wordCount int) {
	objc.CallMethod[objc.Void](t_, objc.SEL("handleTextCheckingResults:forRange:types:options:orthography:wordCount:"), results, range_, checkingTypes, options, objc.ExtractPtr(orthography), wordCount)
}

func (t_ TextView) ToggleAutomaticDashSubstitution(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("toggleAutomaticDashSubstitution:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ToggleAutomaticDataDetection(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("toggleAutomaticDataDetection:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ToggleAutomaticSpellingCorrection(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("toggleAutomaticSpellingCorrection:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ToggleAutomaticTextReplacement(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("toggleAutomaticTextReplacement:"), objc.ExtractPtr(sender))
}

func (t_ TextView) PerformValidatedReplacementInRange_WithAttributedString(range_ foundation.Range, attributedString foundation.IAttributedString) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("performValidatedReplacementInRange:withAttributedString:"), range_, objc.ExtractPtr(attributedString))
	return rv
}

func (t_ TextView) UpdateQuickLookPreviewPanel() {
	objc.CallMethod[objc.Void](t_, objc.SEL("updateQuickLookPreviewPanel"))
}

func (t_ TextView) ToggleQuickLookPreviewPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("toggleQuickLookPreviewPanel:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ChangeLayoutOrientation(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("changeLayoutOrientation:"), objc.ExtractPtr(sender))
}

func (t_ TextView) SetLayoutOrientation(orientation TextLayoutOrientation) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setLayoutOrientation:"), orientation)
}

func (t_ TextView) UpdateTextTouchBarItems() {
	objc.CallMethod[objc.Void](t_, objc.SEL("updateTextTouchBarItems"))
}

func (t_ TextView) UpdateTouchBarItemIdentifiers() {
	objc.CallMethod[objc.Void](t_, objc.SEL("updateTouchBarItemIdentifiers"))
}

func (tc _TextViewClass) StronglyReferencesTextStorage() bool {
	rv := objc.CallMethod[bool](tc, objc.SEL("stronglyReferencesTextStorage"))
	return rv
}

// weak property
func (t_ TextView) TextContainer() TextContainer {
	rv := objc.CallMethod[TextContainer](t_, objc.SEL("textContainer"))
	return rv
}

// weak property
func (t_ TextView) SetTextContainer(value ITextContainer) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTextContainer:"), objc.ExtractPtr(value))
}

func (t_ TextView) TextContainerInset() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.SEL("textContainerInset"))
	return rv
}

func (t_ TextView) SetTextContainerInset(value foundation.Size) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTextContainerInset:"), value)
}

func (t_ TextView) TextContainerOrigin() foundation.Point {
	rv := objc.CallMethod[foundation.Point](t_, objc.SEL("textContainerOrigin"))
	return rv
}

// weak property
func (t_ TextView) TextLayoutManager() TextLayoutManager {
	rv := objc.CallMethod[TextLayoutManager](t_, objc.SEL("textLayoutManager"))
	return rv
}

// weak property
func (t_ TextView) LayoutManager() LayoutManager {
	rv := objc.CallMethod[LayoutManager](t_, objc.SEL("layoutManager"))
	return rv
}

// weak property
func (t_ TextView) TextContentStorage() TextContentStorage {
	rv := objc.CallMethod[TextContentStorage](t_, objc.SEL("textContentStorage"))
	return rv
}

// weak property
func (t_ TextView) TextStorage() TextStorage {
	rv := objc.CallMethod[TextStorage](t_, objc.SEL("textStorage"))
	return rv
}

func (t_ TextView) AllowsDocumentBackgroundColorChange() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("allowsDocumentBackgroundColorChange"))
	return rv
}

func (t_ TextView) SetAllowsDocumentBackgroundColorChange(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAllowsDocumentBackgroundColorChange:"), value)
}

func (t_ TextView) ShouldDrawInsertionPoint() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("shouldDrawInsertionPoint"))
	return rv
}

func (t_ TextView) AllowedInputSourceLocales() []string {
	rv := objc.CallMethod[[]string](t_, objc.SEL("allowedInputSourceLocales"))
	return rv
}

func (t_ TextView) SetAllowedInputSourceLocales(value []string) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAllowedInputSourceLocales:"), value)
}

func (t_ TextView) AllowsUndo() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("allowsUndo"))
	return rv
}

func (t_ TextView) SetAllowsUndo(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAllowsUndo:"), value)
}

func (t_ TextView) DefaultParagraphStyle() ParagraphStyle {
	rv := objc.CallMethod[ParagraphStyle](t_, objc.SEL("defaultParagraphStyle"))
	return rv
}

func (t_ TextView) SetDefaultParagraphStyle(value IParagraphStyle) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDefaultParagraphStyle:"), objc.ExtractPtr(value))
}

func (t_ TextView) AllowsImageEditing() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("allowsImageEditing"))
	return rv
}

func (t_ TextView) SetAllowsImageEditing(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAllowsImageEditing:"), value)
}

func (t_ TextView) IsAutomaticQuoteSubstitutionEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isAutomaticQuoteSubstitutionEnabled"))
	return rv
}

func (t_ TextView) SetAutomaticQuoteSubstitutionEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAutomaticQuoteSubstitutionEnabled:"), value)
}

func (t_ TextView) IsAutomaticLinkDetectionEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isAutomaticLinkDetectionEnabled"))
	return rv
}

func (t_ TextView) SetAutomaticLinkDetectionEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAutomaticLinkDetectionEnabled:"), value)
}

func (t_ TextView) DisplaysLinkToolTips() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("displaysLinkToolTips"))
	return rv
}

func (t_ TextView) SetDisplaysLinkToolTips(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDisplaysLinkToolTips:"), value)
}

func (t_ TextView) IsAutomaticTextCompletionEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isAutomaticTextCompletionEnabled"))
	return rv
}

func (t_ TextView) SetAutomaticTextCompletionEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAutomaticTextCompletionEnabled:"), value)
}

func (t_ TextView) UsesAdaptiveColorMappingForDarkAppearance() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("usesAdaptiveColorMappingForDarkAppearance"))
	return rv
}

func (t_ TextView) SetUsesAdaptiveColorMappingForDarkAppearance(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setUsesAdaptiveColorMappingForDarkAppearance:"), value)
}

func (t_ TextView) UsesRolloverButtonForSelection() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("usesRolloverButtonForSelection"))
	return rv
}

func (t_ TextView) SetUsesRolloverButtonForSelection(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setUsesRolloverButtonForSelection:"), value)
}

func (t_ TextView) UsesRuler() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("usesRuler"))
	return rv
}

func (t_ TextView) SetUsesRuler(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setUsesRuler:"), value)
}

func (t_ TextView) SetRulerVisible(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setRulerVisible:"), value)
}

func (t_ TextView) UsesInspectorBar() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("usesInspectorBar"))
	return rv
}

func (t_ TextView) SetUsesInspectorBar(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setUsesInspectorBar:"), value)
}

func (t_ TextView) SelectedRanges() []foundation.Value {
	rv := objc.CallMethod[[]foundation.Value](t_, objc.SEL("selectedRanges"))
	return rv
}

func (t_ TextView) SetSelectedRanges(value []foundation.IValue) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSelectedRanges:"), value)
}

func (t_ TextView) SelectionAffinity() SelectionAffinity {
	rv := objc.CallMethod[SelectionAffinity](t_, objc.SEL("selectionAffinity"))
	return rv
}

func (t_ TextView) SelectionGranularity() SelectionGranularity {
	rv := objc.CallMethod[SelectionGranularity](t_, objc.SEL("selectionGranularity"))
	return rv
}

func (t_ TextView) SetSelectionGranularity(value SelectionGranularity) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSelectionGranularity:"), value)
}

func (t_ TextView) InsertionPointColor() Color {
	rv := objc.CallMethod[Color](t_, objc.SEL("insertionPointColor"))
	return rv
}

func (t_ TextView) SetInsertionPointColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setInsertionPointColor:"), objc.ExtractPtr(value))
}

func (t_ TextView) SelectedTextAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, objc.SEL("selectedTextAttributes"))
	return rv
}

func (t_ TextView) SetSelectedTextAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSelectedTextAttributes:"), value)
}

func (t_ TextView) MarkedTextAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, objc.SEL("markedTextAttributes"))
	return rv
}

func (t_ TextView) SetMarkedTextAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setMarkedTextAttributes:"), value)
}

func (t_ TextView) LinkTextAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, objc.SEL("linkTextAttributes"))
	return rv
}

func (t_ TextView) SetLinkTextAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setLinkTextAttributes:"), value)
}

func (t_ TextView) ReadablePasteboardTypes() []PasteboardType {
	rv := objc.CallMethod[[]PasteboardType](t_, objc.SEL("readablePasteboardTypes"))
	return rv
}

func (t_ TextView) WritablePasteboardTypes() []PasteboardType {
	rv := objc.CallMethod[[]PasteboardType](t_, objc.SEL("writablePasteboardTypes"))
	return rv
}

func (t_ TextView) TypingAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, objc.SEL("typingAttributes"))
	return rv
}

func (t_ TextView) SetTypingAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTypingAttributes:"), value)
}

func (t_ TextView) IsCoalescingUndo() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isCoalescingUndo"))
	return rv
}

func (t_ TextView) AcceptableDragTypes() []PasteboardType {
	rv := objc.CallMethod[[]PasteboardType](t_, objc.SEL("acceptableDragTypes"))
	return rv
}

func (t_ TextView) RangeForUserCharacterAttributeChange() foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.SEL("rangeForUserCharacterAttributeChange"))
	return rv
}

func (t_ TextView) RangesForUserCharacterAttributeChange() []foundation.Value {
	rv := objc.CallMethod[[]foundation.Value](t_, objc.SEL("rangesForUserCharacterAttributeChange"))
	return rv
}

func (t_ TextView) RangeForUserParagraphAttributeChange() foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.SEL("rangeForUserParagraphAttributeChange"))
	return rv
}

func (t_ TextView) RangesForUserParagraphAttributeChange() []foundation.Value {
	rv := objc.CallMethod[[]foundation.Value](t_, objc.SEL("rangesForUserParagraphAttributeChange"))
	return rv
}

func (t_ TextView) RangeForUserTextChange() foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.SEL("rangeForUserTextChange"))
	return rv
}

func (t_ TextView) RangesForUserTextChange() []foundation.Value {
	rv := objc.CallMethod[[]foundation.Value](t_, objc.SEL("rangesForUserTextChange"))
	return rv
}

func (t_ TextView) SmartInsertDeleteEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("smartInsertDeleteEnabled"))
	return rv
}

func (t_ TextView) SetSmartInsertDeleteEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSmartInsertDeleteEnabled:"), value)
}

func (t_ TextView) IsContinuousSpellCheckingEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isContinuousSpellCheckingEnabled"))
	return rv
}

func (t_ TextView) SetContinuousSpellCheckingEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setContinuousSpellCheckingEnabled:"), value)
}

func (t_ TextView) SpellCheckerDocumentTag() int {
	rv := objc.CallMethod[int](t_, objc.SEL("spellCheckerDocumentTag"))
	return rv
}

func (t_ TextView) IsGrammarCheckingEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isGrammarCheckingEnabled"))
	return rv
}

func (t_ TextView) SetGrammarCheckingEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setGrammarCheckingEnabled:"), value)
}

func (t_ TextView) AcceptsGlyphInfo() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("acceptsGlyphInfo"))
	return rv
}

func (t_ TextView) SetAcceptsGlyphInfo(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAcceptsGlyphInfo:"), value)
}

func (t_ TextView) UsesFindPanel() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("usesFindPanel"))
	return rv
}

func (t_ TextView) SetUsesFindPanel(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setUsesFindPanel:"), value)
}

func (t_ TextView) RangeForUserCompletion() foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.SEL("rangeForUserCompletion"))
	return rv
}

func (t_ TextView) EnabledTextCheckingTypes() foundation.TextCheckingTypes {
	rv := objc.CallMethod[foundation.TextCheckingTypes](t_, objc.SEL("enabledTextCheckingTypes"))
	return rv
}

func (t_ TextView) SetEnabledTextCheckingTypes(value foundation.TextCheckingTypes) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setEnabledTextCheckingTypes:"), value)
}

func (t_ TextView) IsAutomaticDashSubstitutionEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isAutomaticDashSubstitutionEnabled"))
	return rv
}

func (t_ TextView) SetAutomaticDashSubstitutionEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAutomaticDashSubstitutionEnabled:"), value)
}

func (t_ TextView) IsAutomaticDataDetectionEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isAutomaticDataDetectionEnabled"))
	return rv
}

func (t_ TextView) SetAutomaticDataDetectionEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAutomaticDataDetectionEnabled:"), value)
}

func (t_ TextView) IsAutomaticSpellingCorrectionEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isAutomaticSpellingCorrectionEnabled"))
	return rv
}

func (t_ TextView) SetAutomaticSpellingCorrectionEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAutomaticSpellingCorrectionEnabled:"), value)
}

func (t_ TextView) IsAutomaticTextReplacementEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isAutomaticTextReplacementEnabled"))
	return rv
}

func (t_ TextView) SetAutomaticTextReplacementEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAutomaticTextReplacementEnabled:"), value)
}

func (t_ TextView) UsesFindBar() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("usesFindBar"))
	return rv
}

func (t_ TextView) SetUsesFindBar(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setUsesFindBar:"), value)
}

func (t_ TextView) IsIncrementalSearchingEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isIncrementalSearchingEnabled"))
	return rv
}

func (t_ TextView) SetIncrementalSearchingEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setIncrementalSearchingEnabled:"), value)
}

func (t_ TextView) AllowsCharacterPickerTouchBarItem() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("allowsCharacterPickerTouchBarItem"))
	return rv
}

func (t_ TextView) SetAllowsCharacterPickerTouchBarItem(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAllowsCharacterPickerTouchBarItem:"), value)
}
