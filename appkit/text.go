// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var TextClass = _TextClass{objc.GetClass("NSText")}

type _TextClass struct {
	objc.Class
}

type IText interface {
	IView
	ToggleRuler(sender objc.IObject)
	ReplaceCharactersInRange_WithRTF(_range foundation.Range, rtfData []byte)
	ReplaceCharactersInRange_WithRTFD(_range foundation.Range, rtfdData []byte)
	ReplaceCharactersInRange_WithString(_range foundation.Range, _string string)
	SelectAll(sender objc.IObject)
	Copy(sender objc.IObject)
	Cut(sender objc.IObject)
	Paste(sender objc.IObject)
	CopyFont(sender objc.IObject)
	PasteFont(sender objc.IObject)
	CopyRuler(sender objc.IObject)
	PasteRuler(sender objc.IObject)
	Delete(sender objc.IObject)
	ChangeFont(sender objc.IObject)
	SetFont_Range(font IFont, _range foundation.Range)
	AlignCenter(sender objc.IObject)
	AlignLeft(sender objc.IObject)
	AlignRight(sender objc.IObject)
	SetTextColor_Range(color IColor, _range foundation.Range)
	Superscript(sender objc.IObject)
	Subscript(sender objc.IObject)
	Unscript(sender objc.IObject)
	Underline(sender objc.IObject)
	ReadRTFDFromFile(path string) bool
	WriteRTFDToFile_Atomically(path string, flag bool) bool
	RTFDFromRange(_range foundation.Range) []byte
	RTFFromRange(_range foundation.Range) []byte
	CheckSpelling(sender objc.IObject)
	ShowGuessPanel(sender objc.IObject)
	SizeToFit()
	ScrollRangeToVisible(_range foundation.Range)
	String() string
	SetString(value string)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	IsEditable() bool
	SetEditable(value bool)
	IsSelectable() bool
	SetSelectable(value bool)
	IsFieldEditor() bool
	SetFieldEditor(value bool)
	IsRichText() bool
	SetRichText(value bool)
	ImportsGraphics() bool
	SetImportsGraphics(value bool)
	UsesFontPanel() bool
	SetUsesFontPanel(value bool)
	IsRulerVisible() bool
	SelectedRange() foundation.Range
	SetSelectedRange(value foundation.Range)
	Font() Font
	SetFont(value IFont)
	Alignment() TextAlignment
	SetAlignment(value TextAlignment)
	TextColor() Color
	SetTextColor(value IColor)
	BaseWritingDirection() WritingDirection
	SetBaseWritingDirection(value WritingDirection)
	MaxSize() foundation.Size
	SetMaxSize(value foundation.Size)
	MinSize() foundation.Size
	SetMinSize(value foundation.Size)
	IsVerticallyResizable() bool
	SetVerticallyResizable(value bool)
	IsHorizontallyResizable() bool
	SetHorizontallyResizable(value bool)
	Delegate() TextDelegateWrapper
	SetDelegate(value TextDelegate)
}

type Text struct {
	View
}

func MakeText(ptr unsafe.Pointer) Text {
	return Text{
		View: MakeView(ptr),
	}
}

func (t_ Text) InitWithFrame(frameRect foundation.Rect) Text {
	rv := ffi.CallMethod[Text](t_, "initWithFrame:", frameRect)
	return rv
}

func (t_ Text) Init() Text {
	rv := ffi.CallMethod[Text](t_, "init")
	return rv
}

func (tc _TextClass) Alloc() Text {
	rv := ffi.CallMethod[Text](tc, "alloc")
	return rv
}

func (tc _TextClass) New() Text {
	rv := ffi.CallMethod[Text](tc, "new")
	rv.Autorelease()
	return rv
}

func NewText() Text {
	return TextClass.New()
}

func (t_ Text) ToggleRuler(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "toggleRuler:", sender)
}

func (t_ Text) ReplaceCharactersInRange_WithRTF(_range foundation.Range, rtfData []byte) {
	ffi.CallMethod[ffi.Void](t_, "replaceCharactersInRange:withRTF:", _range, rtfData)
}

func (t_ Text) ReplaceCharactersInRange_WithRTFD(_range foundation.Range, rtfdData []byte) {
	ffi.CallMethod[ffi.Void](t_, "replaceCharactersInRange:withRTFD:", _range, rtfdData)
}

func (t_ Text) ReplaceCharactersInRange_WithString(_range foundation.Range, _string string) {
	ffi.CallMethod[ffi.Void](t_, "replaceCharactersInRange:withString:", _range, _string)
}

func (t_ Text) SelectAll(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "selectAll:", sender)
}

func (t_ Text) Copy(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "copy:", sender)
}

func (t_ Text) Cut(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "cut:", sender)
}

func (t_ Text) Paste(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "paste:", sender)
}

func (t_ Text) CopyFont(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "copyFont:", sender)
}

func (t_ Text) PasteFont(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "pasteFont:", sender)
}

func (t_ Text) CopyRuler(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "copyRuler:", sender)
}

func (t_ Text) PasteRuler(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "pasteRuler:", sender)
}

func (t_ Text) Delete(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "delete:", sender)
}

func (t_ Text) ChangeFont(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "changeFont:", sender)
}

func (t_ Text) SetFont_Range(font IFont, _range foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "setFont:range:", font, _range)
}

func (t_ Text) AlignCenter(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "alignCenter:", sender)
}

func (t_ Text) AlignLeft(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "alignLeft:", sender)
}

