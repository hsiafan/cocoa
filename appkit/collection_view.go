// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionViewClass = _CollectionViewClass{objc.GetClass("NSCollectionView")}

type _CollectionViewClass struct {
	objc.Class
}

type ICollectionView interface {
	IView
	MakeItemWithIdentifier_ForIndexPath(identifier UserInterfaceItemIdentifier, indexPath foundation.IIndexPath) CollectionViewItem
	MakeSupplementaryViewOfKind_WithIdentifier_ForIndexPath(elementKind CollectionViewSupplementaryElementKind, identifier UserInterfaceItemIdentifier, indexPath foundation.IIndexPath) View
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
	DataSource() CollectionViewDataSourceWrapper
	SetDataSource(value CollectionViewDataSource)
	Delegate() CollectionViewDelegateWrapper
	SetDelegate(value CollectionViewDelegate)
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
	PrefetchDataSource() CollectionViewPrefetchingWrapper
	SetPrefetchDataSource(value CollectionViewPrefetching)
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
	rv := ffi.CallMethod[CollectionView](c_, "initWithFrame:", frameRect)
	return rv
}

func (c_ CollectionView) Init() CollectionView {
	rv := ffi.CallMethod[CollectionView](c_, "init")
	return rv
}

func (cc _CollectionViewClass) Alloc() CollectionView {
	rv := ffi.CallMethod[CollectionView](cc, "alloc")
	return rv
}

func (cc _CollectionViewClass) New() CollectionView {
	rv := ffi.CallMethod[CollectionView](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionView() CollectionView {
	return CollectionViewClass.New()
}

func (c_ CollectionView) MakeItemWithIdentifier_ForIndexPath(identifier UserInterfaceItemIdentifier, indexPath foundation.IIndexPath) CollectionViewItem {
	rv := ffi.CallMethod[CollectionViewItem](c_, "makeItemWithIdentifier:forIndexPath:", identifier, indexPath)
	return rv
}

func (c_ CollectionView) MakeSupplementaryViewOfKind_WithIdentifier_ForIndexPath(elementKind CollectionViewSupplementaryElementKind, identifier UserInterfaceItemIdentifier, indexPath foundation.IIndexPath) View {
	rv := ffi.CallMethod[View](c_, "makeSupplementaryViewOfKind:withIdentifier:forIndexPath:", elementKind, identifier, indexPath)
	return rv
}

func (c_ CollectionView) ReloadData() {
	ffi.CallMethod[ffi.Void](c_, "reloadData")
}

func (c_ CollectionView) ReloadSections(sections foundation.IIndexSet) {
	ffi.CallMethod[ffi.Void](c_, "reloadSections:", sections)
}

func (c_ CollectionView) ReloadItemsAtIndexPaths(indexPaths foundation.ISet) {
	ffi.CallMethod[ffi.Void](c_, "reloadItemsAtIndexPaths:", indexPaths)
}

func (c_ CollectionView) NumberOfItemsInSection(section int) int {
	rv := ffi.CallMethod[int](c_, "numberOfItemsInSection:", section)
	return rv
}

func (c_ CollectionView) InsertItemsAtIndexPaths(indexPaths foundation.ISet) {
	ffi.CallMethod[ffi.Void](c_, "insertItemsAtIndexPaths:", indexPaths)
}

func (c_ CollectionView) MoveItemAtIndexPath_ToIndexPath(indexPath foundation.IIndexPath, newIndexPath foundation.IIndexPath) {
	ffi.CallMethod[ffi.Void](c_, "moveItemAtIndexPath:toIndexPath:", indexPath, newIndexPath)
}

func (c_ CollectionView) DeleteItemsAtIndexPaths(indexPaths foundation.ISet) {
	ffi.CallMethod[ffi.Void](c_, "deleteItemsAtIndexPaths:", indexPaths)
}

func (c_ CollectionView) InsertSections(sections foundation.IIndexSet) {
	ffi.CallMethod[ffi.Void](c_, "insertSections:", sections)
}

func (c_ CollectionView) MoveSection_ToSection(section int, newSection int) {
	ffi.CallMethod[ffi.Void](c_, "moveSection:toSection:", section, newSection)
}

func (c_ CollectionView) DeleteSections(sections foundation.IIndexSet) {
	ffi.CallMethod[ffi.Void](c_, "deleteSections:", sections)
}

func (c_ CollectionView) ToggleSectionCollapse(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "toggleSectionCollapse:", sender)
}

func (c_ CollectionView) SelectAll(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "selectAll:", sender)
}

func (c_ CollectionView) DeselectAll(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "deselectAll:", sender)
}

func (c_ CollectionView) SelectItemsAtIndexPaths_ScrollPosition(indexPaths foundation.ISet, scrollPosition CollectionViewScrollPosition) {
	ffi.CallMethod[ffi.Void](c_, "selectItemsAtIndexPaths:scrollPosition:", indexPaths, scrollPosition)
}

func (c_ CollectionView) DeselectItemsAtIndexPaths(indexPaths foundation.ISet) {
	ffi.CallMethod[ffi.Void](c_, "deselectItemsAtIndexPaths:", indexPaths)
}

func (c_ CollectionView) VisibleItems() []CollectionViewItem {
	rv := ffi.CallMethod[[]CollectionViewItem](c_, "visibleItems")
	return rv
}

func (c_ CollectionView) IndexPathsForVisibleItems() foundation.Set {
	rv := ffi.CallMethod[foundation.Set](c_, "indexPathsForVisibleItems")
	return rv
}

func (c_ CollectionView) VisibleSupplementaryViewsOfKind(elementKind CollectionViewSupplementaryElementKind) []View {
	rv := ffi.CallMethod[[]View](c_, "visibleSupplementaryViewsOfKind:", elementKind)
	return rv
}

func (c_ CollectionView) IndexPathsForVisibleSupplementaryElementsOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set {
	rv := ffi.CallMethod[foundation.Set](c_, "indexPathsForVisibleSupplementaryElementsOfKind:", elementKind)
	return rv
}

func (c_ CollectionView) IndexPathForItem(item ICollectionViewItem) foundation.IndexPath {
	rv := ffi.CallMethod[foundation.IndexPath](c_, "indexPathForItem:", item)
	return rv
}

func (c_ CollectionView) IndexPathForItemAtPoint(point foundation.Point) foundation.IndexPath {
	rv := ffi.CallMethod[foundation.IndexPath](c_, "indexPathForItemAtPoint:", point)
	return rv
}

func (c_ CollectionView) ItemAtIndexPath(indexPath foundation.IIndexPath) CollectionViewItem {
	rv := ffi.CallMethod[CollectionViewItem](c_, "itemAtIndexPath:", indexPath)
	return rv
}

func (c_ CollectionView) SupplementaryViewForElementKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) View {
	rv := ffi.CallMethod[View](c_, "supplementaryViewForElementKind:atIndexPath:", elementKind, indexPath)
	return rv
}

func (c_ CollectionView) ScrollToItemsAtIndexPaths_ScrollPosition(indexPaths foundation.ISet, scrollPosition CollectionViewScrollPosition) {
	ffi.CallMethod[ffi.Void](c_, "scrollToItemsAtIndexPaths:scrollPosition:", indexPaths, scrollPosition)
}

func (c_ CollectionView) LayoutAttributesForItemAtIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](c_, "layoutAttributesForItemAtIndexPath:", indexPath)
	return rv
}

func (c_ CollectionView) LayoutAttributesForSupplementaryElementOfKind_AtIndexPath(kind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](c_, "layoutAttributesForSupplementaryElementOfKind:atIndexPath:", kind, indexPath)
	return rv
}

func (c_ CollectionView) PerformBatchUpdates_CompletionHandler(updates func(), completionHandler func(finished bool)) {
	ffi.CallMethod[ffi.Void](c_, "performBatchUpdates:completionHandler:", updates, completionHandler)
}

func (c_ CollectionView) DraggingImageForItemsAtIndexPaths_WithEvent_Offset(indexPaths foundation.ISet, event IEvent, dragImageOffset *foundation.Point) Image {
	rv := ffi.CallMethod[Image](c_, "draggingImageForItemsAtIndexPaths:withEvent:offset:", indexPaths, event, dragImageOffset)
	return rv
}

// deprecated
func (c_ CollectionView) NewItemForRepresentedObject(object objc.IObject) CollectionViewItem {
	rv := ffi.CallMethod[CollectionViewItem](c_, "newItemForRepresentedObject:", object)
	rv.Autorelease()
	return rv
}

func (c_ CollectionView) ItemAtIndex(index uint) CollectionViewItem {
	rv := ffi.CallMethod[CollectionViewItem](c_, "itemAtIndex:", index)
	return rv
}

func (c_ CollectionView) FrameForItemAtIndex(index uint) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](c_, "frameForItemAtIndex:", index)
	return rv
}

func (c_ CollectionView) FrameForItemAtIndex_WithNumberOfItems(index uint, numberOfItems uint) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](c_, "frameForItemAtIndex:withNumberOfItems:", index, numberOfItems)
	return rv
}

func (c_ CollectionView) DraggingImageForItemsAtIndexes_WithEvent_Offset(indexes foundation.IIndexSet, event IEvent, dragImageOffset *foundation.Point) Image {
	rv := ffi.CallMethod[Image](c_, "draggingImageForItemsAtIndexes:withEvent:offset:", indexes, event, dragImageOffset)
	return rv
}

func (c_ CollectionView) SetDraggingSourceOperationMask_ForLocal(dragOperationMask DragOperation, localDestination bool) {
	ffi.CallMethod[ffi.Void](c_, "setDraggingSourceOperationMask:forLocal:", dragOperationMask, localDestination)
}

func (c_ CollectionView) DataSource() CollectionViewDataSourceWrapper {
	rv := ffi.CallMethod[CollectionViewDataSourceWrapper](c_, "dataSource")
	return rv
}

func (c_ CollectionView) SetDataSource(value CollectionViewDataSource) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(c_, internal.AssociationKey("setDataSource"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](c_, "setDataSource:", po)
}

func (c_ CollectionView) Delegate() CollectionViewDelegateWrapper {
	rv := ffi.CallMethod[CollectionViewDelegateWrapper](c_, "delegate")
	return rv
}

func (c_ CollectionView) SetDelegate(value CollectionViewDelegate) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(c_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](c_, "setDelegate:", po)
}

func (c_ CollectionView) Content() []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](c_, "content")
	return rv
}

func (c_ CollectionView) SetContent(value []objc.IObject) {
	ffi.CallMethod[ffi.Void](c_, "setContent:", value)
}

func (c_ CollectionView) BackgroundView() View {
	rv := ffi.CallMethod[View](c_, "backgroundView")
	return rv
}

func (c_ CollectionView) SetBackgroundView(value IView) {
	ffi.CallMethod[ffi.Void](c_, "setBackgroundView:", value)
}

func (c_ CollectionView) BackgroundColors() []Color {
	rv := ffi.CallMethod[[]Color](c_, "backgroundColors")
	return rv
}

func (c_ CollectionView) SetBackgroundColors(value []IColor) {
	ffi.CallMethod[ffi.Void](c_, "setBackgroundColors:", value)
}

func (c_ CollectionView) BackgroundViewScrollsWithContent() bool {
	rv := ffi.CallMethod[bool](c_, "backgroundViewScrollsWithContent")
	return rv
}

func (c_ CollectionView) SetBackgroundViewScrollsWithContent(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setBackgroundViewScrollsWithContent:", value)
}

func (c_ CollectionView) CollectionViewLayout() CollectionViewLayout {
	rv := ffi.CallMethod[CollectionViewLayout](c_, "collectionViewLayout")
	return rv
}

func (c_ CollectionView) SetCollectionViewLayout(value ICollectionViewLayout) {
	ffi.CallMethod[ffi.Void](c_, "setCollectionViewLayout:", value)
}

func (c_ CollectionView) PrefetchDataSource() CollectionViewPrefetchingWrapper {
	rv := ffi.CallMethod[CollectionViewPrefetchingWrapper](c_, "prefetchDataSource")
	return rv
}

func (c_ CollectionView) SetPrefetchDataSource(value CollectionViewPrefetching) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(c_, internal.AssociationKey("setPrefetchDataSource"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](c_, "setPrefetchDataSource:", po)
}

func (c_ CollectionView) NumberOfSections() int {
	rv := ffi.CallMethod[int](c_, "numberOfSections")
	return rv
}

func (c_ CollectionView) IsSelectable() bool {
	rv := ffi.CallMethod[bool](c_, "isSelectable")
	return rv
}

func (c_ CollectionView) SetSelectable(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setSelectable:", value)
}

func (c_ CollectionView) AllowsMultipleSelection() bool {
	rv := ffi.CallMethod[bool](c_, "allowsMultipleSelection")
	return rv
}

func (c_ CollectionView) SetAllowsMultipleSelection(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setAllowsMultipleSelection:", value)
}

func (c_ CollectionView) AllowsEmptySelection() bool {
	rv := ffi.CallMethod[bool](c_, "allowsEmptySelection")
	return rv
}

func (c_ CollectionView) SetAllowsEmptySelection(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setAllowsEmptySelection:", value)
}

func (c_ CollectionView) SelectionIndexPaths() foundation.Set {
	rv := ffi.CallMethod[foundation.Set](c_, "selectionIndexPaths")
	return rv
}

func (c_ CollectionView) SetSelectionIndexPaths(value foundation.ISet) {
	ffi.CallMethod[ffi.Void](c_, "setSelectionIndexPaths:", value)
}

func (c_ CollectionView) IsFirstResponder() bool {
	rv := ffi.CallMethod[bool](c_, "isFirstResponder")
	return rv
}

// deprecated
func (c_ CollectionView) ItemPrototype() CollectionViewItem {
	rv := ffi.CallMethod[CollectionViewItem](c_, "itemPrototype")
	return rv
}

// deprecated
func (c_ CollectionView) SetItemPrototype(value ICollectionViewItem) {
	ffi.CallMethod[ffi.Void](c_, "setItemPrototype:", value)
}

func (c_ CollectionView) SelectionIndexes() foundation.IndexSet {
	rv := ffi.CallMethod[foundation.IndexSet](c_, "selectionIndexes")
	return rv
}

func (c_ CollectionView) SetSelectionIndexes(value foundation.IIndexSet) {
	ffi.CallMethod[ffi.Void](c_, "setSelectionIndexes:", value)
}

// deprecated
func (c_ CollectionView) MaxNumberOfRows() uint {
	rv := ffi.CallMethod[uint](c_, "maxNumberOfRows")
	return rv
}

// deprecated
func (c_ CollectionView) SetMaxNumberOfRows(value uint) {
	ffi.CallMethod[ffi.Void](c_, "setMaxNumberOfRows:", value)
}

// deprecated
func (c_ CollectionView) MaxNumberOfColumns() uint {
	rv := ffi.CallMethod[uint](c_, "maxNumberOfColumns")
	return rv
}

// deprecated
func (c_ CollectionView) SetMaxNumberOfColumns(value uint) {
	ffi.CallMethod[ffi.Void](c_, "setMaxNumberOfColumns:", value)
}

// deprecated
func (c_ CollectionView) MinItemSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "minItemSize")
	return rv
}

// deprecated
func (c_ CollectionView) SetMinItemSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](c_, "setMinItemSize:", value)
}

// deprecated
func (c_ CollectionView) MaxItemSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "maxItemSize")
	return rv
}

// deprecated
func (c_ CollectionView) SetMaxItemSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](c_, "setMaxItemSize:", value)
}

var CollectionViewItemClass = _CollectionViewItemClass{objc.GetClass("NSCollectionViewItem")}

type _CollectionViewItemClass struct {
	objc.Class
}

type ICollectionViewItem interface {
	IViewController
	ImageView() ImageView
	SetImageView(value IImageView)
	TextField() TextField
	SetTextField(value ITextField)
	IsSelected() bool
	SetSelected(value bool)
	HighlightState() CollectionViewItemHighlightState
	SetHighlightState(value CollectionViewItemHighlightState)
	CollectionView() CollectionView
	DraggingImageComponents() []DraggingImageComponent
}

type CollectionViewItem struct {
	ViewController
}

func MakeCollectionViewItem(ptr unsafe.Pointer) CollectionViewItem {
	return CollectionViewItem{
		ViewController: MakeViewController(ptr),
	}
}

func (c_ CollectionViewItem) InitWithNibName_Bundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) CollectionViewItem {
	rv := ffi.CallMethod[CollectionViewItem](c_, "initWithNibName:bundle:", nibNameOrNil, nibBundleOrNil)
	return rv
}

func (c_ CollectionViewItem) Init() CollectionViewItem {
	rv := ffi.CallMethod[CollectionViewItem](c_, "init")
	return rv
}

func (cc _CollectionViewItemClass) Alloc() CollectionViewItem {
	rv := ffi.CallMethod[CollectionViewItem](cc, "alloc")
	return rv
}

