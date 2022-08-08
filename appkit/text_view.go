// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
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
	SetBaseWritingDirection_Range(writingDirection WritingDirection, _range foundation.Range)
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
	ReadSelectionFromPasteboard_Type(pboard IPasteboard, _type PasteboardType) bool
	WriteSelectionToPasteboard_Type(pboard IPasteboard, _type PasteboardType) bool
	WriteSelectionToPasteboard_Types(pboard IPasteboard, types []PasteboardType) bool
	AlignJustified(sender objc.IObject)
	ChangeAttributes(sender objc.IObject)
	ChangeColor(sender objc.IObject)
	SetAlignment_Range(alignment TextAlignment, _range foundation.Range)
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
	DragOperationForDraggingInfo_Type(dragInfo DraggingInfo, _type PasteboardType) DragOperation
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
	CheckTextInRange_Types_Options(_range foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject)
	HandleTextCheckingResults_ForRange_Types_Options_Orthography_WordCount(results []foundation.ITextCheckingResult, _range foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject, orthography foundation.IOrthography, wordCount int)
	ToggleAutomaticDashSubstitution(sender objc.IObject)
	ToggleAutomaticDataDetection(sender objc.IObject)
	ToggleAutomaticSpellingCorrection(sender objc.IObject)
	ToggleAutomaticTextReplacement(sender objc.IObject)
	PerformValidatedReplacementInRange_WithAttributedString(_range foundation.Range, attributedString foundation.IAttributedString) bool
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
	rv := ffi.CallMethod[TextView](t_, "initWithFrame:textContainer:", frameRect, container)
	return rv
}

func (t_ TextView) InitWithFrame(frameRect foundation.Rect) TextView {
	rv := ffi.CallMethod[TextView](t_, "initWithFrame:", frameRect)
	return rv
}

func (tc _TextViewClass) FieldEditor() TextView {
	rv := ffi.CallMethod[TextView](tc, "fieldEditor")
	return rv
}

func (t_ TextView) Init() TextView {
	rv := ffi.CallMethod[TextView](t_, "init")
	return rv
}

func (tc _TextViewClass) Alloc() TextView {
	rv := ffi.CallMethod[TextView](tc, "alloc")
	return rv
}

func (tc _TextViewClass) New() TextView {
	rv := ffi.CallMethod[TextView](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextView() TextView {
	return TextViewClass.New()
}

func (tc _TextViewClass) RegisterForServices() {
	ffi.CallMethod[ffi.Void](tc, "registerForServices")
}

func (t_ TextView) ReplaceTextContainer(newContainer ITextContainer) {
	ffi.CallMethod[ffi.Void](t_, "replaceTextContainer:", newContainer)
}

func (t_ TextView) InvalidateTextContainerOrigin() {
	ffi.CallMethod[ffi.Void](t_, "invalidateTextContainerOrigin")
}

func (t_ TextView) ChangeDocumentBackgroundColor(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "changeDocumentBackgroundColor:", sender)
}

func (t_ TextView) SetNeedsDisplayInRect_AvoidAdditionalLayout(rect foundation.Rect, flag bool) {
	ffi.CallMethod[ffi.Void](t_, "setNeedsDisplayInRect:avoidAdditionalLayout:", rect, flag)
}

func (t_ TextView) DrawInsertionPointInRect_Color_TurnedOn(rect foundation.Rect, color IColor, flag bool) {
	ffi.CallMethod[ffi.Void](t_, "drawInsertionPointInRect:color:turnedOn:", rect, color, flag)
}

func (t_ TextView) DrawViewBackgroundInRect(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](t_, "drawViewBackgroundInRect:", rect)
}

func (t_ TextView) SetConstrainedFrameSize(desiredSize foundation.Size) {
	ffi.CallMethod[ffi.Void](t_, "setConstrainedFrameSize:", desiredSize)
}

func (t_ TextView) CleanUpAfterDragOperation() {
	ffi.CallMethod[ffi.Void](t_, "cleanUpAfterDragOperation")
}

func (t_ TextView) ShowFindIndicatorForRange(charRange foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "showFindIndicatorForRange:", charRange)
}

func (tc _TextViewClass) ScrollableDocumentContentTextView() ScrollView {
	rv := ffi.CallMethod[ScrollView](tc, "scrollableDocumentContentTextView")
	return rv
}

func (tc _TextViewClass) ScrollablePlainDocumentContentTextView() ScrollView {
	rv := ffi.CallMethod[ScrollView](tc, "scrollablePlainDocumentContentTextView")
	return rv
}

func (tc _TextViewClass) ScrollableTextView() ScrollView {
	rv := ffi.CallMethod[ScrollView](tc, "scrollableTextView")
	return rv
}

// deprecated
func (t_ TextView) InsertText(insertString objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "insertText:", insertString)
}

func (t_ TextView) SetBaseWritingDirection_Range(writingDirection WritingDirection, _range foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "setBaseWritingDirection:range:", writingDirection, _range)
}

// deprecated
func (t_ TextView) ToggleBaseWritingDirection(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "toggleBaseWritingDirection:", sender)
}

func (t_ TextView) Outline(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "outline:", sender)
}

func (t_ TextView) ToggleAutomaticQuoteSubstitution(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "toggleAutomaticQuoteSubstitution:", sender)
}

func (t_ TextView) ToggleAutomaticLinkDetection(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "toggleAutomaticLinkDetection:", sender)
}

func (t_ TextView) ToggleAutomaticTextCompletion(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "toggleAutomaticTextCompletion:", sender)
}

func (t_ TextView) SetSelectedRange_Affinity_StillSelecting(charRange foundation.Range, affinity SelectionAffinity, stillSelectingFlag bool) {
	ffi.CallMethod[ffi.Void](t_, "setSelectedRange:affinity:stillSelecting:", charRange, affinity, stillSelectingFlag)
}

