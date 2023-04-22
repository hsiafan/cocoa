// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var MatrixClass = _MatrixClass{objc.GetClass("NSMatrix")}

type _MatrixClass struct {
	objc.Class
}

type IMatrix interface {
	IControl
	AddColumn()
	AddColumnWithCells(newCells []ICell)
	AddRow()
	AddRowWithCells(newCells []ICell)
	CellFrameAtRow_Column(row int, col int) foundation.Rect
	GetNumberOfRows_Columns(rowCount *int, colCount *int)
	InsertColumn(column int)
	InsertColumn_WithCells(column int, newCells []ICell)
	InsertRow(row int)
	InsertRow_WithCells(row int, newCells []ICell)
	MakeCellAtRow_Column(row int, col int) Cell
	PutCell_AtRow_Column(newCell ICell, row int, col int)
	RemoveColumn(col int)
	RemoveRow(row int)
	RenewRows_Columns(newRows int, newCols int)
	SortUsingSelector(comparator objc.Selector)
	GetRow_Column_ForPoint(row *int, col *int, point foundation.Point) bool
	GetRow_Column_OfCell(row *int, col *int, cell ICell) bool
	SetState_AtRow_Column(value int, row int, col int)
	SetToolTip_ForCell(toolTipString string, cell ICell)
	ToolTipForCell(cell ICell) string
	SelectCellAtRow_Column(row int, col int)
	SelectCellWithTag(tag int) bool
	SelectAll(sender objc.IObject)
	SetSelectionFrom_To_Anchor_Highlight(startPos int, endPos int, anchorPos int, lit bool)
	DeselectAllCells()
	DeselectSelectedCell()
	CellAtRow_Column(row int, col int) Cell
	CellWithTag(tag int) Cell
	SelectText(sender objc.IObject)
	SelectTextAtRow_Column(row int, col int) Cell
	TextShouldBeginEditing(textObject IText) bool
	TextDidBeginEditing(notification foundation.INotification)
	TextDidChange(notification foundation.INotification)
	TextShouldEndEditing(textObject IText) bool
	TextDidEndEditing(notification foundation.INotification)
	SetValidateSize(flag bool)
	SizeToCells()
	SetScrollable(flag bool)
	ScrollCellToVisibleAtRow_Column(row int, col int)
	DrawCellAtRow_Column(row int, col int)
	HighlightCell_AtRow_Column(flag bool, row int, col int)
	SendAction() bool
	SendAction_To_ForAllCells(selector objc.Selector, object objc.IObject, flag bool)
	SendDoubleAction()
	Mode() MatrixMode
	SetMode(value MatrixMode)
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	IsSelectionByRect() bool
	SetSelectionByRect(value bool)
	Prototype() Cell
	SetPrototype(value ICell)
	CellSize() foundation.Size
	SetCellSize(value foundation.Size)
	IntercellSpacing() foundation.Size
	SetIntercellSpacing(value foundation.Size)
	NumberOfColumns() int
	NumberOfRows() int
	AutorecalculatesCellSize() bool
	SetAutorecalculatesCellSize(value bool)
	KeyCell() Cell
	SetKeyCell(value ICell)
	SelectedCells() []Cell
	SelectedColumn() int
	SelectedRow() int
	Cells() []Cell
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	CellBackgroundColor() Color
	SetCellBackgroundColor(value IColor)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	DrawsCellBackground() bool
	SetDrawsCellBackground(value bool)
	TabKeyTraversesCells() bool
	SetTabKeyTraversesCells(value bool)
	Delegate() MatrixDelegateWrapper
	SetDelegate(value MatrixDelegate)
	SetDelegate0(value objc.IObject)
	AutosizesCells() bool
	SetAutosizesCells(value bool)
	IsAutoscroll() bool
	SetAutoscroll(value bool)
	DoubleAction() objc.Selector
	SetDoubleAction(value objc.Selector)
	MouseDownFlags() int
}

type Matrix struct {
	Control
}

func MakeMatrix(ptr unsafe.Pointer) Matrix {
	return Matrix{
		Control: MakeControl(ptr),
	}
}

