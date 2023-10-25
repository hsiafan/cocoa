// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var BezierPathClass = _BezierPathClass{objc.GetClass("NSBezierPath")}

type _BezierPathClass struct {
	objc.Class
}

type IBezierPath interface {
	objc.IObject
	MoveToPoint(point foundation.Point)
	LineToPoint(point foundation.Point)
	CurveToPoint_ControlPoint1_ControlPoint2(endPoint foundation.Point, controlPoint1 foundation.Point, controlPoint2 foundation.Point)
	ClosePath()
	RelativeMoveToPoint(point foundation.Point)
	RelativeLineToPoint(point foundation.Point)
	RelativeCurveToPoint_ControlPoint1_ControlPoint2(endPoint foundation.Point, controlPoint1 foundation.Point, controlPoint2 foundation.Point)
	AppendBezierPath(path IBezierPath)
	AppendBezierPathWithPoints_Count(points *foundation.Point, count int)
	AppendBezierPathWithOvalInRect(rect foundation.Rect)
	AppendBezierPathWithArcFromPoint_ToPoint_Radius(point1 foundation.Point, point2 foundation.Point, radius float64)
	AppendBezierPathWithArcWithCenter_Radius_StartAngle_EndAngle(center foundation.Point, radius float64, startAngle float64, endAngle float64)
	AppendBezierPathWithArcWithCenter_Radius_StartAngle_EndAngle_Clockwise(center foundation.Point, radius float64, startAngle float64, endAngle float64, clockwise bool)
	AppendBezierPathWithRect(rect foundation.Rect)
	AppendBezierPathWithRoundedRect_XRadius_YRadius(rect foundation.Rect, xRadius float64, yRadius float64)
	AppendBezierPathWithCGGlyph_InFont(glyph coregraphics.Glyph, font IFont)
	AppendBezierPathWithCGGlyphs_Count_InFont(glyphs *coregraphics.Glyph, count int, font IFont)
	// deprecated
	AppendBezierPathWithGlyph_InFont(glyph Glyph, font IFont)
	// deprecated
	AppendBezierPathWithGlyphs_Count_InFont(glyphs *Glyph, count int, font IFont)
	// deprecated
	AppendBezierPathWithPackedGlyphs(packedGlyphs *byte)
	GetLineDash_Count_Phase(pattern *float64, count *int, phase *float64)
	SetLineDash_Count_Phase(pattern *float64, count int, phase float64)
	Stroke()
	Fill()
	AddClip()
	SetClip()
	ContainsPoint(point foundation.Point) bool
	TransformUsingAffineTransform(transform foundation.IAffineTransform)
	ElementAtIndex(index int) BezierPathElement
	ElementAtIndex_AssociatedPoints(index int, points *foundation.Point) BezierPathElement
	RemoveAllPoints()
	SetAssociatedPoints_AtIndex(points *foundation.Point, index int)
	// deprecated
	CachesBezierPath() bool
	// deprecated
	SetCachesBezierPath(flag bool)
	BezierPathByFlatteningPath() BezierPath
	BezierPathByReversingPath() BezierPath
	WindingRule() WindingRule
	SetWindingRule(value WindingRule)
	LineCapStyle() LineCapStyle
	SetLineCapStyle(value LineCapStyle)
	LineJoinStyle() LineJoinStyle
	SetLineJoinStyle(value LineJoinStyle)
	LineWidth() float64
	SetLineWidth(value float64)
	MiterLimit() float64
	SetMiterLimit(value float64)
	Flatness() float64
	SetFlatness(value float64)
	Bounds() foundation.Rect
	ControlPointBounds() foundation.Rect
	CurrentPoint() foundation.Point
	IsEmpty() bool
	ElementCount() int
}

type BezierPath struct {
	objc.Object
}

func MakeBezierPath(ptr unsafe.Pointer) BezierPath {
	return BezierPath{
		Object: objc.MakeObject(ptr),
	}
}

func (bc _BezierPathClass) Alloc() BezierPath {
	rv := objc.CallMethod[BezierPath](bc, objc.SEL("alloc"))
	return rv
}

