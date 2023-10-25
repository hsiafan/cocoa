// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionViewClass = _CollectionViewClass{objc.GetClass("NSCollectionView")}

type _CollectionViewClass struct {
	objc.Class
}

type ICollectionView interface {
	IView
	MakeItemWithIdentifier_ForIndexPath(identifier UserInterfaceItemIdentifier, indexPath foundation.IIndexPath) CollectionViewItem
	RegisterClass_ForItemWithIdentifier(itemClass objc.IClass, identifier UserInterfaceItemIdentifier)
	RegisterNib_ForItemWithIdentifier(nib INib, identifier UserInterfaceItemIdentifier)
	MakeSupplementaryViewOfKind_WithIdentifier_ForIndexPath(elementKind CollectionViewSupplementaryElementKind, identifier UserInterfaceItemIdentifier, indexPath foundation.IIndexPath) View
	RegisterClass_ForSupplementaryViewOfKind_WithIdentifier(viewClass objc.IClass, kind CollectionViewSupplementaryElementKind, identifier UserInterfaceItemIdentifier)
	RegisterNib_ForSupplementaryViewOfKind_WithIdentifier(nib INib, kind CollectionViewSupplementaryElementKind, identifier UserInterfaceItemIdentifier)
	ReloadData()
	ReloadSections(sections foundation.IIndexSet)
	ReloadItemsAtIndexPaths(indexPaths foundation.ISet)
	NumberOfItemsInSection(section int) int
	InsertItemsAtIndexPaths(indexPaths foundation.ISet)
	MoveItemAtIndexPath_ToIndexPath(indexPath foundation.IIndexPath, newIndexPath foundation.IIndexPath)
	DeleteItemsAtIndexPaths(indexPaths foundation.ISet)
	InsertSections(sections foundation.IIndexSet)
	MoveSection_ToSection(section int, newSection int)
	DeleteSections(sections foundation.IIndexSet)
	ToggleSectionCollapse(sender objc.IObject)
	SelectAll(sender objc.IObject)
	DeselectAll(sender objc.IObject)
	SelectItemsAtIndexPaths_ScrollPosition(indexPaths foundation.ISet, scrollPosition CollectionViewScrollPosition)
	DeselectItemsAtIndexPaths(indexPaths foundation.ISet)
	VisibleItems() []CollectionViewItem
	IndexPathsForVisibleItems() foundation.Set
	VisibleSupplementaryViewsOfKind(elementKind CollectionViewSupplementaryElementKind) []View
	IndexPathsForVisibleSupplementaryElementsOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set
	IndexPathForItem(item ICollectionViewItem) foundation.IndexPath
	IndexPathForItemAtPoint(point foundation.Point) foundation.IndexPath
	ItemAtIndexPath(indexPath foundation.IIndexPath) CollectionViewItem
	SupplementaryViewForElementKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) View
	ScrollToItemsAtIndexPaths_ScrollPosition(indexPaths foundation.ISet, scrollPosition CollectionViewScrollPosition)
	LayoutAttributesForItemAtIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	LayoutAttributesForSupplementaryElementOfKind_AtIndexPath(kind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	PerformBatchUpdates_CompletionHandler(updates func(), completionHandler func(finished bool))
	DraggingImageForItemsAtIndexPaths_WithEvent_Offset(indexPaths foundation.ISet, event IEvent, dragImageOffset *foundation.Point) Image
	// deprecated
	NewItemForRepresentedObject(object objc.IObject) CollectionViewItem
	ItemAtIndex(index uint) CollectionViewItem
	FrameForItemAtIndex(index uint) foundation.Rect
	FrameForItemAtIndex_WithNumberOfItems(index uint, numberOfItems uint) foundation.Rect
	DraggingImageForItemsAtIndexes_WithEvent_Offset(indexes foundation.IIndexSet, event IEvent, dragImageOffset *foundation.Point) Image
	SetDraggingSourceOperationMask_ForLocal(dragOperationMask DragOperation, localDestination bool)
	DataSource() objc.Object
	SetDataSource(value objc.IObject)
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
	Content() []objc.Object
	SetContent(value []objc.IObject)
	BackgroundView() View
	SetBackgroundView(value IView)
	BackgroundColors() []Color
	SetBackgroundColors(value []IColor)
	BackgroundViewScrollsWithContent() bool
	SetBackgroundViewScrollsWithContent(value bool)
	CollectionViewLayout() CollectionViewLayout
	SetCollectionViewLayout(value ICollectionViewLayout)
	PrefetchDataSource() objc.Object
	SetPrefetchDataSource(value objc.IObject)
	NumberOfSections() int
	IsSelectable() bool
	SetSelectable(value bool)
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	SelectionIndexPaths() foundation.Set
	SetSelectionIndexPaths(value foundation.ISet)
	IsFirstResponder() bool
	// deprecated
	ItemPrototype() CollectionViewItem
	// deprecated
	SetItemPrototype(value ICollectionViewItem)
	SelectionIndexes() foundation.IndexSet
	SetSelectionIndexes(value foundation.IIndexSet)
	// deprecated
	MaxNumberOfRows() uint
	// deprecated
	SetMaxNumberOfRows(value uint)
	// deprecated
	MaxNumberOfColumns() uint
	// deprecated
	SetMaxNumberOfColumns(value uint)
	// deprecated
	MinItemSize() foundation.Size
	// deprecated
	SetMinItemSize(value foundation.Size)
	// deprecated
	MaxItemSize() foundation.Size
	// deprecated
	SetMaxItemSize(value foundation.Size)
}

