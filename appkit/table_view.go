// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
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
	DataSource() objc.Object
	SetDataSource(value objc.IObject)
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
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
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
	rv := objc.CallMethod[TableView](t_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (t_ TableView) Init() TableView {
	rv := objc.CallMethod[TableView](t_, objc.SEL("init"))
	return rv
}

func (tc _TableViewClass) Alloc() TableView {
	rv := objc.CallMethod[TableView](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TableViewClass) New() TableView {
	rv := objc.CallMethod[TableView](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTableView() TableView {
	return TableViewClass.New()
}

func (t_ TableView) ReloadData() {
	objc.CallMethod[objc.Void](t_, objc.SEL("reloadData"))
}

func (t_ TableView) ReloadDataForRowIndexes_ColumnIndexes(rowIndexes foundation.IIndexSet, columnIndexes foundation.IIndexSet) {
	objc.CallMethod[objc.Void](t_, objc.SEL("reloadDataForRowIndexes:columnIndexes:"), objc.ExtractPtr(rowIndexes), objc.ExtractPtr(columnIndexes))
}

func (t_ TableView) MakeViewWithIdentifier_Owner(identifier UserInterfaceItemIdentifier, owner objc.IObject) View {
	rv := objc.CallMethod[View](t_, objc.SEL("makeViewWithIdentifier:owner:"), identifier, objc.ExtractPtr(owner))
	return rv
}

func (t_ TableView) RowViewAtRow_MakeIfNecessary(row int, makeIfNecessary bool) TableRowView {
	rv := objc.CallMethod[TableRowView](t_, objc.SEL("rowViewAtRow:makeIfNecessary:"), row, makeIfNecessary)
	return rv
}

func (t_ TableView) ViewAtColumn_Row_MakeIfNecessary(column int, row int, makeIfNecessary bool) View {
	rv := objc.CallMethod[View](t_, objc.SEL("viewAtColumn:row:makeIfNecessary:"), column, row, makeIfNecessary)
	return rv
}

func (t_ TableView) BeginUpdates() {
	objc.CallMethod[objc.Void](t_, objc.SEL("beginUpdates"))
}

func (t_ TableView) EndUpdates() {
	objc.CallMethod[objc.Void](t_, objc.SEL("endUpdates"))
}

func (t_ TableView) MoveRowAtIndex_ToIndex(oldIndex int, newIndex int) {
	objc.CallMethod[objc.Void](t_, objc.SEL("moveRowAtIndex:toIndex:"), oldIndex, newIndex)
}

func (t_ TableView) InsertRowsAtIndexes_WithAnimation(indexes foundation.IIndexSet, animationOptions TableViewAnimationOptions) {
	objc.CallMethod[objc.Void](t_, objc.SEL("insertRowsAtIndexes:withAnimation:"), objc.ExtractPtr(indexes), animationOptions)
}

func (t_ TableView) RemoveRowsAtIndexes_WithAnimation(indexes foundation.IIndexSet, animationOptions TableViewAnimationOptions) {
	objc.CallMethod[objc.Void](t_, objc.SEL("removeRowsAtIndexes:withAnimation:"), objc.ExtractPtr(indexes), animationOptions)
}

func (t_ TableView) RowForView(view IView) int {
	rv := objc.CallMethod[int](t_, objc.SEL("rowForView:"), objc.ExtractPtr(view))
	return rv
}

func (t_ TableView) ColumnForView(view IView) int {
	rv := objc.CallMethod[int](t_, objc.SEL("columnForView:"), objc.ExtractPtr(view))
	return rv
}

func (t_ TableView) RegisterNib_ForIdentifier(nib INib, identifier UserInterfaceItemIdentifier) {
	objc.CallMethod[objc.Void](t_, objc.SEL("registerNib:forIdentifier:"), objc.ExtractPtr(nib), identifier)
}

func (t_ TableView) IndicatorImageInTableColumn(tableColumn ITableColumn) Image {
	rv := objc.CallMethod[Image](t_, objc.SEL("indicatorImageInTableColumn:"), objc.ExtractPtr(tableColumn))
	return rv
}

func (t_ TableView) SetIndicatorImage_InTableColumn(image IImage, tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setIndicatorImage:inTableColumn:"), objc.ExtractPtr(image), objc.ExtractPtr(tableColumn))
}

func (t_ TableView) AddTableColumn(tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](t_, objc.SEL("addTableColumn:"), objc.ExtractPtr(tableColumn))
}

func (t_ TableView) RemoveTableColumn(tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](t_, objc.SEL("removeTableColumn:"), objc.ExtractPtr(tableColumn))
}

func (t_ TableView) MoveColumn_ToColumn(oldIndex int, newIndex int) {
	objc.CallMethod[objc.Void](t_, objc.SEL("moveColumn:toColumn:"), oldIndex, newIndex)
}

func (t_ TableView) ColumnWithIdentifier(identifier UserInterfaceItemIdentifier) int {
	rv := objc.CallMethod[int](t_, objc.SEL("columnWithIdentifier:"), identifier)
	return rv
}

func (t_ TableView) TableColumnWithIdentifier(identifier UserInterfaceItemIdentifier) TableColumn {
	rv := objc.CallMethod[TableColumn](t_, objc.SEL("tableColumnWithIdentifier:"), identifier)
	return rv
}

func (t_ TableView) SelectColumnIndexes_ByExtendingSelection(indexes foundation.IIndexSet, extend bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("selectColumnIndexes:byExtendingSelection:"), objc.ExtractPtr(indexes), extend)
}

func (t_ TableView) DeselectColumn(column int) {
	objc.CallMethod[objc.Void](t_, objc.SEL("deselectColumn:"), column)
}

func (t_ TableView) IsColumnSelected(column int) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isColumnSelected:"), column)
	return rv
}

