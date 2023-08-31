// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ArrayControllerClass = _ArrayControllerClass{objc.GetClass("NSArrayController")}

type _ArrayControllerClass struct {
	objc.Class
}

type IArrayController interface {
	IObjectController
	ArrangeObjects(objects []objc.IObject) []objc.Object
	RearrangeObjects()
	SetSelectionIndex(index uint) bool
	SetSelectionIndexes(indexes foundation.IIndexSet) bool
	AddSelectionIndexes(indexes foundation.IIndexSet) bool
	RemoveSelectionIndexes(indexes foundation.IIndexSet) bool
	SetSelectedObjects(objects []objc.IObject) bool
	AddSelectedObjects(objects []objc.IObject) bool
	RemoveSelectedObjects(objects []objc.IObject) bool
	SelectNext(sender objc.IObject)
	SelectPrevious(sender objc.IObject)
	Insert(sender objc.IObject)
	AddObjects(objects []objc.IObject)
	InsertObject_AtArrangedObjectIndex(object objc.IObject, index uint)
	InsertObjects_AtArrangedObjectIndexes(objects []objc.IObject, indexes foundation.IIndexSet)
	RemoveObjectAtArrangedObjectIndex(index uint)
	RemoveObjectsAtArrangedObjectIndexes(indexes foundation.IIndexSet)
	RemoveObjects(objects []objc.IObject)
	DidChangeArrangementCriteria()
	SortDescriptors() []foundation.SortDescriptor
	SetSortDescriptors(value []foundation.ISortDescriptor)
	ArrangedObjects() objc.Object
	AvoidsEmptySelection() bool
	SetAvoidsEmptySelection(value bool)
	PreservesSelection() bool
	SetPreservesSelection(value bool)
	AlwaysUsesMultipleValuesMarker() bool
	SetAlwaysUsesMultipleValuesMarker(value bool)
	SelectionIndex() uint
	SelectsInsertedObjects() bool
	SetSelectsInsertedObjects(value bool)
	SelectionIndexes() foundation.IndexSet
	CanSelectNext() bool
	CanSelectPrevious() bool
	CanInsert() bool
	ClearsFilterPredicateOnInsertion() bool
	SetClearsFilterPredicateOnInsertion(value bool)
	FilterPredicate() foundation.Predicate
	SetFilterPredicate(value foundation.IPredicate)
	AutomaticallyRearrangesObjects() bool
	SetAutomaticallyRearrangesObjects(value bool)
	AutomaticRearrangementKeyPaths() []string
}

type ArrayController struct {
	ObjectController
}

func MakeArrayController(ptr unsafe.Pointer) ArrayController {
	return ArrayController{
		ObjectController: MakeObjectController(ptr),
	}
}

func (a_ ArrayController) InitWithContent(content objc.IObject) ArrayController {
	rv := objc.CallMethod[ArrayController](a_, objc.GetSelector("initWithContent:"), objc.ExtractPtr(content))
	return rv
}

func (a_ ArrayController) Init() ArrayController {
	rv := objc.CallMethod[ArrayController](a_, objc.GetSelector("init"))
	return rv
}

func (ac _ArrayControllerClass) Alloc() ArrayController {
	rv := objc.CallMethod[ArrayController](ac, objc.GetSelector("alloc"))
	return rv
}

func (ac _ArrayControllerClass) New() ArrayController {
	rv := objc.CallMethod[ArrayController](ac, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewArrayController() ArrayController {
	return ArrayControllerClass.New()
}

func (a_ ArrayController) ArrangeObjects(objects []objc.IObject) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](a_, objc.GetSelector("arrangeObjects:"), objects)
	return rv
}

func (a_ ArrayController) RearrangeObjects() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("rearrangeObjects"))
}

func (a_ ArrayController) SetSelectionIndex(index uint) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("setSelectionIndex:"), index)
	return rv
}

func (a_ ArrayController) SetSelectionIndexes(indexes foundation.IIndexSet) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("setSelectionIndexes:"), objc.ExtractPtr(indexes))
	return rv
}

func (a_ ArrayController) AddSelectionIndexes(indexes foundation.IIndexSet) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("addSelectionIndexes:"), objc.ExtractPtr(indexes))
	return rv
}

func (a_ ArrayController) RemoveSelectionIndexes(indexes foundation.IIndexSet) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("removeSelectionIndexes:"), objc.ExtractPtr(indexes))
	return rv
}

func (a_ ArrayController) SetSelectedObjects(objects []objc.IObject) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("setSelectedObjects:"), objects)
	return rv
}

func (a_ ArrayController) AddSelectedObjects(objects []objc.IObject) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("addSelectedObjects:"), objects)
	return rv
}

func (a_ ArrayController) RemoveSelectedObjects(objects []objc.IObject) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("removeSelectedObjects:"), objects)
	return rv
}

