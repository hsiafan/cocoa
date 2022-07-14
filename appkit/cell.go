// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	SetEntryType(_type int)
	// deprecated
	EntryType() int
	// deprecated
	IsEntryAcceptable(_string string) bool
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
	rv := ffi.CallMethod[Cell](c_, "initImageCell:", image)
	rv.Autorelease()
	return rv
}

func (c_ Cell) InitTextCell(_string string) Cell {
	rv := ffi.CallMethod[Cell](c_, "initTextCell:", _string)
	rv.Autorelease()
	return rv
}

func (c_ Cell) Init() Cell {
	rv := ffi.CallMethod[Cell](c_, "init")
	rv.Autorelease()
	return rv
}

func (cc _CellClass) Alloc() Cell {
	rv := ffi.CallMethod[Cell](cc, "alloc")
	return rv
}

func (cc _CellClass) New() Cell {
	rv := ffi.CallMethod[Cell](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCell() Cell {
	return CellClass.New()
}

func (c_ Cell) SetCellAttribute_To(parameter CellAttribute, value int) {
	ffi.CallMethod[ffi.Void](c_, "setCellAttribute:to:", parameter, value)
}

func (c_ Cell) CellAttribute(parameter CellAttribute) int {
	rv := ffi.CallMethod[int](c_, "cellAttribute:", parameter)
	return rv
}

func (c_ Cell) SetNextState() {
	ffi.CallMethod[ffi.Void](c_, "setNextState")
}

func (c_ Cell) SetUpFieldEditorAttributes(textObj IText) Text {
	rv := ffi.CallMethod[Text](c_, "setUpFieldEditorAttributes:", textObj)
	return rv
}

func (c_ Cell) SendActionOn(mask EventMask) int {
	rv := ffi.CallMethod[int](c_, "sendActionOn:", mask)
	return rv
}

func (c_ Cell) MenuForEvent_InRect_OfView(event IEvent, cellFrame foundation.Rect, view IView) Menu {
	rv := ffi.CallMethod[Menu](c_, "menuForEvent:inRect:ofView:", event, cellFrame, view)
	return rv
}

func (c_ Cell) Compare(otherCell objc.IObject) foundation.ComparisonResult {
	rv := ffi.CallMethod[foundation.ComparisonResult](c_, "compare:", otherCell)
	return rv
}

func (c_ Cell) PerformClick(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "performClick:", sender)
}

func (c_ Cell) TakeObjectValueFrom(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "takeObjectValueFrom:", sender)
}

func (c_ Cell) TakeIntegerValueFrom(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "takeIntegerValueFrom:", sender)
}

func (c_ Cell) TakeIntValueFrom(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "takeIntValueFrom:", sender)
}

func (c_ Cell) TakeStringValueFrom(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "takeStringValueFrom:", sender)
}

func (c_ Cell) TakeDoubleValueFrom(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "takeDoubleValueFrom:", sender)
}

func (c_ Cell) TakeFloatValueFrom(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "takeFloatValueFrom:", sender)
}

func (c_ Cell) TrackMouse_InRect_OfView_UntilMouseUp(event IEvent, cellFrame foundation.Rect, controlView IView, flag bool) bool {
	rv := ffi.CallMethod[bool](c_, "trackMouse:inRect:ofView:untilMouseUp:", event, cellFrame, controlView, flag)
	return rv
}

func (c_ Cell) StartTrackingAt_InView(startPoint foundation.Point, controlView IView) bool {
	rv := ffi.CallMethod[bool](c_, "startTrackingAt:inView:", startPoint, controlView)
	return rv
}

func (c_ Cell) ContinueTracking_At_InView(lastPoint foundation.Point, currentPoint foundation.Point, controlView IView) bool {
	rv := ffi.CallMethod[bool](c_, "continueTracking:at:inView:", lastPoint, currentPoint, controlView)
	return rv
}

func (c_ Cell) StopTracking_At_InView_MouseIsUp(lastPoint foundation.Point, stopPoint foundation.Point, controlView IView, flag bool) {
	ffi.CallMethod[ffi.Void](c_, "stopTracking:at:inView:mouseIsUp:", lastPoint, stopPoint, controlView, flag)
}

