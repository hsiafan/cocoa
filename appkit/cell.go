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
	rv := objc.CallMethod[Cell](c_, objc.SEL("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func (c_ Cell) InitTextCell(string_ string) Cell {
	rv := objc.CallMethod[Cell](c_, objc.SEL("initTextCell:"), string_)
	return rv
}

func (c_ Cell) Init() Cell {
	rv := objc.CallMethod[Cell](c_, objc.SEL("init"))
	return rv
}

func (cc _CellClass) Alloc() Cell {
	rv := objc.CallMethod[Cell](cc, objc.SEL("alloc"))
	return rv
}

func (cc _CellClass) New() Cell {
	rv := objc.CallMethod[Cell](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewCell() Cell {
	return CellClass.New()
}

func (c_ Cell) SetCellAttribute_To(parameter CellAttribute, value int) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setCellAttribute:to:"), parameter, value)
}

func (c_ Cell) CellAttribute(parameter CellAttribute) int {
	rv := objc.CallMethod[int](c_, objc.SEL("cellAttribute:"), parameter)
	return rv
}

func (c_ Cell) SetNextState() {
	objc.CallMethod[objc.Void](c_, objc.SEL("setNextState"))
}

func (c_ Cell) SetUpFieldEditorAttributes(textObj IText) Text {
	rv := objc.CallMethod[Text](c_, objc.SEL("setUpFieldEditorAttributes:"), objc.ExtractPtr(textObj))
	return rv
}

func (c_ Cell) SendActionOn(mask EventMask) int {
	rv := objc.CallMethod[int](c_, objc.SEL("sendActionOn:"), mask)
	return rv
}

func (c_ Cell) MenuForEvent_InRect_OfView(event IEvent, cellFrame foundation.Rect, view IView) Menu {
	rv := objc.CallMethod[Menu](c_, objc.SEL("menuForEvent:inRect:ofView:"), objc.ExtractPtr(event), cellFrame, objc.ExtractPtr(view))
	return rv
}

func (c_ Cell) Compare(otherCell objc.IObject) foundation.ComparisonResult {
	rv := objc.CallMethod[foundation.ComparisonResult](c_, objc.SEL("compare:"), objc.ExtractPtr(otherCell))
	return rv
}

func (c_ Cell) PerformClick(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("performClick:"), objc.ExtractPtr(sender))
}

func (c_ Cell) TakeObjectValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("takeObjectValueFrom:"), objc.ExtractPtr(sender))
}

func (c_ Cell) TakeIntegerValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("takeIntegerValueFrom:"), objc.ExtractPtr(sender))
}

func (c_ Cell) TakeIntValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("takeIntValueFrom:"), objc.ExtractPtr(sender))
}

func (c_ Cell) TakeStringValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("takeStringValueFrom:"), objc.ExtractPtr(sender))
}

func (c_ Cell) TakeDoubleValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("takeDoubleValueFrom:"), objc.ExtractPtr(sender))
}

func (c_ Cell) TakeFloatValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("takeFloatValueFrom:"), objc.ExtractPtr(sender))
}

func (c_ Cell) TrackMouse_InRect_OfView_UntilMouseUp(event IEvent, cellFrame foundation.Rect, controlView IView, flag bool) bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("trackMouse:inRect:ofView:untilMouseUp:"), objc.ExtractPtr(event), cellFrame, objc.ExtractPtr(controlView), flag)
	return rv
}

func (c_ Cell) StartTrackingAt_InView(startPoint foundation.Point, controlView IView) bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("startTrackingAt:inView:"), startPoint, objc.ExtractPtr(controlView))
	return rv
}

func (c_ Cell) ContinueTracking_At_InView(lastPoint foundation.Point, currentPoint foundation.Point, controlView IView) bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("continueTracking:at:inView:"), lastPoint, currentPoint, objc.ExtractPtr(controlView))
	return rv
}

func (c_ Cell) StopTracking_At_InView_MouseIsUp(lastPoint foundation.Point, stopPoint foundation.Point, controlView IView, flag bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("stopTracking:at:inView:mouseIsUp:"), lastPoint, stopPoint, objc.ExtractPtr(controlView), flag)
}

func (c_ Cell) GetPeriodicDelay_Interval(delay *float32, interval *float32) {
	objc.CallMethod[objc.Void](c_, objc.SEL("getPeriodicDelay:interval:"), delay, interval)
}

