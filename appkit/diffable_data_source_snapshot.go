// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var DiffableDataSourceSnapshotClass = _DiffableDataSourceSnapshotClass{objc.GetClass("NSDiffableDataSourceSnapshot")}

type _DiffableDataSourceSnapshotClass struct {
	objc.Class
}

type IDiffableDataSourceSnapshot interface {
	objc.IObject
	AppendSectionsWithIdentifiers(sectionIdentifiers []objc.IObject)
	AppendItemsWithIdentifiers_IntoSectionWithIdentifier(identifiers []objc.IObject, sectionIdentifier objc.IObject)
	AppendItemsWithIdentifiers(identifiers []objc.IObject)
	NumberOfItemsInSection(sectionIdentifier objc.IObject) int
	IndexOfItemIdentifier(itemIdentifier objc.IObject) int
	IndexOfSectionIdentifier(sectionIdentifier objc.IObject) int
	ItemIdentifiersInSectionWithIdentifier(sectionIdentifier objc.IObject) []objc.Object
	SectionIdentifierForSectionContainingItemIdentifier(itemIdentifier objc.IObject) objc.Object
	InsertItemsWithIdentifiers_AfterItemWithIdentifier(identifiers []objc.IObject, itemIdentifier objc.IObject)
	InsertItemsWithIdentifiers_BeforeItemWithIdentifier(identifiers []objc.IObject, itemIdentifier objc.IObject)
	InsertSectionsWithIdentifiers_AfterSectionWithIdentifier(sectionIdentifiers []objc.IObject, toSectionIdentifier objc.IObject)
	InsertSectionsWithIdentifiers_BeforeSectionWithIdentifier(sectionIdentifiers []objc.IObject, toSectionIdentifier objc.IObject)
	DeleteAllItems()
	DeleteItemsWithIdentifiers(identifiers []objc.IObject)
	DeleteSectionsWithIdentifiers(sectionIdentifiers []objc.IObject)
	MoveItemWithIdentifier_AfterItemWithIdentifier(fromIdentifier objc.IObject, toIdentifier objc.IObject)
	MoveItemWithIdentifier_BeforeItemWithIdentifier(fromIdentifier objc.IObject, toIdentifier objc.IObject)
	MoveSectionWithIdentifier_AfterSectionWithIdentifier(fromSectionIdentifier objc.IObject, toSectionIdentifier objc.IObject)
	MoveSectionWithIdentifier_BeforeSectionWithIdentifier(fromSectionIdentifier objc.IObject, toSectionIdentifier objc.IObject)
	ReloadItemsWithIdentifiers(identifiers []objc.IObject)
	ReloadSectionsWithIdentifiers(sectionIdentifiers []objc.IObject)
	NumberOfItems() int
	NumberOfSections() int
	ItemIdentifiers() []objc.Object
	SectionIdentifiers() []objc.Object
}

type DiffableDataSourceSnapshot struct {
	objc.Object
}

func MakeDiffableDataSourceSnapshot(ptr unsafe.Pointer) DiffableDataSourceSnapshot {
	return DiffableDataSourceSnapshot{
		Object: objc.MakeObject(ptr),
	}
}

func (dc _DiffableDataSourceSnapshotClass) Alloc() DiffableDataSourceSnapshot {
	rv := ffi.CallMethod[DiffableDataSourceSnapshot](dc, "alloc")
	return rv
}

func (dc _DiffableDataSourceSnapshotClass) New() DiffableDataSourceSnapshot {
	rv := ffi.CallMethod[DiffableDataSourceSnapshot](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDiffableDataSourceSnapshot() DiffableDataSourceSnapshot {
	return DiffableDataSourceSnapshotClass.New()
}

func (d_ DiffableDataSourceSnapshot) Init() DiffableDataSourceSnapshot {
	rv := ffi.CallMethod[DiffableDataSourceSnapshot](d_, "init")
	return rv
}

func (d_ DiffableDataSourceSnapshot) AppendSectionsWithIdentifiers(sectionIdentifiers []objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "appendSectionsWithIdentifiers:", sectionIdentifiers)
}

func (d_ DiffableDataSourceSnapshot) AppendItemsWithIdentifiers_IntoSectionWithIdentifier(identifiers []objc.IObject, sectionIdentifier objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "appendItemsWithIdentifiers:intoSectionWithIdentifier:", identifiers, sectionIdentifier)
}

func (d_ DiffableDataSourceSnapshot) AppendItemsWithIdentifiers(identifiers []objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "appendItemsWithIdentifiers:", identifiers)
}

func (d_ DiffableDataSourceSnapshot) NumberOfItemsInSection(sectionIdentifier objc.IObject) int {
	rv := ffi.CallMethod[int](d_, "numberOfItemsInSection:", sectionIdentifier)
	return rv
}

func (d_ DiffableDataSourceSnapshot) IndexOfItemIdentifier(itemIdentifier objc.IObject) int {
	rv := ffi.CallMethod[int](d_, "indexOfItemIdentifier:", itemIdentifier)
	return rv
}

