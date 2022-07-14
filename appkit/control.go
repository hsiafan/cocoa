// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	SetNeedsDisplay()
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
	rv := ffi.CallMethod[Control](c_, "initWithFrame:", frameRect)
	rv.Autorelease()
	return rv
}

func (c_ Control) Init() Control {
	rv := ffi.CallMethod[Control](c_, "init")
	rv.Autorelease()
	return rv
}

func (cc _ControlClass) Alloc() Control {
	rv := ffi.CallMethod[Control](cc, "alloc")
	return rv
}

func (cc _ControlClass) New() Control {
	rv := ffi.CallMethod[Control](cc, "new")
	rv.Autorelease()
	return rv
}

func NewControl() Control {
	return ControlClass.New()
}

func (c_ Control) TakeDoubleValueFrom(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "takeDoubleValueFrom:", sender)
}

func (c_ Control) TakeFloatValueFrom(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "takeFloatValueFrom:", sender)
}

func (c_ Control) TakeIntValueFrom(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "takeIntValueFrom:", sender)
}

func (c_ Control) TakeIntegerValueFrom(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "takeIntegerValueFrom:", sender)
}

func (c_ Control) TakeObjectValueFrom(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "takeObjectValueFrom:", sender)
}

func (c_ Control) TakeStringValueFrom(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "takeStringValueFrom:", sender)
}

// deprecated
func (c_ Control) SetFloatingPointFormat_Left_Right(autoRange bool, leftDigits uint, rightDigits uint) {
	ffi.CallMethod[ffi.Void](c_, "setFloatingPointFormat:left:right:", autoRange, leftDigits, rightDigits)
}

func (c_ Control) DrawWithExpansionFrame_InView(contentFrame foundation.Rect, view IView) {
	ffi.CallMethod[ffi.Void](c_, "drawWithExpansionFrame:inView:", contentFrame, view)
}

func (c_ Control) ExpansionFrameWithFrame(contentFrame foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](c_, "expansionFrameWithFrame:", contentFrame)
	return rv
}

func (c_ Control) AbortEditing() bool {
	rv := ffi.CallMethod[bool](c_, "abortEditing")
	return rv
}

func (c_ Control) CurrentEditor() Text {
	rv := ffi.CallMethod[Text](c_, "currentEditor")
	return rv
}

func (c_ Control) ValidateEditing() {
	ffi.CallMethod[ffi.Void](c_, "validateEditing")
}

func (c_ Control) EditWithFrame_Editor_Delegate_Event(rect foundation.Rect, textObj IText, delegate objc.IObject, event IEvent) {
	ffi.CallMethod[ffi.Void](c_, "editWithFrame:editor:delegate:event:", rect, textObj, delegate, event)
}

func (c_ Control) EndEditing(textObj IText) {
	ffi.CallMethod[ffi.Void](c_, "endEditing:", textObj)
}

func (c_ Control) SelectWithFrame_Editor_Delegate_Start_Length(rect foundation.Rect, textObj IText, delegate objc.IObject, selStart int, selLength int) {
	ffi.CallMethod[ffi.Void](c_, "selectWithFrame:editor:delegate:start:length:", rect, textObj, delegate, selStart, selLength)
}

func (c_ Control) SizeThatFits(size foundation.Size) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "sizeThatFits:", size)
	return rv
}

func (c_ Control) SizeToFit() {
	ffi.CallMethod[ffi.Void](c_, "sizeToFit")
}

func (c_ Control) SendAction_To(action objc.Selector, target objc.IObject) bool {
	rv := ffi.CallMethod[bool](c_, "sendAction:to:", action, target)
	return rv
}

func (c_ Control) SendActionOn(mask EventMask) int {
	rv := ffi.CallMethod[int](c_, "sendActionOn:", mask)
	return rv
}

func (c_ Control) PerformClick(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "performClick:", sender)
}

func (c_ Control) InvalidateIntrinsicContentSizeForCell(cell ICell) {
	ffi.CallMethod[ffi.Void](c_, "invalidateIntrinsicContentSizeForCell:", cell)
}

