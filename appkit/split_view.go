// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var SplitViewClass = _SplitViewClass{objc.GetClass("NSSplitView")}

type _SplitViewClass struct {
	objc.Class
}

type ISplitView interface {
	IView
	AddArrangedSubview(view IView)
	InsertArrangedSubview_AtIndex(view IView, index int)
	RemoveArrangedSubview(view IView)
	AdjustSubviews()
	IsSubviewCollapsed(subview IView) bool
	HoldingPriorityForSubviewAtIndex(subviewIndex int) LayoutPriority
	SetHoldingPriority_ForSubviewAtIndex(priority LayoutPriority, subviewIndex int)
	DrawDividerInRect(rect foundation.Rect)
	MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64
	MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64
	SetPosition_OfDividerAtIndex(position float64, dividerIndex int)
	// deprecated
	IsPaneSplitter() bool
	// deprecated
	SetIsPaneSplitter(flag bool)
	Delegate() SplitViewDelegateWrapper
	SetDelegate(value SplitViewDelegate)
	ArrangesAllSubviews() bool
	SetArrangesAllSubviews(value bool)
	ArrangedSubviews() []View
	IsVertical() bool
	SetVertical(value bool)
	DividerStyle() SplitViewDividerStyle
	SetDividerStyle(value SplitViewDividerStyle)
	DividerColor() Color
	DividerThickness() float64
	AutosaveName() SplitViewAutosaveName
	SetAutosaveName(value SplitViewAutosaveName)
}

type SplitView struct {
	View
}

func MakeSplitView(ptr unsafe.Pointer) SplitView {
	return SplitView{
		View: MakeView(ptr),
	}
}

func (s_ SplitView) InitWithFrame(frameRect foundation.Rect) SplitView {
	rv := ffi.CallMethod[SplitView](s_, "initWithFrame:", frameRect)
	rv.Autorelease()
	return rv
}

func (s_ SplitView) Init() SplitView {
	rv := ffi.CallMethod[SplitView](s_, "init")
	rv.Autorelease()
	return rv
}

func (sc _SplitViewClass) Alloc() SplitView {
	rv := ffi.CallMethod[SplitView](sc, "alloc")
	return rv
}

func (sc _SplitViewClass) New() SplitView {
	rv := ffi.CallMethod[SplitView](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSplitView() SplitView {
	return SplitViewClass.New()
}

func (s_ SplitView) AddArrangedSubview(view IView) {
	ffi.CallMethod[ffi.Void](s_, "addArrangedSubview:", view)
}

func (s_ SplitView) InsertArrangedSubview_AtIndex(view IView, index int) {
	ffi.CallMethod[ffi.Void](s_, "insertArrangedSubview:atIndex:", view, index)
}

func (s_ SplitView) RemoveArrangedSubview(view IView) {
	ffi.CallMethod[ffi.Void](s_, "removeArrangedSubview:", view)
}

func (s_ SplitView) AdjustSubviews() {
	ffi.CallMethod[ffi.Void](s_, "adjustSubviews")
}

func (s_ SplitView) IsSubviewCollapsed(subview IView) bool {
	rv := ffi.CallMethod[bool](s_, "isSubviewCollapsed:", subview)
	return rv
}

func (s_ SplitView) HoldingPriorityForSubviewAtIndex(subviewIndex int) LayoutPriority {
	rv := ffi.CallMethod[LayoutPriority](s_, "holdingPriorityForSubviewAtIndex:", subviewIndex)
	return rv
}

func (s_ SplitView) SetHoldingPriority_ForSubviewAtIndex(priority LayoutPriority, subviewIndex int) {
	ffi.CallMethod[ffi.Void](s_, "setHoldingPriority:forSubviewAtIndex:", priority, subviewIndex)
}

func (s_ SplitView) DrawDividerInRect(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](s_, "drawDividerInRect:", rect)
}

func (s_ SplitView) MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := ffi.CallMethod[float64](s_, "minPossiblePositionOfDividerAtIndex:", dividerIndex)
	return rv
}

