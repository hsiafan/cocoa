// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
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
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
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
	rv := objc.CallMethod[SplitView](s_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (s_ SplitView) Init() SplitView {
	rv := objc.CallMethod[SplitView](s_, objc.SEL("init"))
	return rv
}

func (sc _SplitViewClass) Alloc() SplitView {
	rv := objc.CallMethod[SplitView](sc, objc.SEL("alloc"))
	return rv
}

func (sc _SplitViewClass) New() SplitView {
	rv := objc.CallMethod[SplitView](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewSplitView() SplitView {
	return SplitViewClass.New()
}

func (s_ SplitView) AddArrangedSubview(view IView) {
	objc.CallMethod[objc.Void](s_, objc.SEL("addArrangedSubview:"), objc.ExtractPtr(view))
}

func (s_ SplitView) InsertArrangedSubview_AtIndex(view IView, index int) {
	objc.CallMethod[objc.Void](s_, objc.SEL("insertArrangedSubview:atIndex:"), objc.ExtractPtr(view), index)
}

func (s_ SplitView) RemoveArrangedSubview(view IView) {
	objc.CallMethod[objc.Void](s_, objc.SEL("removeArrangedSubview:"), objc.ExtractPtr(view))
}

func (s_ SplitView) AdjustSubviews() {
	objc.CallMethod[objc.Void](s_, objc.SEL("adjustSubviews"))
}

func (s_ SplitView) IsSubviewCollapsed(subview IView) bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("isSubviewCollapsed:"), objc.ExtractPtr(subview))
	return rv
}

func (s_ SplitView) HoldingPriorityForSubviewAtIndex(subviewIndex int) LayoutPriority {
	rv := objc.CallMethod[LayoutPriority](s_, objc.SEL("holdingPriorityForSubviewAtIndex:"), subviewIndex)
	return rv
}

func (s_ SplitView) SetHoldingPriority_ForSubviewAtIndex(priority LayoutPriority, subviewIndex int) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setHoldingPriority:forSubviewAtIndex:"), priority, subviewIndex)
}

func (s_ SplitView) DrawDividerInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](s_, objc.SEL("drawDividerInRect:"), rect)
}

func (s_ SplitView) MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("minPossiblePositionOfDividerAtIndex:"), dividerIndex)
	return rv
}

func (s_ SplitView) MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("maxPossiblePositionOfDividerAtIndex:"), dividerIndex)
	return rv
}

func (s_ SplitView) SetPosition_OfDividerAtIndex(position float64, dividerIndex int) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setPosition:ofDividerAtIndex:"), position, dividerIndex)
}

// deprecated
func (s_ SplitView) IsPaneSplitter() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("isPaneSplitter"))
	return rv
}

// deprecated
func (s_ SplitView) SetIsPaneSplitter(flag bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setIsPaneSplitter:"), flag)
}

// weak property
func (s_ SplitView) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.SEL("delegate"))
	return rv
}

// weak property
func (s_ SplitView) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}

func (s_ SplitView) ArrangesAllSubviews() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("arrangesAllSubviews"))
	return rv
}

func (s_ SplitView) SetArrangesAllSubviews(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setArrangesAllSubviews:"), value)
}

func (s_ SplitView) ArrangedSubviews() []View {
	rv := objc.CallMethod[[]View](s_, objc.SEL("arrangedSubviews"))
	return rv
}

func (s_ SplitView) IsVertical() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("isVertical"))
	return rv
}

func (s_ SplitView) SetVertical(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setVertical:"), value)
}

func (s_ SplitView) DividerStyle() SplitViewDividerStyle {
	rv := objc.CallMethod[SplitViewDividerStyle](s_, objc.SEL("dividerStyle"))
	return rv
}

func (s_ SplitView) SetDividerStyle(value SplitViewDividerStyle) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setDividerStyle:"), value)
}

func (s_ SplitView) DividerColor() Color {
	rv := objc.CallMethod[Color](s_, objc.SEL("dividerColor"))
	return rv
}

func (s_ SplitView) DividerThickness() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("dividerThickness"))
	return rv
}

func (s_ SplitView) AutosaveName() SplitViewAutosaveName {
	rv := objc.CallMethod[SplitViewAutosaveName](s_, objc.SEL("autosaveName"))
	return rv
}

func (s_ SplitView) SetAutosaveName(value SplitViewAutosaveName) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setAutosaveName:"), value)
}
