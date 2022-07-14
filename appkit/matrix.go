// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[Matrix](m_, "initWithFrame:", frameRect)
	rv.Autorelease()
	return rv
}

func (m_ Matrix) InitWithFrame_Mode_Prototype_NumberOfRows_NumberOfColumns(frameRect foundation.Rect, mode MatrixMode, cell ICell, rowsHigh int, colsWide int) Matrix {
	rv := ffi.CallMethod[Matrix](m_, "initWithFrame:mode:prototype:numberOfRows:numberOfColumns:", frameRect, mode, cell, rowsHigh, colsWide)
	rv.Autorelease()
	return rv
}

func (m_ Matrix) Init() Matrix {
	rv := ffi.CallMethod[Matrix](m_, "init")
	rv.Autorelease()
	return rv
}

func (mc _MatrixClass) Alloc() Matrix {
	rv := ffi.CallMethod[Matrix](mc, "alloc")
	return rv
}

func (mc _MatrixClass) New() Matrix {
	rv := ffi.CallMethod[Matrix](mc, "new")
	rv.Autorelease()
	return rv
}

func NewMatrix() Matrix {
	return MatrixClass.New()
}

func (m_ Matrix) AddColumn() {
	ffi.CallMethod[ffi.Void](m_, "addColumn")
}

func (m_ Matrix) AddColumnWithCells(newCells []ICell) {
	ffi.CallMethod[ffi.Void](m_, "addColumnWithCells:", newCells)
}

func (m_ Matrix) AddRow() {
	ffi.CallMethod[ffi.Void](m_, "addRow")
}

func (m_ Matrix) AddRowWithCells(newCells []ICell) {
	ffi.CallMethod[ffi.Void](m_, "addRowWithCells:", newCells)
}

func (m_ Matrix) CellFrameAtRow_Column(row int, col int) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](m_, "cellFrameAtRow:column:", row, col)
	return rv
}

func (m_ Matrix) GetNumberOfRows_Columns(rowCount *int, colCount *int) {
	ffi.CallMethod[ffi.Void](m_, "getNumberOfRows:columns:", rowCount, colCount)
}

func (m_ Matrix) InsertColumn(column int) {
	ffi.CallMethod[ffi.Void](m_, "insertColumn:", column)
}

func (m_ Matrix) InsertColumn_WithCells(column int, newCells []ICell) {
	ffi.CallMethod[ffi.Void](m_, "insertColumn:withCells:", column, newCells)
}

func (m_ Matrix) InsertRow(row int) {
	ffi.CallMethod[ffi.Void](m_, "insertRow:", row)
}

func (m_ Matrix) InsertRow_WithCells(row int, newCells []ICell) {
	ffi.CallMethod[ffi.Void](m_, "insertRow:withCells:", row, newCells)
}

func (m_ Matrix) MakeCellAtRow_Column(row int, col int) Cell {
	rv := ffi.CallMethod[Cell](m_, "makeCellAtRow:column:", row, col)
	return rv
}

func (m_ Matrix) PutCell_AtRow_Column(newCell ICell, row int, col int) {
	ffi.CallMethod[ffi.Void](m_, "putCell:atRow:column:", newCell, row, col)
}

func (m_ Matrix) RemoveColumn(col int) {
	ffi.CallMethod[ffi.Void](m_, "removeColumn:", col)
}

func (m_ Matrix) RemoveRow(row int) {
	ffi.CallMethod[ffi.Void](m_, "removeRow:", row)
}

func (m_ Matrix) RenewRows_Columns(newRows int, newCols int) {
	ffi.CallMethod[ffi.Void](m_, "renewRows:columns:", newRows, newCols)
}

func (m_ Matrix) SortUsingSelector(comparator objc.Selector) {
	ffi.CallMethod[ffi.Void](m_, "sortUsingSelector:", comparator)
}

func (m_ Matrix) GetRow_Column_ForPoint(row *int, col *int, point foundation.Point) bool {
	rv := ffi.CallMethod[bool](m_, "getRow:column:forPoint:", row, col, point)
	return rv
}

func (m_ Matrix) GetRow_Column_OfCell(row *int, col *int, cell ICell) bool {
	rv := ffi.CallMethod[bool](m_, "getRow:column:ofCell:", row, col, cell)
	return rv
}

func (m_ Matrix) SetState_AtRow_Column(value int, row int, col int) {
	ffi.CallMethod[ffi.Void](m_, "setState:atRow:column:", value, row, col)
}

func (m_ Matrix) SetToolTip_ForCell(toolTipString string, cell ICell) {
	ffi.CallMethod[ffi.Void](m_, "setToolTip:forCell:", toolTipString, cell)
}

func (m_ Matrix) ToolTipForCell(cell ICell) string {
	rv := ffi.CallMethod[string](m_, "toolTipForCell:", cell)
	return rv
}

