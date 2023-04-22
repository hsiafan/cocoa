// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var CellClass = _CellClass{objc.GetClass("NSCell")}

type _CellClass struct {
	objc.Class
}

type ICell interface {
	objc.IObject
	SetCellAttribute_To(parameter CellAttribute, value int)
	CellAttribute(parameter CellAttribute) int
	SetNextState()
	SetUpFieldEditorAttributes(textObj IText) Text
	SendActionOn(mask EventMask) int
	MenuForEvent_InRect_OfView(event IEvent, cellFrame foundation.Rect, view IView) Menu
	Compare(otherCell objc.IObject) foundation.ComparisonResult
	PerformClick(sender objc.IObject)
	TakeObjectValueFrom(sender objc.IObject)
	TakeIntegerValueFrom(sender objc.IObject)
	TakeIntValueFrom(sender objc.IObject)
	TakeStringValueFrom(sender objc.IObject)
	TakeDoubleValueFrom(sender objc.IObject)
	TakeFloatValueFrom(sender objc.IObject)
	TrackMouse_InRect_OfView_UntilMouseUp(event IEvent, cellFrame foundation.Rect, controlView IView, flag bool) bool
	StartTrackingAt_InView(startPoint foundation.Point, controlView IView) bool
	ContinueTracking_At_InView(lastPoint foundation.Point, currentPoint foundation.Point, controlView IView) bool
	StopTracking_At_InView_MouseIsUp(lastPoint foundation.Point, stopPoint foundation.Point, controlView IView, flag bool)
	GetPeriodicDelay_Interval(delay *float32, interval *float32)
	HitTestForEvent_InRect_OfView(event IEvent, cellFrame foundation.Rect, controlView IView) CellHitResult
	ResetCursorRect_InView(cellFrame foundation.Rect, controlView IView)
	DraggingImageComponentsWithFrame_InView(frame foundation.Rect, view IView) []DraggingImageComponent
	DrawFocusRingMaskWithFrame_InView(cellFrame foundation.Rect, controlView IView)
	FocusRingMaskBoundsForFrame_InView(cellFrame foundation.Rect, controlView IView) foundation.Rect
	CalcDrawInfo(rect foundation.Rect)
	CellSizeForBounds(rect foundation.Rect) foundation.Size
	DrawingRectForBounds(rect foundation.Rect) foundation.Rect
	ImageRectForBounds(rect foundation.Rect) foundation.Rect
	TitleRectForBounds(rect foundation.Rect) foundation.Rect
	DrawWithFrame_InView(cellFrame foundation.Rect, controlView IView)
	HighlightColorWithFrame_InView(cellFrame foundation.Rect, controlView IView) Color
	DrawInteriorWithFrame_InView(cellFrame foundation.Rect, controlView IView)
	Highlight_WithFrame_InView(flag bool, cellFrame foundation.Rect, controlView IView)
	EditWithFrame_InView_Editor_Delegate_Event(rect foundation.Rect, controlView IView, textObj IText, delegate objc.IObject, event IEvent)
	SelectWithFrame_InView_Editor_Delegate_Start_Length(rect foundation.Rect, controlView IView, textObj IText, delegate objc.IObject, selStart int, selLength int)
	EndEditing(textObj IText)
	FieldEditorForView(controlView IView) TextView
	ExpansionFrameWithFrame_InView(cellFrame foundation.Rect, view IView) foundation.Rect
	DrawWithExpansionFrame_InView(cellFrame foundation.Rect, view IView)
	// deprecated
	SetEntryType(type_ int)
	// deprecated
	EntryType() int
	// deprecated
	IsEntryAcceptable(string_ string) bool
	// deprecated
	SetFloatingPointFormat_Left_Right(autoRange bool, leftDigits uint, rightDigits uint)
	// deprecated
	SetTitleWithMnemonic(stringWithAmpersand string)
	// deprecated
	Mnemonic() string
	// deprecated
	SetMnemonicLocation(location uint)
	// deprecated
	MnemonicLocation() uint
	ObjectValue() objc.Object
	SetObjectValue(value objc.IObject)
	HasValidObjectValue() bool
	IntValue() int32
	SetIntValue(value int32)
	IntegerValue() int
	SetIntegerValue(value int)
	StringValue() string
	SetStringValue(value string)
	DoubleValue() float64
	SetDoubleValue(value float64)
	FloatValue() float32
	SetFloatValue(value float32)
	Type() CellType
	SetType(value CellType)
	IsEnabled() bool
	SetEnabled(value bool)
	AllowsUndo() bool
	SetAllowsUndo(value bool)
	IsBezeled() bool
	SetBezeled(value bool)
	IsBordered() bool
	SetBordered(value bool)
	IsOpaque() bool
	// deprecated
	ControlTint() ControlTint
	// deprecated
	SetControlTint(value ControlTint)
	BackgroundStyle() BackgroundStyle
	SetBackgroundStyle(value BackgroundStyle)
	InteriorBackgroundStyle() BackgroundStyle
	AllowsMixedState() bool
	SetAllowsMixedState(value bool)
	NextState() int
	State() ControlStateValue
	SetState(value ControlStateValue)
	IsEditable() bool
	SetEditable(value bool)
	IsSelectable() bool
	SetSelectable(value bool)
	IsScrollable() bool
	SetScrollable(value bool)
	Alignment() TextAlignment
	SetAlignment(value TextAlignment)
	Font() Font
	SetFont(value IFont)
	LineBreakMode() LineBreakMode
	SetLineBreakMode(value LineBreakMode)
	TruncatesLastVisibleLine() bool
	SetTruncatesLastVisibleLine(value bool)
	Wraps() bool
	SetWraps(value bool)
	BaseWritingDirection() WritingDirection
	SetBaseWritingDirection(value WritingDirection)
	AttributedStringValue() foundation.AttributedString
	SetAttributedStringValue(value foundation.IAttributedString)
	AllowsEditingTextAttributes() bool
	SetAllowsEditingTextAttributes(value bool)
	ImportsGraphics() bool
	SetImportsGraphics(value bool)
	Title() string
	SetTitle(value string)
	Action() objc.Selector
	SetAction(value objc.Selector)
	Target() objc.Object
	SetTarget(value objc.IObject)
	IsContinuous() bool
	SetContinuous(value bool)
	Image() Image
	SetImage(value IImage)
	Tag() int
	SetTag(value int)
	Formatter() foundation.Formatter
	SetFormatter(value foundation.IFormatter)
	Menu() Menu
	SetMenu(value IMenu)
	AcceptsFirstResponder() bool
	ShowsFirstResponder() bool
	SetShowsFirstResponder(value bool)
	RefusesFirstResponder() bool
	SetRefusesFirstResponder(value bool)
	RepresentedObject() objc.Object
	SetRepresentedObject(value objc.IObject)
	MouseDownFlags() int
	KeyEquivalent() string
	FocusRingType() FocusRingType
	SetFocusRingType(value FocusRingType)
	CellSize() foundation.Size
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	ControlView() View
	SetControlView(value IView)
	IsHighlighted() bool
	SetHighlighted(value bool)
	SendsActionOnEndEditing() bool
	SetSendsActionOnEndEditing(value bool)
	WantsNotificationForMarkedText() bool
	UsesSingleLineMode() bool
	SetUsesSingleLineMode(value bool)
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
}

