// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	DragOperationForDraggingInfo_Type(dragInfo DraggingInfo, type_ PasteboardType) DragOperation
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

func (t_ TextView) SetBaseWritingDirection_Range(writingDirection WritingDirection, range_ foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "setBaseWritingDirection:range:", writingDirection, range_)
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

func (t_ TextView) ReadSelectionFromPasteboard_Type(pboard IPasteboard, type_ PasteboardType) bool {
	rv := ffi.CallMethod[bool](t_, "readSelectionFromPasteboard:type:", pboard, type_)
	return rv
}

func (t_ TextView) WriteSelectionToPasteboard_Type(pboard IPasteboard, type_ PasteboardType) bool {
	rv := ffi.CallMethod[bool](t_, "writeSelectionToPasteboard:type:", pboard, type_)
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

func (t_ TextView) SetAlignment_Range(alignment TextAlignment, range_ foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "setAlignment:range:", alignment, range_)
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

func (t_ TextView) DragOperationForDraggingInfo_Type(dragInfo DraggingInfo, type_ PasteboardType) DragOperation {
	po := ffi.CreateProtocol("NSDraggingInfo", dragInfo)
	defer po.Release()
	rv := ffi.CallMethod[DragOperation](t_, "dragOperationForDraggingInfo:type:", po, type_)
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

func (t_ TextView) CheckTextInRange_Types_Options(range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "checkTextInRange:types:options:", range_, checkingTypes, options)
}

func (t_ TextView) HandleTextCheckingResults_ForRange_Types_Options_Orthography_WordCount(results []foundation.ITextCheckingResult, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject, orthography foundation.IOrthography, wordCount int) {
	ffi.CallMethod[ffi.Void](t_, "handleTextCheckingResults:forRange:types:options:orthography:wordCount:", results, range_, checkingTypes, options, orthography, wordCount)
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

func (t_ TextView) PerformValidatedReplacementInRange_WithAttributedString(range_ foundation.Range, attributedString foundation.IAttributedString) bool {
	rv := ffi.CallMethod[bool](t_, "performValidatedReplacementInRange:withAttributedString:", range_, attributedString)
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
