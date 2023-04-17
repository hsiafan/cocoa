// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var StackViewClass = _StackViewClass{objc.GetClass("NSStackView")}

type _StackViewClass struct {
	objc.Class
}

type IStackView interface {
	IView
	AddView_InGravity(view IView, gravity StackViewGravity)
	InsertView_AtIndex_InGravity(view IView, index uint, gravity StackViewGravity)
	SetViews_InGravity(views []IView, gravity StackViewGravity)
	RemoveView(view IView)
	AddArrangedSubview(view IView)
	InsertArrangedSubview_AtIndex(view IView, index int)
	RemoveArrangedSubview(view IView)
	ViewsInGravity(gravity StackViewGravity) []View
	ClippingResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	HuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	CustomSpacingAfterView(view IView) float64
	SetCustomSpacing_AfterView(spacing float64, view IView)
	VisibilityPriorityForView(view IView) StackViewVisibilityPriority
	SetVisibilityPriority_ForView(priority StackViewVisibilityPriority, view IView)
	SetClippingResistancePriority_ForOrientation(clippingResistancePriority LayoutPriority, orientation LayoutConstraintOrientation)
	SetHuggingPriority_ForOrientation(huggingPriority LayoutPriority, orientation LayoutConstraintOrientation)
	Delegate() StackViewDelegateWrapper
	SetDelegate(value StackViewDelegate)
	ArrangedSubviews() []View
	Views() []View
	DetachedViews() []View
	Orientation() UserInterfaceLayoutOrientation
	SetOrientation(value UserInterfaceLayoutOrientation)
	Alignment() LayoutAttribute
	SetAlignment(value LayoutAttribute)
	Spacing() float64
	SetSpacing(value float64)
	EdgeInsets() foundation.EdgeInsets
	SetEdgeInsets(value foundation.EdgeInsets)
	// deprecated
	HasEqualSpacing() bool
	// deprecated
	SetHasEqualSpacing(value bool)
	Distribution() StackViewDistribution
	SetDistribution(value StackViewDistribution)
	DetachesHiddenViews() bool
	SetDetachesHiddenViews(value bool)
}

type StackView struct {
	View
}

func MakeStackView(ptr unsafe.Pointer) StackView {
	return StackView{
		View: MakeView(ptr),
	}
}

func (sc _StackViewClass) StackViewWithViews(views []IView) StackView {
	rv := ffi.CallMethod[StackView](sc, "stackViewWithViews:", views)
	return rv
}

func (s_ StackView) InitWithFrame(frameRect foundation.Rect) StackView {
	rv := ffi.CallMethod[StackView](s_, "initWithFrame:", frameRect)
	return rv
}

func (s_ StackView) Init() StackView {
	rv := ffi.CallMethod[StackView](s_, "init")
	return rv
}

func (sc _StackViewClass) Alloc() StackView {
	rv := ffi.CallMethod[StackView](sc, "alloc")
	return rv
}

func (sc _StackViewClass) New() StackView {
	rv := ffi.CallMethod[StackView](sc, "new")
	rv.Autorelease()
	return rv
}

func NewStackView() StackView {
	return StackViewClass.New()
}

func (s_ StackView) AddView_InGravity(view IView, gravity StackViewGravity) {
	ffi.CallMethod[ffi.Void](s_, "addView:inGravity:", view, gravity)
}

func (s_ StackView) InsertView_AtIndex_InGravity(view IView, index uint, gravity StackViewGravity) {
	ffi.CallMethod[ffi.Void](s_, "insertView:atIndex:inGravity:", view, index, gravity)
}

func (s_ StackView) SetViews_InGravity(views []IView, gravity StackViewGravity) {
	ffi.CallMethod[ffi.Void](s_, "setViews:inGravity:", views, gravity)
}

func (s_ StackView) RemoveView(view IView) {
	ffi.CallMethod[ffi.Void](s_, "removeView:", view)
}

func (s_ StackView) AddArrangedSubview(view IView) {
	ffi.CallMethod[ffi.Void](s_, "addArrangedSubview:", view)
}

func (s_ StackView) InsertArrangedSubview_AtIndex(view IView, index int) {
	ffi.CallMethod[ffi.Void](s_, "insertArrangedSubview:atIndex:", view, index)
}

func (s_ StackView) RemoveArrangedSubview(view IView) {
	ffi.CallMethod[ffi.Void](s_, "removeArrangedSubview:", view)
}

func (s_ StackView) ViewsInGravity(gravity StackViewGravity) []View {
	rv := ffi.CallMethod[[]View](s_, "viewsInGravity:", gravity)
	return rv
}