type Cell struct {
	objc.Object
}

func MakeCell(ptr unsafe.Pointer) Cell {
	return Cell{
		Object: objc.MakeObject(ptr),
	}
}

func (c_ Cell) InitImageCell(image IImage) Cell {
	rv := objc.CallMethod[Cell](c_, objc.GetSelector("initImageCell:"), image)
	return rv
}

func (c_ Cell) InitTextCell(string_ string) Cell {
	rv := objc.CallMethod[Cell](c_, objc.GetSelector("initTextCell:"), string_)
	return rv
}

func (c_ Cell) Init() Cell {
	rv := objc.CallMethod[Cell](c_, objc.GetSelector("init"))
	return rv
}

func (cc _CellClass) Alloc() Cell {
	rv := objc.CallMethod[Cell](cc, objc.GetSelector("alloc"))
	return rv
}

func (cc _CellClass) New() Cell {
	rv := objc.CallMethod[Cell](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCell() Cell {
	return CellClass.New()
}

func (c_ Cell) SetCellAttribute_To(parameter CellAttribute, value int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setCellAttribute:to:"), parameter, value)
}

func (c_ Cell) CellAttribute(parameter CellAttribute) int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("cellAttribute:"), parameter)
	return rv
}

func (c_ Cell) SetNextState() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setNextState"))
}