func (t_ TextView) SetSelectedRanges_Affinity_StillSelecting(ranges []foundation.IValue, affinity SelectionAffinity, stillSelectingFlag bool) {
	ffi.CallMethod[ffi.Void](t_, "setSelectedRanges:affinity:stillSelecting:", ranges, affinity, stillSelectingFlag)
}

func (t_ TextView) UpdateInsertionPointStateAndRestartTimer(restartFlag bool) {
	ffi.CallMethod[ffi.Void](t_, "updateInsertionPointStateAndRestartTimer:", restartFlag)
}

func (t_ TextView) CharacterIndexForInsertionAtPoint(point foundation.Point) uint {
	rv := ffi.CallMethod[uint](t_, "characterIndexForInsertionAtPoint:", point)
	return rv
}

func (t_ TextView) UpdateCandidates() {
	ffi.CallMethod[ffi.Void](t_, "updateCandidates")
}

func (t_ TextView) PreferredPasteboardTypeFromArray_RestrictedToTypesFromArray(availableTypes []PasteboardType, allowedTypes []PasteboardType) PasteboardType {
	rv := ffi.CallMethod[PasteboardType](t_, "preferredPasteboardTypeFromArray:restrictedToTypesFromArray:", availableTypes, allowedTypes)
	return rv
}

func (t_ TextView) ReadSelectionFromPasteboard(pboard IPasteboard) bool {
	rv := ffi.CallMethod[bool](t_, "readSelectionFromPasteboard:", pboard)
	return rv
}

func (t_ TextView) ReadSelectionFromPasteboard_Type(pboard IPasteboard, _type PasteboardType) bool {
	rv := ffi.CallMethod[bool](t_, "readSelectionFromPasteboard:type:", pboard, _type)
	return rv
}

func (t_ TextView) WriteSelectionToPasteboard_Type(pboard IPasteboard, _type PasteboardType) bool {
	rv := ffi.CallMethod[bool](t_, "writeSelectionToPasteboard:type:", pboard, _type)
	return rv
}

func (t_ TextView) WriteSelectionToPasteboard_Types(pboard IPasteboard, types []PasteboardType) bool {
	rv := ffi.CallMethod[bool](t_, "writeSelectionToPasteboard:types:", pboard, types)
	return rv
}

func (t_ TextView) AlignJustified(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "alignJustified:", sender)
}

func (t_ TextView) ChangeAttributes(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "changeAttributes:", sender)
}

func (t_ TextView) ChangeColor(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "changeColor:", sender)
}

func (t_ TextView) SetAlignment_Range(alignment TextAlignment, _range foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "setAlignment:range:", alignment, _range)
}

func (t_ TextView) UseStandardKerning(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "useStandardKerning:", sender)
}

func (t_ TextView) LowerBaseline(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "lowerBaseline:", sender)
}

func (t_ TextView) RaiseBaseline(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "raiseBaseline:", sender)
}

func (t_ TextView) TurnOffKerning(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "turnOffKerning:", sender)
}

func (t_ TextView) LoosenKerning(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "loosenKerning:", sender)
}

func (t_ TextView) TightenKerning(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "tightenKerning:", sender)
}

func (t_ TextView) UseStandardLigatures(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "useStandardLigatures:", sender)
}

func (t_ TextView) TurnOffLigatures(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "turnOffLigatures:", sender)
}

func (t_ TextView) UseAllLigatures(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "useAllLigatures:", sender)
}

// deprecated
func (t_ TextView) ToggleTraditionalCharacterShape(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "toggleTraditionalCharacterShape:", sender)
}

func (t_ TextView) ClickedOnLink_AtIndex(link objc.IObject, charIndex uint) {
	ffi.CallMethod[ffi.Void](t_, "clickedOnLink:atIndex:", link, charIndex)
}

func (t_ TextView) PasteAsPlainText(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "pasteAsPlainText:", sender)
}

func (t_ TextView) PasteAsRichText(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "pasteAsRichText:", sender)
}

func (t_ TextView) BreakUndoCoalescing() {
	ffi.CallMethod[ffi.Void](t_, "breakUndoCoalescing")
}

func (t_ TextView) UpdateFontPanel() {
	ffi.CallMethod[ffi.Void](t_, "updateFontPanel")
}

func (t_ TextView) UpdateRuler() {
	ffi.CallMethod[ffi.Void](t_, "updateRuler")
}

func (t_ TextView) UpdateDragTypeRegistration() {
	ffi.CallMethod[ffi.Void](t_, "updateDragTypeRegistration")
}

func (t_ TextView) SelectionRangeForProposedRange_Granularity(proposedCharRange foundation.Range, granularity SelectionGranularity) foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "selectionRangeForProposedRange:granularity:", proposedCharRange, granularity)
	return rv
}

func (t_ TextView) ShouldChangeTextInRange_ReplacementString(affectedCharRange foundation.Range, replacementString string) bool {
	rv := ffi.CallMethod[bool](t_, "shouldChangeTextInRange:replacementString:", affectedCharRange, replacementString)
	return rv
}

func (t_ TextView) ShouldChangeTextInRanges_ReplacementStrings(affectedRanges []foundation.IValue, replacementStrings []string) bool {
	rv := ffi.CallMethod[bool](t_, "shouldChangeTextInRanges:replacementStrings:", affectedRanges, replacementStrings)
	return rv
}

func (t_ TextView) DidChangeText() {
	ffi.CallMethod[ffi.Void](t_, "didChangeText")
}

