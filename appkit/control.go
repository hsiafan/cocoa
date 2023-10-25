// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ControlClass = _ControlClass{objc.GetClass("NSControl")}

type _ControlClass struct {
	objc.Class
}

type IControl interface {
	IView
	TakeDoubleValueFrom(sender objc.IObject)
	TakeFloatValueFrom(sender objc.IObject)
	TakeIntValueFrom(sender objc.IObject)
	TakeIntegerValueFrom(sender objc.IObject)
	TakeObjectValueFrom(sender objc.IObject)
	TakeStringValueFrom(sender objc.IObject)
	// deprecated
	SetFloatingPointFormat_Left_Right(autoRange bool, leftDigits uint, rightDigits uint)
	DrawWithExpansionFrame_InView(contentFrame foundation.Rect, view IView)
	ExpansionFrameWithFrame(contentFrame foundation.Rect) foundation.Rect
	AbortEditing() bool
	CurrentEditor() Text
	ValidateEditing()
	EditWithFrame_Editor_Delegate_Event(rect foundation.Rect, textObj IText, delegate objc.IObject, event IEvent)
	EndEditing(textObj IText)
	SelectWithFrame_Editor_Delegate_Start_Length(rect foundation.Rect, textObj IText, delegate objc.IObject, selStart int, selLength int)
	SizeThatFits(size foundation.Size) foundation.Size
	SizeToFit()
	SendAction_To(action objc.Selector, target objc.IObject) bool
	SendActionOn(mask EventMask) int
	PerformClick(sender objc.IObject)
	InvalidateIntrinsicContentSizeForCell(cell ICell)
	// deprecated
	SelectedCell() Cell
	// deprecated
	SelectedTag() int
	// deprecated
	SetNeedsDisplay0()
	// deprecated
	CalcSize()
	// deprecated
	SelectCell(cell ICell)
	// deprecated
	DrawCell(cell ICell)
	// deprecated
	DrawCellInside(cell ICell)
	// deprecated
	UpdateCell(cell ICell)
	// deprecated
	UpdateCellInside(cell ICell)
	// deprecated
	ControlTextDidBeginEditing(obj foundation.INotification)
	// deprecated
	ControlTextDidChange(obj foundation.INotification)
	// deprecated
	ControlTextDidEndEditing(obj foundation.INotification)
	IsEnabled() bool
	SetEnabled(value bool)
	DoubleValue() float64
	SetDoubleValue(value float64)
	FloatValue() float32
	SetFloatValue(value float32)
	IntValue() int32
	SetIntValue(value int32)
	IntegerValue() int
	SetIntegerValue(value int)
	ObjectValue() objc.Object
	SetObjectValue(value objc.IObject)
	StringValue() string
	SetStringValue(value string)
	AttributedStringValue() foundation.AttributedString
	SetAttributedStringValue(value foundation.IAttributedString)
	Alignment() TextAlignment
	SetAlignment(value TextAlignment)
	Font() Font
	SetFont(value IFont)
	LineBreakMode() LineBreakMode
	SetLineBreakMode(value LineBreakMode)
	UsesSingleLineMode() bool
	SetUsesSingleLineMode(value bool)
	Formatter() foundation.Formatter
	SetFormatter(value foundation.IFormatter)
	BaseWritingDirection() WritingDirection
	SetBaseWritingDirection(value WritingDirection)
	AllowsExpansionToolTips() bool
	SetAllowsExpansionToolTips(value bool)
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	IsHighlighted() bool
	SetHighlighted(value bool)
	Action() objc.Selector
	SetAction(value objc.Selector)
	Target() objc.Object
	SetTarget(value objc.IObject)
	IsContinuous() bool
	SetContinuous(value bool)
	SetTag(value int)
	RefusesFirstResponder() bool
	SetRefusesFirstResponder(value bool)
	IgnoresMultiClick() bool
	SetIgnoresMultiClick(value bool)
	// deprecated
	Cell() Cell
	// deprecated
	SetCell(value ICell)
}

type Control struct {
	View
}