func (t_ Text) AlignRight(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "alignRight:", sender)
}

func (t_ Text) SetTextColor_Range(color IColor, _range foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "setTextColor:range:", color, _range)
}

func (t_ Text) Superscript(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "superscript:", sender)
}

func (t_ Text) Subscript(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "subscript:", sender)
}

func (t_ Text) Unscript(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "unscript:", sender)
}

func (t_ Text) Underline(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "underline:", sender)
}

func (t_ Text) ReadRTFDFromFile(path string) bool {
	rv := ffi.CallMethod[bool](t_, "readRTFDFromFile:", path)
	return rv
}

func (t_ Text) WriteRTFDToFile_Atomically(path string, flag bool) bool {
	rv := ffi.CallMethod[bool](t_, "writeRTFDToFile:atomically:", path, flag)
	return rv
}

func (t_ Text) RTFDFromRange(_range foundation.Range) []byte {
	rv := ffi.CallMethod[[]byte](t_, "RTFDFromRange:", _range)
	return rv
}

func (t_ Text) RTFFromRange(_range foundation.Range) []byte {
	rv := ffi.CallMethod[[]byte](t_, "RTFFromRange:", _range)
	return rv
}

func (t_ Text) CheckSpelling(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "checkSpelling:", sender)
}

func (t_ Text) ShowGuessPanel(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "showGuessPanel:", sender)
}

func (t_ Text) SizeToFit() {
	ffi.CallMethod[ffi.Void](t_, "sizeToFit")
}

func (t_ Text) ScrollRangeToVisible(_range foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "scrollRangeToVisible:", _range)
}

func (t_ Text) String() string {
	rv := ffi.CallMethod[string](t_, "string")
	return rv
}

func (t_ Text) SetString(value string) {
	ffi.CallMethod[ffi.Void](t_, "setString:", value)
}

func (t_ Text) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](t_, "backgroundColor")
	return rv
}

func (t_ Text) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](t_, "setBackgroundColor:", value)
}

func (t_ Text) DrawsBackground() bool {
	rv := ffi.CallMethod[bool](t_, "drawsBackground")
	return rv
}

func (t_ Text) SetDrawsBackground(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setDrawsBackground:", value)
}

func (t_ Text) IsEditable() bool {
	rv := ffi.CallMethod[bool](t_, "isEditable")
	return rv
}

func (t_ Text) SetEditable(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setEditable:", value)
}

func (t_ Text) IsSelectable() bool {
	rv := ffi.CallMethod[bool](t_, "isSelectable")
	return rv
}

func (t_ Text) SetSelectable(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setSelectable:", value)
}

func (t_ Text) IsFieldEditor() bool {
	rv := ffi.CallMethod[bool](t_, "isFieldEditor")
	return rv
}

func (t_ Text) SetFieldEditor(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setFieldEditor:", value)
}

func (t_ Text) IsRichText() bool {
	rv := ffi.CallMethod[bool](t_, "isRichText")
	return rv
}

func (t_ Text) SetRichText(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setRichText:", value)
}

func (t_ Text) ImportsGraphics() bool {
	rv := ffi.CallMethod[bool](t_, "importsGraphics")
	return rv
}

func (t_ Text) SetImportsGraphics(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setImportsGraphics:", value)
}

func (t_ Text) UsesFontPanel() bool {
	rv := ffi.CallMethod[bool](t_, "usesFontPanel")
	return rv
}

func (t_ Text) SetUsesFontPanel(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setUsesFontPanel:", value)
}

func (t_ Text) IsRulerVisible() bool {
	rv := ffi.CallMethod[bool](t_, "isRulerVisible")
	return rv
}

func (t_ Text) SelectedRange() foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "selectedRange")
	return rv
}

func (t_ Text) SetSelectedRange(value foundation.Range) {
	ffi.CallMethod[ffi.Void](t_, "setSelectedRange:", value)
}

func (t_ Text) Font() Font {
	rv := ffi.CallMethod[Font](t_, "font")
	return rv
}

func (t_ Text) SetFont(value IFont) {
	ffi.CallMethod[ffi.Void](t_, "setFont:", value)
}

func (t_ Text) Alignment() TextAlignment {
	rv := ffi.CallMethod[TextAlignment](t_, "alignment")
	return rv
}

func (t_ Text) SetAlignment(value TextAlignment) {
	ffi.CallMethod[ffi.Void](t_, "setAlignment:", value)
}

func (t_ Text) TextColor() Color {
	rv := ffi.CallMethod[Color](t_, "textColor")
	return rv
}

func (t_ Text) SetTextColor(value IColor) {
	ffi.CallMethod[ffi.Void](t_, "setTextColor:", value)
}

func (t_ Text) BaseWritingDirection() WritingDirection {
	rv := ffi.CallMethod[WritingDirection](t_, "baseWritingDirection")
	return rv
}

func (t_ Text) SetBaseWritingDirection(value WritingDirection) {
	ffi.CallMethod[ffi.Void](t_, "setBaseWritingDirection:", value)
}

func (t_ Text) MaxSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](t_, "maxSize")
	return rv
}

func (t_ Text) SetMaxSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](t_, "setMaxSize:", value)
}

func (t_ Text) MinSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](t_, "minSize")
	return rv
}

func (t_ Text) SetMinSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](t_, "setMinSize:", value)
}

func (t_ Text) IsVerticallyResizable() bool {
	rv := ffi.CallMethod[bool](t_, "isVerticallyResizable")
	return rv
}

func (t_ Text) SetVerticallyResizable(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setVerticallyResizable:", value)
}

