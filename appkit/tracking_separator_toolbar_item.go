// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var TrackingSeparatorToolbarItemClass = _TrackingSeparatorToolbarItemClass{objc.GetClass("NSTrackingSeparatorToolbarItem")}

type _TrackingSeparatorToolbarItemClass struct {
	objc.Class
}

type ITrackingSeparatorToolbarItem interface {
	IToolbarItem
	DividerIndex() int
	SetDividerIndex(value int)
	SplitView() SplitView
	SetSplitView(value ISplitView)
}

type TrackingSeparatorToolbarItem struct {
	ToolbarItem
}

func MakeTrackingSeparatorToolbarItem(ptr unsafe.Pointer) TrackingSeparatorToolbarItem {
	return TrackingSeparatorToolbarItem{
		ToolbarItem: MakeToolbarItem(ptr),
	}
}

func (tc _TrackingSeparatorToolbarItemClass) TrackingSeparatorToolbarItemWithIdentifier_SplitView_DividerIndex(identifier ToolbarItemIdentifier, splitView ISplitView, dividerIndex int) TrackingSeparatorToolbarItem {
	rv := objc.CallMethod[TrackingSeparatorToolbarItem](tc, objc.SEL("trackingSeparatorToolbarItemWithIdentifier:splitView:dividerIndex:"), identifier, objc.ExtractPtr(splitView), dividerIndex)
	return rv
}

func (t_ TrackingSeparatorToolbarItem) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) TrackingSeparatorToolbarItem {
	rv := objc.CallMethod[TrackingSeparatorToolbarItem](t_, objc.SEL("initWithItemIdentifier:"), itemIdentifier)
	return rv
}

func (tc _TrackingSeparatorToolbarItemClass) Alloc() TrackingSeparatorToolbarItem {
	rv := objc.CallMethod[TrackingSeparatorToolbarItem](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TrackingSeparatorToolbarItemClass) New() TrackingSeparatorToolbarItem {
	rv := objc.CallMethod[TrackingSeparatorToolbarItem](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTrackingSeparatorToolbarItem() TrackingSeparatorToolbarItem {
	return TrackingSeparatorToolbarItemClass.New()
}

func (t_ TrackingSeparatorToolbarItem) Init() TrackingSeparatorToolbarItem {
	rv := objc.CallMethod[TrackingSeparatorToolbarItem](t_, objc.SEL("init"))
	return rv
}

func (t_ TrackingSeparatorToolbarItem) DividerIndex() int {
	rv := objc.CallMethod[int](t_, objc.SEL("dividerIndex"))
	return rv
}

func (t_ TrackingSeparatorToolbarItem) SetDividerIndex(value int) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDividerIndex:"), value)
}

func (t_ TrackingSeparatorToolbarItem) SplitView() SplitView {
	rv := objc.CallMethod[SplitView](t_, objc.SEL("splitView"))
	return rv
}

func (t_ TrackingSeparatorToolbarItem) SetSplitView(value ISplitView) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSplitView:"), objc.ExtractPtr(value))
}