func MakeControl(ptr unsafe.Pointer) Control {
	return Control{
		View: MakeView(ptr),
	}
}

func (c_ Control) InitWithFrame(frameRect foundation.Rect) Control {
	rv := objc.CallMethod[Control](c_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (c_ Control) Init() Control {
	rv := objc.CallMethod[Control](c_, objc.SEL("init"))
	return rv
}

func (cc _ControlClass) Alloc() Control {
	rv := objc.CallMethod[Control](cc, objc.SEL("alloc"))
	return rv
}

func (cc _ControlClass) New() Control {
	rv := objc.CallMethod[Control](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewControl() Control {
	return ControlClass.New()
}

func (c_ Control) TakeDoubleValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("takeDoubleValueFrom:"), objc.ExtractPtr(sender))
}

func (c_ Control) TakeFloatValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("takeFloatValueFrom:"), objc.ExtractPtr(sender))
}

func (c_ Control) TakeIntValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("takeIntValueFrom:"), objc.ExtractPtr(sender))
}

func (c_ Control) TakeIntegerValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("takeIntegerValueFrom:"), objc.ExtractPtr(sender))
}

func (c_ Control) TakeObjectValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("takeObjectValueFrom:"), objc.ExtractPtr(sender))
}

func (c_ Control) TakeStringValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("takeStringValueFrom:"), objc.ExtractPtr(sender))
}

// deprecated
func (c_ Control) SetFloatingPointFormat_Left_Right(autoRange bool, leftDigits uint, rightDigits uint) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setFloatingPointFormat:left:right:"), autoRange, leftDigits, rightDigits)
}

func (c_ Control) DrawWithExpansionFrame_InView(contentFrame foundation.Rect, view IView) {
	objc.CallMethod[objc.Void](c_, objc.SEL("drawWithExpansionFrame:inView:"), contentFrame, objc.ExtractPtr(view))
}

func (c_ Control) ExpansionFrameWithFrame(contentFrame foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.SEL("expansionFrameWithFrame:"), contentFrame)
	return rv
}

func (c_ Control) AbortEditing() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("abortEditing"))
	return rv
}

func (c_ Control) CurrentEditor() Text {
	rv := objc.CallMethod[Text](c_, objc.SEL("currentEditor"))
	return rv
}

func (c_ Control) ValidateEditing() {
	objc.CallMethod[objc.Void](c_, objc.SEL("validateEditing"))
}

func (c_ Control) EditWithFrame_Editor_Delegate_Event(rect foundation.Rect, textObj IText, delegate objc.IObject, event IEvent) {
	objc.CallMethod[objc.Void](c_, objc.SEL("editWithFrame:editor:delegate:event:"), rect, objc.ExtractPtr(textObj), objc.ExtractPtr(delegate), objc.ExtractPtr(event))
}

func (c_ Control) EndEditing(textObj IText) {
	objc.CallMethod[objc.Void](c_, objc.SEL("endEditing:"), objc.ExtractPtr(textObj))
}

func (c_ Control) SelectWithFrame_Editor_Delegate_Start_Length(rect foundation.Rect, textObj IText, delegate objc.IObject, selStart int, selLength int) {
	objc.CallMethod[objc.Void](c_, objc.SEL("selectWithFrame:editor:delegate:start:length:"), rect, objc.ExtractPtr(textObj), objc.ExtractPtr(delegate), selStart, selLength)
}

func (c_ Control) SizeThatFits(size foundation.Size) foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.SEL("sizeThatFits:"), size)
	return rv
}

func (c_ Control) SizeToFit() {
	objc.CallMethod[objc.Void](c_, objc.SEL("sizeToFit"))
}

func (c_ Control) SendAction_To(action objc.Selector, target objc.IObject) bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("sendAction:to:"), action, objc.ExtractPtr(target))
	return rv
}

func (c_ Control) SendActionOn(mask EventMask) int {
	rv := objc.CallMethod[int](c_, objc.SEL("sendActionOn:"), mask)
	return rv
}

func (c_ Control) PerformClick(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("performClick:"), objc.ExtractPtr(sender))
}

