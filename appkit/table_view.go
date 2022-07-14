// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var TableViewClass = _TableViewClass{objc.GetClass("NSTableView")}

type _TableViewClass struct {
	objc.Class
}

type ITableView interface {
	IControl
	ReloadData()
	ReloadDataForRowIndexes_ColumnIndexes(rowIndexes foundation.IIndexSet, columnIndexes foundation.IIndexSet)
	MakeViewWithIdentifier_Owner(identifier UserInterfaceItemIdentifier, owner objc.IObject) View
	RowViewAtRow_MakeIfNecessary(row int, makeIfNecessary bool) TableRowView
	ViewAtColumn_Row_MakeIfNecessary(column int, row int, makeIfNecessary bool) View
	BeginUpdates()
	EndUpdates()
	MoveRowAtIndex_ToIndex(oldIndex int, newIndex int)
	InsertRowsAtIndexes_WithAnimation(indexes foundation.IIndexSet, animationOptions TableViewAnimationOptions)
	RemoveRowsAtIndexes_WithAnimation(indexes foundation.IIndexSet, animationOptions TableViewAnimationOptions)
	RowForView(view IView) int
	ColumnForView(view IView) int
	RegisterNib_ForIdentifier(nib INib, identifier UserInterfaceItemIdentifier)
	IndicatorImageInTableColumn(tableColumn ITableColumn) Image
	SetIndicatorImage_InTableColumn(image IImage, tableColumn ITableColumn)
	AddTableColumn(tableColumn ITableColumn)
	RemoveTableColumn(tableColumn ITableColumn)
	MoveColumn_ToColumn(oldIndex int, newIndex int)
	ColumnWithIdentifier(identifier UserInterfaceItemIdentifier) int
	TableColumnWithIdentifier(identifier UserInterfaceItemIdentifier) TableColumn
	SelectColumnIndexes_ByExtendingSelection(indexes foundation.IIndexSet, extend bool)
	DeselectColumn(column int)
	IsColumnSelected(column int) bool
	SelectRowIndexes_ByExtendingSelection(indexes foundation.IIndexSet, extend bool)
	DeselectRow(row int)
	IsRowSelected(row int) bool
	SelectAll(sender objc.IObject)
	DeselectAll(sender objc.IObject)
	EnumerateAvailableRowViewsUsingBlock(handler func(rowView TableRowView, row int))
	EditColumn_Row_WithEvent_Select(column int, row int, event IEvent, _select bool)
	DidAddRowView_ForRow(rowView ITableRowView, row int)
	DidRemoveRowView_ForRow(rowView ITableRowView, row int)
	RectOfColumn(column int) foundation.Rect
	RectOfRow(row int) foundation.Rect
	RowsInRect(rect foundation.Rect) foundation.Range
	ColumnIndexesInRect(rect foundation.Rect) foundation.IndexSet
	ColumnAtPoint(point foundation.Point) int
	RowAtPoint(point foundation.Point) int
	FrameOfCellAtColumn_Row(column int, row int) foundation.Rect
	SizeLastColumnToFit()
	NoteNumberOfRowsChanged()
	Tile()
	NoteHeightOfRowsWithIndexesChanged(indexSet foundation.IIndexSet)
	DrawRow_ClipRect(row int, clipRect foundation.Rect)
	DrawGridInClipRect(clipRect foundation.Rect)
	HighlightSelectionInClipRect(clipRect foundation.Rect)
	DrawBackgroundInClipRect(clipRect foundation.Rect)
	ScrollRowToVisible(row int)
	ScrollColumnToVisible(column int)
	DragImageForRowsWithIndexes_TableColumns_Event_Offset(dragRows foundation.IIndexSet, tableColumns []ITableColumn, dragEvent IEvent, dragImageOffset *foundation.Point) Image
	CanDragRowsWithIndexes_AtPoint(rowIndexes foundation.IIndexSet, mouseDownPoint foundation.Point) bool
	SetDraggingSourceOperationMask_ForLocal(mask DragOperation, isLocal bool)
	SetDropRow_DropOperation(row int, dropOperation TableViewDropOperation)
	HideRowsAtIndexes_WithAnimation(indexes foundation.IIndexSet, rowAnimation TableViewAnimationOptions)
	UnhideRowsAtIndexes_WithAnimation(indexes foundation.IIndexSet, rowAnimation TableViewAnimationOptions)
	// deprecated
	DragImageForRows_Event_DragImageOffset(dragRows []objc.IObject, dragEvent IEvent, dragImageOffset *foundation.Point) Image
	// deprecated
	SetAutoresizesAllColumnsToFit(flag bool)
	// deprecated
	AutoresizesAllColumnsToFit() bool
	// deprecated
	SelectColumn_ByExtendingSelection(column int, extend bool)
	// deprecated
	SelectRow_ByExtendingSelection(row int, extend bool)
	// deprecated
	TableView_WriteRows_ToPasteboard(tableView ITableView, rows []objc.IObject, pboard IPasteboard) bool
	// deprecated
	SetDrawsGrid(flag bool)
	// deprecated
	DrawsGrid() bool
	// deprecated
	SelectedColumnEnumerator() foundation.Enumerator
	// deprecated
	SelectedRowEnumerator() foundation.Enumerator
	// deprecated
	FocusedColumn() int
	// deprecated
	SetFocusedColumn(focusedColumn int)
	// deprecated
	ShouldFocusCell_AtColumn_Row(cell ICell, column int, row int) bool
	// deprecated
	PerformClickOnCellAtColumn_Row(column int, row int)
	// deprecated
	PreparedCellAtColumn_Row(column int, row int) Cell
	// deprecated
	ColumnsInRect(rect foundation.Rect) foundation.Range
	// deprecated
	TextShouldBeginEditing(textObject IText) bool
	// deprecated
	TextDidBeginEditing(notification foundation.INotification)
	// deprecated
	TextDidChange(notification foundation.INotification)
	// deprecated
	TextShouldEndEditing(textObject IText) bool
	// deprecated
	TextDidEndEditing(notification foundation.INotification)
	DataSource() TableViewDataSourceWrapper
	SetDataSource(value TableViewDataSource)
	UsesStaticContents() bool
	SetUsesStaticContents(value bool)
	RegisteredNibsByIdentifier() map[UserInterfaceItemIdentifier]Nib
	DoubleAction() objc.Selector
	SetDoubleAction(value objc.Selector)
	ClickedColumn() int
	ClickedRow() int
	AllowsColumnReordering() bool
	SetAllowsColumnReordering(value bool)
	AllowsColumnResizing() bool
	SetAllowsColumnResizing(value bool)
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	AllowsColumnSelection() bool
	SetAllowsColumnSelection(value bool)
	UsesAutomaticRowHeights() bool
	SetUsesAutomaticRowHeights(value bool)
	IntercellSpacing() foundation.Size
	SetIntercellSpacing(value foundation.Size)
	RowHeight() float64
	SetRowHeight(value float64)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	UsesAlternatingRowBackgroundColors() bool
	SetUsesAlternatingRowBackgroundColors(value bool)
	Style() TableViewStyle
	SetStyle(value TableViewStyle)
	EffectiveStyle() TableViewStyle
	SelectionHighlightStyle() TableViewSelectionHighlightStyle
	SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle)
	GridColor() Color
	SetGridColor(value IColor)
	GridStyleMask() TableViewGridLineStyle
	SetGridStyleMask(value TableViewGridLineStyle)
	EffectiveRowSizeStyle() TableViewRowSizeStyle
	RowSizeStyle() TableViewRowSizeStyle
	SetRowSizeStyle(value TableViewRowSizeStyle)
	TableColumns() []TableColumn
	SelectedColumn() int
	SelectedColumnIndexes() foundation.IndexSet
	NumberOfSelectedColumns() int
	SelectedRow() int
	SelectedRowIndexes() foundation.IndexSet
	NumberOfSelectedRows() int
	AllowsTypeSelect() bool
	SetAllowsTypeSelect(value bool)
	NumberOfColumns() int
	NumberOfRows() int
	FloatsGroupRows() bool
	SetFloatsGroupRows(value bool)
	EditedColumn() int
	EditedRow() int
	HeaderView() TableHeaderView
	SetHeaderView(value ITableHeaderView)
	CornerView() View
	SetCornerView(value IView)
	ColumnAutoresizingStyle() TableViewColumnAutoresizingStyle
	SetColumnAutoresizingStyle(value TableViewColumnAutoresizingStyle)
	AutosaveTableColumns() bool
	SetAutosaveTableColumns(value bool)
	AutosaveName() TableViewAutosaveName
	SetAutosaveName(value TableViewAutosaveName)
	Delegate() TableViewDelegateWrapper
	SetDelegate(value TableViewDelegate)
	HighlightedTableColumn() TableColumn
	SetHighlightedTableColumn(value ITableColumn)
	VerticalMotionCanBeginDrag() bool
	SetVerticalMotionCanBeginDrag(value bool)
	DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle
	SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle)
	SortDescriptors() []foundation.SortDescriptor
	SetSortDescriptors(value []foundation.ISortDescriptor)
	RowActionsVisible() bool
	SetRowActionsVisible(value bool)
	HiddenRowIndexes() foundation.IndexSet
}

type TableView struct {
	Control
}

func MakeTableView(ptr unsafe.Pointer) TableView {
	return TableView{
		Control: MakeControl(ptr),
	}
}

func (t_ TableView) InitWithFrame(frameRect foundation.Rect) TableView {
	rv := ffi.CallMethod[TableView](t_, "initWithFrame:", frameRect)
	rv.Autorelease()
	return rv
}

func (t_ TableView) Init() TableView {
	rv := ffi.CallMethod[TableView](t_, "init")
	rv.Autorelease()
	return rv
}

func (tc _TableViewClass) Alloc() TableView {
	rv := ffi.CallMethod[TableView](tc, "alloc")
	return rv
}

func (tc _TableViewClass) New() TableView {
	rv := ffi.CallMethod[TableView](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTableView() TableView {
	return TableViewClass.New()
}

func (t_ TableView) ReloadData() {
	ffi.CallMethod[ffi.Void](t_, "reloadData")
}

func (t_ TableView) ReloadDataForRowIndexes_ColumnIndexes(rowIndexes foundation.IIndexSet, columnIndexes foundation.IIndexSet) {
	ffi.CallMethod[ffi.Void](t_, "reloadDataForRowIndexes:columnIndexes:", rowIndexes, columnIndexes)
}

func (t_ TableView) MakeViewWithIdentifier_Owner(identifier UserInterfaceItemIdentifier, owner objc.IObject) View {
	rv := ffi.CallMethod[View](t_, "makeViewWithIdentifier:owner:", identifier, owner)
	return rv
}

func (t_ TableView) RowViewAtRow_MakeIfNecessary(row int, makeIfNecessary bool) TableRowView {
	rv := ffi.CallMethod[TableRowView](t_, "rowViewAtRow:makeIfNecessary:", row, makeIfNecessary)
	return rv
}

func (t_ TableView) ViewAtColumn_Row_MakeIfNecessary(column int, row int, makeIfNecessary bool) View {
	rv := ffi.CallMethod[View](t_, "viewAtColumn:row:makeIfNecessary:", column, row, makeIfNecessary)
	return rv
}

func (t_ TableView) BeginUpdates() {
	ffi.CallMethod[ffi.Void](t_, "beginUpdates")
}

func (t_ TableView) EndUpdates() {
	ffi.CallMethod[ffi.Void](t_, "endUpdates")
}

func (t_ TableView) MoveRowAtIndex_ToIndex(oldIndex int, newIndex int) {
	ffi.CallMethod[ffi.Void](t_, "moveRowAtIndex:toIndex:", oldIndex, newIndex)
}

func (t_ TableView) InsertRowsAtIndexes_WithAnimation(indexes foundation.IIndexSet, animationOptions TableViewAnimationOptions) {
	ffi.CallMethod[ffi.Void](t_, "insertRowsAtIndexes:withAnimation:", indexes, animationOptions)
}

func (t_ TableView) RemoveRowsAtIndexes_WithAnimation(indexes foundation.IIndexSet, animationOptions TableViewAnimationOptions) {
	ffi.CallMethod[ffi.Void](t_, "removeRowsAtIndexes:withAnimation:", indexes, animationOptions)
}

func (t_ TableView) RowForView(view IView) int {
	rv := ffi.CallMethod[int](t_, "rowForView:", view)
	return rv
}

func (t_ TableView) ColumnForView(view IView) int {
	rv := ffi.CallMethod[int](t_, "columnForView:", view)
	return rv
}

func (t_ TableView) RegisterNib_ForIdentifier(nib INib, identifier UserInterfaceItemIdentifier) {
	ffi.CallMethod[ffi.Void](t_, "registerNib:forIdentifier:", nib, identifier)
}

func (t_ TableView) IndicatorImageInTableColumn(tableColumn ITableColumn) Image {
	rv := ffi.CallMethod[Image](t_, "indicatorImageInTableColumn:", tableColumn)
	return rv
}

func (t_ TableView) SetIndicatorImage_InTableColumn(image IImage, tableColumn ITableColumn) {
	ffi.CallMethod[ffi.Void](t_, "setIndicatorImage:inTableColumn:", image, tableColumn)
}

func (t_ TableView) AddTableColumn(tableColumn ITableColumn) {
	ffi.CallMethod[ffi.Void](t_, "addTableColumn:", tableColumn)
}

func (t_ TableView) RemoveTableColumn(tableColumn ITableColumn) {
	ffi.CallMethod[ffi.Void](t_, "removeTableColumn:", tableColumn)
}

func (t_ TableView) MoveColumn_ToColumn(oldIndex int, newIndex int) {
	ffi.CallMethod[ffi.Void](t_, "moveColumn:toColumn:", oldIndex, newIndex)
}

func (t_ TableView) ColumnWithIdentifier(identifier UserInterfaceItemIdentifier) int {
	rv := ffi.CallMethod[int](t_, "columnWithIdentifier:", identifier)
	return rv
}

func (t_ TableView) TableColumnWithIdentifier(identifier UserInterfaceItemIdentifier) TableColumn {
	rv := ffi.CallMethod[TableColumn](t_, "tableColumnWithIdentifier:", identifier)
	return rv
}

func (t_ TableView) SelectColumnIndexes_ByExtendingSelection(indexes foundation.IIndexSet, extend bool) {
	ffi.CallMethod[ffi.Void](t_, "selectColumnIndexes:byExtendingSelection:", indexes, extend)
}

func (t_ TableView) DeselectColumn(column int) {
	ffi.CallMethod[ffi.Void](t_, "deselectColumn:", column)
}

func (t_ TableView) IsColumnSelected(column int) bool {
	rv := ffi.CallMethod[bool](t_, "isColumnSelected:", column)
	return rv
}

func (t_ TableView) SelectRowIndexes_ByExtendingSelection(indexes foundation.IIndexSet, extend bool) {
	ffi.CallMethod[ffi.Void](t_, "selectRowIndexes:byExtendingSelection:", indexes, extend)
}

func (t_ TableView) DeselectRow(row int) {
	ffi.CallMethod[ffi.Void](t_, "deselectRow:", row)
}

func (t_ TableView) IsRowSelected(row int) bool {
	rv := ffi.CallMethod[bool](t_, "isRowSelected:", row)
	return rv
}

func (t_ TableView) SelectAll(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "selectAll:", sender)
}

func (t_ TableView) DeselectAll(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "deselectAll:", sender)
}