func (t_ TableView) SelectRowIndexes_ByExtendingSelection(indexes foundation.IIndexSet, extend bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("selectRowIndexes:byExtendingSelection:"), objc.ExtractPtr(indexes), extend)
}

func (t_ TableView) DeselectRow(row int) {
	objc.CallMethod[objc.Void](t_, objc.SEL("deselectRow:"), row)
}

func (t_ TableView) IsRowSelected(row int) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isRowSelected:"), row)
	return rv
}

func (t_ TableView) SelectAll(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("selectAll:"), objc.ExtractPtr(sender))
}

func (t_ TableView) DeselectAll(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("deselectAll:"), objc.ExtractPtr(sender))
}

func (t_ TableView) EnumerateAvailableRowViewsUsingBlock(handler func(rowView TableRowView, row int)) {
	objc.CallMethod[objc.Void](t_, objc.SEL("enumerateAvailableRowViewsUsingBlock:"), handler)
}

func (t_ TableView) EditColumn_Row_WithEvent_Select(column int, row int, event IEvent, select_ bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("editColumn:row:withEvent:select:"), column, row, objc.ExtractPtr(event), select_)
}

func (t_ TableView) DidAddRowView_ForRow(rowView ITableRowView, row int) {
	objc.CallMethod[objc.Void](t_, objc.SEL("didAddRowView:forRow:"), objc.ExtractPtr(rowView), row)
}

func (t_ TableView) DidRemoveRowView_ForRow(rowView ITableRowView, row int) {
	objc.CallMethod[objc.Void](t_, objc.SEL("didRemoveRowView:forRow:"), objc.ExtractPtr(rowView), row)
}

func (t_ TableView) RectOfColumn(column int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.SEL("rectOfColumn:"), column)
	return rv
}

func (t_ TableView) RectOfRow(row int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.SEL("rectOfRow:"), row)
	return rv
}

func (t_ TableView) RowsInRect(rect foundation.Rect) foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.SEL("rowsInRect:"), rect)
	return rv
}

func (t_ TableView) ColumnIndexesInRect(rect foundation.Rect) foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](t_, objc.SEL("columnIndexesInRect:"), rect)
	return rv
}

func (t_ TableView) ColumnAtPoint(point foundation.Point) int {
	rv := objc.CallMethod[int](t_, objc.SEL("columnAtPoint:"), point)
	return rv
}

func (t_ TableView) RowAtPoint(point foundation.Point) int {
	rv := objc.CallMethod[int](t_, objc.SEL("rowAtPoint:"), point)
	return rv
}

func (t_ TableView) FrameOfCellAtColumn_Row(column int, row int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.SEL("frameOfCellAtColumn:row:"), column, row)
	return rv
}

func (t_ TableView) SizeLastColumnToFit() {
	objc.CallMethod[objc.Void](t_, objc.SEL("sizeLastColumnToFit"))
}

func (t_ TableView) NoteNumberOfRowsChanged() {
	objc.CallMethod[objc.Void](t_, objc.SEL("noteNumberOfRowsChanged"))
}