func (c_ Control) InvalidateIntrinsicContentSizeForCell(cell ICell) {
	objc.CallMethod[objc.Void](c_, objc.SEL("invalidateIntrinsicContentSizeForCell:"), objc.ExtractPtr(cell))
}

// deprecated
func (c_ Control) SelectedCell() Cell {
	rv := objc.CallMethod[Cell](c_, objc.SEL("selectedCell"))
	return rv
}

// deprecated
func (c_ Control) SelectedTag() int {
	rv := objc.CallMethod[int](c_, objc.SEL("selectedTag"))
	return rv
}

// deprecated
func (c_ Control) SetNeedsDisplay0() {
	objc.CallMethod[objc.Void](c_, objc.SEL("setNeedsDisplay"))
}

// deprecated
func (c_ Control) CalcSize() {
	objc.CallMethod[objc.Void](c_, objc.SEL("calcSize"))
}

// deprecated
func (c_ Control) SelectCell(cell ICell) {
	objc.CallMethod[objc.Void](c_, objc.SEL("selectCell:"), objc.ExtractPtr(cell))
}

// deprecated
func (c_ Control) DrawCell(cell ICell) {
	objc.CallMethod[objc.Void](c_, objc.SEL("drawCell:"), objc.ExtractPtr(cell))
}

// deprecated
func (c_ Control) DrawCellInside(cell ICell) {
	objc.CallMethod[objc.Void](c_, objc.SEL("drawCellInside:"), objc.ExtractPtr(cell))
}

// deprecated
func (c_ Control) UpdateCell(cell ICell) {
	objc.CallMethod[objc.Void](c_, objc.SEL("updateCell:"), objc.ExtractPtr(cell))
}

// deprecated
func (c_ Control) UpdateCellInside(cell ICell) {
	objc.CallMethod[objc.Void](c_, objc.SEL("updateCellInside:"), objc.ExtractPtr(cell))
}

// deprecated
func (c_ Control) ControlTextDidBeginEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.SEL("controlTextDidBeginEditing:"), objc.ExtractPtr(obj))
}

// deprecated
func (c_ Control) ControlTextDidChange(obj foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.SEL("controlTextDidChange:"), objc.ExtractPtr(obj))
}

// deprecated
func (c_ Control) ControlTextDidEndEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](c_, objc.SEL("controlTextDidEndEditing:"), objc.ExtractPtr(obj))
}

func (c_ Control) IsEnabled() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isEnabled"))
	return rv
}

func (c_ Control) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setEnabled:"), value)
}

func (c_ Control) DoubleValue() float64 {
	rv := objc.CallMethod[float64](c_, objc.SEL("doubleValue"))
	return rv
}

func (c_ Control) SetDoubleValue(value float64) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setDoubleValue:"), value)
}

func (c_ Control) FloatValue() float32 {
	rv := objc.CallMethod[float32](c_, objc.SEL("floatValue"))
	return rv
}

func (c_ Control) SetFloatValue(value float32) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setFloatValue:"), value)
}

func (c_ Control) IntValue() int32 {
	rv := objc.CallMethod[int32](c_, objc.SEL("intValue"))
	return rv
}

func (c_ Control) SetIntValue(value int32) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setIntValue:"), value)
}

func (c_ Control) IntegerValue() int {
	rv := objc.CallMethod[int](c_, objc.SEL("integerValue"))
	return rv
}

func (c_ Control) SetIntegerValue(value int) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setIntegerValue:"), value)
}

func (c_ Control) ObjectValue() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.SEL("objectValue"))
	return rv
}

func (c_ Control) SetObjectValue(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setObjectValue:"), objc.ExtractPtr(value))
}

func (c_ Control) StringValue() string {
	rv := objc.CallMethod[string](c_, objc.SEL("stringValue"))
	return rv
}

func (c_ Control) SetStringValue(value string) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setStringValue:"), value)
}

func (c_ Control) AttributedStringValue() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](c_, objc.SEL("attributedStringValue"))
	return rv
}