type CollectionView struct {
	View
}

func MakeCollectionView(ptr unsafe.Pointer) CollectionView {
	return CollectionView{
		View: MakeView(ptr),
	}
}

func (c_ CollectionView) InitWithFrame(frameRect foundation.Rect) CollectionView {
	rv := objc.CallMethod[CollectionView](c_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func (c_ CollectionView) Init() CollectionView {
	rv := objc.CallMethod[CollectionView](c_, objc.GetSelector("init"))
	return rv
}

func (cc _CollectionViewClass) Alloc() CollectionView {
	rv := objc.CallMethod[CollectionView](cc, objc.GetSelector("alloc"))
	return rv
}

func (cc _CollectionViewClass) New() CollectionView {
	rv := objc.CallMethod[CollectionView](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionView() CollectionView {
	return CollectionViewClass.New()
}

func (c_ CollectionView) MakeItemWithIdentifier_ForIndexPath(identifier UserInterfaceItemIdentifier, indexPath foundation.IIndexPath) CollectionViewItem {
	rv := objc.CallMethod[CollectionViewItem](c_, objc.GetSelector("makeItemWithIdentifier:forIndexPath:"), identifier, objc.ExtractPtr(indexPath))
	return rv
}

func (c_ CollectionView) RegisterClass_ForItemWithIdentifier(itemClass objc.IClass, identifier UserInterfaceItemIdentifier) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("registerClass:forItemWithIdentifier:"), objc.ExtractPtr(itemClass), identifier)
}

func (c_ CollectionView) RegisterNib_ForItemWithIdentifier(nib INib, identifier UserInterfaceItemIdentifier) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("registerNib:forItemWithIdentifier:"), objc.ExtractPtr(nib), identifier)
}

func (c_ CollectionView) MakeSupplementaryViewOfKind_WithIdentifier_ForIndexPath(elementKind CollectionViewSupplementaryElementKind, identifier UserInterfaceItemIdentifier, indexPath foundation.IIndexPath) View {
	rv := objc.CallMethod[View](c_, objc.GetSelector("makeSupplementaryViewOfKind:withIdentifier:forIndexPath:"), elementKind, identifier, objc.ExtractPtr(indexPath))
	return rv
}

func (c_ CollectionView) RegisterClass_ForSupplementaryViewOfKind_WithIdentifier(viewClass objc.IClass, kind CollectionViewSupplementaryElementKind, identifier UserInterfaceItemIdentifier) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("registerClass:forSupplementaryViewOfKind:withIdentifier:"), objc.ExtractPtr(viewClass), kind, identifier)
}

func (c_ CollectionView) RegisterNib_ForSupplementaryViewOfKind_WithIdentifier(nib INib, kind CollectionViewSupplementaryElementKind, identifier UserInterfaceItemIdentifier) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("registerNib:forSupplementaryViewOfKind:withIdentifier:"), objc.ExtractPtr(nib), kind, identifier)
}