func (c_ Cell) HitTestForEvent_InRect_OfView(event IEvent, cellFrame foundation.Rect, controlView IView) CellHitResult {
	rv := objc.CallMethod[CellHitResult](c_, objc.SEL("hitTestForEvent:inRect:ofView:"), objc.ExtractPtr(event), cellFrame, objc.ExtractPtr(controlView))
	return rv
}

func (c_ Cell) ResetCursorRect_InView(cellFrame foundation.Rect, controlView IView) {
	objc.CallMethod[objc.Void](c_, objc.SEL("resetCursorRect:inView:"), cellFrame, objc.ExtractPtr(controlView))
}

func (c_ Cell) DraggingImageComponentsWithFrame_InView(frame foundation.Rect, view IView) []DraggingImageComponent {
	rv := objc.CallMethod[[]DraggingImageComponent](c_, objc.SEL("draggingImageComponentsWithFrame:inView:"), frame, objc.ExtractPtr(view))
	return rv
}

func (c_ Cell) DrawFocusRingMaskWithFrame_InView(cellFrame foundation.Rect, controlView IView) {
	objc.CallMethod[objc.Void](c_, objc.SEL("drawFocusRingMaskWithFrame:inView:"), cellFrame, objc.ExtractPtr(controlView))
}

func (c_ Cell) FocusRingMaskBoundsForFrame_InView(cellFrame foundation.Rect, controlView IView) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.SEL("focusRingMaskBoundsForFrame:inView:"), cellFrame, objc.ExtractPtr(controlView))
	return rv
}

func (c_ Cell) CalcDrawInfo(rect foundation.Rect) {
	objc.CallMethod[objc.Void](c_, objc.SEL("calcDrawInfo:"), rect)
}

func (c_ Cell) CellSizeForBounds(rect foundation.Rect) foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.SEL("cellSizeForBounds:"), rect)
	return rv
}

func (c_ Cell) DrawingRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.SEL("drawingRectForBounds:"), rect)
	return rv
}

func (c_ Cell) ImageRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.SEL("imageRectForBounds:"), rect)
	return rv
}

func (c_ Cell) TitleRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.SEL("titleRectForBounds:"), rect)
	return rv
}

func (c_ Cell) DrawWithFrame_InView(cellFrame foundation.Rect, controlView IView) {
	objc.CallMethod[objc.Void](c_, objc.SEL("drawWithFrame:inView:"), cellFrame, objc.ExtractPtr(controlView))
}

func (c_ Cell) HighlightColorWithFrame_InView(cellFrame foundation.Rect, controlView IView) Color {
	rv := objc.CallMethod[Color](c_, objc.SEL("highlightColorWithFrame:inView:"), cellFrame, objc.ExtractPtr(controlView))
	return rv
}

func (c_ Cell) DrawInteriorWithFrame_InView(cellFrame foundation.Rect, controlView IView) {
	objc.CallMethod[objc.Void](c_, objc.SEL("drawInteriorWithFrame:inView:"), cellFrame, objc.ExtractPtr(controlView))
}

func (c_ Cell) Highlight_WithFrame_InView(flag bool, cellFrame foundation.Rect, controlView IView) {
	objc.CallMethod[objc.Void](c_, objc.SEL("highlight:withFrame:inView:"), flag, cellFrame, objc.ExtractPtr(controlView))
}

func (c_ Cell) EditWithFrame_InView_Editor_Delegate_Event(rect foundation.Rect, controlView IView, textObj IText, delegate objc.IObject, event IEvent) {
	objc.CallMethod[objc.Void](c_, objc.SEL("editWithFrame:inView:editor:delegate:event:"), rect, objc.ExtractPtr(controlView), objc.ExtractPtr(textObj), objc.ExtractPtr(delegate), objc.ExtractPtr(event))
}

func (c_ Cell) SelectWithFrame_InView_Editor_Delegate_Start_Length(rect foundation.Rect, controlView IView, textObj IText, delegate objc.IObject, selStart int, selLength int) {
	objc.CallMethod[objc.Void](c_, objc.SEL("selectWithFrame:inView:editor:delegate:start:length:"), rect, objc.ExtractPtr(controlView), objc.ExtractPtr(textObj), objc.ExtractPtr(delegate), selStart, selLength)
}

func (c_ Cell) EndEditing(textObj IText) {
	objc.CallMethod[objc.Void](c_, objc.SEL("endEditing:"), objc.ExtractPtr(textObj))
}

func (c_ Cell) FieldEditorForView(controlView IView) TextView {
	rv := objc.CallMethod[TextView](c_, objc.SEL("fieldEditorForView:"), objc.ExtractPtr(controlView))
	return rv
}

