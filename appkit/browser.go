// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var BrowserClass = _BrowserClass{objc.GetClass("NSBrowser")}

type _BrowserClass struct {
	objc.Class
}

type IBrowser interface {
	IControl
	Tile()
	SelectedRowIndexesInColumn(column int) foundation.IndexSet
	SelectRowIndexes_InColumn(indexes foundation.IIndexSet, column int)
	SelectedCellInColumn(column int) objc.Object
	SelectAll(sender objc.IObject)
	SelectedRowInColumn(column int) int
	SelectRow_InColumn(row int, column int)
	LoadedCellAtRow_Column(row int, col int) objc.Object
	EditItemAtIndexPath_WithEvent_Select(indexPath foundation.IIndexPath, event IEvent, select_ bool)
	ItemAtIndexPath(indexPath foundation.IIndexPath) objc.Object
	ItemAtRow_InColumn(row int, column int) objc.Object
	IndexPathForColumn(column int) foundation.IndexPath
	IsLeafItem(item objc.IObject) bool
	ParentForItemsInColumn(column int) objc.Object
	Path() string
	SetPath(path string) bool
	PathToColumn(column int) string
	AddColumn()
	ValidateVisibleColumns()
	LoadColumnZero()
	ReloadColumn(column int)
	TitleOfColumn(column int) string
	SetTitle_OfColumn(string_ string, column int)
	DrawTitleOfColumn_InRect(column int, rect foundation.Rect)
	TitleFrameOfColumn(column int) foundation.Rect
	NoteHeightOfRowsWithIndexesChanged_InColumn(indexSet foundation.IIndexSet, columnIndex int)
	ReloadDataForRowIndexes_InColumn(rowIndexes foundation.IIndexSet, column int)
	ScrollColumnToVisible(column int)
	ScrollColumnsLeftBy(shiftAmount int)
	ScrollColumnsRightBy(shiftAmount int)
	ScrollRowToVisible_InColumn(row int, column int)
	SetDraggingSourceOperationMask_ForLocal(mask DragOperation, isLocal bool)
	CanDragRowsWithIndexes_InColumn_WithEvent(rowIndexes foundation.IIndexSet, column int, event IEvent) bool
	DraggingImageForRowsWithIndexes_InColumn_WithEvent_Offset(rowIndexes foundation.IIndexSet, column int, event IEvent, dragImageOffset *foundation.Point) Image
	FrameOfColumn(column int) foundation.Rect
	FrameOfInsideOfColumn(column int) foundation.Rect
	FrameOfRow_InColumn(row int, column int) foundation.Rect
	GetRow_Column_ForPoint(row *int, column *int, point foundation.Point) bool
	SendAction() bool
	DoClick(sender objc.IObject)
	DoDoubleClick(sender objc.IObject)
	ColumnContentWidthForColumnWidth(columnWidth float64) float64
	ColumnWidthForColumnContentWidth(columnContentWidth float64) float64
	WidthOfColumn(column int) float64
	SetWidth_OfColumn(columnWidth float64, columnIndex int)
	DefaultColumnWidth() float64
	SetDefaultColumnWidth(columnWidth float64)
	// deprecated
	UpdateScroller()
	// deprecated
	ScrollViaScroller(sender IScroller)
	// deprecated
	DisplayAllColumns()
	// deprecated
	DisplayColumn(column int)
	// deprecated
	ColumnOfMatrix(matrix IMatrix) int
	// deprecated
	MatrixInColumn(column int) Matrix
	// deprecated
	MatrixClass() objc.Class
	// deprecated
	SetMatrixClass(factoryId objc.IClass)
	// deprecated
	AcceptsArrowKeys() bool
	// deprecated
	SetAcceptsArrowKeys(flag bool)
	ReusesColumns() bool
	SetReusesColumns(value bool)
	MaxVisibleColumns() int
	SetMaxVisibleColumns(value int)
	AutohidesScroller() bool
	SetAutohidesScroller(value bool)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	MinColumnWidth() float64
	SetMinColumnWidth(value float64)
	SeparatesColumns() bool
	SetSeparatesColumns(value bool)
	TakesTitleFromPreviousColumn() bool
	SetTakesTitleFromPreviousColumn(value bool)
	Delegate() BrowserDelegateWrapper
	SetDelegate(value BrowserDelegate)
	CellPrototype() objc.Object
	SetCellPrototype(value objc.IObject)
	AllowsBranchSelection() bool
	SetAllowsBranchSelection(value bool)
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	AllowsTypeSelect() bool
	SetAllowsTypeSelect(value bool)
	SelectedCells() []Cell
	SelectionIndexPath() foundation.IndexPath
	SetSelectionIndexPath(value foundation.IIndexPath)
	SelectionIndexPaths() []foundation.IndexPath
	SetSelectionIndexPaths(value []foundation.IIndexPath)
	PathSeparator() string
	SetPathSeparator(value string)
	SelectedColumn() int
	LastColumn() int
	SetLastColumn(value int)
	FirstVisibleColumn() int
	NumberOfVisibleColumns() int
	LastVisibleColumn() int
	IsLoaded() bool
	IsTitled() bool
	SetTitled(value bool)
	TitleHeight() float64
	HasHorizontalScroller() bool
	SetHasHorizontalScroller(value bool)
	DoubleAction() objc.Selector
	SetDoubleAction(value objc.Selector)
	SendsActionOnArrowKeys() bool
	SetSendsActionOnArrowKeys(value bool)
	ClickedColumn() int
	ClickedRow() int
	ColumnsAutosaveName() BrowserColumnsAutosaveName
	SetColumnsAutosaveName(value BrowserColumnsAutosaveName)
	ColumnResizingType() BrowserColumnResizingType
	SetColumnResizingType(value BrowserColumnResizingType)
	PrefersAllColumnUserResizing() bool
	SetPrefersAllColumnUserResizing(value bool)
	RowHeight() float64
	SetRowHeight(value float64)
}

