// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[GraphicsContext](gc, "alloc")
	return rv
}

func (g_ GraphicsContext) Init() GraphicsContext {
	rv := ffi.CallMethod[GraphicsContext](g_, "init")
	return rv
}

func (gc _GraphicsContextClass) New() GraphicsContext {
	rv := ffi.CallMethod[GraphicsContext](gc, "new")
	rv.Autorelease()
	return rv
}

func NewGraphicsContext() GraphicsContext {
	return GraphicsContextClass.New()
}

func (gc _GraphicsContextClass) GraphicsContextWithAttributes(attributes map[GraphicsContextAttributeKey]objc.IObject) GraphicsContext {
	rv := ffi.CallMethod[GraphicsContext](gc, "graphicsContextWithAttributes:", attributes)
	return rv
}

func (gc _GraphicsContextClass) GraphicsContextWithBitmapImageRep(bitmapRep IBitmapImageRep) GraphicsContext {
	rv := ffi.CallMethod[GraphicsContext](gc, "graphicsContextWithBitmapImageRep:", bitmapRep)
	return rv
}

func (gc _GraphicsContextClass) GraphicsContextWithCGContext_Flipped(graphicsPort coregraphics.ContextRef, initialFlippedState bool) GraphicsContext {
	rv := ffi.CallMethod[GraphicsContext](gc, "graphicsContextWithCGContext:flipped:", graphicsPort, initialFlippedState)
	return rv
}

func (gc _GraphicsContextClass) GraphicsContextWithWindow(window IWindow) GraphicsContext {
	rv := ffi.CallMethod[GraphicsContext](gc, "graphicsContextWithWindow:", window)
	return rv
}

// deprecated
func (gc _GraphicsContextClass) GraphicsContextWithGraphicsPort_Flipped(graphicsPort unsafe.Pointer, initialFlippedState bool) GraphicsContext {
	rv := ffi.CallMethod[GraphicsContext](gc, "graphicsContextWithGraphicsPort:flipped:", graphicsPort, initialFlippedState)
	return rv
}

func (gc _GraphicsContextClass) RestoreGraphicsState() {
	ffi.CallMethod[ffi.Void](gc, "restoreGraphicsState")
}

func (gc _GraphicsContextClass) SaveGraphicsState() {
	ffi.CallMethod[ffi.Void](gc, "saveGraphicsState")
}

// deprecated
func (gc _GraphicsContextClass) SetGraphicsState(gState int) {
	ffi.CallMethod[ffi.Void](gc, "setGraphicsState:", gState)
}

func (gc _GraphicsContextClass) CurrentContextDrawingToScreen() bool {
	rv := ffi.CallMethod[bool](gc, "currentContextDrawingToScreen")
	return rv
}

func (g_ GraphicsContext) FlushGraphics() {
	ffi.CallMethod[ffi.Void](g_, "flushGraphics")
}

// deprecated
func (g_ GraphicsContext) FocusStack() objc.Object {
	rv := ffi.CallMethod[objc.Object](g_, "focusStack")
	return rv
}

// deprecated
func (g_ GraphicsContext) SetFocusStack(stack objc.IObject) {
	ffi.CallMethod[ffi.Void](g_, "setFocusStack:", stack)
}

func (gc _GraphicsContextClass) CurrentContext() GraphicsContext {
	rv := ffi.CallMethod[GraphicsContext](gc, "currentContext")
	return rv
}

func (gc _GraphicsContextClass) SetCurrentContext(value IGraphicsContext) {
	ffi.CallMethod[ffi.Void](gc, "setCurrentContext:", value)
}

func (g_ GraphicsContext) CGContext() coregraphics.ContextRef {
	rv := ffi.CallMethod[coregraphics.ContextRef](g_, "CGContext")
	return rv
}

// deprecated
func (g_ GraphicsContext) GraphicsPort() unsafe.Pointer {
	rv := ffi.CallMethod[unsafe.Pointer](g_, "graphicsPort")
	return rv
}

func (g_ GraphicsContext) IsDrawingToScreen() bool {
	rv := ffi.CallMethod[bool](g_, "isDrawingToScreen")
	return rv
}

func (g_ GraphicsContext) Attributes() map[GraphicsContextAttributeKey]objc.Object {
	rv := ffi.CallMethod[map[GraphicsContextAttributeKey]objc.Object](g_, "attributes")
	return rv
}

func (g_ GraphicsContext) IsFlipped() bool {
	rv := ffi.CallMethod[bool](g_, "isFlipped")
	return rv
}

func (g_ GraphicsContext) CompositingOperation() CompositingOperation {
	rv := ffi.CallMethod[CompositingOperation](g_, "compositingOperation")
	return rv
}

func (g_ GraphicsContext) SetCompositingOperation(value CompositingOperation) {
	ffi.CallMethod[ffi.Void](g_, "setCompositingOperation:", value)
}

func (g_ GraphicsContext) ImageInterpolation() ImageInterpolation {
	rv := ffi.CallMethod[ImageInterpolation](g_, "imageInterpolation")
	return rv
}

func (g_ GraphicsContext) SetImageInterpolation(value ImageInterpolation) {
	ffi.CallMethod[ffi.Void](g_, "setImageInterpolation:", value)
}

func (g_ GraphicsContext) ShouldAntialias() bool {
	rv := ffi.CallMethod[bool](g_, "shouldAntialias")
	return rv
}

func (g_ GraphicsContext) SetShouldAntialias(value bool) {
	ffi.CallMethod[ffi.Void](g_, "setShouldAntialias:", value)
}

func (g_ GraphicsContext) PatternPhase() foundation.Point {
	rv := ffi.CallMethod[foundation.Point](g_, "patternPhase")
	return rv
}

func (g_ GraphicsContext) SetPatternPhase(value foundation.Point) {
	ffi.CallMethod[ffi.Void](g_, "setPatternPhase:", value)
}

func (g_ GraphicsContext) ColorRenderingIntent() ColorRenderingIntent {
	rv := ffi.CallMethod[ColorRenderingIntent](g_, "colorRenderingIntent")
	return rv
}

func (g_ GraphicsContext) SetColorRenderingIntent(value ColorRenderingIntent) {
	ffi.CallMethod[ffi.Void](g_, "setColorRenderingIntent:", value)
}
