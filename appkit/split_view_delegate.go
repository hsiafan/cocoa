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

type SplitViewDelegateImpl struct {
	_SplitViewWillResizeSubviews                                    func(notification foundation.Notification)
	_SplitViewDidResizeSubviews                                     func(notification foundation.Notification)
	_SplitView_CanCollapseSubview                                   func(splitView SplitView, subview View) bool
	_SplitView_ShouldCollapseSubview_ForDoubleClickOnDividerAtIndex func(splitView SplitView, subview View, dividerIndex int) bool
	_SplitView_EffectiveRect_ForDrawnRect_OfDividerAtIndex          func(splitView SplitView, proposedEffectiveRect foundation.Rect, drawnRect foundation.Rect, dividerIndex int) foundation.Rect
	_SplitView_ShouldHideDividerAtIndex                             func(splitView SplitView, dividerIndex int) bool
	_SplitView_AdditionalEffectiveRectOfDividerAtIndex              func(splitView SplitView, dividerIndex int) foundation.Rect
	_SplitView_ConstrainSplitPosition_OfSubviewAt                   func(splitView SplitView, proposedPosition float64, dividerIndex int) float64
	_SplitView_ConstrainMinCoordinate_OfSubviewAt                   func(splitView SplitView, proposedMinimumPosition float64, dividerIndex int) float64
	_SplitView_ConstrainMaxCoordinate_OfSubviewAt                   func(splitView SplitView, proposedMaximumPosition float64, dividerIndex int) float64
	_SplitView_ResizeSubviewsWithOldSize                            func(splitView SplitView, oldSize foundation.Size)
	_SplitView_ShouldAdjustSizeOfSubview                            func(splitView SplitView, view View) bool
}

func (di *SplitViewDelegateImpl) ImplementsSplitViewWillResizeSubviews() bool {
	return di._SplitViewWillResizeSubviews != nil
}

func (di *SplitViewDelegateImpl) SetSplitViewWillResizeSubviews(f func(notification foundation.Notification)) {
	di._SplitViewWillResizeSubviews = f
}

func (di *SplitViewDelegateImpl) SplitViewWillResizeSubviews(notification foundation.Notification) {
	di._SplitViewWillResizeSubviews(notification)
}
func (di *SplitViewDelegateImpl) ImplementsSplitViewDidResizeSubviews() bool {
	return di._SplitViewDidResizeSubviews != nil
}

func (di *SplitViewDelegateImpl) SetSplitViewDidResizeSubviews(f func(notification foundation.Notification)) {
	di._SplitViewDidResizeSubviews = f
}

func (di *SplitViewDelegateImpl) SplitViewDidResizeSubviews(notification foundation.Notification) {
	di._SplitViewDidResizeSubviews(notification)
}
func (di *SplitViewDelegateImpl) ImplementsSplitView_CanCollapseSubview() bool {
	return di._SplitView_CanCollapseSubview != nil
}

func (di *SplitViewDelegateImpl) SetSplitView_CanCollapseSubview(f func(splitView SplitView, subview View) bool) {
	di._SplitView_CanCollapseSubview = f
}

func (di *SplitViewDelegateImpl) SplitView_CanCollapseSubview(splitView SplitView, subview View) bool {
	return di._SplitView_CanCollapseSubview(splitView, subview)
}
func (di *SplitViewDelegateImpl) ImplementsSplitView_ShouldCollapseSubview_ForDoubleClickOnDividerAtIndex() bool {
	return di._SplitView_ShouldCollapseSubview_ForDoubleClickOnDividerAtIndex != nil
}

// deprecated
func (di *SplitViewDelegateImpl) SetSplitView_ShouldCollapseSubview_ForDoubleClickOnDividerAtIndex(f func(splitView SplitView, subview View, dividerIndex int) bool) {
	di._SplitView_ShouldCollapseSubview_ForDoubleClickOnDividerAtIndex = f
}

