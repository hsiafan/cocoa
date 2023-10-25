// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
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
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
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
	rv := objc.CallMethod[PathCell](p_, objc.SEL("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func (p_ PathCell) InitTextCell(string_ string) PathCell {
	rv := objc.CallMethod[PathCell](p_, objc.SEL("initTextCell:"), string_)
	return rv
}

func (p_ PathCell) Init() PathCell {
	rv := objc.CallMethod[PathCell](p_, objc.SEL("init"))
	return rv
}

func (pc _PathCellClass) Alloc() PathCell {
	rv := objc.CallMethod[PathCell](pc, objc.SEL("alloc"))
	return rv
}

func (pc _PathCellClass) New() PathCell {
	rv := objc.CallMethod[PathCell](pc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewPathCell() PathCell {
	return PathCellClass.New()
}

func (p_ PathCell) MouseEntered_WithFrame_InView(event IEvent, frame foundation.Rect, view IView) {
	objc.CallMethod[objc.Void](p_, objc.SEL("mouseEntered:withFrame:inView:"), objc.ExtractPtr(event), frame, objc.ExtractPtr(view))
}

func (p_ PathCell) MouseExited_WithFrame_InView(event IEvent, frame foundation.Rect, view IView) {
	objc.CallMethod[objc.Void](p_, objc.SEL("mouseExited:withFrame:inView:"), objc.ExtractPtr(event), frame, objc.ExtractPtr(view))
}

func (p_ PathCell) RectOfPathComponentCell_WithFrame_InView(cell IPathComponentCell, frame foundation.Rect, view IView) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](p_, objc.SEL("rectOfPathComponentCell:withFrame:inView:"), objc.ExtractPtr(cell), frame, objc.ExtractPtr(view))
	return rv
}

func (p_ PathCell) PathComponentCellAtPoint_WithFrame_InView(point foundation.Point, frame foundation.Rect, view IView) PathComponentCell {
	rv := objc.CallMethod[PathComponentCell](p_, objc.SEL("pathComponentCellAtPoint:withFrame:inView:"), point, frame, objc.ExtractPtr(view))
	return rv
}

func (p_ PathCell) AllowedTypes() []string {
	rv := objc.CallMethod[[]string](p_, objc.SEL("allowedTypes"))
	return rv
}

func (p_ PathCell) SetAllowedTypes(value []string) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setAllowedTypes:"), value)
}

func (p_ PathCell) PathStyle() PathStyle {
	rv := objc.CallMethod[PathStyle](p_, objc.SEL("pathStyle"))
	return rv
}

func (p_ PathCell) SetPathStyle(value PathStyle) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setPathStyle:"), value)
}

func (p_ PathCell) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](p_, objc.SEL("placeholderAttributedString"))
	return rv
}

func (p_ PathCell) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setPlaceholderAttributedString:"), objc.ExtractPtr(value))
}

func (p_ PathCell) PlaceholderString() string {
	rv := objc.CallMethod[string](p_, objc.SEL("placeholderString"))
	return rv
}

func (p_ PathCell) SetPlaceholderString(value string) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setPlaceholderString:"), value)
}

func (p_ PathCell) BackgroundColor() Color {
	rv := objc.CallMethod[Color](p_, objc.SEL("backgroundColor"))
	return rv
}

func (p_ PathCell) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (pc _PathCellClass) PathComponentCellClass() objc.Class {
	rv := objc.CallMethod[objc.Class](pc, objc.SEL("pathComponentCellClass"))
	return rv
}

func (p_ PathCell) ClickedPathComponentCell() PathComponentCell {
	rv := objc.CallMethod[PathComponentCell](p_, objc.SEL("clickedPathComponentCell"))
	return rv
}

func (p_ PathCell) PathComponentCells() []PathComponentCell {
	rv := objc.CallMethod[[]PathComponentCell](p_, objc.SEL("pathComponentCells"))
	return rv
}

func (p_ PathCell) SetPathComponentCells(value []IPathComponentCell) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setPathComponentCells:"), value)
}

func (p_ PathCell) DoubleAction() objc.Selector {
	rv := objc.CallMethod[objc.Selector](p_, objc.SEL("doubleAction"))
	return rv
}

func (p_ PathCell) SetDoubleAction(value objc.Selector) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setDoubleAction:"), value)
}

func (p_ PathCell) URL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](p_, objc.SEL("URL"))
	return rv
}

func (p_ PathCell) SetURL(value foundation.IURL) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setURL:"), objc.ExtractPtr(value))
}

// weak property
func (p_ PathCell) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](p_, objc.SEL("delegate"))
	return rv
}

// weak property
func (p_ PathCell) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}
