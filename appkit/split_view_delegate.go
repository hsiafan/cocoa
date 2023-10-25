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
