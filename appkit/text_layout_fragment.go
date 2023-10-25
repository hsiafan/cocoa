// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TextLayoutFragmentClass = _TextLayoutFragmentClass{objc.GetClass("NSTextLayoutFragment")}

type _TextLayoutFragmentClass struct {
	objc.Class
}

type ITextLayoutFragment interface {
	objc.IObject
	DrawAtPoint_InContext(point coregraphics.Point, context coregraphics.ContextRef)
	InvalidateLayout()
	BottomMargin() float64
	LeadingPadding() float64
	TopMargin() float64
	TrailingPadding() float64
	TextLayoutManager() TextLayoutManager
	LayoutQueue() foundation.OperationQueue
	SetLayoutQueue(value foundation.IOperationQueue)
	LayoutFragmentFrame() coregraphics.Rect
	RenderingSurfaceBounds() coregraphics.Rect
	TextAttachmentViewProviders() []TextAttachmentViewProvider
	State() TextLayoutFragmentState
	RangeInElement() TextRange
	TextElement() TextElement
	TextLineFragments() []TextLineFragment
}

type TextLayoutFragment struct {
	objc.Object
}

func MakeTextLayoutFragment(ptr unsafe.Pointer) TextLayoutFragment {
	return TextLayoutFragment{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextLayoutFragment) InitWithTextElement_Range(textElement ITextElement, rangeInElement ITextRange) TextLayoutFragment {
	rv := objc.CallMethod[TextLayoutFragment](t_, objc.SEL("initWithTextElement:range:"), objc.ExtractPtr(textElement), objc.ExtractPtr(rangeInElement))
	return rv
}

func (tc _TextLayoutFragmentClass) Alloc() TextLayoutFragment {
	rv := objc.CallMethod[TextLayoutFragment](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TextLayoutFragmentClass) New() TextLayoutFragment {
	rv := objc.CallMethod[TextLayoutFragment](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTextLayoutFragment() TextLayoutFragment {
	return TextLayoutFragmentClass.New()
}

func (t_ TextLayoutFragment) Init() TextLayoutFragment {
	rv := objc.CallMethod[TextLayoutFragment](t_, objc.SEL("init"))
	return rv
}

func (t_ TextLayoutFragment) DrawAtPoint_InContext(point coregraphics.Point, context coregraphics.ContextRef) {
	objc.CallMethod[objc.Void](t_, objc.SEL("drawAtPoint:inContext:"), point, context)
}

func (t_ TextLayoutFragment) InvalidateLayout() {
	objc.CallMethod[objc.Void](t_, objc.SEL("invalidateLayout"))
}

func (t_ TextLayoutFragment) BottomMargin() float64 {
	rv := objc.CallMethod[float64](t_, objc.SEL("bottomMargin"))
	return rv
}

func (t_ TextLayoutFragment) LeadingPadding() float64 {
	rv := objc.CallMethod[float64](t_, objc.SEL("leadingPadding"))
	return rv
}

func (t_ TextLayoutFragment) TopMargin() float64 {
	rv := objc.CallMethod[float64](t_, objc.SEL("topMargin"))
	return rv
}

func (t_ TextLayoutFragment) TrailingPadding() float64 {
	rv := objc.CallMethod[float64](t_, objc.SEL("trailingPadding"))
	return rv
}

// weak property
func (t_ TextLayoutFragment) TextLayoutManager() TextLayoutManager {
	rv := objc.CallMethod[TextLayoutManager](t_, objc.SEL("textLayoutManager"))
	return rv
}

func (t_ TextLayoutFragment) LayoutQueue() foundation.OperationQueue {
	rv := objc.CallMethod[foundation.OperationQueue](t_, objc.SEL("layoutQueue"))
	return rv
}

func (t_ TextLayoutFragment) SetLayoutQueue(value foundation.IOperationQueue) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setLayoutQueue:"), objc.ExtractPtr(value))
}

func (t_ TextLayoutFragment) LayoutFragmentFrame() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](t_, objc.SEL("layoutFragmentFrame"))
	return rv
}

func (t_ TextLayoutFragment) RenderingSurfaceBounds() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](t_, objc.SEL("renderingSurfaceBounds"))
	return rv
}

func (t_ TextLayoutFragment) TextAttachmentViewProviders() []TextAttachmentViewProvider {
	rv := objc.CallMethod[[]TextAttachmentViewProvider](t_, objc.SEL("textAttachmentViewProviders"))
	return rv
}

func (t_ TextLayoutFragment) State() TextLayoutFragmentState {
	rv := objc.CallMethod[TextLayoutFragmentState](t_, objc.SEL("state"))
	return rv
}

func (t_ TextLayoutFragment) RangeInElement() TextRange {
	rv := objc.CallMethod[TextRange](t_, objc.SEL("rangeInElement"))
	return rv
}

// weak property
func (t_ TextLayoutFragment) TextElement() TextElement {
	rv := objc.CallMethod[TextElement](t_, objc.SEL("textElement"))
	return rv
}

func (t_ TextLayoutFragment) TextLineFragments() []TextLineFragment {
	rv := objc.CallMethod[[]TextLineFragment](t_, objc.SEL("textLineFragments"))
	return rv
}