func (cc _CollectionViewItemClass) New() CollectionViewItem {
	rv := ffi.CallMethod[CollectionViewItem](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewItem() CollectionViewItem {
	return CollectionViewItemClass.New()
}

func (c_ CollectionViewItem) ImageView() ImageView {
	rv := ffi.CallMethod[ImageView](c_, "imageView")
	return rv
}

func (c_ CollectionViewItem) SetImageView(value IImageView) {
	ffi.CallMethod[ffi.Void](c_, "setImageView:", value)
}

func (c_ CollectionViewItem) TextField() TextField {
	rv := ffi.CallMethod[TextField](c_, "textField")
	return rv
}

func (c_ CollectionViewItem) SetTextField(value ITextField) {
	ffi.CallMethod[ffi.Void](c_, "setTextField:", value)
}

func (c_ CollectionViewItem) IsSelected() bool {
	rv := ffi.CallMethod[bool](c_, "isSelected")
	return rv
}

func (c_ CollectionViewItem) SetSelected(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setSelected:", value)
}

func (c_ CollectionViewItem) HighlightState() CollectionViewItemHighlightState {
	rv := ffi.CallMethod[CollectionViewItemHighlightState](c_, "highlightState")
	return rv
}

func (c_ CollectionViewItem) SetHighlightState(value CollectionViewItemHighlightState) {
	ffi.CallMethod[ffi.Void](c_, "setHighlightState:", value)
}

func (c_ CollectionViewItem) CollectionView() CollectionView {
	rv := ffi.CallMethod[CollectionView](c_, "collectionView")
	return rv
}

func (c_ CollectionViewItem) DraggingImageComponents() []DraggingImageComponent {
	rv := ffi.CallMethod[[]DraggingImageComponent](c_, "draggingImageComponents")
	return rv
}

type CollectionViewDataSource interface {
	ImplementsNumberOfSectionsInCollectionView() bool
	// optional
	NumberOfSectionsInCollectionView(collectionView CollectionView) int
	// required
	CollectionView_NumberOfItemsInSection(collectionView CollectionView, section int) int
	// required
	CollectionView_ItemForRepresentedObjectAtIndexPath(collectionView CollectionView, indexPath foundation.IndexPath) ICollectionViewItem
	ImplementsCollectionView_ViewForSupplementaryElementOfKind_AtIndexPath() bool
	// optional
	CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath(collectionView CollectionView, kind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) IView
}

type CollectionViewDataSourceWrapper struct {
	objc.Object
}

func (c_ *CollectionViewDataSourceWrapper) ImplementsNumberOfSectionsInCollectionView() bool {
	return c_.RespondsToSelector(objc.GetSelector("numberOfSectionsInCollectionView:"))
}

func (c_ CollectionViewDataSourceWrapper) NumberOfSectionsInCollectionView(collectionView ICollectionView) int {
	rv := ffi.CallMethod[int](c_, "numberOfSectionsInCollectionView:", collectionView)
	return rv
}

func (c_ CollectionViewDataSourceWrapper) CollectionView_NumberOfItemsInSection(collectionView ICollectionView, section int) int {
	rv := ffi.CallMethod[int](c_, "collectionView:numberOfItemsInSection:", collectionView, section)
	return rv
}

func (c_ CollectionViewDataSourceWrapper) CollectionView_ItemForRepresentedObjectAtIndexPath(collectionView ICollectionView, indexPath foundation.IIndexPath) CollectionViewItem {
	rv := ffi.CallMethod[CollectionViewItem](c_, "collectionView:itemForRepresentedObjectAtIndexPath:", collectionView, indexPath)
	return rv
}

func (c_ *CollectionViewDataSourceWrapper) ImplementsCollectionView_ViewForSupplementaryElementOfKind_AtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:viewForSupplementaryElementOfKind:atIndexPath:"))
}

func (c_ CollectionViewDataSourceWrapper) CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath(collectionView ICollectionView, kind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) View {
	rv := ffi.CallMethod[View](c_, "collectionView:viewForSupplementaryElementOfKind:atIndexPath:", collectionView, kind, indexPath)
	return rv
}

type CollectionViewDelegate interface {
	ImplementsCollectionView_ShouldSelectItemsAtIndexPaths() bool
	// optional
	CollectionView_ShouldSelectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet
	ImplementsCollectionView_DidSelectItemsAtIndexPaths() bool
	// optional
	CollectionView_DidSelectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set)
	ImplementsCollectionView_ShouldDeselectItemsAtIndexPaths() bool
	// optional
	CollectionView_ShouldDeselectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet
	ImplementsCollectionView_DidDeselectItemsAtIndexPaths() bool
	// optional
	CollectionView_DidDeselectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set)
	ImplementsCollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState() bool
	// optional
	CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) foundation.ISet
	ImplementsCollectionView_DidChangeItemsAtIndexPaths_ToHighlightState() bool
	// optional
	CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState)
	ImplementsCollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath() bool
	// optional
	CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)
	ImplementsCollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath() bool
	// optional
	CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)
	ImplementsCollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath() bool
	// optional
	CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)
	ImplementsCollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath() bool
	// optional
	CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)
	ImplementsCollectionView_CanDragItemsAtIndexPaths_WithEvent() bool
	// optional
	CollectionView_CanDragItemsAtIndexPaths_WithEvent(collectionView CollectionView, indexPaths foundation.Set, event Event) bool
	ImplementsCollectionView_PasteboardWriterForItemAtIndexPath() bool
	// optional
	CollectionView_PasteboardWriterForItemAtIndexPath(collectionView CollectionView, indexPath foundation.IndexPath) PasteboardWriting
	ImplementsCollectionView_WriteItemsAtIndexPaths_ToPasteboard() bool
	// optional
	// deprecated
	CollectionView_WriteItemsAtIndexPaths_ToPasteboard(collectionView CollectionView, indexPaths foundation.Set, pasteboard Pasteboard) bool
	ImplementsCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths() bool
	// optional
	// deprecated
	CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths(collectionView CollectionView, dropURL foundation.URL, indexPaths foundation.Set) []string
	ImplementsCollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset() bool
	// optional
	CollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset(collectionView CollectionView, indexPaths foundation.Set, event Event, dragImageOffset *foundation.Point) IImage
	ImplementsCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths() bool
	// optional
	CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexPaths foundation.Set)
	ImplementsCollectionView_DraggingSession_EndedAtPoint_DragOperation() bool
	// optional
	CollectionView_DraggingSession_EndedAtPoint_DragOperation(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	ImplementsCollectionView_UpdateDraggingItemsForDrag() bool
	// optional
	CollectionView_UpdateDraggingItemsForDrag(collectionView CollectionView, draggingInfo DraggingInfoWrapper)
	ImplementsCollectionView_ValidateDrop_ProposedIndexPath_DropOperation() bool
	// optional
	CollectionView_ValidateDrop_ProposedIndexPath_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation
	ImplementsCollectionView_AcceptDrop_IndexPath_DropOperation() bool
	// optional
	CollectionView_AcceptDrop_IndexPath_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool
	ImplementsCollectionView_CanDragItemsAtIndexes_WithEvent() bool
	// optional
	CollectionView_CanDragItemsAtIndexes_WithEvent(collectionView CollectionView, indexes foundation.IndexSet, event Event) bool
	ImplementsCollectionView_PasteboardWriterForItemAtIndex() bool
	// optional
	CollectionView_PasteboardWriterForItemAtIndex(collectionView CollectionView, index uint) PasteboardWriting
	ImplementsCollectionView_WriteItemsAtIndexes_ToPasteboard() bool
	// optional
	// deprecated
	CollectionView_WriteItemsAtIndexes_ToPasteboard(collectionView CollectionView, indexes foundation.IndexSet, pasteboard Pasteboard) bool
	ImplementsCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes() bool
	// optional
	// deprecated
	CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes(collectionView CollectionView, dropURL foundation.URL, indexes foundation.IndexSet) []string
	ImplementsCollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset() bool
	// optional
	CollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset(collectionView CollectionView, indexes foundation.IndexSet, event Event, dragImageOffset *foundation.Point) IImage
	ImplementsCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes() bool
	// optional
	CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexes foundation.IndexSet)
	ImplementsCollectionView_ValidateDrop_ProposedIndex_DropOperation() bool
	// optional
	CollectionView_ValidateDrop_ProposedIndex_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation
	ImplementsCollectionView_AcceptDrop_Index_DropOperation() bool
	// optional
	CollectionView_AcceptDrop_Index_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, index int, dropOperation CollectionViewDropOperation) bool
}

type CollectionViewDelegateImpl struct {
	_CollectionView_ShouldSelectItemsAtIndexPaths                                        func(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet
	_CollectionView_DidSelectItemsAtIndexPaths                                           func(collectionView CollectionView, indexPaths foundation.Set)
	_CollectionView_ShouldDeselectItemsAtIndexPaths                                      func(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet
	_CollectionView_DidDeselectItemsAtIndexPaths                                         func(collectionView CollectionView, indexPaths foundation.Set)
	_CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState                       func(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) foundation.ISet
	_CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState                          func(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState)
	_CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath                      func(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)
	_CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath                 func(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)
	_CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath              func(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)
	_CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath       func(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)
	_CollectionView_CanDragItemsAtIndexPaths_WithEvent                                   func(collectionView CollectionView, indexPaths foundation.Set, event Event) bool
	_CollectionView_PasteboardWriterForItemAtIndexPath                                   func(collectionView CollectionView, indexPath foundation.IndexPath) PasteboardWriting
	_CollectionView_WriteItemsAtIndexPaths_ToPasteboard                                  func(collectionView CollectionView, indexPaths foundation.Set, pasteboard Pasteboard) bool
	_CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths func(collectionView CollectionView, dropURL foundation.URL, indexPaths foundation.Set) []string
	_CollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset                   func(collectionView CollectionView, indexPaths foundation.Set, event Event, dragImageOffset *foundation.Point) IImage
	_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths                func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexPaths foundation.Set)
	_CollectionView_DraggingSession_EndedAtPoint_DragOperation                           func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	_CollectionView_UpdateDraggingItemsForDrag                                           func(collectionView CollectionView, draggingInfo DraggingInfoWrapper)
	_CollectionView_ValidateDrop_ProposedIndexPath_DropOperation                         func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation
	_CollectionView_AcceptDrop_IndexPath_DropOperation                                   func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool
	_CollectionView_CanDragItemsAtIndexes_WithEvent                                      func(collectionView CollectionView, indexes foundation.IndexSet, event Event) bool
	_CollectionView_PasteboardWriterForItemAtIndex                                       func(collectionView CollectionView, index uint) PasteboardWriting
	_CollectionView_WriteItemsAtIndexes_ToPasteboard                                     func(collectionView CollectionView, indexes foundation.IndexSet, pasteboard Pasteboard) bool
	_CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes    func(collectionView CollectionView, dropURL foundation.URL, indexes foundation.IndexSet) []string
	_CollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset                      func(collectionView CollectionView, indexes foundation.IndexSet, event Event, dragImageOffset *foundation.Point) IImage
	_CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes                   func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexes foundation.IndexSet)
	_CollectionView_ValidateDrop_ProposedIndex_DropOperation                             func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation
	_CollectionView_AcceptDrop_Index_DropOperation                                       func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, index int, dropOperation CollectionViewDropOperation) bool
}