func (c_ Cell) GetPeriodicDelay_Interval(delay *float32, interval *float32) {
	ffi.CallMethod[ffi.Void](c_, "getPeriodicDelay:interval:", delay, interval)
}

func (c_ Cell) HitTestForEvent_InRect_OfView(event IEvent, cellFrame foundation.Rect, controlView IView) CellHitResult {
	rv := ffi.CallMethod[CellHitResult](c_, "hitTestForEvent:inRect:ofView:", event, cellFrame, controlView)
	return rv
}

func (c_ Cell) ResetCursorRect_InView(cellFrame foundation.Rect, controlView IView) {
	ffi.CallMethod[ffi.Void](c_, "resetCursorRect:inView:", cellFrame, controlView)
}

func (c_ Cell) DraggingImageComponentsWithFrame_InView(frame foundation.Rect, view IView) []DraggingImageComponent {
	rv := ffi.CallMethod[[]DraggingImageComponent](c_, "draggingImageComponentsWithFrame:inView:", frame, view)
	return rv
}

func (c_ Cell) DrawFocusRingMaskWithFrame_InView(cellFrame foundation.Rect, controlView IView) {
	ffi.CallMethod[ffi.Void](c_, "drawFocusRingMaskWithFrame:inView:", cellFrame, controlView)
}

func (c_ Cell) FocusRingMaskBoundsForFrame_InView(cellFrame foundation.Rect, controlView IView) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](c_, "focusRingMaskBoundsForFrame:inView:", cellFrame, controlView)
	return rv
}

func (c_ Cell) CalcDrawInfo(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](c_, "calcDrawInfo:", rect)
}

func (c_ Cell) CellSizeForBounds(rect foundation.Rect) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "cellSizeForBounds:", rect)
	return rv
}

func (c_ Cell) DrawingRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](c_, "drawingRectForBounds:", rect)
	return rv
}

func (c_ Cell) ImageRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](c_, "imageRectForBounds:", rect)
	return rv
}

func (c_ Cell) TitleRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](c_, "titleRectForBounds:", rect)
	return rv
}

func (c_ Cell) DrawWithFrame_InView(cellFrame foundation.Rect, controlView IView) {
	ffi.CallMethod[ffi.Void](c_, "drawWithFrame:inView:", cellFrame, controlView)
}

func (c_ Cell) HighlightColorWithFrame_InView(cellFrame foundation.Rect, controlView IView) Color {
	rv := ffi.CallMethod[Color](c_, "highlightColorWithFrame:inView:", cellFrame, controlView)
	return rv
}

func (c_ Cell) DrawInteriorWithFrame_InView(cellFrame foundation.Rect, controlView IView) {
	ffi.CallMethod[ffi.Void](c_, "drawInteriorWithFrame:inView:", cellFrame, controlView)
}

func (c_ Cell) Highlight_WithFrame_InView(flag bool, cellFrame foundation.Rect, controlView IView) {
	ffi.CallMethod[ffi.Void](c_, "highlight:withFrame:inView:", flag, cellFrame, controlView)
}

func (c_ Cell) EditWithFrame_InView_Editor_Delegate_Event(rect foundation.Rect, controlView IView, textObj IText, delegate objc.IObject, event IEvent) {
	ffi.CallMethod[ffi.Void](c_, "editWithFrame:inView:editor:delegate:event:", rect, controlView, textObj, delegate, event)
}

func (c_ Cell) SelectWithFrame_InView_Editor_Delegate_Start_Length(rect foundation.Rect, controlView IView, textObj IText, delegate objc.IObject, selStart int, selLength int) {
	ffi.CallMethod[ffi.Void](c_, "selectWithFrame:inView:editor:delegate:start:length:", rect, controlView, textObj, delegate, selStart, selLength)
}

func (c_ Cell) EndEditing(textObj IText) {
	ffi.CallMethod[ffi.Void](c_, "endEditing:", textObj)
}

func (c_ Cell) FieldEditorForView(controlView IView) TextView {
	rv := ffi.CallMethod[TextView](c_, "fieldEditorForView:", controlView)
	return rv
}

func (c_ Cell) ExpansionFrameWithFrame_InView(cellFrame foundation.Rect, view IView) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](c_, "expansionFrameWithFrame:inView:", cellFrame, view)
	return rv
}

