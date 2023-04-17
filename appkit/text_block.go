// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TextBlockClass = _TextBlockClass{objc.GetClass("NSTextBlock")}

type _TextBlockClass struct {
	objc.Class
}

type ITextBlock interface {
	objc.IObject
	SetValue_Type_ForDimension(val float64, type_ TextBlockValueType, dimension TextBlockDimension)
	ValueForDimension(dimension TextBlockDimension) float64
	ValueTypeForDimension(dimension TextBlockDimension) TextBlockValueType
	SetContentWidth_Type(val float64, type_ TextBlockValueType)
	SetWidth_Type_ForLayer_Edge(val float64, type_ TextBlockValueType, layer TextBlockLayer, edge foundation.RectEdge)
	SetWidth_Type_ForLayer(val float64, type_ TextBlockValueType, layer TextBlockLayer)
	WidthForLayer_Edge(layer TextBlockLayer, edge foundation.RectEdge) float64
	WidthValueTypeForLayer_Edge(layer TextBlockLayer, edge foundation.RectEdge) TextBlockValueType
	SetBorderColor_ForEdge(color IColor, edge foundation.RectEdge)
	SetBorderColor(color IColor)
	BorderColorForEdge(edge foundation.RectEdge) Color
	RectForLayoutAtPoint_InRect_TextContainer_CharacterRange(startingPoint foundation.Point, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect
	BoundsRectForContentRect_InRect_TextContainer_CharacterRange(contentRect foundation.Rect, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect
	DrawBackgroundWithFrame_InView_CharacterRange_LayoutManager(frameRect foundation.Rect, controlView IView, charRange foundation.Range, layoutManager ILayoutManager)
	ContentWidth() float64
	ContentWidthValueType() TextBlockValueType
	VerticalAlignment() TextBlockVerticalAlignment
	SetVerticalAlignment(value TextBlockVerticalAlignment)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
}

type TextBlock struct {
	objc.Object
}

func MakeTextBlock(ptr unsafe.Pointer) TextBlock {
	return TextBlock{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextBlock) Init() TextBlock {
	rv := ffi.CallMethod[TextBlock](t_, "init")
	return rv
}

func (tc _TextBlockClass) Alloc() TextBlock {
	rv := ffi.CallMethod[TextBlock](tc, "alloc")
	return rv
}

func (tc _TextBlockClass) New() TextBlock {
	rv := ffi.CallMethod[TextBlock](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextBlock() TextBlock {
	return TextBlockClass.New()
}

func (t_ TextBlock) SetValue_Type_ForDimension(val float64, type_ TextBlockValueType, dimension TextBlockDimension) {
	ffi.CallMethod[ffi.Void](t_, "setValue:type:forDimension:", val, type_, dimension)
}

func (t_ TextBlock) ValueForDimension(dimension TextBlockDimension) float64 {
	rv := ffi.CallMethod[float64](t_, "valueForDimension:", dimension)
	return rv
}

func (t_ TextBlock) ValueTypeForDimension(dimension TextBlockDimension) TextBlockValueType {
	rv := ffi.CallMethod[TextBlockValueType](t_, "valueTypeForDimension:", dimension)
	return rv
}

func (t_ TextBlock) SetContentWidth_Type(val float64, type_ TextBlockValueType) {
	ffi.CallMethod[ffi.Void](t_, "setContentWidth:type:", val, type_)
}

func (t_ TextBlock) SetWidth_Type_ForLayer_Edge(val float64, type_ TextBlockValueType, layer TextBlockLayer, edge foundation.RectEdge) {
	ffi.CallMethod[ffi.Void](t_, "setWidth:type:forLayer:edge:", val, type_, layer, edge)
}

func (t_ TextBlock) SetWidth_Type_ForLayer(val float64, type_ TextBlockValueType, layer TextBlockLayer) {
	ffi.CallMethod[ffi.Void](t_, "setWidth:type:forLayer:", val, type_, layer)
}

func (t_ TextBlock) WidthForLayer_Edge(layer TextBlockLayer, edge foundation.RectEdge) float64 {
	rv := ffi.CallMethod[float64](t_, "widthForLayer:edge:", layer, edge)
	return rv
}

func (t_ TextBlock) WidthValueTypeForLayer_Edge(layer TextBlockLayer, edge foundation.RectEdge) TextBlockValueType {
	rv := ffi.CallMethod[TextBlockValueType](t_, "widthValueTypeForLayer:edge:", layer, edge)
	return rv
}

func (t_ TextBlock) SetBorderColor_ForEdge(color IColor, edge foundation.RectEdge) {
	ffi.CallMethod[ffi.Void](t_, "setBorderColor:forEdge:", color, edge)
}

func (t_ TextBlock) SetBorderColor(color IColor) {
	ffi.CallMethod[ffi.Void](t_, "setBorderColor:", color)
}

func (t_ TextBlock) BorderColorForEdge(edge foundation.RectEdge) Color {
	rv := ffi.CallMethod[Color](t_, "borderColorForEdge:", edge)
	return rv
}

func (t_ TextBlock) RectForLayoutAtPoint_InRect_TextContainer_CharacterRange(startingPoint foundation.Point, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](t_, "rectForLayoutAtPoint:inRect:textContainer:characterRange:", startingPoint, rect, textContainer, charRange)
	return rv
}

func (t_ TextBlock) BoundsRectForContentRect_InRect_TextContainer_CharacterRange(contentRect foundation.Rect, rect foundation.Rect, textContainer ITextContainer, charRange foundation.Range) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](t_, "boundsRectForContentRect:inRect:textContainer:characterRange:", contentRect, rect, textContainer, charRange)
	return rv
}

func (t_ TextBlock) DrawBackgroundWithFrame_InView_CharacterRange_LayoutManager(frameRect foundation.Rect, controlView IView, charRange foundation.Range, layoutManager ILayoutManager) {
	ffi.CallMethod[ffi.Void](t_, "drawBackgroundWithFrame:inView:characterRange:layoutManager:", frameRect, controlView, charRange, layoutManager)
}

func (t_ TextBlock) ContentWidth() float64 {
	rv := ffi.CallMethod[float64](t_, "contentWidth")
	return rv
}

func (t_ TextBlock) ContentWidthValueType() TextBlockValueType {
	rv := ffi.CallMethod[TextBlockValueType](t_, "contentWidthValueType")
	return rv
}

func (t_ TextBlock) VerticalAlignment() TextBlockVerticalAlignment {
	rv := ffi.CallMethod[TextBlockVerticalAlignment](t_, "verticalAlignment")
	return rv
}

func (t_ TextBlock) SetVerticalAlignment(value TextBlockVerticalAlignment) {
	ffi.CallMethod[ffi.Void](t_, "setVerticalAlignment:", value)
}

func (t_ TextBlock) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](t_, "backgroundColor")
	return rv
}

func (t_ TextBlock) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](t_, "setBackgroundColor:", value)
}