type Browser struct {
	Control
}

func MakeBrowser(ptr unsafe.Pointer) Browser {
	return Browser{
		Control: MakeControl(ptr),
	}
}

func (b_ Browser) InitWithFrame(frameRect foundation.Rect) Browser {
	rv := ffi.CallMethod[Browser](b_, "initWithFrame:", frameRect)
	return rv
}

func (b_ Browser) Init() Browser {
	rv := ffi.CallMethod[Browser](b_, "init")
	return rv
}

func (bc _BrowserClass) Alloc() Browser {
	rv := ffi.CallMethod[Browser](bc, "alloc")
	return rv
}

func (bc _BrowserClass) New() Browser {
	rv := ffi.CallMethod[Browser](bc, "new")
	rv.Autorelease()
	return rv
}

func NewBrowser() Browser {
	return BrowserClass.New()
}

func (b_ Browser) Tile() {
	ffi.CallMethod[ffi.Void](b_, "tile")
}

func (b_ Browser) SelectedRowIndexesInColumn(column int) foundation.IndexSet {
	rv := ffi.CallMethod[foundation.IndexSet](b_, "selectedRowIndexesInColumn:", column)
	return rv
}

func (b_ Browser) SelectRowIndexes_InColumn(indexes foundation.IIndexSet, column int) {
	ffi.CallMethod[ffi.Void](b_, "selectRowIndexes:inColumn:", indexes, column)
}

func (b_ Browser) SelectedCellInColumn(column int) objc.Object {
	rv := ffi.CallMethod[objc.Object](b_, "selectedCellInColumn:", column)
	return rv
}

func (b_ Browser) SelectAll(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](b_, "selectAll:", sender)
}

func (b_ Browser) SelectedRowInColumn(column int) int {
	rv := ffi.CallMethod[int](b_, "selectedRowInColumn:", column)
	return rv
}

func (b_ Browser) SelectRow_InColumn(row int, column int) {
	ffi.CallMethod[ffi.Void](b_, "selectRow:inColumn:", row, column)
}

func (b_ Browser) LoadedCellAtRow_Column(row int, col int) objc.Object {
	rv := ffi.CallMethod[objc.Object](b_, "loadedCellAtRow:column:", row, col)
	return rv
}

func (b_ Browser) EditItemAtIndexPath_WithEvent_Select(indexPath foundation.IIndexPath, event IEvent, select_ bool) {
	ffi.CallMethod[ffi.Void](b_, "editItemAtIndexPath:withEvent:select:", indexPath, event, select_)
}

func (b_ Browser) ItemAtIndexPath(indexPath foundation.IIndexPath) objc.Object {
	rv := ffi.CallMethod[objc.Object](b_, "itemAtIndexPath:", indexPath)
	return rv
}

