// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

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

func (mc _MutableParagraphStyleClass) New() MutableParagraphStyle {
	rv := ffi.CallMethod[MutableParagraphStyle](mc, "new")
	rv.Autorelease()
	return rv
}

func NewMutableParagraphStyle() MutableParagraphStyle {
	return MutableParagraphStyleClass.New()
}

func (m_ MutableParagraphStyle) Init() MutableParagraphStyle {
	rv := ffi.CallMethod[MutableParagraphStyle](m_, "init")
	return rv
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