func (t_ TableView) Tile() {
	objc.CallMethod[objc.Void](t_, objc.SEL("tile"))
}

func (t_ TableView) NoteHeightOfRowsWithIndexesChanged(indexSet foundation.IIndexSet) {
	objc.CallMethod[objc.Void](t_, objc.SEL("noteHeightOfRowsWithIndexesChanged:"), objc.ExtractPtr(indexSet))
}

func (t_ TableView) DrawRow_ClipRect(row int, clipRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.SEL("drawRow:clipRect:"), row, clipRect)
}

func (t_ TableView) DrawGridInClipRect(clipRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.SEL("drawGridInClipRect:"), clipRect)
}

func (t_ TableView) HighlightSelectionInClipRect(clipRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.SEL("highlightSelectionInClipRect:"), clipRect)
}

func (t_ TableView) DrawBackgroundInClipRect(clipRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.SEL("drawBackgroundInClipRect:"), clipRect)
}

func (t_ TableView) ScrollRowToVisible(row int) {
	objc.CallMethod[objc.Void](t_, objc.SEL("scrollRowToVisible:"), row)
}

func (t_ TableView) ScrollColumnToVisible(column int) {
	objc.CallMethod[objc.Void](t_, objc.SEL("scrollColumnToVisible:"), column)
}

func (t_ TableView) DragImageForRowsWithIndexes_TableColumns_Event_Offset(dragRows foundation.IIndexSet, tableColumns []ITableColumn, dragEvent IEvent, dragImageOffset *foundation.Point) Image {
	rv := objc.CallMethod[Image](t_, objc.SEL("dragImageForRowsWithIndexes:tableColumns:event:offset:"), objc.ExtractPtr(dragRows), tableColumns, objc.ExtractPtr(dragEvent), dragImageOffset)
	return rv
}

func (t_ TableView) CanDragRowsWithIndexes_AtPoint(rowIndexes foundation.IIndexSet, mouseDownPoint foundation.Point) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("canDragRowsWithIndexes:atPoint:"), objc.ExtractPtr(rowIndexes), mouseDownPoint)
	return rv
}

func (t_ TableView) SetDraggingSourceOperationMask_ForLocal(mask DragOperation, isLocal bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDraggingSourceOperationMask:forLocal:"), mask, isLocal)
}

func (t_ TableView) SetDropRow_DropOperation(row int, dropOperation TableViewDropOperation) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDropRow:dropOperation:"), row, dropOperation)
}

func (t_ TableView) HideRowsAtIndexes_WithAnimation(indexes foundation.IIndexSet, rowAnimation TableViewAnimationOptions) {
	objc.CallMethod[objc.Void](t_, objc.SEL("hideRowsAtIndexes:withAnimation:"), objc.ExtractPtr(indexes), rowAnimation)
}

func (t_ TableView) UnhideRowsAtIndexes_WithAnimation(indexes foundation.IIndexSet, rowAnimation TableViewAnimationOptions) {
	objc.CallMethod[objc.Void](t_, objc.SEL("unhideRowsAtIndexes:withAnimation:"), objc.ExtractPtr(indexes), rowAnimation)
}

// deprecated
func (t_ TableView) DragImageForRows_Event_DragImageOffset(dragRows []objc.IObject, dragEvent IEvent, dragImageOffset *foundation.Point) Image {
	rv := objc.CallMethod[Image](t_, objc.SEL("dragImageForRows:event:dragImageOffset:"), dragRows, objc.ExtractPtr(dragEvent), dragImageOffset)
	return rv
}

// deprecated
func (t_ TableView) SetAutoresizesAllColumnsToFit(flag bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAutoresizesAllColumnsToFit:"), flag)
}

// deprecated
func (t_ TableView) AutoresizesAllColumnsToFit() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("autoresizesAllColumnsToFit"))
	return rv
}

// deprecated
func (t_ TableView) SelectColumn_ByExtendingSelection(column int, extend bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("selectColumn:byExtendingSelection:"), column, extend)
}

// deprecated
func (t_ TableView) SelectRow_ByExtendingSelection(row int, extend bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("selectRow:byExtendingSelection:"), row, extend)
}

// deprecated
func (t_ TableView) TableView_WriteRows_ToPasteboard(tableView ITableView, rows []objc.IObject, pboard IPasteboard) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("tableView:writeRows:toPasteboard:"), objc.ExtractPtr(tableView), rows, objc.ExtractPtr(pboard))
	return rv
}