func (d_ DiffableDataSourceSnapshot) IndexOfSectionIdentifier(sectionIdentifier objc.IObject) int {
	rv := ffi.CallMethod[int](d_, "indexOfSectionIdentifier:", sectionIdentifier)
	return rv
}

func (d_ DiffableDataSourceSnapshot) ItemIdentifiersInSectionWithIdentifier(sectionIdentifier objc.IObject) []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](d_, "itemIdentifiersInSectionWithIdentifier:", sectionIdentifier)
	return rv
}

func (d_ DiffableDataSourceSnapshot) SectionIdentifierForSectionContainingItemIdentifier(itemIdentifier objc.IObject) objc.Object {
	rv := ffi.CallMethod[objc.Object](d_, "sectionIdentifierForSectionContainingItemIdentifier:", itemIdentifier)
	return rv
}

func (d_ DiffableDataSourceSnapshot) InsertItemsWithIdentifiers_AfterItemWithIdentifier(identifiers []objc.IObject, itemIdentifier objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "insertItemsWithIdentifiers:afterItemWithIdentifier:", identifiers, itemIdentifier)
}

func (d_ DiffableDataSourceSnapshot) InsertItemsWithIdentifiers_BeforeItemWithIdentifier(identifiers []objc.IObject, itemIdentifier objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "insertItemsWithIdentifiers:beforeItemWithIdentifier:", identifiers, itemIdentifier)
}

func (d_ DiffableDataSourceSnapshot) InsertSectionsWithIdentifiers_AfterSectionWithIdentifier(sectionIdentifiers []objc.IObject, toSectionIdentifier objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "insertSectionsWithIdentifiers:afterSectionWithIdentifier:", sectionIdentifiers, toSectionIdentifier)
}

func (d_ DiffableDataSourceSnapshot) InsertSectionsWithIdentifiers_BeforeSectionWithIdentifier(sectionIdentifiers []objc.IObject, toSectionIdentifier objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "insertSectionsWithIdentifiers:beforeSectionWithIdentifier:", sectionIdentifiers, toSectionIdentifier)
}

func (d_ DiffableDataSourceSnapshot) DeleteAllItems() {
	ffi.CallMethod[ffi.Void](d_, "deleteAllItems")
}

func (d_ DiffableDataSourceSnapshot) DeleteItemsWithIdentifiers(identifiers []objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "deleteItemsWithIdentifiers:", identifiers)
}

func (d_ DiffableDataSourceSnapshot) DeleteSectionsWithIdentifiers(sectionIdentifiers []objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "deleteSectionsWithIdentifiers:", sectionIdentifiers)
}

func (d_ DiffableDataSourceSnapshot) MoveItemWithIdentifier_AfterItemWithIdentifier(fromIdentifier objc.IObject, toIdentifier objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "moveItemWithIdentifier:afterItemWithIdentifier:", fromIdentifier, toIdentifier)
}

func (d_ DiffableDataSourceSnapshot) MoveItemWithIdentifier_BeforeItemWithIdentifier(fromIdentifier objc.IObject, toIdentifier objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "moveItemWithIdentifier:beforeItemWithIdentifier:", fromIdentifier, toIdentifier)
}

func (d_ DiffableDataSourceSnapshot) MoveSectionWithIdentifier_AfterSectionWithIdentifier(fromSectionIdentifier objc.IObject, toSectionIdentifier objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "moveSectionWithIdentifier:afterSectionWithIdentifier:", fromSectionIdentifier, toSectionIdentifier)
}

func (d_ DiffableDataSourceSnapshot) MoveSectionWithIdentifier_BeforeSectionWithIdentifier(fromSectionIdentifier objc.IObject, toSectionIdentifier objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "moveSectionWithIdentifier:beforeSectionWithIdentifier:", fromSectionIdentifier, toSectionIdentifier)
}

func (d_ DiffableDataSourceSnapshot) ReloadItemsWithIdentifiers(identifiers []objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "reloadItemsWithIdentifiers:", identifiers)
}

func (d_ DiffableDataSourceSnapshot) ReloadSectionsWithIdentifiers(sectionIdentifiers []objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "reloadSectionsWithIdentifiers:", sectionIdentifiers)
}

func (d_ DiffableDataSourceSnapshot) NumberOfItems() int {
	rv := ffi.CallMethod[int](d_, "numberOfItems")
	return rv
}

func (d_ DiffableDataSourceSnapshot) NumberOfSections() int {
	rv := ffi.CallMethod[int](d_, "numberOfSections")
	return rv
}

func (d_ DiffableDataSourceSnapshot) ItemIdentifiers() []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](d_, "itemIdentifiers")
	return rv
}

func (d_ DiffableDataSourceSnapshot) SectionIdentifiers() []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](d_, "sectionIdentifiers")
	return rv
}
