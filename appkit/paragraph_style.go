// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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

func (pc _ParagraphStyleClass) New() ParagraphStyle {
	rv := ffi.CallMethod[ParagraphStyle](pc, "new")
	rv.Autorelease()
	return rv
}

func NewParagraphStyle() ParagraphStyle {
	return ParagraphStyleClass.New()
}

func (p_ ParagraphStyle) Init() ParagraphStyle {
	rv := ffi.CallMethod[ParagraphStyle](p_, "init")
	return rv
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