// deprecated
func (t_ TableView) SetDrawsGrid(flag bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDrawsGrid:"), flag)
}

// deprecated
func (t_ TableView) DrawsGrid() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("drawsGrid"))
	return rv
}

// deprecated
func (t_ TableView) SelectedColumnEnumerator() foundation.Enumerator {
	rv := objc.CallMethod[foundation.Enumerator](t_, objc.SEL("selectedColumnEnumerator"))
	return rv
}

// deprecated
func (t_ TableView) SelectedRowEnumerator() foundation.Enumerator {
	rv := objc.CallMethod[foundation.Enumerator](t_, objc.SEL("selectedRowEnumerator"))
	return rv
}

// deprecated
func (t_ TableView) FocusedColumn() int {
	rv := objc.CallMethod[int](t_, objc.SEL("focusedColumn"))
	return rv
}

// deprecated
func (t_ TableView) SetFocusedColumn(focusedColumn int) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setFocusedColumn:"), focusedColumn)
}

// deprecated
func (t_ TableView) ShouldFocusCell_AtColumn_Row(cell ICell, column int, row int) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("shouldFocusCell:atColumn:row:"), objc.ExtractPtr(cell), column, row)
	return rv
}

// deprecated
func (t_ TableView) PerformClickOnCellAtColumn_Row(column int, row int) {
	objc.CallMethod[objc.Void](t_, objc.SEL("performClickOnCellAtColumn:row:"), column, row)
}

// deprecated
func (t_ TableView) PreparedCellAtColumn_Row(column int, row int) Cell {
	rv := objc.CallMethod[Cell](t_, objc.SEL("preparedCellAtColumn:row:"), column, row)
	return rv
}

// deprecated
func (t_ TableView) ColumnsInRect(rect foundation.Rect) foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.SEL("columnsInRect:"), rect)
	return rv
}

// deprecated
func (t_ TableView) TextShouldBeginEditing(textObject IText) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("textShouldBeginEditing:"), objc.ExtractPtr(textObject))
	return rv
}

// deprecated
func (t_ TableView) TextDidBeginEditing(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.SEL("textDidBeginEditing:"), objc.ExtractPtr(notification))
}

// deprecated
func (t_ TableView) TextDidChange(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.SEL("textDidChange:"), objc.ExtractPtr(notification))
}

// deprecated
func (t_ TableView) TextShouldEndEditing(textObject IText) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("textShouldEndEditing:"), objc.ExtractPtr(textObject))
	return rv
}

// deprecated
func (t_ TableView) TextDidEndEditing(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.SEL("textDidEndEditing:"), objc.ExtractPtr(notification))
}

// weak property
func (t_ TableView) DataSource() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.SEL("dataSource"))
	return rv
}

// weak property
func (t_ TableView) SetDataSource(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDataSource:"), objc.ExtractPtr(value))
}

func (t_ TableView) UsesStaticContents() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("usesStaticContents"))
	return rv
}

func (t_ TableView) SetUsesStaticContents(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setUsesStaticContents:"), value)
}

func (t_ TableView) RegisteredNibsByIdentifier() map[UserInterfaceItemIdentifier]Nib {
	rv := objc.CallMethod[map[UserInterfaceItemIdentifier]Nib](t_, objc.SEL("registeredNibsByIdentifier"))
	return rv
}

func (t_ TableView) DoubleAction() objc.Selector {
	rv := objc.CallMethod[objc.Selector](t_, objc.SEL("doubleAction"))
	return rv
}

func (t_ TableView) SetDoubleAction(value objc.Selector) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDoubleAction:"), value)
}

func (t_ TableView) ClickedColumn() int {
	rv := objc.CallMethod[int](t_, objc.SEL("clickedColumn"))
	return rv
}

func (t_ TableView) ClickedRow() int {
	rv := objc.CallMethod[int](t_, objc.SEL("clickedRow"))
	return rv
}

func (t_ TableView) AllowsColumnReordering() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("allowsColumnReordering"))
	return rv
}

func (t_ TableView) SetAllowsColumnReordering(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAllowsColumnReordering:"), value)
}

func (t_ TableView) AllowsColumnResizing() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("allowsColumnResizing"))
	return rv
}

func (t_ TableView) SetAllowsColumnResizing(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAllowsColumnResizing:"), value)
}

func (t_ TableView) AllowsMultipleSelection() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("allowsMultipleSelection"))
	return rv
}

