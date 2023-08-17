// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

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

func WrapCandidateListTouchBarItemDelegate(v CandidateListTouchBarItemDelegate) objc.Object {
	return objc.WrapAsProtocol("NSCandidateListTouchBarItemDelegate", v)
}

type CandidateListTouchBarItemDelegateBase struct {
}

func (p *CandidateListTouchBarItemDelegateBase) ImplementsCandidateListTouchBarItem_BeginSelectingCandidateAtIndex() bool {
	return false
}

func (p *CandidateListTouchBarItemDelegateBase) CandidateListTouchBarItem_BeginSelectingCandidateAtIndex(anItem CandidateListTouchBarItem, index int) {
	panic("unimpemented protocol method")
}

func (p *CandidateListTouchBarItemDelegateBase) ImplementsCandidateListTouchBarItem_ChangeSelectionFromCandidateAtIndex_ToIndex() bool {
	return false
}

func (p *CandidateListTouchBarItemDelegateBase) CandidateListTouchBarItem_ChangeSelectionFromCandidateAtIndex_ToIndex(anItem CandidateListTouchBarItem, previousIndex int, index int) {
	panic("unimpemented protocol method")
}

func (p *CandidateListTouchBarItemDelegateBase) ImplementsCandidateListTouchBarItem_EndSelectingCandidateAtIndex() bool {
	return false
}

func (p *CandidateListTouchBarItemDelegateBase) CandidateListTouchBarItem_EndSelectingCandidateAtIndex(anItem CandidateListTouchBarItem, index int) {
	panic("unimpemented protocol method")
}

func (p *CandidateListTouchBarItemDelegateBase) ImplementsCandidateListTouchBarItem_ChangedCandidateListVisibility() bool {
	return false
}

func (p *CandidateListTouchBarItemDelegateBase) CandidateListTouchBarItem_ChangedCandidateListVisibility(anItem CandidateListTouchBarItem, isVisible bool) {
	panic("unimpemented protocol method")
}

type CandidateListTouchBarItemDelegateWrapper struct {
	objc.Object
}

func (c_ CandidateListTouchBarItemDelegateWrapper) CandidateListTouchBarItem_BeginSelectingCandidateAtIndex(anItem ICandidateListTouchBarItem, index int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("candidateListTouchBarItem:beginSelectingCandidateAtIndex:"), objc.ExtractPtr(anItem), index)
}

func (c_ CandidateListTouchBarItemDelegateWrapper) CandidateListTouchBarItem_ChangeSelectionFromCandidateAtIndex_ToIndex(anItem ICandidateListTouchBarItem, previousIndex int, index int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("candidateListTouchBarItem:changeSelectionFromCandidateAtIndex:toIndex:"), objc.ExtractPtr(anItem), previousIndex, index)
}

func (c_ CandidateListTouchBarItemDelegateWrapper) CandidateListTouchBarItem_EndSelectingCandidateAtIndex(anItem ICandidateListTouchBarItem, index int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("candidateListTouchBarItem:endSelectingCandidateAtIndex:"), objc.ExtractPtr(anItem), index)
}

func (c_ CandidateListTouchBarItemDelegateWrapper) CandidateListTouchBarItem_ChangedCandidateListVisibility(anItem ICandidateListTouchBarItem, isVisible bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("candidateListTouchBarItem:changedCandidateListVisibility:"), objc.ExtractPtr(anItem), isVisible)
}
