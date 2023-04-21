// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	SetDelegate0(value objc.IObject)
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
	rv := objc.CallMethod[SplitView](s_, "initWithFrame:", frameRect)
	return rv
}

func (s_ SplitView) Init() SplitView {
	rv := objc.CallMethod[SplitView](s_, "init")
	return rv
}

func (sc _SplitViewClass) Alloc() SplitView {
	rv := objc.CallMethod[SplitView](sc, "alloc")
	return rv
}

func (sc _SplitViewClass) New() SplitView {
	rv := objc.CallMethod[SplitView](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSplitView() SplitView {
	return SplitViewClass.New()
}

func (s_ SplitView) AddArrangedSubview(view IView) {
	objc.CallMethod[objc.Void](s_, "addArrangedSubview:", view)
}

func (s_ SplitView) InsertArrangedSubview_AtIndex(view IView, index int) {
	objc.CallMethod[objc.Void](s_, "insertArrangedSubview:atIndex:", view, index)
}

func (s_ SplitView) RemoveArrangedSubview(view IView) {
	objc.CallMethod[objc.Void](s_, "removeArrangedSubview:", view)
}

func (s_ SplitView) AdjustSubviews() {
	objc.CallMethod[objc.Void](s_, "adjustSubviews")
}

func (s_ SplitView) IsSubviewCollapsed(subview IView) bool {
	rv := objc.CallMethod[bool](s_, "isSubviewCollapsed:", subview)
	return rv
}

func (s_ SplitView) HoldingPriorityForSubviewAtIndex(subviewIndex int) LayoutPriority {
	rv := objc.CallMethod[LayoutPriority](s_, "holdingPriorityForSubviewAtIndex:", subviewIndex)
	return rv
}

func (s_ SplitView) SetHoldingPriority_ForSubviewAtIndex(priority LayoutPriority, subviewIndex int) {
	objc.CallMethod[objc.Void](s_, "setHoldingPriority:forSubviewAtIndex:", priority, subviewIndex)
}

func (s_ SplitView) DrawDividerInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](s_, "drawDividerInRect:", rect)
}

func (s_ SplitView) MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := objc.CallMethod[float64](s_, "minPossiblePositionOfDividerAtIndex:", dividerIndex)
	return rv
}

func (s_ SplitView) MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := objc.CallMethod[float64](s_, "maxPossiblePositionOfDividerAtIndex:", dividerIndex)
	return rv
}

func (s_ SplitView) SetPosition_OfDividerAtIndex(position float64, dividerIndex int) {
	objc.CallMethod[objc.Void](s_, "setPosition:ofDividerAtIndex:", position, dividerIndex)
}

// deprecated
func (s_ SplitView) IsPaneSplitter() bool {
	rv := objc.CallMethod[bool](s_, "isPaneSplitter")
	return rv
}

// deprecated
func (s_ SplitView) SetIsPaneSplitter(flag bool) {
	objc.CallMethod[objc.Void](s_, "setIsPaneSplitter:", flag)
}

func (s_ SplitView) Delegate() SplitViewDelegateWrapper {
	rv := objc.CallMethod[SplitViewDelegateWrapper](s_, "delegate")
	return rv
}

func (s_ SplitView) SetDelegate(value SplitViewDelegate) {
	po := objc.CreateProtocol("NSSplitViewDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(s_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](s_, "setDelegate:", po)
}

func (s_ SplitView) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, "setDelegate:", value)
}

func (s_ SplitView) ArrangesAllSubviews() bool {
	rv := objc.CallMethod[bool](s_, "arrangesAllSubviews")
	return rv
}

func (s_ SplitView) SetArrangesAllSubviews(value bool) {
	objc.CallMethod[objc.Void](s_, "setArrangesAllSubviews:", value)
}

func (s_ SplitView) ArrangedSubviews() []View {
	rv := objc.CallMethod[[]View](s_, "arrangedSubviews")
	return rv
}

func (s_ SplitView) IsVertical() bool {
	rv := objc.CallMethod[bool](s_, "isVertical")
	return rv
}

func (s_ SplitView) SetVertical(value bool) {
	objc.CallMethod[objc.Void](s_, "setVertical:", value)
}

func (s_ SplitView) DividerStyle() SplitViewDividerStyle {
	rv := objc.CallMethod[SplitViewDividerStyle](s_, "dividerStyle")
	return rv
}

func (s_ SplitView) SetDividerStyle(value SplitViewDividerStyle) {
	objc.CallMethod[objc.Void](s_, "setDividerStyle:", value)
}

func (s_ SplitView) DividerColor() Color {
	rv := objc.CallMethod[Color](s_, "dividerColor")
	return rv
}

func (s_ SplitView) DividerThickness() float64 {
	rv := objc.CallMethod[float64](s_, "dividerThickness")
	return rv
}

func (s_ SplitView) AutosaveName() SplitViewAutosaveName {
	rv := objc.CallMethod[SplitViewAutosaveName](s_, "autosaveName")
	return rv
}

func (s_ SplitView) SetAutosaveName(value SplitViewAutosaveName) {
	objc.CallMethod[objc.Void](s_, "setAutosaveName:", value)
}