func (c_ CollectionView) ReloadData() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("reloadData"))
}

func (c_ CollectionView) ReloadSections(sections foundation.IIndexSet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("reloadSections:"), objc.ExtractPtr(sections))
}

func (c_ CollectionView) ReloadItemsAtIndexPaths(indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("reloadItemsAtIndexPaths:"), objc.ExtractPtr(indexPaths))
}

func (c_ CollectionView) NumberOfItemsInSection(section int) int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("numberOfItemsInSection:"), section)
	return rv
}

func (c_ CollectionView) InsertItemsAtIndexPaths(indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("insertItemsAtIndexPaths:"), objc.ExtractPtr(indexPaths))
}

func (c_ CollectionView) MoveItemAtIndexPath_ToIndexPath(indexPath foundation.IIndexPath, newIndexPath foundation.IIndexPath) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("moveItemAtIndexPath:toIndexPath:"), objc.ExtractPtr(indexPath), objc.ExtractPtr(newIndexPath))
}

func (c_ CollectionView) DeleteItemsAtIndexPaths(indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("deleteItemsAtIndexPaths:"), objc.ExtractPtr(indexPaths))
}

func (c_ CollectionView) InsertSections(sections foundation.IIndexSet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("insertSections:"), objc.ExtractPtr(sections))
}

func (c_ CollectionView) MoveSection_ToSection(section int, newSection int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("moveSection:toSection:"), section, newSection)
}

func (c_ CollectionView) DeleteSections(sections foundation.IIndexSet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("deleteSections:"), objc.ExtractPtr(sections))
}

func (c_ CollectionView) ToggleSectionCollapse(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("toggleSectionCollapse:"), objc.ExtractPtr(sender))
}

func (c_ CollectionView) SelectAll(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("selectAll:"), objc.ExtractPtr(sender))
}

func (c_ CollectionView) DeselectAll(sender objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("deselectAll:"), objc.ExtractPtr(sender))
}

func (c_ CollectionView) SelectItemsAtIndexPaths_ScrollPosition(indexPaths foundation.ISet, scrollPosition CollectionViewScrollPosition) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("selectItemsAtIndexPaths:scrollPosition:"), objc.ExtractPtr(indexPaths), scrollPosition)
}

func (c_ CollectionView) DeselectItemsAtIndexPaths(indexPaths foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("deselectItemsAtIndexPaths:"), objc.ExtractPtr(indexPaths))
}

func (c_ CollectionView) VisibleItems() []CollectionViewItem {
	rv := objc.CallMethod[[]CollectionViewItem](c_, objc.GetSelector("visibleItems"))
	return rv
}

func (c_ CollectionView) IndexPathsForVisibleItems() foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, objc.GetSelector("indexPathsForVisibleItems"))
	return rv
}

func (c_ CollectionView) VisibleSupplementaryViewsOfKind(elementKind CollectionViewSupplementaryElementKind) []View {
	rv := objc.CallMethod[[]View](c_, objc.GetSelector("visibleSupplementaryViewsOfKind:"), elementKind)
	return rv
}

func (c_ CollectionView) IndexPathsForVisibleSupplementaryElementsOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, objc.GetSelector("indexPathsForVisibleSupplementaryElementsOfKind:"), elementKind)
	return rv
}

func (c_ CollectionView) IndexPathForItem(item ICollectionViewItem) foundation.IndexPath {
	rv := objc.CallMethod[foundation.IndexPath](c_, objc.GetSelector("indexPathForItem:"), objc.ExtractPtr(item))
	return rv
}

func (c_ CollectionView) IndexPathForItemAtPoint(point foundation.Point) foundation.IndexPath {
	rv := objc.CallMethod[foundation.IndexPath](c_, objc.GetSelector("indexPathForItemAtPoint:"), point)
	return rv
}

func (c_ CollectionView) ItemAtIndexPath(indexPath foundation.IIndexPath) CollectionViewItem {
	rv := objc.CallMethod[CollectionViewItem](c_, objc.GetSelector("itemAtIndexPath:"), objc.ExtractPtr(indexPath))
	return rv
}

