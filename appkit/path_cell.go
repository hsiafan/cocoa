// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var PathCellClass = _PathCellClass{objc.GetClass("NSPathCell")}

type _PathCellClass struct {
	objc.Class
}

type IPathCell interface {
	IActionCell
	MouseEntered_WithFrame_InView(event IEvent, frame foundation.Rect, view IView)
	MouseExited_WithFrame_InView(event IEvent, frame foundation.Rect, view IView)
	RectOfPathComponentCell_WithFrame_InView(cell IPathComponentCell, frame foundation.Rect, view IView) foundation.Rect
	PathComponentCellAtPoint_WithFrame_InView(point foundation.Point, frame foundation.Rect, view IView) PathComponentCell
	AllowedTypes() []string
	SetAllowedTypes(value []string)
	PathStyle() PathStyle
	SetPathStyle(value PathStyle)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.IAttributedString)
	PlaceholderString() string
	SetPlaceholderString(value string)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	ClickedPathComponentCell() PathComponentCell
	PathComponentCells() []PathComponentCell
	SetPathComponentCells(value []IPathComponentCell)
	DoubleAction() objc.Selector
	SetDoubleAction(value objc.Selector)
	URL() foundation.URL
	SetURL(value foundation.IURL)
	Delegate() PathCellDelegateWrapper
	SetDelegate(value PathCellDelegate)
}

type PathCell struct {
	ActionCell
}

func MakePathCell(ptr unsafe.Pointer) PathCell {
	return PathCell{
		ActionCell: MakeActionCell(ptr),
	}
}

func (p_ PathCell) InitImageCell(image IImage) PathCell {
	rv := ffi.CallMethod[PathCell](p_, "initImageCell:", image)
	return rv
}

func (p_ PathCell) InitTextCell(string_ string) PathCell {
	rv := ffi.CallMethod[PathCell](p_, "initTextCell:", string_)
	return rv
}

func (p_ PathCell) Init() PathCell {
	rv := ffi.CallMethod[PathCell](p_, "init")
	return rv
}

func (pc _PathCellClass) Alloc() PathCell {
	rv := ffi.CallMethod[PathCell](pc, "alloc")
	return rv
}

func (pc _PathCellClass) New() PathCell {
	rv := ffi.CallMethod[PathCell](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPathCell() PathCell {
	return PathCellClass.New()
}

func (p_ PathCell) MouseEntered_WithFrame_InView(event IEvent, frame foundation.Rect, view IView) {
	ffi.CallMethod[ffi.Void](p_, "mouseEntered:withFrame:inView:", event, frame, view)
}

func (p_ PathCell) MouseExited_WithFrame_InView(event IEvent, frame foundation.Rect, view IView) {
	ffi.CallMethod[ffi.Void](p_, "mouseExited:withFrame:inView:", event, frame, view)
}

func (p_ PathCell) RectOfPathComponentCell_WithFrame_InView(cell IPathComponentCell, frame foundation.Rect, view IView) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](p_, "rectOfPathComponentCell:withFrame:inView:", cell, frame, view)
	return rv
}

func (p_ PathCell) PathComponentCellAtPoint_WithFrame_InView(point foundation.Point, frame foundation.Rect, view IView) PathComponentCell {
	rv := ffi.CallMethod[PathComponentCell](p_, "pathComponentCellAtPoint:withFrame:inView:", point, frame, view)
	return rv
}

func (p_ PathCell) AllowedTypes() []string {
	rv := ffi.CallMethod[[]string](p_, "allowedTypes")
	return rv
}

func (p_ PathCell) SetAllowedTypes(value []string) {
	ffi.CallMethod[ffi.Void](p_, "setAllowedTypes:", value)
}

func (p_ PathCell) PathStyle() PathStyle {
	rv := ffi.CallMethod[PathStyle](p_, "pathStyle")
	return rv
}

func (p_ PathCell) SetPathStyle(value PathStyle) {
	ffi.CallMethod[ffi.Void](p_, "setPathStyle:", value)
}

func (p_ PathCell) PlaceholderAttributedString() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](p_, "placeholderAttributedString")
	return rv
}

func (p_ PathCell) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	ffi.CallMethod[ffi.Void](p_, "setPlaceholderAttributedString:", value)
}

func (p_ PathCell) PlaceholderString() string {
	rv := ffi.CallMethod[string](p_, "placeholderString")
	return rv
}

func (p_ PathCell) SetPlaceholderString(value string) {
	ffi.CallMethod[ffi.Void](p_, "setPlaceholderString:", value)
}

func (p_ PathCell) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](p_, "backgroundColor")
	return rv
}

func (p_ PathCell) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](p_, "setBackgroundColor:", value)
}

func (pc _PathCellClass) PathComponentCellClass() objc.Class {
	rv := ffi.CallMethod[objc.Class](pc, "pathComponentCellClass")
	return rv
}

func (p_ PathCell) ClickedPathComponentCell() PathComponentCell {
	rv := ffi.CallMethod[PathComponentCell](p_, "clickedPathComponentCell")
	return rv
}

func (p_ PathCell) PathComponentCells() []PathComponentCell {
	rv := ffi.CallMethod[[]PathComponentCell](p_, "pathComponentCells")
	return rv
}

func (p_ PathCell) SetPathComponentCells(value []IPathComponentCell) {
	ffi.CallMethod[ffi.Void](p_, "setPathComponentCells:", value)
}

func (p_ PathCell) DoubleAction() objc.Selector {
	rv := ffi.CallMethod[objc.Selector](p_, "doubleAction")
	return rv
}

func (p_ PathCell) SetDoubleAction(value objc.Selector) {
	ffi.CallMethod[ffi.Void](p_, "setDoubleAction:", value)
}

func (p_ PathCell) URL() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](p_, "URL")
	return rv
}

func (p_ PathCell) SetURL(value foundation.IURL) {
	ffi.CallMethod[ffi.Void](p_, "setURL:", value)
}

func (p_ PathCell) Delegate() PathCellDelegateWrapper {
	rv := ffi.CallMethod[PathCellDelegateWrapper](p_, "delegate")
	return rv
}

func (p_ PathCell) SetDelegate(value PathCellDelegate) {
	po := ffi.CreateProtocol("NSPathCellDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(p_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](p_, "setDelegate:", po)
}