func (di *CollectionViewDelegateImpl) ImplementsCollectionView_ShouldSelectItemsAtIndexPaths() bool {
	return di._CollectionView_ShouldSelectItemsAtIndexPaths != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_ShouldSelectItemsAtIndexPaths(f func(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet) {
	di._CollectionView_ShouldSelectItemsAtIndexPaths = f
}

func (di *CollectionViewDelegateImpl) CollectionView_ShouldSelectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet {
	return di._CollectionView_ShouldSelectItemsAtIndexPaths(collectionView, indexPaths)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DidSelectItemsAtIndexPaths() bool {
	return di._CollectionView_DidSelectItemsAtIndexPaths != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DidSelectItemsAtIndexPaths(f func(collectionView CollectionView, indexPaths foundation.Set)) {
	di._CollectionView_DidSelectItemsAtIndexPaths = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DidSelectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) {
	di._CollectionView_DidSelectItemsAtIndexPaths(collectionView, indexPaths)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_ShouldDeselectItemsAtIndexPaths() bool {
	return di._CollectionView_ShouldDeselectItemsAtIndexPaths != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_ShouldDeselectItemsAtIndexPaths(f func(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet) {
	di._CollectionView_ShouldDeselectItemsAtIndexPaths = f
}

func (di *CollectionViewDelegateImpl) CollectionView_ShouldDeselectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) foundation.ISet {
	return di._CollectionView_ShouldDeselectItemsAtIndexPaths(collectionView, indexPaths)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DidDeselectItemsAtIndexPaths() bool {
	return di._CollectionView_DidDeselectItemsAtIndexPaths != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DidDeselectItemsAtIndexPaths(f func(collectionView CollectionView, indexPaths foundation.Set)) {
	di._CollectionView_DidDeselectItemsAtIndexPaths = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DidDeselectItemsAtIndexPaths(collectionView CollectionView, indexPaths foundation.Set) {
	di._CollectionView_DidDeselectItemsAtIndexPaths(collectionView, indexPaths)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState() bool {
	return di._CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(f func(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) foundation.ISet) {
	di._CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState = f
}

func (di *CollectionViewDelegateImpl) CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) foundation.ISet {
	return di._CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(collectionView, indexPaths, highlightState)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DidChangeItemsAtIndexPaths_ToHighlightState() bool {
	return di._CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(f func(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState)) {
	di._CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(collectionView CollectionView, indexPaths foundation.Set, highlightState CollectionViewItemHighlightState) {
	di._CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(collectionView, indexPaths, highlightState)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath() bool {
	return di._CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(f func(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)) {
	di._CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath = f
}

func (di *CollectionViewDelegateImpl) CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath) {
	di._CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(collectionView, item, indexPath)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath() bool {
	return di._CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(f func(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath)) {
	di._CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(collectionView CollectionView, item CollectionViewItem, indexPath foundation.IndexPath) {
	di._CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(collectionView, item, indexPath)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath() bool {
	return di._CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(f func(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)) {
	di._CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath = f
}

func (di *CollectionViewDelegateImpl) CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) {
	di._CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(collectionView, view, elementKind, indexPath)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath() bool {
	return di._CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(f func(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath)) {
	di._CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(collectionView CollectionView, view View, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IndexPath) {
	di._CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(collectionView, view, elementKind, indexPath)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_CanDragItemsAtIndexPaths_WithEvent() bool {
	return di._CollectionView_CanDragItemsAtIndexPaths_WithEvent != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_CanDragItemsAtIndexPaths_WithEvent(f func(collectionView CollectionView, indexPaths foundation.Set, event Event) bool) {
	di._CollectionView_CanDragItemsAtIndexPaths_WithEvent = f
}

func (di *CollectionViewDelegateImpl) CollectionView_CanDragItemsAtIndexPaths_WithEvent(collectionView CollectionView, indexPaths foundation.Set, event Event) bool {
	return di._CollectionView_CanDragItemsAtIndexPaths_WithEvent(collectionView, indexPaths, event)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_PasteboardWriterForItemAtIndexPath() bool {
	return di._CollectionView_PasteboardWriterForItemAtIndexPath != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_PasteboardWriterForItemAtIndexPath(f func(collectionView CollectionView, indexPath foundation.IndexPath) PasteboardWriting) {
	di._CollectionView_PasteboardWriterForItemAtIndexPath = f
}

func (di *CollectionViewDelegateImpl) CollectionView_PasteboardWriterForItemAtIndexPath(collectionView CollectionView, indexPath foundation.IndexPath) PasteboardWriting {
	return di._CollectionView_PasteboardWriterForItemAtIndexPath(collectionView, indexPath)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_WriteItemsAtIndexPaths_ToPasteboard() bool {
	return di._CollectionView_WriteItemsAtIndexPaths_ToPasteboard != nil
}

// deprecated
func (di *CollectionViewDelegateImpl) SetCollectionView_WriteItemsAtIndexPaths_ToPasteboard(f func(collectionView CollectionView, indexPaths foundation.Set, pasteboard Pasteboard) bool) {
	di._CollectionView_WriteItemsAtIndexPaths_ToPasteboard = f
}

// deprecated
func (di *CollectionViewDelegateImpl) CollectionView_WriteItemsAtIndexPaths_ToPasteboard(collectionView CollectionView, indexPaths foundation.Set, pasteboard Pasteboard) bool {
	return di._CollectionView_WriteItemsAtIndexPaths_ToPasteboard(collectionView, indexPaths, pasteboard)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths() bool {
	return di._CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths != nil
}

// deprecated
func (di *CollectionViewDelegateImpl) SetCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths(f func(collectionView CollectionView, dropURL foundation.URL, indexPaths foundation.Set) []string) {
	di._CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths = f
}

// deprecated
func (di *CollectionViewDelegateImpl) CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths(collectionView CollectionView, dropURL foundation.URL, indexPaths foundation.Set) []string {
	return di._CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths(collectionView, dropURL, indexPaths)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset() bool {
	return di._CollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset(f func(collectionView CollectionView, indexPaths foundation.Set, event Event, dragImageOffset *foundation.Point) IImage) {
	di._CollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset(collectionView CollectionView, indexPaths foundation.Set, event Event, dragImageOffset *foundation.Point) IImage {
	return di._CollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset(collectionView, indexPaths, event, dragImageOffset)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths() bool {
	return di._CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(f func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexPaths foundation.Set)) {
	di._CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexPaths foundation.Set) {
	di._CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(collectionView, session, screenPoint, indexPaths)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DraggingSession_EndedAtPoint_DragOperation() bool {
	return di._CollectionView_DraggingSession_EndedAtPoint_DragOperation != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DraggingSession_EndedAtPoint_DragOperation(f func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)) {
	di._CollectionView_DraggingSession_EndedAtPoint_DragOperation = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DraggingSession_EndedAtPoint_DragOperation(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, operation DragOperation) {
	di._CollectionView_DraggingSession_EndedAtPoint_DragOperation(collectionView, session, screenPoint, operation)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_UpdateDraggingItemsForDrag() bool {
	return di._CollectionView_UpdateDraggingItemsForDrag != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_UpdateDraggingItemsForDrag(f func(collectionView CollectionView, draggingInfo DraggingInfoWrapper)) {
	di._CollectionView_UpdateDraggingItemsForDrag = f
}

func (di *CollectionViewDelegateImpl) CollectionView_UpdateDraggingItemsForDrag(collectionView CollectionView, draggingInfo DraggingInfoWrapper) {
	di._CollectionView_UpdateDraggingItemsForDrag(collectionView, draggingInfo)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_ValidateDrop_ProposedIndexPath_DropOperation() bool {
	return di._CollectionView_ValidateDrop_ProposedIndexPath_DropOperation != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_ValidateDrop_ProposedIndexPath_DropOperation(f func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation) {
	di._CollectionView_ValidateDrop_ProposedIndexPath_DropOperation = f
}

func (di *CollectionViewDelegateImpl) CollectionView_ValidateDrop_ProposedIndexPath_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	return di._CollectionView_ValidateDrop_ProposedIndexPath_DropOperation(collectionView, draggingInfo, proposedDropIndexPath, proposedDropOperation)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_AcceptDrop_IndexPath_DropOperation() bool {
	return di._CollectionView_AcceptDrop_IndexPath_DropOperation != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_AcceptDrop_IndexPath_DropOperation(f func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool) {
	di._CollectionView_AcceptDrop_IndexPath_DropOperation = f
}

func (di *CollectionViewDelegateImpl) CollectionView_AcceptDrop_IndexPath_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, indexPath foundation.IndexPath, dropOperation CollectionViewDropOperation) bool {
	return di._CollectionView_AcceptDrop_IndexPath_DropOperation(collectionView, draggingInfo, indexPath, dropOperation)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_CanDragItemsAtIndexes_WithEvent() bool {
	return di._CollectionView_CanDragItemsAtIndexes_WithEvent != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_CanDragItemsAtIndexes_WithEvent(f func(collectionView CollectionView, indexes foundation.IndexSet, event Event) bool) {
	di._CollectionView_CanDragItemsAtIndexes_WithEvent = f
}

func (di *CollectionViewDelegateImpl) CollectionView_CanDragItemsAtIndexes_WithEvent(collectionView CollectionView, indexes foundation.IndexSet, event Event) bool {
	return di._CollectionView_CanDragItemsAtIndexes_WithEvent(collectionView, indexes, event)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_PasteboardWriterForItemAtIndex() bool {
	return di._CollectionView_PasteboardWriterForItemAtIndex != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_PasteboardWriterForItemAtIndex(f func(collectionView CollectionView, index uint) PasteboardWriting) {
	di._CollectionView_PasteboardWriterForItemAtIndex = f
}

func (di *CollectionViewDelegateImpl) CollectionView_PasteboardWriterForItemAtIndex(collectionView CollectionView, index uint) PasteboardWriting {
	return di._CollectionView_PasteboardWriterForItemAtIndex(collectionView, index)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_WriteItemsAtIndexes_ToPasteboard() bool {
	return di._CollectionView_WriteItemsAtIndexes_ToPasteboard != nil
}

// deprecated
func (di *CollectionViewDelegateImpl) SetCollectionView_WriteItemsAtIndexes_ToPasteboard(f func(collectionView CollectionView, indexes foundation.IndexSet, pasteboard Pasteboard) bool) {
	di._CollectionView_WriteItemsAtIndexes_ToPasteboard = f
}

// deprecated
func (di *CollectionViewDelegateImpl) CollectionView_WriteItemsAtIndexes_ToPasteboard(collectionView CollectionView, indexes foundation.IndexSet, pasteboard Pasteboard) bool {
	return di._CollectionView_WriteItemsAtIndexes_ToPasteboard(collectionView, indexes, pasteboard)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes() bool {
	return di._CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes != nil
}

// deprecated
func (di *CollectionViewDelegateImpl) SetCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes(f func(collectionView CollectionView, dropURL foundation.URL, indexes foundation.IndexSet) []string) {
	di._CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes = f
}

// deprecated
func (di *CollectionViewDelegateImpl) CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes(collectionView CollectionView, dropURL foundation.URL, indexes foundation.IndexSet) []string {
	return di._CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes(collectionView, dropURL, indexes)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset() bool {
	return di._CollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset(f func(collectionView CollectionView, indexes foundation.IndexSet, event Event, dragImageOffset *foundation.Point) IImage) {
	di._CollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset(collectionView CollectionView, indexes foundation.IndexSet, event Event, dragImageOffset *foundation.Point) IImage {
	return di._CollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset(collectionView, indexes, event, dragImageOffset)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes() bool {
	return di._CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(f func(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexes foundation.IndexSet)) {
	di._CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes = f
}

func (di *CollectionViewDelegateImpl) CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(collectionView CollectionView, session DraggingSession, screenPoint foundation.Point, indexes foundation.IndexSet) {
	di._CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(collectionView, session, screenPoint, indexes)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_ValidateDrop_ProposedIndex_DropOperation() bool {
	return di._CollectionView_ValidateDrop_ProposedIndex_DropOperation != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_ValidateDrop_ProposedIndex_DropOperation(f func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation) {
	di._CollectionView_ValidateDrop_ProposedIndex_DropOperation = f
}

func (di *CollectionViewDelegateImpl) CollectionView_ValidateDrop_ProposedIndex_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	return di._CollectionView_ValidateDrop_ProposedIndex_DropOperation(collectionView, draggingInfo, proposedDropIndex, proposedDropOperation)
}
func (di *CollectionViewDelegateImpl) ImplementsCollectionView_AcceptDrop_Index_DropOperation() bool {
	return di._CollectionView_AcceptDrop_Index_DropOperation != nil
}

func (di *CollectionViewDelegateImpl) SetCollectionView_AcceptDrop_Index_DropOperation(f func(collectionView CollectionView, draggingInfo DraggingInfoWrapper, index int, dropOperation CollectionViewDropOperation) bool) {
	di._CollectionView_AcceptDrop_Index_DropOperation = f
}

func (di *CollectionViewDelegateImpl) CollectionView_AcceptDrop_Index_DropOperation(collectionView CollectionView, draggingInfo DraggingInfoWrapper, index int, dropOperation CollectionViewDropOperation) bool {
	return di._CollectionView_AcceptDrop_Index_DropOperation(collectionView, draggingInfo, index, dropOperation)
}

type CollectionViewDelegateWrapper struct {
	objc.Object
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_ShouldSelectItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:shouldSelectItemsAtIndexPaths:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_ShouldSelectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) foundation.Set {
	rv := ffi.CallMethod[foundation.Set](c_, "collectionView:shouldSelectItemsAtIndexPaths:", collectionView, indexPaths)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DidSelectItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:didSelectItemsAtIndexPaths:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DidSelectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) {
	ffi.CallMethod[ffi.Void](c_, "collectionView:didSelectItemsAtIndexPaths:", collectionView, indexPaths)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_ShouldDeselectItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:shouldDeselectItemsAtIndexPaths:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_ShouldDeselectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) foundation.Set {
	rv := ffi.CallMethod[foundation.Set](c_, "collectionView:shouldDeselectItemsAtIndexPaths:", collectionView, indexPaths)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DidDeselectItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:didDeselectItemsAtIndexPaths:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DidDeselectItemsAtIndexPaths(collectionView ICollectionView, indexPaths foundation.ISet) {
	ffi.CallMethod[ffi.Void](c_, "collectionView:didDeselectItemsAtIndexPaths:", collectionView, indexPaths)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:shouldChangeItemsAtIndexPaths:toHighlightState:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_ShouldChangeItemsAtIndexPaths_ToHighlightState(collectionView ICollectionView, indexPaths foundation.ISet, highlightState CollectionViewItemHighlightState) foundation.Set {
	rv := ffi.CallMethod[foundation.Set](c_, "collectionView:shouldChangeItemsAtIndexPaths:toHighlightState:", collectionView, indexPaths, highlightState)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DidChangeItemsAtIndexPaths_ToHighlightState() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:didChangeItemsAtIndexPaths:toHighlightState:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DidChangeItemsAtIndexPaths_ToHighlightState(collectionView ICollectionView, indexPaths foundation.ISet, highlightState CollectionViewItemHighlightState) {
	ffi.CallMethod[ffi.Void](c_, "collectionView:didChangeItemsAtIndexPaths:toHighlightState:", collectionView, indexPaths, highlightState)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:willDisplayItem:forRepresentedObjectAtIndexPath:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_WillDisplayItem_ForRepresentedObjectAtIndexPath(collectionView ICollectionView, item ICollectionViewItem, indexPath foundation.IIndexPath) {
	ffi.CallMethod[ffi.Void](c_, "collectionView:willDisplayItem:forRepresentedObjectAtIndexPath:", collectionView, item, indexPath)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:didEndDisplayingItem:forRepresentedObjectAtIndexPath:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DidEndDisplayingItem_ForRepresentedObjectAtIndexPath(collectionView ICollectionView, item ICollectionViewItem, indexPath foundation.IIndexPath) {
	ffi.CallMethod[ffi.Void](c_, "collectionView:didEndDisplayingItem:forRepresentedObjectAtIndexPath:", collectionView, item, indexPath)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:willDisplaySupplementaryView:forElementKind:atIndexPath:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_WillDisplaySupplementaryView_ForElementKind_AtIndexPath(collectionView ICollectionView, view IView, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) {
	ffi.CallMethod[ffi.Void](c_, "collectionView:willDisplaySupplementaryView:forElementKind:atIndexPath:", collectionView, view, elementKind, indexPath)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:didEndDisplayingSupplementaryView:forElementOfKind:atIndexPath:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DidEndDisplayingSupplementaryView_ForElementOfKind_AtIndexPath(collectionView ICollectionView, view IView, elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) {
	ffi.CallMethod[ffi.Void](c_, "collectionView:didEndDisplayingSupplementaryView:forElementOfKind:atIndexPath:", collectionView, view, elementKind, indexPath)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_CanDragItemsAtIndexPaths_WithEvent() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:canDragItemsAtIndexPaths:withEvent:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_CanDragItemsAtIndexPaths_WithEvent(collectionView ICollectionView, indexPaths foundation.ISet, event IEvent) bool {
	rv := ffi.CallMethod[bool](c_, "collectionView:canDragItemsAtIndexPaths:withEvent:", collectionView, indexPaths, event)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_PasteboardWriterForItemAtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:pasteboardWriterForItemAtIndexPath:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_PasteboardWriterForItemAtIndexPath(collectionView ICollectionView, indexPath foundation.IIndexPath) PasteboardWritingWrapper {
	rv := ffi.CallMethod[PasteboardWritingWrapper](c_, "collectionView:pasteboardWriterForItemAtIndexPath:", collectionView, indexPath)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_WriteItemsAtIndexPaths_ToPasteboard() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:writeItemsAtIndexPaths:toPasteboard:"))
}

// deprecated
func (c_ CollectionViewDelegateWrapper) CollectionView_WriteItemsAtIndexPaths_ToPasteboard(collectionView ICollectionView, indexPaths foundation.ISet, pasteboard IPasteboard) bool {
	rv := ffi.CallMethod[bool](c_, "collectionView:writeItemsAtIndexPaths:toPasteboard:", collectionView, indexPaths, pasteboard)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:namesOfPromisedFilesDroppedAtDestination:forDraggedItemsAtIndexPaths:"))
}

// deprecated
func (c_ CollectionViewDelegateWrapper) CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexPaths(collectionView ICollectionView, dropURL foundation.IURL, indexPaths foundation.ISet) []string {
	rv := ffi.CallMethod[[]string](c_, "collectionView:namesOfPromisedFilesDroppedAtDestination:forDraggedItemsAtIndexPaths:", collectionView, dropURL, indexPaths)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:draggingImageForItemsAtIndexPaths:withEvent:offset:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DraggingImageForItemsAtIndexPaths_WithEvent_Offset(collectionView ICollectionView, indexPaths foundation.ISet, event IEvent, dragImageOffset *foundation.Point) Image {
	rv := ffi.CallMethod[Image](c_, "collectionView:draggingImageForItemsAtIndexPaths:withEvent:offset:", collectionView, indexPaths, event, dragImageOffset)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexPaths:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexPaths(collectionView ICollectionView, session IDraggingSession, screenPoint foundation.Point, indexPaths foundation.ISet) {
	ffi.CallMethod[ffi.Void](c_, "collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexPaths:", collectionView, session, screenPoint, indexPaths)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DraggingSession_EndedAtPoint_DragOperation() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:draggingSession:endedAtPoint:dragOperation:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DraggingSession_EndedAtPoint_DragOperation(collectionView ICollectionView, session IDraggingSession, screenPoint foundation.Point, operation DragOperation) {
	ffi.CallMethod[ffi.Void](c_, "collectionView:draggingSession:endedAtPoint:dragOperation:", collectionView, session, screenPoint, operation)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_UpdateDraggingItemsForDrag() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:updateDraggingItemsForDrag:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_UpdateDraggingItemsForDrag(collectionView ICollectionView, draggingInfo DraggingInfo) {
	po := ffi.CreateProtocol(draggingInfo)
	defer po.Release()
	ffi.CallMethod[ffi.Void](c_, "collectionView:updateDraggingItemsForDrag:", collectionView, po)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_ValidateDrop_ProposedIndexPath_DropOperation() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:validateDrop:proposedIndexPath:dropOperation:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_ValidateDrop_ProposedIndexPath_DropOperation(collectionView ICollectionView, draggingInfo DraggingInfo, proposedDropIndexPath *foundation.IndexPath, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	po := ffi.CreateProtocol(draggingInfo)
	defer po.Release()
	rv := ffi.CallMethod[DragOperation](c_, "collectionView:validateDrop:proposedIndexPath:dropOperation:", collectionView, po, unsafe.Pointer(proposedDropIndexPath), proposedDropOperation)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_AcceptDrop_IndexPath_DropOperation() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:acceptDrop:indexPath:dropOperation:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_AcceptDrop_IndexPath_DropOperation(collectionView ICollectionView, draggingInfo DraggingInfo, indexPath foundation.IIndexPath, dropOperation CollectionViewDropOperation) bool {
	po := ffi.CreateProtocol(draggingInfo)
	defer po.Release()
	rv := ffi.CallMethod[bool](c_, "collectionView:acceptDrop:indexPath:dropOperation:", collectionView, po, indexPath, dropOperation)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_CanDragItemsAtIndexes_WithEvent() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:canDragItemsAtIndexes:withEvent:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_CanDragItemsAtIndexes_WithEvent(collectionView ICollectionView, indexes foundation.IIndexSet, event IEvent) bool {
	rv := ffi.CallMethod[bool](c_, "collectionView:canDragItemsAtIndexes:withEvent:", collectionView, indexes, event)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_PasteboardWriterForItemAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:pasteboardWriterForItemAtIndex:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_PasteboardWriterForItemAtIndex(collectionView ICollectionView, index uint) PasteboardWritingWrapper {
	rv := ffi.CallMethod[PasteboardWritingWrapper](c_, "collectionView:pasteboardWriterForItemAtIndex:", collectionView, index)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_WriteItemsAtIndexes_ToPasteboard() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:writeItemsAtIndexes:toPasteboard:"))
}

// deprecated
func (c_ CollectionViewDelegateWrapper) CollectionView_WriteItemsAtIndexes_ToPasteboard(collectionView ICollectionView, indexes foundation.IIndexSet, pasteboard IPasteboard) bool {
	rv := ffi.CallMethod[bool](c_, "collectionView:writeItemsAtIndexes:toPasteboard:", collectionView, indexes, pasteboard)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:namesOfPromisedFilesDroppedAtDestination:forDraggedItemsAtIndexes:"))
}

// deprecated
func (c_ CollectionViewDelegateWrapper) CollectionView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedItemsAtIndexes(collectionView ICollectionView, dropURL foundation.IURL, indexes foundation.IIndexSet) []string {
	rv := ffi.CallMethod[[]string](c_, "collectionView:namesOfPromisedFilesDroppedAtDestination:forDraggedItemsAtIndexes:", collectionView, dropURL, indexes)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:draggingImageForItemsAtIndexes:withEvent:offset:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DraggingImageForItemsAtIndexes_WithEvent_Offset(collectionView ICollectionView, indexes foundation.IIndexSet, event IEvent, dragImageOffset *foundation.Point) Image {
	rv := ffi.CallMethod[Image](c_, "collectionView:draggingImageForItemsAtIndexes:withEvent:offset:", collectionView, indexes, event, dragImageOffset)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexes:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_DraggingSession_WillBeginAtPoint_ForItemsAtIndexes(collectionView ICollectionView, session IDraggingSession, screenPoint foundation.Point, indexes foundation.IIndexSet) {
	ffi.CallMethod[ffi.Void](c_, "collectionView:draggingSession:willBeginAtPoint:forItemsAtIndexes:", collectionView, session, screenPoint, indexes)
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_ValidateDrop_ProposedIndex_DropOperation() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:validateDrop:proposedIndex:dropOperation:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_ValidateDrop_ProposedIndex_DropOperation(collectionView ICollectionView, draggingInfo DraggingInfo, proposedDropIndex *int, proposedDropOperation *CollectionViewDropOperation) DragOperation {
	po := ffi.CreateProtocol(draggingInfo)
	defer po.Release()
	rv := ffi.CallMethod[DragOperation](c_, "collectionView:validateDrop:proposedIndex:dropOperation:", collectionView, po, proposedDropIndex, proposedDropOperation)
	return rv
}

func (c_ *CollectionViewDelegateWrapper) ImplementsCollectionView_AcceptDrop_Index_DropOperation() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:acceptDrop:index:dropOperation:"))
}

func (c_ CollectionViewDelegateWrapper) CollectionView_AcceptDrop_Index_DropOperation(collectionView ICollectionView, draggingInfo DraggingInfo, index int, dropOperation CollectionViewDropOperation) bool {
	po := ffi.CreateProtocol(draggingInfo)
	defer po.Release()
	rv := ffi.CallMethod[bool](c_, "collectionView:acceptDrop:index:dropOperation:", collectionView, po, index, dropOperation)
	return rv
}

var CollectionViewDiffableDataSourceClass = _CollectionViewDiffableDataSourceClass{objc.GetClass("NSCollectionViewDiffableDataSource")}

type _CollectionViewDiffableDataSourceClass struct {
	objc.Class
}

type ICollectionViewDiffableDataSource interface {
	objc.IObject
	ItemIdentifierForIndexPath(indexPath foundation.IIndexPath) objc.Object
	IndexPathForItemIdentifier(identifier objc.IObject) foundation.IndexPath
	Snapshot() DiffableDataSourceSnapshot
	ApplySnapshot_AnimatingDifferences(snapshot IDiffableDataSourceSnapshot, animatingDifferences bool)
	SupplementaryViewProvider() func(param1 ICollectionView, param2 string, param3 foundation.IIndexPath) View
	SetSupplementaryViewProvider(value func(param1 CollectionView, param2 string, param3 foundation.IndexPath) IView)
}

type CollectionViewDiffableDataSource struct {
	objc.Object
}

func MakeCollectionViewDiffableDataSource(ptr unsafe.Pointer) CollectionViewDiffableDataSource {
	return CollectionViewDiffableDataSource{
		Object: objc.MakeObject(ptr),
	}
}

func (c_ CollectionViewDiffableDataSource) InitWithCollectionView_ItemProvider(collectionView ICollectionView, itemProvider func(param1 CollectionView, param2 foundation.IndexPath, param3 objc.Object) ICollectionViewItem) CollectionViewDiffableDataSource {
	rv := ffi.CallMethod[CollectionViewDiffableDataSource](c_, "initWithCollectionView:itemProvider:", collectionView, itemProvider)
	return rv
}

func (cc _CollectionViewDiffableDataSourceClass) Alloc() CollectionViewDiffableDataSource {
	rv := ffi.CallMethod[CollectionViewDiffableDataSource](cc, "alloc")
	return rv
}

func (c_ CollectionViewDiffableDataSource) Init() CollectionViewDiffableDataSource {
	rv := ffi.CallMethod[CollectionViewDiffableDataSource](c_, "init")
	return rv
}

func (cc _CollectionViewDiffableDataSourceClass) New() CollectionViewDiffableDataSource {
	rv := ffi.CallMethod[CollectionViewDiffableDataSource](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewDiffableDataSource() CollectionViewDiffableDataSource {
	return CollectionViewDiffableDataSourceClass.New()
}

func (c_ CollectionViewDiffableDataSource) ItemIdentifierForIndexPath(indexPath foundation.IIndexPath) objc.Object {
	rv := ffi.CallMethod[objc.Object](c_, "itemIdentifierForIndexPath:", indexPath)
	return rv
}

func (c_ CollectionViewDiffableDataSource) IndexPathForItemIdentifier(identifier objc.IObject) foundation.IndexPath {
	rv := ffi.CallMethod[foundation.IndexPath](c_, "indexPathForItemIdentifier:", identifier)
	return rv
}

func (c_ CollectionViewDiffableDataSource) Snapshot() DiffableDataSourceSnapshot {
	rv := ffi.CallMethod[DiffableDataSourceSnapshot](c_, "snapshot")
	return rv
}

func (c_ CollectionViewDiffableDataSource) ApplySnapshot_AnimatingDifferences(snapshot IDiffableDataSourceSnapshot, animatingDifferences bool) {
	ffi.CallMethod[ffi.Void](c_, "applySnapshot:animatingDifferences:", snapshot, animatingDifferences)
}

func (c_ CollectionViewDiffableDataSource) SupplementaryViewProvider() func(param1 ICollectionView, param2 string, param3 foundation.IIndexPath) View {
	rv := ffi.CallMethod[func(param1 ICollectionView, param2 string, param3 foundation.IIndexPath) View](c_, "supplementaryViewProvider")
	return rv
}

func (c_ CollectionViewDiffableDataSource) SetSupplementaryViewProvider(value func(param1 CollectionView, param2 string, param3 foundation.IndexPath) IView) {
	ffi.CallMethod[ffi.Void](c_, "setSupplementaryViewProvider:", value)
}

type CollectionViewSectionHeaderView interface {
	CollectionViewElement
	ImplementsSetSectionCollapseButton() bool
	// optional
	SetSectionCollapseButton(value Button)
	ImplementsSectionCollapseButton() bool
	// optional
	SectionCollapseButton() IButton
}

type CollectionViewSectionHeaderViewWrapper struct {
	CollectionViewElementWrapper
}

func (c_ *CollectionViewSectionHeaderViewWrapper) ImplementsSetSectionCollapseButton() bool {
	return c_.RespondsToSelector(objc.GetSelector("setSectionCollapseButton:"))
}

func (c_ CollectionViewSectionHeaderViewWrapper) SetSectionCollapseButton(value IButton) {
	ffi.CallMethod[ffi.Void](c_, "setSectionCollapseButton:", value)
}

func (c_ *CollectionViewSectionHeaderViewWrapper) ImplementsSectionCollapseButton() bool {
	return c_.RespondsToSelector(objc.GetSelector("sectionCollapseButton"))
}

func (c_ CollectionViewSectionHeaderViewWrapper) SectionCollapseButton() Button {
	rv := ffi.CallMethod[Button](c_, "sectionCollapseButton")
	return rv
}

type CollectionViewElement interface {
	UserInterfaceItemIdentification
	ImplementsPrepareForReuse() bool
	// optional
	PrepareForReuse()
	ImplementsPreferredLayoutAttributesFittingAttributes() bool
	// optional
	PreferredLayoutAttributesFittingAttributes(layoutAttributes CollectionViewLayoutAttributes) ICollectionViewLayoutAttributes
	ImplementsApplyLayoutAttributes() bool
	// optional
	ApplyLayoutAttributes(layoutAttributes CollectionViewLayoutAttributes)
	ImplementsWillTransitionFromLayout_ToLayout() bool
	// optional
	WillTransitionFromLayout_ToLayout(oldLayout CollectionViewLayout, newLayout CollectionViewLayout)
	ImplementsDidTransitionFromLayout_ToLayout() bool
	// optional
	DidTransitionFromLayout_ToLayout(oldLayout CollectionViewLayout, newLayout CollectionViewLayout)
}

type CollectionViewElementWrapper struct {
	objc.Object
	UserInterfaceItemIdentificationWrapper
}

func (c_ *CollectionViewElementWrapper) ImplementsPrepareForReuse() bool {
	return c_.RespondsToSelector(objc.GetSelector("prepareForReuse"))
}

func (c_ CollectionViewElementWrapper) PrepareForReuse() {
	ffi.CallMethod[ffi.Void](c_, "prepareForReuse")
}

func (c_ *CollectionViewElementWrapper) ImplementsPreferredLayoutAttributesFittingAttributes() bool {
	return c_.RespondsToSelector(objc.GetSelector("preferredLayoutAttributesFittingAttributes:"))
}

func (c_ CollectionViewElementWrapper) PreferredLayoutAttributesFittingAttributes(layoutAttributes ICollectionViewLayoutAttributes) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](c_, "preferredLayoutAttributesFittingAttributes:", layoutAttributes)
	return rv
}

func (c_ *CollectionViewElementWrapper) ImplementsApplyLayoutAttributes() bool {
	return c_.RespondsToSelector(objc.GetSelector("applyLayoutAttributes:"))
}

func (c_ CollectionViewElementWrapper) ApplyLayoutAttributes(layoutAttributes ICollectionViewLayoutAttributes) {
	ffi.CallMethod[ffi.Void](c_, "applyLayoutAttributes:", layoutAttributes)
}

func (c_ *CollectionViewElementWrapper) ImplementsWillTransitionFromLayout_ToLayout() bool {
	return c_.RespondsToSelector(objc.GetSelector("willTransitionFromLayout:toLayout:"))
}

func (c_ CollectionViewElementWrapper) WillTransitionFromLayout_ToLayout(oldLayout ICollectionViewLayout, newLayout ICollectionViewLayout) {
	ffi.CallMethod[ffi.Void](c_, "willTransitionFromLayout:toLayout:", oldLayout, newLayout)
}

func (c_ *CollectionViewElementWrapper) ImplementsDidTransitionFromLayout_ToLayout() bool {
	return c_.RespondsToSelector(objc.GetSelector("didTransitionFromLayout:toLayout:"))
}

func (c_ CollectionViewElementWrapper) DidTransitionFromLayout_ToLayout(oldLayout ICollectionViewLayout, newLayout ICollectionViewLayout) {
	ffi.CallMethod[ffi.Void](c_, "didTransitionFromLayout:toLayout:", oldLayout, newLayout)
}

type CollectionViewPrefetching interface {
	ImplementsCollectionView_CancelPrefetchingForItemsAtIndexPaths() bool
	// optional
	CollectionView_CancelPrefetchingForItemsAtIndexPaths(collectionView CollectionView, indexPaths []foundation.IndexPath)
	// required
	CollectionView_PrefetchItemsAtIndexPaths(collectionView CollectionView, indexPaths []foundation.IndexPath)
}

type CollectionViewPrefetchingWrapper struct {
	objc.Object
}

func (c_ *CollectionViewPrefetchingWrapper) ImplementsCollectionView_CancelPrefetchingForItemsAtIndexPaths() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:cancelPrefetchingForItemsAtIndexPaths:"))
}

func (c_ CollectionViewPrefetchingWrapper) CollectionView_CancelPrefetchingForItemsAtIndexPaths(collectionView ICollectionView, indexPaths []foundation.IIndexPath) {
	ffi.CallMethod[ffi.Void](c_, "collectionView:cancelPrefetchingForItemsAtIndexPaths:", collectionView, indexPaths)
}

func (c_ CollectionViewPrefetchingWrapper) CollectionView_PrefetchItemsAtIndexPaths(collectionView ICollectionView, indexPaths []foundation.IIndexPath) {
	ffi.CallMethod[ffi.Void](c_, "collectionView:prefetchItemsAtIndexPaths:", collectionView, indexPaths)
}

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

func (d_ DiffableDataSourceSnapshot) Init() DiffableDataSourceSnapshot {
	rv := ffi.CallMethod[DiffableDataSourceSnapshot](d_, "init")
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

var CollectionViewFlowLayoutClass = _CollectionViewFlowLayoutClass{objc.GetClass("NSCollectionViewFlowLayout")}

type _CollectionViewFlowLayoutClass struct {
	objc.Class
}

type ICollectionViewFlowLayout interface {
	ICollectionViewLayout
	CollapseSectionAtIndex(sectionIndex uint)
	ExpandSectionAtIndex(sectionIndex uint)
	SectionAtIndexIsCollapsed(sectionIndex uint) bool
	ScrollDirection() CollectionViewScrollDirection
	SetScrollDirection(value CollectionViewScrollDirection)
	MinimumLineSpacing() float64
	SetMinimumLineSpacing(value float64)
	MinimumInteritemSpacing() float64
	SetMinimumInteritemSpacing(value float64)
	EstimatedItemSize() foundation.Size
	SetEstimatedItemSize(value foundation.Size)
	ItemSize() foundation.Size
	SetItemSize(value foundation.Size)
	SectionInset() foundation.EdgeInsets
	SetSectionInset(value foundation.EdgeInsets)
	HeaderReferenceSize() foundation.Size
	SetHeaderReferenceSize(value foundation.Size)
	FooterReferenceSize() foundation.Size
	SetFooterReferenceSize(value foundation.Size)
	SectionFootersPinToVisibleBounds() bool
	SetSectionFootersPinToVisibleBounds(value bool)
	SectionHeadersPinToVisibleBounds() bool
	SetSectionHeadersPinToVisibleBounds(value bool)
}

type CollectionViewFlowLayout struct {
	CollectionViewLayout
}

func MakeCollectionViewFlowLayout(ptr unsafe.Pointer) CollectionViewFlowLayout {
	return CollectionViewFlowLayout{
		CollectionViewLayout: MakeCollectionViewLayout(ptr),
	}
}

func (cc _CollectionViewFlowLayoutClass) Alloc() CollectionViewFlowLayout {
	rv := ffi.CallMethod[CollectionViewFlowLayout](cc, "alloc")
	return rv
}

func (c_ CollectionViewFlowLayout) Init() CollectionViewFlowLayout {
	rv := ffi.CallMethod[CollectionViewFlowLayout](c_, "init")
	return rv
}

func (cc _CollectionViewFlowLayoutClass) New() CollectionViewFlowLayout {
	rv := ffi.CallMethod[CollectionViewFlowLayout](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewFlowLayout() CollectionViewFlowLayout {
	return CollectionViewFlowLayoutClass.New()
}

func (c_ CollectionViewFlowLayout) CollapseSectionAtIndex(sectionIndex uint) {
	ffi.CallMethod[ffi.Void](c_, "collapseSectionAtIndex:", sectionIndex)
}

func (c_ CollectionViewFlowLayout) ExpandSectionAtIndex(sectionIndex uint) {
	ffi.CallMethod[ffi.Void](c_, "expandSectionAtIndex:", sectionIndex)
}

func (c_ CollectionViewFlowLayout) SectionAtIndexIsCollapsed(sectionIndex uint) bool {
	rv := ffi.CallMethod[bool](c_, "sectionAtIndexIsCollapsed:", sectionIndex)
	return rv
}

func (c_ CollectionViewFlowLayout) ScrollDirection() CollectionViewScrollDirection {
	rv := ffi.CallMethod[CollectionViewScrollDirection](c_, "scrollDirection")
	return rv
}

func (c_ CollectionViewFlowLayout) SetScrollDirection(value CollectionViewScrollDirection) {
	ffi.CallMethod[ffi.Void](c_, "setScrollDirection:", value)
}

func (c_ CollectionViewFlowLayout) MinimumLineSpacing() float64 {
	rv := ffi.CallMethod[float64](c_, "minimumLineSpacing")
	return rv
}

func (c_ CollectionViewFlowLayout) SetMinimumLineSpacing(value float64) {
	ffi.CallMethod[ffi.Void](c_, "setMinimumLineSpacing:", value)
}

func (c_ CollectionViewFlowLayout) MinimumInteritemSpacing() float64 {
	rv := ffi.CallMethod[float64](c_, "minimumInteritemSpacing")
	return rv
}

func (c_ CollectionViewFlowLayout) SetMinimumInteritemSpacing(value float64) {
	ffi.CallMethod[ffi.Void](c_, "setMinimumInteritemSpacing:", value)
}

func (c_ CollectionViewFlowLayout) EstimatedItemSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "estimatedItemSize")
	return rv
}

func (c_ CollectionViewFlowLayout) SetEstimatedItemSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](c_, "setEstimatedItemSize:", value)
}

func (c_ CollectionViewFlowLayout) ItemSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "itemSize")
	return rv
}

func (c_ CollectionViewFlowLayout) SetItemSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](c_, "setItemSize:", value)
}

func (c_ CollectionViewFlowLayout) SectionInset() foundation.EdgeInsets {
	rv := ffi.CallMethod[foundation.EdgeInsets](c_, "sectionInset")
	return rv
}

func (c_ CollectionViewFlowLayout) SetSectionInset(value foundation.EdgeInsets) {
	ffi.CallMethod[ffi.Void](c_, "setSectionInset:", value)
}

func (c_ CollectionViewFlowLayout) HeaderReferenceSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "headerReferenceSize")
	return rv
}

func (c_ CollectionViewFlowLayout) SetHeaderReferenceSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](c_, "setHeaderReferenceSize:", value)
}

func (c_ CollectionViewFlowLayout) FooterReferenceSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "footerReferenceSize")
	return rv
}

func (c_ CollectionViewFlowLayout) SetFooterReferenceSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](c_, "setFooterReferenceSize:", value)
}

func (c_ CollectionViewFlowLayout) SectionFootersPinToVisibleBounds() bool {
	rv := ffi.CallMethod[bool](c_, "sectionFootersPinToVisibleBounds")
	return rv
}

func (c_ CollectionViewFlowLayout) SetSectionFootersPinToVisibleBounds(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setSectionFootersPinToVisibleBounds:", value)
}

func (c_ CollectionViewFlowLayout) SectionHeadersPinToVisibleBounds() bool {
	rv := ffi.CallMethod[bool](c_, "sectionHeadersPinToVisibleBounds")
	return rv
}

func (c_ CollectionViewFlowLayout) SetSectionHeadersPinToVisibleBounds(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setSectionHeadersPinToVisibleBounds:", value)
}

var CollectionViewFlowLayoutInvalidationContextClass = _CollectionViewFlowLayoutInvalidationContextClass{objc.GetClass("NSCollectionViewFlowLayoutInvalidationContext")}

type _CollectionViewFlowLayoutInvalidationContextClass struct {
	objc.Class
}

type ICollectionViewFlowLayoutInvalidationContext interface {
	ICollectionViewLayoutInvalidationContext
	InvalidateFlowLayoutAttributes() bool
	SetInvalidateFlowLayoutAttributes(value bool)
	InvalidateFlowLayoutDelegateMetrics() bool
	SetInvalidateFlowLayoutDelegateMetrics(value bool)
}

type CollectionViewFlowLayoutInvalidationContext struct {
	CollectionViewLayoutInvalidationContext
}

func MakeCollectionViewFlowLayoutInvalidationContext(ptr unsafe.Pointer) CollectionViewFlowLayoutInvalidationContext {
	return CollectionViewFlowLayoutInvalidationContext{
		CollectionViewLayoutInvalidationContext: MakeCollectionViewLayoutInvalidationContext(ptr),
	}
}

func (cc _CollectionViewFlowLayoutInvalidationContextClass) Alloc() CollectionViewFlowLayoutInvalidationContext {
	rv := ffi.CallMethod[CollectionViewFlowLayoutInvalidationContext](cc, "alloc")
	return rv
}

func (c_ CollectionViewFlowLayoutInvalidationContext) Init() CollectionViewFlowLayoutInvalidationContext {
	rv := ffi.CallMethod[CollectionViewFlowLayoutInvalidationContext](c_, "init")
	return rv
}

func (cc _CollectionViewFlowLayoutInvalidationContextClass) New() CollectionViewFlowLayoutInvalidationContext {
	rv := ffi.CallMethod[CollectionViewFlowLayoutInvalidationContext](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewFlowLayoutInvalidationContext() CollectionViewFlowLayoutInvalidationContext {
	return CollectionViewFlowLayoutInvalidationContextClass.New()
}

func (c_ CollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutAttributes() bool {
	rv := ffi.CallMethod[bool](c_, "invalidateFlowLayoutAttributes")
	return rv
}

func (c_ CollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutAttributes(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setInvalidateFlowLayoutAttributes:", value)
}

func (c_ CollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutDelegateMetrics() bool {
	rv := ffi.CallMethod[bool](c_, "invalidateFlowLayoutDelegateMetrics")
	return rv
}

func (c_ CollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutDelegateMetrics(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setInvalidateFlowLayoutDelegateMetrics:", value)
}

type CollectionViewDelegateFlowLayout interface {
	CollectionViewDelegate
	ImplementsCollectionView_Layout_SizeForItemAtIndexPath() bool
	// optional
	CollectionView_Layout_SizeForItemAtIndexPath(collectionView CollectionView, collectionViewLayout CollectionViewLayout, indexPath foundation.IndexPath) foundation.Size
	ImplementsCollectionView_Layout_InsetForSectionAtIndex() bool
	// optional
	CollectionView_Layout_InsetForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.EdgeInsets
	ImplementsCollectionView_Layout_MinimumLineSpacingForSectionAtIndex() bool
	// optional
	CollectionView_Layout_MinimumLineSpacingForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64
	ImplementsCollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex() bool
	// optional
	CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64
	ImplementsCollectionView_Layout_ReferenceSizeForHeaderInSection() bool
	// optional
	CollectionView_Layout_ReferenceSizeForHeaderInSection(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size
	ImplementsCollectionView_Layout_ReferenceSizeForFooterInSection() bool
	// optional
	CollectionView_Layout_ReferenceSizeForFooterInSection(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size
}

type CollectionViewDelegateFlowLayoutImpl struct {
	CollectionViewDelegateImpl
	_CollectionView_Layout_SizeForItemAtIndexPath                   func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, indexPath foundation.IndexPath) foundation.Size
	_CollectionView_Layout_InsetForSectionAtIndex                   func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.EdgeInsets
	_CollectionView_Layout_MinimumLineSpacingForSectionAtIndex      func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64
	_CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64
	_CollectionView_Layout_ReferenceSizeForHeaderInSection          func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size
	_CollectionView_Layout_ReferenceSizeForFooterInSection          func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size
}

func (di *CollectionViewDelegateFlowLayoutImpl) ImplementsCollectionView_Layout_SizeForItemAtIndexPath() bool {
	return di._CollectionView_Layout_SizeForItemAtIndexPath != nil
}

func (di *CollectionViewDelegateFlowLayoutImpl) SetCollectionView_Layout_SizeForItemAtIndexPath(f func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, indexPath foundation.IndexPath) foundation.Size) {
	di._CollectionView_Layout_SizeForItemAtIndexPath = f
}

func (di *CollectionViewDelegateFlowLayoutImpl) CollectionView_Layout_SizeForItemAtIndexPath(collectionView CollectionView, collectionViewLayout CollectionViewLayout, indexPath foundation.IndexPath) foundation.Size {
	return di._CollectionView_Layout_SizeForItemAtIndexPath(collectionView, collectionViewLayout, indexPath)
}
func (di *CollectionViewDelegateFlowLayoutImpl) ImplementsCollectionView_Layout_InsetForSectionAtIndex() bool {
	return di._CollectionView_Layout_InsetForSectionAtIndex != nil
}

func (di *CollectionViewDelegateFlowLayoutImpl) SetCollectionView_Layout_InsetForSectionAtIndex(f func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.EdgeInsets) {
	di._CollectionView_Layout_InsetForSectionAtIndex = f
}

func (di *CollectionViewDelegateFlowLayoutImpl) CollectionView_Layout_InsetForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.EdgeInsets {
	return di._CollectionView_Layout_InsetForSectionAtIndex(collectionView, collectionViewLayout, section)
}
func (di *CollectionViewDelegateFlowLayoutImpl) ImplementsCollectionView_Layout_MinimumLineSpacingForSectionAtIndex() bool {
	return di._CollectionView_Layout_MinimumLineSpacingForSectionAtIndex != nil
}

func (di *CollectionViewDelegateFlowLayoutImpl) SetCollectionView_Layout_MinimumLineSpacingForSectionAtIndex(f func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64) {
	di._CollectionView_Layout_MinimumLineSpacingForSectionAtIndex = f
}

func (di *CollectionViewDelegateFlowLayoutImpl) CollectionView_Layout_MinimumLineSpacingForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64 {
	return di._CollectionView_Layout_MinimumLineSpacingForSectionAtIndex(collectionView, collectionViewLayout, section)
}
func (di *CollectionViewDelegateFlowLayoutImpl) ImplementsCollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex() bool {
	return di._CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex != nil
}

func (di *CollectionViewDelegateFlowLayoutImpl) SetCollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex(f func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64) {
	di._CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex = f
}

func (di *CollectionViewDelegateFlowLayoutImpl) CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) float64 {
	return di._CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex(collectionView, collectionViewLayout, section)
}
func (di *CollectionViewDelegateFlowLayoutImpl) ImplementsCollectionView_Layout_ReferenceSizeForHeaderInSection() bool {
	return di._CollectionView_Layout_ReferenceSizeForHeaderInSection != nil
}

func (di *CollectionViewDelegateFlowLayoutImpl) SetCollectionView_Layout_ReferenceSizeForHeaderInSection(f func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size) {
	di._CollectionView_Layout_ReferenceSizeForHeaderInSection = f
}

func (di *CollectionViewDelegateFlowLayoutImpl) CollectionView_Layout_ReferenceSizeForHeaderInSection(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size {
	return di._CollectionView_Layout_ReferenceSizeForHeaderInSection(collectionView, collectionViewLayout, section)
}
func (di *CollectionViewDelegateFlowLayoutImpl) ImplementsCollectionView_Layout_ReferenceSizeForFooterInSection() bool {
	return di._CollectionView_Layout_ReferenceSizeForFooterInSection != nil
}

func (di *CollectionViewDelegateFlowLayoutImpl) SetCollectionView_Layout_ReferenceSizeForFooterInSection(f func(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size) {
	di._CollectionView_Layout_ReferenceSizeForFooterInSection = f
}

func (di *CollectionViewDelegateFlowLayoutImpl) CollectionView_Layout_ReferenceSizeForFooterInSection(collectionView CollectionView, collectionViewLayout CollectionViewLayout, section int) foundation.Size {
	return di._CollectionView_Layout_ReferenceSizeForFooterInSection(collectionView, collectionViewLayout, section)
}

type CollectionViewDelegateFlowLayoutWrapper struct {
	CollectionViewDelegateWrapper
}

func (c_ *CollectionViewDelegateFlowLayoutWrapper) ImplementsCollectionView_Layout_SizeForItemAtIndexPath() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:layout:sizeForItemAtIndexPath:"))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_Layout_SizeForItemAtIndexPath(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, indexPath foundation.IIndexPath) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "collectionView:layout:sizeForItemAtIndexPath:", collectionView, collectionViewLayout, indexPath)
	return rv
}

func (c_ *CollectionViewDelegateFlowLayoutWrapper) ImplementsCollectionView_Layout_InsetForSectionAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:layout:insetForSectionAtIndex:"))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_Layout_InsetForSectionAtIndex(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) foundation.EdgeInsets {
	rv := ffi.CallMethod[foundation.EdgeInsets](c_, "collectionView:layout:insetForSectionAtIndex:", collectionView, collectionViewLayout, section)
	return rv
}

func (c_ *CollectionViewDelegateFlowLayoutWrapper) ImplementsCollectionView_Layout_MinimumLineSpacingForSectionAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:layout:minimumLineSpacingForSectionAtIndex:"))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_Layout_MinimumLineSpacingForSectionAtIndex(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) float64 {
	rv := ffi.CallMethod[float64](c_, "collectionView:layout:minimumLineSpacingForSectionAtIndex:", collectionView, collectionViewLayout, section)
	return rv
}

func (c_ *CollectionViewDelegateFlowLayoutWrapper) ImplementsCollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:layout:minimumInteritemSpacingForSectionAtIndex:"))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_Layout_MinimumInteritemSpacingForSectionAtIndex(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) float64 {
	rv := ffi.CallMethod[float64](c_, "collectionView:layout:minimumInteritemSpacingForSectionAtIndex:", collectionView, collectionViewLayout, section)
	return rv
}

func (c_ *CollectionViewDelegateFlowLayoutWrapper) ImplementsCollectionView_Layout_ReferenceSizeForHeaderInSection() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:layout:referenceSizeForHeaderInSection:"))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_Layout_ReferenceSizeForHeaderInSection(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "collectionView:layout:referenceSizeForHeaderInSection:", collectionView, collectionViewLayout, section)
	return rv
}

func (c_ *CollectionViewDelegateFlowLayoutWrapper) ImplementsCollectionView_Layout_ReferenceSizeForFooterInSection() bool {
	return c_.RespondsToSelector(objc.GetSelector("collectionView:layout:referenceSizeForFooterInSection:"))
}

func (c_ CollectionViewDelegateFlowLayoutWrapper) CollectionView_Layout_ReferenceSizeForFooterInSection(collectionView ICollectionView, collectionViewLayout ICollectionViewLayout, section int) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "collectionView:layout:referenceSizeForFooterInSection:", collectionView, collectionViewLayout, section)
	return rv
}

var CollectionViewLayoutClass = _CollectionViewLayoutClass{objc.GetClass("NSCollectionViewLayout")}

type _CollectionViewLayoutClass struct {
	objc.Class
}

type ICollectionViewLayout interface {
	objc.IObject
	PrepareLayout()
	LayoutAttributesForElementsInRect(rect foundation.Rect) []CollectionViewLayoutAttributes
	LayoutAttributesForItemAtIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	LayoutAttributesForSupplementaryViewOfKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	LayoutAttributesForDecorationViewOfKind_AtIndexPath(elementKind CollectionViewDecorationElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	LayoutAttributesForDropTargetAtPoint(pointInCollectionView foundation.Point) CollectionViewLayoutAttributes
	LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	TargetContentOffsetForProposedContentOffset(proposedContentOffset foundation.Point) foundation.Point
	TargetContentOffsetForProposedContentOffset_WithScrollingVelocity(proposedContentOffset foundation.Point, velocity foundation.Point) foundation.Point
	PrepareForCollectionViewUpdates(updateItems []ICollectionViewUpdateItem)
	FinalizeCollectionViewUpdates()
	IndexPathsToInsertForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set
	IndexPathsToInsertForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set
	InitialLayoutAttributesForAppearingItemAtIndexPath(itemIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	InitialLayoutAttributesForAppearingSupplementaryElementOfKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, elementIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	InitialLayoutAttributesForAppearingDecorationElementOfKind_AtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	IndexPathsToDeleteForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set
	IndexPathsToDeleteForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set
	FinalLayoutAttributesForDisappearingItemAtIndexPath(itemIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	FinalLayoutAttributesForDisappearingSupplementaryElementOfKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, elementIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	FinalLayoutAttributesForDisappearingDecorationElementOfKind_AtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes
	InvalidateLayout()
	InvalidateLayoutWithContext(context ICollectionViewLayoutInvalidationContext)
	ShouldInvalidateLayoutForBoundsChange(newBounds foundation.Rect) bool
	ShouldInvalidateLayoutForPreferredLayoutAttributes_WithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) bool
	InvalidationContextForBoundsChange(newBounds foundation.Rect) CollectionViewLayoutInvalidationContext
	InvalidationContextForPreferredLayoutAttributes_WithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) CollectionViewLayoutInvalidationContext
	PrepareForAnimatedBoundsChange(oldBounds foundation.Rect)
	FinalizeAnimatedBoundsChange()
	PrepareForTransitionFromLayout(oldLayout ICollectionViewLayout)
	PrepareForTransitionToLayout(newLayout ICollectionViewLayout)
	FinalizeLayoutTransition()
	CollectionView() CollectionView
	CollectionViewContentSize() foundation.Size
}

type CollectionViewLayout struct {
	objc.Object
}

func MakeCollectionViewLayout(ptr unsafe.Pointer) CollectionViewLayout {
	return CollectionViewLayout{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionViewLayoutClass) Alloc() CollectionViewLayout {
	rv := ffi.CallMethod[CollectionViewLayout](cc, "alloc")
	return rv
}

func (c_ CollectionViewLayout) Init() CollectionViewLayout {
	rv := ffi.CallMethod[CollectionViewLayout](c_, "init")
	return rv
}

func (cc _CollectionViewLayoutClass) New() CollectionViewLayout {
	rv := ffi.CallMethod[CollectionViewLayout](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewLayout() CollectionViewLayout {
	return CollectionViewLayoutClass.New()
}

func (c_ CollectionViewLayout) PrepareLayout() {
	ffi.CallMethod[ffi.Void](c_, "prepareLayout")
}

func (c_ CollectionViewLayout) LayoutAttributesForElementsInRect(rect foundation.Rect) []CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[[]CollectionViewLayoutAttributes](c_, "layoutAttributesForElementsInRect:", rect)
	return rv
}

func (c_ CollectionViewLayout) LayoutAttributesForItemAtIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](c_, "layoutAttributesForItemAtIndexPath:", indexPath)
	return rv
}

func (c_ CollectionViewLayout) LayoutAttributesForSupplementaryViewOfKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](c_, "layoutAttributesForSupplementaryViewOfKind:atIndexPath:", elementKind, indexPath)
	return rv
}

func (c_ CollectionViewLayout) LayoutAttributesForDecorationViewOfKind_AtIndexPath(elementKind CollectionViewDecorationElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](c_, "layoutAttributesForDecorationViewOfKind:atIndexPath:", elementKind, indexPath)
	return rv
}

func (c_ CollectionViewLayout) LayoutAttributesForDropTargetAtPoint(pointInCollectionView foundation.Point) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](c_, "layoutAttributesForDropTargetAtPoint:", pointInCollectionView)
	return rv
}

func (c_ CollectionViewLayout) LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](c_, "layoutAttributesForInterItemGapBeforeIndexPath:", indexPath)
	return rv
}

func (c_ CollectionViewLayout) TargetContentOffsetForProposedContentOffset(proposedContentOffset foundation.Point) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](c_, "targetContentOffsetForProposedContentOffset:", proposedContentOffset)
	return rv
}

func (c_ CollectionViewLayout) TargetContentOffsetForProposedContentOffset_WithScrollingVelocity(proposedContentOffset foundation.Point, velocity foundation.Point) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](c_, "targetContentOffsetForProposedContentOffset:withScrollingVelocity:", proposedContentOffset, velocity)
	return rv
}

func (c_ CollectionViewLayout) PrepareForCollectionViewUpdates(updateItems []ICollectionViewUpdateItem) {
	ffi.CallMethod[ffi.Void](c_, "prepareForCollectionViewUpdates:", updateItems)
}

func (c_ CollectionViewLayout) FinalizeCollectionViewUpdates() {
	ffi.CallMethod[ffi.Void](c_, "finalizeCollectionViewUpdates")
}

func (c_ CollectionViewLayout) IndexPathsToInsertForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set {
	rv := ffi.CallMethod[foundation.Set](c_, "indexPathsToInsertForSupplementaryViewOfKind:", elementKind)
	return rv
}

func (c_ CollectionViewLayout) IndexPathsToInsertForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set {
	rv := ffi.CallMethod[foundation.Set](c_, "indexPathsToInsertForDecorationViewOfKind:", elementKind)
	return rv
}

func (c_ CollectionViewLayout) InitialLayoutAttributesForAppearingItemAtIndexPath(itemIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](c_, "initialLayoutAttributesForAppearingItemAtIndexPath:", itemIndexPath)
	return rv
}

func (c_ CollectionViewLayout) InitialLayoutAttributesForAppearingSupplementaryElementOfKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, elementIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](c_, "initialLayoutAttributesForAppearingSupplementaryElementOfKind:atIndexPath:", elementKind, elementIndexPath)
	return rv
}

func (c_ CollectionViewLayout) InitialLayoutAttributesForAppearingDecorationElementOfKind_AtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](c_, "initialLayoutAttributesForAppearingDecorationElementOfKind:atIndexPath:", elementKind, decorationIndexPath)
	return rv
}

func (c_ CollectionViewLayout) IndexPathsToDeleteForSupplementaryViewOfKind(elementKind CollectionViewSupplementaryElementKind) foundation.Set {
	rv := ffi.CallMethod[foundation.Set](c_, "indexPathsToDeleteForSupplementaryViewOfKind:", elementKind)
	return rv
}

func (c_ CollectionViewLayout) IndexPathsToDeleteForDecorationViewOfKind(elementKind CollectionViewDecorationElementKind) foundation.Set {
	rv := ffi.CallMethod[foundation.Set](c_, "indexPathsToDeleteForDecorationViewOfKind:", elementKind)
	return rv
}

func (c_ CollectionViewLayout) FinalLayoutAttributesForDisappearingItemAtIndexPath(itemIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](c_, "finalLayoutAttributesForDisappearingItemAtIndexPath:", itemIndexPath)
	return rv
}

func (c_ CollectionViewLayout) FinalLayoutAttributesForDisappearingSupplementaryElementOfKind_AtIndexPath(elementKind CollectionViewSupplementaryElementKind, elementIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](c_, "finalLayoutAttributesForDisappearingSupplementaryElementOfKind:atIndexPath:", elementKind, elementIndexPath)
	return rv
}

func (c_ CollectionViewLayout) FinalLayoutAttributesForDisappearingDecorationElementOfKind_AtIndexPath(elementKind CollectionViewDecorationElementKind, decorationIndexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](c_, "finalLayoutAttributesForDisappearingDecorationElementOfKind:atIndexPath:", elementKind, decorationIndexPath)
	return rv
}

func (c_ CollectionViewLayout) InvalidateLayout() {
	ffi.CallMethod[ffi.Void](c_, "invalidateLayout")
}

func (c_ CollectionViewLayout) InvalidateLayoutWithContext(context ICollectionViewLayoutInvalidationContext) {
	ffi.CallMethod[ffi.Void](c_, "invalidateLayoutWithContext:", context)
}

func (c_ CollectionViewLayout) ShouldInvalidateLayoutForBoundsChange(newBounds foundation.Rect) bool {
	rv := ffi.CallMethod[bool](c_, "shouldInvalidateLayoutForBoundsChange:", newBounds)
	return rv
}

func (c_ CollectionViewLayout) ShouldInvalidateLayoutForPreferredLayoutAttributes_WithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) bool {
	rv := ffi.CallMethod[bool](c_, "shouldInvalidateLayoutForPreferredLayoutAttributes:withOriginalAttributes:", preferredAttributes, originalAttributes)
	return rv
}

func (c_ CollectionViewLayout) InvalidationContextForBoundsChange(newBounds foundation.Rect) CollectionViewLayoutInvalidationContext {
	rv := ffi.CallMethod[CollectionViewLayoutInvalidationContext](c_, "invalidationContextForBoundsChange:", newBounds)
	return rv
}

func (c_ CollectionViewLayout) InvalidationContextForPreferredLayoutAttributes_WithOriginalAttributes(preferredAttributes ICollectionViewLayoutAttributes, originalAttributes ICollectionViewLayoutAttributes) CollectionViewLayoutInvalidationContext {
	rv := ffi.CallMethod[CollectionViewLayoutInvalidationContext](c_, "invalidationContextForPreferredLayoutAttributes:withOriginalAttributes:", preferredAttributes, originalAttributes)
	return rv
}

func (c_ CollectionViewLayout) PrepareForAnimatedBoundsChange(oldBounds foundation.Rect) {
	ffi.CallMethod[ffi.Void](c_, "prepareForAnimatedBoundsChange:", oldBounds)
}

func (c_ CollectionViewLayout) FinalizeAnimatedBoundsChange() {
	ffi.CallMethod[ffi.Void](c_, "finalizeAnimatedBoundsChange")
}

func (c_ CollectionViewLayout) PrepareForTransitionFromLayout(oldLayout ICollectionViewLayout) {
	ffi.CallMethod[ffi.Void](c_, "prepareForTransitionFromLayout:", oldLayout)
}

func (c_ CollectionViewLayout) PrepareForTransitionToLayout(newLayout ICollectionViewLayout) {
	ffi.CallMethod[ffi.Void](c_, "prepareForTransitionToLayout:", newLayout)
}

func (c_ CollectionViewLayout) FinalizeLayoutTransition() {
	ffi.CallMethod[ffi.Void](c_, "finalizeLayoutTransition")
}

func (c_ CollectionViewLayout) CollectionView() CollectionView {
	rv := ffi.CallMethod[CollectionView](c_, "collectionView")
	return rv
}

func (c_ CollectionViewLayout) CollectionViewContentSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "collectionViewContentSize")
	return rv
}

var CollectionViewLayoutAttributesClass = _CollectionViewLayoutAttributesClass{objc.GetClass("NSCollectionViewLayoutAttributes")}

type _CollectionViewLayoutAttributesClass struct {
	objc.Class
}

type ICollectionViewLayoutAttributes interface {
	objc.IObject
	RepresentedElementCategory() CollectionElementCategory
	IndexPath() foundation.IndexPath
	SetIndexPath(value foundation.IIndexPath)
	RepresentedElementKind() string
	Frame() foundation.Rect
	SetFrame(value foundation.Rect)
	Size() foundation.Size
	SetSize(value foundation.Size)
	Alpha() float64
	SetAlpha(value float64)
	IsHidden() bool
	SetHidden(value bool)
	ZIndex() int
	SetZIndex(value int)
}

type CollectionViewLayoutAttributes struct {
	objc.Object
}

func MakeCollectionViewLayoutAttributes(ptr unsafe.Pointer) CollectionViewLayoutAttributes {
	return CollectionViewLayoutAttributes{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForItemWithIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](cc, "layoutAttributesForItemWithIndexPath:", indexPath)
	return rv
}

func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForSupplementaryViewOfKind_WithIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](cc, "layoutAttributesForSupplementaryViewOfKind:withIndexPath:", elementKind, indexPath)
	return rv
}

func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForDecorationViewOfKind_WithIndexPath(decorationViewKind CollectionViewDecorationElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](cc, "layoutAttributesForDecorationViewOfKind:withIndexPath:", decorationViewKind, indexPath)
	return rv
}

func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](cc, "layoutAttributesForInterItemGapBeforeIndexPath:", indexPath)
	return rv
}

func (cc _CollectionViewLayoutAttributesClass) Alloc() CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](cc, "alloc")
	return rv
}

func (c_ CollectionViewLayoutAttributes) Init() CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](c_, "init")
	return rv
}

func (cc _CollectionViewLayoutAttributesClass) New() CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewLayoutAttributes() CollectionViewLayoutAttributes {
	return CollectionViewLayoutAttributesClass.New()
}

func (c_ CollectionViewLayoutAttributes) RepresentedElementCategory() CollectionElementCategory {
	rv := ffi.CallMethod[CollectionElementCategory](c_, "representedElementCategory")
	return rv
}

func (c_ CollectionViewLayoutAttributes) IndexPath() foundation.IndexPath {
	rv := ffi.CallMethod[foundation.IndexPath](c_, "indexPath")
	return rv
}

func (c_ CollectionViewLayoutAttributes) SetIndexPath(value foundation.IIndexPath) {
	ffi.CallMethod[ffi.Void](c_, "setIndexPath:", value)
}

func (c_ CollectionViewLayoutAttributes) RepresentedElementKind() string {
	rv := ffi.CallMethod[string](c_, "representedElementKind")
	return rv
}

func (c_ CollectionViewLayoutAttributes) Frame() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](c_, "frame")
	return rv
}

func (c_ CollectionViewLayoutAttributes) SetFrame(value foundation.Rect) {
	ffi.CallMethod[ffi.Void](c_, "setFrame:", value)
}

func (c_ CollectionViewLayoutAttributes) Size() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "size")
	return rv
}

func (c_ CollectionViewLayoutAttributes) SetSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](c_, "setSize:", value)
}

func (c_ CollectionViewLayoutAttributes) Alpha() float64 {
	rv := ffi.CallMethod[float64](c_, "alpha")
	return rv
}

func (c_ CollectionViewLayoutAttributes) SetAlpha(value float64) {
	ffi.CallMethod[ffi.Void](c_, "setAlpha:", value)
}

func (c_ CollectionViewLayoutAttributes) IsHidden() bool {
	rv := ffi.CallMethod[bool](c_, "isHidden")
	return rv
}

func (c_ CollectionViewLayoutAttributes) SetHidden(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setHidden:", value)
}

func (c_ CollectionViewLayoutAttributes) ZIndex() int {
	rv := ffi.CallMethod[int](c_, "zIndex")
	return rv
}

func (c_ CollectionViewLayoutAttributes) SetZIndex(value int) {
	ffi.CallMethod[ffi.Void](c_, "setZIndex:", value)
}

var CollectionViewGridLayoutClass = _CollectionViewGridLayoutClass{objc.GetClass("NSCollectionViewGridLayout")}

type _CollectionViewGridLayoutClass struct {
	objc.Class
}

type ICollectionViewGridLayout interface {
	ICollectionViewLayout
	MaximumNumberOfRows() uint
	SetMaximumNumberOfRows(value uint)
	MaximumNumberOfColumns() uint
	SetMaximumNumberOfColumns(value uint)
	MinimumItemSize() foundation.Size
	SetMinimumItemSize(value foundation.Size)
	MaximumItemSize() foundation.Size
	SetMaximumItemSize(value foundation.Size)
	MinimumInteritemSpacing() float64
	SetMinimumInteritemSpacing(value float64)
	MinimumLineSpacing() float64
	SetMinimumLineSpacing(value float64)
	Margins() foundation.EdgeInsets
	SetMargins(value foundation.EdgeInsets)
	BackgroundColors() []Color
	SetBackgroundColors(value []IColor)
}

type CollectionViewGridLayout struct {
	CollectionViewLayout
}

func MakeCollectionViewGridLayout(ptr unsafe.Pointer) CollectionViewGridLayout {
	return CollectionViewGridLayout{
		CollectionViewLayout: MakeCollectionViewLayout(ptr),
	}
}

func (cc _CollectionViewGridLayoutClass) Alloc() CollectionViewGridLayout {
	rv := ffi.CallMethod[CollectionViewGridLayout](cc, "alloc")
	return rv
}

func (c_ CollectionViewGridLayout) Init() CollectionViewGridLayout {
	rv := ffi.CallMethod[CollectionViewGridLayout](c_, "init")
	return rv
}

func (cc _CollectionViewGridLayoutClass) New() CollectionViewGridLayout {
	rv := ffi.CallMethod[CollectionViewGridLayout](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewGridLayout() CollectionViewGridLayout {
	return CollectionViewGridLayoutClass.New()
}

func (c_ CollectionViewGridLayout) MaximumNumberOfRows() uint {
	rv := ffi.CallMethod[uint](c_, "maximumNumberOfRows")
	return rv
}

func (c_ CollectionViewGridLayout) SetMaximumNumberOfRows(value uint) {
	ffi.CallMethod[ffi.Void](c_, "setMaximumNumberOfRows:", value)
}

func (c_ CollectionViewGridLayout) MaximumNumberOfColumns() uint {
	rv := ffi.CallMethod[uint](c_, "maximumNumberOfColumns")
	return rv
}

func (c_ CollectionViewGridLayout) SetMaximumNumberOfColumns(value uint) {
	ffi.CallMethod[ffi.Void](c_, "setMaximumNumberOfColumns:", value)
}

func (c_ CollectionViewGridLayout) MinimumItemSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "minimumItemSize")
	return rv
}

func (c_ CollectionViewGridLayout) SetMinimumItemSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](c_, "setMinimumItemSize:", value)
}

func (c_ CollectionViewGridLayout) MaximumItemSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "maximumItemSize")
	return rv
}

func (c_ CollectionViewGridLayout) SetMaximumItemSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](c_, "setMaximumItemSize:", value)
}

func (c_ CollectionViewGridLayout) MinimumInteritemSpacing() float64 {
	rv := ffi.CallMethod[float64](c_, "minimumInteritemSpacing")
	return rv
}

func (c_ CollectionViewGridLayout) SetMinimumInteritemSpacing(value float64) {
	ffi.CallMethod[ffi.Void](c_, "setMinimumInteritemSpacing:", value)
}

func (c_ CollectionViewGridLayout) MinimumLineSpacing() float64 {
	rv := ffi.CallMethod[float64](c_, "minimumLineSpacing")
	return rv
}

func (c_ CollectionViewGridLayout) SetMinimumLineSpacing(value float64) {
	ffi.CallMethod[ffi.Void](c_, "setMinimumLineSpacing:", value)
}

func (c_ CollectionViewGridLayout) Margins() foundation.EdgeInsets {
	rv := ffi.CallMethod[foundation.EdgeInsets](c_, "margins")
	return rv
}

func (c_ CollectionViewGridLayout) SetMargins(value foundation.EdgeInsets) {
	ffi.CallMethod[ffi.Void](c_, "setMargins:", value)
}

func (c_ CollectionViewGridLayout) BackgroundColors() []Color {
	rv := ffi.CallMethod[[]Color](c_, "backgroundColors")
	return rv
}

func (c_ CollectionViewGridLayout) SetBackgroundColors(value []IColor) {
	ffi.CallMethod[ffi.Void](c_, "setBackgroundColors:", value)
}

var CollectionViewLayoutInvalidationContextClass = _CollectionViewLayoutInvalidationContextClass{objc.GetClass("NSCollectionViewLayoutInvalidationContext")}

type _CollectionViewLayoutInvalidationContextClass struct {
	objc.Class
}

type ICollectionViewLayoutInvalidationContext interface {
	objc.IObject
	InvalidateItemsAtIndexPaths(indexPaths foundation.ISet)
	InvalidateSupplementaryElementsOfKind_AtIndexPaths(elementKind CollectionViewSupplementaryElementKind, indexPaths foundation.ISet)
	InvalidateDecorationElementsOfKind_AtIndexPaths(elementKind CollectionViewDecorationElementKind, indexPaths foundation.ISet)
	InvalidateEverything() bool
	InvalidateDataSourceCounts() bool
	ContentOffsetAdjustment() foundation.Point
	SetContentOffsetAdjustment(value foundation.Point)
	ContentSizeAdjustment() foundation.Size
	SetContentSizeAdjustment(value foundation.Size)
	InvalidatedItemIndexPaths() foundation.Set
	InvalidatedSupplementaryIndexPaths() map[CollectionViewSupplementaryElementKind]foundation.Set
	InvalidatedDecorationIndexPaths() map[CollectionViewDecorationElementKind]foundation.Set
}

type CollectionViewLayoutInvalidationContext struct {
	objc.Object
}

func MakeCollectionViewLayoutInvalidationContext(ptr unsafe.Pointer) CollectionViewLayoutInvalidationContext {
	return CollectionViewLayoutInvalidationContext{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionViewLayoutInvalidationContextClass) Alloc() CollectionViewLayoutInvalidationContext {
	rv := ffi.CallMethod[CollectionViewLayoutInvalidationContext](cc, "alloc")
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) Init() CollectionViewLayoutInvalidationContext {
	rv := ffi.CallMethod[CollectionViewLayoutInvalidationContext](c_, "init")
	return rv
}

func (cc _CollectionViewLayoutInvalidationContextClass) New() CollectionViewLayoutInvalidationContext {
	rv := ffi.CallMethod[CollectionViewLayoutInvalidationContext](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewLayoutInvalidationContext() CollectionViewLayoutInvalidationContext {
	return CollectionViewLayoutInvalidationContextClass.New()
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateItemsAtIndexPaths(indexPaths foundation.ISet) {
	ffi.CallMethod[ffi.Void](c_, "invalidateItemsAtIndexPaths:", indexPaths)
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateSupplementaryElementsOfKind_AtIndexPaths(elementKind CollectionViewSupplementaryElementKind, indexPaths foundation.ISet) {
	ffi.CallMethod[ffi.Void](c_, "invalidateSupplementaryElementsOfKind:atIndexPaths:", elementKind, indexPaths)
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateDecorationElementsOfKind_AtIndexPaths(elementKind CollectionViewDecorationElementKind, indexPaths foundation.ISet) {
	ffi.CallMethod[ffi.Void](c_, "invalidateDecorationElementsOfKind:atIndexPaths:", elementKind, indexPaths)
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateEverything() bool {
	rv := ffi.CallMethod[bool](c_, "invalidateEverything")
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidateDataSourceCounts() bool {
	rv := ffi.CallMethod[bool](c_, "invalidateDataSourceCounts")
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) ContentOffsetAdjustment() foundation.Point {
	rv := ffi.CallMethod[foundation.Point](c_, "contentOffsetAdjustment")
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) SetContentOffsetAdjustment(value foundation.Point) {
	ffi.CallMethod[ffi.Void](c_, "setContentOffsetAdjustment:", value)
}

func (c_ CollectionViewLayoutInvalidationContext) ContentSizeAdjustment() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "contentSizeAdjustment")
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) SetContentSizeAdjustment(value foundation.Size) {
	ffi.CallMethod[ffi.Void](c_, "setContentSizeAdjustment:", value)
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidatedItemIndexPaths() foundation.Set {
	rv := ffi.CallMethod[foundation.Set](c_, "invalidatedItemIndexPaths")
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidatedSupplementaryIndexPaths() map[CollectionViewSupplementaryElementKind]foundation.Set {
	rv := ffi.CallMethod[map[CollectionViewSupplementaryElementKind]foundation.Set](c_, "invalidatedSupplementaryIndexPaths")
	return rv
}

func (c_ CollectionViewLayoutInvalidationContext) InvalidatedDecorationIndexPaths() map[CollectionViewDecorationElementKind]foundation.Set {
	rv := ffi.CallMethod[map[CollectionViewDecorationElementKind]foundation.Set](c_, "invalidatedDecorationIndexPaths")
	return rv
}

var CollectionViewUpdateItemClass = _CollectionViewUpdateItemClass{objc.GetClass("NSCollectionViewUpdateItem")}

type _CollectionViewUpdateItemClass struct {
	objc.Class
}

type ICollectionViewUpdateItem interface {
	objc.IObject
	IndexPathBeforeUpdate() foundation.IndexPath
	IndexPathAfterUpdate() foundation.IndexPath
	UpdateAction() CollectionUpdateAction
}

type CollectionViewUpdateItem struct {
	objc.Object
}

func MakeCollectionViewUpdateItem(ptr unsafe.Pointer) CollectionViewUpdateItem {
	return CollectionViewUpdateItem{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionViewUpdateItemClass) Alloc() CollectionViewUpdateItem {
	rv := ffi.CallMethod[CollectionViewUpdateItem](cc, "alloc")
	return rv
}

func (c_ CollectionViewUpdateItem) Init() CollectionViewUpdateItem {
	rv := ffi.CallMethod[CollectionViewUpdateItem](c_, "init")
	return rv
}

func (cc _CollectionViewUpdateItemClass) New() CollectionViewUpdateItem {
	rv := ffi.CallMethod[CollectionViewUpdateItem](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewUpdateItem() CollectionViewUpdateItem {
	return CollectionViewUpdateItemClass.New()
}

func (c_ CollectionViewUpdateItem) IndexPathBeforeUpdate() foundation.IndexPath {
	rv := ffi.CallMethod[foundation.IndexPath](c_, "indexPathBeforeUpdate")
	return rv
}

func (c_ CollectionViewUpdateItem) IndexPathAfterUpdate() foundation.IndexPath {
	rv := ffi.CallMethod[foundation.IndexPath](c_, "indexPathAfterUpdate")
	return rv
}

func (c_ CollectionViewUpdateItem) UpdateAction() CollectionUpdateAction {
	rv := ffi.CallMethod[CollectionUpdateAction](c_, "updateAction")
	return rv
}

var CollectionViewCompositionalLayoutClass = _CollectionViewCompositionalLayoutClass{objc.GetClass("NSCollectionViewCompositionalLayout")}

type _CollectionViewCompositionalLayoutClass struct {
	objc.Class
}

type ICollectionViewCompositionalLayout interface {
	ICollectionViewLayout
	Configuration() CollectionViewCompositionalLayoutConfiguration
	SetConfiguration(value ICollectionViewCompositionalLayoutConfiguration)
}

type CollectionViewCompositionalLayout struct {
	CollectionViewLayout
}

func MakeCollectionViewCompositionalLayout(ptr unsafe.Pointer) CollectionViewCompositionalLayout {
	return CollectionViewCompositionalLayout{
		CollectionViewLayout: MakeCollectionViewLayout(ptr),
	}
}

func (c_ CollectionViewCompositionalLayout) InitWithSection(section ICollectionLayoutSection) CollectionViewCompositionalLayout {
	rv := ffi.CallMethod[CollectionViewCompositionalLayout](c_, "initWithSection:", section)
	return rv
}

func (c_ CollectionViewCompositionalLayout) InitWithSection_Configuration(section ICollectionLayoutSection, configuration ICollectionViewCompositionalLayoutConfiguration) CollectionViewCompositionalLayout {
	rv := ffi.CallMethod[CollectionViewCompositionalLayout](c_, "initWithSection:configuration:", section, configuration)
	return rv
}

func (cc _CollectionViewCompositionalLayoutClass) Alloc() CollectionViewCompositionalLayout {
	rv := ffi.CallMethod[CollectionViewCompositionalLayout](cc, "alloc")
	return rv
}

func (c_ CollectionViewCompositionalLayout) Init() CollectionViewCompositionalLayout {
	rv := ffi.CallMethod[CollectionViewCompositionalLayout](c_, "init")
	return rv
}

func (cc _CollectionViewCompositionalLayoutClass) New() CollectionViewCompositionalLayout {
	rv := ffi.CallMethod[CollectionViewCompositionalLayout](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewCompositionalLayout() CollectionViewCompositionalLayout {
	return CollectionViewCompositionalLayoutClass.New()
}

func (c_ CollectionViewCompositionalLayout) Configuration() CollectionViewCompositionalLayoutConfiguration {
	rv := ffi.CallMethod[CollectionViewCompositionalLayoutConfiguration](c_, "configuration")
	return rv
}

func (c_ CollectionViewCompositionalLayout) SetConfiguration(value ICollectionViewCompositionalLayoutConfiguration) {
	ffi.CallMethod[ffi.Void](c_, "setConfiguration:", value)
}

var CollectionViewCompositionalLayoutConfigurationClass = _CollectionViewCompositionalLayoutConfigurationClass{objc.GetClass("NSCollectionViewCompositionalLayoutConfiguration")}

type _CollectionViewCompositionalLayoutConfigurationClass struct {
	objc.Class
}

type ICollectionViewCompositionalLayoutConfiguration interface {
	objc.IObject
	ScrollDirection() CollectionViewScrollDirection
	SetScrollDirection(value CollectionViewScrollDirection)
	InterSectionSpacing() float64
	SetInterSectionSpacing(value float64)
	BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem
	SetBoundarySupplementaryItems(value []ICollectionLayoutBoundarySupplementaryItem)
}

type CollectionViewCompositionalLayoutConfiguration struct {
	objc.Object
}

func MakeCollectionViewCompositionalLayoutConfiguration(ptr unsafe.Pointer) CollectionViewCompositionalLayoutConfiguration {
	return CollectionViewCompositionalLayoutConfiguration{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionViewCompositionalLayoutConfigurationClass) Alloc() CollectionViewCompositionalLayoutConfiguration {
	rv := ffi.CallMethod[CollectionViewCompositionalLayoutConfiguration](cc, "alloc")
	return rv
}

func (c_ CollectionViewCompositionalLayoutConfiguration) Init() CollectionViewCompositionalLayoutConfiguration {
	rv := ffi.CallMethod[CollectionViewCompositionalLayoutConfiguration](c_, "init")
	return rv
}

func (cc _CollectionViewCompositionalLayoutConfigurationClass) New() CollectionViewCompositionalLayoutConfiguration {
	rv := ffi.CallMethod[CollectionViewCompositionalLayoutConfiguration](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewCompositionalLayoutConfiguration() CollectionViewCompositionalLayoutConfiguration {
	return CollectionViewCompositionalLayoutConfigurationClass.New()
}

func (c_ CollectionViewCompositionalLayoutConfiguration) ScrollDirection() CollectionViewScrollDirection {
	rv := ffi.CallMethod[CollectionViewScrollDirection](c_, "scrollDirection")
	return rv
}

func (c_ CollectionViewCompositionalLayoutConfiguration) SetScrollDirection(value CollectionViewScrollDirection) {
	ffi.CallMethod[ffi.Void](c_, "setScrollDirection:", value)
}

func (c_ CollectionViewCompositionalLayoutConfiguration) InterSectionSpacing() float64 {
	rv := ffi.CallMethod[float64](c_, "interSectionSpacing")
	return rv
}

func (c_ CollectionViewCompositionalLayoutConfiguration) SetInterSectionSpacing(value float64) {
	ffi.CallMethod[ffi.Void](c_, "setInterSectionSpacing:", value)
}

func (c_ CollectionViewCompositionalLayoutConfiguration) BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[[]CollectionLayoutBoundarySupplementaryItem](c_, "boundarySupplementaryItems")
	return rv
}

func (c_ CollectionViewCompositionalLayoutConfiguration) SetBoundarySupplementaryItems(value []ICollectionLayoutBoundarySupplementaryItem) {
	ffi.CallMethod[ffi.Void](c_, "setBoundarySupplementaryItems:", value)
}

var CollectionLayoutItemClass = _CollectionLayoutItemClass{objc.GetClass("NSCollectionLayoutItem")}

type _CollectionLayoutItemClass struct {
	objc.Class
}

type ICollectionLayoutItem interface {
	objc.IObject
	LayoutSize() CollectionLayoutSize
	SupplementaryItems() []CollectionLayoutSupplementaryItem
	EdgeSpacing() CollectionLayoutEdgeSpacing
	SetEdgeSpacing(value ICollectionLayoutEdgeSpacing)
	ContentInsets() DirectionalEdgeInsets
	SetContentInsets(value DirectionalEdgeInsets)
}

type CollectionLayoutItem struct {
	objc.Object
}

func MakeCollectionLayoutItem(ptr unsafe.Pointer) CollectionLayoutItem {
	return CollectionLayoutItem{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutItem {
	rv := ffi.CallMethod[CollectionLayoutItem](cc, "itemWithLayoutSize:", layoutSize)
	return rv
}

func (cc _CollectionLayoutItemClass) ItemWithLayoutSize_SupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutItem {
	rv := ffi.CallMethod[CollectionLayoutItem](cc, "itemWithLayoutSize:supplementaryItems:", layoutSize, supplementaryItems)
	return rv
}

func (cc _CollectionLayoutItemClass) Alloc() CollectionLayoutItem {
	rv := ffi.CallMethod[CollectionLayoutItem](cc, "alloc")
	return rv
}

func (c_ CollectionLayoutItem) Init() CollectionLayoutItem {
	rv := ffi.CallMethod[CollectionLayoutItem](c_, "init")
	return rv
}

func (cc _CollectionLayoutItemClass) New() CollectionLayoutItem {
	rv := ffi.CallMethod[CollectionLayoutItem](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutItem() CollectionLayoutItem {
	return CollectionLayoutItemClass.New()
}

func (c_ CollectionLayoutItem) LayoutSize() CollectionLayoutSize {
	rv := ffi.CallMethod[CollectionLayoutSize](c_, "layoutSize")
	return rv
}

func (c_ CollectionLayoutItem) SupplementaryItems() []CollectionLayoutSupplementaryItem {
	rv := ffi.CallMethod[[]CollectionLayoutSupplementaryItem](c_, "supplementaryItems")
	return rv
}

func (c_ CollectionLayoutItem) EdgeSpacing() CollectionLayoutEdgeSpacing {
	rv := ffi.CallMethod[CollectionLayoutEdgeSpacing](c_, "edgeSpacing")
	return rv
}

func (c_ CollectionLayoutItem) SetEdgeSpacing(value ICollectionLayoutEdgeSpacing) {
	ffi.CallMethod[ffi.Void](c_, "setEdgeSpacing:", value)
}

func (c_ CollectionLayoutItem) ContentInsets() DirectionalEdgeInsets {
	rv := ffi.CallMethod[DirectionalEdgeInsets](c_, "contentInsets")
	return rv
}

func (c_ CollectionLayoutItem) SetContentInsets(value DirectionalEdgeInsets) {
	ffi.CallMethod[ffi.Void](c_, "setContentInsets:", value)
}

var CollectionLayoutBoundarySupplementaryItemClass = _CollectionLayoutBoundarySupplementaryItemClass{objc.GetClass("NSCollectionLayoutBoundarySupplementaryItem")}

type _CollectionLayoutBoundarySupplementaryItemClass struct {
	objc.Class
}

type ICollectionLayoutBoundarySupplementaryItem interface {
	ICollectionLayoutSupplementaryItem
	PinToVisibleBounds() bool
	SetPinToVisibleBounds(value bool)
	Offset() foundation.Point
	Alignment() RectAlignment
	ExtendsBoundary() bool
	SetExtendsBoundary(value bool)
}

type CollectionLayoutBoundarySupplementaryItem struct {
	CollectionLayoutSupplementaryItem
}

func MakeCollectionLayoutBoundarySupplementaryItem(ptr unsafe.Pointer) CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItem{
		CollectionLayoutSupplementaryItem: MakeCollectionLayoutSupplementaryItem(ptr),
	}
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment(layoutSize ICollectionLayoutSize, elementKind string, alignment RectAlignment) CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, "boundarySupplementaryItemWithLayoutSize:elementKind:alignment:", layoutSize, elementKind, alignment)
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment_AbsoluteOffset(layoutSize ICollectionLayoutSize, elementKind string, alignment RectAlignment, absoluteOffset foundation.Point) CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, "boundarySupplementaryItemWithLayoutSize:elementKind:alignment:absoluteOffset:", layoutSize, elementKind, alignment, absoluteOffset)
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor) CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, "supplementaryItemWithLayoutSize:elementKind:containerAnchor:", layoutSize, elementKind, containerAnchor)
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor, itemAnchor ICollectionLayoutAnchor) CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, "supplementaryItemWithLayoutSize:elementKind:containerAnchor:itemAnchor:", layoutSize, elementKind, containerAnchor, itemAnchor)
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, "itemWithLayoutSize:", layoutSize)
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) ItemWithLayoutSize_SupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, "itemWithLayoutSize:supplementaryItems:", layoutSize, supplementaryItems)
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) Alloc() CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, "alloc")
	return rv
}

func (c_ CollectionLayoutBoundarySupplementaryItem) Init() CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutBoundarySupplementaryItem](c_, "init")
	return rv
}

func (cc _CollectionLayoutBoundarySupplementaryItemClass) New() CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutBoundarySupplementaryItem](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutBoundarySupplementaryItem() CollectionLayoutBoundarySupplementaryItem {
	return CollectionLayoutBoundarySupplementaryItemClass.New()
}

func (c_ CollectionLayoutBoundarySupplementaryItem) PinToVisibleBounds() bool {
	rv := ffi.CallMethod[bool](c_, "pinToVisibleBounds")
	return rv
}

func (c_ CollectionLayoutBoundarySupplementaryItem) SetPinToVisibleBounds(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setPinToVisibleBounds:", value)
}

func (c_ CollectionLayoutBoundarySupplementaryItem) Offset() foundation.Point {
	rv := ffi.CallMethod[foundation.Point](c_, "offset")
	return rv
}

func (c_ CollectionLayoutBoundarySupplementaryItem) Alignment() RectAlignment {
	rv := ffi.CallMethod[RectAlignment](c_, "alignment")
	return rv
}

func (c_ CollectionLayoutBoundarySupplementaryItem) ExtendsBoundary() bool {
	rv := ffi.CallMethod[bool](c_, "extendsBoundary")
	return rv
}

func (c_ CollectionLayoutBoundarySupplementaryItem) SetExtendsBoundary(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setExtendsBoundary:", value)
}

var CollectionLayoutSpacingClass = _CollectionLayoutSpacingClass{objc.GetClass("NSCollectionLayoutSpacing")}

type _CollectionLayoutSpacingClass struct {
	objc.Class
}

type ICollectionLayoutSpacing interface {
	objc.IObject
	Spacing() float64
	IsFixedSpacing() bool
	IsFlexibleSpacing() bool
}

type CollectionLayoutSpacing struct {
	objc.Object
}

func MakeCollectionLayoutSpacing(ptr unsafe.Pointer) CollectionLayoutSpacing {
	return CollectionLayoutSpacing{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutSpacingClass) FixedSpacing(fixedSpacing float64) CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](cc, "fixedSpacing:", fixedSpacing)
	return rv
}

func (cc _CollectionLayoutSpacingClass) FlexibleSpacing(flexibleSpacing float64) CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](cc, "flexibleSpacing:", flexibleSpacing)
	return rv
}

func (cc _CollectionLayoutSpacingClass) Alloc() CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](cc, "alloc")
	return rv
}

func (c_ CollectionLayoutSpacing) Init() CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](c_, "init")
	return rv
}

func (cc _CollectionLayoutSpacingClass) New() CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutSpacing() CollectionLayoutSpacing {
	return CollectionLayoutSpacingClass.New()
}

func (c_ CollectionLayoutSpacing) Spacing() float64 {
	rv := ffi.CallMethod[float64](c_, "spacing")
	return rv
}

func (c_ CollectionLayoutSpacing) IsFixedSpacing() bool {
	rv := ffi.CallMethod[bool](c_, "isFixedSpacing")
	return rv
}

func (c_ CollectionLayoutSpacing) IsFlexibleSpacing() bool {
	rv := ffi.CallMethod[bool](c_, "isFlexibleSpacing")
	return rv
}

var CollectionLayoutSectionClass = _CollectionLayoutSectionClass{objc.GetClass("NSCollectionLayoutSection")}

type _CollectionLayoutSectionClass struct {
	objc.Class
}

type ICollectionLayoutSection interface {
	objc.IObject
	OrthogonalScrollingBehavior() CollectionLayoutSectionOrthogonalScrollingBehavior
	SetOrthogonalScrollingBehavior(value CollectionLayoutSectionOrthogonalScrollingBehavior)
	InterGroupSpacing() float64
	SetInterGroupSpacing(value float64)
	ContentInsets() DirectionalEdgeInsets
	SetContentInsets(value DirectionalEdgeInsets)
	BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem
	SetBoundarySupplementaryItems(value []ICollectionLayoutBoundarySupplementaryItem)
	DecorationItems() []CollectionLayoutDecorationItem
	SetDecorationItems(value []ICollectionLayoutDecorationItem)
	SupplementariesFollowContentInsets() bool
	SetSupplementariesFollowContentInsets(value bool)
}

type CollectionLayoutSection struct {
	objc.Object
}

func MakeCollectionLayoutSection(ptr unsafe.Pointer) CollectionLayoutSection {
	return CollectionLayoutSection{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutSectionClass) SectionWithGroup(group ICollectionLayoutGroup) CollectionLayoutSection {
	rv := ffi.CallMethod[CollectionLayoutSection](cc, "sectionWithGroup:", group)
	return rv
}

func (cc _CollectionLayoutSectionClass) Alloc() CollectionLayoutSection {
	rv := ffi.CallMethod[CollectionLayoutSection](cc, "alloc")
	return rv
}

func (c_ CollectionLayoutSection) Init() CollectionLayoutSection {
	rv := ffi.CallMethod[CollectionLayoutSection](c_, "init")
	return rv
}

func (cc _CollectionLayoutSectionClass) New() CollectionLayoutSection {
	rv := ffi.CallMethod[CollectionLayoutSection](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutSection() CollectionLayoutSection {
	return CollectionLayoutSectionClass.New()
}

func (c_ CollectionLayoutSection) OrthogonalScrollingBehavior() CollectionLayoutSectionOrthogonalScrollingBehavior {
	rv := ffi.CallMethod[CollectionLayoutSectionOrthogonalScrollingBehavior](c_, "orthogonalScrollingBehavior")
	return rv
}

func (c_ CollectionLayoutSection) SetOrthogonalScrollingBehavior(value CollectionLayoutSectionOrthogonalScrollingBehavior) {
	ffi.CallMethod[ffi.Void](c_, "setOrthogonalScrollingBehavior:", value)
}

func (c_ CollectionLayoutSection) InterGroupSpacing() float64 {
	rv := ffi.CallMethod[float64](c_, "interGroupSpacing")
	return rv
}

func (c_ CollectionLayoutSection) SetInterGroupSpacing(value float64) {
	ffi.CallMethod[ffi.Void](c_, "setInterGroupSpacing:", value)
}

func (c_ CollectionLayoutSection) ContentInsets() DirectionalEdgeInsets {
	rv := ffi.CallMethod[DirectionalEdgeInsets](c_, "contentInsets")
	return rv
}

func (c_ CollectionLayoutSection) SetContentInsets(value DirectionalEdgeInsets) {
	ffi.CallMethod[ffi.Void](c_, "setContentInsets:", value)
}

func (c_ CollectionLayoutSection) BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[[]CollectionLayoutBoundarySupplementaryItem](c_, "boundarySupplementaryItems")
	return rv
}

func (c_ CollectionLayoutSection) SetBoundarySupplementaryItems(value []ICollectionLayoutBoundarySupplementaryItem) {
	ffi.CallMethod[ffi.Void](c_, "setBoundarySupplementaryItems:", value)
}

func (c_ CollectionLayoutSection) DecorationItems() []CollectionLayoutDecorationItem {
	rv := ffi.CallMethod[[]CollectionLayoutDecorationItem](c_, "decorationItems")
	return rv
}

func (c_ CollectionLayoutSection) SetDecorationItems(value []ICollectionLayoutDecorationItem) {
	ffi.CallMethod[ffi.Void](c_, "setDecorationItems:", value)
}

func (c_ CollectionLayoutSection) SupplementariesFollowContentInsets() bool {
	rv := ffi.CallMethod[bool](c_, "supplementariesFollowContentInsets")
	return rv
}

func (c_ CollectionLayoutSection) SetSupplementariesFollowContentInsets(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setSupplementariesFollowContentInsets:", value)
}

var CollectionLayoutGroupCustomItemClass = _CollectionLayoutGroupCustomItemClass{objc.GetClass("NSCollectionLayoutGroupCustomItem")}

type _CollectionLayoutGroupCustomItemClass struct {
	objc.Class
}

type ICollectionLayoutGroupCustomItem interface {
	objc.IObject
	Frame() foundation.Rect
	ZIndex() int
}

type CollectionLayoutGroupCustomItem struct {
	objc.Object
}

func MakeCollectionLayoutGroupCustomItem(ptr unsafe.Pointer) CollectionLayoutGroupCustomItem {
	return CollectionLayoutGroupCustomItem{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutGroupCustomItemClass) CustomItemWithFrame(frame foundation.Rect) CollectionLayoutGroupCustomItem {
	rv := ffi.CallMethod[CollectionLayoutGroupCustomItem](cc, "customItemWithFrame:", frame)
	return rv
}

func (cc _CollectionLayoutGroupCustomItemClass) CustomItemWithFrame_ZIndex(frame foundation.Rect, zIndex int) CollectionLayoutGroupCustomItem {
	rv := ffi.CallMethod[CollectionLayoutGroupCustomItem](cc, "customItemWithFrame:zIndex:", frame, zIndex)
	return rv
}

func (cc _CollectionLayoutGroupCustomItemClass) Alloc() CollectionLayoutGroupCustomItem {
	rv := ffi.CallMethod[CollectionLayoutGroupCustomItem](cc, "alloc")
	return rv
}

func (c_ CollectionLayoutGroupCustomItem) Init() CollectionLayoutGroupCustomItem {
	rv := ffi.CallMethod[CollectionLayoutGroupCustomItem](c_, "init")
	return rv
}

func (cc _CollectionLayoutGroupCustomItemClass) New() CollectionLayoutGroupCustomItem {
	rv := ffi.CallMethod[CollectionLayoutGroupCustomItem](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutGroupCustomItem() CollectionLayoutGroupCustomItem {
	return CollectionLayoutGroupCustomItemClass.New()
}

func (c_ CollectionLayoutGroupCustomItem) Frame() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](c_, "frame")
	return rv
}

func (c_ CollectionLayoutGroupCustomItem) ZIndex() int {
	rv := ffi.CallMethod[int](c_, "zIndex")
	return rv
}

var CollectionLayoutSupplementaryItemClass = _CollectionLayoutSupplementaryItemClass{objc.GetClass("NSCollectionLayoutSupplementaryItem")}

type _CollectionLayoutSupplementaryItemClass struct {
	objc.Class
}

type ICollectionLayoutSupplementaryItem interface {
	ICollectionLayoutItem
	ItemAnchor() CollectionLayoutAnchor
	ContainerAnchor() CollectionLayoutAnchor
	ElementKind() string
	ZIndex() int
	SetZIndex(value int)
}

type CollectionLayoutSupplementaryItem struct {
	CollectionLayoutItem
}

func MakeCollectionLayoutSupplementaryItem(ptr unsafe.Pointer) CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItem{
		CollectionLayoutItem: MakeCollectionLayoutItem(ptr),
	}
}

func (cc _CollectionLayoutSupplementaryItemClass) SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor) CollectionLayoutSupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutSupplementaryItem](cc, "supplementaryItemWithLayoutSize:elementKind:containerAnchor:", layoutSize, elementKind, containerAnchor)
	return rv
}

func (cc _CollectionLayoutSupplementaryItemClass) SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(layoutSize ICollectionLayoutSize, elementKind string, containerAnchor ICollectionLayoutAnchor, itemAnchor ICollectionLayoutAnchor) CollectionLayoutSupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutSupplementaryItem](cc, "supplementaryItemWithLayoutSize:elementKind:containerAnchor:itemAnchor:", layoutSize, elementKind, containerAnchor, itemAnchor)
	return rv
}

func (cc _CollectionLayoutSupplementaryItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutSupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutSupplementaryItem](cc, "itemWithLayoutSize:", layoutSize)
	return rv
}

func (cc _CollectionLayoutSupplementaryItemClass) ItemWithLayoutSize_SupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutSupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutSupplementaryItem](cc, "itemWithLayoutSize:supplementaryItems:", layoutSize, supplementaryItems)
	return rv
}

func (cc _CollectionLayoutSupplementaryItemClass) Alloc() CollectionLayoutSupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutSupplementaryItem](cc, "alloc")
	return rv
}

func (c_ CollectionLayoutSupplementaryItem) Init() CollectionLayoutSupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutSupplementaryItem](c_, "init")
	return rv
}