func (c_ Cell) DrawWithExpansionFrame_InView(cellFrame foundation.Rect, view IView) {
	ffi.CallMethod[ffi.Void](c_, "drawWithExpansionFrame:inView:", cellFrame, view)
}

// deprecated
func (c_ Cell) SetEntryType(_type int) {
	ffi.CallMethod[ffi.Void](c_, "setEntryType:", _type)
}

// deprecated
func (c_ Cell) EntryType() int {
	rv := ffi.CallMethod[int](c_, "entryType")
	return rv
}

// deprecated
func (c_ Cell) IsEntryAcceptable(_string string) bool {
	rv := ffi.CallMethod[bool](c_, "isEntryAcceptable:", _string)
	return rv
}

// deprecated
func (c_ Cell) SetFloatingPointFormat_Left_Right(autoRange bool, leftDigits uint, rightDigits uint) {
	ffi.CallMethod[ffi.Void](c_, "setFloatingPointFormat:left:right:", autoRange, leftDigits, rightDigits)
}

// deprecated
func (c_ Cell) SetTitleWithMnemonic(stringWithAmpersand string) {
	ffi.CallMethod[ffi.Void](c_, "setTitleWithMnemonic:", stringWithAmpersand)
}

// deprecated
func (c_ Cell) Mnemonic() string {
	rv := ffi.CallMethod[string](c_, "mnemonic")
	return rv
}

// deprecated
func (c_ Cell) SetMnemonicLocation(location uint) {
	ffi.CallMethod[ffi.Void](c_, "setMnemonicLocation:", location)
}

// deprecated
func (c_ Cell) MnemonicLocation() uint {
	rv := ffi.CallMethod[uint](c_, "mnemonicLocation")
	return rv
}

func (c_ Cell) ObjectValue() objc.Object {
	rv := ffi.CallMethod[objc.Object](c_, "objectValue")
	return rv
}

func (c_ Cell) SetObjectValue(value objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "setObjectValue:", value)
}

func (c_ Cell) HasValidObjectValue() bool {
	rv := ffi.CallMethod[bool](c_, "hasValidObjectValue")
	return rv
}

func (c_ Cell) IntValue() int32 {
	rv := ffi.CallMethod[int32](c_, "intValue")
	return rv
}

func (c_ Cell) SetIntValue(value int32) {
	ffi.CallMethod[ffi.Void](c_, "setIntValue:", value)
}

func (c_ Cell) IntegerValue() int {
	rv := ffi.CallMethod[int](c_, "integerValue")
	return rv
}

func (c_ Cell) SetIntegerValue(value int) {
	ffi.CallMethod[ffi.Void](c_, "setIntegerValue:", value)
}

func (c_ Cell) StringValue() string {
	rv := ffi.CallMethod[string](c_, "stringValue")
	return rv
}

func (c_ Cell) SetStringValue(value string) {
	ffi.CallMethod[ffi.Void](c_, "setStringValue:", value)
}

func (c_ Cell) DoubleValue() float64 {
	rv := ffi.CallMethod[float64](c_, "doubleValue")
	return rv
}

func (c_ Cell) SetDoubleValue(value float64) {
	ffi.CallMethod[ffi.Void](c_, "setDoubleValue:", value)
}

func (c_ Cell) FloatValue() float32 {
	rv := ffi.CallMethod[float32](c_, "floatValue")
	return rv
}

func (c_ Cell) SetFloatValue(value float32) {
	ffi.CallMethod[ffi.Void](c_, "setFloatValue:", value)
}

func (c_ Cell) Type() CellType {
	rv := ffi.CallMethod[CellType](c_, "type")
	return rv
}

func (c_ Cell) SetType(value CellType) {
	ffi.CallMethod[ffi.Void](c_, "setType:", value)
}

func (c_ Cell) IsEnabled() bool {
	rv := ffi.CallMethod[bool](c_, "isEnabled")
	return rv
}

func (c_ Cell) SetEnabled(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setEnabled:", value)
}

func (c_ Cell) AllowsUndo() bool {
	rv := ffi.CallMethod[bool](c_, "allowsUndo")
	return rv
}

