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

type SplitViewDelegateWrapper struct {
	objc.Object
}

func (s_ SplitViewDelegateWrapper) SplitViewWillResizeSubviews(notification foundation.INotification) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("splitViewWillResizeSubviews:"), objc.ExtractPtr(notification))
}

func (s_ SplitViewDelegateWrapper) SplitViewDidResizeSubviews(notification foundation.INotification) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("splitViewDidResizeSubviews:"), objc.ExtractPtr(notification))
}

func (s_ SplitViewDelegateWrapper) SplitView_CanCollapseSubview(splitView ISplitView, subview IView) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("splitView:canCollapseSubview:"), objc.ExtractPtr(splitView), objc.ExtractPtr(subview))
	return rv
}

// deprecated
func (s_ SplitViewDelegateWrapper) SplitView_ShouldCollapseSubview_ForDoubleClickOnDividerAtIndex(splitView ISplitView, subview IView, dividerIndex int) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("splitView:shouldCollapseSubview:forDoubleClickOnDividerAtIndex:"), objc.ExtractPtr(splitView), objc.ExtractPtr(subview), dividerIndex)
	return rv
}

func (s_ SplitViewDelegateWrapper) SplitView_EffectiveRect_ForDrawnRect_OfDividerAtIndex(splitView ISplitView, proposedEffectiveRect foundation.Rect, drawnRect foundation.Rect, dividerIndex int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("splitView:effectiveRect:forDrawnRect:ofDividerAtIndex:"), objc.ExtractPtr(splitView), proposedEffectiveRect, drawnRect, dividerIndex)
	return rv
}

func (s_ SplitViewDelegateWrapper) SplitView_ShouldHideDividerAtIndex(splitView ISplitView, dividerIndex int) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("splitView:shouldHideDividerAtIndex:"), objc.ExtractPtr(splitView), dividerIndex)
	return rv
}

func (s_ SplitViewDelegateWrapper) SplitView_AdditionalEffectiveRectOfDividerAtIndex(splitView ISplitView, dividerIndex int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("splitView:additionalEffectiveRectOfDividerAtIndex:"), objc.ExtractPtr(splitView), dividerIndex)
	return rv
}

func (s_ SplitViewDelegateWrapper) SplitView_ConstrainSplitPosition_OfSubviewAt(splitView ISplitView, proposedPosition float64, dividerIndex int) float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("splitView:constrainSplitPosition:ofSubviewAt:"), objc.ExtractPtr(splitView), proposedPosition, dividerIndex)
	return rv
}

func (s_ SplitViewDelegateWrapper) SplitView_ConstrainMinCoordinate_OfSubviewAt(splitView ISplitView, proposedMinimumPosition float64, dividerIndex int) float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("splitView:constrainMinCoordinate:ofSubviewAt:"), objc.ExtractPtr(splitView), proposedMinimumPosition, dividerIndex)
	return rv
}

func (s_ SplitViewDelegateWrapper) SplitView_ConstrainMaxCoordinate_OfSubviewAt(splitView ISplitView, proposedMaximumPosition float64, dividerIndex int) float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("splitView:constrainMaxCoordinate:ofSubviewAt:"), objc.ExtractPtr(splitView), proposedMaximumPosition, dividerIndex)
	return rv
}

func (s_ SplitViewDelegateWrapper) SplitView_ResizeSubviewsWithOldSize(splitView ISplitView, oldSize foundation.Size) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("splitView:resizeSubviewsWithOldSize:"), objc.ExtractPtr(splitView), oldSize)
}

func (s_ SplitViewDelegateWrapper) SplitView_ShouldAdjustSizeOfSubview(splitView ISplitView, view IView) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("splitView:shouldAdjustSizeOfSubview:"), objc.ExtractPtr(splitView), objc.ExtractPtr(view))
	return rv
}