func (c_ Cell) SetUpFieldEditorAttributes(textObj IText) Text {
	rv := objc.CallMethod[Text](c_, objc.GetSelector("setUpFieldEditorAttributes:"), textObj)
	return rv
}

func (c_ Cell) SendActionOn(mask EventMask) int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("sendActionOn:"), mask)
	return rv
}

func (c_ Cell) MenuForEvent_InRect_OfView(event IEvent, cellFrame foundation.Rect, view IView) Menu {
	rv := objc.CallMethod[Menu](c_, objc.GetSelector("menuForEvent:inRect:ofView:"), event, cellFrame, view)
	return rv
}

func (c_ Cell) Compare(otherCell objc.IObject) foundation.ComparisonResult {
	rv := objc.CallMethod[foundation.ComparisonResult](c_, objc.GetSelector("compare:"), otherCell)
	return rv
}

func (c_ Cell) PerformClick(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("performClick:"), sender)
}

func (c_ Cell) TakeObjectValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("takeObjectValueFrom:"), sender)
}

func (c_ Cell) TakeIntegerValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("takeIntegerValueFrom:"), sender)
}

func (c_ Cell) TakeIntValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("takeIntValueFrom:"), sender)
}

func (c_ Cell) TakeStringValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("takeStringValueFrom:"), sender)
}

func (c_ Cell) TakeDoubleValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("takeDoubleValueFrom:"), sender)
}

func (c_ Cell) TakeFloatValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("takeFloatValueFrom:"), sender)
}

func (c_ Cell) TrackMouse_InRect_OfView_UntilMouseUp(event IEvent, cellFrame foundation.Rect, controlView IView, flag bool) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("trackMouse:inRect:ofView:untilMouseUp:"), event, cellFrame, controlView, flag)
	return rv
}

func (c_ Cell) StartTrackingAt_InView(startPoint foundation.Point, controlView IView) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("startTrackingAt:inView:"), startPoint, controlView)
	return rv
}

func (c_ Cell) ContinueTracking_At_InView(lastPoint foundation.Point, currentPoint foundation.Point, controlView IView) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("continueTracking:at:inView:"), lastPoint, currentPoint, controlView)
	return rv
}

func (c_ Cell) StopTracking_At_InView_MouseIsUp(lastPoint foundation.Point, stopPoint foundation.Point, controlView IView, flag bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("stopTracking:at:inView:mouseIsUp:"), lastPoint, stopPoint, controlView, flag)
}

func (c_ Cell) GetPeriodicDelay_Interval(delay *float32, interval *float32) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("getPeriodicDelay:interval:"), delay, interval)
}

func (c_ Cell) HitTestForEvent_InRect_OfView(event IEvent, cellFrame foundation.Rect, controlView IView) CellHitResult {
	rv := objc.CallMethod[CellHitResult](c_, objc.GetSelector("hitTestForEvent:inRect:ofView:"), event, cellFrame, controlView)
	return rv
}

