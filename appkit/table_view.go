// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	SetDataSource0(value objc.IObject)
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
	SetDelegate0(value objc.IObject)
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
	rv := objc.CallMethod[TableView](t_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func (t_ TableView) Init() TableView {
	rv := objc.CallMethod[TableView](t_, objc.GetSelector("init"))
	return rv
}

func (tc _TableViewClass) Alloc() TableView {
	rv := objc.CallMethod[TableView](tc, objc.GetSelector("alloc"))
	return rv
}

func (tc _TableViewClass) New() TableView {
	rv := objc.CallMethod[TableView](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTableView() TableView {
	return TableViewClass.New()
}

func (t_ TableView) ReloadData() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("reloadData"))
}

func (t_ TableView) ReloadDataForRowIndexes_ColumnIndexes(rowIndexes foundation.IIndexSet, columnIndexes foundation.IIndexSet) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("reloadDataForRowIndexes:columnIndexes:"), rowIndexes, columnIndexes)
}

func (t_ TableView) MakeViewWithIdentifier_Owner(identifier UserInterfaceItemIdentifier, owner objc.IObject) View {
	rv := objc.CallMethod[View](t_, objc.GetSelector("makeViewWithIdentifier:owner:"), identifier, owner)
	return rv
}

func (t_ TableView) RowViewAtRow_MakeIfNecessary(row int, makeIfNecessary bool) TableRowView {
	rv := objc.CallMethod[TableRowView](t_, objc.GetSelector("rowViewAtRow:makeIfNecessary:"), row, makeIfNecessary)
	return rv
}

func (t_ TableView) ViewAtColumn_Row_MakeIfNecessary(column int, row int, makeIfNecessary bool) View {
	rv := objc.CallMethod[View](t_, objc.GetSelector("viewAtColumn:row:makeIfNecessary:"), column, row, makeIfNecessary)
	return rv
}

func (t_ TableView) BeginUpdates() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("beginUpdates"))
}

func (t_ TableView) EndUpdates() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("endUpdates"))
}

func (t_ TableView) MoveRowAtIndex_ToIndex(oldIndex int, newIndex int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("moveRowAtIndex:toIndex:"), oldIndex, newIndex)
}

func (t_ TableView) InsertRowsAtIndexes_WithAnimation(indexes foundation.IIndexSet, animationOptions TableViewAnimationOptions) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("insertRowsAtIndexes:withAnimation:"), indexes, animationOptions)
}

func (t_ TableView) RemoveRowsAtIndexes_WithAnimation(indexes foundation.IIndexSet, animationOptions TableViewAnimationOptions) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("removeRowsAtIndexes:withAnimation:"), indexes, animationOptions)
}

func (t_ TableView) RowForView(view IView) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("rowForView:"), view)
	return rv
}

func (t_ TableView) ColumnForView(view IView) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("columnForView:"), view)
	return rv
}

func (t_ TableView) RegisterNib_ForIdentifier(nib INib, identifier UserInterfaceItemIdentifier) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("registerNib:forIdentifier:"), nib, identifier)
}

func (t_ TableView) IndicatorImageInTableColumn(tableColumn ITableColumn) Image {
	rv := objc.CallMethod[Image](t_, objc.GetSelector("indicatorImageInTableColumn:"), tableColumn)
	return rv
}

func (t_ TableView) SetIndicatorImage_InTableColumn(image IImage, tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setIndicatorImage:inTableColumn:"), image, tableColumn)
}

func (t_ TableView) AddTableColumn(tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("addTableColumn:"), tableColumn)
}

func (t_ TableView) RemoveTableColumn(tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("removeTableColumn:"), tableColumn)
}

func (t_ TableView) MoveColumn_ToColumn(oldIndex int, newIndex int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("moveColumn:toColumn:"), oldIndex, newIndex)
}

func (t_ TableView) ColumnWithIdentifier(identifier UserInterfaceItemIdentifier) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("columnWithIdentifier:"), identifier)
	return rv
}

