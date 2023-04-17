// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var OutlineViewClass = _OutlineViewClass{objc.GetClass("NSOutlineView")}

type _OutlineViewClass struct {
	objc.Class
}

type IOutlineView interface {
	ITableView
	IsExpandable(item objc.IObject) bool
	IsItemExpanded(item objc.IObject) bool
	ExpandItem(item objc.IObject)
	ExpandItem_ExpandChildren(item objc.IObject, expandChildren bool)
	CollapseItem(item objc.IObject)
	CollapseItem_CollapseChildren(item objc.IObject, collapseChildren bool)
	ReloadItem(item objc.IObject)
	ReloadItem_ReloadChildren(item objc.IObject, reloadChildren bool)
	ItemAtRow(row int) objc.Object
	RowForItem(item objc.IObject) int
	LevelForItem(item objc.IObject) int
	LevelForRow(row int) int
	SetDropItem_DropChildIndex(item objc.IObject, index int)
	ShouldCollapseAutoExpandedItemsForDeposited(deposited bool) bool
	ParentForItem(item objc.IObject) objc.Object
	ChildIndexForItem(item objc.IObject) int
	Child_OfItem(index int, item objc.IObject) objc.Object
	NumberOfChildrenOfItem(item objc.IObject) int
	FrameOfOutlineCellAtRow(row int) foundation.Rect
	InsertItemsAtIndexes_InParent_WithAnimation(indexes foundation.IIndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions)
	MoveItemAtIndex_InParent_ToIndex_InParent(fromIndex int, oldParent objc.IObject, toIndex int, newParent objc.IObject)
	RemoveItemsAtIndexes_InParent_WithAnimation(indexes foundation.IIndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions)
	StronglyReferencesItems() bool
	SetStronglyReferencesItems(value bool)
	OutlineTableColumn() TableColumn
	SetOutlineTableColumn(value ITableColumn)
	AutoresizesOutlineColumn() bool
	SetAutoresizesOutlineColumn(value bool)
	IndentationPerLevel() float64
	SetIndentationPerLevel(value float64)
	IndentationMarkerFollowsCell() bool
	SetIndentationMarkerFollowsCell(value bool)
	AutosaveExpandedItems() bool
	SetAutosaveExpandedItems(value bool)
}

type OutlineView struct {
	TableView
}

func MakeOutlineView(ptr unsafe.Pointer) OutlineView {
	return OutlineView{
		TableView: MakeTableView(ptr),
	}
}

func (o_ OutlineView) InitWithFrame(frameRect foundation.Rect) OutlineView {
	rv := ffi.CallMethod[OutlineView](o_, "initWithFrame:", frameRect)
	return rv
}

func (o_ OutlineView) Init() OutlineView {
	rv := ffi.CallMethod[OutlineView](o_, "init")
	return rv
}

func (oc _OutlineViewClass) Alloc() OutlineView {
	rv := ffi.CallMethod[OutlineView](oc, "alloc")
	return rv
}

func (oc _OutlineViewClass) New() OutlineView {
	rv := ffi.CallMethod[OutlineView](oc, "new")
	rv.Autorelease()
	return rv
}

func NewOutlineView() OutlineView {
	return OutlineViewClass.New()
}

func (o_ OutlineView) IsExpandable(item objc.IObject) bool {
	rv := ffi.CallMethod[bool](o_, "isExpandable:", item)
	return rv
}

func (o_ OutlineView) IsItemExpanded(item objc.IObject) bool {
	rv := ffi.CallMethod[bool](o_, "isItemExpanded:", item)
	return rv
}

func (o_ OutlineView) ExpandItem(item objc.IObject) {
	ffi.CallMethod[ffi.Void](o_, "expandItem:", item)
}

func (o_ OutlineView) ExpandItem_ExpandChildren(item objc.IObject, expandChildren bool) {
	ffi.CallMethod[ffi.Void](o_, "expandItem:expandChildren:", item, expandChildren)
}

func (o_ OutlineView) CollapseItem(item objc.IObject) {
	ffi.CallMethod[ffi.Void](o_, "collapseItem:", item)
}

func (o_ OutlineView) CollapseItem_CollapseChildren(item objc.IObject, collapseChildren bool) {
	ffi.CallMethod[ffi.Void](o_, "collapseItem:collapseChildren:", item, collapseChildren)
}

func (o_ OutlineView) ReloadItem(item objc.IObject) {
	ffi.CallMethod[ffi.Void](o_, "reloadItem:", item)
}

func (o_ OutlineView) ReloadItem_ReloadChildren(item objc.IObject, reloadChildren bool) {
	ffi.CallMethod[ffi.Void](o_, "reloadItem:reloadChildren:", item, reloadChildren)
}

func (o_ OutlineView) ItemAtRow(row int) objc.Object {
	rv := ffi.CallMethod[objc.Object](o_, "itemAtRow:", row)
	return rv
}