func (t_ TableView) EnumerateAvailableRowViewsUsingBlock(handler func(rowView TableRowView, row int)) {
	ffi.CallMethod[ffi.Void](t_, "enumerateAvailableRowViewsUsingBlock:", handler)
}

func (t_ TableView) EditColumn_Row_WithEvent_Select(column int, row int, event IEvent, _select bool) {
	ffi.CallMethod[ffi.Void](t_, "editColumn:row:withEvent:select:", column, row, event, _select)
}

func (t_ TableView) DidAddRowView_ForRow(rowView ITableRowView, row int) {
	ffi.CallMethod[ffi.Void](t_, "didAddRowView:forRow:", rowView, row)
}

func (t_ TableView) DidRemoveRowView_ForRow(rowView ITableRowView, row int) {
	ffi.CallMethod[ffi.Void](t_, "didRemoveRowView:forRow:", rowView, row)
}

func (t_ TableView) RectOfColumn(column int) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](t_, "rectOfColumn:", column)
	return rv
}

func (t_ TableView) RectOfRow(row int) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](t_, "rectOfRow:", row)
	return rv
}

func (t_ TableView) RowsInRect(rect foundation.Rect) foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "rowsInRect:", rect)
	return rv
}

func (t_ TableView) ColumnIndexesInRect(rect foundation.Rect) foundation.IndexSet {
	rv := ffi.CallMethod[foundation.IndexSet](t_, "columnIndexesInRect:", rect)
	return rv
}

func (t_ TableView) ColumnAtPoint(point foundation.Point) int {
	rv := ffi.CallMethod[int](t_, "columnAtPoint:", point)
	return rv
}

func (t_ TableView) RowAtPoint(point foundation.Point) int {
	rv := ffi.CallMethod[int](t_, "rowAtPoint:", point)
	return rv
}

func (t_ TableView) FrameOfCellAtColumn_Row(column int, row int) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](t_, "frameOfCellAtColumn:row:", column, row)
	return rv
}

func (t_ TableView) SizeLastColumnToFit() {
	ffi.CallMethod[ffi.Void](t_, "sizeLastColumnToFit")
}

func (t_ TableView) NoteNumberOfRowsChanged() {
	ffi.CallMethod[ffi.Void](t_, "noteNumberOfRowsChanged")
}

func (t_ TableView) Tile() {
	ffi.CallMethod[ffi.Void](t_, "tile")
}

func (t_ TableView) NoteHeightOfRowsWithIndexesChanged(indexSet foundation.IIndexSet) {
	ffi.CallMethod[ffi.Void](t_, "noteHeightOfRowsWithIndexesChanged:", indexSet)
}

func (t_ TableView) DrawRow_ClipRect(row int, clipRect foundation.Rect) {
	ffi.CallMethod[ffi.Void](t_, "drawRow:clipRect:", row, clipRect)
}

func (t_ TableView) DrawGridInClipRect(clipRect foundation.Rect) {
	ffi.CallMethod[ffi.Void](t_, "drawGridInClipRect:", clipRect)
}

func (t_ TableView) HighlightSelectionInClipRect(clipRect foundation.Rect) {
	ffi.CallMethod[ffi.Void](t_, "highlightSelectionInClipRect:", clipRect)
}

func (t_ TableView) DrawBackgroundInClipRect(clipRect foundation.Rect) {
	ffi.CallMethod[ffi.Void](t_, "drawBackgroundInClipRect:", clipRect)
}

func (t_ TableView) ScrollRowToVisible(row int) {
	ffi.CallMethod[ffi.Void](t_, "scrollRowToVisible:", row)
}

func (t_ TableView) ScrollColumnToVisible(column int) {
	ffi.CallMethod[ffi.Void](t_, "scrollColumnToVisible:", column)
}

func (t_ TableView) DragImageForRowsWithIndexes_TableColumns_Event_Offset(dragRows foundation.IIndexSet, tableColumns []ITableColumn, dragEvent IEvent, dragImageOffset *foundation.Point) Image {
	rv := ffi.CallMethod[Image](t_, "dragImageForRowsWithIndexes:tableColumns:event:offset:", dragRows, tableColumns, dragEvent, dragImageOffset)
	return rv
}

func (t_ TableView) CanDragRowsWithIndexes_AtPoint(rowIndexes foundation.IIndexSet, mouseDownPoint foundation.Point) bool {
	rv := ffi.CallMethod[bool](t_, "canDragRowsWithIndexes:atPoint:", rowIndexes, mouseDownPoint)
	return rv
}

func (t_ TableView) SetDraggingSourceOperationMask_ForLocal(mask DragOperation, isLocal bool) {
	ffi.CallMethod[ffi.Void](t_, "setDraggingSourceOperationMask:forLocal:", mask, isLocal)
}

func (t_ TableView) SetDropRow_DropOperation(row int, dropOperation TableViewDropOperation) {
	ffi.CallMethod[ffi.Void](t_, "setDropRow:dropOperation:", row, dropOperation)
}

func (t_ TableView) HideRowsAtIndexes_WithAnimation(indexes foundation.IIndexSet, rowAnimation TableViewAnimationOptions) {
	ffi.CallMethod[ffi.Void](t_, "hideRowsAtIndexes:withAnimation:", indexes, rowAnimation)
}

func (t_ TableView) UnhideRowsAtIndexes_WithAnimation(indexes foundation.IIndexSet, rowAnimation TableViewAnimationOptions) {
	ffi.CallMethod[ffi.Void](t_, "unhideRowsAtIndexes:withAnimation:", indexes, rowAnimation)
}

// deprecated
func (t_ TableView) DragImageForRows_Event_DragImageOffset(dragRows []objc.IObject, dragEvent IEvent, dragImageOffset *foundation.Point) Image {
	rv := ffi.CallMethod[Image](t_, "dragImageForRows:event:dragImageOffset:", dragRows, dragEvent, dragImageOffset)
	return rv
}

// deprecated
func (t_ TableView) SetAutoresizesAllColumnsToFit(flag bool) {
	ffi.CallMethod[ffi.Void](t_, "setAutoresizesAllColumnsToFit:", flag)
}

// deprecated
func (t_ TableView) AutoresizesAllColumnsToFit() bool {
	rv := ffi.CallMethod[bool](t_, "autoresizesAllColumnsToFit")
	return rv
}

// deprecated
func (t_ TableView) SelectColumn_ByExtendingSelection(column int, extend bool) {
	ffi.CallMethod[ffi.Void](t_, "selectColumn:byExtendingSelection:", column, extend)
}

// deprecated
func (t_ TableView) SelectRow_ByExtendingSelection(row int, extend bool) {
	ffi.CallMethod[ffi.Void](t_, "selectRow:byExtendingSelection:", row, extend)
}

// deprecated
func (t_ TableView) TableView_WriteRows_ToPasteboard(tableView ITableView, rows []objc.IObject, pboard IPasteboard) bool {
	rv := ffi.CallMethod[bool](t_, "tableView:writeRows:toPasteboard:", tableView, rows, pboard)
	return rv
}

// deprecated
func (t_ TableView) SetDrawsGrid(flag bool) {
	ffi.CallMethod[ffi.Void](t_, "setDrawsGrid:", flag)
}

// deprecated
func (t_ TableView) DrawsGrid() bool {
	rv := ffi.CallMethod[bool](t_, "drawsGrid")
	return rv
}

// deprecated
func (t_ TableView) SelectedColumnEnumerator() foundation.Enumerator {
	rv := ffi.CallMethod[foundation.Enumerator](t_, "selectedColumnEnumerator")
	return rv
}

// deprecated
func (t_ TableView) SelectedRowEnumerator() foundation.Enumerator {
	rv := ffi.CallMethod[foundation.Enumerator](t_, "selectedRowEnumerator")
	return rv
}

// deprecated
func (t_ TableView) FocusedColumn() int {
	rv := ffi.CallMethod[int](t_, "focusedColumn")
	return rv
}

// deprecated
func (t_ TableView) SetFocusedColumn(focusedColumn int) {
	ffi.CallMethod[ffi.Void](t_, "setFocusedColumn:", focusedColumn)
}

// deprecated
func (t_ TableView) ShouldFocusCell_AtColumn_Row(cell ICell, column int, row int) bool {
	rv := ffi.CallMethod[bool](t_, "shouldFocusCell:atColumn:row:", cell, column, row)
	return rv
}

// deprecated
func (t_ TableView) PerformClickOnCellAtColumn_Row(column int, row int) {
	ffi.CallMethod[ffi.Void](t_, "performClickOnCellAtColumn:row:", column, row)
}

// deprecated
func (t_ TableView) PreparedCellAtColumn_Row(column int, row int) Cell {
	rv := ffi.CallMethod[Cell](t_, "preparedCellAtColumn:row:", column, row)
	return rv
}

// deprecated
func (t_ TableView) ColumnsInRect(rect foundation.Rect) foundation.Range {
	rv := ffi.CallMethod[foundation.Range](t_, "columnsInRect:", rect)
	return rv
}

// deprecated
func (t_ TableView) TextShouldBeginEditing(textObject IText) bool {
	rv := ffi.CallMethod[bool](t_, "textShouldBeginEditing:", textObject)
	return rv
}

// deprecated
func (t_ TableView) TextDidBeginEditing(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](t_, "textDidBeginEditing:", notification)
}

// deprecated
func (t_ TableView) TextDidChange(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](t_, "textDidChange:", notification)
}

// deprecated
func (t_ TableView) TextShouldEndEditing(textObject IText) bool {
	rv := ffi.CallMethod[bool](t_, "textShouldEndEditing:", textObject)
	return rv
}

// deprecated
func (t_ TableView) TextDidEndEditing(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](t_, "textDidEndEditing:", notification)
}

func (t_ TableView) DataSource() TableViewDataSourceWrapper {
	rv := ffi.CallMethod[TableViewDataSourceWrapper](t_, "dataSource")
	return rv
}

func (t_ TableView) SetDataSource(value TableViewDataSource) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDataSource"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](t_, "setDataSource:", po)
}

func (t_ TableView) UsesStaticContents() bool {
	rv := ffi.CallMethod[bool](t_, "usesStaticContents")
	return rv
}