func (c_ Control) SetAttributedStringValue(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setAttributedStringValue:"), objc.ExtractPtr(value))
}

func (c_ Control) Alignment() TextAlignment {
	rv := objc.CallMethod[TextAlignment](c_, objc.SEL("alignment"))
	return rv
}

func (c_ Control) SetAlignment(value TextAlignment) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setAlignment:"), value)
}

func (c_ Control) Font() Font {
	rv := objc.CallMethod[Font](c_, objc.SEL("font"))
	return rv
}

func (c_ Control) SetFont(value IFont) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setFont:"), objc.ExtractPtr(value))
}

func (c_ Control) LineBreakMode() LineBreakMode {
	rv := objc.CallMethod[LineBreakMode](c_, objc.SEL("lineBreakMode"))
	return rv
}

func (c_ Control) SetLineBreakMode(value LineBreakMode) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setLineBreakMode:"), value)
}

func (c_ Control) UsesSingleLineMode() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("usesSingleLineMode"))
	return rv
}

func (c_ Control) SetUsesSingleLineMode(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setUsesSingleLineMode:"), value)
}

func (c_ Control) Formatter() foundation.Formatter {
	rv := objc.CallMethod[foundation.Formatter](c_, objc.SEL("formatter"))
	return rv
}

func (c_ Control) SetFormatter(value foundation.IFormatter) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setFormatter:"), objc.ExtractPtr(value))
}

func (c_ Control) BaseWritingDirection() WritingDirection {
	rv := objc.CallMethod[WritingDirection](c_, objc.SEL("baseWritingDirection"))
	return rv
}

func (c_ Control) SetBaseWritingDirection(value WritingDirection) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setBaseWritingDirection:"), value)
}

func (c_ Control) AllowsExpansionToolTips() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("allowsExpansionToolTips"))
	return rv
}

func (c_ Control) SetAllowsExpansionToolTips(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setAllowsExpansionToolTips:"), value)
}

func (c_ Control) ControlSize() ControlSize {
	rv := objc.CallMethod[ControlSize](c_, objc.SEL("controlSize"))
	return rv
}

func (c_ Control) SetControlSize(value ControlSize) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setControlSize:"), value)
}

func (c_ Control) IsHighlighted() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isHighlighted"))
	return rv
}

func (c_ Control) SetHighlighted(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setHighlighted:"), value)
}

func (c_ Control) Action() objc.Selector {
	rv := objc.CallMethod[objc.Selector](c_, objc.SEL("action"))
	return rv
}

func (c_ Control) SetAction(value objc.Selector) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setAction:"), value)
}

// weak property
func (c_ Control) Target() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.SEL("target"))
	return rv
}

// weak property
func (c_ Control) SetTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setTarget:"), objc.ExtractPtr(value))
}

func (c_ Control) IsContinuous() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isContinuous"))
	return rv
}

func (c_ Control) SetContinuous(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setContinuous:"), value)
}

func (c_ Control) SetTag(value int) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setTag:"), value)
}

func (c_ Control) RefusesFirstResponder() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("refusesFirstResponder"))
	return rv
}

func (c_ Control) SetRefusesFirstResponder(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setRefusesFirstResponder:"), value)
}

func (c_ Control) IgnoresMultiClick() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("ignoresMultiClick"))
	return rv
}

func (c_ Control) SetIgnoresMultiClick(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setIgnoresMultiClick:"), value)
}

// deprecated
func (c_ Control) Cell() Cell {
	rv := objc.CallMethod[Cell](c_, objc.SEL("cell"))
	return rv
}

// deprecated
func (c_ Control) SetCell(value ICell) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setCell:"), objc.ExtractPtr(value))
}

// deprecated
func (cc _ControlClass) CellClass() objc.Class {
	rv := objc.CallMethod[objc.Class](cc, objc.SEL("cellClass"))
	return rv
}

// deprecated
func (cc _ControlClass) SetCellClass(value objc.IClass) {
	objc.CallMethod[objc.Void](cc, objc.SEL("setCellClass:"), objc.ExtractPtr(value))
}