func (t_ TableView) TableColumnWithIdentifier(identifier UserInterfaceItemIdentifier) TableColumn {
	rv := objc.CallMethod[TableColumn](t_, objc.GetSelector("tableColumnWithIdentifier:"), identifier)
	return rv
}

func (t_ TableView) SelectColumnIndexes_ByExtendingSelection(indexes foundation.IIndexSet, extend bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("selectColumnIndexes:byExtendingSelection:"), indexes, extend)
}

func (t_ TableView) DeselectColumn(column int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("deselectColumn:"), column)
}

func (t_ TableView) IsColumnSelected(column int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isColumnSelected:"), column)
	return rv
}

func (t_ TableView) SelectRowIndexes_ByExtendingSelection(indexes foundation.IIndexSet, extend bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("selectRowIndexes:byExtendingSelection:"), indexes, extend)
}

func (t_ TableView) DeselectRow(row int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("deselectRow:"), row)
}

func (t_ TableView) IsRowSelected(row int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isRowSelected:"), row)
	return rv
}

func (t_ TableView) SelectAll(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("selectAll:"), sender)
}

func (t_ TableView) DeselectAll(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("deselectAll:"), sender)
}

func (t_ TableView) EnumerateAvailableRowViewsUsingBlock(handler func(rowView TableRowView, row int)) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("enumerateAvailableRowViewsUsingBlock:"), handler)
}

func (t_ TableView) EditColumn_Row_WithEvent_Select(column int, row int, event IEvent, select_ bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("editColumn:row:withEvent:select:"), column, row, event, select_)
}

func (t_ TableView) DidAddRowView_ForRow(rowView ITableRowView, row int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("didAddRowView:forRow:"), rowView, row)
}

func (t_ TableView) DidRemoveRowView_ForRow(rowView ITableRowView, row int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("didRemoveRowView:forRow:"), rowView, row)
}

func (t_ TableView) RectOfColumn(column int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.GetSelector("rectOfColumn:"), column)
	return rv
}

func (t_ TableView) RectOfRow(row int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.GetSelector("rectOfRow:"), row)
	return rv
}

func (t_ TableView) RowsInRect(rect foundation.Rect) foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("rowsInRect:"), rect)
	return rv
}

func (t_ TableView) ColumnIndexesInRect(rect foundation.Rect) foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](t_, objc.GetSelector("columnIndexesInRect:"), rect)
	return rv
}

func (t_ TableView) ColumnAtPoint(point foundation.Point) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("columnAtPoint:"), point)
	return rv
}

func (t_ TableView) RowAtPoint(point foundation.Point) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("rowAtPoint:"), point)
	return rv
}

func (t_ TableView) FrameOfCellAtColumn_Row(column int, row int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.GetSelector("frameOfCellAtColumn:row:"), column, row)
	return rv
}

func (t_ TableView) SizeLastColumnToFit() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("sizeLastColumnToFit"))
}

func (t_ TableView) NoteNumberOfRowsChanged() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("noteNumberOfRowsChanged"))
}

func (t_ TableView) Tile() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tile"))
}

func (t_ TableView) NoteHeightOfRowsWithIndexesChanged(indexSet foundation.IIndexSet) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("noteHeightOfRowsWithIndexesChanged:"), indexSet)
}

func (t_ TableView) DrawRow_ClipRect(row int, clipRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("drawRow:clipRect:"), row, clipRect)
}

func (t_ TableView) DrawGridInClipRect(clipRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("drawGridInClipRect:"), clipRect)
}

func (t_ TableView) HighlightSelectionInClipRect(clipRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("highlightSelectionInClipRect:"), clipRect)
}

func (t_ TableView) DrawBackgroundInClipRect(clipRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("drawBackgroundInClipRect:"), clipRect)
}

func (t_ TableView) ScrollRowToVisible(row int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("scrollRowToVisible:"), row)
}