func (t_ TableView) SetUsesStaticContents(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setUsesStaticContents:", value)
}

func (t_ TableView) RegisteredNibsByIdentifier() map[UserInterfaceItemIdentifier]Nib {
	rv := ffi.CallMethod[map[UserInterfaceItemIdentifier]Nib](t_, "registeredNibsByIdentifier")
	return rv
}

func (t_ TableView) DoubleAction() objc.Selector {
	rv := ffi.CallMethod[objc.Selector](t_, "doubleAction")
	return rv
}

func (t_ TableView) SetDoubleAction(value objc.Selector) {
	ffi.CallMethod[ffi.Void](t_, "setDoubleAction:", value)
}

func (t_ TableView) ClickedColumn() int {
	rv := ffi.CallMethod[int](t_, "clickedColumn")
	return rv
}

func (t_ TableView) ClickedRow() int {
	rv := ffi.CallMethod[int](t_, "clickedRow")
	return rv
}

func (t_ TableView) AllowsColumnReordering() bool {
	rv := ffi.CallMethod[bool](t_, "allowsColumnReordering")
	return rv
}

func (t_ TableView) SetAllowsColumnReordering(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAllowsColumnReordering:", value)
}

func (t_ TableView) AllowsColumnResizing() bool {
	rv := ffi.CallMethod[bool](t_, "allowsColumnResizing")
	return rv
}

func (t_ TableView) SetAllowsColumnResizing(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAllowsColumnResizing:", value)
}

func (t_ TableView) AllowsMultipleSelection() bool {
	rv := ffi.CallMethod[bool](t_, "allowsMultipleSelection")
	return rv
}

func (t_ TableView) SetAllowsMultipleSelection(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAllowsMultipleSelection:", value)
}

func (t_ TableView) AllowsEmptySelection() bool {
	rv := ffi.CallMethod[bool](t_, "allowsEmptySelection")
	return rv
}

func (t_ TableView) SetAllowsEmptySelection(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAllowsEmptySelection:", value)
}

func (t_ TableView) AllowsColumnSelection() bool {
	rv := ffi.CallMethod[bool](t_, "allowsColumnSelection")
	return rv
}

func (t_ TableView) SetAllowsColumnSelection(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAllowsColumnSelection:", value)
}

func (t_ TableView) UsesAutomaticRowHeights() bool {
	rv := ffi.CallMethod[bool](t_, "usesAutomaticRowHeights")
	return rv
}

func (t_ TableView) SetUsesAutomaticRowHeights(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setUsesAutomaticRowHeights:", value)
}

func (t_ TableView) IntercellSpacing() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](t_, "intercellSpacing")
	return rv
}

func (t_ TableView) SetIntercellSpacing(value foundation.Size) {
	ffi.CallMethod[ffi.Void](t_, "setIntercellSpacing:", value)
}

func (t_ TableView) RowHeight() float64 {
	rv := ffi.CallMethod[float64](t_, "rowHeight")
	return rv
}

func (t_ TableView) SetRowHeight(value float64) {
	ffi.CallMethod[ffi.Void](t_, "setRowHeight:", value)
}

func (t_ TableView) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](t_, "backgroundColor")
	return rv
}

func (t_ TableView) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](t_, "setBackgroundColor:", value)
}

func (t_ TableView) UsesAlternatingRowBackgroundColors() bool {
	rv := ffi.CallMethod[bool](t_, "usesAlternatingRowBackgroundColors")
	return rv
}

func (t_ TableView) SetUsesAlternatingRowBackgroundColors(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setUsesAlternatingRowBackgroundColors:", value)
}

func (t_ TableView) Style() TableViewStyle {
	rv := ffi.CallMethod[TableViewStyle](t_, "style")
	return rv
}

func (t_ TableView) SetStyle(value TableViewStyle) {
	ffi.CallMethod[ffi.Void](t_, "setStyle:", value)
}

func (t_ TableView) EffectiveStyle() TableViewStyle {
	rv := ffi.CallMethod[TableViewStyle](t_, "effectiveStyle")
	return rv
}

func (t_ TableView) SelectionHighlightStyle() TableViewSelectionHighlightStyle {
	rv := ffi.CallMethod[TableViewSelectionHighlightStyle](t_, "selectionHighlightStyle")
	return rv
}

func (t_ TableView) SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle) {
	ffi.CallMethod[ffi.Void](t_, "setSelectionHighlightStyle:", value)
}

func (t_ TableView) GridColor() Color {
	rv := ffi.CallMethod[Color](t_, "gridColor")
	return rv
}

func (t_ TableView) SetGridColor(value IColor) {
	ffi.CallMethod[ffi.Void](t_, "setGridColor:", value)
}

func (t_ TableView) GridStyleMask() TableViewGridLineStyle {
	rv := ffi.CallMethod[TableViewGridLineStyle](t_, "gridStyleMask")
	return rv
}

func (t_ TableView) SetGridStyleMask(value TableViewGridLineStyle) {
	ffi.CallMethod[ffi.Void](t_, "setGridStyleMask:", value)
}

func (t_ TableView) EffectiveRowSizeStyle() TableViewRowSizeStyle {
	rv := ffi.CallMethod[TableViewRowSizeStyle](t_, "effectiveRowSizeStyle")
	return rv
}

func (t_ TableView) RowSizeStyle() TableViewRowSizeStyle {
	rv := ffi.CallMethod[TableViewRowSizeStyle](t_, "rowSizeStyle")
	return rv
}

func (t_ TableView) SetRowSizeStyle(value TableViewRowSizeStyle) {
	ffi.CallMethod[ffi.Void](t_, "setRowSizeStyle:", value)
}

func (t_ TableView) TableColumns() []TableColumn {
	rv := ffi.CallMethod[[]TableColumn](t_, "tableColumns")
	return rv
}

func (t_ TableView) SelectedColumn() int {
	rv := ffi.CallMethod[int](t_, "selectedColumn")
	return rv
}

func (t_ TableView) SelectedColumnIndexes() foundation.IndexSet {
	rv := ffi.CallMethod[foundation.IndexSet](t_, "selectedColumnIndexes")
	return rv
}

func (t_ TableView) NumberOfSelectedColumns() int {
	rv := ffi.CallMethod[int](t_, "numberOfSelectedColumns")
	return rv
}

func (t_ TableView) SelectedRow() int {
	rv := ffi.CallMethod[int](t_, "selectedRow")
	return rv
}

func (t_ TableView) SelectedRowIndexes() foundation.IndexSet {
	rv := ffi.CallMethod[foundation.IndexSet](t_, "selectedRowIndexes")
	return rv
}

func (t_ TableView) NumberOfSelectedRows() int {
	rv := ffi.CallMethod[int](t_, "numberOfSelectedRows")
	return rv
}

func (t_ TableView) AllowsTypeSelect() bool {
	rv := ffi.CallMethod[bool](t_, "allowsTypeSelect")
	return rv
}

func (t_ TableView) SetAllowsTypeSelect(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAllowsTypeSelect:", value)
}

func (t_ TableView) NumberOfColumns() int {
	rv := ffi.CallMethod[int](t_, "numberOfColumns")
	return rv
}

func (t_ TableView) NumberOfRows() int {
	rv := ffi.CallMethod[int](t_, "numberOfRows")
	return rv
}

func (t_ TableView) FloatsGroupRows() bool {
	rv := ffi.CallMethod[bool](t_, "floatsGroupRows")
	return rv
}

func (t_ TableView) SetFloatsGroupRows(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setFloatsGroupRows:", value)
}

func (t_ TableView) EditedColumn() int {
	rv := ffi.CallMethod[int](t_, "editedColumn")
	return rv
}

func (t_ TableView) EditedRow() int {
	rv := ffi.CallMethod[int](t_, "editedRow")
	return rv
}

func (t_ TableView) HeaderView() TableHeaderView {
	rv := ffi.CallMethod[TableHeaderView](t_, "headerView")
	return rv
}

func (t_ TableView) SetHeaderView(value ITableHeaderView) {
	ffi.CallMethod[ffi.Void](t_, "setHeaderView:", value)
}

func (t_ TableView) CornerView() View {
	rv := ffi.CallMethod[View](t_, "cornerView")
	return rv
}

func (t_ TableView) SetCornerView(value IView) {
	ffi.CallMethod[ffi.Void](t_, "setCornerView:", value)
}

func (t_ TableView) ColumnAutoresizingStyle() TableViewColumnAutoresizingStyle {
	rv := ffi.CallMethod[TableViewColumnAutoresizingStyle](t_, "columnAutoresizingStyle")
	return rv
}

func (t_ TableView) SetColumnAutoresizingStyle(value TableViewColumnAutoresizingStyle) {
	ffi.CallMethod[ffi.Void](t_, "setColumnAutoresizingStyle:", value)
}

func (t_ TableView) AutosaveTableColumns() bool {
	rv := ffi.CallMethod[bool](t_, "autosaveTableColumns")
	return rv
}

func (t_ TableView) SetAutosaveTableColumns(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAutosaveTableColumns:", value)
}

func (t_ TableView) AutosaveName() TableViewAutosaveName {
	rv := ffi.CallMethod[TableViewAutosaveName](t_, "autosaveName")
	return rv
}

func (t_ TableView) SetAutosaveName(value TableViewAutosaveName) {
	ffi.CallMethod[ffi.Void](t_, "setAutosaveName:", value)
}

func (t_ TableView) Delegate() TableViewDelegateWrapper {
	rv := ffi.CallMethod[TableViewDelegateWrapper](t_, "delegate")
	return rv
}

func (t_ TableView) SetDelegate(value TableViewDelegate) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](t_, "setDelegate:", po)
}

func (t_ TableView) HighlightedTableColumn() TableColumn {
	rv := ffi.CallMethod[TableColumn](t_, "highlightedTableColumn")
	return rv
}

func (t_ TableView) SetHighlightedTableColumn(value ITableColumn) {
	ffi.CallMethod[ffi.Void](t_, "setHighlightedTableColumn:", value)
}

func (t_ TableView) VerticalMotionCanBeginDrag() bool {
	rv := ffi.CallMethod[bool](t_, "verticalMotionCanBeginDrag")
	return rv
}

func (t_ TableView) SetVerticalMotionCanBeginDrag(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setVerticalMotionCanBeginDrag:", value)
}

func (t_ TableView) DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle {
	rv := ffi.CallMethod[TableViewDraggingDestinationFeedbackStyle](t_, "draggingDestinationFeedbackStyle")
	return rv
}

func (t_ TableView) SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle) {
	ffi.CallMethod[ffi.Void](t_, "setDraggingDestinationFeedbackStyle:", value)
}

func (t_ TableView) SortDescriptors() []foundation.SortDescriptor {
	rv := ffi.CallMethod[[]foundation.SortDescriptor](t_, "sortDescriptors")
	return rv
}

func (t_ TableView) SetSortDescriptors(value []foundation.ISortDescriptor) {
	ffi.CallMethod[ffi.Void](t_, "setSortDescriptors:", value)
}

func (t_ TableView) RowActionsVisible() bool {
	rv := ffi.CallMethod[bool](t_, "rowActionsVisible")
	return rv
}

func (t_ TableView) SetRowActionsVisible(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setRowActionsVisible:", value)
}

func (t_ TableView) HiddenRowIndexes() foundation.IndexSet {
	rv := ffi.CallMethod[foundation.IndexSet](t_, "hiddenRowIndexes")
	return rv
}