func (s_ SplitView) MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := ffi.CallMethod[float64](s_, "maxPossiblePositionOfDividerAtIndex:", dividerIndex)
	return rv
}

func (s_ SplitView) SetPosition_OfDividerAtIndex(position float64, dividerIndex int) {
	ffi.CallMethod[ffi.Void](s_, "setPosition:ofDividerAtIndex:", position, dividerIndex)
}

// deprecated
func (s_ SplitView) IsPaneSplitter() bool {
	rv := ffi.CallMethod[bool](s_, "isPaneSplitter")
	return rv
}

// deprecated
func (s_ SplitView) SetIsPaneSplitter(flag bool) {
	ffi.CallMethod[ffi.Void](s_, "setIsPaneSplitter:", flag)
}

func (s_ SplitView) Delegate() SplitViewDelegateWrapper {
	rv := ffi.CallMethod[SplitViewDelegateWrapper](s_, "delegate")
	return rv
}

func (s_ SplitView) SetDelegate(value SplitViewDelegate) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(s_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](s_, "setDelegate:", po)
}

func (s_ SplitView) ArrangesAllSubviews() bool {
	rv := ffi.CallMethod[bool](s_, "arrangesAllSubviews")
	return rv
}

func (s_ SplitView) SetArrangesAllSubviews(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setArrangesAllSubviews:", value)
}

func (s_ SplitView) ArrangedSubviews() []View {
	rv := ffi.CallMethod[[]View](s_, "arrangedSubviews")
	return rv
}

func (s_ SplitView) IsVertical() bool {
	rv := ffi.CallMethod[bool](s_, "isVertical")
	return rv
}

func (s_ SplitView) SetVertical(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setVertical:", value)
}

func (s_ SplitView) DividerStyle() SplitViewDividerStyle {
	rv := ffi.CallMethod[SplitViewDividerStyle](s_, "dividerStyle")
	return rv
}

func (s_ SplitView) SetDividerStyle(value SplitViewDividerStyle) {
	ffi.CallMethod[ffi.Void](s_, "setDividerStyle:", value)
}

func (s_ SplitView) DividerColor() Color {
	rv := ffi.CallMethod[Color](s_, "dividerColor")
	return rv
}

func (s_ SplitView) DividerThickness() float64 {
	rv := ffi.CallMethod[float64](s_, "dividerThickness")
	return rv
}

func (s_ SplitView) AutosaveName() SplitViewAutosaveName {
	rv := ffi.CallMethod[SplitViewAutosaveName](s_, "autosaveName")
	return rv
}

func (s_ SplitView) SetAutosaveName(value SplitViewAutosaveName) {
	ffi.CallMethod[ffi.Void](s_, "setAutosaveName:", value)
}

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
	ffi.CallMethod[ffi.Void](s_, "splitViewWillResizeSubviews:", notification)
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitViewDidResizeSubviews() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitViewDidResizeSubviews:"))
}

func (s_ SplitViewDelegateWrapper) SplitViewDidResizeSubviews(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](s_, "splitViewDidResizeSubviews:", notification)
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_CanCollapseSubview() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:canCollapseSubview:"))
}

func (s_ SplitViewDelegateWrapper) SplitView_CanCollapseSubview(splitView ISplitView, subview IView) bool {
	rv := ffi.CallMethod[bool](s_, "splitView:canCollapseSubview:", splitView, subview)
	return rv
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_ShouldCollapseSubview_ForDoubleClickOnDividerAtIndex() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:shouldCollapseSubview:forDoubleClickOnDividerAtIndex:"))
}

