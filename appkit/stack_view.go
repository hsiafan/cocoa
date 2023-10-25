// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
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
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
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
	rv := objc.CallMethod[StackView](sc, objc.SEL("stackViewWithViews:"), views)
	return rv
}

func (s_ StackView) InitWithFrame(frameRect foundation.Rect) StackView {
	rv := objc.CallMethod[StackView](s_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (s_ StackView) Init() StackView {
	rv := objc.CallMethod[StackView](s_, objc.SEL("init"))
	return rv
}

func (sc _StackViewClass) Alloc() StackView {
	rv := objc.CallMethod[StackView](sc, objc.SEL("alloc"))
	return rv
}

func (sc _StackViewClass) New() StackView {
	rv := objc.CallMethod[StackView](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewStackView() StackView {
	return StackViewClass.New()
}

func (s_ StackView) AddView_InGravity(view IView, gravity StackViewGravity) {
	objc.CallMethod[objc.Void](s_, objc.SEL("addView:inGravity:"), objc.ExtractPtr(view), gravity)
}

func (s_ StackView) InsertView_AtIndex_InGravity(view IView, index uint, gravity StackViewGravity) {
	objc.CallMethod[objc.Void](s_, objc.SEL("insertView:atIndex:inGravity:"), objc.ExtractPtr(view), index, gravity)
}

func (s_ StackView) SetViews_InGravity(views []IView, gravity StackViewGravity) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setViews:inGravity:"), views, gravity)
}

func (s_ StackView) RemoveView(view IView) {
	objc.CallMethod[objc.Void](s_, objc.SEL("removeView:"), objc.ExtractPtr(view))
}

func (s_ StackView) AddArrangedSubview(view IView) {
	objc.CallMethod[objc.Void](s_, objc.SEL("addArrangedSubview:"), objc.ExtractPtr(view))
}

func (s_ StackView) InsertArrangedSubview_AtIndex(view IView, index int) {
	objc.CallMethod[objc.Void](s_, objc.SEL("insertArrangedSubview:atIndex:"), objc.ExtractPtr(view), index)
}

func (s_ StackView) RemoveArrangedSubview(view IView) {
	objc.CallMethod[objc.Void](s_, objc.SEL("removeArrangedSubview:"), objc.ExtractPtr(view))
}

func (s_ StackView) ViewsInGravity(gravity StackViewGravity) []View {
	rv := objc.CallMethod[[]View](s_, objc.SEL("viewsInGravity:"), gravity)
	return rv
}

func (s_ StackView) ClippingResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := objc.CallMethod[LayoutPriority](s_, objc.SEL("clippingResistancePriorityForOrientation:"), orientation)
	return rv
}

func (s_ StackView) HuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := objc.CallMethod[LayoutPriority](s_, objc.SEL("huggingPriorityForOrientation:"), orientation)
	return rv
}

func (s_ StackView) CustomSpacingAfterView(view IView) float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("customSpacingAfterView:"), objc.ExtractPtr(view))
	return rv
}

func (s_ StackView) SetCustomSpacing_AfterView(spacing float64, view IView) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setCustomSpacing:afterView:"), spacing, objc.ExtractPtr(view))
}

func (s_ StackView) VisibilityPriorityForView(view IView) StackViewVisibilityPriority {
	rv := objc.CallMethod[StackViewVisibilityPriority](s_, objc.SEL("visibilityPriorityForView:"), objc.ExtractPtr(view))
	return rv
}

func (s_ StackView) SetVisibilityPriority_ForView(priority StackViewVisibilityPriority, view IView) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setVisibilityPriority:forView:"), priority, objc.ExtractPtr(view))
}

func (s_ StackView) SetClippingResistancePriority_ForOrientation(clippingResistancePriority LayoutPriority, orientation LayoutConstraintOrientation) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setClippingResistancePriority:forOrientation:"), clippingResistancePriority, orientation)
}

func (s_ StackView) SetHuggingPriority_ForOrientation(huggingPriority LayoutPriority, orientation LayoutConstraintOrientation) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setHuggingPriority:forOrientation:"), huggingPriority, orientation)
}

// weak property
func (s_ StackView) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.SEL("delegate"))
	return rv
}

// weak property
func (s_ StackView) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}

func (s_ StackView) ArrangedSubviews() []View {
	rv := objc.CallMethod[[]View](s_, objc.SEL("arrangedSubviews"))
	return rv
}

func (s_ StackView) Views() []View {
	rv := objc.CallMethod[[]View](s_, objc.SEL("views"))
	return rv
}

func (s_ StackView) DetachedViews() []View {
	rv := objc.CallMethod[[]View](s_, objc.SEL("detachedViews"))
	return rv
}

func (s_ StackView) Orientation() UserInterfaceLayoutOrientation {
	rv := objc.CallMethod[UserInterfaceLayoutOrientation](s_, objc.SEL("orientation"))
	return rv
}

func (s_ StackView) SetOrientation(value UserInterfaceLayoutOrientation) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setOrientation:"), value)
}

func (s_ StackView) Alignment() LayoutAttribute {
	rv := objc.CallMethod[LayoutAttribute](s_, objc.SEL("alignment"))
	return rv
}

func (s_ StackView) SetAlignment(value LayoutAttribute) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setAlignment:"), value)
}

func (s_ StackView) Spacing() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("spacing"))
	return rv
}

func (s_ StackView) SetSpacing(value float64) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setSpacing:"), value)
}

func (s_ StackView) EdgeInsets() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](s_, objc.SEL("edgeInsets"))
	return rv
}

func (s_ StackView) SetEdgeInsets(value foundation.EdgeInsets) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setEdgeInsets:"), value)
}

// deprecated
func (s_ StackView) HasEqualSpacing() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("hasEqualSpacing"))
	return rv
}

// deprecated
func (s_ StackView) SetHasEqualSpacing(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setHasEqualSpacing:"), value)
}

func (s_ StackView) Distribution() StackViewDistribution {
	rv := objc.CallMethod[StackViewDistribution](s_, objc.SEL("distribution"))
	return rv
}

func (s_ StackView) SetDistribution(value StackViewDistribution) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setDistribution:"), value)
}

func (s_ StackView) DetachesHiddenViews() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("detachesHiddenViews"))
	return rv
}

func (s_ StackView) SetDetachesHiddenViews(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setDetachesHiddenViews:"), value)
}