func (c_ Cell) SetAllowsUndo(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setAllowsUndo:", value)
}

func (c_ Cell) IsBezeled() bool {
	rv := ffi.CallMethod[bool](c_, "isBezeled")
	return rv
}

func (c_ Cell) SetBezeled(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setBezeled:", value)
}

func (c_ Cell) IsBordered() bool {
	rv := ffi.CallMethod[bool](c_, "isBordered")
	return rv
}

func (c_ Cell) SetBordered(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setBordered:", value)
}

func (c_ Cell) IsOpaque() bool {
	rv := ffi.CallMethod[bool](c_, "isOpaque")
	return rv
}

// deprecated
func (c_ Cell) ControlTint() ControlTint {
	rv := ffi.CallMethod[ControlTint](c_, "controlTint")
	return rv
}

// deprecated
func (c_ Cell) SetControlTint(value ControlTint) {
	ffi.CallMethod[ffi.Void](c_, "setControlTint:", value)
}

func (c_ Cell) BackgroundStyle() BackgroundStyle {
	rv := ffi.CallMethod[BackgroundStyle](c_, "backgroundStyle")
	return rv
}

func (c_ Cell) SetBackgroundStyle(value BackgroundStyle) {
	ffi.CallMethod[ffi.Void](c_, "setBackgroundStyle:", value)
}

func (c_ Cell) InteriorBackgroundStyle() BackgroundStyle {
	rv := ffi.CallMethod[BackgroundStyle](c_, "interiorBackgroundStyle")
	return rv
}

func (c_ Cell) AllowsMixedState() bool {
	rv := ffi.CallMethod[bool](c_, "allowsMixedState")
	return rv
}

func (c_ Cell) SetAllowsMixedState(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setAllowsMixedState:", value)
}

func (c_ Cell) NextState() int {
	rv := ffi.CallMethod[int](c_, "nextState")
	return rv
}

func (c_ Cell) State() ControlStateValue {
	rv := ffi.CallMethod[ControlStateValue](c_, "state")
	return rv
}

func (c_ Cell) SetState(value ControlStateValue) {
	ffi.CallMethod[ffi.Void](c_, "setState:", value)
}

func (c_ Cell) IsEditable() bool {
	rv := ffi.CallMethod[bool](c_, "isEditable")
	return rv
}

func (c_ Cell) SetEditable(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setEditable:", value)
}

func (c_ Cell) IsSelectable() bool {
	rv := ffi.CallMethod[bool](c_, "isSelectable")
	return rv
}

func (c_ Cell) SetSelectable(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setSelectable:", value)
}

func (c_ Cell) IsScrollable() bool {
	rv := ffi.CallMethod[bool](c_, "isScrollable")
	return rv
}

func (c_ Cell) SetScrollable(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setScrollable:", value)
}

func (c_ Cell) Alignment() TextAlignment {
	rv := ffi.CallMethod[TextAlignment](c_, "alignment")
	return rv
}

func (c_ Cell) SetAlignment(value TextAlignment) {
	ffi.CallMethod[ffi.Void](c_, "setAlignment:", value)
}

func (c_ Cell) Font() Font {
	rv := ffi.CallMethod[Font](c_, "font")
	return rv
}

func (c_ Cell) SetFont(value IFont) {
	ffi.CallMethod[ffi.Void](c_, "setFont:", value)
}

func (c_ Cell) LineBreakMode() LineBreakMode {
	rv := ffi.CallMethod[LineBreakMode](c_, "lineBreakMode")
	return rv
}

func (c_ Cell) SetLineBreakMode(value LineBreakMode) {
	ffi.CallMethod[ffi.Void](c_, "setLineBreakMode:", value)
}

func (c_ Cell) TruncatesLastVisibleLine() bool {
	rv := ffi.CallMethod[bool](c_, "truncatesLastVisibleLine")
	return rv
}

func (c_ Cell) SetTruncatesLastVisibleLine(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setTruncatesLastVisibleLine:", value)
}

func (c_ Cell) Wraps() bool {
	rv := ffi.CallMethod[bool](c_, "wraps")
	return rv
}

func (c_ Cell) SetWraps(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setWraps:", value)
}