func (t_ TableView) SetAllowsMultipleSelection(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAllowsMultipleSelection:"), value)
}

func (t_ TableView) AllowsEmptySelection() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("allowsEmptySelection"))
	return rv
}

func (t_ TableView) SetAllowsEmptySelection(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAllowsEmptySelection:"), value)
}

func (t_ TableView) AllowsColumnSelection() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("allowsColumnSelection"))
	return rv
}

func (t_ TableView) SetAllowsColumnSelection(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAllowsColumnSelection:"), value)
}

func (t_ TableView) UsesAutomaticRowHeights() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("usesAutomaticRowHeights"))
	return rv
}

func (t_ TableView) SetUsesAutomaticRowHeights(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setUsesAutomaticRowHeights:"), value)
}

func (t_ TableView) IntercellSpacing() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.SEL("intercellSpacing"))
	return rv
}

func (t_ TableView) SetIntercellSpacing(value foundation.Size) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setIntercellSpacing:"), value)
}

func (t_ TableView) RowHeight() float64 {
	rv := objc.CallMethod[float64](t_, objc.SEL("rowHeight"))
	return rv
}

func (t_ TableView) SetRowHeight(value float64) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setRowHeight:"), value)
}

func (t_ TableView) BackgroundColor() Color {
	rv := objc.CallMethod[Color](t_, objc.SEL("backgroundColor"))
	return rv
}

func (t_ TableView) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (t_ TableView) UsesAlternatingRowBackgroundColors() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("usesAlternatingRowBackgroundColors"))
	return rv
}

func (t_ TableView) SetUsesAlternatingRowBackgroundColors(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setUsesAlternatingRowBackgroundColors:"), value)
}

func (t_ TableView) Style() TableViewStyle {
	rv := objc.CallMethod[TableViewStyle](t_, objc.SEL("style"))
	return rv
}

func (t_ TableView) SetStyle(value TableViewStyle) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setStyle:"), value)
}

func (t_ TableView) EffectiveStyle() TableViewStyle {
	rv := objc.CallMethod[TableViewStyle](t_, objc.SEL("effectiveStyle"))
	return rv
}

func (t_ TableView) SelectionHighlightStyle() TableViewSelectionHighlightStyle {
	rv := objc.CallMethod[TableViewSelectionHighlightStyle](t_, objc.SEL("selectionHighlightStyle"))
	return rv
}

func (t_ TableView) SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSelectionHighlightStyle:"), value)
}

func (t_ TableView) GridColor() Color {
	rv := objc.CallMethod[Color](t_, objc.SEL("gridColor"))
	return rv
}

func (t_ TableView) SetGridColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setGridColor:"), objc.ExtractPtr(value))
}

func (t_ TableView) GridStyleMask() TableViewGridLineStyle {
	rv := objc.CallMethod[TableViewGridLineStyle](t_, objc.SEL("gridStyleMask"))
	return rv
}

func (t_ TableView) SetGridStyleMask(value TableViewGridLineStyle) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setGridStyleMask:"), value)
}

func (t_ TableView) EffectiveRowSizeStyle() TableViewRowSizeStyle {
	rv := objc.CallMethod[TableViewRowSizeStyle](t_, objc.SEL("effectiveRowSizeStyle"))
	return rv
}

func (t_ TableView) RowSizeStyle() TableViewRowSizeStyle {
	rv := objc.CallMethod[TableViewRowSizeStyle](t_, objc.SEL("rowSizeStyle"))
	return rv
}

func (t_ TableView) SetRowSizeStyle(value TableViewRowSizeStyle) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setRowSizeStyle:"), value)
}

func (t_ TableView) TableColumns() []TableColumn {
	rv := objc.CallMethod[[]TableColumn](t_, objc.SEL("tableColumns"))
	return rv
}

func (t_ TableView) SelectedColumn() int {
	rv := objc.CallMethod[int](t_, objc.SEL("selectedColumn"))
	return rv
}

func (t_ TableView) SelectedColumnIndexes() foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](t_, objc.SEL("selectedColumnIndexes"))
	return rv
}

func (t_ TableView) NumberOfSelectedColumns() int {
	rv := objc.CallMethod[int](t_, objc.SEL("numberOfSelectedColumns"))
	return rv
}

func (t_ TableView) SelectedRow() int {
	rv := objc.CallMethod[int](t_, objc.SEL("selectedRow"))
	return rv
}

func (t_ TableView) SelectedRowIndexes() foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](t_, objc.SEL("selectedRowIndexes"))
	return rv
}