func (bc _BezierPathClass) New() BezierPath {
	rv := objc.CallMethod[BezierPath](bc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewBezierPath() BezierPath {
	return BezierPathClass.New()
}

func (b_ BezierPath) Init() BezierPath {
	rv := objc.CallMethod[BezierPath](b_, objc.SEL("init"))
	return rv
}

func (bc _BezierPathClass) BezierPath() BezierPath {
	rv := objc.CallMethod[BezierPath](bc, objc.SEL("bezierPath"))
	return rv
}

func (bc _BezierPathClass) BezierPathWithOvalInRect(rect foundation.Rect) BezierPath {
	rv := objc.CallMethod[BezierPath](bc, objc.SEL("bezierPathWithOvalInRect:"), rect)
	return rv
}

func (bc _BezierPathClass) BezierPathWithRect(rect foundation.Rect) BezierPath {
	rv := objc.CallMethod[BezierPath](bc, objc.SEL("bezierPathWithRect:"), rect)
	return rv
}

func (bc _BezierPathClass) BezierPathWithRoundedRect_XRadius_YRadius(rect foundation.Rect, xRadius float64, yRadius float64) BezierPath {
	rv := objc.CallMethod[BezierPath](bc, objc.SEL("bezierPathWithRoundedRect:xRadius:yRadius:"), rect, xRadius, yRadius)
	return rv
}

func (b_ BezierPath) MoveToPoint(point foundation.Point) {
	objc.CallMethod[objc.Void](b_, objc.SEL("moveToPoint:"), point)
}

func (b_ BezierPath) LineToPoint(point foundation.Point) {
	objc.CallMethod[objc.Void](b_, objc.SEL("lineToPoint:"), point)
}

func (b_ BezierPath) CurveToPoint_ControlPoint1_ControlPoint2(endPoint foundation.Point, controlPoint1 foundation.Point, controlPoint2 foundation.Point) {
	objc.CallMethod[objc.Void](b_, objc.SEL("curveToPoint:controlPoint1:controlPoint2:"), endPoint, controlPoint1, controlPoint2)
}

func (b_ BezierPath) ClosePath() {
	objc.CallMethod[objc.Void](b_, objc.SEL("closePath"))
}

func (b_ BezierPath) RelativeMoveToPoint(point foundation.Point) {
	objc.CallMethod[objc.Void](b_, objc.SEL("relativeMoveToPoint:"), point)
}

func (b_ BezierPath) RelativeLineToPoint(point foundation.Point) {
	objc.CallMethod[objc.Void](b_, objc.SEL("relativeLineToPoint:"), point)
}

func (b_ BezierPath) RelativeCurveToPoint_ControlPoint1_ControlPoint2(endPoint foundation.Point, controlPoint1 foundation.Point, controlPoint2 foundation.Point) {
	objc.CallMethod[objc.Void](b_, objc.SEL("relativeCurveToPoint:controlPoint1:controlPoint2:"), endPoint, controlPoint1, controlPoint2)
}

func (b_ BezierPath) AppendBezierPath(path IBezierPath) {
	objc.CallMethod[objc.Void](b_, objc.SEL("appendBezierPath:"), objc.ExtractPtr(path))
}

func (b_ BezierPath) AppendBezierPathWithPoints_Count(points *foundation.Point, count int) {
	objc.CallMethod[objc.Void](b_, objc.SEL("appendBezierPathWithPoints:count:"), points, count)
}

func (b_ BezierPath) AppendBezierPathWithOvalInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](b_, objc.SEL("appendBezierPathWithOvalInRect:"), rect)
}

func (b_ BezierPath) AppendBezierPathWithArcFromPoint_ToPoint_Radius(point1 foundation.Point, point2 foundation.Point, radius float64) {
	objc.CallMethod[objc.Void](b_, objc.SEL("appendBezierPathWithArcFromPoint:toPoint:radius:"), point1, point2, radius)
}

func (b_ BezierPath) AppendBezierPathWithArcWithCenter_Radius_StartAngle_EndAngle(center foundation.Point, radius float64, startAngle float64, endAngle float64) {
	objc.CallMethod[objc.Void](b_, objc.SEL("appendBezierPathWithArcWithCenter:radius:startAngle:endAngle:"), center, radius, startAngle, endAngle)
}

func (b_ BezierPath) AppendBezierPathWithArcWithCenter_Radius_StartAngle_EndAngle_Clockwise(center foundation.Point, radius float64, startAngle float64, endAngle float64, clockwise bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("appendBezierPathWithArcWithCenter:radius:startAngle:endAngle:clockwise:"), center, radius, startAngle, endAngle, clockwise)
}

func (b_ BezierPath) AppendBezierPathWithRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](b_, objc.SEL("appendBezierPathWithRect:"), rect)
}

func (b_ BezierPath) AppendBezierPathWithRoundedRect_XRadius_YRadius(rect foundation.Rect, xRadius float64, yRadius float64) {
	objc.CallMethod[objc.Void](b_, objc.SEL("appendBezierPathWithRoundedRect:xRadius:yRadius:"), rect, xRadius, yRadius)
}

func (b_ BezierPath) AppendBezierPathWithCGGlyph_InFont(glyph coregraphics.Glyph, font IFont) {
	objc.CallMethod[objc.Void](b_, objc.SEL("appendBezierPathWithCGGlyph:inFont:"), glyph, objc.ExtractPtr(font))
}