func (c_ Cell) ResetCursorRect_InView(cellFrame foundation.Rect, controlView IView) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("resetCursorRect:inView:"), cellFrame, controlView)
}

func (c_ Cell) DraggingImageComponentsWithFrame_InView(frame foundation.Rect, view IView) []DraggingImageComponent {
	rv := objc.CallMethod[[]DraggingImageComponent](c_, objc.GetSelector("draggingImageComponentsWithFrame:inView:"), frame, view)
	return rv
}

func (c_ Cell) DrawFocusRingMaskWithFrame_InView(cellFrame foundation.Rect, controlView IView) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("drawFocusRingMaskWithFrame:inView:"), cellFrame, controlView)
}

func (c_ Cell) FocusRingMaskBoundsForFrame_InView(cellFrame foundation.Rect, controlView IView) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.GetSelector("focusRingMaskBoundsForFrame:inView:"), cellFrame, controlView)
	return rv
}

func (c_ Cell) CalcDrawInfo(rect foundation.Rect) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("calcDrawInfo:"), rect)
}

func (c_ Cell) CellSizeForBounds(rect foundation.Rect) foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("cellSizeForBounds:"), rect)
	return rv
}

func (c_ Cell) DrawingRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.GetSelector("drawingRectForBounds:"), rect)
	return rv
}

func (c_ Cell) ImageRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.GetSelector("imageRectForBounds:"), rect)
	return rv
}

func (c_ Cell) TitleRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.GetSelector("titleRectForBounds:"), rect)
	return rv
}

func (c_ Cell) DrawWithFrame_InView(cellFrame foundation.Rect, controlView IView) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("drawWithFrame:inView:"), cellFrame, controlView)
}

func (c_ Cell) HighlightColorWithFrame_InView(cellFrame foundation.Rect, controlView IView) Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("highlightColorWithFrame:inView:"), cellFrame, controlView)
	return rv
}

func (c_ Cell) DrawInteriorWithFrame_InView(cellFrame foundation.Rect, controlView IView) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("drawInteriorWithFrame:inView:"), cellFrame, controlView)
}

func (c_ Cell) Highlight_WithFrame_InView(flag bool, cellFrame foundation.Rect, controlView IView) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("highlight:withFrame:inView:"), flag, cellFrame, controlView)
}

func (c_ Cell) EditWithFrame_InView_Editor_Delegate_Event(rect foundation.Rect, controlView IView, textObj IText, delegate objc.IObject, event IEvent) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("editWithFrame:inView:editor:delegate:event:"), rect, controlView, textObj, delegate, event)
}

func (c_ Cell) SelectWithFrame_InView_Editor_Delegate_Start_Length(rect foundation.Rect, controlView IView, textObj IText, delegate objc.IObject, selStart int, selLength int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("selectWithFrame:inView:editor:delegate:start:length:"), rect, controlView, textObj, delegate, selStart, selLength)
}

func (c_ Cell) EndEditing(textObj IText) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("endEditing:"), textObj)
}

func (c_ Cell) FieldEditorForView(controlView IView) TextView {
	rv := objc.CallMethod[TextView](c_, objc.GetSelector("fieldEditorForView:"), controlView)
	return rv
}

func (c_ Cell) ExpansionFrameWithFrame_InView(cellFrame foundation.Rect, view IView) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.GetSelector("expansionFrameWithFrame:inView:"), cellFrame, view)
	return rv
}

func (c_ Cell) DrawWithExpansionFrame_InView(cellFrame foundation.Rect, view IView) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("drawWithExpansionFrame:inView:"), cellFrame, view)
}

// deprecated
func (c_ Cell) SetEntryType(type_ int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setEntryType:"), type_)
}

// deprecated
func (c_ Cell) EntryType() int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("entryType"))
	return rv
}

// deprecated
func (c_ Cell) IsEntryAcceptable(string_ string) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isEntryAcceptable:"), string_)
	return rv
}

