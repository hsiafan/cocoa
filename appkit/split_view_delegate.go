// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type SplitViewDelegate interface {
	ImplementsSplitViewWillResizeSubviews() bool
	// optional
	SplitViewWillResizeSubviews(notification foundation.Notification)
	ImplementsSplitViewDidResizeSubviews() bool
	// optional
	SplitViewDidResizeSubviews(notification foundation.Notification)
	ImplementsSplitView_CanCollapseSubview() bool
	// optional
	SplitView_CanCollapseSubview(splitView SplitView, subview View) bool
	ImplementsSplitView_ShouldCollapseSubview_ForDoubleClickOnDividerAtIndex() bool
	// optional
	// deprecated
	SplitView_ShouldCollapseSubview_ForDoubleClickOnDividerAtIndex(splitView SplitView, subview View, dividerIndex int) bool
	ImplementsSplitView_EffectiveRect_ForDrawnRect_OfDividerAtIndex() bool
	// optional
	SplitView_EffectiveRect_ForDrawnRect_OfDividerAtIndex(splitView SplitView, proposedEffectiveRect foundation.Rect, drawnRect foundation.Rect, dividerIndex int) foundation.Rect
	ImplementsSplitView_ShouldHideDividerAtIndex() bool
	// optional
	SplitView_ShouldHideDividerAtIndex(splitView SplitView, dividerIndex int) bool
	ImplementsSplitView_AdditionalEffectiveRectOfDividerAtIndex() bool
	// optional
	SplitView_AdditionalEffectiveRectOfDividerAtIndex(splitView SplitView, dividerIndex int) foundation.Rect
	ImplementsSplitView_ConstrainSplitPosition_OfSubviewAt() bool
	// optional
	SplitView_ConstrainSplitPosition_OfSubviewAt(splitView SplitView, proposedPosition float64, dividerIndex int) float64
	ImplementsSplitView_ConstrainMinCoordinate_OfSubviewAt() bool
	// optional
	SplitView_ConstrainMinCoordinate_OfSubviewAt(splitView SplitView, proposedMinimumPosition float64, dividerIndex int) float64
	ImplementsSplitView_ConstrainMaxCoordinate_OfSubviewAt() bool
	// optional
	SplitView_ConstrainMaxCoordinate_OfSubviewAt(splitView SplitView, proposedMaximumPosition float64, dividerIndex int) float64
	ImplementsSplitView_ResizeSubviewsWithOldSize() bool
	// optional
	SplitView_ResizeSubviewsWithOldSize(splitView SplitView, oldSize foundation.Size)
	ImplementsSplitView_ShouldAdjustSizeOfSubview() bool
	// optional
	SplitView_ShouldAdjustSizeOfSubview(splitView SplitView, view View) bool
}

func WrapSplitViewDelegate(v SplitViewDelegate) objc.Object {
	return objc.WrapAsProtocol("NSSplitViewDelegate", v)
}

type SplitViewDelegateBase struct {
}

func (p *SplitViewDelegateBase) ImplementsSplitViewWillResizeSubviews() bool {
	return false
}

func (p *SplitViewDelegateBase) SplitViewWillResizeSubviews(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *SplitViewDelegateBase) ImplementsSplitViewDidResizeSubviews() bool {
	return false
}

func (p *SplitViewDelegateBase) SplitViewDidResizeSubviews(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *SplitViewDelegateBase) ImplementsSplitView_CanCollapseSubview() bool {
	return false
}

func (p *SplitViewDelegateBase) SplitView_CanCollapseSubview(splitView SplitView, subview View) bool {
	panic("unimpemented protocol method")
}

func (p *SplitViewDelegateBase) ImplementsSplitView_ShouldCollapseSubview_ForDoubleClickOnDividerAtIndex() bool {
	return false
}

// deprecated
func (p *SplitViewDelegateBase) SplitView_ShouldCollapseSubview_ForDoubleClickOnDividerAtIndex(splitView SplitView, subview View, dividerIndex int) bool {
	panic("unimpemented protocol method")
}

func (p *SplitViewDelegateBase) ImplementsSplitView_EffectiveRect_ForDrawnRect_OfDividerAtIndex() bool {
	return false
}

func (p *SplitViewDelegateBase) SplitView_EffectiveRect_ForDrawnRect_OfDividerAtIndex(splitView SplitView, proposedEffectiveRect foundation.Rect, drawnRect foundation.Rect, dividerIndex int) foundation.Rect {
	panic("unimpemented protocol method")
}

func (p *SplitViewDelegateBase) ImplementsSplitView_ShouldHideDividerAtIndex() bool {
	return false
}

func (p *SplitViewDelegateBase) SplitView_ShouldHideDividerAtIndex(splitView SplitView, dividerIndex int) bool {
	panic("unimpemented protocol method")
}

func (p *SplitViewDelegateBase) ImplementsSplitView_AdditionalEffectiveRectOfDividerAtIndex() bool {
	return false
}

func (p *SplitViewDelegateBase) SplitView_AdditionalEffectiveRectOfDividerAtIndex(splitView SplitView, dividerIndex int) foundation.Rect {
	panic("unimpemented protocol method")
}

func (p *SplitViewDelegateBase) ImplementsSplitView_ConstrainSplitPosition_OfSubviewAt() bool {
	return false
}

func (p *SplitViewDelegateBase) SplitView_ConstrainSplitPosition_OfSubviewAt(splitView SplitView, proposedPosition float64, dividerIndex int) float64 {
	panic("unimpemented protocol method")
}

