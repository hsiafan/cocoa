// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
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
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
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
	rv := objc.CallMethod[Browser](b_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func (b_ Browser) Init() Browser {
	rv := objc.CallMethod[Browser](b_, objc.GetSelector("init"))
	return rv
}

func (bc _BrowserClass) Alloc() Browser {
	rv := objc.CallMethod[Browser](bc, objc.GetSelector("alloc"))
	return rv
}

func (bc _BrowserClass) New() Browser {
	rv := objc.CallMethod[Browser](bc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewBrowser() Browser {
	return BrowserClass.New()
}

func (b_ Browser) Tile() {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("tile"))
}

func (b_ Browser) SelectedRowIndexesInColumn(column int) foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](b_, objc.GetSelector("selectedRowIndexesInColumn:"), column)
	return rv
}

func (b_ Browser) SelectRowIndexes_InColumn(indexes foundation.IIndexSet, column int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("selectRowIndexes:inColumn:"), objc.ExtractPtr(indexes), column)
}

func (b_ Browser) SelectedCellInColumn(column int) objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("selectedCellInColumn:"), column)
	return rv
}

func (b_ Browser) SelectAll(sender objc.IObject) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("selectAll:"), objc.ExtractPtr(sender))
}

func (b_ Browser) SelectedRowInColumn(column int) int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("selectedRowInColumn:"), column)
	return rv
}

func (b_ Browser) SelectRow_InColumn(row int, column int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("selectRow:inColumn:"), row, column)
}

func (b_ Browser) LoadedCellAtRow_Column(row int, col int) objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("loadedCellAtRow:column:"), row, col)
	return rv
}

func (b_ Browser) EditItemAtIndexPath_WithEvent_Select(indexPath foundation.IIndexPath, event IEvent, select_ bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("editItemAtIndexPath:withEvent:select:"), objc.ExtractPtr(indexPath), objc.ExtractPtr(event), select_)
}

func (b_ Browser) ItemAtIndexPath(indexPath foundation.IIndexPath) objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("itemAtIndexPath:"), objc.ExtractPtr(indexPath))
	return rv
}

func (b_ Browser) ItemAtRow_InColumn(row int, column int) objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("itemAtRow:inColumn:"), row, column)
	return rv
}

func (b_ Browser) IndexPathForColumn(column int) foundation.IndexPath {
	rv := objc.CallMethod[foundation.IndexPath](b_, objc.GetSelector("indexPathForColumn:"), column)
	return rv
}

func (b_ Browser) IsLeafItem(item objc.IObject) bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("isLeafItem:"), objc.ExtractPtr(item))
	return rv
}

func (b_ Browser) ParentForItemsInColumn(column int) objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("parentForItemsInColumn:"), column)
	return rv
}

func (b_ Browser) Path() string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("path"))
	return rv
}

func (b_ Browser) SetPath(path string) bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("setPath:"), path)
	return rv
}

func (b_ Browser) PathToColumn(column int) string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("pathToColumn:"), column)
	return rv
}

func (b_ Browser) AddColumn() {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("addColumn"))
}

func (b_ Browser) ValidateVisibleColumns() {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("validateVisibleColumns"))
}

func (b_ Browser) LoadColumnZero() {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("loadColumnZero"))
}

func (b_ Browser) ReloadColumn(column int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("reloadColumn:"), column)
}

func (b_ Browser) TitleOfColumn(column int) string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("titleOfColumn:"), column)
	return rv
}

func (b_ Browser) SetTitle_OfColumn(string_ string, column int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setTitle:ofColumn:"), string_, column)
}

func (b_ Browser) DrawTitleOfColumn_InRect(column int, rect foundation.Rect) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("drawTitleOfColumn:inRect:"), column, rect)
}

func (b_ Browser) TitleFrameOfColumn(column int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](b_, objc.GetSelector("titleFrameOfColumn:"), column)
	return rv
}

func (b_ Browser) NoteHeightOfRowsWithIndexesChanged_InColumn(indexSet foundation.IIndexSet, columnIndex int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("noteHeightOfRowsWithIndexesChanged:inColumn:"), objc.ExtractPtr(indexSet), columnIndex)
}