func (m_ Matrix) InitWithFrame(frameRect foundation.Rect) Matrix {
	rv := objc.CallMethod[Matrix](m_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func (m_ Matrix) InitWithFrame_Mode_CellClass_NumberOfRows_NumberOfColumns(frameRect foundation.Rect, mode MatrixMode, factoryId objc.IClass, rowsHigh int, colsWide int) Matrix {
	rv := objc.CallMethod[Matrix](m_, objc.GetSelector("initWithFrame:mode:cellClass:numberOfRows:numberOfColumns:"), frameRect, mode, factoryId, rowsHigh, colsWide)
	return rv
}

func (m_ Matrix) InitWithFrame_Mode_Prototype_NumberOfRows_NumberOfColumns(frameRect foundation.Rect, mode MatrixMode, cell ICell, rowsHigh int, colsWide int) Matrix {
	rv := objc.CallMethod[Matrix](m_, objc.GetSelector("initWithFrame:mode:prototype:numberOfRows:numberOfColumns:"), frameRect, mode, cell, rowsHigh, colsWide)
	return rv
}

func (m_ Matrix) Init() Matrix {
	rv := objc.CallMethod[Matrix](m_, objc.GetSelector("init"))
	return rv
}

func (mc _MatrixClass) Alloc() Matrix {
	rv := objc.CallMethod[Matrix](mc, objc.GetSelector("alloc"))
	return rv
}

func (mc _MatrixClass) New() Matrix {
	rv := objc.CallMethod[Matrix](mc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewMatrix() Matrix {
	return MatrixClass.New()
}

func (m_ Matrix) AddColumn() {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("addColumn"))
}

func (m_ Matrix) AddColumnWithCells(newCells []ICell) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("addColumnWithCells:"), newCells)
}

func (m_ Matrix) AddRow() {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("addRow"))
}

func (m_ Matrix) AddRowWithCells(newCells []ICell) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("addRowWithCells:"), newCells)
}

func (m_ Matrix) CellFrameAtRow_Column(row int, col int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](m_, objc.GetSelector("cellFrameAtRow:column:"), row, col)
	return rv
}

func (m_ Matrix) GetNumberOfRows_Columns(rowCount *int, colCount *int) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("getNumberOfRows:columns:"), rowCount, colCount)
}

func (m_ Matrix) InsertColumn(column int) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("insertColumn:"), column)
}

func (m_ Matrix) InsertColumn_WithCells(column int, newCells []ICell) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("insertColumn:withCells:"), column, newCells)
}

func (m_ Matrix) InsertRow(row int) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("insertRow:"), row)
}

func (m_ Matrix) InsertRow_WithCells(row int, newCells []ICell) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("insertRow:withCells:"), row, newCells)
}

func (m_ Matrix) MakeCellAtRow_Column(row int, col int) Cell {
	rv := objc.CallMethod[Cell](m_, objc.GetSelector("makeCellAtRow:column:"), row, col)
	return rv
}

func (m_ Matrix) PutCell_AtRow_Column(newCell ICell, row int, col int) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("putCell:atRow:column:"), newCell, row, col)
}

func (m_ Matrix) RemoveColumn(col int) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("removeColumn:"), col)
}

func (m_ Matrix) RemoveRow(row int) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("removeRow:"), row)
}

func (m_ Matrix) RenewRows_Columns(newRows int, newCols int) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("renewRows:columns:"), newRows, newCols)
}

func (m_ Matrix) SortUsingSelector(comparator objc.Selector) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("sortUsingSelector:"), comparator)
}

func (m_ Matrix) GetRow_Column_ForPoint(row *int, col *int, point foundation.Point) bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("getRow:column:forPoint:"), row, col, point)
	return rv
}

func (m_ Matrix) GetRow_Column_OfCell(row *int, col *int, cell ICell) bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("getRow:column:ofCell:"), row, col, cell)
	return rv
}

func (m_ Matrix) SetState_AtRow_Column(value int, row int, col int) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setState:atRow:column:"), value, row, col)
}

func (m_ Matrix) SetToolTip_ForCell(toolTipString string, cell ICell) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setToolTip:forCell:"), toolTipString, cell)
}

func (m_ Matrix) ToolTipForCell(cell ICell) string {
	rv := objc.CallMethod[string](m_, objc.GetSelector("toolTipForCell:"), cell)
	return rv
}