func (p *SplitViewDelegateBase) ImplementsSplitView_ConstrainMinCoordinate_OfSubviewAt() bool {
	return false
}

func (p *SplitViewDelegateBase) SplitView_ConstrainMinCoordinate_OfSubviewAt(splitView SplitView, proposedMinimumPosition float64, dividerIndex int) float64 {
	panic("unimpemented protocol method")
}

func (p *SplitViewDelegateBase) ImplementsSplitView_ConstrainMaxCoordinate_OfSubviewAt() bool {
	return false
}

func (p *SplitViewDelegateBase) SplitView_ConstrainMaxCoordinate_OfSubviewAt(splitView SplitView, proposedMaximumPosition float64, dividerIndex int) float64 {
	panic("unimpemented protocol method")
}

func (p *SplitViewDelegateBase) ImplementsSplitView_ResizeSubviewsWithOldSize() bool {
	return false
}

func (p *SplitViewDelegateBase) SplitView_ResizeSubviewsWithOldSize(splitView SplitView, oldSize foundation.Size) {
	panic("unimpemented protocol method")
}

func (p *SplitViewDelegateBase) ImplementsSplitView_ShouldAdjustSizeOfSubview() bool {
	return false
}

func (p *SplitViewDelegateBase) SplitView_ShouldAdjustSizeOfSubview(splitView SplitView, view View) bool {
	panic("unimpemented protocol method")
}

type SplitViewDelegateCreator struct {
	className string
	class     objc.Class
}

func NewSplitViewDelegateCreator(name string) *SplitViewDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &SplitViewDelegateCreator{className: name, class: class}
}

func (c *SplitViewDelegateCreator) SetSplitViewWillResizeSubviews(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("splitViewWillResizeSubviews:"), handle)
}

func (c *SplitViewDelegateCreator) SetSplitViewDidResizeSubviews(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("splitViewDidResizeSubviews:"), handle)
}

func (c *SplitViewDelegateCreator) SetSplitView_CanCollapseSubview(handle func(o objc.Object, splitView SplitView, subview View) bool) {
	objc.AddMethod(c.class, objc.SEL("splitView:canCollapseSubview:"), handle)
}

// deprecated
func (c *SplitViewDelegateCreator) SetSplitView_ShouldCollapseSubview_ForDoubleClickOnDividerAtIndex(handle func(o objc.Object, splitView SplitView, subview View, dividerIndex int) bool) {
	objc.AddMethod(c.class, objc.SEL("splitView:shouldCollapseSubview:forDoubleClickOnDividerAtIndex:"), handle)
}

func (c *SplitViewDelegateCreator) SetSplitView_EffectiveRect_ForDrawnRect_OfDividerAtIndex(handle func(o objc.Object, splitView SplitView, proposedEffectiveRect foundation.Rect, drawnRect foundation.Rect, dividerIndex int) foundation.Rect) {
	objc.AddMethod(c.class, objc.SEL("splitView:effectiveRect:forDrawnRect:ofDividerAtIndex:"), handle)
}

func (c *SplitViewDelegateCreator) SetSplitView_ShouldHideDividerAtIndex(handle func(o objc.Object, splitView SplitView, dividerIndex int) bool) {
	objc.AddMethod(c.class, objc.SEL("splitView:shouldHideDividerAtIndex:"), handle)
}

func (c *SplitViewDelegateCreator) SetSplitView_AdditionalEffectiveRectOfDividerAtIndex(handle func(o objc.Object, splitView SplitView, dividerIndex int) foundation.Rect) {
	objc.AddMethod(c.class, objc.SEL("splitView:additionalEffectiveRectOfDividerAtIndex:"), handle)
}

func (c *SplitViewDelegateCreator) SetSplitView_ConstrainSplitPosition_OfSubviewAt(handle func(o objc.Object, splitView SplitView, proposedPosition float64, dividerIndex int) float64) {
	objc.AddMethod(c.class, objc.SEL("splitView:constrainSplitPosition:ofSubviewAt:"), handle)
}

func (c *SplitViewDelegateCreator) SetSplitView_ConstrainMinCoordinate_OfSubviewAt(handle func(o objc.Object, splitView SplitView, proposedMinimumPosition float64, dividerIndex int) float64) {
	objc.AddMethod(c.class, objc.SEL("splitView:constrainMinCoordinate:ofSubviewAt:"), handle)
}

func (c *SplitViewDelegateCreator) SetSplitView_ConstrainMaxCoordinate_OfSubviewAt(handle func(o objc.Object, splitView SplitView, proposedMaximumPosition float64, dividerIndex int) float64) {
	objc.AddMethod(c.class, objc.SEL("splitView:constrainMaxCoordinate:ofSubviewAt:"), handle)
}

func (c *SplitViewDelegateCreator) SetSplitView_ResizeSubviewsWithOldSize(handle func(o objc.Object, splitView SplitView, oldSize foundation.Size)) {
	objc.AddMethod(c.class, objc.SEL("splitView:resizeSubviewsWithOldSize:"), handle)
}

func (c *SplitViewDelegateCreator) SetSplitView_ShouldAdjustSizeOfSubview(handle func(o objc.Object, splitView SplitView, view View) bool) {
	objc.AddMethod(c.class, objc.SEL("splitView:shouldAdjustSizeOfSubview:"), handle)
}

func (c *SplitViewDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