// deprecated
func (c_ Control) SelectedCell() Cell {
	rv := ffi.CallMethod[Cell](c_, "selectedCell")
	return rv
}

// deprecated
func (c_ Control) SelectedTag() int {
	rv := ffi.CallMethod[int](c_, "selectedTag")
	return rv
}

// deprecated
func (c_ Control) SetNeedsDisplay() {
	ffi.CallMethod[ffi.Void](c_, "setNeedsDisplay")
}

// deprecated
func (c_ Control) CalcSize() {
	ffi.CallMethod[ffi.Void](c_, "calcSize")
}

// deprecated
func (c_ Control) SelectCell(cell ICell) {
	ffi.CallMethod[ffi.Void](c_, "selectCell:", cell)
}

// deprecated
func (c_ Control) DrawCell(cell ICell) {
	ffi.CallMethod[ffi.Void](c_, "drawCell:", cell)
}

// deprecated
func (c_ Control) DrawCellInside(cell ICell) {
	ffi.CallMethod[ffi.Void](c_, "drawCellInside:", cell)
}

// deprecated
func (c_ Control) UpdateCell(cell ICell) {
	ffi.CallMethod[ffi.Void](c_, "updateCell:", cell)
}

// deprecated
func (c_ Control) UpdateCellInside(cell ICell) {
	ffi.CallMethod[ffi.Void](c_, "updateCellInside:", cell)
}

// deprecated
func (c_ Control) ControlTextDidBeginEditing(obj foundation.INotification) {
	ffi.CallMethod[ffi.Void](c_, "controlTextDidBeginEditing:", obj)
}

// deprecated
func (c_ Control) ControlTextDidChange(obj foundation.INotification) {
	ffi.CallMethod[ffi.Void](c_, "controlTextDidChange:", obj)
}

// deprecated
func (c_ Control) ControlTextDidEndEditing(obj foundation.INotification) {
	ffi.CallMethod[ffi.Void](c_, "controlTextDidEndEditing:", obj)
}

func (c_ Control) IsEnabled() bool {
	rv := ffi.CallMethod[bool](c_, "isEnabled")
	return rv
}

func (c_ Control) SetEnabled(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setEnabled:", value)
}

func (c_ Control) DoubleValue() float64 {
	rv := ffi.CallMethod[float64](c_, "doubleValue")
	return rv
}

func (c_ Control) SetDoubleValue(value float64) {
	ffi.CallMethod[ffi.Void](c_, "setDoubleValue:", value)
}

func (c_ Control) FloatValue() float32 {
	rv := ffi.CallMethod[float32](c_, "floatValue")
	return rv
}

func (c_ Control) SetFloatValue(value float32) {
	ffi.CallMethod[ffi.Void](c_, "setFloatValue:", value)
}

func (c_ Control) IntValue() int32 {
	rv := ffi.CallMethod[int32](c_, "intValue")
	return rv
}

func (c_ Control) SetIntValue(value int32) {
	ffi.CallMethod[ffi.Void](c_, "setIntValue:", value)
}

func (c_ Control) IntegerValue() int {
	rv := ffi.CallMethod[int](c_, "integerValue")
	return rv
}

func (c_ Control) SetIntegerValue(value int) {
	ffi.CallMethod[ffi.Void](c_, "setIntegerValue:", value)
}

func (c_ Control) ObjectValue() objc.Object {
	rv := ffi.CallMethod[objc.Object](c_, "objectValue")
	return rv
}

func (c_ Control) SetObjectValue(value objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "setObjectValue:", value)
}

func (c_ Control) StringValue() string {
	rv := ffi.CallMethod[string](c_, "stringValue")
	return rv
}

func (c_ Control) SetStringValue(value string) {
	ffi.CallMethod[ffi.Void](c_, "setStringValue:", value)
}

func (c_ Control) AttributedStringValue() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](c_, "attributedStringValue")
	return rv
}

func (c_ Control) SetAttributedStringValue(value foundation.IAttributedString) {
	ffi.CallMethod[ffi.Void](c_, "setAttributedStringValue:", value)
}