func (t_ Text) IsHorizontallyResizable() bool {
	rv := ffi.CallMethod[bool](t_, "isHorizontallyResizable")
	return rv
}

func (t_ Text) SetHorizontallyResizable(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setHorizontallyResizable:", value)
}

func (t_ Text) Delegate() TextDelegateWrapper {
	rv := ffi.CallMethod[TextDelegateWrapper](t_, "delegate")
	return rv
}

func (t_ Text) SetDelegate(value TextDelegate) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](t_, "setDelegate:", po)
}

type TextDelegate interface {
	ImplementsTextDidChange() bool
	// optional
	TextDidChange(notification foundation.Notification)
	ImplementsTextShouldBeginEditing() bool
	// optional
	TextShouldBeginEditing(textObject Text) bool
	ImplementsTextDidBeginEditing() bool
	// optional
	TextDidBeginEditing(notification foundation.Notification)
	ImplementsTextShouldEndEditing() bool
	// optional
	TextShouldEndEditing(textObject Text) bool
	ImplementsTextDidEndEditing() bool
	// optional
	TextDidEndEditing(notification foundation.Notification)
}

type TextDelegateImpl struct {
	_TextDidChange          func(notification foundation.Notification)
	_TextShouldBeginEditing func(textObject Text) bool
	_TextDidBeginEditing    func(notification foundation.Notification)
	_TextShouldEndEditing   func(textObject Text) bool
	_TextDidEndEditing      func(notification foundation.Notification)
}

func (di *TextDelegateImpl) ImplementsTextDidChange() bool {
	return di._TextDidChange != nil
}

func (di *TextDelegateImpl) SetTextDidChange(f func(notification foundation.Notification)) {
	di._TextDidChange = f
}

func (di *TextDelegateImpl) TextDidChange(notification foundation.Notification) {
	di._TextDidChange(notification)
}
func (di *TextDelegateImpl) ImplementsTextShouldBeginEditing() bool {
	return di._TextShouldBeginEditing != nil
}

func (di *TextDelegateImpl) SetTextShouldBeginEditing(f func(textObject Text) bool) {
	di._TextShouldBeginEditing = f
}

func (di *TextDelegateImpl) TextShouldBeginEditing(textObject Text) bool {
	return di._TextShouldBeginEditing(textObject)
}
func (di *TextDelegateImpl) ImplementsTextDidBeginEditing() bool {
	return di._TextDidBeginEditing != nil
}

func (di *TextDelegateImpl) SetTextDidBeginEditing(f func(notification foundation.Notification)) {
	di._TextDidBeginEditing = f
}

func (di *TextDelegateImpl) TextDidBeginEditing(notification foundation.Notification) {
	di._TextDidBeginEditing(notification)
}
func (di *TextDelegateImpl) ImplementsTextShouldEndEditing() bool {
	return di._TextShouldEndEditing != nil
}

func (di *TextDelegateImpl) SetTextShouldEndEditing(f func(textObject Text) bool) {
	di._TextShouldEndEditing = f
}

func (di *TextDelegateImpl) TextShouldEndEditing(textObject Text) bool {
	return di._TextShouldEndEditing(textObject)
}
func (di *TextDelegateImpl) ImplementsTextDidEndEditing() bool {
	return di._TextDidEndEditing != nil
}

func (di *TextDelegateImpl) SetTextDidEndEditing(f func(notification foundation.Notification)) {
	di._TextDidEndEditing = f
}

func (di *TextDelegateImpl) TextDidEndEditing(notification foundation.Notification) {
	di._TextDidEndEditing(notification)
}

type TextDelegateWrapper struct {
	objc.Object
}

func (t_ *TextDelegateWrapper) ImplementsTextDidChange() bool {
	return t_.RespondsToSelector(objc.GetSelector("textDidChange:"))
}

func (t_ TextDelegateWrapper) TextDidChange(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](t_, "textDidChange:", notification)
}

func (t_ *TextDelegateWrapper) ImplementsTextShouldBeginEditing() bool {
	return t_.RespondsToSelector(objc.GetSelector("textShouldBeginEditing:"))
}

func (t_ TextDelegateWrapper) TextShouldBeginEditing(textObject IText) bool {
	rv := ffi.CallMethod[bool](t_, "textShouldBeginEditing:", textObject)
	return rv
}

func (t_ *TextDelegateWrapper) ImplementsTextDidBeginEditing() bool {
	return t_.RespondsToSelector(objc.GetSelector("textDidBeginEditing:"))
}

func (t_ TextDelegateWrapper) TextDidBeginEditing(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](t_, "textDidBeginEditing:", notification)
}

func (t_ *TextDelegateWrapper) ImplementsTextShouldEndEditing() bool {
	return t_.RespondsToSelector(objc.GetSelector("textShouldEndEditing:"))
}

func (t_ TextDelegateWrapper) TextShouldEndEditing(textObject IText) bool {
	rv := ffi.CallMethod[bool](t_, "textShouldEndEditing:", textObject)
	return rv
}

func (t_ *TextDelegateWrapper) ImplementsTextDidEndEditing() bool {
	return t_.RespondsToSelector(objc.GetSelector("textDidEndEditing:"))
}

func (t_ TextDelegateWrapper) TextDidEndEditing(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](t_, "textDidEndEditing:", notification)
}

var TextInputContextClass = _TextInputContextClass{objc.GetClass("NSTextInputContext")}

type _TextInputContextClass struct {
	objc.Class
}