func (t_ TableView) ScrollColumnToVisible(column int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("scrollColumnToVisible:"), column)
}

func (t_ TableView) DragImageForRowsWithIndexes_TableColumns_Event_Offset(dragRows foundation.IIndexSet, tableColumns []ITableColumn, dragEvent IEvent, dragImageOffset *foundation.Point) Image {
	rv := objc.CallMethod[Image](t_, objc.GetSelector("dragImageForRowsWithIndexes:tableColumns:event:offset:"), dragRows, tableColumns, dragEvent, dragImageOffset)
	return rv
}

func (t_ TableView) CanDragRowsWithIndexes_AtPoint(rowIndexes foundation.IIndexSet, mouseDownPoint foundation.Point) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("canDragRowsWithIndexes:atPoint:"), rowIndexes, mouseDownPoint)
	return rv
}

func (t_ TableView) SetDraggingSourceOperationMask_ForLocal(mask DragOperation, isLocal bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDraggingSourceOperationMask:forLocal:"), mask, isLocal)
}

func (t_ TableView) SetDropRow_DropOperation(row int, dropOperation TableViewDropOperation) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDropRow:dropOperation:"), row, dropOperation)
}

func (t_ TableView) HideRowsAtIndexes_WithAnimation(indexes foundation.IIndexSet, rowAnimation TableViewAnimationOptions) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("hideRowsAtIndexes:withAnimation:"), indexes, rowAnimation)
}

func (t_ TableView) UnhideRowsAtIndexes_WithAnimation(indexes foundation.IIndexSet, rowAnimation TableViewAnimationOptions) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("unhideRowsAtIndexes:withAnimation:"), indexes, rowAnimation)
}

// deprecated
func (t_ TableView) DragImageForRows_Event_DragImageOffset(dragRows []objc.IObject, dragEvent IEvent, dragImageOffset *foundation.Point) Image {
	rv := objc.CallMethod[Image](t_, objc.GetSelector("dragImageForRows:event:dragImageOffset:"), dragRows, dragEvent, dragImageOffset)
	return rv
}

// deprecated
func (t_ TableView) SetAutoresizesAllColumnsToFit(flag bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAutoresizesAllColumnsToFit:"), flag)
}

// deprecated
func (t_ TableView) AutoresizesAllColumnsToFit() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("autoresizesAllColumnsToFit"))
	return rv
}

// deprecated
func (t_ TableView) SelectColumn_ByExtendingSelection(column int, extend bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("selectColumn:byExtendingSelection:"), column, extend)
}

// deprecated
func (t_ TableView) SelectRow_ByExtendingSelection(row int, extend bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("selectRow:byExtendingSelection:"), row, extend)
}

// deprecated
func (t_ TableView) TableView_WriteRows_ToPasteboard(tableView ITableView, rows []objc.IObject, pboard IPasteboard) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tableView:writeRows:toPasteboard:"), tableView, rows, pboard)
	return rv
}

// deprecated
func (t_ TableView) SetDrawsGrid(flag bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDrawsGrid:"), flag)
}

// deprecated
func (t_ TableView) DrawsGrid() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("drawsGrid"))
	return rv
}

// deprecated
func (t_ TableView) SelectedColumnEnumerator() foundation.Enumerator {
	rv := objc.CallMethod[foundation.Enumerator](t_, objc.GetSelector("selectedColumnEnumerator"))
	return rv
}

// deprecated
func (t_ TableView) SelectedRowEnumerator() foundation.Enumerator {
	rv := objc.CallMethod[foundation.Enumerator](t_, objc.GetSelector("selectedRowEnumerator"))
	return rv
}

// deprecated
func (t_ TableView) FocusedColumn() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("focusedColumn"))
	return rv
}

// deprecated
func (t_ TableView) SetFocusedColumn(focusedColumn int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setFocusedColumn:"), focusedColumn)
}