type TableViewDelegate interface {
	ControlTextEditingDelegate
	ImplementsTableView_ViewForTableColumn_Row() bool
	// optional
	TableView_ViewForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) IView
	ImplementsTableView_RowViewForRow() bool
	// optional
	TableView_RowViewForRow(tableView TableView, row int) ITableRowView
	ImplementsTableView_DidAddRowView_ForRow() bool
	// optional
	TableView_DidAddRowView_ForRow(tableView TableView, rowView TableRowView, row int)
	ImplementsTableView_DidRemoveRowView_ForRow() bool
	// optional
	TableView_DidRemoveRowView_ForRow(tableView TableView, rowView TableRowView, row int)
	ImplementsTableView_IsGroupRow() bool
	// optional
	TableView_IsGroupRow(tableView TableView, row int) bool
	ImplementsTableView_WillDisplayCell_ForTableColumn_Row() bool
	// optional
	TableView_WillDisplayCell_ForTableColumn_Row(tableView TableView, cell objc.Object, tableColumn TableColumn, row int)
	ImplementsTableView_DataCellForTableColumn_Row() bool
	// optional
	TableView_DataCellForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) ICell
	ImplementsTableView_ShouldShowCellExpansionForTableColumn_Row() bool
	// optional
	TableView_ShouldShowCellExpansionForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) bool
	ImplementsTableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation() bool
	// optional
	TableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation(tableView TableView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, row int, mouseLocation foundation.Point) string
	ImplementsTableView_ShouldEditTableColumn_Row() bool
	// optional
	TableView_ShouldEditTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) bool
	ImplementsTableView_HeightOfRow() bool
	// optional
	TableView_HeightOfRow(tableView TableView, row int) float64
	ImplementsTableView_SizeToFitWidthOfColumn() bool
	// optional
	TableView_SizeToFitWidthOfColumn(tableView TableView, column int) float64
	ImplementsSelectionShouldChangeInTableView() bool
	// optional
	SelectionShouldChangeInTableView(tableView TableView) bool
	ImplementsTableView_ShouldSelectRow() bool
	// optional
	TableView_ShouldSelectRow(tableView TableView, row int) bool
	ImplementsTableView_SelectionIndexesForProposedSelection() bool
	// optional
	TableView_SelectionIndexesForProposedSelection(tableView TableView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet
	ImplementsTableView_ShouldSelectTableColumn() bool
	// optional
	TableView_ShouldSelectTableColumn(tableView TableView, tableColumn TableColumn) bool
	ImplementsTableViewSelectionIsChanging() bool
	// optional
	TableViewSelectionIsChanging(notification foundation.Notification)
	ImplementsTableViewSelectionDidChange() bool
	// optional
	TableViewSelectionDidChange(notification foundation.Notification)
	ImplementsTableView_ShouldTypeSelectForEvent_WithCurrentSearchString() bool
	// optional
	TableView_ShouldTypeSelectForEvent_WithCurrentSearchString(tableView TableView, event Event, searchString string) bool
	ImplementsTableView_TypeSelectStringForTableColumn_Row() bool
	// optional
	TableView_TypeSelectStringForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) string
	ImplementsTableView_NextTypeSelectMatchFromRow_ToRow_ForString() bool
	// optional
	TableView_NextTypeSelectMatchFromRow_ToRow_ForString(tableView TableView, startRow int, endRow int, searchString string) int
	ImplementsTableView_ShouldReorderColumn_ToColumn() bool
	// optional
	TableView_ShouldReorderColumn_ToColumn(tableView TableView, columnIndex int, newColumnIndex int) bool
	ImplementsTableView_DidDragTableColumn() bool
	// optional
	TableView_DidDragTableColumn(tableView TableView, tableColumn TableColumn)
	ImplementsTableViewColumnDidMove() bool
	// optional
	TableViewColumnDidMove(notification foundation.Notification)
	ImplementsTableViewColumnDidResize() bool
	// optional
	TableViewColumnDidResize(notification foundation.Notification)
	ImplementsTableView_DidClickTableColumn() bool
	// optional
	TableView_DidClickTableColumn(tableView TableView, tableColumn TableColumn)
	ImplementsTableView_MouseDownInHeaderOfTableColumn() bool
	// optional
	TableView_MouseDownInHeaderOfTableColumn(tableView TableView, tableColumn TableColumn)
	ImplementsTableView_ShouldTrackCell_ForTableColumn_Row() bool
	// optional
	TableView_ShouldTrackCell_ForTableColumn_Row(tableView TableView, cell Cell, tableColumn TableColumn, row int) bool
	ImplementsTableView_RowActionsForRow_Edge() bool
	// optional
	TableView_RowActionsForRow_Edge(tableView TableView, row int, edge TableRowActionEdge) []ITableViewRowAction
}