// deprecated
func (s_ SplitViewDelegateWrapper) SplitView_ShouldCollapseSubview_ForDoubleClickOnDividerAtIndex(splitView ISplitView, subview IView, dividerIndex int) bool {
	rv := ffi.CallMethod[bool](s_, "splitView:shouldCollapseSubview:forDoubleClickOnDividerAtIndex:", splitView, subview, dividerIndex)
	return rv
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_EffectiveRect_ForDrawnRect_OfDividerAtIndex() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:effectiveRect:forDrawnRect:ofDividerAtIndex:"))
}

func (s_ SplitViewDelegateWrapper) SplitView_EffectiveRect_ForDrawnRect_OfDividerAtIndex(splitView ISplitView, proposedEffectiveRect foundation.Rect, drawnRect foundation.Rect, dividerIndex int) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "splitView:effectiveRect:forDrawnRect:ofDividerAtIndex:", splitView, proposedEffectiveRect, drawnRect, dividerIndex)
	return rv
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_ShouldHideDividerAtIndex() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:shouldHideDividerAtIndex:"))
}

func (s_ SplitViewDelegateWrapper) SplitView_ShouldHideDividerAtIndex(splitView ISplitView, dividerIndex int) bool {
	rv := ffi.CallMethod[bool](s_, "splitView:shouldHideDividerAtIndex:", splitView, dividerIndex)
	return rv
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_AdditionalEffectiveRectOfDividerAtIndex() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:additionalEffectiveRectOfDividerAtIndex:"))
}

func (s_ SplitViewDelegateWrapper) SplitView_AdditionalEffectiveRectOfDividerAtIndex(splitView ISplitView, dividerIndex int) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "splitView:additionalEffectiveRectOfDividerAtIndex:", splitView, dividerIndex)
	return rv
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_ConstrainSplitPosition_OfSubviewAt() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:constrainSplitPosition:ofSubviewAt:"))
}

func (s_ SplitViewDelegateWrapper) SplitView_ConstrainSplitPosition_OfSubviewAt(splitView ISplitView, proposedPosition float64, dividerIndex int) float64 {
	rv := ffi.CallMethod[float64](s_, "splitView:constrainSplitPosition:ofSubviewAt:", splitView, proposedPosition, dividerIndex)
	return rv
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_ConstrainMinCoordinate_OfSubviewAt() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:constrainMinCoordinate:ofSubviewAt:"))
}

func (s_ SplitViewDelegateWrapper) SplitView_ConstrainMinCoordinate_OfSubviewAt(splitView ISplitView, proposedMinimumPosition float64, dividerIndex int) float64 {
	rv := ffi.CallMethod[float64](s_, "splitView:constrainMinCoordinate:ofSubviewAt:", splitView, proposedMinimumPosition, dividerIndex)
	return rv
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_ConstrainMaxCoordinate_OfSubviewAt() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:constrainMaxCoordinate:ofSubviewAt:"))
}

func (s_ SplitViewDelegateWrapper) SplitView_ConstrainMaxCoordinate_OfSubviewAt(splitView ISplitView, proposedMaximumPosition float64, dividerIndex int) float64 {
	rv := ffi.CallMethod[float64](s_, "splitView:constrainMaxCoordinate:ofSubviewAt:", splitView, proposedMaximumPosition, dividerIndex)
	return rv
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_ResizeSubviewsWithOldSize() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:resizeSubviewsWithOldSize:"))
}

func (s_ SplitViewDelegateWrapper) SplitView_ResizeSubviewsWithOldSize(splitView ISplitView, oldSize foundation.Size) {
	ffi.CallMethod[ffi.Void](s_, "splitView:resizeSubviewsWithOldSize:", splitView, oldSize)
}

func (s_ *SplitViewDelegateWrapper) ImplementsSplitView_ShouldAdjustSizeOfSubview() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:shouldAdjustSizeOfSubview:"))
}

func (s_ SplitViewDelegateWrapper) SplitView_ShouldAdjustSizeOfSubview(splitView ISplitView, view IView) bool {
	rv := ffi.CallMethod[bool](s_, "splitView:shouldAdjustSizeOfSubview:", splitView, view)
	return rv
}
