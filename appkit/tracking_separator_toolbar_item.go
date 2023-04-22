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
	rv := objc.CallMethod[TrackingSeparatorToolbarItem](tc, objc.GetSelector("trackingSeparatorToolbarItemWithIdentifier:splitView:dividerIndex:"), identifier, splitView, dividerIndex)
	return rv
}

func (t_ TrackingSeparatorToolbarItem) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) TrackingSeparatorToolbarItem {
	rv := objc.CallMethod[TrackingSeparatorToolbarItem](t_, objc.GetSelector("initWithItemIdentifier:"), itemIdentifier)
	return rv
}

func (tc _TrackingSeparatorToolbarItemClass) Alloc() TrackingSeparatorToolbarItem {
	rv := objc.CallMethod[TrackingSeparatorToolbarItem](tc, objc.GetSelector("alloc"))
	return rv
}

func (tc _TrackingSeparatorToolbarItemClass) New() TrackingSeparatorToolbarItem {
	rv := objc.CallMethod[TrackingSeparatorToolbarItem](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTrackingSeparatorToolbarItem() TrackingSeparatorToolbarItem {
	return TrackingSeparatorToolbarItemClass.New()
}

func (t_ TrackingSeparatorToolbarItem) Init() TrackingSeparatorToolbarItem {
	rv := objc.CallMethod[TrackingSeparatorToolbarItem](t_, objc.GetSelector("init"))
	return rv
}

func (t_ TrackingSeparatorToolbarItem) DividerIndex() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("dividerIndex"))
	return rv
}

func (t_ TrackingSeparatorToolbarItem) SetDividerIndex(value int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDividerIndex:"), value)
}

func (t_ TrackingSeparatorToolbarItem) SplitView() SplitView {
	rv := objc.CallMethod[SplitView](t_, objc.GetSelector("splitView"))
	return rv
}

func (t_ TrackingSeparatorToolbarItem) SetSplitView(value ISplitView) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSplitView:"), value)
}