func (c_ Control) Alignment() TextAlignment {
	rv := ffi.CallMethod[TextAlignment](c_, "alignment")
	return rv
}

func (c_ Control) SetAlignment(value TextAlignment) {
	ffi.CallMethod[ffi.Void](c_, "setAlignment:", value)
}

func (c_ Control) Font() Font {
	rv := ffi.CallMethod[Font](c_, "font")
	return rv
}

func (c_ Control) SetFont(value IFont) {
	ffi.CallMethod[ffi.Void](c_, "setFont:", value)
}

func (c_ Control) LineBreakMode() LineBreakMode {
	rv := ffi.CallMethod[LineBreakMode](c_, "lineBreakMode")
	return rv
}

func (c_ Control) SetLineBreakMode(value LineBreakMode) {
	ffi.CallMethod[ffi.Void](c_, "setLineBreakMode:", value)
}

func (c_ Control) UsesSingleLineMode() bool {
	rv := ffi.CallMethod[bool](c_, "usesSingleLineMode")
	return rv
}

func (c_ Control) SetUsesSingleLineMode(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setUsesSingleLineMode:", value)
}

func (c_ Control) Formatter() foundation.Formatter {
	rv := ffi.CallMethod[foundation.Formatter](c_, "formatter")
	return rv
}

func (c_ Control) SetFormatter(value foundation.IFormatter) {
	ffi.CallMethod[ffi.Void](c_, "setFormatter:", value)
}

func (c_ Control) BaseWritingDirection() WritingDirection {
	rv := ffi.CallMethod[WritingDirection](c_, "baseWritingDirection")
	return rv
}

func (c_ Control) SetBaseWritingDirection(value WritingDirection) {
	ffi.CallMethod[ffi.Void](c_, "setBaseWritingDirection:", value)
}

func (c_ Control) AllowsExpansionToolTips() bool {
	rv := ffi.CallMethod[bool](c_, "allowsExpansionToolTips")
	return rv
}

func (c_ Control) SetAllowsExpansionToolTips(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setAllowsExpansionToolTips:", value)
}

func (c_ Control) ControlSize() ControlSize {
	rv := ffi.CallMethod[ControlSize](c_, "controlSize")
	return rv
}

func (c_ Control) SetControlSize(value ControlSize) {
	ffi.CallMethod[ffi.Void](c_, "setControlSize:", value)
}

func (c_ Control) IsHighlighted() bool {
	rv := ffi.CallMethod[bool](c_, "isHighlighted")
	return rv
}

func (c_ Control) SetHighlighted(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setHighlighted:", value)
}

func (c_ Control) Action() objc.Selector {
	rv := ffi.CallMethod[objc.Selector](c_, "action")
	return rv
}

func (c_ Control) SetAction(value objc.Selector) {
	ffi.CallMethod[ffi.Void](c_, "setAction:", value)
}

func (c_ Control) Target() objc.Object {
	rv := ffi.CallMethod[objc.Object](c_, "target")
	return rv
}

func (c_ Control) SetTarget(value objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "setTarget:", value)
}

func (c_ Control) IsContinuous() bool {
	rv := ffi.CallMethod[bool](c_, "isContinuous")
	return rv
}

func (c_ Control) SetContinuous(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setContinuous:", value)
}

func (c_ Control) SetTag(value int) {
	ffi.CallMethod[ffi.Void](c_, "setTag:", value)
}

func (c_ Control) RefusesFirstResponder() bool {
	rv := ffi.CallMethod[bool](c_, "refusesFirstResponder")
	return rv
}

func (c_ Control) SetRefusesFirstResponder(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setRefusesFirstResponder:", value)
}

func (c_ Control) IgnoresMultiClick() bool {
	rv := ffi.CallMethod[bool](c_, "ignoresMultiClick")
	return rv
}

func (c_ Control) SetIgnoresMultiClick(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setIgnoresMultiClick:", value)
}

// deprecated
func (c_ Control) Cell() Cell {
	rv := ffi.CallMethod[Cell](c_, "cell")
	return rv
}