func (m_ Matrix) SelectCellAtRow_Column(row int, col int) {
	ffi.CallMethod[ffi.Void](m_, "selectCellAtRow:column:", row, col)
}

func (m_ Matrix) SelectCellWithTag(tag int) bool {
	rv := ffi.CallMethod[bool](m_, "selectCellWithTag:", tag)
	return rv
}

func (m_ Matrix) SelectAll(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](m_, "selectAll:", sender)
}

func (m_ Matrix) SetSelectionFrom_To_Anchor_Highlight(startPos int, endPos int, anchorPos int, lit bool) {
	ffi.CallMethod[ffi.Void](m_, "setSelectionFrom:to:anchor:highlight:", startPos, endPos, anchorPos, lit)
}

func (m_ Matrix) DeselectAllCells() {
	ffi.CallMethod[ffi.Void](m_, "deselectAllCells")
}

func (m_ Matrix) DeselectSelectedCell() {
	ffi.CallMethod[ffi.Void](m_, "deselectSelectedCell")
}

func (m_ Matrix) CellAtRow_Column(row int, col int) Cell {
	rv := ffi.CallMethod[Cell](m_, "cellAtRow:column:", row, col)
	return rv
}

func (m_ Matrix) CellWithTag(tag int) Cell {
	rv := ffi.CallMethod[Cell](m_, "cellWithTag:", tag)
	return rv
}

func (m_ Matrix) SelectText(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](m_, "selectText:", sender)
}

func (m_ Matrix) SelectTextAtRow_Column(row int, col int) Cell {
	rv := ffi.CallMethod[Cell](m_, "selectTextAtRow:column:", row, col)
	return rv
}

func (m_ Matrix) TextShouldBeginEditing(textObject IText) bool {
	rv := ffi.CallMethod[bool](m_, "textShouldBeginEditing:", textObject)
	return rv
}

func (m_ Matrix) TextDidBeginEditing(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](m_, "textDidBeginEditing:", notification)
}

func (m_ Matrix) TextDidChange(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](m_, "textDidChange:", notification)
}

func (m_ Matrix) TextShouldEndEditing(textObject IText) bool {
	rv := ffi.CallMethod[bool](m_, "textShouldEndEditing:", textObject)
	return rv
}

func (m_ Matrix) TextDidEndEditing(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](m_, "textDidEndEditing:", notification)
}

func (m_ Matrix) SetValidateSize(flag bool) {
	ffi.CallMethod[ffi.Void](m_, "setValidateSize:", flag)
}

func (m_ Matrix) SizeToCells() {
	ffi.CallMethod[ffi.Void](m_, "sizeToCells")
}

func (m_ Matrix) SetScrollable(flag bool) {
	ffi.CallMethod[ffi.Void](m_, "setScrollable:", flag)
}

func (m_ Matrix) ScrollCellToVisibleAtRow_Column(row int, col int) {
	ffi.CallMethod[ffi.Void](m_, "scrollCellToVisibleAtRow:column:", row, col)
}

func (m_ Matrix) DrawCellAtRow_Column(row int, col int) {
	ffi.CallMethod[ffi.Void](m_, "drawCellAtRow:column:", row, col)
}

func (m_ Matrix) HighlightCell_AtRow_Column(flag bool, row int, col int) {
	ffi.CallMethod[ffi.Void](m_, "highlightCell:atRow:column:", flag, row, col)
}

func (m_ Matrix) SendAction() bool {
	rv := ffi.CallMethod[bool](m_, "sendAction")
	return rv
}

func (m_ Matrix) SendAction_To_ForAllCells(selector objc.Selector, object objc.IObject, flag bool) {
	ffi.CallMethod[ffi.Void](m_, "sendAction:to:forAllCells:", selector, object, flag)
}

func (m_ Matrix) SendDoubleAction() {
	ffi.CallMethod[ffi.Void](m_, "sendDoubleAction")
}

func (m_ Matrix) Mode() MatrixMode {
	rv := ffi.CallMethod[MatrixMode](m_, "mode")
	return rv
}

func (m_ Matrix) SetMode(value MatrixMode) {
	ffi.CallMethod[ffi.Void](m_, "setMode:", value)
}

func (m_ Matrix) AllowsEmptySelection() bool {
	rv := ffi.CallMethod[bool](m_, "allowsEmptySelection")
	return rv
}

func (m_ Matrix) SetAllowsEmptySelection(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setAllowsEmptySelection:", value)
}

func (m_ Matrix) IsSelectionByRect() bool {
	rv := ffi.CallMethod[bool](m_, "isSelectionByRect")
	return rv
}

func (m_ Matrix) SetSelectionByRect(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setSelectionByRect:", value)
}

func (m_ Matrix) Prototype() Cell {
	rv := ffi.CallMethod[Cell](m_, "prototype")
	return rv
}