func (b_ Browser) ItemAtRow_InColumn(row int, column int) objc.Object {
	rv := ffi.CallMethod[objc.Object](b_, "itemAtRow:inColumn:", row, column)
	return rv
}

func (b_ Browser) IndexPathForColumn(column int) foundation.IndexPath {
	rv := ffi.CallMethod[foundation.IndexPath](b_, "indexPathForColumn:", column)
	return rv
}

func (b_ Browser) IsLeafItem(item objc.IObject) bool {
	rv := ffi.CallMethod[bool](b_, "isLeafItem:", item)
	return rv
}

func (b_ Browser) ParentForItemsInColumn(column int) objc.Object {
	rv := ffi.CallMethod[objc.Object](b_, "parentForItemsInColumn:", column)
	return rv
}

func (b_ Browser) Path() string {
	rv := ffi.CallMethod[string](b_, "path")
	return rv
}

func (b_ Browser) SetPath(path string) bool {
	rv := ffi.CallMethod[bool](b_, "setPath:", path)
	return rv
}

func (b_ Browser) PathToColumn(column int) string {
	rv := ffi.CallMethod[string](b_, "pathToColumn:", column)
	return rv
}

func (b_ Browser) AddColumn() {
	ffi.CallMethod[ffi.Void](b_, "addColumn")
}

func (b_ Browser) ValidateVisibleColumns() {
	ffi.CallMethod[ffi.Void](b_, "validateVisibleColumns")
}

func (b_ Browser) LoadColumnZero() {
	ffi.CallMethod[ffi.Void](b_, "loadColumnZero")
}

func (b_ Browser) ReloadColumn(column int) {
	ffi.CallMethod[ffi.Void](b_, "reloadColumn:", column)
}

func (b_ Browser) TitleOfColumn(column int) string {
	rv := ffi.CallMethod[string](b_, "titleOfColumn:", column)
	return rv
}

func (b_ Browser) SetTitle_OfColumn(string_ string, column int) {
	ffi.CallMethod[ffi.Void](b_, "setTitle:ofColumn:", string_, column)
}

func (b_ Browser) DrawTitleOfColumn_InRect(column int, rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](b_, "drawTitleOfColumn:inRect:", column, rect)
}

func (b_ Browser) TitleFrameOfColumn(column int) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](b_, "titleFrameOfColumn:", column)
	return rv
}

func (b_ Browser) NoteHeightOfRowsWithIndexesChanged_InColumn(indexSet foundation.IIndexSet, columnIndex int) {
	ffi.CallMethod[ffi.Void](b_, "noteHeightOfRowsWithIndexesChanged:inColumn:", indexSet, columnIndex)
}

func (b_ Browser) ReloadDataForRowIndexes_InColumn(rowIndexes foundation.IIndexSet, column int) {
	ffi.CallMethod[ffi.Void](b_, "reloadDataForRowIndexes:inColumn:", rowIndexes, column)
}

func (b_ Browser) ScrollColumnToVisible(column int) {
	ffi.CallMethod[ffi.Void](b_, "scrollColumnToVisible:", column)
}

func (b_ Browser) ScrollColumnsLeftBy(shiftAmount int) {
	ffi.CallMethod[ffi.Void](b_, "scrollColumnsLeftBy:", shiftAmount)
}

func (b_ Browser) ScrollColumnsRightBy(shiftAmount int) {
	ffi.CallMethod[ffi.Void](b_, "scrollColumnsRightBy:", shiftAmount)
}

func (b_ Browser) ScrollRowToVisible_InColumn(row int, column int) {
	ffi.CallMethod[ffi.Void](b_, "scrollRowToVisible:inColumn:", row, column)
}

func (b_ Browser) SetDraggingSourceOperationMask_ForLocal(mask DragOperation, isLocal bool) {
	ffi.CallMethod[ffi.Void](b_, "setDraggingSourceOperationMask:forLocal:", mask, isLocal)
}

func (b_ Browser) CanDragRowsWithIndexes_InColumn_WithEvent(rowIndexes foundation.IIndexSet, column int, event IEvent) bool {
	rv := ffi.CallMethod[bool](b_, "canDragRowsWithIndexes:inColumn:withEvent:", rowIndexes, column, event)
	return rv
}