func (b_ Browser) ReloadDataForRowIndexes_InColumn(rowIndexes foundation.IIndexSet, column int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("reloadDataForRowIndexes:inColumn:"), objc.ExtractPtr(rowIndexes), column)
}

func (b_ Browser) ScrollColumnToVisible(column int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("scrollColumnToVisible:"), column)
}

func (b_ Browser) ScrollColumnsLeftBy(shiftAmount int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("scrollColumnsLeftBy:"), shiftAmount)
}

func (b_ Browser) ScrollColumnsRightBy(shiftAmount int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("scrollColumnsRightBy:"), shiftAmount)
}

func (b_ Browser) ScrollRowToVisible_InColumn(row int, column int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("scrollRowToVisible:inColumn:"), row, column)
}

func (b_ Browser) SetDraggingSourceOperationMask_ForLocal(mask DragOperation, isLocal bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setDraggingSourceOperationMask:forLocal:"), mask, isLocal)
}

func (b_ Browser) CanDragRowsWithIndexes_InColumn_WithEvent(rowIndexes foundation.IIndexSet, column int, event IEvent) bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("canDragRowsWithIndexes:inColumn:withEvent:"), objc.ExtractPtr(rowIndexes), column, objc.ExtractPtr(event))
	return rv
}

func (b_ Browser) DraggingImageForRowsWithIndexes_InColumn_WithEvent_Offset(rowIndexes foundation.IIndexSet, column int, event IEvent, dragImageOffset *foundation.Point) Image {
	rv := objc.CallMethod[Image](b_, objc.GetSelector("draggingImageForRowsWithIndexes:inColumn:withEvent:offset:"), objc.ExtractPtr(rowIndexes), column, objc.ExtractPtr(event), dragImageOffset)
	return rv
}

func (b_ Browser) FrameOfColumn(column int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](b_, objc.GetSelector("frameOfColumn:"), column)
	return rv
}

func (b_ Browser) FrameOfInsideOfColumn(column int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](b_, objc.GetSelector("frameOfInsideOfColumn:"), column)
	return rv
}

func (b_ Browser) FrameOfRow_InColumn(row int, column int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](b_, objc.GetSelector("frameOfRow:inColumn:"), row, column)
	return rv
}

func (b_ Browser) GetRow_Column_ForPoint(row *int, column *int, point foundation.Point) bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("getRow:column:forPoint:"), row, column, point)
	return rv
}

func (b_ Browser) SendAction() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("sendAction"))
	return rv
}

func (b_ Browser) DoClick(sender objc.IObject) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("doClick:"), objc.ExtractPtr(sender))
}

func (b_ Browser) DoDoubleClick(sender objc.IObject) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("doDoubleClick:"), objc.ExtractPtr(sender))
}

func (bc _BrowserClass) RemoveSavedColumnsWithAutosaveName(name BrowserColumnsAutosaveName) {
	objc.CallMethod[objc.Void](bc, objc.GetSelector("removeSavedColumnsWithAutosaveName:"), name)
}

func (b_ Browser) ColumnContentWidthForColumnWidth(columnWidth float64) float64 {
	rv := objc.CallMethod[float64](b_, objc.GetSelector("columnContentWidthForColumnWidth:"), columnWidth)
	return rv
}

func (b_ Browser) ColumnWidthForColumnContentWidth(columnContentWidth float64) float64 {
	rv := objc.CallMethod[float64](b_, objc.GetSelector("columnWidthForColumnContentWidth:"), columnContentWidth)
	return rv
}

func (b_ Browser) WidthOfColumn(column int) float64 {
	rv := objc.CallMethod[float64](b_, objc.GetSelector("widthOfColumn:"), column)
	return rv
}

func (b_ Browser) SetWidth_OfColumn(columnWidth float64, columnIndex int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setWidth:ofColumn:"), columnWidth, columnIndex)
}

func (b_ Browser) DefaultColumnWidth() float64 {
	rv := objc.CallMethod[float64](b_, objc.GetSelector("defaultColumnWidth"))
	return rv
}

func (b_ Browser) SetDefaultColumnWidth(columnWidth float64) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setDefaultColumnWidth:"), columnWidth)
}

