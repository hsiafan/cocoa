// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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

func (tc _TextContainerClass) New() TextContainer {
	rv := ffi.CallMethod[TextContainer](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextContainer() TextContainer {
	return TextContainerClass.New()
}

func (t_ TextContainer) Init() TextContainer {
	rv := ffi.CallMethod[TextContainer](t_, "init")
	return rv
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