// deprecated
func (c_ Control) SetCell(value ICell) {
	ffi.CallMethod[ffi.Void](c_, "setCell:", value)
}

type ControlTextEditingDelegate interface {
	ImplementsControl_IsValidObject() bool
	// optional
	Control_IsValidObject(control Control, obj objc.Object) bool
	ImplementsControl_DidFailToValidatePartialString_ErrorDescription() bool
	// optional
	Control_DidFailToValidatePartialString_ErrorDescription(control Control, _string string, error string)
	ImplementsControl_DidFailToFormatString_ErrorDescription() bool
	// optional
	Control_DidFailToFormatString_ErrorDescription(control Control, _string string, error string) bool
	ImplementsControl_TextShouldBeginEditing() bool
	// optional
	Control_TextShouldBeginEditing(control Control, fieldEditor Text) bool
	ImplementsControl_TextShouldEndEditing() bool
	// optional
	Control_TextShouldEndEditing(control Control, fieldEditor Text) bool
	ImplementsControl_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem() bool
	// optional
	Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(control Control, textView TextView, words []string, charRange foundation.Range, index *int) []string
	ImplementsControl_TextView_DoCommandBySelector() bool
	// optional
	Control_TextView_DoCommandBySelector(control Control, textView TextView, commandSelector objc.Selector) bool
	ImplementsControlTextDidBeginEditing() bool
	// optional
	ControlTextDidBeginEditing(obj foundation.Notification)
	ImplementsControlTextDidChange() bool
	// optional
	ControlTextDidChange(obj foundation.Notification)
	ImplementsControlTextDidEndEditing() bool
	// optional
	ControlTextDidEndEditing(obj foundation.Notification)
}

type ControlTextEditingDelegateImpl struct {
	_Control_IsValidObject                                                func(control Control, obj objc.Object) bool
	_Control_DidFailToValidatePartialString_ErrorDescription              func(control Control, _string string, error string)
	_Control_DidFailToFormatString_ErrorDescription                       func(control Control, _string string, error string) bool
	_Control_TextShouldBeginEditing                                       func(control Control, fieldEditor Text) bool
	_Control_TextShouldEndEditing                                         func(control Control, fieldEditor Text) bool
	_Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem func(control Control, textView TextView, words []string, charRange foundation.Range, index *int) []string
	_Control_TextView_DoCommandBySelector                                 func(control Control, textView TextView, commandSelector objc.Selector) bool
	_ControlTextDidBeginEditing                                           func(obj foundation.Notification)
	_ControlTextDidChange                                                 func(obj foundation.Notification)
	_ControlTextDidEndEditing                                             func(obj foundation.Notification)
}

func (di *ControlTextEditingDelegateImpl) ImplementsControl_IsValidObject() bool {
	return di._Control_IsValidObject != nil
}

func (di *ControlTextEditingDelegateImpl) SetControl_IsValidObject(f func(control Control, obj objc.Object) bool) {
	di._Control_IsValidObject = f
}

func (di *ControlTextEditingDelegateImpl) Control_IsValidObject(control Control, obj objc.Object) bool {
	return di._Control_IsValidObject(control, obj)
}
func (di *ControlTextEditingDelegateImpl) ImplementsControl_DidFailToValidatePartialString_ErrorDescription() bool {
	return di._Control_DidFailToValidatePartialString_ErrorDescription != nil
}

func (di *ControlTextEditingDelegateImpl) SetControl_DidFailToValidatePartialString_ErrorDescription(f func(control Control, _string string, error string)) {
	di._Control_DidFailToValidatePartialString_ErrorDescription = f
}

func (di *ControlTextEditingDelegateImpl) Control_DidFailToValidatePartialString_ErrorDescription(control Control, _string string, error string) {
	di._Control_DidFailToValidatePartialString_ErrorDescription(control, _string, error)
}
func (di *ControlTextEditingDelegateImpl) ImplementsControl_DidFailToFormatString_ErrorDescription() bool {
	return di._Control_DidFailToFormatString_ErrorDescription != nil
}