// deprecated
func (t_ TableView) ShouldFocusCell_AtColumn_Row(cell ICell, column int, row int) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("shouldFocusCell:atColumn:row:"), cell, column, row)
	return rv
}

// deprecated
func (t_ TableView) PerformClickOnCellAtColumn_Row(column int, row int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("performClickOnCellAtColumn:row:"), column, row)
}

// deprecated
func (t_ TableView) PreparedCellAtColumn_Row(column int, row int) Cell {
	rv := objc.CallMethod[Cell](t_, objc.GetSelector("preparedCellAtColumn:row:"), column, row)
	return rv
}

// deprecated
func (t_ TableView) ColumnsInRect(rect foundation.Rect) foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("columnsInRect:"), rect)
	return rv
}

// deprecated
func (t_ TableView) TextShouldBeginEditing(textObject IText) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textShouldBeginEditing:"), textObject)
	return rv
}

// deprecated
func (t_ TableView) TextDidBeginEditing(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textDidBeginEditing:"), notification)
}

// deprecated
func (t_ TableView) TextDidChange(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textDidChange:"), notification)
}

// deprecated
func (t_ TableView) TextShouldEndEditing(textObject IText) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textShouldEndEditing:"), textObject)
	return rv
}

// deprecated
func (t_ TableView) TextDidEndEditing(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textDidEndEditing:"), notification)
}

func (t_ TableView) DataSource() TableViewDataSourceWrapper {
	rv := objc.CallMethod[TableViewDataSourceWrapper](t_, objc.GetSelector("dataSource"))
	return rv
}

func (t_ TableView) SetDataSource(value TableViewDataSource) {
	po := objc.CreateProtocol("NSTableViewDataSource", value)
	defer po.Release()
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDataSource"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDataSource:"), po)
}

func (t_ TableView) SetDataSource0(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDataSource:"), value)
}

func (t_ TableView) UsesStaticContents() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("usesStaticContents"))
	return rv
}

func (t_ TableView) SetUsesStaticContents(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setUsesStaticContents:"), value)
}

func (t_ TableView) RegisteredNibsByIdentifier() map[UserInterfaceItemIdentifier]Nib {
	rv := objc.CallMethod[map[UserInterfaceItemIdentifier]Nib](t_, objc.GetSelector("registeredNibsByIdentifier"))
	return rv
}

func (t_ TableView) DoubleAction() objc.Selector {
	rv := objc.CallMethod[objc.Selector](t_, objc.GetSelector("doubleAction"))
	return rv
}

func (t_ TableView) SetDoubleAction(value objc.Selector) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDoubleAction:"), value)
}

func (t_ TableView) ClickedColumn() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("clickedColumn"))
	return rv
}

func (t_ TableView) ClickedRow() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("clickedRow"))
	return rv
}

func (t_ TableView) AllowsColumnReordering() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsColumnReordering"))
	return rv
}

func (t_ TableView) SetAllowsColumnReordering(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsColumnReordering:"), value)
}

func (t_ TableView) AllowsColumnResizing() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsColumnResizing"))
	return rv
}

func (t_ TableView) SetAllowsColumnResizing(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsColumnResizing:"), value)
}

func (t_ TableView) AllowsMultipleSelection() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsMultipleSelection"))
	return rv
}

func (t_ TableView) SetAllowsMultipleSelection(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsMultipleSelection:"), value)
}

func (t_ TableView) AllowsEmptySelection() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsEmptySelection"))
	return rv
}

func (t_ TableView) SetAllowsEmptySelection(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsEmptySelection:"), value)
}

func (t_ TableView) AllowsColumnSelection() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsColumnSelection"))
	return rv
}

func (t_ TableView) SetAllowsColumnSelection(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsColumnSelection:"), value)
}

func (t_ TableView) UsesAutomaticRowHeights() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("usesAutomaticRowHeights"))
	return rv
}

func (t_ TableView) SetUsesAutomaticRowHeights(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setUsesAutomaticRowHeights:"), value)
}

