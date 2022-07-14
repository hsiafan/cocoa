// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[BezierPath](bc, "alloc")
	return rv
}

func (b_ BezierPath) Init() BezierPath {
	rv := ffi.CallMethod[BezierPath](b_, "init")
	rv.Autorelease()
	return rv
}

func (bc _BezierPathClass) New() BezierPath {
	rv := ffi.CallMethod[BezierPath](bc, "new")
	rv.Autorelease()
	return rv
}

func NewBezierPath() BezierPath {
	return BezierPathClass.New()
}

func (bc _BezierPathClass) BezierPath() BezierPath {
	rv := ffi.CallMethod[BezierPath](bc, "bezierPath")
	return rv
}

func (bc _BezierPathClass) BezierPathWithOvalInRect(rect foundation.Rect) BezierPath {
	rv := ffi.CallMethod[BezierPath](bc, "bezierPathWithOvalInRect:", rect)
	return rv
}

func (bc _BezierPathClass) BezierPathWithRect(rect foundation.Rect) BezierPath {
	rv := ffi.CallMethod[BezierPath](bc, "bezierPathWithRect:", rect)
	return rv
}

func (bc _BezierPathClass) BezierPathWithRoundedRect_XRadius_YRadius(rect foundation.Rect, xRadius float64, yRadius float64) BezierPath {
	rv := ffi.CallMethod[BezierPath](bc, "bezierPathWithRoundedRect:xRadius:yRadius:", rect, xRadius, yRadius)
	return rv
}

func (b_ BezierPath) MoveToPoint(point foundation.Point) {
	ffi.CallMethod[ffi.Void](b_, "moveToPoint:", point)
}

func (b_ BezierPath) LineToPoint(point foundation.Point) {
	ffi.CallMethod[ffi.Void](b_, "lineToPoint:", point)
}

func (b_ BezierPath) CurveToPoint_ControlPoint1_ControlPoint2(endPoint foundation.Point, controlPoint1 foundation.Point, controlPoint2 foundation.Point) {
	ffi.CallMethod[ffi.Void](b_, "curveToPoint:controlPoint1:controlPoint2:", endPoint, controlPoint1, controlPoint2)
}

func (b_ BezierPath) ClosePath() {
	ffi.CallMethod[ffi.Void](b_, "closePath")
}

func (b_ BezierPath) RelativeMoveToPoint(point foundation.Point) {
	ffi.CallMethod[ffi.Void](b_, "relativeMoveToPoint:", point)
}

func (b_ BezierPath) RelativeLineToPoint(point foundation.Point) {
	ffi.CallMethod[ffi.Void](b_, "relativeLineToPoint:", point)
}

func (b_ BezierPath) RelativeCurveToPoint_ControlPoint1_ControlPoint2(endPoint foundation.Point, controlPoint1 foundation.Point, controlPoint2 foundation.Point) {
	ffi.CallMethod[ffi.Void](b_, "relativeCurveToPoint:controlPoint1:controlPoint2:", endPoint, controlPoint1, controlPoint2)
}

func (b_ BezierPath) AppendBezierPath(path IBezierPath) {
	ffi.CallMethod[ffi.Void](b_, "appendBezierPath:", path)
}

func (b_ BezierPath) AppendBezierPathWithPoints_Count(points *foundation.Point, count int) {
	ffi.CallMethod[ffi.Void](b_, "appendBezierPathWithPoints:count:", points, count)
}

func (b_ BezierPath) AppendBezierPathWithOvalInRect(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](b_, "appendBezierPathWithOvalInRect:", rect)
}

func (b_ BezierPath) AppendBezierPathWithArcFromPoint_ToPoint_Radius(point1 foundation.Point, point2 foundation.Point, radius float64) {
	ffi.CallMethod[ffi.Void](b_, "appendBezierPathWithArcFromPoint:toPoint:radius:", point1, point2, radius)
}

func (b_ BezierPath) AppendBezierPathWithArcWithCenter_Radius_StartAngle_EndAngle(center foundation.Point, radius float64, startAngle float64, endAngle float64) {
	ffi.CallMethod[ffi.Void](b_, "appendBezierPathWithArcWithCenter:radius:startAngle:endAngle:", center, radius, startAngle, endAngle)
}

func (b_ BezierPath) AppendBezierPathWithArcWithCenter_Radius_StartAngle_EndAngle_Clockwise(center foundation.Point, radius float64, startAngle float64, endAngle float64, clockwise bool) {
	ffi.CallMethod[ffi.Void](b_, "appendBezierPathWithArcWithCenter:radius:startAngle:endAngle:clockwise:", center, radius, startAngle, endAngle, clockwise)
}

func (b_ BezierPath) AppendBezierPathWithRect(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](b_, "appendBezierPathWithRect:", rect)
}

func (b_ BezierPath) AppendBezierPathWithRoundedRect_XRadius_YRadius(rect foundation.Rect, xRadius float64, yRadius float64) {
	ffi.CallMethod[ffi.Void](b_, "appendBezierPathWithRoundedRect:xRadius:yRadius:", rect, xRadius, yRadius)
}

