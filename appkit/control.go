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
	rv := objc.CallMethod[Control](c_, "initWithFrame:", frameRect)
	return rv
}

func (c_ Control) Init() Control {
	rv := objc.CallMethod[Control](c_, "init")
	return rv
}

func (cc _ControlClass) Alloc() Control {
	rv := objc.CallMethod[Control](cc, "alloc")
	return rv
}

func (cc _ControlClass) New() Control {
	rv := objc.CallMethod[Control](cc, "new")
	rv.Autorelease()
	return rv
}

func NewControl() Control {
	return ControlClass.New()
}

func (c_ Control) TakeDoubleValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, "takeDoubleValueFrom:", sender)
}

func (c_ Control) TakeFloatValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, "takeFloatValueFrom:", sender)
}

func (c_ Control) TakeIntValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, "takeIntValueFrom:", sender)
}

func (c_ Control) TakeIntegerValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, "takeIntegerValueFrom:", sender)
}

func (c_ Control) TakeObjectValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, "takeObjectValueFrom:", sender)
}

func (c_ Control) TakeStringValueFrom(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, "takeStringValueFrom:", sender)
}

// deprecated
func (c_ Control) SetFloatingPointFormat_Left_Right(autoRange bool, leftDigits uint, rightDigits uint) {
	objc.CallMethod[objc.Void](c_, "setFloatingPointFormat:left:right:", autoRange, leftDigits, rightDigits)
}

func (c_ Control) DrawWithExpansionFrame_InView(contentFrame foundation.Rect, view IView) {
	objc.CallMethod[objc.Void](c_, "drawWithExpansionFrame:inView:", contentFrame, view)
}

func (c_ Control) ExpansionFrameWithFrame(contentFrame foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, "expansionFrameWithFrame:", contentFrame)
	return rv
}

func (c_ Control) AbortEditing() bool {
	rv := objc.CallMethod[bool](c_, "abortEditing")
	return rv
}

func (c_ Control) CurrentEditor() Text {
	rv := objc.CallMethod[Text](c_, "currentEditor")
	return rv
}

func (c_ Control) ValidateEditing() {
	objc.CallMethod[objc.Void](c_, "validateEditing")
}

func (c_ Control) EditWithFrame_Editor_Delegate_Event(rect foundation.Rect, textObj IText, delegate objc.IObject, event IEvent) {
	objc.CallMethod[objc.Void](c_, "editWithFrame:editor:delegate:event:", rect, textObj, delegate, event)
}

func (c_ Control) EndEditing(textObj IText) {
	objc.CallMethod[objc.Void](c_, "endEditing:", textObj)
}

func (c_ Control) SelectWithFrame_Editor_Delegate_Start_Length(rect foundation.Rect, textObj IText, delegate objc.IObject, selStart int, selLength int) {
	objc.CallMethod[objc.Void](c_, "selectWithFrame:editor:delegate:start:length:", rect, textObj, delegate, selStart, selLength)
}

func (c_ Control) SizeThatFits(size foundation.Size) foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, "sizeThatFits:", size)
	return rv
}

func (c_ Control) SizeToFit() {
	objc.CallMethod[objc.Void](c_, "sizeToFit")
}

func (c_ Control) SendAction_To(action objc.Selector, target objc.IObject) bool {
	rv := objc.CallMethod[bool](c_, "sendAction:to:", action, target)
	return rv
}

func (c_ Control) SendActionOn(mask EventMask) int {
	rv := objc.CallMethod[int](c_, "sendActionOn:", mask)
	return rv
}

func (c_ Control) PerformClick(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, "performClick:", sender)
}

func (c_ Control) InvalidateIntrinsicContentSizeForCell(cell ICell) {
	objc.CallMethod[objc.Void](c_, "invalidateIntrinsicContentSizeForCell:", cell)
}

// deprecated
func (c_ Control) SelectedCell() Cell {
	rv := objc.CallMethod[Cell](c_, "selectedCell")
	return rv
}

// deprecated
func (c_ Control) SelectedTag() int {
	rv := objc.CallMethod[int](c_, "selectedTag")
	return rv
}