func (t_ TextView) SmartDeleteRangeForProposedRange(proposedCharRange foundation.Range) foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "smartDeleteRangeForProposedRange:", proposedCharRange)
	return rv
}

func (t_ TextView) SmartInsertAfterStringForString_ReplacingRange(pasteString string, charRangeToReplace foundation.Range) string {
	rv := ffi.CallMethod[string](t_, "smartInsertAfterStringForString:replacingRange:", pasteString, charRangeToReplace)
	return rv
}

func (t_ TextView) SmartInsertBeforeStringForString_ReplacingRange(pasteString string, charRangeToReplace foundation.Range) string {
	rv := ffi.CallMethod[string](t_, "smartInsertBeforeStringForString:replacingRange:", pasteString, charRangeToReplace)
	return rv
}

func (t_ TextView) SmartInsertForString_ReplacingRange_BeforeString_AfterString(pasteString string, charRangeToReplace foundation.Range, beforeString *foundation.String, afterString *foundation.String) {
	ffi.CallMethod[ffi.Void](t_, "smartInsertForString:replacingRange:beforeString:afterString:", pasteString, charRangeToReplace, beforeString, afterString)
}

func (t_ TextView) ToggleSmartInsertDelete(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "toggleSmartInsertDelete:", sender)
}

func (t_ TextView) ToggleContinuousSpellChecking(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "toggleContinuousSpellChecking:", sender)
}

func (t_ TextView) ToggleGrammarChecking(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "toggleGrammarChecking:", sender)
}

func (t_ TextView) SetSpellingState_Range(value int, charRange foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "setSpellingState:range:", value, charRange)
}

func (t_ TextView) OrderFrontSharingServicePicker(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "orderFrontSharingServicePicker:", sender)
}

func (t_ TextView) DragImageForSelectionWithEvent_Origin(event IEvent, origin *foundation.Point) Image {
	rv := ffi.CallMethod[Image](t_, "dragImageForSelectionWithEvent:origin:", event, origin)
	return rv
}

func (t_ TextView) DragOperationForDraggingInfo_Type(dragInfo DraggingInfo, _type PasteboardType) DragOperation {
	po := ffi.CreateProtocol(dragInfo)
	defer po.Release()
	rv := ffi.CallMethod[DragOperation](t_, "dragOperationForDraggingInfo:type:", po, _type)
	return rv
}

func (t_ TextView) DragSelectionWithEvent_Offset_SlideBack(event IEvent, mouseOffset foundation.Size, slideBack bool) bool {
	rv := ffi.CallMethod[bool](t_, "dragSelectionWithEvent:offset:slideBack:", event, mouseOffset, slideBack)
	return rv
}

func (t_ TextView) StartSpeaking(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "startSpeaking:", sender)
}

func (t_ TextView) StopSpeaking(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "stopSpeaking:", sender)
}

func (t_ TextView) PerformFindPanelAction(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "performFindPanelAction:", sender)
}

func (t_ TextView) OrderFrontLinkPanel(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "orderFrontLinkPanel:", sender)
}

func (t_ TextView) OrderFrontListPanel(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "orderFrontListPanel:", sender)
}

func (t_ TextView) OrderFrontSpacingPanel(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "orderFrontSpacingPanel:", sender)
}

func (t_ TextView) OrderFrontTablePanel(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "orderFrontTablePanel:", sender)
}

func (t_ TextView) OrderFrontSubstitutionsPanel(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "orderFrontSubstitutionsPanel:", sender)
}

func (t_ TextView) Complete(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "complete:", sender)
}

func (t_ TextView) CompletionsForPartialWordRange_IndexOfSelectedItem(charRange foundation.Range, index *int) []string {
	rv := ffi.CallMethod[[]string](t_, "completionsForPartialWordRange:indexOfSelectedItem:", charRange, index)
	return rv
}

func (t_ TextView) InsertCompletion_ForPartialWordRange_Movement_IsFinal(word string, charRange foundation.Range, movement int, flag bool) {
	ffi.CallMethod[ffi.Void](t_, "insertCompletion:forPartialWordRange:movement:isFinal:", word, charRange, movement, flag)
}

func (t_ TextView) CheckTextInDocument(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "checkTextInDocument:", sender)
}

func (t_ TextView) CheckTextInSelection(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "checkTextInSelection:", sender)
}

func (t_ TextView) CheckTextInRange_Types_Options(_range foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "checkTextInRange:types:options:", _range, checkingTypes, options)
}

func (t_ TextView) HandleTextCheckingResults_ForRange_Types_Options_Orthography_WordCount(results []foundation.ITextCheckingResult, _range foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject, orthography foundation.IOrthography, wordCount int) {
	ffi.CallMethod[ffi.Void](t_, "handleTextCheckingResults:forRange:types:options:orthography:wordCount:", results, _range, checkingTypes, options, orthography, wordCount)
}

func (t_ TextView) ToggleAutomaticDashSubstitution(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "toggleAutomaticDashSubstitution:", sender)
}

func (t_ TextView) ToggleAutomaticDataDetection(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "toggleAutomaticDataDetection:", sender)
}

func (t_ TextView) ToggleAutomaticSpellingCorrection(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "toggleAutomaticSpellingCorrection:", sender)
}

func (t_ TextView) ToggleAutomaticTextReplacement(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "toggleAutomaticTextReplacement:", sender)
}

func (t_ TextView) PerformValidatedReplacementInRange_WithAttributedString(_range foundation.Range, attributedString foundation.IAttributedString) bool {
	rv := ffi.CallMethod[bool](t_, "performValidatedReplacementInRange:withAttributedString:", _range, attributedString)
	return rv
}

func (t_ TextView) UpdateQuickLookPreviewPanel() {
	ffi.CallMethod[ffi.Void](t_, "updateQuickLookPreviewPanel")
}