func (t_ TableView) IntercellSpacing() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.GetSelector("intercellSpacing"))
	return rv
}

func (t_ TableView) SetIntercellSpacing(value foundation.Size) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setIntercellSpacing:"), value)
}

func (t_ TableView) RowHeight() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("rowHeight"))
	return rv
}

func (t_ TableView) SetRowHeight(value float64) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setRowHeight:"), value)
}

func (t_ TableView) BackgroundColor() Color {
	rv := objc.CallMethod[Color](t_, objc.GetSelector("backgroundColor"))
	return rv
}

func (t_ TableView) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setBackgroundColor:"), value)
}

func (t_ TableView) UsesAlternatingRowBackgroundColors() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("usesAlternatingRowBackgroundColors"))
	return rv
}

func (t_ TableView) SetUsesAlternatingRowBackgroundColors(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setUsesAlternatingRowBackgroundColors:"), value)
}

func (t_ TableView) Style() TableViewStyle {
	rv := objc.CallMethod[TableViewStyle](t_, objc.GetSelector("style"))
	return rv
}

func (t_ TableView) SetStyle(value TableViewStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setStyle:"), value)
}

func (t_ TableView) EffectiveStyle() TableViewStyle {
	rv := objc.CallMethod[TableViewStyle](t_, objc.GetSelector("effectiveStyle"))
	return rv
}

func (t_ TableView) SelectionHighlightStyle() TableViewSelectionHighlightStyle {
	rv := objc.CallMethod[TableViewSelectionHighlightStyle](t_, objc.GetSelector("selectionHighlightStyle"))
	return rv
}

func (t_ TableView) SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSelectionHighlightStyle:"), value)
}

func (t_ TableView) GridColor() Color {
	rv := objc.CallMethod[Color](t_, objc.GetSelector("gridColor"))
	return rv
}

func (t_ TableView) SetGridColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setGridColor:"), value)
}

func (t_ TableView) GridStyleMask() TableViewGridLineStyle {
	rv := objc.CallMethod[TableViewGridLineStyle](t_, objc.GetSelector("gridStyleMask"))
	return rv
}

func (t_ TableView) SetGridStyleMask(value TableViewGridLineStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setGridStyleMask:"), value)
}

func (t_ TableView) EffectiveRowSizeStyle() TableViewRowSizeStyle {
	rv := objc.CallMethod[TableViewRowSizeStyle](t_, objc.GetSelector("effectiveRowSizeStyle"))
	return rv
}

func (t_ TableView) RowSizeStyle() TableViewRowSizeStyle {
	rv := objc.CallMethod[TableViewRowSizeStyle](t_, objc.GetSelector("rowSizeStyle"))
	return rv
}

func (t_ TableView) SetRowSizeStyle(value TableViewRowSizeStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setRowSizeStyle:"), value)
}

func (t_ TableView) TableColumns() []TableColumn {
	rv := objc.CallMethod[[]TableColumn](t_, objc.GetSelector("tableColumns"))
	return rv
}

func (t_ TableView) SelectedColumn() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("selectedColumn"))
	return rv
}

func (t_ TableView) SelectedColumnIndexes() foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](t_, objc.GetSelector("selectedColumnIndexes"))
	return rv
}

func (t_ TableView) NumberOfSelectedColumns() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("numberOfSelectedColumns"))
	return rv
}

func (t_ TableView) SelectedRow() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("selectedRow"))
	return rv
}

func (t_ TableView) SelectedRowIndexes() foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](t_, objc.GetSelector("selectedRowIndexes"))
	return rv
}

func (t_ TableView) NumberOfSelectedRows() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("numberOfSelectedRows"))
	return rv
}

func (t_ TableView) AllowsTypeSelect() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsTypeSelect"))
	return rv
}

func (t_ TableView) SetAllowsTypeSelect(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsTypeSelect:"), value)
}

func (t_ TableView) NumberOfColumns() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("numberOfColumns"))
	return rv
}

