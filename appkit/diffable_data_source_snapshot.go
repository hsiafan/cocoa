// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[DiffableDataSourceSnapshot](dc, objc.SEL("alloc"))
	return rv
}

func (dc _DiffableDataSourceSnapshotClass) New() DiffableDataSourceSnapshot {
	rv := objc.CallMethod[DiffableDataSourceSnapshot](dc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewDiffableDataSourceSnapshot() DiffableDataSourceSnapshot {
	return DiffableDataSourceSnapshotClass.New()
}

func (d_ DiffableDataSourceSnapshot) Init() DiffableDataSourceSnapshot {
	rv := objc.CallMethod[DiffableDataSourceSnapshot](d_, objc.SEL("init"))
	return rv
}

func (d_ DiffableDataSourceSnapshot) AppendSectionsWithIdentifiers(sectionIdentifiers []objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.SEL("appendSectionsWithIdentifiers:"), sectionIdentifiers)
}

func (d_ DiffableDataSourceSnapshot) AppendItemsWithIdentifiers_IntoSectionWithIdentifier(identifiers []objc.IObject, sectionIdentifier objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.SEL("appendItemsWithIdentifiers:intoSectionWithIdentifier:"), identifiers, objc.ExtractPtr(sectionIdentifier))
}

func (d_ DiffableDataSourceSnapshot) AppendItemsWithIdentifiers(identifiers []objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.SEL("appendItemsWithIdentifiers:"), identifiers)
}

func (d_ DiffableDataSourceSnapshot) NumberOfItemsInSection(sectionIdentifier objc.IObject) int {
	rv := objc.CallMethod[int](d_, objc.SEL("numberOfItemsInSection:"), objc.ExtractPtr(sectionIdentifier))
	return rv
}

func (d_ DiffableDataSourceSnapshot) IndexOfItemIdentifier(itemIdentifier objc.IObject) int {
	rv := objc.CallMethod[int](d_, objc.SEL("indexOfItemIdentifier:"), objc.ExtractPtr(itemIdentifier))
	return rv
}

func (d_ DiffableDataSourceSnapshot) IndexOfSectionIdentifier(sectionIdentifier objc.IObject) int {
	rv := objc.CallMethod[int](d_, objc.SEL("indexOfSectionIdentifier:"), objc.ExtractPtr(sectionIdentifier))
	return rv
}

func (d_ DiffableDataSourceSnapshot) ItemIdentifiersInSectionWithIdentifier(sectionIdentifier objc.IObject) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](d_, objc.SEL("itemIdentifiersInSectionWithIdentifier:"), objc.ExtractPtr(sectionIdentifier))
	return rv
}

func (d_ DiffableDataSourceSnapshot) SectionIdentifierForSectionContainingItemIdentifier(itemIdentifier objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.SEL("sectionIdentifierForSectionContainingItemIdentifier:"), objc.ExtractPtr(itemIdentifier))
	return rv
}

func (d_ DiffableDataSourceSnapshot) InsertItemsWithIdentifiers_AfterItemWithIdentifier(identifiers []objc.IObject, itemIdentifier objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.SEL("insertItemsWithIdentifiers:afterItemWithIdentifier:"), identifiers, objc.ExtractPtr(itemIdentifier))
}

func (d_ DiffableDataSourceSnapshot) InsertItemsWithIdentifiers_BeforeItemWithIdentifier(identifiers []objc.IObject, itemIdentifier objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.SEL("insertItemsWithIdentifiers:beforeItemWithIdentifier:"), identifiers, objc.ExtractPtr(itemIdentifier))
}

func (d_ DiffableDataSourceSnapshot) InsertSectionsWithIdentifiers_AfterSectionWithIdentifier(sectionIdentifiers []objc.IObject, toSectionIdentifier objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.SEL("insertSectionsWithIdentifiers:afterSectionWithIdentifier:"), sectionIdentifiers, objc.ExtractPtr(toSectionIdentifier))
}

func (d_ DiffableDataSourceSnapshot) InsertSectionsWithIdentifiers_BeforeSectionWithIdentifier(sectionIdentifiers []objc.IObject, toSectionIdentifier objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.SEL("insertSectionsWithIdentifiers:beforeSectionWithIdentifier:"), sectionIdentifiers, objc.ExtractPtr(toSectionIdentifier))
}

func (d_ DiffableDataSourceSnapshot) DeleteAllItems() {
	objc.CallMethod[objc.Void](d_, objc.SEL("deleteAllItems"))
}

func (d_ DiffableDataSourceSnapshot) DeleteItemsWithIdentifiers(identifiers []objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.SEL("deleteItemsWithIdentifiers:"), identifiers)
}

func (d_ DiffableDataSourceSnapshot) DeleteSectionsWithIdentifiers(sectionIdentifiers []objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.SEL("deleteSectionsWithIdentifiers:"), sectionIdentifiers)
}

func (d_ DiffableDataSourceSnapshot) MoveItemWithIdentifier_AfterItemWithIdentifier(fromIdentifier objc.IObject, toIdentifier objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.SEL("moveItemWithIdentifier:afterItemWithIdentifier:"), objc.ExtractPtr(fromIdentifier), objc.ExtractPtr(toIdentifier))
}

func (d_ DiffableDataSourceSnapshot) MoveItemWithIdentifier_BeforeItemWithIdentifier(fromIdentifier objc.IObject, toIdentifier objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.SEL("moveItemWithIdentifier:beforeItemWithIdentifier:"), objc.ExtractPtr(fromIdentifier), objc.ExtractPtr(toIdentifier))
}

func (d_ DiffableDataSourceSnapshot) MoveSectionWithIdentifier_AfterSectionWithIdentifier(fromSectionIdentifier objc.IObject, toSectionIdentifier objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.SEL("moveSectionWithIdentifier:afterSectionWithIdentifier:"), objc.ExtractPtr(fromSectionIdentifier), objc.ExtractPtr(toSectionIdentifier))
}

func (d_ DiffableDataSourceSnapshot) MoveSectionWithIdentifier_BeforeSectionWithIdentifier(fromSectionIdentifier objc.IObject, toSectionIdentifier objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.SEL("moveSectionWithIdentifier:beforeSectionWithIdentifier:"), objc.ExtractPtr(fromSectionIdentifier), objc.ExtractPtr(toSectionIdentifier))
}

func (d_ DiffableDataSourceSnapshot) ReloadItemsWithIdentifiers(identifiers []objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.SEL("reloadItemsWithIdentifiers:"), identifiers)
}

func (d_ DiffableDataSourceSnapshot) ReloadSectionsWithIdentifiers(sectionIdentifiers []objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.SEL("reloadSectionsWithIdentifiers:"), sectionIdentifiers)
}

func (d_ DiffableDataSourceSnapshot) NumberOfItems() int {
	rv := objc.CallMethod[int](d_, objc.SEL("numberOfItems"))
	return rv
}

func (d_ DiffableDataSourceSnapshot) NumberOfSections() int {
	rv := objc.CallMethod[int](d_, objc.SEL("numberOfSections"))
	return rv
}

func (d_ DiffableDataSourceSnapshot) ItemIdentifiers() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](d_, objc.SEL("itemIdentifiers"))
	return rv
}

func (d_ DiffableDataSourceSnapshot) SectionIdentifiers() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](d_, objc.SEL("sectionIdentifiers"))
	return rv
}