// deprecated
func (di *SplitViewDelegateImpl) SplitView_ShouldCollapseSubview_ForDoubleClickOnDividerAtIndex(splitView SplitView, subview View, dividerIndex int) bool {
	return di._SplitView_ShouldCollapseSubview_ForDoubleClickOnDividerAtIndex(splitView, subview, dividerIndex)
}
func (di *SplitViewDelegateImpl) ImplementsSplitView_EffectiveRect_ForDrawnRect_OfDividerAtIndex() bool {
	return di._SplitView_EffectiveRect_ForDrawnRect_OfDividerAtIndex != nil
}

func (di *SplitViewDelegateImpl) SetSplitView_EffectiveRect_ForDrawnRect_OfDividerAtIndex(f func(splitView SplitView, proposedEffectiveRect foundation.Rect, drawnRect foundation.Rect, dividerIndex int) foundation.Rect) {
	di._SplitView_EffectiveRect_ForDrawnRect_OfDividerAtIndex = f
}

func (di *SplitViewDelegateImpl) SplitView_EffectiveRect_ForDrawnRect_OfDividerAtIndex(splitView SplitView, proposedEffectiveRect foundation.Rect, drawnRect foundation.Rect, dividerIndex int) foundation.Rect {
	return di._SplitView_EffectiveRect_ForDrawnRect_OfDividerAtIndex(splitView, proposedEffectiveRect, drawnRect, dividerIndex)
}
func (di *SplitViewDelegateImpl) ImplementsSplitView_ShouldHideDividerAtIndex() bool {
	return di._SplitView_ShouldHideDividerAtIndex != nil
}

func (di *SplitViewDelegateImpl) SetSplitView_ShouldHideDividerAtIndex(f func(splitView SplitView, dividerIndex int) bool) {
	di._SplitView_ShouldHideDividerAtIndex = f
}

func (di *SplitViewDelegateImpl) SplitView_ShouldHideDividerAtIndex(splitView SplitView, dividerIndex int) bool {
	return di._SplitView_ShouldHideDividerAtIndex(splitView, dividerIndex)
}
func (di *SplitViewDelegateImpl) ImplementsSplitView_AdditionalEffectiveRectOfDividerAtIndex() bool {
	return di._SplitView_AdditionalEffectiveRectOfDividerAtIndex != nil
}

func (di *SplitViewDelegateImpl) SetSplitView_AdditionalEffectiveRectOfDividerAtIndex(f func(splitView SplitView, dividerIndex int) foundation.Rect) {
	di._SplitView_AdditionalEffectiveRectOfDividerAtIndex = f
}

func (di *SplitViewDelegateImpl) SplitView_AdditionalEffectiveRectOfDividerAtIndex(splitView SplitView, dividerIndex int) foundation.Rect {
	return di._SplitView_AdditionalEffectiveRectOfDividerAtIndex(splitView, dividerIndex)
}
func (di *SplitViewDelegateImpl) ImplementsSplitView_ConstrainSplitPosition_OfSubviewAt() bool {
	return di._SplitView_ConstrainSplitPosition_OfSubviewAt != nil
}

func (di *SplitViewDelegateImpl) SetSplitView_ConstrainSplitPosition_OfSubviewAt(f func(splitView SplitView, proposedPosition float64, dividerIndex int) float64) {
	di._SplitView_ConstrainSplitPosition_OfSubviewAt = f
}

func (di *SplitViewDelegateImpl) SplitView_ConstrainSplitPosition_OfSubviewAt(splitView SplitView, proposedPosition float64, dividerIndex int) float64 {
	return di._SplitView_ConstrainSplitPosition_OfSubviewAt(splitView, proposedPosition, dividerIndex)
}
func (di *SplitViewDelegateImpl) ImplementsSplitView_ConstrainMinCoordinate_OfSubviewAt() bool {
	return di._SplitView_ConstrainMinCoordinate_OfSubviewAt != nil
}