type ITextInputContext interface {
	objc.IObject
	Activate()
	Deactivate()
	HandleEvent(event IEvent) bool
	DiscardMarkedText()
	InvalidateCharacterCoordinates()
	AcceptsGlyphInfo() bool
	SetAcceptsGlyphInfo(value bool)
	AllowedInputSourceLocales() []string
	SetAllowedInputSourceLocales(value []string)
	KeyboardInputSources() []TextInputSourceIdentifier
	SelectedKeyboardInputSource() TextInputSourceIdentifier
	SetSelectedKeyboardInputSource(value TextInputSourceIdentifier)
}

type TextInputContext struct {
	objc.Object
}

func MakeTextInputContext(ptr unsafe.Pointer) TextInputContext {
	return TextInputContext{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TextInputContextClass) Alloc() TextInputContext {
	rv := ffi.CallMethod[TextInputContext](tc, "alloc")
	return rv
}

func (t_ TextInputContext) Init() TextInputContext {
	rv := ffi.CallMethod[TextInputContext](t_, "init")
	return rv
}

func (tc _TextInputContextClass) New() TextInputContext {
	rv := ffi.CallMethod[TextInputContext](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextInputContext() TextInputContext {
	return TextInputContextClass.New()
}

func (t_ TextInputContext) Activate() {
	ffi.CallMethod[ffi.Void](t_, "activate")
}

func (t_ TextInputContext) Deactivate() {
	ffi.CallMethod[ffi.Void](t_, "deactivate")
}

func (t_ TextInputContext) HandleEvent(event IEvent) bool {
	rv := ffi.CallMethod[bool](t_, "handleEvent:", event)
	return rv
}

func (t_ TextInputContext) DiscardMarkedText() {
	ffi.CallMethod[ffi.Void](t_, "discardMarkedText")
}

func (t_ TextInputContext) InvalidateCharacterCoordinates() {
	ffi.CallMethod[ffi.Void](t_, "invalidateCharacterCoordinates")
}

func (tc _TextInputContextClass) LocalizedNameForInputSource(inputSourceIdentifier TextInputSourceIdentifier) string {
	rv := ffi.CallMethod[string](tc, "localizedNameForInputSource:", inputSourceIdentifier)
	return rv
}

func (tc _TextInputContextClass) CurrentInputContext() TextInputContext {
	rv := ffi.CallMethod[TextInputContext](tc, "currentInputContext")
	return rv
}

func (t_ TextInputContext) AcceptsGlyphInfo() bool {
	rv := ffi.CallMethod[bool](t_, "acceptsGlyphInfo")
	return rv
}

func (t_ TextInputContext) SetAcceptsGlyphInfo(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAcceptsGlyphInfo:", value)
}

func (t_ TextInputContext) AllowedInputSourceLocales() []string {
	rv := ffi.CallMethod[[]string](t_, "allowedInputSourceLocales")
	return rv
}

func (t_ TextInputContext) SetAllowedInputSourceLocales(value []string) {
	ffi.CallMethod[ffi.Void](t_, "setAllowedInputSourceLocales:", value)
}

func (t_ TextInputContext) KeyboardInputSources() []TextInputSourceIdentifier {
	rv := ffi.CallMethod[[]TextInputSourceIdentifier](t_, "keyboardInputSources")
	return rv
}

func (t_ TextInputContext) SelectedKeyboardInputSource() TextInputSourceIdentifier {
	rv := ffi.CallMethod[TextInputSourceIdentifier](t_, "selectedKeyboardInputSource")
	return rv
}

func (t_ TextInputContext) SetSelectedKeyboardInputSource(value TextInputSourceIdentifier) {
	ffi.CallMethod[ffi.Void](t_, "setSelectedKeyboardInputSource:", value)
}

var TextContainerClass = _TextContainerClass{objc.GetClass("NSTextContainer")}

type _TextContainerClass struct {
	objc.Class
}

type ITextContainer interface {
	objc.IObject
	ReplaceLayoutManager(newLayoutManager ILayoutManager)
	LineFragmentRectForProposedRect_AtIndex_WritingDirection_RemainingRect(proposedRect foundation.Rect, characterIndex uint, baseWritingDirection WritingDirection, remainingRect *foundation.Rect) foundation.Rect
	// deprecated
	LineFragmentRectForProposedRect_SweepDirection_MovementDirection_RemainingRect(proposedRect foundation.Rect, sweepDirection LineSweepDirection, movementDirection LineMovementDirection, remainingRect *foundation.Rect) foundation.Rect
	// deprecated
	ContainsPoint(point foundation.Point) bool
	LayoutManager() LayoutManager
	SetLayoutManager(value ILayoutManager)
	TextLayoutManager() TextLayoutManager
	TextView() TextView
	SetTextView(value ITextView)
	Size() foundation.Size
	SetSize(value foundation.Size)
	ExclusionPaths() []BezierPath
	SetExclusionPaths(value []IBezierPath)
	LineBreakMode() LineBreakMode
	SetLineBreakMode(value LineBreakMode)
	WidthTracksTextView() bool
	SetWidthTracksTextView(value bool)
	HeightTracksTextView() bool
	SetHeightTracksTextView(value bool)
	MaximumNumberOfLines() uint
	SetMaximumNumberOfLines(value uint)
	LineFragmentPadding() float64
	SetLineFragmentPadding(value float64)
	IsSimpleRectangularTextContainer() bool
	// deprecated
	ContainerSize() foundation.Size
	// deprecated
	SetContainerSize(value foundation.Size)
}

type TextContainer struct {
	objc.Object
}

func MakeTextContainer(ptr unsafe.Pointer) TextContainer {
	return TextContainer{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextContainer) InitWithSize(size foundation.Size) TextContainer {
	rv := ffi.CallMethod[TextContainer](t_, "initWithSize:", size)
	return rv
}

func (t_ TextContainer) InitWithContainerSize(aContainerSize foundation.Size) TextContainer {
	rv := ffi.CallMethod[TextContainer](t_, "initWithContainerSize:", aContainerSize)
	return rv
}

func (tc _TextContainerClass) Alloc() TextContainer {
	rv := ffi.CallMethod[TextContainer](tc, "alloc")
	return rv
}

func (t_ TextContainer) Init() TextContainer {
	rv := ffi.CallMethod[TextContainer](t_, "init")
	return rv
}

func (tc _TextContainerClass) New() TextContainer {
	rv := ffi.CallMethod[TextContainer](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextContainer() TextContainer {
	return TextContainerClass.New()
}

func (t_ TextContainer) ReplaceLayoutManager(newLayoutManager ILayoutManager) {
	ffi.CallMethod[ffi.Void](t_, "replaceLayoutManager:", newLayoutManager)
}

func (t_ TextContainer) LineFragmentRectForProposedRect_AtIndex_WritingDirection_RemainingRect(proposedRect foundation.Rect, characterIndex uint, baseWritingDirection WritingDirection, remainingRect *foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](t_, "lineFragmentRectForProposedRect:atIndex:writingDirection:remainingRect:", proposedRect, characterIndex, baseWritingDirection, remainingRect)
	return rv
}

// deprecated
func (t_ TextContainer) LineFragmentRectForProposedRect_SweepDirection_MovementDirection_RemainingRect(proposedRect foundation.Rect, sweepDirection LineSweepDirection, movementDirection LineMovementDirection, remainingRect *foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](t_, "lineFragmentRectForProposedRect:sweepDirection:movementDirection:remainingRect:", proposedRect, sweepDirection, movementDirection, remainingRect)
	return rv
}

// deprecated
func (t_ TextContainer) ContainsPoint(point foundation.Point) bool {
	rv := ffi.CallMethod[bool](t_, "containsPoint:", point)
	return rv
}

func (t_ TextContainer) LayoutManager() LayoutManager {
	rv := ffi.CallMethod[LayoutManager](t_, "layoutManager")
	return rv
}

func (t_ TextContainer) SetLayoutManager(value ILayoutManager) {
	ffi.CallMethod[ffi.Void](t_, "setLayoutManager:", value)
}

func (t_ TextContainer) TextLayoutManager() TextLayoutManager {
	rv := ffi.CallMethod[TextLayoutManager](t_, "textLayoutManager")
	return rv
}

func (t_ TextContainer) TextView() TextView {
	rv := ffi.CallMethod[TextView](t_, "textView")
	return rv
}

func (t_ TextContainer) SetTextView(value ITextView) {
	ffi.CallMethod[ffi.Void](t_, "setTextView:", value)
}

func (t_ TextContainer) Size() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](t_, "size")
	return rv
}

func (t_ TextContainer) SetSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](t_, "setSize:", value)
}

