// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

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
	return rv
}

func (t_ TableRowView) Init() TableRowView {
	rv := ffi.CallMethod[TableRowView](t_, "init")
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