// deprecated
func (b_ Browser) UpdateScroller() {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("updateScroller"))
}

// deprecated
func (b_ Browser) ScrollViaScroller(sender IScroller) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("scrollViaScroller:"), objc.ExtractPtr(sender))
}

// deprecated
func (b_ Browser) DisplayAllColumns() {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("displayAllColumns"))
}

// deprecated
func (b_ Browser) DisplayColumn(column int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("displayColumn:"), column)
}

// deprecated
func (b_ Browser) ColumnOfMatrix(matrix IMatrix) int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("columnOfMatrix:"), objc.ExtractPtr(matrix))
	return rv
}

// deprecated
func (b_ Browser) MatrixInColumn(column int) Matrix {
	rv := objc.CallMethod[Matrix](b_, objc.GetSelector("matrixInColumn:"), column)
	return rv
}

// deprecated
func (b_ Browser) MatrixClass() objc.Class {
	rv := objc.CallMethod[objc.Class](b_, objc.GetSelector("matrixClass"))
	return rv
}

// deprecated
func (b_ Browser) SetMatrixClass(factoryId objc.IClass) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setMatrixClass:"), objc.ExtractPtr(factoryId))
}

// deprecated
func (b_ Browser) AcceptsArrowKeys() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("acceptsArrowKeys"))
	return rv
}

// deprecated
func (b_ Browser) SetAcceptsArrowKeys(flag bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setAcceptsArrowKeys:"), flag)
}

func (b_ Browser) ReusesColumns() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("reusesColumns"))
	return rv
}

func (b_ Browser) SetReusesColumns(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setReusesColumns:"), value)
}

func (b_ Browser) MaxVisibleColumns() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("maxVisibleColumns"))
	return rv
}

func (b_ Browser) SetMaxVisibleColumns(value int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setMaxVisibleColumns:"), value)
}

func (b_ Browser) AutohidesScroller() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("autohidesScroller"))
	return rv
}

func (b_ Browser) SetAutohidesScroller(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setAutohidesScroller:"), value)
}

func (b_ Browser) BackgroundColor() Color {
	rv := objc.CallMethod[Color](b_, objc.GetSelector("backgroundColor"))
	return rv
}

func (b_ Browser) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (b_ Browser) MinColumnWidth() float64 {
	rv := objc.CallMethod[float64](b_, objc.GetSelector("minColumnWidth"))
	return rv
}

func (b_ Browser) SetMinColumnWidth(value float64) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setMinColumnWidth:"), value)
}

func (b_ Browser) SeparatesColumns() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("separatesColumns"))
	return rv
}

func (b_ Browser) SetSeparatesColumns(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setSeparatesColumns:"), value)
}

func (b_ Browser) TakesTitleFromPreviousColumn() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("takesTitleFromPreviousColumn"))
	return rv
}

func (b_ Browser) SetTakesTitleFromPreviousColumn(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setTakesTitleFromPreviousColumn:"), value)
}

// weak property
func (b_ Browser) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("delegate"))
	return rv
}

// weak property
func (b_ Browser) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (b_ Browser) CellPrototype() objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("cellPrototype"))
	return rv
}

func (b_ Browser) SetCellPrototype(value objc.IObject) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setCellPrototype:"), objc.ExtractPtr(value))
}

func (b_ Browser) AllowsBranchSelection() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("allowsBranchSelection"))
	return rv
}

func (b_ Browser) SetAllowsBranchSelection(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setAllowsBranchSelection:"), value)
}

func (b_ Browser) AllowsEmptySelection() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("allowsEmptySelection"))
	return rv
}

func (b_ Browser) SetAllowsEmptySelection(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setAllowsEmptySelection:"), value)
}

func (b_ Browser) AllowsMultipleSelection() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("allowsMultipleSelection"))
	return rv
}

func (b_ Browser) SetAllowsMultipleSelection(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setAllowsMultipleSelection:"), value)
}

func (b_ Browser) AllowsTypeSelect() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("allowsTypeSelect"))
	return rv
}

func (b_ Browser) SetAllowsTypeSelect(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setAllowsTypeSelect:"), value)
}