// deprecated
func (c_ Cell) SetFloatingPointFormat_Left_Right(autoRange bool, leftDigits uint, rightDigits uint) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setFloatingPointFormat:left:right:"), autoRange, leftDigits, rightDigits)
}

// deprecated
func (c_ Cell) SetTitleWithMnemonic(stringWithAmpersand string) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setTitleWithMnemonic:"), stringWithAmpersand)
}

// deprecated
func (c_ Cell) Mnemonic() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("mnemonic"))
	return rv
}

// deprecated
func (c_ Cell) SetMnemonicLocation(location uint) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setMnemonicLocation:"), location)
}

// deprecated
func (c_ Cell) MnemonicLocation() uint {
	rv := objc.CallMethod[uint](c_, objc.GetSelector("mnemonicLocation"))
	return rv
}

func (c_ Cell) ObjectValue() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.GetSelector("objectValue"))
	return rv
}

func (c_ Cell) SetObjectValue(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setObjectValue:"), value)
}

func (c_ Cell) HasValidObjectValue() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("hasValidObjectValue"))
	return rv
}

func (c_ Cell) IntValue() int32 {
	rv := objc.CallMethod[int32](c_, objc.GetSelector("intValue"))
	return rv
}

func (c_ Cell) SetIntValue(value int32) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setIntValue:"), value)
}

func (c_ Cell) IntegerValue() int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("integerValue"))
	return rv
}

func (c_ Cell) SetIntegerValue(value int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setIntegerValue:"), value)
}

func (c_ Cell) StringValue() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("stringValue"))
	return rv
}

func (c_ Cell) SetStringValue(value string) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setStringValue:"), value)
}

func (c_ Cell) DoubleValue() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("doubleValue"))
	return rv
}

func (c_ Cell) SetDoubleValue(value float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setDoubleValue:"), value)
}

func (c_ Cell) FloatValue() float32 {
	rv := objc.CallMethod[float32](c_, objc.GetSelector("floatValue"))
	return rv
}

func (c_ Cell) SetFloatValue(value float32) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setFloatValue:"), value)
}

func (c_ Cell) Type() CellType {
	rv := objc.CallMethod[CellType](c_, objc.GetSelector("type"))
	return rv
}

func (c_ Cell) SetType(value CellType) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setType:"), value)
}

func (c_ Cell) IsEnabled() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isEnabled"))
	return rv
}

func (c_ Cell) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setEnabled:"), value)
}

func (c_ Cell) AllowsUndo() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("allowsUndo"))
	return rv
}

func (c_ Cell) SetAllowsUndo(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setAllowsUndo:"), value)
}

func (c_ Cell) IsBezeled() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isBezeled"))
	return rv
}

func (c_ Cell) SetBezeled(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setBezeled:"), value)
}

func (c_ Cell) IsBordered() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isBordered"))
	return rv
}

func (c_ Cell) SetBordered(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setBordered:"), value)
}

func (c_ Cell) IsOpaque() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isOpaque"))
	return rv
}

// deprecated
func (c_ Cell) ControlTint() ControlTint {
	rv := objc.CallMethod[ControlTint](c_, objc.GetSelector("controlTint"))
	return rv
}

// deprecated
func (c_ Cell) SetControlTint(value ControlTint) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setControlTint:"), value)
}

func (c_ Cell) BackgroundStyle() BackgroundStyle {
	rv := objc.CallMethod[BackgroundStyle](c_, objc.GetSelector("backgroundStyle"))
	return rv
}

func (c_ Cell) SetBackgroundStyle(value BackgroundStyle) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setBackgroundStyle:"), value)
}

func (c_ Cell) InteriorBackgroundStyle() BackgroundStyle {
	rv := objc.CallMethod[BackgroundStyle](c_, objc.GetSelector("interiorBackgroundStyle"))
	return rv
}

func (c_ Cell) AllowsMixedState() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("allowsMixedState"))
	return rv
}