func (di *SplitViewDelegateImpl) SetSplitView_ConstrainMinCoordinate_OfSubviewAt(f func(splitView SplitView, proposedMinimumPosition float64, dividerIndex int) float64) {
	di._SplitView_ConstrainMinCoordinate_OfSubviewAt = f
}

func (di *SplitViewDelegateImpl) SplitView_ConstrainMinCoordinate_OfSubviewAt(splitView SplitView, proposedMinimumPosition float64, dividerIndex int) float64 {
	return di._SplitView_ConstrainMinCoordinate_OfSubviewAt(splitView, proposedMinimumPosition, dividerIndex)
}
func (di *SplitViewDelegateImpl) ImplementsSplitView_ConstrainMaxCoordinate_OfSubviewAt() bool {
	return di._SplitView_ConstrainMaxCoordinate_OfSubviewAt != nil
}

func (di *SplitViewDelegateImpl) SetSplitView_ConstrainMaxCoordinate_OfSubviewAt(f func(splitView SplitView, proposedMaximumPosition float64, dividerIndex int) float64) {
	di._SplitView_ConstrainMaxCoordinate_OfSubviewAt = f
}

func (di *SplitViewDelegateImpl) SplitView_ConstrainMaxCoordinate_OfSubviewAt(splitView SplitView, proposedMaximumPosition float64, dividerIndex int) float64 {
	return di._SplitView_ConstrainMaxCoordinate_OfSubviewAt(splitView, proposedMaximumPosition, dividerIndex)
}
func (di *SplitViewDelegateImpl) ImplementsSplitView_ResizeSubviewsWithOldSize() bool {
	return di._SplitView_ResizeSubviewsWithOldSize != nil
}

func (di *SplitViewDelegateImpl) SetSplitView_ResizeSubviewsWithOldSize(f func(splitView SplitView, oldSize foundation.Size)) {
	di._SplitView_ResizeSubviewsWithOldSize = f
}

func (di *SplitViewDelegateImpl) SplitView_ResizeSubviewsWithOldSize(splitView SplitView, oldSize foundation.Size) {
	di._SplitView_ResizeSubviewsWithOldSize(splitView, oldSize)
}
func (di *SplitViewDelegateImpl) ImplementsSplitView_ShouldAdjustSizeOfSubview() bool {
	return di._SplitView_ShouldAdjustSizeOfSubview != nil
}

func (di *SplitViewDelegateImpl) SetSplitView_ShouldAdjustSizeOfSubview(f func(splitView SplitView, view View) bool) {
	di._SplitView_ShouldAdjustSizeOfSubview = f
}

func (di *SplitViewDelegateImpl) SplitView_ShouldAdjustSizeOfSubview(splitView SplitView, view View) bool {
	return di._SplitView_ShouldAdjustSizeOfSubview(splitView, view)
}

type SplitViewDelegateWrapper struct {
	objc.Object
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitViewWillResizeSubviews() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitViewWillResizeSubviews:"))
}

func (s_ SplitViewDelegateWrapper) SplitViewWillResizeSubviews(notification foundation.INotification) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("splitViewWillResizeSubviews:"), objc.ExtractPtr(notification))
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitViewDidResizeSubviews() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitViewDidResizeSubviews:"))
}

func (s_ SplitViewDelegateWrapper) SplitViewDidResizeSubviews(notification foundation.INotification) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("splitViewDidResizeSubviews:"), objc.ExtractPtr(notification))
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_CanCollapseSubview() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:canCollapseSubview:"))
}

func (s_ SplitViewDelegateWrapper) SplitView_CanCollapseSubview(splitView ISplitView, subview IView) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("splitView:canCollapseSubview:"), objc.ExtractPtr(splitView), objc.ExtractPtr(subview))
	return rv
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_ShouldCollapseSubview_ForDoubleClickOnDividerAtIndex() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:shouldCollapseSubview:forDoubleClickOnDividerAtIndex:"))
}