func (t_ TextView) ToggleQuickLookPreviewPanel(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "toggleQuickLookPreviewPanel:", sender)
}

func (t_ TextView) ChangeLayoutOrientation(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "changeLayoutOrientation:", sender)
}

func (t_ TextView) SetLayoutOrientation(orientation TextLayoutOrientation) {
	ffi.CallMethod[ffi.Void](t_, "setLayoutOrientation:", orientation)
}

func (t_ TextView) UpdateTextTouchBarItems() {
	ffi.CallMethod[ffi.Void](t_, "updateTextTouchBarItems")
}

func (t_ TextView) UpdateTouchBarItemIdentifiers() {
	ffi.CallMethod[ffi.Void](t_, "updateTouchBarItemIdentifiers")
}

func (tc _TextViewClass) StronglyReferencesTextStorage() bool {
	rv := ffi.CallMethod[bool](tc, "stronglyReferencesTextStorage")
	return rv
}

func (t_ TextView) TextContainer() TextContainer {
	rv := ffi.CallMethod[TextContainer](t_, "textContainer")
	return rv
}

func (t_ TextView) SetTextContainer(value ITextContainer) {
	ffi.CallMethod[ffi.Void](t_, "setTextContainer:", value)
}

func (t_ TextView) TextContainerInset() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](t_, "textContainerInset")
	return rv
}

func (t_ TextView) SetTextContainerInset(value foundation.Size) {
	ffi.CallMethod[ffi.Void](t_, "setTextContainerInset:", value)
}

func (t_ TextView) TextContainerOrigin() foundation.Point {
	rv := ffi.CallMethod[foundation.Point](t_, "textContainerOrigin")
	return rv
}

func (t_ TextView) TextLayoutManager() TextLayoutManager {
	rv := ffi.CallMethod[TextLayoutManager](t_, "textLayoutManager")
	return rv
}

func (t_ TextView) LayoutManager() LayoutManager {
	rv := ffi.CallMethod[LayoutManager](t_, "layoutManager")
	return rv
}

func (t_ TextView) TextContentStorage() TextContentStorage {
	rv := ffi.CallMethod[TextContentStorage](t_, "textContentStorage")
	return rv
}

func (t_ TextView) TextStorage() TextStorage {
	rv := ffi.CallMethod[TextStorage](t_, "textStorage")
	return rv
}

func (t_ TextView) AllowsDocumentBackgroundColorChange() bool {
	rv := ffi.CallMethod[bool](t_, "allowsDocumentBackgroundColorChange")
	return rv
}

func (t_ TextView) SetAllowsDocumentBackgroundColorChange(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAllowsDocumentBackgroundColorChange:", value)
}

func (t_ TextView) ShouldDrawInsertionPoint() bool {
	rv := ffi.CallMethod[bool](t_, "shouldDrawInsertionPoint")
	return rv
}

func (t_ TextView) AllowedInputSourceLocales() []string {
	rv := ffi.CallMethod[[]string](t_, "allowedInputSourceLocales")
	return rv
}

func (t_ TextView) SetAllowedInputSourceLocales(value []string) {
	ffi.CallMethod[ffi.Void](t_, "setAllowedInputSourceLocales:", value)
}

func (t_ TextView) AllowsUndo() bool {
	rv := ffi.CallMethod[bool](t_, "allowsUndo")
	return rv
}

func (t_ TextView) SetAllowsUndo(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAllowsUndo:", value)
}

func (t_ TextView) DefaultParagraphStyle() ParagraphStyle {
	rv := ffi.CallMethod[ParagraphStyle](t_, "defaultParagraphStyle")
	return rv
}

func (t_ TextView) SetDefaultParagraphStyle(value IParagraphStyle) {
	ffi.CallMethod[ffi.Void](t_, "setDefaultParagraphStyle:", value)
}

func (t_ TextView) AllowsImageEditing() bool {
	rv := ffi.CallMethod[bool](t_, "allowsImageEditing")
	return rv
}

func (t_ TextView) SetAllowsImageEditing(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAllowsImageEditing:", value)
}

func (t_ TextView) IsAutomaticQuoteSubstitutionEnabled() bool {
	rv := ffi.CallMethod[bool](t_, "isAutomaticQuoteSubstitutionEnabled")
	return rv
}

func (t_ TextView) SetAutomaticQuoteSubstitutionEnabled(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAutomaticQuoteSubstitutionEnabled:", value)
}

func (t_ TextView) IsAutomaticLinkDetectionEnabled() bool {
	rv := ffi.CallMethod[bool](t_, "isAutomaticLinkDetectionEnabled")
	return rv
}

func (t_ TextView) SetAutomaticLinkDetectionEnabled(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAutomaticLinkDetectionEnabled:", value)
}

func (t_ TextView) DisplaysLinkToolTips() bool {
	rv := ffi.CallMethod[bool](t_, "displaysLinkToolTips")
	return rv
}

func (t_ TextView) SetDisplaysLinkToolTips(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setDisplaysLinkToolTips:", value)
}

func (t_ TextView) IsAutomaticTextCompletionEnabled() bool {
	rv := ffi.CallMethod[bool](t_, "isAutomaticTextCompletionEnabled")
	return rv
}

func (t_ TextView) SetAutomaticTextCompletionEnabled(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAutomaticTextCompletionEnabled:", value)
}

func (t_ TextView) UsesAdaptiveColorMappingForDarkAppearance() bool {
	rv := ffi.CallMethod[bool](t_, "usesAdaptiveColorMappingForDarkAppearance")
	return rv
}

func (t_ TextView) SetUsesAdaptiveColorMappingForDarkAppearance(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setUsesAdaptiveColorMappingForDarkAppearance:", value)
}

