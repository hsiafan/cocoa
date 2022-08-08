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

func (c_ CandidateListTouchBarItem) Init() CandidateListTouchBarItem {
	rv := ffi.CallMethod[CandidateListTouchBarItem](c_, "init")
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
	po := ffi.CreateProtocol(value)
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

type CandidateListTouchBarItemDelegate interface {
	ImplementsCandidateListTouchBarItem_BeginSelectingCandidateAtIndex() bool
	// optional
	CandidateListTouchBarItem_BeginSelectingCandidateAtIndex(anItem CandidateListTouchBarItem, index int)
	ImplementsCandidateListTouchBarItem_ChangeSelectionFromCandidateAtIndex_ToIndex() bool
	// optional
	CandidateListTouchBarItem_ChangeSelectionFromCandidateAtIndex_ToIndex(anItem CandidateListTouchBarItem, previousIndex int, index int)
	ImplementsCandidateListTouchBarItem_EndSelectingCandidateAtIndex() bool
	// optional
	CandidateListTouchBarItem_EndSelectingCandidateAtIndex(anItem CandidateListTouchBarItem, index int)
	ImplementsCandidateListTouchBarItem_ChangedCandidateListVisibility() bool
	// optional
	CandidateListTouchBarItem_ChangedCandidateListVisibility(anItem CandidateListTouchBarItem, isVisible bool)
}

type CandidateListTouchBarItemDelegateImpl struct {
	_CandidateListTouchBarItem_BeginSelectingCandidateAtIndex              func(anItem CandidateListTouchBarItem, index int)
	_CandidateListTouchBarItem_ChangeSelectionFromCandidateAtIndex_ToIndex func(anItem CandidateListTouchBarItem, previousIndex int, index int)
	_CandidateListTouchBarItem_EndSelectingCandidateAtIndex                func(anItem CandidateListTouchBarItem, index int)
	_CandidateListTouchBarItem_ChangedCandidateListVisibility              func(anItem CandidateListTouchBarItem, isVisible bool)
}

func (di *CandidateListTouchBarItemDelegateImpl) ImplementsCandidateListTouchBarItem_BeginSelectingCandidateAtIndex() bool {
	return di._CandidateListTouchBarItem_BeginSelectingCandidateAtIndex != nil
}

func (di *CandidateListTouchBarItemDelegateImpl) SetCandidateListTouchBarItem_BeginSelectingCandidateAtIndex(f func(anItem CandidateListTouchBarItem, index int)) {
	di._CandidateListTouchBarItem_BeginSelectingCandidateAtIndex = f
}

func (di *CandidateListTouchBarItemDelegateImpl) CandidateListTouchBarItem_BeginSelectingCandidateAtIndex(anItem CandidateListTouchBarItem, index int) {
	di._CandidateListTouchBarItem_BeginSelectingCandidateAtIndex(anItem, index)
}
func (di *CandidateListTouchBarItemDelegateImpl) ImplementsCandidateListTouchBarItem_ChangeSelectionFromCandidateAtIndex_ToIndex() bool {
	return di._CandidateListTouchBarItem_ChangeSelectionFromCandidateAtIndex_ToIndex != nil
}

func (di *CandidateListTouchBarItemDelegateImpl) SetCandidateListTouchBarItem_ChangeSelectionFromCandidateAtIndex_ToIndex(f func(anItem CandidateListTouchBarItem, previousIndex int, index int)) {
	di._CandidateListTouchBarItem_ChangeSelectionFromCandidateAtIndex_ToIndex = f
}

func (di *CandidateListTouchBarItemDelegateImpl) CandidateListTouchBarItem_ChangeSelectionFromCandidateAtIndex_ToIndex(anItem CandidateListTouchBarItem, previousIndex int, index int) {
	di._CandidateListTouchBarItem_ChangeSelectionFromCandidateAtIndex_ToIndex(anItem, previousIndex, index)
}
func (di *CandidateListTouchBarItemDelegateImpl) ImplementsCandidateListTouchBarItem_EndSelectingCandidateAtIndex() bool {
	return di._CandidateListTouchBarItem_EndSelectingCandidateAtIndex != nil
}

func (di *CandidateListTouchBarItemDelegateImpl) SetCandidateListTouchBarItem_EndSelectingCandidateAtIndex(f func(anItem CandidateListTouchBarItem, index int)) {
	di._CandidateListTouchBarItem_EndSelectingCandidateAtIndex = f
}

func (di *CandidateListTouchBarItemDelegateImpl) CandidateListTouchBarItem_EndSelectingCandidateAtIndex(anItem CandidateListTouchBarItem, index int) {
	di._CandidateListTouchBarItem_EndSelectingCandidateAtIndex(anItem, index)
}
func (di *CandidateListTouchBarItemDelegateImpl) ImplementsCandidateListTouchBarItem_ChangedCandidateListVisibility() bool {
	return di._CandidateListTouchBarItem_ChangedCandidateListVisibility != nil
}

func (di *CandidateListTouchBarItemDelegateImpl) SetCandidateListTouchBarItem_ChangedCandidateListVisibility(f func(anItem CandidateListTouchBarItem, isVisible bool)) {
	di._CandidateListTouchBarItem_ChangedCandidateListVisibility = f
}

func (di *CandidateListTouchBarItemDelegateImpl) CandidateListTouchBarItem_ChangedCandidateListVisibility(anItem CandidateListTouchBarItem, isVisible bool) {
	di._CandidateListTouchBarItem_ChangedCandidateListVisibility(anItem, isVisible)
}

type CandidateListTouchBarItemDelegateWrapper struct {
	objc.Object
}

func (c_ *CandidateListTouchBarItemDelegateWrapper) ImplementsCandidateListTouchBarItem_BeginSelectingCandidateAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("candidateListTouchBarItem:beginSelectingCandidateAtIndex:"))
}