// deprecated
func (s_ SplitViewDelegateWrapper) SplitView_ShouldCollapseSubview_ForDoubleClickOnDividerAtIndex(splitView ISplitView, subview IView, dividerIndex int) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("splitView:shouldCollapseSubview:forDoubleClickOnDividerAtIndex:"), objc.ExtractPtr(splitView), objc.ExtractPtr(subview), dividerIndex)
	return rv
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_EffectiveRect_ForDrawnRect_OfDividerAtIndex() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:effectiveRect:forDrawnRect:ofDividerAtIndex:"))
}

func (s_ SplitViewDelegateWrapper) SplitView_EffectiveRect_ForDrawnRect_OfDividerAtIndex(splitView ISplitView, proposedEffectiveRect foundation.Rect, drawnRect foundation.Rect, dividerIndex int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("splitView:effectiveRect:forDrawnRect:ofDividerAtIndex:"), objc.ExtractPtr(splitView), proposedEffectiveRect, drawnRect, dividerIndex)
	return rv
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_ShouldHideDividerAtIndex() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:shouldHideDividerAtIndex:"))
}

func (s_ SplitViewDelegateWrapper) SplitView_ShouldHideDividerAtIndex(splitView ISplitView, dividerIndex int) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("splitView:shouldHideDividerAtIndex:"), objc.ExtractPtr(splitView), dividerIndex)
	return rv
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_AdditionalEffectiveRectOfDividerAtIndex() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:additionalEffectiveRectOfDividerAtIndex:"))
}

func (s_ SplitViewDelegateWrapper) SplitView_AdditionalEffectiveRectOfDividerAtIndex(splitView ISplitView, dividerIndex int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("splitView:additionalEffectiveRectOfDividerAtIndex:"), objc.ExtractPtr(splitView), dividerIndex)
	return rv
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_ConstrainSplitPosition_OfSubviewAt() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:constrainSplitPosition:ofSubviewAt:"))
}

func (s_ SplitViewDelegateWrapper) SplitView_ConstrainSplitPosition_OfSubviewAt(splitView ISplitView, proposedPosition float64, dividerIndex int) float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("splitView:constrainSplitPosition:ofSubviewAt:"), objc.ExtractPtr(splitView), proposedPosition, dividerIndex)
	return rv
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_ConstrainMinCoordinate_OfSubviewAt() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:constrainMinCoordinate:ofSubviewAt:"))
}

func (s_ SplitViewDelegateWrapper) SplitView_ConstrainMinCoordinate_OfSubviewAt(splitView ISplitView, proposedMinimumPosition float64, dividerIndex int) float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("splitView:constrainMinCoordinate:ofSubviewAt:"), objc.ExtractPtr(splitView), proposedMinimumPosition, dividerIndex)
	return rv
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_ConstrainMaxCoordinate_OfSubviewAt() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:constrainMaxCoordinate:ofSubviewAt:"))
}

func (s_ SplitViewDelegateWrapper) SplitView_ConstrainMaxCoordinate_OfSubviewAt(splitView ISplitView, proposedMaximumPosition float64, dividerIndex int) float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("splitView:constrainMaxCoordinate:ofSubviewAt:"), objc.ExtractPtr(splitView), proposedMaximumPosition, dividerIndex)
	return rv
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_ResizeSubviewsWithOldSize() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:resizeSubviewsWithOldSize:"))
}

func (s_ SplitViewDelegateWrapper) SplitView_ResizeSubviewsWithOldSize(splitView ISplitView, oldSize foundation.Size) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("splitView:resizeSubviewsWithOldSize:"), objc.ExtractPtr(splitView), oldSize)
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_ShouldAdjustSizeOfSubview() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:shouldAdjustSizeOfSubview:"))
}

func (s_ SplitViewDelegateWrapper) SplitView_ShouldAdjustSizeOfSubview(splitView ISplitView, view IView) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("splitView:shouldAdjustSizeOfSubview:"), objc.ExtractPtr(splitView), objc.ExtractPtr(view))
	return rv
}
