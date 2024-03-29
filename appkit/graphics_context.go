// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var GraphicsContextClass = _GraphicsContextClass{objc.GetClass("NSGraphicsContext")}

type _GraphicsContextClass struct {
	objc.Class
}

type IGraphicsContext interface {
	objc.IObject
	FlushGraphics()
	// deprecated
	FocusStack() objc.Object
	// deprecated
	SetFocusStack(stack objc.IObject)
	CGContext() coregraphics.ContextRef
	// deprecated
	GraphicsPort() unsafe.Pointer
	IsDrawingToScreen() bool
	Attributes() map[GraphicsContextAttributeKey]objc.Object
	IsFlipped() bool
	CompositingOperation() CompositingOperation
	SetCompositingOperation(value CompositingOperation)
	ImageInterpolation() ImageInterpolation
	SetImageInterpolation(value ImageInterpolation)
	ShouldAntialias() bool
	SetShouldAntialias(value bool)
	PatternPhase() foundation.Point
	SetPatternPhase(value foundation.Point)
	ColorRenderingIntent() ColorRenderingIntent
	SetColorRenderingIntent(value ColorRenderingIntent)
}

type GraphicsContext struct {
	objc.Object
}

func MakeGraphicsContext(ptr unsafe.Pointer) GraphicsContext {
	return GraphicsContext{
		Object: objc.MakeObject(ptr),
	}
}

func (gc _GraphicsContextClass) Alloc() GraphicsContext {
	rv := objc.CallMethod[GraphicsContext](gc, objc.SEL("alloc"))
	return rv
}

func (gc _GraphicsContextClass) New() GraphicsContext {
	rv := objc.CallMethod[GraphicsContext](gc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewGraphicsContext() GraphicsContext {
	return GraphicsContextClass.New()
}

func (g_ GraphicsContext) Init() GraphicsContext {
	rv := objc.CallMethod[GraphicsContext](g_, objc.SEL("init"))
	return rv
}

func (gc _GraphicsContextClass) GraphicsContextWithAttributes(attributes map[GraphicsContextAttributeKey]objc.IObject) GraphicsContext {
	rv := objc.CallMethod[GraphicsContext](gc, objc.SEL("graphicsContextWithAttributes:"), attributes)
	return rv
}

func (gc _GraphicsContextClass) GraphicsContextWithBitmapImageRep(bitmapRep IBitmapImageRep) GraphicsContext {
	rv := objc.CallMethod[GraphicsContext](gc, objc.SEL("graphicsContextWithBitmapImageRep:"), objc.ExtractPtr(bitmapRep))
	return rv
}

func (gc _GraphicsContextClass) GraphicsContextWithCGContext_Flipped(graphicsPort coregraphics.ContextRef, initialFlippedState bool) GraphicsContext {
	rv := objc.CallMethod[GraphicsContext](gc, objc.SEL("graphicsContextWithCGContext:flipped:"), graphicsPort, initialFlippedState)
	return rv
}

func (gc _GraphicsContextClass) GraphicsContextWithWindow(window IWindow) GraphicsContext {
	rv := objc.CallMethod[GraphicsContext](gc, objc.SEL("graphicsContextWithWindow:"), objc.ExtractPtr(window))
	return rv
}

// deprecated
func (gc _GraphicsContextClass) GraphicsContextWithGraphicsPort_Flipped(graphicsPort unsafe.Pointer, initialFlippedState bool) GraphicsContext {
	rv := objc.CallMethod[GraphicsContext](gc, objc.SEL("graphicsContextWithGraphicsPort:flipped:"), graphicsPort, initialFlippedState)
	return rv
}

func (gc _GraphicsContextClass) RestoreGraphicsState() {
	objc.CallMethod[objc.Void](gc, objc.SEL("restoreGraphicsState"))
}

func (gc _GraphicsContextClass) SaveGraphicsState() {
	objc.CallMethod[objc.Void](gc, objc.SEL("saveGraphicsState"))
}

// deprecated
func (gc _GraphicsContextClass) SetGraphicsState(gState int) {
	objc.CallMethod[objc.Void](gc, objc.SEL("setGraphicsState:"), gState)
}

func (gc _GraphicsContextClass) CurrentContextDrawingToScreen() bool {
	rv := objc.CallMethod[bool](gc, objc.SEL("currentContextDrawingToScreen"))
	return rv
}

func (g_ GraphicsContext) FlushGraphics() {
	objc.CallMethod[objc.Void](g_, objc.SEL("flushGraphics"))
}

// deprecated
func (g_ GraphicsContext) FocusStack() objc.Object {
	rv := objc.CallMethod[objc.Object](g_, objc.SEL("focusStack"))
	return rv
}

// deprecated
func (g_ GraphicsContext) SetFocusStack(stack objc.IObject) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setFocusStack:"), objc.ExtractPtr(stack))
}

func (gc _GraphicsContextClass) CurrentContext() GraphicsContext {
	rv := objc.CallMethod[GraphicsContext](gc, objc.SEL("currentContext"))
	return rv
}

func (gc _GraphicsContextClass) SetCurrentContext(value IGraphicsContext) {
	objc.CallMethod[objc.Void](gc, objc.SEL("setCurrentContext:"), objc.ExtractPtr(value))
}

func (g_ GraphicsContext) CGContext() coregraphics.ContextRef {
	rv := objc.CallMethod[coregraphics.ContextRef](g_, objc.SEL("CGContext"))
	return rv
}

// deprecated
func (g_ GraphicsContext) GraphicsPort() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](g_, objc.SEL("graphicsPort"))
	return rv
}