func (c_ Cell) BaseWritingDirection() WritingDirection {
	rv := ffi.CallMethod[WritingDirection](c_, "baseWritingDirection")
	return rv
}

func (c_ Cell) SetBaseWritingDirection(value WritingDirection) {
	ffi.CallMethod[ffi.Void](c_, "setBaseWritingDirection:", value)
}

func (c_ Cell) AttributedStringValue() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](c_, "attributedStringValue")
	return rv
}

func (c_ Cell) SetAttributedStringValue(value foundation.IAttributedString) {
	ffi.CallMethod[ffi.Void](c_, "setAttributedStringValue:", value)
}

func (c_ Cell) AllowsEditingTextAttributes() bool {
	rv := ffi.CallMethod[bool](c_, "allowsEditingTextAttributes")
	return rv
}

func (c_ Cell) SetAllowsEditingTextAttributes(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setAllowsEditingTextAttributes:", value)
}

func (c_ Cell) ImportsGraphics() bool {
	rv := ffi.CallMethod[bool](c_, "importsGraphics")
	return rv
}

func (c_ Cell) SetImportsGraphics(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setImportsGraphics:", value)
}

func (c_ Cell) Title() string {
	rv := ffi.CallMethod[string](c_, "title")
	return rv
}

func (c_ Cell) SetTitle(value string) {
	ffi.CallMethod[ffi.Void](c_, "setTitle:", value)
}

func (c_ Cell) Action() objc.Selector {
	rv := ffi.CallMethod[objc.Selector](c_, "action")
	return rv
}

func (c_ Cell) SetAction(value objc.Selector) {
	ffi.CallMethod[ffi.Void](c_, "setAction:", value)
}

func (c_ Cell) Target() objc.Object {
	rv := ffi.CallMethod[objc.Object](c_, "target")
	return rv
}

func (c_ Cell) SetTarget(value objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "setTarget:", value)
}

func (c_ Cell) IsContinuous() bool {
	rv := ffi.CallMethod[bool](c_, "isContinuous")
	return rv
}

func (c_ Cell) SetContinuous(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setContinuous:", value)
}

func (c_ Cell) Image() Image {
	rv := ffi.CallMethod[Image](c_, "image")
	return rv
}

func (c_ Cell) SetImage(value IImage) {
	ffi.CallMethod[ffi.Void](c_, "setImage:", value)
}

func (c_ Cell) Tag() int {
	rv := ffi.CallMethod[int](c_, "tag")
	return rv
}

func (c_ Cell) SetTag(value int) {
	ffi.CallMethod[ffi.Void](c_, "setTag:", value)
}

func (c_ Cell) Formatter() foundation.Formatter {
	rv := ffi.CallMethod[foundation.Formatter](c_, "formatter")
	return rv
}

func (c_ Cell) SetFormatter(value foundation.IFormatter) {
	ffi.CallMethod[ffi.Void](c_, "setFormatter:", value)
}

func (cc _CellClass) DefaultMenu() Menu {
	rv := ffi.CallMethod[Menu](cc, "defaultMenu")
	return rv
}

func (c_ Cell) Menu() Menu {
	rv := ffi.CallMethod[Menu](c_, "menu")
	return rv
}

func (c_ Cell) SetMenu(value IMenu) {
	ffi.CallMethod[ffi.Void](c_, "setMenu:", value)
}

func (c_ Cell) AcceptsFirstResponder() bool {
	rv := ffi.CallMethod[bool](c_, "acceptsFirstResponder")
	return rv
}

func (c_ Cell) ShowsFirstResponder() bool {
	rv := ffi.CallMethod[bool](c_, "showsFirstResponder")
	return rv
}

func (c_ Cell) SetShowsFirstResponder(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setShowsFirstResponder:", value)
}

func (c_ Cell) RefusesFirstResponder() bool {
	rv := ffi.CallMethod[bool](c_, "refusesFirstResponder")
	return rv
}

func (c_ Cell) SetRefusesFirstResponder(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setRefusesFirstResponder:", value)
}

func (c_ Cell) RepresentedObject() objc.Object {
	rv := ffi.CallMethod[objc.Object](c_, "representedObject")
	return rv
}

func (c_ Cell) SetRepresentedObject(value objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "setRepresentedObject:", value)
}