func (b_ Browser) SelectedCells() []Cell {
	rv := objc.CallMethod[[]Cell](b_, objc.GetSelector("selectedCells"))
	return rv
}

func (b_ Browser) SelectionIndexPath() foundation.IndexPath {
	rv := objc.CallMethod[foundation.IndexPath](b_, objc.GetSelector("selectionIndexPath"))
	return rv
}

func (b_ Browser) SetSelectionIndexPath(value foundation.IIndexPath) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setSelectionIndexPath:"), objc.ExtractPtr(value))
}

func (b_ Browser) SelectionIndexPaths() []foundation.IndexPath {
	rv := objc.CallMethod[[]foundation.IndexPath](b_, objc.GetSelector("selectionIndexPaths"))
	return rv
}

func (b_ Browser) SetSelectionIndexPaths(value []foundation.IIndexPath) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setSelectionIndexPaths:"), value)
}

func (b_ Browser) PathSeparator() string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("pathSeparator"))
	return rv
}

func (b_ Browser) SetPathSeparator(value string) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setPathSeparator:"), value)
}

func (b_ Browser) SelectedColumn() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("selectedColumn"))
	return rv
}

func (b_ Browser) LastColumn() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("lastColumn"))
	return rv
}

func (b_ Browser) SetLastColumn(value int) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setLastColumn:"), value)
}

func (b_ Browser) FirstVisibleColumn() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("firstVisibleColumn"))
	return rv
}

func (b_ Browser) NumberOfVisibleColumns() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("numberOfVisibleColumns"))
	return rv
}

func (b_ Browser) LastVisibleColumn() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("lastVisibleColumn"))
	return rv
}

func (b_ Browser) IsLoaded() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("isLoaded"))
	return rv
}

func (b_ Browser) IsTitled() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("isTitled"))
	return rv
}

func (b_ Browser) SetTitled(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setTitled:"), value)
}

func (b_ Browser) TitleHeight() float64 {
	rv := objc.CallMethod[float64](b_, objc.GetSelector("titleHeight"))
	return rv
}

func (b_ Browser) HasHorizontalScroller() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("hasHorizontalScroller"))
	return rv
}

func (b_ Browser) SetHasHorizontalScroller(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setHasHorizontalScroller:"), value)
}

func (b_ Browser) DoubleAction() objc.Selector {
	rv := objc.CallMethod[objc.Selector](b_, objc.GetSelector("doubleAction"))
	return rv
}

func (b_ Browser) SetDoubleAction(value objc.Selector) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setDoubleAction:"), value)
}

func (b_ Browser) SendsActionOnArrowKeys() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("sendsActionOnArrowKeys"))
	return rv
}

func (b_ Browser) SetSendsActionOnArrowKeys(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setSendsActionOnArrowKeys:"), value)
}

func (b_ Browser) ClickedColumn() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("clickedColumn"))
	return rv
}

func (b_ Browser) ClickedRow() int {
	rv := objc.CallMethod[int](b_, objc.GetSelector("clickedRow"))
	return rv
}

func (b_ Browser) ColumnsAutosaveName() BrowserColumnsAutosaveName {
	rv := objc.CallMethod[BrowserColumnsAutosaveName](b_, objc.GetSelector("columnsAutosaveName"))
	return rv
}

func (b_ Browser) SetColumnsAutosaveName(value BrowserColumnsAutosaveName) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setColumnsAutosaveName:"), value)
}

func (b_ Browser) ColumnResizingType() BrowserColumnResizingType {
	rv := objc.CallMethod[BrowserColumnResizingType](b_, objc.GetSelector("columnResizingType"))
	return rv
}

func (b_ Browser) SetColumnResizingType(value BrowserColumnResizingType) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setColumnResizingType:"), value)
}

func (b_ Browser) PrefersAllColumnUserResizing() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("prefersAllColumnUserResizing"))
	return rv
}

func (b_ Browser) SetPrefersAllColumnUserResizing(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setPrefersAllColumnUserResizing:"), value)
}

func (b_ Browser) RowHeight() float64 {
	rv := objc.CallMethod[float64](b_, objc.GetSelector("rowHeight"))
	return rv
}

func (b_ Browser) SetRowHeight(value float64) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setRowHeight:"), value)
}