func (b_ BezierPath) AppendBezierPathWithCGGlyphs_Count_InFont(glyphs *coregraphics.Glyph, count int, font IFont) {
	objc.CallMethod[objc.Void](b_, objc.SEL("appendBezierPathWithCGGlyphs:count:inFont:"), glyphs, count, objc.ExtractPtr(font))
}

// deprecated
func (b_ BezierPath) AppendBezierPathWithGlyph_InFont(glyph Glyph, font IFont) {
	objc.CallMethod[objc.Void](b_, objc.SEL("appendBezierPathWithGlyph:inFont:"), glyph, objc.ExtractPtr(font))
}

// deprecated
func (b_ BezierPath) AppendBezierPathWithGlyphs_Count_InFont(glyphs *Glyph, count int, font IFont) {
	objc.CallMethod[objc.Void](b_, objc.SEL("appendBezierPathWithGlyphs:count:inFont:"), glyphs, count, objc.ExtractPtr(font))
}

// deprecated
func (b_ BezierPath) AppendBezierPathWithPackedGlyphs(packedGlyphs *byte) {
	objc.CallMethod[objc.Void](b_, objc.SEL("appendBezierPathWithPackedGlyphs:"), packedGlyphs)
}

func (b_ BezierPath) GetLineDash_Count_Phase(pattern *float64, count *int, phase *float64) {
	objc.CallMethod[objc.Void](b_, objc.SEL("getLineDash:count:phase:"), pattern, count, phase)
}

func (b_ BezierPath) SetLineDash_Count_Phase(pattern *float64, count int, phase float64) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setLineDash:count:phase:"), pattern, count, phase)
}

func (b_ BezierPath) Stroke() {
	objc.CallMethod[objc.Void](b_, objc.SEL("stroke"))
}

func (b_ BezierPath) Fill() {
	objc.CallMethod[objc.Void](b_, objc.SEL("fill"))
}

func (bc _BezierPathClass) FillRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](bc, objc.SEL("fillRect:"), rect)
}

func (bc _BezierPathClass) StrokeRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](bc, objc.SEL("strokeRect:"), rect)
}

func (bc _BezierPathClass) StrokeLineFromPoint_ToPoint(point1 foundation.Point, point2 foundation.Point) {
	objc.CallMethod[objc.Void](bc, objc.SEL("strokeLineFromPoint:toPoint:"), point1, point2)
}

func (bc _BezierPathClass) DrawPackedGlyphs_AtPoint(packedGlyphs *byte, point foundation.Point) {
	objc.CallMethod[objc.Void](bc, objc.SEL("drawPackedGlyphs:atPoint:"), packedGlyphs, point)
}

func (b_ BezierPath) AddClip() {
	objc.CallMethod[objc.Void](b_, objc.SEL("addClip"))
}

func (b_ BezierPath) SetClip() {
	objc.CallMethod[objc.Void](b_, objc.SEL("setClip"))
}

func (bc _BezierPathClass) ClipRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](bc, objc.SEL("clipRect:"), rect)
}

func (b_ BezierPath) ContainsPoint(point foundation.Point) bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("containsPoint:"), point)
	return rv
}

func (b_ BezierPath) TransformUsingAffineTransform(transform foundation.IAffineTransform) {
	objc.CallMethod[objc.Void](b_, objc.SEL("transformUsingAffineTransform:"), objc.ExtractPtr(transform))
}

func (b_ BezierPath) ElementAtIndex(index int) BezierPathElement {
	rv := objc.CallMethod[BezierPathElement](b_, objc.SEL("elementAtIndex:"), index)
	return rv
}

func (b_ BezierPath) ElementAtIndex_AssociatedPoints(index int, points *foundation.Point) BezierPathElement {
	rv := objc.CallMethod[BezierPathElement](b_, objc.SEL("elementAtIndex:associatedPoints:"), index, points)
	return rv
}

func (b_ BezierPath) RemoveAllPoints() {
	objc.CallMethod[objc.Void](b_, objc.SEL("removeAllPoints"))
}

func (b_ BezierPath) SetAssociatedPoints_AtIndex(points *foundation.Point, index int) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setAssociatedPoints:atIndex:"), points, index)
}

// deprecated
func (b_ BezierPath) CachesBezierPath() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("cachesBezierPath"))
	return rv
}

// deprecated
func (b_ BezierPath) SetCachesBezierPath(flag bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setCachesBezierPath:"), flag)
}

func (b_ BezierPath) BezierPathByFlatteningPath() BezierPath {
	rv := objc.CallMethod[BezierPath](b_, objc.SEL("bezierPathByFlatteningPath"))
	return rv
}

func (b_ BezierPath) BezierPathByReversingPath() BezierPath {
	rv := objc.CallMethod[BezierPath](b_, objc.SEL("bezierPathByReversingPath"))
	return rv
}