func (c_ Cell) MouseDownFlags() int {
	rv := ffi.CallMethod[int](c_, "mouseDownFlags")
	return rv
}

func (cc _CellClass) PrefersTrackingUntilMouseUp() bool {
	rv := ffi.CallMethod[bool](cc, "prefersTrackingUntilMouseUp")
	return rv
}

func (c_ Cell) KeyEquivalent() string {
	rv := ffi.CallMethod[string](c_, "keyEquivalent")
	return rv
}

func (cc _CellClass) DefaultFocusRingType() FocusRingType {
	rv := ffi.CallMethod[FocusRingType](cc, "defaultFocusRingType")
	return rv
}

func (c_ Cell) FocusRingType() FocusRingType {
	rv := ffi.CallMethod[FocusRingType](c_, "focusRingType")
	return rv
}

func (c_ Cell) SetFocusRingType(value FocusRingType) {
	ffi.CallMethod[ffi.Void](c_, "setFocusRingType:", value)
}

func (c_ Cell) CellSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "cellSize")
	return rv
}

func (c_ Cell) ControlSize() ControlSize {
	rv := ffi.CallMethod[ControlSize](c_, "controlSize")
	return rv
}

func (c_ Cell) SetControlSize(value ControlSize) {
	ffi.CallMethod[ffi.Void](c_, "setControlSize:", value)
}

func (c_ Cell) ControlView() View {
	rv := ffi.CallMethod[View](c_, "controlView")
	return rv
}

func (c_ Cell) SetControlView(value IView) {
	ffi.CallMethod[ffi.Void](c_, "setControlView:", value)
}

func (c_ Cell) IsHighlighted() bool {
	rv := ffi.CallMethod[bool](c_, "isHighlighted")
	return rv
}

func (c_ Cell) SetHighlighted(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setHighlighted:", value)
}

func (c_ Cell) SendsActionOnEndEditing() bool {
	rv := ffi.CallMethod[bool](c_, "sendsActionOnEndEditing")
	return rv
}

func (c_ Cell) SetSendsActionOnEndEditing(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setSendsActionOnEndEditing:", value)
}

func (c_ Cell) WantsNotificationForMarkedText() bool {
	rv := ffi.CallMethod[bool](c_, "wantsNotificationForMarkedText")
	return rv
}

func (c_ Cell) UsesSingleLineMode() bool {
	rv := ffi.CallMethod[bool](c_, "usesSingleLineMode")
	return rv
}

func (c_ Cell) SetUsesSingleLineMode(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setUsesSingleLineMode:", value)
}

func (c_ Cell) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := ffi.CallMethod[UserInterfaceLayoutDirection](c_, "userInterfaceLayoutDirection")
	return rv
}

func (c_ Cell) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	ffi.CallMethod[ffi.Void](c_, "setUserInterfaceLayoutDirection:", value)
}

var ActionCellClass = _ActionCellClass{objc.GetClass("NSActionCell")}

type _ActionCellClass struct {
	objc.Class
}

type IActionCell interface {
	ICell
}

type ActionCell struct {
	Cell
}

func MakeActionCell(ptr unsafe.Pointer) ActionCell {
	return ActionCell{
		Cell: MakeCell(ptr),
	}
}

func (a_ ActionCell) InitImageCell(image IImage) ActionCell {
	rv := ffi.CallMethod[ActionCell](a_, "initImageCell:", image)
	rv.Autorelease()
	return rv
}

func (a_ ActionCell) InitTextCell(_string string) ActionCell {
	rv := ffi.CallMethod[ActionCell](a_, "initTextCell:", _string)
	rv.Autorelease()
	return rv
}

func (a_ ActionCell) Init() ActionCell {
	rv := ffi.CallMethod[ActionCell](a_, "init")
	rv.Autorelease()
	return rv
}

func (ac _ActionCellClass) Alloc() ActionCell {
	rv := ffi.CallMethod[ActionCell](ac, "alloc")
	return rv
}

func (ac _ActionCellClass) New() ActionCell {
	rv := ffi.CallMethod[ActionCell](ac, "new")
	rv.Autorelease()
	return rv
}

func NewActionCell() ActionCell {
	return ActionCellClass.New()
}