func (c_ Cell) ExpansionFrameWithFrame_InView(cellFrame foundation.Rect, view IView) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.SEL("expansionFrameWithFrame:inView:"), cellFrame, objc.ExtractPtr(view))
	return rv
}

func (c_ Cell) DrawWithExpansionFrame_InView(cellFrame foundation.Rect, view IView) {
	objc.CallMethod[objc.Void](c_, objc.SEL("drawWithExpansionFrame:inView:"), cellFrame, objc.ExtractPtr(view))
}

// deprecated
func (c_ Cell) SetEntryType(type_ int) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setEntryType:"), type_)
}

// deprecated
func (c_ Cell) EntryType() int {
	rv := objc.CallMethod[int](c_, objc.SEL("entryType"))
	return rv
}

// deprecated
func (c_ Cell) IsEntryAcceptable(string_ string) bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isEntryAcceptable:"), string_)
	return rv
}

// deprecated
func (c_ Cell) SetFloatingPointFormat_Left_Right(autoRange bool, leftDigits uint, rightDigits uint) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setFloatingPointFormat:left:right:"), autoRange, leftDigits, rightDigits)
}

// deprecated
func (c_ Cell) SetTitleWithMnemonic(stringWithAmpersand string) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setTitleWithMnemonic:"), stringWithAmpersand)
}

// deprecated
func (c_ Cell) Mnemonic() string {
	rv := objc.CallMethod[string](c_, objc.SEL("mnemonic"))
	return rv
}

// deprecated
func (c_ Cell) SetMnemonicLocation(location uint) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setMnemonicLocation:"), location)
}

// deprecated
func (c_ Cell) MnemonicLocation() uint {
	rv := objc.CallMethod[uint](c_, objc.SEL("mnemonicLocation"))
	return rv
}

func (c_ Cell) ObjectValue() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.SEL("objectValue"))
	return rv
}

func (c_ Cell) SetObjectValue(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setObjectValue:"), objc.ExtractPtr(value))
}

func (c_ Cell) HasValidObjectValue() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("hasValidObjectValue"))
	return rv
}

func (c_ Cell) IntValue() int32 {
	rv := objc.CallMethod[int32](c_, objc.SEL("intValue"))
	return rv
}

func (c_ Cell) SetIntValue(value int32) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setIntValue:"), value)
}

func (c_ Cell) IntegerValue() int {
	rv := objc.CallMethod[int](c_, objc.SEL("integerValue"))
	return rv
}

func (c_ Cell) SetIntegerValue(value int) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setIntegerValue:"), value)
}

func (c_ Cell) StringValue() string {
	rv := objc.CallMethod[string](c_, objc.SEL("stringValue"))
	return rv
}

func (c_ Cell) SetStringValue(value string) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setStringValue:"), value)
}

func (c_ Cell) DoubleValue() float64 {
	rv := objc.CallMethod[float64](c_, objc.SEL("doubleValue"))
	return rv
}

func (c_ Cell) SetDoubleValue(value float64) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setDoubleValue:"), value)
}

func (c_ Cell) FloatValue() float32 {
	rv := objc.CallMethod[float32](c_, objc.SEL("floatValue"))
	return rv
}

func (c_ Cell) SetFloatValue(value float32) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setFloatValue:"), value)
}

func (c_ Cell) Type() CellType {
	rv := objc.CallMethod[CellType](c_, objc.SEL("type"))
	return rv
}

func (c_ Cell) SetType(value CellType) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setType:"), value)
}

func (c_ Cell) IsEnabled() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isEnabled"))
	return rv
}

func (c_ Cell) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setEnabled:"), value)
}

func (c_ Cell) AllowsUndo() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("allowsUndo"))
	return rv
}

func (c_ Cell) SetAllowsUndo(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setAllowsUndo:"), value)
}

func (c_ Cell) IsBezeled() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isBezeled"))
	return rv
}

func (c_ Cell) SetBezeled(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setBezeled:"), value)
}

func (c_ Cell) IsBordered() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isBordered"))
	return rv
}

func (c_ Cell) SetBordered(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setBordered:"), value)
}

func (c_ Cell) IsOpaque() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isOpaque"))
	return rv
}

// deprecated
func (c_ Cell) ControlTint() ControlTint {
	rv := objc.CallMethod[ControlTint](c_, objc.SEL("controlTint"))
	return rv
}

// deprecated
func (c_ Cell) SetControlTint(value ControlTint) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setControlTint:"), value)
}