func (b_ BezierPath) WindingRule() WindingRule {
	rv := objc.CallMethod[WindingRule](b_, objc.SEL("windingRule"))
	return rv
}

func (b_ BezierPath) SetWindingRule(value WindingRule) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setWindingRule:"), value)
}

func (b_ BezierPath) LineCapStyle() LineCapStyle {
	rv := objc.CallMethod[LineCapStyle](b_, objc.SEL("lineCapStyle"))
	return rv
}

func (b_ BezierPath) SetLineCapStyle(value LineCapStyle) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setLineCapStyle:"), value)
}

func (b_ BezierPath) LineJoinStyle() LineJoinStyle {
	rv := objc.CallMethod[LineJoinStyle](b_, objc.SEL("lineJoinStyle"))
	return rv
}

func (b_ BezierPath) SetLineJoinStyle(value LineJoinStyle) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setLineJoinStyle:"), value)
}

func (b_ BezierPath) LineWidth() float64 {
	rv := objc.CallMethod[float64](b_, objc.SEL("lineWidth"))
	return rv
}

func (b_ BezierPath) SetLineWidth(value float64) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setLineWidth:"), value)
}

func (b_ BezierPath) MiterLimit() float64 {
	rv := objc.CallMethod[float64](b_, objc.SEL("miterLimit"))
	return rv
}

func (b_ BezierPath) SetMiterLimit(value float64) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setMiterLimit:"), value)
}

func (b_ BezierPath) Flatness() float64 {
	rv := objc.CallMethod[float64](b_, objc.SEL("flatness"))
	return rv
}

func (b_ BezierPath) SetFlatness(value float64) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setFlatness:"), value)
}

func (bc _BezierPathClass) DefaultWindingRule() WindingRule {
	rv := objc.CallMethod[WindingRule](bc, objc.SEL("defaultWindingRule"))
	return rv
}

func (bc _BezierPathClass) SetDefaultWindingRule(value WindingRule) {
	objc.CallMethod[objc.Void](bc, objc.SEL("setDefaultWindingRule:"), value)
}

func (bc _BezierPathClass) DefaultLineCapStyle() LineCapStyle {
	rv := objc.CallMethod[LineCapStyle](bc, objc.SEL("defaultLineCapStyle"))
	return rv
}

func (bc _BezierPathClass) SetDefaultLineCapStyle(value LineCapStyle) {
	objc.CallMethod[objc.Void](bc, objc.SEL("setDefaultLineCapStyle:"), value)
}

func (bc _BezierPathClass) DefaultLineJoinStyle() LineJoinStyle {
	rv := objc.CallMethod[LineJoinStyle](bc, objc.SEL("defaultLineJoinStyle"))
	return rv
}

func (bc _BezierPathClass) SetDefaultLineJoinStyle(value LineJoinStyle) {
	objc.CallMethod[objc.Void](bc, objc.SEL("setDefaultLineJoinStyle:"), value)
}

func (bc _BezierPathClass) DefaultLineWidth() float64 {
	rv := objc.CallMethod[float64](bc, objc.SEL("defaultLineWidth"))
	return rv
}

func (bc _BezierPathClass) SetDefaultLineWidth(value float64) {
	objc.CallMethod[objc.Void](bc, objc.SEL("setDefaultLineWidth:"), value)
}

func (bc _BezierPathClass) DefaultMiterLimit() float64 {
	rv := objc.CallMethod[float64](bc, objc.SEL("defaultMiterLimit"))
	return rv
}

func (bc _BezierPathClass) SetDefaultMiterLimit(value float64) {
	objc.CallMethod[objc.Void](bc, objc.SEL("setDefaultMiterLimit:"), value)
}

func (bc _BezierPathClass) DefaultFlatness() float64 {
	rv := objc.CallMethod[float64](bc, objc.SEL("defaultFlatness"))
	return rv
}

func (bc _BezierPathClass) SetDefaultFlatness(value float64) {
	objc.CallMethod[objc.Void](bc, objc.SEL("setDefaultFlatness:"), value)
}

func (b_ BezierPath) Bounds() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](b_, objc.SEL("bounds"))
	return rv
}

func (b_ BezierPath) ControlPointBounds() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](b_, objc.SEL("controlPointBounds"))
	return rv
}

func (b_ BezierPath) CurrentPoint() foundation.Point {
	rv := objc.CallMethod[foundation.Point](b_, objc.SEL("currentPoint"))
	return rv
}

func (b_ BezierPath) IsEmpty() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("isEmpty"))
	return rv
}

func (b_ BezierPath) ElementCount() int {
	rv := objc.CallMethod[int](b_, objc.SEL("elementCount"))
	return rv
}