// deprecated
func (c_ Control) SetNeedsDisplay0() {
	objc.CallMethod[objc.Void](c_, "setNeedsDisplay")
}

// deprecated
func (c_ Control) CalcSize() {
	objc.CallMethod[objc.Void](c_, "calcSize")
}

// deprecated
func (c_ Control) SelectCell(cell ICell) {
	objc.CallMethod[objc.Void](c_, "selectCell:", cell)
}

// deprecated
func (c_ Control) DrawCell(cell ICell) {
	objc.CallMethod[objc.Void](c_, "drawCell:", cell)
}

// deprecated
func (c_ Control) DrawCellInside(cell ICell) {
	objc.CallMethod[objc.Void](c_, "drawCellInside:", cell)
}

// deprecated
func (c_ Control) UpdateCell(cell ICell) {
	objc.CallMethod[objc.Void](c_, "updateCell:", cell)
}

// deprecated
func (c_ Control) UpdateCellInside(cell ICell) {
	objc.CallMethod[objc.Void](c_, "updateCellInside:", cell)
}

// deprecated
func (c_ Control) ControlTextDidBeginEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](c_, "controlTextDidBeginEditing:", obj)
}

// deprecated
func (c_ Control) ControlTextDidChange(obj foundation.INotification) {
	objc.CallMethod[objc.Void](c_, "controlTextDidChange:", obj)
}

// deprecated
func (c_ Control) ControlTextDidEndEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](c_, "controlTextDidEndEditing:", obj)
}

func (c_ Control) IsEnabled() bool {
	rv := objc.CallMethod[bool](c_, "isEnabled")
	return rv
}

func (c_ Control) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](c_, "setEnabled:", value)
}

func (c_ Control) DoubleValue() float64 {
	rv := objc.CallMethod[float64](c_, "doubleValue")
	return rv
}

func (c_ Control) SetDoubleValue(value float64) {
	objc.CallMethod[objc.Void](c_, "setDoubleValue:", value)
}

func (c_ Control) FloatValue() float32 {
	rv := objc.CallMethod[float32](c_, "floatValue")
	return rv
}

func (c_ Control) SetFloatValue(value float32) {
	objc.CallMethod[objc.Void](c_, "setFloatValue:", value)
}

func (c_ Control) IntValue() int32 {
	rv := objc.CallMethod[int32](c_, "intValue")
	return rv
}

func (c_ Control) SetIntValue(value int32) {
	objc.CallMethod[objc.Void](c_, "setIntValue:", value)
}

func (c_ Control) IntegerValue() int {
	rv := objc.CallMethod[int](c_, "integerValue")
	return rv
}

func (c_ Control) SetIntegerValue(value int) {
	objc.CallMethod[objc.Void](c_, "setIntegerValue:", value)
}

func (c_ Control) ObjectValue() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, "objectValue")
	return rv
}

func (c_ Control) SetObjectValue(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, "setObjectValue:", value)
}

func (c_ Control) StringValue() string {
	rv := objc.CallMethod[string](c_, "stringValue")
	return rv
}

func (c_ Control) SetStringValue(value string) {
	objc.CallMethod[objc.Void](c_, "setStringValue:", value)
}

func (c_ Control) AttributedStringValue() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](c_, "attributedStringValue")
	return rv
}

func (c_ Control) SetAttributedStringValue(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](c_, "setAttributedStringValue:", value)
}

func (c_ Control) Alignment() TextAlignment {
	rv := objc.CallMethod[TextAlignment](c_, "alignment")
	return rv
}

func (c_ Control) SetAlignment(value TextAlignment) {
	objc.CallMethod[objc.Void](c_, "setAlignment:", value)
}

func (c_ Control) Font() Font {
	rv := objc.CallMethod[Font](c_, "font")
	return rv
}

func (c_ Control) SetFont(value IFont) {
	objc.CallMethod[objc.Void](c_, "setFont:", value)
}

func (c_ Control) LineBreakMode() LineBreakMode {
	rv := objc.CallMethod[LineBreakMode](c_, "lineBreakMode")
	return rv
}

