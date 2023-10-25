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
	rv := objc.CallMethod[ArrayController](a_, objc.SEL("initWithContent:"), objc.ExtractPtr(content))
	return rv
}

func (a_ ArrayController) Init() ArrayController {
	rv := objc.CallMethod[ArrayController](a_, objc.SEL("init"))
	return rv
}

func (ac _ArrayControllerClass) Alloc() ArrayController {
	rv := objc.CallMethod[ArrayController](ac, objc.SEL("alloc"))
	return rv
}

func (ac _ArrayControllerClass) New() ArrayController {
	rv := objc.CallMethod[ArrayController](ac, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewArrayController() ArrayController {
	return ArrayControllerClass.New()
}

func (a_ ArrayController) ArrangeObjects(objects []objc.IObject) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](a_, objc.SEL("arrangeObjects:"), objects)
	return rv
}

func (a_ ArrayController) RearrangeObjects() {
	objc.CallMethod[objc.Void](a_, objc.SEL("rearrangeObjects"))
}

func (a_ ArrayController) SetSelectionIndex(index uint) bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("setSelectionIndex:"), index)
	return rv
}

func (a_ ArrayController) SetSelectionIndexes(indexes foundation.IIndexSet) bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("setSelectionIndexes:"), objc.ExtractPtr(indexes))
	return rv
}

func (a_ ArrayController) AddSelectionIndexes(indexes foundation.IIndexSet) bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("addSelectionIndexes:"), objc.ExtractPtr(indexes))
	return rv
}

func (a_ ArrayController) RemoveSelectionIndexes(indexes foundation.IIndexSet) bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("removeSelectionIndexes:"), objc.ExtractPtr(indexes))
	return rv
}

func (a_ ArrayController) SetSelectedObjects(objects []objc.IObject) bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("setSelectedObjects:"), objects)
	return rv
}

func (a_ ArrayController) AddSelectedObjects(objects []objc.IObject) bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("addSelectedObjects:"), objects)
	return rv
}

func (a_ ArrayController) RemoveSelectedObjects(objects []objc.IObject) bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("removeSelectedObjects:"), objects)
	return rv
}

func (a_ ArrayController) SelectNext(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.SEL("selectNext:"), objc.ExtractPtr(sender))
}

func (a_ ArrayController) SelectPrevious(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.SEL("selectPrevious:"), objc.ExtractPtr(sender))
}

func (a_ ArrayController) Insert(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.SEL("insert:"), objc.ExtractPtr(sender))
}

func (a_ ArrayController) AddObjects(objects []objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.SEL("addObjects:"), objects)
}

func (a_ ArrayController) InsertObject_AtArrangedObjectIndex(object objc.IObject, index uint) {
	objc.CallMethod[objc.Void](a_, objc.SEL("insertObject:atArrangedObjectIndex:"), objc.ExtractPtr(object), index)
}

func (a_ ArrayController) InsertObjects_AtArrangedObjectIndexes(objects []objc.IObject, indexes foundation.IIndexSet) {
	objc.CallMethod[objc.Void](a_, objc.SEL("insertObjects:atArrangedObjectIndexes:"), objects, objc.ExtractPtr(indexes))
}

func (a_ ArrayController) RemoveObjectAtArrangedObjectIndex(index uint) {
	objc.CallMethod[objc.Void](a_, objc.SEL("removeObjectAtArrangedObjectIndex:"), index)
}

func (a_ ArrayController) RemoveObjectsAtArrangedObjectIndexes(indexes foundation.IIndexSet) {
	objc.CallMethod[objc.Void](a_, objc.SEL("removeObjectsAtArrangedObjectIndexes:"), objc.ExtractPtr(indexes))
}

func (a_ ArrayController) RemoveObjects(objects []objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.SEL("removeObjects:"), objects)
}

func (a_ ArrayController) DidChangeArrangementCriteria() {
	objc.CallMethod[objc.Void](a_, objc.SEL("didChangeArrangementCriteria"))
}

func (a_ ArrayController) SortDescriptors() []foundation.SortDescriptor {
	rv := objc.CallMethod[[]foundation.SortDescriptor](a_, objc.SEL("sortDescriptors"))
	return rv
}

func (a_ ArrayController) SetSortDescriptors(value []foundation.ISortDescriptor) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setSortDescriptors:"), value)
}

func (a_ ArrayController) ArrangedObjects() objc.Object {
	rv := objc.CallMethod[objc.Object](a_, objc.SEL("arrangedObjects"))
	return rv
}

func (a_ ArrayController) AvoidsEmptySelection() bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("avoidsEmptySelection"))
	return rv
}

func (a_ ArrayController) SetAvoidsEmptySelection(value bool) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setAvoidsEmptySelection:"), value)
}

func (a_ ArrayController) PreservesSelection() bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("preservesSelection"))
	return rv
}

func (a_ ArrayController) SetPreservesSelection(value bool) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setPreservesSelection:"), value)
}

func (a_ ArrayController) AlwaysUsesMultipleValuesMarker() bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("alwaysUsesMultipleValuesMarker"))
	return rv
}

func (a_ ArrayController) SetAlwaysUsesMultipleValuesMarker(value bool) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setAlwaysUsesMultipleValuesMarker:"), value)
}

func (a_ ArrayController) SelectionIndex() uint {
	rv := objc.CallMethod[uint](a_, objc.SEL("selectionIndex"))
	return rv
}

func (a_ ArrayController) SelectsInsertedObjects() bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("selectsInsertedObjects"))
	return rv
}

func (a_ ArrayController) SetSelectsInsertedObjects(value bool) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setSelectsInsertedObjects:"), value)
}

func (a_ ArrayController) SelectionIndexes() foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](a_, objc.SEL("selectionIndexes"))
	return rv
}

func (a_ ArrayController) CanSelectNext() bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("canSelectNext"))
	return rv
}

func (a_ ArrayController) CanSelectPrevious() bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("canSelectPrevious"))
	return rv
}

func (a_ ArrayController) CanInsert() bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("canInsert"))
	return rv
}

func (a_ ArrayController) ClearsFilterPredicateOnInsertion() bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("clearsFilterPredicateOnInsertion"))
	return rv
}

func (a_ ArrayController) SetClearsFilterPredicateOnInsertion(value bool) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setClearsFilterPredicateOnInsertion:"), value)
}

func (a_ ArrayController) FilterPredicate() foundation.Predicate {
	rv := objc.CallMethod[foundation.Predicate](a_, objc.SEL("filterPredicate"))
	return rv
}

func (a_ ArrayController) SetFilterPredicate(value foundation.IPredicate) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setFilterPredicate:"), objc.ExtractPtr(value))
}

func (a_ ArrayController) AutomaticallyRearrangesObjects() bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("automaticallyRearrangesObjects"))
	return rv
}

func (a_ ArrayController) SetAutomaticallyRearrangesObjects(value bool) {
	objc.CallMethod[objc.Void](a_, objc.SEL("setAutomaticallyRearrangesObjects:"), value)
}

func (a_ ArrayController) AutomaticRearrangementKeyPaths() []string {
	rv := objc.CallMethod[[]string](a_, objc.SEL("automaticRearrangementKeyPaths"))
	return rv
}