func (cc _CollectionLayoutSupplementaryItemClass) New() CollectionLayoutSupplementaryItem {
	rv := ffi.CallMethod[CollectionLayoutSupplementaryItem](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutSupplementaryItem() CollectionLayoutSupplementaryItem {
	return CollectionLayoutSupplementaryItemClass.New()
}

func (c_ CollectionLayoutSupplementaryItem) ItemAnchor() CollectionLayoutAnchor {
	rv := ffi.CallMethod[CollectionLayoutAnchor](c_, "itemAnchor")
	return rv
}

func (c_ CollectionLayoutSupplementaryItem) ContainerAnchor() CollectionLayoutAnchor {
	rv := ffi.CallMethod[CollectionLayoutAnchor](c_, "containerAnchor")
	return rv
}

func (c_ CollectionLayoutSupplementaryItem) ElementKind() string {
	rv := ffi.CallMethod[string](c_, "elementKind")
	return rv
}

func (c_ CollectionLayoutSupplementaryItem) ZIndex() int {
	rv := ffi.CallMethod[int](c_, "zIndex")
	return rv
}

func (c_ CollectionLayoutSupplementaryItem) SetZIndex(value int) {
	ffi.CallMethod[ffi.Void](c_, "setZIndex:", value)
}

var CollectionLayoutSizeClass = _CollectionLayoutSizeClass{objc.GetClass("NSCollectionLayoutSize")}

type _CollectionLayoutSizeClass struct {
	objc.Class
}

type ICollectionLayoutSize interface {
	objc.IObject
	WidthDimension() CollectionLayoutDimension
	HeightDimension() CollectionLayoutDimension
}

type CollectionLayoutSize struct {
	objc.Object
}

func MakeCollectionLayoutSize(ptr unsafe.Pointer) CollectionLayoutSize {
	return CollectionLayoutSize{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutSizeClass) SizeWithWidthDimension_HeightDimension(width ICollectionLayoutDimension, height ICollectionLayoutDimension) CollectionLayoutSize {
	rv := ffi.CallMethod[CollectionLayoutSize](cc, "sizeWithWidthDimension:heightDimension:", width, height)
	return rv
}

func (cc _CollectionLayoutSizeClass) Alloc() CollectionLayoutSize {
	rv := ffi.CallMethod[CollectionLayoutSize](cc, "alloc")
	return rv
}

func (c_ CollectionLayoutSize) Init() CollectionLayoutSize {
	rv := ffi.CallMethod[CollectionLayoutSize](c_, "init")
	return rv
}

func (cc _CollectionLayoutSizeClass) New() CollectionLayoutSize {
	rv := ffi.CallMethod[CollectionLayoutSize](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutSize() CollectionLayoutSize {
	return CollectionLayoutSizeClass.New()
}

func (c_ CollectionLayoutSize) WidthDimension() CollectionLayoutDimension {
	rv := ffi.CallMethod[CollectionLayoutDimension](c_, "widthDimension")
	return rv
}

func (c_ CollectionLayoutSize) HeightDimension() CollectionLayoutDimension {
	rv := ffi.CallMethod[CollectionLayoutDimension](c_, "heightDimension")
	return rv
}

var CollectionLayoutEdgeSpacingClass = _CollectionLayoutEdgeSpacingClass{objc.GetClass("NSCollectionLayoutEdgeSpacing")}

type _CollectionLayoutEdgeSpacingClass struct {
	objc.Class
}

type ICollectionLayoutEdgeSpacing interface {
	objc.IObject
	Leading() CollectionLayoutSpacing
	Top() CollectionLayoutSpacing
	Trailing() CollectionLayoutSpacing
	Bottom() CollectionLayoutSpacing
}

type CollectionLayoutEdgeSpacing struct {
	objc.Object
}

func MakeCollectionLayoutEdgeSpacing(ptr unsafe.Pointer) CollectionLayoutEdgeSpacing {
	return CollectionLayoutEdgeSpacing{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutEdgeSpacingClass) SpacingForLeading_Top_Trailing_Bottom(leading ICollectionLayoutSpacing, top ICollectionLayoutSpacing, trailing ICollectionLayoutSpacing, bottom ICollectionLayoutSpacing) CollectionLayoutEdgeSpacing {
	rv := ffi.CallMethod[CollectionLayoutEdgeSpacing](cc, "spacingForLeading:top:trailing:bottom:", leading, top, trailing, bottom)
	return rv
}

func (cc _CollectionLayoutEdgeSpacingClass) Alloc() CollectionLayoutEdgeSpacing {
	rv := ffi.CallMethod[CollectionLayoutEdgeSpacing](cc, "alloc")
	return rv
}

func (c_ CollectionLayoutEdgeSpacing) Init() CollectionLayoutEdgeSpacing {
	rv := ffi.CallMethod[CollectionLayoutEdgeSpacing](c_, "init")
	return rv
}

func (cc _CollectionLayoutEdgeSpacingClass) New() CollectionLayoutEdgeSpacing {
	rv := ffi.CallMethod[CollectionLayoutEdgeSpacing](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutEdgeSpacing() CollectionLayoutEdgeSpacing {
	return CollectionLayoutEdgeSpacingClass.New()
}

func (c_ CollectionLayoutEdgeSpacing) Leading() CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](c_, "leading")
	return rv
}

func (c_ CollectionLayoutEdgeSpacing) Top() CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](c_, "top")
	return rv
}

func (c_ CollectionLayoutEdgeSpacing) Trailing() CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](c_, "trailing")
	return rv
}

func (c_ CollectionLayoutEdgeSpacing) Bottom() CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](c_, "bottom")
	return rv
}