func (c_ CandidateListTouchBarItemDelegateWrapper) CandidateListTouchBarItem_BeginSelectingCandidateAtIndex(anItem ICandidateListTouchBarItem, index int) {
	ffi.CallMethod[ffi.Void](c_, "candidateListTouchBarItem:beginSelectingCandidateAtIndex:", anItem, index)
}

func (c_ *CandidateListTouchBarItemDelegateWrapper) ImplementsCandidateListTouchBarItem_ChangeSelectionFromCandidateAtIndex_ToIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("candidateListTouchBarItem:changeSelectionFromCandidateAtIndex:toIndex:"))
}

func (c_ CandidateListTouchBarItemDelegateWrapper) CandidateListTouchBarItem_ChangeSelectionFromCandidateAtIndex_ToIndex(anItem ICandidateListTouchBarItem, previousIndex int, index int) {
	ffi.CallMethod[ffi.Void](c_, "candidateListTouchBarItem:changeSelectionFromCandidateAtIndex:toIndex:", anItem, previousIndex, index)
}

func (c_ *CandidateListTouchBarItemDelegateWrapper) ImplementsCandidateListTouchBarItem_EndSelectingCandidateAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("candidateListTouchBarItem:endSelectingCandidateAtIndex:"))
}

func (c_ CandidateListTouchBarItemDelegateWrapper) CandidateListTouchBarItem_EndSelectingCandidateAtIndex(anItem ICandidateListTouchBarItem, index int) {
	ffi.CallMethod[ffi.Void](c_, "candidateListTouchBarItem:endSelectingCandidateAtIndex:", anItem, index)
}

func (c_ *CandidateListTouchBarItemDelegateWrapper) ImplementsCandidateListTouchBarItem_ChangedCandidateListVisibility() bool {
	return c_.RespondsToSelector(objc.GetSelector("candidateListTouchBarItem:changedCandidateListVisibility:"))
}

func (c_ CandidateListTouchBarItemDelegateWrapper) CandidateListTouchBarItem_ChangedCandidateListVisibility(anItem ICandidateListTouchBarItem, isVisible bool) {
	ffi.CallMethod[ffi.Void](c_, "candidateListTouchBarItem:changedCandidateListVisibility:", anItem, isVisible)
}