func (c_ CollectionView) SupplementaryViewForElementKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) View {
	rv := objc.CallMethod[View](c_, objc.GetSelector("supplementaryViewForElementKind:atIndexPath:"), elementKind, objc.ExtractPtr(indexPath))
	return rv
}

func (c_ CollectionView) ScrollToItemsAtIndexPaths_ScrollPosition(indexPaths foundation.ISet, scrollPosition CollectionViewScrollPosition) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("scrollToItemsAtIndexPaths:scrollPosition:"), objc.ExtractPtr(indexPaths), scrollPosition)
}

func (c_ CollectionView) LayoutAttributesForItemAtIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, objc.GetSelector("layoutAttributesForItemAtIndexPath:"), objc.ExtractPtr(indexPath))
	return rv
}

func (c_ CollectionView) LayoutAttributesForSupplementaryElementOfKind_AtIndexPath(kind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, objc.GetSelector("layoutAttributesForSupplementaryElementOfKind:atIndexPath:"), kind, objc.ExtractPtr(indexPath))
	return rv
}

func (c_ CollectionView) PerformBatchUpdates_CompletionHandler(updates func(), completionHandler func(finished bool)) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("performBatchUpdates:completionHandler:"), updates, completionHandler)
}

func (c_ CollectionView) DraggingImageForItemsAtIndexPaths_WithEvent_Offset(indexPaths foundation.ISet, event IEvent, dragImageOffset *foundation.Point) Image {
	rv := objc.CallMethod[Image](c_, objc.GetSelector("draggingImageForItemsAtIndexPaths:withEvent:offset:"), objc.ExtractPtr(indexPaths), objc.ExtractPtr(event), dragImageOffset)
	return rv
}

// deprecated
func (c_ CollectionView) NewItemForRepresentedObject(object objc.IObject) CollectionViewItem {
	rv := objc.CallMethod[CollectionViewItem](c_, objc.GetSelector("newItemForRepresentedObject:"), objc.ExtractPtr(object))
	rv.Autorelease()
	return rv
}

func (c_ CollectionView) ItemAtIndex(index uint) CollectionViewItem {
	rv := objc.CallMethod[CollectionViewItem](c_, objc.GetSelector("itemAtIndex:"), index)
	return rv
}

func (c_ CollectionView) FrameForItemAtIndex(index uint) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.GetSelector("frameForItemAtIndex:"), index)
	return rv
}

func (c_ CollectionView) FrameForItemAtIndex_WithNumberOfItems(index uint, numberOfItems uint) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](c_, objc.GetSelector("frameForItemAtIndex:withNumberOfItems:"), index, numberOfItems)
	return rv
}

func (c_ CollectionView) DraggingImageForItemsAtIndexes_WithEvent_Offset(indexes foundation.IIndexSet, event IEvent, dragImageOffset *foundation.Point) Image {
	rv := objc.CallMethod[Image](c_, objc.GetSelector("draggingImageForItemsAtIndexes:withEvent:offset:"), objc.ExtractPtr(indexes), objc.ExtractPtr(event), dragImageOffset)
	return rv
}

func (c_ CollectionView) SetDraggingSourceOperationMask_ForLocal(dragOperationMask DragOperation, localDestination bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setDraggingSourceOperationMask:forLocal:"), dragOperationMask, localDestination)
}

// weak property
func (c_ CollectionView) DataSource() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.GetSelector("dataSource"))
	return rv
}

// weak property
func (c_ CollectionView) SetDataSource(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setDataSource:"), objc.ExtractPtr(value))
}

// weak property
func (c_ CollectionView) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.GetSelector("delegate"))
	return rv
}

// weak property
func (c_ CollectionView) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (c_ CollectionView) Content() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](c_, objc.GetSelector("content"))
	return rv
}

func (c_ CollectionView) SetContent(value []objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setContent:"), value)
}

func (c_ CollectionView) BackgroundView() View {
	rv := objc.CallMethod[View](c_, objc.GetSelector("backgroundView"))
	return rv
}

func (c_ CollectionView) SetBackgroundView(value IView) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setBackgroundView:"), objc.ExtractPtr(value))
}

