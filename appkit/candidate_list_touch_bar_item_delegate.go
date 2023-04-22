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
	objc.CallMethod[objc.Void](c_, objc.GetSelector("candidateListTouchBarItem:beginSelectingCandidateAtIndex:"), anItem, index)
}

func (c_ *CandidateListTouchBarItemDelegateWrapper) ImplementsCandidateListTouchBarItem_ChangeSelectionFromCandidateAtIndex_ToIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("candidateListTouchBarItem:changeSelectionFromCandidateAtIndex:toIndex:"))
}

func (c_ CandidateListTouchBarItemDelegateWrapper) CandidateListTouchBarItem_ChangeSelectionFromCandidateAtIndex_ToIndex(anItem ICandidateListTouchBarItem, previousIndex int, index int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("candidateListTouchBarItem:changeSelectionFromCandidateAtIndex:toIndex:"), anItem, previousIndex, index)
}

func (c_ *CandidateListTouchBarItemDelegateWrapper) ImplementsCandidateListTouchBarItem_EndSelectingCandidateAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("candidateListTouchBarItem:endSelectingCandidateAtIndex:"))
}

func (c_ CandidateListTouchBarItemDelegateWrapper) CandidateListTouchBarItem_EndSelectingCandidateAtIndex(anItem ICandidateListTouchBarItem, index int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("candidateListTouchBarItem:endSelectingCandidateAtIndex:"), anItem, index)
}

func (c_ *CandidateListTouchBarItemDelegateWrapper) ImplementsCandidateListTouchBarItem_ChangedCandidateListVisibility() bool {
	return c_.RespondsToSelector(objc.GetSelector("candidateListTouchBarItem:changedCandidateListVisibility:"))
}

func (c_ CandidateListTouchBarItemDelegateWrapper) CandidateListTouchBarItem_ChangedCandidateListVisibility(anItem ICandidateListTouchBarItem, isVisible bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("candidateListTouchBarItem:changedCandidateListVisibility:"), anItem, isVisible)
}
