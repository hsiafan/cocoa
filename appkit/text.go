// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TextClass = _TextClass{objc.GetClass("NSText")}

type _TextClass struct {
	objc.Class
}

type IText interface {
	IView
	ToggleRuler(sender objc.IObject)
	ReplaceCharactersInRange_WithRTF(range_ foundation.Range, rtfData []byte)
	ReplaceCharactersInRange_WithRTFD(range_ foundation.Range, rtfdData []byte)
	ReplaceCharactersInRange_WithString(range_ foundation.Range, string_ string)
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
	SetFont_Range(font IFont, range_ foundation.Range)
	AlignCenter(sender objc.IObject)
	AlignLeft(sender objc.IObject)
	AlignRight(sender objc.IObject)
	SetTextColor_Range(color IColor, range_ foundation.Range)
	Superscript(sender objc.IObject)
	Subscript(sender objc.IObject)
	Unscript(sender objc.IObject)
	Underline(sender objc.IObject)
	ReadRTFDFromFile(path string) bool
	WriteRTFDToFile_Atomically(path string, flag bool) bool
	RTFDFromRange(range_ foundation.Range) []byte
	RTFFromRange(range_ foundation.Range) []byte
	CheckSpelling(sender objc.IObject)
	ShowGuessPanel(sender objc.IObject)
	SizeToFit()
	ScrollRangeToVisible(range_ foundation.Range)
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
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
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
	rv := objc.CallMethod[Text](t_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (t_ Text) Init() Text {
	rv := objc.CallMethod[Text](t_, objc.SEL("init"))
	return rv
}

func (tc _TextClass) Alloc() Text {
	rv := objc.CallMethod[Text](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TextClass) New() Text {
	rv := objc.CallMethod[Text](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewText() Text {
	return TextClass.New()
}

func (t_ Text) ToggleRuler(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("toggleRuler:"), objc.ExtractPtr(sender))
}

func (t_ Text) ReplaceCharactersInRange_WithRTF(range_ foundation.Range, rtfData []byte) {
	objc.CallMethod[objc.Void](t_, objc.SEL("replaceCharactersInRange:withRTF:"), range_, rtfData)
}

func (t_ Text) ReplaceCharactersInRange_WithRTFD(range_ foundation.Range, rtfdData []byte) {
	objc.CallMethod[objc.Void](t_, objc.SEL("replaceCharactersInRange:withRTFD:"), range_, rtfdData)
}

func (t_ Text) ReplaceCharactersInRange_WithString(range_ foundation.Range, string_ string) {
	objc.CallMethod[objc.Void](t_, objc.SEL("replaceCharactersInRange:withString:"), range_, string_)
}

func (t_ Text) SelectAll(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("selectAll:"), objc.ExtractPtr(sender))
}

func (t_ Text) Copy(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("copy:"), objc.ExtractPtr(sender))
}

func (t_ Text) Cut(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("cut:"), objc.ExtractPtr(sender))
}

func (t_ Text) Paste(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("paste:"), objc.ExtractPtr(sender))
}

func (t_ Text) CopyFont(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("copyFont:"), objc.ExtractPtr(sender))
}

func (t_ Text) PasteFont(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("pasteFont:"), objc.ExtractPtr(sender))
}

func (t_ Text) CopyRuler(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("copyRuler:"), objc.ExtractPtr(sender))
}

func (t_ Text) PasteRuler(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("pasteRuler:"), objc.ExtractPtr(sender))
}

func (t_ Text) Delete(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("delete:"), objc.ExtractPtr(sender))
}

func (t_ Text) ChangeFont(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("changeFont:"), objc.ExtractPtr(sender))
}

func (t_ Text) SetFont_Range(font IFont, range_ foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setFont:range:"), objc.ExtractPtr(font), range_)
}

func (t_ Text) AlignCenter(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("alignCenter:"), objc.ExtractPtr(sender))
}

func (t_ Text) AlignLeft(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("alignLeft:"), objc.ExtractPtr(sender))
}

func (t_ Text) AlignRight(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("alignRight:"), objc.ExtractPtr(sender))
}

func (t_ Text) SetTextColor_Range(color IColor, range_ foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTextColor:range:"), objc.ExtractPtr(color), range_)
}

func (t_ Text) Superscript(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("superscript:"), objc.ExtractPtr(sender))
}

func (t_ Text) Subscript(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("subscript:"), objc.ExtractPtr(sender))
}

func (t_ Text) Unscript(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("unscript:"), objc.ExtractPtr(sender))
}

func (t_ Text) Underline(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("underline:"), objc.ExtractPtr(sender))
}

func (t_ Text) ReadRTFDFromFile(path string) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("readRTFDFromFile:"), path)
	return rv
}