func (m_ Matrix) SelectCellAtRow_Column(row int, col int) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("selectCellAtRow:column:"), row, col)
}

func (m_ Matrix) SelectCellWithTag(tag int) bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("selectCellWithTag:"), tag)
	return rv
}

func (m_ Matrix) SelectAll(sender objc.IObject) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("selectAll:"), sender)
}

func (m_ Matrix) SetSelectionFrom_To_Anchor_Highlight(startPos int, endPos int, anchorPos int, lit bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setSelectionFrom:to:anchor:highlight:"), startPos, endPos, anchorPos, lit)
}

func (m_ Matrix) DeselectAllCells() {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("deselectAllCells"))
}

func (m_ Matrix) DeselectSelectedCell() {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("deselectSelectedCell"))
}

func (m_ Matrix) CellAtRow_Column(row int, col int) Cell {
	rv := objc.CallMethod[Cell](m_, objc.GetSelector("cellAtRow:column:"), row, col)
	return rv
}

func (m_ Matrix) CellWithTag(tag int) Cell {
	rv := objc.CallMethod[Cell](m_, objc.GetSelector("cellWithTag:"), tag)
	return rv
}

func (m_ Matrix) SelectText(sender objc.IObject) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("selectText:"), sender)
}

func (m_ Matrix) SelectTextAtRow_Column(row int, col int) Cell {
	rv := objc.CallMethod[Cell](m_, objc.GetSelector("selectTextAtRow:column:"), row, col)
	return rv
}

func (m_ Matrix) TextShouldBeginEditing(textObject IText) bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("textShouldBeginEditing:"), textObject)
	return rv
}

func (m_ Matrix) TextDidBeginEditing(notification foundation.INotification) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("textDidBeginEditing:"), notification)
}

func (m_ Matrix) TextDidChange(notification foundation.INotification) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("textDidChange:"), notification)
}

func (m_ Matrix) TextShouldEndEditing(textObject IText) bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("textShouldEndEditing:"), textObject)
	return rv
}

func (m_ Matrix) TextDidEndEditing(notification foundation.INotification) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("textDidEndEditing:"), notification)
}

func (m_ Matrix) SetValidateSize(flag bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setValidateSize:"), flag)
}

func (m_ Matrix) SizeToCells() {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("sizeToCells"))
}

func (m_ Matrix) SetScrollable(flag bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setScrollable:"), flag)
}

func (m_ Matrix) ScrollCellToVisibleAtRow_Column(row int, col int) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("scrollCellToVisibleAtRow:column:"), row, col)
}

func (m_ Matrix) DrawCellAtRow_Column(row int, col int) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("drawCellAtRow:column:"), row, col)
}

func (m_ Matrix) HighlightCell_AtRow_Column(flag bool, row int, col int) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("highlightCell:atRow:column:"), flag, row, col)
}

func (m_ Matrix) SendAction() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("sendAction"))
	return rv
}

func (m_ Matrix) SendAction_To_ForAllCells(selector objc.Selector, object objc.IObject, flag bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("sendAction:to:forAllCells:"), selector, object, flag)
}

func (m_ Matrix) SendDoubleAction() {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("sendDoubleAction"))
}

func (m_ Matrix) Mode() MatrixMode {
	rv := objc.CallMethod[MatrixMode](m_, objc.GetSelector("mode"))
	return rv
}

func (m_ Matrix) SetMode(value MatrixMode) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setMode:"), value)
}

func (m_ Matrix) AllowsEmptySelection() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("allowsEmptySelection"))
	return rv
}

func (m_ Matrix) SetAllowsEmptySelection(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAllowsEmptySelection:"), value)
}

func (m_ Matrix) IsSelectionByRect() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("isSelectionByRect"))
	return rv
}

func (m_ Matrix) SetSelectionByRect(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setSelectionByRect:"), value)
}

func (m_ Matrix) Prototype() Cell {
	rv := objc.CallMethod[Cell](m_, objc.GetSelector("prototype"))
	return rv
}

func (m_ Matrix) SetPrototype(value ICell) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setPrototype:"), value)
}

func (m_ Matrix) CellSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](m_, objc.GetSelector("cellSize"))
	return rv
}

func (m_ Matrix) SetCellSize(value foundation.Size) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setCellSize:"), value)
}