func (g_ GraphicsContext) IsDrawingToScreen() bool {
	rv := objc.CallMethod[bool](g_, objc.SEL("isDrawingToScreen"))
	return rv
}

func (g_ GraphicsContext) Attributes() map[GraphicsContextAttributeKey]objc.Object {
	rv := objc.CallMethod[map[GraphicsContextAttributeKey]objc.Object](g_, objc.SEL("attributes"))
	return rv
}

func (g_ GraphicsContext) IsFlipped() bool {
	rv := objc.CallMethod[bool](g_, objc.SEL("isFlipped"))
	return rv
}

func (g_ GraphicsContext) CompositingOperation() CompositingOperation {
	rv := objc.CallMethod[CompositingOperation](g_, objc.SEL("compositingOperation"))
	return rv
}

func (g_ GraphicsContext) SetCompositingOperation(value CompositingOperation) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setCompositingOperation:"), value)
}

func (g_ GraphicsContext) ImageInterpolation() ImageInterpolation {
	rv := objc.CallMethod[ImageInterpolation](g_, objc.SEL("imageInterpolation"))
	return rv
}

func (g_ GraphicsContext) SetImageInterpolation(value ImageInterpolation) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setImageInterpolation:"), value)
}

func (g_ GraphicsContext) ShouldAntialias() bool {
	rv := objc.CallMethod[bool](g_, objc.SEL("shouldAntialias"))
	return rv
}

func (g_ GraphicsContext) SetShouldAntialias(value bool) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setShouldAntialias:"), value)
}

func (g_ GraphicsContext) PatternPhase() foundation.Point {
	rv := objc.CallMethod[foundation.Point](g_, objc.SEL("patternPhase"))
	return rv
}

func (g_ GraphicsContext) SetPatternPhase(value foundation.Point) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setPatternPhase:"), value)
}

func (g_ GraphicsContext) ColorRenderingIntent() ColorRenderingIntent {
	rv := objc.CallMethod[ColorRenderingIntent](g_, objc.SEL("colorRenderingIntent"))
	return rv
}

func (g_ GraphicsContext) SetColorRenderingIntent(value ColorRenderingIntent) {
	objc.CallMethod[objc.Void](g_, objc.SEL("setColorRenderingIntent:"), value)
}