func (c_ CollectionView) BackgroundColors() []Color {
	rv := objc.CallMethod[[]Color](c_, objc.GetSelector("backgroundColors"))
	return rv
}

func (c_ CollectionView) SetBackgroundColors(value []IColor) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setBackgroundColors:"), value)
}

func (c_ CollectionView) BackgroundViewScrollsWithContent() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("backgroundViewScrollsWithContent"))
	return rv
}

func (c_ CollectionView) SetBackgroundViewScrollsWithContent(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setBackgroundViewScrollsWithContent:"), value)
}

func (c_ CollectionView) CollectionViewLayout() CollectionViewLayout {
	rv := objc.CallMethod[CollectionViewLayout](c_, objc.GetSelector("collectionViewLayout"))
	return rv
}

func (c_ CollectionView) SetCollectionViewLayout(value ICollectionViewLayout) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setCollectionViewLayout:"), objc.ExtractPtr(value))
}

// weak property
func (c_ CollectionView) PrefetchDataSource() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.GetSelector("prefetchDataSource"))
	return rv
}

// weak property
func (c_ CollectionView) SetPrefetchDataSource(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setPrefetchDataSource:"), objc.ExtractPtr(value))
}

func (c_ CollectionView) NumberOfSections() int {
	rv := objc.CallMethod[int](c_, objc.GetSelector("numberOfSections"))
	return rv
}

func (c_ CollectionView) IsSelectable() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isSelectable"))
	return rv
}

func (c_ CollectionView) SetSelectable(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setSelectable:"), value)
}

func (c_ CollectionView) AllowsMultipleSelection() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("allowsMultipleSelection"))
	return rv
}

func (c_ CollectionView) SetAllowsMultipleSelection(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setAllowsMultipleSelection:"), value)
}

func (c_ CollectionView) AllowsEmptySelection() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("allowsEmptySelection"))
	return rv
}

func (c_ CollectionView) SetAllowsEmptySelection(value bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setAllowsEmptySelection:"), value)
}

func (c_ CollectionView) SelectionIndexPaths() foundation.Set {
	rv := objc.CallMethod[foundation.Set](c_, objc.GetSelector("selectionIndexPaths"))
	return rv
}

func (c_ CollectionView) SetSelectionIndexPaths(value foundation.ISet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setSelectionIndexPaths:"), objc.ExtractPtr(value))
}

func (c_ CollectionView) IsFirstResponder() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isFirstResponder"))
	return rv
}

// deprecated
func (c_ CollectionView) ItemPrototype() CollectionViewItem {
	rv := objc.CallMethod[CollectionViewItem](c_, objc.GetSelector("itemPrototype"))
	return rv
}

// deprecated
func (c_ CollectionView) SetItemPrototype(value ICollectionViewItem) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setItemPrototype:"), objc.ExtractPtr(value))
}

func (c_ CollectionView) SelectionIndexes() foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](c_, objc.GetSelector("selectionIndexes"))
	return rv
}

func (c_ CollectionView) SetSelectionIndexes(value foundation.IIndexSet) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setSelectionIndexes:"), objc.ExtractPtr(value))
}

// deprecated
func (c_ CollectionView) MaxNumberOfRows() uint {
	rv := objc.CallMethod[uint](c_, objc.GetSelector("maxNumberOfRows"))
	return rv
}

// deprecated
func (c_ CollectionView) SetMaxNumberOfRows(value uint) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setMaxNumberOfRows:"), value)
}

// deprecated
func (c_ CollectionView) MaxNumberOfColumns() uint {
	rv := objc.CallMethod[uint](c_, objc.GetSelector("maxNumberOfColumns"))
	return rv
}

// deprecated
func (c_ CollectionView) SetMaxNumberOfColumns(value uint) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setMaxNumberOfColumns:"), value)
}

// deprecated
func (c_ CollectionView) MinItemSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("minItemSize"))
	return rv
}

// deprecated
func (c_ CollectionView) SetMinItemSize(value foundation.Size) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setMinItemSize:"), value)
}

// deprecated
func (c_ CollectionView) MaxItemSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("maxItemSize"))
	return rv
}

// deprecated
func (c_ CollectionView) SetMaxItemSize(value foundation.Size) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setMaxItemSize:"), value)
}