func (b_ BezierPath) AppendBezierPathWithCGGlyph_InFont(glyph coregraphics.Glyph, font IFont) {
	ffi.CallMethod[ffi.Void](b_, "appendBezierPathWithCGGlyph:inFont:", glyph, font)
}

func (b_ BezierPath) AppendBezierPathWithCGGlyphs_Count_InFont(glyphs *coregraphics.Glyph, count int, font IFont) {
	ffi.CallMethod[ffi.Void](b_, "appendBezierPathWithCGGlyphs:count:inFont:", glyphs, count, font)
}

// deprecated
func (b_ BezierPath) AppendBezierPathWithGlyph_InFont(glyph Glyph, font IFont) {
	ffi.CallMethod[ffi.Void](b_, "appendBezierPathWithGlyph:inFont:", glyph, font)
}

// deprecated
func (b_ BezierPath) AppendBezierPathWithGlyphs_Count_InFont(glyphs *Glyph, count int, font IFont) {
	ffi.CallMethod[ffi.Void](b_, "appendBezierPathWithGlyphs:count:inFont:", glyphs, count, font)
}

// deprecated
func (b_ BezierPath) AppendBezierPathWithPackedGlyphs(packedGlyphs *byte) {
	ffi.CallMethod[ffi.Void](b_, "appendBezierPathWithPackedGlyphs:", packedGlyphs)
}

func (b_ BezierPath) GetLineDash_Count_Phase(pattern *float64, count *int, phase *float64) {
	ffi.CallMethod[ffi.Void](b_, "getLineDash:count:phase:", pattern, count, phase)
}

func (b_ BezierPath) SetLineDash_Count_Phase(pattern *float64, count int, phase float64) {
	ffi.CallMethod[ffi.Void](b_, "setLineDash:count:phase:", pattern, count, phase)
}

func (b_ BezierPath) Stroke() {
	ffi.CallMethod[ffi.Void](b_, "stroke")
}

func (b_ BezierPath) Fill() {
	ffi.CallMethod[ffi.Void](b_, "fill")
}

func (bc _BezierPathClass) FillRect(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](bc, "fillRect:", rect)
}

func (bc _BezierPathClass) StrokeRect(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](bc, "strokeRect:", rect)
}

func (bc _BezierPathClass) StrokeLineFromPoint_ToPoint(point1 foundation.Point, point2 foundation.Point) {
	ffi.CallMethod[ffi.Void](bc, "strokeLineFromPoint:toPoint:", point1, point2)
}

func (bc _BezierPathClass) DrawPackedGlyphs_AtPoint(packedGlyphs *byte, point foundation.Point) {
	ffi.CallMethod[ffi.Void](bc, "drawPackedGlyphs:atPoint:", packedGlyphs, point)
}

func (b_ BezierPath) AddClip() {
	ffi.CallMethod[ffi.Void](b_, "addClip")
}

func (b_ BezierPath) SetClip() {
	ffi.CallMethod[ffi.Void](b_, "setClip")
}

func (bc _BezierPathClass) ClipRect(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](bc, "clipRect:", rect)
}

func (b_ BezierPath) ContainsPoint(point foundation.Point) bool {
	rv := ffi.CallMethod[bool](b_, "containsPoint:", point)
	return rv
}

func (b_ BezierPath) TransformUsingAffineTransform(transform foundation.IAffineTransform) {
	ffi.CallMethod[ffi.Void](b_, "transformUsingAffineTransform:", transform)
}

func (b_ BezierPath) ElementAtIndex(index int) BezierPathElement {
	rv := ffi.CallMethod[BezierPathElement](b_, "elementAtIndex:", index)
	return rv
}

func (b_ BezierPath) ElementAtIndex_AssociatedPoints(index int, points *foundation.Point) BezierPathElement {
	rv := ffi.CallMethod[BezierPathElement](b_, "elementAtIndex:associatedPoints:", index, points)
	return rv
}

func (b_ BezierPath) RemoveAllPoints() {
	ffi.CallMethod[ffi.Void](b_, "removeAllPoints")
}

func (b_ BezierPath) SetAssociatedPoints_AtIndex(points *foundation.Point, index int) {
	ffi.CallMethod[ffi.Void](b_, "setAssociatedPoints:atIndex:", points, index)
}

// deprecated
func (b_ BezierPath) CachesBezierPath() bool {
	rv := ffi.CallMethod[bool](b_, "cachesBezierPath")
	return rv
}

// deprecated
func (b_ BezierPath) SetCachesBezierPath(flag bool) {
	ffi.CallMethod[ffi.Void](b_, "setCachesBezierPath:", flag)
}

func (b_ BezierPath) BezierPathByFlatteningPath() BezierPath {
	rv := ffi.CallMethod[BezierPath](b_, "bezierPathByFlatteningPath")
	return rv
}