func (b_ Browser) DraggingImageForRowsWithIndexes_InColumn_WithEvent_Offset(rowIndexes foundation.IIndexSet, column int, event IEvent, dragImageOffset *foundation.Point) Image {
	rv := ffi.CallMethod[Image](b_, "draggingImageForRowsWithIndexes:inColumn:withEvent:offset:", rowIndexes, column, event, dragImageOffset)
	return rv
}

func (b_ Browser) FrameOfColumn(column int) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](b_, "frameOfColumn:", column)
	return rv
}

func (b_ Browser) FrameOfInsideOfColumn(column int) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](b_, "frameOfInsideOfColumn:", column)
	return rv
}

func (b_ Browser) FrameOfRow_InColumn(row int, column int) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](b_, "frameOfRow:inColumn:", row, column)
	return rv
}

func (b_ Browser) GetRow_Column_ForPoint(row *int, column *int, point foundation.Point) bool {
	rv := ffi.CallMethod[bool](b_, "getRow:column:forPoint:", row, column, point)
	return rv
}

func (b_ Browser) SendAction() bool {
	rv := ffi.CallMethod[bool](b_, "sendAction")
	return rv
}

func (b_ Browser) DoClick(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](b_, "doClick:", sender)
}

func (b_ Browser) DoDoubleClick(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](b_, "doDoubleClick:", sender)
}

func (bc _BrowserClass) RemoveSavedColumnsWithAutosaveName(name BrowserColumnsAutosaveName) {
	ffi.CallMethod[ffi.Void](bc, "removeSavedColumnsWithAutosaveName:", name)
}

func (b_ Browser) ColumnContentWidthForColumnWidth(columnWidth float64) float64 {
	rv := ffi.CallMethod[float64](b_, "columnContentWidthForColumnWidth:", columnWidth)
	return rv
}

func (b_ Browser) ColumnWidthForColumnContentWidth(columnContentWidth float64) float64 {
	rv := ffi.CallMethod[float64](b_, "columnWidthForColumnContentWidth:", columnContentWidth)
	return rv
}

func (b_ Browser) WidthOfColumn(column int) float64 {
	rv := ffi.CallMethod[float64](b_, "widthOfColumn:", column)
	return rv
}

func (b_ Browser) SetWidth_OfColumn(columnWidth float64, columnIndex int) {
	ffi.CallMethod[ffi.Void](b_, "setWidth:ofColumn:", columnWidth, columnIndex)
}

func (b_ Browser) DefaultColumnWidth() float64 {
	rv := ffi.CallMethod[float64](b_, "defaultColumnWidth")
	return rv
}

func (b_ Browser) SetDefaultColumnWidth(columnWidth float64) {
	ffi.CallMethod[ffi.Void](b_, "setDefaultColumnWidth:", columnWidth)
}

// deprecated
func (b_ Browser) UpdateScroller() {
	ffi.CallMethod[ffi.Void](b_, "updateScroller")
}

// deprecated
func (b_ Browser) ScrollViaScroller(sender IScroller) {
	ffi.CallMethod[ffi.Void](b_, "scrollViaScroller:", sender)
}

// deprecated
func (b_ Browser) DisplayAllColumns() {
	ffi.CallMethod[ffi.Void](b_, "displayAllColumns")
}

// deprecated
func (b_ Browser) DisplayColumn(column int) {
	ffi.CallMethod[ffi.Void](b_, "displayColumn:", column)
}

// deprecated
func (b_ Browser) ColumnOfMatrix(matrix IMatrix) int {
	rv := ffi.CallMethod[int](b_, "columnOfMatrix:", matrix)
	return rv
}

// deprecated
func (b_ Browser) MatrixInColumn(column int) Matrix {
	rv := ffi.CallMethod[Matrix](b_, "matrixInColumn:", column)
	return rv
}

// deprecated
func (b_ Browser) MatrixClass() objc.Class {
	rv := ffi.CallMethod[objc.Class](b_, "matrixClass")
	return rv
}

// deprecated
func (b_ Browser) SetMatrixClass(factoryId objc.IClass) {
	ffi.CallMethod[ffi.Void](b_, "setMatrixClass:", factoryId)
}

// deprecated
func (b_ Browser) AcceptsArrowKeys() bool {
	rv := ffi.CallMethod[bool](b_, "acceptsArrowKeys")
	return rv
}

// deprecated
func (b_ Browser) SetAcceptsArrowKeys(flag bool) {
	ffi.CallMethod[ffi.Void](b_, "setAcceptsArrowKeys:", flag)
}