func (t_ TextContainer) ExclusionPaths() []BezierPath {
	rv := ffi.CallMethod[[]BezierPath](t_, "exclusionPaths")
	return rv
}

func (t_ TextContainer) SetExclusionPaths(value []IBezierPath) {
	ffi.CallMethod[ffi.Void](t_, "setExclusionPaths:", value)
}

func (t_ TextContainer) LineBreakMode() LineBreakMode {
	rv := ffi.CallMethod[LineBreakMode](t_, "lineBreakMode")
	return rv
}

func (t_ TextContainer) SetLineBreakMode(value LineBreakMode) {
	ffi.CallMethod[ffi.Void](t_, "setLineBreakMode:", value)
}

func (t_ TextContainer) WidthTracksTextView() bool {
	rv := ffi.CallMethod[bool](t_, "widthTracksTextView")
	return rv
}

func (t_ TextContainer) SetWidthTracksTextView(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setWidthTracksTextView:", value)
}

func (t_ TextContainer) HeightTracksTextView() bool {
	rv := ffi.CallMethod[bool](t_, "heightTracksTextView")
	return rv
}

func (t_ TextContainer) SetHeightTracksTextView(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setHeightTracksTextView:", value)
}

func (t_ TextContainer) MaximumNumberOfLines() uint {
	rv := ffi.CallMethod[uint](t_, "maximumNumberOfLines")
	return rv
}

func (t_ TextContainer) SetMaximumNumberOfLines(value uint) {
	ffi.CallMethod[ffi.Void](t_, "setMaximumNumberOfLines:", value)
}

func (t_ TextContainer) LineFragmentPadding() float64 {
	rv := ffi.CallMethod[float64](t_, "lineFragmentPadding")
	return rv
}

func (t_ TextContainer) SetLineFragmentPadding(value float64) {
	ffi.CallMethod[ffi.Void](t_, "setLineFragmentPadding:", value)
}

func (t_ TextContainer) IsSimpleRectangularTextContainer() bool {
	rv := ffi.CallMethod[bool](t_, "isSimpleRectangularTextContainer")
	return rv
}

// deprecated
func (t_ TextContainer) ContainerSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](t_, "containerSize")
	return rv
}

// deprecated
func (t_ TextContainer) SetContainerSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](t_, "setContainerSize:", value)
}

