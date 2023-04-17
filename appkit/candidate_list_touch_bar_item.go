// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var CandidateListTouchBarItemClass = _CandidateListTouchBarItemClass{objc.GetClass("NSCandidateListTouchBarItem")}

type _CandidateListTouchBarItemClass struct {
	objc.Class
}

type ICandidateListTouchBarItem interface {
	ITouchBarItem
	UpdateWithInsertionPointVisibility(isVisible bool)
	Client() View
	SetClient(value IView)
	Delegate() CandidateListTouchBarItemDelegateWrapper
	SetDelegate(value CandidateListTouchBarItemDelegate)
	AllowsTextInputContextCandidates() bool
	SetAllowsTextInputContextCandidates(value bool)
	AllowsCollapsing() bool
	SetAllowsCollapsing(value bool)
	IsCollapsed() bool
	SetCollapsed(value bool)
	IsCandidateListVisible() bool
	SetCustomizationLabel(value string)
}

type CandidateListTouchBarItem struct {
	TouchBarItem
}

func MakeCandidateListTouchBarItem(ptr unsafe.Pointer) CandidateListTouchBarItem {
	return CandidateListTouchBarItem{
		TouchBarItem: MakeTouchBarItem(ptr),
	}
}

func (c_ CandidateListTouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) CandidateListTouchBarItem {
	rv := ffi.CallMethod[CandidateListTouchBarItem](c_, "initWithIdentifier:", identifier)
	return rv
}

func (cc _CandidateListTouchBarItemClass) Alloc() CandidateListTouchBarItem {
	rv := ffi.CallMethod[CandidateListTouchBarItem](cc, "alloc")
	return rv
}

func (cc _CandidateListTouchBarItemClass) New() CandidateListTouchBarItem {
	rv := ffi.CallMethod[CandidateListTouchBarItem](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCandidateListTouchBarItem() CandidateListTouchBarItem {
	return CandidateListTouchBarItemClass.New()
}

func (c_ CandidateListTouchBarItem) Init() CandidateListTouchBarItem {
	rv := ffi.CallMethod[CandidateListTouchBarItem](c_, "init")
	return rv
}

func (c_ CandidateListTouchBarItem) UpdateWithInsertionPointVisibility(isVisible bool) {
	ffi.CallMethod[ffi.Void](c_, "updateWithInsertionPointVisibility:", isVisible)
}

func (c_ CandidateListTouchBarItem) Client() View {
	rv := ffi.CallMethod[View](c_, "client")
	return rv
}

func (c_ CandidateListTouchBarItem) SetClient(value IView) {
	ffi.CallMethod[ffi.Void](c_, "setClient:", value)
}

func (c_ CandidateListTouchBarItem) Delegate() CandidateListTouchBarItemDelegateWrapper {
	rv := ffi.CallMethod[CandidateListTouchBarItemDelegateWrapper](c_, "delegate")
	return rv
}

func (c_ CandidateListTouchBarItem) SetDelegate(value CandidateListTouchBarItemDelegate) {
	po := ffi.CreateProtocol("NSCandidateListTouchBarItemDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(c_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](c_, "setDelegate:", po)
}

func (c_ CandidateListTouchBarItem) AllowsTextInputContextCandidates() bool {
	rv := ffi.CallMethod[bool](c_, "allowsTextInputContextCandidates")
	return rv
}

func (c_ CandidateListTouchBarItem) SetAllowsTextInputContextCandidates(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setAllowsTextInputContextCandidates:", value)
}

func (c_ CandidateListTouchBarItem) AllowsCollapsing() bool {
	rv := ffi.CallMethod[bool](c_, "allowsCollapsing")
	return rv
}

func (c_ CandidateListTouchBarItem) SetAllowsCollapsing(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setAllowsCollapsing:", value)
}

func (c_ CandidateListTouchBarItem) IsCollapsed() bool {
	rv := ffi.CallMethod[bool](c_, "isCollapsed")
	return rv
}

func (c_ CandidateListTouchBarItem) SetCollapsed(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setCollapsed:", value)
}

func (c_ CandidateListTouchBarItem) IsCandidateListVisible() bool {
	rv := ffi.CallMethod[bool](c_, "isCandidateListVisible")
	return rv
}

func (c_ CandidateListTouchBarItem) SetCustomizationLabel(value string) {
	ffi.CallMethod[ffi.Void](c_, "setCustomizationLabel:", value)
}