func (c_ Cell) SetAllowsMixedState(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setAllowsMixedState:"), value)
}

func (c_ Cell) NextState() int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("nextState"))
	return rv
}

func (c_ Cell) State() ControlStateValue {
	rv := objc.CallMethod[ControlStateValue](c_, objc.GetSelector("state"))
	return rv
}

func (c_ Cell) SetState(value ControlStateValue) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setState:"), value)
}

func (c_ Cell) IsEditable() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isEditable"))
	return rv
}

func (c_ Cell) SetEditable(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setEditable:"), value)
}

func (c_ Cell) IsSelectable() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isSelectable"))
	return rv
}

func (c_ Cell) SetSelectable(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setSelectable:"), value)
}

func (c_ Cell) IsScrollable() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isScrollable"))
	return rv
}

func (c_ Cell) SetScrollable(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setScrollable:"), value)
}

func (c_ Cell) Alignment() TextAlignment {
	rv := objc.CallMethod[TextAlignment](c_, objc.GetSelector("alignment"))
	return rv
}

func (c_ Cell) SetAlignment(value TextAlignment) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setAlignment:"), value)
}

func (c_ Cell) Font() Font {
	rv := objc.CallMethod[Font](c_, objc.GetSelector("font"))
	return rv
}

func (c_ Cell) SetFont(value IFont) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setFont:"), value)
}

func (c_ Cell) LineBreakMode() LineBreakMode {
	rv := objc.CallMethod[LineBreakMode](c_, objc.GetSelector("lineBreakMode"))
	return rv
}

func (c_ Cell) SetLineBreakMode(value LineBreakMode) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setLineBreakMode:"), value)
}

func (c_ Cell) TruncatesLastVisibleLine() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("truncatesLastVisibleLine"))
	return rv
}

func (c_ Cell) SetTruncatesLastVisibleLine(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setTruncatesLastVisibleLine:"), value)
}

func (c_ Cell) Wraps() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("wraps"))
	return rv
}

func (c_ Cell) SetWraps(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setWraps:"), value)
}

func (c_ Cell) BaseWritingDirection() WritingDirection {
	rv := objc.CallMethod[WritingDirection](c_, objc.GetSelector("baseWritingDirection"))
	return rv
}

func (c_ Cell) SetBaseWritingDirection(value WritingDirection) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setBaseWritingDirection:"), value)
}

func (c_ Cell) AttributedStringValue() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](c_, objc.GetSelector("attributedStringValue"))
	return rv
}

func (c_ Cell) SetAttributedStringValue(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setAttributedStringValue:"), value)
}

func (c_ Cell) AllowsEditingTextAttributes() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("allowsEditingTextAttributes"))
	return rv
}

func (c_ Cell) SetAllowsEditingTextAttributes(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setAllowsEditingTextAttributes:"), value)
}

func (c_ Cell) ImportsGraphics() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("importsGraphics"))
	return rv
}

func (c_ Cell) SetImportsGraphics(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setImportsGraphics:"), value)
}

func (c_ Cell) Title() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("title"))
	return rv
}

func (c_ Cell) SetTitle(value string) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setTitle:"), value)
}

func (c_ Cell) Action() objc.Selector {
	rv := objc.CallMethod[objc.Selector](c_, objc.GetSelector("action"))
	return rv
}

func (c_ Cell) SetAction(value objc.Selector) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setAction:"), value)
}

func (c_ Cell) Target() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.GetSelector("target"))
	return rv
}

func (c_ Cell) SetTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setTarget:"), value)
}

func (c_ Cell) IsContinuous() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isContinuous"))
	return rv
}

func (c_ Cell) SetContinuous(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setContinuous:"), value)
}

func (c_ Cell) Image() Image {
	rv := objc.CallMethod[Image](c_, objc.GetSelector("image"))
	return rv
}

func (c_ Cell) SetImage(value IImage) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setImage:"), value)
}

func (c_ Cell) Tag() int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("tag"))
	return rv
}