var TextRangeClass = _TextRangeClass{objc.GetClass("NSTextRange")}

type _TextRangeClass struct {
	objc.Class
}

type ITextRange interface {
	objc.IObject
	IntersectsWithTextRange(textRange ITextRange) bool
	IsEqualToTextRange(textRange ITextRange) bool
	ContainsRange(textRange ITextRange) bool
	IsEmpty() bool
}

type TextRange struct {
	objc.Object
}

func MakeTextRange(ptr unsafe.Pointer) TextRange {
	return TextRange{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextRange) TextRangeByIntersectingWithTextRange(textRange ITextRange) TextRange {
	rv := ffi.CallMethod[TextRange](t_, "textRangeByIntersectingWithTextRange:", textRange)
	return rv
}

func (t_ TextRange) TextRangeByFormingUnionWithTextRange(textRange ITextRange) TextRange {
	rv := ffi.CallMethod[TextRange](t_, "textRangeByFormingUnionWithTextRange:", textRange)
	return rv
}

func (tc _TextRangeClass) Alloc() TextRange {
	rv := ffi.CallMethod[TextRange](tc, "alloc")
	return rv
}

func (t_ TextRange) Init() TextRange {
	rv := ffi.CallMethod[TextRange](t_, "init")
	return rv
}

func (tc _TextRangeClass) New() TextRange {
	rv := ffi.CallMethod[TextRange](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextRange() TextRange {
	return TextRangeClass.New()
}

func (t_ TextRange) IntersectsWithTextRange(textRange ITextRange) bool {
	rv := ffi.CallMethod[bool](t_, "intersectsWithTextRange:", textRange)
	return rv
}

func (t_ TextRange) IsEqualToTextRange(textRange ITextRange) bool {
	rv := ffi.CallMethod[bool](t_, "isEqualToTextRange:", textRange)
	return rv
}

func (t_ TextRange) ContainsRange(textRange ITextRange) bool {
	rv := ffi.CallMethod[bool](t_, "containsRange:", textRange)
	return rv
}

func (t_ TextRange) IsEmpty() bool {
	rv := ffi.CallMethod[bool](t_, "isEmpty")
	return rv
}

var TextElementClass = _TextElementClass{objc.GetClass("NSTextElement")}

type _TextElementClass struct {
	objc.Class
}

type ITextElement interface {
	objc.IObject
	TextContentManager() TextContentManager
	SetTextContentManager(value ITextContentManager)
	ElementRange() TextRange
	SetElementRange(value ITextRange)
	IsRepresentedElement() bool
	ParentElement() TextElement
	ChildElements() []TextElement
}

type TextElement struct {
	objc.Object
}

func MakeTextElement(ptr unsafe.Pointer) TextElement {
	return TextElement{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextElement) InitWithTextContentManager(textContentManager ITextContentManager) TextElement {
	rv := ffi.CallMethod[TextElement](t_, "initWithTextContentManager:", textContentManager)
	return rv
}

func (tc _TextElementClass) Alloc() TextElement {
	rv := ffi.CallMethod[TextElement](tc, "alloc")
	return rv
}

func (t_ TextElement) Init() TextElement {
	rv := ffi.CallMethod[TextElement](t_, "init")
	return rv
}

func (tc _TextElementClass) New() TextElement {
	rv := ffi.CallMethod[TextElement](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextElement() TextElement {
	return TextElementClass.New()
}

func (t_ TextElement) TextContentManager() TextContentManager {
	rv := ffi.CallMethod[TextContentManager](t_, "textContentManager")
	return rv
}

func (t_ TextElement) SetTextContentManager(value ITextContentManager) {
	ffi.CallMethod[ffi.Void](t_, "setTextContentManager:", value)
}

func (t_ TextElement) ElementRange() TextRange {
	rv := ffi.CallMethod[TextRange](t_, "elementRange")
	return rv
}

func (t_ TextElement) SetElementRange(value ITextRange) {
	ffi.CallMethod[ffi.Void](t_, "setElementRange:", value)
}

func (t_ TextElement) IsRepresentedElement() bool {
	rv := ffi.CallMethod[bool](t_, "isRepresentedElement")
	return rv
}

func (t_ TextElement) ParentElement() TextElement {
	rv := ffi.CallMethod[TextElement](t_, "parentElement")
	return rv
}

func (t_ TextElement) ChildElements() []TextElement {
	rv := ffi.CallMethod[[]TextElement](t_, "childElements")
	return rv
}

var TextSelectionClass = _TextSelectionClass{objc.GetClass("NSTextSelection")}

type _TextSelectionClass struct {
	objc.Class
}

type ITextSelection interface {
	objc.IObject
	TextSelectionWithTextRanges(textRanges []ITextRange) TextSelection
	Affinity() TextSelectionAffinity
	AnchorPositionOffset() float64
	SetAnchorPositionOffset(value float64)
	Granularity() TextSelectionGranularity
	IsLogical() bool
	SetLogical(value bool)
	IsTransient() bool
	TextRanges() []TextRange
	TypingAttributes() map[foundation.AttributedStringKey]objc.Object
	SetTypingAttributes(value map[foundation.AttributedStringKey]objc.IObject)
}

type TextSelection struct {
	objc.Object
}

func MakeTextSelection(ptr unsafe.Pointer) TextSelection {
	return TextSelection{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextSelection) InitWithRange_Affinity_Granularity(_range ITextRange, affinity TextSelectionAffinity, granularity TextSelectionGranularity) TextSelection {
	rv := ffi.CallMethod[TextSelection](t_, "initWithRange:affinity:granularity:", _range, affinity, granularity)
	return rv
}

func (t_ TextSelection) InitWithRanges_Affinity_Granularity(textRanges []ITextRange, affinity TextSelectionAffinity, granularity TextSelectionGranularity) TextSelection {
	rv := ffi.CallMethod[TextSelection](t_, "initWithRanges:affinity:granularity:", textRanges, affinity, granularity)
	return rv
}

func (tc _TextSelectionClass) Alloc() TextSelection {
	rv := ffi.CallMethod[TextSelection](tc, "alloc")
	return rv
}

func (t_ TextSelection) Init() TextSelection {
	rv := ffi.CallMethod[TextSelection](t_, "init")
	return rv
}

func (tc _TextSelectionClass) New() TextSelection {
	rv := ffi.CallMethod[TextSelection](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextSelection() TextSelection {
	return TextSelectionClass.New()
}

func (t_ TextSelection) TextSelectionWithTextRanges(textRanges []ITextRange) TextSelection {
	rv := ffi.CallMethod[TextSelection](t_, "textSelectionWithTextRanges:", textRanges)
	return rv
}

func (t_ TextSelection) Affinity() TextSelectionAffinity {
	rv := ffi.CallMethod[TextSelectionAffinity](t_, "affinity")
	return rv
}

func (t_ TextSelection) AnchorPositionOffset() float64 {
	rv := ffi.CallMethod[float64](t_, "anchorPositionOffset")
	return rv
}

func (t_ TextSelection) SetAnchorPositionOffset(value float64) {
	ffi.CallMethod[ffi.Void](t_, "setAnchorPositionOffset:", value)
}

func (t_ TextSelection) Granularity() TextSelectionGranularity {
	rv := ffi.CallMethod[TextSelectionGranularity](t_, "granularity")
	return rv
}

func (t_ TextSelection) IsLogical() bool {
	rv := ffi.CallMethod[bool](t_, "isLogical")
	return rv
}

func (t_ TextSelection) SetLogical(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setLogical:", value)
}

func (t_ TextSelection) IsTransient() bool {
	rv := ffi.CallMethod[bool](t_, "isTransient")
	return rv
}

func (t_ TextSelection) TextRanges() []TextRange {
	rv := ffi.CallMethod[[]TextRange](t_, "textRanges")
	return rv
}

func (t_ TextSelection) TypingAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := ffi.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, "typingAttributes")
	return rv
}

func (t_ TextSelection) SetTypingAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "setTypingAttributes:", value)
}

var TextSelectionNavigationClass = _TextSelectionNavigationClass{objc.GetClass("NSTextSelectionNavigation")}

type _TextSelectionNavigationClass struct {
	objc.Class
}

type ITextSelectionNavigation interface {
	objc.IObject
	TextSelectionForSelectionGranularity_EnclosingTextSelection(selectionGranularity TextSelectionGranularity, textSelection ITextSelection) TextSelection
	DestinationSelectionForTextSelection_Direction_Destination_Extending_Confined(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, extending bool, confined bool) TextSelection
	FlushLayoutCache()
	DeletionRangesForTextSelection_Direction_Destination_AllowsDecomposition(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, allowsDecomposition bool) []TextRange
	AllowsNonContiguousRanges() bool
	SetAllowsNonContiguousRanges(value bool)
	RotatesCoordinateSystemForLayoutOrientation() bool
	SetRotatesCoordinateSystemForLayoutOrientation(value bool)
}

type TextSelectionNavigation struct {
	objc.Object
}

func MakeTextSelectionNavigation(ptr unsafe.Pointer) TextSelectionNavigation {
	return TextSelectionNavigation{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TextSelectionNavigationClass) Alloc() TextSelectionNavigation {
	rv := ffi.CallMethod[TextSelectionNavigation](tc, "alloc")
	return rv
}

func (t_ TextSelectionNavigation) Init() TextSelectionNavigation {
	rv := ffi.CallMethod[TextSelectionNavigation](t_, "init")
	return rv
}

func (tc _TextSelectionNavigationClass) New() TextSelectionNavigation {
	rv := ffi.CallMethod[TextSelectionNavigation](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextSelectionNavigation() TextSelectionNavigation {
	return TextSelectionNavigationClass.New()
}

func (t_ TextSelectionNavigation) TextSelectionForSelectionGranularity_EnclosingTextSelection(selectionGranularity TextSelectionGranularity, textSelection ITextSelection) TextSelection {
	rv := ffi.CallMethod[TextSelection](t_, "textSelectionForSelectionGranularity:enclosingTextSelection:", selectionGranularity, textSelection)
	return rv
}

func (t_ TextSelectionNavigation) DestinationSelectionForTextSelection_Direction_Destination_Extending_Confined(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, extending bool, confined bool) TextSelection {
	rv := ffi.CallMethod[TextSelection](t_, "destinationSelectionForTextSelection:direction:destination:extending:confined:", textSelection, direction, destination, extending, confined)
	return rv
}

func (t_ TextSelectionNavigation) FlushLayoutCache() {
	ffi.CallMethod[ffi.Void](t_, "flushLayoutCache")
}

func (t_ TextSelectionNavigation) DeletionRangesForTextSelection_Direction_Destination_AllowsDecomposition(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, allowsDecomposition bool) []TextRange {
	rv := ffi.CallMethod[[]TextRange](t_, "deletionRangesForTextSelection:direction:destination:allowsDecomposition:", textSelection, direction, destination, allowsDecomposition)
	return rv
}

func (t_ TextSelectionNavigation) AllowsNonContiguousRanges() bool {
	rv := ffi.CallMethod[bool](t_, "allowsNonContiguousRanges")
	return rv
}

func (t_ TextSelectionNavigation) SetAllowsNonContiguousRanges(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAllowsNonContiguousRanges:", value)
}

func (t_ TextSelectionNavigation) RotatesCoordinateSystemForLayoutOrientation() bool {
	rv := ffi.CallMethod[bool](t_, "rotatesCoordinateSystemForLayoutOrientation")
	return rv
}

func (t_ TextSelectionNavigation) SetRotatesCoordinateSystemForLayoutOrientation(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setRotatesCoordinateSystemForLayoutOrientation:", value)
}

var TextContentManagerClass = _TextContentManagerClass{objc.GetClass("NSTextContentManager")}

type _TextContentManagerClass struct {
	objc.Class
}

type ITextContentManager interface {
	objc.IObject
	PerformEditingTransactionUsingBlock(transaction func())
	RecordEditActionInRange_NewTextRange(originalTextRange ITextRange, newTextRange ITextRange)
	AddTextLayoutManager(textLayoutManager ITextLayoutManager)
	RemoveTextLayoutManager(textLayoutManager ITextLayoutManager)
	SynchronizeTextLayoutManagers(completionHandler func(error foundation.Error))
	TextElementsForRange(_range ITextRange) []TextElement
	AutomaticallySynchronizesToBackingStore() bool
	SetAutomaticallySynchronizesToBackingStore(value bool)
	HasEditingTransaction() bool
	PrimaryTextLayoutManager() TextLayoutManager
	SetPrimaryTextLayoutManager(value ITextLayoutManager)
	TextLayoutManagers() []TextLayoutManager
	AutomaticallySynchronizesTextLayoutManagers() bool
	SetAutomaticallySynchronizesTextLayoutManagers(value bool)
}

type TextContentManager struct {
	objc.Object
}

func MakeTextContentManager(ptr unsafe.Pointer) TextContentManager {
	return TextContentManager{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextContentManager) Init() TextContentManager {
	rv := ffi.CallMethod[TextContentManager](t_, "init")
	return rv
}

func (tc _TextContentManagerClass) Alloc() TextContentManager {
	rv := ffi.CallMethod[TextContentManager](tc, "alloc")
	return rv
}

func (tc _TextContentManagerClass) New() TextContentManager {
	rv := ffi.CallMethod[TextContentManager](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextContentManager() TextContentManager {
	return TextContentManagerClass.New()
}

func (t_ TextContentManager) PerformEditingTransactionUsingBlock(transaction func()) {
	ffi.CallMethod[ffi.Void](t_, "performEditingTransactionUsingBlock:", transaction)
}

func (t_ TextContentManager) RecordEditActionInRange_NewTextRange(originalTextRange ITextRange, newTextRange ITextRange) {
	ffi.CallMethod[ffi.Void](t_, "recordEditActionInRange:newTextRange:", originalTextRange, newTextRange)
}

func (t_ TextContentManager) AddTextLayoutManager(textLayoutManager ITextLayoutManager) {
	ffi.CallMethod[ffi.Void](t_, "addTextLayoutManager:", textLayoutManager)
}

func (t_ TextContentManager) RemoveTextLayoutManager(textLayoutManager ITextLayoutManager) {
	ffi.CallMethod[ffi.Void](t_, "removeTextLayoutManager:", textLayoutManager)
}

func (t_ TextContentManager) SynchronizeTextLayoutManagers(completionHandler func(error foundation.Error)) {
	ffi.CallMethod[ffi.Void](t_, "synchronizeTextLayoutManagers:", completionHandler)
}

func (t_ TextContentManager) TextElementsForRange(_range ITextRange) []TextElement {
	rv := ffi.CallMethod[[]TextElement](t_, "textElementsForRange:", _range)
	return rv
}

func (t_ TextContentManager) AutomaticallySynchronizesToBackingStore() bool {
	rv := ffi.CallMethod[bool](t_, "automaticallySynchronizesToBackingStore")
	return rv
}

func (t_ TextContentManager) SetAutomaticallySynchronizesToBackingStore(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAutomaticallySynchronizesToBackingStore:", value)
}

func (t_ TextContentManager) HasEditingTransaction() bool {
	rv := ffi.CallMethod[bool](t_, "hasEditingTransaction")
	return rv
}

func (t_ TextContentManager) PrimaryTextLayoutManager() TextLayoutManager {
	rv := ffi.CallMethod[TextLayoutManager](t_, "primaryTextLayoutManager")
	return rv
}

func (t_ TextContentManager) SetPrimaryTextLayoutManager(value ITextLayoutManager) {
	ffi.CallMethod[ffi.Void](t_, "setPrimaryTextLayoutManager:", value)
}

func (t_ TextContentManager) TextLayoutManagers() []TextLayoutManager {
	rv := ffi.CallMethod[[]TextLayoutManager](t_, "textLayoutManagers")
	return rv
}

func (t_ TextContentManager) AutomaticallySynchronizesTextLayoutManagers() bool {
	rv := ffi.CallMethod[bool](t_, "automaticallySynchronizesTextLayoutManagers")
	return rv
}

func (t_ TextContentManager) SetAutomaticallySynchronizesTextLayoutManagers(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAutomaticallySynchronizesTextLayoutManagers:", value)
}