func (o_ OutlineView) RowForItem(item objc.IObject) int {
	rv := ffi.CallMethod[int](o_, "rowForItem:", item)
	return rv
}

func (o_ OutlineView) LevelForItem(item objc.IObject) int {
	rv := ffi.CallMethod[int](o_, "levelForItem:", item)
	return rv
}

func (o_ OutlineView) LevelForRow(row int) int {
	rv := ffi.CallMethod[int](o_, "levelForRow:", row)
	return rv
}

func (o_ OutlineView) SetDropItem_DropChildIndex(item objc.IObject, index int) {
	ffi.CallMethod[ffi.Void](o_, "setDropItem:dropChildIndex:", item, index)
}

func (o_ OutlineView) ShouldCollapseAutoExpandedItemsForDeposited(deposited bool) bool {
	rv := ffi.CallMethod[bool](o_, "shouldCollapseAutoExpandedItemsForDeposited:", deposited)
	return rv
}

func (o_ OutlineView) ParentForItem(item objc.IObject) objc.Object {
	rv := ffi.CallMethod[objc.Object](o_, "parentForItem:", item)
	return rv
}

func (o_ OutlineView) ChildIndexForItem(item objc.IObject) int {
	rv := ffi.CallMethod[int](o_, "childIndexForItem:", item)
	return rv
}

func (o_ OutlineView) Child_OfItem(index int, item objc.IObject) objc.Object {
	rv := ffi.CallMethod[objc.Object](o_, "child:ofItem:", index, item)
	return rv
}

func (o_ OutlineView) NumberOfChildrenOfItem(item objc.IObject) int {
	rv := ffi.CallMethod[int](o_, "numberOfChildrenOfItem:", item)
	return rv
}

func (o_ OutlineView) FrameOfOutlineCellAtRow(row int) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](o_, "frameOfOutlineCellAtRow:", row)
	return rv
}

func (o_ OutlineView) InsertItemsAtIndexes_InParent_WithAnimation(indexes foundation.IIndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions) {
	ffi.CallMethod[ffi.Void](o_, "insertItemsAtIndexes:inParent:withAnimation:", indexes, parent, animationOptions)
}

func (o_ OutlineView) MoveItemAtIndex_InParent_ToIndex_InParent(fromIndex int, oldParent objc.IObject, toIndex int, newParent objc.IObject) {
	ffi.CallMethod[ffi.Void](o_, "moveItemAtIndex:inParent:toIndex:inParent:", fromIndex, oldParent, toIndex, newParent)
}

func (o_ OutlineView) RemoveItemsAtIndexes_InParent_WithAnimation(indexes foundation.IIndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions) {
	ffi.CallMethod[ffi.Void](o_, "removeItemsAtIndexes:inParent:withAnimation:", indexes, parent, animationOptions)
}

func (o_ OutlineView) StronglyReferencesItems() bool {
	rv := ffi.CallMethod[bool](o_, "stronglyReferencesItems")
	return rv
}

func (o_ OutlineView) SetStronglyReferencesItems(value bool) {
	ffi.CallMethod[ffi.Void](o_, "setStronglyReferencesItems:", value)
}

func (o_ OutlineView) OutlineTableColumn() TableColumn {
	rv := ffi.CallMethod[TableColumn](o_, "outlineTableColumn")
	return rv
}

func (o_ OutlineView) SetOutlineTableColumn(value ITableColumn) {
	ffi.CallMethod[ffi.Void](o_, "setOutlineTableColumn:", value)
}

func (o_ OutlineView) AutoresizesOutlineColumn() bool {
	rv := ffi.CallMethod[bool](o_, "autoresizesOutlineColumn")
	return rv
}

func (o_ OutlineView) SetAutoresizesOutlineColumn(value bool) {
	ffi.CallMethod[ffi.Void](o_, "setAutoresizesOutlineColumn:", value)
}

func (o_ OutlineView) IndentationPerLevel() float64 {
	rv := ffi.CallMethod[float64](o_, "indentationPerLevel")
	return rv
}

func (o_ OutlineView) SetIndentationPerLevel(value float64) {
	ffi.CallMethod[ffi.Void](o_, "setIndentationPerLevel:", value)
}

func (o_ OutlineView) IndentationMarkerFollowsCell() bool {
	rv := ffi.CallMethod[bool](o_, "indentationMarkerFollowsCell")
	return rv
}

func (o_ OutlineView) SetIndentationMarkerFollowsCell(value bool) {
	ffi.CallMethod[ffi.Void](o_, "setIndentationMarkerFollowsCell:", value)
}

func (o_ OutlineView) AutosaveExpandedItems() bool {
	rv := ffi.CallMethod[bool](o_, "autosaveExpandedItems")
	return rv
}

func (o_ OutlineView) SetAutosaveExpandedItems(value bool) {
	ffi.CallMethod[ffi.Void](o_, "setAutosaveExpandedItems:", value)
}