func (b_ BezierPath) BezierPathByReversingPath() BezierPath {
	rv := ffi.CallMethod[BezierPath](b_, "bezierPathByReversingPath")
	return rv
}

func (b_ BezierPath) WindingRule() WindingRule {
	rv := ffi.CallMethod[WindingRule](b_, "windingRule")
	return rv
}

func (b_ BezierPath) SetWindingRule(value WindingRule) {
	ffi.CallMethod[ffi.Void](b_, "setWindingRule:", value)
}

func (b_ BezierPath) LineCapStyle() LineCapStyle {
	rv := ffi.CallMethod[LineCapStyle](b_, "lineCapStyle")
	return rv
}

func (b_ BezierPath) SetLineCapStyle(value LineCapStyle) {
	ffi.CallMethod[ffi.Void](b_, "setLineCapStyle:", value)
}

func (b_ BezierPath) LineJoinStyle() LineJoinStyle {
	rv := ffi.CallMethod[LineJoinStyle](b_, "lineJoinStyle")
	return rv
}

func (b_ BezierPath) SetLineJoinStyle(value LineJoinStyle) {
	ffi.CallMethod[ffi.Void](b_, "setLineJoinStyle:", value)
}

func (b_ BezierPath) LineWidth() float64 {
	rv := ffi.CallMethod[float64](b_, "lineWidth")
	return rv
}

func (b_ BezierPath) SetLineWidth(value float64) {
	ffi.CallMethod[ffi.Void](b_, "setLineWidth:", value)
}

func (b_ BezierPath) MiterLimit() float64 {
	rv := ffi.CallMethod[float64](b_, "miterLimit")
	return rv
}

func (b_ BezierPath) SetMiterLimit(value float64) {
	ffi.CallMethod[ffi.Void](b_, "setMiterLimit:", value)
}

func (b_ BezierPath) Flatness() float64 {
	rv := ffi.CallMethod[float64](b_, "flatness")
	return rv
}

func (b_ BezierPath) SetFlatness(value float64) {
	ffi.CallMethod[ffi.Void](b_, "setFlatness:", value)
}

func (bc _BezierPathClass) DefaultWindingRule() WindingRule {
	rv := ffi.CallMethod[WindingRule](bc, "defaultWindingRule")
	return rv
}

func (bc _BezierPathClass) SetDefaultWindingRule(value WindingRule) {
	ffi.CallMethod[ffi.Void](bc, "setDefaultWindingRule:", value)
}

func (bc _BezierPathClass) DefaultLineCapStyle() LineCapStyle {
	rv := ffi.CallMethod[LineCapStyle](bc, "defaultLineCapStyle")
	return rv
}

func (bc _BezierPathClass) SetDefaultLineCapStyle(value LineCapStyle) {
	ffi.CallMethod[ffi.Void](bc, "setDefaultLineCapStyle:", value)
}

func (bc _BezierPathClass) DefaultLineJoinStyle() LineJoinStyle {
	rv := ffi.CallMethod[LineJoinStyle](bc, "defaultLineJoinStyle")
	return rv
}

func (bc _BezierPathClass) SetDefaultLineJoinStyle(value LineJoinStyle) {
	ffi.CallMethod[ffi.Void](bc, "setDefaultLineJoinStyle:", value)
}

func (bc _BezierPathClass) DefaultLineWidth() float64 {
	rv := ffi.CallMethod[float64](bc, "defaultLineWidth")
	return rv
}

func (bc _BezierPathClass) SetDefaultLineWidth(value float64) {
	ffi.CallMethod[ffi.Void](bc, "setDefaultLineWidth:", value)
}

func (bc _BezierPathClass) DefaultMiterLimit() float64 {
	rv := ffi.CallMethod[float64](bc, "defaultMiterLimit")
	return rv
}

func (bc _BezierPathClass) SetDefaultMiterLimit(value float64) {
	ffi.CallMethod[ffi.Void](bc, "setDefaultMiterLimit:", value)
}

func (bc _BezierPathClass) DefaultFlatness() float64 {
	rv := ffi.CallMethod[float64](bc, "defaultFlatness")
	return rv
}

func (bc _BezierPathClass) SetDefaultFlatness(value float64) {
	ffi.CallMethod[ffi.Void](bc, "setDefaultFlatness:", value)
}

func (b_ BezierPath) Bounds() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](b_, "bounds")
	return rv
}

func (b_ BezierPath) ControlPointBounds() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](b_, "controlPointBounds")
	return rv
}

func (b_ BezierPath) CurrentPoint() foundation.Point {
	rv := ffi.CallMethod[foundation.Point](b_, "currentPoint")
	return rv
}

func (b_ BezierPath) IsEmpty() bool {
	rv := ffi.CallMethod[bool](b_, "isEmpty")
	return rv
}

func (b_ BezierPath) ElementCount() int {
	rv := ffi.CallMethod[int](b_, "elementCount")
	return rv
}