func (c_ Control) SetLineBreakMode(value LineBreakMode) {
	objc.CallMethod[objc.Void](c_, "setLineBreakMode:", value)
}

func (c_ Control) UsesSingleLineMode() bool {
	rv := objc.CallMethod[bool](c_, "usesSingleLineMode")
	return rv
}

func (c_ Control) SetUsesSingleLineMode(value bool) {
	objc.CallMethod[objc.Void](c_, "setUsesSingleLineMode:", value)
}

func (c_ Control) Formatter() foundation.Formatter {
	rv := objc.CallMethod[foundation.Formatter](c_, "formatter")
	return rv
}

func (c_ Control) SetFormatter(value foundation.IFormatter) {
	objc.CallMethod[objc.Void](c_, "setFormatter:", value)
}

func (c_ Control) BaseWritingDirection() WritingDirection {
	rv := objc.CallMethod[WritingDirection](c_, "baseWritingDirection")
	return rv
}

func (c_ Control) SetBaseWritingDirection(value WritingDirection) {
	objc.CallMethod[objc.Void](c_, "setBaseWritingDirection:", value)
}

func (c_ Control) AllowsExpansionToolTips() bool {
	rv := objc.CallMethod[bool](c_, "allowsExpansionToolTips")
	return rv
}

func (c_ Control) SetAllowsExpansionToolTips(value bool) {
	objc.CallMethod[objc.Void](c_, "setAllowsExpansionToolTips:", value)
}

func (c_ Control) ControlSize() ControlSize {
	rv := objc.CallMethod[ControlSize](c_, "controlSize")
	return rv
}

func (c_ Control) SetControlSize(value ControlSize) {
	objc.CallMethod[objc.Void](c_, "setControlSize:", value)
}

func (c_ Control) IsHighlighted() bool {
	rv := objc.CallMethod[bool](c_, "isHighlighted")
	return rv
}

func (c_ Control) SetHighlighted(value bool) {
	objc.CallMethod[objc.Void](c_, "setHighlighted:", value)
}

func (c_ Control) Action() objc.Selector {
	rv := objc.CallMethod[objc.Selector](c_, "action")
	return rv
}

func (c_ Control) SetAction(value objc.Selector) {
	objc.CallMethod[objc.Void](c_, "setAction:", value)
}

func (c_ Control) Target() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, "target")
	return rv
}

func (c_ Control) SetTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, "setTarget:", value)
}

func (c_ Control) IsContinuous() bool {
	rv := objc.CallMethod[bool](c_, "isContinuous")
	return rv
}

func (c_ Control) SetContinuous(value bool) {
	objc.CallMethod[objc.Void](c_, "setContinuous:", value)
}

func (c_ Control) SetTag(value int) {
	objc.CallMethod[objc.Void](c_, "setTag:", value)
}

func (c_ Control) RefusesFirstResponder() bool {
	rv := objc.CallMethod[bool](c_, "refusesFirstResponder")
	return rv
}

func (c_ Control) SetRefusesFirstResponder(value bool) {
	objc.CallMethod[objc.Void](c_, "setRefusesFirstResponder:", value)
}

func (c_ Control) IgnoresMultiClick() bool {
	rv := objc.CallMethod[bool](c_, "ignoresMultiClick")
	return rv
}

func (c_ Control) SetIgnoresMultiClick(value bool) {
	objc.CallMethod[objc.Void](c_, "setIgnoresMultiClick:", value)
}

// deprecated
func (c_ Control) Cell() Cell {
	rv := objc.CallMethod[Cell](c_, "cell")
	return rv
}

// deprecated
func (c_ Control) SetCell(value ICell) {
	objc.CallMethod[objc.Void](c_, "setCell:", value)
}

// deprecated
func (cc _ControlClass) CellClass() objc.Class {
	rv := objc.CallMethod[objc.Class](cc, "cellClass")
	return rv
}

// deprecated
func (cc _ControlClass) SetCellClass(value objc.IClass) {
	objc.CallMethod[objc.Void](cc, "setCellClass:", value)
}