func (s_ StackView) ClippingResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := ffi.CallMethod[LayoutPriority](s_, "clippingResistancePriorityForOrientation:", orientation)
	return rv
}

func (s_ StackView) HuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := ffi.CallMethod[LayoutPriority](s_, "huggingPriorityForOrientation:", orientation)
	return rv
}

func (s_ StackView) CustomSpacingAfterView(view IView) float64 {
	rv := ffi.CallMethod[float64](s_, "customSpacingAfterView:", view)
	return rv
}

func (s_ StackView) SetCustomSpacing_AfterView(spacing float64, view IView) {
	ffi.CallMethod[ffi.Void](s_, "setCustomSpacing:afterView:", spacing, view)
}

func (s_ StackView) VisibilityPriorityForView(view IView) StackViewVisibilityPriority {
	rv := ffi.CallMethod[StackViewVisibilityPriority](s_, "visibilityPriorityForView:", view)
	return rv
}

func (s_ StackView) SetVisibilityPriority_ForView(priority StackViewVisibilityPriority, view IView) {
	ffi.CallMethod[ffi.Void](s_, "setVisibilityPriority:forView:", priority, view)
}

func (s_ StackView) SetClippingResistancePriority_ForOrientation(clippingResistancePriority LayoutPriority, orientation LayoutConstraintOrientation) {
	ffi.CallMethod[ffi.Void](s_, "setClippingResistancePriority:forOrientation:", clippingResistancePriority, orientation)
}

func (s_ StackView) SetHuggingPriority_ForOrientation(huggingPriority LayoutPriority, orientation LayoutConstraintOrientation) {
	ffi.CallMethod[ffi.Void](s_, "setHuggingPriority:forOrientation:", huggingPriority, orientation)
}

func (s_ StackView) Delegate() StackViewDelegateWrapper {
	rv := ffi.CallMethod[StackViewDelegateWrapper](s_, "delegate")
	return rv
}

func (s_ StackView) SetDelegate(value StackViewDelegate) {
	po := ffi.CreateProtocol("NSStackViewDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(s_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](s_, "setDelegate:", po)
}

func (s_ StackView) ArrangedSubviews() []View {
	rv := ffi.CallMethod[[]View](s_, "arrangedSubviews")
	return rv
}

func (s_ StackView) Views() []View {
	rv := ffi.CallMethod[[]View](s_, "views")
	return rv
}

func (s_ StackView) DetachedViews() []View {
	rv := ffi.CallMethod[[]View](s_, "detachedViews")
	return rv
}

func (s_ StackView) Orientation() UserInterfaceLayoutOrientation {
	rv := ffi.CallMethod[UserInterfaceLayoutOrientation](s_, "orientation")
	return rv
}

func (s_ StackView) SetOrientation(value UserInterfaceLayoutOrientation) {
	ffi.CallMethod[ffi.Void](s_, "setOrientation:", value)
}

func (s_ StackView) Alignment() LayoutAttribute {
	rv := ffi.CallMethod[LayoutAttribute](s_, "alignment")
	return rv
}

func (s_ StackView) SetAlignment(value LayoutAttribute) {
	ffi.CallMethod[ffi.Void](s_, "setAlignment:", value)
}

func (s_ StackView) Spacing() float64 {
	rv := ffi.CallMethod[float64](s_, "spacing")
	return rv
}

func (s_ StackView) SetSpacing(value float64) {
	ffi.CallMethod[ffi.Void](s_, "setSpacing:", value)
}

func (s_ StackView) EdgeInsets() foundation.EdgeInsets {
	rv := ffi.CallMethod[foundation.EdgeInsets](s_, "edgeInsets")
	return rv
}

func (s_ StackView) SetEdgeInsets(value foundation.EdgeInsets) {
	ffi.CallMethod[ffi.Void](s_, "setEdgeInsets:", value)
}

// deprecated
func (s_ StackView) HasEqualSpacing() bool {
	rv := ffi.CallMethod[bool](s_, "hasEqualSpacing")
	return rv
}

// deprecated
func (s_ StackView) SetHasEqualSpacing(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setHasEqualSpacing:", value)
}

func (s_ StackView) Distribution() StackViewDistribution {
	rv := ffi.CallMethod[StackViewDistribution](s_, "distribution")
	return rv
}

func (s_ StackView) SetDistribution(value StackViewDistribution) {
	ffi.CallMethod[ffi.Void](s_, "setDistribution:", value)
}

func (s_ StackView) DetachesHiddenViews() bool {
	rv := ffi.CallMethod[bool](s_, "detachesHiddenViews")
	return rv
}

func (s_ StackView) SetDetachesHiddenViews(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setDetachesHiddenViews:", value)
}
