// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[TableRowView](t_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (t_ TableRowView) Init() TableRowView {
	rv := objc.CallMethod[TableRowView](t_, objc.SEL("init"))
	return rv
}

func (tc _TableRowViewClass) Alloc() TableRowView {
	rv := objc.CallMethod[TableRowView](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TableRowViewClass) New() TableRowView {
	rv := objc.CallMethod[TableRowView](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTableRowView() TableRowView {
	return TableRowViewClass.New()
}

func (t_ TableRowView) DrawBackgroundInRect(dirtyRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.SEL("drawBackgroundInRect:"), dirtyRect)
}

func (t_ TableRowView) DrawDraggingDestinationFeedbackInRect(dirtyRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.SEL("drawDraggingDestinationFeedbackInRect:"), dirtyRect)
}

func (t_ TableRowView) DrawSelectionInRect(dirtyRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.SEL("drawSelectionInRect:"), dirtyRect)
}

func (t_ TableRowView) DrawSeparatorInRect(dirtyRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.SEL("drawSeparatorInRect:"), dirtyRect)
}

func (t_ TableRowView) ViewAtColumn(column int) objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.SEL("viewAtColumn:"), column)
	return rv
}

func (t_ TableRowView) IsEmphasized() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isEmphasized"))
	return rv
}

func (t_ TableRowView) SetEmphasized(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setEmphasized:"), value)
}

func (t_ TableRowView) InteriorBackgroundStyle() BackgroundStyle {
	rv := objc.CallMethod[BackgroundStyle](t_, objc.SEL("interiorBackgroundStyle"))
	return rv
}

func (t_ TableRowView) IsFloating() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isFloating"))
	return rv
}

func (t_ TableRowView) SetFloating(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setFloating:"), value)
}

func (t_ TableRowView) IsSelected() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isSelected"))
	return rv
}

func (t_ TableRowView) SetSelected(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSelected:"), value)
}

func (t_ TableRowView) SelectionHighlightStyle() TableViewSelectionHighlightStyle {
	rv := objc.CallMethod[TableViewSelectionHighlightStyle](t_, objc.SEL("selectionHighlightStyle"))
	return rv
}

func (t_ TableRowView) SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setSelectionHighlightStyle:"), value)
}

func (t_ TableRowView) DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle {
	rv := objc.CallMethod[TableViewDraggingDestinationFeedbackStyle](t_, objc.SEL("draggingDestinationFeedbackStyle"))
	return rv
}

func (t_ TableRowView) SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setDraggingDestinationFeedbackStyle:"), value)
}

func (t_ TableRowView) IndentationForDropOperation() float64 {
	rv := objc.CallMethod[float64](t_, objc.SEL("indentationForDropOperation"))
	return rv
}

func (t_ TableRowView) SetIndentationForDropOperation(value float64) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setIndentationForDropOperation:"), value)
}

func (t_ TableRowView) IsTargetForDropOperation() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isTargetForDropOperation"))
	return rv
}

func (t_ TableRowView) SetTargetForDropOperation(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTargetForDropOperation:"), value)
}

func (t_ TableRowView) IsGroupRowStyle() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isGroupRowStyle"))
	return rv
}

func (t_ TableRowView) SetGroupRowStyle(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setGroupRowStyle:"), value)
}

func (t_ TableRowView) NumberOfColumns() int {
	rv := objc.CallMethod[int](t_, objc.SEL("numberOfColumns"))
	return rv
}

func (t_ TableRowView) BackgroundColor() Color {
	rv := objc.CallMethod[Color](t_, objc.SEL("backgroundColor"))
	return rv
}

func (t_ TableRowView) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (t_ TableRowView) IsNextRowSelected() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isNextRowSelected"))
	return rv
}

func (t_ TableRowView) SetNextRowSelected(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setNextRowSelected:"), value)
}

func (t_ TableRowView) IsPreviousRowSelected() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isPreviousRowSelected"))
	return rv
}

func (t_ TableRowView) SetPreviousRowSelected(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setPreviousRowSelected:"), value)
}
