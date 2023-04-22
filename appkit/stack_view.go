// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	SetDelegate0(value objc.IObject)
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
	rv := objc.CallMethod[StackView](sc, objc.GetSelector("stackViewWithViews:"), views)
	return rv
}

func (s_ StackView) InitWithFrame(frameRect foundation.Rect) StackView {
	rv := objc.CallMethod[StackView](s_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func (s_ StackView) Init() StackView {
	rv := objc.CallMethod[StackView](s_, objc.GetSelector("init"))
	return rv
}

func (sc _StackViewClass) Alloc() StackView {
	rv := objc.CallMethod[StackView](sc, objc.GetSelector("alloc"))
	return rv
}

func (sc _StackViewClass) New() StackView {
	rv := objc.CallMethod[StackView](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewStackView() StackView {
	return StackViewClass.New()
}

func (s_ StackView) AddView_InGravity(view IView, gravity StackViewGravity) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("addView:inGravity:"), view, gravity)
}

func (s_ StackView) InsertView_AtIndex_InGravity(view IView, index uint, gravity StackViewGravity) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("insertView:atIndex:inGravity:"), view, index, gravity)
}

func (s_ StackView) SetViews_InGravity(views []IView, gravity StackViewGravity) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setViews:inGravity:"), views, gravity)
}

func (s_ StackView) RemoveView(view IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("removeView:"), view)
}

func (s_ StackView) AddArrangedSubview(view IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("addArrangedSubview:"), view)
}

func (s_ StackView) InsertArrangedSubview_AtIndex(view IView, index int) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("insertArrangedSubview:atIndex:"), view, index)
}

func (s_ StackView) RemoveArrangedSubview(view IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("removeArrangedSubview:"), view)
}

func (s_ StackView) ViewsInGravity(gravity StackViewGravity) []View {
	rv := objc.CallMethod[[]View](s_, objc.GetSelector("viewsInGravity:"), gravity)
	return rv
}

func (s_ StackView) ClippingResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := objc.CallMethod[LayoutPriority](s_, objc.GetSelector("clippingResistancePriorityForOrientation:"), orientation)
	return rv
}

func (s_ StackView) HuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := objc.CallMethod[LayoutPriority](s_, objc.GetSelector("huggingPriorityForOrientation:"), orientation)
	return rv
}

func (s_ StackView) CustomSpacingAfterView(view IView) float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("customSpacingAfterView:"), view)
	return rv
}

func (s_ StackView) SetCustomSpacing_AfterView(spacing float64, view IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setCustomSpacing:afterView:"), spacing, view)
}

func (s_ StackView) VisibilityPriorityForView(view IView) StackViewVisibilityPriority {
	rv := objc.CallMethod[StackViewVisibilityPriority](s_, objc.GetSelector("visibilityPriorityForView:"), view)
	return rv
}

func (s_ StackView) SetVisibilityPriority_ForView(priority StackViewVisibilityPriority, view IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setVisibilityPriority:forView:"), priority, view)
}

func (s_ StackView) SetClippingResistancePriority_ForOrientation(clippingResistancePriority LayoutPriority, orientation LayoutConstraintOrientation) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setClippingResistancePriority:forOrientation:"), clippingResistancePriority, orientation)
}

func (s_ StackView) SetHuggingPriority_ForOrientation(huggingPriority LayoutPriority, orientation LayoutConstraintOrientation) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setHuggingPriority:forOrientation:"), huggingPriority, orientation)
}

func (s_ StackView) Delegate() StackViewDelegateWrapper {
	rv := objc.CallMethod[StackViewDelegateWrapper](s_, objc.GetSelector("delegate"))
	return rv
}

func (s_ StackView) SetDelegate(value StackViewDelegate) {
	po := objc.CreateProtocol("NSStackViewDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(s_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDelegate:"), po)
}

func (s_ StackView) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDelegate:"), value)
}

func (s_ StackView) ArrangedSubviews() []View {
	rv := objc.CallMethod[[]View](s_, objc.GetSelector("arrangedSubviews"))
	return rv
}

func (s_ StackView) Views() []View {
	rv := objc.CallMethod[[]View](s_, objc.GetSelector("views"))
	return rv
}

func (s_ StackView) DetachedViews() []View {
	rv := objc.CallMethod[[]View](s_, objc.GetSelector("detachedViews"))
	return rv
}

func (s_ StackView) Orientation() UserInterfaceLayoutOrientation {
	rv := objc.CallMethod[UserInterfaceLayoutOrientation](s_, objc.GetSelector("orientation"))
	return rv
}

func (s_ StackView) SetOrientation(value UserInterfaceLayoutOrientation) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setOrientation:"), value)
}

func (s_ StackView) Alignment() LayoutAttribute {
	rv := objc.CallMethod[LayoutAttribute](s_, objc.GetSelector("alignment"))
	return rv
}

func (s_ StackView) SetAlignment(value LayoutAttribute) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAlignment:"), value)
}

func (s_ StackView) Spacing() float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("spacing"))
	return rv
}

func (s_ StackView) SetSpacing(value float64) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSpacing:"), value)
}

func (s_ StackView) EdgeInsets() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](s_, objc.GetSelector("edgeInsets"))
	return rv
}

func (s_ StackView) SetEdgeInsets(value foundation.EdgeInsets) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setEdgeInsets:"), value)
}

// deprecated
func (s_ StackView) HasEqualSpacing() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("hasEqualSpacing"))
	return rv
}

// deprecated
func (s_ StackView) SetHasEqualSpacing(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setHasEqualSpacing:"), value)
}

func (s_ StackView) Distribution() StackViewDistribution {
	rv := objc.CallMethod[StackViewDistribution](s_, objc.GetSelector("distribution"))
	return rv
}

func (s_ StackView) SetDistribution(value StackViewDistribution) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDistribution:"), value)
}

func (s_ StackView) DetachesHiddenViews() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("detachesHiddenViews"))
	return rv
}

func (s_ StackView) SetDetachesHiddenViews(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDetachesHiddenViews:"), value)
}
