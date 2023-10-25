// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

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
	rv := objc.CallMethod[TextContainer](t_, objc.SEL("initWithSize:"), size)
	return rv
}

func (t_ TextContainer) InitWithContainerSize(aContainerSize foundation.Size) TextContainer {
	rv := objc.CallMethod[TextContainer](t_, objc.SEL("initWithContainerSize:"), aContainerSize)
	return rv
}

func (tc _TextContainerClass) Alloc() TextContainer {
	rv := objc.CallMethod[TextContainer](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TextContainerClass) New() TextContainer {
	rv := objc.CallMethod[TextContainer](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTextContainer() TextContainer {
	return TextContainerClass.New()
}

func (t_ TextContainer) Init() TextContainer {
	rv := objc.CallMethod[TextContainer](t_, objc.SEL("init"))
	return rv
}

func (t_ TextContainer) ReplaceLayoutManager(newLayoutManager ILayoutManager) {
	objc.CallMethod[objc.Void](t_, objc.SEL("replaceLayoutManager:"), objc.ExtractPtr(newLayoutManager))
}

func (t_ TextContainer) LineFragmentRectForProposedRect_AtIndex_WritingDirection_RemainingRect(proposedRect foundation.Rect, characterIndex uint, baseWritingDirection WritingDirection, remainingRect *foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.SEL("lineFragmentRectForProposedRect:atIndex:writingDirection:remainingRect:"), proposedRect, characterIndex, baseWritingDirection, remainingRect)
	return rv
}

// deprecated
func (t_ TextContainer) LineFragmentRectForProposedRect_SweepDirection_MovementDirection_RemainingRect(proposedRect foundation.Rect, sweepDirection LineSweepDirection, movementDirection LineMovementDirection, remainingRect *foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.SEL("lineFragmentRectForProposedRect:sweepDirection:movementDirection:remainingRect:"), proposedRect, sweepDirection, movementDirection, remainingRect)
	return rv
}

// deprecated
func (t_ TextContainer) ContainsPoint(point foundation.Point) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("containsPoint:"), point)
	return rv
}

// weak property
func (t_ TextContainer) LayoutManager() LayoutManager {
	rv := objc.CallMethod[LayoutManager](t_, objc.SEL("layoutManager"))
	return rv
}

// weak property
func (t_ TextContainer) SetLayoutManager(value ILayoutManager) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setLayoutManager:"), objc.ExtractPtr(value))
}

// weak property
func (t_ TextContainer) TextLayoutManager() TextLayoutManager {
	rv := objc.CallMethod[TextLayoutManager](t_, objc.SEL("textLayoutManager"))
	return rv
}

// weak property
func (t_ TextContainer) TextView() TextView {
	rv := objc.CallMethod[TextView](t_, objc.SEL("textView"))
	return rv
}

// weak property
func (t_ TextContainer) SetTextView(value ITextView) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTextView:"), objc.ExtractPtr(value))
}

func (t_ TextContainer) Size() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.SEL("size"))
	return rv
}

func (t_ TextContainer) SetSize(value foundation.Size) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSize:"), value)
}

func (t_ TextContainer) ExclusionPaths() []BezierPath {
	rv := objc.CallMethod[[]BezierPath](t_, objc.SEL("exclusionPaths"))
	return rv
}

func (t_ TextContainer) SetExclusionPaths(value []IBezierPath) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setExclusionPaths:"), value)
}

func (t_ TextContainer) LineBreakMode() LineBreakMode {
	rv := objc.CallMethod[LineBreakMode](t_, objc.SEL("lineBreakMode"))
	return rv
}

func (t_ TextContainer) SetLineBreakMode(value LineBreakMode) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setLineBreakMode:"), value)
}

func (t_ TextContainer) WidthTracksTextView() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("widthTracksTextView"))
	return rv
}

func (t_ TextContainer) SetWidthTracksTextView(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setWidthTracksTextView:"), value)
}

func (t_ TextContainer) HeightTracksTextView() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("heightTracksTextView"))
	return rv
}

func (t_ TextContainer) SetHeightTracksTextView(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setHeightTracksTextView:"), value)
}

func (t_ TextContainer) MaximumNumberOfLines() uint {
	rv := objc.CallMethod[uint](t_, objc.SEL("maximumNumberOfLines"))
	return rv
}

func (t_ TextContainer) SetMaximumNumberOfLines(value uint) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setMaximumNumberOfLines:"), value)
}

func (t_ TextContainer) LineFragmentPadding() float64 {
	rv := objc.CallMethod[float64](t_, objc.SEL("lineFragmentPadding"))
	return rv
}

func (t_ TextContainer) SetLineFragmentPadding(value float64) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setLineFragmentPadding:"), value)
}

func (t_ TextContainer) IsSimpleRectangularTextContainer() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isSimpleRectangularTextContainer"))
	return rv
}

// deprecated
func (t_ TextContainer) ContainerSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.SEL("containerSize"))
	return rv
}

// deprecated
func (t_ TextContainer) SetContainerSize(value foundation.Size) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setContainerSize:"), value)
}