func (b_ Browser) ReusesColumns() bool {
	rv := ffi.CallMethod[bool](b_, "reusesColumns")
	return rv
}

func (b_ Browser) SetReusesColumns(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setReusesColumns:", value)
}

func (b_ Browser) MaxVisibleColumns() int {
	rv := ffi.CallMethod[int](b_, "maxVisibleColumns")
	return rv
}

func (b_ Browser) SetMaxVisibleColumns(value int) {
	ffi.CallMethod[ffi.Void](b_, "setMaxVisibleColumns:", value)
}

func (b_ Browser) AutohidesScroller() bool {
	rv := ffi.CallMethod[bool](b_, "autohidesScroller")
	return rv
}

func (b_ Browser) SetAutohidesScroller(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setAutohidesScroller:", value)
}

func (b_ Browser) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](b_, "backgroundColor")
	return rv
}

func (b_ Browser) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](b_, "setBackgroundColor:", value)
}

func (b_ Browser) MinColumnWidth() float64 {
	rv := ffi.CallMethod[float64](b_, "minColumnWidth")
	return rv
}

func (b_ Browser) SetMinColumnWidth(value float64) {
	ffi.CallMethod[ffi.Void](b_, "setMinColumnWidth:", value)
}

func (b_ Browser) SeparatesColumns() bool {
	rv := ffi.CallMethod[bool](b_, "separatesColumns")
	return rv
}

func (b_ Browser) SetSeparatesColumns(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setSeparatesColumns:", value)
}

func (b_ Browser) TakesTitleFromPreviousColumn() bool {
	rv := ffi.CallMethod[bool](b_, "takesTitleFromPreviousColumn")
	return rv
}

func (b_ Browser) SetTakesTitleFromPreviousColumn(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setTakesTitleFromPreviousColumn:", value)
}

func (b_ Browser) Delegate() BrowserDelegateWrapper {
	rv := ffi.CallMethod[BrowserDelegateWrapper](b_, "delegate")
	return rv
}