func (t_ Text) WriteRTFDToFile_Atomically(path string, flag bool) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("writeRTFDToFile:atomically:"), path, flag)
	return rv
}

func (t_ Text) RTFDFromRange(range_ foundation.Range) []byte {
	rv := objc.CallMethod[[]byte](t_, objc.SEL("RTFDFromRange:"), range_)
	return rv
}

func (t_ Text) RTFFromRange(range_ foundation.Range) []byte {
	rv := objc.CallMethod[[]byte](t_, objc.SEL("RTFFromRange:"), range_)
	return rv
}

func (t_ Text) CheckSpelling(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("checkSpelling:"), objc.ExtractPtr(sender))
}

func (t_ Text) ShowGuessPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("showGuessPanel:"), objc.ExtractPtr(sender))
}

func (t_ Text) SizeToFit() {
	objc.CallMethod[objc.Void](t_, objc.SEL("sizeToFit"))
}

func (t_ Text) ScrollRangeToVisible(range_ foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.SEL("scrollRangeToVisible:"), range_)
}

func (t_ Text) String() string {
	rv := objc.CallMethod[string](t_, objc.SEL("string"))
	return rv
}

func (t_ Text) SetString(value string) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setString:"), value)
}

func (t_ Text) BackgroundColor() Color {
	rv := objc.CallMethod[Color](t_, objc.SEL("backgroundColor"))
	return rv
}

func (t_ Text) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (t_ Text) DrawsBackground() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("drawsBackground"))
	return rv
}

func (t_ Text) SetDrawsBackground(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDrawsBackground:"), value)
}

func (t_ Text) IsEditable() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isEditable"))
	return rv
}

func (t_ Text) SetEditable(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setEditable:"), value)
}

func (t_ Text) IsSelectable() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isSelectable"))
	return rv
}

func (t_ Text) SetSelectable(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSelectable:"), value)
}

func (t_ Text) IsFieldEditor() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isFieldEditor"))
	return rv
}

func (t_ Text) SetFieldEditor(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setFieldEditor:"), value)
}

func (t_ Text) IsRichText() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isRichText"))
	return rv
}

func (t_ Text) SetRichText(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setRichText:"), value)
}

func (t_ Text) ImportsGraphics() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("importsGraphics"))
	return rv
}

func (t_ Text) SetImportsGraphics(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setImportsGraphics:"), value)
}

func (t_ Text) UsesFontPanel() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("usesFontPanel"))
	return rv
}

func (t_ Text) SetUsesFontPanel(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setUsesFontPanel:"), value)
}

func (t_ Text) IsRulerVisible() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isRulerVisible"))
	return rv
}

func (t_ Text) SelectedRange() foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.SEL("selectedRange"))
	return rv
}

func (t_ Text) SetSelectedRange(value foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSelectedRange:"), value)
}

func (t_ Text) Font() Font {
	rv := objc.CallMethod[Font](t_, objc.SEL("font"))
	return rv
}

func (t_ Text) SetFont(value IFont) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setFont:"), objc.ExtractPtr(value))
}

func (t_ Text) Alignment() TextAlignment {
	rv := objc.CallMethod[TextAlignment](t_, objc.SEL("alignment"))
	return rv
}

func (t_ Text) SetAlignment(value TextAlignment) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAlignment:"), value)
}

func (t_ Text) TextColor() Color {
	rv := objc.CallMethod[Color](t_, objc.SEL("textColor"))
	return rv
}

func (t_ Text) SetTextColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTextColor:"), objc.ExtractPtr(value))
}

func (t_ Text) BaseWritingDirection() WritingDirection {
	rv := objc.CallMethod[WritingDirection](t_, objc.SEL("baseWritingDirection"))
	return rv
}

func (t_ Text) SetBaseWritingDirection(value WritingDirection) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setBaseWritingDirection:"), value)
}

func (t_ Text) MaxSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.SEL("maxSize"))
	return rv
}

func (t_ Text) SetMaxSize(value foundation.Size) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setMaxSize:"), value)
}

func (t_ Text) MinSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.SEL("minSize"))
	return rv
}

func (t_ Text) SetMinSize(value foundation.Size) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setMinSize:"), value)
}

func (t_ Text) IsVerticallyResizable() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isVerticallyResizable"))
	return rv
}

func (t_ Text) SetVerticallyResizable(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setVerticallyResizable:"), value)
}

func (t_ Text) IsHorizontallyResizable() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isHorizontallyResizable"))
	return rv
}

func (t_ Text) SetHorizontallyResizable(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setHorizontallyResizable:"), value)
}

// weak property
func (t_ Text) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.SEL("delegate"))
	return rv
}

// weak property
func (t_ Text) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}
