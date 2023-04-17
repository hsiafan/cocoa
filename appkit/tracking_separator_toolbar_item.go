// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[TrackingSeparatorToolbarItem](tc, "trackingSeparatorToolbarItemWithIdentifier:splitView:dividerIndex:", identifier, splitView, dividerIndex)
	return rv
}

func (t_ TrackingSeparatorToolbarItem) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) TrackingSeparatorToolbarItem {
	rv := ffi.CallMethod[TrackingSeparatorToolbarItem](t_, "initWithItemIdentifier:", itemIdentifier)
	return rv
}

func (tc _TrackingSeparatorToolbarItemClass) Alloc() TrackingSeparatorToolbarItem {
	rv := ffi.CallMethod[TrackingSeparatorToolbarItem](tc, "alloc")
	return rv
}

func (tc _TrackingSeparatorToolbarItemClass) New() TrackingSeparatorToolbarItem {
	rv := ffi.CallMethod[TrackingSeparatorToolbarItem](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTrackingSeparatorToolbarItem() TrackingSeparatorToolbarItem {
	return TrackingSeparatorToolbarItemClass.New()
}

func (t_ TrackingSeparatorToolbarItem) Init() TrackingSeparatorToolbarItem {
	rv := ffi.CallMethod[TrackingSeparatorToolbarItem](t_, "init")
	return rv
}

func (t_ TrackingSeparatorToolbarItem) DividerIndex() int {
	rv := ffi.CallMethod[int](t_, "dividerIndex")
	return rv
}

func (t_ TrackingSeparatorToolbarItem) SetDividerIndex(value int) {
	ffi.CallMethod[ffi.Void](t_, "setDividerIndex:", value)
}

func (t_ TrackingSeparatorToolbarItem) SplitView() SplitView {
	rv := ffi.CallMethod[SplitView](t_, "splitView")
	return rv
}

func (t_ TrackingSeparatorToolbarItem) SetSplitView(value ISplitView) {
	ffi.CallMethod[ffi.Void](t_, "setSplitView:", value)
}