func (t_ TableView) NumberOfSelectedRows() int {
	rv := objc.CallMethod[int](t_, objc.SEL("numberOfSelectedRows"))
	return rv
}

func (t_ TableView) AllowsTypeSelect() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("allowsTypeSelect"))
	return rv
}

func (t_ TableView) SetAllowsTypeSelect(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAllowsTypeSelect:"), value)
}

func (t_ TableView) NumberOfColumns() int {
	rv := objc.CallMethod[int](t_, objc.SEL("numberOfColumns"))
	return rv
}

func (t_ TableView) NumberOfRows() int {
	rv := objc.CallMethod[int](t_, objc.SEL("numberOfRows"))
	return rv
}

func (t_ TableView) FloatsGroupRows() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("floatsGroupRows"))
	return rv
}

func (t_ TableView) SetFloatsGroupRows(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setFloatsGroupRows:"), value)
}

func (t_ TableView) EditedColumn() int {
	rv := objc.CallMethod[int](t_, objc.SEL("editedColumn"))
	return rv
}

func (t_ TableView) EditedRow() int {
	rv := objc.CallMethod[int](t_, objc.SEL("editedRow"))
	return rv
}

func (t_ TableView) HeaderView() TableHeaderView {
	rv := objc.CallMethod[TableHeaderView](t_, objc.SEL("headerView"))
	return rv
}

func (t_ TableView) SetHeaderView(value ITableHeaderView) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setHeaderView:"), objc.ExtractPtr(value))
}

func (t_ TableView) CornerView() View {
	rv := objc.CallMethod[View](t_, objc.SEL("cornerView"))
	return rv
}

func (t_ TableView) SetCornerView(value IView) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setCornerView:"), objc.ExtractPtr(value))
}

func (t_ TableView) ColumnAutoresizingStyle() TableViewColumnAutoresizingStyle {
	rv := objc.CallMethod[TableViewColumnAutoresizingStyle](t_, objc.SEL("columnAutoresizingStyle"))
	return rv
}

func (t_ TableView) SetColumnAutoresizingStyle(value TableViewColumnAutoresizingStyle) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setColumnAutoresizingStyle:"), value)
}

func (t_ TableView) AutosaveTableColumns() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("autosaveTableColumns"))
	return rv
}

func (t_ TableView) SetAutosaveTableColumns(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAutosaveTableColumns:"), value)
}

func (t_ TableView) AutosaveName() TableViewAutosaveName {
	rv := objc.CallMethod[TableViewAutosaveName](t_, objc.SEL("autosaveName"))
	return rv
}

func (t_ TableView) SetAutosaveName(value TableViewAutosaveName) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAutosaveName:"), value)
}

// weak property
func (t_ TableView) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.SEL("delegate"))
	return rv
}

// weak property
func (t_ TableView) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}

// weak property
func (t_ TableView) HighlightedTableColumn() TableColumn {
	rv := objc.CallMethod[TableColumn](t_, objc.SEL("highlightedTableColumn"))
	return rv
}

// weak property
func (t_ TableView) SetHighlightedTableColumn(value ITableColumn) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setHighlightedTableColumn:"), objc.ExtractPtr(value))
}

func (t_ TableView) VerticalMotionCanBeginDrag() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("verticalMotionCanBeginDrag"))
	return rv
}

func (t_ TableView) SetVerticalMotionCanBeginDrag(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setVerticalMotionCanBeginDrag:"), value)
}

func (t_ TableView) DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle {
	rv := objc.CallMethod[TableViewDraggingDestinationFeedbackStyle](t_, objc.SEL("draggingDestinationFeedbackStyle"))
	return rv
}

func (t_ TableView) SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDraggingDestinationFeedbackStyle:"), value)
}

func (t_ TableView) SortDescriptors() []foundation.SortDescriptor {
	rv := objc.CallMethod[[]foundation.SortDescriptor](t_, objc.SEL("sortDescriptors"))
	return rv
}

func (t_ TableView) SetSortDescriptors(value []foundation.ISortDescriptor) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSortDescriptors:"), value)
}

func (t_ TableView) RowActionsVisible() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("rowActionsVisible"))
	return rv
}

func (t_ TableView) SetRowActionsVisible(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setRowActionsVisible:"), value)
}

func (t_ TableView) HiddenRowIndexes() foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](t_, objc.SEL("hiddenRowIndexes"))
	return rv
}
