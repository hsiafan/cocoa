// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ParagraphStyleClass = _ParagraphStyleClass{objc.GetClass("NSParagraphStyle")}

type _ParagraphStyleClass struct {
	objc.Class
}

type IParagraphStyle interface {
	objc.IObject
	Alignment() TextAlignment
	FirstLineHeadIndent() float64
	HeadIndent() float64
	TailIndent() float64
	LineHeightMultiple() float64
	MaximumLineHeight() float64
	MinimumLineHeight() float64
	LineSpacing() float64
	ParagraphSpacing() float64
	ParagraphSpacingBefore() float64
	TabStops() []TextTab
	DefaultTabInterval() float64
	TextBlocks() []TextBlock
	TextLists() []TextList
	LineBreakMode() LineBreakMode
	LineBreakStrategy() LineBreakStrategy
	HyphenationFactor() float32
	UsesDefaultHyphenation() bool
	TighteningFactorForTruncation() float32
	AllowsDefaultTighteningForTruncation() bool
	HeaderLevel() int
	BaseWritingDirection() WritingDirection
}

type ParagraphStyle struct {
	objc.Object
}

func MakeParagraphStyle(ptr unsafe.Pointer) ParagraphStyle {
	return ParagraphStyle{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _ParagraphStyleClass) Alloc() ParagraphStyle {
	rv := ffi.CallMethod[ParagraphStyle](pc, "alloc")
	return rv
}

func (p_ ParagraphStyle) Init() ParagraphStyle {
	rv := ffi.CallMethod[ParagraphStyle](p_, "init")
	rv.Autorelease()
	return rv
}

func (pc _ParagraphStyleClass) New() ParagraphStyle {
	rv := ffi.CallMethod[ParagraphStyle](pc, "new")
	rv.Autorelease()
	return rv
}

func NewParagraphStyle() ParagraphStyle {
	return ParagraphStyleClass.New()
}

func (pc _ParagraphStyleClass) DefaultWritingDirectionForLanguage(languageName string) WritingDirection {
	rv := ffi.CallMethod[WritingDirection](pc, "defaultWritingDirectionForLanguage:", languageName)
	return rv
}

func (pc _ParagraphStyleClass) DefaultParagraphStyle() ParagraphStyle {
	rv := ffi.CallMethod[ParagraphStyle](pc, "defaultParagraphStyle")
	return rv
}

func (p_ ParagraphStyle) Alignment() TextAlignment {
	rv := ffi.CallMethod[TextAlignment](p_, "alignment")
	return rv
}

func (p_ ParagraphStyle) FirstLineHeadIndent() float64 {
	rv := ffi.CallMethod[float64](p_, "firstLineHeadIndent")
	return rv
}

func (p_ ParagraphStyle) HeadIndent() float64 {
	rv := ffi.CallMethod[float64](p_, "headIndent")
	return rv
}

func (p_ ParagraphStyle) TailIndent() float64 {
	rv := ffi.CallMethod[float64](p_, "tailIndent")
	return rv
}

func (p_ ParagraphStyle) LineHeightMultiple() float64 {
	rv := ffi.CallMethod[float64](p_, "lineHeightMultiple")
	return rv
}

func (p_ ParagraphStyle) MaximumLineHeight() float64 {
	rv := ffi.CallMethod[float64](p_, "maximumLineHeight")
	return rv
}

func (p_ ParagraphStyle) MinimumLineHeight() float64 {
	rv := ffi.CallMethod[float64](p_, "minimumLineHeight")
	return rv
}

func (p_ ParagraphStyle) LineSpacing() float64 {
	rv := ffi.CallMethod[float64](p_, "lineSpacing")
	return rv
}

func (p_ ParagraphStyle) ParagraphSpacing() float64 {
	rv := ffi.CallMethod[float64](p_, "paragraphSpacing")
	return rv
}

func (p_ ParagraphStyle) ParagraphSpacingBefore() float64 {
	rv := ffi.CallMethod[float64](p_, "paragraphSpacingBefore")
	return rv
}

func (p_ ParagraphStyle) TabStops() []TextTab {
	rv := ffi.CallMethod[[]TextTab](p_, "tabStops")
	return rv
}

func (p_ ParagraphStyle) DefaultTabInterval() float64 {
	rv := ffi.CallMethod[float64](p_, "defaultTabInterval")
	return rv
}

func (p_ ParagraphStyle) TextBlocks() []TextBlock {
	rv := ffi.CallMethod[[]TextBlock](p_, "textBlocks")
	return rv
}

func (p_ ParagraphStyle) TextLists() []TextList {
	rv := ffi.CallMethod[[]TextList](p_, "textLists")
	return rv
}

func (p_ ParagraphStyle) LineBreakMode() LineBreakMode {
	rv := ffi.CallMethod[LineBreakMode](p_, "lineBreakMode")
	return rv
}

func (p_ ParagraphStyle) LineBreakStrategy() LineBreakStrategy {
	rv := ffi.CallMethod[LineBreakStrategy](p_, "lineBreakStrategy")
	return rv
}

func (p_ ParagraphStyle) HyphenationFactor() float32 {
	rv := ffi.CallMethod[float32](p_, "hyphenationFactor")
	return rv
}

func (p_ ParagraphStyle) UsesDefaultHyphenation() bool {
	rv := ffi.CallMethod[bool](p_, "usesDefaultHyphenation")
	return rv
}

func (p_ ParagraphStyle) TighteningFactorForTruncation() float32 {
	rv := ffi.CallMethod[float32](p_, "tighteningFactorForTruncation")
	return rv
}

func (p_ ParagraphStyle) AllowsDefaultTighteningForTruncation() bool {
	rv := ffi.CallMethod[bool](p_, "allowsDefaultTighteningForTruncation")
	return rv
}

func (p_ ParagraphStyle) HeaderLevel() int {
	rv := ffi.CallMethod[int](p_, "headerLevel")
	return rv
}

func (p_ ParagraphStyle) BaseWritingDirection() WritingDirection {
	rv := ffi.CallMethod[WritingDirection](p_, "baseWritingDirection")
	return rv
}

var MutableParagraphStyleClass = _MutableParagraphStyleClass{objc.GetClass("NSMutableParagraphStyle")}

type _MutableParagraphStyleClass struct {
	objc.Class
}

type IMutableParagraphStyle interface {
	IParagraphStyle
	SetParagraphStyle(obj IParagraphStyle)
	AddTabStop(anObject ITextTab)
	RemoveTabStop(anObject ITextTab)
	SetAlignment(value TextAlignment)
	SetFirstLineHeadIndent(value float64)
	SetHeadIndent(value float64)
	SetTailIndent(value float64)
	SetLineHeightMultiple(value float64)
	SetMaximumLineHeight(value float64)
	SetMinimumLineHeight(value float64)
	SetLineSpacing(value float64)
	SetParagraphSpacing(value float64)
	SetParagraphSpacingBefore(value float64)
	SetBaseWritingDirection(value WritingDirection)
	SetTabStops(value []ITextTab)
	SetDefaultTabInterval(value float64)
	SetTextBlocks(value []ITextBlock)
	SetTextLists(value []ITextList)
	SetLineBreakMode(value LineBreakMode)
	SetLineBreakStrategy(value LineBreakStrategy)
	SetHyphenationFactor(value float32)
	SetUsesDefaultHyphenation(value bool)
	SetTighteningFactorForTruncation(value float32)
	SetAllowsDefaultTighteningForTruncation(value bool)
	SetHeaderLevel(value int)
}

type MutableParagraphStyle struct {
	ParagraphStyle
}

func MakeMutableParagraphStyle(ptr unsafe.Pointer) MutableParagraphStyle {
	return MutableParagraphStyle{
		ParagraphStyle: MakeParagraphStyle(ptr),
	}
}

func (mc _MutableParagraphStyleClass) Alloc() MutableParagraphStyle {
	rv := ffi.CallMethod[MutableParagraphStyle](mc, "alloc")
	return rv
}

func (m_ MutableParagraphStyle) Init() MutableParagraphStyle {
	rv := ffi.CallMethod[MutableParagraphStyle](m_, "init")
	rv.Autorelease()
	return rv
}

func (mc _MutableParagraphStyleClass) New() MutableParagraphStyle {
	rv := ffi.CallMethod[MutableParagraphStyle](mc, "new")
	rv.Autorelease()
	return rv
}

func NewMutableParagraphStyle() MutableParagraphStyle {
	return MutableParagraphStyleClass.New()
}

func (m_ MutableParagraphStyle) SetParagraphStyle(obj IParagraphStyle) {
	ffi.CallMethod[ffi.Void](m_, "setParagraphStyle:", obj)
}

func (m_ MutableParagraphStyle) AddTabStop(anObject ITextTab) {
	ffi.CallMethod[ffi.Void](m_, "addTabStop:", anObject)
}

func (m_ MutableParagraphStyle) RemoveTabStop(anObject ITextTab) {
	ffi.CallMethod[ffi.Void](m_, "removeTabStop:", anObject)
}

func (m_ MutableParagraphStyle) SetAlignment(value TextAlignment) {
	ffi.CallMethod[ffi.Void](m_, "setAlignment:", value)
}

func (m_ MutableParagraphStyle) SetFirstLineHeadIndent(value float64) {
	ffi.CallMethod[ffi.Void](m_, "setFirstLineHeadIndent:", value)
}

func (m_ MutableParagraphStyle) SetHeadIndent(value float64) {
	ffi.CallMethod[ffi.Void](m_, "setHeadIndent:", value)
}

func (m_ MutableParagraphStyle) SetTailIndent(value float64) {
	ffi.CallMethod[ffi.Void](m_, "setTailIndent:", value)
}

func (m_ MutableParagraphStyle) SetLineHeightMultiple(value float64) {
	ffi.CallMethod[ffi.Void](m_, "setLineHeightMultiple:", value)
}

func (m_ MutableParagraphStyle) SetMaximumLineHeight(value float64) {
	ffi.CallMethod[ffi.Void](m_, "setMaximumLineHeight:", value)
}

func (m_ MutableParagraphStyle) SetMinimumLineHeight(value float64) {
	ffi.CallMethod[ffi.Void](m_, "setMinimumLineHeight:", value)
}

func (m_ MutableParagraphStyle) SetLineSpacing(value float64) {
	ffi.CallMethod[ffi.Void](m_, "setLineSpacing:", value)
}

func (m_ MutableParagraphStyle) SetParagraphSpacing(value float64) {
	ffi.CallMethod[ffi.Void](m_, "setParagraphSpacing:", value)
}

func (m_ MutableParagraphStyle) SetParagraphSpacingBefore(value float64) {
	ffi.CallMethod[ffi.Void](m_, "setParagraphSpacingBefore:", value)
}

func (m_ MutableParagraphStyle) SetBaseWritingDirection(value WritingDirection) {
	ffi.CallMethod[ffi.Void](m_, "setBaseWritingDirection:", value)
}

func (m_ MutableParagraphStyle) SetTabStops(value []ITextTab) {
	ffi.CallMethod[ffi.Void](m_, "setTabStops:", value)
}

func (m_ MutableParagraphStyle) SetDefaultTabInterval(value float64) {
	ffi.CallMethod[ffi.Void](m_, "setDefaultTabInterval:", value)
}

func (m_ MutableParagraphStyle) SetTextBlocks(value []ITextBlock) {
	ffi.CallMethod[ffi.Void](m_, "setTextBlocks:", value)
}

func (m_ MutableParagraphStyle) SetTextLists(value []ITextList) {
	ffi.CallMethod[ffi.Void](m_, "setTextLists:", value)
}

func (m_ MutableParagraphStyle) SetLineBreakMode(value LineBreakMode) {
	ffi.CallMethod[ffi.Void](m_, "setLineBreakMode:", value)
}

func (m_ MutableParagraphStyle) SetLineBreakStrategy(value LineBreakStrategy) {
	ffi.CallMethod[ffi.Void](m_, "setLineBreakStrategy:", value)
}

func (m_ MutableParagraphStyle) SetHyphenationFactor(value float32) {
	ffi.CallMethod[ffi.Void](m_, "setHyphenationFactor:", value)
}

func (m_ MutableParagraphStyle) SetUsesDefaultHyphenation(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setUsesDefaultHyphenation:", value)
}

func (m_ MutableParagraphStyle) SetTighteningFactorForTruncation(value float32) {
	ffi.CallMethod[ffi.Void](m_, "setTighteningFactorForTruncation:", value)
}

func (m_ MutableParagraphStyle) SetAllowsDefaultTighteningForTruncation(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setAllowsDefaultTighteningForTruncation:", value)
}

func (m_ MutableParagraphStyle) SetHeaderLevel(value int) {
	ffi.CallMethod[ffi.Void](m_, "setHeaderLevel:", value)
}

var TextTabClass = _TextTabClass{objc.GetClass("NSTextTab")}

type _TextTabClass struct {
	objc.Class
}

type ITextTab interface {
	objc.IObject
	Location() float64
	Alignment() TextAlignment
	Options() map[TextTabOptionKey]objc.Object
	// deprecated
	TabStopType() TextTabType
}

type TextTab struct {
	objc.Object
}

func MakeTextTab(ptr unsafe.Pointer) TextTab {
	return TextTab{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextTab) InitWithTextAlignment_Location_Options(alignment TextAlignment, loc float64, options map[TextTabOptionKey]objc.IObject) TextTab {
	rv := ffi.CallMethod[TextTab](t_, "initWithTextAlignment:location:options:", alignment, loc, options)
	rv.Autorelease()
	return rv
}

func (t_ TextTab) InitWithType_Location(_type TextTabType, loc float64) TextTab {
	rv := ffi.CallMethod[TextTab](t_, "initWithType:location:", _type, loc)
	rv.Autorelease()
	return rv
}

func (tc _TextTabClass) Alloc() TextTab {
	rv := ffi.CallMethod[TextTab](tc, "alloc")
	return rv
}

func (t_ TextTab) Init() TextTab {
	rv := ffi.CallMethod[TextTab](t_, "init")
	rv.Autorelease()
	return rv
}

func (tc _TextTabClass) New() TextTab {
	rv := ffi.CallMethod[TextTab](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextTab() TextTab {
	return TextTabClass.New()
}

func (tc _TextTabClass) ColumnTerminatorsForLocale(aLocale foundation.ILocale) foundation.CharacterSet {
	rv := ffi.CallMethod[foundation.CharacterSet](tc, "columnTerminatorsForLocale:", aLocale)
	return rv
}

func (t_ TextTab) Location() float64 {
	rv := ffi.CallMethod[float64](t_, "location")
	return rv
}

func (t_ TextTab) Alignment() TextAlignment {
	rv := ffi.CallMethod[TextAlignment](t_, "alignment")
	return rv
}

func (t_ TextTab) Options() map[TextTabOptionKey]objc.Object {
	rv := ffi.CallMethod[map[TextTabOptionKey]objc.Object](t_, "options")
	return rv
}

// deprecated
func (t_ TextTab) TabStopType() TextTabType {
	rv := ffi.CallMethod[TextTabType](t_, "tabStopType")
	return rv
}

var TextListClass = _TextListClass{objc.GetClass("NSTextList")}

type _TextListClass struct {
	objc.Class
}

type ITextList interface {
	objc.IObject
	MarkerForItemNumber(itemNumber int) string
	MarkerFormat() TextListMarkerFormat
	IsOrdered() bool
	ListOptions() TextListOptions
	StartingItemNumber() int
	SetStartingItemNumber(value int)
}

type TextList struct {
	objc.Object
}

func MakeTextList(ptr unsafe.Pointer) TextList {
	return TextList{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextList) InitWithMarkerFormat_Options(markerFormat TextListMarkerFormat, options uint) TextList {
	rv := ffi.CallMethod[TextList](t_, "initWithMarkerFormat:options:", markerFormat, options)
	rv.Autorelease()
	return rv
}

func (t_ TextList) InitWithMarkerFormat_Options_StartingItemNumber(markerFormat TextListMarkerFormat, options TextListOptions, startingItemNumber int) TextList {
	rv := ffi.CallMethod[TextList](t_, "initWithMarkerFormat:options:startingItemNumber:", markerFormat, options, startingItemNumber)
	rv.Autorelease()
	return rv
}

func (tc _TextListClass) Alloc() TextList {
	rv := ffi.CallMethod[TextList](tc, "alloc")
	return rv
}

func (t_ TextList) Init() TextList {
	rv := ffi.CallMethod[TextList](t_, "init")
	rv.Autorelease()
	return rv
}

func (tc _TextListClass) New() TextList {
	rv := ffi.CallMethod[TextList](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextList() TextList {
	return TextListClass.New()
}

func (t_ TextList) MarkerForItemNumber(itemNumber int) string {
	rv := ffi.CallMethod[string](t_, "markerForItemNumber:", itemNumber)
	return rv
}

func (t_ TextList) MarkerFormat() TextListMarkerFormat {
	rv := ffi.CallMethod[TextListMarkerFormat](t_, "markerFormat")
	return rv
}

func (t_ TextList) IsOrdered() bool {
	rv := ffi.CallMethod[bool](t_, "isOrdered")
	return rv
}

func (t_ TextList) ListOptions() TextListOptions {
	rv := ffi.CallMethod[TextListOptions](t_, "listOptions")
	return rv
}

func (t_ TextList) StartingItemNumber() int {
	rv := ffi.CallMethod[int](t_, "startingItemNumber")
	return rv
}

func (t_ TextList) SetStartingItemNumber(value int) {
	ffi.CallMethod[ffi.Void](t_, "setStartingItemNumber:", value)
}