func (di *ControlTextEditingDelegateImpl) SetControl_DidFailToFormatString_ErrorDescription(f func(control Control, _string string, error string) bool) {
	di._Control_DidFailToFormatString_ErrorDescription = f
}

func (di *ControlTextEditingDelegateImpl) Control_DidFailToFormatString_ErrorDescription(control Control, _string string, error string) bool {
	return di._Control_DidFailToFormatString_ErrorDescription(control, _string, error)
}
func (di *ControlTextEditingDelegateImpl) ImplementsControl_TextShouldBeginEditing() bool {
	return di._Control_TextShouldBeginEditing != nil
}

func (di *ControlTextEditingDelegateImpl) SetControl_TextShouldBeginEditing(f func(control Control, fieldEditor Text) bool) {
	di._Control_TextShouldBeginEditing = f
}

func (di *ControlTextEditingDelegateImpl) Control_TextShouldBeginEditing(control Control, fieldEditor Text) bool {
	return di._Control_TextShouldBeginEditing(control, fieldEditor)
}
func (di *ControlTextEditingDelegateImpl) ImplementsControl_TextShouldEndEditing() bool {
	return di._Control_TextShouldEndEditing != nil
}

func (di *ControlTextEditingDelegateImpl) SetControl_TextShouldEndEditing(f func(control Control, fieldEditor Text) bool) {
	di._Control_TextShouldEndEditing = f
}

func (di *ControlTextEditingDelegateImpl) Control_TextShouldEndEditing(control Control, fieldEditor Text) bool {
	return di._Control_TextShouldEndEditing(control, fieldEditor)
}
func (di *ControlTextEditingDelegateImpl) ImplementsControl_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem() bool {
	return di._Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem != nil
}

func (di *ControlTextEditingDelegateImpl) SetControl_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(f func(control Control, textView TextView, words []string, charRange foundation.Range, index *int) []string) {
	di._Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem = f
}

func (di *ControlTextEditingDelegateImpl) Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(control Control, textView TextView, words []string, charRange foundation.Range, index *int) []string {
	return di._Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(control, textView, words, charRange, index)
}
func (di *ControlTextEditingDelegateImpl) ImplementsControl_TextView_DoCommandBySelector() bool {
	return di._Control_TextView_DoCommandBySelector != nil
}

func (di *ControlTextEditingDelegateImpl) SetControl_TextView_DoCommandBySelector(f func(control Control, textView TextView, commandSelector objc.Selector) bool) {
	di._Control_TextView_DoCommandBySelector = f
}

func (di *ControlTextEditingDelegateImpl) Control_TextView_DoCommandBySelector(control Control, textView TextView, commandSelector objc.Selector) bool {
	return di._Control_TextView_DoCommandBySelector(control, textView, commandSelector)
}
func (di *ControlTextEditingDelegateImpl) ImplementsControlTextDidBeginEditing() bool {
	return di._ControlTextDidBeginEditing != nil
}

func (di *ControlTextEditingDelegateImpl) SetControlTextDidBeginEditing(f func(obj foundation.Notification)) {
	di._ControlTextDidBeginEditing = f
}

func (di *ControlTextEditingDelegateImpl) ControlTextDidBeginEditing(obj foundation.Notification) {
	di._ControlTextDidBeginEditing(obj)
}
func (di *ControlTextEditingDelegateImpl) ImplementsControlTextDidChange() bool {
	return di._ControlTextDidChange != nil
}

func (di *ControlTextEditingDelegateImpl) SetControlTextDidChange(f func(obj foundation.Notification)) {
	di._ControlTextDidChange = f
}

func (di *ControlTextEditingDelegateImpl) ControlTextDidChange(obj foundation.Notification) {
	di._ControlTextDidChange(obj)
}
func (di *ControlTextEditingDelegateImpl) ImplementsControlTextDidEndEditing() bool {
	return di._ControlTextDidEndEditing != nil
}

func (di *ControlTextEditingDelegateImpl) SetControlTextDidEndEditing(f func(obj foundation.Notification)) {
	di._ControlTextDidEndEditing = f
}