func (t_ TextView) UsesRolloverButtonForSelection() bool {
	rv := ffi.CallMethod[bool](t_, "usesRolloverButtonForSelection")
	return rv
}

func (t_ TextView) SetUsesRolloverButtonForSelection(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setUsesRolloverButtonForSelection:", value)
}

func (t_ TextView) UsesRuler() bool {
	rv := ffi.CallMethod[bool](t_, "usesRuler")
	return rv
}

func (t_ TextView) SetUsesRuler(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setUsesRuler:", value)
}

func (t_ TextView) SetRulerVisible(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setRulerVisible:", value)
}

func (t_ TextView) UsesInspectorBar() bool {
	rv := ffi.CallMethod[bool](t_, "usesInspectorBar")
	return rv
}

func (t_ TextView) SetUsesInspectorBar(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setUsesInspectorBar:", value)
}

func (t_ TextView) SelectedRanges() []foundation.Value {
	rv := ffi.CallMethod[[]foundation.Value](t_, "selectedRanges")
	return rv
}

func (t_ TextView) SetSelectedRanges(value []foundation.IValue) {
	ffi.CallMethod[ffi.Void](t_, "setSelectedRanges:", value)
}

func (t_ TextView) SelectionAffinity() SelectionAffinity {
	rv := ffi.CallMethod[SelectionAffinity](t_, "selectionAffinity")
	return rv
}

func (t_ TextView) SelectionGranularity() SelectionGranularity {
	rv := ffi.CallMethod[SelectionGranularity](t_, "selectionGranularity")
	return rv
}

func (t_ TextView) SetSelectionGranularity(value SelectionGranularity) {
	ffi.CallMethod[ffi.Void](t_, "setSelectionGranularity:", value)
}

func (t_ TextView) InsertionPointColor() Color {
	rv := ffi.CallMethod[Color](t_, "insertionPointColor")
	return rv
}

func (t_ TextView) SetInsertionPointColor(value IColor) {
	ffi.CallMethod[ffi.Void](t_, "setInsertionPointColor:", value)
}

func (t_ TextView) SelectedTextAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := ffi.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, "selectedTextAttributes")
	return rv
}

func (t_ TextView) SetSelectedTextAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "setSelectedTextAttributes:", value)
}

func (t_ TextView) MarkedTextAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := ffi.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, "markedTextAttributes")
	return rv
}

func (t_ TextView) SetMarkedTextAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "setMarkedTextAttributes:", value)
}

func (t_ TextView) LinkTextAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := ffi.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, "linkTextAttributes")
	return rv
}

func (t_ TextView) SetLinkTextAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "setLinkTextAttributes:", value)
}

func (t_ TextView) ReadablePasteboardTypes() []PasteboardType {
	rv := ffi.CallMethod[[]PasteboardType](t_, "readablePasteboardTypes")
	return rv
}

func (t_ TextView) WritablePasteboardTypes() []PasteboardType {
	rv := ffi.CallMethod[[]PasteboardType](t_, "writablePasteboardTypes")
	return rv
}

func (t_ TextView) TypingAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := ffi.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, "typingAttributes")
	return rv
}

func (t_ TextView) SetTypingAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "setTypingAttributes:", value)
}

func (t_ TextView) IsCoalescingUndo() bool {
	rv := ffi.CallMethod[bool](t_, "isCoalescingUndo")
	return rv
}

func (t_ TextView) AcceptableDragTypes() []PasteboardType {
	rv := ffi.CallMethod[[]PasteboardType](t_, "acceptableDragTypes")
	return rv
}

func (t_ TextView) RangeForUserCharacterAttributeChange() foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "rangeForUserCharacterAttributeChange")
	return rv
}

func (t_ TextView) RangesForUserCharacterAttributeChange() []foundation.Value {
	rv := ffi.CallMethod[[]foundation.Value](t_, "rangesForUserCharacterAttributeChange")
	return rv
}

func (t_ TextView) RangeForUserParagraphAttributeChange() foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "rangeForUserParagraphAttributeChange")
	return rv
}

func (t_ TextView) RangesForUserParagraphAttributeChange() []foundation.Value {
	rv := ffi.CallMethod[[]foundation.Value](t_, "rangesForUserParagraphAttributeChange")
	return rv
}

func (t_ TextView) RangeForUserTextChange() foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "rangeForUserTextChange")
	return rv
}

func (t_ TextView) RangesForUserTextChange() []foundation.Value {
	rv := ffi.CallMethod[[]foundation.Value](t_, "rangesForUserTextChange")
	return rv
}

func (t_ TextView) SmartInsertDeleteEnabled() bool {
	rv := ffi.CallMethod[bool](t_, "smartInsertDeleteEnabled")
	return rv
}

func (t_ TextView) SetSmartInsertDeleteEnabled(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setSmartInsertDeleteEnabled:", value)
}

func (t_ TextView) IsContinuousSpellCheckingEnabled() bool {
	rv := ffi.CallMethod[bool](t_, "isContinuousSpellCheckingEnabled")
	return rv
}

func (t_ TextView) SetContinuousSpellCheckingEnabled(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setContinuousSpellCheckingEnabled:", value)
}

func (t_ TextView) SpellCheckerDocumentTag() int {
	rv := ffi.CallMethod[int](t_, "spellCheckerDocumentTag")
	return rv
}

func (t_ TextView) IsGrammarCheckingEnabled() bool {
	rv := ffi.CallMethod[bool](t_, "isGrammarCheckingEnabled")
	return rv
}

func (t_ TextView) SetGrammarCheckingEnabled(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setGrammarCheckingEnabled:", value)
}