func (c_ Cell) BackgroundStyle() BackgroundStyle {
	rv := objc.CallMethod[BackgroundStyle](c_, objc.SEL("backgroundStyle"))
	return rv
}

func (c_ Cell) SetBackgroundStyle(value BackgroundStyle) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setBackgroundStyle:"), value)
}

func (c_ Cell) InteriorBackgroundStyle() BackgroundStyle {
	rv := objc.CallMethod[BackgroundStyle](c_, objc.SEL("interiorBackgroundStyle"))
	return rv
}

func (c_ Cell) AllowsMixedState() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("allowsMixedState"))
	return rv
}

func (c_ Cell) SetAllowsMixedState(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setAllowsMixedState:"), value)
}

func (c_ Cell) NextState() int {
	rv := objc.CallMethod[int](c_, objc.SEL("nextState"))
	return rv
}

func (c_ Cell) State() ControlStateValue {
	rv := objc.CallMethod[ControlStateValue](c_, objc.SEL("state"))
	return rv
}

func (c_ Cell) SetState(value ControlStateValue) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setState:"), value)
}

func (c_ Cell) IsEditable() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isEditable"))
	return rv
}

func (c_ Cell) SetEditable(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setEditable:"), value)
}

func (c_ Cell) IsSelectable() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isSelectable"))
	return rv
}

func (c_ Cell) SetSelectable(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setSelectable:"), value)
}

func (c_ Cell) IsScrollable() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isScrollable"))
	return rv
}

func (c_ Cell) SetScrollable(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setScrollable:"), value)
}

func (c_ Cell) Alignment() TextAlignment {
	rv := objc.CallMethod[TextAlignment](c_, objc.SEL("alignment"))
	return rv
}

func (c_ Cell) SetAlignment(value TextAlignment) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setAlignment:"), value)
}

func (c_ Cell) Font() Font {
	rv := objc.CallMethod[Font](c_, objc.SEL("font"))
	return rv
}

func (c_ Cell) SetFont(value IFont) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setFont:"), objc.ExtractPtr(value))
}

func (c_ Cell) LineBreakMode() LineBreakMode {
	rv := objc.CallMethod[LineBreakMode](c_, objc.SEL("lineBreakMode"))
	return rv
}

func (c_ Cell) SetLineBreakMode(value LineBreakMode) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setLineBreakMode:"), value)
}

func (c_ Cell) TruncatesLastVisibleLine() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("truncatesLastVisibleLine"))
	return rv
}

func (c_ Cell) SetTruncatesLastVisibleLine(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setTruncatesLastVisibleLine:"), value)
}

func (c_ Cell) Wraps() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("wraps"))
	return rv
}

func (c_ Cell) SetWraps(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setWraps:"), value)
}

func (c_ Cell) BaseWritingDirection() WritingDirection {
	rv := objc.CallMethod[WritingDirection](c_, objc.SEL("baseWritingDirection"))
	return rv
}

func (c_ Cell) SetBaseWritingDirection(value WritingDirection) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setBaseWritingDirection:"), value)
}

func (c_ Cell) AttributedStringValue() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](c_, objc.SEL("attributedStringValue"))
	return rv
}

func (c_ Cell) SetAttributedStringValue(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setAttributedStringValue:"), objc.ExtractPtr(value))
}

func (c_ Cell) AllowsEditingTextAttributes() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("allowsEditingTextAttributes"))
	return rv
}

func (c_ Cell) SetAllowsEditingTextAttributes(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setAllowsEditingTextAttributes:"), value)
}

func (c_ Cell) ImportsGraphics() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("importsGraphics"))
	return rv
}

func (c_ Cell) SetImportsGraphics(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setImportsGraphics:"), value)
}

func (c_ Cell) Title() string {
	rv := objc.CallMethod[string](c_, objc.SEL("title"))
	return rv
}

func (c_ Cell) SetTitle(value string) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setTitle:"), value)
}

func (c_ Cell) Action() objc.Selector {
	rv := objc.CallMethod[objc.Selector](c_, objc.SEL("action"))
	return rv
}

func (c_ Cell) SetAction(value objc.Selector) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setAction:"), value)
}

// weak property
func (c_ Cell) Target() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.SEL("target"))
	return rv
}

// weak property
func (c_ Cell) SetTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setTarget:"), objc.ExtractPtr(value))
}

func (c_ Cell) IsContinuous() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isContinuous"))
	return rv
}

func (c_ Cell) SetContinuous(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setContinuous:"), value)
}