func (di *ControlTextEditingDelegateImpl) ControlTextDidEndEditing(obj foundation.Notification) {
	di._ControlTextDidEndEditing(obj)
}

type ControlTextEditingDelegateWrapper struct {
	objc.Object
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControl_IsValidObject() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:isValidObject:"))
}

func (c_ ControlTextEditingDelegateWrapper) Control_IsValidObject(control IControl, obj objc.IObject) bool {
	rv := ffi.CallMethod[bool](c_, "control:isValidObject:", control, obj)
	return rv
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControl_DidFailToValidatePartialString_ErrorDescription() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:didFailToValidatePartialString:errorDescription:"))
}

func (c_ ControlTextEditingDelegateWrapper) Control_DidFailToValidatePartialString_ErrorDescription(control IControl, _string string, error string) {
	ffi.CallMethod[ffi.Void](c_, "control:didFailToValidatePartialString:errorDescription:", control, _string, error)
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControl_DidFailToFormatString_ErrorDescription() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:didFailToFormatString:errorDescription:"))
}

func (c_ ControlTextEditingDelegateWrapper) Control_DidFailToFormatString_ErrorDescription(control IControl, _string string, error string) bool {
	rv := ffi.CallMethod[bool](c_, "control:didFailToFormatString:errorDescription:", control, _string, error)
	return rv
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControl_TextShouldBeginEditing() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:textShouldBeginEditing:"))
}

func (c_ ControlTextEditingDelegateWrapper) Control_TextShouldBeginEditing(control IControl, fieldEditor IText) bool {
	rv := ffi.CallMethod[bool](c_, "control:textShouldBeginEditing:", control, fieldEditor)
	return rv
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControl_TextShouldEndEditing() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:textShouldEndEditing:"))
}

func (c_ ControlTextEditingDelegateWrapper) Control_TextShouldEndEditing(control IControl, fieldEditor IText) bool {
	rv := ffi.CallMethod[bool](c_, "control:textShouldEndEditing:", control, fieldEditor)
	return rv
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControl_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:textView:completions:forPartialWordRange:indexOfSelectedItem:"))
}

func (c_ ControlTextEditingDelegateWrapper) Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(control IControl, textView ITextView, words []string, charRange foundation.Range, index *int) []string {
	rv := ffi.CallMethod[[]string](c_, "control:textView:completions:forPartialWordRange:indexOfSelectedItem:", control, textView, words, charRange, index)
	return rv
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControl_TextView_DoCommandBySelector() bool {
	return c_.RespondsToSelector(objc.GetSelector("control:textView:doCommandBySelector:"))
}

func (c_ ControlTextEditingDelegateWrapper) Control_TextView_DoCommandBySelector(control IControl, textView ITextView, commandSelector objc.Selector) bool {
	rv := ffi.CallMethod[bool](c_, "control:textView:doCommandBySelector:", control, textView, commandSelector)
	return rv
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControlTextDidBeginEditing() bool {
	return c_.RespondsToSelector(objc.GetSelector("controlTextDidBeginEditing:"))
}

func (c_ ControlTextEditingDelegateWrapper) ControlTextDidBeginEditing(obj foundation.INotification) {
	ffi.CallMethod[ffi.Void](c_, "controlTextDidBeginEditing:", obj)
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControlTextDidChange() bool {
	return c_.RespondsToSelector(objc.GetSelector("controlTextDidChange:"))
}

func (c_ ControlTextEditingDelegateWrapper) ControlTextDidChange(obj foundation.INotification) {
	ffi.CallMethod[ffi.Void](c_, "controlTextDidChange:", obj)
}

func (c_ *ControlTextEditingDelegateWrapper) ImplementsControlTextDidEndEditing() bool {
	return c_.RespondsToSelector(objc.GetSelector("controlTextDidEndEditing:"))
}

func (c_ ControlTextEditingDelegateWrapper) ControlTextDidEndEditing(obj foundation.INotification) {
	ffi.CallMethod[ffi.Void](c_, "controlTextDidEndEditing:", obj)
}