func (m_ Matrix) SetPrototype(value ICell) {
	ffi.CallMethod[ffi.Void](m_, "setPrototype:", value)
}

func (m_ Matrix) CellSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](m_, "cellSize")
	return rv
}

func (m_ Matrix) SetCellSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](m_, "setCellSize:", value)
}

func (m_ Matrix) IntercellSpacing() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](m_, "intercellSpacing")
	return rv
}

func (m_ Matrix) SetIntercellSpacing(value foundation.Size) {
	ffi.CallMethod[ffi.Void](m_, "setIntercellSpacing:", value)
}

func (m_ Matrix) NumberOfColumns() int {
	rv := ffi.CallMethod[int](m_, "numberOfColumns")
	return rv
}

func (m_ Matrix) NumberOfRows() int {
	rv := ffi.CallMethod[int](m_, "numberOfRows")
	return rv
}

func (m_ Matrix) AutorecalculatesCellSize() bool {
	rv := ffi.CallMethod[bool](m_, "autorecalculatesCellSize")
	return rv
}

func (m_ Matrix) SetAutorecalculatesCellSize(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setAutorecalculatesCellSize:", value)
}

func (m_ Matrix) KeyCell() Cell {
	rv := ffi.CallMethod[Cell](m_, "keyCell")
	return rv
}

func (m_ Matrix) SetKeyCell(value ICell) {
	ffi.CallMethod[ffi.Void](m_, "setKeyCell:", value)
}

func (m_ Matrix) SelectedCells() []Cell {
	rv := ffi.CallMethod[[]Cell](m_, "selectedCells")
	return rv
}

func (m_ Matrix) SelectedColumn() int {
	rv := ffi.CallMethod[int](m_, "selectedColumn")
	return rv
}

func (m_ Matrix) SelectedRow() int {
	rv := ffi.CallMethod[int](m_, "selectedRow")
	return rv
}

func (m_ Matrix) Cells() []Cell {
	rv := ffi.CallMethod[[]Cell](m_, "cells")
	return rv
}

func (m_ Matrix) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](m_, "backgroundColor")
	return rv
}

func (m_ Matrix) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](m_, "setBackgroundColor:", value)
}

func (m_ Matrix) CellBackgroundColor() Color {
	rv := ffi.CallMethod[Color](m_, "cellBackgroundColor")
	return rv
}

func (m_ Matrix) SetCellBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](m_, "setCellBackgroundColor:", value)
}

func (m_ Matrix) DrawsBackground() bool {
	rv := ffi.CallMethod[bool](m_, "drawsBackground")
	return rv
}

func (m_ Matrix) SetDrawsBackground(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setDrawsBackground:", value)
}

func (m_ Matrix) DrawsCellBackground() bool {
	rv := ffi.CallMethod[bool](m_, "drawsCellBackground")
	return rv
}

func (m_ Matrix) SetDrawsCellBackground(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setDrawsCellBackground:", value)
}

func (m_ Matrix) TabKeyTraversesCells() bool {
	rv := ffi.CallMethod[bool](m_, "tabKeyTraversesCells")
	return rv
}

func (m_ Matrix) SetTabKeyTraversesCells(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setTabKeyTraversesCells:", value)
}

func (m_ Matrix) Delegate() MatrixDelegateWrapper {
	rv := ffi.CallMethod[MatrixDelegateWrapper](m_, "delegate")
	return rv
}

func (m_ Matrix) SetDelegate(value MatrixDelegate) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(m_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](m_, "setDelegate:", po)
}

func (m_ Matrix) AutosizesCells() bool {
	rv := ffi.CallMethod[bool](m_, "autosizesCells")
	return rv
}

func (m_ Matrix) SetAutosizesCells(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setAutosizesCells:", value)
}

func (m_ Matrix) IsAutoscroll() bool {
	rv := ffi.CallMethod[bool](m_, "isAutoscroll")
	return rv
}

func (m_ Matrix) SetAutoscroll(value bool) {
	ffi.CallMethod[ffi.Void](m_, "setAutoscroll:", value)
}

func (m_ Matrix) DoubleAction() objc.Selector {
	rv := ffi.CallMethod[objc.Selector](m_, "doubleAction")
	return rv
}

func (m_ Matrix) SetDoubleAction(value objc.Selector) {
	ffi.CallMethod[ffi.Void](m_, "setDoubleAction:", value)
}

func (m_ Matrix) MouseDownFlags() int {
	rv := ffi.CallMethod[int](m_, "mouseDownFlags")
	return rv
}

type MatrixDelegate interface {
	ControlTextEditingDelegate
}

type MatrixDelegateImpl struct {
	ControlTextEditingDelegateImpl
}

type MatrixDelegateWrapper struct {
	ControlTextEditingDelegateWrapper
}