func (t_ TableView) NumberOfRows() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("numberOfRows"))
	return rv
}

func (t_ TableView) FloatsGroupRows() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("floatsGroupRows"))
	return rv
}

func (t_ TableView) SetFloatsGroupRows(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setFloatsGroupRows:"), value)
}

func (t_ TableView) EditedColumn() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("editedColumn"))
	return rv
}

func (t_ TableView) EditedRow() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("editedRow"))
	return rv
}

func (t_ TableView) HeaderView() TableHeaderView {
	rv := objc.CallMethod[TableHeaderView](t_, objc.GetSelector("headerView"))
	return rv
}

func (t_ TableView) SetHeaderView(value ITableHeaderView) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setHeaderView:"), value)
}

func (t_ TableView) CornerView() View {
	rv := objc.CallMethod[View](t_, objc.GetSelector("cornerView"))
	return rv
}

func (t_ TableView) SetCornerView(value IView) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setCornerView:"), value)
}

func (t_ TableView) ColumnAutoresizingStyle() TableViewColumnAutoresizingStyle {
	rv := objc.CallMethod[TableViewColumnAutoresizingStyle](t_, objc.GetSelector("columnAutoresizingStyle"))
	return rv
}

func (t_ TableView) SetColumnAutoresizingStyle(value TableViewColumnAutoresizingStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setColumnAutoresizingStyle:"), value)
}

func (t_ TableView) AutosaveTableColumns() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("autosaveTableColumns"))
	return rv
}

func (t_ TableView) SetAutosaveTableColumns(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAutosaveTableColumns:"), value)
}

func (t_ TableView) AutosaveName() TableViewAutosaveName {
	rv := objc.CallMethod[TableViewAutosaveName](t_, objc.GetSelector("autosaveName"))
	return rv
}

func (t_ TableView) SetAutosaveName(value TableViewAutosaveName) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAutosaveName:"), value)
}

func (t_ TableView) Delegate() TableViewDelegateWrapper {
	rv := objc.CallMethod[TableViewDelegateWrapper](t_, objc.GetSelector("delegate"))
	return rv
}

func (t_ TableView) SetDelegate(value TableViewDelegate) {
	po := objc.CreateProtocol("NSTableViewDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(t_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDelegate:"), po)
}

func (t_ TableView) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDelegate:"), value)
}

func (t_ TableView) HighlightedTableColumn() TableColumn {
	rv := objc.CallMethod[TableColumn](t_, objc.GetSelector("highlightedTableColumn"))
	return rv
}

func (t_ TableView) SetHighlightedTableColumn(value ITableColumn) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setHighlightedTableColumn:"), value)
}

func (t_ TableView) VerticalMotionCanBeginDrag() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("verticalMotionCanBeginDrag"))
	return rv
}

func (t_ TableView) SetVerticalMotionCanBeginDrag(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setVerticalMotionCanBeginDrag:"), value)
}

func (t_ TableView) DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle {
	rv := objc.CallMethod[TableViewDraggingDestinationFeedbackStyle](t_, objc.GetSelector("draggingDestinationFeedbackStyle"))
	return rv
}

func (t_ TableView) SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDraggingDestinationFeedbackStyle:"), value)
}

func (t_ TableView) SortDescriptors() []foundation.SortDescriptor {
	rv := objc.CallMethod[[]foundation.SortDescriptor](t_, objc.GetSelector("sortDescriptors"))
	return rv
}

func (t_ TableView) SetSortDescriptors(value []foundation.ISortDescriptor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSortDescriptors:"), value)
}

func (t_ TableView) RowActionsVisible() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("rowActionsVisible"))
	return rv
}

func (t_ TableView) SetRowActionsVisible(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setRowActionsVisible:"), value)
}

func (t_ TableView) HiddenRowIndexes() foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](t_, objc.GetSelector("hiddenRowIndexes"))
	return rv
}