func (t_ TextView) AcceptsGlyphInfo() bool {
	rv := ffi.CallMethod[bool](t_, "acceptsGlyphInfo")
	return rv
}

func (t_ TextView) SetAcceptsGlyphInfo(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAcceptsGlyphInfo:", value)
}

func (t_ TextView) UsesFindPanel() bool {
	rv := ffi.CallMethod[bool](t_, "usesFindPanel")
	return rv
}

func (t_ TextView) SetUsesFindPanel(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setUsesFindPanel:", value)
}

func (t_ TextView) RangeForUserCompletion() foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "rangeForUserCompletion")
	return rv
}

func (t_ TextView) EnabledTextCheckingTypes() foundation.TextCheckingTypes {
	rv := ffi.CallMethod[foundation.TextCheckingTypes](t_, "enabledTextCheckingTypes")
	return rv
}

func (t_ TextView) SetEnabledTextCheckingTypes(value foundation.TextCheckingTypes) {
	ffi.CallMethod[ffi.Void](t_, "setEnabledTextCheckingTypes:", value)
}

func (t_ TextView) IsAutomaticDashSubstitutionEnabled() bool {
	rv := ffi.CallMethod[bool](t_, "isAutomaticDashSubstitutionEnabled")
	return rv
}

func (t_ TextView) SetAutomaticDashSubstitutionEnabled(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAutomaticDashSubstitutionEnabled:", value)
}

func (t_ TextView) IsAutomaticDataDetectionEnabled() bool {
	rv := ffi.CallMethod[bool](t_, "isAutomaticDataDetectionEnabled")
	return rv
}

func (t_ TextView) SetAutomaticDataDetectionEnabled(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAutomaticDataDetectionEnabled:", value)
}

func (t_ TextView) IsAutomaticSpellingCorrectionEnabled() bool {
	rv := ffi.CallMethod[bool](t_, "isAutomaticSpellingCorrectionEnabled")
	return rv
}

func (t_ TextView) SetAutomaticSpellingCorrectionEnabled(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAutomaticSpellingCorrectionEnabled:", value)
}

func (t_ TextView) IsAutomaticTextReplacementEnabled() bool {
	rv := ffi.CallMethod[bool](t_, "isAutomaticTextReplacementEnabled")
	return rv
}

func (t_ TextView) SetAutomaticTextReplacementEnabled(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAutomaticTextReplacementEnabled:", value)
}

func (t_ TextView) UsesFindBar() bool {
	rv := ffi.CallMethod[bool](t_, "usesFindBar")
	return rv
}

func (t_ TextView) SetUsesFindBar(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setUsesFindBar:", value)
}

func (t_ TextView) IsIncrementalSearchingEnabled() bool {
	rv := ffi.CallMethod[bool](t_, "isIncrementalSearchingEnabled")
	return rv
}

func (t_ TextView) SetIncrementalSearchingEnabled(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setIncrementalSearchingEnabled:", value)
}

func (t_ TextView) AllowsCharacterPickerTouchBarItem() bool {
	rv := ffi.CallMethod[bool](t_, "allowsCharacterPickerTouchBarItem")
	return rv
}

