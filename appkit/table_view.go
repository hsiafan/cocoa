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
	EditColumn_Row_WithEvent_Select(column int, row int, event IEvent, select_ bool)
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
	return rv
}

func (t_ TableView) Init() TableView {
	rv := ffi.CallMethod[TableView](t_, "init")
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

func (t_ TableView) EditColumn_Row_WithEvent_Select(column int, row int, event IEvent, select_ bool) {
	ffi.CallMethod[ffi.Void](t_, "editColumn:row:withEvent:select:", column, row, event, select_)
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
	po := ffi.CreateProtocol("NSTableViewDataSource", value)
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
	po := ffi.CreateProtocol("NSTableViewDelegate", value)
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
