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

type CandidateListTouchBarItemDelegateCreator struct {
	className string
	class     objc.Class
}

func NewCandidateListTouchBarItemDelegateCreator(name string) *CandidateListTouchBarItemDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &CandidateListTouchBarItemDelegateCreator{className: name, class: class}
}

func (c *CandidateListTouchBarItemDelegateCreator) SetCandidateListTouchBarItem_BeginSelectingCandidateAtIndex(handle func(o objc.Object, anItem CandidateListTouchBarItem, index int)) {
	objc.AddMethod(c.class, objc.SEL("candidateListTouchBarItem:beginSelectingCandidateAtIndex:"), handle)
}

func (c *CandidateListTouchBarItemDelegateCreator) SetCandidateListTouchBarItem_ChangeSelectionFromCandidateAtIndex_ToIndex(handle func(o objc.Object, anItem CandidateListTouchBarItem, previousIndex int, index int)) {
	objc.AddMethod(c.class, objc.SEL("candidateListTouchBarItem:changeSelectionFromCandidateAtIndex:toIndex:"), handle)
}

func (c *CandidateListTouchBarItemDelegateCreator) SetCandidateListTouchBarItem_EndSelectingCandidateAtIndex(handle func(o objc.Object, anItem CandidateListTouchBarItem, index int)) {
	objc.AddMethod(c.class, objc.SEL("candidateListTouchBarItem:endSelectingCandidateAtIndex:"), handle)
}

func (c *CandidateListTouchBarItemDelegateCreator) SetCandidateListTouchBarItem_ChangedCandidateListVisibility(handle func(o objc.Object, anItem CandidateListTouchBarItem, isVisible bool)) {
	objc.AddMethod(c.class, objc.SEL("candidateListTouchBarItem:changedCandidateListVisibility:"), handle)
}

func (c *CandidateListTouchBarItemDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