func (t_ TextView) SetAllowsCharacterPickerTouchBarItem(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAllowsCharacterPickerTouchBarItem:", value)
}

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
	TextView_WillCheckTextInRange_Options_Types(view TextView, _range foundation.Range, options map[TextCheckingOptionKey]objc.Object, checkingTypes *foundation.TextCheckingTypes) map[TextCheckingOptionKey]objc.IObject
	ImplementsTextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount() bool
	// optional
	TextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount(view TextView, _range foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.Object, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.ITextCheckingResult
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
	_TextView_WillCheckTextInRange_Options_Types                              func(view TextView, _range foundation.Range, options map[TextCheckingOptionKey]objc.Object, checkingTypes *foundation.TextCheckingTypes) map[TextCheckingOptionKey]objc.IObject
	_TextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount func(view TextView, _range foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.Object, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.ITextCheckingResult
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

func (di *TextViewDelegateImpl) SetTextView_WillCheckTextInRange_Options_Types(f func(view TextView, _range foundation.Range, options map[TextCheckingOptionKey]objc.Object, checkingTypes *foundation.TextCheckingTypes) map[TextCheckingOptionKey]objc.IObject) {
	di._TextView_WillCheckTextInRange_Options_Types = f
}

func (di *TextViewDelegateImpl) TextView_WillCheckTextInRange_Options_Types(view TextView, _range foundation.Range, options map[TextCheckingOptionKey]objc.Object, checkingTypes *foundation.TextCheckingTypes) map[TextCheckingOptionKey]objc.IObject {
	return di._TextView_WillCheckTextInRange_Options_Types(view, _range, options, checkingTypes)
}
func (di *TextViewDelegateImpl) ImplementsTextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount() bool {
	return di._TextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount != nil
}

func (di *TextViewDelegateImpl) SetTextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount(f func(view TextView, _range foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.Object, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.ITextCheckingResult) {
	di._TextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount = f
}

func (di *TextViewDelegateImpl) TextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount(view TextView, _range foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.Object, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.ITextCheckingResult {
	return di._TextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount(view, _range, checkingTypes, options, results, orthography, wordCount)
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
	rv := ffi.CallMethod[foundation.UndoManager](t_, "undoManagerForTextView:", view)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_WillDisplayToolTip_ForCharacterAtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:willDisplayToolTip:forCharacterAtIndex:"))
}

func (t_ TextViewDelegateWrapper) TextView_WillDisplayToolTip_ForCharacterAtIndex(textView ITextView, tooltip string, characterIndex uint) string {
	rv := ffi.CallMethod[string](t_, "textView:willDisplayToolTip:forCharacterAtIndex:", textView, tooltip, characterIndex)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_URLForContentsOfTextAttachment_AtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:URLForContentsOfTextAttachment:atIndex:"))
}

func (t_ TextViewDelegateWrapper) TextView_URLForContentsOfTextAttachment_AtIndex(textView ITextView, textAttachment ITextAttachment, charIndex uint) foundation.URL {
	rv := ffi.CallMethod[foundation.URL](t_, "textView:URLForContentsOfTextAttachment:atIndex:", textView, textAttachment, charIndex)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_WillChangeSelectionFromCharacterRange_ToCharacterRange() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:willChangeSelectionFromCharacterRange:toCharacterRange:"))
}

func (t_ TextViewDelegateWrapper) TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange(textView ITextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "textView:willChangeSelectionFromCharacterRange:toCharacterRange:", textView, oldSelectedCharRange, newSelectedCharRange)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:willChangeSelectionFromCharacterRanges:toCharacterRanges:"))
}

func (t_ TextViewDelegateWrapper) TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges(textView ITextView, oldSelectedCharRanges []foundation.IValue, newSelectedCharRanges []foundation.IValue) []foundation.Value {
	rv := ffi.CallMethod[[]foundation.Value](t_, "textView:willChangeSelectionFromCharacterRanges:toCharacterRanges:", textView, oldSelectedCharRanges, newSelectedCharRanges)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextViewDidChangeSelection() bool {
	return t_.RespondsToSelector(objc.GetSelector("textViewDidChangeSelection:"))
}

func (t_ TextViewDelegateWrapper) TextViewDidChangeSelection(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](t_, "textViewDidChangeSelection:", notification)
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_Candidates_ForSelectedRange() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:candidates:forSelectedRange:"))
}

func (t_ TextViewDelegateWrapper) TextView_Candidates_ForSelectedRange(textView ITextView, candidates []foundation.ITextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult {
	rv := ffi.CallMethod[[]foundation.TextCheckingResult](t_, "textView:candidates:forSelectedRange:", textView, candidates, selectedRange)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_CandidatesForSelectedRange() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:candidatesForSelectedRange:"))
}

func (t_ TextViewDelegateWrapper) TextView_CandidatesForSelectedRange(textView ITextView, selectedRange foundation.Range) []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](t_, "textView:candidatesForSelectedRange:", textView, selectedRange)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_ShouldSelectCandidateAtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:shouldSelectCandidateAtIndex:"))
}

func (t_ TextViewDelegateWrapper) TextView_ShouldSelectCandidateAtIndex(textView ITextView, index uint) bool {
	rv := ffi.CallMethod[bool](t_, "textView:shouldSelectCandidateAtIndex:", textView, index)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_ShouldUpdateTouchBarItemIdentifiers() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:shouldUpdateTouchBarItemIdentifiers:"))
}

func (t_ TextViewDelegateWrapper) TextView_ShouldUpdateTouchBarItemIdentifiers(textView ITextView, identifiers []TouchBarItemIdentifier) []TouchBarItemIdentifier {
	rv := ffi.CallMethod[[]TouchBarItemIdentifier](t_, "textView:shouldUpdateTouchBarItemIdentifiers:", textView, identifiers)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_ShouldChangeTextInRange_ReplacementString() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:shouldChangeTextInRange:replacementString:"))
}

func (t_ TextViewDelegateWrapper) TextView_ShouldChangeTextInRange_ReplacementString(textView ITextView, affectedCharRange foundation.Range, replacementString string) bool {
	rv := ffi.CallMethod[bool](t_, "textView:shouldChangeTextInRange:replacementString:", textView, affectedCharRange, replacementString)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_ShouldChangeTextInRanges_ReplacementStrings() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:shouldChangeTextInRanges:replacementStrings:"))
}

func (t_ TextViewDelegateWrapper) TextView_ShouldChangeTextInRanges_ReplacementStrings(textView ITextView, affectedRanges []foundation.IValue, replacementStrings []string) bool {
	rv := ffi.CallMethod[bool](t_, "textView:shouldChangeTextInRanges:replacementStrings:", textView, affectedRanges, replacementStrings)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_ShouldChangeTypingAttributes_ToAttributes() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:shouldChangeTypingAttributes:toAttributes:"))
}

func (t_ TextViewDelegateWrapper) TextView_ShouldChangeTypingAttributes_ToAttributes(textView ITextView, oldTypingAttributes map[string]objc.IObject, newTypingAttributes map[foundation.AttributedStringKey]objc.IObject) map[foundation.AttributedStringKey]objc.Object {
	rv := ffi.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, "textView:shouldChangeTypingAttributes:toAttributes:", textView, oldTypingAttributes, newTypingAttributes)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextViewDidChangeTypingAttributes() bool {
	return t_.RespondsToSelector(objc.GetSelector("textViewDidChangeTypingAttributes:"))
}

func (t_ TextViewDelegateWrapper) TextViewDidChangeTypingAttributes(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](t_, "textViewDidChangeTypingAttributes:", notification)
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_ClickedOnLink_AtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:clickedOnLink:atIndex:"))
}

func (t_ TextViewDelegateWrapper) TextView_ClickedOnLink_AtIndex(textView ITextView, link objc.IObject, charIndex uint) bool {
	rv := ffi.CallMethod[bool](t_, "textView:clickedOnLink:atIndex:", textView, link, charIndex)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_ShouldSetSpellingState_Range() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:shouldSetSpellingState:range:"))
}

func (t_ TextViewDelegateWrapper) TextView_ShouldSetSpellingState_Range(textView ITextView, value int, affectedCharRange foundation.Range) int {
	rv := ffi.CallMethod[int](t_, "textView:shouldSetSpellingState:range:", textView, value, affectedCharRange)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_WillCheckTextInRange_Options_Types() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:willCheckTextInRange:options:types:"))
}

func (t_ TextViewDelegateWrapper) TextView_WillCheckTextInRange_Options_Types(view ITextView, _range foundation.Range, options map[TextCheckingOptionKey]objc.IObject, checkingTypes *foundation.TextCheckingTypes) map[TextCheckingOptionKey]objc.Object {
	rv := ffi.CallMethod[map[TextCheckingOptionKey]objc.Object](t_, "textView:willCheckTextInRange:options:types:", view, _range, options, checkingTypes)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:didCheckTextInRange:types:options:results:orthography:wordCount:"))
}

func (t_ TextViewDelegateWrapper) TextView_DidCheckTextInRange_Types_Options_Results_Orthography_WordCount(view ITextView, _range foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject, results []foundation.ITextCheckingResult, orthography foundation.IOrthography, wordCount int) []foundation.TextCheckingResult {
	rv := ffi.CallMethod[[]foundation.TextCheckingResult](t_, "textView:didCheckTextInRange:types:options:results:orthography:wordCount:", view, _range, checkingTypes, options, results, orthography, wordCount)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_Completions_ForPartialWordRange_IndexOfSelectedItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:completions:forPartialWordRange:indexOfSelectedItem:"))
}

func (t_ TextViewDelegateWrapper) TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(textView ITextView, words []string, charRange foundation.Range, index *int) []string {
	rv := ffi.CallMethod[[]string](t_, "textView:completions:forPartialWordRange:indexOfSelectedItem:", textView, words, charRange, index)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_WillShowSharingServicePicker_ForItems() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:willShowSharingServicePicker:forItems:"))
}

func (t_ TextViewDelegateWrapper) TextView_WillShowSharingServicePicker_ForItems(textView ITextView, servicePicker ISharingServicePicker, items []objc.IObject) SharingServicePicker {
	rv := ffi.CallMethod[SharingServicePicker](t_, "textView:willShowSharingServicePicker:forItems:", textView, servicePicker, items)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_DoCommandBySelector() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:doCommandBySelector:"))
}

func (t_ TextViewDelegateWrapper) TextView_DoCommandBySelector(textView ITextView, commandSelector objc.Selector) bool {
	rv := ffi.CallMethod[bool](t_, "textView:doCommandBySelector:", textView, commandSelector)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_Menu_ForEvent_AtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:menu:forEvent:atIndex:"))
}

func (t_ TextViewDelegateWrapper) TextView_Menu_ForEvent_AtIndex(view ITextView, menu IMenu, event IEvent, charIndex uint) Menu {
	rv := ffi.CallMethod[Menu](t_, "textView:menu:forEvent:atIndex:", view, menu, event, charIndex)
	return rv
}

func (t_ *TextViewDelegateWrapper) ImplementsTextView_ClickedOnLink() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:clickedOnLink:"))
}