func (c_ Cell) Image() Image {
	rv := objc.CallMethod[Image](c_, objc.SEL("image"))
	return rv
}

func (c_ Cell) SetImage(value IImage) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setImage:"), objc.ExtractPtr(value))
}

func (c_ Cell) Tag() int {
	rv := objc.CallMethod[int](c_, objc.SEL("tag"))
	return rv
}

func (c_ Cell) SetTag(value int) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setTag:"), value)
}

func (c_ Cell) Formatter() foundation.Formatter {
	rv := objc.CallMethod[foundation.Formatter](c_, objc.SEL("formatter"))
	return rv
}

func (c_ Cell) SetFormatter(value foundation.IFormatter) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setFormatter:"), objc.ExtractPtr(value))
}

func (cc _CellClass) DefaultMenu() Menu {
	rv := objc.CallMethod[Menu](cc, objc.SEL("defaultMenu"))
	return rv
}

func (c_ Cell) Menu() Menu {
	rv := objc.CallMethod[Menu](c_, objc.SEL("menu"))
	return rv
}

func (c_ Cell) SetMenu(value IMenu) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setMenu:"), objc.ExtractPtr(value))
}

func (c_ Cell) AcceptsFirstResponder() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("acceptsFirstResponder"))
	return rv
}

func (c_ Cell) ShowsFirstResponder() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("showsFirstResponder"))
	return rv
}

func (c_ Cell) SetShowsFirstResponder(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setShowsFirstResponder:"), value)
}

func (c_ Cell) RefusesFirstResponder() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("refusesFirstResponder"))
	return rv
}

func (c_ Cell) SetRefusesFirstResponder(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setRefusesFirstResponder:"), value)
}

func (c_ Cell) RepresentedObject() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.SEL("representedObject"))
	return rv
}

func (c_ Cell) SetRepresentedObject(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setRepresentedObject:"), objc.ExtractPtr(value))
}

func (c_ Cell) MouseDownFlags() int {
	rv := objc.CallMethod[int](c_, objc.SEL("mouseDownFlags"))
	return rv
}

func (cc _CellClass) PrefersTrackingUntilMouseUp() bool {
	rv := objc.CallMethod[bool](cc, objc.SEL("prefersTrackingUntilMouseUp"))
	return rv
}

func (c_ Cell) KeyEquivalent() string {
	rv := objc.CallMethod[string](c_, objc.SEL("keyEquivalent"))
	return rv
}

func (cc _CellClass) DefaultFocusRingType() FocusRingType {
	rv := objc.CallMethod[FocusRingType](cc, objc.SEL("defaultFocusRingType"))
	return rv
}

func (c_ Cell) FocusRingType() FocusRingType {
	rv := objc.CallMethod[FocusRingType](c_, objc.SEL("focusRingType"))
	return rv
}

func (c_ Cell) SetFocusRingType(value FocusRingType) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setFocusRingType:"), value)
}

func (c_ Cell) CellSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.SEL("cellSize"))
	return rv
}

func (c_ Cell) ControlSize() ControlSize {
	rv := objc.CallMethod[ControlSize](c_, objc.SEL("controlSize"))
	return rv
}

func (c_ Cell) SetControlSize(value ControlSize) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setControlSize:"), value)
}

// weak property
func (c_ Cell) ControlView() View {
	rv := objc.CallMethod[View](c_, objc.SEL("controlView"))
	return rv
}

// weak property
func (c_ Cell) SetControlView(value IView) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setControlView:"), objc.ExtractPtr(value))
}

func (c_ Cell) IsHighlighted() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isHighlighted"))
	return rv
}

func (c_ Cell) SetHighlighted(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setHighlighted:"), value)
}

func (c_ Cell) SendsActionOnEndEditing() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("sendsActionOnEndEditing"))
	return rv
}

func (c_ Cell) SetSendsActionOnEndEditing(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setSendsActionOnEndEditing:"), value)
}

func (c_ Cell) WantsNotificationForMarkedText() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("wantsNotificationForMarkedText"))
	return rv
}

func (c_ Cell) UsesSingleLineMode() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("usesSingleLineMode"))
	return rv
}

func (c_ Cell) SetUsesSingleLineMode(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setUsesSingleLineMode:"), value)
}

func (c_ Cell) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.CallMethod[UserInterfaceLayoutDirection](c_, objc.SEL("userInterfaceLayoutDirection"))
	return rv
}

func (c_ Cell) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setUserInterfaceLayoutDirection:"), value)
}