type TableViewDelegateImpl struct {
	ControlTextEditingDelegateImpl
	_TableView_ViewForTableColumn_Row                            func(tableView TableView, tableColumn TableColumn, row int) IView
	_TableView_RowViewForRow                                     func(tableView TableView, row int) ITableRowView
	_TableView_DidAddRowView_ForRow                              func(tableView TableView, rowView TableRowView, row int)
	_TableView_DidRemoveRowView_ForRow                           func(tableView TableView, rowView TableRowView, row int)
	_TableView_IsGroupRow                                        func(tableView TableView, row int) bool
	_TableView_WillDisplayCell_ForTableColumn_Row                func(tableView TableView, cell objc.Object, tableColumn TableColumn, row int)
	_TableView_DataCellForTableColumn_Row                        func(tableView TableView, tableColumn TableColumn, row int) ICell
	_TableView_ShouldShowCellExpansionForTableColumn_Row         func(tableView TableView, tableColumn TableColumn, row int) bool
	_TableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation func(tableView TableView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, row int, mouseLocation foundation.Point) string
	_TableView_ShouldEditTableColumn_Row                         func(tableView TableView, tableColumn TableColumn, row int) bool
	_TableView_HeightOfRow                                       func(tableView TableView, row int) float64
	_TableView_SizeToFitWidthOfColumn                            func(tableView TableView, column int) float64
	_SelectionShouldChangeInTableView                            func(tableView TableView) bool
	_TableView_ShouldSelectRow                                   func(tableView TableView, row int) bool
	_TableView_SelectionIndexesForProposedSelection              func(tableView TableView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet
	_TableView_ShouldSelectTableColumn                           func(tableView TableView, tableColumn TableColumn) bool
	_TableViewSelectionIsChanging                                func(notification foundation.Notification)
	_TableViewSelectionDidChange                                 func(notification foundation.Notification)
	_TableView_ShouldTypeSelectForEvent_WithCurrentSearchString  func(tableView TableView, event Event, searchString string) bool
	_TableView_TypeSelectStringForTableColumn_Row                func(tableView TableView, tableColumn TableColumn, row int) string
	_TableView_NextTypeSelectMatchFromRow_ToRow_ForString        func(tableView TableView, startRow int, endRow int, searchString string) int
	_TableView_ShouldReorderColumn_ToColumn                      func(tableView TableView, columnIndex int, newColumnIndex int) bool
	_TableView_DidDragTableColumn                                func(tableView TableView, tableColumn TableColumn)
	_TableViewColumnDidMove                                      func(notification foundation.Notification)
	_TableViewColumnDidResize                                    func(notification foundation.Notification)
	_TableView_DidClickTableColumn                               func(tableView TableView, tableColumn TableColumn)
	_TableView_MouseDownInHeaderOfTableColumn                    func(tableView TableView, tableColumn TableColumn)
	_TableView_ShouldTrackCell_ForTableColumn_Row                func(tableView TableView, cell Cell, tableColumn TableColumn, row int) bool
	_TableView_RowActionsForRow_Edge                             func(tableView TableView, row int, edge TableRowActionEdge) []ITableViewRowAction
}

func (di *TableViewDelegateImpl) ImplementsTableView_ViewForTableColumn_Row() bool {
	return di._TableView_ViewForTableColumn_Row != nil
}

func (di *TableViewDelegateImpl) SetTableView_ViewForTableColumn_Row(f func(tableView TableView, tableColumn TableColumn, row int) IView) {
	di._TableView_ViewForTableColumn_Row = f
}

func (di *TableViewDelegateImpl) TableView_ViewForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) IView {
	return di._TableView_ViewForTableColumn_Row(tableView, tableColumn, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_RowViewForRow() bool {
	return di._TableView_RowViewForRow != nil
}

func (di *TableViewDelegateImpl) SetTableView_RowViewForRow(f func(tableView TableView, row int) ITableRowView) {
	di._TableView_RowViewForRow = f
}

func (di *TableViewDelegateImpl) TableView_RowViewForRow(tableView TableView, row int) ITableRowView {
	return di._TableView_RowViewForRow(tableView, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_DidAddRowView_ForRow() bool {
	return di._TableView_DidAddRowView_ForRow != nil
}

func (di *TableViewDelegateImpl) SetTableView_DidAddRowView_ForRow(f func(tableView TableView, rowView TableRowView, row int)) {
	di._TableView_DidAddRowView_ForRow = f
}

func (di *TableViewDelegateImpl) TableView_DidAddRowView_ForRow(tableView TableView, rowView TableRowView, row int) {
	di._TableView_DidAddRowView_ForRow(tableView, rowView, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_DidRemoveRowView_ForRow() bool {
	return di._TableView_DidRemoveRowView_ForRow != nil
}

func (di *TableViewDelegateImpl) SetTableView_DidRemoveRowView_ForRow(f func(tableView TableView, rowView TableRowView, row int)) {
	di._TableView_DidRemoveRowView_ForRow = f
}

func (di *TableViewDelegateImpl) TableView_DidRemoveRowView_ForRow(tableView TableView, rowView TableRowView, row int) {
	di._TableView_DidRemoveRowView_ForRow(tableView, rowView, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_IsGroupRow() bool {
	return di._TableView_IsGroupRow != nil
}

func (di *TableViewDelegateImpl) SetTableView_IsGroupRow(f func(tableView TableView, row int) bool) {
	di._TableView_IsGroupRow = f
}

func (di *TableViewDelegateImpl) TableView_IsGroupRow(tableView TableView, row int) bool {
	return di._TableView_IsGroupRow(tableView, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_WillDisplayCell_ForTableColumn_Row() bool {
	return di._TableView_WillDisplayCell_ForTableColumn_Row != nil
}

func (di *TableViewDelegateImpl) SetTableView_WillDisplayCell_ForTableColumn_Row(f func(tableView TableView, cell objc.Object, tableColumn TableColumn, row int)) {
	di._TableView_WillDisplayCell_ForTableColumn_Row = f
}

func (di *TableViewDelegateImpl) TableView_WillDisplayCell_ForTableColumn_Row(tableView TableView, cell objc.Object, tableColumn TableColumn, row int) {
	di._TableView_WillDisplayCell_ForTableColumn_Row(tableView, cell, tableColumn, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_DataCellForTableColumn_Row() bool {
	return di._TableView_DataCellForTableColumn_Row != nil
}

func (di *TableViewDelegateImpl) SetTableView_DataCellForTableColumn_Row(f func(tableView TableView, tableColumn TableColumn, row int) ICell) {
	di._TableView_DataCellForTableColumn_Row = f
}

func (di *TableViewDelegateImpl) TableView_DataCellForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) ICell {
	return di._TableView_DataCellForTableColumn_Row(tableView, tableColumn, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_ShouldShowCellExpansionForTableColumn_Row() bool {
	return di._TableView_ShouldShowCellExpansionForTableColumn_Row != nil
}

func (di *TableViewDelegateImpl) SetTableView_ShouldShowCellExpansionForTableColumn_Row(f func(tableView TableView, tableColumn TableColumn, row int) bool) {
	di._TableView_ShouldShowCellExpansionForTableColumn_Row = f
}

func (di *TableViewDelegateImpl) TableView_ShouldShowCellExpansionForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) bool {
	return di._TableView_ShouldShowCellExpansionForTableColumn_Row(tableView, tableColumn, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation() bool {
	return di._TableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation != nil
}

func (di *TableViewDelegateImpl) SetTableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation(f func(tableView TableView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, row int, mouseLocation foundation.Point) string) {
	di._TableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation = f
}

func (di *TableViewDelegateImpl) TableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation(tableView TableView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, row int, mouseLocation foundation.Point) string {
	return di._TableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation(tableView, cell, rect, tableColumn, row, mouseLocation)
}
func (di *TableViewDelegateImpl) ImplementsTableView_ShouldEditTableColumn_Row() bool {
	return di._TableView_ShouldEditTableColumn_Row != nil
}

func (di *TableViewDelegateImpl) SetTableView_ShouldEditTableColumn_Row(f func(tableView TableView, tableColumn TableColumn, row int) bool) {
	di._TableView_ShouldEditTableColumn_Row = f
}

func (di *TableViewDelegateImpl) TableView_ShouldEditTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) bool {
	return di._TableView_ShouldEditTableColumn_Row(tableView, tableColumn, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_HeightOfRow() bool {
	return di._TableView_HeightOfRow != nil
}

func (di *TableViewDelegateImpl) SetTableView_HeightOfRow(f func(tableView TableView, row int) float64) {
	di._TableView_HeightOfRow = f
}

func (di *TableViewDelegateImpl) TableView_HeightOfRow(tableView TableView, row int) float64 {
	return di._TableView_HeightOfRow(tableView, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_SizeToFitWidthOfColumn() bool {
	return di._TableView_SizeToFitWidthOfColumn != nil
}

func (di *TableViewDelegateImpl) SetTableView_SizeToFitWidthOfColumn(f func(tableView TableView, column int) float64) {
	di._TableView_SizeToFitWidthOfColumn = f
}

func (di *TableViewDelegateImpl) TableView_SizeToFitWidthOfColumn(tableView TableView, column int) float64 {
	return di._TableView_SizeToFitWidthOfColumn(tableView, column)
}
func (di *TableViewDelegateImpl) ImplementsSelectionShouldChangeInTableView() bool {
	return di._SelectionShouldChangeInTableView != nil
}

func (di *TableViewDelegateImpl) SetSelectionShouldChangeInTableView(f func(tableView TableView) bool) {
	di._SelectionShouldChangeInTableView = f
}

func (di *TableViewDelegateImpl) SelectionShouldChangeInTableView(tableView TableView) bool {
	return di._SelectionShouldChangeInTableView(tableView)
}
func (di *TableViewDelegateImpl) ImplementsTableView_ShouldSelectRow() bool {
	return di._TableView_ShouldSelectRow != nil
}

func (di *TableViewDelegateImpl) SetTableView_ShouldSelectRow(f func(tableView TableView, row int) bool) {
	di._TableView_ShouldSelectRow = f
}

func (di *TableViewDelegateImpl) TableView_ShouldSelectRow(tableView TableView, row int) bool {
	return di._TableView_ShouldSelectRow(tableView, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_SelectionIndexesForProposedSelection() bool {
	return di._TableView_SelectionIndexesForProposedSelection != nil
}

func (di *TableViewDelegateImpl) SetTableView_SelectionIndexesForProposedSelection(f func(tableView TableView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet) {
	di._TableView_SelectionIndexesForProposedSelection = f
}

func (di *TableViewDelegateImpl) TableView_SelectionIndexesForProposedSelection(tableView TableView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet {
	return di._TableView_SelectionIndexesForProposedSelection(tableView, proposedSelectionIndexes)
}
func (di *TableViewDelegateImpl) ImplementsTableView_ShouldSelectTableColumn() bool {
	return di._TableView_ShouldSelectTableColumn != nil
}

func (di *TableViewDelegateImpl) SetTableView_ShouldSelectTableColumn(f func(tableView TableView, tableColumn TableColumn) bool) {
	di._TableView_ShouldSelectTableColumn = f
}

func (di *TableViewDelegateImpl) TableView_ShouldSelectTableColumn(tableView TableView, tableColumn TableColumn) bool {
	return di._TableView_ShouldSelectTableColumn(tableView, tableColumn)
}
func (di *TableViewDelegateImpl) ImplementsTableViewSelectionIsChanging() bool {
	return di._TableViewSelectionIsChanging != nil
}

func (di *TableViewDelegateImpl) SetTableViewSelectionIsChanging(f func(notification foundation.Notification)) {
	di._TableViewSelectionIsChanging = f
}

func (di *TableViewDelegateImpl) TableViewSelectionIsChanging(notification foundation.Notification) {
	di._TableViewSelectionIsChanging(notification)
}
func (di *TableViewDelegateImpl) ImplementsTableViewSelectionDidChange() bool {
	return di._TableViewSelectionDidChange != nil
}

func (di *TableViewDelegateImpl) SetTableViewSelectionDidChange(f func(notification foundation.Notification)) {
	di._TableViewSelectionDidChange = f
}

func (di *TableViewDelegateImpl) TableViewSelectionDidChange(notification foundation.Notification) {
	di._TableViewSelectionDidChange(notification)
}
func (di *TableViewDelegateImpl) ImplementsTableView_ShouldTypeSelectForEvent_WithCurrentSearchString() bool {
	return di._TableView_ShouldTypeSelectForEvent_WithCurrentSearchString != nil
}

func (di *TableViewDelegateImpl) SetTableView_ShouldTypeSelectForEvent_WithCurrentSearchString(f func(tableView TableView, event Event, searchString string) bool) {
	di._TableView_ShouldTypeSelectForEvent_WithCurrentSearchString = f
}

func (di *TableViewDelegateImpl) TableView_ShouldTypeSelectForEvent_WithCurrentSearchString(tableView TableView, event Event, searchString string) bool {
	return di._TableView_ShouldTypeSelectForEvent_WithCurrentSearchString(tableView, event, searchString)
}
func (di *TableViewDelegateImpl) ImplementsTableView_TypeSelectStringForTableColumn_Row() bool {
	return di._TableView_TypeSelectStringForTableColumn_Row != nil
}

func (di *TableViewDelegateImpl) SetTableView_TypeSelectStringForTableColumn_Row(f func(tableView TableView, tableColumn TableColumn, row int) string) {
	di._TableView_TypeSelectStringForTableColumn_Row = f
}

func (di *TableViewDelegateImpl) TableView_TypeSelectStringForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) string {
	return di._TableView_TypeSelectStringForTableColumn_Row(tableView, tableColumn, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_NextTypeSelectMatchFromRow_ToRow_ForString() bool {
	return di._TableView_NextTypeSelectMatchFromRow_ToRow_ForString != nil
}

func (di *TableViewDelegateImpl) SetTableView_NextTypeSelectMatchFromRow_ToRow_ForString(f func(tableView TableView, startRow int, endRow int, searchString string) int) {
	di._TableView_NextTypeSelectMatchFromRow_ToRow_ForString = f
}

func (di *TableViewDelegateImpl) TableView_NextTypeSelectMatchFromRow_ToRow_ForString(tableView TableView, startRow int, endRow int, searchString string) int {
	return di._TableView_NextTypeSelectMatchFromRow_ToRow_ForString(tableView, startRow, endRow, searchString)
}
func (di *TableViewDelegateImpl) ImplementsTableView_ShouldReorderColumn_ToColumn() bool {
	return di._TableView_ShouldReorderColumn_ToColumn != nil
}

func (di *TableViewDelegateImpl) SetTableView_ShouldReorderColumn_ToColumn(f func(tableView TableView, columnIndex int, newColumnIndex int) bool) {
	di._TableView_ShouldReorderColumn_ToColumn = f
}

func (di *TableViewDelegateImpl) TableView_ShouldReorderColumn_ToColumn(tableView TableView, columnIndex int, newColumnIndex int) bool {
	return di._TableView_ShouldReorderColumn_ToColumn(tableView, columnIndex, newColumnIndex)
}
func (di *TableViewDelegateImpl) ImplementsTableView_DidDragTableColumn() bool {
	return di._TableView_DidDragTableColumn != nil
}

func (di *TableViewDelegateImpl) SetTableView_DidDragTableColumn(f func(tableView TableView, tableColumn TableColumn)) {
	di._TableView_DidDragTableColumn = f
}

func (di *TableViewDelegateImpl) TableView_DidDragTableColumn(tableView TableView, tableColumn TableColumn) {
	di._TableView_DidDragTableColumn(tableView, tableColumn)
}
func (di *TableViewDelegateImpl) ImplementsTableViewColumnDidMove() bool {
	return di._TableViewColumnDidMove != nil
}

func (di *TableViewDelegateImpl) SetTableViewColumnDidMove(f func(notification foundation.Notification)) {
	di._TableViewColumnDidMove = f
}

func (di *TableViewDelegateImpl) TableViewColumnDidMove(notification foundation.Notification) {
	di._TableViewColumnDidMove(notification)
}
func (di *TableViewDelegateImpl) ImplementsTableViewColumnDidResize() bool {
	return di._TableViewColumnDidResize != nil
}

func (di *TableViewDelegateImpl) SetTableViewColumnDidResize(f func(notification foundation.Notification)) {
	di._TableViewColumnDidResize = f
}

func (di *TableViewDelegateImpl) TableViewColumnDidResize(notification foundation.Notification) {
	di._TableViewColumnDidResize(notification)
}
func (di *TableViewDelegateImpl) ImplementsTableView_DidClickTableColumn() bool {
	return di._TableView_DidClickTableColumn != nil
}

func (di *TableViewDelegateImpl) SetTableView_DidClickTableColumn(f func(tableView TableView, tableColumn TableColumn)) {
	di._TableView_DidClickTableColumn = f
}

func (di *TableViewDelegateImpl) TableView_DidClickTableColumn(tableView TableView, tableColumn TableColumn) {
	di._TableView_DidClickTableColumn(tableView, tableColumn)
}
func (di *TableViewDelegateImpl) ImplementsTableView_MouseDownInHeaderOfTableColumn() bool {
	return di._TableView_MouseDownInHeaderOfTableColumn != nil
}

func (di *TableViewDelegateImpl) SetTableView_MouseDownInHeaderOfTableColumn(f func(tableView TableView, tableColumn TableColumn)) {
	di._TableView_MouseDownInHeaderOfTableColumn = f
}

func (di *TableViewDelegateImpl) TableView_MouseDownInHeaderOfTableColumn(tableView TableView, tableColumn TableColumn) {
	di._TableView_MouseDownInHeaderOfTableColumn(tableView, tableColumn)
}
func (di *TableViewDelegateImpl) ImplementsTableView_ShouldTrackCell_ForTableColumn_Row() bool {
	return di._TableView_ShouldTrackCell_ForTableColumn_Row != nil
}

func (di *TableViewDelegateImpl) SetTableView_ShouldTrackCell_ForTableColumn_Row(f func(tableView TableView, cell Cell, tableColumn TableColumn, row int) bool) {
	di._TableView_ShouldTrackCell_ForTableColumn_Row = f
}

func (di *TableViewDelegateImpl) TableView_ShouldTrackCell_ForTableColumn_Row(tableView TableView, cell Cell, tableColumn TableColumn, row int) bool {
	return di._TableView_ShouldTrackCell_ForTableColumn_Row(tableView, cell, tableColumn, row)
}
func (di *TableViewDelegateImpl) ImplementsTableView_RowActionsForRow_Edge() bool {
	return di._TableView_RowActionsForRow_Edge != nil
}

func (di *TableViewDelegateImpl) SetTableView_RowActionsForRow_Edge(f func(tableView TableView, row int, edge TableRowActionEdge) []ITableViewRowAction) {
	di._TableView_RowActionsForRow_Edge = f
}

func (di *TableViewDelegateImpl) TableView_RowActionsForRow_Edge(tableView TableView, row int, edge TableRowActionEdge) []ITableViewRowAction {
	return di._TableView_RowActionsForRow_Edge(tableView, row, edge)
}

type TableViewDelegateWrapper struct {
	ControlTextEditingDelegateWrapper
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_ViewForTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:viewForTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableView_ViewForTableColumn_Row(tableView ITableView, tableColumn ITableColumn, row int) View {
	rv := ffi.CallMethod[View](t_, "tableView:viewForTableColumn:row:", tableView, tableColumn, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_RowViewForRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:rowViewForRow:"))
}

func (t_ TableViewDelegateWrapper) TableView_RowViewForRow(tableView ITableView, row int) TableRowView {
	rv := ffi.CallMethod[TableRowView](t_, "tableView:rowViewForRow:", tableView, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_DidAddRowView_ForRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:didAddRowView:forRow:"))
}

func (t_ TableViewDelegateWrapper) TableView_DidAddRowView_ForRow(tableView ITableView, rowView ITableRowView, row int) {
	ffi.CallMethod[ffi.Void](t_, "tableView:didAddRowView:forRow:", tableView, rowView, row)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_DidRemoveRowView_ForRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:didRemoveRowView:forRow:"))
}

func (t_ TableViewDelegateWrapper) TableView_DidRemoveRowView_ForRow(tableView ITableView, rowView ITableRowView, row int) {
	ffi.CallMethod[ffi.Void](t_, "tableView:didRemoveRowView:forRow:", tableView, rowView, row)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_IsGroupRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:isGroupRow:"))
}

func (t_ TableViewDelegateWrapper) TableView_IsGroupRow(tableView ITableView, row int) bool {
	rv := ffi.CallMethod[bool](t_, "tableView:isGroupRow:", tableView, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_WillDisplayCell_ForTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:willDisplayCell:forTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableView_WillDisplayCell_ForTableColumn_Row(tableView ITableView, cell objc.IObject, tableColumn ITableColumn, row int) {
	ffi.CallMethod[ffi.Void](t_, "tableView:willDisplayCell:forTableColumn:row:", tableView, cell, tableColumn, row)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_DataCellForTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:dataCellForTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableView_DataCellForTableColumn_Row(tableView ITableView, tableColumn ITableColumn, row int) Cell {
	rv := ffi.CallMethod[Cell](t_, "tableView:dataCellForTableColumn:row:", tableView, tableColumn, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_ShouldShowCellExpansionForTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldShowCellExpansionForTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableView_ShouldShowCellExpansionForTableColumn_Row(tableView ITableView, tableColumn ITableColumn, row int) bool {
	rv := ffi.CallMethod[bool](t_, "tableView:shouldShowCellExpansionForTableColumn:row:", tableView, tableColumn, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:toolTipForCell:rect:tableColumn:row:mouseLocation:"))
}

func (t_ TableViewDelegateWrapper) TableView_ToolTipForCell_Rect_TableColumn_Row_MouseLocation(tableView ITableView, cell ICell, rect *foundation.Rect, tableColumn ITableColumn, row int, mouseLocation foundation.Point) string {
	rv := ffi.CallMethod[string](t_, "tableView:toolTipForCell:rect:tableColumn:row:mouseLocation:", tableView, cell, rect, tableColumn, row, mouseLocation)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_ShouldEditTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldEditTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableView_ShouldEditTableColumn_Row(tableView ITableView, tableColumn ITableColumn, row int) bool {
	rv := ffi.CallMethod[bool](t_, "tableView:shouldEditTableColumn:row:", tableView, tableColumn, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_HeightOfRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:heightOfRow:"))
}

func (t_ TableViewDelegateWrapper) TableView_HeightOfRow(tableView ITableView, row int) float64 {
	rv := ffi.CallMethod[float64](t_, "tableView:heightOfRow:", tableView, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_SizeToFitWidthOfColumn() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:sizeToFitWidthOfColumn:"))
}

func (t_ TableViewDelegateWrapper) TableView_SizeToFitWidthOfColumn(tableView ITableView, column int) float64 {
	rv := ffi.CallMethod[float64](t_, "tableView:sizeToFitWidthOfColumn:", tableView, column)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsSelectionShouldChangeInTableView() bool {
	return t_.RespondsToSelector(objc.GetSelector("selectionShouldChangeInTableView:"))
}

func (t_ TableViewDelegateWrapper) SelectionShouldChangeInTableView(tableView ITableView) bool {
	rv := ffi.CallMethod[bool](t_, "selectionShouldChangeInTableView:", tableView)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_ShouldSelectRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldSelectRow:"))
}

func (t_ TableViewDelegateWrapper) TableView_ShouldSelectRow(tableView ITableView, row int) bool {
	rv := ffi.CallMethod[bool](t_, "tableView:shouldSelectRow:", tableView, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_SelectionIndexesForProposedSelection() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:selectionIndexesForProposedSelection:"))
}

func (t_ TableViewDelegateWrapper) TableView_SelectionIndexesForProposedSelection(tableView ITableView, proposedSelectionIndexes foundation.IIndexSet) foundation.IndexSet {
	rv := ffi.CallMethod[foundation.IndexSet](t_, "tableView:selectionIndexesForProposedSelection:", tableView, proposedSelectionIndexes)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_ShouldSelectTableColumn() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldSelectTableColumn:"))
}

func (t_ TableViewDelegateWrapper) TableView_ShouldSelectTableColumn(tableView ITableView, tableColumn ITableColumn) bool {
	rv := ffi.CallMethod[bool](t_, "tableView:shouldSelectTableColumn:", tableView, tableColumn)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableViewSelectionIsChanging() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableViewSelectionIsChanging:"))
}

func (t_ TableViewDelegateWrapper) TableViewSelectionIsChanging(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](t_, "tableViewSelectionIsChanging:", notification)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableViewSelectionDidChange() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableViewSelectionDidChange:"))
}

func (t_ TableViewDelegateWrapper) TableViewSelectionDidChange(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](t_, "tableViewSelectionDidChange:", notification)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_ShouldTypeSelectForEvent_WithCurrentSearchString() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldTypeSelectForEvent:withCurrentSearchString:"))
}

func (t_ TableViewDelegateWrapper) TableView_ShouldTypeSelectForEvent_WithCurrentSearchString(tableView ITableView, event IEvent, searchString string) bool {
	rv := ffi.CallMethod[bool](t_, "tableView:shouldTypeSelectForEvent:withCurrentSearchString:", tableView, event, searchString)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_TypeSelectStringForTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:typeSelectStringForTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableView_TypeSelectStringForTableColumn_Row(tableView ITableView, tableColumn ITableColumn, row int) string {
	rv := ffi.CallMethod[string](t_, "tableView:typeSelectStringForTableColumn:row:", tableView, tableColumn, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_NextTypeSelectMatchFromRow_ToRow_ForString() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:nextTypeSelectMatchFromRow:toRow:forString:"))
}

func (t_ TableViewDelegateWrapper) TableView_NextTypeSelectMatchFromRow_ToRow_ForString(tableView ITableView, startRow int, endRow int, searchString string) int {
	rv := ffi.CallMethod[int](t_, "tableView:nextTypeSelectMatchFromRow:toRow:forString:", tableView, startRow, endRow, searchString)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_ShouldReorderColumn_ToColumn() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldReorderColumn:toColumn:"))
}

func (t_ TableViewDelegateWrapper) TableView_ShouldReorderColumn_ToColumn(tableView ITableView, columnIndex int, newColumnIndex int) bool {
	rv := ffi.CallMethod[bool](t_, "tableView:shouldReorderColumn:toColumn:", tableView, columnIndex, newColumnIndex)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_DidDragTableColumn() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:didDragTableColumn:"))
}

func (t_ TableViewDelegateWrapper) TableView_DidDragTableColumn(tableView ITableView, tableColumn ITableColumn) {
	ffi.CallMethod[ffi.Void](t_, "tableView:didDragTableColumn:", tableView, tableColumn)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableViewColumnDidMove() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableViewColumnDidMove:"))
}

func (t_ TableViewDelegateWrapper) TableViewColumnDidMove(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](t_, "tableViewColumnDidMove:", notification)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableViewColumnDidResize() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableViewColumnDidResize:"))
}

func (t_ TableViewDelegateWrapper) TableViewColumnDidResize(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](t_, "tableViewColumnDidResize:", notification)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_DidClickTableColumn() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:didClickTableColumn:"))
}

func (t_ TableViewDelegateWrapper) TableView_DidClickTableColumn(tableView ITableView, tableColumn ITableColumn) {
	ffi.CallMethod[ffi.Void](t_, "tableView:didClickTableColumn:", tableView, tableColumn)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_MouseDownInHeaderOfTableColumn() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:mouseDownInHeaderOfTableColumn:"))
}

func (t_ TableViewDelegateWrapper) TableView_MouseDownInHeaderOfTableColumn(tableView ITableView, tableColumn ITableColumn) {
	ffi.CallMethod[ffi.Void](t_, "tableView:mouseDownInHeaderOfTableColumn:", tableView, tableColumn)
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_ShouldTrackCell_ForTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:shouldTrackCell:forTableColumn:row:"))
}

func (t_ TableViewDelegateWrapper) TableView_ShouldTrackCell_ForTableColumn_Row(tableView ITableView, cell ICell, tableColumn ITableColumn, row int) bool {
	rv := ffi.CallMethod[bool](t_, "tableView:shouldTrackCell:forTableColumn:row:", tableView, cell, tableColumn, row)
	return rv
}

func (t_ *TableViewDelegateWrapper) ImplementsTableView_RowActionsForRow_Edge() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:rowActionsForRow:edge:"))
}

func (t_ TableViewDelegateWrapper) TableView_RowActionsForRow_Edge(tableView ITableView, row int, edge TableRowActionEdge) []TableViewRowAction {
	rv := ffi.CallMethod[[]TableViewRowAction](t_, "tableView:rowActionsForRow:edge:", tableView, row, edge)
	return rv
}

type TableViewDataSource interface {
	ImplementsNumberOfRowsInTableView() bool
	// optional
	NumberOfRowsInTableView(tableView TableView) int
	ImplementsTableView_ObjectValueForTableColumn_Row() bool
	// optional
	TableView_ObjectValueForTableColumn_Row(tableView TableView, tableColumn TableColumn, row int) objc.IObject
	ImplementsTableView_SetObjectValue_ForTableColumn_Row() bool
	// optional
	TableView_SetObjectValue_ForTableColumn_Row(tableView TableView, object objc.Object, tableColumn TableColumn, row int)
	ImplementsTableView_PasteboardWriterForRow() bool
	// optional
	TableView_PasteboardWriterForRow(tableView TableView, row int) PasteboardWriting
	ImplementsTableView_AcceptDrop_Row_DropOperation() bool
	// optional
	TableView_AcceptDrop_Row_DropOperation(tableView TableView, info DraggingInfoWrapper, row int, dropOperation TableViewDropOperation) bool
	ImplementsTableView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes() bool
	// optional
	// deprecated
	TableView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes(tableView TableView, dropDestination foundation.URL, indexSet foundation.IndexSet) []string
	ImplementsTableView_ValidateDrop_ProposedRow_ProposedDropOperation() bool
	// optional
	TableView_ValidateDrop_ProposedRow_ProposedDropOperation(tableView TableView, info DraggingInfoWrapper, row int, dropOperation TableViewDropOperation) DragOperation
	ImplementsTableView_WriteRowsWithIndexes_ToPasteboard() bool
	// optional
	// deprecated
	TableView_WriteRowsWithIndexes_ToPasteboard(tableView TableView, rowIndexes foundation.IndexSet, pboard Pasteboard) bool
	ImplementsTableView_DraggingSession_WillBeginAtPoint_ForRowIndexes() bool
	// optional
	TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes(tableView TableView, session DraggingSession, screenPoint foundation.Point, rowIndexes foundation.IndexSet)
	ImplementsTableView_UpdateDraggingItemsForDrag() bool
	// optional
	TableView_UpdateDraggingItemsForDrag(tableView TableView, draggingInfo DraggingInfoWrapper)
	ImplementsTableView_DraggingSession_EndedAtPoint_Operation() bool
	// optional
	TableView_DraggingSession_EndedAtPoint_Operation(tableView TableView, session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	ImplementsTableView_SortDescriptorsDidChange() bool
	// optional
	TableView_SortDescriptorsDidChange(tableView TableView, oldDescriptors []foundation.SortDescriptor)
}

type TableViewDataSourceWrapper struct {
	objc.Object
}

func (t_ *TableViewDataSourceWrapper) ImplementsNumberOfRowsInTableView() bool {
	return t_.RespondsToSelector(objc.GetSelector("numberOfRowsInTableView:"))
}

func (t_ TableViewDataSourceWrapper) NumberOfRowsInTableView(tableView ITableView) int {
	rv := ffi.CallMethod[int](t_, "numberOfRowsInTableView:", tableView)
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_ObjectValueForTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:objectValueForTableColumn:row:"))
}

func (t_ TableViewDataSourceWrapper) TableView_ObjectValueForTableColumn_Row(tableView ITableView, tableColumn ITableColumn, row int) objc.Object {
	rv := ffi.CallMethod[objc.Object](t_, "tableView:objectValueForTableColumn:row:", tableView, tableColumn, row)
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_SetObjectValue_ForTableColumn_Row() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:setObjectValue:forTableColumn:row:"))
}

func (t_ TableViewDataSourceWrapper) TableView_SetObjectValue_ForTableColumn_Row(tableView ITableView, object objc.IObject, tableColumn ITableColumn, row int) {
	ffi.CallMethod[ffi.Void](t_, "tableView:setObjectValue:forTableColumn:row:", tableView, object, tableColumn, row)
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_PasteboardWriterForRow() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:pasteboardWriterForRow:"))
}

func (t_ TableViewDataSourceWrapper) TableView_PasteboardWriterForRow(tableView ITableView, row int) PasteboardWritingWrapper {
	rv := ffi.CallMethod[PasteboardWritingWrapper](t_, "tableView:pasteboardWriterForRow:", tableView, row)
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_AcceptDrop_Row_DropOperation() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:acceptDrop:row:dropOperation:"))
}

func (t_ TableViewDataSourceWrapper) TableView_AcceptDrop_Row_DropOperation(tableView ITableView, info DraggingInfo, row int, dropOperation TableViewDropOperation) bool {
	po := ffi.CreateProtocol(info)
	defer po.Release()
	rv := ffi.CallMethod[bool](t_, "tableView:acceptDrop:row:dropOperation:", tableView, po, row, dropOperation)
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:namesOfPromisedFilesDroppedAtDestination:forDraggedRowsWithIndexes:"))
}

// deprecated
func (t_ TableViewDataSourceWrapper) TableView_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes(tableView ITableView, dropDestination foundation.IURL, indexSet foundation.IIndexSet) []string {
	rv := ffi.CallMethod[[]string](t_, "tableView:namesOfPromisedFilesDroppedAtDestination:forDraggedRowsWithIndexes:", tableView, dropDestination, indexSet)
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_ValidateDrop_ProposedRow_ProposedDropOperation() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:validateDrop:proposedRow:proposedDropOperation:"))
}

func (t_ TableViewDataSourceWrapper) TableView_ValidateDrop_ProposedRow_ProposedDropOperation(tableView ITableView, info DraggingInfo, row int, dropOperation TableViewDropOperation) DragOperation {
	po := ffi.CreateProtocol(info)
	defer po.Release()
	rv := ffi.CallMethod[DragOperation](t_, "tableView:validateDrop:proposedRow:proposedDropOperation:", tableView, po, row, dropOperation)
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_WriteRowsWithIndexes_ToPasteboard() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:writeRowsWithIndexes:toPasteboard:"))
}

// deprecated
func (t_ TableViewDataSourceWrapper) TableView_WriteRowsWithIndexes_ToPasteboard(tableView ITableView, rowIndexes foundation.IIndexSet, pboard IPasteboard) bool {
	rv := ffi.CallMethod[bool](t_, "tableView:writeRowsWithIndexes:toPasteboard:", tableView, rowIndexes, pboard)
	return rv
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_DraggingSession_WillBeginAtPoint_ForRowIndexes() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:draggingSession:willBeginAtPoint:forRowIndexes:"))
}

func (t_ TableViewDataSourceWrapper) TableView_DraggingSession_WillBeginAtPoint_ForRowIndexes(tableView ITableView, session IDraggingSession, screenPoint foundation.Point, rowIndexes foundation.IIndexSet) {
	ffi.CallMethod[ffi.Void](t_, "tableView:draggingSession:willBeginAtPoint:forRowIndexes:", tableView, session, screenPoint, rowIndexes)
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_UpdateDraggingItemsForDrag() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:updateDraggingItemsForDrag:"))
}

func (t_ TableViewDataSourceWrapper) TableView_UpdateDraggingItemsForDrag(tableView ITableView, draggingInfo DraggingInfo) {
	po := ffi.CreateProtocol(draggingInfo)
	defer po.Release()
	ffi.CallMethod[ffi.Void](t_, "tableView:updateDraggingItemsForDrag:", tableView, po)
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_DraggingSession_EndedAtPoint_Operation() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:draggingSession:endedAtPoint:operation:"))
}

func (t_ TableViewDataSourceWrapper) TableView_DraggingSession_EndedAtPoint_Operation(tableView ITableView, session IDraggingSession, screenPoint foundation.Point, operation DragOperation) {
	ffi.CallMethod[ffi.Void](t_, "tableView:draggingSession:endedAtPoint:operation:", tableView, session, screenPoint, operation)
}

func (t_ *TableViewDataSourceWrapper) ImplementsTableView_SortDescriptorsDidChange() bool {
	return t_.RespondsToSelector(objc.GetSelector("tableView:sortDescriptorsDidChange:"))
}

func (t_ TableViewDataSourceWrapper) TableView_SortDescriptorsDidChange(tableView ITableView, oldDescriptors []foundation.ISortDescriptor) {
	ffi.CallMethod[ffi.Void](t_, "tableView:sortDescriptorsDidChange:", tableView, oldDescriptors)
}

var TableCellViewClass = _TableCellViewClass{objc.GetClass("NSTableCellView")}

type _TableCellViewClass struct {
	objc.Class
}

type ITableCellView interface {
	IView
	ObjectValue() objc.Object
	SetObjectValue(value objc.IObject)
	ImageView() ImageView
	SetImageView(value IImageView)
	TextField() TextField
	SetTextField(value ITextField)
	BackgroundStyle() BackgroundStyle
	SetBackgroundStyle(value BackgroundStyle)
	RowSizeStyle() TableViewRowSizeStyle
	SetRowSizeStyle(value TableViewRowSizeStyle)
	DraggingImageComponents() []DraggingImageComponent
}

type TableCellView struct {
	View
}

func MakeTableCellView(ptr unsafe.Pointer) TableCellView {
	return TableCellView{
		View: MakeView(ptr),
	}
}

func (t_ TableCellView) InitWithFrame(frameRect foundation.Rect) TableCellView {
	rv := ffi.CallMethod[TableCellView](t_, "initWithFrame:", frameRect)
	rv.Autorelease()
	return rv
}

func (t_ TableCellView) Init() TableCellView {
	rv := ffi.CallMethod[TableCellView](t_, "init")
	rv.Autorelease()
	return rv
}

func (tc _TableCellViewClass) Alloc() TableCellView {
	rv := ffi.CallMethod[TableCellView](tc, "alloc")
	return rv
}

func (tc _TableCellViewClass) New() TableCellView {
	rv := ffi.CallMethod[TableCellView](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTableCellView() TableCellView {
	return TableCellViewClass.New()
}

func (t_ TableCellView) ObjectValue() objc.Object {
	rv := ffi.CallMethod[objc.Object](t_, "objectValue")
	return rv
}

func (t_ TableCellView) SetObjectValue(value objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "setObjectValue:", value)
}

func (t_ TableCellView) ImageView() ImageView {
	rv := ffi.CallMethod[ImageView](t_, "imageView")
	return rv
}

func (t_ TableCellView) SetImageView(value IImageView) {
	ffi.CallMethod[ffi.Void](t_, "setImageView:", value)
}

func (t_ TableCellView) TextField() TextField {
	rv := ffi.CallMethod[TextField](t_, "textField")
	return rv
}

func (t_ TableCellView) SetTextField(value ITextField) {
	ffi.CallMethod[ffi.Void](t_, "setTextField:", value)
}

func (t_ TableCellView) BackgroundStyle() BackgroundStyle {
	rv := ffi.CallMethod[BackgroundStyle](t_, "backgroundStyle")
	return rv
}

func (t_ TableCellView) SetBackgroundStyle(value BackgroundStyle) {
	ffi.CallMethod[ffi.Void](t_, "setBackgroundStyle:", value)
}

func (t_ TableCellView) RowSizeStyle() TableViewRowSizeStyle {
	rv := ffi.CallMethod[TableViewRowSizeStyle](t_, "rowSizeStyle")
	return rv
}

func (t_ TableCellView) SetRowSizeStyle(value TableViewRowSizeStyle) {
	ffi.CallMethod[ffi.Void](t_, "setRowSizeStyle:", value)
}

func (t_ TableCellView) DraggingImageComponents() []DraggingImageComponent {
	rv := ffi.CallMethod[[]DraggingImageComponent](t_, "draggingImageComponents")
	return rv
}

var TableHeaderCellClass = _TableHeaderCellClass{objc.GetClass("NSTableHeaderCell")}

type _TableHeaderCellClass struct {
	objc.Class
}

type ITableHeaderCell interface {
	ITextFieldCell
	DrawSortIndicatorWithFrame_InView_Ascending_Priority(cellFrame foundation.Rect, controlView IView, ascending bool, priority int)
	SortIndicatorRectForBounds(rect foundation.Rect) foundation.Rect
}

type TableHeaderCell struct {
	TextFieldCell
}

func MakeTableHeaderCell(ptr unsafe.Pointer) TableHeaderCell {
	return TableHeaderCell{
		TextFieldCell: MakeTextFieldCell(ptr),
	}
}

func (t_ TableHeaderCell) InitTextCell(_string string) TableHeaderCell {
	rv := ffi.CallMethod[TableHeaderCell](t_, "initTextCell:", _string)
	rv.Autorelease()
	return rv
}

func (t_ TableHeaderCell) InitImageCell(image IImage) TableHeaderCell {
	rv := ffi.CallMethod[TableHeaderCell](t_, "initImageCell:", image)
	rv.Autorelease()
	return rv
}

func (t_ TableHeaderCell) Init() TableHeaderCell {
	rv := ffi.CallMethod[TableHeaderCell](t_, "init")
	rv.Autorelease()
	return rv
}

func (tc _TableHeaderCellClass) Alloc() TableHeaderCell {
	rv := ffi.CallMethod[TableHeaderCell](tc, "alloc")
	return rv
}

func (tc _TableHeaderCellClass) New() TableHeaderCell {
	rv := ffi.CallMethod[TableHeaderCell](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTableHeaderCell() TableHeaderCell {
	return TableHeaderCellClass.New()
}

func (t_ TableHeaderCell) DrawSortIndicatorWithFrame_InView_Ascending_Priority(cellFrame foundation.Rect, controlView IView, ascending bool, priority int) {
	ffi.CallMethod[ffi.Void](t_, "drawSortIndicatorWithFrame:inView:ascending:priority:", cellFrame, controlView, ascending, priority)
}

func (t_ TableHeaderCell) SortIndicatorRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](t_, "sortIndicatorRectForBounds:", rect)
	return rv
}

var TableColumnClass = _TableColumnClass{objc.GetClass("NSTableColumn")}

type _TableColumnClass struct {
	objc.Class
}

type ITableColumn interface {
	objc.IObject
	SizeToFit()
	// deprecated
	IsResizable() bool
	// deprecated
	SetResizable(flag bool)
	// deprecated
	DataCellForRow(row int) objc.Object
	TableView() TableView
	SetTableView(value ITableView)
	Width() float64
	SetWidth(value float64)
	MinWidth() float64
	SetMinWidth(value float64)
	MaxWidth() float64
	SetMaxWidth(value float64)
	ResizingMask() TableColumnResizingOptions
	SetResizingMask(value TableColumnResizingOptions)
	Title() string
	SetTitle(value string)
	HeaderCell() TableHeaderCell
	SetHeaderCell(value ITableHeaderCell)
	Identifier() UserInterfaceItemIdentifier
	SetIdentifier(value UserInterfaceItemIdentifier)
	IsEditable() bool
	SetEditable(value bool)
	SortDescriptorPrototype() foundation.SortDescriptor
	SetSortDescriptorPrototype(value foundation.ISortDescriptor)
	IsHidden() bool
	SetHidden(value bool)
	HeaderToolTip() string
	SetHeaderToolTip(value string)
	// deprecated
	DataCell() objc.Object
	// deprecated
	SetDataCell(value objc.IObject)
}

type TableColumn struct {
	objc.Object
}

func MakeTableColumn(ptr unsafe.Pointer) TableColumn {
	return TableColumn{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TableColumn) InitWithIdentifier(identifier UserInterfaceItemIdentifier) TableColumn {
	rv := ffi.CallMethod[TableColumn](t_, "initWithIdentifier:", identifier)
	rv.Autorelease()
	return rv
}

func (tc _TableColumnClass) Alloc() TableColumn {
	rv := ffi.CallMethod[TableColumn](tc, "alloc")
	return rv
}

func (t_ TableColumn) Init() TableColumn {
	rv := ffi.CallMethod[TableColumn](t_, "init")
	rv.Autorelease()
	return rv
}

func (tc _TableColumnClass) New() TableColumn {
	rv := ffi.CallMethod[TableColumn](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTableColumn() TableColumn {
	return TableColumnClass.New()
}

func (t_ TableColumn) SizeToFit() {
	ffi.CallMethod[ffi.Void](t_, "sizeToFit")
}

// deprecated
func (t_ TableColumn) IsResizable() bool {
	rv := ffi.CallMethod[bool](t_, "isResizable")
	return rv
}

// deprecated
func (t_ TableColumn) SetResizable(flag bool) {
	ffi.CallMethod[ffi.Void](t_, "setResizable:", flag)
}

// deprecated
func (t_ TableColumn) DataCellForRow(row int) objc.Object {
	rv := ffi.CallMethod[objc.Object](t_, "dataCellForRow:", row)
	return rv
}

func (t_ TableColumn) TableView() TableView {
	rv := ffi.CallMethod[TableView](t_, "tableView")
	return rv
}

func (t_ TableColumn) SetTableView(value ITableView) {
	ffi.CallMethod[ffi.Void](t_, "setTableView:", value)
}

func (t_ TableColumn) Width() float64 {
	rv := ffi.CallMethod[float64](t_, "width")
	return rv
}

func (t_ TableColumn) SetWidth(value float64) {
	ffi.CallMethod[ffi.Void](t_, "setWidth:", value)
}

func (t_ TableColumn) MinWidth() float64 {
	rv := ffi.CallMethod[float64](t_, "minWidth")
	return rv
}

func (t_ TableColumn) SetMinWidth(value float64) {
	ffi.CallMethod[ffi.Void](t_, "setMinWidth:", value)
}

func (t_ TableColumn) MaxWidth() float64 {
	rv := ffi.CallMethod[float64](t_, "maxWidth")
	return rv
}

func (t_ TableColumn) SetMaxWidth(value float64) {
	ffi.CallMethod[ffi.Void](t_, "setMaxWidth:", value)
}

func (t_ TableColumn) ResizingMask() TableColumnResizingOptions {
	rv := ffi.CallMethod[TableColumnResizingOptions](t_, "resizingMask")
	return rv
}

func (t_ TableColumn) SetResizingMask(value TableColumnResizingOptions) {
	ffi.CallMethod[ffi.Void](t_, "setResizingMask:", value)
}

func (t_ TableColumn) Title() string {
	rv := ffi.CallMethod[string](t_, "title")
	return rv
}

func (t_ TableColumn) SetTitle(value string) {
	ffi.CallMethod[ffi.Void](t_, "setTitle:", value)
}

func (t_ TableColumn) HeaderCell() TableHeaderCell {
	rv := ffi.CallMethod[TableHeaderCell](t_, "headerCell")
	return rv
}

func (t_ TableColumn) SetHeaderCell(value ITableHeaderCell) {
	ffi.CallMethod[ffi.Void](t_, "setHeaderCell:", value)
}

func (t_ TableColumn) Identifier() UserInterfaceItemIdentifier {
	rv := ffi.CallMethod[UserInterfaceItemIdentifier](t_, "identifier")
	return rv
}

func (t_ TableColumn) SetIdentifier(value UserInterfaceItemIdentifier) {
	ffi.CallMethod[ffi.Void](t_, "setIdentifier:", value)
}

func (t_ TableColumn) IsEditable() bool {
	rv := ffi.CallMethod[bool](t_, "isEditable")
	return rv
}

func (t_ TableColumn) SetEditable(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setEditable:", value)
}

func (t_ TableColumn) SortDescriptorPrototype() foundation.SortDescriptor {
	rv := ffi.CallMethod[foundation.SortDescriptor](t_, "sortDescriptorPrototype")
	return rv
}

func (t_ TableColumn) SetSortDescriptorPrototype(value foundation.ISortDescriptor) {
	ffi.CallMethod[ffi.Void](t_, "setSortDescriptorPrototype:", value)
}

func (t_ TableColumn) IsHidden() bool {
	rv := ffi.CallMethod[bool](t_, "isHidden")
	return rv
}

func (t_ TableColumn) SetHidden(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setHidden:", value)
}

func (t_ TableColumn) HeaderToolTip() string {
	rv := ffi.CallMethod[string](t_, "headerToolTip")
	return rv
}

func (t_ TableColumn) SetHeaderToolTip(value string) {
	ffi.CallMethod[ffi.Void](t_, "setHeaderToolTip:", value)
}

// deprecated
func (t_ TableColumn) DataCell() objc.Object {
	rv := ffi.CallMethod[objc.Object](t_, "dataCell")
	return rv
}

// deprecated
func (t_ TableColumn) SetDataCell(value objc.IObject) {
	ffi.CallMethod[ffi.Void](t_, "setDataCell:", value)
}

var TableRowViewClass = _TableRowViewClass{objc.GetClass("NSTableRowView")}

type _TableRowViewClass struct {
	objc.Class
}

type ITableRowView interface {
	IView
	DrawBackgroundInRect(dirtyRect foundation.Rect)
	DrawDraggingDestinationFeedbackInRect(dirtyRect foundation.Rect)
	DrawSelectionInRect(dirtyRect foundation.Rect)
	DrawSeparatorInRect(dirtyRect foundation.Rect)
	ViewAtColumn(column int) objc.Object
	IsEmphasized() bool
	SetEmphasized(value bool)
	InteriorBackgroundStyle() BackgroundStyle
	IsFloating() bool
	SetFloating(value bool)
	IsSelected() bool
	SetSelected(value bool)
	SelectionHighlightStyle() TableViewSelectionHighlightStyle
	SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle)
	DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle
	SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle)
	IndentationForDropOperation() float64
	SetIndentationForDropOperation(value float64)
	IsTargetForDropOperation() bool
	SetTargetForDropOperation(value bool)
	IsGroupRowStyle() bool
	SetGroupRowStyle(value bool)
	NumberOfColumns() int
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	IsNextRowSelected() bool
	SetNextRowSelected(value bool)
	IsPreviousRowSelected() bool
	SetPreviousRowSelected(value bool)
}

type TableRowView struct {
	View
}

func MakeTableRowView(ptr unsafe.Pointer) TableRowView {
	return TableRowView{
		View: MakeView(ptr),
	}
}

func (t_ TableRowView) InitWithFrame(frameRect foundation.Rect) TableRowView {
	rv := ffi.CallMethod[TableRowView](t_, "initWithFrame:", frameRect)
	rv.Autorelease()
	return rv
}

func (t_ TableRowView) Init() TableRowView {
	rv := ffi.CallMethod[TableRowView](t_, "init")
	rv.Autorelease()
	return rv
}

func (tc _TableRowViewClass) Alloc() TableRowView {
	rv := ffi.CallMethod[TableRowView](tc, "alloc")
	return rv
}

func (tc _TableRowViewClass) New() TableRowView {
	rv := ffi.CallMethod[TableRowView](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTableRowView() TableRowView {
	return TableRowViewClass.New()
}

func (t_ TableRowView) DrawBackgroundInRect(dirtyRect foundation.Rect) {
	ffi.CallMethod[ffi.Void](t_, "drawBackgroundInRect:", dirtyRect)
}

func (t_ TableRowView) DrawDraggingDestinationFeedbackInRect(dirtyRect foundation.Rect) {
	ffi.CallMethod[ffi.Void](t_, "drawDraggingDestinationFeedbackInRect:", dirtyRect)
}

func (t_ TableRowView) DrawSelectionInRect(dirtyRect foundation.Rect) {
	ffi.CallMethod[ffi.Void](t_, "drawSelectionInRect:", dirtyRect)
}

func (t_ TableRowView) DrawSeparatorInRect(dirtyRect foundation.Rect) {
	ffi.CallMethod[ffi.Void](t_, "drawSeparatorInRect:", dirtyRect)
}

func (t_ TableRowView) ViewAtColumn(column int) objc.Object {
	rv := ffi.CallMethod[objc.Object](t_, "viewAtColumn:", column)
	return rv
}

func (t_ TableRowView) IsEmphasized() bool {
	rv := ffi.CallMethod[bool](t_, "isEmphasized")
	return rv
}

func (t_ TableRowView) SetEmphasized(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setEmphasized:", value)
}

func (t_ TableRowView) InteriorBackgroundStyle() BackgroundStyle {
	rv := ffi.CallMethod[BackgroundStyle](t_, "interiorBackgroundStyle")
	return rv
}

func (t_ TableRowView) IsFloating() bool {
	rv := ffi.CallMethod[bool](t_, "isFloating")
	return rv
}

func (t_ TableRowView) SetFloating(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setFloating:", value)
}

func (t_ TableRowView) IsSelected() bool {
	rv := ffi.CallMethod[bool](t_, "isSelected")
	return rv
}

func (t_ TableRowView) SetSelected(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setSelected:", value)
}

func (t_ TableRowView) SelectionHighlightStyle() TableViewSelectionHighlightStyle {
	rv := ffi.CallMethod[TableViewSelectionHighlightStyle](t_, "selectionHighlightStyle")
	return rv
}

func (t_ TableRowView) SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle) {
	ffi.CallMethod[ffi.Void](t_, "setSelectionHighlightStyle:", value)
}

func (t_ TableRowView) DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle {
	rv := ffi.CallMethod[TableViewDraggingDestinationFeedbackStyle](t_, "draggingDestinationFeedbackStyle")
	return rv
}

func (t_ TableRowView) SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle) {
	ffi.CallMethod[ffi.Void](t_, "setDraggingDestinationFeedbackStyle:", value)
}

func (t_ TableRowView) IndentationForDropOperation() float64 {
	rv := ffi.CallMethod[float64](t_, "indentationForDropOperation")
	return rv
}

func (t_ TableRowView) SetIndentationForDropOperation(value float64) {
	ffi.CallMethod[ffi.Void](t_, "setIndentationForDropOperation:", value)
}

func (t_ TableRowView) IsTargetForDropOperation() bool {
	rv := ffi.CallMethod[bool](t_, "isTargetForDropOperation")
	return rv
}

func (t_ TableRowView) SetTargetForDropOperation(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setTargetForDropOperation:", value)
}

func (t_ TableRowView) IsGroupRowStyle() bool {
	rv := ffi.CallMethod[bool](t_, "isGroupRowStyle")
	return rv
}

func (t_ TableRowView) SetGroupRowStyle(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setGroupRowStyle:", value)
}

func (t_ TableRowView) NumberOfColumns() int {
	rv := ffi.CallMethod[int](t_, "numberOfColumns")
	return rv
}

func (t_ TableRowView) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](t_, "backgroundColor")
	return rv
}

func (t_ TableRowView) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](t_, "setBackgroundColor:", value)
}

func (t_ TableRowView) IsNextRowSelected() bool {
	rv := ffi.CallMethod[bool](t_, "isNextRowSelected")
	return rv
}

func (t_ TableRowView) SetNextRowSelected(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setNextRowSelected:", value)
}

func (t_ TableRowView) IsPreviousRowSelected() bool {
	rv := ffi.CallMethod[bool](t_, "isPreviousRowSelected")
	return rv
}

func (t_ TableRowView) SetPreviousRowSelected(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setPreviousRowSelected:", value)
}

var TableHeaderViewClass = _TableHeaderViewClass{objc.GetClass("NSTableHeaderView")}

type _TableHeaderViewClass struct {
	objc.Class
}

type ITableHeaderView interface {
	IView
	ColumnAtPoint(point foundation.Point) int
	HeaderRectOfColumn(column int) foundation.Rect
	TableView() TableView
	SetTableView(value ITableView)
	DraggedColumn() int
	DraggedDistance() float64
	ResizedColumn() int
}

type TableHeaderView struct {
	View
}

func MakeTableHeaderView(ptr unsafe.Pointer) TableHeaderView {
	return TableHeaderView{
		View: MakeView(ptr),
	}
}

func (t_ TableHeaderView) InitWithFrame(frameRect foundation.Rect) TableHeaderView {
	rv := ffi.CallMethod[TableHeaderView](t_, "initWithFrame:", frameRect)
	rv.Autorelease()
	return rv
}

func (t_ TableHeaderView) Init() TableHeaderView {
	rv := ffi.CallMethod[TableHeaderView](t_, "init")
	rv.Autorelease()
	return rv
}

func (tc _TableHeaderViewClass) Alloc() TableHeaderView {
	rv := ffi.CallMethod[TableHeaderView](tc, "alloc")
	return rv
}

func (tc _TableHeaderViewClass) New() TableHeaderView {
	rv := ffi.CallMethod[TableHeaderView](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTableHeaderView() TableHeaderView {
	return TableHeaderViewClass.New()
}

func (t_ TableHeaderView) ColumnAtPoint(point foundation.Point) int {
	rv := ffi.CallMethod[int](t_, "columnAtPoint:", point)
	return rv
}

func (t_ TableHeaderView) HeaderRectOfColumn(column int) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](t_, "headerRectOfColumn:", column)
	return rv
}

func (t_ TableHeaderView) TableView() TableView {
	rv := ffi.CallMethod[TableView](t_, "tableView")
	return rv
}

func (t_ TableHeaderView) SetTableView(value ITableView) {
	ffi.CallMethod[ffi.Void](t_, "setTableView:", value)
}

func (t_ TableHeaderView) DraggedColumn() int {
	rv := ffi.CallMethod[int](t_, "draggedColumn")
	return rv
}

func (t_ TableHeaderView) DraggedDistance() float64 {
	rv := ffi.CallMethod[float64](t_, "draggedDistance")
	return rv
}

func (t_ TableHeaderView) ResizedColumn() int {
	rv := ffi.CallMethod[int](t_, "resizedColumn")
	return rv
}

var TableViewRowActionClass = _TableViewRowActionClass{objc.GetClass("NSTableViewRowAction")}

type _TableViewRowActionClass struct {
	objc.Class
}

type ITableViewRowAction interface {
	objc.IObject
	Style() TableViewRowActionStyle
	Title() string
	SetTitle(value string)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	Image() Image
	SetImage(value IImage)
}

type TableViewRowAction struct {
	objc.Object
}

func MakeTableViewRowAction(ptr unsafe.Pointer) TableViewRowAction {
	return TableViewRowAction{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TableViewRowActionClass) RowActionWithStyle_Title_Handler(style TableViewRowActionStyle, title string, handler func(action TableViewRowAction, row int)) TableViewRowAction {
	rv := ffi.CallMethod[TableViewRowAction](tc, "rowActionWithStyle:title:handler:", style, title, handler)
	return rv
}

func (tc _TableViewRowActionClass) Alloc() TableViewRowAction {
	rv := ffi.CallMethod[TableViewRowAction](tc, "alloc")
	return rv
}

func (t_ TableViewRowAction) Init() TableViewRowAction {
	rv := ffi.CallMethod[TableViewRowAction](t_, "init")
	rv.Autorelease()
	return rv
}

func (tc _TableViewRowActionClass) New() TableViewRowAction {
	rv := ffi.CallMethod[TableViewRowAction](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTableViewRowAction() TableViewRowAction {
	return TableViewRowActionClass.New()
}

func (t_ TableViewRowAction) Style() TableViewRowActionStyle {
	rv := ffi.CallMethod[TableViewRowActionStyle](t_, "style")
	return rv
}

func (t_ TableViewRowAction) Title() string {
	rv := ffi.CallMethod[string](t_, "title")
	return rv
}

func (t_ TableViewRowAction) SetTitle(value string) {
	ffi.CallMethod[ffi.Void](t_, "setTitle:", value)
}

func (t_ TableViewRowAction) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](t_, "backgroundColor")
	return rv
}

func (t_ TableViewRowAction) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](t_, "setBackgroundColor:", value)
}

func (t_ TableViewRowAction) Image() Image {
	rv := ffi.CallMethod[Image](t_, "image")
	return rv
}

func (t_ TableViewRowAction) SetImage(value IImage) {
	ffi.CallMethod[ffi.Void](t_, "setImage:", value)
}