func (a_ ArrayController) SelectNext(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("selectNext:"), objc.ExtractPtr(sender))
}

func (a_ ArrayController) SelectPrevious(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("selectPrevious:"), objc.ExtractPtr(sender))
}

func (a_ ArrayController) Insert(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("insert:"), objc.ExtractPtr(sender))
}

func (a_ ArrayController) AddObjects(objects []objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("addObjects:"), objects)
}

func (a_ ArrayController) InsertObject_AtArrangedObjectIndex(object objc.IObject, index uint) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("insertObject:atArrangedObjectIndex:"), objc.ExtractPtr(object), index)
}

func (a_ ArrayController) InsertObjects_AtArrangedObjectIndexes(objects []objc.IObject, indexes foundation.IIndexSet) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("insertObjects:atArrangedObjectIndexes:"), objects, objc.ExtractPtr(indexes))
}

func (a_ ArrayController) RemoveObjectAtArrangedObjectIndex(index uint) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("removeObjectAtArrangedObjectIndex:"), index)
}

func (a_ ArrayController) RemoveObjectsAtArrangedObjectIndexes(indexes foundation.IIndexSet) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("removeObjectsAtArrangedObjectIndexes:"), objc.ExtractPtr(indexes))
}

func (a_ ArrayController) RemoveObjects(objects []objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("removeObjects:"), objects)
}

func (a_ ArrayController) DidChangeArrangementCriteria() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("didChangeArrangementCriteria"))
}

func (a_ ArrayController) SortDescriptors() []foundation.SortDescriptor {
	rv := objc.CallMethod[[]foundation.SortDescriptor](a_, objc.GetSelector("sortDescriptors"))
	return rv
}

func (a_ ArrayController) SetSortDescriptors(value []foundation.ISortDescriptor) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setSortDescriptors:"), value)
}

func (a_ ArrayController) ArrangedObjects() objc.Object {
	rv := objc.CallMethod[objc.Object](a_, objc.GetSelector("arrangedObjects"))
	return rv
}

func (a_ ArrayController) AvoidsEmptySelection() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("avoidsEmptySelection"))
	return rv
}

func (a_ ArrayController) SetAvoidsEmptySelection(value bool) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setAvoidsEmptySelection:"), value)
}

func (a_ ArrayController) PreservesSelection() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("preservesSelection"))
	return rv
}

func (a_ ArrayController) SetPreservesSelection(value bool) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setPreservesSelection:"), value)
}

func (a_ ArrayController) AlwaysUsesMultipleValuesMarker() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("alwaysUsesMultipleValuesMarker"))
	return rv
}

func (a_ ArrayController) SetAlwaysUsesMultipleValuesMarker(value bool) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setAlwaysUsesMultipleValuesMarker:"), value)
}

func (a_ ArrayController) SelectionIndex() uint {
	rv := objc.CallMethod[uint](a_, objc.GetSelector("selectionIndex"))
	return rv
}

func (a_ ArrayController) SelectsInsertedObjects() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("selectsInsertedObjects"))
	return rv
}

func (a_ ArrayController) SetSelectsInsertedObjects(value bool) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setSelectsInsertedObjects:"), value)
}

func (a_ ArrayController) SelectionIndexes() foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](a_, objc.GetSelector("selectionIndexes"))
	return rv
}

func (a_ ArrayController) CanSelectNext() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("canSelectNext"))
	return rv
}

func (a_ ArrayController) CanSelectPrevious() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("canSelectPrevious"))
	return rv
}

func (a_ ArrayController) CanInsert() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("canInsert"))
	return rv
}

func (a_ ArrayController) ClearsFilterPredicateOnInsertion() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("clearsFilterPredicateOnInsertion"))
	return rv
}

func (a_ ArrayController) SetClearsFilterPredicateOnInsertion(value bool) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setClearsFilterPredicateOnInsertion:"), value)
}

func (a_ ArrayController) FilterPredicate() foundation.Predicate {
	rv := objc.CallMethod[foundation.Predicate](a_, objc.GetSelector("filterPredicate"))
	return rv
}

func (a_ ArrayController) SetFilterPredicate(value foundation.IPredicate) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setFilterPredicate:"), objc.ExtractPtr(value))
}

func (a_ ArrayController) AutomaticallyRearrangesObjects() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("automaticallyRearrangesObjects"))
	return rv
}

func (a_ ArrayController) SetAutomaticallyRearrangesObjects(value bool) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setAutomaticallyRearrangesObjects:"), value)
}

func (a_ ArrayController) AutomaticRearrangementKeyPaths() []string {
	rv := objc.CallMethod[[]string](a_, objc.GetSelector("automaticRearrangementKeyPaths"))
	return rv
}