func (b_ Browser) SetDelegate(value BrowserDelegate) {
	po := ffi.CreateProtocol("NSBrowserDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(b_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](b_, "setDelegate:", po)
}

func (b_ Browser) CellPrototype() objc.Object {
	rv := ffi.CallMethod[objc.Object](b_, "cellPrototype")
	return rv
}

func (b_ Browser) SetCellPrototype(value objc.IObject) {
	ffi.CallMethod[ffi.Void](b_, "setCellPrototype:", value)
}

func (b_ Browser) AllowsBranchSelection() bool {
	rv := ffi.CallMethod[bool](b_, "allowsBranchSelection")
	return rv
}

func (b_ Browser) SetAllowsBranchSelection(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setAllowsBranchSelection:", value)
}

func (b_ Browser) AllowsEmptySelection() bool {
	rv := ffi.CallMethod[bool](b_, "allowsEmptySelection")
	return rv
}

func (b_ Browser) SetAllowsEmptySelection(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setAllowsEmptySelection:", value)
}

func (b_ Browser) AllowsMultipleSelection() bool {
	rv := ffi.CallMethod[bool](b_, "allowsMultipleSelection")
	return rv
}

func (b_ Browser) SetAllowsMultipleSelection(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setAllowsMultipleSelection:", value)
}

func (b_ Browser) AllowsTypeSelect() bool {
	rv := ffi.CallMethod[bool](b_, "allowsTypeSelect")
	return rv
}

func (b_ Browser) SetAllowsTypeSelect(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setAllowsTypeSelect:", value)
}

func (b_ Browser) SelectedCells() []Cell {
	rv := ffi.CallMethod[[]Cell](b_, "selectedCells")
	return rv
}

func (b_ Browser) SelectionIndexPath() foundation.IndexPath {
	rv := ffi.CallMethod[foundation.IndexPath](b_, "selectionIndexPath")
	return rv
}

func (b_ Browser) SetSelectionIndexPath(value foundation.IIndexPath) {
	ffi.CallMethod[ffi.Void](b_, "setSelectionIndexPath:", value)
}

func (b_ Browser) SelectionIndexPaths() []foundation.IndexPath {
	rv := ffi.CallMethod[[]foundation.IndexPath](b_, "selectionIndexPaths")
	return rv
}

func (b_ Browser) SetSelectionIndexPaths(value []foundation.IIndexPath) {
	ffi.CallMethod[ffi.Void](b_, "setSelectionIndexPaths:", value)
}

func (b_ Browser) PathSeparator() string {
	rv := ffi.CallMethod[string](b_, "pathSeparator")
	return rv
}

func (b_ Browser) SetPathSeparator(value string) {
	ffi.CallMethod[ffi.Void](b_, "setPathSeparator:", value)
}

func (b_ Browser) SelectedColumn() int {
	rv := ffi.CallMethod[int](b_, "selectedColumn")
	return rv
}

func (b_ Browser) LastColumn() int {
	rv := ffi.CallMethod[int](b_, "lastColumn")
	return rv
}

func (b_ Browser) SetLastColumn(value int) {
	ffi.CallMethod[ffi.Void](b_, "setLastColumn:", value)
}

func (b_ Browser) FirstVisibleColumn() int {
	rv := ffi.CallMethod[int](b_, "firstVisibleColumn")
	return rv
}

func (b_ Browser) NumberOfVisibleColumns() int {
	rv := ffi.CallMethod[int](b_, "numberOfVisibleColumns")
	return rv
}

func (b_ Browser) LastVisibleColumn() int {
	rv := ffi.CallMethod[int](b_, "lastVisibleColumn")
	return rv
}

func (b_ Browser) IsLoaded() bool {
	rv := ffi.CallMethod[bool](b_, "isLoaded")
	return rv
}

func (b_ Browser) IsTitled() bool {
	rv := ffi.CallMethod[bool](b_, "isTitled")
	return rv
}

func (b_ Browser) SetTitled(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setTitled:", value)
}

func (b_ Browser) TitleHeight() float64 {
	rv := ffi.CallMethod[float64](b_, "titleHeight")
	return rv
}

func (b_ Browser) HasHorizontalScroller() bool {
	rv := ffi.CallMethod[bool](b_, "hasHorizontalScroller")
	return rv
}

func (b_ Browser) SetHasHorizontalScroller(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setHasHorizontalScroller:", value)
}

func (b_ Browser) DoubleAction() objc.Selector {
	rv := ffi.CallMethod[objc.Selector](b_, "doubleAction")
	return rv
}

func (b_ Browser) SetDoubleAction(value objc.Selector) {
	ffi.CallMethod[ffi.Void](b_, "setDoubleAction:", value)
}

func (b_ Browser) SendsActionOnArrowKeys() bool {
	rv := ffi.CallMethod[bool](b_, "sendsActionOnArrowKeys")
	return rv
}

func (b_ Browser) SetSendsActionOnArrowKeys(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setSendsActionOnArrowKeys:", value)
}

func (b_ Browser) ClickedColumn() int {
	rv := ffi.CallMethod[int](b_, "clickedColumn")
	return rv
}

func (b_ Browser) ClickedRow() int {
	rv := ffi.CallMethod[int](b_, "clickedRow")
	return rv
}

func (b_ Browser) ColumnsAutosaveName() BrowserColumnsAutosaveName {
	rv := ffi.CallMethod[BrowserColumnsAutosaveName](b_, "columnsAutosaveName")
	return rv
}

func (b_ Browser) SetColumnsAutosaveName(value BrowserColumnsAutosaveName) {
	ffi.CallMethod[ffi.Void](b_, "setColumnsAutosaveName:", value)
}

func (b_ Browser) ColumnResizingType() BrowserColumnResizingType {
	rv := ffi.CallMethod[BrowserColumnResizingType](b_, "columnResizingType")
	return rv
}

func (b_ Browser) SetColumnResizingType(value BrowserColumnResizingType) {
	ffi.CallMethod[ffi.Void](b_, "setColumnResizingType:", value)
}

func (b_ Browser) PrefersAllColumnUserResizing() bool {
	rv := ffi.CallMethod[bool](b_, "prefersAllColumnUserResizing")
	return rv
}

func (b_ Browser) SetPrefersAllColumnUserResizing(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setPrefersAllColumnUserResizing:", value)
}

func (b_ Browser) RowHeight() float64 {
	rv := ffi.CallMethod[float64](b_, "rowHeight")
	return rv
}

func (b_ Browser) SetRowHeight(value float64) {
	ffi.CallMethod[ffi.Void](b_, "setRowHeight:", value)
}
