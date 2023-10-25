// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[ParagraphStyle](pc, objc.SEL("alloc"))
	return rv
}

func (pc _ParagraphStyleClass) New() ParagraphStyle {
	rv := objc.CallMethod[ParagraphStyle](pc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewParagraphStyle() ParagraphStyle {
	return ParagraphStyleClass.New()
}

func (p_ ParagraphStyle) Init() ParagraphStyle {
	rv := objc.CallMethod[ParagraphStyle](p_, objc.SEL("init"))
	return rv
}

func (pc _ParagraphStyleClass) DefaultWritingDirectionForLanguage(languageName string) WritingDirection {
	rv := objc.CallMethod[WritingDirection](pc, objc.SEL("defaultWritingDirectionForLanguage:"), languageName)
	return rv
}

func (pc _ParagraphStyleClass) DefaultParagraphStyle() ParagraphStyle {
	rv := objc.CallMethod[ParagraphStyle](pc, objc.SEL("defaultParagraphStyle"))
	return rv
}

func (p_ ParagraphStyle) Alignment() TextAlignment {
	rv := objc.CallMethod[TextAlignment](p_, objc.SEL("alignment"))
	return rv
}

func (p_ ParagraphStyle) FirstLineHeadIndent() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("firstLineHeadIndent"))
	return rv
}

func (p_ ParagraphStyle) HeadIndent() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("headIndent"))
	return rv
}

func (p_ ParagraphStyle) TailIndent() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("tailIndent"))
	return rv
}

func (p_ ParagraphStyle) LineHeightMultiple() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("lineHeightMultiple"))
	return rv
}

func (p_ ParagraphStyle) MaximumLineHeight() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("maximumLineHeight"))
	return rv
}

func (p_ ParagraphStyle) MinimumLineHeight() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("minimumLineHeight"))
	return rv
}

func (p_ ParagraphStyle) LineSpacing() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("lineSpacing"))
	return rv
}

func (p_ ParagraphStyle) ParagraphSpacing() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("paragraphSpacing"))
	return rv
}

func (p_ ParagraphStyle) ParagraphSpacingBefore() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("paragraphSpacingBefore"))
	return rv
}

func (p_ ParagraphStyle) TabStops() []TextTab {
	rv := objc.CallMethod[[]TextTab](p_, objc.SEL("tabStops"))
	return rv
}

func (p_ ParagraphStyle) DefaultTabInterval() float64 {
	rv := objc.CallMethod[float64](p_, objc.SEL("defaultTabInterval"))
	return rv
}

func (p_ ParagraphStyle) TextBlocks() []TextBlock {
	rv := objc.CallMethod[[]TextBlock](p_, objc.SEL("textBlocks"))
	return rv
}

func (p_ ParagraphStyle) TextLists() []TextList {
	rv := objc.CallMethod[[]TextList](p_, objc.SEL("textLists"))
	return rv
}

func (p_ ParagraphStyle) LineBreakMode() LineBreakMode {
	rv := objc.CallMethod[LineBreakMode](p_, objc.SEL("lineBreakMode"))
	return rv
}

func (p_ ParagraphStyle) LineBreakStrategy() LineBreakStrategy {
	rv := objc.CallMethod[LineBreakStrategy](p_, objc.SEL("lineBreakStrategy"))
	return rv
}

func (p_ ParagraphStyle) HyphenationFactor() float32 {
	rv := objc.CallMethod[float32](p_, objc.SEL("hyphenationFactor"))
	return rv
}

func (p_ ParagraphStyle) UsesDefaultHyphenation() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("usesDefaultHyphenation"))
	return rv
}

func (p_ ParagraphStyle) TighteningFactorForTruncation() float32 {
	rv := objc.CallMethod[float32](p_, objc.SEL("tighteningFactorForTruncation"))
	return rv
}

func (p_ ParagraphStyle) AllowsDefaultTighteningForTruncation() bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("allowsDefaultTighteningForTruncation"))
	return rv
}

func (p_ ParagraphStyle) HeaderLevel() int {
	rv := objc.CallMethod[int](p_, objc.SEL("headerLevel"))
	return rv
}

func (p_ ParagraphStyle) BaseWritingDirection() WritingDirection {
	rv := objc.CallMethod[WritingDirection](p_, objc.SEL("baseWritingDirection"))
	return rv
}