func (m_ Matrix) IntercellSpacing() foundation.Size {
	rv := objc.CallMethod[foundation.Size](m_, objc.GetSelector("intercellSpacing"))
	return rv
}

func (m_ Matrix) SetIntercellSpacing(value foundation.Size) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setIntercellSpacing:"), value)
}

func (m_ Matrix) NumberOfColumns() int {
	rv := objc.CallMethod[int](m_, objc.GetSelector("numberOfColumns"))
	return rv
}

func (m_ Matrix) NumberOfRows() int {
	rv := objc.CallMethod[int](m_, objc.GetSelector("numberOfRows"))
	return rv
}

func (m_ Matrix) AutorecalculatesCellSize() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("autorecalculatesCellSize"))
	return rv
}

func (m_ Matrix) SetAutorecalculatesCellSize(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAutorecalculatesCellSize:"), value)
}

func (m_ Matrix) KeyCell() Cell {
	rv := objc.CallMethod[Cell](m_, objc.GetSelector("keyCell"))
	return rv
}

func (m_ Matrix) SetKeyCell(value ICell) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setKeyCell:"), value)
}

func (m_ Matrix) SelectedCells() []Cell {
	rv := objc.CallMethod[[]Cell](m_, objc.GetSelector("selectedCells"))
	return rv
}

func (m_ Matrix) SelectedColumn() int {
	rv := objc.CallMethod[int](m_, objc.GetSelector("selectedColumn"))
	return rv
}

func (m_ Matrix) SelectedRow() int {
	rv := objc.CallMethod[int](m_, objc.GetSelector("selectedRow"))
	return rv
}

func (m_ Matrix) Cells() []Cell {
	rv := objc.CallMethod[[]Cell](m_, objc.GetSelector("cells"))
	return rv
}

func (m_ Matrix) BackgroundColor() Color {
	rv := objc.CallMethod[Color](m_, objc.GetSelector("backgroundColor"))
	return rv
}

func (m_ Matrix) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setBackgroundColor:"), value)
}

func (m_ Matrix) CellBackgroundColor() Color {
	rv := objc.CallMethod[Color](m_, objc.GetSelector("cellBackgroundColor"))
	return rv
}

func (m_ Matrix) SetCellBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setCellBackgroundColor:"), value)
}

func (m_ Matrix) DrawsBackground() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("drawsBackground"))
	return rv
}

func (m_ Matrix) SetDrawsBackground(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setDrawsBackground:"), value)
}

func (m_ Matrix) DrawsCellBackground() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("drawsCellBackground"))
	return rv
}

func (m_ Matrix) SetDrawsCellBackground(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setDrawsCellBackground:"), value)
}

func (m_ Matrix) TabKeyTraversesCells() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("tabKeyTraversesCells"))
	return rv
}

func (m_ Matrix) SetTabKeyTraversesCells(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setTabKeyTraversesCells:"), value)
}

func (m_ Matrix) Delegate() MatrixDelegateWrapper {
	rv := objc.CallMethod[MatrixDelegateWrapper](m_, objc.GetSelector("delegate"))
	return rv
}

func (m_ Matrix) SetDelegate(value MatrixDelegate) {
	po := objc.CreateProtocol("NSMatrixDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(m_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setDelegate:"), po)
}

func (m_ Matrix) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setDelegate:"), value)
}

func (m_ Matrix) AutosizesCells() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("autosizesCells"))
	return rv
}

func (m_ Matrix) SetAutosizesCells(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAutosizesCells:"), value)
}

func (m_ Matrix) IsAutoscroll() bool {
	rv := objc.CallMethod[bool](m_, objc.GetSelector("isAutoscroll"))
	return rv
}

func (m_ Matrix) SetAutoscroll(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAutoscroll:"), value)
}

func (m_ Matrix) DoubleAction() objc.Selector {
	rv := objc.CallMethod[objc.Selector](m_, objc.GetSelector("doubleAction"))
	return rv
}

func (m_ Matrix) SetDoubleAction(value objc.Selector) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setDoubleAction:"), value)
}

func (m_ Matrix) MouseDownFlags() int {
	rv := objc.CallMethod[int](m_, objc.GetSelector("mouseDownFlags"))
	return rv
}