func (c_ Cell) SetTag(value int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setTag:"), value)
}

func (c_ Cell) Formatter() foundation.Formatter {
	rv := objc.CallMethod[foundation.Formatter](c_, objc.GetSelector("formatter"))
	return rv
}

func (c_ Cell) SetFormatter(value foundation.IFormatter) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setFormatter:"), value)
}

func (cc _CellClass) DefaultMenu() Menu {
	rv := objc.CallMethod[Menu](cc, objc.GetSelector("defaultMenu"))
	return rv
}

func (c_ Cell) Menu() Menu {
	rv := objc.CallMethod[Menu](c_, objc.GetSelector("menu"))
	return rv
}

func (c_ Cell) SetMenu(value IMenu) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setMenu:"), value)
}

func (c_ Cell) AcceptsFirstResponder() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("acceptsFirstResponder"))
	return rv
}

func (c_ Cell) ShowsFirstResponder() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("showsFirstResponder"))
	return rv
}

func (c_ Cell) SetShowsFirstResponder(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setShowsFirstResponder:"), value)
}

func (c_ Cell) RefusesFirstResponder() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("refusesFirstResponder"))
	return rv
}

func (c_ Cell) SetRefusesFirstResponder(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setRefusesFirstResponder:"), value)
}

func (c_ Cell) RepresentedObject() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.GetSelector("representedObject"))
	return rv
}

func (c_ Cell) SetRepresentedObject(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setRepresentedObject:"), value)
}

func (c_ Cell) MouseDownFlags() int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("mouseDownFlags"))
	return rv
}

func (cc _CellClass) PrefersTrackingUntilMouseUp() bool {
	rv := objc.CallMethod[bool](cc, objc.GetSelector("prefersTrackingUntilMouseUp"))
	return rv
}

func (c_ Cell) KeyEquivalent() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("keyEquivalent"))
	return rv
}

func (cc _CellClass) DefaultFocusRingType() FocusRingType {
	rv := objc.CallMethod[FocusRingType](cc, objc.GetSelector("defaultFocusRingType"))
	return rv
}

func (c_ Cell) FocusRingType() FocusRingType {
	rv := objc.CallMethod[FocusRingType](c_, objc.GetSelector("focusRingType"))
	return rv
}

func (c_ Cell) SetFocusRingType(value FocusRingType) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setFocusRingType:"), value)
}

func (c_ Cell) CellSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("cellSize"))
	return rv
}

func (c_ Cell) ControlSize() ControlSize {
	rv := objc.CallMethod[ControlSize](c_, objc.GetSelector("controlSize"))
	return rv
}

func (c_ Cell) SetControlSize(value ControlSize) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setControlSize:"), value)
}

func (c_ Cell) ControlView() View {
	rv := objc.CallMethod[View](c_, objc.GetSelector("controlView"))
	return rv
}

func (c_ Cell) SetControlView(value IView) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setControlView:"), value)
}

func (c_ Cell) IsHighlighted() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isHighlighted"))
	return rv
}

func (c_ Cell) SetHighlighted(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setHighlighted:"), value)
}

func (c_ Cell) SendsActionOnEndEditing() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("sendsActionOnEndEditing"))
	return rv
}

func (c_ Cell) SetSendsActionOnEndEditing(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setSendsActionOnEndEditing:"), value)
}

func (c_ Cell) WantsNotificationForMarkedText() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("wantsNotificationForMarkedText"))
	return rv
}

func (c_ Cell) UsesSingleLineMode() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("usesSingleLineMode"))
	return rv
}

func (c_ Cell) SetUsesSingleLineMode(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setUsesSingleLineMode:"), value)
}

func (c_ Cell) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.CallMethod[UserInterfaceLayoutDirection](c_, objc.GetSelector("userInterfaceLayoutDirection"))
	return rv
}

func (c_ Cell) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setUserInterfaceLayoutDirection:"), value)
}