// deprecated
func (t_ TextViewDelegateWrapper) TextView_ClickedOnLink(textView ITextView, link objc.IObject) bool {
	rv := ffi.CallMethod[bool](t_, "textView:clickedOnLink:", textView, link)
	return rv
}

var TextViewportLayoutControllerClass = _TextViewportLayoutControllerClass{objc.GetClass("NSTextViewportLayoutController")}

type _TextViewportLayoutControllerClass struct {
	objc.Class
}

type ITextViewportLayoutController interface {
	objc.IObject
	AdjustViewportByVerticalOffset(verticalOffset float64)
	LayoutViewport()
	TextLayoutManager() TextLayoutManager
	ViewportBounds() coregraphics.Rect
	ViewportRange() TextRange
}

type TextViewportLayoutController struct {
	objc.Object
}

func MakeTextViewportLayoutController(ptr unsafe.Pointer) TextViewportLayoutController {
	return TextViewportLayoutController{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextViewportLayoutController) InitWithTextLayoutManager(textLayoutManager ITextLayoutManager) TextViewportLayoutController {
	rv := ffi.CallMethod[TextViewportLayoutController](t_, "initWithTextLayoutManager:", textLayoutManager)
	return rv
}

func (tc _TextViewportLayoutControllerClass) Alloc() TextViewportLayoutController {
	rv := ffi.CallMethod[TextViewportLayoutController](tc, "alloc")
	return rv
}

func (t_ TextViewportLayoutController) Init() TextViewportLayoutController {
	rv := ffi.CallMethod[TextViewportLayoutController](t_, "init")
	return rv
}

func (tc _TextViewportLayoutControllerClass) New() TextViewportLayoutController {
	rv := ffi.CallMethod[TextViewportLayoutController](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextViewportLayoutController() TextViewportLayoutController {
	return TextViewportLayoutControllerClass.New()
}

func (t_ TextViewportLayoutController) AdjustViewportByVerticalOffset(verticalOffset float64) {
	ffi.CallMethod[ffi.Void](t_, "adjustViewportByVerticalOffset:", verticalOffset)
}

func (t_ TextViewportLayoutController) LayoutViewport() {
	ffi.CallMethod[ffi.Void](t_, "layoutViewport")
}

func (t_ TextViewportLayoutController) TextLayoutManager() TextLayoutManager {
	rv := ffi.CallMethod[TextLayoutManager](t_, "textLayoutManager")
	return rv
}

func (t_ TextViewportLayoutController) ViewportBounds() coregraphics.Rect {
	rv := ffi.CallMethod[coregraphics.Rect](t_, "viewportBounds")
	return rv
}

func (t_ TextViewportLayoutController) ViewportRange() TextRange {
	rv := ffi.CallMethod[TextRange](t_, "viewportRange")
	return rv
}