var CollectionLayoutAnchorClass = _CollectionLayoutAnchorClass{objc.GetClass("NSCollectionLayoutAnchor")}

type _CollectionLayoutAnchorClass struct {
	objc.Class
}

type ICollectionLayoutAnchor interface {
	objc.IObject
	Edges() DirectionalRectEdge
	Offset() foundation.Point
	IsAbsoluteOffset() bool
	IsFractionalOffset() bool
}

type CollectionLayoutAnchor struct {
	objc.Object
}

func MakeCollectionLayoutAnchor(ptr unsafe.Pointer) CollectionLayoutAnchor {
	return CollectionLayoutAnchor{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutAnchorClass) LayoutAnchorWithEdges(edges DirectionalRectEdge) CollectionLayoutAnchor {
	rv := ffi.CallMethod[CollectionLayoutAnchor](cc, "layoutAnchorWithEdges:", edges)
	return rv
}

func (cc _CollectionLayoutAnchorClass) LayoutAnchorWithEdges_AbsoluteOffset(edges DirectionalRectEdge, absoluteOffset foundation.Point) CollectionLayoutAnchor {
	rv := ffi.CallMethod[CollectionLayoutAnchor](cc, "layoutAnchorWithEdges:absoluteOffset:", edges, absoluteOffset)
	return rv
}

func (cc _CollectionLayoutAnchorClass) LayoutAnchorWithEdges_FractionalOffset(edges DirectionalRectEdge, fractionalOffset foundation.Point) CollectionLayoutAnchor {
	rv := ffi.CallMethod[CollectionLayoutAnchor](cc, "layoutAnchorWithEdges:fractionalOffset:", edges, fractionalOffset)
	return rv
}

func (cc _CollectionLayoutAnchorClass) Alloc() CollectionLayoutAnchor {
	rv := ffi.CallMethod[CollectionLayoutAnchor](cc, "alloc")
	return rv
}

func (c_ CollectionLayoutAnchor) Init() CollectionLayoutAnchor {
	rv := ffi.CallMethod[CollectionLayoutAnchor](c_, "init")
	return rv
}

func (cc _CollectionLayoutAnchorClass) New() CollectionLayoutAnchor {
	rv := ffi.CallMethod[CollectionLayoutAnchor](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutAnchor() CollectionLayoutAnchor {
	return CollectionLayoutAnchorClass.New()
}

func (c_ CollectionLayoutAnchor) Edges() DirectionalRectEdge {
	rv := ffi.CallMethod[DirectionalRectEdge](c_, "edges")
	return rv
}

func (c_ CollectionLayoutAnchor) Offset() foundation.Point {
	rv := ffi.CallMethod[foundation.Point](c_, "offset")
	return rv
}

func (c_ CollectionLayoutAnchor) IsAbsoluteOffset() bool {
	rv := ffi.CallMethod[bool](c_, "isAbsoluteOffset")
	return rv
}

func (c_ CollectionLayoutAnchor) IsFractionalOffset() bool {
	rv := ffi.CallMethod[bool](c_, "isFractionalOffset")
	return rv
}

var CollectionLayoutDimensionClass = _CollectionLayoutDimensionClass{objc.GetClass("NSCollectionLayoutDimension")}

type _CollectionLayoutDimensionClass struct {
	objc.Class
}

type ICollectionLayoutDimension interface {
	objc.IObject
	Dimension() float64
	IsAbsolute() bool
	IsEstimated() bool
	IsFractionalHeight() bool
	IsFractionalWidth() bool
}

type CollectionLayoutDimension struct {
	objc.Object
}

func MakeCollectionLayoutDimension(ptr unsafe.Pointer) CollectionLayoutDimension {
	return CollectionLayoutDimension{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutDimensionClass) AbsoluteDimension(absoluteDimension float64) CollectionLayoutDimension {
	rv := ffi.CallMethod[CollectionLayoutDimension](cc, "absoluteDimension:", absoluteDimension)
	return rv
}

func (cc _CollectionLayoutDimensionClass) EstimatedDimension(estimatedDimension float64) CollectionLayoutDimension {
	rv := ffi.CallMethod[CollectionLayoutDimension](cc, "estimatedDimension:", estimatedDimension)
	return rv
}

func (cc _CollectionLayoutDimensionClass) FractionalHeightDimension(fractionalHeight float64) CollectionLayoutDimension {
	rv := ffi.CallMethod[CollectionLayoutDimension](cc, "fractionalHeightDimension:", fractionalHeight)
	return rv
}

func (cc _CollectionLayoutDimensionClass) FractionalWidthDimension(fractionalWidth float64) CollectionLayoutDimension {
	rv := ffi.CallMethod[CollectionLayoutDimension](cc, "fractionalWidthDimension:", fractionalWidth)
	return rv
}

func (cc _CollectionLayoutDimensionClass) Alloc() CollectionLayoutDimension {
	rv := ffi.CallMethod[CollectionLayoutDimension](cc, "alloc")
	return rv
}

func (c_ CollectionLayoutDimension) Init() CollectionLayoutDimension {
	rv := ffi.CallMethod[CollectionLayoutDimension](c_, "init")
	return rv
}

func (cc _CollectionLayoutDimensionClass) New() CollectionLayoutDimension {
	rv := ffi.CallMethod[CollectionLayoutDimension](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutDimension() CollectionLayoutDimension {
	return CollectionLayoutDimensionClass.New()
}

func (c_ CollectionLayoutDimension) Dimension() float64 {
	rv := ffi.CallMethod[float64](c_, "dimension")
	return rv
}

func (c_ CollectionLayoutDimension) IsAbsolute() bool {
	rv := ffi.CallMethod[bool](c_, "isAbsolute")
	return rv
}

func (c_ CollectionLayoutDimension) IsEstimated() bool {
	rv := ffi.CallMethod[bool](c_, "isEstimated")
	return rv
}

func (c_ CollectionLayoutDimension) IsFractionalHeight() bool {
	rv := ffi.CallMethod[bool](c_, "isFractionalHeight")
	return rv
}

func (c_ CollectionLayoutDimension) IsFractionalWidth() bool {
	rv := ffi.CallMethod[bool](c_, "isFractionalWidth")
	return rv
}

var CollectionLayoutGroupClass = _CollectionLayoutGroupClass{objc.GetClass("NSCollectionLayoutGroup")}

type _CollectionLayoutGroupClass struct {
	objc.Class
}

type ICollectionLayoutGroup interface {
	ICollectionLayoutItem
	VisualDescription() string
	Subitems() []CollectionLayoutItem
	SetSupplementaryItems(value []ICollectionLayoutSupplementaryItem)
	InterItemSpacing() CollectionLayoutSpacing
	SetInterItemSpacing(value ICollectionLayoutSpacing)
}

type CollectionLayoutGroup struct {
	CollectionLayoutItem
}

func MakeCollectionLayoutGroup(ptr unsafe.Pointer) CollectionLayoutGroup {
	return CollectionLayoutGroup{
		CollectionLayoutItem: MakeCollectionLayoutItem(ptr),
	}
}

func (cc _CollectionLayoutGroupClass) HorizontalGroupWithLayoutSize_Subitems(layoutSize ICollectionLayoutSize, subitems []ICollectionLayoutItem) CollectionLayoutGroup {
	rv := ffi.CallMethod[CollectionLayoutGroup](cc, "horizontalGroupWithLayoutSize:subitems:", layoutSize, subitems)
	return rv
}

func (cc _CollectionLayoutGroupClass) VerticalGroupWithLayoutSize_Subitems(layoutSize ICollectionLayoutSize, subitems []ICollectionLayoutItem) CollectionLayoutGroup {
	rv := ffi.CallMethod[CollectionLayoutGroup](cc, "verticalGroupWithLayoutSize:subitems:", layoutSize, subitems)
	return rv
}

func (cc _CollectionLayoutGroupClass) HorizontalGroupWithLayoutSize_Subitem_Count(layoutSize ICollectionLayoutSize, subitem ICollectionLayoutItem, count int) CollectionLayoutGroup {
	rv := ffi.CallMethod[CollectionLayoutGroup](cc, "horizontalGroupWithLayoutSize:subitem:count:", layoutSize, subitem, count)
	return rv
}

func (cc _CollectionLayoutGroupClass) VerticalGroupWithLayoutSize_Subitem_Count(layoutSize ICollectionLayoutSize, subitem ICollectionLayoutItem, count int) CollectionLayoutGroup {
	rv := ffi.CallMethod[CollectionLayoutGroup](cc, "verticalGroupWithLayoutSize:subitem:count:", layoutSize, subitem, count)
	return rv
}

func (cc _CollectionLayoutGroupClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutGroup {
	rv := ffi.CallMethod[CollectionLayoutGroup](cc, "itemWithLayoutSize:", layoutSize)
	return rv
}

func (cc _CollectionLayoutGroupClass) ItemWithLayoutSize_SupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutGroup {
	rv := ffi.CallMethod[CollectionLayoutGroup](cc, "itemWithLayoutSize:supplementaryItems:", layoutSize, supplementaryItems)
	return rv
}

func (cc _CollectionLayoutGroupClass) Alloc() CollectionLayoutGroup {
	rv := ffi.CallMethod[CollectionLayoutGroup](cc, "alloc")
	return rv
}

func (c_ CollectionLayoutGroup) Init() CollectionLayoutGroup {
	rv := ffi.CallMethod[CollectionLayoutGroup](c_, "init")
	return rv
}

func (cc _CollectionLayoutGroupClass) New() CollectionLayoutGroup {
	rv := ffi.CallMethod[CollectionLayoutGroup](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutGroup() CollectionLayoutGroup {
	return CollectionLayoutGroupClass.New()
}

func (c_ CollectionLayoutGroup) VisualDescription() string {
	rv := ffi.CallMethod[string](c_, "visualDescription")
	return rv
}

func (c_ CollectionLayoutGroup) Subitems() []CollectionLayoutItem {
	rv := ffi.CallMethod[[]CollectionLayoutItem](c_, "subitems")
	return rv
}

func (c_ CollectionLayoutGroup) SetSupplementaryItems(value []ICollectionLayoutSupplementaryItem) {
	ffi.CallMethod[ffi.Void](c_, "setSupplementaryItems:", value)
}

func (c_ CollectionLayoutGroup) InterItemSpacing() CollectionLayoutSpacing {
	rv := ffi.CallMethod[CollectionLayoutSpacing](c_, "interItemSpacing")
	return rv
}

func (c_ CollectionLayoutGroup) SetInterItemSpacing(value ICollectionLayoutSpacing) {
	ffi.CallMethod[ffi.Void](c_, "setInterItemSpacing:", value)
}

var CollectionLayoutDecorationItemClass = _CollectionLayoutDecorationItemClass{objc.GetClass("NSCollectionLayoutDecorationItem")}

type _CollectionLayoutDecorationItemClass struct {
	objc.Class
}

type ICollectionLayoutDecorationItem interface {
	ICollectionLayoutItem
	ElementKind() string
	ZIndex() int
	SetZIndex(value int)
}

type CollectionLayoutDecorationItem struct {
	CollectionLayoutItem
}

func MakeCollectionLayoutDecorationItem(ptr unsafe.Pointer) CollectionLayoutDecorationItem {
	return CollectionLayoutDecorationItem{
		CollectionLayoutItem: MakeCollectionLayoutItem(ptr),
	}
}

func (cc _CollectionLayoutDecorationItemClass) BackgroundDecorationItemWithElementKind(elementKind string) CollectionLayoutDecorationItem {
	rv := ffi.CallMethod[CollectionLayoutDecorationItem](cc, "backgroundDecorationItemWithElementKind:", elementKind)
	return rv
}

func (cc _CollectionLayoutDecorationItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutDecorationItem {
	rv := ffi.CallMethod[CollectionLayoutDecorationItem](cc, "itemWithLayoutSize:", layoutSize)
	return rv
}

func (cc _CollectionLayoutDecorationItemClass) ItemWithLayoutSize_SupplementaryItems(layoutSize ICollectionLayoutSize, supplementaryItems []ICollectionLayoutSupplementaryItem) CollectionLayoutDecorationItem {
	rv := ffi.CallMethod[CollectionLayoutDecorationItem](cc, "itemWithLayoutSize:supplementaryItems:", layoutSize, supplementaryItems)
	return rv
}

func (cc _CollectionLayoutDecorationItemClass) Alloc() CollectionLayoutDecorationItem {
	rv := ffi.CallMethod[CollectionLayoutDecorationItem](cc, "alloc")
	return rv
}

func (c_ CollectionLayoutDecorationItem) Init() CollectionLayoutDecorationItem {
	rv := ffi.CallMethod[CollectionLayoutDecorationItem](c_, "init")
	return rv
}

func (cc _CollectionLayoutDecorationItemClass) New() CollectionLayoutDecorationItem {
	rv := ffi.CallMethod[CollectionLayoutDecorationItem](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutDecorationItem() CollectionLayoutDecorationItem {
	return CollectionLayoutDecorationItemClass.New()
}

func (c_ CollectionLayoutDecorationItem) ElementKind() string {
	rv := ffi.CallMethod[string](c_, "elementKind")
	return rv
}

func (c_ CollectionLayoutDecorationItem) ZIndex() int {
	rv := ffi.CallMethod[int](c_, "zIndex")
	return rv
}

func (c_ CollectionLayoutDecorationItem) SetZIndex(value int) {
	ffi.CallMethod[ffi.Void](c_, "setZIndex:", value)
}
